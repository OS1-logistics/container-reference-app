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

// PostErrorResponse struct for PostErrorResponse
type PostErrorResponse struct {
	Error *Error `json:"error,omitempty"`
	Request *Request `json:"request,omitempty"`
}

// NewPostErrorResponse instantiates a new PostErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostErrorResponse() *PostErrorResponse {
	this := PostErrorResponse{}
	return &this
}

// NewPostErrorResponseWithDefaults instantiates a new PostErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostErrorResponseWithDefaults() *PostErrorResponse {
	this := PostErrorResponse{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *PostErrorResponse) GetError() Error {
	if o == nil || isNil(o.Error) {
		var ret Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostErrorResponse) GetErrorOk() (*Error, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *PostErrorResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given Error and assigns it to the Error field.
func (o *PostErrorResponse) SetError(v Error) {
	o.Error = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *PostErrorResponse) GetRequest() Request {
	if o == nil || isNil(o.Request) {
		var ret Request
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostErrorResponse) GetRequestOk() (*Request, bool) {
	if o == nil || isNil(o.Request) {
    return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *PostErrorResponse) HasRequest() bool {
	if o != nil && !isNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given Request and assigns it to the Request field.
func (o *PostErrorResponse) SetRequest(v Request) {
	o.Request = &v
}

func (o PostErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !isNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	return json.Marshal(toSerialize)
}

type NullablePostErrorResponse struct {
	value *PostErrorResponse
	isSet bool
}

func (v NullablePostErrorResponse) Get() *PostErrorResponse {
	return v.value
}

func (v *NullablePostErrorResponse) Set(val *PostErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePostErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePostErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostErrorResponse(val *PostErrorResponse) *NullablePostErrorResponse {
	return &NullablePostErrorResponse{value: val, isSet: true}
}

func (v NullablePostErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


