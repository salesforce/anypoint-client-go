/*
 * API Manager API
 *
 * API Manager API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apim

import (
	"encoding/json"
)

// Spec struct for Spec
type Spec struct {
	AssetId *string `json:"assetId,omitempty"`
	GroupId *string `json:"groupId,omitempty"`
	Version *string `json:"version,omitempty"`
}

// NewSpec instantiates a new Spec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpec() *Spec {
	this := Spec{}
	return &this
}

// NewSpecWithDefaults instantiates a new Spec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpecWithDefaults() *Spec {
	this := Spec{}
	return &this
}

// GetAssetId returns the AssetId field value if set, zero value otherwise.
func (o *Spec) GetAssetId() string {
	if o == nil || o.AssetId == nil {
		var ret string
		return ret
	}
	return *o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Spec) GetAssetIdOk() (*string, bool) {
	if o == nil || o.AssetId == nil {
		return nil, false
	}
	return o.AssetId, true
}

// HasAssetId returns a boolean if a field has been set.
func (o *Spec) HasAssetId() bool {
	if o != nil && o.AssetId != nil {
		return true
	}

	return false
}

// SetAssetId gets a reference to the given string and assigns it to the AssetId field.
func (o *Spec) SetAssetId(v string) {
	o.AssetId = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *Spec) GetGroupId() string {
	if o == nil || o.GroupId == nil {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Spec) GetGroupIdOk() (*string, bool) {
	if o == nil || o.GroupId == nil {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *Spec) HasGroupId() bool {
	if o != nil && o.GroupId != nil {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *Spec) SetGroupId(v string) {
	o.GroupId = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Spec) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Spec) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Spec) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Spec) SetVersion(v string) {
	o.Version = &v
}

func (o Spec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AssetId != nil {
		toSerialize["assetId"] = o.AssetId
	}
	if o.GroupId != nil {
		toSerialize["groupId"] = o.GroupId
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableSpec struct {
	value *Spec
	isSet bool
}

func (v NullableSpec) Get() *Spec {
	return v.value
}

func (v *NullableSpec) Set(val *Spec) {
	v.value = val
	v.isSet = true
}

func (v NullableSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpec(val *Spec) *NullableSpec {
	return &NullableSpec{value: val, isSet: true}
}

func (v NullableSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


