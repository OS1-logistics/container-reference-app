/*
Authentication And Authorization (AAA) Service

This swagger documentation provides all AAA API details. AAA service provides authentication and authorization capabilities for users.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aaadomain

import (
	"encoding/json"
)

// FileStatus Status of the file uploaded
type FileStatus struct {
	// Unique Id of tenant
	TenantId *string `json:"tenantId,omitempty"`
	// Id of the yaml upload
	StatusId *string `json:"statusId,omitempty"`
	// Current status of uploaded file
	Status *string `json:"status,omitempty"`
	Reason []string `json:"reason,omitempty"`
}

// NewFileStatus instantiates a new FileStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileStatus() *FileStatus {
	this := FileStatus{}
	return &this
}

// NewFileStatusWithDefaults instantiates a new FileStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileStatusWithDefaults() *FileStatus {
	this := FileStatus{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *FileStatus) GetTenantId() string {
	if o == nil || o.TenantId == nil {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileStatus) GetTenantIdOk() (*string, bool) {
	if o == nil || o.TenantId == nil {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *FileStatus) HasTenantId() bool {
	if o != nil && o.TenantId != nil {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *FileStatus) SetTenantId(v string) {
	o.TenantId = &v
}

// GetStatusId returns the StatusId field value if set, zero value otherwise.
func (o *FileStatus) GetStatusId() string {
	if o == nil || o.StatusId == nil {
		var ret string
		return ret
	}
	return *o.StatusId
}

// GetStatusIdOk returns a tuple with the StatusId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileStatus) GetStatusIdOk() (*string, bool) {
	if o == nil || o.StatusId == nil {
		return nil, false
	}
	return o.StatusId, true
}

// HasStatusId returns a boolean if a field has been set.
func (o *FileStatus) HasStatusId() bool {
	if o != nil && o.StatusId != nil {
		return true
	}

	return false
}

// SetStatusId gets a reference to the given string and assigns it to the StatusId field.
func (o *FileStatus) SetStatusId(v string) {
	o.StatusId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FileStatus) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileStatus) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FileStatus) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *FileStatus) SetStatus(v string) {
	o.Status = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *FileStatus) GetReason() []string {
	if o == nil || o.Reason == nil {
		var ret []string
		return ret
	}
	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileStatus) GetReasonOk() ([]string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *FileStatus) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given []string and assigns it to the Reason field.
func (o *FileStatus) SetReason(v []string) {
	o.Reason = v
}

func (o FileStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TenantId != nil {
		toSerialize["tenantId"] = o.TenantId
	}
	if o.StatusId != nil {
		toSerialize["statusId"] = o.StatusId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	return json.Marshal(toSerialize)
}

type NullableFileStatus struct {
	value *FileStatus
	isSet bool
}

func (v NullableFileStatus) Get() *FileStatus {
	return v.value
}

func (v *NullableFileStatus) Set(val *FileStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableFileStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableFileStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileStatus(val *FileStatus) *NullableFileStatus {
	return &NullableFileStatus{value: val, isSet: true}
}

func (v NullableFileStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


