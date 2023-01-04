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

// EventUpdateRequest struct for EventUpdateRequest
type EventUpdateRequest struct {
	// Short description of the event. Max length 64 characters.
	Description string `json:"description"`
	EventType *EventType `json:"eventType,omitempty"`
	// List of valid reason codes that can raise this event code
	ReasonCodes []string `json:"reasonCodes,omitempty"`
	// List of valid containerTypeNames where this event can be applied
	ContainerTypeNames []string `json:"containerTypeNames,omitempty"`
	// True - Event Code is valid & usable. False - Event Code is not usable
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// Data expected along with the event and their respective validations, if any.
	DataValidations []EventDataItem `json:"dataValidations,omitempty"`
}

// NewEventUpdateRequest instantiates a new EventUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventUpdateRequest(description string) *EventUpdateRequest {
	this := EventUpdateRequest{}
	this.Description = description
	var eventType EventType = EVENTTYPE_TRANSITIONAL
	this.EventType = &eventType
	var isEnabled bool = true
	this.IsEnabled = &isEnabled
	return &this
}

// NewEventUpdateRequestWithDefaults instantiates a new EventUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventUpdateRequestWithDefaults() *EventUpdateRequest {
	this := EventUpdateRequest{}
	var eventType EventType = EVENTTYPE_TRANSITIONAL
	this.EventType = &eventType
	var isEnabled bool = true
	this.IsEnabled = &isEnabled
	return &this
}

// GetDescription returns the Description field value
func (o *EventUpdateRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *EventUpdateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *EventUpdateRequest) SetDescription(v string) {
	o.Description = v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *EventUpdateRequest) GetEventType() EventType {
	if o == nil || isNil(o.EventType) {
		var ret EventType
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventUpdateRequest) GetEventTypeOk() (*EventType, bool) {
	if o == nil || isNil(o.EventType) {
    return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *EventUpdateRequest) HasEventType() bool {
	if o != nil && !isNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given EventType and assigns it to the EventType field.
func (o *EventUpdateRequest) SetEventType(v EventType) {
	o.EventType = &v
}

// GetReasonCodes returns the ReasonCodes field value if set, zero value otherwise.
func (o *EventUpdateRequest) GetReasonCodes() []string {
	if o == nil || isNil(o.ReasonCodes) {
		var ret []string
		return ret
	}
	return o.ReasonCodes
}

// GetReasonCodesOk returns a tuple with the ReasonCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventUpdateRequest) GetReasonCodesOk() ([]string, bool) {
	if o == nil || isNil(o.ReasonCodes) {
    return nil, false
	}
	return o.ReasonCodes, true
}

// HasReasonCodes returns a boolean if a field has been set.
func (o *EventUpdateRequest) HasReasonCodes() bool {
	if o != nil && !isNil(o.ReasonCodes) {
		return true
	}

	return false
}

// SetReasonCodes gets a reference to the given []string and assigns it to the ReasonCodes field.
func (o *EventUpdateRequest) SetReasonCodes(v []string) {
	o.ReasonCodes = v
}

// GetContainerTypeNames returns the ContainerTypeNames field value if set, zero value otherwise.
func (o *EventUpdateRequest) GetContainerTypeNames() []string {
	if o == nil || isNil(o.ContainerTypeNames) {
		var ret []string
		return ret
	}
	return o.ContainerTypeNames
}

// GetContainerTypeNamesOk returns a tuple with the ContainerTypeNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventUpdateRequest) GetContainerTypeNamesOk() ([]string, bool) {
	if o == nil || isNil(o.ContainerTypeNames) {
    return nil, false
	}
	return o.ContainerTypeNames, true
}

// HasContainerTypeNames returns a boolean if a field has been set.
func (o *EventUpdateRequest) HasContainerTypeNames() bool {
	if o != nil && !isNil(o.ContainerTypeNames) {
		return true
	}

	return false
}

// SetContainerTypeNames gets a reference to the given []string and assigns it to the ContainerTypeNames field.
func (o *EventUpdateRequest) SetContainerTypeNames(v []string) {
	o.ContainerTypeNames = v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *EventUpdateRequest) GetIsEnabled() bool {
	if o == nil || isNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventUpdateRequest) GetIsEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.IsEnabled) {
    return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *EventUpdateRequest) HasIsEnabled() bool {
	if o != nil && !isNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *EventUpdateRequest) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetDataValidations returns the DataValidations field value if set, zero value otherwise.
func (o *EventUpdateRequest) GetDataValidations() []EventDataItem {
	if o == nil || isNil(o.DataValidations) {
		var ret []EventDataItem
		return ret
	}
	return o.DataValidations
}

// GetDataValidationsOk returns a tuple with the DataValidations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventUpdateRequest) GetDataValidationsOk() ([]EventDataItem, bool) {
	if o == nil || isNil(o.DataValidations) {
    return nil, false
	}
	return o.DataValidations, true
}

// HasDataValidations returns a boolean if a field has been set.
func (o *EventUpdateRequest) HasDataValidations() bool {
	if o != nil && !isNil(o.DataValidations) {
		return true
	}

	return false
}

// SetDataValidations gets a reference to the given []EventDataItem and assigns it to the DataValidations field.
func (o *EventUpdateRequest) SetDataValidations(v []EventDataItem) {
	o.DataValidations = v
}

func (o EventUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	if !isNil(o.ReasonCodes) {
		toSerialize["reasonCodes"] = o.ReasonCodes
	}
	if !isNil(o.ContainerTypeNames) {
		toSerialize["containerTypeNames"] = o.ContainerTypeNames
	}
	if !isNil(o.IsEnabled) {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if !isNil(o.DataValidations) {
		toSerialize["dataValidations"] = o.DataValidations
	}
	return json.Marshal(toSerialize)
}

type NullableEventUpdateRequest struct {
	value *EventUpdateRequest
	isSet bool
}

func (v NullableEventUpdateRequest) Get() *EventUpdateRequest {
	return v.value
}

func (v *NullableEventUpdateRequest) Set(val *EventUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEventUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEventUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventUpdateRequest(val *EventUpdateRequest) *NullableEventUpdateRequest {
	return &NullableEventUpdateRequest{value: val, isSet: true}
}

func (v NullableEventUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


