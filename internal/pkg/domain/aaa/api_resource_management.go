/*
Authentication And Authorization (AAA) Service

This swagger documentation provides all AAA API details. AAA service provides authentication and authorization capabilities for users.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aaadomain

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// ResourceManagementApiService ResourceManagementApi service
type ResourceManagementApiService service

type ApiCreateResourceRequest struct {
	ctx context.Context
	ApiService *ResourceManagementApiService
	appId string
	xCOREOSREQUESTID *string
	xCOREOSACCESS *string
	xCOREOSTID *string
	resourceCreateRequest *ResourceCreateRequest
	xCOREOSAUTH *string
	xCOREOSUSERINFO *string
}

// Unique request ID
func (r ApiCreateResourceRequest) XCOREOSREQUESTID(xCOREOSREQUESTID string) ApiCreateResourceRequest {
	r.xCOREOSREQUESTID = &xCOREOSREQUESTID
	return r
}

// Core-os access token
func (r ApiCreateResourceRequest) XCOREOSACCESS(xCOREOSACCESS string) ApiCreateResourceRequest {
	r.xCOREOSACCESS = &xCOREOSACCESS
	return r
}

// Tenant ID
func (r ApiCreateResourceRequest) XCOREOSTID(xCOREOSTID string) ApiCreateResourceRequest {
	r.xCOREOSTID = &xCOREOSTID
	return r
}

// 
func (r ApiCreateResourceRequest) ResourceCreateRequest(resourceCreateRequest ResourceCreateRequest) ApiCreateResourceRequest {
	r.resourceCreateRequest = &resourceCreateRequest
	return r
}

// Core-os authentication token.
func (r ApiCreateResourceRequest) XCOREOSAUTH(xCOREOSAUTH string) ApiCreateResourceRequest {
	r.xCOREOSAUTH = &xCOREOSAUTH
	return r
}

// User info
func (r ApiCreateResourceRequest) XCOREOSUSERINFO(xCOREOSUSERINFO string) ApiCreateResourceRequest {
	r.xCOREOSUSERINFO = &xCOREOSUSERINFO
	return r
}

func (r ApiCreateResourceRequest) Execute() (*SuccessResponse, *http.Response, error) {
	return r.ApiService.CreateResourceExecute(r)
}

/*
CreateResource Create a resource

API is used to create a resource for an app. Based on resource we can grant access to APIs

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appId
 @return ApiCreateResourceRequest
*/
func (a *ResourceManagementApiService) CreateResource(ctx context.Context, appId string) ApiCreateResourceRequest {
	return ApiCreateResourceRequest{
		ApiService: a,
		ctx: ctx,
		appId: appId,
	}
}

// Execute executes the request
//  @return SuccessResponse
func (a *ResourceManagementApiService) CreateResourceExecute(r ApiCreateResourceRequest) (*SuccessResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SuccessResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourceManagementApiService.CreateResource")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/aaa/apps/{appId}/resources"
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterToString(r.appId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xCOREOSREQUESTID == nil {
		return localVarReturnValue, nil, reportError("xCOREOSREQUESTID is required and must be specified")
	}
	if r.xCOREOSACCESS == nil {
		return localVarReturnValue, nil, reportError("xCOREOSACCESS is required and must be specified")
	}
	if r.xCOREOSTID == nil {
		return localVarReturnValue, nil, reportError("xCOREOSTID is required and must be specified")
	}
	if r.resourceCreateRequest == nil {
		return localVarReturnValue, nil, reportError("resourceCreateRequest is required and must be specified")
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
	localVarHeaderParams["X-COREOS-ACCESS"] = parameterToString(*r.xCOREOSACCESS, "")
	if r.xCOREOSAUTH != nil {
		localVarHeaderParams["X-COREOS-AUTH"] = parameterToString(*r.xCOREOSAUTH, "")
	}
	if r.xCOREOSUSERINFO != nil {
		localVarHeaderParams["X-COREOS-USERINFO"] = parameterToString(*r.xCOREOSUSERINFO, "")
	}
	localVarHeaderParams["X-COREOS-TID"] = parameterToString(*r.xCOREOSTID, "")
	// body params
	localVarPostBody = r.resourceCreateRequest
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
		if localVarHTTPResponse.StatusCode == 403 {
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

type ApiGetAllResourcesRequest struct {
	ctx context.Context
	ApiService *ResourceManagementApiService
	appId string
	xCOREOSREQUESTID *string
	xCOREOSACCESS *string
	xCOREOSTID *string
	xCOREOSAUTH *string
	xCOREOSUSERINFO *string
}

// Unique request ID
func (r ApiGetAllResourcesRequest) XCOREOSREQUESTID(xCOREOSREQUESTID string) ApiGetAllResourcesRequest {
	r.xCOREOSREQUESTID = &xCOREOSREQUESTID
	return r
}

// Core-os access token
func (r ApiGetAllResourcesRequest) XCOREOSACCESS(xCOREOSACCESS string) ApiGetAllResourcesRequest {
	r.xCOREOSACCESS = &xCOREOSACCESS
	return r
}

// Tenant ID
func (r ApiGetAllResourcesRequest) XCOREOSTID(xCOREOSTID string) ApiGetAllResourcesRequest {
	r.xCOREOSTID = &xCOREOSTID
	return r
}

// Core-os authentication token.
func (r ApiGetAllResourcesRequest) XCOREOSAUTH(xCOREOSAUTH string) ApiGetAllResourcesRequest {
	r.xCOREOSAUTH = &xCOREOSAUTH
	return r
}

// User info
func (r ApiGetAllResourcesRequest) XCOREOSUSERINFO(xCOREOSUSERINFO string) ApiGetAllResourcesRequest {
	r.xCOREOSUSERINFO = &xCOREOSUSERINFO
	return r
}

func (r ApiGetAllResourcesRequest) Execute() (*SuccessResourcesResponse, *http.Response, error) {
	return r.ApiService.GetAllResourcesExecute(r)
}

/*
GetAllResources Get all resource's information

API is used to fetch information about all resources

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appId
 @return ApiGetAllResourcesRequest
*/
func (a *ResourceManagementApiService) GetAllResources(ctx context.Context, appId string) ApiGetAllResourcesRequest {
	return ApiGetAllResourcesRequest{
		ApiService: a,
		ctx: ctx,
		appId: appId,
	}
}

// Execute executes the request
//  @return SuccessResourcesResponse
func (a *ResourceManagementApiService) GetAllResourcesExecute(r ApiGetAllResourcesRequest) (*SuccessResourcesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SuccessResourcesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourceManagementApiService.GetAllResources")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/aaa/apps/{appId}/resources"
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterToString(r.appId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xCOREOSREQUESTID == nil {
		return localVarReturnValue, nil, reportError("xCOREOSREQUESTID is required and must be specified")
	}
	if r.xCOREOSACCESS == nil {
		return localVarReturnValue, nil, reportError("xCOREOSACCESS is required and must be specified")
	}
	if r.xCOREOSTID == nil {
		return localVarReturnValue, nil, reportError("xCOREOSTID is required and must be specified")
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
	localVarHeaderParams["X-COREOS-ACCESS"] = parameterToString(*r.xCOREOSACCESS, "")
	if r.xCOREOSAUTH != nil {
		localVarHeaderParams["X-COREOS-AUTH"] = parameterToString(*r.xCOREOSAUTH, "")
	}
	if r.xCOREOSUSERINFO != nil {
		localVarHeaderParams["X-COREOS-USERINFO"] = parameterToString(*r.xCOREOSUSERINFO, "")
	}
	localVarHeaderParams["X-COREOS-TID"] = parameterToString(*r.xCOREOSTID, "")
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
		if localVarHTTPResponse.StatusCode == 403 {
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

type ApiGetResourceByIdRequest struct {
	ctx context.Context
	ApiService *ResourceManagementApiService
	appId string
	resourceId string
	xCOREOSREQUESTID *string
	xCOREOSACCESS *string
	xCOREOSTID *string
	xCOREOSAUTH *string
	xCOREOSUSERINFO *string
}

// Unique request ID
func (r ApiGetResourceByIdRequest) XCOREOSREQUESTID(xCOREOSREQUESTID string) ApiGetResourceByIdRequest {
	r.xCOREOSREQUESTID = &xCOREOSREQUESTID
	return r
}

// Core-os access token
func (r ApiGetResourceByIdRequest) XCOREOSACCESS(xCOREOSACCESS string) ApiGetResourceByIdRequest {
	r.xCOREOSACCESS = &xCOREOSACCESS
	return r
}

// Tenant ID
func (r ApiGetResourceByIdRequest) XCOREOSTID(xCOREOSTID string) ApiGetResourceByIdRequest {
	r.xCOREOSTID = &xCOREOSTID
	return r
}

// Core-os authentication token.
func (r ApiGetResourceByIdRequest) XCOREOSAUTH(xCOREOSAUTH string) ApiGetResourceByIdRequest {
	r.xCOREOSAUTH = &xCOREOSAUTH
	return r
}

// User info
func (r ApiGetResourceByIdRequest) XCOREOSUSERINFO(xCOREOSUSERINFO string) ApiGetResourceByIdRequest {
	r.xCOREOSUSERINFO = &xCOREOSUSERINFO
	return r
}

func (r ApiGetResourceByIdRequest) Execute() (*SuccessResourceResponse, *http.Response, error) {
	return r.ApiService.GetResourceByIdExecute(r)
}

/*
GetResourceById Get resource information

API is used to fetch information about a resource

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appId
 @param resourceId
 @return ApiGetResourceByIdRequest
*/
func (a *ResourceManagementApiService) GetResourceById(ctx context.Context, appId string, resourceId string) ApiGetResourceByIdRequest {
	return ApiGetResourceByIdRequest{
		ApiService: a,
		ctx: ctx,
		appId: appId,
		resourceId: resourceId,
	}
}

// Execute executes the request
//  @return SuccessResourceResponse
func (a *ResourceManagementApiService) GetResourceByIdExecute(r ApiGetResourceByIdRequest) (*SuccessResourceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SuccessResourceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourceManagementApiService.GetResourceById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/aaa/apps/{appId}/resources/{resourceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterToString(r.appId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resourceId"+"}", url.PathEscape(parameterToString(r.resourceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xCOREOSREQUESTID == nil {
		return localVarReturnValue, nil, reportError("xCOREOSREQUESTID is required and must be specified")
	}
	if r.xCOREOSACCESS == nil {
		return localVarReturnValue, nil, reportError("xCOREOSACCESS is required and must be specified")
	}
	if r.xCOREOSTID == nil {
		return localVarReturnValue, nil, reportError("xCOREOSTID is required and must be specified")
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
	localVarHeaderParams["X-COREOS-ACCESS"] = parameterToString(*r.xCOREOSACCESS, "")
	if r.xCOREOSAUTH != nil {
		localVarHeaderParams["X-COREOS-AUTH"] = parameterToString(*r.xCOREOSAUTH, "")
	}
	if r.xCOREOSUSERINFO != nil {
		localVarHeaderParams["X-COREOS-USERINFO"] = parameterToString(*r.xCOREOSUSERINFO, "")
	}
	localVarHeaderParams["X-COREOS-TID"] = parameterToString(*r.xCOREOSTID, "")
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
		if localVarHTTPResponse.StatusCode == 403 {
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

type ApiUpdateResourceRequest struct {
	ctx context.Context
	ApiService *ResourceManagementApiService
	appId string
	resourceId string
	xCOREOSREQUESTID *string
	xCOREOSACCESS *string
	xCOREOSTID *string
	resourceUpdateRequest *ResourceUpdateRequest
	xCOREOSAUTH *string
	xCOREOSUSERINFO *string
}

// Unique request ID
func (r ApiUpdateResourceRequest) XCOREOSREQUESTID(xCOREOSREQUESTID string) ApiUpdateResourceRequest {
	r.xCOREOSREQUESTID = &xCOREOSREQUESTID
	return r
}

// Core-os access token
func (r ApiUpdateResourceRequest) XCOREOSACCESS(xCOREOSACCESS string) ApiUpdateResourceRequest {
	r.xCOREOSACCESS = &xCOREOSACCESS
	return r
}

// Tenant ID
func (r ApiUpdateResourceRequest) XCOREOSTID(xCOREOSTID string) ApiUpdateResourceRequest {
	r.xCOREOSTID = &xCOREOSTID
	return r
}

// 
func (r ApiUpdateResourceRequest) ResourceUpdateRequest(resourceUpdateRequest ResourceUpdateRequest) ApiUpdateResourceRequest {
	r.resourceUpdateRequest = &resourceUpdateRequest
	return r
}

// Core-os authentication token.
func (r ApiUpdateResourceRequest) XCOREOSAUTH(xCOREOSAUTH string) ApiUpdateResourceRequest {
	r.xCOREOSAUTH = &xCOREOSAUTH
	return r
}

// User info
func (r ApiUpdateResourceRequest) XCOREOSUSERINFO(xCOREOSUSERINFO string) ApiUpdateResourceRequest {
	r.xCOREOSUSERINFO = &xCOREOSUSERINFO
	return r
}

func (r ApiUpdateResourceRequest) Execute() (*SuccessResponse, *http.Response, error) {
	return r.ApiService.UpdateResourceExecute(r)
}

/*
UpdateResource Update resource information

API is used to update a resource

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appId
 @param resourceId
 @return ApiUpdateResourceRequest
*/
func (a *ResourceManagementApiService) UpdateResource(ctx context.Context, appId string, resourceId string) ApiUpdateResourceRequest {
	return ApiUpdateResourceRequest{
		ApiService: a,
		ctx: ctx,
		appId: appId,
		resourceId: resourceId,
	}
}

// Execute executes the request
//  @return SuccessResponse
func (a *ResourceManagementApiService) UpdateResourceExecute(r ApiUpdateResourceRequest) (*SuccessResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SuccessResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourceManagementApiService.UpdateResource")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/aaa/apps/{appId}/resources/{resourceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterToString(r.appId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resourceId"+"}", url.PathEscape(parameterToString(r.resourceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xCOREOSREQUESTID == nil {
		return localVarReturnValue, nil, reportError("xCOREOSREQUESTID is required and must be specified")
	}
	if r.xCOREOSACCESS == nil {
		return localVarReturnValue, nil, reportError("xCOREOSACCESS is required and must be specified")
	}
	if r.xCOREOSTID == nil {
		return localVarReturnValue, nil, reportError("xCOREOSTID is required and must be specified")
	}
	if r.resourceUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("resourceUpdateRequest is required and must be specified")
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
	localVarHeaderParams["X-COREOS-ACCESS"] = parameterToString(*r.xCOREOSACCESS, "")
	if r.xCOREOSAUTH != nil {
		localVarHeaderParams["X-COREOS-AUTH"] = parameterToString(*r.xCOREOSAUTH, "")
	}
	if r.xCOREOSUSERINFO != nil {
		localVarHeaderParams["X-COREOS-USERINFO"] = parameterToString(*r.xCOREOSUSERINFO, "")
	}
	localVarHeaderParams["X-COREOS-TID"] = parameterToString(*r.xCOREOSTID, "")
	// body params
	localVarPostBody = r.resourceUpdateRequest
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
		if localVarHTTPResponse.StatusCode == 403 {
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