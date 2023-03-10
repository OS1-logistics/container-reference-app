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

// ContainerTypeBasicProperties struct for ContainerTypeBasicProperties
type ContainerTypeBasicProperties struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewContainerTypeBasicProperties instantiates a new ContainerTypeBasicProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerTypeBasicProperties() *ContainerTypeBasicProperties {
	this := ContainerTypeBasicProperties{}
	return &this
}

// NewContainerTypeBasicPropertiesWithDefaults instantiates a new ContainerTypeBasicProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerTypeBasicPropertiesWithDefaults() *ContainerTypeBasicProperties {
	this := ContainerTypeBasicProperties{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ContainerTypeBasicProperties) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerTypeBasicProperties) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ContainerTypeBasicProperties) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ContainerTypeBasicProperties) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ContainerTypeBasicProperties) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerTypeBasicProperties) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ContainerTypeBasicProperties) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ContainerTypeBasicProperties) SetName(v string) {
	o.Name = &v
}

func (o ContainerTypeBasicProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableContainerTypeBasicProperties struct {
	value *ContainerTypeBasicProperties
	isSet bool
}

func (v NullableContainerTypeBasicProperties) Get() *ContainerTypeBasicProperties {
	return v.value
}

func (v *NullableContainerTypeBasicProperties) Set(val *ContainerTypeBasicProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerTypeBasicProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerTypeBasicProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerTypeBasicProperties(val *ContainerTypeBasicProperties) *NullableContainerTypeBasicProperties {
	return &NullableContainerTypeBasicProperties{value: val, isSet: true}
}

func (v NullableContainerTypeBasicProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerTypeBasicProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


