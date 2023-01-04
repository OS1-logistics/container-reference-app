package service

import (
	"fmt"

	container "github.com/os1-logistics/container-reference-app/internal/pkg/domain/container"
)

type ShipmentServiceInterface interface {
	GetShipments()
}

type ShipmentService struct {
	containerClient container.APIClient
}

func NewShipmentService() ShipmentService {
	return ShipmentService{}
}

func (s ShipmentService) GetShipments() {
	fmt.Println("invoked GetShipments")
}
