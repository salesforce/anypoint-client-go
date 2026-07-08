# DeployAppBodyJSON

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationInfo** | Pointer to [**NewDeploymentStruct**](NewDeploymentStruct.md) |  | [optional] 
**ApplicationSource** | Pointer to [**ExchangeApplicationSource**](ExchangeApplicationSource.md) |  | [optional] 
**AutoStart** | Pointer to **bool** | Indicates whether the application should be automatically started after deployment. | [optional] 

## Methods

### NewDeployAppBodyJSON

`func NewDeployAppBodyJSON() *DeployAppBodyJSON`

NewDeployAppBodyJSON instantiates a new DeployAppBodyJSON object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployAppBodyJSONWithDefaults

`func NewDeployAppBodyJSONWithDefaults() *DeployAppBodyJSON`

NewDeployAppBodyJSONWithDefaults instantiates a new DeployAppBodyJSON object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationInfo

`func (o *DeployAppBodyJSON) GetApplicationInfo() NewDeploymentStruct`

GetApplicationInfo returns the ApplicationInfo field if non-nil, zero value otherwise.

### GetApplicationInfoOk

`func (o *DeployAppBodyJSON) GetApplicationInfoOk() (*NewDeploymentStruct, bool)`

GetApplicationInfoOk returns a tuple with the ApplicationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationInfo

`func (o *DeployAppBodyJSON) SetApplicationInfo(v NewDeploymentStruct)`

SetApplicationInfo sets ApplicationInfo field to given value.

### HasApplicationInfo

`func (o *DeployAppBodyJSON) HasApplicationInfo() bool`

HasApplicationInfo returns a boolean if a field has been set.

### GetApplicationSource

`func (o *DeployAppBodyJSON) GetApplicationSource() ExchangeApplicationSource`

GetApplicationSource returns the ApplicationSource field if non-nil, zero value otherwise.

### GetApplicationSourceOk

`func (o *DeployAppBodyJSON) GetApplicationSourceOk() (*ExchangeApplicationSource, bool)`

GetApplicationSourceOk returns a tuple with the ApplicationSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationSource

`func (o *DeployAppBodyJSON) SetApplicationSource(v ExchangeApplicationSource)`

SetApplicationSource sets ApplicationSource field to given value.

### HasApplicationSource

`func (o *DeployAppBodyJSON) HasApplicationSource() bool`

HasApplicationSource returns a boolean if a field has been set.

### GetAutoStart

`func (o *DeployAppBodyJSON) GetAutoStart() bool`

GetAutoStart returns the AutoStart field if non-nil, zero value otherwise.

### GetAutoStartOk

`func (o *DeployAppBodyJSON) GetAutoStartOk() (*bool, bool)`

GetAutoStartOk returns a tuple with the AutoStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoStart

`func (o *DeployAppBodyJSON) SetAutoStart(v bool)`

SetAutoStart sets AutoStart field to given value.

### HasAutoStart

`func (o *DeployAppBodyJSON) HasAutoStart() bool`

HasAutoStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


