package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/golang/glog"
	"github.com/os1-logistics/container-reference-app/internal/pkg/domain"
	containerdomain "github.com/os1-logistics/container-reference-app/internal/pkg/domain/container"
)

func GetErrorResponse(inputBuff *io.ReadCloser) containerdomain.ErrorResponse {
	r := containerdomain.ErrorResponse{}
	buf := new(bytes.Buffer)
	buf.ReadFrom(*inputBuff)
	json.Unmarshal([]byte(buf.String()), &r)
	return r
}

func Initialize(tenantId string) error {

	glog.Info("Initializing the application with container types")

	// package type for container sample app
	const packageOrderContainerTypeName = "PackageCsa"
	// bag type for container sample app
	const bagContainerTypeName = "BagCsa"
	containerClient := domain.NewContainerClient(tenantId)
	ctx := context.Background()

	token, _ := domain.GetToken(tenantId)

	getContainerTypeRequest1 := containerClient.ContainerTypeApi.GetContainerTypeById(ctx, packageOrderContainerTypeName)
	getContainerTypeRequest1 = getContainerTypeRequest1.XCOREOSACCESS(token)
	getContainerTypeRequest1 = getContainerTypeRequest1.XCOREOSTID(tenantId)
	getContainerTypeRequest1 = getContainerTypeRequest1.XCOREOSREQUESTID("1234")
	getContainerTypeRequest1 = getContainerTypeRequest1.XCOREOSUSERINFO("1234")

	_, getresponse1, _ := containerClient.ContainerTypeApi.GetContainerTypeByIdExecute(getContainerTypeRequest1)

	if getresponse1.StatusCode == 400 {
		getErrResponse := GetErrorResponse(&getresponse1.Body)
		// The container type you are searching/updating undefined doesnt exists
		if getErrResponse.Error.Code == "100920043304" {
			// create container type PackageCsa
			containerTypeRequest := containerClient.ContainerTypeApi.CreateContainerType(ctx)
			containerTypeRequest = containerTypeRequest.XCOREOSACCESS(token)
			containerTypeRequest = containerTypeRequest.XCOREOSTID(tenantId)
			containerTypeRequest = containerTypeRequest.XCOREOSREQUESTID("1234")
			containerTypeRequest = containerTypeRequest.XCOREOSUSERINFO("1234")

			PackageCsaContainerTypeCreateRequest := containerdomain.ContainerTypeCreateRequest{
				Name:   packageOrderContainerTypeName,
				IsLeaf: containerdomain.PtrBool(true),
			}
			containerTypeRequest = containerTypeRequest.ContainerTypeCreateRequest(PackageCsaContainerTypeCreateRequest)

			_, response, err := containerClient.ContainerTypeApi.CreateContainerTypeExecute(containerTypeRequest)
			glog.Info("====================Create PackageCsa type=====================")
			glog.Info(response)

			if err != nil {
				createErrResponse := GetErrorResponse(&getresponse1.Body)
				// Request failed with status code 400
				if createErrResponse.Error.Code != "100920021050" {
					return err
				}
			}
		}
	}

	glog.Info("====================PackageCsa Type Created Idempotently=====================")

	getContainerTypeRequest2 := containerClient.ContainerTypeApi.GetContainerTypeById(ctx, bagContainerTypeName)
	getContainerTypeRequest2 = getContainerTypeRequest2.XCOREOSACCESS(token)
	getContainerTypeRequest2 = getContainerTypeRequest2.XCOREOSTID(tenantId)
	getContainerTypeRequest2 = getContainerTypeRequest2.XCOREOSREQUESTID("1234")
	getContainerTypeRequest2 = getContainerTypeRequest2.XCOREOSUSERINFO("1234")

	_, getresponse2, _ := containerClient.ContainerTypeApi.GetContainerTypeByIdExecute(getContainerTypeRequest2)

	if getresponse2.StatusCode == 400 {
		errResponse := GetErrorResponse(&getresponse2.Body)
		if errResponse.Error.Description == "The container type you are searching/updating undefined doesnt exists" {
			// create container type PackageCsa
			containerTypeRequest := containerClient.ContainerTypeApi.CreateContainerType(ctx)
			containerTypeRequest = containerTypeRequest.XCOREOSACCESS(token)
			containerTypeRequest = containerTypeRequest.XCOREOSTID(tenantId)
			containerTypeRequest = containerTypeRequest.XCOREOSREQUESTID("1234")
			containerTypeRequest = containerTypeRequest.XCOREOSUSERINFO("1234")

			BagContainerTypeCreateRequest := containerdomain.ContainerTypeCreateRequest{
				Name:   bagContainerTypeName,
				IsLeaf: containerdomain.PtrBool(false),
			}

			containerTypeRequest = containerTypeRequest.ContainerTypeCreateRequest(BagContainerTypeCreateRequest)

			_, response, err := containerClient.ContainerTypeApi.CreateContainerTypeExecute(containerTypeRequest)
			glog.Info("====================Create BagCsa type=====================")
			glog.Info(response)

			if err != nil {
				createErrResponse := GetErrorResponse(&getresponse2.Body)
				// Request failed with status code 400
				if createErrResponse.Error.Code != "100920021050" {
					return err
				}
			}
		} else {
			return fmt.Errorf("Error while getting container type Bag: %s", errResponse.Error.Description)
		}
	}
	glog.Info("====================BagCsa Type Created Idempotently=====================")
	return nil
}
