/*
Identity Provider Management API

Description of Identity Provider API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package idp

import (
	"encoding/json"
)

// checks if the Urls1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Urls1{}

// Urls1 struct for Urls1
type Urls1 struct {
	Register *string `json:"register,omitempty"`
}

// NewUrls1 instantiates a new Urls1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUrls1() *Urls1 {
	this := Urls1{}
	return &this
}

// NewUrls1WithDefaults instantiates a new Urls1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUrls1WithDefaults() *Urls1 {
	this := Urls1{}
	return &this
}

// GetRegister returns the Register field value if set, zero value otherwise.
func (o *Urls1) GetRegister() string {
	if o == nil || IsNil(o.Register) {
		var ret string
		return ret
	}
	return *o.Register
}

// GetRegisterOk returns a tuple with the Register field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Urls1) GetRegisterOk() (*string, bool) {
	if o == nil || IsNil(o.Register) {
		return nil, false
	}
	return o.Register, true
}

// HasRegister returns a boolean if a field has been set.
func (o *Urls1) HasRegister() bool {
	if o != nil && !IsNil(o.Register) {
		return true
	}

	return false
}

// SetRegister gets a reference to the given string and assigns it to the Register field.
func (o *Urls1) SetRegister(v string) {
	o.Register = &v
}

func (o Urls1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Urls1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Register) {
		toSerialize["register"] = o.Register
	}
	return toSerialize, nil
}

type NullableUrls1 struct {
	value *Urls1
	isSet bool
}

func (v NullableUrls1) Get() *Urls1 {
	return v.value
}

func (v *NullableUrls1) Set(val *Urls1) {
	v.value = val
	v.isSet = true
}

func (v NullableUrls1) IsSet() bool {
	return v.isSet
}

func (v *NullableUrls1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUrls1(val *Urls1) *NullableUrls1 {
	return &NullableUrls1{value: val, isSet: true}
}

func (v NullableUrls1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUrls1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


