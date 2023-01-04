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

// TenantIdInfo struct for TenantIdInfo
type TenantIdInfo struct {
	// Tenant ID pertaining to the provided alias.
	TenantId *string `json:"tenantId,omitempty"`
	// Denotes the auth0 organization this platform Tenant belongs to.
	OrgId *string `json:"orgId,omitempty"`
}

// NewTenantIdInfo instantiates a new TenantIdInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantIdInfo() *TenantIdInfo {
	this := TenantIdInfo{}
	return &this
}

// NewTenantIdInfoWithDefaults instantiates a new TenantIdInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantIdInfoWithDefaults() *TenantIdInfo {
	this := TenantIdInfo{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *TenantIdInfo) GetTenantId() string {
	if o == nil || o.TenantId == nil {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantIdInfo) GetTenantIdOk() (*string, bool) {
	if o == nil || o.TenantId == nil {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *TenantIdInfo) HasTenantId() bool {
	if o != nil && o.TenantId != nil {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *TenantIdInfo) SetTenantId(v string) {
	o.TenantId = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *TenantIdInfo) GetOrgId() string {
	if o == nil || o.OrgId == nil {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantIdInfo) GetOrgIdOk() (*string, bool) {
	if o == nil || o.OrgId == nil {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *TenantIdInfo) HasOrgId() bool {
	if o != nil && o.OrgId != nil {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *TenantIdInfo) SetOrgId(v string) {
	o.OrgId = &v
}

func (o TenantIdInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TenantId != nil {
		toSerialize["tenantId"] = o.TenantId
	}
	if o.OrgId != nil {
		toSerialize["orgId"] = o.OrgId
	}
	return json.Marshal(toSerialize)
}

type NullableTenantIdInfo struct {
	value *TenantIdInfo
	isSet bool
}

func (v NullableTenantIdInfo) Get() *TenantIdInfo {
	return v.value
}

func (v *NullableTenantIdInfo) Set(val *TenantIdInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantIdInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantIdInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantIdInfo(val *TenantIdInfo) *NullableTenantIdInfo {
	return &NullableTenantIdInfo{value: val, isSet: true}
}

func (v NullableTenantIdInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantIdInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


