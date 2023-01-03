package service

import (
	"fmt"

	container_client "github.com/os1-logistics/container-reference-app/internal/pkg/clients/container"
)

type ShipmentServiceInterface interface {
	GetShipments()
}

type ShipmentService struct {
	containerClient container_client.APIClient
}

func NewShipmentService() ShipmentService {
	return ShipmentService{}
}

func (s ShipmentService) GetShipments() {
	fmt.Println("invoked GetShipments")
}
