/*
Container Service

**API documentation for Container Service**

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package containerdomain

import (
	"encoding/json"
)

// State struct for State
type State struct {
	// Name of the state
	Name string `json:"name"`
	// Represents default substate for this state
	DefaultSubstate string `json:"defaultSubstate"`
	SubStates []SubState `json:"subStates"`
}

// NewState instantiates a new State object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewState(name string, defaultSubstate string, subStates []SubState) *State {
	this := State{}
	this.Name = name
	this.DefaultSubstate = defaultSubstate
	this.SubStates = subStates
	return &this
}

// NewStateWithDefaults instantiates a new State object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStateWithDefaults() *State {
	this := State{}
	return &this
}

// GetName returns the Name field value
func (o *State) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *State) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *State) SetName(v string) {
	o.Name = v
}

// GetDefaultSubstate returns the DefaultSubstate field value
func (o *State) GetDefaultSubstate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultSubstate
}

// GetDefaultSubstateOk returns a tuple with the DefaultSubstate field value
// and a boolean to check if the value has been set.
func (o *State) GetDefaultSubstateOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DefaultSubstate, true
}

// SetDefaultSubstate sets field value
func (o *State) SetDefaultSubstate(v string) {
	o.DefaultSubstate = v
}

// GetSubStates returns the SubStates field value
func (o *State) GetSubStates() []SubState {
	if o == nil {
		var ret []SubState
		return ret
	}

	return o.SubStates
}

// GetSubStatesOk returns a tuple with the SubStates field value
// and a boolean to check if the value has been set.
func (o *State) GetSubStatesOk() ([]SubState, bool) {
	if o == nil {
    return nil, false
	}
	return o.SubStates, true
}

// SetSubStates sets field value
func (o *State) SetSubStates(v []SubState) {
	o.SubStates = v
}

func (o State) MarshalJSON() ([]byte, error) {
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

type NullableState struct {
	value *State
	isSet bool
}

func (v NullableState) Get() *State {
	return v.value
}

func (v *NullableState) Set(val *State) {
	v.value = val
	v.isSet = true
}

func (v NullableState) IsSet() bool {
	return v.isSet
}

func (v *NullableState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableState(val *State) *NullableState {
	return &NullableState{value: val, isSet: true}
}

func (v NullableState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

