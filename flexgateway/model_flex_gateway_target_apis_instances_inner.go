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

// checks if the FlexGatewayTargetApisInstancesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlexGatewayTargetApisInstancesInner{}

// FlexGatewayTargetApisInstancesInner struct for FlexGatewayTargetApisInstancesInner
type FlexGatewayTargetApisInstancesInner struct {
	Id *int32 `json:"id,omitempty"`
	Port *string `json:"port,omitempty"`
	Path *string `json:"path,omitempty"`
}

// NewFlexGatewayTargetApisInstancesInner instantiates a new FlexGatewayTargetApisInstancesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlexGatewayTargetApisInstancesInner() *FlexGatewayTargetApisInstancesInner {
	this := FlexGatewayTargetApisInstancesInner{}
	return &this
}

// NewFlexGatewayTargetApisInstancesInnerWithDefaults instantiates a new FlexGatewayTargetApisInstancesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlexGatewayTargetApisInstancesInnerWithDefaults() *FlexGatewayTargetApisInstancesInner {
	this := FlexGatewayTargetApisInstancesInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FlexGatewayTargetApisInstancesInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexGatewayTargetApisInstancesInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FlexGatewayTargetApisInstancesInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *FlexGatewayTargetApisInstancesInner) SetId(v int32) {
	o.Id = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *FlexGatewayTargetApisInstancesInner) GetPort() string {
	if o == nil || IsNil(o.Port) {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexGatewayTargetApisInstancesInner) GetPortOk() (*string, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *FlexGatewayTargetApisInstancesInner) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *FlexGatewayTargetApisInstancesInner) SetPort(v string) {
	o.Port = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *FlexGatewayTargetApisInstancesInner) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexGatewayTargetApisInstancesInner) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *FlexGatewayTargetApisInstancesInner) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *FlexGatewayTargetApisInstancesInner) SetPath(v string) {
	o.Path = &v
}

func (o FlexGatewayTargetApisInstancesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlexGatewayTargetApisInstancesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	return toSerialize, nil
}

type NullableFlexGatewayTargetApisInstancesInner struct {
	value *FlexGatewayTargetApisInstancesInner
	isSet bool
}

func (v NullableFlexGatewayTargetApisInstancesInner) Get() *FlexGatewayTargetApisInstancesInner {
	return v.value
}

func (v *NullableFlexGatewayTargetApisInstancesInner) Set(val *FlexGatewayTargetApisInstancesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableFlexGatewayTargetApisInstancesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableFlexGatewayTargetApisInstancesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlexGatewayTargetApisInstancesInner(val *FlexGatewayTargetApisInstancesInner) *NullableFlexGatewayTargetApisInstancesInner {
	return &NullableFlexGatewayTargetApisInstancesInner{value: val, isSet: true}
}

func (v NullableFlexGatewayTargetApisInstancesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlexGatewayTargetApisInstancesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


