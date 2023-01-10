package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"

	api "github.com/os1-logistics/container-reference-app/api"
	apiv1 "github.com/os1-logistics/container-reference-app/api/v1"
	config "github.com/os1-logistics/container-reference-app/configs"
	cache "github.com/os1-logistics/container-reference-app/internal/pkg/cache"
	"github.com/os1-logistics/container-reference-app/internal/pkg/common"
)

func main() {

	glog.Info("Starting Server")

	// startup initializations
	config.LoadConfig()
	cache.LoadCache()

	/**
	Uncomment the below line to run the initialization as part of the tenant that
	you are working on so tha the container type is created in start up and easy for
	developers to test
	*/

	// TODO: add a api to initialize the container type
	// err := inits.Initialize("alpha")
	// if err != nil {
	// 	glog.Error(err)
	// 	panic("Unable to initialize the application. Container type creation failed")
	// }

	r := gin.New()

	// interceptors for all requests
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(api.CORS())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	})

	// create a new v1 server instance
	s := api.NewServerV1()

	// serveer options
	o := apiv1.GinServerOptions{
		BaseURL: common.BasePath,
	}

	// register generated handlers with v1 server implementation
	apiv1.RegisterHandlersWithOptions(r, s, o)

	r.Run(fmt.Sprintf(":%d", 3000))
}
