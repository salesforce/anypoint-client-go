# ApimInstanceCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**Assets** | Pointer to [**[]ApimInstanceCollectionAssetsInner**](ApimInstanceCollectionAssetsInner.md) |  | [optional] 

## Methods

### NewApimInstanceCollection

`func NewApimInstanceCollection() *ApimInstanceCollection`

NewApimInstanceCollection instantiates a new ApimInstanceCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApimInstanceCollectionWithDefaults

`func NewApimInstanceCollectionWithDefaults() *ApimInstanceCollection`

NewApimInstanceCollectionWithDefaults instantiates a new ApimInstanceCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *ApimInstanceCollection) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ApimInstanceCollection) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ApimInstanceCollection) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ApimInstanceCollection) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetAssets

`func (o *ApimInstanceCollection) GetAssets() []ApimInstanceCollectionAssetsInner`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *ApimInstanceCollection) GetAssetsOk() (*[]ApimInstanceCollectionAssetsInner, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *ApimInstanceCollection) SetAssets(v []ApimInstanceCollectionAssetsInner)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *ApimInstanceCollection) HasAssets() bool`

HasAssets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


