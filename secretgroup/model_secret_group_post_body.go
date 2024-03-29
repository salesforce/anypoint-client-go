/*
Secret Group API

Secret Group API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package secretgroup

import (
	"encoding/json"
)

// checks if the SecretGroupPostBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecretGroupPostBody{}

// SecretGroupPostBody struct for SecretGroupPostBody
type SecretGroupPostBody struct {
	Name *string `json:"name,omitempty"`
	Downloadable *bool `json:"downloadable,omitempty"`
}

// NewSecretGroupPostBody instantiates a new SecretGroupPostBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretGroupPostBody() *SecretGroupPostBody {
	this := SecretGroupPostBody{}
	return &this
}

// NewSecretGroupPostBodyWithDefaults instantiates a new SecretGroupPostBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretGroupPostBodyWithDefaults() *SecretGroupPostBody {
	this := SecretGroupPostBody{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SecretGroupPostBody) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretGroupPostBody) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SecretGroupPostBody) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SecretGroupPostBody) SetName(v string) {
	o.Name = &v
}

// GetDownloadable returns the Downloadable field value if set, zero value otherwise.
func (o *SecretGroupPostBody) GetDownloadable() bool {
	if o == nil || IsNil(o.Downloadable) {
		var ret bool
		return ret
	}
	return *o.Downloadable
}

// GetDownloadableOk returns a tuple with the Downloadable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretGroupPostBody) GetDownloadableOk() (*bool, bool) {
	if o == nil || IsNil(o.Downloadable) {
		return nil, false
	}
	return o.Downloadable, true
}

// HasDownloadable returns a boolean if a field has been set.
func (o *SecretGroupPostBody) HasDownloadable() bool {
	if o != nil && !IsNil(o.Downloadable) {
		return true
	}

	return false
}

// SetDownloadable gets a reference to the given bool and assigns it to the Downloadable field.
func (o *SecretGroupPostBody) SetDownloadable(v bool) {
	o.Downloadable = &v
}

func (o SecretGroupPostBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecretGroupPostBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Downloadable) {
		toSerialize["downloadable"] = o.Downloadable
	}
	return toSerialize, nil
}

type NullableSecretGroupPostBody struct {
	value *SecretGroupPostBody
	isSet bool
}

func (v NullableSecretGroupPostBody) Get() *SecretGroupPostBody {
	return v.value
}

func (v *NullableSecretGroupPostBody) Set(val *SecretGroupPostBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretGroupPostBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretGroupPostBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretGroupPostBody(val *SecretGroupPostBody) *NullableSecretGroupPostBody {
	return &NullableSecretGroupPostBody{value: val, isSet: true}
}

func (v NullableSecretGroupPostBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretGroupPostBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


