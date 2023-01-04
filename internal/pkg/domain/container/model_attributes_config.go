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

// AttributesConfig struct for AttributesConfig
type AttributesConfig struct {
	Attributes []AttributeConfig `json:"attributes"`
}

// NewAttributesConfig instantiates a new AttributesConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttributesConfig(attributes []AttributeConfig) *AttributesConfig {
	this := AttributesConfig{}
	this.Attributes = attributes
	return &this
}

// NewAttributesConfigWithDefaults instantiates a new AttributesConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttributesConfigWithDefaults() *AttributesConfig {
	this := AttributesConfig{}
	return &this
}

// GetAttributes returns the Attributes field value
func (o *AttributesConfig) GetAttributes() []AttributeConfig {
	if o == nil {
		var ret []AttributeConfig
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AttributesConfig) GetAttributesOk() ([]AttributeConfig, bool) {
	if o == nil {
    return nil, false
	}
	return o.Attributes, true
}

// SetAttributes sets field value
func (o *AttributesConfig) SetAttributes(v []AttributeConfig) {
	o.Attributes = v
}

func (o AttributesConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableAttributesConfig struct {
	value *AttributesConfig
	isSet bool
}

func (v NullableAttributesConfig) Get() *AttributesConfig {
	return v.value
}

func (v *NullableAttributesConfig) Set(val *AttributesConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAttributesConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAttributesConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttributesConfig(val *AttributesConfig) *NullableAttributesConfig {
	return &NullableAttributesConfig{value: val, isSet: true}
}

func (v NullableAttributesConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttributesConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


