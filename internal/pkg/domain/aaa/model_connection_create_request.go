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

// ConnectionCreateRequest struct for ConnectionCreateRequest
type ConnectionCreateRequest struct {
	// Unique name of the connection.
	ConnectionName string `json:"connectionName"`
	// Client ID of the social application.
	ConnectionClientId string `json:"connectionClientId"`
	// Client Secret of the social application.
	ConnectionClientSecret string `json:"connectionClientSecret"`
}

// NewConnectionCreateRequest instantiates a new ConnectionCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionCreateRequest(connectionName string, connectionClientId string, connectionClientSecret string) *ConnectionCreateRequest {
	this := ConnectionCreateRequest{}
	this.ConnectionName = connectionName
	this.ConnectionClientId = connectionClientId
	this.ConnectionClientSecret = connectionClientSecret
	return &this
}

// NewConnectionCreateRequestWithDefaults instantiates a new ConnectionCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionCreateRequestWithDefaults() *ConnectionCreateRequest {
	this := ConnectionCreateRequest{}
	return &this
}

// GetConnectionName returns the ConnectionName field value
func (o *ConnectionCreateRequest) GetConnectionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionName
}

// GetConnectionNameOk returns a tuple with the ConnectionName field value
// and a boolean to check if the value has been set.
func (o *ConnectionCreateRequest) GetConnectionNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ConnectionName, true
}

// SetConnectionName sets field value
func (o *ConnectionCreateRequest) SetConnectionName(v string) {
	o.ConnectionName = v
}

// GetConnectionClientId returns the ConnectionClientId field value
func (o *ConnectionCreateRequest) GetConnectionClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionClientId
}

// GetConnectionClientIdOk returns a tuple with the ConnectionClientId field value
// and a boolean to check if the value has been set.
func (o *ConnectionCreateRequest) GetConnectionClientIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ConnectionClientId, true
}

// SetConnectionClientId sets field value
func (o *ConnectionCreateRequest) SetConnectionClientId(v string) {
	o.ConnectionClientId = v
}

// GetConnectionClientSecret returns the ConnectionClientSecret field value
func (o *ConnectionCreateRequest) GetConnectionClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionClientSecret
}

// GetConnectionClientSecretOk returns a tuple with the ConnectionClientSecret field value
// and a boolean to check if the value has been set.
func (o *ConnectionCreateRequest) GetConnectionClientSecretOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ConnectionClientSecret, true
}

// SetConnectionClientSecret sets field value
func (o *ConnectionCreateRequest) SetConnectionClientSecret(v string) {
	o.ConnectionClientSecret = v
}

func (o ConnectionCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["connectionName"] = o.ConnectionName
	}
	if true {
		toSerialize["connectionClientId"] = o.ConnectionClientId
	}
	if true {
		toSerialize["connectionClientSecret"] = o.ConnectionClientSecret
	}
	return json.Marshal(toSerialize)
}

type NullableConnectionCreateRequest struct {
	value *ConnectionCreateRequest
	isSet bool
}

func (v NullableConnectionCreateRequest) Get() *ConnectionCreateRequest {
	return v.value
}

func (v *NullableConnectionCreateRequest) Set(val *ConnectionCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionCreateRequest(val *ConnectionCreateRequest) *NullableConnectionCreateRequest {
	return &NullableConnectionCreateRequest{value: val, isSet: true}
}

func (v NullableConnectionCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


