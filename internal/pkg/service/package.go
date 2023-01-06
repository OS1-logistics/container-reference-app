package service

import (
	"context"
	"encoding/json"
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

func StructToMap(obj interface{}) (newMap *map[string]interface{}, err error) {
	data, err := json.Marshal(obj) // Convert to a json string
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &newMap) // Convert to a map
	return
}

func (s PackageService) GetPackage(tenantId string, packageId string) (*api_v1.GetPackageResponse, error) {
	token, _ := domain.GetToken(tenantId)
	ctx := context.Background()
	ApiGetContainerByIdRequest := s.containerApiClient.ContainerApi.GetContainerById(ctx, packageId)
	ApiGetContainerByIdRequest = ApiGetContainerByIdRequest.XCOREOSACCESS(token)
	ApiGetContainerByIdRequest = ApiGetContainerByIdRequest.XCOREOSTID(tenantId)
	ApiGetContainerByIdRequest = ApiGetContainerByIdRequest.XCOREOSREQUESTID("1234")
	ApiGetContainerByIdRequest = ApiGetContainerByIdRequest.XCOREOSUSERINFO("1234")

	d, r, e := s.containerApiClient.ContainerApi.GetContainerByIdExecute(ApiGetContainerByIdRequest)

	if e != nil {
		return nil, e
	}

	if r.StatusCode == 200 {
		data, _ := StructToMap(d.Data)
		response := &api_v1.GetPackageResponse{
			Data: data,
		}
		return response, nil
	}

	return nil, e

}

func (s PackageService) GetPackages(tenantId string) (*api_v1.GetPackagesResponse, error) {

	token, _ := domain.GetToken(tenantId)
	ctx := context.Background()
	ApiGetContainersRequest := s.containerApiClient.ContainerApi.GetContainers(ctx, common.PackageContainerTypeName)
	ApiGetContainersRequest = ApiGetContainersRequest.XCOREOSACCESS(token)
	ApiGetContainersRequest = ApiGetContainersRequest.XCOREOSTID(tenantId)
	ApiGetContainersRequest = ApiGetContainersRequest.XCOREOSREQUESTID("1234")
	ApiGetContainersRequest = ApiGetContainersRequest.XCOREOSUSERINFO("1234")

	d, r, e := s.containerApiClient.ContainerApi.GetContainersExecute(ApiGetContainersRequest)

	if e != nil {
		return nil, e
	}

	if r.StatusCode == 200 {
		data, _ := StructToMap(d.Data)
		response := &api_v1.GetPackagesResponse{
			Data: data,
		}
		return response, nil
	}

	return nil, e

}

func (s PackageService) CreatePackage(tenantId string, request api_v1.CreatePackageJSONRequestBody) (*string, error) {

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
