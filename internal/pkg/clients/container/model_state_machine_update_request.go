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

// StateMachineUpdateRequest Update state machine config
type StateMachineUpdateRequest struct {
	States States `json:"states"`
}

// NewStateMachineUpdateRequest instantiates a new StateMachineUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStateMachineUpdateRequest(states States) *StateMachineUpdateRequest {
	this := StateMachineUpdateRequest{}
	this.States = states
	return &this
}

// NewStateMachineUpdateRequestWithDefaults instantiates a new StateMachineUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStateMachineUpdateRequestWithDefaults() *StateMachineUpdateRequest {
	this := StateMachineUpdateRequest{}
	return &this
}

// GetStates returns the States field value
func (o *StateMachineUpdateRequest) GetStates() States {
	if o == nil {
		var ret States
		return ret
	}

	return o.States
}

// GetStatesOk returns a tuple with the States field value
// and a boolean to check if the value has been set.
func (o *StateMachineUpdateRequest) GetStatesOk() (*States, bool) {
	if o == nil {
		return nil, false
	}
	return &o.States, true
}

// SetStates sets field value
func (o *StateMachineUpdateRequest) SetStates(v States) {
	o.States = v
}

func (o StateMachineUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["states"] = o.States
	}
	return json.Marshal(toSerialize)
}

type NullableStateMachineUpdateRequest struct {
	value *StateMachineUpdateRequest
	isSet bool
}

func (v NullableStateMachineUpdateRequest) Get() *StateMachineUpdateRequest {
	return v.value
}

func (v *NullableStateMachineUpdateRequest) Set(val *StateMachineUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStateMachineUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStateMachineUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStateMachineUpdateRequest(val *StateMachineUpdateRequest) *NullableStateMachineUpdateRequest {
	return &NullableStateMachineUpdateRequest{value: val, isSet: true}
}

func (v NullableStateMachineUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStateMachineUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


