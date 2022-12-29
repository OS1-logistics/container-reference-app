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

// ContainerTypeFilterQuery struct for ContainerTypeFilterQuery
type ContainerTypeFilterQuery struct {
	// The match query is the standard query for performing a full-text search.
	Match []Match `json:"match,omitempty"`
	// The Term query returns fields that contain an exact term in a provided field.
	Term []Term `json:"term,omitempty"`
	// The Terms query returns field that contain one or more exact terms in a provided field. The terms query is the same as the term query, except you can search for multiple values.
	Terms []Terms `json:"terms,omitempty"`
	// The Range query returns fields that contain terms within a provided range.
	Range []RangeQuery `json:"range,omitempty"`
	// The Regexp query returns fields that contain terms matching a regular expression.
	Regexp []Regex `json:"regexp,omitempty"`
	Exists []Exists `json:"exists,omitempty"`
	Prefix []Prefix `json:"prefix,omitempty"`
}

// NewContainerTypeFilterQuery instantiates a new ContainerTypeFilterQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerTypeFilterQuery() *ContainerTypeFilterQuery {
	this := ContainerTypeFilterQuery{}
	return &this
}

// NewContainerTypeFilterQueryWithDefaults instantiates a new ContainerTypeFilterQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerTypeFilterQueryWithDefaults() *ContainerTypeFilterQuery {
	this := ContainerTypeFilterQuery{}
	return &this
}

// GetMatch returns the Match field value if set, zero value otherwise.
func (o *ContainerTypeFilterQuery) GetMatch() []Match {
	if o == nil || o.Match == nil {
		var ret []Match
		return ret
	}
	return o.Match
}

// GetMatchOk returns a tuple with the Match field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerTypeFilterQuery) GetMatchOk() ([]Match, bool) {
	if o == nil || o.Match == nil {
		return nil, false
	}
	return o.Match, true
}

// HasMatch returns a boolean if a field has been set.
func (o *ContainerTypeFilterQuery) HasMatch() bool {
	if o != nil && o.Match != nil {
		return true
	}

	return false
}

// SetMatch gets a reference to the given []Match and assigns it to the Match field.
func (o *ContainerTypeFilterQuery) SetMatch(v []Match) {
	o.Match = v
}

// GetTerm returns the Term field value if set, zero value otherwise.
func (o *ContainerTypeFilterQuery) GetTerm() []Term {
	if o == nil || o.Term == nil {
		var ret []Term
		return ret
	}
	return o.Term
}

// GetTermOk returns a tuple with the Term field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerTypeFilterQuery) GetTermOk() ([]Term, bool) {
	if o == nil || o.Term == nil {
		return nil, false
	}
	return o.Term, true
}

// HasTerm returns a boolean if a field has been set.
func (o *ContainerTypeFilterQuery) HasTerm() bool {
	if o != nil && o.Term != nil {
		return true
	}

	return false
}

// SetTerm gets a reference to the given []Term and assigns it to the Term field.
func (o *ContainerTypeFilterQuery) SetTerm(v []Term) {
	o.Term = v
}

// GetTerms returns the Terms field value if set, zero value otherwise.
func (o *ContainerTypeFilterQuery) GetTerms() []Terms {
	if o == nil || o.Terms == nil {
		var ret []Terms
		return ret
	}
	return o.Terms
}

// GetTermsOk returns a tuple with the Terms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerTypeFilterQuery) GetTermsOk() ([]Terms, bool) {
	if o == nil || o.Terms == nil {
		return nil, false
	}
	return o.Terms, true
}

// HasTerms returns a boolean if a field has been set.
func (o *ContainerTypeFilterQuery) HasTerms() bool {
	if o != nil && o.Terms != nil {
		return true
	}

	return false
}

// SetTerms gets a reference to the given []Terms and assigns it to the Terms field.
func (o *ContainerTypeFilterQuery) SetTerms(v []Terms) {
	o.Terms = v
}

// GetRange returns the Range field value if set, zero value otherwise.
func (o *ContainerTypeFilterQuery) GetRange() []RangeQuery {
	if o == nil || o.Range == nil {
		var ret []RangeQuery
		return ret
	}
	return o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerTypeFilterQuery) GetRangeOk() ([]RangeQuery, bool) {
	if o == nil || o.Range == nil {
		return nil, false
	}
	return o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *ContainerTypeFilterQuery) HasRange() bool {
	if o != nil && o.Range != nil {
		return true
	}

	return false
}

// SetRange gets a reference to the given []RangeQuery and assigns it to the Range field.
func (o *ContainerTypeFilterQuery) SetRange(v []RangeQuery) {
	o.Range = v
}

// GetRegexp returns the Regexp field value if set, zero value otherwise.
func (o *ContainerTypeFilterQuery) GetRegexp() []Regex {
	if o == nil || o.Regexp == nil {
		var ret []Regex
		return ret
	}
	return o.Regexp
}

// GetRegexpOk returns a tuple with the Regexp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerTypeFilterQuery) GetRegexpOk() ([]Regex, bool) {
	if o == nil || o.Regexp == nil {
		return nil, false
	}
	return o.Regexp, true
}

// HasRegexp returns a boolean if a field has been set.
func (o *ContainerTypeFilterQuery) HasRegexp() bool {
	if o != nil && o.Regexp != nil {
		return true
	}

	return false
}

// SetRegexp gets a reference to the given []Regex and assigns it to the Regexp field.
func (o *ContainerTypeFilterQuery) SetRegexp(v []Regex) {
	o.Regexp = v
}

// GetExists returns the Exists field value if set, zero value otherwise.
func (o *ContainerTypeFilterQuery) GetExists() []Exists {
	if o == nil || o.Exists == nil {
		var ret []Exists
		return ret
	}
	return o.Exists
}

// GetExistsOk returns a tuple with the Exists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerTypeFilterQuery) GetExistsOk() ([]Exists, bool) {
	if o == nil || o.Exists == nil {
		return nil, false
	}
	return o.Exists, true
}

// HasExists returns a boolean if a field has been set.
func (o *ContainerTypeFilterQuery) HasExists() bool {
	if o != nil && o.Exists != nil {
		return true
	}

	return false
}

// SetExists gets a reference to the given []Exists and assigns it to the Exists field.
func (o *ContainerTypeFilterQuery) SetExists(v []Exists) {
	o.Exists = v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *ContainerTypeFilterQuery) GetPrefix() []Prefix {
	if o == nil || o.Prefix == nil {
		var ret []Prefix
		return ret
	}
	return o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerTypeFilterQuery) GetPrefixOk() ([]Prefix, bool) {
	if o == nil || o.Prefix == nil {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *ContainerTypeFilterQuery) HasPrefix() bool {
	if o != nil && o.Prefix != nil {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given []Prefix and assigns it to the Prefix field.
func (o *ContainerTypeFilterQuery) SetPrefix(v []Prefix) {
	o.Prefix = v
}

func (o ContainerTypeFilterQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Match != nil {
		toSerialize["match"] = o.Match
	}
	if o.Term != nil {
		toSerialize["term"] = o.Term
	}
	if o.Terms != nil {
		toSerialize["terms"] = o.Terms
	}
	if o.Range != nil {
		toSerialize["range"] = o.Range
	}
	if o.Regexp != nil {
		toSerialize["regexp"] = o.Regexp
	}
	if o.Exists != nil {
		toSerialize["exists"] = o.Exists
	}
	if o.Prefix != nil {
		toSerialize["prefix"] = o.Prefix
	}
	return json.Marshal(toSerialize)
}

type NullableContainerTypeFilterQuery struct {
	value *ContainerTypeFilterQuery
	isSet bool
}

func (v NullableContainerTypeFilterQuery) Get() *ContainerTypeFilterQuery {
	return v.value
}

func (v *NullableContainerTypeFilterQuery) Set(val *ContainerTypeFilterQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerTypeFilterQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerTypeFilterQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerTypeFilterQuery(val *ContainerTypeFilterQuery) *NullableContainerTypeFilterQuery {
	return &NullableContainerTypeFilterQuery{value: val, isSet: true}
}

func (v NullableContainerTypeFilterQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerTypeFilterQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


