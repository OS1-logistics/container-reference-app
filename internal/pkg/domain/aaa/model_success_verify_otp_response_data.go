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

// SuccessVerifyOtpResponseData struct for SuccessVerifyOtpResponseData
type SuccessVerifyOtpResponseData struct {
	// UserId of the user which is authenticated using email or mobile Number.
	UserId *string `json:"userId,omitempty"`
}

// NewSuccessVerifyOtpResponseData instantiates a new SuccessVerifyOtpResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuccessVerifyOtpResponseData() *SuccessVerifyOtpResponseData {
	this := SuccessVerifyOtpResponseData{}
	return &this
}

// NewSuccessVerifyOtpResponseDataWithDefaults instantiates a new SuccessVerifyOtpResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuccessVerifyOtpResponseDataWithDefaults() *SuccessVerifyOtpResponseData {
	this := SuccessVerifyOtpResponseData{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *SuccessVerifyOtpResponseData) GetUserId() string {
	if o == nil || isNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessVerifyOtpResponseData) GetUserIdOk() (*string, bool) {
	if o == nil || isNil(o.UserId) {
    return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *SuccessVerifyOtpResponseData) HasUserId() bool {
	if o != nil && !isNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *SuccessVerifyOtpResponseData) SetUserId(v string) {
	o.UserId = &v
}

func (o SuccessVerifyOtpResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return json.Marshal(toSerialize)
}

type NullableSuccessVerifyOtpResponseData struct {
	value *SuccessVerifyOtpResponseData
	isSet bool
}

func (v NullableSuccessVerifyOtpResponseData) Get() *SuccessVerifyOtpResponseData {
	return v.value
}

func (v *NullableSuccessVerifyOtpResponseData) Set(val *SuccessVerifyOtpResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableSuccessVerifyOtpResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableSuccessVerifyOtpResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuccessVerifyOtpResponseData(val *SuccessVerifyOtpResponseData) *NullableSuccessVerifyOtpResponseData {
	return &NullableSuccessVerifyOtpResponseData{value: val, isSet: true}
}

func (v NullableSuccessVerifyOtpResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuccessVerifyOtpResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

