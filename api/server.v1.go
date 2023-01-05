package api

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	api_v1 "github.com/os1-logistics/container-reference-app/api/v1"
	"github.com/os1-logistics/container-reference-app/internal/pkg/service"
)

type ServerV1 struct {
	bagcsa     service.BagServiceInterface
	packagecsa service.PackageService
}

func NewServerV1() ServerV1 {
	return ServerV1{
		bagcsa:     service.NewBagService(),
		packagecsa: service.NewPackageService(),
	}
}

func (s ServerV1) GetBags(c *gin.Context, params api_v1.GetBagsParams) {
	s.bagcsa.GetBags()
	c.JSON(200, gin.H{"message": "ok"})
}
func (s ServerV1) CreateBag(c *gin.Context, params api_v1.CreateBagParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) GetAttributes(c *gin.Context, params api_v1.GetAttributesParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) UpdateAttribute(c *gin.Context, params api_v1.UpdateAttributeParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) BagContainerization(c *gin.Context, params api_v1.BagContainerizationParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) GetStateTemplateConfig(c *gin.Context) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) PutStateTemplateConfig(c *gin.Context) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) GetBag(c *gin.Context, bagId string, params api_v1.GetBagParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) UpdateBag(c *gin.Context, bagId string, params api_v1.UpdateBagParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) SealBag(c *gin.Context, bagId string, params api_v1.SealBagParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) GetBagState(c *gin.Context, bagId string, params api_v1.GetBagStateParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) UpdateBagState(c *gin.Context, bagId string, params api_v1.UpdateBagStateParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) UnsealBag(c *gin.Context, bagId string, params api_v1.UnsealBagParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) GetPackages(c *gin.Context, params api_v1.GetPackagesParams) {
	s.packagecsa.GetPackages(params.XCOREOSTENANTID)
	c.JSON(501, gin.H{"message": "Not Implemented"})
}

func (s ServerV1) CreatePackage(c *gin.Context, params api_v1.CreatePackageParams) {
	Response := &api_v1.CreatedResponse{}
	Response.RequestSchema = &api_v1.RequestSchema{}
	Response.RequestSchema.Uri = &c.Request.URL.Path
	Response.RequestSchema.QueryString = &c.Request.URL.RawQuery

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

func (s ServerV1) UpdatePackagebyClient(c *gin.Context, packageId string, params api_v1.UpdatePackagebyClientParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) GetPackageAttributes(c *gin.Context, params api_v1.GetPackageAttributesParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) UpdatePackageAttribute(c *gin.Context, params api_v1.UpdatePackageAttributeParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) GetPackageStateTemplateConfig(c *gin.Context) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) PutPackageStateTemplateConfig(c *gin.Context) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) Containerize(c *gin.Context, containerId string, params api_v1.ContainerizeParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) ReopenPackage(c *gin.Context, packageId string, params api_v1.ReopenPackageParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) GetPackage(c *gin.Context, packageId string, params api_v1.GetPackageParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) UpdatePackage(c *gin.Context, packageId string, params api_v1.UpdatePackageParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) GetPackageState(c *gin.Context, packageId string, params api_v1.GetPackageStateParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) UpdatePackageState(c *gin.Context, packageId string, params api_v1.UpdatePackageStateParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
