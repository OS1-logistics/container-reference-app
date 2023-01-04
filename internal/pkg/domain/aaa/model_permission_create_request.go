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

// PermissionCreateRequest struct for PermissionCreateRequest
type PermissionCreateRequest struct {
	// Named action that describes the operation permitted to be performed.
	Action string `json:"action"`
	ResourceId string `json:"resourceId"`
	HttpMethod string `json:"httpMethod"`
	// Description of the permission.
	Description string `json:"description"`
}

// NewPermissionCreateRequest instantiates a new PermissionCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionCreateRequest(action string, resourceId string, httpMethod string, description string) *PermissionCreateRequest {
	this := PermissionCreateRequest{}
	this.Action = action
	this.ResourceId = resourceId
	this.HttpMethod = httpMethod
	this.Description = description
	return &this
}

// NewPermissionCreateRequestWithDefaults instantiates a new PermissionCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionCreateRequestWithDefaults() *PermissionCreateRequest {
	this := PermissionCreateRequest{}
	return &this
}

// GetAction returns the Action field value
func (o *PermissionCreateRequest) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *PermissionCreateRequest) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *PermissionCreateRequest) SetAction(v string) {
	o.Action = v
}

// GetResourceId returns the ResourceId field value
func (o *PermissionCreateRequest) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *PermissionCreateRequest) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *PermissionCreateRequest) SetResourceId(v string) {
	o.ResourceId = v
}

// GetHttpMethod returns the HttpMethod field value
func (o *PermissionCreateRequest) GetHttpMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HttpMethod
}

// GetHttpMethodOk returns a tuple with the HttpMethod field value
// and a boolean to check if the value has been set.
func (o *PermissionCreateRequest) GetHttpMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HttpMethod, true
}

// SetHttpMethod sets field value
func (o *PermissionCreateRequest) SetHttpMethod(v string) {
	o.HttpMethod = v
}

// GetDescription returns the Description field value
func (o *PermissionCreateRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *PermissionCreateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *PermissionCreateRequest) SetDescription(v string) {
	o.Description = v
}

func (o PermissionCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["action"] = o.Action
	}
	if true {
		toSerialize["resourceId"] = o.ResourceId
	}
	if true {
		toSerialize["httpMethod"] = o.HttpMethod
	}
	if true {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullablePermissionCreateRequest struct {
	value *PermissionCreateRequest
	isSet bool
}

func (v NullablePermissionCreateRequest) Get() *PermissionCreateRequest {
	return v.value
}

func (v *NullablePermissionCreateRequest) Set(val *PermissionCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionCreateRequest(val *PermissionCreateRequest) *NullablePermissionCreateRequest {
	return &NullablePermissionCreateRequest{value: val, isSet: true}
}

func (v NullablePermissionCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


