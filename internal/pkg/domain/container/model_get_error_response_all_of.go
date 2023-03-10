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

// GetErrorResponseAllOf struct for GetErrorResponseAllOf
type GetErrorResponseAllOf struct {
	Data *GetErrorResponseAllOfData `json:"data,omitempty"`
}

// NewGetErrorResponseAllOf instantiates a new GetErrorResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetErrorResponseAllOf() *GetErrorResponseAllOf {
	this := GetErrorResponseAllOf{}
	return &this
}

// NewGetErrorResponseAllOfWithDefaults instantiates a new GetErrorResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetErrorResponseAllOfWithDefaults() *GetErrorResponseAllOf {
	this := GetErrorResponseAllOf{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetErrorResponseAllOf) GetData() GetErrorResponseAllOfData {
	if o == nil || isNil(o.Data) {
		var ret GetErrorResponseAllOfData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetErrorResponseAllOf) GetDataOk() (*GetErrorResponseAllOfData, bool) {
	if o == nil || isNil(o.Data) {
    return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetErrorResponseAllOf) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GetErrorResponseAllOfData and assigns it to the Data field.
func (o *GetErrorResponseAllOf) SetData(v GetErrorResponseAllOfData) {
	o.Data = &v
}

func (o GetErrorResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetErrorResponseAllOf struct {
	value *GetErrorResponseAllOf
	isSet bool
}

func (v NullableGetErrorResponseAllOf) Get() *GetErrorResponseAllOf {
	return v.value
}

func (v *NullableGetErrorResponseAllOf) Set(val *GetErrorResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGetErrorResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGetErrorResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetErrorResponseAllOf(val *GetErrorResponseAllOf) *NullableGetErrorResponseAllOf {
	return &NullableGetErrorResponseAllOf{value: val, isSet: true}
}

func (v NullableGetErrorResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetErrorResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


