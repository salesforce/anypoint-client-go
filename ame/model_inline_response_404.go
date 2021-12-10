/*
 * Anypoint MQ Exchange specfication
 *
 * Anypoint MQ Exchange API specification
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ame

import (
	"encoding/json"
)

// InlineResponse404 struct for InlineResponse404
type InlineResponse404 struct {
	Status *int32 `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewInlineResponse404 instantiates a new InlineResponse404 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse404() *InlineResponse404 {
	this := InlineResponse404{}
	var status int32 = 404
	this.Status = &status
	return &this
}

// NewInlineResponse404WithDefaults instantiates a new InlineResponse404 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse404WithDefaults() *InlineResponse404 {
	this := InlineResponse404{}
	var status int32 = 404
	this.Status = &status
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse404) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse404) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse404) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *InlineResponse404) SetStatus(v int32) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *InlineResponse404) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse404) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *InlineResponse404) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *InlineResponse404) SetMessage(v string) {
	o.Message = &v
}

func (o InlineResponse404) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse404 struct {
	value *InlineResponse404
	isSet bool
}

func (v NullableInlineResponse404) Get() *InlineResponse404 {
	return v.value
}

func (v *NullableInlineResponse404) Set(val *InlineResponse404) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse404) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse404) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse404(val *InlineResponse404) *NullableInlineResponse404 {
	return &NullableInlineResponse404{value: val, isSet: true}
}

func (v NullableInlineResponse404) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse404) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


