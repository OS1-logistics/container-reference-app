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

// GroupChange user change information.
type GroupChange struct {
	Action ChangeAction `json:"action"`
	// Represents a date time as number of seconds elapsed since 00:00Hrs of 1st January 1970 UTC.
	ActionTime int64 `json:"actionTime"`
	ActionBy map[string]interface{} `json:"actionBy"`
	Channel string `json:"channel"`
	Cdc GroupResponse `json:"cdc"`
}

// NewGroupChange instantiates a new GroupChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupChange(action ChangeAction, actionTime int64, actionBy map[string]interface{}, channel string, cdc GroupResponse) *GroupChange {
	this := GroupChange{}
	this.Action = action
	this.ActionTime = actionTime
	this.ActionBy = actionBy
	this.Channel = channel
	this.Cdc = cdc
	return &this
}

// NewGroupChangeWithDefaults instantiates a new GroupChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupChangeWithDefaults() *GroupChange {
	this := GroupChange{}
	return &this
}

// GetAction returns the Action field value
func (o *GroupChange) GetAction() ChangeAction {
	if o == nil {
		var ret ChangeAction
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *GroupChange) GetActionOk() (*ChangeAction, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *GroupChange) SetAction(v ChangeAction) {
	o.Action = v
}

// GetActionTime returns the ActionTime field value
func (o *GroupChange) GetActionTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ActionTime
}

// GetActionTimeOk returns a tuple with the ActionTime field value
// and a boolean to check if the value has been set.
func (o *GroupChange) GetActionTimeOk() (*int64, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ActionTime, true
}

// SetActionTime sets field value
func (o *GroupChange) SetActionTime(v int64) {
	o.ActionTime = v
}

// GetActionBy returns the ActionBy field value
func (o *GroupChange) GetActionBy() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.ActionBy
}

// GetActionByOk returns a tuple with the ActionBy field value
// and a boolean to check if the value has been set.
func (o *GroupChange) GetActionByOk() (map[string]interface{}, bool) {
	if o == nil {
    return map[string]interface{}{}, false
	}
	return o.ActionBy, true
}

// SetActionBy sets field value
func (o *GroupChange) SetActionBy(v map[string]interface{}) {
	o.ActionBy = v
}

// GetChannel returns the Channel field value
func (o *GroupChange) GetChannel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *GroupChange) GetChannelOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value
func (o *GroupChange) SetChannel(v string) {
	o.Channel = v
}

// GetCdc returns the Cdc field value
func (o *GroupChange) GetCdc() GroupResponse {
	if o == nil {
		var ret GroupResponse
		return ret
	}

	return o.Cdc
}

// GetCdcOk returns a tuple with the Cdc field value
// and a boolean to check if the value has been set.
func (o *GroupChange) GetCdcOk() (*GroupResponse, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Cdc, true
}

// SetCdc sets field value
func (o *GroupChange) SetCdc(v GroupResponse) {
	o.Cdc = v
}

func (o GroupChange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["action"] = o.Action
	}
	if true {
		toSerialize["actionTime"] = o.ActionTime
	}
	if true {
		toSerialize["actionBy"] = o.ActionBy
	}
	if true {
		toSerialize["channel"] = o.Channel
	}
	if true {
		toSerialize["cdc"] = o.Cdc
	}
	return json.Marshal(toSerialize)
}

type NullableGroupChange struct {
	value *GroupChange
	isSet bool
}

func (v NullableGroupChange) Get() *GroupChange {
	return v.value
}

func (v *NullableGroupChange) Set(val *GroupChange) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupChange) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupChange(val *GroupChange) *NullableGroupChange {
	return &NullableGroupChange{value: val, isSet: true}
}

func (v NullableGroupChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


