/*
Secret Group Truststore API

Secret Group Truststore API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package secretgroup_truststore

import (
	"encoding/json"
)

// checks if the CertificateEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertificateEntry{}

// CertificateEntry struct for CertificateEntry
type CertificateEntry struct {
	// Alias associated with the certificate entry
	Alias *string `json:"alias,omitempty"`
	Certificate *CertificateDetails `json:"certificate,omitempty"`
}

// NewCertificateEntry instantiates a new CertificateEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateEntry() *CertificateEntry {
	this := CertificateEntry{}
	return &this
}

// NewCertificateEntryWithDefaults instantiates a new CertificateEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateEntryWithDefaults() *CertificateEntry {
	this := CertificateEntry{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *CertificateEntry) GetAlias() string {
	if o == nil || IsNil(o.Alias) {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateEntry) GetAliasOk() (*string, bool) {
	if o == nil || IsNil(o.Alias) {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *CertificateEntry) HasAlias() bool {
	if o != nil && !IsNil(o.Alias) {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *CertificateEntry) SetAlias(v string) {
	o.Alias = &v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *CertificateEntry) GetCertificate() CertificateDetails {
	if o == nil || IsNil(o.Certificate) {
		var ret CertificateDetails
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateEntry) GetCertificateOk() (*CertificateDetails, bool) {
	if o == nil || IsNil(o.Certificate) {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *CertificateEntry) HasCertificate() bool {
	if o != nil && !IsNil(o.Certificate) {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given CertificateDetails and assigns it to the Certificate field.
func (o *CertificateEntry) SetCertificate(v CertificateDetails) {
	o.Certificate = &v
}

func (o CertificateEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificateEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Alias) {
		toSerialize["alias"] = o.Alias
	}
	if !IsNil(o.Certificate) {
		toSerialize["certificate"] = o.Certificate
	}
	return toSerialize, nil
}

type NullableCertificateEntry struct {
	value *CertificateEntry
	isSet bool
}

func (v NullableCertificateEntry) Get() *CertificateEntry {
	return v.value
}

func (v *NullableCertificateEntry) Set(val *CertificateEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateEntry(val *CertificateEntry) *NullableCertificateEntry {
	return &NullableCertificateEntry{value: val, isSet: true}
}

func (v NullableCertificateEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

