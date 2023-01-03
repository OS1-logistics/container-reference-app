/*
Authentication And Authorization (AAA) Service

This swagger documentation provides all AAA API details. AAA service provides authentication and authorization capabilities for users.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aaa_client

import (
	"encoding/json"
)

// SuccessUserResponse struct for SuccessUserResponse
type SuccessUserResponse struct {
	Error *Error `json:"error,omitempty"`
	Data *UserResponse `json:"data,omitempty"`
	Request *Request `json:"request,omitempty"`
}

// NewSuccessUserResponse instantiates a new SuccessUserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuccessUserResponse() *SuccessUserResponse {
	this := SuccessUserResponse{}
	return &this
}

// NewSuccessUserResponseWithDefaults instantiates a new SuccessUserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuccessUserResponseWithDefaults() *SuccessUserResponse {
	this := SuccessUserResponse{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SuccessUserResponse) GetError() Error {
	if o == nil || o.Error == nil {
		var ret Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessUserResponse) GetErrorOk() (*Error, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SuccessUserResponse) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given Error and assigns it to the Error field.
func (o *SuccessUserResponse) SetError(v Error) {
	o.Error = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SuccessUserResponse) GetData() UserResponse {
	if o == nil || o.Data == nil {
		var ret UserResponse
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessUserResponse) GetDataOk() (*UserResponse, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SuccessUserResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given UserResponse and assigns it to the Data field.
func (o *SuccessUserResponse) SetData(v UserResponse) {
	o.Data = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *SuccessUserResponse) GetRequest() Request {
	if o == nil || o.Request == nil {
		var ret Request
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessUserResponse) GetRequestOk() (*Request, bool) {
	if o == nil || o.Request == nil {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *SuccessUserResponse) HasRequest() bool {
	if o != nil && o.Request != nil {
		return true
	}

	return false
}

// SetRequest gets a reference to the given Request and assigns it to the Request field.
func (o *SuccessUserResponse) SetRequest(v Request) {
	o.Request = &v
}

func (o SuccessUserResponse) MarshalJSON() ([]byte, error) {
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

type NullableSuccessUserResponse struct {
	value *SuccessUserResponse
	isSet bool
}

func (v NullableSuccessUserResponse) Get() *SuccessUserResponse {
	return v.value
}

func (v *NullableSuccessUserResponse) Set(val *SuccessUserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSuccessUserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSuccessUserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuccessUserResponse(val *SuccessUserResponse) *NullableSuccessUserResponse {
	return &NullableSuccessUserResponse{value: val, isSet: true}
}

func (v NullableSuccessUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuccessUserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


