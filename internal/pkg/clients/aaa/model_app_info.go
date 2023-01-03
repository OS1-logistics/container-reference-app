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

// AppInfo struct for AppInfo
type AppInfo struct {
	// Description of the application.
	Description NullableString `json:"description,omitempty"`
	// Set if app is private and can be allocated only to the Tenant specified in the field `privateTenantId`.
	IsPrivateApp bool `json:"isPrivateApp"`
	// Will be ignored if `isPrivate` is false.
	PrivateTenantId *string `json:"privateTenantId,omitempty"`
	// When `isActive` = False OR `isDeleted` = False, the role will be ignored for granting permissions.
	IsActive bool `json:"isActive"`
	// Array of allowed redirect URIs.
	RedirectUri []string `json:"redirectUri,omitempty"`
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
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

// NewAppInfo instantiates a new AppInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppInfo(isPrivateApp bool, isActive bool, appId interface{}, appType string, appName string) *AppInfo {
	this := AppInfo{}
	this.IsPrivateApp = isPrivateApp
	this.IsActive = isActive
	this.AppId = appId
	this.AppType = appType
	this.AppName = appName
	var isDeleted bool = false
	this.IsDeleted = &isDeleted
	return &this
}

// NewAppInfoWithDefaults instantiates a new AppInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppInfoWithDefaults() *AppInfo {
	this := AppInfo{}
	var isPrivateApp bool = false
	this.IsPrivateApp = isPrivateApp
	var isDeleted bool = false
	this.IsDeleted = &isDeleted
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppInfo) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppInfo) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *AppInfo) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *AppInfo) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *AppInfo) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *AppInfo) UnsetDescription() {
	o.Description.Unset()
}

// GetIsPrivateApp returns the IsPrivateApp field value
func (o *AppInfo) GetIsPrivateApp() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPrivateApp
}

// GetIsPrivateAppOk returns a tuple with the IsPrivateApp field value
// and a boolean to check if the value has been set.
func (o *AppInfo) GetIsPrivateAppOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPrivateApp, true
}

// SetIsPrivateApp sets field value
func (o *AppInfo) SetIsPrivateApp(v bool) {
	o.IsPrivateApp = v
}

// GetPrivateTenantId returns the PrivateTenantId field value if set, zero value otherwise.
func (o *AppInfo) GetPrivateTenantId() string {
	if o == nil || o.PrivateTenantId == nil {
		var ret string
		return ret
	}
	return *o.PrivateTenantId
}

// GetPrivateTenantIdOk returns a tuple with the PrivateTenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppInfo) GetPrivateTenantIdOk() (*string, bool) {
	if o == nil || o.PrivateTenantId == nil {
		return nil, false
	}
	return o.PrivateTenantId, true
}

// HasPrivateTenantId returns a boolean if a field has been set.
func (o *AppInfo) HasPrivateTenantId() bool {
	if o != nil && o.PrivateTenantId != nil {
		return true
	}

	return false
}

// SetPrivateTenantId gets a reference to the given string and assigns it to the PrivateTenantId field.
func (o *AppInfo) SetPrivateTenantId(v string) {
	o.PrivateTenantId = &v
}

// GetIsActive returns the IsActive field value
func (o *AppInfo) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *AppInfo) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value
func (o *AppInfo) SetIsActive(v bool) {
	o.IsActive = v
}

// GetRedirectUri returns the RedirectUri field value if set, zero value otherwise.
func (o *AppInfo) GetRedirectUri() []string {
	if o == nil || o.RedirectUri == nil {
		var ret []string
		return ret
	}
	return o.RedirectUri
}

// GetRedirectUriOk returns a tuple with the RedirectUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppInfo) GetRedirectUriOk() ([]string, bool) {
	if o == nil || o.RedirectUri == nil {
		return nil, false
	}
	return o.RedirectUri, true
}

// HasRedirectUri returns a boolean if a field has been set.
func (o *AppInfo) HasRedirectUri() bool {
	if o != nil && o.RedirectUri != nil {
		return true
	}

	return false
}

// SetRedirectUri gets a reference to the given []string and assigns it to the RedirectUri field.
func (o *AppInfo) SetRedirectUri(v []string) {
	o.RedirectUri = v
}

// GetAppId returns the AppId field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *AppInfo) GetAppId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppInfo) GetAppIdOk() (*interface{}, bool) {
	if o == nil || o.AppId == nil {
		return nil, false
	}
	return &o.AppId, true
}

// SetAppId sets field value
func (o *AppInfo) SetAppId(v interface{}) {
	o.AppId = v
}

// GetAppType returns the AppType field value
func (o *AppInfo) GetAppType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppType
}

// GetAppTypeOk returns a tuple with the AppType field value
// and a boolean to check if the value has been set.
func (o *AppInfo) GetAppTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppType, true
}

// SetAppType sets field value
func (o *AppInfo) SetAppType(v string) {
	o.AppType = v
}

// GetAppName returns the AppName field value
func (o *AppInfo) GetAppName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value
// and a boolean to check if the value has been set.
func (o *AppInfo) GetAppNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppName, true
}

// SetAppName sets field value
func (o *AppInfo) SetAppName(v string) {
	o.AppName = v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *AppInfo) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppInfo) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *AppInfo) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *AppInfo) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecretId returns the SecretId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppInfo) GetSecretId() string {
	if o == nil || o.SecretId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SecretId.Get()
}

// GetSecretIdOk returns a tuple with the SecretId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppInfo) GetSecretIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecretId.Get(), o.SecretId.IsSet()
}

// HasSecretId returns a boolean if a field has been set.
func (o *AppInfo) HasSecretId() bool {
	if o != nil && o.SecretId.IsSet() {
		return true
	}

	return false
}

// SetSecretId gets a reference to the given NullableString and assigns it to the SecretId field.
func (o *AppInfo) SetSecretId(v string) {
	o.SecretId.Set(&v)
}
// SetSecretIdNil sets the value for SecretId to be an explicit nil
func (o *AppInfo) SetSecretIdNil() {
	o.SecretId.Set(nil)
}

// UnsetSecretId ensures that no value is present for SecretId, not even an explicit nil
func (o *AppInfo) UnsetSecretId() {
	o.SecretId.Unset()
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *AppInfo) GetIsDeleted() bool {
	if o == nil || o.IsDeleted == nil {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppInfo) GetIsDeletedOk() (*bool, bool) {
	if o == nil || o.IsDeleted == nil {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *AppInfo) HasIsDeleted() bool {
	if o != nil && o.IsDeleted != nil {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *AppInfo) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

func (o AppInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if true {
		toSerialize["isPrivateApp"] = o.IsPrivateApp
	}
	if o.PrivateTenantId != nil {
		toSerialize["privateTenantId"] = o.PrivateTenantId
	}
	if true {
		toSerialize["isActive"] = o.IsActive
	}
	if o.RedirectUri != nil {
		toSerialize["redirectUri"] = o.RedirectUri
	}
	if o.AppId != nil {
		toSerialize["appId"] = o.AppId
	}
	if true {
		toSerialize["appType"] = o.AppType
	}
	if true {
		toSerialize["appName"] = o.AppName
	}
	if o.ClientId != nil {
		toSerialize["clientId"] = o.ClientId
	}
	if o.SecretId.IsSet() {
		toSerialize["secretId"] = o.SecretId.Get()
	}
	if o.IsDeleted != nil {
		toSerialize["isDeleted"] = o.IsDeleted
	}
	return json.Marshal(toSerialize)
}

type NullableAppInfo struct {
	value *AppInfo
	isSet bool
}

func (v NullableAppInfo) Get() *AppInfo {
	return v.value
}

func (v *NullableAppInfo) Set(val *AppInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAppInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAppInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppInfo(val *AppInfo) *NullableAppInfo {
	return &NullableAppInfo{value: val, isSet: true}
}

func (v NullableAppInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


