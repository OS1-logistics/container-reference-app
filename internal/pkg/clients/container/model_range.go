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

// Range Field to specify the range that the attribute value must lie. In case of \"string\" datatype, range will be the length of string value. this validation allowed only for string and number type of attributes.
type Range struct {
	// The min value, inclusive.
	Min *float32 `json:"min,omitempty"`
	// The max value, inclusive.
	Max *float32 `json:"max,omitempty"`
}

// NewRange instantiates a new Range object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRange() *Range {
	this := Range{}
	return &this
}

// NewRangeWithDefaults instantiates a new Range object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRangeWithDefaults() *Range {
	this := Range{}
	return &this
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *Range) GetMin() float32 {
	if o == nil || o.Min == nil {
		var ret float32
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Range) GetMinOk() (*float32, bool) {
	if o == nil || o.Min == nil {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *Range) HasMin() bool {
	if o != nil && o.Min != nil {
		return true
	}

	return false
}

// SetMin gets a reference to the given float32 and assigns it to the Min field.
func (o *Range) SetMin(v float32) {
	o.Min = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *Range) GetMax() float32 {
	if o == nil || o.Max == nil {
		var ret float32
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Range) GetMaxOk() (*float32, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *Range) HasMax() bool {
	if o != nil && o.Max != nil {
		return true
	}

	return false
}

// SetMax gets a reference to the given float32 and assigns it to the Max field.
func (o *Range) SetMax(v float32) {
	o.Max = &v
}

func (o Range) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Min != nil {
		toSerialize["min"] = o.Min
	}
	if o.Max != nil {
		toSerialize["max"] = o.Max
	}
	return json.Marshal(toSerialize)
}

type NullableRange struct {
	value *Range
	isSet bool
}

func (v NullableRange) Get() *Range {
	return v.value
}

func (v *NullableRange) Set(val *Range) {
	v.value = val
	v.isSet = true
}

func (v NullableRange) IsSet() bool {
	return v.isSet
}

func (v *NullableRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRange(val *Range) *NullableRange {
	return &NullableRange{value: val, isSet: true}
}

func (v NullableRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


