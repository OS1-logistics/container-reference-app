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

// SMSSendRequest struct for SMSSendRequest
type SMSSendRequest struct {
	Recipient string `json:"recipient"`
	Body string `json:"body"`
	Sender string `json:"sender"`
	Req *SMSSendRequestReq `json:"req,omitempty"`
}

// NewSMSSendRequest instantiates a new SMSSendRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSMSSendRequest(recipient string, body string, sender string) *SMSSendRequest {
	this := SMSSendRequest{}
	this.Recipient = recipient
	this.Body = body
	this.Sender = sender
	return &this
}

// NewSMSSendRequestWithDefaults instantiates a new SMSSendRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSMSSendRequestWithDefaults() *SMSSendRequest {
	this := SMSSendRequest{}
	return &this
}

// GetRecipient returns the Recipient field value
func (o *SMSSendRequest) GetRecipient() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value
// and a boolean to check if the value has been set.
func (o *SMSSendRequest) GetRecipientOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Recipient, true
}

// SetRecipient sets field value
func (o *SMSSendRequest) SetRecipient(v string) {
	o.Recipient = v
}

// GetBody returns the Body field value
func (o *SMSSendRequest) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *SMSSendRequest) GetBodyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *SMSSendRequest) SetBody(v string) {
	o.Body = v
}

// GetSender returns the Sender field value
func (o *SMSSendRequest) GetSender() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sender
}

// GetSenderOk returns a tuple with the Sender field value
// and a boolean to check if the value has been set.
func (o *SMSSendRequest) GetSenderOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Sender, true
}

// SetSender sets field value
func (o *SMSSendRequest) SetSender(v string) {
	o.Sender = v
}

// GetReq returns the Req field value if set, zero value otherwise.
func (o *SMSSendRequest) GetReq() SMSSendRequestReq {
	if o == nil || isNil(o.Req) {
		var ret SMSSendRequestReq
		return ret
	}
	return *o.Req
}

// GetReqOk returns a tuple with the Req field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SMSSendRequest) GetReqOk() (*SMSSendRequestReq, bool) {
	if o == nil || isNil(o.Req) {
    return nil, false
	}
	return o.Req, true
}

// HasReq returns a boolean if a field has been set.
func (o *SMSSendRequest) HasReq() bool {
	if o != nil && !isNil(o.Req) {
		return true
	}

	return false
}

// SetReq gets a reference to the given SMSSendRequestReq and assigns it to the Req field.
func (o *SMSSendRequest) SetReq(v SMSSendRequestReq) {
	o.Req = &v
}

func (o SMSSendRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["recipient"] = o.Recipient
	}
	if true {
		toSerialize["body"] = o.Body
	}
	if true {
		toSerialize["sender"] = o.Sender
	}
	if !isNil(o.Req) {
		toSerialize["req"] = o.Req
	}
	return json.Marshal(toSerialize)
}

type NullableSMSSendRequest struct {
	value *SMSSendRequest
	isSet bool
}

func (v NullableSMSSendRequest) Get() *SMSSendRequest {
	return v.value
}

func (v *NullableSMSSendRequest) Set(val *SMSSendRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSMSSendRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSMSSendRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSMSSendRequest(val *SMSSendRequest) *NullableSMSSendRequest {
	return &NullableSMSSendRequest{value: val, isSet: true}
}

func (v NullableSMSSendRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSMSSendRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


