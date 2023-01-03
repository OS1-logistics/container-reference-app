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

// AbacConfigRequest Setting for ABAC Configuration for a tenant.
type AbacConfigRequest struct {
	// url on which ABAC authorization requests will be sent.
	Url *string `json:"url,omitempty"`
	// Content-type header of the request. Currently application/json, application/x-www-form-urlencoded are supported.
	ContentType *string `json:"contentType,omitempty"`
	// Headers required for sending the validation request.
	Headers map[string]interface{} `json:"headers,omitempty"`
	// Whether the ABAC is enabled during authorization.
	IsActive *bool `json:"isActive,omitempty"`
}

// NewAbacConfigRequest instantiates a new AbacConfigRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAbacConfigRequest() *AbacConfigRequest {
	this := AbacConfigRequest{}
	var isActive bool = false
	this.IsActive = &isActive
	return &this
}

// NewAbacConfigRequestWithDefaults instantiates a new AbacConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAbacConfigRequestWithDefaults() *AbacConfigRequest {
	this := AbacConfigRequest{}
	var isActive bool = false
	this.IsActive = &isActive
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AbacConfigRequest) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AbacConfigRequest) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AbacConfigRequest) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AbacConfigRequest) SetUrl(v string) {
	o.Url = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *AbacConfigRequest) GetContentType() string {
	if o == nil || o.ContentType == nil {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AbacConfigRequest) GetContentTypeOk() (*string, bool) {
	if o == nil || o.ContentType == nil {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *AbacConfigRequest) HasContentType() bool {
	if o != nil && o.ContentType != nil {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *AbacConfigRequest) SetContentType(v string) {
	o.ContentType = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *AbacConfigRequest) GetHeaders() map[string]interface{} {
	if o == nil || o.Headers == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AbacConfigRequest) GetHeadersOk() (map[string]interface{}, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *AbacConfigRequest) HasHeaders() bool {
	if o != nil && o.Headers != nil {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string]interface{} and assigns it to the Headers field.
func (o *AbacConfigRequest) SetHeaders(v map[string]interface{}) {
	o.Headers = v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *AbacConfigRequest) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AbacConfigRequest) GetIsActiveOk() (*bool, bool) {
	if o == nil || o.IsActive == nil {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *AbacConfigRequest) HasIsActive() bool {
	if o != nil && o.IsActive != nil {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *AbacConfigRequest) SetIsActive(v bool) {
	o.IsActive = &v
}

func (o AbacConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.ContentType != nil {
		toSerialize["contentType"] = o.ContentType
	}
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	if o.IsActive != nil {
		toSerialize["isActive"] = o.IsActive
	}
	return json.Marshal(toSerialize)
}

type NullableAbacConfigRequest struct {
	value *AbacConfigRequest
	isSet bool
}

func (v NullableAbacConfigRequest) Get() *AbacConfigRequest {
	return v.value
}

func (v *NullableAbacConfigRequest) Set(val *AbacConfigRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAbacConfigRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAbacConfigRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAbacConfigRequest(val *AbacConfigRequest) *NullableAbacConfigRequest {
	return &NullableAbacConfigRequest{value: val, isSet: true}
}

func (v NullableAbacConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAbacConfigRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


