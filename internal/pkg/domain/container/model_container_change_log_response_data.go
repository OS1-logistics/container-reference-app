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

// ContainerChangeLogResponseData struct for ContainerChangeLogResponseData
type ContainerChangeLogResponseData struct {
	Meta *ContainerTypesResponseDataMeta `json:"meta,omitempty"`
	Log []ContainerBaseChangeLogResponse `json:"log,omitempty"`
}

// NewContainerChangeLogResponseData instantiates a new ContainerChangeLogResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerChangeLogResponseData() *ContainerChangeLogResponseData {
	this := ContainerChangeLogResponseData{}
	return &this
}

// NewContainerChangeLogResponseDataWithDefaults instantiates a new ContainerChangeLogResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerChangeLogResponseDataWithDefaults() *ContainerChangeLogResponseData {
	this := ContainerChangeLogResponseData{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ContainerChangeLogResponseData) GetMeta() ContainerTypesResponseDataMeta {
	if o == nil || isNil(o.Meta) {
		var ret ContainerTypesResponseDataMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerChangeLogResponseData) GetMetaOk() (*ContainerTypesResponseDataMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ContainerChangeLogResponseData) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ContainerTypesResponseDataMeta and assigns it to the Meta field.
func (o *ContainerChangeLogResponseData) SetMeta(v ContainerTypesResponseDataMeta) {
	o.Meta = &v
}

// GetLog returns the Log field value if set, zero value otherwise.
func (o *ContainerChangeLogResponseData) GetLog() []ContainerBaseChangeLogResponse {
	if o == nil || isNil(o.Log) {
		var ret []ContainerBaseChangeLogResponse
		return ret
	}
	return o.Log
}

// GetLogOk returns a tuple with the Log field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerChangeLogResponseData) GetLogOk() ([]ContainerBaseChangeLogResponse, bool) {
	if o == nil || isNil(o.Log) {
    return nil, false
	}
	return o.Log, true
}

// HasLog returns a boolean if a field has been set.
func (o *ContainerChangeLogResponseData) HasLog() bool {
	if o != nil && !isNil(o.Log) {
		return true
	}

	return false
}

// SetLog gets a reference to the given []ContainerBaseChangeLogResponse and assigns it to the Log field.
func (o *ContainerChangeLogResponseData) SetLog(v []ContainerBaseChangeLogResponse) {
	o.Log = v
}

func (o ContainerChangeLogResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !isNil(o.Log) {
		toSerialize["log"] = o.Log
	}
	return json.Marshal(toSerialize)
}

type NullableContainerChangeLogResponseData struct {
	value *ContainerChangeLogResponseData
	isSet bool
}

func (v NullableContainerChangeLogResponseData) Get() *ContainerChangeLogResponseData {
	return v.value
}

func (v *NullableContainerChangeLogResponseData) Set(val *ContainerChangeLogResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerChangeLogResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerChangeLogResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerChangeLogResponseData(val *ContainerChangeLogResponseData) *NullableContainerChangeLogResponseData {
	return &NullableContainerChangeLogResponseData{value: val, isSet: true}
}

func (v NullableContainerChangeLogResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerChangeLogResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


