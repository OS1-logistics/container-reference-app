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

// SuccessAppSecretCreateResponse struct for SuccessAppSecretCreateResponse
type SuccessAppSecretCreateResponse struct {
	Error *Error `json:"error,omitempty"`
	Data *AppSecretResponse `json:"data,omitempty"`
	Request *Request `json:"request,omitempty"`
}

// NewSuccessAppSecretCreateResponse instantiates a new SuccessAppSecretCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuccessAppSecretCreateResponse() *SuccessAppSecretCreateResponse {
	this := SuccessAppSecretCreateResponse{}
	return &this
}

// NewSuccessAppSecretCreateResponseWithDefaults instantiates a new SuccessAppSecretCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuccessAppSecretCreateResponseWithDefaults() *SuccessAppSecretCreateResponse {
	this := SuccessAppSecretCreateResponse{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SuccessAppSecretCreateResponse) GetError() Error {
	if o == nil || o.Error == nil {
		var ret Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessAppSecretCreateResponse) GetErrorOk() (*Error, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SuccessAppSecretCreateResponse) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given Error and assigns it to the Error field.
func (o *SuccessAppSecretCreateResponse) SetError(v Error) {
	o.Error = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SuccessAppSecretCreateResponse) GetData() AppSecretResponse {
	if o == nil || o.Data == nil {
		var ret AppSecretResponse
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessAppSecretCreateResponse) GetDataOk() (*AppSecretResponse, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SuccessAppSecretCreateResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given AppSecretResponse and assigns it to the Data field.
func (o *SuccessAppSecretCreateResponse) SetData(v AppSecretResponse) {
	o.Data = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *SuccessAppSecretCreateResponse) GetRequest() Request {
	if o == nil || o.Request == nil {
		var ret Request
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessAppSecretCreateResponse) GetRequestOk() (*Request, bool) {
	if o == nil || o.Request == nil {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *SuccessAppSecretCreateResponse) HasRequest() bool {
	if o != nil && o.Request != nil {
		return true
	}

	return false
}

// SetRequest gets a reference to the given Request and assigns it to the Request field.
func (o *SuccessAppSecretCreateResponse) SetRequest(v Request) {
	o.Request = &v
}

func (o SuccessAppSecretCreateResponse) MarshalJSON() ([]byte, error) {
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

type NullableSuccessAppSecretCreateResponse struct {
	value *SuccessAppSecretCreateResponse
	isSet bool
}

func (v NullableSuccessAppSecretCreateResponse) Get() *SuccessAppSecretCreateResponse {
	return v.value
}

func (v *NullableSuccessAppSecretCreateResponse) Set(val *SuccessAppSecretCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSuccessAppSecretCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSuccessAppSecretCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuccessAppSecretCreateResponse(val *SuccessAppSecretCreateResponse) *NullableSuccessAppSecretCreateResponse {
	return &NullableSuccessAppSecretCreateResponse{value: val, isSet: true}
}

func (v NullableSuccessAppSecretCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuccessAppSecretCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


