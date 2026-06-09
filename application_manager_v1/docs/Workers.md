# Workers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**WorkerType**](WorkerType.md) |  | [optional] 
**Amount** | Pointer to **int32** | The number of worker instances allocated. | [optional] 
**RemainingOrgWorkers** | Pointer to **float32** | Remaining worker count available to the organization. | [optional] 
**TotalOrgWorkers** | Pointer to **float32** | Total number of workers allocated to the organization. | [optional] 
**RecentStatistics** | Pointer to [**WorkersRecentStatistics**](WorkersRecentStatistics.md) |  | [optional] 

## Methods

### NewWorkers

`func NewWorkers() *Workers`

NewWorkers instantiates a new Workers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkersWithDefaults

`func NewWorkersWithDefaults() *Workers`

NewWorkersWithDefaults instantiates a new Workers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Workers) GetType() WorkerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Workers) GetTypeOk() (*WorkerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Workers) SetType(v WorkerType)`

SetType sets Type field to given value.

### HasType

`func (o *Workers) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAmount

`func (o *Workers) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Workers) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Workers) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Workers) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetRemainingOrgWorkers

`func (o *Workers) GetRemainingOrgWorkers() float32`

GetRemainingOrgWorkers returns the RemainingOrgWorkers field if non-nil, zero value otherwise.

### GetRemainingOrgWorkersOk

`func (o *Workers) GetRemainingOrgWorkersOk() (*float32, bool)`

GetRemainingOrgWorkersOk returns a tuple with the RemainingOrgWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingOrgWorkers

`func (o *Workers) SetRemainingOrgWorkers(v float32)`

SetRemainingOrgWorkers sets RemainingOrgWorkers field to given value.

### HasRemainingOrgWorkers

`func (o *Workers) HasRemainingOrgWorkers() bool`

HasRemainingOrgWorkers returns a boolean if a field has been set.

### GetTotalOrgWorkers

`func (o *Workers) GetTotalOrgWorkers() float32`

GetTotalOrgWorkers returns the TotalOrgWorkers field if non-nil, zero value otherwise.

### GetTotalOrgWorkersOk

`func (o *Workers) GetTotalOrgWorkersOk() (*float32, bool)`

GetTotalOrgWorkersOk returns a tuple with the TotalOrgWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOrgWorkers

`func (o *Workers) SetTotalOrgWorkers(v float32)`

SetTotalOrgWorkers sets TotalOrgWorkers field to given value.

### HasTotalOrgWorkers

`func (o *Workers) HasTotalOrgWorkers() bool`

HasTotalOrgWorkers returns a boolean if a field has been set.

### GetRecentStatistics

`func (o *Workers) GetRecentStatistics() WorkersRecentStatistics`

GetRecentStatistics returns the RecentStatistics field if non-nil, zero value otherwise.

### GetRecentStatisticsOk

`func (o *Workers) GetRecentStatisticsOk() (*WorkersRecentStatistics, bool)`

GetRecentStatisticsOk returns a tuple with the RecentStatistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentStatistics

`func (o *Workers) SetRecentStatistics(v WorkersRecentStatistics)`

SetRecentStatistics sets RecentStatistics field to given value.

### HasRecentStatistics

`func (o *Workers) HasRecentStatistics() bool`

HasRecentStatistics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


