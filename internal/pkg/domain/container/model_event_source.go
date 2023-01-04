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

// EventSource Represents the source which triggered the event. It can be an app, a user or some location from where the event was triggered.
type EventSource struct {
	// application ID which is responsible for calling this event.
	AppId string `json:"appId"`
	// user ID which is responsible for calling this event.
	UserId *string `json:"userId,omitempty"`
	// locationId of the event.
	LocId *string `json:"locId,omitempty"`
}

// NewEventSource instantiates a new EventSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventSource(appId string) *EventSource {
	this := EventSource{}
	this.AppId = appId
	return &this
}

// NewEventSourceWithDefaults instantiates a new EventSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventSourceWithDefaults() *EventSource {
	this := EventSource{}
	return &this
}

// GetAppId returns the AppId field value
func (o *EventSource) GetAppId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value
// and a boolean to check if the value has been set.
func (o *EventSource) GetAppIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AppId, true
}

// SetAppId sets field value
func (o *EventSource) SetAppId(v string) {
	o.AppId = v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *EventSource) GetUserId() string {
	if o == nil || isNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSource) GetUserIdOk() (*string, bool) {
	if o == nil || isNil(o.UserId) {
    return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *EventSource) HasUserId() bool {
	if o != nil && !isNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *EventSource) SetUserId(v string) {
	o.UserId = &v
}

// GetLocId returns the LocId field value if set, zero value otherwise.
func (o *EventSource) GetLocId() string {
	if o == nil || isNil(o.LocId) {
		var ret string
		return ret
	}
	return *o.LocId
}

// GetLocIdOk returns a tuple with the LocId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSource) GetLocIdOk() (*string, bool) {
	if o == nil || isNil(o.LocId) {
    return nil, false
	}
	return o.LocId, true
}

// HasLocId returns a boolean if a field has been set.
func (o *EventSource) HasLocId() bool {
	if o != nil && !isNil(o.LocId) {
		return true
	}

	return false
}

// SetLocId gets a reference to the given string and assigns it to the LocId field.
func (o *EventSource) SetLocId(v string) {
	o.LocId = &v
}

func (o EventSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["appId"] = o.AppId
	}
	if !isNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !isNil(o.LocId) {
		toSerialize["locId"] = o.LocId
	}
	return json.Marshal(toSerialize)
}

type NullableEventSource struct {
	value *EventSource
	isSet bool
}

func (v NullableEventSource) Get() *EventSource {
	return v.value
}

func (v *NullableEventSource) Set(val *EventSource) {
	v.value = val
	v.isSet = true
}

func (v NullableEventSource) IsSet() bool {
	return v.isSet
}

func (v *NullableEventSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventSource(val *EventSource) *NullableEventSource {
	return &NullableEventSource{value: val, isSet: true}
}

func (v NullableEventSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

