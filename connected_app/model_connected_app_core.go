/*
 * Connected App API
 *
 * Description of the Connected App API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package connected_app

import (
	"encoding/json"
)

// ConnectedAppCore struct for ConnectedAppCore
type ConnectedAppCore struct {
	ClientName *string `json:"client_name,omitempty"`
	GrantTypes *[]string `json:"grant_types,omitempty"`
	RedirectUris *[]string `json:"redirect_uris,omitempty"`
	Scopes *[]string `json:"scopes,omitempty"`
	// Application public key (PEM format). Used to validate JWT authorization grants.
	PublicKeys *[]string `json:"public_keys,omitempty"`
	ClientUri *string `json:"client_uri,omitempty"`
	Audience *string `json:"audience,omitempty"`
}

// NewConnectedAppCore instantiates a new ConnectedAppCore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectedAppCore() *ConnectedAppCore {
	this := ConnectedAppCore{}
	return &this
}

// NewConnectedAppCoreWithDefaults instantiates a new ConnectedAppCore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectedAppCoreWithDefaults() *ConnectedAppCore {
	this := ConnectedAppCore{}
	return &this
}

// GetClientName returns the ClientName field value if set, zero value otherwise.
func (o *ConnectedAppCore) GetClientName() string {
	if o == nil || o.ClientName == nil {
		var ret string
		return ret
	}
	return *o.ClientName
}

// GetClientNameOk returns a tuple with the ClientName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedAppCore) GetClientNameOk() (*string, bool) {
	if o == nil || o.ClientName == nil {
		return nil, false
	}
	return o.ClientName, true
}

// HasClientName returns a boolean if a field has been set.
func (o *ConnectedAppCore) HasClientName() bool {
	if o != nil && o.ClientName != nil {
		return true
	}

	return false
}

// SetClientName gets a reference to the given string and assigns it to the ClientName field.
func (o *ConnectedAppCore) SetClientName(v string) {
	o.ClientName = &v
}

// GetGrantTypes returns the GrantTypes field value if set, zero value otherwise.
func (o *ConnectedAppCore) GetGrantTypes() []string {
	if o == nil || o.GrantTypes == nil {
		var ret []string
		return ret
	}
	return *o.GrantTypes
}

// GetGrantTypesOk returns a tuple with the GrantTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedAppCore) GetGrantTypesOk() (*[]string, bool) {
	if o == nil || o.GrantTypes == nil {
		return nil, false
	}
	return o.GrantTypes, true
}

// HasGrantTypes returns a boolean if a field has been set.
func (o *ConnectedAppCore) HasGrantTypes() bool {
	if o != nil && o.GrantTypes != nil {
		return true
	}

	return false
}

// SetGrantTypes gets a reference to the given []string and assigns it to the GrantTypes field.
func (o *ConnectedAppCore) SetGrantTypes(v []string) {
	o.GrantTypes = &v
}

// GetRedirectUris returns the RedirectUris field value if set, zero value otherwise.
func (o *ConnectedAppCore) GetRedirectUris() []string {
	if o == nil || o.RedirectUris == nil {
		var ret []string
		return ret
	}
	return *o.RedirectUris
}

// GetRedirectUrisOk returns a tuple with the RedirectUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedAppCore) GetRedirectUrisOk() (*[]string, bool) {
	if o == nil || o.RedirectUris == nil {
		return nil, false
	}
	return o.RedirectUris, true
}

// HasRedirectUris returns a boolean if a field has been set.
func (o *ConnectedAppCore) HasRedirectUris() bool {
	if o != nil && o.RedirectUris != nil {
		return true
	}

	return false
}

// SetRedirectUris gets a reference to the given []string and assigns it to the RedirectUris field.
func (o *ConnectedAppCore) SetRedirectUris(v []string) {
	o.RedirectUris = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *ConnectedAppCore) GetScopes() []string {
	if o == nil || o.Scopes == nil {
		var ret []string
		return ret
	}
	return *o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedAppCore) GetScopesOk() (*[]string, bool) {
	if o == nil || o.Scopes == nil {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *ConnectedAppCore) HasScopes() bool {
	if o != nil && o.Scopes != nil {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *ConnectedAppCore) SetScopes(v []string) {
	o.Scopes = &v
}

// GetPublicKeys returns the PublicKeys field value if set, zero value otherwise.
func (o *ConnectedAppCore) GetPublicKeys() []string {
	if o == nil || o.PublicKeys == nil {
		var ret []string
		return ret
	}
	return *o.PublicKeys
}

// GetPublicKeysOk returns a tuple with the PublicKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedAppCore) GetPublicKeysOk() (*[]string, bool) {
	if o == nil || o.PublicKeys == nil {
		return nil, false
	}
	return o.PublicKeys, true
}

// HasPublicKeys returns a boolean if a field has been set.
func (o *ConnectedAppCore) HasPublicKeys() bool {
	if o != nil && o.PublicKeys != nil {
		return true
	}

	return false
}

// SetPublicKeys gets a reference to the given []string and assigns it to the PublicKeys field.
func (o *ConnectedAppCore) SetPublicKeys(v []string) {
	o.PublicKeys = &v
}

// GetClientUri returns the ClientUri field value if set, zero value otherwise.
func (o *ConnectedAppCore) GetClientUri() string {
	if o == nil || o.ClientUri == nil {
		var ret string
		return ret
	}
	return *o.ClientUri
}

// GetClientUriOk returns a tuple with the ClientUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedAppCore) GetClientUriOk() (*string, bool) {
	if o == nil || o.ClientUri == nil {
		return nil, false
	}
	return o.ClientUri, true
}

// HasClientUri returns a boolean if a field has been set.
func (o *ConnectedAppCore) HasClientUri() bool {
	if o != nil && o.ClientUri != nil {
		return true
	}

	return false
}

// SetClientUri gets a reference to the given string and assigns it to the ClientUri field.
func (o *ConnectedAppCore) SetClientUri(v string) {
	o.ClientUri = &v
}

// GetAudience returns the Audience field value if set, zero value otherwise.
func (o *ConnectedAppCore) GetAudience() string {
	if o == nil || o.Audience == nil {
		var ret string
		return ret
	}
	return *o.Audience
}

// GetAudienceOk returns a tuple with the Audience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedAppCore) GetAudienceOk() (*string, bool) {
	if o == nil || o.Audience == nil {
		return nil, false
	}
	return o.Audience, true
}

// HasAudience returns a boolean if a field has been set.
func (o *ConnectedAppCore) HasAudience() bool {
	if o != nil && o.Audience != nil {
		return true
	}

	return false
}

// SetAudience gets a reference to the given string and assigns it to the Audience field.
func (o *ConnectedAppCore) SetAudience(v string) {
	o.Audience = &v
}

func (o ConnectedAppCore) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientName != nil {
		toSerialize["client_name"] = o.ClientName
	}
	if o.GrantTypes != nil {
		toSerialize["grant_types"] = o.GrantTypes
	}
	if o.RedirectUris != nil {
		toSerialize["redirect_uris"] = o.RedirectUris
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	if o.PublicKeys != nil {
		toSerialize["public_keys"] = o.PublicKeys
	}
	if o.ClientUri != nil {
		toSerialize["client_uri"] = o.ClientUri
	}
	if o.Audience != nil {
		toSerialize["audience"] = o.Audience
	}
	return json.Marshal(toSerialize)
}

type NullableConnectedAppCore struct {
	value *ConnectedAppCore
	isSet bool
}

func (v NullableConnectedAppCore) Get() *ConnectedAppCore {
	return v.value
}

func (v *NullableConnectedAppCore) Set(val *ConnectedAppCore) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectedAppCore) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectedAppCore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectedAppCore(val *ConnectedAppCore) *NullableConnectedAppCore {
	return &NullableConnectedAppCore{value: val, isSet: true}
}

func (v NullableConnectedAppCore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectedAppCore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

