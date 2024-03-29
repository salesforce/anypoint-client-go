/*
Flex Gateway API

Description of the Flex Gateway API

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package flexgateway

import (
	"encoding/json"
)

// checks if the FlexGatewayTargetApis type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlexGatewayTargetApis{}

// FlexGatewayTargetApis struct for FlexGatewayTargetApis
type FlexGatewayTargetApis struct {
	Instances []FlexGatewayTargetApisInstancesInner `json:"instances,omitempty"`
	TargetAllowsPortSharing *bool `json:"targetAllowsPortSharing,omitempty"`
}

// NewFlexGatewayTargetApis instantiates a new FlexGatewayTargetApis object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlexGatewayTargetApis() *FlexGatewayTargetApis {
	this := FlexGatewayTargetApis{}
	return &this
}

// NewFlexGatewayTargetApisWithDefaults instantiates a new FlexGatewayTargetApis object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlexGatewayTargetApisWithDefaults() *FlexGatewayTargetApis {
	this := FlexGatewayTargetApis{}
	return &this
}

// GetInstances returns the Instances field value if set, zero value otherwise.
func (o *FlexGatewayTargetApis) GetInstances() []FlexGatewayTargetApisInstancesInner {
	if o == nil || IsNil(o.Instances) {
		var ret []FlexGatewayTargetApisInstancesInner
		return ret
	}
	return o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexGatewayTargetApis) GetInstancesOk() ([]FlexGatewayTargetApisInstancesInner, bool) {
	if o == nil || IsNil(o.Instances) {
		return nil, false
	}
	return o.Instances, true
}

// HasInstances returns a boolean if a field has been set.
func (o *FlexGatewayTargetApis) HasInstances() bool {
	if o != nil && !IsNil(o.Instances) {
		return true
	}

	return false
}

// SetInstances gets a reference to the given []FlexGatewayTargetApisInstancesInner and assigns it to the Instances field.
func (o *FlexGatewayTargetApis) SetInstances(v []FlexGatewayTargetApisInstancesInner) {
	o.Instances = v
}

// GetTargetAllowsPortSharing returns the TargetAllowsPortSharing field value if set, zero value otherwise.
func (o *FlexGatewayTargetApis) GetTargetAllowsPortSharing() bool {
	if o == nil || IsNil(o.TargetAllowsPortSharing) {
		var ret bool
		return ret
	}
	return *o.TargetAllowsPortSharing
}

// GetTargetAllowsPortSharingOk returns a tuple with the TargetAllowsPortSharing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexGatewayTargetApis) GetTargetAllowsPortSharingOk() (*bool, bool) {
	if o == nil || IsNil(o.TargetAllowsPortSharing) {
		return nil, false
	}
	return o.TargetAllowsPortSharing, true
}

// HasTargetAllowsPortSharing returns a boolean if a field has been set.
func (o *FlexGatewayTargetApis) HasTargetAllowsPortSharing() bool {
	if o != nil && !IsNil(o.TargetAllowsPortSharing) {
		return true
	}

	return false
}

// SetTargetAllowsPortSharing gets a reference to the given bool and assigns it to the TargetAllowsPortSharing field.
func (o *FlexGatewayTargetApis) SetTargetAllowsPortSharing(v bool) {
	o.TargetAllowsPortSharing = &v
}

func (o FlexGatewayTargetApis) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlexGatewayTargetApis) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Instances) {
		toSerialize["instances"] = o.Instances
	}
	if !IsNil(o.TargetAllowsPortSharing) {
		toSerialize["targetAllowsPortSharing"] = o.TargetAllowsPortSharing
	}
	return toSerialize, nil
}

type NullableFlexGatewayTargetApis struct {
	value *FlexGatewayTargetApis
	isSet bool
}

func (v NullableFlexGatewayTargetApis) Get() *FlexGatewayTargetApis {
	return v.value
}

func (v *NullableFlexGatewayTargetApis) Set(val *FlexGatewayTargetApis) {
	v.value = val
	v.isSet = true
}

func (v NullableFlexGatewayTargetApis) IsSet() bool {
	return v.isSet
}

func (v *NullableFlexGatewayTargetApis) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlexGatewayTargetApis(val *FlexGatewayTargetApis) *NullableFlexGatewayTargetApis {
	return &NullableFlexGatewayTargetApis{value: val, isSet: true}
}

func (v NullableFlexGatewayTargetApis) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlexGatewayTargetApis) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


