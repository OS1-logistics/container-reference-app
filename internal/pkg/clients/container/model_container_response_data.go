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

// ContainerResponseData struct for ContainerResponseData
type ContainerResponseData struct {
	Id string `json:"id"`
	// Container type defined using container type configuration APIs.
	ContainerType string `json:"containerType"`
	// Defines whether container is re-usable or not
	IsReusable *bool `json:"isReusable,omitempty"`
	// List of itmes can be added into container only if container-type is a leaf.
	Items []Item `json:"items,omitempty"`
	// Field to represent tracking ID (such as AWB) for a container which is usually pasted on the container and used for scanning it. There can be multiple tracking IDs for a container, owned by different operators and hence it will be stored as a list.
	TrackingDetails []TrackingDetail `json:"trackingDetails,omitempty"`
	// Field to represent container ID of the parent container. During the containization process, this value is assigned to represent which container contains this container.
	ParentId *string `json:"parentId,omitempty"`
	// A map to provide values for the attributes defined in Container-type configuration APIs.
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// Represents whether this container can contain hazardous materials or not.
	IsHazmat *bool `json:"isHazmat,omitempty"`
	// Defines whether container can be put into other containers or not
	IsContainerizable *bool `json:"isContainerizable,omitempty"`
	CreatedAt int64 `json:"createdAt"`
	CreatedBy ActionBy `json:"createdBy"`
	UpdatedAt int64 `json:"updatedAt"`
	UpdatedBy ActionBy `json:"updatedBy"`
}

// NewContainerResponseData instantiates a new ContainerResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerResponseData(id string, containerType string, createdAt int64, createdBy ActionBy, updatedAt int64, updatedBy ActionBy) *ContainerResponseData {
	this := ContainerResponseData{}
	this.Id = id
	this.ContainerType = containerType
	var isReusable bool = false
	this.IsReusable = &isReusable
	var isHazmat bool = false
	this.IsHazmat = &isHazmat
	var isContainerizable bool = true
	this.IsContainerizable = &isContainerizable
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.UpdatedAt = updatedAt
	this.UpdatedBy = updatedBy
	return &this
}

// NewContainerResponseDataWithDefaults instantiates a new ContainerResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerResponseDataWithDefaults() *ContainerResponseData {
	this := ContainerResponseData{}
	var isReusable bool = false
	this.IsReusable = &isReusable
	var isHazmat bool = false
	this.IsHazmat = &isHazmat
	var isContainerizable bool = true
	this.IsContainerizable = &isContainerizable
	return &this
}

// GetId returns the Id field value
func (o *ContainerResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ContainerResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ContainerResponseData) SetId(v string) {
	o.Id = v
}

// GetContainerType returns the ContainerType field value
func (o *ContainerResponseData) GetContainerType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContainerType
}

// GetContainerTypeOk returns a tuple with the ContainerType field value
// and a boolean to check if the value has been set.
func (o *ContainerResponseData) GetContainerTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContainerType, true
}

// SetContainerType sets field value
func (o *ContainerResponseData) SetContainerType(v string) {
	o.ContainerType = v
}

// GetIsReusable returns the IsReusable field value if set, zero value otherwise.
func (o *ContainerResponseData) GetIsReusable() bool {
	if o == nil || o.IsReusable == nil {
		var ret bool
		return ret
	}
	return *o.IsReusable
}

// GetIsReusableOk returns a tuple with the IsReusable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerResponseData) GetIsReusableOk() (*bool, bool) {
	if o == nil || o.IsReusable == nil {
		return nil, false
	}
	return o.IsReusable, true
}

// HasIsReusable returns a boolean if a field has been set.
func (o *ContainerResponseData) HasIsReusable() bool {
	if o != nil && o.IsReusable != nil {
		return true
	}

	return false
}

// SetIsReusable gets a reference to the given bool and assigns it to the IsReusable field.
func (o *ContainerResponseData) SetIsReusable(v bool) {
	o.IsReusable = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ContainerResponseData) GetItems() []Item {
	if o == nil || o.Items == nil {
		var ret []Item
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerResponseData) GetItemsOk() ([]Item, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ContainerResponseData) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Item and assigns it to the Items field.
func (o *ContainerResponseData) SetItems(v []Item) {
	o.Items = v
}

// GetTrackingDetails returns the TrackingDetails field value if set, zero value otherwise.
func (o *ContainerResponseData) GetTrackingDetails() []TrackingDetail {
	if o == nil || o.TrackingDetails == nil {
		var ret []TrackingDetail
		return ret
	}
	return o.TrackingDetails
}

// GetTrackingDetailsOk returns a tuple with the TrackingDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerResponseData) GetTrackingDetailsOk() ([]TrackingDetail, bool) {
	if o == nil || o.TrackingDetails == nil {
		return nil, false
	}
	return o.TrackingDetails, true
}

// HasTrackingDetails returns a boolean if a field has been set.
func (o *ContainerResponseData) HasTrackingDetails() bool {
	if o != nil && o.TrackingDetails != nil {
		return true
	}

	return false
}

// SetTrackingDetails gets a reference to the given []TrackingDetail and assigns it to the TrackingDetails field.
func (o *ContainerResponseData) SetTrackingDetails(v []TrackingDetail) {
	o.TrackingDetails = v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *ContainerResponseData) GetParentId() string {
	if o == nil || o.ParentId == nil {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerResponseData) GetParentIdOk() (*string, bool) {
	if o == nil || o.ParentId == nil {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *ContainerResponseData) HasParentId() bool {
	if o != nil && o.ParentId != nil {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *ContainerResponseData) SetParentId(v string) {
	o.ParentId = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ContainerResponseData) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerResponseData) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ContainerResponseData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *ContainerResponseData) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetIsHazmat returns the IsHazmat field value if set, zero value otherwise.
func (o *ContainerResponseData) GetIsHazmat() bool {
	if o == nil || o.IsHazmat == nil {
		var ret bool
		return ret
	}
	return *o.IsHazmat
}

// GetIsHazmatOk returns a tuple with the IsHazmat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerResponseData) GetIsHazmatOk() (*bool, bool) {
	if o == nil || o.IsHazmat == nil {
		return nil, false
	}
	return o.IsHazmat, true
}

// HasIsHazmat returns a boolean if a field has been set.
func (o *ContainerResponseData) HasIsHazmat() bool {
	if o != nil && o.IsHazmat != nil {
		return true
	}

	return false
}

// SetIsHazmat gets a reference to the given bool and assigns it to the IsHazmat field.
func (o *ContainerResponseData) SetIsHazmat(v bool) {
	o.IsHazmat = &v
}

// GetIsContainerizable returns the IsContainerizable field value if set, zero value otherwise.
func (o *ContainerResponseData) GetIsContainerizable() bool {
	if o == nil || o.IsContainerizable == nil {
		var ret bool
		return ret
	}
	return *o.IsContainerizable
}

// GetIsContainerizableOk returns a tuple with the IsContainerizable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerResponseData) GetIsContainerizableOk() (*bool, bool) {
	if o == nil || o.IsContainerizable == nil {
		return nil, false
	}
	return o.IsContainerizable, true
}

// HasIsContainerizable returns a boolean if a field has been set.
func (o *ContainerResponseData) HasIsContainerizable() bool {
	if o != nil && o.IsContainerizable != nil {
		return true
	}

	return false
}

// SetIsContainerizable gets a reference to the given bool and assigns it to the IsContainerizable field.
func (o *ContainerResponseData) SetIsContainerizable(v bool) {
	o.IsContainerizable = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ContainerResponseData) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ContainerResponseData) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ContainerResponseData) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *ContainerResponseData) GetCreatedBy() ActionBy {
	if o == nil {
		var ret ActionBy
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *ContainerResponseData) GetCreatedByOk() (*ActionBy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *ContainerResponseData) SetCreatedBy(v ActionBy) {
	o.CreatedBy = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ContainerResponseData) GetUpdatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ContainerResponseData) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ContainerResponseData) SetUpdatedAt(v int64) {
	o.UpdatedAt = v
}

// GetUpdatedBy returns the UpdatedBy field value
func (o *ContainerResponseData) GetUpdatedBy() ActionBy {
	if o == nil {
		var ret ActionBy
		return ret
	}

	return o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value
// and a boolean to check if the value has been set.
func (o *ContainerResponseData) GetUpdatedByOk() (*ActionBy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedBy, true
}

// SetUpdatedBy sets field value
func (o *ContainerResponseData) SetUpdatedBy(v ActionBy) {
	o.UpdatedBy = v
}

func (o ContainerResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["containerType"] = o.ContainerType
	}
	if o.IsReusable != nil {
		toSerialize["isReusable"] = o.IsReusable
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.TrackingDetails != nil {
		toSerialize["trackingDetails"] = o.TrackingDetails
	}
	if o.ParentId != nil {
		toSerialize["parentId"] = o.ParentId
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.IsHazmat != nil {
		toSerialize["isHazmat"] = o.IsHazmat
	}
	if o.IsContainerizable != nil {
		toSerialize["isContainerizable"] = o.IsContainerizable
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

type NullableContainerResponseData struct {
	value *ContainerResponseData
	isSet bool
}

func (v NullableContainerResponseData) Get() *ContainerResponseData {
	return v.value
}

func (v *NullableContainerResponseData) Set(val *ContainerResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerResponseData(val *ContainerResponseData) *NullableContainerResponseData {
	return &NullableContainerResponseData{value: val, isSet: true}
}

func (v NullableContainerResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


