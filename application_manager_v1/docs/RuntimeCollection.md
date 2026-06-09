# RuntimeCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Runtime**](Runtime.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewRuntimeCollection

`func NewRuntimeCollection() *RuntimeCollection`

NewRuntimeCollection instantiates a new RuntimeCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuntimeCollectionWithDefaults

`func NewRuntimeCollectionWithDefaults() *RuntimeCollection`

NewRuntimeCollectionWithDefaults instantiates a new RuntimeCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RuntimeCollection) GetData() []Runtime`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RuntimeCollection) GetDataOk() (*[]Runtime, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RuntimeCollection) SetData(v []Runtime)`

SetData sets Data field to given value.

### HasData

`func (o *RuntimeCollection) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotal

`func (o *RuntimeCollection) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *RuntimeCollection) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *RuntimeCollection) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *RuntimeCollection) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


