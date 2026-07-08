# Application

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audit** | Pointer to [**Audit**](Audit.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ClientProvider** | Pointer to **map[string]interface{}** |  | [optional] 
**CoreServicesId** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewApplication

`func NewApplication() *Application`

NewApplication instantiates a new Application object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationWithDefaults

`func NewApplicationWithDefaults() *Application`

NewApplicationWithDefaults instantiates a new Application object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudit

`func (o *Application) GetAudit() Audit`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *Application) GetAuditOk() (*Audit, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *Application) SetAudit(v Audit)`

SetAudit sets Audit field to given value.

### HasAudit

`func (o *Application) HasAudit() bool`

HasAudit returns a boolean if a field has been set.

### GetId

`func (o *Application) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Application) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Application) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Application) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Application) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Application) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Application) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Application) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Application) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Application) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Application) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Application) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Application) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Application) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetClientProvider

`func (o *Application) GetClientProvider() map[string]interface{}`

GetClientProvider returns the ClientProvider field if non-nil, zero value otherwise.

### GetClientProviderOk

`func (o *Application) GetClientProviderOk() (*map[string]interface{}, bool)`

GetClientProviderOk returns a tuple with the ClientProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientProvider

`func (o *Application) SetClientProvider(v map[string]interface{})`

SetClientProvider sets ClientProvider field to given value.

### HasClientProvider

`func (o *Application) HasClientProvider() bool`

HasClientProvider returns a boolean if a field has been set.

### GetCoreServicesId

`func (o *Application) GetCoreServicesId() string`

GetCoreServicesId returns the CoreServicesId field if non-nil, zero value otherwise.

### GetCoreServicesIdOk

`func (o *Application) GetCoreServicesIdOk() (*string, bool)`

GetCoreServicesIdOk returns a tuple with the CoreServicesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreServicesId

`func (o *Application) SetCoreServicesId(v string)`

SetCoreServicesId sets CoreServicesId field to given value.

### HasCoreServicesId

`func (o *Application) HasCoreServicesId() bool`

HasCoreServicesId returns a boolean if a field has been set.

### GetUrl

`func (o *Application) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Application) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Application) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Application) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *Application) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *Application) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


