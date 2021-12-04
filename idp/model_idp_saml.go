/*
 * Identity Provider Management API
 *
 * Description of Identity Provider API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package idp

import (
	"encoding/json"
)

// IdpSAML struct for IdpSAML
type IdpSAML struct {
	Saml *Saml `json:"saml,omitempty"`
}

// NewIdpSAML instantiates a new IdpSAML object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdpSAML() *IdpSAML {
	this := IdpSAML{}
	return &this
}

// NewIdpSAMLWithDefaults instantiates a new IdpSAML object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdpSAMLWithDefaults() *IdpSAML {
	this := IdpSAML{}
	return &this
}

// GetSaml returns the Saml field value if set, zero value otherwise.
func (o *IdpSAML) GetSaml() Saml {
	if o == nil || o.Saml == nil {
		var ret Saml
		return ret
	}
	return *o.Saml
}

// GetSamlOk returns a tuple with the Saml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpSAML) GetSamlOk() (*Saml, bool) {
	if o == nil || o.Saml == nil {
		return nil, false
	}
	return o.Saml, true
}

// HasSaml returns a boolean if a field has been set.
func (o *IdpSAML) HasSaml() bool {
	if o != nil && o.Saml != nil {
		return true
	}

	return false
}

// SetSaml gets a reference to the given Saml and assigns it to the Saml field.
func (o *IdpSAML) SetSaml(v Saml) {
	o.Saml = &v
}

func (o IdpSAML) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Saml != nil {
		toSerialize["saml"] = o.Saml
	}
	return json.Marshal(toSerialize)
}

type NullableIdpSAML struct {
	value *IdpSAML
	isSet bool
}

func (v NullableIdpSAML) Get() *IdpSAML {
	return v.value
}

func (v *NullableIdpSAML) Set(val *IdpSAML) {
	v.value = val
	v.isSet = true
}

func (v NullableIdpSAML) IsSet() bool {
	return v.isSet
}

func (v *NullableIdpSAML) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdpSAML(val *IdpSAML) *NullableIdpSAML {
	return &NullableIdpSAML{value: val, isSet: true}
}

func (v NullableIdpSAML) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdpSAML) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


