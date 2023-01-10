package inits

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/golang/glog"
	"github.com/os1-logistics/container-reference-app/internal/pkg/common"
	"github.com/os1-logistics/container-reference-app/internal/pkg/domain"
	containerdomain "github.com/os1-logistics/container-reference-app/internal/pkg/domain/container"
)

func getErrorResponse(inputBuff *io.ReadCloser) containerdomain.ErrorResponse {
	r := containerdomain.ErrorResponse{}
	buf := new(bytes.Buffer)
	buf.ReadFrom(*inputBuff)
	json.Unmarshal([]byte(buf.String()), &r)
	return r
}

func Initialize(tenantId string) error {

	glog.Info("Initializing the application with container types")

	containerClient := domain.NewContainerClient(tenantId)
	ctx := context.Background()

	token, _ := domain.GetToken(tenantId)

	// Package type

	getContainerTypeRequest1 := containerClient.ContainerTypeApi.GetContainerTypeById(ctx, common.PackageContainerTypeName)
	getContainerTypeRequest1 = getContainerTypeRequest1.XCOREOSACCESS(token)
	getContainerTypeRequest1 = getContainerTypeRequest1.XCOREOSTID(tenantId)
	getContainerTypeRequest1 = getContainerTypeRequest1.XCOREOSREQUESTID(common.UUIDv4())
	getContainerTypeRequest1 = getContainerTypeRequest1.XCOREOSUSERINFO("1234")

	_, getresponse1, _ := containerClient.ContainerTypeApi.GetContainerTypeByIdExecute(getContainerTypeRequest1)

	if getresponse1.StatusCode == 400 {
		getErrResponse := getErrorResponse(&getresponse1.Body)
		// The container type you are searching/updating undefined doesnt exists
		if getErrResponse.Error.Code == "100920043304" {
			// create container type PackageCsa
			containerTypeRequest := containerClient.ContainerTypeApi.CreateContainerType(ctx)
			containerTypeRequest = containerTypeRequest.XCOREOSACCESS(token)
			containerTypeRequest = containerTypeRequest.XCOREOSTID(tenantId)
			containerTypeRequest = containerTypeRequest.XCOREOSREQUESTID(common.UUIDv4())
			containerTypeRequest = containerTypeRequest.XCOREOSUSERINFO("1234")

			PackageCsaContainerTypeCreateRequest := containerdomain.ContainerTypeCreateRequest{
				Name:   common.PackageContainerTypeName,
				IsLeaf: containerdomain.PtrBool(true),
			}
			containerTypeRequest = containerTypeRequest.ContainerTypeCreateRequest(PackageCsaContainerTypeCreateRequest)

			_, response, err := containerClient.ContainerTypeApi.CreateContainerTypeExecute(containerTypeRequest)
			glog.Info("====================Create PackageCsa type=====================")
			glog.Info(response)

			if err != nil {
				createErrResponse := getErrorResponse(&getresponse1.Body)
				// Request failed with status code 400
				if createErrResponse.Error.Code != "100920021050" {
					return err
				}
			}
		}

	}

	glog.Info("====================PackageCsa Type Created Idempotently=====================")

	glog.Info("====================Create PackageCsa Type attributes=====================")
	ApiUpdateAttributesConfigRequest := containerClient.ContainerTypeAttributesConfigApi.UpdateAttributesConfig(ctx, common.PackageContainerTypeName)
	ApiUpdateAttributesConfigRequest = ApiUpdateAttributesConfigRequest.XCOREOSACCESS(token)
	ApiUpdateAttributesConfigRequest = ApiUpdateAttributesConfigRequest.XCOREOSORIGINTOKEN(token)
	ApiUpdateAttributesConfigRequest = ApiUpdateAttributesConfigRequest.XCOREOSTID(tenantId)
	ApiUpdateAttributesConfigRequest = ApiUpdateAttributesConfigRequest.XCOREOSREQUESTID(common.UUIDv4())
	ApiUpdateAttributesConfigRequest = ApiUpdateAttributesConfigRequest.XCOREOSUSERINFO("1234")

	attributesConfigUpdateRequest := containerdomain.AttributesConfigUpdateRequest{
		Attributes: []containerdomain.AttributeConfig{
			{
				Name:        "containerName",
				DataType:    containerdomain.DATATYPE_STRING,
				Description: containerdomain.PtrString("Container Name - package/ bag"),
				Indexed:     containerdomain.PtrBool(true),
				Validation: &containerdomain.AttributeValidation{
					Regex:    containerdomain.PtrString("^package$"),
					Required: containerdomain.PtrBool(true),
				},
			},
			{
				Name:        "containerType",
				DataType:    containerdomain.DATATYPE_STRING,
				Description: containerdomain.PtrString("Container Type - reusable/ non-reusable"),
				Indexed:     containerdomain.PtrBool(true),
				Validation: &containerdomain.AttributeValidation{
					Regex:    containerdomain.PtrString("^non-reusable$"),
					Required: containerdomain.PtrBool(true),
				},
			},
			{
				Name:        "leaf",
				DataType:    containerdomain.DATATYPE_BOOLEAN,
				Description: containerdomain.PtrString("Is this a leaf container or not"),
				Indexed:     containerdomain.PtrBool(false),
				Validation: &containerdomain.AttributeValidation{
					Required: containerdomain.PtrBool(true),
				},
			},
			{
				Name:        "origin",
				DataType:    containerdomain.DATATYPE_STRING,
				Description: containerdomain.PtrString("Originating location of the container"),
				Indexed:     containerdomain.PtrBool(false),
				Validation: &containerdomain.AttributeValidation{
					Required: containerdomain.PtrBool(true),
				},
			},
			{
				Name:        "destination",
				DataType:    containerdomain.DATATYPE_STRING,
				Description: containerdomain.PtrString("Destination location of the container"),
				Indexed:     containerdomain.PtrBool(false),
				Validation: &containerdomain.AttributeValidation{
					Required: containerdomain.PtrBool(true),
				},
			},
			{
				Name:        "trackingId",
				DataType:    containerdomain.DATATYPE_STRING,
				Description: containerdomain.PtrString("Tracking/ Scannable IDs (can be list - Order ID, LS ID, Container ID)"),
				Indexed:     containerdomain.PtrBool(true),
				Validation: &containerdomain.AttributeValidation{
					Required: containerdomain.PtrBool(true),
				},
			},
			{
				Name:        "volume",
				DataType:    containerdomain.DATATYPE_NUMBER,
				Description: containerdomain.PtrString("Volume of the container"),
				Indexed:     containerdomain.PtrBool(false),
				Validation: &containerdomain.AttributeValidation{
					Required: containerdomain.PtrBool(false),
				},
			},
			{
				Name:        "deadWeight",
				DataType:    containerdomain.DATATYPE_NUMBER,
				Description: containerdomain.PtrString("Dead weight of the container"),
				Indexed:     containerdomain.PtrBool(false),
				Validation: &containerdomain.AttributeValidation{
					Required: containerdomain.PtrBool(false),
				},
			},
			{
				Name:        "deadWeignt",
				DataType:    containerdomain.DATATYPE_NUMBER,
				Description: containerdomain.PtrString("Dead weight of the container"),
				Indexed:     containerdomain.PtrBool(false),
				Validation: &containerdomain.AttributeValidation{
					Required: containerdomain.PtrBool(false),
				},
			},
			{
				Name:        "orderDescription",
				DataType:    containerdomain.DATATYPE_STRING,
				Description: containerdomain.PtrString("Description of the order"),
				Indexed:     containerdomain.PtrBool(false),
				Validation: &containerdomain.AttributeValidation{
					Required: containerdomain.PtrBool(false),
				},
			},
			{
				Name:        "promiseDateOfDelivery",
				DataType:    containerdomain.DATATYPE_STRING,
				Description: containerdomain.PtrString("Promise date of delivery"),
				Indexed:     containerdomain.PtrBool(false),
				Validation: &containerdomain.AttributeValidation{
					Required: containerdomain.PtrBool(false),
				},
			},
			{
				Name:        "nextTransitCenter",
				DataType:    containerdomain.DATATYPE_STRING,
				Description: containerdomain.PtrString("Next transit center"),
				Indexed:     containerdomain.PtrBool(false),
				Validation: &containerdomain.AttributeValidation{
					Required: containerdomain.PtrBool(false),
				},
			},
		},
	}

	ApiUpdateAttributesConfigRequest = ApiUpdateAttributesConfigRequest.AttributesConfigUpdateRequest(attributesConfigUpdateRequest)
	_, r, e := containerClient.ContainerTypeAttributesConfigApi.UpdateAttributesConfigExecute(ApiUpdateAttributesConfigRequest)
	if e != nil {
		glog.Error(e)
	}

	if r.StatusCode == 202 {
		glog.Info("Package Attributes updated accepted successfully")
	}

	// Bag type

	getContainerTypeRequest2 := containerClient.ContainerTypeApi.GetContainerTypeById(ctx, common.BagContainerTypeName)
	getContainerTypeRequest2 = getContainerTypeRequest2.XCOREOSACCESS(token)
	getContainerTypeRequest2 = getContainerTypeRequest2.XCOREOSTID(tenantId)
	getContainerTypeRequest2 = getContainerTypeRequest2.XCOREOSREQUESTID(common.UUIDv4())
	getContainerTypeRequest2 = getContainerTypeRequest2.XCOREOSUSERINFO("1234")

	_, getresponse2, _ := containerClient.ContainerTypeApi.GetContainerTypeByIdExecute(getContainerTypeRequest2)

	if getresponse2.StatusCode == 400 {
		errResponse := getErrorResponse(&getresponse2.Body)
		if errResponse.Error.Description == "The container type you are searching/updating undefined doesnt exists" {
			// create container type BagCsa
			containerTypeRequest := containerClient.ContainerTypeApi.CreateContainerType(ctx)
			containerTypeRequest = containerTypeRequest.XCOREOSACCESS(token)
			containerTypeRequest = containerTypeRequest.XCOREOSTID(tenantId)
			containerTypeRequest = containerTypeRequest.XCOREOSREQUESTID(common.UUIDv4())
			containerTypeRequest = containerTypeRequest.XCOREOSUSERINFO("1234")

			BagContainerTypeCreateRequest := containerdomain.ContainerTypeCreateRequest{
				Name:   common.BagContainerTypeName,
				IsLeaf: containerdomain.PtrBool(false),
			}

			containerTypeRequest = containerTypeRequest.ContainerTypeCreateRequest(BagContainerTypeCreateRequest)

			_, response, err := containerClient.ContainerTypeApi.CreateContainerTypeExecute(containerTypeRequest)
			glog.Info("====================Create BagCsa type=====================")
			glog.Info(response)

			if err != nil {
				createErrResponse := getErrorResponse(&getresponse2.Body)
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

	glog.Info("====================Create BagCsa Type attributes=====================")
	ApiUpdateAttributesConfigRequest = containerClient.ContainerTypeAttributesConfigApi.UpdateAttributesConfig(ctx, common.BagContainerTypeName)
	ApiUpdateAttributesConfigRequest = ApiUpdateAttributesConfigRequest.XCOREOSACCESS(token)
	ApiUpdateAttributesConfigRequest = ApiUpdateAttributesConfigRequest.XCOREOSORIGINTOKEN(token)
	ApiUpdateAttributesConfigRequest = ApiUpdateAttributesConfigRequest.XCOREOSTID(tenantId)
	ApiUpdateAttributesConfigRequest = ApiUpdateAttributesConfigRequest.XCOREOSREQUESTID(common.UUIDv4())
	ApiUpdateAttributesConfigRequest = ApiUpdateAttributesConfigRequest.XCOREOSUSERINFO("1234")

	attributesConfigUpdateRequest = containerdomain.AttributesConfigUpdateRequest{
		Attributes: []containerdomain.AttributeConfig{
			{
				Name:        "containerName",
				DataType:    containerdomain.DATATYPE_STRING,
				Description: containerdomain.PtrString("Container Name - package/ bag"),
				Indexed:     containerdomain.PtrBool(true),
				Validation: &containerdomain.AttributeValidation{
					Regex:    containerdomain.PtrString("^bag$"),
					Required: containerdomain.PtrBool(true),
				},
			},
			{
				Name:        "containerType",
				DataType:    containerdomain.DATATYPE_STRING,
				Description: containerdomain.PtrString("Container Type - reusable/ non-reusable"),
				Indexed:     containerdomain.PtrBool(true),
				Validation: &containerdomain.AttributeValidation{
					Regex:    containerdomain.PtrString("^reusable$"),
					Required: containerdomain.PtrBool(true),
				},
			},
			{
				Name:        "leaf",
				DataType:    containerdomain.DATATYPE_BOOLEAN,
				Description: containerdomain.PtrString("Is this a leaf container or not"),
				Indexed:     containerdomain.PtrBool(false),
				Validation: &containerdomain.AttributeValidation{
					Required: containerdomain.PtrBool(true),
				},
			},
			{
				Name:        "allowableContainers",
				DataType:    containerdomain.DATATYPE_STRING,
				Description: containerdomain.PtrString("Allowable containers - package/ bag"),
				Indexed:     containerdomain.PtrBool(false),
				Validation: &containerdomain.AttributeValidation{
					Required: containerdomain.PtrBool(false),
				},
			},
			{
				Name:        "origin",
				DataType:    containerdomain.DATATYPE_STRING,
				Description: containerdomain.PtrString("Originating location of the container"),
				Indexed:     containerdomain.PtrBool(false),
				Validation: &containerdomain.AttributeValidation{
					Required: containerdomain.PtrBool(true),
				},
			},
			{
				Name:        "destination",
				DataType:    containerdomain.DATATYPE_STRING,
				Description: containerdomain.PtrString("Destination location of the container"),
				Indexed:     containerdomain.PtrBool(false),
				Validation: &containerdomain.AttributeValidation{
					Required: containerdomain.PtrBool(true),
				},
			},
			{
				Name:        "trackingId",
				DataType:    containerdomain.DATATYPE_STRING,
				Description: containerdomain.PtrString("Tracking/ Scannable IDs (can be list - Order ID, LS ID, Container ID)"),
				Indexed:     containerdomain.PtrBool(true),
				Validation: &containerdomain.AttributeValidation{
					Required: containerdomain.PtrBool(true),
				},
			},
			{
				Name:        "volume",
				DataType:    containerdomain.DATATYPE_NUMBER,
				Description: containerdomain.PtrString("Volume of the container"),
				Indexed:     containerdomain.PtrBool(false),
				Validation: &containerdomain.AttributeValidation{
					Required: containerdomain.PtrBool(false),
				},
			},
			{
				Name:        "deadWeight",
				DataType:    containerdomain.DATATYPE_NUMBER,
				Description: containerdomain.PtrString("Dead weight of the container"),
				Indexed:     containerdomain.PtrBool(false),
				Validation: &containerdomain.AttributeValidation{
					Required: containerdomain.PtrBool(false),
				},
			},
			{
				Name:        "packagesInsideTheBag",
				DataType:    containerdomain.DATATYPE_ARRAY,
				Description: containerdomain.PtrString("Packages inside the bag"),
				Indexed:     containerdomain.PtrBool(false),
				Validation: &containerdomain.AttributeValidation{
					Required: containerdomain.PtrBool(false),
				},
			},
			{
				Name:        "nextTransitCenter",
				DataType:    containerdomain.DATATYPE_STRING,
				Description: containerdomain.PtrString("Next transit center"),
				Indexed:     containerdomain.PtrBool(false),
				Validation: &containerdomain.AttributeValidation{
					Required: containerdomain.PtrBool(false),
				},
			},
		},
	}

	ApiUpdateAttributesConfigRequest = ApiUpdateAttributesConfigRequest.AttributesConfigUpdateRequest(attributesConfigUpdateRequest)
	_, r, e = containerClient.ContainerTypeAttributesConfigApi.UpdateAttributesConfigExecute(ApiUpdateAttributesConfigRequest)
	if e != nil {
		glog.Error(e)
	}

	if r.StatusCode == 202 {
		glog.Info("Bag Attributes updated accepted successfully")
	}

	return nil
}
