/*
Authentication And Authorization (AAA) Service

This swagger documentation provides all AAA API details. AAA service provides authentication and authorization capabilities for users.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aaadomain

import (
	"encoding/json"
	"os"
)

// MultipartFileUploadRequest struct for MultipartFileUploadRequest
type MultipartFileUploadRequest struct {
	File *os.File `json:"file"`
	// URL to notify the outcome of the request.
	Callback string `json:"callback"`
}

// NewMultipartFileUploadRequest instantiates a new MultipartFileUploadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultipartFileUploadRequest(file *os.File, callback string) *MultipartFileUploadRequest {
	this := MultipartFileUploadRequest{}
	this.File = file
	this.Callback = callback
	return &this
}

// NewMultipartFileUploadRequestWithDefaults instantiates a new MultipartFileUploadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultipartFileUploadRequestWithDefaults() *MultipartFileUploadRequest {
	this := MultipartFileUploadRequest{}
	return &this
}

// GetFile returns the File field value
func (o *MultipartFileUploadRequest) GetFile() *os.File {
	if o == nil {
		var ret *os.File
		return ret
	}

	return o.File
}

// GetFileOk returns a tuple with the File field value
// and a boolean to check if the value has been set.
func (o *MultipartFileUploadRequest) GetFileOk() (**os.File, bool) {
	if o == nil {
    return nil, false
	}
	return &o.File, true
}

// SetFile sets field value
func (o *MultipartFileUploadRequest) SetFile(v *os.File) {
	o.File = v
}

// GetCallback returns the Callback field value
func (o *MultipartFileUploadRequest) GetCallback() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Callback
}

// GetCallbackOk returns a tuple with the Callback field value
// and a boolean to check if the value has been set.
func (o *MultipartFileUploadRequest) GetCallbackOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Callback, true
}

// SetCallback sets field value
func (o *MultipartFileUploadRequest) SetCallback(v string) {
	o.Callback = v
}

func (o MultipartFileUploadRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["file"] = o.File
	}
	if true {
		toSerialize["callback"] = o.Callback
	}
	return json.Marshal(toSerialize)
}

type NullableMultipartFileUploadRequest struct {
	value *MultipartFileUploadRequest
	isSet bool
}

func (v NullableMultipartFileUploadRequest) Get() *MultipartFileUploadRequest {
	return v.value
}

func (v *NullableMultipartFileUploadRequest) Set(val *MultipartFileUploadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMultipartFileUploadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMultipartFileUploadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultipartFileUploadRequest(val *MultipartFileUploadRequest) *NullableMultipartFileUploadRequest {
	return &NullableMultipartFileUploadRequest{value: val, isSet: true}
}

func (v NullableMultipartFileUploadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultipartFileUploadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


