# TeamCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Team**](Team.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewTeamCollection

`func NewTeamCollection() *TeamCollection`

NewTeamCollection instantiates a new TeamCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamCollectionWithDefaults

`func NewTeamCollectionWithDefaults() *TeamCollection`

NewTeamCollectionWithDefaults instantiates a new TeamCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TeamCollection) GetData() []Team`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TeamCollection) GetDataOk() (*[]Team, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TeamCollection) SetData(v []Team)`

SetData sets Data field to given value.

### HasData

`func (o *TeamCollection) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotal

`func (o *TeamCollection) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *TeamCollection) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *TeamCollection) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *TeamCollection) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


