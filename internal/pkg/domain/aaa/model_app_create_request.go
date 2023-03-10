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

// AppCreateRequest struct for AppCreateRequest
type AppCreateRequest struct {
	// Type of app.
	AppType string `json:"appType"`
	// Name of the app. Starts with alphabet. Only alphabets & hyphens allowed in the name.
	AppName string `json:"appName"`
	// Description of the app.
	Description *string `json:"description,omitempty"`
	// Set if app is private and can be allocated only to the Tenant specified in the field `privateTenantId`.
	IsPrivateApp *bool `json:"isPrivateApp,omitempty"`
	// Will be ignored if `isPrivate` is false.
	PrivateTenantId *string `json:"privateTenantId,omitempty"`
	// Array of redirect URIs allowed.
	RedirectUri []string `json:"redirectUri,omitempty"`
	// `clientId` of the app.
	ClientId *string `json:"clientId,omitempty"`
	ClientSecret *AppCreateRequestClientSecret `json:"clientSecret,omitempty"`
}

// NewAppCreateRequest instantiates a new AppCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppCreateRequest(appType string, appName string) *AppCreateRequest {
	this := AppCreateRequest{}
	this.AppType = appType
	this.AppName = appName
	return &this
}

// NewAppCreateRequestWithDefaults instantiates a new AppCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppCreateRequestWithDefaults() *AppCreateRequest {
	this := AppCreateRequest{}
	return &this
}

// GetAppType returns the AppType field value
func (o *AppCreateRequest) GetAppType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppType
}

// GetAppTypeOk returns a tuple with the AppType field value
// and a boolean to check if the value has been set.
func (o *AppCreateRequest) GetAppTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AppType, true
}

// SetAppType sets field value
func (o *AppCreateRequest) SetAppType(v string) {
	o.AppType = v
}

// GetAppName returns the AppName field value
func (o *AppCreateRequest) GetAppName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value
// and a boolean to check if the value has been set.
func (o *AppCreateRequest) GetAppNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AppName, true
}

// SetAppName sets field value
func (o *AppCreateRequest) SetAppName(v string) {
	o.AppName = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AppCreateRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCreateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AppCreateRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AppCreateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetIsPrivateApp returns the IsPrivateApp field value if set, zero value otherwise.
func (o *AppCreateRequest) GetIsPrivateApp() bool {
	if o == nil || isNil(o.IsPrivateApp) {
		var ret bool
		return ret
	}
	return *o.IsPrivateApp
}

// GetIsPrivateAppOk returns a tuple with the IsPrivateApp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCreateRequest) GetIsPrivateAppOk() (*bool, bool) {
	if o == nil || isNil(o.IsPrivateApp) {
    return nil, false
	}
	return o.IsPrivateApp, true
}

// HasIsPrivateApp returns a boolean if a field has been set.
func (o *AppCreateRequest) HasIsPrivateApp() bool {
	if o != nil && !isNil(o.IsPrivateApp) {
		return true
	}

	return false
}

// SetIsPrivateApp gets a reference to the given bool and assigns it to the IsPrivateApp field.
func (o *AppCreateRequest) SetIsPrivateApp(v bool) {
	o.IsPrivateApp = &v
}

// GetPrivateTenantId returns the PrivateTenantId field value if set, zero value otherwise.
func (o *AppCreateRequest) GetPrivateTenantId() string {
	if o == nil || isNil(o.PrivateTenantId) {
		var ret string
		return ret
	}
	return *o.PrivateTenantId
}

// GetPrivateTenantIdOk returns a tuple with the PrivateTenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCreateRequest) GetPrivateTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.PrivateTenantId) {
    return nil, false
	}
	return o.PrivateTenantId, true
}

// HasPrivateTenantId returns a boolean if a field has been set.
func (o *AppCreateRequest) HasPrivateTenantId() bool {
	if o != nil && !isNil(o.PrivateTenantId) {
		return true
	}

	return false
}

// SetPrivateTenantId gets a reference to the given string and assigns it to the PrivateTenantId field.
func (o *AppCreateRequest) SetPrivateTenantId(v string) {
	o.PrivateTenantId = &v
}

// GetRedirectUri returns the RedirectUri field value if set, zero value otherwise.
func (o *AppCreateRequest) GetRedirectUri() []string {
	if o == nil || isNil(o.RedirectUri) {
		var ret []string
		return ret
	}
	return o.RedirectUri
}

// GetRedirectUriOk returns a tuple with the RedirectUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCreateRequest) GetRedirectUriOk() ([]string, bool) {
	if o == nil || isNil(o.RedirectUri) {
    return nil, false
	}
	return o.RedirectUri, true
}

// HasRedirectUri returns a boolean if a field has been set.
func (o *AppCreateRequest) HasRedirectUri() bool {
	if o != nil && !isNil(o.RedirectUri) {
		return true
	}

	return false
}

// SetRedirectUri gets a reference to the given []string and assigns it to the RedirectUri field.
func (o *AppCreateRequest) SetRedirectUri(v []string) {
	o.RedirectUri = v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *AppCreateRequest) GetClientId() string {
	if o == nil || isNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCreateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || isNil(o.ClientId) {
    return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *AppCreateRequest) HasClientId() bool {
	if o != nil && !isNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *AppCreateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *AppCreateRequest) GetClientSecret() AppCreateRequestClientSecret {
	if o == nil || isNil(o.ClientSecret) {
		var ret AppCreateRequestClientSecret
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCreateRequest) GetClientSecretOk() (*AppCreateRequestClientSecret, bool) {
	if o == nil || isNil(o.ClientSecret) {
    return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *AppCreateRequest) HasClientSecret() bool {
	if o != nil && !isNil(o.ClientSecret) {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given AppCreateRequestClientSecret and assigns it to the ClientSecret field.
func (o *AppCreateRequest) SetClientSecret(v AppCreateRequestClientSecret) {
	o.ClientSecret = &v
}

func (o AppCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["appType"] = o.AppType
	}
	if true {
		toSerialize["appName"] = o.AppName
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.IsPrivateApp) {
		toSerialize["isPrivateApp"] = o.IsPrivateApp
	}
	if !isNil(o.PrivateTenantId) {
		toSerialize["privateTenantId"] = o.PrivateTenantId
	}
	if !isNil(o.RedirectUri) {
		toSerialize["redirectUri"] = o.RedirectUri
	}
	if !isNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !isNil(o.ClientSecret) {
		toSerialize["clientSecret"] = o.ClientSecret
	}
	return json.Marshal(toSerialize)
}

type NullableAppCreateRequest struct {
	value *AppCreateRequest
	isSet bool
}

func (v NullableAppCreateRequest) Get() *AppCreateRequest {
	return v.value
}

func (v *NullableAppCreateRequest) Set(val *AppCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCreateRequest(val *AppCreateRequest) *NullableAppCreateRequest {
	return &NullableAppCreateRequest{value: val, isSet: true}
}

func (v NullableAppCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


