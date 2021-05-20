# Env

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The env Id | 
**Name** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**IsProduction** | Pointer to **bool** |  | [optional] [default to false]
**Type** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 

## Methods

### NewEnv

`func NewEnv(id string, ) *Env`

NewEnv instantiates a new Env object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvWithDefaults

`func NewEnvWithDefaults() *Env`

NewEnvWithDefaults instantiates a new Env object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Env) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Env) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Env) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Env) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Env) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Env) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Env) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganizationId

`func (o *Env) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Env) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *Env) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *Env) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetIsProduction

`func (o *Env) GetIsProduction() bool`

GetIsProduction returns the IsProduction field if non-nil, zero value otherwise.

### GetIsProductionOk

`func (o *Env) GetIsProductionOk() (*bool, bool)`

GetIsProductionOk returns a tuple with the IsProduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProduction

`func (o *Env) SetIsProduction(v bool)`

SetIsProduction sets IsProduction field to given value.

### HasIsProduction

`func (o *Env) HasIsProduction() bool`

HasIsProduction returns a boolean if a field has been set.

### GetType

`func (o *Env) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Env) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Env) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Env) HasType() bool`

HasType returns a boolean if a field has been set.

### GetClientId

`func (o *Env) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Env) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Env) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *Env) HasClientId() bool`

HasClientId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


