/*
VPN API

Description of the VPN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vpn

import (
	"encoding/json"
)

// checks if the State1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &State1{}

// State1 struct for State1
type State1 struct {
	VpnConnectionStatus string `json:"vpnConnectionStatus"`
}

// NewState1 instantiates a new State1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewState1(vpnConnectionStatus string) *State1 {
	this := State1{}
	this.VpnConnectionStatus = vpnConnectionStatus
	return &this
}

// NewState1WithDefaults instantiates a new State1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewState1WithDefaults() *State1 {
	this := State1{}
	return &this
}

// GetVpnConnectionStatus returns the VpnConnectionStatus field value
func (o *State1) GetVpnConnectionStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VpnConnectionStatus
}

// GetVpnConnectionStatusOk returns a tuple with the VpnConnectionStatus field value
// and a boolean to check if the value has been set.
func (o *State1) GetVpnConnectionStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VpnConnectionStatus, true
}

// SetVpnConnectionStatus sets field value
func (o *State1) SetVpnConnectionStatus(v string) {
	o.VpnConnectionStatus = v
}

func (o State1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o State1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["vpnConnectionStatus"] = o.VpnConnectionStatus
	return toSerialize, nil
}

type NullableState1 struct {
	value *State1
	isSet bool
}

func (v NullableState1) Get() *State1 {
	return v.value
}

func (v *NullableState1) Set(val *State1) {
	v.value = val
	v.isSet = true
}

func (v NullableState1) IsSet() bool {
	return v.isSet
}

func (v *NullableState1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableState1(val *State1) *NullableState1 {
	return &NullableState1{value: val, isSet: true}
}

func (v NullableState1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableState1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


