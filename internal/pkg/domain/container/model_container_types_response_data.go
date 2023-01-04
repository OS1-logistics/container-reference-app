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

// ContainerTypesResponseData struct for ContainerTypesResponseData
type ContainerTypesResponseData struct {
	Meta *ContainerTypesResponseDataMeta `json:"meta,omitempty"`
	ContainerTypes []ContainerTypeConfigResponse `json:"containerTypes,omitempty"`
}

// NewContainerTypesResponseData instantiates a new ContainerTypesResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerTypesResponseData() *ContainerTypesResponseData {
	this := ContainerTypesResponseData{}
	return &this
}

// NewContainerTypesResponseDataWithDefaults instantiates a new ContainerTypesResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerTypesResponseDataWithDefaults() *ContainerTypesResponseData {
	this := ContainerTypesResponseData{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ContainerTypesResponseData) GetMeta() ContainerTypesResponseDataMeta {
	if o == nil || isNil(o.Meta) {
		var ret ContainerTypesResponseDataMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerTypesResponseData) GetMetaOk() (*ContainerTypesResponseDataMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ContainerTypesResponseData) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ContainerTypesResponseDataMeta and assigns it to the Meta field.
func (o *ContainerTypesResponseData) SetMeta(v ContainerTypesResponseDataMeta) {
	o.Meta = &v
}

// GetContainerTypes returns the ContainerTypes field value if set, zero value otherwise.
func (o *ContainerTypesResponseData) GetContainerTypes() []ContainerTypeConfigResponse {
	if o == nil || isNil(o.ContainerTypes) {
		var ret []ContainerTypeConfigResponse
		return ret
	}
	return o.ContainerTypes
}

// GetContainerTypesOk returns a tuple with the ContainerTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerTypesResponseData) GetContainerTypesOk() ([]ContainerTypeConfigResponse, bool) {
	if o == nil || isNil(o.ContainerTypes) {
    return nil, false
	}
	return o.ContainerTypes, true
}

// HasContainerTypes returns a boolean if a field has been set.
func (o *ContainerTypesResponseData) HasContainerTypes() bool {
	if o != nil && !isNil(o.ContainerTypes) {
		return true
	}

	return false
}

// SetContainerTypes gets a reference to the given []ContainerTypeConfigResponse and assigns it to the ContainerTypes field.
func (o *ContainerTypesResponseData) SetContainerTypes(v []ContainerTypeConfigResponse) {
	o.ContainerTypes = v
}

func (o ContainerTypesResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !isNil(o.ContainerTypes) {
		toSerialize["containerTypes"] = o.ContainerTypes
	}
	return json.Marshal(toSerialize)
}

type NullableContainerTypesResponseData struct {
	value *ContainerTypesResponseData
	isSet bool
}

func (v NullableContainerTypesResponseData) Get() *ContainerTypesResponseData {
	return v.value
}

func (v *NullableContainerTypesResponseData) Set(val *ContainerTypesResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerTypesResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerTypesResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerTypesResponseData(val *ContainerTypesResponseData) *NullableContainerTypesResponseData {
	return &NullableContainerTypesResponseData{value: val, isSet: true}
}

func (v NullableContainerTypesResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerTypesResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


