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

// TenantCreateRequest struct for TenantCreateRequest
type TenantCreateRequest struct {
	// Name of the Tenant.
	Name string `json:"name"`
	// tenant id of the tenant
	TenantId string `json:"tenantId"`
	// Tenant FPA Tenant UUID having format tenantId:uuid
	Id string `json:"id"`
	// Display name of the Tenant.
	DisplayName *string `json:"displayName,omitempty"`
	// Metadata associated with the Tenant.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// whether Tenant is enabled or not.
	Enabled *bool `json:"enabled,omitempty"`
	AbacConfig *AbacConfigRequest `json:"abacConfig,omitempty"`
}

// NewTenantCreateRequest instantiates a new TenantCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantCreateRequest(name string, tenantId string, id string) *TenantCreateRequest {
	this := TenantCreateRequest{}
	this.Name = name
	this.TenantId = tenantId
	this.Id = id
	return &this
}

// NewTenantCreateRequestWithDefaults instantiates a new TenantCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantCreateRequestWithDefaults() *TenantCreateRequest {
	this := TenantCreateRequest{}
	return &this
}

// GetName returns the Name field value
func (o *TenantCreateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TenantCreateRequest) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TenantCreateRequest) SetName(v string) {
	o.Name = v
}

// GetTenantId returns the TenantId field value
func (o *TenantCreateRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *TenantCreateRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *TenantCreateRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetId returns the Id field value
func (o *TenantCreateRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TenantCreateRequest) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TenantCreateRequest) SetId(v string) {
	o.Id = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *TenantCreateRequest) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantCreateRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
    return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *TenantCreateRequest) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *TenantCreateRequest) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *TenantCreateRequest) GetMetadata() map[string]interface{} {
	if o == nil || isNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantCreateRequest) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Metadata) {
    return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *TenantCreateRequest) HasMetadata() bool {
	if o != nil && !isNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *TenantCreateRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *TenantCreateRequest) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantCreateRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *TenantCreateRequest) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *TenantCreateRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAbacConfig returns the AbacConfig field value if set, zero value otherwise.
func (o *TenantCreateRequest) GetAbacConfig() AbacConfigRequest {
	if o == nil || isNil(o.AbacConfig) {
		var ret AbacConfigRequest
		return ret
	}
	return *o.AbacConfig
}

// GetAbacConfigOk returns a tuple with the AbacConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantCreateRequest) GetAbacConfigOk() (*AbacConfigRequest, bool) {
	if o == nil || isNil(o.AbacConfig) {
    return nil, false
	}
	return o.AbacConfig, true
}

// HasAbacConfig returns a boolean if a field has been set.
func (o *TenantCreateRequest) HasAbacConfig() bool {
	if o != nil && !isNil(o.AbacConfig) {
		return true
	}

	return false
}

// SetAbacConfig gets a reference to the given AbacConfigRequest and assigns it to the AbacConfig field.
func (o *TenantCreateRequest) SetAbacConfig(v AbacConfigRequest) {
	o.AbacConfig = &v
}

func (o TenantCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["tenantId"] = o.TenantId
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !isNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.AbacConfig) {
		toSerialize["abacConfig"] = o.AbacConfig
	}
	return json.Marshal(toSerialize)
}

type NullableTenantCreateRequest struct {
	value *TenantCreateRequest
	isSet bool
}

func (v NullableTenantCreateRequest) Get() *TenantCreateRequest {
	return v.value
}

func (v *NullableTenantCreateRequest) Set(val *TenantCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantCreateRequest(val *TenantCreateRequest) *NullableTenantCreateRequest {
	return &NullableTenantCreateRequest{value: val, isSet: true}
}

func (v NullableTenantCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

