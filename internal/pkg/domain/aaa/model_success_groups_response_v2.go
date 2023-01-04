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

// SuccessGroupsResponseV2 struct for SuccessGroupsResponseV2
type SuccessGroupsResponseV2 struct {
	Error *Error `json:"error,omitempty"`
	Data *SuccessGroupsResponseV2Data `json:"data,omitempty"`
	Request *Request `json:"request,omitempty"`
}

// NewSuccessGroupsResponseV2 instantiates a new SuccessGroupsResponseV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuccessGroupsResponseV2() *SuccessGroupsResponseV2 {
	this := SuccessGroupsResponseV2{}
	return &this
}

// NewSuccessGroupsResponseV2WithDefaults instantiates a new SuccessGroupsResponseV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuccessGroupsResponseV2WithDefaults() *SuccessGroupsResponseV2 {
	this := SuccessGroupsResponseV2{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SuccessGroupsResponseV2) GetError() Error {
	if o == nil || isNil(o.Error) {
		var ret Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessGroupsResponseV2) GetErrorOk() (*Error, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SuccessGroupsResponseV2) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given Error and assigns it to the Error field.
func (o *SuccessGroupsResponseV2) SetError(v Error) {
	o.Error = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SuccessGroupsResponseV2) GetData() SuccessGroupsResponseV2Data {
	if o == nil || isNil(o.Data) {
		var ret SuccessGroupsResponseV2Data
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessGroupsResponseV2) GetDataOk() (*SuccessGroupsResponseV2Data, bool) {
	if o == nil || isNil(o.Data) {
    return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SuccessGroupsResponseV2) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given SuccessGroupsResponseV2Data and assigns it to the Data field.
func (o *SuccessGroupsResponseV2) SetData(v SuccessGroupsResponseV2Data) {
	o.Data = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *SuccessGroupsResponseV2) GetRequest() Request {
	if o == nil || isNil(o.Request) {
		var ret Request
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessGroupsResponseV2) GetRequestOk() (*Request, bool) {
	if o == nil || isNil(o.Request) {
    return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *SuccessGroupsResponseV2) HasRequest() bool {
	if o != nil && !isNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given Request and assigns it to the Request field.
func (o *SuccessGroupsResponseV2) SetRequest(v Request) {
	o.Request = &v
}

func (o SuccessGroupsResponseV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !isNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	return json.Marshal(toSerialize)
}

type NullableSuccessGroupsResponseV2 struct {
	value *SuccessGroupsResponseV2
	isSet bool
}

func (v NullableSuccessGroupsResponseV2) Get() *SuccessGroupsResponseV2 {
	return v.value
}

func (v *NullableSuccessGroupsResponseV2) Set(val *SuccessGroupsResponseV2) {
	v.value = val
	v.isSet = true
}

func (v NullableSuccessGroupsResponseV2) IsSet() bool {
	return v.isSet
}

func (v *NullableSuccessGroupsResponseV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuccessGroupsResponseV2(val *SuccessGroupsResponseV2) *NullableSuccessGroupsResponseV2 {
	return &NullableSuccessGroupsResponseV2{value: val, isSet: true}
}

func (v NullableSuccessGroupsResponseV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuccessGroupsResponseV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


