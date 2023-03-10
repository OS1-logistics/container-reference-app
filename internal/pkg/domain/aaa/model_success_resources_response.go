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

// SuccessResourcesResponse struct for SuccessResourcesResponse
type SuccessResourcesResponse struct {
	Error *Error `json:"error,omitempty"`
	Data []ResourceResponse `json:"data,omitempty"`
	Request *Request `json:"request,omitempty"`
}

// NewSuccessResourcesResponse instantiates a new SuccessResourcesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuccessResourcesResponse() *SuccessResourcesResponse {
	this := SuccessResourcesResponse{}
	return &this
}

// NewSuccessResourcesResponseWithDefaults instantiates a new SuccessResourcesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuccessResourcesResponseWithDefaults() *SuccessResourcesResponse {
	this := SuccessResourcesResponse{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SuccessResourcesResponse) GetError() Error {
	if o == nil || isNil(o.Error) {
		var ret Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessResourcesResponse) GetErrorOk() (*Error, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SuccessResourcesResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given Error and assigns it to the Error field.
func (o *SuccessResourcesResponse) SetError(v Error) {
	o.Error = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SuccessResourcesResponse) GetData() []ResourceResponse {
	if o == nil || isNil(o.Data) {
		var ret []ResourceResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessResourcesResponse) GetDataOk() ([]ResourceResponse, bool) {
	if o == nil || isNil(o.Data) {
    return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SuccessResourcesResponse) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []ResourceResponse and assigns it to the Data field.
func (o *SuccessResourcesResponse) SetData(v []ResourceResponse) {
	o.Data = v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *SuccessResourcesResponse) GetRequest() Request {
	if o == nil || isNil(o.Request) {
		var ret Request
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessResourcesResponse) GetRequestOk() (*Request, bool) {
	if o == nil || isNil(o.Request) {
    return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *SuccessResourcesResponse) HasRequest() bool {
	if o != nil && !isNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given Request and assigns it to the Request field.
func (o *SuccessResourcesResponse) SetRequest(v Request) {
	o.Request = &v
}

func (o SuccessResourcesResponse) MarshalJSON() ([]byte, error) {
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

type NullableSuccessResourcesResponse struct {
	value *SuccessResourcesResponse
	isSet bool
}

func (v NullableSuccessResourcesResponse) Get() *SuccessResourcesResponse {
	return v.value
}

func (v *NullableSuccessResourcesResponse) Set(val *SuccessResourcesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSuccessResourcesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSuccessResourcesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuccessResourcesResponse(val *SuccessResourcesResponse) *NullableSuccessResourcesResponse {
	return &NullableSuccessResourcesResponse{value: val, isSet: true}
}

func (v NullableSuccessResourcesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuccessResourcesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


