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

// ContainerTypesResponseDataMeta struct for ContainerTypesResponseDataMeta
type ContainerTypesResponseDataMeta struct {
	TotalElements *float32 `json:"totalElements,omitempty"`
}

// NewContainerTypesResponseDataMeta instantiates a new ContainerTypesResponseDataMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerTypesResponseDataMeta() *ContainerTypesResponseDataMeta {
	this := ContainerTypesResponseDataMeta{}
	return &this
}

// NewContainerTypesResponseDataMetaWithDefaults instantiates a new ContainerTypesResponseDataMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerTypesResponseDataMetaWithDefaults() *ContainerTypesResponseDataMeta {
	this := ContainerTypesResponseDataMeta{}
	return &this
}

// GetTotalElements returns the TotalElements field value if set, zero value otherwise.
func (o *ContainerTypesResponseDataMeta) GetTotalElements() float32 {
	if o == nil || o.TotalElements == nil {
		var ret float32
		return ret
	}
	return *o.TotalElements
}

// GetTotalElementsOk returns a tuple with the TotalElements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerTypesResponseDataMeta) GetTotalElementsOk() (*float32, bool) {
	if o == nil || o.TotalElements == nil {
		return nil, false
	}
	return o.TotalElements, true
}

// HasTotalElements returns a boolean if a field has been set.
func (o *ContainerTypesResponseDataMeta) HasTotalElements() bool {
	if o != nil && o.TotalElements != nil {
		return true
	}

	return false
}

// SetTotalElements gets a reference to the given float32 and assigns it to the TotalElements field.
func (o *ContainerTypesResponseDataMeta) SetTotalElements(v float32) {
	o.TotalElements = &v
}

func (o ContainerTypesResponseDataMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TotalElements != nil {
		toSerialize["totalElements"] = o.TotalElements
	}
	return json.Marshal(toSerialize)
}

type NullableContainerTypesResponseDataMeta struct {
	value *ContainerTypesResponseDataMeta
	isSet bool
}

func (v NullableContainerTypesResponseDataMeta) Get() *ContainerTypesResponseDataMeta {
	return v.value
}

func (v *NullableContainerTypesResponseDataMeta) Set(val *ContainerTypesResponseDataMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerTypesResponseDataMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerTypesResponseDataMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerTypesResponseDataMeta(val *ContainerTypesResponseDataMeta) *NullableContainerTypesResponseDataMeta {
	return &NullableContainerTypesResponseDataMeta{value: val, isSet: true}
}

func (v NullableContainerTypesResponseDataMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerTypesResponseDataMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


