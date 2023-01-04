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


// PermissionManagementApiService PermissionManagementApi service
type PermissionManagementApiService service

type ApiCreatePermissionRequest struct {
	ctx context.Context
	ApiService *PermissionManagementApiService
	appId string
	xCOREOSREQUESTID *string
	xCOREOSTID *string
	xCOREOSACCESS *string
	permissionCreateRequest *PermissionCreateRequest
	xCOREOSAUTH *string
	xCOREOSUSERINFO *string
}

// Unique request ID
func (r ApiCreatePermissionRequest) XCOREOSREQUESTID(xCOREOSREQUESTID string) ApiCreatePermissionRequest {
	r.xCOREOSREQUESTID = &xCOREOSREQUESTID
	return r
}

// Tenant ID
func (r ApiCreatePermissionRequest) XCOREOSTID(xCOREOSTID string) ApiCreatePermissionRequest {
	r.xCOREOSTID = &xCOREOSTID
	return r
}

// Core-os access token
func (r ApiCreatePermissionRequest) XCOREOSACCESS(xCOREOSACCESS string) ApiCreatePermissionRequest {
	r.xCOREOSACCESS = &xCOREOSACCESS
	return r
}

// 
func (r ApiCreatePermissionRequest) PermissionCreateRequest(permissionCreateRequest PermissionCreateRequest) ApiCreatePermissionRequest {
	r.permissionCreateRequest = &permissionCreateRequest
	return r
}

// Core-os authentication token.
func (r ApiCreatePermissionRequest) XCOREOSAUTH(xCOREOSAUTH string) ApiCreatePermissionRequest {
	r.xCOREOSAUTH = &xCOREOSAUTH
	return r
}

// User info
func (r ApiCreatePermissionRequest) XCOREOSUSERINFO(xCOREOSUSERINFO string) ApiCreatePermissionRequest {
	r.xCOREOSUSERINFO = &xCOREOSUSERINFO
	return r
}

func (r ApiCreatePermissionRequest) Execute() (*SuccessResponse, *http.Response, error) {
	return r.ApiService.CreatePermissionExecute(r)
}

/*
CreatePermission Create a permission

API is used to create a permission

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appId
 @return ApiCreatePermissionRequest
*/
func (a *PermissionManagementApiService) CreatePermission(ctx context.Context, appId string) ApiCreatePermissionRequest {
	return ApiCreatePermissionRequest{
		ApiService: a,
		ctx: ctx,
		appId: appId,
	}
}

// Execute executes the request
//  @return SuccessResponse
func (a *PermissionManagementApiService) CreatePermissionExecute(r ApiCreatePermissionRequest) (*SuccessResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SuccessResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionManagementApiService.CreatePermission")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/aaa/apps/{appId}/permissions"
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterToString(r.appId, "")), -1)

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
	if r.permissionCreateRequest == nil {
		return localVarReturnValue, nil, reportError("permissionCreateRequest is required and must be specified")
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
	if r.xCOREOSAUTH != nil {
		localVarHeaderParams["X-COREOS-AUTH"] = parameterToString(*r.xCOREOSAUTH, "")
	}
	if r.xCOREOSUSERINFO != nil {
		localVarHeaderParams["X-COREOS-USERINFO"] = parameterToString(*r.xCOREOSUSERINFO, "")
	}
	// body params
	localVarPostBody = r.permissionCreateRequest
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

type ApiGetPermissionByIdRequest struct {
	ctx context.Context
	ApiService *PermissionManagementApiService
	appId string
	permissionId string
	xCOREOSREQUESTID *string
	xCOREOSTID *string
	xCOREOSACCESS *string
	xCOREOSAUTH *string
	xCOREOSUSERINFO *string
}

// Unique request ID
func (r ApiGetPermissionByIdRequest) XCOREOSREQUESTID(xCOREOSREQUESTID string) ApiGetPermissionByIdRequest {
	r.xCOREOSREQUESTID = &xCOREOSREQUESTID
	return r
}

// Tenant ID
func (r ApiGetPermissionByIdRequest) XCOREOSTID(xCOREOSTID string) ApiGetPermissionByIdRequest {
	r.xCOREOSTID = &xCOREOSTID
	return r
}

// Core-os access token
func (r ApiGetPermissionByIdRequest) XCOREOSACCESS(xCOREOSACCESS string) ApiGetPermissionByIdRequest {
	r.xCOREOSACCESS = &xCOREOSACCESS
	return r
}

// Core-os authentication token.
func (r ApiGetPermissionByIdRequest) XCOREOSAUTH(xCOREOSAUTH string) ApiGetPermissionByIdRequest {
	r.xCOREOSAUTH = &xCOREOSAUTH
	return r
}

// User info
func (r ApiGetPermissionByIdRequest) XCOREOSUSERINFO(xCOREOSUSERINFO string) ApiGetPermissionByIdRequest {
	r.xCOREOSUSERINFO = &xCOREOSUSERINFO
	return r
}

func (r ApiGetPermissionByIdRequest) Execute() (*SuccessPermissionResponse, *http.Response, error) {
	return r.ApiService.GetPermissionByIdExecute(r)
}

/*
GetPermissionById Get permission information

API is used to fetch information about a permission

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appId
 @param permissionId
 @return ApiGetPermissionByIdRequest
*/
func (a *PermissionManagementApiService) GetPermissionById(ctx context.Context, appId string, permissionId string) ApiGetPermissionByIdRequest {
	return ApiGetPermissionByIdRequest{
		ApiService: a,
		ctx: ctx,
		appId: appId,
		permissionId: permissionId,
	}
}

// Execute executes the request
//  @return SuccessPermissionResponse
func (a *PermissionManagementApiService) GetPermissionByIdExecute(r ApiGetPermissionByIdRequest) (*SuccessPermissionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SuccessPermissionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionManagementApiService.GetPermissionById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/aaa/apps/{appId}/permissions/{permissionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterToString(r.appId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"permissionId"+"}", url.PathEscape(parameterToString(r.permissionId, "")), -1)

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
	if r.xCOREOSAUTH != nil {
		localVarHeaderParams["X-COREOS-AUTH"] = parameterToString(*r.xCOREOSAUTH, "")
	}
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

type ApiGetPermissionsRequest struct {
	ctx context.Context
	ApiService *PermissionManagementApiService
	appId string
	xCOREOSTID *string
	xCOREOSREQUESTID *string
	xCOREOSACCESS *string
	xCOREOSAUTH *string
	xCOREOSUSERINFO *string
}

// Tenant ID
func (r ApiGetPermissionsRequest) XCOREOSTID(xCOREOSTID string) ApiGetPermissionsRequest {
	r.xCOREOSTID = &xCOREOSTID
	return r
}

// Unique request ID
func (r ApiGetPermissionsRequest) XCOREOSREQUESTID(xCOREOSREQUESTID string) ApiGetPermissionsRequest {
	r.xCOREOSREQUESTID = &xCOREOSREQUESTID
	return r
}

// Core-os access token
func (r ApiGetPermissionsRequest) XCOREOSACCESS(xCOREOSACCESS string) ApiGetPermissionsRequest {
	r.xCOREOSACCESS = &xCOREOSACCESS
	return r
}

// Core-os authentication token.
func (r ApiGetPermissionsRequest) XCOREOSAUTH(xCOREOSAUTH string) ApiGetPermissionsRequest {
	r.xCOREOSAUTH = &xCOREOSAUTH
	return r
}

// User info
func (r ApiGetPermissionsRequest) XCOREOSUSERINFO(xCOREOSUSERINFO string) ApiGetPermissionsRequest {
	r.xCOREOSUSERINFO = &xCOREOSUSERINFO
	return r
}

func (r ApiGetPermissionsRequest) Execute() (*SuccessPermissionsResponse, *http.Response, error) {
	return r.ApiService.GetPermissionsExecute(r)
}

/*
GetPermissions Get permission's information

API is used to fetch information about all permissions

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appId
 @return ApiGetPermissionsRequest
*/
func (a *PermissionManagementApiService) GetPermissions(ctx context.Context, appId string) ApiGetPermissionsRequest {
	return ApiGetPermissionsRequest{
		ApiService: a,
		ctx: ctx,
		appId: appId,
	}
}

// Execute executes the request
//  @return SuccessPermissionsResponse
func (a *PermissionManagementApiService) GetPermissionsExecute(r ApiGetPermissionsRequest) (*SuccessPermissionsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SuccessPermissionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionManagementApiService.GetPermissions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/aaa/apps/{appId}/permissions"
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterToString(r.appId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xCOREOSTID == nil {
		return localVarReturnValue, nil, reportError("xCOREOSTID is required and must be specified")
	}
	if r.xCOREOSREQUESTID == nil {
		return localVarReturnValue, nil, reportError("xCOREOSREQUESTID is required and must be specified")
	}
	if r.xCOREOSACCESS == nil {
		return localVarReturnValue, nil, reportError("xCOREOSACCESS is required and must be specified")
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
	localVarHeaderParams["X-COREOS-TID"] = parameterToString(*r.xCOREOSTID, "")
	localVarHeaderParams["X-COREOS-REQUEST-ID"] = parameterToString(*r.xCOREOSREQUESTID, "")
	localVarHeaderParams["X-COREOS-ACCESS"] = parameterToString(*r.xCOREOSACCESS, "")
	if r.xCOREOSAUTH != nil {
		localVarHeaderParams["X-COREOS-AUTH"] = parameterToString(*r.xCOREOSAUTH, "")
	}
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

type ApiUpdatePermissionRequest struct {
	ctx context.Context
	ApiService *PermissionManagementApiService
	appId string
	permissionId string
	xCOREOSREQUESTID *string
	xCOREOSTID *string
	xCOREOSACCESS *string
	permissionUpdateRequest *PermissionUpdateRequest
	xCOREOSAUTH *string
	xCOREOSUSERINFO *string
}

// Unique request ID
func (r ApiUpdatePermissionRequest) XCOREOSREQUESTID(xCOREOSREQUESTID string) ApiUpdatePermissionRequest {
	r.xCOREOSREQUESTID = &xCOREOSREQUESTID
	return r
}

// Tenant ID
func (r ApiUpdatePermissionRequest) XCOREOSTID(xCOREOSTID string) ApiUpdatePermissionRequest {
	r.xCOREOSTID = &xCOREOSTID
	return r
}

// Core-os access token
func (r ApiUpdatePermissionRequest) XCOREOSACCESS(xCOREOSACCESS string) ApiUpdatePermissionRequest {
	r.xCOREOSACCESS = &xCOREOSACCESS
	return r
}

// 
func (r ApiUpdatePermissionRequest) PermissionUpdateRequest(permissionUpdateRequest PermissionUpdateRequest) ApiUpdatePermissionRequest {
	r.permissionUpdateRequest = &permissionUpdateRequest
	return r
}

// Core-os authentication token.
func (r ApiUpdatePermissionRequest) XCOREOSAUTH(xCOREOSAUTH string) ApiUpdatePermissionRequest {
	r.xCOREOSAUTH = &xCOREOSAUTH
	return r
}

// User info
func (r ApiUpdatePermissionRequest) XCOREOSUSERINFO(xCOREOSUSERINFO string) ApiUpdatePermissionRequest {
	r.xCOREOSUSERINFO = &xCOREOSUSERINFO
	return r
}

func (r ApiUpdatePermissionRequest) Execute() (*SuccessResponse, *http.Response, error) {
	return r.ApiService.UpdatePermissionExecute(r)
}

/*
UpdatePermission Update permission information

API is used to update a permission

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appId
 @param permissionId
 @return ApiUpdatePermissionRequest
*/
func (a *PermissionManagementApiService) UpdatePermission(ctx context.Context, appId string, permissionId string) ApiUpdatePermissionRequest {
	return ApiUpdatePermissionRequest{
		ApiService: a,
		ctx: ctx,
		appId: appId,
		permissionId: permissionId,
	}
}

// Execute executes the request
//  @return SuccessResponse
func (a *PermissionManagementApiService) UpdatePermissionExecute(r ApiUpdatePermissionRequest) (*SuccessResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SuccessResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionManagementApiService.UpdatePermission")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/aaa/apps/{appId}/permissions/{permissionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterToString(r.appId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"permissionId"+"}", url.PathEscape(parameterToString(r.permissionId, "")), -1)

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
	if r.permissionUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("permissionUpdateRequest is required and must be specified")
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
	if r.xCOREOSAUTH != nil {
		localVarHeaderParams["X-COREOS-AUTH"] = parameterToString(*r.xCOREOSAUTH, "")
	}
	if r.xCOREOSUSERINFO != nil {
		localVarHeaderParams["X-COREOS-USERINFO"] = parameterToString(*r.xCOREOSUSERINFO, "")
	}
	// body params
	localVarPostBody = r.permissionUpdateRequest
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
