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

// ContainerBaseAttributes Base attributes that defines a container. This attributes are set on container creation and cannot be updated.
type ContainerBaseAttributes struct {
	// Container type defined using container type configuration APIs.
	ContainerType string `json:"containerType"`
	// Defines whether container is re-usable or not
	IsReusable *bool `json:"isReusable,omitempty"`
}

// NewContainerBaseAttributes instantiates a new ContainerBaseAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerBaseAttributes(containerType string) *ContainerBaseAttributes {
	this := ContainerBaseAttributes{}
	this.ContainerType = containerType
	var isReusable bool = false
	this.IsReusable = &isReusable
	return &this
}

// NewContainerBaseAttributesWithDefaults instantiates a new ContainerBaseAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerBaseAttributesWithDefaults() *ContainerBaseAttributes {
	this := ContainerBaseAttributes{}
	var isReusable bool = false
	this.IsReusable = &isReusable
	return &this
}

// GetContainerType returns the ContainerType field value
func (o *ContainerBaseAttributes) GetContainerType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContainerType
}

// GetContainerTypeOk returns a tuple with the ContainerType field value
// and a boolean to check if the value has been set.
func (o *ContainerBaseAttributes) GetContainerTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ContainerType, true
}

// SetContainerType sets field value
func (o *ContainerBaseAttributes) SetContainerType(v string) {
	o.ContainerType = v
}

// GetIsReusable returns the IsReusable field value if set, zero value otherwise.
func (o *ContainerBaseAttributes) GetIsReusable() bool {
	if o == nil || isNil(o.IsReusable) {
		var ret bool
		return ret
	}
	return *o.IsReusable
}

// GetIsReusableOk returns a tuple with the IsReusable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerBaseAttributes) GetIsReusableOk() (*bool, bool) {
	if o == nil || isNil(o.IsReusable) {
    return nil, false
	}
	return o.IsReusable, true
}

// HasIsReusable returns a boolean if a field has been set.
func (o *ContainerBaseAttributes) HasIsReusable() bool {
	if o != nil && !isNil(o.IsReusable) {
		return true
	}

	return false
}

// SetIsReusable gets a reference to the given bool and assigns it to the IsReusable field.
func (o *ContainerBaseAttributes) SetIsReusable(v bool) {
	o.IsReusable = &v
}

func (o ContainerBaseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["containerType"] = o.ContainerType
	}
	if !isNil(o.IsReusable) {
		toSerialize["isReusable"] = o.IsReusable
	}
	return json.Marshal(toSerialize)
}

type NullableContainerBaseAttributes struct {
	value *ContainerBaseAttributes
	isSet bool
}

func (v NullableContainerBaseAttributes) Get() *ContainerBaseAttributes {
	return v.value
}

func (v *NullableContainerBaseAttributes) Set(val *ContainerBaseAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerBaseAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerBaseAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerBaseAttributes(val *ContainerBaseAttributes) *NullableContainerBaseAttributes {
	return &NullableContainerBaseAttributes{value: val, isSet: true}
}

func (v NullableContainerBaseAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerBaseAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


