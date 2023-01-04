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

// ReasonsSuccessResponse struct for ReasonsSuccessResponse
type ReasonsSuccessResponse struct {
	Data *ReasonsSuccessResponseAllOfData `json:"data,omitempty"`
	Request *Request `json:"request,omitempty"`
}

// NewReasonsSuccessResponse instantiates a new ReasonsSuccessResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReasonsSuccessResponse() *ReasonsSuccessResponse {
	this := ReasonsSuccessResponse{}
	return &this
}

// NewReasonsSuccessResponseWithDefaults instantiates a new ReasonsSuccessResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReasonsSuccessResponseWithDefaults() *ReasonsSuccessResponse {
	this := ReasonsSuccessResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ReasonsSuccessResponse) GetData() ReasonsSuccessResponseAllOfData {
	if o == nil || isNil(o.Data) {
		var ret ReasonsSuccessResponseAllOfData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReasonsSuccessResponse) GetDataOk() (*ReasonsSuccessResponseAllOfData, bool) {
	if o == nil || isNil(o.Data) {
    return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ReasonsSuccessResponse) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ReasonsSuccessResponseAllOfData and assigns it to the Data field.
func (o *ReasonsSuccessResponse) SetData(v ReasonsSuccessResponseAllOfData) {
	o.Data = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *ReasonsSuccessResponse) GetRequest() Request {
	if o == nil || isNil(o.Request) {
		var ret Request
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReasonsSuccessResponse) GetRequestOk() (*Request, bool) {
	if o == nil || isNil(o.Request) {
    return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *ReasonsSuccessResponse) HasRequest() bool {
	if o != nil && !isNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given Request and assigns it to the Request field.
func (o *ReasonsSuccessResponse) SetRequest(v Request) {
	o.Request = &v
}

func (o ReasonsSuccessResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !isNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	return json.Marshal(toSerialize)
}

type NullableReasonsSuccessResponse struct {
	value *ReasonsSuccessResponse
	isSet bool
}

func (v NullableReasonsSuccessResponse) Get() *ReasonsSuccessResponse {
	return v.value
}

func (v *NullableReasonsSuccessResponse) Set(val *ReasonsSuccessResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReasonsSuccessResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReasonsSuccessResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReasonsSuccessResponse(val *ReasonsSuccessResponse) *NullableReasonsSuccessResponse {
	return &NullableReasonsSuccessResponse{value: val, isSet: true}
}

func (v NullableReasonsSuccessResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReasonsSuccessResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


