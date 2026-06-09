# Environment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | The client id of the environment. | [optional] [default to ""]
**Id** | Pointer to **string** | The id of the environment. | [optional] [default to ""]
**IsProduction** | Pointer to **bool** | Whether the environment is a production environment. | [optional] [default to false]
**Name** | Pointer to **string** | The name of the environment. | [optional] [default to ""]
**OrganizationId** | Pointer to **string** | The id of the organization. | [optional] [default to ""]
**Type** | Pointer to **string** | The type of the environment. | [optional] [default to ""]

## Methods

### NewEnvironment

`func NewEnvironment() *Environment`

NewEnvironment instantiates a new Environment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentWithDefaults

`func NewEnvironmentWithDefaults() *Environment`

NewEnvironmentWithDefaults instantiates a new Environment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *Environment) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Environment) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Environment) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *Environment) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetId

`func (o *Environment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Environment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Environment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Environment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsProduction

`func (o *Environment) GetIsProduction() bool`

GetIsProduction returns the IsProduction field if non-nil, zero value otherwise.

### GetIsProductionOk

`func (o *Environment) GetIsProductionOk() (*bool, bool)`

GetIsProductionOk returns a tuple with the IsProduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProduction

`func (o *Environment) SetIsProduction(v bool)`

SetIsProduction sets IsProduction field to given value.

### HasIsProduction

`func (o *Environment) HasIsProduction() bool`

HasIsProduction returns a boolean if a field has been set.

### GetName

`func (o *Environment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Environment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Environment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Environment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganizationId

`func (o *Environment) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Environment) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *Environment) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *Environment) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetType

`func (o *Environment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Environment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Environment) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Environment) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


