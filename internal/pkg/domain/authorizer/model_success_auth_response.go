/*
AAA - External Authorizer

This swagger documentation provides all Authorization api details. External Authorizer service provides authentication and authorization capabilities for users and machines.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authorizerdomain

import (
	"encoding/json"
)

// SuccessAuthResponse struct for SuccessAuthResponse
type SuccessAuthResponse struct {
	Data *SuccessAuthResponseData `json:"data,omitempty"`
	Error *Error `json:"error,omitempty"`
	Request *Request `json:"request,omitempty"`
}

// NewSuccessAuthResponse instantiates a new SuccessAuthResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuccessAuthResponse() *SuccessAuthResponse {
	this := SuccessAuthResponse{}
	return &this
}

// NewSuccessAuthResponseWithDefaults instantiates a new SuccessAuthResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuccessAuthResponseWithDefaults() *SuccessAuthResponse {
	this := SuccessAuthResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SuccessAuthResponse) GetData() SuccessAuthResponseData {
	if o == nil || isNil(o.Data) {
		var ret SuccessAuthResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessAuthResponse) GetDataOk() (*SuccessAuthResponseData, bool) {
	if o == nil || isNil(o.Data) {
    return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SuccessAuthResponse) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given SuccessAuthResponseData and assigns it to the Data field.
func (o *SuccessAuthResponse) SetData(v SuccessAuthResponseData) {
	o.Data = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SuccessAuthResponse) GetError() Error {
	if o == nil || isNil(o.Error) {
		var ret Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessAuthResponse) GetErrorOk() (*Error, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SuccessAuthResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given Error and assigns it to the Error field.
func (o *SuccessAuthResponse) SetError(v Error) {
	o.Error = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *SuccessAuthResponse) GetRequest() Request {
	if o == nil || isNil(o.Request) {
		var ret Request
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessAuthResponse) GetRequestOk() (*Request, bool) {
	if o == nil || isNil(o.Request) {
    return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *SuccessAuthResponse) HasRequest() bool {
	if o != nil && !isNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given Request and assigns it to the Request field.
func (o *SuccessAuthResponse) SetRequest(v Request) {
	o.Request = &v
}

func (o SuccessAuthResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !isNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	return json.Marshal(toSerialize)
}

type NullableSuccessAuthResponse struct {
	value *SuccessAuthResponse
	isSet bool
}

func (v NullableSuccessAuthResponse) Get() *SuccessAuthResponse {
	return v.value
}

func (v *NullableSuccessAuthResponse) Set(val *SuccessAuthResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSuccessAuthResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSuccessAuthResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuccessAuthResponse(val *SuccessAuthResponse) *NullableSuccessAuthResponse {
	return &NullableSuccessAuthResponse{value: val, isSet: true}
}

func (v NullableSuccessAuthResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuccessAuthResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

