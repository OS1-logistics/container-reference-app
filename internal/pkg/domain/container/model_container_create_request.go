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

// ContainerCreateRequest struct for ContainerCreateRequest
type ContainerCreateRequest struct {
	Callback *map[string]interface{} `json:"callback,omitempty"`
	// Field to represent tracking ID (such as AWB) for a container which is usually pasted on the container and used for scanning it. There can be multiple tracking IDs for a container, owned by different operators and hence it will be stored as a list.
	TrackingDetails []ContainerCreateAttributesTrackingDetailsInner `json:"trackingDetails,omitempty"`
	// A map to provide values for the attributes defined in Container-type configuration APIs.
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// List of itmes can be added into container only if container-type is a leaf.
	Items []Item `json:"items,omitempty"`
	ScannableId string `json:"scannableId"`
	// Represents whether this container can contain hazardous materials or not.
	IsHazmat *bool `json:"isHazmat,omitempty"`
	// Defines whether container can be put into other containers or not
	IsContainerizable *bool `json:"isContainerizable,omitempty"`
	// Defines whether container is re-usable or not
	IsReusable *bool `json:"isReusable,omitempty"`
}

// NewContainerCreateRequest instantiates a new ContainerCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerCreateRequest(scannableId string) *ContainerCreateRequest {
	this := ContainerCreateRequest{}
	this.ScannableId = scannableId
	var isHazmat bool = false
	this.IsHazmat = &isHazmat
	var isContainerizable bool = true
	this.IsContainerizable = &isContainerizable
	var isReusable bool = false
	this.IsReusable = &isReusable
	return &this
}

// NewContainerCreateRequestWithDefaults instantiates a new ContainerCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerCreateRequestWithDefaults() *ContainerCreateRequest {
	this := ContainerCreateRequest{}
	var isHazmat bool = false
	this.IsHazmat = &isHazmat
	var isContainerizable bool = true
	this.IsContainerizable = &isContainerizable
	var isReusable bool = false
	this.IsReusable = &isReusable
	return &this
}

// GetCallback returns the Callback field value if set, zero value otherwise.
func (o *ContainerCreateRequest) GetCallback() map[string]interface{} {
	if o == nil || isNil(o.Callback) {
		var ret map[string]interface{}
		return ret
	}
	return *o.Callback
}

// GetCallbackOk returns a tuple with the Callback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerCreateRequest) GetCallbackOk() (*map[string]interface{}, bool) {
	if o == nil || isNil(o.Callback) {
    return nil, false
	}
	return o.Callback, true
}

// HasCallback returns a boolean if a field has been set.
func (o *ContainerCreateRequest) HasCallback() bool {
	if o != nil && !isNil(o.Callback) {
		return true
	}

	return false
}

// SetCallback gets a reference to the given map[string]interface{} and assigns it to the Callback field.
func (o *ContainerCreateRequest) SetCallback(v map[string]interface{}) {
	o.Callback = &v
}

// GetTrackingDetails returns the TrackingDetails field value if set, zero value otherwise.
func (o *ContainerCreateRequest) GetTrackingDetails() []ContainerCreateAttributesTrackingDetailsInner {
	if o == nil || isNil(o.TrackingDetails) {
		var ret []ContainerCreateAttributesTrackingDetailsInner
		return ret
	}
	return o.TrackingDetails
}

// GetTrackingDetailsOk returns a tuple with the TrackingDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerCreateRequest) GetTrackingDetailsOk() ([]ContainerCreateAttributesTrackingDetailsInner, bool) {
	if o == nil || isNil(o.TrackingDetails) {
    return nil, false
	}
	return o.TrackingDetails, true
}

// HasTrackingDetails returns a boolean if a field has been set.
func (o *ContainerCreateRequest) HasTrackingDetails() bool {
	if o != nil && !isNil(o.TrackingDetails) {
		return true
	}

	return false
}

// SetTrackingDetails gets a reference to the given []ContainerCreateAttributesTrackingDetailsInner and assigns it to the TrackingDetails field.
func (o *ContainerCreateRequest) SetTrackingDetails(v []ContainerCreateAttributesTrackingDetailsInner) {
	o.TrackingDetails = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ContainerCreateRequest) GetAttributes() map[string]interface{} {
	if o == nil || isNil(o.Attributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerCreateRequest) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Attributes) {
    return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ContainerCreateRequest) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *ContainerCreateRequest) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ContainerCreateRequest) GetItems() []Item {
	if o == nil || isNil(o.Items) {
		var ret []Item
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerCreateRequest) GetItemsOk() ([]Item, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ContainerCreateRequest) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Item and assigns it to the Items field.
func (o *ContainerCreateRequest) SetItems(v []Item) {
	o.Items = v
}

// GetScannableId returns the ScannableId field value
func (o *ContainerCreateRequest) GetScannableId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScannableId
}

// GetScannableIdOk returns a tuple with the ScannableId field value
// and a boolean to check if the value has been set.
func (o *ContainerCreateRequest) GetScannableIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ScannableId, true
}

// SetScannableId sets field value
func (o *ContainerCreateRequest) SetScannableId(v string) {
	o.ScannableId = v
}

// GetIsHazmat returns the IsHazmat field value if set, zero value otherwise.
func (o *ContainerCreateRequest) GetIsHazmat() bool {
	if o == nil || isNil(o.IsHazmat) {
		var ret bool
		return ret
	}
	return *o.IsHazmat
}

// GetIsHazmatOk returns a tuple with the IsHazmat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerCreateRequest) GetIsHazmatOk() (*bool, bool) {
	if o == nil || isNil(o.IsHazmat) {
    return nil, false
	}
	return o.IsHazmat, true
}

// HasIsHazmat returns a boolean if a field has been set.
func (o *ContainerCreateRequest) HasIsHazmat() bool {
	if o != nil && !isNil(o.IsHazmat) {
		return true
	}

	return false
}

// SetIsHazmat gets a reference to the given bool and assigns it to the IsHazmat field.
func (o *ContainerCreateRequest) SetIsHazmat(v bool) {
	o.IsHazmat = &v
}

// GetIsContainerizable returns the IsContainerizable field value if set, zero value otherwise.
func (o *ContainerCreateRequest) GetIsContainerizable() bool {
	if o == nil || isNil(o.IsContainerizable) {
		var ret bool
		return ret
	}
	return *o.IsContainerizable
}

// GetIsContainerizableOk returns a tuple with the IsContainerizable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerCreateRequest) GetIsContainerizableOk() (*bool, bool) {
	if o == nil || isNil(o.IsContainerizable) {
    return nil, false
	}
	return o.IsContainerizable, true
}

// HasIsContainerizable returns a boolean if a field has been set.
func (o *ContainerCreateRequest) HasIsContainerizable() bool {
	if o != nil && !isNil(o.IsContainerizable) {
		return true
	}

	return false
}

// SetIsContainerizable gets a reference to the given bool and assigns it to the IsContainerizable field.
func (o *ContainerCreateRequest) SetIsContainerizable(v bool) {
	o.IsContainerizable = &v
}

// GetIsReusable returns the IsReusable field value if set, zero value otherwise.
func (o *ContainerCreateRequest) GetIsReusable() bool {
	if o == nil || isNil(o.IsReusable) {
		var ret bool
		return ret
	}
	return *o.IsReusable
}

// GetIsReusableOk returns a tuple with the IsReusable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerCreateRequest) GetIsReusableOk() (*bool, bool) {
	if o == nil || isNil(o.IsReusable) {
    return nil, false
	}
	return o.IsReusable, true
}

// HasIsReusable returns a boolean if a field has been set.
func (o *ContainerCreateRequest) HasIsReusable() bool {
	if o != nil && !isNil(o.IsReusable) {
		return true
	}

	return false
}

// SetIsReusable gets a reference to the given bool and assigns it to the IsReusable field.
func (o *ContainerCreateRequest) SetIsReusable(v bool) {
	o.IsReusable = &v
}

func (o ContainerCreateRequest) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["scannableId"] = o.ScannableId
	}
	if !isNil(o.IsHazmat) {
		toSerialize["isHazmat"] = o.IsHazmat
	}
	if !isNil(o.IsContainerizable) {
		toSerialize["isContainerizable"] = o.IsContainerizable
	}
	if !isNil(o.IsReusable) {
		toSerialize["isReusable"] = o.IsReusable
	}
	return json.Marshal(toSerialize)
}

type NullableContainerCreateRequest struct {
	value *ContainerCreateRequest
	isSet bool
}

func (v NullableContainerCreateRequest) Get() *ContainerCreateRequest {
	return v.value
}

func (v *NullableContainerCreateRequest) Set(val *ContainerCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerCreateRequest(val *ContainerCreateRequest) *NullableContainerCreateRequest {
	return &NullableContainerCreateRequest{value: val, isSet: true}
}

func (v NullableContainerCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


