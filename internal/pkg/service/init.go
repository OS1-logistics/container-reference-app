package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"

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

	fmt.Println("Initializing the application with container types")

	const shipmentOrderContainerTypeName = "ShipmentOrder"
	const bagContainerTypeName = "Bag"
	containerClient := domain.NewContainerClient(tenantId)
	ctx := context.Background()

	token, _ := domain.GetToken(tenantId)

	getContainerTypeRequest1 := containerClient.ContainerTypeApi.GetContainerTypeById(ctx, shipmentOrderContainerTypeName)
	getContainerTypeRequest1 = getContainerTypeRequest1.XCOREOSACCESS(token)
	getContainerTypeRequest1 = getContainerTypeRequest1.XCOREOSTID(tenantId)
	getContainerTypeRequest1 = getContainerTypeRequest1.XCOREOSREQUESTID("1234")
	getContainerTypeRequest1 = getContainerTypeRequest1.XCOREOSUSERINFO("1234")

	_, getresponse1, _ := containerClient.ContainerTypeApi.GetContainerTypeByIdExecute(getContainerTypeRequest1)

	if getresponse1.StatusCode == 400 {
		getErrResponse := GetErrorResponse(&getresponse1.Body)
		// The container type you are searching/updating undefined doesnt exists
		if getErrResponse.Error.Code == "100920043304" {
			// create container type ShipmentOrder
			containerTypeRequest := containerClient.ContainerTypeApi.CreateContainerType(ctx)
			containerTypeRequest = containerTypeRequest.XCOREOSACCESS(token)
			containerTypeRequest = containerTypeRequest.XCOREOSTID(tenantId)
			containerTypeRequest = containerTypeRequest.XCOREOSREQUESTID("1234")
			containerTypeRequest = containerTypeRequest.XCOREOSUSERINFO("1234")

			ShipmentOrderContainerTypeCreateRequest := containerdomain.ContainerTypeCreateRequest{
				Name:   shipmentOrderContainerTypeName,
				IsLeaf: containerdomain.PtrBool(true),
			}
			containerTypeRequest = containerTypeRequest.ContainerTypeCreateRequest(ShipmentOrderContainerTypeCreateRequest)

			_, response, err := containerClient.ContainerTypeApi.CreateContainerTypeExecute(containerTypeRequest)
			fmt.Println("====================Create ShipmentOrder type=====================")
			fmt.Println(response)

			if err != nil {
				createErrResponse := GetErrorResponse(&getresponse1.Body)
				// Request failed with status code 400
				if createErrResponse.Error.Code != "100920021050" {
					return err
				}
			}
		}
	}

	fmt.Println("====================ShipmentOrder Type Created Idempotently=====================")

	getContainerTypeRequest2 := containerClient.ContainerTypeApi.GetContainerTypeById(ctx, bagContainerTypeName)
	getContainerTypeRequest2 = getContainerTypeRequest2.XCOREOSACCESS(token)
	getContainerTypeRequest2 = getContainerTypeRequest2.XCOREOSTID(tenantId)
	getContainerTypeRequest2 = getContainerTypeRequest2.XCOREOSREQUESTID("1234")
	getContainerTypeRequest2 = getContainerTypeRequest2.XCOREOSUSERINFO("1234")

	_, getresponse2, _ := containerClient.ContainerTypeApi.GetContainerTypeByIdExecute(getContainerTypeRequest2)

	if getresponse2.StatusCode == 400 {
		errResponse := GetErrorResponse(&getresponse2.Body)
		if errResponse.Error.Description == "The container type you are searching/updating undefined doesnt exists" {
			// create container type ShipmentOrder
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
			fmt.Println("====================Create Bag type=====================")
			fmt.Println(response)

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
	fmt.Println("====================Bag Type Created Idempotently=====================")
	return nil
}
