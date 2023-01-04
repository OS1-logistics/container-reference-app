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

// AttributesConfigResponseData struct for AttributesConfigResponseData
type AttributesConfigResponseData struct {
	// A unique name to represent type of a container.
	Name string `json:"name"`
	// Defines whether this container-type will have actual physical items or not. Such type of containers is termed as leaf container.
	IsLeaf *bool `json:"isLeaf,omitempty"`
	// Define the particular container type is active or not.
	IsActive *bool `json:"isActive,omitempty"`
	AllowedParent *ContainerTypeAllowedParent `json:"allowedParent,omitempty"`
	OwnerAppId *string `json:"ownerAppId,omitempty"`
	Attributes []AttributeConfigGet `json:"attributes"`
	CreatedAt int64 `json:"createdAt"`
	CreatedBy ActionBy `json:"createdBy"`
	UpdatedAt int64 `json:"updatedAt"`
	UpdatedBy ActionBy `json:"updatedBy"`
}

// NewAttributesConfigResponseData instantiates a new AttributesConfigResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttributesConfigResponseData(name string, attributes []AttributeConfigGet, createdAt int64, createdBy ActionBy, updatedAt int64, updatedBy ActionBy) *AttributesConfigResponseData {
	this := AttributesConfigResponseData{}
	this.Name = name
	var isLeaf bool = false
	this.IsLeaf = &isLeaf
	var isActive bool = false
	this.IsActive = &isActive
	this.Attributes = attributes
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.UpdatedAt = updatedAt
	this.UpdatedBy = updatedBy
	return &this
}

// NewAttributesConfigResponseDataWithDefaults instantiates a new AttributesConfigResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttributesConfigResponseDataWithDefaults() *AttributesConfigResponseData {
	this := AttributesConfigResponseData{}
	var isLeaf bool = false
	this.IsLeaf = &isLeaf
	var isActive bool = false
	this.IsActive = &isActive
	return &this
}

// GetName returns the Name field value
func (o *AttributesConfigResponseData) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AttributesConfigResponseData) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AttributesConfigResponseData) SetName(v string) {
	o.Name = v
}

// GetIsLeaf returns the IsLeaf field value if set, zero value otherwise.
func (o *AttributesConfigResponseData) GetIsLeaf() bool {
	if o == nil || isNil(o.IsLeaf) {
		var ret bool
		return ret
	}
	return *o.IsLeaf
}

// GetIsLeafOk returns a tuple with the IsLeaf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributesConfigResponseData) GetIsLeafOk() (*bool, bool) {
	if o == nil || isNil(o.IsLeaf) {
    return nil, false
	}
	return o.IsLeaf, true
}

// HasIsLeaf returns a boolean if a field has been set.
func (o *AttributesConfigResponseData) HasIsLeaf() bool {
	if o != nil && !isNil(o.IsLeaf) {
		return true
	}

	return false
}

// SetIsLeaf gets a reference to the given bool and assigns it to the IsLeaf field.
func (o *AttributesConfigResponseData) SetIsLeaf(v bool) {
	o.IsLeaf = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *AttributesConfigResponseData) GetIsActive() bool {
	if o == nil || isNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributesConfigResponseData) GetIsActiveOk() (*bool, bool) {
	if o == nil || isNil(o.IsActive) {
    return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *AttributesConfigResponseData) HasIsActive() bool {
	if o != nil && !isNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *AttributesConfigResponseData) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetAllowedParent returns the AllowedParent field value if set, zero value otherwise.
func (o *AttributesConfigResponseData) GetAllowedParent() ContainerTypeAllowedParent {
	if o == nil || isNil(o.AllowedParent) {
		var ret ContainerTypeAllowedParent
		return ret
	}
	return *o.AllowedParent
}

// GetAllowedParentOk returns a tuple with the AllowedParent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributesConfigResponseData) GetAllowedParentOk() (*ContainerTypeAllowedParent, bool) {
	if o == nil || isNil(o.AllowedParent) {
    return nil, false
	}
	return o.AllowedParent, true
}

// HasAllowedParent returns a boolean if a field has been set.
func (o *AttributesConfigResponseData) HasAllowedParent() bool {
	if o != nil && !isNil(o.AllowedParent) {
		return true
	}

	return false
}

// SetAllowedParent gets a reference to the given ContainerTypeAllowedParent and assigns it to the AllowedParent field.
func (o *AttributesConfigResponseData) SetAllowedParent(v ContainerTypeAllowedParent) {
	o.AllowedParent = &v
}

// GetOwnerAppId returns the OwnerAppId field value if set, zero value otherwise.
func (o *AttributesConfigResponseData) GetOwnerAppId() string {
	if o == nil || isNil(o.OwnerAppId) {
		var ret string
		return ret
	}
	return *o.OwnerAppId
}

// GetOwnerAppIdOk returns a tuple with the OwnerAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributesConfigResponseData) GetOwnerAppIdOk() (*string, bool) {
	if o == nil || isNil(o.OwnerAppId) {
    return nil, false
	}
	return o.OwnerAppId, true
}

// HasOwnerAppId returns a boolean if a field has been set.
func (o *AttributesConfigResponseData) HasOwnerAppId() bool {
	if o != nil && !isNil(o.OwnerAppId) {
		return true
	}

	return false
}

// SetOwnerAppId gets a reference to the given string and assigns it to the OwnerAppId field.
func (o *AttributesConfigResponseData) SetOwnerAppId(v string) {
	o.OwnerAppId = &v
}

// GetAttributes returns the Attributes field value
func (o *AttributesConfigResponseData) GetAttributes() []AttributeConfigGet {
	if o == nil {
		var ret []AttributeConfigGet
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AttributesConfigResponseData) GetAttributesOk() ([]AttributeConfigGet, bool) {
	if o == nil {
    return nil, false
	}
	return o.Attributes, true
}

// SetAttributes sets field value
func (o *AttributesConfigResponseData) SetAttributes(v []AttributeConfigGet) {
	o.Attributes = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *AttributesConfigResponseData) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AttributesConfigResponseData) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AttributesConfigResponseData) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *AttributesConfigResponseData) GetCreatedBy() ActionBy {
	if o == nil {
		var ret ActionBy
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *AttributesConfigResponseData) GetCreatedByOk() (*ActionBy, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *AttributesConfigResponseData) SetCreatedBy(v ActionBy) {
	o.CreatedBy = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *AttributesConfigResponseData) GetUpdatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *AttributesConfigResponseData) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *AttributesConfigResponseData) SetUpdatedAt(v int64) {
	o.UpdatedAt = v
}

// GetUpdatedBy returns the UpdatedBy field value
func (o *AttributesConfigResponseData) GetUpdatedBy() ActionBy {
	if o == nil {
		var ret ActionBy
		return ret
	}

	return o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value
// and a boolean to check if the value has been set.
func (o *AttributesConfigResponseData) GetUpdatedByOk() (*ActionBy, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UpdatedBy, true
}

// SetUpdatedBy sets field value
func (o *AttributesConfigResponseData) SetUpdatedBy(v ActionBy) {
	o.UpdatedBy = v
}

func (o AttributesConfigResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.IsLeaf) {
		toSerialize["isLeaf"] = o.IsLeaf
	}
	if !isNil(o.IsActive) {
		toSerialize["isActive"] = o.IsActive
	}
	if !isNil(o.AllowedParent) {
		toSerialize["allowedParent"] = o.AllowedParent
	}
	if !isNil(o.OwnerAppId) {
		toSerialize["ownerAppId"] = o.OwnerAppId
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if true {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if true {
		toSerialize["updatedBy"] = o.UpdatedBy
	}
	return json.Marshal(toSerialize)
}

type NullableAttributesConfigResponseData struct {
	value *AttributesConfigResponseData
	isSet bool
}

func (v NullableAttributesConfigResponseData) Get() *AttributesConfigResponseData {
	return v.value
}

func (v *NullableAttributesConfigResponseData) Set(val *AttributesConfigResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableAttributesConfigResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableAttributesConfigResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttributesConfigResponseData(val *AttributesConfigResponseData) *NullableAttributesConfigResponseData {
	return &NullableAttributesConfigResponseData{value: val, isSet: true}
}

func (v NullableAttributesConfigResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttributesConfigResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


