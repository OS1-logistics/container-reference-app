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

// BaseContainerTypeCreateRequest struct for BaseContainerTypeCreateRequest
type BaseContainerTypeCreateRequest struct {
	// A unique name to represent type of a container.
	Name string `json:"name"`
	// Defines whether this container-type will have actual physical items or not. Such type of containers is termed as leaf container.
	IsLeaf *bool `json:"isLeaf,omitempty"`
	AllowedParent *ContainerTypeAllowedParent `json:"allowedParent,omitempty"`
}

// NewBaseContainerTypeCreateRequest instantiates a new BaseContainerTypeCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseContainerTypeCreateRequest(name string) *BaseContainerTypeCreateRequest {
	this := BaseContainerTypeCreateRequest{}
	this.Name = name
	var isLeaf bool = false
	this.IsLeaf = &isLeaf
	return &this
}

// NewBaseContainerTypeCreateRequestWithDefaults instantiates a new BaseContainerTypeCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseContainerTypeCreateRequestWithDefaults() *BaseContainerTypeCreateRequest {
	this := BaseContainerTypeCreateRequest{}
	var isLeaf bool = false
	this.IsLeaf = &isLeaf
	return &this
}

// GetName returns the Name field value
func (o *BaseContainerTypeCreateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BaseContainerTypeCreateRequest) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BaseContainerTypeCreateRequest) SetName(v string) {
	o.Name = v
}

// GetIsLeaf returns the IsLeaf field value if set, zero value otherwise.
func (o *BaseContainerTypeCreateRequest) GetIsLeaf() bool {
	if o == nil || isNil(o.IsLeaf) {
		var ret bool
		return ret
	}
	return *o.IsLeaf
}

// GetIsLeafOk returns a tuple with the IsLeaf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseContainerTypeCreateRequest) GetIsLeafOk() (*bool, bool) {
	if o == nil || isNil(o.IsLeaf) {
    return nil, false
	}
	return o.IsLeaf, true
}

// HasIsLeaf returns a boolean if a field has been set.
func (o *BaseContainerTypeCreateRequest) HasIsLeaf() bool {
	if o != nil && !isNil(o.IsLeaf) {
		return true
	}

	return false
}

// SetIsLeaf gets a reference to the given bool and assigns it to the IsLeaf field.
func (o *BaseContainerTypeCreateRequest) SetIsLeaf(v bool) {
	o.IsLeaf = &v
}

// GetAllowedParent returns the AllowedParent field value if set, zero value otherwise.
func (o *BaseContainerTypeCreateRequest) GetAllowedParent() ContainerTypeAllowedParent {
	if o == nil || isNil(o.AllowedParent) {
		var ret ContainerTypeAllowedParent
		return ret
	}
	return *o.AllowedParent
}

// GetAllowedParentOk returns a tuple with the AllowedParent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseContainerTypeCreateRequest) GetAllowedParentOk() (*ContainerTypeAllowedParent, bool) {
	if o == nil || isNil(o.AllowedParent) {
    return nil, false
	}
	return o.AllowedParent, true
}

// HasAllowedParent returns a boolean if a field has been set.
func (o *BaseContainerTypeCreateRequest) HasAllowedParent() bool {
	if o != nil && !isNil(o.AllowedParent) {
		return true
	}

	return false
}

// SetAllowedParent gets a reference to the given ContainerTypeAllowedParent and assigns it to the AllowedParent field.
func (o *BaseContainerTypeCreateRequest) SetAllowedParent(v ContainerTypeAllowedParent) {
	o.AllowedParent = &v
}

func (o BaseContainerTypeCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.IsLeaf) {
		toSerialize["isLeaf"] = o.IsLeaf
	}
	if !isNil(o.AllowedParent) {
		toSerialize["allowedParent"] = o.AllowedParent
	}
	return json.Marshal(toSerialize)
}

type NullableBaseContainerTypeCreateRequest struct {
	value *BaseContainerTypeCreateRequest
	isSet bool
}

func (v NullableBaseContainerTypeCreateRequest) Get() *BaseContainerTypeCreateRequest {
	return v.value
}

func (v *NullableBaseContainerTypeCreateRequest) Set(val *BaseContainerTypeCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseContainerTypeCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseContainerTypeCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseContainerTypeCreateRequest(val *BaseContainerTypeCreateRequest) *NullableBaseContainerTypeCreateRequest {
	return &NullableBaseContainerTypeCreateRequest{value: val, isSet: true}
}

func (v NullableBaseContainerTypeCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseContainerTypeCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

