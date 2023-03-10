/*
Authentication And Authorization (AAA) Service

This swagger documentation provides all AAA API details. AAA service provides authentication and authorization capabilities for users.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aaadomain

import (
	"encoding/json"
)

// AcceptedResponse struct for AcceptedResponse
type AcceptedResponse struct {
	Error *Error `json:"error,omitempty"`
	Request *Request `json:"request,omitempty"`
	Data *SuccessResponseAllOfData `json:"data,omitempty"`
}

// NewAcceptedResponse instantiates a new AcceptedResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcceptedResponse() *AcceptedResponse {
	this := AcceptedResponse{}
	return &this
}

// NewAcceptedResponseWithDefaults instantiates a new AcceptedResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcceptedResponseWithDefaults() *AcceptedResponse {
	this := AcceptedResponse{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *AcceptedResponse) GetError() Error {
	if o == nil || isNil(o.Error) {
		var ret Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptedResponse) GetErrorOk() (*Error, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *AcceptedResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given Error and assigns it to the Error field.
func (o *AcceptedResponse) SetError(v Error) {
	o.Error = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *AcceptedResponse) GetRequest() Request {
	if o == nil || isNil(o.Request) {
		var ret Request
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptedResponse) GetRequestOk() (*Request, bool) {
	if o == nil || isNil(o.Request) {
    return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *AcceptedResponse) HasRequest() bool {
	if o != nil && !isNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given Request and assigns it to the Request field.
func (o *AcceptedResponse) SetRequest(v Request) {
	o.Request = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AcceptedResponse) GetData() SuccessResponseAllOfData {
	if o == nil || isNil(o.Data) {
		var ret SuccessResponseAllOfData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptedResponse) GetDataOk() (*SuccessResponseAllOfData, bool) {
	if o == nil || isNil(o.Data) {
    return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AcceptedResponse) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given SuccessResponseAllOfData and assigns it to the Data field.
func (o *AcceptedResponse) SetData(v SuccessResponseAllOfData) {
	o.Data = &v
}

func (o AcceptedResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !isNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableAcceptedResponse struct {
	value *AcceptedResponse
	isSet bool
}

func (v NullableAcceptedResponse) Get() *AcceptedResponse {
	return v.value
}

func (v *NullableAcceptedResponse) Set(val *AcceptedResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAcceptedResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAcceptedResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcceptedResponse(val *AcceptedResponse) *NullableAcceptedResponse {
	return &NullableAcceptedResponse{value: val, isSet: true}
}

func (v NullableAcceptedResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcceptedResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


