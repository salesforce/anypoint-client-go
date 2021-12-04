# IdpOIDC

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Urls** | Pointer to [**Urls**](Urls.md) |  | [optional] 
**Client** | Pointer to **string** |  | [optional] 
**Issuer** | Pointer to **string** |  | [optional] 
**GroupScope** | Pointer to **string** |  | [optional] 
**AllowUntrustedCertificates** | Pointer to **bool** |  | [optional] 

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

### GetUrls

`func (o *IdpOIDC) GetUrls() Urls`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *IdpOIDC) GetUrlsOk() (*Urls, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *IdpOIDC) SetUrls(v Urls)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *IdpOIDC) HasUrls() bool`

HasUrls returns a boolean if a field has been set.

### GetClient

`func (o *IdpOIDC) GetClient() string`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *IdpOIDC) GetClientOk() (*string, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *IdpOIDC) SetClient(v string)`

SetClient sets Client field to given value.

### HasClient

`func (o *IdpOIDC) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetIssuer

`func (o *IdpOIDC) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *IdpOIDC) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *IdpOIDC) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *IdpOIDC) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetGroupScope

`func (o *IdpOIDC) GetGroupScope() string`

GetGroupScope returns the GroupScope field if non-nil, zero value otherwise.

### GetGroupScopeOk

`func (o *IdpOIDC) GetGroupScopeOk() (*string, bool)`

GetGroupScopeOk returns a tuple with the GroupScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupScope

`func (o *IdpOIDC) SetGroupScope(v string)`

SetGroupScope sets GroupScope field to given value.

### HasGroupScope

`func (o *IdpOIDC) HasGroupScope() bool`

HasGroupScope returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


