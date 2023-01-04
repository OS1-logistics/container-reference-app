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

// BatchResponseMetaData struct for BatchResponseMetaData
type BatchResponseMetaData struct {
	// status of batch request
	Status *string `json:"status,omitempty"`
	// total number of successful record.
	Success *int64 `json:"success,omitempty"`
	// total number of failed records.
	Failed *int64 `json:"failed,omitempty"`
	// total number of record.
	Total *int64 `json:"total,omitempty"`
}

// NewBatchResponseMetaData instantiates a new BatchResponseMetaData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchResponseMetaData() *BatchResponseMetaData {
	this := BatchResponseMetaData{}
	return &this
}

// NewBatchResponseMetaDataWithDefaults instantiates a new BatchResponseMetaData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchResponseMetaDataWithDefaults() *BatchResponseMetaData {
	this := BatchResponseMetaData{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BatchResponseMetaData) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchResponseMetaData) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BatchResponseMetaData) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BatchResponseMetaData) SetStatus(v string) {
	o.Status = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *BatchResponseMetaData) GetSuccess() int64 {
	if o == nil || isNil(o.Success) {
		var ret int64
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchResponseMetaData) GetSuccessOk() (*int64, bool) {
	if o == nil || isNil(o.Success) {
    return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *BatchResponseMetaData) HasSuccess() bool {
	if o != nil && !isNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given int64 and assigns it to the Success field.
func (o *BatchResponseMetaData) SetSuccess(v int64) {
	o.Success = &v
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *BatchResponseMetaData) GetFailed() int64 {
	if o == nil || isNil(o.Failed) {
		var ret int64
		return ret
	}
	return *o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchResponseMetaData) GetFailedOk() (*int64, bool) {
	if o == nil || isNil(o.Failed) {
    return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *BatchResponseMetaData) HasFailed() bool {
	if o != nil && !isNil(o.Failed) {
		return true
	}

	return false
}

// SetFailed gets a reference to the given int64 and assigns it to the Failed field.
func (o *BatchResponseMetaData) SetFailed(v int64) {
	o.Failed = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *BatchResponseMetaData) GetTotal() int64 {
	if o == nil || isNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchResponseMetaData) GetTotalOk() (*int64, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *BatchResponseMetaData) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *BatchResponseMetaData) SetTotal(v int64) {
	o.Total = &v
}

func (o BatchResponseMetaData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !isNil(o.Failed) {
		toSerialize["failed"] = o.Failed
	}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableBatchResponseMetaData struct {
	value *BatchResponseMetaData
	isSet bool
}

func (v NullableBatchResponseMetaData) Get() *BatchResponseMetaData {
	return v.value
}

func (v *NullableBatchResponseMetaData) Set(val *BatchResponseMetaData) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchResponseMetaData) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchResponseMetaData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchResponseMetaData(val *BatchResponseMetaData) *NullableBatchResponseMetaData {
	return &NullableBatchResponseMetaData{value: val, isSet: true}
}

func (v NullableBatchResponseMetaData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchResponseMetaData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


