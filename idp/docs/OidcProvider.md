# OidcProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Client** | Pointer to [**Client**](Client.md) |  | [optional] 
**Issuer** | Pointer to **string** |  | [optional] 
**Urls** | Pointer to [**Urls3**](Urls3.md) |  | [optional] 
**GroupScope** | Pointer to **string** |  | [optional] 
**ClaimsMapping** | Pointer to [**ClaimsMapping1**](ClaimsMapping1.md) |  | [optional] 

## Methods

### NewOidcProvider

`func NewOidcProvider() *OidcProvider`

NewOidcProvider instantiates a new OidcProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOidcProviderWithDefaults

`func NewOidcProviderWithDefaults() *OidcProvider`

NewOidcProviderWithDefaults instantiates a new OidcProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClient

`func (o *OidcProvider) GetClient() Client`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *OidcProvider) GetClientOk() (*Client, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *OidcProvider) SetClient(v Client)`

SetClient sets Client field to given value.

### HasClient

`func (o *OidcProvider) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetIssuer

`func (o *OidcProvider) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *OidcProvider) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *OidcProvider) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *OidcProvider) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetUrls

`func (o *OidcProvider) GetUrls() Urls3`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *OidcProvider) GetUrlsOk() (*Urls3, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *OidcProvider) SetUrls(v Urls3)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *OidcProvider) HasUrls() bool`

HasUrls returns a boolean if a field has been set.

### GetGroupScope

`func (o *OidcProvider) GetGroupScope() string`

GetGroupScope returns the GroupScope field if non-nil, zero value otherwise.

### GetGroupScopeOk

`func (o *OidcProvider) GetGroupScopeOk() (*string, bool)`

GetGroupScopeOk returns a tuple with the GroupScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupScope

`func (o *OidcProvider) SetGroupScope(v string)`

SetGroupScope sets GroupScope field to given value.

### HasGroupScope

`func (o *OidcProvider) HasGroupScope() bool`

HasGroupScope returns a boolean if a field has been set.

### GetClaimsMapping

`func (o *OidcProvider) GetClaimsMapping() ClaimsMapping1`

GetClaimsMapping returns the ClaimsMapping field if non-nil, zero value otherwise.

### GetClaimsMappingOk

`func (o *OidcProvider) GetClaimsMappingOk() (*ClaimsMapping1, bool)`

GetClaimsMappingOk returns a tuple with the ClaimsMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimsMapping

`func (o *OidcProvider) SetClaimsMapping(v ClaimsMapping1)`

SetClaimsMapping sets ClaimsMapping field to given value.

### HasClaimsMapping

`func (o *OidcProvider) HasClaimsMapping() bool`

HasClaimsMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


