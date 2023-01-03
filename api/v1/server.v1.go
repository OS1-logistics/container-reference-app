package api

import (
	"github.com/os1-logistics/container-reference-app/internal/pkg/service"

	"github.com/gin-gonic/gin"
)

type ServerV1 struct {
	bag      service.BagServiceInterface
	shipment service.ShipmentServiceInterface
}

func NewServerV1() ServerV1 {
	return ServerV1{
		bag:      service.NewBagService(),
		shipment: service.NewShipmentService(),
	}
}

func (s ServerV1) GetBags(c *gin.Context, params GetBagsParams) {
	s.bag.GetBags()
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
func (s ServerV1) GetShipments(c *gin.Context, params GetShipmentsParams) {
	s.shipment.GetShipments()
	c.JSON(200, gin.H{"message": "ok"})
}
func (s ServerV1) CreateShipment(c *gin.Context, params CreateShipmentParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) UpdateShipmentbyClient(c *gin.Context, shipmentId string, params UpdateShipmentbyClientParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) GetShipmentAttributes(c *gin.Context, params GetShipmentAttributesParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) UpdateShipmentAttribute(c *gin.Context, params UpdateShipmentAttributeParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) GetShipmentStateTemplateConfig(c *gin.Context) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) PutShipmentStateTemplateConfig(c *gin.Context) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) Containerize(c *gin.Context, containerId string, params ContainerizeParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) ReopenShipment(c *gin.Context, shipmentId string, params ReopenShipmentParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) GetShipment(c *gin.Context, shipmentId string, params GetShipmentParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) UpdateShipment(c *gin.Context, shipmentId string, params UpdateShipmentParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) GetShipmentState(c *gin.Context, shipmentId string, params GetShipmentStateParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
func (s ServerV1) UpdateShipmentState(c *gin.Context, shipmentId string, params UpdateShipmentStateParams) {
	c.JSON(501, gin.H{"message": "Not Implemented"})
}
