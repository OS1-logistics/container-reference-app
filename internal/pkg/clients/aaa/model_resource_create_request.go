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

// ResourceCreateRequest struct for ResourceCreateRequest
type ResourceCreateRequest struct {
	// Name of the entity exposed by this resource.
	Name string `json:"name"`
	//  Description of the entity exposed by this resource.
	Description *string `json:"description,omitempty"`
	// URI for resource. Example \"/core/api/v1/containers/_**\"
	ResourcePath string `json:"resourcePath"`
	// Formatted resourcePath. Example \"/core/api/v1/containers/{{containerId}}\"
	ResourcePathFormatted NullableString `json:"resourcePathFormatted,omitempty"`
	// JSON path in ValidateAPI response for requestParameter value. JSON path value will be matched with requestParameter value. This field will only be present in abac related resources.
	AttributePath NullableString `json:"attributePath,omitempty"`
	// Request Path Param name for the attribute. It will save the name of the ABAC attribute.
	RequestParameter NullableString `json:"requestParameter,omitempty"`
	// Scopes of resource: get, update, create.
	AllowedHttpMethods []string `json:"allowedHttpMethods"`
}

// NewResourceCreateRequest instantiates a new ResourceCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceCreateRequest(name string, resourcePath string, allowedHttpMethods []string) *ResourceCreateRequest {
	this := ResourceCreateRequest{}
	this.Name = name
	this.ResourcePath = resourcePath
	this.AllowedHttpMethods = allowedHttpMethods
	return &this
}

// NewResourceCreateRequestWithDefaults instantiates a new ResourceCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceCreateRequestWithDefaults() *ResourceCreateRequest {
	this := ResourceCreateRequest{}
	return &this
}

// GetName returns the Name field value
func (o *ResourceCreateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ResourceCreateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ResourceCreateRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ResourceCreateRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceCreateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ResourceCreateRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ResourceCreateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetResourcePath returns the ResourcePath field value
func (o *ResourceCreateRequest) GetResourcePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourcePath
}

// GetResourcePathOk returns a tuple with the ResourcePath field value
// and a boolean to check if the value has been set.
func (o *ResourceCreateRequest) GetResourcePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourcePath, true
}

// SetResourcePath sets field value
func (o *ResourceCreateRequest) SetResourcePath(v string) {
	o.ResourcePath = v
}

// GetResourcePathFormatted returns the ResourcePathFormatted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceCreateRequest) GetResourcePathFormatted() string {
	if o == nil || o.ResourcePathFormatted.Get() == nil {
		var ret string
		return ret
	}
	return *o.ResourcePathFormatted.Get()
}

// GetResourcePathFormattedOk returns a tuple with the ResourcePathFormatted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceCreateRequest) GetResourcePathFormattedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourcePathFormatted.Get(), o.ResourcePathFormatted.IsSet()
}

// HasResourcePathFormatted returns a boolean if a field has been set.
func (o *ResourceCreateRequest) HasResourcePathFormatted() bool {
	if o != nil && o.ResourcePathFormatted.IsSet() {
		return true
	}

	return false
}

// SetResourcePathFormatted gets a reference to the given NullableString and assigns it to the ResourcePathFormatted field.
func (o *ResourceCreateRequest) SetResourcePathFormatted(v string) {
	o.ResourcePathFormatted.Set(&v)
}
// SetResourcePathFormattedNil sets the value for ResourcePathFormatted to be an explicit nil
func (o *ResourceCreateRequest) SetResourcePathFormattedNil() {
	o.ResourcePathFormatted.Set(nil)
}

// UnsetResourcePathFormatted ensures that no value is present for ResourcePathFormatted, not even an explicit nil
func (o *ResourceCreateRequest) UnsetResourcePathFormatted() {
	o.ResourcePathFormatted.Unset()
}

// GetAttributePath returns the AttributePath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceCreateRequest) GetAttributePath() string {
	if o == nil || o.AttributePath.Get() == nil {
		var ret string
		return ret
	}
	return *o.AttributePath.Get()
}

// GetAttributePathOk returns a tuple with the AttributePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceCreateRequest) GetAttributePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AttributePath.Get(), o.AttributePath.IsSet()
}

// HasAttributePath returns a boolean if a field has been set.
func (o *ResourceCreateRequest) HasAttributePath() bool {
	if o != nil && o.AttributePath.IsSet() {
		return true
	}

	return false
}

// SetAttributePath gets a reference to the given NullableString and assigns it to the AttributePath field.
func (o *ResourceCreateRequest) SetAttributePath(v string) {
	o.AttributePath.Set(&v)
}
// SetAttributePathNil sets the value for AttributePath to be an explicit nil
func (o *ResourceCreateRequest) SetAttributePathNil() {
	o.AttributePath.Set(nil)
}

// UnsetAttributePath ensures that no value is present for AttributePath, not even an explicit nil
func (o *ResourceCreateRequest) UnsetAttributePath() {
	o.AttributePath.Unset()
}

// GetRequestParameter returns the RequestParameter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceCreateRequest) GetRequestParameter() string {
	if o == nil || o.RequestParameter.Get() == nil {
		var ret string
		return ret
	}
	return *o.RequestParameter.Get()
}

// GetRequestParameterOk returns a tuple with the RequestParameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceCreateRequest) GetRequestParameterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestParameter.Get(), o.RequestParameter.IsSet()
}

// HasRequestParameter returns a boolean if a field has been set.
func (o *ResourceCreateRequest) HasRequestParameter() bool {
	if o != nil && o.RequestParameter.IsSet() {
		return true
	}

	return false
}

// SetRequestParameter gets a reference to the given NullableString and assigns it to the RequestParameter field.
func (o *ResourceCreateRequest) SetRequestParameter(v string) {
	o.RequestParameter.Set(&v)
}
// SetRequestParameterNil sets the value for RequestParameter to be an explicit nil
func (o *ResourceCreateRequest) SetRequestParameterNil() {
	o.RequestParameter.Set(nil)
}

// UnsetRequestParameter ensures that no value is present for RequestParameter, not even an explicit nil
func (o *ResourceCreateRequest) UnsetRequestParameter() {
	o.RequestParameter.Unset()
}

// GetAllowedHttpMethods returns the AllowedHttpMethods field value
func (o *ResourceCreateRequest) GetAllowedHttpMethods() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AllowedHttpMethods
}

// GetAllowedHttpMethodsOk returns a tuple with the AllowedHttpMethods field value
// and a boolean to check if the value has been set.
func (o *ResourceCreateRequest) GetAllowedHttpMethodsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowedHttpMethods, true
}

// SetAllowedHttpMethods sets field value
func (o *ResourceCreateRequest) SetAllowedHttpMethods(v []string) {
	o.AllowedHttpMethods = v
}

func (o ResourceCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["resourcePath"] = o.ResourcePath
	}
	if o.ResourcePathFormatted.IsSet() {
		toSerialize["resourcePathFormatted"] = o.ResourcePathFormatted.Get()
	}
	if o.AttributePath.IsSet() {
		toSerialize["attributePath"] = o.AttributePath.Get()
	}
	if o.RequestParameter.IsSet() {
		toSerialize["requestParameter"] = o.RequestParameter.Get()
	}
	if true {
		toSerialize["allowedHttpMethods"] = o.AllowedHttpMethods
	}
	return json.Marshal(toSerialize)
}

type NullableResourceCreateRequest struct {
	value *ResourceCreateRequest
	isSet bool
}

func (v NullableResourceCreateRequest) Get() *ResourceCreateRequest {
	return v.value
}

func (v *NullableResourceCreateRequest) Set(val *ResourceCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceCreateRequest(val *ResourceCreateRequest) *NullableResourceCreateRequest {
	return &NullableResourceCreateRequest{value: val, isSet: true}
}

func (v NullableResourceCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


