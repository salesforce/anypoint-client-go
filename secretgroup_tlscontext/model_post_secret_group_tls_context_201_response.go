/*
Secret Group TLS Context API

Secret Group TLS Context API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package secretgroup_tlscontext

import (
	"encoding/json"
)

// checks if the PostSecretGroupTlsContext201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostSecretGroupTlsContext201Response{}

// PostSecretGroupTlsContext201Response struct for PostSecretGroupTlsContext201Response
type PostSecretGroupTlsContext201Response struct {
	Message *string `json:"message,omitempty"`
	Id *string `json:"id,omitempty"`
}

// NewPostSecretGroupTlsContext201Response instantiates a new PostSecretGroupTlsContext201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostSecretGroupTlsContext201Response() *PostSecretGroupTlsContext201Response {
	this := PostSecretGroupTlsContext201Response{}
	return &this
}

// NewPostSecretGroupTlsContext201ResponseWithDefaults instantiates a new PostSecretGroupTlsContext201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostSecretGroupTlsContext201ResponseWithDefaults() *PostSecretGroupTlsContext201Response {
	this := PostSecretGroupTlsContext201Response{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *PostSecretGroupTlsContext201Response) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostSecretGroupTlsContext201Response) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *PostSecretGroupTlsContext201Response) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *PostSecretGroupTlsContext201Response) SetMessage(v string) {
	o.Message = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PostSecretGroupTlsContext201Response) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostSecretGroupTlsContext201Response) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PostSecretGroupTlsContext201Response) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PostSecretGroupTlsContext201Response) SetId(v string) {
	o.Id = &v
}

func (o PostSecretGroupTlsContext201Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostSecretGroupTlsContext201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePostSecretGroupTlsContext201Response struct {
	value *PostSecretGroupTlsContext201Response
	isSet bool
}

func (v NullablePostSecretGroupTlsContext201Response) Get() *PostSecretGroupTlsContext201Response {
	return v.value
}

func (v *NullablePostSecretGroupTlsContext201Response) Set(val *PostSecretGroupTlsContext201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePostSecretGroupTlsContext201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePostSecretGroupTlsContext201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostSecretGroupTlsContext201Response(val *PostSecretGroupTlsContext201Response) *NullablePostSecretGroupTlsContext201Response {
	return &NullablePostSecretGroupTlsContext201Response{value: val, isSet: true}
}

func (v NullablePostSecretGroupTlsContext201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostSecretGroupTlsContext201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


