# IdpPatchBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**IdpPatchBodyType**](IdpPatchBodyType.md) |  | [optional] 
**OidcProvider** | Pointer to [**OidcProvider1**](OidcProvider1.md) |  | [optional] 
**AllowUntrustedCertificates** | Pointer to **bool** |  | [optional] 
**Saml** | Pointer to [**Saml1**](Saml1.md) |  | [optional] 
**ServiceProvider** | Pointer to [**ServiceProvider1**](ServiceProvider1.md) |  | [optional] 

## Methods

### NewIdpPatchBody

`func NewIdpPatchBody() *IdpPatchBody`

NewIdpPatchBody instantiates a new IdpPatchBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpPatchBodyWithDefaults

`func NewIdpPatchBodyWithDefaults() *IdpPatchBody`

NewIdpPatchBodyWithDefaults instantiates a new IdpPatchBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IdpPatchBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdpPatchBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdpPatchBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdpPatchBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *IdpPatchBody) GetType() IdpPatchBodyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IdpPatchBody) GetTypeOk() (*IdpPatchBodyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IdpPatchBody) SetType(v IdpPatchBodyType)`

SetType sets Type field to given value.

### HasType

`func (o *IdpPatchBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOidcProvider

`func (o *IdpPatchBody) GetOidcProvider() OidcProvider1`

GetOidcProvider returns the OidcProvider field if non-nil, zero value otherwise.

### GetOidcProviderOk

`func (o *IdpPatchBody) GetOidcProviderOk() (*OidcProvider1, bool)`

GetOidcProviderOk returns a tuple with the OidcProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcProvider

`func (o *IdpPatchBody) SetOidcProvider(v OidcProvider1)`

SetOidcProvider sets OidcProvider field to given value.

### HasOidcProvider

`func (o *IdpPatchBody) HasOidcProvider() bool`

HasOidcProvider returns a boolean if a field has been set.

### GetAllowUntrustedCertificates

`func (o *IdpPatchBody) GetAllowUntrustedCertificates() bool`

GetAllowUntrustedCertificates returns the AllowUntrustedCertificates field if non-nil, zero value otherwise.

### GetAllowUntrustedCertificatesOk

`func (o *IdpPatchBody) GetAllowUntrustedCertificatesOk() (*bool, bool)`

GetAllowUntrustedCertificatesOk returns a tuple with the AllowUntrustedCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUntrustedCertificates

`func (o *IdpPatchBody) SetAllowUntrustedCertificates(v bool)`

SetAllowUntrustedCertificates sets AllowUntrustedCertificates field to given value.

### HasAllowUntrustedCertificates

`func (o *IdpPatchBody) HasAllowUntrustedCertificates() bool`

HasAllowUntrustedCertificates returns a boolean if a field has been set.

### GetSaml

`func (o *IdpPatchBody) GetSaml() Saml1`

GetSaml returns the Saml field if non-nil, zero value otherwise.

### GetSamlOk

`func (o *IdpPatchBody) GetSamlOk() (*Saml1, bool)`

GetSamlOk returns a tuple with the Saml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml

`func (o *IdpPatchBody) SetSaml(v Saml1)`

SetSaml sets Saml field to given value.

### HasSaml

`func (o *IdpPatchBody) HasSaml() bool`

HasSaml returns a boolean if a field has been set.

### GetServiceProvider

`func (o *IdpPatchBody) GetServiceProvider() ServiceProvider1`

GetServiceProvider returns the ServiceProvider field if non-nil, zero value otherwise.

### GetServiceProviderOk

`func (o *IdpPatchBody) GetServiceProviderOk() (*ServiceProvider1, bool)`

GetServiceProviderOk returns a tuple with the ServiceProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProvider

`func (o *IdpPatchBody) SetServiceProvider(v ServiceProvider1)`

SetServiceProvider sets ServiceProvider field to given value.

### HasServiceProvider

`func (o *IdpPatchBody) HasServiceProvider() bool`

HasServiceProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


