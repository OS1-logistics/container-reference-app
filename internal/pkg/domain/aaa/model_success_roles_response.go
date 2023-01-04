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

// SuccessRolesResponse struct for SuccessRolesResponse
type SuccessRolesResponse struct {
	Error *Error `json:"error,omitempty"`
	Data []RoleResponse `json:"data,omitempty"`
	Request *Request `json:"request,omitempty"`
}

// NewSuccessRolesResponse instantiates a new SuccessRolesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuccessRolesResponse() *SuccessRolesResponse {
	this := SuccessRolesResponse{}
	return &this
}

// NewSuccessRolesResponseWithDefaults instantiates a new SuccessRolesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuccessRolesResponseWithDefaults() *SuccessRolesResponse {
	this := SuccessRolesResponse{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SuccessRolesResponse) GetError() Error {
	if o == nil || isNil(o.Error) {
		var ret Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessRolesResponse) GetErrorOk() (*Error, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SuccessRolesResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given Error and assigns it to the Error field.
func (o *SuccessRolesResponse) SetError(v Error) {
	o.Error = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SuccessRolesResponse) GetData() []RoleResponse {
	if o == nil || isNil(o.Data) {
		var ret []RoleResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessRolesResponse) GetDataOk() ([]RoleResponse, bool) {
	if o == nil || isNil(o.Data) {
    return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SuccessRolesResponse) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []RoleResponse and assigns it to the Data field.
func (o *SuccessRolesResponse) SetData(v []RoleResponse) {
	o.Data = v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *SuccessRolesResponse) GetRequest() Request {
	if o == nil || isNil(o.Request) {
		var ret Request
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessRolesResponse) GetRequestOk() (*Request, bool) {
	if o == nil || isNil(o.Request) {
    return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *SuccessRolesResponse) HasRequest() bool {
	if o != nil && !isNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given Request and assigns it to the Request field.
func (o *SuccessRolesResponse) SetRequest(v Request) {
	o.Request = &v
}

func (o SuccessRolesResponse) MarshalJSON() ([]byte, error) {
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

type NullableSuccessRolesResponse struct {
	value *SuccessRolesResponse
	isSet bool
}

func (v NullableSuccessRolesResponse) Get() *SuccessRolesResponse {
	return v.value
}

func (v *NullableSuccessRolesResponse) Set(val *SuccessRolesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSuccessRolesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSuccessRolesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuccessRolesResponse(val *SuccessRolesResponse) *NullableSuccessRolesResponse {
	return &NullableSuccessRolesResponse{value: val, isSet: true}
}

func (v NullableSuccessRolesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuccessRolesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


