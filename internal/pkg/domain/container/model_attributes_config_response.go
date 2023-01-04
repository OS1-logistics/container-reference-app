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

// AttributesConfigResponse struct for AttributesConfigResponse
type AttributesConfigResponse struct {
	Data *AttributesConfigResponseData `json:"data,omitempty"`
	Request *Request `json:"request,omitempty"`
}

// NewAttributesConfigResponse instantiates a new AttributesConfigResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttributesConfigResponse() *AttributesConfigResponse {
	this := AttributesConfigResponse{}
	return &this
}

// NewAttributesConfigResponseWithDefaults instantiates a new AttributesConfigResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttributesConfigResponseWithDefaults() *AttributesConfigResponse {
	this := AttributesConfigResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AttributesConfigResponse) GetData() AttributesConfigResponseData {
	if o == nil || isNil(o.Data) {
		var ret AttributesConfigResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributesConfigResponse) GetDataOk() (*AttributesConfigResponseData, bool) {
	if o == nil || isNil(o.Data) {
    return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AttributesConfigResponse) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AttributesConfigResponseData and assigns it to the Data field.
func (o *AttributesConfigResponse) SetData(v AttributesConfigResponseData) {
	o.Data = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *AttributesConfigResponse) GetRequest() Request {
	if o == nil || isNil(o.Request) {
		var ret Request
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributesConfigResponse) GetRequestOk() (*Request, bool) {
	if o == nil || isNil(o.Request) {
    return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *AttributesConfigResponse) HasRequest() bool {
	if o != nil && !isNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given Request and assigns it to the Request field.
func (o *AttributesConfigResponse) SetRequest(v Request) {
	o.Request = &v
}

func (o AttributesConfigResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !isNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	return json.Marshal(toSerialize)
}

type NullableAttributesConfigResponse struct {
	value *AttributesConfigResponse
	isSet bool
}

func (v NullableAttributesConfigResponse) Get() *AttributesConfigResponse {
	return v.value
}

func (v *NullableAttributesConfigResponse) Set(val *AttributesConfigResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAttributesConfigResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAttributesConfigResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttributesConfigResponse(val *AttributesConfigResponse) *NullableAttributesConfigResponse {
	return &NullableAttributesConfigResponse{value: val, isSet: true}
}

func (v NullableAttributesConfigResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttributesConfigResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

