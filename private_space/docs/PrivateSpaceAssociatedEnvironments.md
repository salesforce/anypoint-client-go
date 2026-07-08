# PrivateSpaceAssociatedEnvironments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of environments. Possible values are all, sandbox, production. | [optional] 
**BusinessGroups** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPrivateSpaceAssociatedEnvironments

`func NewPrivateSpaceAssociatedEnvironments() *PrivateSpaceAssociatedEnvironments`

NewPrivateSpaceAssociatedEnvironments instantiates a new PrivateSpaceAssociatedEnvironments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateSpaceAssociatedEnvironmentsWithDefaults

`func NewPrivateSpaceAssociatedEnvironmentsWithDefaults() *PrivateSpaceAssociatedEnvironments`

NewPrivateSpaceAssociatedEnvironmentsWithDefaults instantiates a new PrivateSpaceAssociatedEnvironments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PrivateSpaceAssociatedEnvironments) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PrivateSpaceAssociatedEnvironments) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PrivateSpaceAssociatedEnvironments) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PrivateSpaceAssociatedEnvironments) HasType() bool`

HasType returns a boolean if a field has been set.

### GetBusinessGroups

`func (o *PrivateSpaceAssociatedEnvironments) GetBusinessGroups() []string`

GetBusinessGroups returns the BusinessGroups field if non-nil, zero value otherwise.

### GetBusinessGroupsOk

`func (o *PrivateSpaceAssociatedEnvironments) GetBusinessGroupsOk() (*[]string, bool)`

GetBusinessGroupsOk returns a tuple with the BusinessGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessGroups

`func (o *PrivateSpaceAssociatedEnvironments) SetBusinessGroups(v []string)`

SetBusinessGroups sets BusinessGroups field to given value.

### HasBusinessGroups

`func (o *PrivateSpaceAssociatedEnvironments) HasBusinessGroups() bool`

HasBusinessGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


