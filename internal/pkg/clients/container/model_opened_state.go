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

// OpenedState It's signifies readiness of container to begin containerization process.
type OpenedState struct {
	// Name of the state
	Name string `json:"name"`
	// Represents default substate for this state
	DefaultSubstate string `json:"defaultSubstate"`
	SubStates []SubState `json:"subStates"`
}

// NewOpenedState instantiates a new OpenedState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenedState(name string, defaultSubstate string, subStates []SubState) *OpenedState {
	this := OpenedState{}
	this.Name = name
	this.DefaultSubstate = defaultSubstate
	this.SubStates = subStates
	return &this
}

// NewOpenedStateWithDefaults instantiates a new OpenedState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenedStateWithDefaults() *OpenedState {
	this := OpenedState{}
	return &this
}

// GetName returns the Name field value
func (o *OpenedState) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OpenedState) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OpenedState) SetName(v string) {
	o.Name = v
}

// GetDefaultSubstate returns the DefaultSubstate field value
func (o *OpenedState) GetDefaultSubstate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultSubstate
}

// GetDefaultSubstateOk returns a tuple with the DefaultSubstate field value
// and a boolean to check if the value has been set.
func (o *OpenedState) GetDefaultSubstateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultSubstate, true
}

// SetDefaultSubstate sets field value
func (o *OpenedState) SetDefaultSubstate(v string) {
	o.DefaultSubstate = v
}

// GetSubStates returns the SubStates field value
func (o *OpenedState) GetSubStates() []SubState {
	if o == nil {
		var ret []SubState
		return ret
	}

	return o.SubStates
}

// GetSubStatesOk returns a tuple with the SubStates field value
// and a boolean to check if the value has been set.
func (o *OpenedState) GetSubStatesOk() ([]SubState, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubStates, true
}

// SetSubStates sets field value
func (o *OpenedState) SetSubStates(v []SubState) {
	o.SubStates = v
}

func (o OpenedState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["defaultSubstate"] = o.DefaultSubstate
	}
	if true {
		toSerialize["subStates"] = o.SubStates
	}
	return json.Marshal(toSerialize)
}

type NullableOpenedState struct {
	value *OpenedState
	isSet bool
}

func (v NullableOpenedState) Get() *OpenedState {
	return v.value
}

func (v *NullableOpenedState) Set(val *OpenedState) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenedState) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenedState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenedState(val *OpenedState) *NullableOpenedState {
	return &NullableOpenedState{value: val, isSet: true}
}

func (v NullableOpenedState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenedState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


