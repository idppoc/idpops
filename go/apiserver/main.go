package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type Product struct {
	Product     string `json:"product"`
	GitLoc      string `json:"gitLoc"`
	ClusterName string `json:"clusterName"`
	Cloud       string `json:"cloud"`
	Account     string `json:"account"`
	Env         string `json:"env"`
	Region      string `json:"region"`
}

var jsonData = `[
  {
    "product": "app1",
    "gitLoc": "https://stefanprodan.github.io/podinfo",
    "clusterName": "java-app-cluster",
    "cloud": "aws",
    "account": "account1",
    "env": "qa",
    "region": "us-west-2"
  },
  {
    "product": "app2",
    "gitLoc": "https://stefanprodan.github.io/podinfo1",
    "clusterName": "java-app-cluster1",
    "cloud": "aws",
    "account": "account2",
    "env": "qa",
    "region": "us-west-2"
  }
]
`

func main() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"Origin", "Access-Control-Allow-Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))

	dev := os.Getenv("DEVELOPEMENT")
	// CreateFoundationBlueprint AWS Session

	localFs := static.LocalFile("/webapp/build", true)
	if dev == "true" {
		//router.Use(static.Serve("/", static.LocalFile("/Users/mseelam/gopath/src/github.com/infacloud/kubeyard-dashboard/webapp/build", true)))
		localFs = static.LocalFile("/Users/mseelam/gopath/src/github.com/idppoc/first/build", true)
	}

	router.Use(static.Serve("/projects", localFs))
	router.Use(static.Serve("/home", localFs))

	router.Use(static.Serve("/", localFs))

	products := make([]Product, 0)
	err := json.Unmarshal([]byte(jsonData), &products)
	if err != nil {
		fmt.Println("Error parsing json", err)
	}

	router.GET("/getProducts", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"products": products,
		})
	})
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
