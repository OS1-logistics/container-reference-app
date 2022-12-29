/*
Container Service

**API documentation for Container Service**

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package container_client

import (
	"encoding/json"
	"fmt"
)

// EventDataType the model 'EventDataType'
type EventDataType string

// List of EventDataType
const (
	EVENTDATATYPE_STRING EventDataType = "string"
	EVENTDATATYPE_NUMBER EventDataType = "number"
	EVENTDATATYPE_BOOLEAN EventDataType = "boolean"
	EVENTDATATYPE_OBJECT EventDataType = "object"
	EVENTDATATYPE_ARRAY EventDataType = "array"
)

// All allowed values of EventDataType enum
var AllowedEventDataTypeEnumValues = []EventDataType{
	"string",
	"number",
	"boolean",
	"object",
	"array",
}

func (v *EventDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EventDataType(value)
	for _, existing := range AllowedEventDataTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EventDataType", value)
}

// NewEventDataTypeFromValue returns a pointer to a valid EventDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventDataTypeFromValue(v string) (*EventDataType, error) {
	ev := EventDataType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventDataType: valid values are %v", v, AllowedEventDataTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventDataType) IsValid() bool {
	for _, existing := range AllowedEventDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventDataType value
func (v EventDataType) Ptr() *EventDataType {
	return &v
}

type NullableEventDataType struct {
	value *EventDataType
	isSet bool
}

func (v NullableEventDataType) Get() *EventDataType {
	return v.value
}

func (v *NullableEventDataType) Set(val *EventDataType) {
	v.value = val
	v.isSet = true
}

func (v NullableEventDataType) IsSet() bool {
	return v.isSet
}

func (v *NullableEventDataType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventDataType(val *EventDataType) *NullableEventDataType {
	return &NullableEventDataType{value: val, isSet: true}
}

func (v NullableEventDataType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventDataType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

