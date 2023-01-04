/*
Container Service

**API documentation for Container Service**

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package containerdomain

import (
	"encoding/json"
)

// ContainerTypeAllowedParent Rules defining the allowed container types as parent of a container type
type ContainerTypeAllowedParent struct {
	// List of container types allowed as parent of this container type
	OneOf []string `json:"oneOf,omitempty"`
	// List of container types *not* allowed as parent of this container type
	Not []string `json:"not,omitempty"`
}

// NewContainerTypeAllowedParent instantiates a new ContainerTypeAllowedParent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerTypeAllowedParent() *ContainerTypeAllowedParent {
	this := ContainerTypeAllowedParent{}
	return &this
}

// NewContainerTypeAllowedParentWithDefaults instantiates a new ContainerTypeAllowedParent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerTypeAllowedParentWithDefaults() *ContainerTypeAllowedParent {
	this := ContainerTypeAllowedParent{}
	return &this
}

// GetOneOf returns the OneOf field value if set, zero value otherwise.
func (o *ContainerTypeAllowedParent) GetOneOf() []string {
	if o == nil || o.OneOf == nil {
		var ret []string
		return ret
	}
	return o.OneOf
}

// GetOneOfOk returns a tuple with the OneOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerTypeAllowedParent) GetOneOfOk() ([]string, bool) {
	if o == nil || o.OneOf == nil {
		return nil, false
	}
	return o.OneOf, true
}

// HasOneOf returns a boolean if a field has been set.
func (o *ContainerTypeAllowedParent) HasOneOf() bool {
	if o != nil && o.OneOf != nil {
		return true
	}

	return false
}

// SetOneOf gets a reference to the given []string and assigns it to the OneOf field.
func (o *ContainerTypeAllowedParent) SetOneOf(v []string) {
	o.OneOf = v
}

// GetNot returns the Not field value if set, zero value otherwise.
func (o *ContainerTypeAllowedParent) GetNot() []string {
	if o == nil || o.Not == nil {
		var ret []string
		return ret
	}
	return o.Not
}

// GetNotOk returns a tuple with the Not field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerTypeAllowedParent) GetNotOk() ([]string, bool) {
	if o == nil || o.Not == nil {
		return nil, false
	}
	return o.Not, true
}

// HasNot returns a boolean if a field has been set.
func (o *ContainerTypeAllowedParent) HasNot() bool {
	if o != nil && o.Not != nil {
		return true
	}

	return false
}

// SetNot gets a reference to the given []string and assigns it to the Not field.
func (o *ContainerTypeAllowedParent) SetNot(v []string) {
	o.Not = v
}

func (o ContainerTypeAllowedParent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OneOf != nil {
		toSerialize["oneOf"] = o.OneOf
	}
	if o.Not != nil {
		toSerialize["not"] = o.Not
	}
	return json.Marshal(toSerialize)
}

type NullableContainerTypeAllowedParent struct {
	value *ContainerTypeAllowedParent
	isSet bool
}

func (v NullableContainerTypeAllowedParent) Get() *ContainerTypeAllowedParent {
	return v.value
}

func (v *NullableContainerTypeAllowedParent) Set(val *ContainerTypeAllowedParent) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerTypeAllowedParent) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerTypeAllowedParent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerTypeAllowedParent(val *ContainerTypeAllowedParent) *NullableContainerTypeAllowedParent {
	return &NullableContainerTypeAllowedParent{value: val, isSet: true}
}

func (v NullableContainerTypeAllowedParent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerTypeAllowedParent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


