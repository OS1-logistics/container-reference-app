package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	api "github.com/os1-logistics/container-reference-app/api/v1"
)

func main() {
	glog.Info("Starting Server")
	g := gin.Default()
	s := api.NewServerV1()
	api.RegisterHandlers(g, s)
	g.Run(fmt.Sprintf(":%d", 3000))
}
