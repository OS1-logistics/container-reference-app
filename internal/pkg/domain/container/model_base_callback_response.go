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

// BaseCallbackResponse struct for BaseCallbackResponse
type BaseCallbackResponse struct {
	// tenant id.
	TenantId *string `json:"tenantId,omitempty"`
	Status *CallbackStatus `json:"status,omitempty"`
	// failure reason in case of execution is failed.
	Reason *string `json:"reason,omitempty"`
	// request data
	Data map[string]interface{} `json:"data,omitempty"`
}

// NewBaseCallbackResponse instantiates a new BaseCallbackResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseCallbackResponse() *BaseCallbackResponse {
	this := BaseCallbackResponse{}
	return &this
}

// NewBaseCallbackResponseWithDefaults instantiates a new BaseCallbackResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseCallbackResponseWithDefaults() *BaseCallbackResponse {
	this := BaseCallbackResponse{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *BaseCallbackResponse) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCallbackResponse) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *BaseCallbackResponse) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *BaseCallbackResponse) SetTenantId(v string) {
	o.TenantId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BaseCallbackResponse) GetStatus() CallbackStatus {
	if o == nil || isNil(o.Status) {
		var ret CallbackStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCallbackResponse) GetStatusOk() (*CallbackStatus, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BaseCallbackResponse) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given CallbackStatus and assigns it to the Status field.
func (o *BaseCallbackResponse) SetStatus(v CallbackStatus) {
	o.Status = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *BaseCallbackResponse) GetReason() string {
	if o == nil || isNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCallbackResponse) GetReasonOk() (*string, bool) {
	if o == nil || isNil(o.Reason) {
    return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *BaseCallbackResponse) HasReason() bool {
	if o != nil && !isNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *BaseCallbackResponse) SetReason(v string) {
	o.Reason = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BaseCallbackResponse) GetData() map[string]interface{} {
	if o == nil || isNil(o.Data) {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCallbackResponse) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Data) {
    return map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BaseCallbackResponse) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *BaseCallbackResponse) SetData(v map[string]interface{}) {
	o.Data = v
}

func (o BaseCallbackResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableBaseCallbackResponse struct {
	value *BaseCallbackResponse
	isSet bool
}

func (v NullableBaseCallbackResponse) Get() *BaseCallbackResponse {
	return v.value
}

func (v *NullableBaseCallbackResponse) Set(val *BaseCallbackResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseCallbackResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseCallbackResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseCallbackResponse(val *BaseCallbackResponse) *NullableBaseCallbackResponse {
	return &NullableBaseCallbackResponse{value: val, isSet: true}
}

func (v NullableBaseCallbackResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseCallbackResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


