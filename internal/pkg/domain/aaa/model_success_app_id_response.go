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

// SuccessAppIdResponse struct for SuccessAppIdResponse
type SuccessAppIdResponse struct {
	Error *Error `json:"error,omitempty"`
	Data *AppIdResponse `json:"data,omitempty"`
	Request *Request `json:"request,omitempty"`
}

// NewSuccessAppIdResponse instantiates a new SuccessAppIdResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuccessAppIdResponse() *SuccessAppIdResponse {
	this := SuccessAppIdResponse{}
	return &this
}

// NewSuccessAppIdResponseWithDefaults instantiates a new SuccessAppIdResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuccessAppIdResponseWithDefaults() *SuccessAppIdResponse {
	this := SuccessAppIdResponse{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SuccessAppIdResponse) GetError() Error {
	if o == nil || o.Error == nil {
		var ret Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessAppIdResponse) GetErrorOk() (*Error, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SuccessAppIdResponse) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given Error and assigns it to the Error field.
func (o *SuccessAppIdResponse) SetError(v Error) {
	o.Error = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SuccessAppIdResponse) GetData() AppIdResponse {
	if o == nil || o.Data == nil {
		var ret AppIdResponse
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessAppIdResponse) GetDataOk() (*AppIdResponse, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SuccessAppIdResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given AppIdResponse and assigns it to the Data field.
func (o *SuccessAppIdResponse) SetData(v AppIdResponse) {
	o.Data = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *SuccessAppIdResponse) GetRequest() Request {
	if o == nil || o.Request == nil {
		var ret Request
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessAppIdResponse) GetRequestOk() (*Request, bool) {
	if o == nil || o.Request == nil {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *SuccessAppIdResponse) HasRequest() bool {
	if o != nil && o.Request != nil {
		return true
	}

	return false
}

// SetRequest gets a reference to the given Request and assigns it to the Request field.
func (o *SuccessAppIdResponse) SetRequest(v Request) {
	o.Request = &v
}

func (o SuccessAppIdResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Request != nil {
		toSerialize["request"] = o.Request
	}
	return json.Marshal(toSerialize)
}

type NullableSuccessAppIdResponse struct {
	value *SuccessAppIdResponse
	isSet bool
}

func (v NullableSuccessAppIdResponse) Get() *SuccessAppIdResponse {
	return v.value
}

func (v *NullableSuccessAppIdResponse) Set(val *SuccessAppIdResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSuccessAppIdResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSuccessAppIdResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuccessAppIdResponse(val *SuccessAppIdResponse) *NullableSuccessAppIdResponse {
	return &NullableSuccessAppIdResponse{value: val, isSet: true}
}

func (v NullableSuccessAppIdResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuccessAppIdResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


