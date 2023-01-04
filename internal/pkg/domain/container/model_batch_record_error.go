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

// BatchRecordError struct for BatchRecordError
type BatchRecordError struct {
	// reference ID of container create batch request.
	RefID *string `json:"refID,omitempty"`
	// failure reason of record.
	Reason *string `json:"reason,omitempty"`
}

// NewBatchRecordError instantiates a new BatchRecordError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchRecordError() *BatchRecordError {
	this := BatchRecordError{}
	return &this
}

// NewBatchRecordErrorWithDefaults instantiates a new BatchRecordError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchRecordErrorWithDefaults() *BatchRecordError {
	this := BatchRecordError{}
	return &this
}

// GetRefID returns the RefID field value if set, zero value otherwise.
func (o *BatchRecordError) GetRefID() string {
	if o == nil || isNil(o.RefID) {
		var ret string
		return ret
	}
	return *o.RefID
}

// GetRefIDOk returns a tuple with the RefID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchRecordError) GetRefIDOk() (*string, bool) {
	if o == nil || isNil(o.RefID) {
    return nil, false
	}
	return o.RefID, true
}

// HasRefID returns a boolean if a field has been set.
func (o *BatchRecordError) HasRefID() bool {
	if o != nil && !isNil(o.RefID) {
		return true
	}

	return false
}

// SetRefID gets a reference to the given string and assigns it to the RefID field.
func (o *BatchRecordError) SetRefID(v string) {
	o.RefID = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *BatchRecordError) GetReason() string {
	if o == nil || isNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchRecordError) GetReasonOk() (*string, bool) {
	if o == nil || isNil(o.Reason) {
    return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *BatchRecordError) HasReason() bool {
	if o != nil && !isNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *BatchRecordError) SetReason(v string) {
	o.Reason = &v
}

func (o BatchRecordError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RefID) {
		toSerialize["refID"] = o.RefID
	}
	if !isNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return json.Marshal(toSerialize)
}

type NullableBatchRecordError struct {
	value *BatchRecordError
	isSet bool
}

func (v NullableBatchRecordError) Get() *BatchRecordError {
	return v.value
}

func (v *NullableBatchRecordError) Set(val *BatchRecordError) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchRecordError) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchRecordError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchRecordError(val *BatchRecordError) *NullableBatchRecordError {
	return &NullableBatchRecordError{value: val, isSet: true}
}

func (v NullableBatchRecordError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchRecordError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

