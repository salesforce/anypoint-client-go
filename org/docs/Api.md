# Api

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Production** | Pointer to [**Production**](Production.md) |  | [optional] 
**Sandbox** | Pointer to [**Sandbox**](Sandbox.md) |  | [optional] 

## Methods

### NewApi

`func NewApi() *Api`

NewApi instantiates a new Api object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiWithDefaults

`func NewApiWithDefaults() *Api`

NewApiWithDefaults instantiates a new Api object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduction

`func (o *Api) GetProduction() Production`

GetProduction returns the Production field if non-nil, zero value otherwise.

### GetProductionOk

`func (o *Api) GetProductionOk() (*Production, bool)`

GetProductionOk returns a tuple with the Production field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduction

`func (o *Api) SetProduction(v Production)`

SetProduction sets Production field to given value.

### HasProduction

`func (o *Api) HasProduction() bool`

HasProduction returns a boolean if a field has been set.

### GetSandbox

`func (o *Api) GetSandbox() Sandbox`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *Api) GetSandboxOk() (*Sandbox, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandbox

`func (o *Api) SetSandbox(v Sandbox)`

SetSandbox sets Sandbox field to given value.

### HasSandbox

`func (o *Api) HasSandbox() bool`

HasSandbox returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


