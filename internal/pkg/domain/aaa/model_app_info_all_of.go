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

// AppInfoAllOf struct for AppInfoAllOf
type AppInfoAllOf struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

// NewAppInfoAllOf instantiates a new AppInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppInfoAllOf() *AppInfoAllOf {
	this := AppInfoAllOf{}
	var isDeleted bool = false
	this.IsDeleted = &isDeleted
	return &this
}

// NewAppInfoAllOfWithDefaults instantiates a new AppInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppInfoAllOfWithDefaults() *AppInfoAllOf {
	this := AppInfoAllOf{}
	var isDeleted bool = false
	this.IsDeleted = &isDeleted
	return &this
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *AppInfoAllOf) GetIsDeleted() bool {
	if o == nil || isNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppInfoAllOf) GetIsDeletedOk() (*bool, bool) {
	if o == nil || isNil(o.IsDeleted) {
    return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *AppInfoAllOf) HasIsDeleted() bool {
	if o != nil && !isNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *AppInfoAllOf) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

func (o AppInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IsDeleted) {
		toSerialize["isDeleted"] = o.IsDeleted
	}
	return json.Marshal(toSerialize)
}

type NullableAppInfoAllOf struct {
	value *AppInfoAllOf
	isSet bool
}

func (v NullableAppInfoAllOf) Get() *AppInfoAllOf {
	return v.value
}

func (v *NullableAppInfoAllOf) Set(val *AppInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAppInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAppInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppInfoAllOf(val *AppInfoAllOf) *NullableAppInfoAllOf {
	return &NullableAppInfoAllOf{value: val, isSet: true}
}

func (v NullableAppInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


