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

// ContainerBatchResponse struct for ContainerBatchResponse
type ContainerBatchResponse struct {
	Data *ContainerBatchResponseData `json:"data,omitempty"`
	Request *Request `json:"request,omitempty"`
}

// NewContainerBatchResponse instantiates a new ContainerBatchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerBatchResponse() *ContainerBatchResponse {
	this := ContainerBatchResponse{}
	return &this
}

// NewContainerBatchResponseWithDefaults instantiates a new ContainerBatchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerBatchResponseWithDefaults() *ContainerBatchResponse {
	this := ContainerBatchResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ContainerBatchResponse) GetData() ContainerBatchResponseData {
	if o == nil || isNil(o.Data) {
		var ret ContainerBatchResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerBatchResponse) GetDataOk() (*ContainerBatchResponseData, bool) {
	if o == nil || isNil(o.Data) {
    return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ContainerBatchResponse) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ContainerBatchResponseData and assigns it to the Data field.
func (o *ContainerBatchResponse) SetData(v ContainerBatchResponseData) {
	o.Data = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *ContainerBatchResponse) GetRequest() Request {
	if o == nil || isNil(o.Request) {
		var ret Request
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerBatchResponse) GetRequestOk() (*Request, bool) {
	if o == nil || isNil(o.Request) {
    return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *ContainerBatchResponse) HasRequest() bool {
	if o != nil && !isNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given Request and assigns it to the Request field.
func (o *ContainerBatchResponse) SetRequest(v Request) {
	o.Request = &v
}

func (o ContainerBatchResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !isNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	return json.Marshal(toSerialize)
}

type NullableContainerBatchResponse struct {
	value *ContainerBatchResponse
	isSet bool
}

func (v NullableContainerBatchResponse) Get() *ContainerBatchResponse {
	return v.value
}

func (v *NullableContainerBatchResponse) Set(val *ContainerBatchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerBatchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerBatchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerBatchResponse(val *ContainerBatchResponse) *NullableContainerBatchResponse {
	return &NullableContainerBatchResponse{value: val, isSet: true}
}

func (v NullableContainerBatchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerBatchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


