package service

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	api_v1 "github.com/os1-logistics/container-reference-app/api/v1"
)

type ServiceV1 struct {
	packagecsa PackageService
}

func NewServiceV1() ServiceV1 {
	return ServiceV1{
		packagecsa: NewPackageService(),
	}
}

func (s ServiceV1) CreatePackage(c *gin.Context, params api_v1.CreatePackageParams) {
	Response := &api_v1.CreatedResponse{}

	r := api_v1.CreatePackageJSONRequestBody{}
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	json.Unmarshal([]byte(buf.String()), &r)
	id, e := s.packagecsa.CreatePackage(params.XCOREOSTENANTID, r)
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

func (s ServiceV1) GetPackage(c *gin.Context, packageId string, params api_v1.GetPackageParams) {

	r, e := s.packagecsa.GetPackage(params.XCOREOSTENANTID, packageId)
	if e != nil {
		Response := &api_v1.CreatedResponse{}
		Response.ErrorSchema = &api_v1.ErrorSchema{}
		Response.ErrorSchema.Description = e.Error()
		c.JSON(http.StatusInternalServerError, Response)
		return
	}

	c.JSON(200, r)
}

func (s ServiceV1) GetPackages(c *gin.Context, params api_v1.GetPackagesParams) {

	r, e := s.packagecsa.GetPackages(params.XCOREOSTENANTID)
	if e != nil {
		Response := &api_v1.CreatedResponse{}
		Response.ErrorSchema = &api_v1.ErrorSchema{}
		Response.ErrorSchema.Description = e.Error()
		c.JSON(http.StatusInternalServerError, Response)
		return
	}

	c.JSON(200, r)
}
