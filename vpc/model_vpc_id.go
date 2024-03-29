/*
VPC API

Description of the VPC API

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vpc

import (
	"encoding/json"
)

// checks if the VpcId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VpcId{}

// VpcId struct for VpcId
type VpcId struct {
	// The vpc Id
	Id string `json:"id"`
}

// NewVpcId instantiates a new VpcId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVpcId(id string) *VpcId {
	this := VpcId{}
	this.Id = id
	return &this
}

// NewVpcIdWithDefaults instantiates a new VpcId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVpcIdWithDefaults() *VpcId {
	this := VpcId{}
	return &this
}

// GetId returns the Id field value
func (o *VpcId) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VpcId) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VpcId) SetId(v string) {
	o.Id = v
}

func (o VpcId) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VpcId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableVpcId struct {
	value *VpcId
	isSet bool
}

func (v NullableVpcId) Get() *VpcId {
	return v.value
}

func (v *NullableVpcId) Set(val *VpcId) {
	v.value = val
	v.isSet = true
}

func (v NullableVpcId) IsSet() bool {
	return v.isSet
}

func (v *NullableVpcId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVpcId(val *VpcId) *NullableVpcId {
	return &NullableVpcId{value: val, isSet: true}
}

func (v NullableVpcId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVpcId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


