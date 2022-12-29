/*
Container Service

**API documentation for Container Service**

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package container_client

import (
	"encoding/json"
)

// ParentId struct for ParentId
type ParentId struct {
	// Field to represent container ID of the parent container. During the containerization process, this value is assigned to represent which container contains this container.
	ParentId string `json:"parentId"`
	// Field to represent action to be performed on the container.
	Action string `json:"action"`
}

// NewParentId instantiates a new ParentId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParentId(parentId string, action string) *ParentId {
	this := ParentId{}
	this.ParentId = parentId
	this.Action = action
	return &this
}

// NewParentIdWithDefaults instantiates a new ParentId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParentIdWithDefaults() *ParentId {
	this := ParentId{}
	return &this
}

// GetParentId returns the ParentId field value
func (o *ParentId) GetParentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value
// and a boolean to check if the value has been set.
func (o *ParentId) GetParentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentId, true
}

// SetParentId sets field value
func (o *ParentId) SetParentId(v string) {
	o.ParentId = v
}

// GetAction returns the Action field value
func (o *ParentId) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ParentId) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *ParentId) SetAction(v string) {
	o.Action = v
}

func (o ParentId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["parentId"] = o.ParentId
	}
	if true {
		toSerialize["action"] = o.Action
	}
	return json.Marshal(toSerialize)
}

type NullableParentId struct {
	value *ParentId
	isSet bool
}

func (v NullableParentId) Get() *ParentId {
	return v.value
}

func (v *NullableParentId) Set(val *ParentId) {
	v.value = val
	v.isSet = true
}

func (v NullableParentId) IsSet() bool {
	return v.isSet
}

func (v *NullableParentId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParentId(val *ParentId) *NullableParentId {
	return &NullableParentId{value: val, isSet: true}
}

func (v NullableParentId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParentId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


