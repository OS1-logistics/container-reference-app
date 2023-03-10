/*
Container Service

**API documentation for Container Service**

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package containerdomain

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// ContainerStateApiService ContainerStateApi service
type ContainerStateApiService service

type ApiGetContainerStateRequest struct {
	ctx context.Context
	ApiService *ContainerStateApiService
	xCOREOSREQUESTID *string
	xCOREOSTID *string
	xCOREOSACCESS *string
	containerId string
	containerTypeName string
	xCOREOSUSERINFO *string
}

// Unique request id.
func (r ApiGetContainerStateRequest) XCOREOSREQUESTID(xCOREOSREQUESTID string) ApiGetContainerStateRequest {
	r.xCOREOSREQUESTID = &xCOREOSREQUESTID
	return r
}

// Tenant id
func (r ApiGetContainerStateRequest) XCOREOSTID(xCOREOSTID string) ApiGetContainerStateRequest {
	r.xCOREOSTID = &xCOREOSTID
	return r
}

// Core-os access token
func (r ApiGetContainerStateRequest) XCOREOSACCESS(xCOREOSACCESS string) ApiGetContainerStateRequest {
	r.xCOREOSACCESS = &xCOREOSACCESS
	return r
}

// Core-os user info
func (r ApiGetContainerStateRequest) XCOREOSUSERINFO(xCOREOSUSERINFO string) ApiGetContainerStateRequest {
	r.xCOREOSUSERINFO = &xCOREOSUSERINFO
	return r
}

func (r ApiGetContainerStateRequest) Execute() (*ContainerStateResponse, *http.Response, error) {
	return r.ApiService.GetContainerStateExecute(r)
}

/*
GetContainerState Get the current state of a container

**API to get the current state of a container.**

 | HEADER/PATH | DESCRIPTION | TYPE |
 | ---- | ---- | ----- |
 | `X-COREOS-REQUEST-ID`*(header) | Unique request ID. | string |
 | `X-COREOS-TID`*(header) | Tenant ID. | string |
 | `X-COREOS-ACCESS`*(header) | Core-OS access token. | string |
 
 | `containerId`*(path) | Unique ID of Container. | string|
 *This is required.
   <br><br>
   **HTTP Status Code Summary**
   | Code | Description |
   | ---- | ---------- |
   | **200 - OK** | Everything worked as expected. |
   | **400 - Bad Request** | The request was unacceptable, often due to missing a required parameter. |
   | **401 - Unauthorized** | No valid API key provided. |
   | **404 - Not Found** | The requested resource doesn't exist. |
   | **503 - Server Error** | Something went wrong.. (These are rare.) |


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param containerId Unique id of Container
 @param containerTypeName A unique name to represent type of a container.
 @return ApiGetContainerStateRequest
*/
func (a *ContainerStateApiService) GetContainerState(ctx context.Context, containerId string, containerTypeName string) ApiGetContainerStateRequest {
	return ApiGetContainerStateRequest{
		ApiService: a,
		ctx: ctx,
		containerId: containerId,
		containerTypeName: containerTypeName,
	}
}

// Execute executes the request
//  @return ContainerStateResponse
func (a *ContainerStateApiService) GetContainerStateExecute(r ApiGetContainerStateRequest) (*ContainerStateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ContainerStateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContainerStateApiService.GetContainerState")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{containerTypeName}/{containerId}/state"
	localVarPath = strings.Replace(localVarPath, "{"+"containerId"+"}", url.PathEscape(parameterToString(r.containerId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"containerTypeName"+"}", url.PathEscape(parameterToString(r.containerTypeName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xCOREOSREQUESTID == nil {
		return localVarReturnValue, nil, reportError("xCOREOSREQUESTID is required and must be specified")
	}
	if r.xCOREOSTID == nil {
		return localVarReturnValue, nil, reportError("xCOREOSTID is required and must be specified")
	}
	if r.xCOREOSACCESS == nil {
		return localVarReturnValue, nil, reportError("xCOREOSACCESS is required and must be specified")
	}
	if strlen(r.containerTypeName) < 3 {
		return localVarReturnValue, nil, reportError("containerTypeName must have at least 3 elements")
	}
	if strlen(r.containerTypeName) > 64 {
		return localVarReturnValue, nil, reportError("containerTypeName must have less than 64 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["X-COREOS-REQUEST-ID"] = parameterToString(*r.xCOREOSREQUESTID, "")
	localVarHeaderParams["X-COREOS-TID"] = parameterToString(*r.xCOREOSTID, "")
	localVarHeaderParams["X-COREOS-ACCESS"] = parameterToString(*r.xCOREOSACCESS, "")
	if r.xCOREOSUSERINFO != nil {
		localVarHeaderParams["X-COREOS-USERINFO"] = parameterToString(*r.xCOREOSUSERINFO, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateContainerStateRequest struct {
	ctx context.Context
	ApiService *ContainerStateApiService
	xCOREOSREQUESTID *string
	xCOREOSTID *string
	xCOREOSACCESS *string
	containerId string
	containerTypeName string
	containerStateUpdateRequest *ContainerStateUpdateRequest
	xCOREOSUSERINFO *string
}

// Unique request id.
func (r ApiUpdateContainerStateRequest) XCOREOSREQUESTID(xCOREOSREQUESTID string) ApiUpdateContainerStateRequest {
	r.xCOREOSREQUESTID = &xCOREOSREQUESTID
	return r
}

// Tenant id
func (r ApiUpdateContainerStateRequest) XCOREOSTID(xCOREOSTID string) ApiUpdateContainerStateRequest {
	r.xCOREOSTID = &xCOREOSTID
	return r
}

// Core-os access token
func (r ApiUpdateContainerStateRequest) XCOREOSACCESS(xCOREOSACCESS string) ApiUpdateContainerStateRequest {
	r.xCOREOSACCESS = &xCOREOSACCESS
	return r
}

// Event for container state transition.
func (r ApiUpdateContainerStateRequest) ContainerStateUpdateRequest(containerStateUpdateRequest ContainerStateUpdateRequest) ApiUpdateContainerStateRequest {
	r.containerStateUpdateRequest = &containerStateUpdateRequest
	return r
}

// Core-os user info
func (r ApiUpdateContainerStateRequest) XCOREOSUSERINFO(xCOREOSUSERINFO string) ApiUpdateContainerStateRequest {
	r.xCOREOSUSERINFO = &xCOREOSUSERINFO
	return r
}

func (r ApiUpdateContainerStateRequest) Execute() (*ContainerStateResponse, *http.Response, error) {
	return r.ApiService.UpdateContainerStateExecute(r)
}

/*
UpdateContainerState Apply event on the container.

**This API will apply an event to an instance. If the event is
transitionable, then the instance will transition to the particular
state. Otherwise it will remain in the existing state.**

| PARAMETER | DESCRIPTION | TYPE | VALIDATION |
| ---- | ---- | ----- | ---- |
| `name`* | Represents name of the event for state transition. | string | **minLength**: 3 **maxLength**: 64 |
| `timestamp` | Represents the timestamp of the event occured. | integer |
| `propagate` | Propagate flag indicates whether these events need to propagate on child containers or not. | boolean | **Default**: false |
| `data` | Represents event data values, if any, for this event. |
| `source` | Represents the source which triggered the event. It can be an app, a user or some location from where the event was triggered. |
| ??????`appId`* | Application ID which is responsible for calling this event. | string | **minLength**: 1 **maxLength**: 64 |
 | ??????`userId` | User ID which is responsible for calling this event. | string | **minLength**: 1 **maxLength**: 64 |
 | ??????`locId` | LocationId of the event. | string |
 | `callback` | Represents callback URL, which can be called to get any info about state change of a container and notify status(success/failed) of API. | string | A valid URL. |
 *This is required.
  <br><br>
  **HTTP Status Code Summary**
  | Code | Description |
  | ---- | ---------- |
  | **200 - OK** | Everything worked as expected. |
  | **400 - Bad Request** | The request was unacceptable, often due to missing a required parameter. |
  | **401 - Unauthorized** | No valid API key provided. |
  | **404 - Not Found** | The requested resource doesn't exist. |
  | **503 - Server Error** | Something went wrong.. (These are rare.) |


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param containerId Unique id of Container
 @param containerTypeName A unique name to represent type of a container.
 @return ApiUpdateContainerStateRequest
*/
func (a *ContainerStateApiService) UpdateContainerState(ctx context.Context, containerId string, containerTypeName string) ApiUpdateContainerStateRequest {
	return ApiUpdateContainerStateRequest{
		ApiService: a,
		ctx: ctx,
		containerId: containerId,
		containerTypeName: containerTypeName,
	}
}

// Execute executes the request
//  @return ContainerStateResponse
func (a *ContainerStateApiService) UpdateContainerStateExecute(r ApiUpdateContainerStateRequest) (*ContainerStateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ContainerStateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContainerStateApiService.UpdateContainerState")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{containerTypeName}/{containerId}/state/event"
	localVarPath = strings.Replace(localVarPath, "{"+"containerId"+"}", url.PathEscape(parameterToString(r.containerId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"containerTypeName"+"}", url.PathEscape(parameterToString(r.containerTypeName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xCOREOSREQUESTID == nil {
		return localVarReturnValue, nil, reportError("xCOREOSREQUESTID is required and must be specified")
	}
	if r.xCOREOSTID == nil {
		return localVarReturnValue, nil, reportError("xCOREOSTID is required and must be specified")
	}
	if r.xCOREOSACCESS == nil {
		return localVarReturnValue, nil, reportError("xCOREOSACCESS is required and must be specified")
	}
	if strlen(r.containerTypeName) < 3 {
		return localVarReturnValue, nil, reportError("containerTypeName must have at least 3 elements")
	}
	if strlen(r.containerTypeName) > 64 {
		return localVarReturnValue, nil, reportError("containerTypeName must have less than 64 elements")
	}
	if r.containerStateUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("containerStateUpdateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["X-COREOS-REQUEST-ID"] = parameterToString(*r.xCOREOSREQUESTID, "")
	localVarHeaderParams["X-COREOS-TID"] = parameterToString(*r.xCOREOSTID, "")
	localVarHeaderParams["X-COREOS-ACCESS"] = parameterToString(*r.xCOREOSACCESS, "")
	if r.xCOREOSUSERINFO != nil {
		localVarHeaderParams["X-COREOS-USERINFO"] = parameterToString(*r.xCOREOSUSERINFO, "")
	}
	// body params
	localVarPostBody = r.containerStateUpdateRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
