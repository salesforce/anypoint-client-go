# EnvCore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**IsProduction** | Pointer to **bool** |  | [optional] [default to false]
**Type** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 

## Methods

### NewEnvCore

`func NewEnvCore() *EnvCore`

NewEnvCore instantiates a new EnvCore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvCoreWithDefaults

`func NewEnvCoreWithDefaults() *EnvCore`

NewEnvCoreWithDefaults instantiates a new EnvCore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnvCore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvCore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvCore) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnvCore) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganizationId

`func (o *EnvCore) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *EnvCore) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *EnvCore) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *EnvCore) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetIsProduction

`func (o *EnvCore) GetIsProduction() bool`

GetIsProduction returns the IsProduction field if non-nil, zero value otherwise.

### GetIsProductionOk

`func (o *EnvCore) GetIsProductionOk() (*bool, bool)`

GetIsProductionOk returns a tuple with the IsProduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProduction

`func (o *EnvCore) SetIsProduction(v bool)`

SetIsProduction sets IsProduction field to given value.

### HasIsProduction

`func (o *EnvCore) HasIsProduction() bool`

HasIsProduction returns a boolean if a field has been set.

### GetType

`func (o *EnvCore) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnvCore) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnvCore) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EnvCore) HasType() bool`

HasType returns a boolean if a field has been set.

### GetClientId

`func (o *EnvCore) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *EnvCore) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *EnvCore) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *EnvCore) HasClientId() bool`

HasClientId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


