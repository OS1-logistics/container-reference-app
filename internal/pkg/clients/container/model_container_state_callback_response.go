/*
Container Service

**API documentation for Container Service**

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package container_client

import (
	"encoding/json"
)

// ContainerStateCallbackResponse Represents callback url, which can be called to notify the status of API request.
type ContainerStateCallbackResponse struct {
	// tenant id.
	TenantId *string `json:"tenantId,omitempty"`
	// container type.
	ContainerType *string `json:"containerType,omitempty"`
	Status *CallbackStatus `json:"status,omitempty"`
	// failure reason in case of execution is failed.
	Reason *string `json:"reason,omitempty"`
	Data *ContainerStateCallbackResponseData `json:"data,omitempty"`
}

// NewContainerStateCallbackResponse instantiates a new ContainerStateCallbackResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerStateCallbackResponse() *ContainerStateCallbackResponse {
	this := ContainerStateCallbackResponse{}
	return &this
}

// NewContainerStateCallbackResponseWithDefaults instantiates a new ContainerStateCallbackResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerStateCallbackResponseWithDefaults() *ContainerStateCallbackResponse {
	this := ContainerStateCallbackResponse{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ContainerStateCallbackResponse) GetTenantId() string {
	if o == nil || o.TenantId == nil {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerStateCallbackResponse) GetTenantIdOk() (*string, bool) {
	if o == nil || o.TenantId == nil {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ContainerStateCallbackResponse) HasTenantId() bool {
	if o != nil && o.TenantId != nil {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ContainerStateCallbackResponse) SetTenantId(v string) {
	o.TenantId = &v
}

// GetContainerType returns the ContainerType field value if set, zero value otherwise.
func (o *ContainerStateCallbackResponse) GetContainerType() string {
	if o == nil || o.ContainerType == nil {
		var ret string
		return ret
	}
	return *o.ContainerType
}

// GetContainerTypeOk returns a tuple with the ContainerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerStateCallbackResponse) GetContainerTypeOk() (*string, bool) {
	if o == nil || o.ContainerType == nil {
		return nil, false
	}
	return o.ContainerType, true
}

// HasContainerType returns a boolean if a field has been set.
func (o *ContainerStateCallbackResponse) HasContainerType() bool {
	if o != nil && o.ContainerType != nil {
		return true
	}

	return false
}

// SetContainerType gets a reference to the given string and assigns it to the ContainerType field.
func (o *ContainerStateCallbackResponse) SetContainerType(v string) {
	o.ContainerType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ContainerStateCallbackResponse) GetStatus() CallbackStatus {
	if o == nil || o.Status == nil {
		var ret CallbackStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerStateCallbackResponse) GetStatusOk() (*CallbackStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ContainerStateCallbackResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given CallbackStatus and assigns it to the Status field.
func (o *ContainerStateCallbackResponse) SetStatus(v CallbackStatus) {
	o.Status = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ContainerStateCallbackResponse) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerStateCallbackResponse) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ContainerStateCallbackResponse) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *ContainerStateCallbackResponse) SetReason(v string) {
	o.Reason = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ContainerStateCallbackResponse) GetData() ContainerStateCallbackResponseData {
	if o == nil || o.Data == nil {
		var ret ContainerStateCallbackResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerStateCallbackResponse) GetDataOk() (*ContainerStateCallbackResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ContainerStateCallbackResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given ContainerStateCallbackResponseData and assigns it to the Data field.
func (o *ContainerStateCallbackResponse) SetData(v ContainerStateCallbackResponseData) {
	o.Data = &v
}

func (o ContainerStateCallbackResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TenantId != nil {
		toSerialize["tenantId"] = o.TenantId
	}
	if o.ContainerType != nil {
		toSerialize["containerType"] = o.ContainerType
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableContainerStateCallbackResponse struct {
	value *ContainerStateCallbackResponse
	isSet bool
}

func (v NullableContainerStateCallbackResponse) Get() *ContainerStateCallbackResponse {
	return v.value
}

func (v *NullableContainerStateCallbackResponse) Set(val *ContainerStateCallbackResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerStateCallbackResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerStateCallbackResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerStateCallbackResponse(val *ContainerStateCallbackResponse) *NullableContainerStateCallbackResponse {
	return &NullableContainerStateCallbackResponse{value: val, isSet: true}
}

func (v NullableContainerStateCallbackResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerStateCallbackResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


