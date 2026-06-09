# DeploymentDetailsWorkersType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the worker type. | [optional] 
**Weight** | Pointer to **float32** | The weight factor for the worker type. | [optional] 
**Cpu** | Pointer to **string** | The CPU allocation description (e.g., &#39;1 vCore&#39;). | [optional] 
**Memory** | Pointer to **string** | The memory allocation description (e.g., &#39;1.5 GB memory&#39;). | [optional] 

## Methods

### NewDeploymentDetailsWorkersType

`func NewDeploymentDetailsWorkersType() *DeploymentDetailsWorkersType`

NewDeploymentDetailsWorkersType instantiates a new DeploymentDetailsWorkersType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentDetailsWorkersTypeWithDefaults

`func NewDeploymentDetailsWorkersTypeWithDefaults() *DeploymentDetailsWorkersType`

NewDeploymentDetailsWorkersTypeWithDefaults instantiates a new DeploymentDetailsWorkersType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeploymentDetailsWorkersType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentDetailsWorkersType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentDetailsWorkersType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeploymentDetailsWorkersType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetWeight

`func (o *DeploymentDetailsWorkersType) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *DeploymentDetailsWorkersType) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *DeploymentDetailsWorkersType) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *DeploymentDetailsWorkersType) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetCpu

`func (o *DeploymentDetailsWorkersType) GetCpu() string`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *DeploymentDetailsWorkersType) GetCpuOk() (*string, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *DeploymentDetailsWorkersType) SetCpu(v string)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *DeploymentDetailsWorkersType) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMemory

`func (o *DeploymentDetailsWorkersType) GetMemory() string`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *DeploymentDetailsWorkersType) GetMemoryOk() (*string, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *DeploymentDetailsWorkersType) SetMemory(v string)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *DeploymentDetailsWorkersType) HasMemory() bool`

HasMemory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


