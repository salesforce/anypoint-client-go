# TeamRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**RoleId** | Pointer to **string** |  | [optional] 
**ContextParams** | Pointer to [**ContextParams**](ContextParams.md) |  | [optional] 

## Methods

### NewTeamRole

`func NewTeamRole() *TeamRole`

NewTeamRole instantiates a new TeamRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamRoleWithDefaults

`func NewTeamRoleWithDefaults() *TeamRole`

NewTeamRoleWithDefaults instantiates a new TeamRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TeamRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamRole) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TeamRole) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRoleId

`func (o *TeamRole) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *TeamRole) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *TeamRole) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *TeamRole) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetContextParams

`func (o *TeamRole) GetContextParams() ContextParams`

GetContextParams returns the ContextParams field if non-nil, zero value otherwise.

### GetContextParamsOk

`func (o *TeamRole) GetContextParamsOk() (*ContextParams, bool)`

GetContextParamsOk returns a tuple with the ContextParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextParams

`func (o *TeamRole) SetContextParams(v ContextParams)`

SetContextParams sets ContextParams field to given value.

### HasContextParams

`func (o *TeamRole) HasContextParams() bool`

HasContextParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


