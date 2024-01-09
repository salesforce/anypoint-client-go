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

// Audit struct for Audit
type Audit struct {
	Created *AuditCreated `json:"created,omitempty"`
	Updated *AuditCreated `json:"updated,omitempty"`
}

// NewAudit instantiates a new Audit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAudit() *Audit {
	this := Audit{}
	return &this
}

// NewAuditWithDefaults instantiates a new Audit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditWithDefaults() *Audit {
	this := Audit{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Audit) GetCreated() AuditCreated {
	if o == nil || o.Created == nil {
		var ret AuditCreated
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Audit) GetCreatedOk() (*AuditCreated, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Audit) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given AuditCreated and assigns it to the Created field.
func (o *Audit) SetCreated(v AuditCreated) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *Audit) GetUpdated() AuditCreated {
	if o == nil || o.Updated == nil {
		var ret AuditCreated
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Audit) GetUpdatedOk() (*AuditCreated, bool) {
	if o == nil || o.Updated == nil {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *Audit) HasUpdated() bool {
	if o != nil && o.Updated != nil {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given AuditCreated and assigns it to the Updated field.
func (o *Audit) SetUpdated(v AuditCreated) {
	o.Updated = &v
}

func (o Audit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Updated != nil {
		toSerialize["updated"] = o.Updated
	}
	return json.Marshal(toSerialize)
}

type NullableAudit struct {
	value *Audit
	isSet bool
}

func (v NullableAudit) Get() *Audit {
	return v.value
}

func (v *NullableAudit) Set(val *Audit) {
	v.value = val
	v.isSet = true
}

func (v NullableAudit) IsSet() bool {
	return v.isSet
}

func (v *NullableAudit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAudit(val *Audit) *NullableAudit {
	return &NullableAudit{value: val, isSet: true}
}

func (v NullableAudit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAudit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

