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

// checks if the Spec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Spec{}

// Spec struct for Spec
type Spec struct {
	Name string `json:"name"`
	RemoteAsn *float32 `json:"remoteAsn,omitempty"`
	RemoteIpAddress string `json:"remoteIpAddress"`
	TunnelConfigs []TunnelConfig1 `json:"tunnelConfigs,omitempty"`
	RemoteNetworks []string `json:"remoteNetworks"`
}

// NewSpec instantiates a new Spec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpec(name string, remoteIpAddress string, remoteNetworks []string) *Spec {
	this := Spec{}
	this.Name = name
	this.RemoteIpAddress = remoteIpAddress
	this.RemoteNetworks = remoteNetworks
	return &this
}

// NewSpecWithDefaults instantiates a new Spec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpecWithDefaults() *Spec {
	this := Spec{}
	return &this
}

// GetName returns the Name field value
func (o *Spec) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Spec) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Spec) SetName(v string) {
	o.Name = v
}

// GetRemoteAsn returns the RemoteAsn field value if set, zero value otherwise.
func (o *Spec) GetRemoteAsn() float32 {
	if o == nil || IsNil(o.RemoteAsn) {
		var ret float32
		return ret
	}
	return *o.RemoteAsn
}

// GetRemoteAsnOk returns a tuple with the RemoteAsn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Spec) GetRemoteAsnOk() (*float32, bool) {
	if o == nil || IsNil(o.RemoteAsn) {
		return nil, false
	}
	return o.RemoteAsn, true
}

// HasRemoteAsn returns a boolean if a field has been set.
func (o *Spec) HasRemoteAsn() bool {
	if o != nil && !IsNil(o.RemoteAsn) {
		return true
	}

	return false
}

// SetRemoteAsn gets a reference to the given float32 and assigns it to the RemoteAsn field.
func (o *Spec) SetRemoteAsn(v float32) {
	o.RemoteAsn = &v
}

// GetRemoteIpAddress returns the RemoteIpAddress field value
func (o *Spec) GetRemoteIpAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RemoteIpAddress
}

// GetRemoteIpAddressOk returns a tuple with the RemoteIpAddress field value
// and a boolean to check if the value has been set.
func (o *Spec) GetRemoteIpAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemoteIpAddress, true
}

// SetRemoteIpAddress sets field value
func (o *Spec) SetRemoteIpAddress(v string) {
	o.RemoteIpAddress = v
}

// GetTunnelConfigs returns the TunnelConfigs field value if set, zero value otherwise.
func (o *Spec) GetTunnelConfigs() []TunnelConfig1 {
	if o == nil || IsNil(o.TunnelConfigs) {
		var ret []TunnelConfig1
		return ret
	}
	return o.TunnelConfigs
}

// GetTunnelConfigsOk returns a tuple with the TunnelConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Spec) GetTunnelConfigsOk() ([]TunnelConfig1, bool) {
	if o == nil || IsNil(o.TunnelConfigs) {
		return nil, false
	}
	return o.TunnelConfigs, true
}

// HasTunnelConfigs returns a boolean if a field has been set.
func (o *Spec) HasTunnelConfigs() bool {
	if o != nil && !IsNil(o.TunnelConfigs) {
		return true
	}

	return false
}

// SetTunnelConfigs gets a reference to the given []TunnelConfig1 and assigns it to the TunnelConfigs field.
func (o *Spec) SetTunnelConfigs(v []TunnelConfig1) {
	o.TunnelConfigs = v
}

// GetRemoteNetworks returns the RemoteNetworks field value
func (o *Spec) GetRemoteNetworks() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RemoteNetworks
}

// GetRemoteNetworksOk returns a tuple with the RemoteNetworks field value
// and a boolean to check if the value has been set.
func (o *Spec) GetRemoteNetworksOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemoteNetworks, true
}

// SetRemoteNetworks sets field value
func (o *Spec) SetRemoteNetworks(v []string) {
	o.RemoteNetworks = v
}

func (o Spec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Spec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.RemoteAsn) {
		toSerialize["remoteAsn"] = o.RemoteAsn
	}
	toSerialize["remoteIpAddress"] = o.RemoteIpAddress
	if !IsNil(o.TunnelConfigs) {
		toSerialize["tunnelConfigs"] = o.TunnelConfigs
	}
	toSerialize["remoteNetworks"] = o.RemoteNetworks
	return toSerialize, nil
}

type NullableSpec struct {
	value *Spec
	isSet bool
}

func (v NullableSpec) Get() *Spec {
	return v.value
}

func (v *NullableSpec) Set(val *Spec) {
	v.value = val
	v.isSet = true
}

func (v NullableSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpec(val *Spec) *NullableSpec {
	return &NullableSpec{value: val, isSet: true}
}

func (v NullableSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


