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

// AuthReqOtpRequestMobile Mobile object has to be present if mode is SMS.
type AuthReqOtpRequestMobile struct {
	// Country code for mobile number.
	CountryCode *string `json:"countryCode,omitempty"`
	// Mobile number.
	Number *string `json:"number,omitempty"`
}

// NewAuthReqOtpRequestMobile instantiates a new AuthReqOtpRequestMobile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthReqOtpRequestMobile() *AuthReqOtpRequestMobile {
	this := AuthReqOtpRequestMobile{}
	return &this
}

// NewAuthReqOtpRequestMobileWithDefaults instantiates a new AuthReqOtpRequestMobile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthReqOtpRequestMobileWithDefaults() *AuthReqOtpRequestMobile {
	this := AuthReqOtpRequestMobile{}
	return &this
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *AuthReqOtpRequestMobile) GetCountryCode() string {
	if o == nil || o.CountryCode == nil {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthReqOtpRequestMobile) GetCountryCodeOk() (*string, bool) {
	if o == nil || o.CountryCode == nil {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *AuthReqOtpRequestMobile) HasCountryCode() bool {
	if o != nil && o.CountryCode != nil {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *AuthReqOtpRequestMobile) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *AuthReqOtpRequestMobile) GetNumber() string {
	if o == nil || o.Number == nil {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthReqOtpRequestMobile) GetNumberOk() (*string, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *AuthReqOtpRequestMobile) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *AuthReqOtpRequestMobile) SetNumber(v string) {
	o.Number = &v
}

func (o AuthReqOtpRequestMobile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CountryCode != nil {
		toSerialize["countryCode"] = o.CountryCode
	}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	return json.Marshal(toSerialize)
}

type NullableAuthReqOtpRequestMobile struct {
	value *AuthReqOtpRequestMobile
	isSet bool
}

func (v NullableAuthReqOtpRequestMobile) Get() *AuthReqOtpRequestMobile {
	return v.value
}

func (v *NullableAuthReqOtpRequestMobile) Set(val *AuthReqOtpRequestMobile) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthReqOtpRequestMobile) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthReqOtpRequestMobile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthReqOtpRequestMobile(val *AuthReqOtpRequestMobile) *NullableAuthReqOtpRequestMobile {
	return &NullableAuthReqOtpRequestMobile{value: val, isSet: true}
}

func (v NullableAuthReqOtpRequestMobile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthReqOtpRequestMobile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


