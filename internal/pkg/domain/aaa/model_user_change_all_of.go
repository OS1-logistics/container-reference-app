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

// UserChangeAllOf struct for UserChangeAllOf
type UserChangeAllOf struct {
	Cdc UserResponse `json:"cdc"`
}

// NewUserChangeAllOf instantiates a new UserChangeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserChangeAllOf(cdc UserResponse) *UserChangeAllOf {
	this := UserChangeAllOf{}
	this.Cdc = cdc
	return &this
}

// NewUserChangeAllOfWithDefaults instantiates a new UserChangeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserChangeAllOfWithDefaults() *UserChangeAllOf {
	this := UserChangeAllOf{}
	return &this
}

// GetCdc returns the Cdc field value
func (o *UserChangeAllOf) GetCdc() UserResponse {
	if o == nil {
		var ret UserResponse
		return ret
	}

	return o.Cdc
}

// GetCdcOk returns a tuple with the Cdc field value
// and a boolean to check if the value has been set.
func (o *UserChangeAllOf) GetCdcOk() (*UserResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cdc, true
}

// SetCdc sets field value
func (o *UserChangeAllOf) SetCdc(v UserResponse) {
	o.Cdc = v
}

func (o UserChangeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cdc"] = o.Cdc
	}
	return json.Marshal(toSerialize)
}

type NullableUserChangeAllOf struct {
	value *UserChangeAllOf
	isSet bool
}

func (v NullableUserChangeAllOf) Get() *UserChangeAllOf {
	return v.value
}

func (v *NullableUserChangeAllOf) Set(val *UserChangeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUserChangeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUserChangeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserChangeAllOf(val *UserChangeAllOf) *NullableUserChangeAllOf {
	return &NullableUserChangeAllOf{value: val, isSet: true}
}

func (v NullableUserChangeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserChangeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


