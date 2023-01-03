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

// Event struct for Event
type Event struct {
	// Event Code of the format:E-yyy. E - Prefix indicating that this is an event code.  yyy - 3 digit event code string ranging from 001 to 999
	EventCode string `json:"eventCode"`
	// Reason Code of the format:R-nnnn R - Prefix indicating that this is a reason code nnnn - 4 digit reason code string ranging from 0001 to 9999
	ReasonCode *string `json:"reasonCode,omitempty"`
	// Represents the timestamp of the event occured.
	Timestamp int32 `json:"timestamp"`
	// propagate flag indicate whether these events need to propogate on child containers or not.
	Propagate *bool `json:"propagate,omitempty"`
	// Represents event data values, if any, for this event
	Data map[string]interface{} `json:"data"`
	Source EventSource `json:"source"`
}

// NewEvent instantiates a new Event object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvent(eventCode string, timestamp int32, data map[string]interface{}, source EventSource) *Event {
	this := Event{}
	this.EventCode = eventCode
	this.Timestamp = timestamp
	var propagate bool = false
	this.Propagate = &propagate
	this.Data = data
	this.Source = source
	return &this
}

// NewEventWithDefaults instantiates a new Event object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventWithDefaults() *Event {
	this := Event{}
	var propagate bool = false
	this.Propagate = &propagate
	return &this
}

// GetEventCode returns the EventCode field value
func (o *Event) GetEventCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventCode
}

// GetEventCodeOk returns a tuple with the EventCode field value
// and a boolean to check if the value has been set.
func (o *Event) GetEventCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventCode, true
}

// SetEventCode sets field value
func (o *Event) SetEventCode(v string) {
	o.EventCode = v
}

// GetReasonCode returns the ReasonCode field value if set, zero value otherwise.
func (o *Event) GetReasonCode() string {
	if o == nil || o.ReasonCode == nil {
		var ret string
		return ret
	}
	return *o.ReasonCode
}

// GetReasonCodeOk returns a tuple with the ReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetReasonCodeOk() (*string, bool) {
	if o == nil || o.ReasonCode == nil {
		return nil, false
	}
	return o.ReasonCode, true
}

// HasReasonCode returns a boolean if a field has been set.
func (o *Event) HasReasonCode() bool {
	if o != nil && o.ReasonCode != nil {
		return true
	}

	return false
}

// SetReasonCode gets a reference to the given string and assigns it to the ReasonCode field.
func (o *Event) SetReasonCode(v string) {
	o.ReasonCode = &v
}

// GetTimestamp returns the Timestamp field value
func (o *Event) GetTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *Event) GetTimestampOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *Event) SetTimestamp(v int32) {
	o.Timestamp = v
}

// GetPropagate returns the Propagate field value if set, zero value otherwise.
func (o *Event) GetPropagate() bool {
	if o == nil || o.Propagate == nil {
		var ret bool
		return ret
	}
	return *o.Propagate
}

// GetPropagateOk returns a tuple with the Propagate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetPropagateOk() (*bool, bool) {
	if o == nil || o.Propagate == nil {
		return nil, false
	}
	return o.Propagate, true
}

// HasPropagate returns a boolean if a field has been set.
func (o *Event) HasPropagate() bool {
	if o != nil && o.Propagate != nil {
		return true
	}

	return false
}

// SetPropagate gets a reference to the given bool and assigns it to the Propagate field.
func (o *Event) SetPropagate(v bool) {
	o.Propagate = &v
}

// GetData returns the Data field value
func (o *Event) GetData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *Event) GetDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *Event) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetSource returns the Source field value
func (o *Event) GetSource() EventSource {
	if o == nil {
		var ret EventSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *Event) GetSourceOk() (*EventSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *Event) SetSource(v EventSource) {
	o.Source = v
}

func (o Event) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["eventCode"] = o.EventCode
	}
	if o.ReasonCode != nil {
		toSerialize["reasonCode"] = o.ReasonCode
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.Propagate != nil {
		toSerialize["propagate"] = o.Propagate
	}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["source"] = o.Source
	}
	return json.Marshal(toSerialize)
}

type NullableEvent struct {
	value *Event
	isSet bool
}

func (v NullableEvent) Get() *Event {
	return v.value
}

func (v *NullableEvent) Set(val *Event) {
	v.value = val
	v.isSet = true
}

func (v NullableEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvent(val *Event) *NullableEvent {
	return &NullableEvent{value: val, isSet: true}
}

func (v NullableEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


