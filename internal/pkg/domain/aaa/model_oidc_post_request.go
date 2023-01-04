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

// OIDCPostRequest struct for OIDCPostRequest
type OIDCPostRequest struct {
	// Alias of the OIDC provider.
	ConnectionName string `json:"connectionName"`
	// Display name of the OIDC provider.
	DisplayName *string `json:"displayName,omitempty"`
	// The provider is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Order of the provider in GUI.
	GuiOrder *float32 `json:"guiOrder,omitempty"`
	// User info URL.
	UserInfoUrl *string `json:"userInfoUrl,omitempty"`
	// Log out URL.
	LogoutUrl *string `json:"logoutUrl,omitempty"`
	// Authorization URL.
	AuthorizationUrl string `json:"authorizationUrl"`
	// Token URL.
	TokenUrl string `json:"tokenUrl"`
	// Client identifier registered with identity provider.
	ClientId string `json:"clientId"`
	// Client secret.
	ClientSecret string `json:"clientSecret"`
	// Issuer of response.
	Issuer *string `json:"issuer,omitempty"`
	// JWKS URL.
	JwksUrl *string `json:"jwksUrl,omitempty"`
}

// NewOIDCPostRequest instantiates a new OIDCPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOIDCPostRequest(connectionName string, authorizationUrl string, tokenUrl string, clientId string, clientSecret string) *OIDCPostRequest {
	this := OIDCPostRequest{}
	this.ConnectionName = connectionName
	this.AuthorizationUrl = authorizationUrl
	this.TokenUrl = tokenUrl
	this.ClientId = clientId
	this.ClientSecret = clientSecret
	return &this
}

// NewOIDCPostRequestWithDefaults instantiates a new OIDCPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOIDCPostRequestWithDefaults() *OIDCPostRequest {
	this := OIDCPostRequest{}
	return &this
}

// GetConnectionName returns the ConnectionName field value
func (o *OIDCPostRequest) GetConnectionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionName
}

// GetConnectionNameOk returns a tuple with the ConnectionName field value
// and a boolean to check if the value has been set.
func (o *OIDCPostRequest) GetConnectionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionName, true
}

// SetConnectionName sets field value
func (o *OIDCPostRequest) SetConnectionName(v string) {
	o.ConnectionName = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *OIDCPostRequest) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCPostRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *OIDCPostRequest) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *OIDCPostRequest) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *OIDCPostRequest) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCPostRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *OIDCPostRequest) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *OIDCPostRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetGuiOrder returns the GuiOrder field value if set, zero value otherwise.
func (o *OIDCPostRequest) GetGuiOrder() float32 {
	if o == nil || o.GuiOrder == nil {
		var ret float32
		return ret
	}
	return *o.GuiOrder
}

// GetGuiOrderOk returns a tuple with the GuiOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCPostRequest) GetGuiOrderOk() (*float32, bool) {
	if o == nil || o.GuiOrder == nil {
		return nil, false
	}
	return o.GuiOrder, true
}

// HasGuiOrder returns a boolean if a field has been set.
func (o *OIDCPostRequest) HasGuiOrder() bool {
	if o != nil && o.GuiOrder != nil {
		return true
	}

	return false
}

// SetGuiOrder gets a reference to the given float32 and assigns it to the GuiOrder field.
func (o *OIDCPostRequest) SetGuiOrder(v float32) {
	o.GuiOrder = &v
}

// GetUserInfoUrl returns the UserInfoUrl field value if set, zero value otherwise.
func (o *OIDCPostRequest) GetUserInfoUrl() string {
	if o == nil || o.UserInfoUrl == nil {
		var ret string
		return ret
	}
	return *o.UserInfoUrl
}

// GetUserInfoUrlOk returns a tuple with the UserInfoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCPostRequest) GetUserInfoUrlOk() (*string, bool) {
	if o == nil || o.UserInfoUrl == nil {
		return nil, false
	}
	return o.UserInfoUrl, true
}

// HasUserInfoUrl returns a boolean if a field has been set.
func (o *OIDCPostRequest) HasUserInfoUrl() bool {
	if o != nil && o.UserInfoUrl != nil {
		return true
	}

	return false
}

// SetUserInfoUrl gets a reference to the given string and assigns it to the UserInfoUrl field.
func (o *OIDCPostRequest) SetUserInfoUrl(v string) {
	o.UserInfoUrl = &v
}

// GetLogoutUrl returns the LogoutUrl field value if set, zero value otherwise.
func (o *OIDCPostRequest) GetLogoutUrl() string {
	if o == nil || o.LogoutUrl == nil {
		var ret string
		return ret
	}
	return *o.LogoutUrl
}

// GetLogoutUrlOk returns a tuple with the LogoutUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCPostRequest) GetLogoutUrlOk() (*string, bool) {
	if o == nil || o.LogoutUrl == nil {
		return nil, false
	}
	return o.LogoutUrl, true
}

// HasLogoutUrl returns a boolean if a field has been set.
func (o *OIDCPostRequest) HasLogoutUrl() bool {
	if o != nil && o.LogoutUrl != nil {
		return true
	}

	return false
}

// SetLogoutUrl gets a reference to the given string and assigns it to the LogoutUrl field.
func (o *OIDCPostRequest) SetLogoutUrl(v string) {
	o.LogoutUrl = &v
}

// GetAuthorizationUrl returns the AuthorizationUrl field value
func (o *OIDCPostRequest) GetAuthorizationUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthorizationUrl
}

// GetAuthorizationUrlOk returns a tuple with the AuthorizationUrl field value
// and a boolean to check if the value has been set.
func (o *OIDCPostRequest) GetAuthorizationUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizationUrl, true
}

// SetAuthorizationUrl sets field value
func (o *OIDCPostRequest) SetAuthorizationUrl(v string) {
	o.AuthorizationUrl = v
}

// GetTokenUrl returns the TokenUrl field value
func (o *OIDCPostRequest) GetTokenUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenUrl
}

// GetTokenUrlOk returns a tuple with the TokenUrl field value
// and a boolean to check if the value has been set.
func (o *OIDCPostRequest) GetTokenUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenUrl, true
}

// SetTokenUrl sets field value
func (o *OIDCPostRequest) SetTokenUrl(v string) {
	o.TokenUrl = v
}

// GetClientId returns the ClientId field value
func (o *OIDCPostRequest) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *OIDCPostRequest) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *OIDCPostRequest) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value
func (o *OIDCPostRequest) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *OIDCPostRequest) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *OIDCPostRequest) SetClientSecret(v string) {
	o.ClientSecret = v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *OIDCPostRequest) GetIssuer() string {
	if o == nil || o.Issuer == nil {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCPostRequest) GetIssuerOk() (*string, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *OIDCPostRequest) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *OIDCPostRequest) SetIssuer(v string) {
	o.Issuer = &v
}

// GetJwksUrl returns the JwksUrl field value if set, zero value otherwise.
func (o *OIDCPostRequest) GetJwksUrl() string {
	if o == nil || o.JwksUrl == nil {
		var ret string
		return ret
	}
	return *o.JwksUrl
}

// GetJwksUrlOk returns a tuple with the JwksUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCPostRequest) GetJwksUrlOk() (*string, bool) {
	if o == nil || o.JwksUrl == nil {
		return nil, false
	}
	return o.JwksUrl, true
}

// HasJwksUrl returns a boolean if a field has been set.
func (o *OIDCPostRequest) HasJwksUrl() bool {
	if o != nil && o.JwksUrl != nil {
		return true
	}

	return false
}

// SetJwksUrl gets a reference to the given string and assigns it to the JwksUrl field.
func (o *OIDCPostRequest) SetJwksUrl(v string) {
	o.JwksUrl = &v
}

func (o OIDCPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["connectionName"] = o.ConnectionName
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.GuiOrder != nil {
		toSerialize["guiOrder"] = o.GuiOrder
	}
	if o.UserInfoUrl != nil {
		toSerialize["userInfoUrl"] = o.UserInfoUrl
	}
	if o.LogoutUrl != nil {
		toSerialize["logoutUrl"] = o.LogoutUrl
	}
	if true {
		toSerialize["authorizationUrl"] = o.AuthorizationUrl
	}
	if true {
		toSerialize["tokenUrl"] = o.TokenUrl
	}
	if true {
		toSerialize["clientId"] = o.ClientId
	}
	if true {
		toSerialize["clientSecret"] = o.ClientSecret
	}
	if o.Issuer != nil {
		toSerialize["issuer"] = o.Issuer
	}
	if o.JwksUrl != nil {
		toSerialize["jwksUrl"] = o.JwksUrl
	}
	return json.Marshal(toSerialize)
}

type NullableOIDCPostRequest struct {
	value *OIDCPostRequest
	isSet bool
}

func (v NullableOIDCPostRequest) Get() *OIDCPostRequest {
	return v.value
}

func (v *NullableOIDCPostRequest) Set(val *OIDCPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOIDCPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOIDCPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOIDCPostRequest(val *OIDCPostRequest) *NullableOIDCPostRequest {
	return &NullableOIDCPostRequest{value: val, isSet: true}
}

func (v NullableOIDCPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOIDCPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


