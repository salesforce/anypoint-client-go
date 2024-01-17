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

// checks if the IssuerSubject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IssuerSubject{}

// IssuerSubject struct for IssuerSubject
type IssuerSubject struct {
	CommonName *string `json:"commonName,omitempty"`
	OrganizationName *string `json:"organizationName,omitempty"`
	LocalityName *string `json:"localityName,omitempty"`
	OrganizationUnit *string `json:"organizationUnit,omitempty"`
	State *string `json:"state,omitempty"`
	CountryName *string `json:"countryName,omitempty"`
}

// NewIssuerSubject instantiates a new IssuerSubject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssuerSubject() *IssuerSubject {
	this := IssuerSubject{}
	return &this
}

// NewIssuerSubjectWithDefaults instantiates a new IssuerSubject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssuerSubjectWithDefaults() *IssuerSubject {
	this := IssuerSubject{}
	return &this
}

// GetCommonName returns the CommonName field value if set, zero value otherwise.
func (o *IssuerSubject) GetCommonName() string {
	if o == nil || IsNil(o.CommonName) {
		var ret string
		return ret
	}
	return *o.CommonName
}

// GetCommonNameOk returns a tuple with the CommonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuerSubject) GetCommonNameOk() (*string, bool) {
	if o == nil || IsNil(o.CommonName) {
		return nil, false
	}
	return o.CommonName, true
}

// HasCommonName returns a boolean if a field has been set.
func (o *IssuerSubject) HasCommonName() bool {
	if o != nil && !IsNil(o.CommonName) {
		return true
	}

	return false
}

// SetCommonName gets a reference to the given string and assigns it to the CommonName field.
func (o *IssuerSubject) SetCommonName(v string) {
	o.CommonName = &v
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise.
func (o *IssuerSubject) GetOrganizationName() string {
	if o == nil || IsNil(o.OrganizationName) {
		var ret string
		return ret
	}
	return *o.OrganizationName
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuerSubject) GetOrganizationNameOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationName) {
		return nil, false
	}
	return o.OrganizationName, true
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *IssuerSubject) HasOrganizationName() bool {
	if o != nil && !IsNil(o.OrganizationName) {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given string and assigns it to the OrganizationName field.
func (o *IssuerSubject) SetOrganizationName(v string) {
	o.OrganizationName = &v
}

// GetLocalityName returns the LocalityName field value if set, zero value otherwise.
func (o *IssuerSubject) GetLocalityName() string {
	if o == nil || IsNil(o.LocalityName) {
		var ret string
		return ret
	}
	return *o.LocalityName
}

// GetLocalityNameOk returns a tuple with the LocalityName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuerSubject) GetLocalityNameOk() (*string, bool) {
	if o == nil || IsNil(o.LocalityName) {
		return nil, false
	}
	return o.LocalityName, true
}

// HasLocalityName returns a boolean if a field has been set.
func (o *IssuerSubject) HasLocalityName() bool {
	if o != nil && !IsNil(o.LocalityName) {
		return true
	}

	return false
}

// SetLocalityName gets a reference to the given string and assigns it to the LocalityName field.
func (o *IssuerSubject) SetLocalityName(v string) {
	o.LocalityName = &v
}

// GetOrganizationUnit returns the OrganizationUnit field value if set, zero value otherwise.
func (o *IssuerSubject) GetOrganizationUnit() string {
	if o == nil || IsNil(o.OrganizationUnit) {
		var ret string
		return ret
	}
	return *o.OrganizationUnit
}

// GetOrganizationUnitOk returns a tuple with the OrganizationUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuerSubject) GetOrganizationUnitOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationUnit) {
		return nil, false
	}
	return o.OrganizationUnit, true
}

// HasOrganizationUnit returns a boolean if a field has been set.
func (o *IssuerSubject) HasOrganizationUnit() bool {
	if o != nil && !IsNil(o.OrganizationUnit) {
		return true
	}

	return false
}

// SetOrganizationUnit gets a reference to the given string and assigns it to the OrganizationUnit field.
func (o *IssuerSubject) SetOrganizationUnit(v string) {
	o.OrganizationUnit = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *IssuerSubject) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuerSubject) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *IssuerSubject) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *IssuerSubject) SetState(v string) {
	o.State = &v
}

// GetCountryName returns the CountryName field value if set, zero value otherwise.
func (o *IssuerSubject) GetCountryName() string {
	if o == nil || IsNil(o.CountryName) {
		var ret string
		return ret
	}
	return *o.CountryName
}

// GetCountryNameOk returns a tuple with the CountryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuerSubject) GetCountryNameOk() (*string, bool) {
	if o == nil || IsNil(o.CountryName) {
		return nil, false
	}
	return o.CountryName, true
}

// HasCountryName returns a boolean if a field has been set.
func (o *IssuerSubject) HasCountryName() bool {
	if o != nil && !IsNil(o.CountryName) {
		return true
	}

	return false
}

// SetCountryName gets a reference to the given string and assigns it to the CountryName field.
func (o *IssuerSubject) SetCountryName(v string) {
	o.CountryName = &v
}

func (o IssuerSubject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IssuerSubject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CommonName) {
		toSerialize["commonName"] = o.CommonName
	}
	if !IsNil(o.OrganizationName) {
		toSerialize["organizationName"] = o.OrganizationName
	}
	if !IsNil(o.LocalityName) {
		toSerialize["localityName"] = o.LocalityName
	}
	if !IsNil(o.OrganizationUnit) {
		toSerialize["organizationUnit"] = o.OrganizationUnit
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.CountryName) {
		toSerialize["countryName"] = o.CountryName
	}
	return toSerialize, nil
}

type NullableIssuerSubject struct {
	value *IssuerSubject
	isSet bool
}

func (v NullableIssuerSubject) Get() *IssuerSubject {
	return v.value
}

func (v *NullableIssuerSubject) Set(val *IssuerSubject) {
	v.value = val
	v.isSet = true
}

func (v NullableIssuerSubject) IsSet() bool {
	return v.isSet
}

func (v *NullableIssuerSubject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssuerSubject(val *IssuerSubject) *NullableIssuerSubject {
	return &NullableIssuerSubject{value: val, isSet: true}
}

func (v NullableIssuerSubject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssuerSubject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


