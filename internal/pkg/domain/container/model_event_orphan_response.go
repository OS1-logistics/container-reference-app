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

// EventOrphanResponse struct for EventOrphanResponse
type EventOrphanResponse struct {
	// Short description of the event. Max length 64 characters.
	Description *string `json:"description,omitempty"`
	EventType *EventType `json:"eventType,omitempty"`
	// List of valid current states of the entity when this event code can be applied.
	ValidCurrentStates []string `json:"validCurrentStates,omitempty"`
	// True - Event Code is valid & usable. False - Event Code is not usable
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// Data expected along with the event and their respective validations, if any.
	DataValidations []EventDataItem `json:"dataValidations,omitempty"`
	// event code.
	EventCode *string `json:"eventCode,omitempty"`
	// Entity Code
	EnityCode *string `json:"enityCode,omitempty"`
}

// NewEventOrphanResponse instantiates a new EventOrphanResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventOrphanResponse() *EventOrphanResponse {
	this := EventOrphanResponse{}
	var eventType EventType = EVENTTYPE_TRANSITIONAL
	this.EventType = &eventType
	var isEnabled bool = true
	this.IsEnabled = &isEnabled
	return &this
}

// NewEventOrphanResponseWithDefaults instantiates a new EventOrphanResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventOrphanResponseWithDefaults() *EventOrphanResponse {
	this := EventOrphanResponse{}
	var eventType EventType = EVENTTYPE_TRANSITIONAL
	this.EventType = &eventType
	var isEnabled bool = true
	this.IsEnabled = &isEnabled
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EventOrphanResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventOrphanResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EventOrphanResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EventOrphanResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *EventOrphanResponse) GetEventType() EventType {
	if o == nil || isNil(o.EventType) {
		var ret EventType
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventOrphanResponse) GetEventTypeOk() (*EventType, bool) {
	if o == nil || isNil(o.EventType) {
    return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *EventOrphanResponse) HasEventType() bool {
	if o != nil && !isNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given EventType and assigns it to the EventType field.
func (o *EventOrphanResponse) SetEventType(v EventType) {
	o.EventType = &v
}

// GetValidCurrentStates returns the ValidCurrentStates field value if set, zero value otherwise.
func (o *EventOrphanResponse) GetValidCurrentStates() []string {
	if o == nil || isNil(o.ValidCurrentStates) {
		var ret []string
		return ret
	}
	return o.ValidCurrentStates
}

// GetValidCurrentStatesOk returns a tuple with the ValidCurrentStates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventOrphanResponse) GetValidCurrentStatesOk() ([]string, bool) {
	if o == nil || isNil(o.ValidCurrentStates) {
    return nil, false
	}
	return o.ValidCurrentStates, true
}

// HasValidCurrentStates returns a boolean if a field has been set.
func (o *EventOrphanResponse) HasValidCurrentStates() bool {
	if o != nil && !isNil(o.ValidCurrentStates) {
		return true
	}

	return false
}

// SetValidCurrentStates gets a reference to the given []string and assigns it to the ValidCurrentStates field.
func (o *EventOrphanResponse) SetValidCurrentStates(v []string) {
	o.ValidCurrentStates = v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *EventOrphanResponse) GetIsEnabled() bool {
	if o == nil || isNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventOrphanResponse) GetIsEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.IsEnabled) {
    return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *EventOrphanResponse) HasIsEnabled() bool {
	if o != nil && !isNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *EventOrphanResponse) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetDataValidations returns the DataValidations field value if set, zero value otherwise.
func (o *EventOrphanResponse) GetDataValidations() []EventDataItem {
	if o == nil || isNil(o.DataValidations) {
		var ret []EventDataItem
		return ret
	}
	return o.DataValidations
}

// GetDataValidationsOk returns a tuple with the DataValidations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventOrphanResponse) GetDataValidationsOk() ([]EventDataItem, bool) {
	if o == nil || isNil(o.DataValidations) {
    return nil, false
	}
	return o.DataValidations, true
}

// HasDataValidations returns a boolean if a field has been set.
func (o *EventOrphanResponse) HasDataValidations() bool {
	if o != nil && !isNil(o.DataValidations) {
		return true
	}

	return false
}

// SetDataValidations gets a reference to the given []EventDataItem and assigns it to the DataValidations field.
func (o *EventOrphanResponse) SetDataValidations(v []EventDataItem) {
	o.DataValidations = v
}

// GetEventCode returns the EventCode field value if set, zero value otherwise.
func (o *EventOrphanResponse) GetEventCode() string {
	if o == nil || isNil(o.EventCode) {
		var ret string
		return ret
	}
	return *o.EventCode
}

// GetEventCodeOk returns a tuple with the EventCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventOrphanResponse) GetEventCodeOk() (*string, bool) {
	if o == nil || isNil(o.EventCode) {
    return nil, false
	}
	return o.EventCode, true
}

// HasEventCode returns a boolean if a field has been set.
func (o *EventOrphanResponse) HasEventCode() bool {
	if o != nil && !isNil(o.EventCode) {
		return true
	}

	return false
}

// SetEventCode gets a reference to the given string and assigns it to the EventCode field.
func (o *EventOrphanResponse) SetEventCode(v string) {
	o.EventCode = &v
}

// GetEnityCode returns the EnityCode field value if set, zero value otherwise.
func (o *EventOrphanResponse) GetEnityCode() string {
	if o == nil || isNil(o.EnityCode) {
		var ret string
		return ret
	}
	return *o.EnityCode
}

// GetEnityCodeOk returns a tuple with the EnityCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventOrphanResponse) GetEnityCodeOk() (*string, bool) {
	if o == nil || isNil(o.EnityCode) {
    return nil, false
	}
	return o.EnityCode, true
}

// HasEnityCode returns a boolean if a field has been set.
func (o *EventOrphanResponse) HasEnityCode() bool {
	if o != nil && !isNil(o.EnityCode) {
		return true
	}

	return false
}

// SetEnityCode gets a reference to the given string and assigns it to the EnityCode field.
func (o *EventOrphanResponse) SetEnityCode(v string) {
	o.EnityCode = &v
}

func (o EventOrphanResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	if !isNil(o.ValidCurrentStates) {
		toSerialize["validCurrentStates"] = o.ValidCurrentStates
	}
	if !isNil(o.IsEnabled) {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if !isNil(o.DataValidations) {
		toSerialize["dataValidations"] = o.DataValidations
	}
	if !isNil(o.EventCode) {
		toSerialize["eventCode"] = o.EventCode
	}
	if !isNil(o.EnityCode) {
		toSerialize["enityCode"] = o.EnityCode
	}
	return json.Marshal(toSerialize)
}

type NullableEventOrphanResponse struct {
	value *EventOrphanResponse
	isSet bool
}

func (v NullableEventOrphanResponse) Get() *EventOrphanResponse {
	return v.value
}

func (v *NullableEventOrphanResponse) Set(val *EventOrphanResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEventOrphanResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEventOrphanResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventOrphanResponse(val *EventOrphanResponse) *NullableEventOrphanResponse {
	return &NullableEventOrphanResponse{value: val, isSet: true}
}

func (v NullableEventOrphanResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventOrphanResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


