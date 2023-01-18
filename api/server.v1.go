package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	api_v1 "github.com/os1-logistics/container-reference-app/api/v1"
	"github.com/os1-logistics/container-reference-app/internal/pkg/common"
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

func prepareContext(c *gin.Context) context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, common.ContextTenantId, c.GetString(common.ContextTenantId))
	ctx = context.WithValue(ctx, common.ContextRequestID, c.GetString(common.ContextRequestID))
	ctx = context.WithValue(ctx, common.ContextUserinfo, fmt.Sprintf("%#v", c.MustGet(common.ContextUserinfo)))
	ctx = context.WithValue(ctx, common.ContextMiddlewareAccessToken, c.GetString(common.ContextMiddlewareAccessToken))
	return ctx
}

// Init

func (s ServerV1) InitTenant(c *gin.Context, params api_v1.InitTenantParams) {
	ctx := prepareContext(c)
	Response := &api_v1.DefaultResponse{}
	e := s.ser.Init(ctx)
	if e != nil {
		Response.Error = &api_v1.ErrorSchema{}
		Response.Error.Description = e.Error()
		c.JSON(http.StatusInternalServerError, Response)
		return
	}
	c.JSON(200, Response)
}

// Packages

func (s ServerV1) GetPackage(c *gin.Context, packageId string, params api_v1.GetPackageParams) {
	ctx := prepareContext(c)
	r, e := s.ser.GetPackage(ctx, packageId)
	if e != nil {
		Response := &api_v1.CreatedResponse{}
		Response.Error = &api_v1.ErrorSchema{}
		Response.Error.Description = e.Error()
		c.JSON(http.StatusInternalServerError, Response)
		return
	}

	c.JSON(200, r)
}

func (s ServerV1) GetPackages(c *gin.Context, params api_v1.GetPackagesParams) {
	ctx := prepareContext(c)
	r, e := s.ser.GetPackages(ctx)
	if e != nil {
		Response := &api_v1.CreatedResponse{}
		Response.Error = &api_v1.ErrorSchema{}
		Response.Error.Description = e.Error()
		c.JSON(http.StatusInternalServerError, Response)
		return
	}

	c.JSON(200, r)
}

func (s ServerV1) CreatePackage(c *gin.Context, params api_v1.CreatePackageParams) {
	ctx := prepareContext(c)
	Response := &api_v1.CreatedResponse{}
	r := api_v1.CreatePackageJSONRequestBody{}
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	json.Unmarshal([]byte(buf.String()), &r)
	id, e := s.ser.CreatePackage(ctx, r)
	if e != nil {
		Response.Error = &api_v1.ErrorSchema{}
		Response.Error.Description = e.Error()
		c.JSON(http.StatusInternalServerError, Response)
		return
	}

	Response.Data = &api_v1.DataSchema{}
	Response.Data.Id = id
	c.JSON(http.StatusAccepted, Response)
}

func (s ServerV1) ChangePackageState(c *gin.Context, packageId string, command string, params api_v1.ChangePackageStateParams) {
	ctx := prepareContext(c)
	Response := &api_v1.CreatedResponse{}
	e := s.ser.UpdateContainerState(ctx, packageId, command, common.PackageContainerTypeName)
	if e != nil {
		Response.Error = &api_v1.ErrorSchema{}
		Response.Error.Description = e.Error()
		c.JSON(http.StatusInternalServerError, Response)
		return
	}

	c.JSON(http.StatusAccepted, Response)
}

// Bags

func (s ServerV1) GetBag(c *gin.Context, bagId string, params api_v1.GetBagParams) {
	ctx := prepareContext(c)
	r, e := s.ser.GetBag(ctx, bagId)
	if e != nil {
		Response := &api_v1.CreatedResponse{}
		Response.Error = &api_v1.ErrorSchema{}
		Response.Error.Description = e.Error()
		c.JSON(http.StatusInternalServerError, Response)
		return
	}

	c.JSON(200, r)

}

func (s ServerV1) GetBags(c *gin.Context, params api_v1.GetBagsParams) {
	ctx := prepareContext(c)
	r, e := s.ser.GetBags(ctx)
	if e != nil {
		Response := &api_v1.CreatedResponse{}
		Response.Error = &api_v1.ErrorSchema{}
		Response.Error.Description = e.Error()
		c.JSON(http.StatusInternalServerError, Response)
		return
	}

	c.JSON(200, r)

}

func (s ServerV1) CreateBag(c *gin.Context, params api_v1.CreateBagParams) {
	ctx := prepareContext(c)
	Response := &api_v1.CreatedResponse{}
	r := api_v1.CreateBagJSONRequestBody{}
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	json.Unmarshal([]byte(buf.String()), &r)
	id, e := s.ser.CreateBag(ctx, r)
	if e != nil {
		Response.Error = &api_v1.ErrorSchema{}
		Response.Error.Description = e.Error()
		c.JSON(http.StatusInternalServerError, Response)
		return
	}

	Response.Data = &api_v1.DataSchema{}
	Response.Data.Id = id
	c.JSON(http.StatusAccepted, Response)

}

func (s ServerV1) ChangeBagState(c *gin.Context, bagId string, command string, params api_v1.ChangeBagStateParams) {
	ctx := prepareContext(c)
	Response := &api_v1.DefaultResponse{}
	e := s.ser.UpdateContainerState(ctx, bagId, command, common.BagContainerTypeName)
	if e != nil {
		Response.Error = &api_v1.ErrorSchema{}
		Response.Error.Description = e.Error()
		c.JSON(http.StatusInternalServerError, Response)
		return
	}

	c.JSON(http.StatusOK, Response)
}

func (s ServerV1) AddPackageToBag(c *gin.Context, bagId string, packageId string, params api_v1.AddPackageToBagParams) {
	ctx := prepareContext(c)
	Response := &api_v1.DefaultResponse{}
	e := s.ser.ContainerizeOperations(ctx, bagId, packageId, common.CONTAINER_OPERATION_CONTAINERIZE)
	if e != nil {
		Response.Error = e
		c.JSON(http.StatusInternalServerError, Response)
		return
	}

	c.JSON(http.StatusOK, Response)
}

func (s ServerV1) RemovePackageFromBag(c *gin.Context, bagId string, packageId string, params api_v1.RemovePackageFromBagParams) {
	ctx := prepareContext(c)
	Response := &api_v1.DefaultResponse{}
	e := s.ser.ContainerizeOperations(ctx, bagId, packageId, common.CONTAINER_OPERATION_DECONTAINERIZE)
	if e != nil {
		Response.Error = e
		c.JSON(http.StatusInternalServerError, Response)
		return
	}

	c.JSON(http.StatusOK, Response)

}
