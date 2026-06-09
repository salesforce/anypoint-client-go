# IngressConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReadResponseTimeout** | Pointer to **int32** | The read response timeout of the ingress configuration. | [optional] 
**Logs** | Pointer to [**IngressConfigurationLogs**](IngressConfigurationLogs.md) |  | [optional] 
**Protocol** | Pointer to **string** | The protocol of the ingress configuration. | [optional] 
**Deployment** | Pointer to [**IngressConfigurationDeployment**](IngressConfigurationDeployment.md) |  | [optional] 

## Methods

### NewIngressConfiguration

`func NewIngressConfiguration() *IngressConfiguration`

NewIngressConfiguration instantiates a new IngressConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngressConfigurationWithDefaults

`func NewIngressConfigurationWithDefaults() *IngressConfiguration`

NewIngressConfigurationWithDefaults instantiates a new IngressConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReadResponseTimeout

`func (o *IngressConfiguration) GetReadResponseTimeout() int32`

GetReadResponseTimeout returns the ReadResponseTimeout field if non-nil, zero value otherwise.

### GetReadResponseTimeoutOk

`func (o *IngressConfiguration) GetReadResponseTimeoutOk() (*int32, bool)`

GetReadResponseTimeoutOk returns a tuple with the ReadResponseTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadResponseTimeout

`func (o *IngressConfiguration) SetReadResponseTimeout(v int32)`

SetReadResponseTimeout sets ReadResponseTimeout field to given value.

### HasReadResponseTimeout

`func (o *IngressConfiguration) HasReadResponseTimeout() bool`

HasReadResponseTimeout returns a boolean if a field has been set.

### GetLogs

`func (o *IngressConfiguration) GetLogs() IngressConfigurationLogs`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *IngressConfiguration) GetLogsOk() (*IngressConfigurationLogs, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *IngressConfiguration) SetLogs(v IngressConfigurationLogs)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *IngressConfiguration) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetProtocol

`func (o *IngressConfiguration) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *IngressConfiguration) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *IngressConfiguration) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *IngressConfiguration) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetDeployment

`func (o *IngressConfiguration) GetDeployment() IngressConfigurationDeployment`

GetDeployment returns the Deployment field if non-nil, zero value otherwise.

### GetDeploymentOk

`func (o *IngressConfiguration) GetDeploymentOk() (*IngressConfigurationDeployment, bool)`

GetDeploymentOk returns a tuple with the Deployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployment

`func (o *IngressConfiguration) SetDeployment(v IngressConfigurationDeployment)`

SetDeployment sets Deployment field to given value.

### HasDeployment

`func (o *IngressConfiguration) HasDeployment() bool`

HasDeployment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


