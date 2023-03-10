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

// UserResponse struct for UserResponse
type UserResponse struct {
	// Unique ID of user created at the time of creation.
	UserId *string `json:"userId,omitempty"`
	// User email. Email is in masked format, it reveals first 2 characters of email masking remaining characters with asterisk. However, the domain part is not masked.
	Email *string `json:"email,omitempty"`
	// First name of user.
	FirstName *string `json:"firstName,omitempty"`
	// Middle name of user.
	MiddleName *string `json:"middleName,omitempty"`
	// Last name of user.
	LastName *string `json:"lastName,omitempty"`
	// The participant identifier of the user. There are no validations for this field.
	ParticipantId NullableString `json:"participantId,omitempty"`
	AllowSocialAuth *SocialAuth `json:"allowSocialAuth,omitempty"`
	// A Boolean value that specifies whether SAML 2.0 Login is allowed.
	AllowSAMLAuth NullableBool `json:"allowSAMLAuth,omitempty"`
	SamlConnector *string `json:"samlConnector,omitempty"`
	SamlUserId *string `json:"samlUserId,omitempty"`
	// Whether user is active or not.
	IsActive *bool `json:"isActive,omitempty"`
	PrimaryMobile *MobileNumber `json:"primaryMobile,omitempty"`
	SecondaryMobile *MobileNumber `json:"secondaryMobile,omitempty"`
	// List of the group IDs user belong to.
	Groups []string `json:"groups,omitempty"`
}

// NewUserResponse instantiates a new UserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserResponse() *UserResponse {
	this := UserResponse{}
	var isActive bool = true
	this.IsActive = &isActive
	return &this
}

// NewUserResponseWithDefaults instantiates a new UserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserResponseWithDefaults() *UserResponse {
	this := UserResponse{}
	var isActive bool = true
	this.IsActive = &isActive
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UserResponse) GetUserId() string {
	if o == nil || isNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetUserIdOk() (*string, bool) {
	if o == nil || isNil(o.UserId) {
    return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UserResponse) HasUserId() bool {
	if o != nil && !isNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *UserResponse) SetUserId(v string) {
	o.UserId = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserResponse) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
    return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserResponse) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UserResponse) SetEmail(v string) {
	o.Email = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *UserResponse) GetFirstName() string {
	if o == nil || isNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetFirstNameOk() (*string, bool) {
	if o == nil || isNil(o.FirstName) {
    return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *UserResponse) HasFirstName() bool {
	if o != nil && !isNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *UserResponse) SetFirstName(v string) {
	o.FirstName = &v
}

// GetMiddleName returns the MiddleName field value if set, zero value otherwise.
func (o *UserResponse) GetMiddleName() string {
	if o == nil || isNil(o.MiddleName) {
		var ret string
		return ret
	}
	return *o.MiddleName
}

// GetMiddleNameOk returns a tuple with the MiddleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetMiddleNameOk() (*string, bool) {
	if o == nil || isNil(o.MiddleName) {
    return nil, false
	}
	return o.MiddleName, true
}

// HasMiddleName returns a boolean if a field has been set.
func (o *UserResponse) HasMiddleName() bool {
	if o != nil && !isNil(o.MiddleName) {
		return true
	}

	return false
}

// SetMiddleName gets a reference to the given string and assigns it to the MiddleName field.
func (o *UserResponse) SetMiddleName(v string) {
	o.MiddleName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *UserResponse) GetLastName() string {
	if o == nil || isNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetLastNameOk() (*string, bool) {
	if o == nil || isNil(o.LastName) {
    return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *UserResponse) HasLastName() bool {
	if o != nil && !isNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *UserResponse) SetLastName(v string) {
	o.LastName = &v
}

// GetParticipantId returns the ParticipantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserResponse) GetParticipantId() string {
	if o == nil || isNil(o.ParticipantId.Get()) {
		var ret string
		return ret
	}
	return *o.ParticipantId.Get()
}

// GetParticipantIdOk returns a tuple with the ParticipantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserResponse) GetParticipantIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ParticipantId.Get(), o.ParticipantId.IsSet()
}

// HasParticipantId returns a boolean if a field has been set.
func (o *UserResponse) HasParticipantId() bool {
	if o != nil && o.ParticipantId.IsSet() {
		return true
	}

	return false
}

// SetParticipantId gets a reference to the given NullableString and assigns it to the ParticipantId field.
func (o *UserResponse) SetParticipantId(v string) {
	o.ParticipantId.Set(&v)
}
// SetParticipantIdNil sets the value for ParticipantId to be an explicit nil
func (o *UserResponse) SetParticipantIdNil() {
	o.ParticipantId.Set(nil)
}

// UnsetParticipantId ensures that no value is present for ParticipantId, not even an explicit nil
func (o *UserResponse) UnsetParticipantId() {
	o.ParticipantId.Unset()
}

// GetAllowSocialAuth returns the AllowSocialAuth field value if set, zero value otherwise.
func (o *UserResponse) GetAllowSocialAuth() SocialAuth {
	if o == nil || isNil(o.AllowSocialAuth) {
		var ret SocialAuth
		return ret
	}
	return *o.AllowSocialAuth
}

// GetAllowSocialAuthOk returns a tuple with the AllowSocialAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetAllowSocialAuthOk() (*SocialAuth, bool) {
	if o == nil || isNil(o.AllowSocialAuth) {
    return nil, false
	}
	return o.AllowSocialAuth, true
}

// HasAllowSocialAuth returns a boolean if a field has been set.
func (o *UserResponse) HasAllowSocialAuth() bool {
	if o != nil && !isNil(o.AllowSocialAuth) {
		return true
	}

	return false
}

// SetAllowSocialAuth gets a reference to the given SocialAuth and assigns it to the AllowSocialAuth field.
func (o *UserResponse) SetAllowSocialAuth(v SocialAuth) {
	o.AllowSocialAuth = &v
}

// GetAllowSAMLAuth returns the AllowSAMLAuth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserResponse) GetAllowSAMLAuth() bool {
	if o == nil || isNil(o.AllowSAMLAuth.Get()) {
		var ret bool
		return ret
	}
	return *o.AllowSAMLAuth.Get()
}

// GetAllowSAMLAuthOk returns a tuple with the AllowSAMLAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserResponse) GetAllowSAMLAuthOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.AllowSAMLAuth.Get(), o.AllowSAMLAuth.IsSet()
}

// HasAllowSAMLAuth returns a boolean if a field has been set.
func (o *UserResponse) HasAllowSAMLAuth() bool {
	if o != nil && o.AllowSAMLAuth.IsSet() {
		return true
	}

	return false
}

// SetAllowSAMLAuth gets a reference to the given NullableBool and assigns it to the AllowSAMLAuth field.
func (o *UserResponse) SetAllowSAMLAuth(v bool) {
	o.AllowSAMLAuth.Set(&v)
}
// SetAllowSAMLAuthNil sets the value for AllowSAMLAuth to be an explicit nil
func (o *UserResponse) SetAllowSAMLAuthNil() {
	o.AllowSAMLAuth.Set(nil)
}

// UnsetAllowSAMLAuth ensures that no value is present for AllowSAMLAuth, not even an explicit nil
func (o *UserResponse) UnsetAllowSAMLAuth() {
	o.AllowSAMLAuth.Unset()
}

// GetSamlConnector returns the SamlConnector field value if set, zero value otherwise.
func (o *UserResponse) GetSamlConnector() string {
	if o == nil || isNil(o.SamlConnector) {
		var ret string
		return ret
	}
	return *o.SamlConnector
}

// GetSamlConnectorOk returns a tuple with the SamlConnector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetSamlConnectorOk() (*string, bool) {
	if o == nil || isNil(o.SamlConnector) {
    return nil, false
	}
	return o.SamlConnector, true
}

// HasSamlConnector returns a boolean if a field has been set.
func (o *UserResponse) HasSamlConnector() bool {
	if o != nil && !isNil(o.SamlConnector) {
		return true
	}

	return false
}

// SetSamlConnector gets a reference to the given string and assigns it to the SamlConnector field.
func (o *UserResponse) SetSamlConnector(v string) {
	o.SamlConnector = &v
}

// GetSamlUserId returns the SamlUserId field value if set, zero value otherwise.
func (o *UserResponse) GetSamlUserId() string {
	if o == nil || isNil(o.SamlUserId) {
		var ret string
		return ret
	}
	return *o.SamlUserId
}

// GetSamlUserIdOk returns a tuple with the SamlUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetSamlUserIdOk() (*string, bool) {
	if o == nil || isNil(o.SamlUserId) {
    return nil, false
	}
	return o.SamlUserId, true
}

// HasSamlUserId returns a boolean if a field has been set.
func (o *UserResponse) HasSamlUserId() bool {
	if o != nil && !isNil(o.SamlUserId) {
		return true
	}

	return false
}

// SetSamlUserId gets a reference to the given string and assigns it to the SamlUserId field.
func (o *UserResponse) SetSamlUserId(v string) {
	o.SamlUserId = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *UserResponse) GetIsActive() bool {
	if o == nil || isNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetIsActiveOk() (*bool, bool) {
	if o == nil || isNil(o.IsActive) {
    return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *UserResponse) HasIsActive() bool {
	if o != nil && !isNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *UserResponse) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetPrimaryMobile returns the PrimaryMobile field value if set, zero value otherwise.
func (o *UserResponse) GetPrimaryMobile() MobileNumber {
	if o == nil || isNil(o.PrimaryMobile) {
		var ret MobileNumber
		return ret
	}
	return *o.PrimaryMobile
}

// GetPrimaryMobileOk returns a tuple with the PrimaryMobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetPrimaryMobileOk() (*MobileNumber, bool) {
	if o == nil || isNil(o.PrimaryMobile) {
    return nil, false
	}
	return o.PrimaryMobile, true
}

// HasPrimaryMobile returns a boolean if a field has been set.
func (o *UserResponse) HasPrimaryMobile() bool {
	if o != nil && !isNil(o.PrimaryMobile) {
		return true
	}

	return false
}

// SetPrimaryMobile gets a reference to the given MobileNumber and assigns it to the PrimaryMobile field.
func (o *UserResponse) SetPrimaryMobile(v MobileNumber) {
	o.PrimaryMobile = &v
}

// GetSecondaryMobile returns the SecondaryMobile field value if set, zero value otherwise.
func (o *UserResponse) GetSecondaryMobile() MobileNumber {
	if o == nil || isNil(o.SecondaryMobile) {
		var ret MobileNumber
		return ret
	}
	return *o.SecondaryMobile
}

// GetSecondaryMobileOk returns a tuple with the SecondaryMobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetSecondaryMobileOk() (*MobileNumber, bool) {
	if o == nil || isNil(o.SecondaryMobile) {
    return nil, false
	}
	return o.SecondaryMobile, true
}

// HasSecondaryMobile returns a boolean if a field has been set.
func (o *UserResponse) HasSecondaryMobile() bool {
	if o != nil && !isNil(o.SecondaryMobile) {
		return true
	}

	return false
}

// SetSecondaryMobile gets a reference to the given MobileNumber and assigns it to the SecondaryMobile field.
func (o *UserResponse) SetSecondaryMobile(v MobileNumber) {
	o.SecondaryMobile = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *UserResponse) GetGroups() []string {
	if o == nil || isNil(o.Groups) {
		var ret []string
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetGroupsOk() ([]string, bool) {
	if o == nil || isNil(o.Groups) {
    return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *UserResponse) HasGroups() bool {
	if o != nil && !isNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []string and assigns it to the Groups field.
func (o *UserResponse) SetGroups(v []string) {
	o.Groups = v
}

func (o UserResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !isNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !isNil(o.MiddleName) {
		toSerialize["middleName"] = o.MiddleName
	}
	if !isNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	if o.ParticipantId.IsSet() {
		toSerialize["participantId"] = o.ParticipantId.Get()
	}
	if !isNil(o.AllowSocialAuth) {
		toSerialize["allowSocialAuth"] = o.AllowSocialAuth
	}
	if o.AllowSAMLAuth.IsSet() {
		toSerialize["allowSAMLAuth"] = o.AllowSAMLAuth.Get()
	}
	if !isNil(o.SamlConnector) {
		toSerialize["samlConnector"] = o.SamlConnector
	}
	if !isNil(o.SamlUserId) {
		toSerialize["samlUserId"] = o.SamlUserId
	}
	if !isNil(o.IsActive) {
		toSerialize["isActive"] = o.IsActive
	}
	if !isNil(o.PrimaryMobile) {
		toSerialize["primaryMobile"] = o.PrimaryMobile
	}
	if !isNil(o.SecondaryMobile) {
		toSerialize["secondaryMobile"] = o.SecondaryMobile
	}
	if !isNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	return json.Marshal(toSerialize)
}

type NullableUserResponse struct {
	value *UserResponse
	isSet bool
}

func (v NullableUserResponse) Get() *UserResponse {
	return v.value
}

func (v *NullableUserResponse) Set(val *UserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserResponse(val *UserResponse) *NullableUserResponse {
	return &NullableUserResponse{value: val, isSet: true}
}

func (v NullableUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


