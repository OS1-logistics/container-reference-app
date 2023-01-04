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

// MobileNumber struct for MobileNumber
type MobileNumber struct {
	// Country code for mobile number.
	CountryCode string `json:"countryCode"`
	// Mobile number. In `Get` calls, mobile number is masked excluding last 4 digits of number.
	Number string `json:"number"`
}

// NewMobileNumber instantiates a new MobileNumber object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMobileNumber(countryCode string, number string) *MobileNumber {
	this := MobileNumber{}
	this.CountryCode = countryCode
	this.Number = number
	return &this
}

// NewMobileNumberWithDefaults instantiates a new MobileNumber object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMobileNumberWithDefaults() *MobileNumber {
	this := MobileNumber{}
	return &this
}

// GetCountryCode returns the CountryCode field value
func (o *MobileNumber) GetCountryCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value
// and a boolean to check if the value has been set.
func (o *MobileNumber) GetCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryCode, true
}

// SetCountryCode sets field value
func (o *MobileNumber) SetCountryCode(v string) {
	o.CountryCode = v
}

// GetNumber returns the Number field value
func (o *MobileNumber) GetNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *MobileNumber) GetNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *MobileNumber) SetNumber(v string) {
	o.Number = v
}

func (o MobileNumber) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["countryCode"] = o.CountryCode
	}
	if true {
		toSerialize["number"] = o.Number
	}
	return json.Marshal(toSerialize)
}

type NullableMobileNumber struct {
	value *MobileNumber
	isSet bool
}

func (v NullableMobileNumber) Get() *MobileNumber {
	return v.value
}

func (v *NullableMobileNumber) Set(val *MobileNumber) {
	v.value = val
	v.isSet = true
}

func (v NullableMobileNumber) IsSet() bool {
	return v.isSet
}

func (v *NullableMobileNumber) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMobileNumber(val *MobileNumber) *NullableMobileNumber {
	return &NullableMobileNumber{value: val, isSet: true}
}

func (v NullableMobileNumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMobileNumber) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


