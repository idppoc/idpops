package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"

	//appv1 "github.com/idppoc/idpops/go/api/v1"
	"net/http"
)

const VERSION_V1 = "v1"

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

func GetVersionHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"version": VERSION_V1,
		})
	}
}

func GetProductsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		//products := make([]appv1.Product, 0)
		//err := json.Unmarshal([]byte(jsonData), &products)
		if len(products) == 0 {
			fmt.Println("No Products found")
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "No Products found",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"products": products,
		})
	}
}

func GetProductDetailsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		if len(products) == 0 {
			fmt.Println("No Products found")
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "No Products found",
			})
			return
		}
		//cloud=&account=&clusterName=&region=&product=
		cloud, cexists := c.GetQuery("cloud")
		acc, accExists := c.GetQuery("account")
		cname, cnameExists := c.GetQuery("clusterName")
		region, regionExists := c.GetQuery("region")
		product, productExists := c.GetQuery("product")
		if cexists && accExists && cnameExists && regionExists && productExists {
			for _, p := range products {
				if p.Cloud == cloud && p.Product == product && p.Account == acc &&
					p.ClusterName == cname && p.Region == region {
					c.JSON(http.StatusOK, gin.H{
						"product": p,
					})
					return
				}
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "product not found",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Query prameters \"cloud, account, clusterName,region, product\" shoule be preesnt",
			})
		}
	}
}
