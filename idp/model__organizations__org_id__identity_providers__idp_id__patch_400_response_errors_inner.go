/*
Identity Provider Management API

Description of Identity Provider API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package idp

import (
	"encoding/json"
)

// checks if the OrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner{}

// OrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner struct for OrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner
type OrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner struct {
	Type *string `json:"type,omitempty"`
	Keyword *string `json:"keyword,omitempty"`
	Schemas *string `json:"schemas,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewOrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner instantiates a new OrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner() *OrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner {
	this := OrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner{}
	return &this
}

// NewOrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInnerWithDefaults instantiates a new OrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInnerWithDefaults() *OrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner {
	this := OrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner) SetType(v string) {
	o.Type = &v
}

// GetKeyword returns the Keyword field value if set, zero value otherwise.
func (o *OrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner) GetKeyword() string {
	if o == nil || IsNil(o.Keyword) {
		var ret string
		return ret
	}
	return *o.Keyword
}

// GetKeywordOk returns a tuple with the Keyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner) GetKeywordOk() (*string, bool) {
	if o == nil || IsNil(o.Keyword) {
		return nil, false
	}
	return o.Keyword, true
}

// HasKeyword returns a boolean if a field has been set.
func (o *OrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner) HasKeyword() bool {
	if o != nil && !IsNil(o.Keyword) {
		return true
	}

	return false
}

// SetKeyword gets a reference to the given string and assigns it to the Keyword field.
func (o *OrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner) SetKeyword(v string) {
	o.Keyword = &v
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *OrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner) GetSchemas() string {
	if o == nil || IsNil(o.Schemas) {
		var ret string
		return ret
	}
	return *o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner) GetSchemasOk() (*string, bool) {
	if o == nil || IsNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *OrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given string and assigns it to the Schemas field.
func (o *OrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner) SetSchemas(v string) {
	o.Schemas = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *OrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *OrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *OrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner) SetMessage(v string) {
	o.Message = &v
}

func (o OrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Keyword) {
		toSerialize["keyword"] = o.Keyword
	}
	if !IsNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableOrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner struct {
	value *OrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner
	isSet bool
}

func (v NullableOrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner) Get() *OrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner {
	return v.value
}

func (v *NullableOrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner) Set(val *OrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner(val *OrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner) *NullableOrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner {
	return &NullableOrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner{value: val, isSet: true}
}

func (v NullableOrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


