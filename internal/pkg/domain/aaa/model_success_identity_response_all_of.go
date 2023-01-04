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

// SuccessIdentityResponseAllOf struct for SuccessIdentityResponseAllOf
type SuccessIdentityResponseAllOf struct {
	Data *SuccessIdentityResponseAllOfData `json:"data,omitempty"`
	Request *Request `json:"request,omitempty"`
}

// NewSuccessIdentityResponseAllOf instantiates a new SuccessIdentityResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuccessIdentityResponseAllOf() *SuccessIdentityResponseAllOf {
	this := SuccessIdentityResponseAllOf{}
	return &this
}

// NewSuccessIdentityResponseAllOfWithDefaults instantiates a new SuccessIdentityResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuccessIdentityResponseAllOfWithDefaults() *SuccessIdentityResponseAllOf {
	this := SuccessIdentityResponseAllOf{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SuccessIdentityResponseAllOf) GetData() SuccessIdentityResponseAllOfData {
	if o == nil || o.Data == nil {
		var ret SuccessIdentityResponseAllOfData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessIdentityResponseAllOf) GetDataOk() (*SuccessIdentityResponseAllOfData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SuccessIdentityResponseAllOf) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given SuccessIdentityResponseAllOfData and assigns it to the Data field.
func (o *SuccessIdentityResponseAllOf) SetData(v SuccessIdentityResponseAllOfData) {
	o.Data = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *SuccessIdentityResponseAllOf) GetRequest() Request {
	if o == nil || o.Request == nil {
		var ret Request
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessIdentityResponseAllOf) GetRequestOk() (*Request, bool) {
	if o == nil || o.Request == nil {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *SuccessIdentityResponseAllOf) HasRequest() bool {
	if o != nil && o.Request != nil {
		return true
	}

	return false
}

// SetRequest gets a reference to the given Request and assigns it to the Request field.
func (o *SuccessIdentityResponseAllOf) SetRequest(v Request) {
	o.Request = &v
}

func (o SuccessIdentityResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Request != nil {
		toSerialize["request"] = o.Request
	}
	return json.Marshal(toSerialize)
}

type NullableSuccessIdentityResponseAllOf struct {
	value *SuccessIdentityResponseAllOf
	isSet bool
}

func (v NullableSuccessIdentityResponseAllOf) Get() *SuccessIdentityResponseAllOf {
	return v.value
}

func (v *NullableSuccessIdentityResponseAllOf) Set(val *SuccessIdentityResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSuccessIdentityResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSuccessIdentityResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuccessIdentityResponseAllOf(val *SuccessIdentityResponseAllOf) *NullableSuccessIdentityResponseAllOf {
	return &NullableSuccessIdentityResponseAllOf{value: val, isSet: true}
}

func (v NullableSuccessIdentityResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuccessIdentityResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


