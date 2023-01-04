package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	api "github.com/os1-logistics/container-reference-app/api/v1"
	config "github.com/os1-logistics/container-reference-app/configs"
	cache "github.com/os1-logistics/container-reference-app/internal/pkg/cache"
)


func main() {

	glog.Info("Starting Server")

	// load configurations from environment valiables
	config.LoadConfig()
	cache.LoadCache()

	g := gin.Default()

	g.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	})

	s := api.NewServerV1()
	api.RegisterHandlers(g, s)

	g.Run(fmt.Sprintf(":%d", 3000))
}
