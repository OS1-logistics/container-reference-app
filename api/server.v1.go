package api

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	api_v1 "github.com/os1-logistics/container-reference-app/api/v1"
	service "github.com/os1-logistics/container-reference-app/internal/pkg/service"
)

type ServerV1 struct {
	ser service.Service
}

func NewServerV1() ServerV1 {
	return ServerV1{
		ser: service.NewPackageService(),
	}
}

// Packages

func (s ServerV1) GetPackage(c *gin.Context, packageId string, params api_v1.GetPackageParams) {

	r, e := s.ser.GetPackage(params.XCOREOSTENANTID, packageId)
	if e != nil {
		Response := &api_v1.CreatedResponse{}
		Response.ErrorSchema = &api_v1.ErrorSchema{}
		Response.ErrorSchema.Description = e.Error()
		c.JSON(http.StatusInternalServerError, Response)
		return
	}

	c.JSON(200, r)
}

func (s ServerV1) GetPackages(c *gin.Context, params api_v1.GetPackagesParams) {

	r, e := s.ser.GetPackages(params.XCOREOSTENANTID)
	if e != nil {
		Response := &api_v1.CreatedResponse{}
		Response.ErrorSchema = &api_v1.ErrorSchema{}
		Response.ErrorSchema.Description = e.Error()
		c.JSON(http.StatusInternalServerError, Response)
		return
	}

	c.JSON(200, r)
}

func (s ServerV1) CreatePackage(c *gin.Context, params api_v1.CreatePackageParams) {
	Response := &api_v1.CreatedResponse{}

	r := api_v1.CreatePackageJSONRequestBody{}
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	json.Unmarshal([]byte(buf.String()), &r)
	id, e := s.ser.CreatePackage(params.XCOREOSTENANTID, r)
	if e != nil {
		Response.ErrorSchema = &api_v1.ErrorSchema{}
		Response.ErrorSchema.Description = e.Error()
		c.JSON(http.StatusInternalServerError, Response)
		return
	}

	Response.DataSchema = &api_v1.DataSchema{}
	Response.DataSchema.Id = id
	c.JSON(http.StatusAccepted, Response)
}

func (s ServerV1) OpenPackage(c *gin.Context, packageId string, command string, params api_v1.OpenPackageParams) {
	Response := &api_v1.CreatedResponse{}

	e := s.ser.UpdateContainerState(params.XCOREOSTENANTID, packageId, command)
	if e != nil {
		Response.ErrorSchema = &api_v1.ErrorSchema{}
		Response.ErrorSchema.Description = e.Error()
		c.JSON(http.StatusInternalServerError, Response)
		return
	}

	c.JSON(http.StatusAccepted, Response)
}

// Bags

func (s ServerV1) GetBag(c *gin.Context, bagId string, params api_v1.GetBagParams) {

	r, e := s.ser.GetBag(params.XCOREOSTENANTID, bagId)
	if e != nil {
		Response := &api_v1.CreatedResponse{}
		Response.ErrorSchema = &api_v1.ErrorSchema{}
		Response.ErrorSchema.Description = e.Error()
		c.JSON(http.StatusInternalServerError, Response)
		return
	}

	c.JSON(200, r)

}

func (s ServerV1) GetBags(c *gin.Context, params api_v1.GetBagsParams) {

	r, e := s.ser.GetBags(params.XCOREOSTENANTID)
	if e != nil {
		Response := &api_v1.CreatedResponse{}
		Response.ErrorSchema = &api_v1.ErrorSchema{}
		Response.ErrorSchema.Description = e.Error()
		c.JSON(http.StatusInternalServerError, Response)
		return
	}

	c.JSON(200, r)

}

func (s ServerV1) CreateBag(c *gin.Context, params api_v1.CreateBagParams) {
	Response := &api_v1.CreatedResponse{}

	r := api_v1.CreateBagJSONRequestBody{}
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	json.Unmarshal([]byte(buf.String()), &r)
	id, e := s.ser.CreateBag(params.XCOREOSTENANTID, r)
	if e != nil {
		Response.ErrorSchema = &api_v1.ErrorSchema{}
		Response.ErrorSchema.Description = e.Error()
		c.JSON(http.StatusInternalServerError, Response)
		return
	}

	Response.DataSchema = &api_v1.DataSchema{}
	Response.DataSchema.Id = id
	c.JSON(http.StatusAccepted, Response)

}

func (s ServerV1) OpenBag(c *gin.Context, bagId string, command string, params api_v1.OpenBagParams) {
	Response := &api_v1.CreatedResponse{}

	e := s.ser.UpdateContainerState(params.XCOREOSTENANTID, bagId, command)
	if e != nil {
		Response.ErrorSchema = &api_v1.ErrorSchema{}
		Response.ErrorSchema.Description = e.Error()
		c.JSON(http.StatusInternalServerError, Response)
		return
	}

	c.JSON(http.StatusAccepted, Response)
}
