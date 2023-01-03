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

// PostErrorResponseAllOf struct for PostErrorResponseAllOf
type PostErrorResponseAllOf struct {
	Request *Request `json:"request,omitempty"`
}

// NewPostErrorResponseAllOf instantiates a new PostErrorResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostErrorResponseAllOf() *PostErrorResponseAllOf {
	this := PostErrorResponseAllOf{}
	return &this
}

// NewPostErrorResponseAllOfWithDefaults instantiates a new PostErrorResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostErrorResponseAllOfWithDefaults() *PostErrorResponseAllOf {
	this := PostErrorResponseAllOf{}
	return &this
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *PostErrorResponseAllOf) GetRequest() Request {
	if o == nil || o.Request == nil {
		var ret Request
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostErrorResponseAllOf) GetRequestOk() (*Request, bool) {
	if o == nil || o.Request == nil {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *PostErrorResponseAllOf) HasRequest() bool {
	if o != nil && o.Request != nil {
		return true
	}

	return false
}

// SetRequest gets a reference to the given Request and assigns it to the Request field.
func (o *PostErrorResponseAllOf) SetRequest(v Request) {
	o.Request = &v
}

func (o PostErrorResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Request != nil {
		toSerialize["request"] = o.Request
	}
	return json.Marshal(toSerialize)
}

type NullablePostErrorResponseAllOf struct {
	value *PostErrorResponseAllOf
	isSet bool
}

func (v NullablePostErrorResponseAllOf) Get() *PostErrorResponseAllOf {
	return v.value
}

func (v *NullablePostErrorResponseAllOf) Set(val *PostErrorResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePostErrorResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePostErrorResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostErrorResponseAllOf(val *PostErrorResponseAllOf) *NullablePostErrorResponseAllOf {
	return &NullablePostErrorResponseAllOf{value: val, isSet: true}
}

func (v NullablePostErrorResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostErrorResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


