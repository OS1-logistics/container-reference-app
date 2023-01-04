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


// RoleManagementApiService RoleManagementApi service
type RoleManagementApiService service

type ApiCreateAppRoleRequest struct {
	ctx context.Context
	ApiService *RoleManagementApiService
	appId string
	xCOREOSREQUESTID *string
	xCOREOSACCESS *string
	xCOREOSTID *string
	roleCreateRequest *RoleCreateRequest
	xCOREOSAUTH *string
	xCOREOSUSERINFO *string
}

// Unique request ID
func (r ApiCreateAppRoleRequest) XCOREOSREQUESTID(xCOREOSREQUESTID string) ApiCreateAppRoleRequest {
	r.xCOREOSREQUESTID = &xCOREOSREQUESTID
	return r
}

// Core-os access token
func (r ApiCreateAppRoleRequest) XCOREOSACCESS(xCOREOSACCESS string) ApiCreateAppRoleRequest {
	r.xCOREOSACCESS = &xCOREOSACCESS
	return r
}

// Tenant ID
func (r ApiCreateAppRoleRequest) XCOREOSTID(xCOREOSTID string) ApiCreateAppRoleRequest {
	r.xCOREOSTID = &xCOREOSTID
	return r
}

// 
func (r ApiCreateAppRoleRequest) RoleCreateRequest(roleCreateRequest RoleCreateRequest) ApiCreateAppRoleRequest {
	r.roleCreateRequest = &roleCreateRequest
	return r
}

// Core-os authentication token.
func (r ApiCreateAppRoleRequest) XCOREOSAUTH(xCOREOSAUTH string) ApiCreateAppRoleRequest {
	r.xCOREOSAUTH = &xCOREOSAUTH
	return r
}

// User info
func (r ApiCreateAppRoleRequest) XCOREOSUSERINFO(xCOREOSUSERINFO string) ApiCreateAppRoleRequest {
	r.xCOREOSUSERINFO = &xCOREOSUSERINFO
	return r
}

func (r ApiCreateAppRoleRequest) Execute() (*SuccessResponse, *http.Response, error) {
	return r.ApiService.CreateAppRoleExecute(r)
}

/*
CreateAppRole Create App Role

This endpoint creates an App Role. Role is a generic representation of a set of permissions or resources of a specific app or service.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appId
 @return ApiCreateAppRoleRequest
*/
func (a *RoleManagementApiService) CreateAppRole(ctx context.Context, appId string) ApiCreateAppRoleRequest {
	return ApiCreateAppRoleRequest{
		ApiService: a,
		ctx: ctx,
		appId: appId,
	}
}

// Execute executes the request
//  @return SuccessResponse
func (a *RoleManagementApiService) CreateAppRoleExecute(r ApiCreateAppRoleRequest) (*SuccessResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SuccessResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleManagementApiService.CreateAppRole")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/aaa/apps/{appId}/roles"
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
	if r.roleCreateRequest == nil {
		return localVarReturnValue, nil, reportError("roleCreateRequest is required and must be specified")
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
	localVarPostBody = r.roleCreateRequest
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

type ApiGetAllRolesRequest struct {
	ctx context.Context
	ApiService *RoleManagementApiService
	appId string
	xCOREOSREQUESTID *string
	xCOREOSACCESS *string
	xCOREOSTID *string
	xCOREOSAUTH *string
	xCOREOSUSERINFO *string
}

// Unique request ID
func (r ApiGetAllRolesRequest) XCOREOSREQUESTID(xCOREOSREQUESTID string) ApiGetAllRolesRequest {
	r.xCOREOSREQUESTID = &xCOREOSREQUESTID
	return r
}

// Core-os access token
func (r ApiGetAllRolesRequest) XCOREOSACCESS(xCOREOSACCESS string) ApiGetAllRolesRequest {
	r.xCOREOSACCESS = &xCOREOSACCESS
	return r
}

// Tenant ID
func (r ApiGetAllRolesRequest) XCOREOSTID(xCOREOSTID string) ApiGetAllRolesRequest {
	r.xCOREOSTID = &xCOREOSTID
	return r
}

// Core-os authentication token.
func (r ApiGetAllRolesRequest) XCOREOSAUTH(xCOREOSAUTH string) ApiGetAllRolesRequest {
	r.xCOREOSAUTH = &xCOREOSAUTH
	return r
}

// User info
func (r ApiGetAllRolesRequest) XCOREOSUSERINFO(xCOREOSUSERINFO string) ApiGetAllRolesRequest {
	r.xCOREOSUSERINFO = &xCOREOSUSERINFO
	return r
}

func (r ApiGetAllRolesRequest) Execute() (*SuccessRolesResponse, *http.Response, error) {
	return r.ApiService.GetAllRolesExecute(r)
}

/*
GetAllRoles Get all Roles

This endpoint retrieves details of all created Roles.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appId
 @return ApiGetAllRolesRequest
*/
func (a *RoleManagementApiService) GetAllRoles(ctx context.Context, appId string) ApiGetAllRolesRequest {
	return ApiGetAllRolesRequest{
		ApiService: a,
		ctx: ctx,
		appId: appId,
	}
}

// Execute executes the request
//  @return SuccessRolesResponse
func (a *RoleManagementApiService) GetAllRolesExecute(r ApiGetAllRolesRequest) (*SuccessRolesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SuccessRolesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleManagementApiService.GetAllRoles")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/aaa/apps/{appId}/roles"
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

type ApiGetRoleDetailsRequest struct {
	ctx context.Context
	ApiService *RoleManagementApiService
	appId string
	roleId string
	xCOREOSTID *string
	xCOREOSREQUESTID *string
	xCOREOSACCESS *string
	xCOREOSAUTH *string
	xCOREOSUSERINFO *string
}

// Tenant ID
func (r ApiGetRoleDetailsRequest) XCOREOSTID(xCOREOSTID string) ApiGetRoleDetailsRequest {
	r.xCOREOSTID = &xCOREOSTID
	return r
}

// Unique request ID
func (r ApiGetRoleDetailsRequest) XCOREOSREQUESTID(xCOREOSREQUESTID string) ApiGetRoleDetailsRequest {
	r.xCOREOSREQUESTID = &xCOREOSREQUESTID
	return r
}

// Core-os access token
func (r ApiGetRoleDetailsRequest) XCOREOSACCESS(xCOREOSACCESS string) ApiGetRoleDetailsRequest {
	r.xCOREOSACCESS = &xCOREOSACCESS
	return r
}

// Core-os authentication token.
func (r ApiGetRoleDetailsRequest) XCOREOSAUTH(xCOREOSAUTH string) ApiGetRoleDetailsRequest {
	r.xCOREOSAUTH = &xCOREOSAUTH
	return r
}

// User info
func (r ApiGetRoleDetailsRequest) XCOREOSUSERINFO(xCOREOSUSERINFO string) ApiGetRoleDetailsRequest {
	r.xCOREOSUSERINFO = &xCOREOSUSERINFO
	return r
}

func (r ApiGetRoleDetailsRequest) Execute() (*SuccessRoleResponse, *http.Response, error) {
	return r.ApiService.GetRoleDetailsExecute(r)
}

/*
GetRoleDetails Get specific role information.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appId
 @param roleId The App for which the permissions in the role belong.
 @return ApiGetRoleDetailsRequest
*/
func (a *RoleManagementApiService) GetRoleDetails(ctx context.Context, appId string, roleId string) ApiGetRoleDetailsRequest {
	return ApiGetRoleDetailsRequest{
		ApiService: a,
		ctx: ctx,
		appId: appId,
		roleId: roleId,
	}
}

// Execute executes the request
//  @return SuccessRoleResponse
func (a *RoleManagementApiService) GetRoleDetailsExecute(r ApiGetRoleDetailsRequest) (*SuccessRoleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SuccessRoleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleManagementApiService.GetRoleDetails")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/aaa/apps/{appId}/roles/{roleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterToString(r.appId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleId"+"}", url.PathEscape(parameterToString(r.roleId, "")), -1)

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

type ApiUpdateGlobalRoleRequest struct {
	ctx context.Context
	ApiService *RoleManagementApiService
	appId string
	roleId string
	xCOREOSREQUESTID *string
	xCOREOSTID *string
	xCOREOSACCESS *string
	roleUpdateRequest *RoleUpdateRequest
	xCOREOSAUTH *string
	xCOREOSUSERINFO *string
}

// Unique request ID
func (r ApiUpdateGlobalRoleRequest) XCOREOSREQUESTID(xCOREOSREQUESTID string) ApiUpdateGlobalRoleRequest {
	r.xCOREOSREQUESTID = &xCOREOSREQUESTID
	return r
}

// Tenant ID
func (r ApiUpdateGlobalRoleRequest) XCOREOSTID(xCOREOSTID string) ApiUpdateGlobalRoleRequest {
	r.xCOREOSTID = &xCOREOSTID
	return r
}

// Core-os access token
func (r ApiUpdateGlobalRoleRequest) XCOREOSACCESS(xCOREOSACCESS string) ApiUpdateGlobalRoleRequest {
	r.xCOREOSACCESS = &xCOREOSACCESS
	return r
}

// 
func (r ApiUpdateGlobalRoleRequest) RoleUpdateRequest(roleUpdateRequest RoleUpdateRequest) ApiUpdateGlobalRoleRequest {
	r.roleUpdateRequest = &roleUpdateRequest
	return r
}

// Core-os authentication token.
func (r ApiUpdateGlobalRoleRequest) XCOREOSAUTH(xCOREOSAUTH string) ApiUpdateGlobalRoleRequest {
	r.xCOREOSAUTH = &xCOREOSAUTH
	return r
}

// User info
func (r ApiUpdateGlobalRoleRequest) XCOREOSUSERINFO(xCOREOSUSERINFO string) ApiUpdateGlobalRoleRequest {
	r.xCOREOSUSERINFO = &xCOREOSUSERINFO
	return r
}

func (r ApiUpdateGlobalRoleRequest) Execute() (*SuccessResponse, *http.Response, error) {
	return r.ApiService.UpdateGlobalRoleExecute(r)
}

/*
UpdateGlobalRole Update a Role

This endpoint updates an existing Role.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appId
 @param roleId The App for which the permissions in the role belong.
 @return ApiUpdateGlobalRoleRequest
*/
func (a *RoleManagementApiService) UpdateGlobalRole(ctx context.Context, appId string, roleId string) ApiUpdateGlobalRoleRequest {
	return ApiUpdateGlobalRoleRequest{
		ApiService: a,
		ctx: ctx,
		appId: appId,
		roleId: roleId,
	}
}

// Execute executes the request
//  @return SuccessResponse
func (a *RoleManagementApiService) UpdateGlobalRoleExecute(r ApiUpdateGlobalRoleRequest) (*SuccessResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SuccessResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleManagementApiService.UpdateGlobalRole")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/aaa/apps/{appId}/roles/{roleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterToString(r.appId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleId"+"}", url.PathEscape(parameterToString(r.roleId, "")), -1)

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
	if r.roleUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("roleUpdateRequest is required and must be specified")
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
	localVarPostBody = r.roleUpdateRequest
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
