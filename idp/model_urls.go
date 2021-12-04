/*
 * Identity Provider Management API
 *
 * Description of Identity Provider API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package idp

import (
	"encoding/json"
)

// Urls struct for Urls
type Urls struct {
	Token *string `json:"token,omitempty"`
	Redirect *string `json:"redirect,omitempty"`
	Userinfo *string `json:"userinfo,omitempty"`
	Authorize *string `json:"authorize,omitempty"`
}

// NewUrls instantiates a new Urls object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUrls() *Urls {
	this := Urls{}
	return &this
}

// NewUrlsWithDefaults instantiates a new Urls object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUrlsWithDefaults() *Urls {
	this := Urls{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *Urls) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Urls) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *Urls) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *Urls) SetToken(v string) {
	o.Token = &v
}

// GetRedirect returns the Redirect field value if set, zero value otherwise.
func (o *Urls) GetRedirect() string {
	if o == nil || o.Redirect == nil {
		var ret string
		return ret
	}
	return *o.Redirect
}

// GetRedirectOk returns a tuple with the Redirect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Urls) GetRedirectOk() (*string, bool) {
	if o == nil || o.Redirect == nil {
		return nil, false
	}
	return o.Redirect, true
}

// HasRedirect returns a boolean if a field has been set.
func (o *Urls) HasRedirect() bool {
	if o != nil && o.Redirect != nil {
		return true
	}

	return false
}

// SetRedirect gets a reference to the given string and assigns it to the Redirect field.
func (o *Urls) SetRedirect(v string) {
	o.Redirect = &v
}

// GetUserinfo returns the Userinfo field value if set, zero value otherwise.
func (o *Urls) GetUserinfo() string {
	if o == nil || o.Userinfo == nil {
		var ret string
		return ret
	}
	return *o.Userinfo
}

// GetUserinfoOk returns a tuple with the Userinfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Urls) GetUserinfoOk() (*string, bool) {
	if o == nil || o.Userinfo == nil {
		return nil, false
	}
	return o.Userinfo, true
}

// HasUserinfo returns a boolean if a field has been set.
func (o *Urls) HasUserinfo() bool {
	if o != nil && o.Userinfo != nil {
		return true
	}

	return false
}

// SetUserinfo gets a reference to the given string and assigns it to the Userinfo field.
func (o *Urls) SetUserinfo(v string) {
	o.Userinfo = &v
}

// GetAuthorize returns the Authorize field value if set, zero value otherwise.
func (o *Urls) GetAuthorize() string {
	if o == nil || o.Authorize == nil {
		var ret string
		return ret
	}
	return *o.Authorize
}

// GetAuthorizeOk returns a tuple with the Authorize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Urls) GetAuthorizeOk() (*string, bool) {
	if o == nil || o.Authorize == nil {
		return nil, false
	}
	return o.Authorize, true
}

// HasAuthorize returns a boolean if a field has been set.
func (o *Urls) HasAuthorize() bool {
	if o != nil && o.Authorize != nil {
		return true
	}

	return false
}

// SetAuthorize gets a reference to the given string and assigns it to the Authorize field.
func (o *Urls) SetAuthorize(v string) {
	o.Authorize = &v
}

func (o Urls) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.Redirect != nil {
		toSerialize["redirect"] = o.Redirect
	}
	if o.Userinfo != nil {
		toSerialize["userinfo"] = o.Userinfo
	}
	if o.Authorize != nil {
		toSerialize["authorize"] = o.Authorize
	}
	return json.Marshal(toSerialize)
}

type NullableUrls struct {
	value *Urls
	isSet bool
}

func (v NullableUrls) Get() *Urls {
	return v.value
}

func (v *NullableUrls) Set(val *Urls) {
	v.value = val
	v.isSet = true
}

func (v NullableUrls) IsSet() bool {
	return v.isSet
}

func (v *NullableUrls) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUrls(val *Urls) *NullableUrls {
	return &NullableUrls{value: val, isSet: true}
}

func (v NullableUrls) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUrls) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


