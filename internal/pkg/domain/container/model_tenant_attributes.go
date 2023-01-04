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

// TenantAttributes Optional attributes provided by tenants, that can be updated
type TenantAttributes struct {
	// Specifies the time to live for a data entry in the database. Specified for a tenant.
	TerminalTTL *string `json:"terminalTTL,omitempty"`
}

// NewTenantAttributes instantiates a new TenantAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantAttributes() *TenantAttributes {
	this := TenantAttributes{}
	return &this
}

// NewTenantAttributesWithDefaults instantiates a new TenantAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantAttributesWithDefaults() *TenantAttributes {
	this := TenantAttributes{}
	return &this
}

// GetTerminalTTL returns the TerminalTTL field value if set, zero value otherwise.
func (o *TenantAttributes) GetTerminalTTL() string {
	if o == nil || isNil(o.TerminalTTL) {
		var ret string
		return ret
	}
	return *o.TerminalTTL
}

// GetTerminalTTLOk returns a tuple with the TerminalTTL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantAttributes) GetTerminalTTLOk() (*string, bool) {
	if o == nil || isNil(o.TerminalTTL) {
    return nil, false
	}
	return o.TerminalTTL, true
}

// HasTerminalTTL returns a boolean if a field has been set.
func (o *TenantAttributes) HasTerminalTTL() bool {
	if o != nil && !isNil(o.TerminalTTL) {
		return true
	}

	return false
}

// SetTerminalTTL gets a reference to the given string and assigns it to the TerminalTTL field.
func (o *TenantAttributes) SetTerminalTTL(v string) {
	o.TerminalTTL = &v
}

func (o TenantAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TerminalTTL) {
		toSerialize["terminalTTL"] = o.TerminalTTL
	}
	return json.Marshal(toSerialize)
}

type NullableTenantAttributes struct {
	value *TenantAttributes
	isSet bool
}

func (v NullableTenantAttributes) Get() *TenantAttributes {
	return v.value
}

func (v *NullableTenantAttributes) Set(val *TenantAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantAttributes(val *TenantAttributes) *NullableTenantAttributes {
	return &NullableTenantAttributes{value: val, isSet: true}
}

func (v NullableTenantAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

