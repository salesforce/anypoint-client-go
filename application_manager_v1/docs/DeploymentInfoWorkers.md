# DeploymentInfoWorkers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | The number or replicas | [optional] 
**Type** | Pointer to [**DeploymentInfoWorkersType**](DeploymentInfoWorkersType.md) |  | [optional] 

## Methods

### NewDeploymentInfoWorkers

`func NewDeploymentInfoWorkers() *DeploymentInfoWorkers`

NewDeploymentInfoWorkers instantiates a new DeploymentInfoWorkers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentInfoWorkersWithDefaults

`func NewDeploymentInfoWorkersWithDefaults() *DeploymentInfoWorkers`

NewDeploymentInfoWorkersWithDefaults instantiates a new DeploymentInfoWorkers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *DeploymentInfoWorkers) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DeploymentInfoWorkers) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DeploymentInfoWorkers) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *DeploymentInfoWorkers) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetType

`func (o *DeploymentInfoWorkers) GetType() DeploymentInfoWorkersType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeploymentInfoWorkers) GetTypeOk() (*DeploymentInfoWorkersType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeploymentInfoWorkers) SetType(v DeploymentInfoWorkersType)`

SetType sets Type field to given value.

### HasType

`func (o *DeploymentInfoWorkers) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


