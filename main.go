package main

import (
	"container/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.GET("/albums", controller.GetAllContainers)
		router.GET("/albums/:id", controller.GetContainer)
		router.POST("/albums", controller.CreateContainer)

    router.Run("localhost:3000")
}