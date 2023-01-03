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

// SuccessConnectionResponse struct for SuccessConnectionResponse
type SuccessConnectionResponse struct {
	Error *Error `json:"error,omitempty"`
	Request *Request `json:"request,omitempty"`
	Data *SuccessConnectionResponseAllOfData `json:"data,omitempty"`
}

// NewSuccessConnectionResponse instantiates a new SuccessConnectionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuccessConnectionResponse() *SuccessConnectionResponse {
	this := SuccessConnectionResponse{}
	return &this
}

// NewSuccessConnectionResponseWithDefaults instantiates a new SuccessConnectionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuccessConnectionResponseWithDefaults() *SuccessConnectionResponse {
	this := SuccessConnectionResponse{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SuccessConnectionResponse) GetError() Error {
	if o == nil || o.Error == nil {
		var ret Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessConnectionResponse) GetErrorOk() (*Error, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SuccessConnectionResponse) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given Error and assigns it to the Error field.
func (o *SuccessConnectionResponse) SetError(v Error) {
	o.Error = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *SuccessConnectionResponse) GetRequest() Request {
	if o == nil || o.Request == nil {
		var ret Request
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessConnectionResponse) GetRequestOk() (*Request, bool) {
	if o == nil || o.Request == nil {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *SuccessConnectionResponse) HasRequest() bool {
	if o != nil && o.Request != nil {
		return true
	}

	return false
}

// SetRequest gets a reference to the given Request and assigns it to the Request field.
func (o *SuccessConnectionResponse) SetRequest(v Request) {
	o.Request = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SuccessConnectionResponse) GetData() SuccessConnectionResponseAllOfData {
	if o == nil || o.Data == nil {
		var ret SuccessConnectionResponseAllOfData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessConnectionResponse) GetDataOk() (*SuccessConnectionResponseAllOfData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SuccessConnectionResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given SuccessConnectionResponseAllOfData and assigns it to the Data field.
func (o *SuccessConnectionResponse) SetData(v SuccessConnectionResponseAllOfData) {
	o.Data = &v
}

func (o SuccessConnectionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.Request != nil {
		toSerialize["request"] = o.Request
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableSuccessConnectionResponse struct {
	value *SuccessConnectionResponse
	isSet bool
}

func (v NullableSuccessConnectionResponse) Get() *SuccessConnectionResponse {
	return v.value
}

func (v *NullableSuccessConnectionResponse) Set(val *SuccessConnectionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSuccessConnectionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSuccessConnectionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuccessConnectionResponse(val *SuccessConnectionResponse) *NullableSuccessConnectionResponse {
	return &NullableSuccessConnectionResponse{value: val, isSet: true}
}

func (v NullableSuccessConnectionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuccessConnectionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


