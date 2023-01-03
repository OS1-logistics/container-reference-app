/*
Authentication And Authorization (AAA) Service

This swagger documentation provides all AAA API details. AAA service provides authentication and authorization capabilities for users.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aaa_client

import (
	"encoding/json"
)

// AppIdResponse struct for AppIdResponse
type AppIdResponse struct {
	// Unique ID assigned to each role at the time of creation.
	AppId *string `json:"appId,omitempty"`
}

// NewAppIdResponse instantiates a new AppIdResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppIdResponse() *AppIdResponse {
	this := AppIdResponse{}
	return &this
}

// NewAppIdResponseWithDefaults instantiates a new AppIdResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppIdResponseWithDefaults() *AppIdResponse {
	this := AppIdResponse{}
	return &this
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *AppIdResponse) GetAppId() string {
	if o == nil || o.AppId == nil {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppIdResponse) GetAppIdOk() (*string, bool) {
	if o == nil || o.AppId == nil {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *AppIdResponse) HasAppId() bool {
	if o != nil && o.AppId != nil {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *AppIdResponse) SetAppId(v string) {
	o.AppId = &v
}

func (o AppIdResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppId != nil {
		toSerialize["appId"] = o.AppId
	}
	return json.Marshal(toSerialize)
}

type NullableAppIdResponse struct {
	value *AppIdResponse
	isSet bool
}

func (v NullableAppIdResponse) Get() *AppIdResponse {
	return v.value
}

func (v *NullableAppIdResponse) Set(val *AppIdResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppIdResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppIdResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppIdResponse(val *AppIdResponse) *NullableAppIdResponse {
	return &NullableAppIdResponse{value: val, isSet: true}
}

func (v NullableAppIdResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppIdResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


