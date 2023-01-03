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

// AppMutableDataAllOf struct for AppMutableDataAllOf
type AppMutableDataAllOf struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

// NewAppMutableDataAllOf instantiates a new AppMutableDataAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppMutableDataAllOf() *AppMutableDataAllOf {
	this := AppMutableDataAllOf{}
	return &this
}

// NewAppMutableDataAllOfWithDefaults instantiates a new AppMutableDataAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppMutableDataAllOfWithDefaults() *AppMutableDataAllOf {
	this := AppMutableDataAllOf{}
	return &this
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *AppMutableDataAllOf) GetIsDeleted() bool {
	if o == nil || o.IsDeleted == nil {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppMutableDataAllOf) GetIsDeletedOk() (*bool, bool) {
	if o == nil || o.IsDeleted == nil {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *AppMutableDataAllOf) HasIsDeleted() bool {
	if o != nil && o.IsDeleted != nil {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *AppMutableDataAllOf) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

func (o AppMutableDataAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsDeleted != nil {
		toSerialize["isDeleted"] = o.IsDeleted
	}
	return json.Marshal(toSerialize)
}

type NullableAppMutableDataAllOf struct {
	value *AppMutableDataAllOf
	isSet bool
}

func (v NullableAppMutableDataAllOf) Get() *AppMutableDataAllOf {
	return v.value
}

func (v *NullableAppMutableDataAllOf) Set(val *AppMutableDataAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAppMutableDataAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAppMutableDataAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppMutableDataAllOf(val *AppMutableDataAllOf) *NullableAppMutableDataAllOf {
	return &NullableAppMutableDataAllOf{value: val, isSet: true}
}

func (v NullableAppMutableDataAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppMutableDataAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


