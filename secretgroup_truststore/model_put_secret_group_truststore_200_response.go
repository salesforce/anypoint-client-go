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

// checks if the PutSecretGroupTruststore200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutSecretGroupTruststore200Response{}

// PutSecretGroupTruststore200Response struct for PutSecretGroupTruststore200Response
type PutSecretGroupTruststore200Response struct {
	Message *string `json:"message,omitempty"`
}

// NewPutSecretGroupTruststore200Response instantiates a new PutSecretGroupTruststore200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutSecretGroupTruststore200Response() *PutSecretGroupTruststore200Response {
	this := PutSecretGroupTruststore200Response{}
	return &this
}

// NewPutSecretGroupTruststore200ResponseWithDefaults instantiates a new PutSecretGroupTruststore200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutSecretGroupTruststore200ResponseWithDefaults() *PutSecretGroupTruststore200Response {
	this := PutSecretGroupTruststore200Response{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *PutSecretGroupTruststore200Response) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutSecretGroupTruststore200Response) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *PutSecretGroupTruststore200Response) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *PutSecretGroupTruststore200Response) SetMessage(v string) {
	o.Message = &v
}

func (o PutSecretGroupTruststore200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutSecretGroupTruststore200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullablePutSecretGroupTruststore200Response struct {
	value *PutSecretGroupTruststore200Response
	isSet bool
}

func (v NullablePutSecretGroupTruststore200Response) Get() *PutSecretGroupTruststore200Response {
	return v.value
}

func (v *NullablePutSecretGroupTruststore200Response) Set(val *PutSecretGroupTruststore200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePutSecretGroupTruststore200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePutSecretGroupTruststore200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutSecretGroupTruststore200Response(val *PutSecretGroupTruststore200Response) *NullablePutSecretGroupTruststore200Response {
	return &NullablePutSecretGroupTruststore200Response{value: val, isSet: true}
}

func (v NullablePutSecretGroupTruststore200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutSecretGroupTruststore200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


