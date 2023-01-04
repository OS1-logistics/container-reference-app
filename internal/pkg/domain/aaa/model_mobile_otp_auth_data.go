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

// MobileOTPAuthData struct for MobileOTPAuthData
type MobileOTPAuthData struct {
	Mobile *MobileNumber `json:"mobile,omitempty"`
	Otp *string `json:"otp,omitempty"`
	Audience string `json:"audience"`
}

// NewMobileOTPAuthData instantiates a new MobileOTPAuthData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMobileOTPAuthData(audience string) *MobileOTPAuthData {
	this := MobileOTPAuthData{}
	this.Audience = audience
	return &this
}

// NewMobileOTPAuthDataWithDefaults instantiates a new MobileOTPAuthData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMobileOTPAuthDataWithDefaults() *MobileOTPAuthData {
	this := MobileOTPAuthData{}
	return &this
}

// GetMobile returns the Mobile field value if set, zero value otherwise.
func (o *MobileOTPAuthData) GetMobile() MobileNumber {
	if o == nil || o.Mobile == nil {
		var ret MobileNumber
		return ret
	}
	return *o.Mobile
}

// GetMobileOk returns a tuple with the Mobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileOTPAuthData) GetMobileOk() (*MobileNumber, bool) {
	if o == nil || o.Mobile == nil {
		return nil, false
	}
	return o.Mobile, true
}

// HasMobile returns a boolean if a field has been set.
func (o *MobileOTPAuthData) HasMobile() bool {
	if o != nil && o.Mobile != nil {
		return true
	}

	return false
}

// SetMobile gets a reference to the given MobileNumber and assigns it to the Mobile field.
func (o *MobileOTPAuthData) SetMobile(v MobileNumber) {
	o.Mobile = &v
}

// GetOtp returns the Otp field value if set, zero value otherwise.
func (o *MobileOTPAuthData) GetOtp() string {
	if o == nil || o.Otp == nil {
		var ret string
		return ret
	}
	return *o.Otp
}

// GetOtpOk returns a tuple with the Otp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileOTPAuthData) GetOtpOk() (*string, bool) {
	if o == nil || o.Otp == nil {
		return nil, false
	}
	return o.Otp, true
}

// HasOtp returns a boolean if a field has been set.
func (o *MobileOTPAuthData) HasOtp() bool {
	if o != nil && o.Otp != nil {
		return true
	}

	return false
}

// SetOtp gets a reference to the given string and assigns it to the Otp field.
func (o *MobileOTPAuthData) SetOtp(v string) {
	o.Otp = &v
}

// GetAudience returns the Audience field value
func (o *MobileOTPAuthData) GetAudience() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Audience
}

// GetAudienceOk returns a tuple with the Audience field value
// and a boolean to check if the value has been set.
func (o *MobileOTPAuthData) GetAudienceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Audience, true
}

// SetAudience sets field value
func (o *MobileOTPAuthData) SetAudience(v string) {
	o.Audience = v
}

func (o MobileOTPAuthData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Mobile != nil {
		toSerialize["mobile"] = o.Mobile
	}
	if o.Otp != nil {
		toSerialize["otp"] = o.Otp
	}
	if true {
		toSerialize["audience"] = o.Audience
	}
	return json.Marshal(toSerialize)
}

type NullableMobileOTPAuthData struct {
	value *MobileOTPAuthData
	isSet bool
}

func (v NullableMobileOTPAuthData) Get() *MobileOTPAuthData {
	return v.value
}

func (v *NullableMobileOTPAuthData) Set(val *MobileOTPAuthData) {
	v.value = val
	v.isSet = true
}

func (v NullableMobileOTPAuthData) IsSet() bool {
	return v.isSet
}

func (v *NullableMobileOTPAuthData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMobileOTPAuthData(val *MobileOTPAuthData) *NullableMobileOTPAuthData {
	return &NullableMobileOTPAuthData{value: val, isSet: true}
}

func (v NullableMobileOTPAuthData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMobileOTPAuthData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


