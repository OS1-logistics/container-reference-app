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

// ContainerUpdateRequest struct for ContainerUpdateRequest
type ContainerUpdateRequest struct {
	Callback *map[string]interface{} `json:"callback,omitempty"`
	// Field to represent tracking ID (such as AWB) for a container which is usually pasted on the container and used for scanning it. There can be multiple tracking IDs for a container, owned by different operators and hence it will be stored as a list.
	TrackingDetails []ContainerUpdateAttributesTrackingDetailsInner `json:"trackingDetails,omitempty"`
	// A map to provide values for the attributes defined in Container-type configuration APIs.
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// List of itmes can be added into container only if container-type is a leaf.
	Items []Item `json:"items,omitempty"`
}

// NewContainerUpdateRequest instantiates a new ContainerUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerUpdateRequest() *ContainerUpdateRequest {
	this := ContainerUpdateRequest{}
	return &this
}

// NewContainerUpdateRequestWithDefaults instantiates a new ContainerUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerUpdateRequestWithDefaults() *ContainerUpdateRequest {
	this := ContainerUpdateRequest{}
	return &this
}

// GetCallback returns the Callback field value if set, zero value otherwise.
func (o *ContainerUpdateRequest) GetCallback() map[string]interface{} {
	if o == nil || isNil(o.Callback) {
		var ret map[string]interface{}
		return ret
	}
	return *o.Callback
}

// GetCallbackOk returns a tuple with the Callback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerUpdateRequest) GetCallbackOk() (*map[string]interface{}, bool) {
	if o == nil || isNil(o.Callback) {
    return nil, false
	}
	return o.Callback, true
}

// HasCallback returns a boolean if a field has been set.
func (o *ContainerUpdateRequest) HasCallback() bool {
	if o != nil && !isNil(o.Callback) {
		return true
	}

	return false
}

// SetCallback gets a reference to the given map[string]interface{} and assigns it to the Callback field.
func (o *ContainerUpdateRequest) SetCallback(v map[string]interface{}) {
	o.Callback = &v
}

// GetTrackingDetails returns the TrackingDetails field value if set, zero value otherwise.
func (o *ContainerUpdateRequest) GetTrackingDetails() []ContainerUpdateAttributesTrackingDetailsInner {
	if o == nil || isNil(o.TrackingDetails) {
		var ret []ContainerUpdateAttributesTrackingDetailsInner
		return ret
	}
	return o.TrackingDetails
}

// GetTrackingDetailsOk returns a tuple with the TrackingDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerUpdateRequest) GetTrackingDetailsOk() ([]ContainerUpdateAttributesTrackingDetailsInner, bool) {
	if o == nil || isNil(o.TrackingDetails) {
    return nil, false
	}
	return o.TrackingDetails, true
}

// HasTrackingDetails returns a boolean if a field has been set.
func (o *ContainerUpdateRequest) HasTrackingDetails() bool {
	if o != nil && !isNil(o.TrackingDetails) {
		return true
	}

	return false
}

// SetTrackingDetails gets a reference to the given []ContainerUpdateAttributesTrackingDetailsInner and assigns it to the TrackingDetails field.
func (o *ContainerUpdateRequest) SetTrackingDetails(v []ContainerUpdateAttributesTrackingDetailsInner) {
	o.TrackingDetails = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ContainerUpdateRequest) GetAttributes() map[string]interface{} {
	if o == nil || isNil(o.Attributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerUpdateRequest) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Attributes) {
    return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ContainerUpdateRequest) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *ContainerUpdateRequest) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ContainerUpdateRequest) GetItems() []Item {
	if o == nil || isNil(o.Items) {
		var ret []Item
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerUpdateRequest) GetItemsOk() ([]Item, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ContainerUpdateRequest) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Item and assigns it to the Items field.
func (o *ContainerUpdateRequest) SetItems(v []Item) {
	o.Items = v
}

func (o ContainerUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Callback) {
		toSerialize["callback"] = o.Callback
	}
	if !isNil(o.TrackingDetails) {
		toSerialize["trackingDetails"] = o.TrackingDetails
	}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableContainerUpdateRequest struct {
	value *ContainerUpdateRequest
	isSet bool
}

func (v NullableContainerUpdateRequest) Get() *ContainerUpdateRequest {
	return v.value
}

func (v *NullableContainerUpdateRequest) Set(val *ContainerUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerUpdateRequest(val *ContainerUpdateRequest) *NullableContainerUpdateRequest {
	return &NullableContainerUpdateRequest{value: val, isSet: true}
}

func (v NullableContainerUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


