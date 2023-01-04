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

// RoleMappingRequest struct for RoleMappingRequest
type RoleMappingRequest struct {
	// Boolean field to identify whether to assign or remove role. The same action (assign or remove) will be applicable on all the roles.
	Assign bool `json:"assign"`
	// List of roles to be added or removed from group.
	Roles []RoleMapping `json:"roles"`
}

// NewRoleMappingRequest instantiates a new RoleMappingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleMappingRequest(assign bool, roles []RoleMapping) *RoleMappingRequest {
	this := RoleMappingRequest{}
	this.Assign = assign
	this.Roles = roles
	return &this
}

// NewRoleMappingRequestWithDefaults instantiates a new RoleMappingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleMappingRequestWithDefaults() *RoleMappingRequest {
	this := RoleMappingRequest{}
	return &this
}

// GetAssign returns the Assign field value
func (o *RoleMappingRequest) GetAssign() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Assign
}

// GetAssignOk returns a tuple with the Assign field value
// and a boolean to check if the value has been set.
func (o *RoleMappingRequest) GetAssignOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Assign, true
}

// SetAssign sets field value
func (o *RoleMappingRequest) SetAssign(v bool) {
	o.Assign = v
}

// GetRoles returns the Roles field value
func (o *RoleMappingRequest) GetRoles() []RoleMapping {
	if o == nil {
		var ret []RoleMapping
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *RoleMappingRequest) GetRolesOk() ([]RoleMapping, bool) {
	if o == nil {
    return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *RoleMappingRequest) SetRoles(v []RoleMapping) {
	o.Roles = v
}

func (o RoleMappingRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["assign"] = o.Assign
	}
	if true {
		toSerialize["roles"] = o.Roles
	}
	return json.Marshal(toSerialize)
}

type NullableRoleMappingRequest struct {
	value *RoleMappingRequest
	isSet bool
}

func (v NullableRoleMappingRequest) Get() *RoleMappingRequest {
	return v.value
}

func (v *NullableRoleMappingRequest) Set(val *RoleMappingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleMappingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleMappingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleMappingRequest(val *RoleMappingRequest) *NullableRoleMappingRequest {
	return &NullableRoleMappingRequest{value: val, isSet: true}
}

func (v NullableRoleMappingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleMappingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


