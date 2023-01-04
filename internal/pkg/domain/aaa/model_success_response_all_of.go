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

// SuccessResponseAllOf struct for SuccessResponseAllOf
type SuccessResponseAllOf struct {
	Data *SuccessResponseAllOfData `json:"data,omitempty"`
	Request *Request `json:"request,omitempty"`
}

// NewSuccessResponseAllOf instantiates a new SuccessResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuccessResponseAllOf() *SuccessResponseAllOf {
	this := SuccessResponseAllOf{}
	return &this
}

// NewSuccessResponseAllOfWithDefaults instantiates a new SuccessResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuccessResponseAllOfWithDefaults() *SuccessResponseAllOf {
	this := SuccessResponseAllOf{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SuccessResponseAllOf) GetData() SuccessResponseAllOfData {
	if o == nil || isNil(o.Data) {
		var ret SuccessResponseAllOfData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessResponseAllOf) GetDataOk() (*SuccessResponseAllOfData, bool) {
	if o == nil || isNil(o.Data) {
    return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SuccessResponseAllOf) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given SuccessResponseAllOfData and assigns it to the Data field.
func (o *SuccessResponseAllOf) SetData(v SuccessResponseAllOfData) {
	o.Data = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *SuccessResponseAllOf) GetRequest() Request {
	if o == nil || isNil(o.Request) {
		var ret Request
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessResponseAllOf) GetRequestOk() (*Request, bool) {
	if o == nil || isNil(o.Request) {
    return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *SuccessResponseAllOf) HasRequest() bool {
	if o != nil && !isNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given Request and assigns it to the Request field.
func (o *SuccessResponseAllOf) SetRequest(v Request) {
	o.Request = &v
}

func (o SuccessResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !isNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	return json.Marshal(toSerialize)
}

type NullableSuccessResponseAllOf struct {
	value *SuccessResponseAllOf
	isSet bool
}

func (v NullableSuccessResponseAllOf) Get() *SuccessResponseAllOf {
	return v.value
}

func (v *NullableSuccessResponseAllOf) Set(val *SuccessResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSuccessResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSuccessResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuccessResponseAllOf(val *SuccessResponseAllOf) *NullableSuccessResponseAllOf {
	return &NullableSuccessResponseAllOf{value: val, isSet: true}
}

func (v NullableSuccessResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuccessResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


