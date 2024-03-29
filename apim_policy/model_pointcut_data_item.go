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

// checks if the PointcutDataItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PointcutDataItem{}

// PointcutDataItem struct for PointcutDataItem
type PointcutDataItem struct {
	MethodRegex *string `json:"methodRegex,omitempty"`
	UriTemplateRegex *string `json:"uriTemplateRegex,omitempty"`
}

// NewPointcutDataItem instantiates a new PointcutDataItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPointcutDataItem() *PointcutDataItem {
	this := PointcutDataItem{}
	return &this
}

// NewPointcutDataItemWithDefaults instantiates a new PointcutDataItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPointcutDataItemWithDefaults() *PointcutDataItem {
	this := PointcutDataItem{}
	return &this
}

// GetMethodRegex returns the MethodRegex field value if set, zero value otherwise.
func (o *PointcutDataItem) GetMethodRegex() string {
	if o == nil || IsNil(o.MethodRegex) {
		var ret string
		return ret
	}
	return *o.MethodRegex
}

// GetMethodRegexOk returns a tuple with the MethodRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PointcutDataItem) GetMethodRegexOk() (*string, bool) {
	if o == nil || IsNil(o.MethodRegex) {
		return nil, false
	}
	return o.MethodRegex, true
}

// HasMethodRegex returns a boolean if a field has been set.
func (o *PointcutDataItem) HasMethodRegex() bool {
	if o != nil && !IsNil(o.MethodRegex) {
		return true
	}

	return false
}

// SetMethodRegex gets a reference to the given string and assigns it to the MethodRegex field.
func (o *PointcutDataItem) SetMethodRegex(v string) {
	o.MethodRegex = &v
}

// GetUriTemplateRegex returns the UriTemplateRegex field value if set, zero value otherwise.
func (o *PointcutDataItem) GetUriTemplateRegex() string {
	if o == nil || IsNil(o.UriTemplateRegex) {
		var ret string
		return ret
	}
	return *o.UriTemplateRegex
}

// GetUriTemplateRegexOk returns a tuple with the UriTemplateRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PointcutDataItem) GetUriTemplateRegexOk() (*string, bool) {
	if o == nil || IsNil(o.UriTemplateRegex) {
		return nil, false
	}
	return o.UriTemplateRegex, true
}

// HasUriTemplateRegex returns a boolean if a field has been set.
func (o *PointcutDataItem) HasUriTemplateRegex() bool {
	if o != nil && !IsNil(o.UriTemplateRegex) {
		return true
	}

	return false
}

// SetUriTemplateRegex gets a reference to the given string and assigns it to the UriTemplateRegex field.
func (o *PointcutDataItem) SetUriTemplateRegex(v string) {
	o.UriTemplateRegex = &v
}

func (o PointcutDataItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PointcutDataItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MethodRegex) {
		toSerialize["methodRegex"] = o.MethodRegex
	}
	if !IsNil(o.UriTemplateRegex) {
		toSerialize["uriTemplateRegex"] = o.UriTemplateRegex
	}
	return toSerialize, nil
}

type NullablePointcutDataItem struct {
	value *PointcutDataItem
	isSet bool
}

func (v NullablePointcutDataItem) Get() *PointcutDataItem {
	return v.value
}

func (v *NullablePointcutDataItem) Set(val *PointcutDataItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePointcutDataItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePointcutDataItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePointcutDataItem(val *PointcutDataItem) *NullablePointcutDataItem {
	return &NullablePointcutDataItem{value: val, isSet: true}
}

func (v NullablePointcutDataItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePointcutDataItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


