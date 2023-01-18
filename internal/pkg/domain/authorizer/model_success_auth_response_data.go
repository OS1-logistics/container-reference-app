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

// SuccessAuthResponseData struct for SuccessAuthResponseData
type SuccessAuthResponseData struct {
	// Status of operation.
	Response *string `json:"response,omitempty"`
}

// NewSuccessAuthResponseData instantiates a new SuccessAuthResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuccessAuthResponseData() *SuccessAuthResponseData {
	this := SuccessAuthResponseData{}
	return &this
}

// NewSuccessAuthResponseDataWithDefaults instantiates a new SuccessAuthResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuccessAuthResponseDataWithDefaults() *SuccessAuthResponseData {
	this := SuccessAuthResponseData{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *SuccessAuthResponseData) GetResponse() string {
	if o == nil || isNil(o.Response) {
		var ret string
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessAuthResponseData) GetResponseOk() (*string, bool) {
	if o == nil || isNil(o.Response) {
    return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *SuccessAuthResponseData) HasResponse() bool {
	if o != nil && !isNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given string and assigns it to the Response field.
func (o *SuccessAuthResponseData) SetResponse(v string) {
	o.Response = &v
}

func (o SuccessAuthResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Response) {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullableSuccessAuthResponseData struct {
	value *SuccessAuthResponseData
	isSet bool
}

func (v NullableSuccessAuthResponseData) Get() *SuccessAuthResponseData {
	return v.value
}

func (v *NullableSuccessAuthResponseData) Set(val *SuccessAuthResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableSuccessAuthResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableSuccessAuthResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuccessAuthResponseData(val *SuccessAuthResponseData) *NullableSuccessAuthResponseData {
	return &NullableSuccessAuthResponseData{value: val, isSet: true}
}

func (v NullableSuccessAuthResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuccessAuthResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


