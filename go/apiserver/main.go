package main

import (
	"apiserver/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"os"
)

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

	go handler.SyncGit()

	v1 := router.Group("/idpops/api/v1")

	v1.GET("/version", handler.GetVersionHandler())
	{
		v1.GET("/getProducts", handler.GetProductsHandler())
		v1.GET("/getProductDetails", handler.GetProductDetailsHandler())
	}

	router.Run(":30002") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
