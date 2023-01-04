package api

import (
	"github.com/os1-logistics/container-reference-app/internal/pkg/service"

	"github.com/gin-gonic/gin"
)

type ServerV1 struct {
	bagcsa     service.BagServiceInterface
	packagecsa service.PackageServiceInterface
}

func NewServerV1() ServerV1 {
	return ServerV1{
		bagcsa:     service.NewBagService(),
		packagecsa: service.NewPackageService(),
	}
}

func (s ServerV1) GetBags(c *gin.Context, params GetBagsParams) {
	s.bagcsa.GetBags()
	c.JSON(200, gin.H{"message": "ok"})
}
func (s ServerV1) CreateBag(c *gin.Context, params CreateBagParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) GetAttributes(c *gin.Context, params GetAttributesParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) UpdateAttribute(c *gin.Context, params UpdateAttributeParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) BagContainerization(c *gin.Context, params BagContainerizationParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) GetStateTemplateConfig(c *gin.Context) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) PutStateTemplateConfig(c *gin.Context) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) GetBag(c *gin.Context, bagId string, params GetBagParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) UpdateBag(c *gin.Context, bagId string, params UpdateBagParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) SealBag(c *gin.Context, bagId string, params SealBagParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) GetBagState(c *gin.Context, bagId string, params GetBagStateParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) UpdateBagState(c *gin.Context, bagId string, params UpdateBagStateParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) UnsealBag(c *gin.Context, bagId string, params UnsealBagParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) GetPackages(c *gin.Context, params GetPackagesParams) {
	s.packagecsa.GetPackages()
	c.JSON(200, gin.H{"message": "ok"})
}
func (s ServerV1) CreatePackage(c *gin.Context, params CreatePackageParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) UpdatePackagebyClient(c *gin.Context, packageId string, params UpdatePackagebyClientParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) GetPackageAttributes(c *gin.Context, params GetPackageAttributesParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) UpdatePackageAttribute(c *gin.Context, params UpdatePackageAttributeParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) GetPackageStateTemplateConfig(c *gin.Context) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) PutPackageStateTemplateConfig(c *gin.Context) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) Containerize(c *gin.Context, containerId string, params ContainerizeParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) ReopenPackage(c *gin.Context, packageId string, params ReopenPackageParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) GetPackage(c *gin.Context, packageId string, params GetPackageParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) UpdatePackage(c *gin.Context, packageId string, params UpdatePackageParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) GetPackageState(c *gin.Context, packageId string, params GetPackageStateParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) UpdatePackageState(c *gin.Context, packageId string, params UpdatePackageStateParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
