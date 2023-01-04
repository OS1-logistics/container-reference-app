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

// SuccessAuthorizationResponse struct for SuccessAuthorizationResponse
type SuccessAuthorizationResponse struct {
	Error *Error `json:"error,omitempty"`
	Data *SuccessAuthorizationResponseData `json:"data,omitempty"`
	Request *Request `json:"request,omitempty"`
}

// NewSuccessAuthorizationResponse instantiates a new SuccessAuthorizationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuccessAuthorizationResponse() *SuccessAuthorizationResponse {
	this := SuccessAuthorizationResponse{}
	return &this
}

// NewSuccessAuthorizationResponseWithDefaults instantiates a new SuccessAuthorizationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuccessAuthorizationResponseWithDefaults() *SuccessAuthorizationResponse {
	this := SuccessAuthorizationResponse{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SuccessAuthorizationResponse) GetError() Error {
	if o == nil || isNil(o.Error) {
		var ret Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessAuthorizationResponse) GetErrorOk() (*Error, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SuccessAuthorizationResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given Error and assigns it to the Error field.
func (o *SuccessAuthorizationResponse) SetError(v Error) {
	o.Error = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SuccessAuthorizationResponse) GetData() SuccessAuthorizationResponseData {
	if o == nil || isNil(o.Data) {
		var ret SuccessAuthorizationResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessAuthorizationResponse) GetDataOk() (*SuccessAuthorizationResponseData, bool) {
	if o == nil || isNil(o.Data) {
    return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SuccessAuthorizationResponse) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given SuccessAuthorizationResponseData and assigns it to the Data field.
func (o *SuccessAuthorizationResponse) SetData(v SuccessAuthorizationResponseData) {
	o.Data = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *SuccessAuthorizationResponse) GetRequest() Request {
	if o == nil || isNil(o.Request) {
		var ret Request
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessAuthorizationResponse) GetRequestOk() (*Request, bool) {
	if o == nil || isNil(o.Request) {
    return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *SuccessAuthorizationResponse) HasRequest() bool {
	if o != nil && !isNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given Request and assigns it to the Request field.
func (o *SuccessAuthorizationResponse) SetRequest(v Request) {
	o.Request = &v
}

func (o SuccessAuthorizationResponse) MarshalJSON() ([]byte, error) {
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

type NullableSuccessAuthorizationResponse struct {
	value *SuccessAuthorizationResponse
	isSet bool
}

func (v NullableSuccessAuthorizationResponse) Get() *SuccessAuthorizationResponse {
	return v.value
}

func (v *NullableSuccessAuthorizationResponse) Set(val *SuccessAuthorizationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSuccessAuthorizationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSuccessAuthorizationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuccessAuthorizationResponse(val *SuccessAuthorizationResponse) *NullableSuccessAuthorizationResponse {
	return &NullableSuccessAuthorizationResponse{value: val, isSet: true}
}

func (v NullableSuccessAuthorizationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuccessAuthorizationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


