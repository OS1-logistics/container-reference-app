package controller

import (
	"container/pkg/models/item"

	"github.com/gin-gonic/gin"
)

type CreateContainerInput struct {
	State string 			`json:"state"`
	IsReuseable bool 	`json:"isReusable"`
	Items []item.Item `json:"items"`
}

func GetAllContainers(c *gin.Context) {
	//TODO
}

func GetContainer(c *gin.Context) {
	//TODO
	//id := c.Param("id")
}

func CreateContainer(c *gin.Context) {
	//TODO
}

func UpdateContainer(c *gin.Context) {
	//TODO
}
