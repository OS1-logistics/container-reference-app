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

// SuccessGroupsResponseV2Data struct for SuccessGroupsResponseV2Data
type SuccessGroupsResponseV2Data struct {
	Meta *Meta `json:"meta,omitempty"`
	Groups []GroupsResponseV2Inner `json:"groups,omitempty"`
}

// NewSuccessGroupsResponseV2Data instantiates a new SuccessGroupsResponseV2Data object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuccessGroupsResponseV2Data() *SuccessGroupsResponseV2Data {
	this := SuccessGroupsResponseV2Data{}
	return &this
}

// NewSuccessGroupsResponseV2DataWithDefaults instantiates a new SuccessGroupsResponseV2Data object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuccessGroupsResponseV2DataWithDefaults() *SuccessGroupsResponseV2Data {
	this := SuccessGroupsResponseV2Data{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SuccessGroupsResponseV2Data) GetMeta() Meta {
	if o == nil || o.Meta == nil {
		var ret Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessGroupsResponseV2Data) GetMetaOk() (*Meta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SuccessGroupsResponseV2Data) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *SuccessGroupsResponseV2Data) SetMeta(v Meta) {
	o.Meta = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *SuccessGroupsResponseV2Data) GetGroups() []GroupsResponseV2Inner {
	if o == nil || o.Groups == nil {
		var ret []GroupsResponseV2Inner
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessGroupsResponseV2Data) GetGroupsOk() ([]GroupsResponseV2Inner, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *SuccessGroupsResponseV2Data) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []GroupsResponseV2Inner and assigns it to the Groups field.
func (o *SuccessGroupsResponseV2Data) SetGroups(v []GroupsResponseV2Inner) {
	o.Groups = v
}

func (o SuccessGroupsResponseV2Data) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	return json.Marshal(toSerialize)
}

type NullableSuccessGroupsResponseV2Data struct {
	value *SuccessGroupsResponseV2Data
	isSet bool
}

func (v NullableSuccessGroupsResponseV2Data) Get() *SuccessGroupsResponseV2Data {
	return v.value
}

func (v *NullableSuccessGroupsResponseV2Data) Set(val *SuccessGroupsResponseV2Data) {
	v.value = val
	v.isSet = true
}

func (v NullableSuccessGroupsResponseV2Data) IsSet() bool {
	return v.isSet
}

func (v *NullableSuccessGroupsResponseV2Data) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuccessGroupsResponseV2Data(val *SuccessGroupsResponseV2Data) *NullableSuccessGroupsResponseV2Data {
	return &NullableSuccessGroupsResponseV2Data{value: val, isSet: true}
}

func (v NullableSuccessGroupsResponseV2Data) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuccessGroupsResponseV2Data) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


