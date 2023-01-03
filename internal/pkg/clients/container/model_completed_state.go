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

// CompletedState End of the lifecycle of a container.
type CompletedState struct {
	// Name of the state
	Name string `json:"name"`
	// Represents default substate for this state
	DefaultSubstate string `json:"defaultSubstate"`
	SubStates []SubState `json:"subStates"`
	// terminal-state
	TerminalStates []string `json:"terminalStates,omitempty"`
}

// NewCompletedState instantiates a new CompletedState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompletedState(name string, defaultSubstate string, subStates []SubState) *CompletedState {
	this := CompletedState{}
	this.Name = name
	this.DefaultSubstate = defaultSubstate
	this.SubStates = subStates
	return &this
}

// NewCompletedStateWithDefaults instantiates a new CompletedState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompletedStateWithDefaults() *CompletedState {
	this := CompletedState{}
	return &this
}

// GetName returns the Name field value
func (o *CompletedState) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CompletedState) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CompletedState) SetName(v string) {
	o.Name = v
}

// GetDefaultSubstate returns the DefaultSubstate field value
func (o *CompletedState) GetDefaultSubstate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultSubstate
}

// GetDefaultSubstateOk returns a tuple with the DefaultSubstate field value
// and a boolean to check if the value has been set.
func (o *CompletedState) GetDefaultSubstateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultSubstate, true
}

// SetDefaultSubstate sets field value
func (o *CompletedState) SetDefaultSubstate(v string) {
	o.DefaultSubstate = v
}

// GetSubStates returns the SubStates field value
func (o *CompletedState) GetSubStates() []SubState {
	if o == nil {
		var ret []SubState
		return ret
	}

	return o.SubStates
}

// GetSubStatesOk returns a tuple with the SubStates field value
// and a boolean to check if the value has been set.
func (o *CompletedState) GetSubStatesOk() ([]SubState, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubStates, true
}

// SetSubStates sets field value
func (o *CompletedState) SetSubStates(v []SubState) {
	o.SubStates = v
}

// GetTerminalStates returns the TerminalStates field value if set, zero value otherwise.
func (o *CompletedState) GetTerminalStates() []string {
	if o == nil || o.TerminalStates == nil {
		var ret []string
		return ret
	}
	return o.TerminalStates
}

// GetTerminalStatesOk returns a tuple with the TerminalStates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompletedState) GetTerminalStatesOk() ([]string, bool) {
	if o == nil || o.TerminalStates == nil {
		return nil, false
	}
	return o.TerminalStates, true
}

// HasTerminalStates returns a boolean if a field has been set.
func (o *CompletedState) HasTerminalStates() bool {
	if o != nil && o.TerminalStates != nil {
		return true
	}

	return false
}

// SetTerminalStates gets a reference to the given []string and assigns it to the TerminalStates field.
func (o *CompletedState) SetTerminalStates(v []string) {
	o.TerminalStates = v
}

func (o CompletedState) MarshalJSON() ([]byte, error) {
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
	if o.TerminalStates != nil {
		toSerialize["terminalStates"] = o.TerminalStates
	}
	return json.Marshal(toSerialize)
}

type NullableCompletedState struct {
	value *CompletedState
	isSet bool
}

func (v NullableCompletedState) Get() *CompletedState {
	return v.value
}

func (v *NullableCompletedState) Set(val *CompletedState) {
	v.value = val
	v.isSet = true
}

func (v NullableCompletedState) IsSet() bool {
	return v.isSet
}

func (v *NullableCompletedState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompletedState(val *CompletedState) *NullableCompletedState {
	return &NullableCompletedState{value: val, isSet: true}
}

func (v NullableCompletedState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompletedState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


