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

// TrackingDetail struct for TrackingDetail
type TrackingDetail struct {
	// Field to specify the owner of the tracking ID.
	Operator string `json:"operator"`
	TrackingId string `json:"trackingId"`
	// it defines whether it is a primary tracking id or not. If not defined we assume first tracking id as primary tracking id.
	IsPrimary *bool `json:"isPrimary,omitempty"`
	// Represents a date time as number of seconds elapsed since 00:00Hrs of 1st January 1970 UTC.
	Timestamp *int64 `json:"timestamp,omitempty"`
}

// NewTrackingDetail instantiates a new TrackingDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrackingDetail(operator string, trackingId string) *TrackingDetail {
	this := TrackingDetail{}
	this.Operator = operator
	this.TrackingId = trackingId
	var isPrimary bool = false
	this.IsPrimary = &isPrimary
	return &this
}

// NewTrackingDetailWithDefaults instantiates a new TrackingDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrackingDetailWithDefaults() *TrackingDetail {
	this := TrackingDetail{}
	var isPrimary bool = false
	this.IsPrimary = &isPrimary
	return &this
}

// GetOperator returns the Operator field value
func (o *TrackingDetail) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *TrackingDetail) GetOperatorOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *TrackingDetail) SetOperator(v string) {
	o.Operator = v
}

// GetTrackingId returns the TrackingId field value
func (o *TrackingDetail) GetTrackingId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TrackingId
}

// GetTrackingIdOk returns a tuple with the TrackingId field value
// and a boolean to check if the value has been set.
func (o *TrackingDetail) GetTrackingIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TrackingId, true
}

// SetTrackingId sets field value
func (o *TrackingDetail) SetTrackingId(v string) {
	o.TrackingId = v
}

// GetIsPrimary returns the IsPrimary field value if set, zero value otherwise.
func (o *TrackingDetail) GetIsPrimary() bool {
	if o == nil || isNil(o.IsPrimary) {
		var ret bool
		return ret
	}
	return *o.IsPrimary
}

// GetIsPrimaryOk returns a tuple with the IsPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingDetail) GetIsPrimaryOk() (*bool, bool) {
	if o == nil || isNil(o.IsPrimary) {
    return nil, false
	}
	return o.IsPrimary, true
}

// HasIsPrimary returns a boolean if a field has been set.
func (o *TrackingDetail) HasIsPrimary() bool {
	if o != nil && !isNil(o.IsPrimary) {
		return true
	}

	return false
}

// SetIsPrimary gets a reference to the given bool and assigns it to the IsPrimary field.
func (o *TrackingDetail) SetIsPrimary(v bool) {
	o.IsPrimary = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *TrackingDetail) GetTimestamp() int64 {
	if o == nil || isNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingDetail) GetTimestampOk() (*int64, bool) {
	if o == nil || isNil(o.Timestamp) {
    return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *TrackingDetail) HasTimestamp() bool {
	if o != nil && !isNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *TrackingDetail) SetTimestamp(v int64) {
	o.Timestamp = &v
}

func (o TrackingDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["operator"] = o.Operator
	}
	if true {
		toSerialize["trackingId"] = o.TrackingId
	}
	if !isNil(o.IsPrimary) {
		toSerialize["isPrimary"] = o.IsPrimary
	}
	if !isNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	return json.Marshal(toSerialize)
}

type NullableTrackingDetail struct {
	value *TrackingDetail
	isSet bool
}

func (v NullableTrackingDetail) Get() *TrackingDetail {
	return v.value
}

func (v *NullableTrackingDetail) Set(val *TrackingDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableTrackingDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableTrackingDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrackingDetail(val *TrackingDetail) *NullableTrackingDetail {
	return &NullableTrackingDetail{value: val, isSet: true}
}

func (v NullableTrackingDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrackingDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


