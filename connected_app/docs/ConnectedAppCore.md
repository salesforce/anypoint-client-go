# ConnectedAppCore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientName** | Pointer to **string** |  | [optional] 
**GrantTypes** | Pointer to **[]string** |  | [optional] 
**RedirectUris** | Pointer to **[]string** |  | [optional] 
**Scopes** | Pointer to **[]string** |  | [optional] 
**PublicKeys** | Pointer to **[]string** | Application public key (PEM format). Used to validate JWT authorization grants. | [optional] 
**ClientUri** | Pointer to **string** |  | [optional] 
**Audience** | Pointer to **string** |  | [optional] 

## Methods

### NewConnectedAppCore

`func NewConnectedAppCore() *ConnectedAppCore`

NewConnectedAppCore instantiates a new ConnectedAppCore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectedAppCoreWithDefaults

`func NewConnectedAppCoreWithDefaults() *ConnectedAppCore`

NewConnectedAppCoreWithDefaults instantiates a new ConnectedAppCore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientName

`func (o *ConnectedAppCore) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *ConnectedAppCore) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *ConnectedAppCore) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *ConnectedAppCore) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetGrantTypes

`func (o *ConnectedAppCore) GetGrantTypes() []string`

GetGrantTypes returns the GrantTypes field if non-nil, zero value otherwise.

### GetGrantTypesOk

`func (o *ConnectedAppCore) GetGrantTypesOk() (*[]string, bool)`

GetGrantTypesOk returns a tuple with the GrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypes

`func (o *ConnectedAppCore) SetGrantTypes(v []string)`

SetGrantTypes sets GrantTypes field to given value.

### HasGrantTypes

`func (o *ConnectedAppCore) HasGrantTypes() bool`

HasGrantTypes returns a boolean if a field has been set.

### GetRedirectUris

`func (o *ConnectedAppCore) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *ConnectedAppCore) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *ConnectedAppCore) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *ConnectedAppCore) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.

### GetScopes

`func (o *ConnectedAppCore) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ConnectedAppCore) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ConnectedAppCore) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ConnectedAppCore) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetPublicKeys

`func (o *ConnectedAppCore) GetPublicKeys() []string`

GetPublicKeys returns the PublicKeys field if non-nil, zero value otherwise.

### GetPublicKeysOk

`func (o *ConnectedAppCore) GetPublicKeysOk() (*[]string, bool)`

GetPublicKeysOk returns a tuple with the PublicKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeys

`func (o *ConnectedAppCore) SetPublicKeys(v []string)`

SetPublicKeys sets PublicKeys field to given value.

### HasPublicKeys

`func (o *ConnectedAppCore) HasPublicKeys() bool`

HasPublicKeys returns a boolean if a field has been set.

### GetClientUri

`func (o *ConnectedAppCore) GetClientUri() string`

GetClientUri returns the ClientUri field if non-nil, zero value otherwise.

### GetClientUriOk

`func (o *ConnectedAppCore) GetClientUriOk() (*string, bool)`

GetClientUriOk returns a tuple with the ClientUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientUri

`func (o *ConnectedAppCore) SetClientUri(v string)`

SetClientUri sets ClientUri field to given value.

### HasClientUri

`func (o *ConnectedAppCore) HasClientUri() bool`

HasClientUri returns a boolean if a field has been set.

### GetAudience

`func (o *ConnectedAppCore) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *ConnectedAppCore) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *ConnectedAppCore) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *ConnectedAppCore) HasAudience() bool`

HasAudience returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


