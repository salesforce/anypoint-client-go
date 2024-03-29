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

// checks if the UpdateObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateObject{}

// UpdateObject struct for UpdateObject
type UpdateObject struct {
	Op *string `json:"op,omitempty"`
	Path *string `json:"path,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

// NewUpdateObject instantiates a new UpdateObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateObject() *UpdateObject {
	this := UpdateObject{}
	return &this
}

// NewUpdateObjectWithDefaults instantiates a new UpdateObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateObjectWithDefaults() *UpdateObject {
	this := UpdateObject{}
	return &this
}

// GetOp returns the Op field value if set, zero value otherwise.
func (o *UpdateObject) GetOp() string {
	if o == nil || IsNil(o.Op) {
		var ret string
		return ret
	}
	return *o.Op
}

// GetOpOk returns a tuple with the Op field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateObject) GetOpOk() (*string, bool) {
	if o == nil || IsNil(o.Op) {
		return nil, false
	}
	return o.Op, true
}

// HasOp returns a boolean if a field has been set.
func (o *UpdateObject) HasOp() bool {
	if o != nil && !IsNil(o.Op) {
		return true
	}

	return false
}

// SetOp gets a reference to the given string and assigns it to the Op field.
func (o *UpdateObject) SetOp(v string) {
	o.Op = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *UpdateObject) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateObject) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *UpdateObject) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *UpdateObject) SetPath(v string) {
	o.Path = &v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateObject) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateObject) GetValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *UpdateObject) HasValue() bool {
	if o != nil && IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *UpdateObject) SetValue(v interface{}) {
	o.Value = v
}

func (o UpdateObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Op) {
		toSerialize["op"] = o.Op
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableUpdateObject struct {
	value *UpdateObject
	isSet bool
}

func (v NullableUpdateObject) Get() *UpdateObject {
	return v.value
}

func (v *NullableUpdateObject) Set(val *UpdateObject) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateObject) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateObject(val *UpdateObject) *NullableUpdateObject {
	return &NullableUpdateObject{value: val, isSet: true}
}

func (v NullableUpdateObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


