# HighAvailability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clustering** | Pointer to **bool** | Clustering | [optional] [default to false]

## Methods

### NewHighAvailability

`func NewHighAvailability() *HighAvailability`

NewHighAvailability instantiates a new HighAvailability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHighAvailabilityWithDefaults

`func NewHighAvailabilityWithDefaults() *HighAvailability`

NewHighAvailabilityWithDefaults instantiates a new HighAvailability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClustering

`func (o *HighAvailability) GetClustering() bool`

GetClustering returns the Clustering field if non-nil, zero value otherwise.

### GetClusteringOk

`func (o *HighAvailability) GetClusteringOk() (*bool, bool)`

GetClusteringOk returns a tuple with the Clustering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClustering

`func (o *HighAvailability) SetClustering(v bool)`

SetClustering sets Clustering field to given value.

### HasClustering

`func (o *HighAvailability) HasClustering() bool`

HasClustering returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


