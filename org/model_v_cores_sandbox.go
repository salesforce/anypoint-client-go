/*
 * Organization API
 *
 * Description of the Organization API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package org

import (
	"encoding/json"
)

// VCoresSandbox An explanation about the purpose of this instance.
type VCoresSandbox struct {
	// An explanation about the purpose of this instance.
	Assigned float32 `json:"assigned"`
	// An explanation about the purpose of this instance.
	Reassigned float32 `json:"reassigned"`
}

// NewVCoresSandbox instantiates a new VCoresSandbox object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVCoresSandbox(assigned float32, reassigned float32) *VCoresSandbox {
	this := VCoresSandbox{}
	this.Assigned = assigned
	this.Reassigned = reassigned
	return &this
}

// NewVCoresSandboxWithDefaults instantiates a new VCoresSandbox object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVCoresSandboxWithDefaults() *VCoresSandbox {
	this := VCoresSandbox{}
	var assigned float32 = 0
	this.Assigned = assigned
	var reassigned float32 = 0.0
	this.Reassigned = reassigned
	return &this
}

// GetAssigned returns the Assigned field value
func (o *VCoresSandbox) GetAssigned() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Assigned
}

// GetAssignedOk returns a tuple with the Assigned field value
// and a boolean to check if the value has been set.
func (o *VCoresSandbox) GetAssignedOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Assigned, true
}

// SetAssigned sets field value
func (o *VCoresSandbox) SetAssigned(v float32) {
	o.Assigned = v
}

// GetReassigned returns the Reassigned field value
func (o *VCoresSandbox) GetReassigned() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Reassigned
}

// GetReassignedOk returns a tuple with the Reassigned field value
// and a boolean to check if the value has been set.
func (o *VCoresSandbox) GetReassignedOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Reassigned, true
}

// SetReassigned sets field value
func (o *VCoresSandbox) SetReassigned(v float32) {
	o.Reassigned = v
}

func (o VCoresSandbox) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["assigned"] = o.Assigned
	}
	if true {
		toSerialize["reassigned"] = o.Reassigned
	}
	return json.Marshal(toSerialize)
}

type NullableVCoresSandbox struct {
	value *VCoresSandbox
	isSet bool
}

func (v NullableVCoresSandbox) Get() *VCoresSandbox {
	return v.value
}

func (v *NullableVCoresSandbox) Set(val *VCoresSandbox) {
	v.value = val
	v.isSet = true
}

func (v NullableVCoresSandbox) IsSet() bool {
	return v.isSet
}

func (v *NullableVCoresSandbox) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVCoresSandbox(val *VCoresSandbox) *NullableVCoresSandbox {
	return &NullableVCoresSandbox{value: val, isSet: true}
}

func (v NullableVCoresSandbox) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVCoresSandbox) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


