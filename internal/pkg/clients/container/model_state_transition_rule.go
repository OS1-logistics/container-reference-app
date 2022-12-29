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

// StateTransitionRule State Transition Rule
type StateTransitionRule struct {
	// Event Code of the format:E-yyy. E - Prefix indicating that this is an event code.  yyy - 3 digit event code string ranging from 001 to 999
	EventCode string `json:"eventCode"`
	// Name of the main state and its substate of destination state
	Destination string `json:"destination"`
	// Reason Code of the format:R-nnnn R - Prefix indicating that this is a reason code nnnn - 4 digit reason code string ranging from 0001 to 9999
	ReasonCode *string `json:"reasonCode,omitempty"`
}

// NewStateTransitionRule instantiates a new StateTransitionRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStateTransitionRule(eventCode string, destination string) *StateTransitionRule {
	this := StateTransitionRule{}
	this.EventCode = eventCode
	this.Destination = destination
	return &this
}

// NewStateTransitionRuleWithDefaults instantiates a new StateTransitionRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStateTransitionRuleWithDefaults() *StateTransitionRule {
	this := StateTransitionRule{}
	return &this
}

// GetEventCode returns the EventCode field value
func (o *StateTransitionRule) GetEventCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventCode
}

// GetEventCodeOk returns a tuple with the EventCode field value
// and a boolean to check if the value has been set.
func (o *StateTransitionRule) GetEventCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventCode, true
}

// SetEventCode sets field value
func (o *StateTransitionRule) SetEventCode(v string) {
	o.EventCode = v
}

// GetDestination returns the Destination field value
func (o *StateTransitionRule) GetDestination() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *StateTransitionRule) GetDestinationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *StateTransitionRule) SetDestination(v string) {
	o.Destination = v
}

// GetReasonCode returns the ReasonCode field value if set, zero value otherwise.
func (o *StateTransitionRule) GetReasonCode() string {
	if o == nil || o.ReasonCode == nil {
		var ret string
		return ret
	}
	return *o.ReasonCode
}

// GetReasonCodeOk returns a tuple with the ReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StateTransitionRule) GetReasonCodeOk() (*string, bool) {
	if o == nil || o.ReasonCode == nil {
		return nil, false
	}
	return o.ReasonCode, true
}

// HasReasonCode returns a boolean if a field has been set.
func (o *StateTransitionRule) HasReasonCode() bool {
	if o != nil && o.ReasonCode != nil {
		return true
	}

	return false
}

// SetReasonCode gets a reference to the given string and assigns it to the ReasonCode field.
func (o *StateTransitionRule) SetReasonCode(v string) {
	o.ReasonCode = &v
}

func (o StateTransitionRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["eventCode"] = o.EventCode
	}
	if true {
		toSerialize["destination"] = o.Destination
	}
	if o.ReasonCode != nil {
		toSerialize["reasonCode"] = o.ReasonCode
	}
	return json.Marshal(toSerialize)
}

type NullableStateTransitionRule struct {
	value *StateTransitionRule
	isSet bool
}

func (v NullableStateTransitionRule) Get() *StateTransitionRule {
	return v.value
}

func (v *NullableStateTransitionRule) Set(val *StateTransitionRule) {
	v.value = val
	v.isSet = true
}

func (v NullableStateTransitionRule) IsSet() bool {
	return v.isSet
}

func (v *NullableStateTransitionRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStateTransitionRule(val *StateTransitionRule) *NullableStateTransitionRule {
	return &NullableStateTransitionRule{value: val, isSet: true}
}

func (v NullableStateTransitionRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStateTransitionRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


