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

// ChildContainerResponseData struct for ChildContainerResponseData
type ChildContainerResponseData struct {
	// container id
	Id *string `json:"id,omitempty"`
	Operation *Child `json:"operation,omitempty"`
}

// NewChildContainerResponseData instantiates a new ChildContainerResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChildContainerResponseData() *ChildContainerResponseData {
	this := ChildContainerResponseData{}
	return &this
}

// NewChildContainerResponseDataWithDefaults instantiates a new ChildContainerResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChildContainerResponseDataWithDefaults() *ChildContainerResponseData {
	this := ChildContainerResponseData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ChildContainerResponseData) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChildContainerResponseData) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ChildContainerResponseData) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ChildContainerResponseData) SetId(v string) {
	o.Id = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *ChildContainerResponseData) GetOperation() Child {
	if o == nil || isNil(o.Operation) {
		var ret Child
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChildContainerResponseData) GetOperationOk() (*Child, bool) {
	if o == nil || isNil(o.Operation) {
    return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *ChildContainerResponseData) HasOperation() bool {
	if o != nil && !isNil(o.Operation) {
		return true
	}

	return false
}

// SetOperation gets a reference to the given Child and assigns it to the Operation field.
func (o *ChildContainerResponseData) SetOperation(v Child) {
	o.Operation = &v
}

func (o ChildContainerResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Operation) {
		toSerialize["operation"] = o.Operation
	}
	return json.Marshal(toSerialize)
}

type NullableChildContainerResponseData struct {
	value *ChildContainerResponseData
	isSet bool
}

func (v NullableChildContainerResponseData) Get() *ChildContainerResponseData {
	return v.value
}

func (v *NullableChildContainerResponseData) Set(val *ChildContainerResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableChildContainerResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableChildContainerResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChildContainerResponseData(val *ChildContainerResponseData) *NullableChildContainerResponseData {
	return &NullableChildContainerResponseData{value: val, isSet: true}
}

func (v NullableChildContainerResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChildContainerResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


