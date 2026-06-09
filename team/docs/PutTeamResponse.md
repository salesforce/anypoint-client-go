# PutTeamResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | Pointer to **string** |  | [optional] 
**TeamId** | Pointer to **string** |  | [optional] 
**TeamName** | Pointer to **string** |  | [optional] 
**TeamType** | Pointer to **string** |  | [optional] 
**AncestorTeamIds** | Pointer to **[]string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**PreviousAncestorTeamIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPutTeamResponse

`func NewPutTeamResponse() *PutTeamResponse`

NewPutTeamResponse instantiates a new PutTeamResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutTeamResponseWithDefaults

`func NewPutTeamResponseWithDefaults() *PutTeamResponse`

NewPutTeamResponseWithDefaults instantiates a new PutTeamResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *PutTeamResponse) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *PutTeamResponse) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *PutTeamResponse) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *PutTeamResponse) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetTeamId

`func (o *PutTeamResponse) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *PutTeamResponse) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *PutTeamResponse) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *PutTeamResponse) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetTeamName

`func (o *PutTeamResponse) GetTeamName() string`

GetTeamName returns the TeamName field if non-nil, zero value otherwise.

### GetTeamNameOk

`func (o *PutTeamResponse) GetTeamNameOk() (*string, bool)`

GetTeamNameOk returns a tuple with the TeamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamName

`func (o *PutTeamResponse) SetTeamName(v string)`

SetTeamName sets TeamName field to given value.

### HasTeamName

`func (o *PutTeamResponse) HasTeamName() bool`

HasTeamName returns a boolean if a field has been set.

### GetTeamType

`func (o *PutTeamResponse) GetTeamType() string`

GetTeamType returns the TeamType field if non-nil, zero value otherwise.

### GetTeamTypeOk

`func (o *PutTeamResponse) GetTeamTypeOk() (*string, bool)`

GetTeamTypeOk returns a tuple with the TeamType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamType

`func (o *PutTeamResponse) SetTeamType(v string)`

SetTeamType sets TeamType field to given value.

### HasTeamType

`func (o *PutTeamResponse) HasTeamType() bool`

HasTeamType returns a boolean if a field has been set.

### GetAncestorTeamIds

`func (o *PutTeamResponse) GetAncestorTeamIds() []string`

GetAncestorTeamIds returns the AncestorTeamIds field if non-nil, zero value otherwise.

### GetAncestorTeamIdsOk

`func (o *PutTeamResponse) GetAncestorTeamIdsOk() (*[]string, bool)`

GetAncestorTeamIdsOk returns a tuple with the AncestorTeamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestorTeamIds

`func (o *PutTeamResponse) SetAncestorTeamIds(v []string)`

SetAncestorTeamIds sets AncestorTeamIds field to given value.

### HasAncestorTeamIds

`func (o *PutTeamResponse) HasAncestorTeamIds() bool`

HasAncestorTeamIds returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PutTeamResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PutTeamResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PutTeamResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PutTeamResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PutTeamResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PutTeamResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PutTeamResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PutTeamResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetPreviousAncestorTeamIds

`func (o *PutTeamResponse) GetPreviousAncestorTeamIds() []string`

GetPreviousAncestorTeamIds returns the PreviousAncestorTeamIds field if non-nil, zero value otherwise.

### GetPreviousAncestorTeamIdsOk

`func (o *PutTeamResponse) GetPreviousAncestorTeamIdsOk() (*[]string, bool)`

GetPreviousAncestorTeamIdsOk returns a tuple with the PreviousAncestorTeamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAncestorTeamIds

`func (o *PutTeamResponse) SetPreviousAncestorTeamIds(v []string)`

SetPreviousAncestorTeamIds sets PreviousAncestorTeamIds field to given value.

### HasPreviousAncestorTeamIds

`func (o *PutTeamResponse) HasPreviousAncestorTeamIds() bool`

HasPreviousAncestorTeamIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


