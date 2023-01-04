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

// OIDCUpdateRequest struct for OIDCUpdateRequest
type OIDCUpdateRequest struct {
	// Display name of the OIDC provider.
	DisplayName *string `json:"displayName,omitempty"`
	// When on the provider is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Order of the provider in GUI.
	GuiOrder *float32 `json:"guiOrder,omitempty"`
	// User info URL.
	UserInfoUrl *string `json:"userInfoUrl,omitempty"`
	// Log out URL.
	LogoutUrl *string `json:"logoutUrl,omitempty"`
	// Authorization URL.
	AuthorizationUrl *string `json:"authorizationUrl,omitempty"`
	// Token URL.
	TokenUrl *string `json:"tokenUrl,omitempty"`
	// Client identifier registered with identity provider.
	ClientId *string `json:"clientId,omitempty"`
	// Client secret.
	ClientSecret *string `json:"clientSecret,omitempty"`
	// Issuer of response.
	Issuer *string `json:"issuer,omitempty"`
	// JWKS URL.
	JwksUrl *string `json:"jwksUrl,omitempty"`
}

// NewOIDCUpdateRequest instantiates a new OIDCUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOIDCUpdateRequest() *OIDCUpdateRequest {
	this := OIDCUpdateRequest{}
	return &this
}

// NewOIDCUpdateRequestWithDefaults instantiates a new OIDCUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOIDCUpdateRequestWithDefaults() *OIDCUpdateRequest {
	this := OIDCUpdateRequest{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *OIDCUpdateRequest) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCUpdateRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
    return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *OIDCUpdateRequest) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *OIDCUpdateRequest) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *OIDCUpdateRequest) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCUpdateRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *OIDCUpdateRequest) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *OIDCUpdateRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetGuiOrder returns the GuiOrder field value if set, zero value otherwise.
func (o *OIDCUpdateRequest) GetGuiOrder() float32 {
	if o == nil || isNil(o.GuiOrder) {
		var ret float32
		return ret
	}
	return *o.GuiOrder
}

// GetGuiOrderOk returns a tuple with the GuiOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCUpdateRequest) GetGuiOrderOk() (*float32, bool) {
	if o == nil || isNil(o.GuiOrder) {
    return nil, false
	}
	return o.GuiOrder, true
}

// HasGuiOrder returns a boolean if a field has been set.
func (o *OIDCUpdateRequest) HasGuiOrder() bool {
	if o != nil && !isNil(o.GuiOrder) {
		return true
	}

	return false
}

// SetGuiOrder gets a reference to the given float32 and assigns it to the GuiOrder field.
func (o *OIDCUpdateRequest) SetGuiOrder(v float32) {
	o.GuiOrder = &v
}

// GetUserInfoUrl returns the UserInfoUrl field value if set, zero value otherwise.
func (o *OIDCUpdateRequest) GetUserInfoUrl() string {
	if o == nil || isNil(o.UserInfoUrl) {
		var ret string
		return ret
	}
	return *o.UserInfoUrl
}

// GetUserInfoUrlOk returns a tuple with the UserInfoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCUpdateRequest) GetUserInfoUrlOk() (*string, bool) {
	if o == nil || isNil(o.UserInfoUrl) {
    return nil, false
	}
	return o.UserInfoUrl, true
}

// HasUserInfoUrl returns a boolean if a field has been set.
func (o *OIDCUpdateRequest) HasUserInfoUrl() bool {
	if o != nil && !isNil(o.UserInfoUrl) {
		return true
	}

	return false
}

// SetUserInfoUrl gets a reference to the given string and assigns it to the UserInfoUrl field.
func (o *OIDCUpdateRequest) SetUserInfoUrl(v string) {
	o.UserInfoUrl = &v
}

// GetLogoutUrl returns the LogoutUrl field value if set, zero value otherwise.
func (o *OIDCUpdateRequest) GetLogoutUrl() string {
	if o == nil || isNil(o.LogoutUrl) {
		var ret string
		return ret
	}
	return *o.LogoutUrl
}

// GetLogoutUrlOk returns a tuple with the LogoutUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCUpdateRequest) GetLogoutUrlOk() (*string, bool) {
	if o == nil || isNil(o.LogoutUrl) {
    return nil, false
	}
	return o.LogoutUrl, true
}

// HasLogoutUrl returns a boolean if a field has been set.
func (o *OIDCUpdateRequest) HasLogoutUrl() bool {
	if o != nil && !isNil(o.LogoutUrl) {
		return true
	}

	return false
}

// SetLogoutUrl gets a reference to the given string and assigns it to the LogoutUrl field.
func (o *OIDCUpdateRequest) SetLogoutUrl(v string) {
	o.LogoutUrl = &v
}

// GetAuthorizationUrl returns the AuthorizationUrl field value if set, zero value otherwise.
func (o *OIDCUpdateRequest) GetAuthorizationUrl() string {
	if o == nil || isNil(o.AuthorizationUrl) {
		var ret string
		return ret
	}
	return *o.AuthorizationUrl
}

// GetAuthorizationUrlOk returns a tuple with the AuthorizationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCUpdateRequest) GetAuthorizationUrlOk() (*string, bool) {
	if o == nil || isNil(o.AuthorizationUrl) {
    return nil, false
	}
	return o.AuthorizationUrl, true
}

// HasAuthorizationUrl returns a boolean if a field has been set.
func (o *OIDCUpdateRequest) HasAuthorizationUrl() bool {
	if o != nil && !isNil(o.AuthorizationUrl) {
		return true
	}

	return false
}

// SetAuthorizationUrl gets a reference to the given string and assigns it to the AuthorizationUrl field.
func (o *OIDCUpdateRequest) SetAuthorizationUrl(v string) {
	o.AuthorizationUrl = &v
}

// GetTokenUrl returns the TokenUrl field value if set, zero value otherwise.
func (o *OIDCUpdateRequest) GetTokenUrl() string {
	if o == nil || isNil(o.TokenUrl) {
		var ret string
		return ret
	}
	return *o.TokenUrl
}

// GetTokenUrlOk returns a tuple with the TokenUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCUpdateRequest) GetTokenUrlOk() (*string, bool) {
	if o == nil || isNil(o.TokenUrl) {
    return nil, false
	}
	return o.TokenUrl, true
}

// HasTokenUrl returns a boolean if a field has been set.
func (o *OIDCUpdateRequest) HasTokenUrl() bool {
	if o != nil && !isNil(o.TokenUrl) {
		return true
	}

	return false
}

// SetTokenUrl gets a reference to the given string and assigns it to the TokenUrl field.
func (o *OIDCUpdateRequest) SetTokenUrl(v string) {
	o.TokenUrl = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *OIDCUpdateRequest) GetClientId() string {
	if o == nil || isNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCUpdateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || isNil(o.ClientId) {
    return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *OIDCUpdateRequest) HasClientId() bool {
	if o != nil && !isNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *OIDCUpdateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *OIDCUpdateRequest) GetClientSecret() string {
	if o == nil || isNil(o.ClientSecret) {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCUpdateRequest) GetClientSecretOk() (*string, bool) {
	if o == nil || isNil(o.ClientSecret) {
    return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *OIDCUpdateRequest) HasClientSecret() bool {
	if o != nil && !isNil(o.ClientSecret) {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *OIDCUpdateRequest) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *OIDCUpdateRequest) GetIssuer() string {
	if o == nil || isNil(o.Issuer) {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCUpdateRequest) GetIssuerOk() (*string, bool) {
	if o == nil || isNil(o.Issuer) {
    return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *OIDCUpdateRequest) HasIssuer() bool {
	if o != nil && !isNil(o.Issuer) {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *OIDCUpdateRequest) SetIssuer(v string) {
	o.Issuer = &v
}

// GetJwksUrl returns the JwksUrl field value if set, zero value otherwise.
func (o *OIDCUpdateRequest) GetJwksUrl() string {
	if o == nil || isNil(o.JwksUrl) {
		var ret string
		return ret
	}
	return *o.JwksUrl
}

// GetJwksUrlOk returns a tuple with the JwksUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCUpdateRequest) GetJwksUrlOk() (*string, bool) {
	if o == nil || isNil(o.JwksUrl) {
    return nil, false
	}
	return o.JwksUrl, true
}

// HasJwksUrl returns a boolean if a field has been set.
func (o *OIDCUpdateRequest) HasJwksUrl() bool {
	if o != nil && !isNil(o.JwksUrl) {
		return true
	}

	return false
}

// SetJwksUrl gets a reference to the given string and assigns it to the JwksUrl field.
func (o *OIDCUpdateRequest) SetJwksUrl(v string) {
	o.JwksUrl = &v
}

func (o OIDCUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.GuiOrder) {
		toSerialize["guiOrder"] = o.GuiOrder
	}
	if !isNil(o.UserInfoUrl) {
		toSerialize["userInfoUrl"] = o.UserInfoUrl
	}
	if !isNil(o.LogoutUrl) {
		toSerialize["logoutUrl"] = o.LogoutUrl
	}
	if !isNil(o.AuthorizationUrl) {
		toSerialize["authorizationUrl"] = o.AuthorizationUrl
	}
	if !isNil(o.TokenUrl) {
		toSerialize["tokenUrl"] = o.TokenUrl
	}
	if !isNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !isNil(o.ClientSecret) {
		toSerialize["clientSecret"] = o.ClientSecret
	}
	if !isNil(o.Issuer) {
		toSerialize["issuer"] = o.Issuer
	}
	if !isNil(o.JwksUrl) {
		toSerialize["jwksUrl"] = o.JwksUrl
	}
	return json.Marshal(toSerialize)
}

type NullableOIDCUpdateRequest struct {
	value *OIDCUpdateRequest
	isSet bool
}

func (v NullableOIDCUpdateRequest) Get() *OIDCUpdateRequest {
	return v.value
}

func (v *NullableOIDCUpdateRequest) Set(val *OIDCUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOIDCUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOIDCUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOIDCUpdateRequest(val *OIDCUpdateRequest) *NullableOIDCUpdateRequest {
	return &NullableOIDCUpdateRequest{value: val, isSet: true}
}

func (v NullableOIDCUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOIDCUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


