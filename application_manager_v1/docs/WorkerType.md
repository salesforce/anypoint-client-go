# WorkerType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the worker type. | [optional] 
**Weight** | Pointer to **float32** | The weight factor for the worker type. | [optional] 
**Cpu** | Pointer to **string** | The CPU allocation description (e.g., &#39;1 vCore&#39;). | [optional] 
**Memory** | Pointer to **string** | The memory allocation description (e.g., &#39;1.5 GB memory&#39;). | [optional] 

## Methods

### NewWorkerType

`func NewWorkerType() *WorkerType`

NewWorkerType instantiates a new WorkerType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkerTypeWithDefaults

`func NewWorkerTypeWithDefaults() *WorkerType`

NewWorkerTypeWithDefaults instantiates a new WorkerType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WorkerType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkerType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkerType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkerType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetWeight

`func (o *WorkerType) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *WorkerType) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *WorkerType) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *WorkerType) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetCpu

`func (o *WorkerType) GetCpu() string`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *WorkerType) GetCpuOk() (*string, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *WorkerType) SetCpu(v string)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *WorkerType) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMemory

`func (o *WorkerType) GetMemory() string`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *WorkerType) GetMemoryOk() (*string, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *WorkerType) SetMemory(v string)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *WorkerType) HasMemory() bool`

HasMemory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


