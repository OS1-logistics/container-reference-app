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

// GroupCreateRequest struct for GroupCreateRequest
type GroupCreateRequest struct {
	// Name of the group. Group name is unique for application.
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewGroupCreateRequest instantiates a new GroupCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupCreateRequest() *GroupCreateRequest {
	this := GroupCreateRequest{}
	return &this
}

// NewGroupCreateRequestWithDefaults instantiates a new GroupCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupCreateRequestWithDefaults() *GroupCreateRequest {
	this := GroupCreateRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GroupCreateRequest) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupCreateRequest) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GroupCreateRequest) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GroupCreateRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GroupCreateRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupCreateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GroupCreateRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GroupCreateRequest) SetDescription(v string) {
	o.Description = &v
}

func (o GroupCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableGroupCreateRequest struct {
	value *GroupCreateRequest
	isSet bool
}

func (v NullableGroupCreateRequest) Get() *GroupCreateRequest {
	return v.value
}

func (v *NullableGroupCreateRequest) Set(val *GroupCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupCreateRequest(val *GroupCreateRequest) *NullableGroupCreateRequest {
	return &NullableGroupCreateRequest{value: val, isSet: true}
}

func (v NullableGroupCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


