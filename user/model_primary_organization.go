/*
 * User API
 *
 * Description of the User API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package user

import (
	"encoding/json"
)

// PrimaryOrganization struct for PrimaryOrganization
type PrimaryOrganization struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewPrimaryOrganization instantiates a new PrimaryOrganization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrimaryOrganization() *PrimaryOrganization {
	this := PrimaryOrganization{}
	return &this
}

// NewPrimaryOrganizationWithDefaults instantiates a new PrimaryOrganization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrimaryOrganizationWithDefaults() *PrimaryOrganization {
	this := PrimaryOrganization{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PrimaryOrganization) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrimaryOrganization) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PrimaryOrganization) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PrimaryOrganization) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PrimaryOrganization) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrimaryOrganization) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PrimaryOrganization) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PrimaryOrganization) SetName(v string) {
	o.Name = &v
}

func (o PrimaryOrganization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullablePrimaryOrganization struct {
	value *PrimaryOrganization
	isSet bool
}

func (v NullablePrimaryOrganization) Get() *PrimaryOrganization {
	return v.value
}

func (v *NullablePrimaryOrganization) Set(val *PrimaryOrganization) {
	v.value = val
	v.isSet = true
}

func (v NullablePrimaryOrganization) IsSet() bool {
	return v.isSet
}

func (v *NullablePrimaryOrganization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrimaryOrganization(val *PrimaryOrganization) *NullablePrimaryOrganization {
	return &NullablePrimaryOrganization{value: val, isSet: true}
}

func (v NullablePrimaryOrganization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrimaryOrganization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


