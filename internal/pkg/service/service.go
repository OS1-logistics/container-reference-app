package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/golang/glog"
	api_v1 "github.com/os1-logistics/container-reference-app/api/v1"
	inits "github.com/os1-logistics/container-reference-app/init"
	"github.com/os1-logistics/container-reference-app/internal/pkg/common"
	domain "github.com/os1-logistics/container-reference-app/internal/pkg/domain"
	containerdomain "github.com/os1-logistics/container-reference-app/internal/pkg/domain/container"
)

type Service struct {
}

func NewPackageService() Service {
	return Service{}
}

func StructToMap(obj interface{}) (newMap *map[string]interface{}, err error) {
	data, err := json.Marshal(obj) // Convert to a json string
	if err != nil {
		glog.Error("StructToMap : ", err)
		return
	}
	err = json.Unmarshal(data, &newMap) // Convert to a map
	glog.Error("StructToMap :", err)
	return
}

func (s Service) Init(tenantId string) error {
	e := inits.Initialize(tenantId)
	if e != nil {
		return e
	}
	return nil
}

func (s Service) GetPackage(tenantId string, packageId string) (*api_v1.GetPackageResponse, error) {
	token, _ := domain.GetToken(tenantId)
	ctx := context.Background()
	ApiGetContainerByIdRequest := domain.NewContainerClient(tenantId).ContainerApi.
		GetContainerById(ctx, packageId, common.PackageContainerTypeName).
		XCOREOSACCESS(token).
		XCOREOSTID(tenantId).
		XCOREOSREQUESTID(common.UUIDv4()).
		XCOREOSUSERINFO("1234")

	d, r, e := domain.NewContainerClient(tenantId).ContainerApi.GetContainerByIdExecute(ApiGetContainerByIdRequest)

	if e != nil {
		glog.Error(r)
		glog.Error(e)
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
	ApiGetContainersRequest := domain.NewContainerClient(tenantId).ContainerApi.
		GetContainers(ctx, common.PackageContainerTypeName).
		XCOREOSACCESS(token).
		XCOREOSTID(tenantId).
		XCOREOSREQUESTID(common.UUIDv4()).
		XCOREOSUSERINFO("1234")

	d, r, e := domain.NewContainerClient(tenantId).ContainerApi.GetContainersExecute(ApiGetContainersRequest)

	if e != nil {
		glog.Error(r)
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

	containerCreateRequest := containerdomain.ContainerCreateRequest{
		ScannableId:       *request.ScannableId,
		IsHazmat:          request.IsHazmat,
		IsContainerizable: containerdomain.PtrBool(true),
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

	ApiCreateContainerRequest := domain.NewContainerClient(tenantId).ContainerApi.
		CreateContainer(ctx, common.PackageContainerTypeName).
		XCOREOSACCESS(token).
		XCOREOSORIGINTOKEN(token).
		XCOREOSTID(tenantId).
		XCOREOSREQUESTID(common.UUIDv4()).
		XCOREOSUSERINFO("1234").
		ContainerCreateRequest(containerCreateRequest)

	d, r, e := domain.NewContainerClient(tenantId).ContainerApi.CreateContainerExecute(ApiCreateContainerRequest)

	if e != nil {
		glog.Error(r)
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

	containerCreateRequest := containerdomain.ContainerCreateRequest{
		ScannableId:       *request.ScannableId,
		IsHazmat:          request.IsHazmat,
		IsContainerizable: containerdomain.PtrBool(false),
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

	ApiCreateContainerRequest := domain.NewContainerClient(tenantId).ContainerApi.
		CreateContainer(ctx, common.BagContainerTypeName).
		XCOREOSACCESS(token).
		XCOREOSORIGINTOKEN(token).
		XCOREOSTID(tenantId).
		XCOREOSREQUESTID(common.UUIDv4()).
		XCOREOSUSERINFO("1234").
		ContainerCreateRequest(containerCreateRequest)

	d, r, e := domain.NewContainerClient(tenantId).ContainerApi.CreateContainerExecute(ApiCreateContainerRequest)

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
	ApiGetContainerByIdRequest := domain.NewContainerClient(tenantId).ContainerApi.
		GetContainerById(ctx, bagId, common.BagContainerTypeName).
		XCOREOSACCESS(token).
		XCOREOSTID(tenantId).
		XCOREOSREQUESTID(common.UUIDv4()).
		XCOREOSUSERINFO("1234")

	d, r, e := domain.NewContainerClient(tenantId).ContainerApi.GetContainerByIdExecute(ApiGetContainerByIdRequest)

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
	ApiGetContainersRequest := domain.NewContainerClient(tenantId).ContainerApi.
		GetContainers(ctx, common.BagContainerTypeName).
		XCOREOSACCESS(token).
		XCOREOSTID(tenantId).
		XCOREOSREQUESTID(common.UUIDv4()).
		XCOREOSUSERINFO("1234")

	d, r, e := domain.NewContainerClient(tenantId).ContainerApi.GetContainersExecute(ApiGetContainersRequest)

	if e != nil {
		glog.Error(r)
		glog.Error(e)
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
func (s Service) UpdateContainerState(tenantId string, containerId string, command string, containerTypeName string) error {

	var operation *common.ContainerStateOperation = nil

	if command == common.STATE_OPERATION_OPEN.Name {
		operation = &common.STATE_OPERATION_OPEN
	} else if command == common.STATE_OPERATION_CLOSE.Name {
		operation = &common.STATE_OPERATION_CLOSE
	} else if command == common.STATE_OPERATION_COMPLETE.Name {
		operation = &common.STATE_OPERATION_COMPLETE
	} else if command == common.STATE_OPERATION_DEAD.Name {
		operation = &common.STATE_OPERATION_DEAD
	} else {
		return fmt.Errorf("Command %s is not supported", command)
	}

	token, _ := domain.GetToken(tenantId)
	ctx := context.Background()

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

	apiUpdateContainerStateRequest := domain.NewContainerClient(tenantId).ContainerStateApi.
		UpdateContainerState(ctx, containerId, containerTypeName).
		XCOREOSACCESS(token).
		XCOREOSTID(tenantId).
		XCOREOSREQUESTID(common.UUIDv4()).
		XCOREOSUSERINFO("1234").
		ContainerStateUpdateRequest(containerStateUpdateRequest)

	_, r, e := domain.NewContainerClient(tenantId).ContainerStateApi.UpdateContainerStateExecute(apiUpdateContainerStateRequest)

	if e != nil {
		return e
	}

	if r.StatusCode == 200 || r.StatusCode == 201 || r.StatusCode == 202 {
		return nil
	}

	return e

}

// containerize / decontainerize operations

func (s Service) ContainerizeOperations(tenantId string, parentId string, containerId string, operation string) error {
	glog.Info("parent: ", parentId)
	glog.Info("child: ", containerId)
	token, _ := domain.GetToken(tenantId)
	ctx := context.Background()

	ParentIdRequest := containerdomain.ParentIdRequest{
		ParentId: parentId,
		Action:   operation,
	}

	apiContainerizeRequest := domain.NewContainerClient(tenantId).ContainerApi.
		ContainerizeContainerById(ctx, containerId).
		XCOREOSACCESS(token).
		XCOREOSTID(tenantId).
		XCOREOSREQUESTID(common.UUIDv4()).
		XCOREOSUSERINFO("1234").
		ParentIdRequest(ParentIdRequest)

	_, r, e := domain.NewContainerClient(tenantId).ContainerApi.ContainerizeContainerByIdExecute(apiContainerizeRequest)

	if e != nil {
		glog.Error(e)
		glog.Infof("%v", r)
		return e
	}

	if r.StatusCode == 200 || r.StatusCode == 201 || r.StatusCode == 202 {
		return nil
	}

	return e

}
