# TeamRolePostBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleId** | Pointer to **string** |  | [optional] 
**ContextParams** | Pointer to [**ContextParams**](ContextParams.md) |  | [optional] 

## Methods

### NewTeamRolePostBody

`func NewTeamRolePostBody() *TeamRolePostBody`

NewTeamRolePostBody instantiates a new TeamRolePostBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamRolePostBodyWithDefaults

`func NewTeamRolePostBodyWithDefaults() *TeamRolePostBody`

NewTeamRolePostBodyWithDefaults instantiates a new TeamRolePostBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleId

`func (o *TeamRolePostBody) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *TeamRolePostBody) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *TeamRolePostBody) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *TeamRolePostBody) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetContextParams

`func (o *TeamRolePostBody) GetContextParams() ContextParams`

GetContextParams returns the ContextParams field if non-nil, zero value otherwise.

### GetContextParamsOk

`func (o *TeamRolePostBody) GetContextParamsOk() (*ContextParams, bool)`

GetContextParamsOk returns a tuple with the ContextParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextParams

`func (o *TeamRolePostBody) SetContextParams(v ContextParams)`

SetContextParams sets ContextParams field to given value.

### HasContextParams

`func (o *TeamRolePostBody) HasContextParams() bool`

HasContextParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


