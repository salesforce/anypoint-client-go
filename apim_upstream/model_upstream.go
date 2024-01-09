/*
 * API Manager Upstream API
 *
 * API Manager Upstream API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apim_upstream

import (
	"encoding/json"
)

// Upstream struct for Upstream
type Upstream struct {
	Id *string `json:"id,omitempty"`
	Label *string `json:"label,omitempty"`
	Uri *string `json:"uri,omitempty"`
	TlsContext *UpstreamTlsContext `json:"tlsContext,omitempty"`
}

// NewUpstream instantiates a new Upstream object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpstream() *Upstream {
	this := Upstream{}
	return &this
}

// NewUpstreamWithDefaults instantiates a new Upstream object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpstreamWithDefaults() *Upstream {
	this := Upstream{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Upstream) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Upstream) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Upstream) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Upstream) SetId(v string) {
	o.Id = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *Upstream) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Upstream) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *Upstream) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *Upstream) SetLabel(v string) {
	o.Label = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *Upstream) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Upstream) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *Upstream) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *Upstream) SetUri(v string) {
	o.Uri = &v
}

// GetTlsContext returns the TlsContext field value if set, zero value otherwise.
func (o *Upstream) GetTlsContext() UpstreamTlsContext {
	if o == nil || o.TlsContext == nil {
		var ret UpstreamTlsContext
		return ret
	}
	return *o.TlsContext
}

// GetTlsContextOk returns a tuple with the TlsContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Upstream) GetTlsContextOk() (*UpstreamTlsContext, bool) {
	if o == nil || o.TlsContext == nil {
		return nil, false
	}
	return o.TlsContext, true
}

// HasTlsContext returns a boolean if a field has been set.
func (o *Upstream) HasTlsContext() bool {
	if o != nil && o.TlsContext != nil {
		return true
	}

	return false
}

// SetTlsContext gets a reference to the given UpstreamTlsContext and assigns it to the TlsContext field.
func (o *Upstream) SetTlsContext(v UpstreamTlsContext) {
	o.TlsContext = &v
}

func (o Upstream) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Uri != nil {
		toSerialize["uri"] = o.Uri
	}
	if o.TlsContext != nil {
		toSerialize["tlsContext"] = o.TlsContext
	}
	return json.Marshal(toSerialize)
}

type NullableUpstream struct {
	value *Upstream
	isSet bool
}

func (v NullableUpstream) Get() *Upstream {
	return v.value
}

func (v *NullableUpstream) Set(val *Upstream) {
	v.value = val
	v.isSet = true
}

func (v NullableUpstream) IsSet() bool {
	return v.isSet
}

func (v *NullableUpstream) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpstream(val *Upstream) *NullableUpstream {
	return &NullableUpstream{value: val, isSet: true}
}

func (v NullableUpstream) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpstream) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

