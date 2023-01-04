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

// StateMachineConfigResponse struct for StateMachineConfigResponse
type StateMachineConfigResponse struct {
	Data *StateMachineConfigResponseData `json:"data,omitempty"`
	Request *Request `json:"request,omitempty"`
}

// NewStateMachineConfigResponse instantiates a new StateMachineConfigResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStateMachineConfigResponse() *StateMachineConfigResponse {
	this := StateMachineConfigResponse{}
	return &this
}

// NewStateMachineConfigResponseWithDefaults instantiates a new StateMachineConfigResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStateMachineConfigResponseWithDefaults() *StateMachineConfigResponse {
	this := StateMachineConfigResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *StateMachineConfigResponse) GetData() StateMachineConfigResponseData {
	if o == nil || isNil(o.Data) {
		var ret StateMachineConfigResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StateMachineConfigResponse) GetDataOk() (*StateMachineConfigResponseData, bool) {
	if o == nil || isNil(o.Data) {
    return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *StateMachineConfigResponse) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given StateMachineConfigResponseData and assigns it to the Data field.
func (o *StateMachineConfigResponse) SetData(v StateMachineConfigResponseData) {
	o.Data = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *StateMachineConfigResponse) GetRequest() Request {
	if o == nil || isNil(o.Request) {
		var ret Request
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StateMachineConfigResponse) GetRequestOk() (*Request, bool) {
	if o == nil || isNil(o.Request) {
    return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *StateMachineConfigResponse) HasRequest() bool {
	if o != nil && !isNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given Request and assigns it to the Request field.
func (o *StateMachineConfigResponse) SetRequest(v Request) {
	o.Request = &v
}

func (o StateMachineConfigResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !isNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	return json.Marshal(toSerialize)
}

type NullableStateMachineConfigResponse struct {
	value *StateMachineConfigResponse
	isSet bool
}

func (v NullableStateMachineConfigResponse) Get() *StateMachineConfigResponse {
	return v.value
}

func (v *NullableStateMachineConfigResponse) Set(val *StateMachineConfigResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStateMachineConfigResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStateMachineConfigResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStateMachineConfigResponse(val *StateMachineConfigResponse) *NullableStateMachineConfigResponse {
	return &NullableStateMachineConfigResponse{value: val, isSet: true}
}

func (v NullableStateMachineConfigResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStateMachineConfigResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

