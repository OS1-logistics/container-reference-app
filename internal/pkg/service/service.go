package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/golang/glog"
	api_v1 "github.com/os1-logistics/container-reference-app/api/v1"
	"github.com/os1-logistics/container-reference-app/internal/pkg/common"
	domain "github.com/os1-logistics/container-reference-app/internal/pkg/domain"
	container "github.com/os1-logistics/container-reference-app/internal/pkg/domain/container"
	containerdomain "github.com/os1-logistics/container-reference-app/internal/pkg/domain/container"
)

type Service struct {
	containerApiClient container.APIClient
}

func NewPackageService() Service {
	containerClient := domain.NewContainerClient("alpha")
	return Service{
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

func (s Service) GetPackage(tenantId string, packageId string) (*api_v1.GetPackageResponse, error) {
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

func (s Service) GetPackages(tenantId string) (*api_v1.GetPackagesResponse, error) {

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

func (s Service) CreatePackage(tenantId string, request api_v1.CreatePackageJSONRequestBody) (*string, error) {

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
		return nil, e
	}

	if r.StatusCode == 202 {
		return containerdomain.PtrString(*d.Data.Id), nil
	}

	return nil, e
}

// bag

func (s Service) CreateBag(tenantId string, request api_v1.CreateBagJSONRequestBody) (*string, error) {

	token, _ := domain.GetToken(tenantId)
	ctx := context.Background()
	ApiCreateContainerRequest := s.containerApiClient.ContainerApi.CreateContainer(ctx, common.BagContainerTypeName)
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
			"containerName":       "bag",
			"containerType":       "reusable",
			"leaf":                false,
			"origin":              request.Origin,
			"destination":         request.Destination,
			"trackingId":          request.TrackingId,
			"allowableContainers": "",
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
		return nil, e
	}

	if r.StatusCode == 202 {
		return containerdomain.PtrString(*d.Data.Id), nil
	}

	return nil, e
}

func (s Service) GetBag(tenantId string, bagId string) (*api_v1.GetBagResponse, error) {
	token, _ := domain.GetToken(tenantId)
	ctx := context.Background()
	ApiGetContainerByIdRequest := s.containerApiClient.ContainerApi.GetContainerById(ctx, bagId)
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
		response := &api_v1.GetBagResponse{
			Data: data,
		}
		return response, nil
	}

	return nil, e

}

func (s Service) GetBags(tenantId string) (*api_v1.GetBagsResponse, error) {

	token, _ := domain.GetToken(tenantId)
	ctx := context.Background()
	ApiGetContainersRequest := s.containerApiClient.ContainerApi.GetContainers(ctx, common.BagContainerTypeName)
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
		response := &api_v1.GetBagsResponse{
			Data: data,
		}
		return response, nil
	}

	return nil, e

}

// container operations
func (s Service) UpdateContainerState(tenantId string, containerId string, command string) error {

	var operation *common.Operation = nil

	if command == common.OPERATION_OPEN.Name {
		operation = &common.OPERATION_OPEN
	} else if command == common.OPERATION_CLOSE.Name {
		operation = &common.OPERATION_CLOSE
	} else if command == common.OPERATION_COMPLETE.Name {
		operation = &common.OPERATION_COMPLETE
	} else if command == common.OPERATION_DEAD.Name {
		operation = &common.OPERATION_DEAD
	} else {
		return fmt.Errorf("Command %s is not supported", command)
	}

	token, _ := domain.GetToken(tenantId)
	ctx := context.Background()
	apiUpdateContainerStateRequest := s.containerApiClient.ContainerStateApi.UpdateContainerState(ctx, containerId)
	apiUpdateContainerStateRequest = apiUpdateContainerStateRequest.XCOREOSACCESS(token)
	apiUpdateContainerStateRequest = apiUpdateContainerStateRequest.XCOREOSTID(tenantId)
	apiUpdateContainerStateRequest = apiUpdateContainerStateRequest.XCOREOSREQUESTID("1234")
	apiUpdateContainerStateRequest = apiUpdateContainerStateRequest.XCOREOSUSERINFO("1234")

	containerStateUpdateRequest := containerdomain.ContainerStateUpdateRequest{
		EventCode:  operation.EventCode,
		ReasonCode: containerdomain.PtrString(operation.ReasonCode),
		Timestamp:  int32(time.Now().UTC().UnixMicro()),
		Data:       map[string]interface{}{},
		Source: containerdomain.EventSource{
			AppId:  common.AppName,
			UserId: containerdomain.PtrString("1234"),
			LocId:  containerdomain.PtrString("1234"),
		},
	}
	apiUpdateContainerStateRequest = apiUpdateContainerStateRequest.ContainerStateUpdateRequest(containerStateUpdateRequest)

	_, r, e := s.containerApiClient.ContainerStateApi.UpdateContainerStateExecute(apiUpdateContainerStateRequest)

	if e != nil {
		return e
	}

	if r.StatusCode == 200 || r.StatusCode == 201 || r.StatusCode == 202 {
		return nil
	}

	return e

}
