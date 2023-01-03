package service

import "fmt"

type ShipmentServiceInterface interface {
	GetShipments()
}

type ShipmentService struct {
}

func NewShipmentService() ShipmentService {
	return ShipmentService{}
}

func (s ShipmentService) GetShipments() {
	fmt.Println("invoked GetShipments")
}
