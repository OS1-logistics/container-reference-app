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

// ReasonCreateRequest Create reason
type ReasonCreateRequest struct {
	// Short description of the reason. Max length 64 characters.
	Description string `json:"description"`
	// True - Event Code is valid & usable. False - Event Code is not usable
	IsEnabled *bool `json:"isEnabled,omitempty"`
}

// NewReasonCreateRequest instantiates a new ReasonCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReasonCreateRequest(description string) *ReasonCreateRequest {
	this := ReasonCreateRequest{}
	this.Description = description
	var isEnabled bool = true
	this.IsEnabled = &isEnabled
	return &this
}

// NewReasonCreateRequestWithDefaults instantiates a new ReasonCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReasonCreateRequestWithDefaults() *ReasonCreateRequest {
	this := ReasonCreateRequest{}
	var isEnabled bool = true
	this.IsEnabled = &isEnabled
	return &this
}

// GetDescription returns the Description field value
func (o *ReasonCreateRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ReasonCreateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ReasonCreateRequest) SetDescription(v string) {
	o.Description = v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *ReasonCreateRequest) GetIsEnabled() bool {
	if o == nil || isNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReasonCreateRequest) GetIsEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.IsEnabled) {
    return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *ReasonCreateRequest) HasIsEnabled() bool {
	if o != nil && !isNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *ReasonCreateRequest) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

func (o ReasonCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.IsEnabled) {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableReasonCreateRequest struct {
	value *ReasonCreateRequest
	isSet bool
}

func (v NullableReasonCreateRequest) Get() *ReasonCreateRequest {
	return v.value
}

func (v *NullableReasonCreateRequest) Set(val *ReasonCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReasonCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReasonCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReasonCreateRequest(val *ReasonCreateRequest) *NullableReasonCreateRequest {
	return &NullableReasonCreateRequest{value: val, isSet: true}
}

func (v NullableReasonCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReasonCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


