# UpdateDeploymentBodyJSON

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationInfo** | Pointer to [**NewDeploymentStruct**](NewDeploymentStruct.md) |  | [optional] 
**ApplicationSource** | Pointer to [**ExchangeApplicationSource**](ExchangeApplicationSource.md) |  | [optional] 

## Methods

### NewUpdateDeploymentBodyJSON

`func NewUpdateDeploymentBodyJSON() *UpdateDeploymentBodyJSON`

NewUpdateDeploymentBodyJSON instantiates a new UpdateDeploymentBodyJSON object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeploymentBodyJSONWithDefaults

`func NewUpdateDeploymentBodyJSONWithDefaults() *UpdateDeploymentBodyJSON`

NewUpdateDeploymentBodyJSONWithDefaults instantiates a new UpdateDeploymentBodyJSON object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationInfo

`func (o *UpdateDeploymentBodyJSON) GetApplicationInfo() NewDeploymentStruct`

GetApplicationInfo returns the ApplicationInfo field if non-nil, zero value otherwise.

### GetApplicationInfoOk

`func (o *UpdateDeploymentBodyJSON) GetApplicationInfoOk() (*NewDeploymentStruct, bool)`

GetApplicationInfoOk returns a tuple with the ApplicationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationInfo

`func (o *UpdateDeploymentBodyJSON) SetApplicationInfo(v NewDeploymentStruct)`

SetApplicationInfo sets ApplicationInfo field to given value.

### HasApplicationInfo

`func (o *UpdateDeploymentBodyJSON) HasApplicationInfo() bool`

HasApplicationInfo returns a boolean if a field has been set.

### GetApplicationSource

`func (o *UpdateDeploymentBodyJSON) GetApplicationSource() ExchangeApplicationSource`

GetApplicationSource returns the ApplicationSource field if non-nil, zero value otherwise.

### GetApplicationSourceOk

`func (o *UpdateDeploymentBodyJSON) GetApplicationSourceOk() (*ExchangeApplicationSource, bool)`

GetApplicationSourceOk returns a tuple with the ApplicationSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationSource

`func (o *UpdateDeploymentBodyJSON) SetApplicationSource(v ExchangeApplicationSource)`

SetApplicationSource sets ApplicationSource field to given value.

### HasApplicationSource

`func (o *UpdateDeploymentBodyJSON) HasApplicationSource() bool`

HasApplicationSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


