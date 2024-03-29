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

// checks if the MasterBGSpecificDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MasterBGSpecificDetails{}

// MasterBGSpecificDetails struct for MasterBGSpecificDetails
type MasterBGSpecificDetails struct {
	// An explanation about the purpose of this instance.
	SessionTimeout *int32 `json:"sessionTimeout,omitempty"`
	Subscription *Subscription `json:"subscription,omitempty"`
}

// NewMasterBGSpecificDetails instantiates a new MasterBGSpecificDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMasterBGSpecificDetails() *MasterBGSpecificDetails {
	this := MasterBGSpecificDetails{}
	var sessionTimeout int32 = 0
	this.SessionTimeout = &sessionTimeout
	return &this
}

// NewMasterBGSpecificDetailsWithDefaults instantiates a new MasterBGSpecificDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMasterBGSpecificDetailsWithDefaults() *MasterBGSpecificDetails {
	this := MasterBGSpecificDetails{}
	var sessionTimeout int32 = 0
	this.SessionTimeout = &sessionTimeout
	return &this
}

// GetSessionTimeout returns the SessionTimeout field value if set, zero value otherwise.
func (o *MasterBGSpecificDetails) GetSessionTimeout() int32 {
	if o == nil || IsNil(o.SessionTimeout) {
		var ret int32
		return ret
	}
	return *o.SessionTimeout
}

// GetSessionTimeoutOk returns a tuple with the SessionTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGSpecificDetails) GetSessionTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.SessionTimeout) {
		return nil, false
	}
	return o.SessionTimeout, true
}

// HasSessionTimeout returns a boolean if a field has been set.
func (o *MasterBGSpecificDetails) HasSessionTimeout() bool {
	if o != nil && !IsNil(o.SessionTimeout) {
		return true
	}

	return false
}

// SetSessionTimeout gets a reference to the given int32 and assigns it to the SessionTimeout field.
func (o *MasterBGSpecificDetails) SetSessionTimeout(v int32) {
	o.SessionTimeout = &v
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *MasterBGSpecificDetails) GetSubscription() Subscription {
	if o == nil || IsNil(o.Subscription) {
		var ret Subscription
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGSpecificDetails) GetSubscriptionOk() (*Subscription, bool) {
	if o == nil || IsNil(o.Subscription) {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *MasterBGSpecificDetails) HasSubscription() bool {
	if o != nil && !IsNil(o.Subscription) {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given Subscription and assigns it to the Subscription field.
func (o *MasterBGSpecificDetails) SetSubscription(v Subscription) {
	o.Subscription = &v
}

func (o MasterBGSpecificDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MasterBGSpecificDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SessionTimeout) {
		toSerialize["sessionTimeout"] = o.SessionTimeout
	}
	if !IsNil(o.Subscription) {
		toSerialize["subscription"] = o.Subscription
	}
	return toSerialize, nil
}

type NullableMasterBGSpecificDetails struct {
	value *MasterBGSpecificDetails
	isSet bool
}

func (v NullableMasterBGSpecificDetails) Get() *MasterBGSpecificDetails {
	return v.value
}

func (v *NullableMasterBGSpecificDetails) Set(val *MasterBGSpecificDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableMasterBGSpecificDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableMasterBGSpecificDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMasterBGSpecificDetails(val *MasterBGSpecificDetails) *NullableMasterBGSpecificDetails {
	return &NullableMasterBGSpecificDetails{value: val, isSet: true}
}

func (v NullableMasterBGSpecificDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMasterBGSpecificDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


