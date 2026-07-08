# ApplicationContractResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MasterOrganizationId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**CoreServicesId** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**ClientProvider** | Pointer to [**ApplicationContractResponseClientProvider**](ApplicationContractResponseClientProvider.md) |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**RedirectUri** | Pointer to **[]string** |  | [optional] 
**GrantTypes** | Pointer to **[]string** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**Owners** | Pointer to [**[]ApplicationOwner**](ApplicationOwner.md) |  | [optional] 

## Methods

### NewApplicationContractResponse

`func NewApplicationContractResponse() *ApplicationContractResponse`

NewApplicationContractResponse instantiates a new ApplicationContractResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationContractResponseWithDefaults

`func NewApplicationContractResponseWithDefaults() *ApplicationContractResponse`

NewApplicationContractResponseWithDefaults instantiates a new ApplicationContractResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMasterOrganizationId

`func (o *ApplicationContractResponse) GetMasterOrganizationId() string`

GetMasterOrganizationId returns the MasterOrganizationId field if non-nil, zero value otherwise.

### GetMasterOrganizationIdOk

`func (o *ApplicationContractResponse) GetMasterOrganizationIdOk() (*string, bool)`

GetMasterOrganizationIdOk returns a tuple with the MasterOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterOrganizationId

`func (o *ApplicationContractResponse) SetMasterOrganizationId(v string)`

SetMasterOrganizationId sets MasterOrganizationId field to given value.

### HasMasterOrganizationId

`func (o *ApplicationContractResponse) HasMasterOrganizationId() bool`

HasMasterOrganizationId returns a boolean if a field has been set.

### GetId

`func (o *ApplicationContractResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationContractResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationContractResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ApplicationContractResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ApplicationContractResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationContractResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationContractResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplicationContractResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ApplicationContractResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationContractResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationContractResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationContractResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ApplicationContractResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ApplicationContractResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCoreServicesId

`func (o *ApplicationContractResponse) GetCoreServicesId() string`

GetCoreServicesId returns the CoreServicesId field if non-nil, zero value otherwise.

### GetCoreServicesIdOk

`func (o *ApplicationContractResponse) GetCoreServicesIdOk() (*string, bool)`

GetCoreServicesIdOk returns a tuple with the CoreServicesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreServicesId

`func (o *ApplicationContractResponse) SetCoreServicesId(v string)`

SetCoreServicesId sets CoreServicesId field to given value.

### HasCoreServicesId

`func (o *ApplicationContractResponse) HasCoreServicesId() bool`

HasCoreServicesId returns a boolean if a field has been set.

### GetClientId

`func (o *ApplicationContractResponse) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ApplicationContractResponse) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ApplicationContractResponse) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ApplicationContractResponse) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientProvider

`func (o *ApplicationContractResponse) GetClientProvider() ApplicationContractResponseClientProvider`

GetClientProvider returns the ClientProvider field if non-nil, zero value otherwise.

### GetClientProviderOk

`func (o *ApplicationContractResponse) GetClientProviderOk() (*ApplicationContractResponseClientProvider, bool)`

GetClientProviderOk returns a tuple with the ClientProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientProvider

`func (o *ApplicationContractResponse) SetClientProvider(v ApplicationContractResponseClientProvider)`

SetClientProvider sets ClientProvider field to given value.

### HasClientProvider

`func (o *ApplicationContractResponse) HasClientProvider() bool`

HasClientProvider returns a boolean if a field has been set.

### GetOwner

`func (o *ApplicationContractResponse) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ApplicationContractResponse) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ApplicationContractResponse) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ApplicationContractResponse) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetEmail

`func (o *ApplicationContractResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ApplicationContractResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ApplicationContractResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ApplicationContractResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetRedirectUri

`func (o *ApplicationContractResponse) GetRedirectUri() []string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *ApplicationContractResponse) GetRedirectUriOk() (*[]string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *ApplicationContractResponse) SetRedirectUri(v []string)`

SetRedirectUri sets RedirectUri field to given value.

### HasRedirectUri

`func (o *ApplicationContractResponse) HasRedirectUri() bool`

HasRedirectUri returns a boolean if a field has been set.

### GetGrantTypes

`func (o *ApplicationContractResponse) GetGrantTypes() []string`

GetGrantTypes returns the GrantTypes field if non-nil, zero value otherwise.

### GetGrantTypesOk

`func (o *ApplicationContractResponse) GetGrantTypesOk() (*[]string, bool)`

GetGrantTypesOk returns a tuple with the GrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypes

`func (o *ApplicationContractResponse) SetGrantTypes(v []string)`

SetGrantTypes sets GrantTypes field to given value.

### HasGrantTypes

`func (o *ApplicationContractResponse) HasGrantTypes() bool`

HasGrantTypes returns a boolean if a field has been set.

### GetUrl

`func (o *ApplicationContractResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ApplicationContractResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ApplicationContractResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ApplicationContractResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ApplicationContractResponse) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ApplicationContractResponse) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetOwners

`func (o *ApplicationContractResponse) GetOwners() []ApplicationOwner`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *ApplicationContractResponse) GetOwnersOk() (*[]ApplicationOwner, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *ApplicationContractResponse) SetOwners(v []ApplicationOwner)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *ApplicationContractResponse) HasOwners() bool`

HasOwners returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


