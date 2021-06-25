# InvitePostBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**RoleGroups** | Pointer to [**RoleGroups**](RoleGroups.md) |  | [optional] 

## Methods

### NewInvitePostBody

`func NewInvitePostBody() *InvitePostBody`

NewInvitePostBody instantiates a new InvitePostBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvitePostBodyWithDefaults

`func NewInvitePostBodyWithDefaults() *InvitePostBody`

NewInvitePostBodyWithDefaults instantiates a new InvitePostBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *InvitePostBody) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InvitePostBody) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InvitePostBody) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *InvitePostBody) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetRoleGroups

`func (o *InvitePostBody) GetRoleGroups() RoleGroups`

GetRoleGroups returns the RoleGroups field if non-nil, zero value otherwise.

### GetRoleGroupsOk

`func (o *InvitePostBody) GetRoleGroupsOk() (*RoleGroups, bool)`

GetRoleGroupsOk returns a tuple with the RoleGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleGroups

`func (o *InvitePostBody) SetRoleGroups(v RoleGroups)`

SetRoleGroups sets RoleGroups field to given value.

### HasRoleGroups

`func (o *InvitePostBody) HasRoleGroups() bool`

HasRoleGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


