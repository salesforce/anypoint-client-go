# TeamPatchBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TeamName** | Pointer to **string** |  | [optional] 
**TeamType** | Pointer to **string** |  | [optional] [default to "internal"]

## Methods

### NewTeamPatchBody

`func NewTeamPatchBody() *TeamPatchBody`

NewTeamPatchBody instantiates a new TeamPatchBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamPatchBodyWithDefaults

`func NewTeamPatchBodyWithDefaults() *TeamPatchBody`

NewTeamPatchBodyWithDefaults instantiates a new TeamPatchBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTeamName

`func (o *TeamPatchBody) GetTeamName() string`

GetTeamName returns the TeamName field if non-nil, zero value otherwise.

### GetTeamNameOk

`func (o *TeamPatchBody) GetTeamNameOk() (*string, bool)`

GetTeamNameOk returns a tuple with the TeamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamName

`func (o *TeamPatchBody) SetTeamName(v string)`

SetTeamName sets TeamName field to given value.

### HasTeamName

`func (o *TeamPatchBody) HasTeamName() bool`

HasTeamName returns a boolean if a field has been set.

### GetTeamType

`func (o *TeamPatchBody) GetTeamType() string`

GetTeamType returns the TeamType field if non-nil, zero value otherwise.

### GetTeamTypeOk

`func (o *TeamPatchBody) GetTeamTypeOk() (*string, bool)`

GetTeamTypeOk returns a tuple with the TeamType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamType

`func (o *TeamPatchBody) SetTeamType(v string)`

SetTeamType sets TeamType field to given value.

### HasTeamType

`func (o *TeamPatchBody) HasTeamType() bool`

HasTeamType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


