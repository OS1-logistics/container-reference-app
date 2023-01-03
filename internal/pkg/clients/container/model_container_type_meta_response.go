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

// ContainerTypeMetaResponse struct for ContainerTypeMetaResponse
type ContainerTypeMetaResponse struct {
	// uniqueCode for each container type
	EntityCode *string `json:"entityCode,omitempty"`
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

// NewContainerTypeMetaResponse instantiates a new ContainerTypeMetaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerTypeMetaResponse(name string, createdAt int64, createdBy ActionBy, updatedAt int64, updatedBy ActionBy) *ContainerTypeMetaResponse {
	this := ContainerTypeMetaResponse{}
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

// NewContainerTypeMetaResponseWithDefaults instantiates a new ContainerTypeMetaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerTypeMetaResponseWithDefaults() *ContainerTypeMetaResponse {
	this := ContainerTypeMetaResponse{}
	var isLeaf bool = false
	this.IsLeaf = &isLeaf
	var isActive bool = false
	this.IsActive = &isActive
	return &this
}

// GetEntityCode returns the EntityCode field value if set, zero value otherwise.
func (o *ContainerTypeMetaResponse) GetEntityCode() string {
	if o == nil || o.EntityCode == nil {
		var ret string
		return ret
	}
	return *o.EntityCode
}

// GetEntityCodeOk returns a tuple with the EntityCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerTypeMetaResponse) GetEntityCodeOk() (*string, bool) {
	if o == nil || o.EntityCode == nil {
		return nil, false
	}
	return o.EntityCode, true
}

// HasEntityCode returns a boolean if a field has been set.
func (o *ContainerTypeMetaResponse) HasEntityCode() bool {
	if o != nil && o.EntityCode != nil {
		return true
	}

	return false
}

// SetEntityCode gets a reference to the given string and assigns it to the EntityCode field.
func (o *ContainerTypeMetaResponse) SetEntityCode(v string) {
	o.EntityCode = &v
}

// GetName returns the Name field value
func (o *ContainerTypeMetaResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContainerTypeMetaResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContainerTypeMetaResponse) SetName(v string) {
	o.Name = v
}

// GetIsLeaf returns the IsLeaf field value if set, zero value otherwise.
func (o *ContainerTypeMetaResponse) GetIsLeaf() bool {
	if o == nil || o.IsLeaf == nil {
		var ret bool
		return ret
	}
	return *o.IsLeaf
}

// GetIsLeafOk returns a tuple with the IsLeaf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerTypeMetaResponse) GetIsLeafOk() (*bool, bool) {
	if o == nil || o.IsLeaf == nil {
		return nil, false
	}
	return o.IsLeaf, true
}

// HasIsLeaf returns a boolean if a field has been set.
func (o *ContainerTypeMetaResponse) HasIsLeaf() bool {
	if o != nil && o.IsLeaf != nil {
		return true
	}

	return false
}

// SetIsLeaf gets a reference to the given bool and assigns it to the IsLeaf field.
func (o *ContainerTypeMetaResponse) SetIsLeaf(v bool) {
	o.IsLeaf = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *ContainerTypeMetaResponse) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerTypeMetaResponse) GetIsActiveOk() (*bool, bool) {
	if o == nil || o.IsActive == nil {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *ContainerTypeMetaResponse) HasIsActive() bool {
	if o != nil && o.IsActive != nil {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *ContainerTypeMetaResponse) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetAllowedParent returns the AllowedParent field value if set, zero value otherwise.
func (o *ContainerTypeMetaResponse) GetAllowedParent() ContainerTypeAllowedParent {
	if o == nil || o.AllowedParent == nil {
		var ret ContainerTypeAllowedParent
		return ret
	}
	return *o.AllowedParent
}

// GetAllowedParentOk returns a tuple with the AllowedParent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerTypeMetaResponse) GetAllowedParentOk() (*ContainerTypeAllowedParent, bool) {
	if o == nil || o.AllowedParent == nil {
		return nil, false
	}
	return o.AllowedParent, true
}

// HasAllowedParent returns a boolean if a field has been set.
func (o *ContainerTypeMetaResponse) HasAllowedParent() bool {
	if o != nil && o.AllowedParent != nil {
		return true
	}

	return false
}

// SetAllowedParent gets a reference to the given ContainerTypeAllowedParent and assigns it to the AllowedParent field.
func (o *ContainerTypeMetaResponse) SetAllowedParent(v ContainerTypeAllowedParent) {
	o.AllowedParent = &v
}

// GetOwnerAppId returns the OwnerAppId field value if set, zero value otherwise.
func (o *ContainerTypeMetaResponse) GetOwnerAppId() string {
	if o == nil || o.OwnerAppId == nil {
		var ret string
		return ret
	}
	return *o.OwnerAppId
}

// GetOwnerAppIdOk returns a tuple with the OwnerAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerTypeMetaResponse) GetOwnerAppIdOk() (*string, bool) {
	if o == nil || o.OwnerAppId == nil {
		return nil, false
	}
	return o.OwnerAppId, true
}

// HasOwnerAppId returns a boolean if a field has been set.
func (o *ContainerTypeMetaResponse) HasOwnerAppId() bool {
	if o != nil && o.OwnerAppId != nil {
		return true
	}

	return false
}

// SetOwnerAppId gets a reference to the given string and assigns it to the OwnerAppId field.
func (o *ContainerTypeMetaResponse) SetOwnerAppId(v string) {
	o.OwnerAppId = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ContainerTypeMetaResponse) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ContainerTypeMetaResponse) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ContainerTypeMetaResponse) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *ContainerTypeMetaResponse) GetCreatedBy() ActionBy {
	if o == nil {
		var ret ActionBy
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *ContainerTypeMetaResponse) GetCreatedByOk() (*ActionBy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *ContainerTypeMetaResponse) SetCreatedBy(v ActionBy) {
	o.CreatedBy = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ContainerTypeMetaResponse) GetUpdatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ContainerTypeMetaResponse) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ContainerTypeMetaResponse) SetUpdatedAt(v int64) {
	o.UpdatedAt = v
}

// GetUpdatedBy returns the UpdatedBy field value
func (o *ContainerTypeMetaResponse) GetUpdatedBy() ActionBy {
	if o == nil {
		var ret ActionBy
		return ret
	}

	return o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value
// and a boolean to check if the value has been set.
func (o *ContainerTypeMetaResponse) GetUpdatedByOk() (*ActionBy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedBy, true
}

// SetUpdatedBy sets field value
func (o *ContainerTypeMetaResponse) SetUpdatedBy(v ActionBy) {
	o.UpdatedBy = v
}

func (o ContainerTypeMetaResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EntityCode != nil {
		toSerialize["entityCode"] = o.EntityCode
	}
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

type NullableContainerTypeMetaResponse struct {
	value *ContainerTypeMetaResponse
	isSet bool
}

func (v NullableContainerTypeMetaResponse) Get() *ContainerTypeMetaResponse {
	return v.value
}

func (v *NullableContainerTypeMetaResponse) Set(val *ContainerTypeMetaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerTypeMetaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerTypeMetaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerTypeMetaResponse(val *ContainerTypeMetaResponse) *NullableContainerTypeMetaResponse {
	return &NullableContainerTypeMetaResponse{value: val, isSet: true}
}

func (v NullableContainerTypeMetaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerTypeMetaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


