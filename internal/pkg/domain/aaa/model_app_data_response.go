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

// AppDataResponse struct for AppDataResponse
type AppDataResponse struct {
	// Unique app ID of the app.
	Id *string `json:"id,omitempty"`
	// Description of the application.
	Description NullableString `json:"description,omitempty"`
	// Set if app is private and can be allocated only to the Tenant specified in the field `privateTenantId`.
	IsPrivateApp *bool `json:"isPrivateApp,omitempty"`
	// Will be ignored if `isPrivate` is false.
	PrivateTenantId *string `json:"privateTenantId,omitempty"`
	// When `isActive` = False OR `isDeleted` = False, the role will be ignored for granting permissions.
	IsActive *bool `json:"isActive,omitempty"`
	// Array of allowed redirect URIs.
	RedirectUri []string `json:"redirectUri,omitempty"`
}

// NewAppDataResponse instantiates a new AppDataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppDataResponse() *AppDataResponse {
	this := AppDataResponse{}
	var isPrivateApp bool = false
	this.IsPrivateApp = &isPrivateApp
	return &this
}

// NewAppDataResponseWithDefaults instantiates a new AppDataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppDataResponseWithDefaults() *AppDataResponse {
	this := AppDataResponse{}
	var isPrivateApp bool = false
	this.IsPrivateApp = &isPrivateApp
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AppDataResponse) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDataResponse) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AppDataResponse) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AppDataResponse) SetId(v string) {
	o.Id = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppDataResponse) GetDescription() string {
	if o == nil || isNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppDataResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *AppDataResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *AppDataResponse) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *AppDataResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *AppDataResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetIsPrivateApp returns the IsPrivateApp field value if set, zero value otherwise.
func (o *AppDataResponse) GetIsPrivateApp() bool {
	if o == nil || isNil(o.IsPrivateApp) {
		var ret bool
		return ret
	}
	return *o.IsPrivateApp
}

// GetIsPrivateAppOk returns a tuple with the IsPrivateApp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDataResponse) GetIsPrivateAppOk() (*bool, bool) {
	if o == nil || isNil(o.IsPrivateApp) {
    return nil, false
	}
	return o.IsPrivateApp, true
}

// HasIsPrivateApp returns a boolean if a field has been set.
func (o *AppDataResponse) HasIsPrivateApp() bool {
	if o != nil && !isNil(o.IsPrivateApp) {
		return true
	}

	return false
}

// SetIsPrivateApp gets a reference to the given bool and assigns it to the IsPrivateApp field.
func (o *AppDataResponse) SetIsPrivateApp(v bool) {
	o.IsPrivateApp = &v
}

// GetPrivateTenantId returns the PrivateTenantId field value if set, zero value otherwise.
func (o *AppDataResponse) GetPrivateTenantId() string {
	if o == nil || isNil(o.PrivateTenantId) {
		var ret string
		return ret
	}
	return *o.PrivateTenantId
}

// GetPrivateTenantIdOk returns a tuple with the PrivateTenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDataResponse) GetPrivateTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.PrivateTenantId) {
    return nil, false
	}
	return o.PrivateTenantId, true
}

// HasPrivateTenantId returns a boolean if a field has been set.
func (o *AppDataResponse) HasPrivateTenantId() bool {
	if o != nil && !isNil(o.PrivateTenantId) {
		return true
	}

	return false
}

// SetPrivateTenantId gets a reference to the given string and assigns it to the PrivateTenantId field.
func (o *AppDataResponse) SetPrivateTenantId(v string) {
	o.PrivateTenantId = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *AppDataResponse) GetIsActive() bool {
	if o == nil || isNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDataResponse) GetIsActiveOk() (*bool, bool) {
	if o == nil || isNil(o.IsActive) {
    return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *AppDataResponse) HasIsActive() bool {
	if o != nil && !isNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *AppDataResponse) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetRedirectUri returns the RedirectUri field value if set, zero value otherwise.
func (o *AppDataResponse) GetRedirectUri() []string {
	if o == nil || isNil(o.RedirectUri) {
		var ret []string
		return ret
	}
	return o.RedirectUri
}

// GetRedirectUriOk returns a tuple with the RedirectUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDataResponse) GetRedirectUriOk() ([]string, bool) {
	if o == nil || isNil(o.RedirectUri) {
    return nil, false
	}
	return o.RedirectUri, true
}

// HasRedirectUri returns a boolean if a field has been set.
func (o *AppDataResponse) HasRedirectUri() bool {
	if o != nil && !isNil(o.RedirectUri) {
		return true
	}

	return false
}

// SetRedirectUri gets a reference to the given []string and assigns it to the RedirectUri field.
func (o *AppDataResponse) SetRedirectUri(v []string) {
	o.RedirectUri = v
}

func (o AppDataResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !isNil(o.IsPrivateApp) {
		toSerialize["isPrivateApp"] = o.IsPrivateApp
	}
	if !isNil(o.PrivateTenantId) {
		toSerialize["privateTenantId"] = o.PrivateTenantId
	}
	if !isNil(o.IsActive) {
		toSerialize["isActive"] = o.IsActive
	}
	if !isNil(o.RedirectUri) {
		toSerialize["redirectUri"] = o.RedirectUri
	}
	return json.Marshal(toSerialize)
}

type NullableAppDataResponse struct {
	value *AppDataResponse
	isSet bool
}

func (v NullableAppDataResponse) Get() *AppDataResponse {
	return v.value
}

func (v *NullableAppDataResponse) Set(val *AppDataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppDataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppDataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppDataResponse(val *AppDataResponse) *NullableAppDataResponse {
	return &NullableAppDataResponse{value: val, isSet: true}
}

func (v NullableAppDataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppDataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


