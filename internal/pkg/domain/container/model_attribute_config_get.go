/*
Container Service

**API documentation for Container Service**

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package containerdomain

import (
	"encoding/json"
)

// AttributeConfigGet struct for AttributeConfigGet
type AttributeConfigGet struct {
	// Field to define attribute name.
	Name string `json:"name"`
	DataType DataType `json:"dataType"`
	Description *string `json:"description,omitempty"`
	// Field to specify if the attribute is to be indexed. Filter or search operation on basis of a custom attribute will be only allowed if this field is set as TRUE.
	Indexed *bool `json:"indexed,omitempty"`
	// Access Parameter for the Attribute
	IsReadPublic *bool `json:"isReadPublic,omitempty"`
	OwnerAppId *string `json:"ownerAppId,omitempty"`
	DefaultValue *AttributeConfigDefaultValue `json:"defaultValue,omitempty"`
	Validation *AttributeValidation `json:"validation,omitempty"`
}

// NewAttributeConfigGet instantiates a new AttributeConfigGet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttributeConfigGet(name string, dataType DataType) *AttributeConfigGet {
	this := AttributeConfigGet{}
	this.Name = name
	this.DataType = dataType
	var isReadPublic bool = false
	this.IsReadPublic = &isReadPublic
	return &this
}

// NewAttributeConfigGetWithDefaults instantiates a new AttributeConfigGet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttributeConfigGetWithDefaults() *AttributeConfigGet {
	this := AttributeConfigGet{}
	var isReadPublic bool = false
	this.IsReadPublic = &isReadPublic
	return &this
}

// GetName returns the Name field value
func (o *AttributeConfigGet) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AttributeConfigGet) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AttributeConfigGet) SetName(v string) {
	o.Name = v
}

// GetDataType returns the DataType field value
func (o *AttributeConfigGet) GetDataType() DataType {
	if o == nil {
		var ret DataType
		return ret
	}

	return o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value
// and a boolean to check if the value has been set.
func (o *AttributeConfigGet) GetDataTypeOk() (*DataType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DataType, true
}

// SetDataType sets field value
func (o *AttributeConfigGet) SetDataType(v DataType) {
	o.DataType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AttributeConfigGet) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeConfigGet) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AttributeConfigGet) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AttributeConfigGet) SetDescription(v string) {
	o.Description = &v
}

// GetIndexed returns the Indexed field value if set, zero value otherwise.
func (o *AttributeConfigGet) GetIndexed() bool {
	if o == nil || isNil(o.Indexed) {
		var ret bool
		return ret
	}
	return *o.Indexed
}

// GetIndexedOk returns a tuple with the Indexed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeConfigGet) GetIndexedOk() (*bool, bool) {
	if o == nil || isNil(o.Indexed) {
    return nil, false
	}
	return o.Indexed, true
}

// HasIndexed returns a boolean if a field has been set.
func (o *AttributeConfigGet) HasIndexed() bool {
	if o != nil && !isNil(o.Indexed) {
		return true
	}

	return false
}

// SetIndexed gets a reference to the given bool and assigns it to the Indexed field.
func (o *AttributeConfigGet) SetIndexed(v bool) {
	o.Indexed = &v
}

// GetIsReadPublic returns the IsReadPublic field value if set, zero value otherwise.
func (o *AttributeConfigGet) GetIsReadPublic() bool {
	if o == nil || isNil(o.IsReadPublic) {
		var ret bool
		return ret
	}
	return *o.IsReadPublic
}

// GetIsReadPublicOk returns a tuple with the IsReadPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeConfigGet) GetIsReadPublicOk() (*bool, bool) {
	if o == nil || isNil(o.IsReadPublic) {
    return nil, false
	}
	return o.IsReadPublic, true
}

// HasIsReadPublic returns a boolean if a field has been set.
func (o *AttributeConfigGet) HasIsReadPublic() bool {
	if o != nil && !isNil(o.IsReadPublic) {
		return true
	}

	return false
}

// SetIsReadPublic gets a reference to the given bool and assigns it to the IsReadPublic field.
func (o *AttributeConfigGet) SetIsReadPublic(v bool) {
	o.IsReadPublic = &v
}

// GetOwnerAppId returns the OwnerAppId field value if set, zero value otherwise.
func (o *AttributeConfigGet) GetOwnerAppId() string {
	if o == nil || isNil(o.OwnerAppId) {
		var ret string
		return ret
	}
	return *o.OwnerAppId
}

// GetOwnerAppIdOk returns a tuple with the OwnerAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeConfigGet) GetOwnerAppIdOk() (*string, bool) {
	if o == nil || isNil(o.OwnerAppId) {
    return nil, false
	}
	return o.OwnerAppId, true
}

// HasOwnerAppId returns a boolean if a field has been set.
func (o *AttributeConfigGet) HasOwnerAppId() bool {
	if o != nil && !isNil(o.OwnerAppId) {
		return true
	}

	return false
}

// SetOwnerAppId gets a reference to the given string and assigns it to the OwnerAppId field.
func (o *AttributeConfigGet) SetOwnerAppId(v string) {
	o.OwnerAppId = &v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *AttributeConfigGet) GetDefaultValue() AttributeConfigDefaultValue {
	if o == nil || isNil(o.DefaultValue) {
		var ret AttributeConfigDefaultValue
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeConfigGet) GetDefaultValueOk() (*AttributeConfigDefaultValue, bool) {
	if o == nil || isNil(o.DefaultValue) {
    return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *AttributeConfigGet) HasDefaultValue() bool {
	if o != nil && !isNil(o.DefaultValue) {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given AttributeConfigDefaultValue and assigns it to the DefaultValue field.
func (o *AttributeConfigGet) SetDefaultValue(v AttributeConfigDefaultValue) {
	o.DefaultValue = &v
}

// GetValidation returns the Validation field value if set, zero value otherwise.
func (o *AttributeConfigGet) GetValidation() AttributeValidation {
	if o == nil || isNil(o.Validation) {
		var ret AttributeValidation
		return ret
	}
	return *o.Validation
}

// GetValidationOk returns a tuple with the Validation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeConfigGet) GetValidationOk() (*AttributeValidation, bool) {
	if o == nil || isNil(o.Validation) {
    return nil, false
	}
	return o.Validation, true
}

// HasValidation returns a boolean if a field has been set.
func (o *AttributeConfigGet) HasValidation() bool {
	if o != nil && !isNil(o.Validation) {
		return true
	}

	return false
}

// SetValidation gets a reference to the given AttributeValidation and assigns it to the Validation field.
func (o *AttributeConfigGet) SetValidation(v AttributeValidation) {
	o.Validation = &v
}

func (o AttributeConfigGet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["dataType"] = o.DataType
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Indexed) {
		toSerialize["indexed"] = o.Indexed
	}
	if !isNil(o.IsReadPublic) {
		toSerialize["isReadPublic"] = o.IsReadPublic
	}
	if !isNil(o.OwnerAppId) {
		toSerialize["ownerAppId"] = o.OwnerAppId
	}
	if !isNil(o.DefaultValue) {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	if !isNil(o.Validation) {
		toSerialize["validation"] = o.Validation
	}
	return json.Marshal(toSerialize)
}

type NullableAttributeConfigGet struct {
	value *AttributeConfigGet
	isSet bool
}

func (v NullableAttributeConfigGet) Get() *AttributeConfigGet {
	return v.value
}

func (v *NullableAttributeConfigGet) Set(val *AttributeConfigGet) {
	v.value = val
	v.isSet = true
}

func (v NullableAttributeConfigGet) IsSet() bool {
	return v.isSet
}

func (v *NullableAttributeConfigGet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttributeConfigGet(val *AttributeConfigGet) *NullableAttributeConfigGet {
	return &NullableAttributeConfigGet{value: val, isSet: true}
}

func (v NullableAttributeConfigGet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttributeConfigGet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


