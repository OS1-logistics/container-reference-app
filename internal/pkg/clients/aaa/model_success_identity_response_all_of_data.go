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

// SuccessIdentityResponseAllOfData struct for SuccessIdentityResponseAllOfData
type SuccessIdentityResponseAllOfData struct {
	// Login URL for the identity login flows.
	LoginUrl *string `json:"loginUrl,omitempty"`
}

// NewSuccessIdentityResponseAllOfData instantiates a new SuccessIdentityResponseAllOfData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuccessIdentityResponseAllOfData() *SuccessIdentityResponseAllOfData {
	this := SuccessIdentityResponseAllOfData{}
	return &this
}

// NewSuccessIdentityResponseAllOfDataWithDefaults instantiates a new SuccessIdentityResponseAllOfData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuccessIdentityResponseAllOfDataWithDefaults() *SuccessIdentityResponseAllOfData {
	this := SuccessIdentityResponseAllOfData{}
	return &this
}

// GetLoginUrl returns the LoginUrl field value if set, zero value otherwise.
func (o *SuccessIdentityResponseAllOfData) GetLoginUrl() string {
	if o == nil || o.LoginUrl == nil {
		var ret string
		return ret
	}
	return *o.LoginUrl
}

// GetLoginUrlOk returns a tuple with the LoginUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessIdentityResponseAllOfData) GetLoginUrlOk() (*string, bool) {
	if o == nil || o.LoginUrl == nil {
		return nil, false
	}
	return o.LoginUrl, true
}

// HasLoginUrl returns a boolean if a field has been set.
func (o *SuccessIdentityResponseAllOfData) HasLoginUrl() bool {
	if o != nil && o.LoginUrl != nil {
		return true
	}

	return false
}

// SetLoginUrl gets a reference to the given string and assigns it to the LoginUrl field.
func (o *SuccessIdentityResponseAllOfData) SetLoginUrl(v string) {
	o.LoginUrl = &v
}

func (o SuccessIdentityResponseAllOfData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LoginUrl != nil {
		toSerialize["loginUrl"] = o.LoginUrl
	}
	return json.Marshal(toSerialize)
}

type NullableSuccessIdentityResponseAllOfData struct {
	value *SuccessIdentityResponseAllOfData
	isSet bool
}

func (v NullableSuccessIdentityResponseAllOfData) Get() *SuccessIdentityResponseAllOfData {
	return v.value
}

func (v *NullableSuccessIdentityResponseAllOfData) Set(val *SuccessIdentityResponseAllOfData) {
	v.value = val
	v.isSet = true
}

func (v NullableSuccessIdentityResponseAllOfData) IsSet() bool {
	return v.isSet
}

func (v *NullableSuccessIdentityResponseAllOfData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuccessIdentityResponseAllOfData(val *SuccessIdentityResponseAllOfData) *NullableSuccessIdentityResponseAllOfData {
	return &NullableSuccessIdentityResponseAllOfData{value: val, isSet: true}
}

func (v NullableSuccessIdentityResponseAllOfData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuccessIdentityResponseAllOfData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


