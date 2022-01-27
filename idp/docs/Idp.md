# Idp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderId** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ModelType**](ModelType.md) |  | [optional] 
**Saml** | Pointer to [**Saml**](Saml.md) |  | [optional] 
**AllowUntrustedCertificates** | Pointer to **bool** |  | [optional] 
**OidcProvider** | Pointer to [**OidcProvider**](OidcProvider.md) |  | [optional] 
**ServiceProvider** | Pointer to [**ServiceProvider**](ServiceProvider.md) |  | [optional] 

## Methods

### NewIdp

`func NewIdp() *Idp`

NewIdp instantiates a new Idp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpWithDefaults

`func NewIdpWithDefaults() *Idp`

NewIdpWithDefaults instantiates a new Idp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderId

`func (o *Idp) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *Idp) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *Idp) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *Idp) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetOrgId

`func (o *Idp) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Idp) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Idp) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *Idp) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetName

`func (o *Idp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Idp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Idp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Idp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *Idp) GetType() ModelType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Idp) GetTypeOk() (*ModelType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Idp) SetType(v ModelType)`

SetType sets Type field to given value.

### HasType

`func (o *Idp) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSaml

`func (o *Idp) GetSaml() Saml`

GetSaml returns the Saml field if non-nil, zero value otherwise.

### GetSamlOk

`func (o *Idp) GetSamlOk() (*Saml, bool)`

GetSamlOk returns a tuple with the Saml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml

`func (o *Idp) SetSaml(v Saml)`

SetSaml sets Saml field to given value.

### HasSaml

`func (o *Idp) HasSaml() bool`

HasSaml returns a boolean if a field has been set.

### GetAllowUntrustedCertificates

`func (o *Idp) GetAllowUntrustedCertificates() bool`

GetAllowUntrustedCertificates returns the AllowUntrustedCertificates field if non-nil, zero value otherwise.

### GetAllowUntrustedCertificatesOk

`func (o *Idp) GetAllowUntrustedCertificatesOk() (*bool, bool)`

GetAllowUntrustedCertificatesOk returns a tuple with the AllowUntrustedCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUntrustedCertificates

`func (o *Idp) SetAllowUntrustedCertificates(v bool)`

SetAllowUntrustedCertificates sets AllowUntrustedCertificates field to given value.

### HasAllowUntrustedCertificates

`func (o *Idp) HasAllowUntrustedCertificates() bool`

HasAllowUntrustedCertificates returns a boolean if a field has been set.

### GetOidcProvider

`func (o *Idp) GetOidcProvider() OidcProvider`

GetOidcProvider returns the OidcProvider field if non-nil, zero value otherwise.

### GetOidcProviderOk

`func (o *Idp) GetOidcProviderOk() (*OidcProvider, bool)`

GetOidcProviderOk returns a tuple with the OidcProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcProvider

`func (o *Idp) SetOidcProvider(v OidcProvider)`

SetOidcProvider sets OidcProvider field to given value.

### HasOidcProvider

`func (o *Idp) HasOidcProvider() bool`

HasOidcProvider returns a boolean if a field has been set.

### GetServiceProvider

`func (o *Idp) GetServiceProvider() ServiceProvider`

GetServiceProvider returns the ServiceProvider field if non-nil, zero value otherwise.

### GetServiceProviderOk

`func (o *Idp) GetServiceProviderOk() (*ServiceProvider, bool)`

GetServiceProviderOk returns a tuple with the ServiceProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProvider

`func (o *Idp) SetServiceProvider(v ServiceProvider)`

SetServiceProvider sets ServiceProvider field to given value.

### HasServiceProvider

`func (o *Idp) HasServiceProvider() bool`

HasServiceProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


