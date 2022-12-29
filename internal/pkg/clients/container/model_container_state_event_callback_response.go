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

// ContainerStateEventCallbackResponse Represents callback url, which can be called to notify the status of API request.
type ContainerStateEventCallbackResponse struct {
	// tenant id.
	TenantId *string `json:"tenantId,omitempty"`
	// container type.
	ContainerType *string `json:"containerType,omitempty"`
	Status *CallbackStatus `json:"status,omitempty"`
	// failure reason in case of execution is failed.
	Reason *string `json:"reason,omitempty"`
	Data *ContainerStateEventCallbackResponseData `json:"data,omitempty"`
}

// NewContainerStateEventCallbackResponse instantiates a new ContainerStateEventCallbackResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerStateEventCallbackResponse() *ContainerStateEventCallbackResponse {
	this := ContainerStateEventCallbackResponse{}
	return &this
}

// NewContainerStateEventCallbackResponseWithDefaults instantiates a new ContainerStateEventCallbackResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerStateEventCallbackResponseWithDefaults() *ContainerStateEventCallbackResponse {
	this := ContainerStateEventCallbackResponse{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ContainerStateEventCallbackResponse) GetTenantId() string {
	if o == nil || o.TenantId == nil {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerStateEventCallbackResponse) GetTenantIdOk() (*string, bool) {
	if o == nil || o.TenantId == nil {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ContainerStateEventCallbackResponse) HasTenantId() bool {
	if o != nil && o.TenantId != nil {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ContainerStateEventCallbackResponse) SetTenantId(v string) {
	o.TenantId = &v
}

// GetContainerType returns the ContainerType field value if set, zero value otherwise.
func (o *ContainerStateEventCallbackResponse) GetContainerType() string {
	if o == nil || o.ContainerType == nil {
		var ret string
		return ret
	}
	return *o.ContainerType
}

// GetContainerTypeOk returns a tuple with the ContainerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerStateEventCallbackResponse) GetContainerTypeOk() (*string, bool) {
	if o == nil || o.ContainerType == nil {
		return nil, false
	}
	return o.ContainerType, true
}

// HasContainerType returns a boolean if a field has been set.
func (o *ContainerStateEventCallbackResponse) HasContainerType() bool {
	if o != nil && o.ContainerType != nil {
		return true
	}

	return false
}

// SetContainerType gets a reference to the given string and assigns it to the ContainerType field.
func (o *ContainerStateEventCallbackResponse) SetContainerType(v string) {
	o.ContainerType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ContainerStateEventCallbackResponse) GetStatus() CallbackStatus {
	if o == nil || o.Status == nil {
		var ret CallbackStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerStateEventCallbackResponse) GetStatusOk() (*CallbackStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ContainerStateEventCallbackResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given CallbackStatus and assigns it to the Status field.
func (o *ContainerStateEventCallbackResponse) SetStatus(v CallbackStatus) {
	o.Status = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ContainerStateEventCallbackResponse) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerStateEventCallbackResponse) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ContainerStateEventCallbackResponse) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *ContainerStateEventCallbackResponse) SetReason(v string) {
	o.Reason = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ContainerStateEventCallbackResponse) GetData() ContainerStateEventCallbackResponseData {
	if o == nil || o.Data == nil {
		var ret ContainerStateEventCallbackResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerStateEventCallbackResponse) GetDataOk() (*ContainerStateEventCallbackResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ContainerStateEventCallbackResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given ContainerStateEventCallbackResponseData and assigns it to the Data field.
func (o *ContainerStateEventCallbackResponse) SetData(v ContainerStateEventCallbackResponseData) {
	o.Data = &v
}

func (o ContainerStateEventCallbackResponse) MarshalJSON() ([]byte, error) {
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

type NullableContainerStateEventCallbackResponse struct {
	value *ContainerStateEventCallbackResponse
	isSet bool
}

func (v NullableContainerStateEventCallbackResponse) Get() *ContainerStateEventCallbackResponse {
	return v.value
}

func (v *NullableContainerStateEventCallbackResponse) Set(val *ContainerStateEventCallbackResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerStateEventCallbackResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerStateEventCallbackResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerStateEventCallbackResponse(val *ContainerStateEventCallbackResponse) *NullableContainerStateEventCallbackResponse {
	return &NullableContainerStateEventCallbackResponse{value: val, isSet: true}
}

func (v NullableContainerStateEventCallbackResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerStateEventCallbackResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


