# TeamPostBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentTeamId** | Pointer to **string** |  | [optional] 
**TeamName** | Pointer to **string** |  | [optional] 
**TeamType** | Pointer to **string** |  | [optional] [default to "internal"]

## Methods

### NewTeamPostBody

`func NewTeamPostBody() *TeamPostBody`

NewTeamPostBody instantiates a new TeamPostBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamPostBodyWithDefaults

`func NewTeamPostBodyWithDefaults() *TeamPostBody`

NewTeamPostBodyWithDefaults instantiates a new TeamPostBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentTeamId

`func (o *TeamPostBody) GetParentTeamId() string`

GetParentTeamId returns the ParentTeamId field if non-nil, zero value otherwise.

### GetParentTeamIdOk

`func (o *TeamPostBody) GetParentTeamIdOk() (*string, bool)`

GetParentTeamIdOk returns a tuple with the ParentTeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTeamId

`func (o *TeamPostBody) SetParentTeamId(v string)`

SetParentTeamId sets ParentTeamId field to given value.

### HasParentTeamId

`func (o *TeamPostBody) HasParentTeamId() bool`

HasParentTeamId returns a boolean if a field has been set.

### GetTeamName

`func (o *TeamPostBody) GetTeamName() string`

GetTeamName returns the TeamName field if non-nil, zero value otherwise.

### GetTeamNameOk

`func (o *TeamPostBody) GetTeamNameOk() (*string, bool)`

GetTeamNameOk returns a tuple with the TeamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamName

`func (o *TeamPostBody) SetTeamName(v string)`

SetTeamName sets TeamName field to given value.

### HasTeamName

`func (o *TeamPostBody) HasTeamName() bool`

HasTeamName returns a boolean if a field has been set.

### GetTeamType

`func (o *TeamPostBody) GetTeamType() string`

GetTeamType returns the TeamType field if non-nil, zero value otherwise.

### GetTeamTypeOk

`func (o *TeamPostBody) GetTeamTypeOk() (*string, bool)`

GetTeamTypeOk returns a tuple with the TeamType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamType

`func (o *TeamPostBody) SetTeamType(v string)`

SetTeamType sets TeamType field to given value.

### HasTeamType

`func (o *TeamPostBody) HasTeamType() bool`

HasTeamType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


