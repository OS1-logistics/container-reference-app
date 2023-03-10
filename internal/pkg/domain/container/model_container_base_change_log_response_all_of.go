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

// ContainerBaseChangeLogResponseAllOf struct for ContainerBaseChangeLogResponseAllOf
type ContainerBaseChangeLogResponseAllOf struct {
	Cdc ContainerCDC `json:"cdc"`
}

// NewContainerBaseChangeLogResponseAllOf instantiates a new ContainerBaseChangeLogResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerBaseChangeLogResponseAllOf(cdc ContainerCDC) *ContainerBaseChangeLogResponseAllOf {
	this := ContainerBaseChangeLogResponseAllOf{}
	this.Cdc = cdc
	return &this
}

// NewContainerBaseChangeLogResponseAllOfWithDefaults instantiates a new ContainerBaseChangeLogResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerBaseChangeLogResponseAllOfWithDefaults() *ContainerBaseChangeLogResponseAllOf {
	this := ContainerBaseChangeLogResponseAllOf{}
	return &this
}

// GetCdc returns the Cdc field value
func (o *ContainerBaseChangeLogResponseAllOf) GetCdc() ContainerCDC {
	if o == nil {
		var ret ContainerCDC
		return ret
	}

	return o.Cdc
}

// GetCdcOk returns a tuple with the Cdc field value
// and a boolean to check if the value has been set.
func (o *ContainerBaseChangeLogResponseAllOf) GetCdcOk() (*ContainerCDC, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Cdc, true
}

// SetCdc sets field value
func (o *ContainerBaseChangeLogResponseAllOf) SetCdc(v ContainerCDC) {
	o.Cdc = v
}

func (o ContainerBaseChangeLogResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cdc"] = o.Cdc
	}
	return json.Marshal(toSerialize)
}

type NullableContainerBaseChangeLogResponseAllOf struct {
	value *ContainerBaseChangeLogResponseAllOf
	isSet bool
}

func (v NullableContainerBaseChangeLogResponseAllOf) Get() *ContainerBaseChangeLogResponseAllOf {
	return v.value
}

func (v *NullableContainerBaseChangeLogResponseAllOf) Set(val *ContainerBaseChangeLogResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerBaseChangeLogResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerBaseChangeLogResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerBaseChangeLogResponseAllOf(val *ContainerBaseChangeLogResponseAllOf) *NullableContainerBaseChangeLogResponseAllOf {
	return &NullableContainerBaseChangeLogResponseAllOf{value: val, isSet: true}
}

func (v NullableContainerBaseChangeLogResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerBaseChangeLogResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


