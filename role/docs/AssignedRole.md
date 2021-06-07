# AssignedRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContextParams** | Pointer to [**ContextParams**](ContextParams.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**RoleGroupAssignmentId** | Pointer to **string** |  | [optional] 
**RoleGroupId** | Pointer to **string** |  | [optional] 
**RoleId** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Internal** | Pointer to **bool** |  | [optional] 

## Methods

### NewAssignedRole

`func NewAssignedRole() *AssignedRole`

NewAssignedRole instantiates a new AssignedRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignedRoleWithDefaults

`func NewAssignedRoleWithDefaults() *AssignedRole`

NewAssignedRoleWithDefaults instantiates a new AssignedRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContextParams

`func (o *AssignedRole) GetContextParams() ContextParams`

GetContextParams returns the ContextParams field if non-nil, zero value otherwise.

### GetContextParamsOk

`func (o *AssignedRole) GetContextParamsOk() (*ContextParams, bool)`

GetContextParamsOk returns a tuple with the ContextParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextParams

`func (o *AssignedRole) SetContextParams(v ContextParams)`

SetContextParams sets ContextParams field to given value.

### HasContextParams

`func (o *AssignedRole) HasContextParams() bool`

HasContextParams returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AssignedRole) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AssignedRole) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AssignedRole) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AssignedRole) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetRoleGroupAssignmentId

`func (o *AssignedRole) GetRoleGroupAssignmentId() string`

GetRoleGroupAssignmentId returns the RoleGroupAssignmentId field if non-nil, zero value otherwise.

### GetRoleGroupAssignmentIdOk

`func (o *AssignedRole) GetRoleGroupAssignmentIdOk() (*string, bool)`

GetRoleGroupAssignmentIdOk returns a tuple with the RoleGroupAssignmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleGroupAssignmentId

`func (o *AssignedRole) SetRoleGroupAssignmentId(v string)`

SetRoleGroupAssignmentId sets RoleGroupAssignmentId field to given value.

### HasRoleGroupAssignmentId

`func (o *AssignedRole) HasRoleGroupAssignmentId() bool`

HasRoleGroupAssignmentId returns a boolean if a field has been set.

### GetRoleGroupId

`func (o *AssignedRole) GetRoleGroupId() string`

GetRoleGroupId returns the RoleGroupId field if non-nil, zero value otherwise.

### GetRoleGroupIdOk

`func (o *AssignedRole) GetRoleGroupIdOk() (*string, bool)`

GetRoleGroupIdOk returns a tuple with the RoleGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleGroupId

`func (o *AssignedRole) SetRoleGroupId(v string)`

SetRoleGroupId sets RoleGroupId field to given value.

### HasRoleGroupId

`func (o *AssignedRole) HasRoleGroupId() bool`

HasRoleGroupId returns a boolean if a field has been set.

### GetRoleId

`func (o *AssignedRole) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *AssignedRole) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *AssignedRole) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *AssignedRole) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetOrgId

`func (o *AssignedRole) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *AssignedRole) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *AssignedRole) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *AssignedRole) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetName

`func (o *AssignedRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssignedRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssignedRole) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AssignedRole) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AssignedRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AssignedRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AssignedRole) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AssignedRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInternal

`func (o *AssignedRole) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *AssignedRole) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *AssignedRole) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *AssignedRole) HasInternal() bool`

HasInternal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


