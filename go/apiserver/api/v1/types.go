package v1

type Product struct {
	Product     string `json:"product"`
	GitLoc      string `json:"gitLoc"`
	ClusterName string `json:"clusterName"`
	Cloud       string `json:"cloud"`
	Account     string `json:"account"`
	Env         string `json:"env"`
	Region      string `json:"region"`
}

type GitRepoType struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Spec       struct {
		URL string `yaml:"url"`
	} `yaml:"spec"`
}
