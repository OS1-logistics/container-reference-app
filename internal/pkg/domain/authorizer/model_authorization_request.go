/*
AAA - External Authorizer

This swagger documentation provides all Authorization api details. External Authorizer service provides authentication and authorization capabilities for users and machines.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authorizerdomain

import (
	"encoding/json"
)

// AuthorizationRequest struct for AuthorizationRequest
type AuthorizationRequest struct {
	Path string `json:"path"`
	Method string `json:"method"`
}

// NewAuthorizationRequest instantiates a new AuthorizationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationRequest(path string, method string) *AuthorizationRequest {
	this := AuthorizationRequest{}
	this.Path = path
	this.Method = method
	return &this
}

// NewAuthorizationRequestWithDefaults instantiates a new AuthorizationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationRequestWithDefaults() *AuthorizationRequest {
	this := AuthorizationRequest{}
	return &this
}

// GetPath returns the Path field value
func (o *AuthorizationRequest) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *AuthorizationRequest) GetPathOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *AuthorizationRequest) SetPath(v string) {
	o.Path = v
}

// GetMethod returns the Method field value
func (o *AuthorizationRequest) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *AuthorizationRequest) GetMethodOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *AuthorizationRequest) SetMethod(v string) {
	o.Method = v
}

func (o AuthorizationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["path"] = o.Path
	}
	if true {
		toSerialize["method"] = o.Method
	}
	return json.Marshal(toSerialize)
}

type NullableAuthorizationRequest struct {
	value *AuthorizationRequest
	isSet bool
}

func (v NullableAuthorizationRequest) Get() *AuthorizationRequest {
	return v.value
}

func (v *NullableAuthorizationRequest) Set(val *AuthorizationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationRequest(val *AuthorizationRequest) *NullableAuthorizationRequest {
	return &NullableAuthorizationRequest{value: val, isSet: true}
}

func (v NullableAuthorizationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


