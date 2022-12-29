package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	util "github.com/os1-logistics/container-reference-app/internal/pkg"
)

type Server struct {
}

func (t Server) GetContainerById(c *gin.Context, id string) {
	// our logic to retrieve all containers from a persistent OS1 Platform
	container := Container{
		Id:        &id,
		Name:      &id,
		CreatedAt: util.ToTimePtr(time.Now()),
	}
	c.JSON(http.StatusOK, gin.H{"data": container})
}
