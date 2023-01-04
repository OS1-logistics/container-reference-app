package service

import (
	"github.com/golang/glog"
	domain "github.com/os1-logistics/container-reference-app/internal/pkg/domain"
	aaa "github.com/os1-logistics/container-reference-app/internal/pkg/domain/aaa"
	container "github.com/os1-logistics/container-reference-app/internal/pkg/domain/container"
)

type BagServiceInterface interface {
	GetBags()
}

type BagService struct {
	containerApiClient container.APIClient
	aaaApiClient       aaa.APIClient
}

func NewBagService() BagService {

	containerClient := domain.NewContainerClient("alpha")
	aaaClient := domain.NewAAAClient("alpha")

	domain.GetToken("alpha")

	return BagService{
		aaaApiClient:       *aaaClient,
		containerApiClient: *containerClient,
	}
}

func (s BagService) GetBags() {
	domain.GetToken("alpha")
	glog.Info("invoked GetBags")
}
