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

// RoleBase struct for RoleBase
type RoleBase struct {
	// Unique ID assigned to each role at the time of creation.
	RoleId *string `json:"roleId,omitempty"`
	// Name of  the role
	RoleName *string `json:"roleName,omitempty"`
	// Description of the role (for example, role created for admin users).
	Description *string `json:"description,omitempty"`
	// Permissions attached with the role.
	Permissions []string `json:"permissions,omitempty"`
	// When `isActive` = False OR `isDeleted` = False, the role will be ignored for granting permissions.
	IsActive *bool `json:"isActive,omitempty"`
	// Whether the role is deleted or not. When `isActive` = False OR `isDeleted` = False, the role will be ignored for granting permissions.
	IsDeleted *bool `json:"isDeleted,omitempty"`
	// Boolean value indicates if this Role can be granted to user groups.
	CanGrantToUsers *bool `json:"canGrantToUsers,omitempty"`
	// Security level of the the role. Default is OPEN. Used to restrict the access level for the role.
	SecurityLevel *string `json:"securityLevel,omitempty"`
}

// NewRoleBase instantiates a new RoleBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleBase() *RoleBase {
	this := RoleBase{}
	return &this
}

// NewRoleBaseWithDefaults instantiates a new RoleBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleBaseWithDefaults() *RoleBase {
	this := RoleBase{}
	return &this
}

// GetRoleId returns the RoleId field value if set, zero value otherwise.
func (o *RoleBase) GetRoleId() string {
	if o == nil || isNil(o.RoleId) {
		var ret string
		return ret
	}
	return *o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleBase) GetRoleIdOk() (*string, bool) {
	if o == nil || isNil(o.RoleId) {
    return nil, false
	}
	return o.RoleId, true
}

// HasRoleId returns a boolean if a field has been set.
func (o *RoleBase) HasRoleId() bool {
	if o != nil && !isNil(o.RoleId) {
		return true
	}

	return false
}

// SetRoleId gets a reference to the given string and assigns it to the RoleId field.
func (o *RoleBase) SetRoleId(v string) {
	o.RoleId = &v
}

// GetRoleName returns the RoleName field value if set, zero value otherwise.
func (o *RoleBase) GetRoleName() string {
	if o == nil || isNil(o.RoleName) {
		var ret string
		return ret
	}
	return *o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleBase) GetRoleNameOk() (*string, bool) {
	if o == nil || isNil(o.RoleName) {
    return nil, false
	}
	return o.RoleName, true
}

// HasRoleName returns a boolean if a field has been set.
func (o *RoleBase) HasRoleName() bool {
	if o != nil && !isNil(o.RoleName) {
		return true
	}

	return false
}

// SetRoleName gets a reference to the given string and assigns it to the RoleName field.
func (o *RoleBase) SetRoleName(v string) {
	o.RoleName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RoleBase) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleBase) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RoleBase) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RoleBase) SetDescription(v string) {
	o.Description = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoleBase) GetPermissions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoleBase) GetPermissionsOk() ([]string, bool) {
	if o == nil || isNil(o.Permissions) {
    return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *RoleBase) HasPermissions() bool {
	if o != nil && isNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []string and assigns it to the Permissions field.
func (o *RoleBase) SetPermissions(v []string) {
	o.Permissions = v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *RoleBase) GetIsActive() bool {
	if o == nil || isNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleBase) GetIsActiveOk() (*bool, bool) {
	if o == nil || isNil(o.IsActive) {
    return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *RoleBase) HasIsActive() bool {
	if o != nil && !isNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *RoleBase) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *RoleBase) GetIsDeleted() bool {
	if o == nil || isNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleBase) GetIsDeletedOk() (*bool, bool) {
	if o == nil || isNil(o.IsDeleted) {
    return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *RoleBase) HasIsDeleted() bool {
	if o != nil && !isNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *RoleBase) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetCanGrantToUsers returns the CanGrantToUsers field value if set, zero value otherwise.
func (o *RoleBase) GetCanGrantToUsers() bool {
	if o == nil || isNil(o.CanGrantToUsers) {
		var ret bool
		return ret
	}
	return *o.CanGrantToUsers
}

// GetCanGrantToUsersOk returns a tuple with the CanGrantToUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleBase) GetCanGrantToUsersOk() (*bool, bool) {
	if o == nil || isNil(o.CanGrantToUsers) {
    return nil, false
	}
	return o.CanGrantToUsers, true
}

// HasCanGrantToUsers returns a boolean if a field has been set.
func (o *RoleBase) HasCanGrantToUsers() bool {
	if o != nil && !isNil(o.CanGrantToUsers) {
		return true
	}

	return false
}

// SetCanGrantToUsers gets a reference to the given bool and assigns it to the CanGrantToUsers field.
func (o *RoleBase) SetCanGrantToUsers(v bool) {
	o.CanGrantToUsers = &v
}

// GetSecurityLevel returns the SecurityLevel field value if set, zero value otherwise.
func (o *RoleBase) GetSecurityLevel() string {
	if o == nil || isNil(o.SecurityLevel) {
		var ret string
		return ret
	}
	return *o.SecurityLevel
}

// GetSecurityLevelOk returns a tuple with the SecurityLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleBase) GetSecurityLevelOk() (*string, bool) {
	if o == nil || isNil(o.SecurityLevel) {
    return nil, false
	}
	return o.SecurityLevel, true
}

// HasSecurityLevel returns a boolean if a field has been set.
func (o *RoleBase) HasSecurityLevel() bool {
	if o != nil && !isNil(o.SecurityLevel) {
		return true
	}

	return false
}

// SetSecurityLevel gets a reference to the given string and assigns it to the SecurityLevel field.
func (o *RoleBase) SetSecurityLevel(v string) {
	o.SecurityLevel = &v
}

func (o RoleBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RoleId) {
		toSerialize["roleId"] = o.RoleId
	}
	if !isNil(o.RoleName) {
		toSerialize["roleName"] = o.RoleName
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	if !isNil(o.IsActive) {
		toSerialize["isActive"] = o.IsActive
	}
	if !isNil(o.IsDeleted) {
		toSerialize["isDeleted"] = o.IsDeleted
	}
	if !isNil(o.CanGrantToUsers) {
		toSerialize["canGrantToUsers"] = o.CanGrantToUsers
	}
	if !isNil(o.SecurityLevel) {
		toSerialize["securityLevel"] = o.SecurityLevel
	}
	return json.Marshal(toSerialize)
}

type NullableRoleBase struct {
	value *RoleBase
	isSet bool
}

func (v NullableRoleBase) Get() *RoleBase {
	return v.value
}

func (v *NullableRoleBase) Set(val *RoleBase) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleBase) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleBase(val *RoleBase) *NullableRoleBase {
	return &NullableRoleBase{value: val, isSet: true}
}

func (v NullableRoleBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


