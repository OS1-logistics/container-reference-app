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

// Cost cost of the item.
type Cost struct {
	Unit *CostUnit `json:"unit,omitempty"`
	Total CostTotal `json:"total"`
}

// NewCost instantiates a new Cost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCost(total CostTotal) *Cost {
	this := Cost{}
	this.Total = total
	return &this
}

// NewCostWithDefaults instantiates a new Cost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCostWithDefaults() *Cost {
	this := Cost{}
	return &this
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *Cost) GetUnit() CostUnit {
	if o == nil || isNil(o.Unit) {
		var ret CostUnit
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cost) GetUnitOk() (*CostUnit, bool) {
	if o == nil || isNil(o.Unit) {
    return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *Cost) HasUnit() bool {
	if o != nil && !isNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given CostUnit and assigns it to the Unit field.
func (o *Cost) SetUnit(v CostUnit) {
	o.Unit = &v
}

// GetTotal returns the Total field value
func (o *Cost) GetTotal() CostTotal {
	if o == nil {
		var ret CostTotal
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *Cost) GetTotalOk() (*CostTotal, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *Cost) SetTotal(v CostTotal) {
	o.Total = v
}

func (o Cost) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	if true {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableCost struct {
	value *Cost
	isSet bool
}

func (v NullableCost) Get() *Cost {
	return v.value
}

func (v *NullableCost) Set(val *Cost) {
	v.value = val
	v.isSet = true
}

func (v NullableCost) IsSet() bool {
	return v.isSet
}

func (v *NullableCost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCost(val *Cost) *NullableCost {
	return &NullableCost{value: val, isSet: true}
}

func (v NullableCost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


