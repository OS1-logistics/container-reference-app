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

// EntityStateValidationRequest Validate an entity state
type EntityStateValidationRequest struct {
	// reasonCode.
	ReasonCode string `json:"reasonCode"`
	// event code.
	EventCode string `json:"eventCode"`
	// given state.
	State string `json:"state"`
}

// NewEntityStateValidationRequest instantiates a new EntityStateValidationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityStateValidationRequest(reasonCode string, eventCode string, state string) *EntityStateValidationRequest {
	this := EntityStateValidationRequest{}
	this.ReasonCode = reasonCode
	this.EventCode = eventCode
	this.State = state
	return &this
}

// NewEntityStateValidationRequestWithDefaults instantiates a new EntityStateValidationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityStateValidationRequestWithDefaults() *EntityStateValidationRequest {
	this := EntityStateValidationRequest{}
	return &this
}

// GetReasonCode returns the ReasonCode field value
func (o *EntityStateValidationRequest) GetReasonCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReasonCode
}

// GetReasonCodeOk returns a tuple with the ReasonCode field value
// and a boolean to check if the value has been set.
func (o *EntityStateValidationRequest) GetReasonCodeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ReasonCode, true
}

// SetReasonCode sets field value
func (o *EntityStateValidationRequest) SetReasonCode(v string) {
	o.ReasonCode = v
}

// GetEventCode returns the EventCode field value
func (o *EntityStateValidationRequest) GetEventCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventCode
}

// GetEventCodeOk returns a tuple with the EventCode field value
// and a boolean to check if the value has been set.
func (o *EntityStateValidationRequest) GetEventCodeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EventCode, true
}

// SetEventCode sets field value
func (o *EntityStateValidationRequest) SetEventCode(v string) {
	o.EventCode = v
}

// GetState returns the State field value
func (o *EntityStateValidationRequest) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *EntityStateValidationRequest) GetStateOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *EntityStateValidationRequest) SetState(v string) {
	o.State = v
}

func (o EntityStateValidationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["reasonCode"] = o.ReasonCode
	}
	if true {
		toSerialize["eventCode"] = o.EventCode
	}
	if true {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableEntityStateValidationRequest struct {
	value *EntityStateValidationRequest
	isSet bool
}

func (v NullableEntityStateValidationRequest) Get() *EntityStateValidationRequest {
	return v.value
}

func (v *NullableEntityStateValidationRequest) Set(val *EntityStateValidationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityStateValidationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityStateValidationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityStateValidationRequest(val *EntityStateValidationRequest) *NullableEntityStateValidationRequest {
	return &NullableEntityStateValidationRequest{value: val, isSet: true}
}

func (v NullableEntityStateValidationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityStateValidationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


