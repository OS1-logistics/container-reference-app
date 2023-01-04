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

// Validation Validation rules for a string or number value
type Validation struct {
	Range *Range `json:"range,omitempty"`
	// Regex that the data should match
	Regex *string `json:"regex,omitempty"`
	// Specifies whether the value is required or optional
	Required *bool `json:"required,omitempty"`
}

// NewValidation instantiates a new Validation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidation() *Validation {
	this := Validation{}
	var required bool = false
	this.Required = &required
	return &this
}

// NewValidationWithDefaults instantiates a new Validation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidationWithDefaults() *Validation {
	this := Validation{}
	var required bool = false
	this.Required = &required
	return &this
}

// GetRange returns the Range field value if set, zero value otherwise.
func (o *Validation) GetRange() Range {
	if o == nil || isNil(o.Range) {
		var ret Range
		return ret
	}
	return *o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Validation) GetRangeOk() (*Range, bool) {
	if o == nil || isNil(o.Range) {
    return nil, false
	}
	return o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *Validation) HasRange() bool {
	if o != nil && !isNil(o.Range) {
		return true
	}

	return false
}

// SetRange gets a reference to the given Range and assigns it to the Range field.
func (o *Validation) SetRange(v Range) {
	o.Range = &v
}

// GetRegex returns the Regex field value if set, zero value otherwise.
func (o *Validation) GetRegex() string {
	if o == nil || isNil(o.Regex) {
		var ret string
		return ret
	}
	return *o.Regex
}

// GetRegexOk returns a tuple with the Regex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Validation) GetRegexOk() (*string, bool) {
	if o == nil || isNil(o.Regex) {
    return nil, false
	}
	return o.Regex, true
}

// HasRegex returns a boolean if a field has been set.
func (o *Validation) HasRegex() bool {
	if o != nil && !isNil(o.Regex) {
		return true
	}

	return false
}

// SetRegex gets a reference to the given string and assigns it to the Regex field.
func (o *Validation) SetRegex(v string) {
	o.Regex = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *Validation) GetRequired() bool {
	if o == nil || isNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Validation) GetRequiredOk() (*bool, bool) {
	if o == nil || isNil(o.Required) {
    return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *Validation) HasRequired() bool {
	if o != nil && !isNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *Validation) SetRequired(v bool) {
	o.Required = &v
}

func (o Validation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Range) {
		toSerialize["range"] = o.Range
	}
	if !isNil(o.Regex) {
		toSerialize["regex"] = o.Regex
	}
	if !isNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	return json.Marshal(toSerialize)
}

type NullableValidation struct {
	value *Validation
	isSet bool
}

func (v NullableValidation) Get() *Validation {
	return v.value
}

func (v *NullableValidation) Set(val *Validation) {
	v.value = val
	v.isSet = true
}

func (v NullableValidation) IsSet() bool {
	return v.isSet
}

func (v *NullableValidation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidation(val *Validation) *NullableValidation {
	return &NullableValidation{value: val, isSet: true}
}

func (v NullableValidation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


