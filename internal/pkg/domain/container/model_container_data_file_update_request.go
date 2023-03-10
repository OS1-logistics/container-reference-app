/*
Container Service

**API documentation for Container Service**

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package containerdomain

import (
	"encoding/json"
	"os"
)

// ContainerDataFileUpdateRequest struct for ContainerDataFileUpdateRequest
type ContainerDataFileUpdateRequest struct {
	File *os.File `json:"file"`
	Callback *ContainerTypeUpdateRequestCallback `json:"callback,omitempty"`
}

// NewContainerDataFileUpdateRequest instantiates a new ContainerDataFileUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerDataFileUpdateRequest(file *os.File) *ContainerDataFileUpdateRequest {
	this := ContainerDataFileUpdateRequest{}
	this.File = file
	return &this
}

// NewContainerDataFileUpdateRequestWithDefaults instantiates a new ContainerDataFileUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerDataFileUpdateRequestWithDefaults() *ContainerDataFileUpdateRequest {
	this := ContainerDataFileUpdateRequest{}
	return &this
}

// GetFile returns the File field value
func (o *ContainerDataFileUpdateRequest) GetFile() *os.File {
	if o == nil {
		var ret *os.File
		return ret
	}

	return o.File
}

// GetFileOk returns a tuple with the File field value
// and a boolean to check if the value has been set.
func (o *ContainerDataFileUpdateRequest) GetFileOk() (**os.File, bool) {
	if o == nil {
    return nil, false
	}
	return &o.File, true
}

// SetFile sets field value
func (o *ContainerDataFileUpdateRequest) SetFile(v *os.File) {
	o.File = v
}

// GetCallback returns the Callback field value if set, zero value otherwise.
func (o *ContainerDataFileUpdateRequest) GetCallback() ContainerTypeUpdateRequestCallback {
	if o == nil || isNil(o.Callback) {
		var ret ContainerTypeUpdateRequestCallback
		return ret
	}
	return *o.Callback
}

// GetCallbackOk returns a tuple with the Callback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerDataFileUpdateRequest) GetCallbackOk() (*ContainerTypeUpdateRequestCallback, bool) {
	if o == nil || isNil(o.Callback) {
    return nil, false
	}
	return o.Callback, true
}

// HasCallback returns a boolean if a field has been set.
func (o *ContainerDataFileUpdateRequest) HasCallback() bool {
	if o != nil && !isNil(o.Callback) {
		return true
	}

	return false
}

// SetCallback gets a reference to the given ContainerTypeUpdateRequestCallback and assigns it to the Callback field.
func (o *ContainerDataFileUpdateRequest) SetCallback(v ContainerTypeUpdateRequestCallback) {
	o.Callback = &v
}

func (o ContainerDataFileUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["file"] = o.File
	}
	if !isNil(o.Callback) {
		toSerialize["callback"] = o.Callback
	}
	return json.Marshal(toSerialize)
}

type NullableContainerDataFileUpdateRequest struct {
	value *ContainerDataFileUpdateRequest
	isSet bool
}

func (v NullableContainerDataFileUpdateRequest) Get() *ContainerDataFileUpdateRequest {
	return v.value
}

func (v *NullableContainerDataFileUpdateRequest) Set(val *ContainerDataFileUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerDataFileUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerDataFileUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerDataFileUpdateRequest(val *ContainerDataFileUpdateRequest) *NullableContainerDataFileUpdateRequest {
	return &NullableContainerDataFileUpdateRequest{value: val, isSet: true}
}

func (v NullableContainerDataFileUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerDataFileUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


