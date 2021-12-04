# Saml

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Issuer** | Pointer to **string** |  | [optional] 
**Audience** | Pointer to **string** |  | [optional] 
**PublicKey** | Pointer to **[]string** |  | [optional] 
**ClaimsMapping** | Pointer to [**ClaimsMapping**](ClaimsMapping.md) |  | [optional] 
**SpInitiatedSsoEnabled** | Pointer to **bool** |  | [optional] 
**IdpInitiatedSsoEnabled** | Pointer to **bool** |  | [optional] 
**RequireEncryptedSamlAssertions** | Pointer to **bool** |  | [optional] 

## Methods

### NewSaml

`func NewSaml() *Saml`

NewSaml instantiates a new Saml object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlWithDefaults

`func NewSamlWithDefaults() *Saml`

NewSamlWithDefaults instantiates a new Saml object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuer

`func (o *Saml) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *Saml) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *Saml) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *Saml) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetAudience

`func (o *Saml) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *Saml) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *Saml) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *Saml) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetPublicKey

`func (o *Saml) GetPublicKey() []string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *Saml) GetPublicKeyOk() (*[]string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *Saml) SetPublicKey(v []string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *Saml) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetClaimsMapping

`func (o *Saml) GetClaimsMapping() ClaimsMapping`

GetClaimsMapping returns the ClaimsMapping field if non-nil, zero value otherwise.

### GetClaimsMappingOk

`func (o *Saml) GetClaimsMappingOk() (*ClaimsMapping, bool)`

GetClaimsMappingOk returns a tuple with the ClaimsMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimsMapping

`func (o *Saml) SetClaimsMapping(v ClaimsMapping)`

SetClaimsMapping sets ClaimsMapping field to given value.

### HasClaimsMapping

`func (o *Saml) HasClaimsMapping() bool`

HasClaimsMapping returns a boolean if a field has been set.

### GetSpInitiatedSsoEnabled

`func (o *Saml) GetSpInitiatedSsoEnabled() bool`

GetSpInitiatedSsoEnabled returns the SpInitiatedSsoEnabled field if non-nil, zero value otherwise.

### GetSpInitiatedSsoEnabledOk

`func (o *Saml) GetSpInitiatedSsoEnabledOk() (*bool, bool)`

GetSpInitiatedSsoEnabledOk returns a tuple with the SpInitiatedSsoEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpInitiatedSsoEnabled

`func (o *Saml) SetSpInitiatedSsoEnabled(v bool)`

SetSpInitiatedSsoEnabled sets SpInitiatedSsoEnabled field to given value.

### HasSpInitiatedSsoEnabled

`func (o *Saml) HasSpInitiatedSsoEnabled() bool`

HasSpInitiatedSsoEnabled returns a boolean if a field has been set.

### GetIdpInitiatedSsoEnabled

`func (o *Saml) GetIdpInitiatedSsoEnabled() bool`

GetIdpInitiatedSsoEnabled returns the IdpInitiatedSsoEnabled field if non-nil, zero value otherwise.

### GetIdpInitiatedSsoEnabledOk

`func (o *Saml) GetIdpInitiatedSsoEnabledOk() (*bool, bool)`

GetIdpInitiatedSsoEnabledOk returns a tuple with the IdpInitiatedSsoEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpInitiatedSsoEnabled

`func (o *Saml) SetIdpInitiatedSsoEnabled(v bool)`

SetIdpInitiatedSsoEnabled sets IdpInitiatedSsoEnabled field to given value.

### HasIdpInitiatedSsoEnabled

`func (o *Saml) HasIdpInitiatedSsoEnabled() bool`

HasIdpInitiatedSsoEnabled returns a boolean if a field has been set.

### GetRequireEncryptedSamlAssertions

`func (o *Saml) GetRequireEncryptedSamlAssertions() bool`

GetRequireEncryptedSamlAssertions returns the RequireEncryptedSamlAssertions field if non-nil, zero value otherwise.

### GetRequireEncryptedSamlAssertionsOk

`func (o *Saml) GetRequireEncryptedSamlAssertionsOk() (*bool, bool)`

GetRequireEncryptedSamlAssertionsOk returns a tuple with the RequireEncryptedSamlAssertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireEncryptedSamlAssertions

`func (o *Saml) SetRequireEncryptedSamlAssertions(v bool)`

SetRequireEncryptedSamlAssertions sets RequireEncryptedSamlAssertions field to given value.

### HasRequireEncryptedSamlAssertions

`func (o *Saml) HasRequireEncryptedSamlAssertions() bool`

HasRequireEncryptedSamlAssertions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


