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

// SuccessUserInfoResponseAllOf struct for SuccessUserInfoResponseAllOf
type SuccessUserInfoResponseAllOf struct {
	Data *UserInfo `json:"data,omitempty"`
	Request *Request `json:"request,omitempty"`
}

// NewSuccessUserInfoResponseAllOf instantiates a new SuccessUserInfoResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuccessUserInfoResponseAllOf() *SuccessUserInfoResponseAllOf {
	this := SuccessUserInfoResponseAllOf{}
	return &this
}

// NewSuccessUserInfoResponseAllOfWithDefaults instantiates a new SuccessUserInfoResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuccessUserInfoResponseAllOfWithDefaults() *SuccessUserInfoResponseAllOf {
	this := SuccessUserInfoResponseAllOf{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SuccessUserInfoResponseAllOf) GetData() UserInfo {
	if o == nil || isNil(o.Data) {
		var ret UserInfo
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessUserInfoResponseAllOf) GetDataOk() (*UserInfo, bool) {
	if o == nil || isNil(o.Data) {
    return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SuccessUserInfoResponseAllOf) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given UserInfo and assigns it to the Data field.
func (o *SuccessUserInfoResponseAllOf) SetData(v UserInfo) {
	o.Data = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *SuccessUserInfoResponseAllOf) GetRequest() Request {
	if o == nil || isNil(o.Request) {
		var ret Request
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessUserInfoResponseAllOf) GetRequestOk() (*Request, bool) {
	if o == nil || isNil(o.Request) {
    return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *SuccessUserInfoResponseAllOf) HasRequest() bool {
	if o != nil && !isNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given Request and assigns it to the Request field.
func (o *SuccessUserInfoResponseAllOf) SetRequest(v Request) {
	o.Request = &v
}

func (o SuccessUserInfoResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !isNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	return json.Marshal(toSerialize)
}

type NullableSuccessUserInfoResponseAllOf struct {
	value *SuccessUserInfoResponseAllOf
	isSet bool
}

func (v NullableSuccessUserInfoResponseAllOf) Get() *SuccessUserInfoResponseAllOf {
	return v.value
}

func (v *NullableSuccessUserInfoResponseAllOf) Set(val *SuccessUserInfoResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSuccessUserInfoResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSuccessUserInfoResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuccessUserInfoResponseAllOf(val *SuccessUserInfoResponseAllOf) *NullableSuccessUserInfoResponseAllOf {
	return &NullableSuccessUserInfoResponseAllOf{value: val, isSet: true}
}

func (v NullableSuccessUserInfoResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuccessUserInfoResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

