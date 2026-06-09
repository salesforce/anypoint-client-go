# WorkersRecentStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transactions** | Pointer to **int32** | The number of transactions processed. | [optional] 
**Cpu** | Pointer to **int32** | The current CPU usage percentage. | [optional] 

## Methods

### NewWorkersRecentStatistics

`func NewWorkersRecentStatistics() *WorkersRecentStatistics`

NewWorkersRecentStatistics instantiates a new WorkersRecentStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkersRecentStatisticsWithDefaults

`func NewWorkersRecentStatisticsWithDefaults() *WorkersRecentStatistics`

NewWorkersRecentStatisticsWithDefaults instantiates a new WorkersRecentStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactions

`func (o *WorkersRecentStatistics) GetTransactions() int32`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *WorkersRecentStatistics) GetTransactionsOk() (*int32, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *WorkersRecentStatistics) SetTransactions(v int32)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *WorkersRecentStatistics) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.

### GetCpu

`func (o *WorkersRecentStatistics) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *WorkersRecentStatistics) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *WorkersRecentStatistics) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *WorkersRecentStatistics) HasCpu() bool`

HasCpu returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


