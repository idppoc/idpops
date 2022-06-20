package handler

import (
	v1 "apiserver/api/v1"
	"apiserver/utils"
	"bytes"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"

	"io"
	"io/ioutil"
	"os"
	"os/user"
	"sync"
	"time"
)

var GitRepo = "https://github.com/idppoc/idpops.git"

var products = make([]v1.Product, 0)

func SyncGit() {
	for {
		GetClustersFromGitRepo()
		time.Sleep(15 * time.Minute)
	}
}

var mutex = sync.Mutex{}

func GetClustersFromGitRepo() error {
	log := utils.GetLogger().WithField("func", "GetClustersFromGitRepo")

	user, err := user.Current()
	if err != nil {
		return fmt.Errorf("unable to get the user defails [%v]", err)
	}

	//Get the current time to make prefix folder
	//prefix := fmt.Sprint(time.Now().UTC().Format("2006-01-02T15:04:05.999Z"))
	//prefix = strings.ReplaceAll(prefix, ":", "")
	gitRootFolder := fmt.Sprintf("%s/%s", user.HomeDir, "idpops")

	if _, err := os.Stat(gitRootFolder); err != nil {
		if os.IsNotExist(err) {
			log.Debug("Cloning the github hub repo :", GitRepo, gitRootFolder)
			_, err = git.PlainClone(gitRootFolder, false, &git.CloneOptions{
				URL:      GitRepo,
				Progress: os.Stdout,
				//ReferenceName: "madhu-changes",
				//Auth: &http.BasicAuth{
				//Username: config.GitConfig.GitUser,
				//	Password: config.GitConfig.Token,
				//},
			})

			if err != nil {
				log.Errorf("unable to clone the git repo [%v]", err)
				return fmt.Errorf("unable to clone the git repo [%v]", err)
			}
			log.Debug("Git Clone success")
		} else {

		}
	}

	clouds := getDirs(fmt.Sprintf("%s/products", gitRootFolder), log)

	mutex.Lock()
	defer mutex.Unlock()
	//var gitopsMap = make(map[string]types.InputYamlInterface)
	products = make([]v1.Product, 0)

	//products = products[:0]
	for _, c := range clouds {
		path := fmt.Sprintf("%s/products/%s", gitRootFolder, c)
		accounts := getDirs(path, log)

		for _, account := range accounts {
			accPath := fmt.Sprintf("%s/%s", path, account)
			regions := getDirs(accPath, log)
			for _, region := range regions {
				regionPath := fmt.Sprintf("%s/%s", accPath, region)
				clusters := getDirs(regionPath, log)
				for _, cluster := range clusters {
					clusterPath := fmt.Sprintf("%s/%s", regionPath, cluster)
					prodLocs := getDirs(clusterPath, log)
					for _, prod := range prodLocs {
						if prod == "foundation" {
							continue
						}

						data, err := os.ReadFile(clusterPath + "/" + prod + "/flux-sync.yaml")
						if err != nil {
							log.Errorf(clusterPath + "/flux-sync.yaml not found")
							continue
						}
						decoder := yaml.NewDecoder(bytes.NewBufferString(string(data)))
						var d v1.GitRepoType
						for {
							if err = decoder.Decode(&d); err != nil {
								if err == io.EOF {
									break
								}
								log.Errorf("Document decode failed: %w", err)
							}
							if d.Kind == "GitRepository" {
								fmt.Printf("%+v\n", d)
								break
							}

						}
						fmt.Printf("All documents decoded")
						product := v1.Product{
							Product:     prod,
							GitLoc:      d.Spec.URL,
							ClusterName: cluster,
							Cloud:       c,
							Account:     account,
							Env:         "qa",
							Region:      region,
						}
						products = append(products, product)
					}
				}
			}
		}
	}

	err = os.RemoveAll(gitRootFolder)
	if err != nil {
		log.Errorf("Unable to remove folder %s error:[%v]", gitRootFolder, err)
	}

	fmt.Println("----", products)
	return nil
}

func getDirs(root string, log *logrus.Entry) []string {
	var dirs []string
	files, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.IsDir() {
			dirs = append(dirs, f.Name())
		}
	}
	return dirs
}
