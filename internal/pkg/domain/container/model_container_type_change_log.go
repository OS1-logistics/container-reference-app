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

// ContainerTypeChangeLog Container type change information.
type ContainerTypeChangeLog struct {
	Action ChangeAction `json:"action"`
	ActionTime int64 `json:"actionTime"`
	ActionBy ActionBy `json:"actionBy"`
	// Container type change log data.
	Cdc map[string]interface{} `json:"cdc"`
}

// NewContainerTypeChangeLog instantiates a new ContainerTypeChangeLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerTypeChangeLog(action ChangeAction, actionTime int64, actionBy ActionBy, cdc map[string]interface{}) *ContainerTypeChangeLog {
	this := ContainerTypeChangeLog{}
	this.Action = action
	this.ActionTime = actionTime
	this.ActionBy = actionBy
	this.Cdc = cdc
	return &this
}

// NewContainerTypeChangeLogWithDefaults instantiates a new ContainerTypeChangeLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerTypeChangeLogWithDefaults() *ContainerTypeChangeLog {
	this := ContainerTypeChangeLog{}
	return &this
}

// GetAction returns the Action field value
func (o *ContainerTypeChangeLog) GetAction() ChangeAction {
	if o == nil {
		var ret ChangeAction
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ContainerTypeChangeLog) GetActionOk() (*ChangeAction, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *ContainerTypeChangeLog) SetAction(v ChangeAction) {
	o.Action = v
}

// GetActionTime returns the ActionTime field value
func (o *ContainerTypeChangeLog) GetActionTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ActionTime
}

// GetActionTimeOk returns a tuple with the ActionTime field value
// and a boolean to check if the value has been set.
func (o *ContainerTypeChangeLog) GetActionTimeOk() (*int64, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ActionTime, true
}

// SetActionTime sets field value
func (o *ContainerTypeChangeLog) SetActionTime(v int64) {
	o.ActionTime = v
}

// GetActionBy returns the ActionBy field value
func (o *ContainerTypeChangeLog) GetActionBy() ActionBy {
	if o == nil {
		var ret ActionBy
		return ret
	}

	return o.ActionBy
}

// GetActionByOk returns a tuple with the ActionBy field value
// and a boolean to check if the value has been set.
func (o *ContainerTypeChangeLog) GetActionByOk() (*ActionBy, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ActionBy, true
}

// SetActionBy sets field value
func (o *ContainerTypeChangeLog) SetActionBy(v ActionBy) {
	o.ActionBy = v
}

// GetCdc returns the Cdc field value
func (o *ContainerTypeChangeLog) GetCdc() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Cdc
}

// GetCdcOk returns a tuple with the Cdc field value
// and a boolean to check if the value has been set.
func (o *ContainerTypeChangeLog) GetCdcOk() (map[string]interface{}, bool) {
	if o == nil {
    return map[string]interface{}{}, false
	}
	return o.Cdc, true
}

// SetCdc sets field value
func (o *ContainerTypeChangeLog) SetCdc(v map[string]interface{}) {
	o.Cdc = v
}

func (o ContainerTypeChangeLog) MarshalJSON() ([]byte, error) {
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
		toSerialize["cdc"] = o.Cdc
	}
	return json.Marshal(toSerialize)
}

type NullableContainerTypeChangeLog struct {
	value *ContainerTypeChangeLog
	isSet bool
}

func (v NullableContainerTypeChangeLog) Get() *ContainerTypeChangeLog {
	return v.value
}

func (v *NullableContainerTypeChangeLog) Set(val *ContainerTypeChangeLog) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerTypeChangeLog) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerTypeChangeLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerTypeChangeLog(val *ContainerTypeChangeLog) *NullableContainerTypeChangeLog {
	return &NullableContainerTypeChangeLog{value: val, isSet: true}
}

func (v NullableContainerTypeChangeLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerTypeChangeLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


