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

// SuccessSamlConnectionResponseAllOfData struct for SuccessSamlConnectionResponseAllOfData
type SuccessSamlConnectionResponseAllOfData struct {
	// Unique ID of connection.
	ConnectionId *string `json:"connectionId,omitempty"`
	// Name of connection.
	ConnectionName *string `json:"connectionName,omitempty"`
	// Display Name of the connection
	DisplayName *string `json:"displayName,omitempty"`
	// Type of the connection (Google, Github, etc.).
	ConnectionType *string `json:"connectionType,omitempty"`
	// Sign In endpoint of connection
	SignInEndpoint *string `json:"signInEndpoint,omitempty"`
	// Sign Out Endpoint of connection
	SignOutEndpoint *string `json:"signOutEndpoint,omitempty"`
	// The certificate required for saml connection should be of x509Cert format
	PublicKeyCertificate *string `json:"publicKeyCertificate,omitempty"`
	// The Entity ID that will be used to uniquely identify this SAML Service Provider
	EntityId *string `json:"entityId,omitempty"`
	// The redirect uri to use when configuring the identity provider.
	RedirectUri *string `json:"redirectUri,omitempty"`
	// The URL on which saml connection metadata is present
	MetadataUri *string `json:"metadataUri,omitempty"`
	// GUI order of connection on Login page.
	GuiOrder *string `json:"guiOrder,omitempty"`
}

// NewSuccessSamlConnectionResponseAllOfData instantiates a new SuccessSamlConnectionResponseAllOfData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuccessSamlConnectionResponseAllOfData() *SuccessSamlConnectionResponseAllOfData {
	this := SuccessSamlConnectionResponseAllOfData{}
	return &this
}

// NewSuccessSamlConnectionResponseAllOfDataWithDefaults instantiates a new SuccessSamlConnectionResponseAllOfData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuccessSamlConnectionResponseAllOfDataWithDefaults() *SuccessSamlConnectionResponseAllOfData {
	this := SuccessSamlConnectionResponseAllOfData{}
	return &this
}

// GetConnectionId returns the ConnectionId field value if set, zero value otherwise.
func (o *SuccessSamlConnectionResponseAllOfData) GetConnectionId() string {
	if o == nil || o.ConnectionId == nil {
		var ret string
		return ret
	}
	return *o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessSamlConnectionResponseAllOfData) GetConnectionIdOk() (*string, bool) {
	if o == nil || o.ConnectionId == nil {
		return nil, false
	}
	return o.ConnectionId, true
}

// HasConnectionId returns a boolean if a field has been set.
func (o *SuccessSamlConnectionResponseAllOfData) HasConnectionId() bool {
	if o != nil && o.ConnectionId != nil {
		return true
	}

	return false
}

// SetConnectionId gets a reference to the given string and assigns it to the ConnectionId field.
func (o *SuccessSamlConnectionResponseAllOfData) SetConnectionId(v string) {
	o.ConnectionId = &v
}

// GetConnectionName returns the ConnectionName field value if set, zero value otherwise.
func (o *SuccessSamlConnectionResponseAllOfData) GetConnectionName() string {
	if o == nil || o.ConnectionName == nil {
		var ret string
		return ret
	}
	return *o.ConnectionName
}

// GetConnectionNameOk returns a tuple with the ConnectionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessSamlConnectionResponseAllOfData) GetConnectionNameOk() (*string, bool) {
	if o == nil || o.ConnectionName == nil {
		return nil, false
	}
	return o.ConnectionName, true
}

// HasConnectionName returns a boolean if a field has been set.
func (o *SuccessSamlConnectionResponseAllOfData) HasConnectionName() bool {
	if o != nil && o.ConnectionName != nil {
		return true
	}

	return false
}

// SetConnectionName gets a reference to the given string and assigns it to the ConnectionName field.
func (o *SuccessSamlConnectionResponseAllOfData) SetConnectionName(v string) {
	o.ConnectionName = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *SuccessSamlConnectionResponseAllOfData) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessSamlConnectionResponseAllOfData) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *SuccessSamlConnectionResponseAllOfData) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *SuccessSamlConnectionResponseAllOfData) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetConnectionType returns the ConnectionType field value if set, zero value otherwise.
func (o *SuccessSamlConnectionResponseAllOfData) GetConnectionType() string {
	if o == nil || o.ConnectionType == nil {
		var ret string
		return ret
	}
	return *o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessSamlConnectionResponseAllOfData) GetConnectionTypeOk() (*string, bool) {
	if o == nil || o.ConnectionType == nil {
		return nil, false
	}
	return o.ConnectionType, true
}

// HasConnectionType returns a boolean if a field has been set.
func (o *SuccessSamlConnectionResponseAllOfData) HasConnectionType() bool {
	if o != nil && o.ConnectionType != nil {
		return true
	}

	return false
}

// SetConnectionType gets a reference to the given string and assigns it to the ConnectionType field.
func (o *SuccessSamlConnectionResponseAllOfData) SetConnectionType(v string) {
	o.ConnectionType = &v
}

// GetSignInEndpoint returns the SignInEndpoint field value if set, zero value otherwise.
func (o *SuccessSamlConnectionResponseAllOfData) GetSignInEndpoint() string {
	if o == nil || o.SignInEndpoint == nil {
		var ret string
		return ret
	}
	return *o.SignInEndpoint
}

// GetSignInEndpointOk returns a tuple with the SignInEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessSamlConnectionResponseAllOfData) GetSignInEndpointOk() (*string, bool) {
	if o == nil || o.SignInEndpoint == nil {
		return nil, false
	}
	return o.SignInEndpoint, true
}

// HasSignInEndpoint returns a boolean if a field has been set.
func (o *SuccessSamlConnectionResponseAllOfData) HasSignInEndpoint() bool {
	if o != nil && o.SignInEndpoint != nil {
		return true
	}

	return false
}

// SetSignInEndpoint gets a reference to the given string and assigns it to the SignInEndpoint field.
func (o *SuccessSamlConnectionResponseAllOfData) SetSignInEndpoint(v string) {
	o.SignInEndpoint = &v
}

// GetSignOutEndpoint returns the SignOutEndpoint field value if set, zero value otherwise.
func (o *SuccessSamlConnectionResponseAllOfData) GetSignOutEndpoint() string {
	if o == nil || o.SignOutEndpoint == nil {
		var ret string
		return ret
	}
	return *o.SignOutEndpoint
}

// GetSignOutEndpointOk returns a tuple with the SignOutEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessSamlConnectionResponseAllOfData) GetSignOutEndpointOk() (*string, bool) {
	if o == nil || o.SignOutEndpoint == nil {
		return nil, false
	}
	return o.SignOutEndpoint, true
}

// HasSignOutEndpoint returns a boolean if a field has been set.
func (o *SuccessSamlConnectionResponseAllOfData) HasSignOutEndpoint() bool {
	if o != nil && o.SignOutEndpoint != nil {
		return true
	}

	return false
}

// SetSignOutEndpoint gets a reference to the given string and assigns it to the SignOutEndpoint field.
func (o *SuccessSamlConnectionResponseAllOfData) SetSignOutEndpoint(v string) {
	o.SignOutEndpoint = &v
}

// GetPublicKeyCertificate returns the PublicKeyCertificate field value if set, zero value otherwise.
func (o *SuccessSamlConnectionResponseAllOfData) GetPublicKeyCertificate() string {
	if o == nil || o.PublicKeyCertificate == nil {
		var ret string
		return ret
	}
	return *o.PublicKeyCertificate
}

// GetPublicKeyCertificateOk returns a tuple with the PublicKeyCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessSamlConnectionResponseAllOfData) GetPublicKeyCertificateOk() (*string, bool) {
	if o == nil || o.PublicKeyCertificate == nil {
		return nil, false
	}
	return o.PublicKeyCertificate, true
}

// HasPublicKeyCertificate returns a boolean if a field has been set.
func (o *SuccessSamlConnectionResponseAllOfData) HasPublicKeyCertificate() bool {
	if o != nil && o.PublicKeyCertificate != nil {
		return true
	}

	return false
}

// SetPublicKeyCertificate gets a reference to the given string and assigns it to the PublicKeyCertificate field.
func (o *SuccessSamlConnectionResponseAllOfData) SetPublicKeyCertificate(v string) {
	o.PublicKeyCertificate = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *SuccessSamlConnectionResponseAllOfData) GetEntityId() string {
	if o == nil || o.EntityId == nil {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessSamlConnectionResponseAllOfData) GetEntityIdOk() (*string, bool) {
	if o == nil || o.EntityId == nil {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *SuccessSamlConnectionResponseAllOfData) HasEntityId() bool {
	if o != nil && o.EntityId != nil {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *SuccessSamlConnectionResponseAllOfData) SetEntityId(v string) {
	o.EntityId = &v
}

// GetRedirectUri returns the RedirectUri field value if set, zero value otherwise.
func (o *SuccessSamlConnectionResponseAllOfData) GetRedirectUri() string {
	if o == nil || o.RedirectUri == nil {
		var ret string
		return ret
	}
	return *o.RedirectUri
}

// GetRedirectUriOk returns a tuple with the RedirectUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessSamlConnectionResponseAllOfData) GetRedirectUriOk() (*string, bool) {
	if o == nil || o.RedirectUri == nil {
		return nil, false
	}
	return o.RedirectUri, true
}

// HasRedirectUri returns a boolean if a field has been set.
func (o *SuccessSamlConnectionResponseAllOfData) HasRedirectUri() bool {
	if o != nil && o.RedirectUri != nil {
		return true
	}

	return false
}

// SetRedirectUri gets a reference to the given string and assigns it to the RedirectUri field.
func (o *SuccessSamlConnectionResponseAllOfData) SetRedirectUri(v string) {
	o.RedirectUri = &v
}

// GetMetadataUri returns the MetadataUri field value if set, zero value otherwise.
func (o *SuccessSamlConnectionResponseAllOfData) GetMetadataUri() string {
	if o == nil || o.MetadataUri == nil {
		var ret string
		return ret
	}
	return *o.MetadataUri
}

// GetMetadataUriOk returns a tuple with the MetadataUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessSamlConnectionResponseAllOfData) GetMetadataUriOk() (*string, bool) {
	if o == nil || o.MetadataUri == nil {
		return nil, false
	}
	return o.MetadataUri, true
}

// HasMetadataUri returns a boolean if a field has been set.
func (o *SuccessSamlConnectionResponseAllOfData) HasMetadataUri() bool {
	if o != nil && o.MetadataUri != nil {
		return true
	}

	return false
}

// SetMetadataUri gets a reference to the given string and assigns it to the MetadataUri field.
func (o *SuccessSamlConnectionResponseAllOfData) SetMetadataUri(v string) {
	o.MetadataUri = &v
}

// GetGuiOrder returns the GuiOrder field value if set, zero value otherwise.
func (o *SuccessSamlConnectionResponseAllOfData) GetGuiOrder() string {
	if o == nil || o.GuiOrder == nil {
		var ret string
		return ret
	}
	return *o.GuiOrder
}

// GetGuiOrderOk returns a tuple with the GuiOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessSamlConnectionResponseAllOfData) GetGuiOrderOk() (*string, bool) {
	if o == nil || o.GuiOrder == nil {
		return nil, false
	}
	return o.GuiOrder, true
}

// HasGuiOrder returns a boolean if a field has been set.
func (o *SuccessSamlConnectionResponseAllOfData) HasGuiOrder() bool {
	if o != nil && o.GuiOrder != nil {
		return true
	}

	return false
}

// SetGuiOrder gets a reference to the given string and assigns it to the GuiOrder field.
func (o *SuccessSamlConnectionResponseAllOfData) SetGuiOrder(v string) {
	o.GuiOrder = &v
}

func (o SuccessSamlConnectionResponseAllOfData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConnectionId != nil {
		toSerialize["connectionId"] = o.ConnectionId
	}
	if o.ConnectionName != nil {
		toSerialize["connectionName"] = o.ConnectionName
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.ConnectionType != nil {
		toSerialize["connectionType"] = o.ConnectionType
	}
	if o.SignInEndpoint != nil {
		toSerialize["signInEndpoint"] = o.SignInEndpoint
	}
	if o.SignOutEndpoint != nil {
		toSerialize["signOutEndpoint"] = o.SignOutEndpoint
	}
	if o.PublicKeyCertificate != nil {
		toSerialize["publicKeyCertificate"] = o.PublicKeyCertificate
	}
	if o.EntityId != nil {
		toSerialize["entityId"] = o.EntityId
	}
	if o.RedirectUri != nil {
		toSerialize["redirectUri"] = o.RedirectUri
	}
	if o.MetadataUri != nil {
		toSerialize["metadataUri"] = o.MetadataUri
	}
	if o.GuiOrder != nil {
		toSerialize["guiOrder"] = o.GuiOrder
	}
	return json.Marshal(toSerialize)
}

type NullableSuccessSamlConnectionResponseAllOfData struct {
	value *SuccessSamlConnectionResponseAllOfData
	isSet bool
}

func (v NullableSuccessSamlConnectionResponseAllOfData) Get() *SuccessSamlConnectionResponseAllOfData {
	return v.value
}

func (v *NullableSuccessSamlConnectionResponseAllOfData) Set(val *SuccessSamlConnectionResponseAllOfData) {
	v.value = val
	v.isSet = true
}

func (v NullableSuccessSamlConnectionResponseAllOfData) IsSet() bool {
	return v.isSet
}

func (v *NullableSuccessSamlConnectionResponseAllOfData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuccessSamlConnectionResponseAllOfData(val *SuccessSamlConnectionResponseAllOfData) *NullableSuccessSamlConnectionResponseAllOfData {
	return &NullableSuccessSamlConnectionResponseAllOfData{value: val, isSet: true}
}

func (v NullableSuccessSamlConnectionResponseAllOfData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuccessSamlConnectionResponseAllOfData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


