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

// Wildcard struct for Wildcard
type Wildcard struct {
	Key *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewWildcard instantiates a new Wildcard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWildcard() *Wildcard {
	this := Wildcard{}
	return &this
}

// NewWildcardWithDefaults instantiates a new Wildcard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWildcardWithDefaults() *Wildcard {
	this := Wildcard{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *Wildcard) GetKey() string {
	if o == nil || isNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Wildcard) GetKeyOk() (*string, bool) {
	if o == nil || isNil(o.Key) {
    return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *Wildcard) HasKey() bool {
	if o != nil && !isNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *Wildcard) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Wildcard) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Wildcard) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Wildcard) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *Wildcard) SetValue(v string) {
	o.Value = &v
}

func (o Wildcard) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableWildcard struct {
	value *Wildcard
	isSet bool
}

func (v NullableWildcard) Get() *Wildcard {
	return v.value
}

func (v *NullableWildcard) Set(val *Wildcard) {
	v.value = val
	v.isSet = true
}

func (v NullableWildcard) IsSet() bool {
	return v.isSet
}

func (v *NullableWildcard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWildcard(val *Wildcard) *NullableWildcard {
	return &NullableWildcard{value: val, isSet: true}
}

func (v NullableWildcard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWildcard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


