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

// GroupUpdateRequestUsers struct for GroupUpdateRequestUsers
type GroupUpdateRequestUsers struct {
	// list of user IDs.
	UserIds []string `json:"userIds,omitempty"`
	// This is used for add/remove users membership of a group
	Membership *bool `json:"membership,omitempty"`
}

// NewGroupUpdateRequestUsers instantiates a new GroupUpdateRequestUsers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupUpdateRequestUsers() *GroupUpdateRequestUsers {
	this := GroupUpdateRequestUsers{}
	return &this
}

// NewGroupUpdateRequestUsersWithDefaults instantiates a new GroupUpdateRequestUsers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupUpdateRequestUsersWithDefaults() *GroupUpdateRequestUsers {
	this := GroupUpdateRequestUsers{}
	return &this
}

// GetUserIds returns the UserIds field value if set, zero value otherwise.
func (o *GroupUpdateRequestUsers) GetUserIds() []string {
	if o == nil || isNil(o.UserIds) {
		var ret []string
		return ret
	}
	return o.UserIds
}

// GetUserIdsOk returns a tuple with the UserIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupUpdateRequestUsers) GetUserIdsOk() ([]string, bool) {
	if o == nil || isNil(o.UserIds) {
    return nil, false
	}
	return o.UserIds, true
}

// HasUserIds returns a boolean if a field has been set.
func (o *GroupUpdateRequestUsers) HasUserIds() bool {
	if o != nil && !isNil(o.UserIds) {
		return true
	}

	return false
}

// SetUserIds gets a reference to the given []string and assigns it to the UserIds field.
func (o *GroupUpdateRequestUsers) SetUserIds(v []string) {
	o.UserIds = v
}

// GetMembership returns the Membership field value if set, zero value otherwise.
func (o *GroupUpdateRequestUsers) GetMembership() bool {
	if o == nil || isNil(o.Membership) {
		var ret bool
		return ret
	}
	return *o.Membership
}

// GetMembershipOk returns a tuple with the Membership field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupUpdateRequestUsers) GetMembershipOk() (*bool, bool) {
	if o == nil || isNil(o.Membership) {
    return nil, false
	}
	return o.Membership, true
}

// HasMembership returns a boolean if a field has been set.
func (o *GroupUpdateRequestUsers) HasMembership() bool {
	if o != nil && !isNil(o.Membership) {
		return true
	}

	return false
}

// SetMembership gets a reference to the given bool and assigns it to the Membership field.
func (o *GroupUpdateRequestUsers) SetMembership(v bool) {
	o.Membership = &v
}

func (o GroupUpdateRequestUsers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.UserIds) {
		toSerialize["userIds"] = o.UserIds
	}
	if !isNil(o.Membership) {
		toSerialize["membership"] = o.Membership
	}
	return json.Marshal(toSerialize)
}

type NullableGroupUpdateRequestUsers struct {
	value *GroupUpdateRequestUsers
	isSet bool
}

func (v NullableGroupUpdateRequestUsers) Get() *GroupUpdateRequestUsers {
	return v.value
}

func (v *NullableGroupUpdateRequestUsers) Set(val *GroupUpdateRequestUsers) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupUpdateRequestUsers) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupUpdateRequestUsers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupUpdateRequestUsers(val *GroupUpdateRequestUsers) *NullableGroupUpdateRequestUsers {
	return &NullableGroupUpdateRequestUsers{value: val, isSet: true}
}

func (v NullableGroupUpdateRequestUsers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupUpdateRequestUsers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


