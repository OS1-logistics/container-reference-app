package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	api "github.com/os1-logistics/container-reference-app/api/v1"
)

var (
	env = os.Getenv("ENVIRONMENT")
)

func main() {
	glog.Info("Starting Server")
	g := gin.Default()

	g.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	})

	s := api.NewServerV1()
	api.RegisterHandlers(g, s)

	g.Run(fmt.Sprintf(":%d", 3000))
}
