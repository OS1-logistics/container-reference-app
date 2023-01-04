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

// AppAllOf struct for AppAllOf
type AppAllOf struct {
	// Unique ID of the App. Format of ID is platform:app:<appName>
	AppId interface{} `json:"appId"`
	// Type of app.
	AppType string `json:"appType"`
	// Name of the app. Starts with alphabet and may contain hypen.
	AppName string `json:"appName"`
	// `clientId` of the app.
	ClientId *string `json:"clientId,omitempty"`
	// SDS ID where client secret is stored. Secrets are applicable for backend and internal apps only.
	SecretId NullableString `json:"secretId,omitempty"`
}

// NewAppAllOf instantiates a new AppAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppAllOf(appId interface{}, appType string, appName string) *AppAllOf {
	this := AppAllOf{}
	this.AppId = appId
	this.AppType = appType
	this.AppName = appName
	return &this
}

// NewAppAllOfWithDefaults instantiates a new AppAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppAllOfWithDefaults() *AppAllOf {
	this := AppAllOf{}
	return &this
}

// GetAppId returns the AppId field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *AppAllOf) GetAppId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppAllOf) GetAppIdOk() (*interface{}, bool) {
	if o == nil || isNil(o.AppId) {
    return nil, false
	}
	return &o.AppId, true
}

// SetAppId sets field value
func (o *AppAllOf) SetAppId(v interface{}) {
	o.AppId = v
}

// GetAppType returns the AppType field value
func (o *AppAllOf) GetAppType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppType
}

// GetAppTypeOk returns a tuple with the AppType field value
// and a boolean to check if the value has been set.
func (o *AppAllOf) GetAppTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AppType, true
}

// SetAppType sets field value
func (o *AppAllOf) SetAppType(v string) {
	o.AppType = v
}

// GetAppName returns the AppName field value
func (o *AppAllOf) GetAppName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value
// and a boolean to check if the value has been set.
func (o *AppAllOf) GetAppNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AppName, true
}

// SetAppName sets field value
func (o *AppAllOf) SetAppName(v string) {
	o.AppName = v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *AppAllOf) GetClientId() string {
	if o == nil || isNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAllOf) GetClientIdOk() (*string, bool) {
	if o == nil || isNil(o.ClientId) {
    return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *AppAllOf) HasClientId() bool {
	if o != nil && !isNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *AppAllOf) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecretId returns the SecretId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppAllOf) GetSecretId() string {
	if o == nil || isNil(o.SecretId.Get()) {
		var ret string
		return ret
	}
	return *o.SecretId.Get()
}

// GetSecretIdOk returns a tuple with the SecretId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppAllOf) GetSecretIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.SecretId.Get(), o.SecretId.IsSet()
}

// HasSecretId returns a boolean if a field has been set.
func (o *AppAllOf) HasSecretId() bool {
	if o != nil && o.SecretId.IsSet() {
		return true
	}

	return false
}

// SetSecretId gets a reference to the given NullableString and assigns it to the SecretId field.
func (o *AppAllOf) SetSecretId(v string) {
	o.SecretId.Set(&v)
}
// SetSecretIdNil sets the value for SecretId to be an explicit nil
func (o *AppAllOf) SetSecretIdNil() {
	o.SecretId.Set(nil)
}

// UnsetSecretId ensures that no value is present for SecretId, not even an explicit nil
func (o *AppAllOf) UnsetSecretId() {
	o.SecretId.Unset()
}

func (o AppAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppId != nil {
		toSerialize["appId"] = o.AppId
	}
	if true {
		toSerialize["appType"] = o.AppType
	}
	if true {
		toSerialize["appName"] = o.AppName
	}
	if !isNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if o.SecretId.IsSet() {
		toSerialize["secretId"] = o.SecretId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAppAllOf struct {
	value *AppAllOf
	isSet bool
}

func (v NullableAppAllOf) Get() *AppAllOf {
	return v.value
}

func (v *NullableAppAllOf) Set(val *AppAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAppAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAppAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppAllOf(val *AppAllOf) *NullableAppAllOf {
	return &NullableAppAllOf{value: val, isSet: true}
}

func (v NullableAppAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


