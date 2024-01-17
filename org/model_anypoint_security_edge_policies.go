/*
Organization API

Description of the Organization API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package org

import (
	"encoding/json"
)

// checks if the AnypointSecurityEdgePolicies type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnypointSecurityEdgePolicies{}

// AnypointSecurityEdgePolicies An explanation about the purpose of this instance.
type AnypointSecurityEdgePolicies struct {
	// An explanation about the purpose of this instance.
	Enabled bool `json:"enabled"`
}

// NewAnypointSecurityEdgePolicies instantiates a new AnypointSecurityEdgePolicies object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnypointSecurityEdgePolicies(enabled bool) *AnypointSecurityEdgePolicies {
	this := AnypointSecurityEdgePolicies{}
	this.Enabled = enabled
	return &this
}

// NewAnypointSecurityEdgePoliciesWithDefaults instantiates a new AnypointSecurityEdgePolicies object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnypointSecurityEdgePoliciesWithDefaults() *AnypointSecurityEdgePolicies {
	this := AnypointSecurityEdgePolicies{}
	var enabled bool = false
	this.Enabled = enabled
	return &this
}

// GetEnabled returns the Enabled field value
func (o *AnypointSecurityEdgePolicies) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AnypointSecurityEdgePolicies) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AnypointSecurityEdgePolicies) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AnypointSecurityEdgePolicies) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnypointSecurityEdgePolicies) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

type NullableAnypointSecurityEdgePolicies struct {
	value *AnypointSecurityEdgePolicies
	isSet bool
}

func (v NullableAnypointSecurityEdgePolicies) Get() *AnypointSecurityEdgePolicies {
	return v.value
}

func (v *NullableAnypointSecurityEdgePolicies) Set(val *AnypointSecurityEdgePolicies) {
	v.value = val
	v.isSet = true
}

func (v NullableAnypointSecurityEdgePolicies) IsSet() bool {
	return v.isSet
}

func (v *NullableAnypointSecurityEdgePolicies) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnypointSecurityEdgePolicies(val *AnypointSecurityEdgePolicies) *NullableAnypointSecurityEdgePolicies {
	return &NullableAnypointSecurityEdgePolicies{value: val, isSet: true}
}

func (v NullableAnypointSecurityEdgePolicies) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnypointSecurityEdgePolicies) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


