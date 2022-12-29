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

// EventsResponse struct for EventsResponse
type EventsResponse struct {
	// Short description of the event. Max length 64 characters.
	Description *string `json:"description,omitempty"`
	EventType *EventType `json:"eventType,omitempty"`
	// List of valid reason codes that can raise this event code
	Reasons []ReasonMetaData `json:"reasons,omitempty"`
	// List of valid entities where this event can be applied
	Entities []EntitiesMetaData `json:"entities,omitempty"`
	// True - If the event is system generated. False - If the event is custom generated
	IsSystemDefined *bool `json:"isSystemDefined,omitempty"`
	// True - Event Code is valid & usable. False - Event Code is not usable
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// Data expected along with the event and their respective validations, if any.
	DataValidations []EventDataItem `json:"dataValidations,omitempty"`
	// Event Code
	EventCode *string `json:"eventCode,omitempty"`
}

// NewEventsResponse instantiates a new EventsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventsResponse() *EventsResponse {
	this := EventsResponse{}
	var eventType EventType = EVENTTYPE_TRANSITIONAL
	this.EventType = &eventType
	var isEnabled bool = true
	this.IsEnabled = &isEnabled
	return &this
}

// NewEventsResponseWithDefaults instantiates a new EventsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventsResponseWithDefaults() *EventsResponse {
	this := EventsResponse{}
	var eventType EventType = EVENTTYPE_TRANSITIONAL
	this.EventType = &eventType
	var isEnabled bool = true
	this.IsEnabled = &isEnabled
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EventsResponse) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EventsResponse) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EventsResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *EventsResponse) GetEventType() EventType {
	if o == nil || o.EventType == nil {
		var ret EventType
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsResponse) GetEventTypeOk() (*EventType, bool) {
	if o == nil || o.EventType == nil {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *EventsResponse) HasEventType() bool {
	if o != nil && o.EventType != nil {
		return true
	}

	return false
}

// SetEventType gets a reference to the given EventType and assigns it to the EventType field.
func (o *EventsResponse) SetEventType(v EventType) {
	o.EventType = &v
}

// GetReasons returns the Reasons field value if set, zero value otherwise.
func (o *EventsResponse) GetReasons() []ReasonMetaData {
	if o == nil || o.Reasons == nil {
		var ret []ReasonMetaData
		return ret
	}
	return o.Reasons
}

// GetReasonsOk returns a tuple with the Reasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsResponse) GetReasonsOk() ([]ReasonMetaData, bool) {
	if o == nil || o.Reasons == nil {
		return nil, false
	}
	return o.Reasons, true
}

// HasReasons returns a boolean if a field has been set.
func (o *EventsResponse) HasReasons() bool {
	if o != nil && o.Reasons != nil {
		return true
	}

	return false
}

// SetReasons gets a reference to the given []ReasonMetaData and assigns it to the Reasons field.
func (o *EventsResponse) SetReasons(v []ReasonMetaData) {
	o.Reasons = v
}

// GetEntities returns the Entities field value if set, zero value otherwise.
func (o *EventsResponse) GetEntities() []EntitiesMetaData {
	if o == nil || o.Entities == nil {
		var ret []EntitiesMetaData
		return ret
	}
	return o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsResponse) GetEntitiesOk() ([]EntitiesMetaData, bool) {
	if o == nil || o.Entities == nil {
		return nil, false
	}
	return o.Entities, true
}

// HasEntities returns a boolean if a field has been set.
func (o *EventsResponse) HasEntities() bool {
	if o != nil && o.Entities != nil {
		return true
	}

	return false
}

// SetEntities gets a reference to the given []EntitiesMetaData and assigns it to the Entities field.
func (o *EventsResponse) SetEntities(v []EntitiesMetaData) {
	o.Entities = v
}

// GetIsSystemDefined returns the IsSystemDefined field value if set, zero value otherwise.
func (o *EventsResponse) GetIsSystemDefined() bool {
	if o == nil || o.IsSystemDefined == nil {
		var ret bool
		return ret
	}
	return *o.IsSystemDefined
}

// GetIsSystemDefinedOk returns a tuple with the IsSystemDefined field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsResponse) GetIsSystemDefinedOk() (*bool, bool) {
	if o == nil || o.IsSystemDefined == nil {
		return nil, false
	}
	return o.IsSystemDefined, true
}

// HasIsSystemDefined returns a boolean if a field has been set.
func (o *EventsResponse) HasIsSystemDefined() bool {
	if o != nil && o.IsSystemDefined != nil {
		return true
	}

	return false
}

// SetIsSystemDefined gets a reference to the given bool and assigns it to the IsSystemDefined field.
func (o *EventsResponse) SetIsSystemDefined(v bool) {
	o.IsSystemDefined = &v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *EventsResponse) GetIsEnabled() bool {
	if o == nil || o.IsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsResponse) GetIsEnabledOk() (*bool, bool) {
	if o == nil || o.IsEnabled == nil {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *EventsResponse) HasIsEnabled() bool {
	if o != nil && o.IsEnabled != nil {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *EventsResponse) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetDataValidations returns the DataValidations field value if set, zero value otherwise.
func (o *EventsResponse) GetDataValidations() []EventDataItem {
	if o == nil || o.DataValidations == nil {
		var ret []EventDataItem
		return ret
	}
	return o.DataValidations
}

// GetDataValidationsOk returns a tuple with the DataValidations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsResponse) GetDataValidationsOk() ([]EventDataItem, bool) {
	if o == nil || o.DataValidations == nil {
		return nil, false
	}
	return o.DataValidations, true
}

// HasDataValidations returns a boolean if a field has been set.
func (o *EventsResponse) HasDataValidations() bool {
	if o != nil && o.DataValidations != nil {
		return true
	}

	return false
}

// SetDataValidations gets a reference to the given []EventDataItem and assigns it to the DataValidations field.
func (o *EventsResponse) SetDataValidations(v []EventDataItem) {
	o.DataValidations = v
}

// GetEventCode returns the EventCode field value if set, zero value otherwise.
func (o *EventsResponse) GetEventCode() string {
	if o == nil || o.EventCode == nil {
		var ret string
		return ret
	}
	return *o.EventCode
}

// GetEventCodeOk returns a tuple with the EventCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsResponse) GetEventCodeOk() (*string, bool) {
	if o == nil || o.EventCode == nil {
		return nil, false
	}
	return o.EventCode, true
}

// HasEventCode returns a boolean if a field has been set.
func (o *EventsResponse) HasEventCode() bool {
	if o != nil && o.EventCode != nil {
		return true
	}

	return false
}

// SetEventCode gets a reference to the given string and assigns it to the EventCode field.
func (o *EventsResponse) SetEventCode(v string) {
	o.EventCode = &v
}

func (o EventsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.EventType != nil {
		toSerialize["eventType"] = o.EventType
	}
	if o.Reasons != nil {
		toSerialize["reasons"] = o.Reasons
	}
	if o.Entities != nil {
		toSerialize["entities"] = o.Entities
	}
	if o.IsSystemDefined != nil {
		toSerialize["isSystemDefined"] = o.IsSystemDefined
	}
	if o.IsEnabled != nil {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if o.DataValidations != nil {
		toSerialize["dataValidations"] = o.DataValidations
	}
	if o.EventCode != nil {
		toSerialize["eventCode"] = o.EventCode
	}
	return json.Marshal(toSerialize)
}

type NullableEventsResponse struct {
	value *EventsResponse
	isSet bool
}

func (v NullableEventsResponse) Get() *EventsResponse {
	return v.value
}

func (v *NullableEventsResponse) Set(val *EventsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEventsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEventsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventsResponse(val *EventsResponse) *NullableEventsResponse {
	return &NullableEventsResponse{value: val, isSet: true}
}

func (v NullableEventsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


