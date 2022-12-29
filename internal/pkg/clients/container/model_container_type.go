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

// ContainerType struct for ContainerType
type ContainerType struct {
	// A unique name to represent type of a container.
	Name string `json:"name"`
	// Defines whether this container-type will have actual physical items or not. Such type of containers is termed as leaf container.
	IsLeaf *bool `json:"isLeaf,omitempty"`
	// Define the particular container type is active or not.
	IsActive *bool `json:"isActive,omitempty"`
	AllowedParent *ContainerTypeAllowedParent `json:"allowedParent,omitempty"`
	OwnerAppId *string `json:"ownerAppId,omitempty"`
}

// NewContainerType instantiates a new ContainerType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerType(name string) *ContainerType {
	this := ContainerType{}
	this.Name = name
	var isLeaf bool = false
	this.IsLeaf = &isLeaf
	var isActive bool = false
	this.IsActive = &isActive
	return &this
}

// NewContainerTypeWithDefaults instantiates a new ContainerType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerTypeWithDefaults() *ContainerType {
	this := ContainerType{}
	var isLeaf bool = false
	this.IsLeaf = &isLeaf
	var isActive bool = false
	this.IsActive = &isActive
	return &this
}

// GetName returns the Name field value
func (o *ContainerType) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContainerType) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContainerType) SetName(v string) {
	o.Name = v
}

// GetIsLeaf returns the IsLeaf field value if set, zero value otherwise.
func (o *ContainerType) GetIsLeaf() bool {
	if o == nil || o.IsLeaf == nil {
		var ret bool
		return ret
	}
	return *o.IsLeaf
}

// GetIsLeafOk returns a tuple with the IsLeaf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerType) GetIsLeafOk() (*bool, bool) {
	if o == nil || o.IsLeaf == nil {
		return nil, false
	}
	return o.IsLeaf, true
}

// HasIsLeaf returns a boolean if a field has been set.
func (o *ContainerType) HasIsLeaf() bool {
	if o != nil && o.IsLeaf != nil {
		return true
	}

	return false
}

// SetIsLeaf gets a reference to the given bool and assigns it to the IsLeaf field.
func (o *ContainerType) SetIsLeaf(v bool) {
	o.IsLeaf = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *ContainerType) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerType) GetIsActiveOk() (*bool, bool) {
	if o == nil || o.IsActive == nil {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *ContainerType) HasIsActive() bool {
	if o != nil && o.IsActive != nil {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *ContainerType) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetAllowedParent returns the AllowedParent field value if set, zero value otherwise.
func (o *ContainerType) GetAllowedParent() ContainerTypeAllowedParent {
	if o == nil || o.AllowedParent == nil {
		var ret ContainerTypeAllowedParent
		return ret
	}
	return *o.AllowedParent
}

// GetAllowedParentOk returns a tuple with the AllowedParent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerType) GetAllowedParentOk() (*ContainerTypeAllowedParent, bool) {
	if o == nil || o.AllowedParent == nil {
		return nil, false
	}
	return o.AllowedParent, true
}

// HasAllowedParent returns a boolean if a field has been set.
func (o *ContainerType) HasAllowedParent() bool {
	if o != nil && o.AllowedParent != nil {
		return true
	}

	return false
}

// SetAllowedParent gets a reference to the given ContainerTypeAllowedParent and assigns it to the AllowedParent field.
func (o *ContainerType) SetAllowedParent(v ContainerTypeAllowedParent) {
	o.AllowedParent = &v
}

// GetOwnerAppId returns the OwnerAppId field value if set, zero value otherwise.
func (o *ContainerType) GetOwnerAppId() string {
	if o == nil || o.OwnerAppId == nil {
		var ret string
		return ret
	}
	return *o.OwnerAppId
}

// GetOwnerAppIdOk returns a tuple with the OwnerAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerType) GetOwnerAppIdOk() (*string, bool) {
	if o == nil || o.OwnerAppId == nil {
		return nil, false
	}
	return o.OwnerAppId, true
}

// HasOwnerAppId returns a boolean if a field has been set.
func (o *ContainerType) HasOwnerAppId() bool {
	if o != nil && o.OwnerAppId != nil {
		return true
	}

	return false
}

// SetOwnerAppId gets a reference to the given string and assigns it to the OwnerAppId field.
func (o *ContainerType) SetOwnerAppId(v string) {
	o.OwnerAppId = &v
}

func (o ContainerType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.IsLeaf != nil {
		toSerialize["isLeaf"] = o.IsLeaf
	}
	if o.IsActive != nil {
		toSerialize["isActive"] = o.IsActive
	}
	if o.AllowedParent != nil {
		toSerialize["allowedParent"] = o.AllowedParent
	}
	if o.OwnerAppId != nil {
		toSerialize["ownerAppId"] = o.OwnerAppId
	}
	return json.Marshal(toSerialize)
}

type NullableContainerType struct {
	value *ContainerType
	isSet bool
}

func (v NullableContainerType) Get() *ContainerType {
	return v.value
}

func (v *NullableContainerType) Set(val *ContainerType) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerType) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerType(val *ContainerType) *NullableContainerType {
	return &NullableContainerType{value: val, isSet: true}
}

func (v NullableContainerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


