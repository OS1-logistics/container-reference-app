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

// ContainerCreateAttributes Attributes that define a container and can be modified after container creation.
type ContainerCreateAttributes struct {
	// Field to represent tracking ID (such as AWB) for a container which is usually pasted on the container and used for scanning it. There can be multiple tracking IDs for a container, owned by different operators and hence it will be stored as a list.
	TrackingDetails []ContainerCreateAttributesTrackingDetailsInner `json:"trackingDetails,omitempty"`
	// A map to provide values for the attributes defined in Container-type configuration APIs.
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// List of itmes can be added into container only if container-type is a leaf.
	Items []Item `json:"items,omitempty"`
	// Represents whether this container can contain hazardous materials or not.
	IsHazmat *bool `json:"isHazmat,omitempty"`
	// Defines whether container can be put into other containers or not
	IsContainerizable *bool `json:"isContainerizable,omitempty"`
	// Defines whether container is re-usable or not
	IsReusable *bool `json:"isReusable,omitempty"`
}

// NewContainerCreateAttributes instantiates a new ContainerCreateAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerCreateAttributes() *ContainerCreateAttributes {
	this := ContainerCreateAttributes{}
	var isHazmat bool = false
	this.IsHazmat = &isHazmat
	var isContainerizable bool = true
	this.IsContainerizable = &isContainerizable
	var isReusable bool = false
	this.IsReusable = &isReusable
	return &this
}

// NewContainerCreateAttributesWithDefaults instantiates a new ContainerCreateAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerCreateAttributesWithDefaults() *ContainerCreateAttributes {
	this := ContainerCreateAttributes{}
	var isHazmat bool = false
	this.IsHazmat = &isHazmat
	var isContainerizable bool = true
	this.IsContainerizable = &isContainerizable
	var isReusable bool = false
	this.IsReusable = &isReusable
	return &this
}

// GetTrackingDetails returns the TrackingDetails field value if set, zero value otherwise.
func (o *ContainerCreateAttributes) GetTrackingDetails() []ContainerCreateAttributesTrackingDetailsInner {
	if o == nil || o.TrackingDetails == nil {
		var ret []ContainerCreateAttributesTrackingDetailsInner
		return ret
	}
	return o.TrackingDetails
}

// GetTrackingDetailsOk returns a tuple with the TrackingDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerCreateAttributes) GetTrackingDetailsOk() ([]ContainerCreateAttributesTrackingDetailsInner, bool) {
	if o == nil || o.TrackingDetails == nil {
		return nil, false
	}
	return o.TrackingDetails, true
}

// HasTrackingDetails returns a boolean if a field has been set.
func (o *ContainerCreateAttributes) HasTrackingDetails() bool {
	if o != nil && o.TrackingDetails != nil {
		return true
	}

	return false
}

// SetTrackingDetails gets a reference to the given []ContainerCreateAttributesTrackingDetailsInner and assigns it to the TrackingDetails field.
func (o *ContainerCreateAttributes) SetTrackingDetails(v []ContainerCreateAttributesTrackingDetailsInner) {
	o.TrackingDetails = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ContainerCreateAttributes) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerCreateAttributes) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ContainerCreateAttributes) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *ContainerCreateAttributes) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ContainerCreateAttributes) GetItems() []Item {
	if o == nil || o.Items == nil {
		var ret []Item
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerCreateAttributes) GetItemsOk() ([]Item, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ContainerCreateAttributes) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Item and assigns it to the Items field.
func (o *ContainerCreateAttributes) SetItems(v []Item) {
	o.Items = v
}

// GetIsHazmat returns the IsHazmat field value if set, zero value otherwise.
func (o *ContainerCreateAttributes) GetIsHazmat() bool {
	if o == nil || o.IsHazmat == nil {
		var ret bool
		return ret
	}
	return *o.IsHazmat
}

// GetIsHazmatOk returns a tuple with the IsHazmat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerCreateAttributes) GetIsHazmatOk() (*bool, bool) {
	if o == nil || o.IsHazmat == nil {
		return nil, false
	}
	return o.IsHazmat, true
}

// HasIsHazmat returns a boolean if a field has been set.
func (o *ContainerCreateAttributes) HasIsHazmat() bool {
	if o != nil && o.IsHazmat != nil {
		return true
	}

	return false
}

// SetIsHazmat gets a reference to the given bool and assigns it to the IsHazmat field.
func (o *ContainerCreateAttributes) SetIsHazmat(v bool) {
	o.IsHazmat = &v
}

// GetIsContainerizable returns the IsContainerizable field value if set, zero value otherwise.
func (o *ContainerCreateAttributes) GetIsContainerizable() bool {
	if o == nil || o.IsContainerizable == nil {
		var ret bool
		return ret
	}
	return *o.IsContainerizable
}

// GetIsContainerizableOk returns a tuple with the IsContainerizable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerCreateAttributes) GetIsContainerizableOk() (*bool, bool) {
	if o == nil || o.IsContainerizable == nil {
		return nil, false
	}
	return o.IsContainerizable, true
}

// HasIsContainerizable returns a boolean if a field has been set.
func (o *ContainerCreateAttributes) HasIsContainerizable() bool {
	if o != nil && o.IsContainerizable != nil {
		return true
	}

	return false
}

// SetIsContainerizable gets a reference to the given bool and assigns it to the IsContainerizable field.
func (o *ContainerCreateAttributes) SetIsContainerizable(v bool) {
	o.IsContainerizable = &v
}

// GetIsReusable returns the IsReusable field value if set, zero value otherwise.
func (o *ContainerCreateAttributes) GetIsReusable() bool {
	if o == nil || o.IsReusable == nil {
		var ret bool
		return ret
	}
	return *o.IsReusable
}

// GetIsReusableOk returns a tuple with the IsReusable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerCreateAttributes) GetIsReusableOk() (*bool, bool) {
	if o == nil || o.IsReusable == nil {
		return nil, false
	}
	return o.IsReusable, true
}

// HasIsReusable returns a boolean if a field has been set.
func (o *ContainerCreateAttributes) HasIsReusable() bool {
	if o != nil && o.IsReusable != nil {
		return true
	}

	return false
}

// SetIsReusable gets a reference to the given bool and assigns it to the IsReusable field.
func (o *ContainerCreateAttributes) SetIsReusable(v bool) {
	o.IsReusable = &v
}

func (o ContainerCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TrackingDetails != nil {
		toSerialize["trackingDetails"] = o.TrackingDetails
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.IsHazmat != nil {
		toSerialize["isHazmat"] = o.IsHazmat
	}
	if o.IsContainerizable != nil {
		toSerialize["isContainerizable"] = o.IsContainerizable
	}
	if o.IsReusable != nil {
		toSerialize["isReusable"] = o.IsReusable
	}
	return json.Marshal(toSerialize)
}

type NullableContainerCreateAttributes struct {
	value *ContainerCreateAttributes
	isSet bool
}

func (v NullableContainerCreateAttributes) Get() *ContainerCreateAttributes {
	return v.value
}

func (v *NullableContainerCreateAttributes) Set(val *ContainerCreateAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerCreateAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerCreateAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerCreateAttributes(val *ContainerCreateAttributes) *NullableContainerCreateAttributes {
	return &NullableContainerCreateAttributes{value: val, isSet: true}
}

func (v NullableContainerCreateAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerCreateAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


