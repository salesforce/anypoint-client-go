/*
RoleGroup API

Description of the RoleGroup API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rolegroup

import (
	"encoding/json"
)

// checks if the RolegroupPostBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RolegroupPostBody{}

// RolegroupPostBody struct for RolegroupPostBody
type RolegroupPostBody struct {
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	ExternalNames []string `json:"external_names,omitempty"`
}

// NewRolegroupPostBody instantiates a new RolegroupPostBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRolegroupPostBody() *RolegroupPostBody {
	this := RolegroupPostBody{}
	return &this
}

// NewRolegroupPostBodyWithDefaults instantiates a new RolegroupPostBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRolegroupPostBodyWithDefaults() *RolegroupPostBody {
	this := RolegroupPostBody{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RolegroupPostBody) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolegroupPostBody) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RolegroupPostBody) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RolegroupPostBody) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RolegroupPostBody) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolegroupPostBody) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RolegroupPostBody) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RolegroupPostBody) SetDescription(v string) {
	o.Description = &v
}

// GetExternalNames returns the ExternalNames field value if set, zero value otherwise.
func (o *RolegroupPostBody) GetExternalNames() []string {
	if o == nil || IsNil(o.ExternalNames) {
		var ret []string
		return ret
	}
	return o.ExternalNames
}

// GetExternalNamesOk returns a tuple with the ExternalNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolegroupPostBody) GetExternalNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.ExternalNames) {
		return nil, false
	}
	return o.ExternalNames, true
}

// HasExternalNames returns a boolean if a field has been set.
func (o *RolegroupPostBody) HasExternalNames() bool {
	if o != nil && !IsNil(o.ExternalNames) {
		return true
	}

	return false
}

// SetExternalNames gets a reference to the given []string and assigns it to the ExternalNames field.
func (o *RolegroupPostBody) SetExternalNames(v []string) {
	o.ExternalNames = v
}

func (o RolegroupPostBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RolegroupPostBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ExternalNames) {
		toSerialize["external_names"] = o.ExternalNames
	}
	return toSerialize, nil
}

type NullableRolegroupPostBody struct {
	value *RolegroupPostBody
	isSet bool
}

func (v NullableRolegroupPostBody) Get() *RolegroupPostBody {
	return v.value
}

func (v *NullableRolegroupPostBody) Set(val *RolegroupPostBody) {
	v.value = val
	v.isSet = true
}

func (v NullableRolegroupPostBody) IsSet() bool {
	return v.isSet
}

func (v *NullableRolegroupPostBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRolegroupPostBody(val *RolegroupPostBody) *NullableRolegroupPostBody {
	return &NullableRolegroupPostBody{value: val, isSet: true}
}

func (v NullableRolegroupPostBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRolegroupPostBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


