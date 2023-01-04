/*
Authentication And Authorization (AAA) Service

This swagger documentation provides all AAA API details. AAA service provides authentication and authorization capabilities for users.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aaadomain

import (
	"encoding/json"
)

// SuccessPermissionsResponse struct for SuccessPermissionsResponse
type SuccessPermissionsResponse struct {
	Error *Error `json:"error,omitempty"`
	Data []PermissionResponse `json:"data,omitempty"`
	Request *Request `json:"request,omitempty"`
}

// NewSuccessPermissionsResponse instantiates a new SuccessPermissionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuccessPermissionsResponse() *SuccessPermissionsResponse {
	this := SuccessPermissionsResponse{}
	return &this
}

// NewSuccessPermissionsResponseWithDefaults instantiates a new SuccessPermissionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuccessPermissionsResponseWithDefaults() *SuccessPermissionsResponse {
	this := SuccessPermissionsResponse{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SuccessPermissionsResponse) GetError() Error {
	if o == nil || isNil(o.Error) {
		var ret Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessPermissionsResponse) GetErrorOk() (*Error, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SuccessPermissionsResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given Error and assigns it to the Error field.
func (o *SuccessPermissionsResponse) SetError(v Error) {
	o.Error = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SuccessPermissionsResponse) GetData() []PermissionResponse {
	if o == nil || isNil(o.Data) {
		var ret []PermissionResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessPermissionsResponse) GetDataOk() ([]PermissionResponse, bool) {
	if o == nil || isNil(o.Data) {
    return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SuccessPermissionsResponse) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []PermissionResponse and assigns it to the Data field.
func (o *SuccessPermissionsResponse) SetData(v []PermissionResponse) {
	o.Data = v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *SuccessPermissionsResponse) GetRequest() Request {
	if o == nil || isNil(o.Request) {
		var ret Request
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessPermissionsResponse) GetRequestOk() (*Request, bool) {
	if o == nil || isNil(o.Request) {
    return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *SuccessPermissionsResponse) HasRequest() bool {
	if o != nil && !isNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given Request and assigns it to the Request field.
func (o *SuccessPermissionsResponse) SetRequest(v Request) {
	o.Request = &v
}

func (o SuccessPermissionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !isNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	return json.Marshal(toSerialize)
}

type NullableSuccessPermissionsResponse struct {
	value *SuccessPermissionsResponse
	isSet bool
}

func (v NullableSuccessPermissionsResponse) Get() *SuccessPermissionsResponse {
	return v.value
}

func (v *NullableSuccessPermissionsResponse) Set(val *SuccessPermissionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSuccessPermissionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSuccessPermissionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuccessPermissionsResponse(val *SuccessPermissionsResponse) *NullableSuccessPermissionsResponse {
	return &NullableSuccessPermissionsResponse{value: val, isSet: true}
}

func (v NullableSuccessPermissionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuccessPermissionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

