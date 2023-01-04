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

// TenantConfig Attributes associated with each tenant.
type TenantConfig struct {
	Name *string `json:"name,omitempty"`
	IsActive *bool `json:"isActive,omitempty"`
	Attributes *TenantAttributes `json:"attributes,omitempty"`
}

// NewTenantConfig instantiates a new TenantConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantConfig() *TenantConfig {
	this := TenantConfig{}
	return &this
}

// NewTenantConfigWithDefaults instantiates a new TenantConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantConfigWithDefaults() *TenantConfig {
	this := TenantConfig{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TenantConfig) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantConfig) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TenantConfig) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TenantConfig) SetName(v string) {
	o.Name = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *TenantConfig) GetIsActive() bool {
	if o == nil || isNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantConfig) GetIsActiveOk() (*bool, bool) {
	if o == nil || isNil(o.IsActive) {
    return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *TenantConfig) HasIsActive() bool {
	if o != nil && !isNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *TenantConfig) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TenantConfig) GetAttributes() TenantAttributes {
	if o == nil || isNil(o.Attributes) {
		var ret TenantAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantConfig) GetAttributesOk() (*TenantAttributes, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TenantConfig) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given TenantAttributes and assigns it to the Attributes field.
func (o *TenantConfig) SetAttributes(v TenantAttributes) {
	o.Attributes = &v
}

func (o TenantConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.IsActive) {
		toSerialize["isActive"] = o.IsActive
	}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableTenantConfig struct {
	value *TenantConfig
	isSet bool
}

func (v NullableTenantConfig) Get() *TenantConfig {
	return v.value
}

func (v *NullableTenantConfig) Set(val *TenantConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantConfig(val *TenantConfig) *NullableTenantConfig {
	return &NullableTenantConfig{value: val, isSet: true}
}

func (v NullableTenantConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

