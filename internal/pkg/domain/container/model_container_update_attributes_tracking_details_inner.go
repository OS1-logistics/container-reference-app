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

// ContainerUpdateAttributesTrackingDetailsInner struct for ContainerUpdateAttributesTrackingDetailsInner
type ContainerUpdateAttributesTrackingDetailsInner struct {
	// Field to specify the owner of the tracking ID.
	Operator string `json:"operator"`
	TrackingId string `json:"trackingId"`
	// it defines whether it is a primary tracking id or not. If not defined we assume first tracking id as primary tracking id.
	IsPrimary *bool `json:"isPrimary,omitempty"`
}

// NewContainerUpdateAttributesTrackingDetailsInner instantiates a new ContainerUpdateAttributesTrackingDetailsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerUpdateAttributesTrackingDetailsInner(operator string, trackingId string) *ContainerUpdateAttributesTrackingDetailsInner {
	this := ContainerUpdateAttributesTrackingDetailsInner{}
	this.Operator = operator
	this.TrackingId = trackingId
	var isPrimary bool = false
	this.IsPrimary = &isPrimary
	return &this
}

// NewContainerUpdateAttributesTrackingDetailsInnerWithDefaults instantiates a new ContainerUpdateAttributesTrackingDetailsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerUpdateAttributesTrackingDetailsInnerWithDefaults() *ContainerUpdateAttributesTrackingDetailsInner {
	this := ContainerUpdateAttributesTrackingDetailsInner{}
	var isPrimary bool = false
	this.IsPrimary = &isPrimary
	return &this
}

// GetOperator returns the Operator field value
func (o *ContainerUpdateAttributesTrackingDetailsInner) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *ContainerUpdateAttributesTrackingDetailsInner) GetOperatorOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *ContainerUpdateAttributesTrackingDetailsInner) SetOperator(v string) {
	o.Operator = v
}

// GetTrackingId returns the TrackingId field value
func (o *ContainerUpdateAttributesTrackingDetailsInner) GetTrackingId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TrackingId
}

// GetTrackingIdOk returns a tuple with the TrackingId field value
// and a boolean to check if the value has been set.
func (o *ContainerUpdateAttributesTrackingDetailsInner) GetTrackingIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TrackingId, true
}

// SetTrackingId sets field value
func (o *ContainerUpdateAttributesTrackingDetailsInner) SetTrackingId(v string) {
	o.TrackingId = v
}

// GetIsPrimary returns the IsPrimary field value if set, zero value otherwise.
func (o *ContainerUpdateAttributesTrackingDetailsInner) GetIsPrimary() bool {
	if o == nil || isNil(o.IsPrimary) {
		var ret bool
		return ret
	}
	return *o.IsPrimary
}

// GetIsPrimaryOk returns a tuple with the IsPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerUpdateAttributesTrackingDetailsInner) GetIsPrimaryOk() (*bool, bool) {
	if o == nil || isNil(o.IsPrimary) {
    return nil, false
	}
	return o.IsPrimary, true
}

// HasIsPrimary returns a boolean if a field has been set.
func (o *ContainerUpdateAttributesTrackingDetailsInner) HasIsPrimary() bool {
	if o != nil && !isNil(o.IsPrimary) {
		return true
	}

	return false
}

// SetIsPrimary gets a reference to the given bool and assigns it to the IsPrimary field.
func (o *ContainerUpdateAttributesTrackingDetailsInner) SetIsPrimary(v bool) {
	o.IsPrimary = &v
}

func (o ContainerUpdateAttributesTrackingDetailsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["operator"] = o.Operator
	}
	if true {
		toSerialize["trackingId"] = o.TrackingId
	}
	if !isNil(o.IsPrimary) {
		toSerialize["isPrimary"] = o.IsPrimary
	}
	return json.Marshal(toSerialize)
}

type NullableContainerUpdateAttributesTrackingDetailsInner struct {
	value *ContainerUpdateAttributesTrackingDetailsInner
	isSet bool
}

func (v NullableContainerUpdateAttributesTrackingDetailsInner) Get() *ContainerUpdateAttributesTrackingDetailsInner {
	return v.value
}

func (v *NullableContainerUpdateAttributesTrackingDetailsInner) Set(val *ContainerUpdateAttributesTrackingDetailsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerUpdateAttributesTrackingDetailsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerUpdateAttributesTrackingDetailsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerUpdateAttributesTrackingDetailsInner(val *ContainerUpdateAttributesTrackingDetailsInner) *NullableContainerUpdateAttributesTrackingDetailsInner {
	return &NullableContainerUpdateAttributesTrackingDetailsInner{value: val, isSet: true}
}

func (v NullableContainerUpdateAttributesTrackingDetailsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerUpdateAttributesTrackingDetailsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

