package service

import (
	"context"
	"fmt"

	"github.com/golang/glog"
	api_v1 "github.com/os1-logistics/container-reference-app/api/v1"
	"github.com/os1-logistics/container-reference-app/internal/pkg/common"
	domain "github.com/os1-logistics/container-reference-app/internal/pkg/domain"
	container "github.com/os1-logistics/container-reference-app/internal/pkg/domain/container"
	containerdomain "github.com/os1-logistics/container-reference-app/internal/pkg/domain/container"
)

type PackageService struct {
	containerApiClient container.APIClient
}

func NewPackageService() PackageService {
	containerClient := domain.NewContainerClient("alpha")
	return PackageService{
		containerApiClient: *containerClient,
	}
}

func (s PackageService) GetPackages(tenantId string) {
	glog.Infof("invoked GetPackages with tenant %s", tenantId)
}

func (s PackageService) CreatePackage(tenantId string, request api_v1.CreatePackageJSONRequestBody) (*string, error) {
	glog.Info("invoked CreatePackage")
	fmt.Printf("request: %v", request)

	token, _ := domain.GetToken(tenantId)
	ctx := context.Background()
	ApiCreateContainerRequest := s.containerApiClient.ContainerApi.CreateContainer(ctx, common.PackageContainerTypeName)
	ApiCreateContainerRequest = ApiCreateContainerRequest.XCOREOSACCESS(token)
	ApiCreateContainerRequest = ApiCreateContainerRequest.XCOREOSORIGINTOKEN(token)
	ApiCreateContainerRequest = ApiCreateContainerRequest.XCOREOSTID(tenantId)
	ApiCreateContainerRequest = ApiCreateContainerRequest.XCOREOSREQUESTID("1234")
	ApiCreateContainerRequest = ApiCreateContainerRequest.XCOREOSUSERINFO("1234")

	containerCreateRequest := containerdomain.ContainerCreateRequest{
		IsHazmat:          request.IsHazmat,
		IsContainerizable: request.IsContainerizable,
		IsReusable:        request.IsReusable,
		TrackingDetails:   []containerdomain.ContainerCreateAttributesTrackingDetailsInner{},
		Attributes: map[string]interface{}{
			"containerName": "package",
			"containerType": "non-reusable",
			"leaf":          true,
			"origin":        request.Origin,
			"destination":   request.Destination,
			"trackingId":    request.TrackingId,
		},
	}
	containerCreateRequest.TrackingDetails = append(containerCreateRequest.TrackingDetails, containerdomain.ContainerCreateAttributesTrackingDetailsInner{
		Operator:   "default",
		TrackingId: request.TrackingId,
		IsPrimary:  containerdomain.PtrBool(true),
	})

	ApiCreateContainerRequest = ApiCreateContainerRequest.ContainerCreateRequest(containerCreateRequest)

	d, r, e := s.containerApiClient.ContainerApi.CreateContainerExecute(ApiCreateContainerRequest)

	if e != nil {
		glog.Error(e)
		fmt.Printf("%v", r)
		return nil, e
	}

	if r.StatusCode == 202 {
		return containerdomain.PtrString(*d.Data.Id), nil
	}

	return nil, e

}
