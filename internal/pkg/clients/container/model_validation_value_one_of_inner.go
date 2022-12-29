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

// ValidationValueOneOfInner struct for ValidationValueOneOfInner
type ValidationValueOneOfInner struct {
	[]interface{} *[]interface{}
	bool *bool
	float32 *float32
	map[string]interface{} *map[string]interface{}
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ValidationValueOneOfInner) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into []interface{}
	err = json.Unmarshal(data, &dst.[]interface{});
	if err == nil {
		json[]interface{}, _ := json.Marshal(dst.[]interface{})
		if string(json[]interface{}) == "{}" { // empty struct
			dst.[]interface{} = nil
		} else {
			return nil // data stored in dst.[]interface{}, return on the first match
		}
	} else {
		dst.[]interface{} = nil
	}

	// try to unmarshal JSON data into bool
	err = json.Unmarshal(data, &dst.bool);
	if err == nil {
		jsonbool, _ := json.Marshal(dst.bool)
		if string(jsonbool) == "{}" { // empty struct
			dst.bool = nil
		} else {
			return nil // data stored in dst.bool, return on the first match
		}
	} else {
		dst.bool = nil
	}

	// try to unmarshal JSON data into float32
	err = json.Unmarshal(data, &dst.float32);
	if err == nil {
		jsonfloat32, _ := json.Marshal(dst.float32)
		if string(jsonfloat32) == "{}" { // empty struct
			dst.float32 = nil
		} else {
			return nil // data stored in dst.float32, return on the first match
		}
	} else {
		dst.float32 = nil
	}

	// try to unmarshal JSON data into map[string]interface{}
	err = json.Unmarshal(data, &dst.map[string]interface{});
	if err == nil {
		jsonmap[string]interface{}, _ := json.Marshal(dst.map[string]interface{})
		if string(jsonmap[string]interface{}) == "{}" { // empty struct
			dst.map[string]interface{} = nil
		} else {
			return nil // data stored in dst.map[string]interface{}, return on the first match
		}
	} else {
		dst.map[string]interface{} = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("Data failed to match schemas in anyOf(ValidationValueOneOfInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ValidationValueOneOfInner) MarshalJSON() ([]byte, error) {
	if src.[]interface{} != nil {
		return json.Marshal(&src.[]interface{})
	}

	if src.bool != nil {
		return json.Marshal(&src.bool)
	}

	if src.float32 != nil {
		return json.Marshal(&src.float32)
	}

	if src.map[string]interface{} != nil {
		return json.Marshal(&src.map[string]interface{})
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableValidationValueOneOfInner struct {
	value *ValidationValueOneOfInner
	isSet bool
}

func (v NullableValidationValueOneOfInner) Get() *ValidationValueOneOfInner {
	return v.value
}

func (v *NullableValidationValueOneOfInner) Set(val *ValidationValueOneOfInner) {
	v.value = val
	v.isSet = true
}

func (v NullableValidationValueOneOfInner) IsSet() bool {
	return v.isSet
}

func (v *NullableValidationValueOneOfInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidationValueOneOfInner(val *ValidationValueOneOfInner) *NullableValidationValueOneOfInner {
	return &NullableValidationValueOneOfInner{value: val, isSet: true}
}

func (v NullableValidationValueOneOfInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidationValueOneOfInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


