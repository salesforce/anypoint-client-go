# IdpOIDC

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowUntrustedCertificates** | Pointer to **bool** |  | [optional] 
**OidcProvider** | Pointer to [**OidcProvider**](OidcProvider.md) |  | [optional] 

## Methods

### NewIdpOIDC

`func NewIdpOIDC() *IdpOIDC`

NewIdpOIDC instantiates a new IdpOIDC object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpOIDCWithDefaults

`func NewIdpOIDCWithDefaults() *IdpOIDC`

NewIdpOIDCWithDefaults instantiates a new IdpOIDC object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowUntrustedCertificates

`func (o *IdpOIDC) GetAllowUntrustedCertificates() bool`

GetAllowUntrustedCertificates returns the AllowUntrustedCertificates field if non-nil, zero value otherwise.

### GetAllowUntrustedCertificatesOk

`func (o *IdpOIDC) GetAllowUntrustedCertificatesOk() (*bool, bool)`

GetAllowUntrustedCertificatesOk returns a tuple with the AllowUntrustedCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUntrustedCertificates

`func (o *IdpOIDC) SetAllowUntrustedCertificates(v bool)`

SetAllowUntrustedCertificates sets AllowUntrustedCertificates field to given value.

### HasAllowUntrustedCertificates

`func (o *IdpOIDC) HasAllowUntrustedCertificates() bool`

HasAllowUntrustedCertificates returns a boolean if a field has been set.

### GetOidcProvider

`func (o *IdpOIDC) GetOidcProvider() OidcProvider`

GetOidcProvider returns the OidcProvider field if non-nil, zero value otherwise.

### GetOidcProviderOk

`func (o *IdpOIDC) GetOidcProviderOk() (*OidcProvider, bool)`

GetOidcProviderOk returns a tuple with the OidcProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcProvider

`func (o *IdpOIDC) SetOidcProvider(v OidcProvider)`

SetOidcProvider sets OidcProvider field to given value.

### HasOidcProvider

`func (o *IdpOIDC) HasOidcProvider() bool`

HasOidcProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


