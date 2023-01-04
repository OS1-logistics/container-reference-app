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

// SuccessStatusResponseAllOfData struct for SuccessStatusResponseAllOfData
type SuccessStatusResponseAllOfData struct {
	// Status message of the operation.
	Status *string `json:"status,omitempty"`
}

// NewSuccessStatusResponseAllOfData instantiates a new SuccessStatusResponseAllOfData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuccessStatusResponseAllOfData() *SuccessStatusResponseAllOfData {
	this := SuccessStatusResponseAllOfData{}
	return &this
}

// NewSuccessStatusResponseAllOfDataWithDefaults instantiates a new SuccessStatusResponseAllOfData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuccessStatusResponseAllOfDataWithDefaults() *SuccessStatusResponseAllOfData {
	this := SuccessStatusResponseAllOfData{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SuccessStatusResponseAllOfData) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessStatusResponseAllOfData) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SuccessStatusResponseAllOfData) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SuccessStatusResponseAllOfData) SetStatus(v string) {
	o.Status = &v
}

func (o SuccessStatusResponseAllOfData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableSuccessStatusResponseAllOfData struct {
	value *SuccessStatusResponseAllOfData
	isSet bool
}

func (v NullableSuccessStatusResponseAllOfData) Get() *SuccessStatusResponseAllOfData {
	return v.value
}

func (v *NullableSuccessStatusResponseAllOfData) Set(val *SuccessStatusResponseAllOfData) {
	v.value = val
	v.isSet = true
}

func (v NullableSuccessStatusResponseAllOfData) IsSet() bool {
	return v.isSet
}

func (v *NullableSuccessStatusResponseAllOfData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuccessStatusResponseAllOfData(val *SuccessStatusResponseAllOfData) *NullableSuccessStatusResponseAllOfData {
	return &NullableSuccessStatusResponseAllOfData{value: val, isSet: true}
}

func (v NullableSuccessStatusResponseAllOfData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuccessStatusResponseAllOfData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


