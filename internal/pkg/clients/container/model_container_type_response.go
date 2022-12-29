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

// ContainerTypeResponse struct for ContainerTypeResponse
type ContainerTypeResponse struct {
	Data *ContainerTypeResponseData `json:"data,omitempty"`
	Request *Request `json:"request,omitempty"`
}

// NewContainerTypeResponse instantiates a new ContainerTypeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerTypeResponse() *ContainerTypeResponse {
	this := ContainerTypeResponse{}
	return &this
}

// NewContainerTypeResponseWithDefaults instantiates a new ContainerTypeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerTypeResponseWithDefaults() *ContainerTypeResponse {
	this := ContainerTypeResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ContainerTypeResponse) GetData() ContainerTypeResponseData {
	if o == nil || o.Data == nil {
		var ret ContainerTypeResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerTypeResponse) GetDataOk() (*ContainerTypeResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ContainerTypeResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given ContainerTypeResponseData and assigns it to the Data field.
func (o *ContainerTypeResponse) SetData(v ContainerTypeResponseData) {
	o.Data = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *ContainerTypeResponse) GetRequest() Request {
	if o == nil || o.Request == nil {
		var ret Request
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerTypeResponse) GetRequestOk() (*Request, bool) {
	if o == nil || o.Request == nil {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *ContainerTypeResponse) HasRequest() bool {
	if o != nil && o.Request != nil {
		return true
	}

	return false
}

// SetRequest gets a reference to the given Request and assigns it to the Request field.
func (o *ContainerTypeResponse) SetRequest(v Request) {
	o.Request = &v
}

func (o ContainerTypeResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Request != nil {
		toSerialize["request"] = o.Request
	}
	return json.Marshal(toSerialize)
}

type NullableContainerTypeResponse struct {
	value *ContainerTypeResponse
	isSet bool
}

func (v NullableContainerTypeResponse) Get() *ContainerTypeResponse {
	return v.value
}

func (v *NullableContainerTypeResponse) Set(val *ContainerTypeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerTypeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerTypeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerTypeResponse(val *ContainerTypeResponse) *NullableContainerTypeResponse {
	return &NullableContainerTypeResponse{value: val, isSet: true}
}

func (v NullableContainerTypeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerTypeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


