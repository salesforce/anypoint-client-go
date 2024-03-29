/*
API Manager Policy API

API Manager Policy API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apim_policy

import (
	"encoding/json"
)

// checks if the ApimPolicyFullTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApimPolicyFullTemplate{}

// ApimPolicyFullTemplate struct for ApimPolicyFullTemplate
type ApimPolicyFullTemplate struct {
	GroupId *string `json:"groupId,omitempty"`
	AssetId *string `json:"assetId,omitempty"`
	AssetVersion *string `json:"assetVersion,omitempty"`
}

// NewApimPolicyFullTemplate instantiates a new ApimPolicyFullTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApimPolicyFullTemplate() *ApimPolicyFullTemplate {
	this := ApimPolicyFullTemplate{}
	return &this
}

// NewApimPolicyFullTemplateWithDefaults instantiates a new ApimPolicyFullTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApimPolicyFullTemplateWithDefaults() *ApimPolicyFullTemplate {
	this := ApimPolicyFullTemplate{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *ApimPolicyFullTemplate) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimPolicyFullTemplate) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *ApimPolicyFullTemplate) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *ApimPolicyFullTemplate) SetGroupId(v string) {
	o.GroupId = &v
}

// GetAssetId returns the AssetId field value if set, zero value otherwise.
func (o *ApimPolicyFullTemplate) GetAssetId() string {
	if o == nil || IsNil(o.AssetId) {
		var ret string
		return ret
	}
	return *o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimPolicyFullTemplate) GetAssetIdOk() (*string, bool) {
	if o == nil || IsNil(o.AssetId) {
		return nil, false
	}
	return o.AssetId, true
}

// HasAssetId returns a boolean if a field has been set.
func (o *ApimPolicyFullTemplate) HasAssetId() bool {
	if o != nil && !IsNil(o.AssetId) {
		return true
	}

	return false
}

// SetAssetId gets a reference to the given string and assigns it to the AssetId field.
func (o *ApimPolicyFullTemplate) SetAssetId(v string) {
	o.AssetId = &v
}

// GetAssetVersion returns the AssetVersion field value if set, zero value otherwise.
func (o *ApimPolicyFullTemplate) GetAssetVersion() string {
	if o == nil || IsNil(o.AssetVersion) {
		var ret string
		return ret
	}
	return *o.AssetVersion
}

// GetAssetVersionOk returns a tuple with the AssetVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimPolicyFullTemplate) GetAssetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.AssetVersion) {
		return nil, false
	}
	return o.AssetVersion, true
}

// HasAssetVersion returns a boolean if a field has been set.
func (o *ApimPolicyFullTemplate) HasAssetVersion() bool {
	if o != nil && !IsNil(o.AssetVersion) {
		return true
	}

	return false
}

// SetAssetVersion gets a reference to the given string and assigns it to the AssetVersion field.
func (o *ApimPolicyFullTemplate) SetAssetVersion(v string) {
	o.AssetVersion = &v
}

func (o ApimPolicyFullTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApimPolicyFullTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !IsNil(o.AssetId) {
		toSerialize["assetId"] = o.AssetId
	}
	if !IsNil(o.AssetVersion) {
		toSerialize["assetVersion"] = o.AssetVersion
	}
	return toSerialize, nil
}

type NullableApimPolicyFullTemplate struct {
	value *ApimPolicyFullTemplate
	isSet bool
}

func (v NullableApimPolicyFullTemplate) Get() *ApimPolicyFullTemplate {
	return v.value
}

func (v *NullableApimPolicyFullTemplate) Set(val *ApimPolicyFullTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableApimPolicyFullTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableApimPolicyFullTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApimPolicyFullTemplate(val *ApimPolicyFullTemplate) *NullableApimPolicyFullTemplate {
	return &NullableApimPolicyFullTemplate{value: val, isSet: true}
}

func (v NullableApimPolicyFullTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApimPolicyFullTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


