/*
Invite API

Description of the Invite API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package invite

import (
	"encoding/json"
)

// checks if the OrganizationsOrgIdInvitesGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationsOrgIdInvitesGet200Response{}

// OrganizationsOrgIdInvitesGet200Response struct for OrganizationsOrgIdInvitesGet200Response
type OrganizationsOrgIdInvitesGet200Response struct {
	Data []Invite `json:"data,omitempty"`
	Total *int32 `json:"total,omitempty"`
}

// NewOrganizationsOrgIdInvitesGet200Response instantiates a new OrganizationsOrgIdInvitesGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrgIdInvitesGet200Response() *OrganizationsOrgIdInvitesGet200Response {
	this := OrganizationsOrgIdInvitesGet200Response{}
	return &this
}

// NewOrganizationsOrgIdInvitesGet200ResponseWithDefaults instantiates a new OrganizationsOrgIdInvitesGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrgIdInvitesGet200ResponseWithDefaults() *OrganizationsOrgIdInvitesGet200Response {
	this := OrganizationsOrgIdInvitesGet200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *OrganizationsOrgIdInvitesGet200Response) GetData() []Invite {
	if o == nil || IsNil(o.Data) {
		var ret []Invite
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrgIdInvitesGet200Response) GetDataOk() ([]Invite, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *OrganizationsOrgIdInvitesGet200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Invite and assigns it to the Data field.
func (o *OrganizationsOrgIdInvitesGet200Response) SetData(v []Invite) {
	o.Data = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *OrganizationsOrgIdInvitesGet200Response) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrgIdInvitesGet200Response) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *OrganizationsOrgIdInvitesGet200Response) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *OrganizationsOrgIdInvitesGet200Response) SetTotal(v int32) {
	o.Total = &v
}

func (o OrganizationsOrgIdInvitesGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationsOrgIdInvitesGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableOrganizationsOrgIdInvitesGet200Response struct {
	value *OrganizationsOrgIdInvitesGet200Response
	isSet bool
}

func (v NullableOrganizationsOrgIdInvitesGet200Response) Get() *OrganizationsOrgIdInvitesGet200Response {
	return v.value
}

func (v *NullableOrganizationsOrgIdInvitesGet200Response) Set(val *OrganizationsOrgIdInvitesGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrgIdInvitesGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrgIdInvitesGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrgIdInvitesGet200Response(val *OrganizationsOrgIdInvitesGet200Response) *NullableOrganizationsOrgIdInvitesGet200Response {
	return &NullableOrganizationsOrgIdInvitesGet200Response{value: val, isSet: true}
}

func (v NullableOrganizationsOrgIdInvitesGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrgIdInvitesGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


