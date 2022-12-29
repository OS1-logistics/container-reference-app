/*
Container Service

**API documentation for Container Service**

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package container_client

import (
	"encoding/json"
)

// MatchPhrasePrefix struct for MatchPhrasePrefix
type MatchPhrasePrefix struct {
	Key *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewMatchPhrasePrefix instantiates a new MatchPhrasePrefix object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMatchPhrasePrefix() *MatchPhrasePrefix {
	this := MatchPhrasePrefix{}
	return &this
}

// NewMatchPhrasePrefixWithDefaults instantiates a new MatchPhrasePrefix object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMatchPhrasePrefixWithDefaults() *MatchPhrasePrefix {
	this := MatchPhrasePrefix{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *MatchPhrasePrefix) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatchPhrasePrefix) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *MatchPhrasePrefix) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *MatchPhrasePrefix) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *MatchPhrasePrefix) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatchPhrasePrefix) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *MatchPhrasePrefix) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *MatchPhrasePrefix) SetValue(v string) {
	o.Value = &v
}

func (o MatchPhrasePrefix) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableMatchPhrasePrefix struct {
	value *MatchPhrasePrefix
	isSet bool
}

func (v NullableMatchPhrasePrefix) Get() *MatchPhrasePrefix {
	return v.value
}

func (v *NullableMatchPhrasePrefix) Set(val *MatchPhrasePrefix) {
	v.value = val
	v.isSet = true
}

func (v NullableMatchPhrasePrefix) IsSet() bool {
	return v.isSet
}

func (v *NullableMatchPhrasePrefix) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMatchPhrasePrefix(val *MatchPhrasePrefix) *NullableMatchPhrasePrefix {
	return &NullableMatchPhrasePrefix{value: val, isSet: true}
}

func (v NullableMatchPhrasePrefix) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMatchPhrasePrefix) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


