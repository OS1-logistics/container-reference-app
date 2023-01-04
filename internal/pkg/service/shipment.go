package service

import (
	"github.com/golang/glog"
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
	glog.Info("invoked GetShipments")
}
