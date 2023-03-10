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

// SuccessPermissionResponse struct for SuccessPermissionResponse
type SuccessPermissionResponse struct {
	Error *Error `json:"error,omitempty"`
	Data *PermissionResponse `json:"data,omitempty"`
	Request *Request `json:"request,omitempty"`
}

// NewSuccessPermissionResponse instantiates a new SuccessPermissionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuccessPermissionResponse() *SuccessPermissionResponse {
	this := SuccessPermissionResponse{}
	return &this
}

// NewSuccessPermissionResponseWithDefaults instantiates a new SuccessPermissionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuccessPermissionResponseWithDefaults() *SuccessPermissionResponse {
	this := SuccessPermissionResponse{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SuccessPermissionResponse) GetError() Error {
	if o == nil || isNil(o.Error) {
		var ret Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessPermissionResponse) GetErrorOk() (*Error, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SuccessPermissionResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given Error and assigns it to the Error field.
func (o *SuccessPermissionResponse) SetError(v Error) {
	o.Error = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SuccessPermissionResponse) GetData() PermissionResponse {
	if o == nil || isNil(o.Data) {
		var ret PermissionResponse
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessPermissionResponse) GetDataOk() (*PermissionResponse, bool) {
	if o == nil || isNil(o.Data) {
    return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SuccessPermissionResponse) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PermissionResponse and assigns it to the Data field.
func (o *SuccessPermissionResponse) SetData(v PermissionResponse) {
	o.Data = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *SuccessPermissionResponse) GetRequest() Request {
	if o == nil || isNil(o.Request) {
		var ret Request
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessPermissionResponse) GetRequestOk() (*Request, bool) {
	if o == nil || isNil(o.Request) {
    return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *SuccessPermissionResponse) HasRequest() bool {
	if o != nil && !isNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given Request and assigns it to the Request field.
func (o *SuccessPermissionResponse) SetRequest(v Request) {
	o.Request = &v
}

func (o SuccessPermissionResponse) MarshalJSON() ([]byte, error) {
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

type NullableSuccessPermissionResponse struct {
	value *SuccessPermissionResponse
	isSet bool
}

func (v NullableSuccessPermissionResponse) Get() *SuccessPermissionResponse {
	return v.value
}

func (v *NullableSuccessPermissionResponse) Set(val *SuccessPermissionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSuccessPermissionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSuccessPermissionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuccessPermissionResponse(val *SuccessPermissionResponse) *NullableSuccessPermissionResponse {
	return &NullableSuccessPermissionResponse{value: val, isSet: true}
}

func (v NullableSuccessPermissionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuccessPermissionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


