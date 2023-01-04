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

// ReadOnlyAttributes ReadOnly attributes of a container.
type ReadOnlyAttributes struct {
	Id string `json:"id"`
}

// NewReadOnlyAttributes instantiates a new ReadOnlyAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadOnlyAttributes(id string) *ReadOnlyAttributes {
	this := ReadOnlyAttributes{}
	this.Id = id
	return &this
}

// NewReadOnlyAttributesWithDefaults instantiates a new ReadOnlyAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadOnlyAttributesWithDefaults() *ReadOnlyAttributes {
	this := ReadOnlyAttributes{}
	return &this
}

// GetId returns the Id field value
func (o *ReadOnlyAttributes) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ReadOnlyAttributes) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ReadOnlyAttributes) SetId(v string) {
	o.Id = v
}

func (o ReadOnlyAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableReadOnlyAttributes struct {
	value *ReadOnlyAttributes
	isSet bool
}

func (v NullableReadOnlyAttributes) Get() *ReadOnlyAttributes {
	return v.value
}

func (v *NullableReadOnlyAttributes) Set(val *ReadOnlyAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableReadOnlyAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableReadOnlyAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadOnlyAttributes(val *ReadOnlyAttributes) *NullableReadOnlyAttributes {
	return &NullableReadOnlyAttributes{value: val, isSet: true}
}

func (v NullableReadOnlyAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadOnlyAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

