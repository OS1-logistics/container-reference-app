package service

import (
	"github.com/golang/glog"
	container "github.com/os1-logistics/container-reference-app/internal/pkg/domain/container"
)

type PackageServiceInterface interface {
	GetPackages()
}

type PackageService struct {
	containerClient container.APIClient
}

func NewPackageService() PackageService {
	return PackageService{}
}

func (s PackageService) GetPackages() {
	glog.Info("invoked GetPackages")
}
