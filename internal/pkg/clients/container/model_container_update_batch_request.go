/*
Container Service

**API documentation for Container Service**

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package container_client

import (
	"encoding/json"
)

// ContainerUpdateBatchRequest Object containing payload and call back url for container update request
type ContainerUpdateBatchRequest struct {
	// Array of container update request
	Payload []ContainerBatchUpdate `json:"payload,omitempty"`
	Callback *BatchCallbackURL `json:"callback,omitempty"`
}

// NewContainerUpdateBatchRequest instantiates a new ContainerUpdateBatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerUpdateBatchRequest() *ContainerUpdateBatchRequest {
	this := ContainerUpdateBatchRequest{}
	return &this
}

// NewContainerUpdateBatchRequestWithDefaults instantiates a new ContainerUpdateBatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerUpdateBatchRequestWithDefaults() *ContainerUpdateBatchRequest {
	this := ContainerUpdateBatchRequest{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *ContainerUpdateBatchRequest) GetPayload() []ContainerBatchUpdate {
	if o == nil || o.Payload == nil {
		var ret []ContainerBatchUpdate
		return ret
	}
	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerUpdateBatchRequest) GetPayloadOk() ([]ContainerBatchUpdate, bool) {
	if o == nil || o.Payload == nil {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *ContainerUpdateBatchRequest) HasPayload() bool {
	if o != nil && o.Payload != nil {
		return true
	}

	return false
}

// SetPayload gets a reference to the given []ContainerBatchUpdate and assigns it to the Payload field.
func (o *ContainerUpdateBatchRequest) SetPayload(v []ContainerBatchUpdate) {
	o.Payload = v
}

// GetCallback returns the Callback field value if set, zero value otherwise.
func (o *ContainerUpdateBatchRequest) GetCallback() BatchCallbackURL {
	if o == nil || o.Callback == nil {
		var ret BatchCallbackURL
		return ret
	}
	return *o.Callback
}

// GetCallbackOk returns a tuple with the Callback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerUpdateBatchRequest) GetCallbackOk() (*BatchCallbackURL, bool) {
	if o == nil || o.Callback == nil {
		return nil, false
	}
	return o.Callback, true
}

// HasCallback returns a boolean if a field has been set.
func (o *ContainerUpdateBatchRequest) HasCallback() bool {
	if o != nil && o.Callback != nil {
		return true
	}

	return false
}

// SetCallback gets a reference to the given BatchCallbackURL and assigns it to the Callback field.
func (o *ContainerUpdateBatchRequest) SetCallback(v BatchCallbackURL) {
	o.Callback = &v
}

func (o ContainerUpdateBatchRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Payload != nil {
		toSerialize["payload"] = o.Payload
	}
	if o.Callback != nil {
		toSerialize["callback"] = o.Callback
	}
	return json.Marshal(toSerialize)
}

type NullableContainerUpdateBatchRequest struct {
	value *ContainerUpdateBatchRequest
	isSet bool
}

func (v NullableContainerUpdateBatchRequest) Get() *ContainerUpdateBatchRequest {
	return v.value
}

func (v *NullableContainerUpdateBatchRequest) Set(val *ContainerUpdateBatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerUpdateBatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerUpdateBatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerUpdateBatchRequest(val *ContainerUpdateBatchRequest) *NullableContainerUpdateBatchRequest {
	return &NullableContainerUpdateBatchRequest{value: val, isSet: true}
}

func (v NullableContainerUpdateBatchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerUpdateBatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


