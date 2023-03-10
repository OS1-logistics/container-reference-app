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

// UserChangeLogResponseDataMeta struct for UserChangeLogResponseDataMeta
type UserChangeLogResponseDataMeta struct {
	TotalElements *int64 `json:"totalElements,omitempty"`
}

// NewUserChangeLogResponseDataMeta instantiates a new UserChangeLogResponseDataMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserChangeLogResponseDataMeta() *UserChangeLogResponseDataMeta {
	this := UserChangeLogResponseDataMeta{}
	return &this
}

// NewUserChangeLogResponseDataMetaWithDefaults instantiates a new UserChangeLogResponseDataMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserChangeLogResponseDataMetaWithDefaults() *UserChangeLogResponseDataMeta {
	this := UserChangeLogResponseDataMeta{}
	return &this
}

// GetTotalElements returns the TotalElements field value if set, zero value otherwise.
func (o *UserChangeLogResponseDataMeta) GetTotalElements() int64 {
	if o == nil || isNil(o.TotalElements) {
		var ret int64
		return ret
	}
	return *o.TotalElements
}

// GetTotalElementsOk returns a tuple with the TotalElements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserChangeLogResponseDataMeta) GetTotalElementsOk() (*int64, bool) {
	if o == nil || isNil(o.TotalElements) {
    return nil, false
	}
	return o.TotalElements, true
}

// HasTotalElements returns a boolean if a field has been set.
func (o *UserChangeLogResponseDataMeta) HasTotalElements() bool {
	if o != nil && !isNil(o.TotalElements) {
		return true
	}

	return false
}

// SetTotalElements gets a reference to the given int64 and assigns it to the TotalElements field.
func (o *UserChangeLogResponseDataMeta) SetTotalElements(v int64) {
	o.TotalElements = &v
}

func (o UserChangeLogResponseDataMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TotalElements) {
		toSerialize["totalElements"] = o.TotalElements
	}
	return json.Marshal(toSerialize)
}

type NullableUserChangeLogResponseDataMeta struct {
	value *UserChangeLogResponseDataMeta
	isSet bool
}

func (v NullableUserChangeLogResponseDataMeta) Get() *UserChangeLogResponseDataMeta {
	return v.value
}

func (v *NullableUserChangeLogResponseDataMeta) Set(val *UserChangeLogResponseDataMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableUserChangeLogResponseDataMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableUserChangeLogResponseDataMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserChangeLogResponseDataMeta(val *UserChangeLogResponseDataMeta) *NullableUserChangeLogResponseDataMeta {
	return &NullableUserChangeLogResponseDataMeta{value: val, isSet: true}
}

func (v NullableUserChangeLogResponseDataMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserChangeLogResponseDataMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


