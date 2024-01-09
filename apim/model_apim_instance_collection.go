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

// ApimInstanceCollection struct for ApimInstanceCollection
type ApimInstanceCollection struct {
	Total *int32 `json:"total,omitempty"`
	Assets *[]ApimInstanceCollectionAssets `json:"assets,omitempty"`
}

// NewApimInstanceCollection instantiates a new ApimInstanceCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApimInstanceCollection() *ApimInstanceCollection {
	this := ApimInstanceCollection{}
	return &this
}

// NewApimInstanceCollectionWithDefaults instantiates a new ApimInstanceCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApimInstanceCollectionWithDefaults() *ApimInstanceCollection {
	this := ApimInstanceCollection{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *ApimInstanceCollection) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstanceCollection) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *ApimInstanceCollection) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *ApimInstanceCollection) SetTotal(v int32) {
	o.Total = &v
}

// GetAssets returns the Assets field value if set, zero value otherwise.
func (o *ApimInstanceCollection) GetAssets() []ApimInstanceCollectionAssets {
	if o == nil || o.Assets == nil {
		var ret []ApimInstanceCollectionAssets
		return ret
	}
	return *o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstanceCollection) GetAssetsOk() (*[]ApimInstanceCollectionAssets, bool) {
	if o == nil || o.Assets == nil {
		return nil, false
	}
	return o.Assets, true
}

// HasAssets returns a boolean if a field has been set.
func (o *ApimInstanceCollection) HasAssets() bool {
	if o != nil && o.Assets != nil {
		return true
	}

	return false
}

// SetAssets gets a reference to the given []ApimInstanceCollectionAssets and assigns it to the Assets field.
func (o *ApimInstanceCollection) SetAssets(v []ApimInstanceCollectionAssets) {
	o.Assets = &v
}

func (o ApimInstanceCollection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	if o.Assets != nil {
		toSerialize["assets"] = o.Assets
	}
	return json.Marshal(toSerialize)
}

type NullableApimInstanceCollection struct {
	value *ApimInstanceCollection
	isSet bool
}

func (v NullableApimInstanceCollection) Get() *ApimInstanceCollection {
	return v.value
}

func (v *NullableApimInstanceCollection) Set(val *ApimInstanceCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableApimInstanceCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableApimInstanceCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApimInstanceCollection(val *ApimInstanceCollection) *NullableApimInstanceCollection {
	return &NullableApimInstanceCollection{value: val, isSet: true}
}

func (v NullableApimInstanceCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApimInstanceCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

