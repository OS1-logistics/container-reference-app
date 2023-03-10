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

// ContainerTrackingResponse struct for ContainerTrackingResponse
type ContainerTrackingResponse struct {
	Data *TrackingData `json:"data,omitempty"`
	Request *Request `json:"request,omitempty"`
}

// NewContainerTrackingResponse instantiates a new ContainerTrackingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerTrackingResponse() *ContainerTrackingResponse {
	this := ContainerTrackingResponse{}
	return &this
}

// NewContainerTrackingResponseWithDefaults instantiates a new ContainerTrackingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerTrackingResponseWithDefaults() *ContainerTrackingResponse {
	this := ContainerTrackingResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ContainerTrackingResponse) GetData() TrackingData {
	if o == nil || isNil(o.Data) {
		var ret TrackingData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerTrackingResponse) GetDataOk() (*TrackingData, bool) {
	if o == nil || isNil(o.Data) {
    return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ContainerTrackingResponse) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given TrackingData and assigns it to the Data field.
func (o *ContainerTrackingResponse) SetData(v TrackingData) {
	o.Data = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *ContainerTrackingResponse) GetRequest() Request {
	if o == nil || isNil(o.Request) {
		var ret Request
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerTrackingResponse) GetRequestOk() (*Request, bool) {
	if o == nil || isNil(o.Request) {
    return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *ContainerTrackingResponse) HasRequest() bool {
	if o != nil && !isNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given Request and assigns it to the Request field.
func (o *ContainerTrackingResponse) SetRequest(v Request) {
	o.Request = &v
}

func (o ContainerTrackingResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !isNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	return json.Marshal(toSerialize)
}

type NullableContainerTrackingResponse struct {
	value *ContainerTrackingResponse
	isSet bool
}

func (v NullableContainerTrackingResponse) Get() *ContainerTrackingResponse {
	return v.value
}

func (v *NullableContainerTrackingResponse) Set(val *ContainerTrackingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerTrackingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerTrackingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerTrackingResponse(val *ContainerTrackingResponse) *NullableContainerTrackingResponse {
	return &NullableContainerTrackingResponse{value: val, isSet: true}
}

func (v NullableContainerTrackingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerTrackingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


