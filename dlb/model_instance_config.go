/*
Dedicated Load Balancer API

Description of the DLB API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dlb

import (
	"encoding/json"
)

// checks if the InstanceConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceConfig{}

// InstanceConfig struct for InstanceConfig
type InstanceConfig struct {
	ImageName *string `json:"imageName,omitempty"`
}

// NewInstanceConfig instantiates a new InstanceConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceConfig() *InstanceConfig {
	this := InstanceConfig{}
	return &this
}

// NewInstanceConfigWithDefaults instantiates a new InstanceConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceConfigWithDefaults() *InstanceConfig {
	this := InstanceConfig{}
	return &this
}

// GetImageName returns the ImageName field value if set, zero value otherwise.
func (o *InstanceConfig) GetImageName() string {
	if o == nil || IsNil(o.ImageName) {
		var ret string
		return ret
	}
	return *o.ImageName
}

// GetImageNameOk returns a tuple with the ImageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceConfig) GetImageNameOk() (*string, bool) {
	if o == nil || IsNil(o.ImageName) {
		return nil, false
	}
	return o.ImageName, true
}

// HasImageName returns a boolean if a field has been set.
func (o *InstanceConfig) HasImageName() bool {
	if o != nil && !IsNil(o.ImageName) {
		return true
	}

	return false
}

// SetImageName gets a reference to the given string and assigns it to the ImageName field.
func (o *InstanceConfig) SetImageName(v string) {
	o.ImageName = &v
}

func (o InstanceConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstanceConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ImageName) {
		toSerialize["imageName"] = o.ImageName
	}
	return toSerialize, nil
}

type NullableInstanceConfig struct {
	value *InstanceConfig
	isSet bool
}

func (v NullableInstanceConfig) Get() *InstanceConfig {
	return v.value
}

func (v *NullableInstanceConfig) Set(val *InstanceConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceConfig(val *InstanceConfig) *NullableInstanceConfig {
	return &NullableInstanceConfig{value: val, isSet: true}
}

func (v NullableInstanceConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


