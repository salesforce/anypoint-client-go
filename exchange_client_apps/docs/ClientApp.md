# ClientApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**MasterOrganizationId** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**GrantTypes** | Pointer to **[]string** |  | [optional] 
**RedirectUri** | Pointer to **[]string** |  | [optional] 
**ClientProvider** | Pointer to [**ClientAppClientProvider**](ClientAppClientProvider.md) |  | [optional] 

## Methods

### NewClientApp

`func NewClientApp() *ClientApp`

NewClientApp instantiates a new ClientApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientAppWithDefaults

`func NewClientAppWithDefaults() *ClientApp`

NewClientAppWithDefaults instantiates a new ClientApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClientApp) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClientApp) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClientApp) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ClientApp) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ClientApp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClientApp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClientApp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClientApp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *ClientApp) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ClientApp) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ClientApp) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ClientApp) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ClientApp) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ClientApp) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetDescription

`func (o *ClientApp) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClientApp) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClientApp) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClientApp) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ClientApp) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ClientApp) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMasterOrganizationId

`func (o *ClientApp) GetMasterOrganizationId() string`

GetMasterOrganizationId returns the MasterOrganizationId field if non-nil, zero value otherwise.

### GetMasterOrganizationIdOk

`func (o *ClientApp) GetMasterOrganizationIdOk() (*string, bool)`

GetMasterOrganizationIdOk returns a tuple with the MasterOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterOrganizationId

`func (o *ClientApp) SetMasterOrganizationId(v string)`

SetMasterOrganizationId sets MasterOrganizationId field to given value.

### HasMasterOrganizationId

`func (o *ClientApp) HasMasterOrganizationId() bool`

HasMasterOrganizationId returns a boolean if a field has been set.

### GetClientId

`func (o *ClientApp) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ClientApp) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ClientApp) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ClientApp) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *ClientApp) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *ClientApp) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *ClientApp) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *ClientApp) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetGrantTypes

`func (o *ClientApp) GetGrantTypes() []string`

GetGrantTypes returns the GrantTypes field if non-nil, zero value otherwise.

### GetGrantTypesOk

`func (o *ClientApp) GetGrantTypesOk() (*[]string, bool)`

GetGrantTypesOk returns a tuple with the GrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypes

`func (o *ClientApp) SetGrantTypes(v []string)`

SetGrantTypes sets GrantTypes field to given value.

### HasGrantTypes

`func (o *ClientApp) HasGrantTypes() bool`

HasGrantTypes returns a boolean if a field has been set.

### GetRedirectUri

`func (o *ClientApp) GetRedirectUri() []string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *ClientApp) GetRedirectUriOk() (*[]string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *ClientApp) SetRedirectUri(v []string)`

SetRedirectUri sets RedirectUri field to given value.

### HasRedirectUri

`func (o *ClientApp) HasRedirectUri() bool`

HasRedirectUri returns a boolean if a field has been set.

### GetClientProvider

`func (o *ClientApp) GetClientProvider() ClientAppClientProvider`

GetClientProvider returns the ClientProvider field if non-nil, zero value otherwise.

### GetClientProviderOk

`func (o *ClientApp) GetClientProviderOk() (*ClientAppClientProvider, bool)`

GetClientProviderOk returns a tuple with the ClientProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientProvider

`func (o *ClientApp) SetClientProvider(v ClientAppClientProvider)`

SetClientProvider sets ClientProvider field to given value.

### HasClientProvider

`func (o *ClientApp) HasClientProvider() bool`

HasClientProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


