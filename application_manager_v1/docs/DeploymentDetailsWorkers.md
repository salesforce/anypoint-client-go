# DeploymentDetailsWorkers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**DeploymentDetailsWorkersType**](DeploymentDetailsWorkersType.md) |  | [optional] 
**Amount** | Pointer to **int32** | The number of worker instances allocated. | [optional] 
**RemainingOrgWorkers** | Pointer to **int32** | Remaining worker count available to the organization. | [optional] 
**TotalOrgWorkers** | Pointer to **int32** | Total number of workers allocated to the organization. | [optional] 
**RecentStatistics** | Pointer to [**DeploymentDetailsWorkersRecentStatistics**](DeploymentDetailsWorkersRecentStatistics.md) |  | [optional] 

## Methods

### NewDeploymentDetailsWorkers

`func NewDeploymentDetailsWorkers() *DeploymentDetailsWorkers`

NewDeploymentDetailsWorkers instantiates a new DeploymentDetailsWorkers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentDetailsWorkersWithDefaults

`func NewDeploymentDetailsWorkersWithDefaults() *DeploymentDetailsWorkers`

NewDeploymentDetailsWorkersWithDefaults instantiates a new DeploymentDetailsWorkers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DeploymentDetailsWorkers) GetType() DeploymentDetailsWorkersType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeploymentDetailsWorkers) GetTypeOk() (*DeploymentDetailsWorkersType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeploymentDetailsWorkers) SetType(v DeploymentDetailsWorkersType)`

SetType sets Type field to given value.

### HasType

`func (o *DeploymentDetailsWorkers) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAmount

`func (o *DeploymentDetailsWorkers) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DeploymentDetailsWorkers) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DeploymentDetailsWorkers) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *DeploymentDetailsWorkers) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetRemainingOrgWorkers

`func (o *DeploymentDetailsWorkers) GetRemainingOrgWorkers() int32`

GetRemainingOrgWorkers returns the RemainingOrgWorkers field if non-nil, zero value otherwise.

### GetRemainingOrgWorkersOk

`func (o *DeploymentDetailsWorkers) GetRemainingOrgWorkersOk() (*int32, bool)`

GetRemainingOrgWorkersOk returns a tuple with the RemainingOrgWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingOrgWorkers

`func (o *DeploymentDetailsWorkers) SetRemainingOrgWorkers(v int32)`

SetRemainingOrgWorkers sets RemainingOrgWorkers field to given value.

### HasRemainingOrgWorkers

`func (o *DeploymentDetailsWorkers) HasRemainingOrgWorkers() bool`

HasRemainingOrgWorkers returns a boolean if a field has been set.

### GetTotalOrgWorkers

`func (o *DeploymentDetailsWorkers) GetTotalOrgWorkers() int32`

GetTotalOrgWorkers returns the TotalOrgWorkers field if non-nil, zero value otherwise.

### GetTotalOrgWorkersOk

`func (o *DeploymentDetailsWorkers) GetTotalOrgWorkersOk() (*int32, bool)`

GetTotalOrgWorkersOk returns a tuple with the TotalOrgWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOrgWorkers

`func (o *DeploymentDetailsWorkers) SetTotalOrgWorkers(v int32)`

SetTotalOrgWorkers sets TotalOrgWorkers field to given value.

### HasTotalOrgWorkers

`func (o *DeploymentDetailsWorkers) HasTotalOrgWorkers() bool`

HasTotalOrgWorkers returns a boolean if a field has been set.

### GetRecentStatistics

`func (o *DeploymentDetailsWorkers) GetRecentStatistics() DeploymentDetailsWorkersRecentStatistics`

GetRecentStatistics returns the RecentStatistics field if non-nil, zero value otherwise.

### GetRecentStatisticsOk

`func (o *DeploymentDetailsWorkers) GetRecentStatisticsOk() (*DeploymentDetailsWorkersRecentStatistics, bool)`

GetRecentStatisticsOk returns a tuple with the RecentStatistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentStatistics

`func (o *DeploymentDetailsWorkers) SetRecentStatistics(v DeploymentDetailsWorkersRecentStatistics)`

SetRecentStatistics sets RecentStatistics field to given value.

### HasRecentStatistics

`func (o *DeploymentDetailsWorkers) HasRecentStatistics() bool`

HasRecentStatistics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


