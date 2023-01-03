/*
Authentication And Authorization (AAA) Service

This swagger documentation provides all AAA API details. AAA service provides authentication and authorization capabilities for users.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aaa_client

import (
	"encoding/json"
)

// SuccessResponseAllOfData struct for SuccessResponseAllOfData
type SuccessResponseAllOfData struct {
	// Unique ID of resource.
	Id *string `json:"id,omitempty"`
}

// NewSuccessResponseAllOfData instantiates a new SuccessResponseAllOfData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuccessResponseAllOfData() *SuccessResponseAllOfData {
	this := SuccessResponseAllOfData{}
	return &this
}

// NewSuccessResponseAllOfDataWithDefaults instantiates a new SuccessResponseAllOfData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuccessResponseAllOfDataWithDefaults() *SuccessResponseAllOfData {
	this := SuccessResponseAllOfData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SuccessResponseAllOfData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessResponseAllOfData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SuccessResponseAllOfData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SuccessResponseAllOfData) SetId(v string) {
	o.Id = &v
}

func (o SuccessResponseAllOfData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableSuccessResponseAllOfData struct {
	value *SuccessResponseAllOfData
	isSet bool
}

func (v NullableSuccessResponseAllOfData) Get() *SuccessResponseAllOfData {
	return v.value
}

func (v *NullableSuccessResponseAllOfData) Set(val *SuccessResponseAllOfData) {
	v.value = val
	v.isSet = true
}

func (v NullableSuccessResponseAllOfData) IsSet() bool {
	return v.isSet
}

func (v *NullableSuccessResponseAllOfData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuccessResponseAllOfData(val *SuccessResponseAllOfData) *NullableSuccessResponseAllOfData {
	return &NullableSuccessResponseAllOfData{value: val, isSet: true}
}

func (v NullableSuccessResponseAllOfData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuccessResponseAllOfData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


