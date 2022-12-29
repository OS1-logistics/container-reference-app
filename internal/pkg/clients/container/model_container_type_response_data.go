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

// ContainerTypeResponseData struct for ContainerTypeResponseData
type ContainerTypeResponseData struct {
	// A unique name to represent type of a container.
	Name string `json:"name"`
	// Defines whether this container-type will have actual physical items or not. Such type of containers is termed as leaf container.
	IsLeaf *bool `json:"isLeaf,omitempty"`
	// Define the particular container type is active or not.
	IsActive *bool `json:"isActive,omitempty"`
	AllowedParent *ContainerTypeAllowedParent `json:"allowedParent,omitempty"`
	OwnerAppId *string `json:"ownerAppId,omitempty"`
	CreatedAt int64 `json:"createdAt"`
	CreatedBy ActionBy `json:"createdBy"`
	UpdatedAt int64 `json:"updatedAt"`
	UpdatedBy ActionBy `json:"updatedBy"`
}

// NewContainerTypeResponseData instantiates a new ContainerTypeResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerTypeResponseData(name string, createdAt int64, createdBy ActionBy, updatedAt int64, updatedBy ActionBy) *ContainerTypeResponseData {
	this := ContainerTypeResponseData{}
	this.Name = name
	var isLeaf bool = false
	this.IsLeaf = &isLeaf
	var isActive bool = false
	this.IsActive = &isActive
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.UpdatedAt = updatedAt
	this.UpdatedBy = updatedBy
	return &this
}

// NewContainerTypeResponseDataWithDefaults instantiates a new ContainerTypeResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerTypeResponseDataWithDefaults() *ContainerTypeResponseData {
	this := ContainerTypeResponseData{}
	var isLeaf bool = false
	this.IsLeaf = &isLeaf
	var isActive bool = false
	this.IsActive = &isActive
	return &this
}

// GetName returns the Name field value
func (o *ContainerTypeResponseData) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContainerTypeResponseData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContainerTypeResponseData) SetName(v string) {
	o.Name = v
}

// GetIsLeaf returns the IsLeaf field value if set, zero value otherwise.
func (o *ContainerTypeResponseData) GetIsLeaf() bool {
	if o == nil || o.IsLeaf == nil {
		var ret bool
		return ret
	}
	return *o.IsLeaf
}

// GetIsLeafOk returns a tuple with the IsLeaf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerTypeResponseData) GetIsLeafOk() (*bool, bool) {
	if o == nil || o.IsLeaf == nil {
		return nil, false
	}
	return o.IsLeaf, true
}

// HasIsLeaf returns a boolean if a field has been set.
func (o *ContainerTypeResponseData) HasIsLeaf() bool {
	if o != nil && o.IsLeaf != nil {
		return true
	}

	return false
}

// SetIsLeaf gets a reference to the given bool and assigns it to the IsLeaf field.
func (o *ContainerTypeResponseData) SetIsLeaf(v bool) {
	o.IsLeaf = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *ContainerTypeResponseData) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerTypeResponseData) GetIsActiveOk() (*bool, bool) {
	if o == nil || o.IsActive == nil {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *ContainerTypeResponseData) HasIsActive() bool {
	if o != nil && o.IsActive != nil {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *ContainerTypeResponseData) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetAllowedParent returns the AllowedParent field value if set, zero value otherwise.
func (o *ContainerTypeResponseData) GetAllowedParent() ContainerTypeAllowedParent {
	if o == nil || o.AllowedParent == nil {
		var ret ContainerTypeAllowedParent
		return ret
	}
	return *o.AllowedParent
}

// GetAllowedParentOk returns a tuple with the AllowedParent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerTypeResponseData) GetAllowedParentOk() (*ContainerTypeAllowedParent, bool) {
	if o == nil || o.AllowedParent == nil {
		return nil, false
	}
	return o.AllowedParent, true
}

// HasAllowedParent returns a boolean if a field has been set.
func (o *ContainerTypeResponseData) HasAllowedParent() bool {
	if o != nil && o.AllowedParent != nil {
		return true
	}

	return false
}

// SetAllowedParent gets a reference to the given ContainerTypeAllowedParent and assigns it to the AllowedParent field.
func (o *ContainerTypeResponseData) SetAllowedParent(v ContainerTypeAllowedParent) {
	o.AllowedParent = &v
}

// GetOwnerAppId returns the OwnerAppId field value if set, zero value otherwise.
func (o *ContainerTypeResponseData) GetOwnerAppId() string {
	if o == nil || o.OwnerAppId == nil {
		var ret string
		return ret
	}
	return *o.OwnerAppId
}

// GetOwnerAppIdOk returns a tuple with the OwnerAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerTypeResponseData) GetOwnerAppIdOk() (*string, bool) {
	if o == nil || o.OwnerAppId == nil {
		return nil, false
	}
	return o.OwnerAppId, true
}

// HasOwnerAppId returns a boolean if a field has been set.
func (o *ContainerTypeResponseData) HasOwnerAppId() bool {
	if o != nil && o.OwnerAppId != nil {
		return true
	}

	return false
}

// SetOwnerAppId gets a reference to the given string and assigns it to the OwnerAppId field.
func (o *ContainerTypeResponseData) SetOwnerAppId(v string) {
	o.OwnerAppId = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ContainerTypeResponseData) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ContainerTypeResponseData) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ContainerTypeResponseData) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *ContainerTypeResponseData) GetCreatedBy() ActionBy {
	if o == nil {
		var ret ActionBy
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *ContainerTypeResponseData) GetCreatedByOk() (*ActionBy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *ContainerTypeResponseData) SetCreatedBy(v ActionBy) {
	o.CreatedBy = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ContainerTypeResponseData) GetUpdatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ContainerTypeResponseData) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ContainerTypeResponseData) SetUpdatedAt(v int64) {
	o.UpdatedAt = v
}

// GetUpdatedBy returns the UpdatedBy field value
func (o *ContainerTypeResponseData) GetUpdatedBy() ActionBy {
	if o == nil {
		var ret ActionBy
		return ret
	}

	return o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value
// and a boolean to check if the value has been set.
func (o *ContainerTypeResponseData) GetUpdatedByOk() (*ActionBy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedBy, true
}

// SetUpdatedBy sets field value
func (o *ContainerTypeResponseData) SetUpdatedBy(v ActionBy) {
	o.UpdatedBy = v
}

func (o ContainerTypeResponseData) MarshalJSON() ([]byte, error) {
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

type NullableContainerTypeResponseData struct {
	value *ContainerTypeResponseData
	isSet bool
}

func (v NullableContainerTypeResponseData) Get() *ContainerTypeResponseData {
	return v.value
}

func (v *NullableContainerTypeResponseData) Set(val *ContainerTypeResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerTypeResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerTypeResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerTypeResponseData(val *ContainerTypeResponseData) *NullableContainerTypeResponseData {
	return &NullableContainerTypeResponseData{value: val, isSet: true}
}

func (v NullableContainerTypeResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerTypeResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


