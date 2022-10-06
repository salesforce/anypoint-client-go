# ConnectedAppPatchExt

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
**ClientId** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewConnectedAppPatchExt

`func NewConnectedAppPatchExt() *ConnectedAppPatchExt`

NewConnectedAppPatchExt instantiates a new ConnectedAppPatchExt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectedAppPatchExtWithDefaults

`func NewConnectedAppPatchExtWithDefaults() *ConnectedAppPatchExt`

NewConnectedAppPatchExtWithDefaults instantiates a new ConnectedAppPatchExt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientName

`func (o *ConnectedAppPatchExt) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *ConnectedAppPatchExt) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *ConnectedAppPatchExt) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *ConnectedAppPatchExt) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetGrantTypes

`func (o *ConnectedAppPatchExt) GetGrantTypes() []string`

GetGrantTypes returns the GrantTypes field if non-nil, zero value otherwise.

### GetGrantTypesOk

`func (o *ConnectedAppPatchExt) GetGrantTypesOk() (*[]string, bool)`

GetGrantTypesOk returns a tuple with the GrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypes

`func (o *ConnectedAppPatchExt) SetGrantTypes(v []string)`

SetGrantTypes sets GrantTypes field to given value.

### HasGrantTypes

`func (o *ConnectedAppPatchExt) HasGrantTypes() bool`

HasGrantTypes returns a boolean if a field has been set.

### GetRedirectUris

`func (o *ConnectedAppPatchExt) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *ConnectedAppPatchExt) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *ConnectedAppPatchExt) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *ConnectedAppPatchExt) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.

### GetScopes

`func (o *ConnectedAppPatchExt) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ConnectedAppPatchExt) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ConnectedAppPatchExt) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ConnectedAppPatchExt) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetPublicKeys

`func (o *ConnectedAppPatchExt) GetPublicKeys() []string`

GetPublicKeys returns the PublicKeys field if non-nil, zero value otherwise.

### GetPublicKeysOk

`func (o *ConnectedAppPatchExt) GetPublicKeysOk() (*[]string, bool)`

GetPublicKeysOk returns a tuple with the PublicKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeys

`func (o *ConnectedAppPatchExt) SetPublicKeys(v []string)`

SetPublicKeys sets PublicKeys field to given value.

### HasPublicKeys

`func (o *ConnectedAppPatchExt) HasPublicKeys() bool`

HasPublicKeys returns a boolean if a field has been set.

### GetClientUri

`func (o *ConnectedAppPatchExt) GetClientUri() string`

GetClientUri returns the ClientUri field if non-nil, zero value otherwise.

### GetClientUriOk

`func (o *ConnectedAppPatchExt) GetClientUriOk() (*string, bool)`

GetClientUriOk returns a tuple with the ClientUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientUri

`func (o *ConnectedAppPatchExt) SetClientUri(v string)`

SetClientUri sets ClientUri field to given value.

### HasClientUri

`func (o *ConnectedAppPatchExt) HasClientUri() bool`

HasClientUri returns a boolean if a field has been set.

### GetAudience

`func (o *ConnectedAppPatchExt) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *ConnectedAppPatchExt) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *ConnectedAppPatchExt) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *ConnectedAppPatchExt) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetClientId

`func (o *ConnectedAppPatchExt) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ConnectedAppPatchExt) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ConnectedAppPatchExt) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ConnectedAppPatchExt) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *ConnectedAppPatchExt) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *ConnectedAppPatchExt) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *ConnectedAppPatchExt) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *ConnectedAppPatchExt) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetEnabled

`func (o *ConnectedAppPatchExt) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ConnectedAppPatchExt) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ConnectedAppPatchExt) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ConnectedAppPatchExt) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


