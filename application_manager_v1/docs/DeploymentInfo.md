# DeploymentInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** |  | [optional] 
**MuleVersion** | Pointer to [**DeploymentInfoMuleVersion**](DeploymentInfoMuleVersion.md) |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**MonitoringEnabled** | Pointer to **bool** |  | [optional] 
**MonitoringAutoRestart** | Pointer to **bool** |  | [optional] 
**Workers** | Pointer to [**DeploymentInfoWorkers**](DeploymentInfoWorkers.md) |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** | Application properties | [optional] 
**PropertiesOptions** | Pointer to [**map[string]DeploymentInfoPropertiesOptionsValue**](DeploymentInfoPropertiesOptionsValue.md) | Additional option on the properties, like security configuratoin for each prop | [optional] 
**LogLevels** | Pointer to [**[]DeploymentInfoLogLevelsInner**](DeploymentInfoLogLevelsInner.md) |  | [optional] 
**SecureDataGatewayEnabled** | Pointer to **bool** |  | [optional] 
**StaticIPsEnabled** | Pointer to **bool** |  | [optional] 
**IpAddresses** | Pointer to **[]string** |  | [optional] 
**LoggingCustomLog4JEnabled** | Pointer to **bool** |  | [optional] 
**LoggingNgEnabled** | Pointer to **bool** |  | [optional] 
**PersistentQueues** | Pointer to **bool** |  | [optional] 
**PersistentQueuesEncrypted** | Pointer to **bool** |  | [optional] 
**ObjectStoreV1** | Pointer to **bool** | if set to false, Object Store v2 will be used. | [optional] 
**Runtime** | Pointer to [**DeploymentInfoRuntime**](DeploymentInfoRuntime.md) |  | [optional] 

## Methods

### NewDeploymentInfo

`func NewDeploymentInfo() *DeploymentInfo`

NewDeploymentInfo instantiates a new DeploymentInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentInfoWithDefaults

`func NewDeploymentInfoWithDefaults() *DeploymentInfo`

NewDeploymentInfoWithDefaults instantiates a new DeploymentInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *DeploymentInfo) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DeploymentInfo) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DeploymentInfo) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *DeploymentInfo) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetMuleVersion

`func (o *DeploymentInfo) GetMuleVersion() DeploymentInfoMuleVersion`

GetMuleVersion returns the MuleVersion field if non-nil, zero value otherwise.

### GetMuleVersionOk

`func (o *DeploymentInfo) GetMuleVersionOk() (*DeploymentInfoMuleVersion, bool)`

GetMuleVersionOk returns a tuple with the MuleVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuleVersion

`func (o *DeploymentInfo) SetMuleVersion(v DeploymentInfoMuleVersion)`

SetMuleVersion sets MuleVersion field to given value.

### HasMuleVersion

`func (o *DeploymentInfo) HasMuleVersion() bool`

HasMuleVersion returns a boolean if a field has been set.

### GetRegion

`func (o *DeploymentInfo) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *DeploymentInfo) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *DeploymentInfo) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *DeploymentInfo) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetMonitoringEnabled

`func (o *DeploymentInfo) GetMonitoringEnabled() bool`

GetMonitoringEnabled returns the MonitoringEnabled field if non-nil, zero value otherwise.

### GetMonitoringEnabledOk

`func (o *DeploymentInfo) GetMonitoringEnabledOk() (*bool, bool)`

GetMonitoringEnabledOk returns a tuple with the MonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringEnabled

`func (o *DeploymentInfo) SetMonitoringEnabled(v bool)`

SetMonitoringEnabled sets MonitoringEnabled field to given value.

### HasMonitoringEnabled

`func (o *DeploymentInfo) HasMonitoringEnabled() bool`

HasMonitoringEnabled returns a boolean if a field has been set.

### GetMonitoringAutoRestart

`func (o *DeploymentInfo) GetMonitoringAutoRestart() bool`

GetMonitoringAutoRestart returns the MonitoringAutoRestart field if non-nil, zero value otherwise.

### GetMonitoringAutoRestartOk

`func (o *DeploymentInfo) GetMonitoringAutoRestartOk() (*bool, bool)`

GetMonitoringAutoRestartOk returns a tuple with the MonitoringAutoRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringAutoRestart

`func (o *DeploymentInfo) SetMonitoringAutoRestart(v bool)`

SetMonitoringAutoRestart sets MonitoringAutoRestart field to given value.

### HasMonitoringAutoRestart

`func (o *DeploymentInfo) HasMonitoringAutoRestart() bool`

HasMonitoringAutoRestart returns a boolean if a field has been set.

### GetWorkers

`func (o *DeploymentInfo) GetWorkers() DeploymentInfoWorkers`

GetWorkers returns the Workers field if non-nil, zero value otherwise.

### GetWorkersOk

`func (o *DeploymentInfo) GetWorkersOk() (*DeploymentInfoWorkers, bool)`

GetWorkersOk returns a tuple with the Workers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkers

`func (o *DeploymentInfo) SetWorkers(v DeploymentInfoWorkers)`

SetWorkers sets Workers field to given value.

### HasWorkers

`func (o *DeploymentInfo) HasWorkers() bool`

HasWorkers returns a boolean if a field has been set.

### GetProperties

`func (o *DeploymentInfo) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *DeploymentInfo) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *DeploymentInfo) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *DeploymentInfo) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetPropertiesOptions

`func (o *DeploymentInfo) GetPropertiesOptions() map[string]DeploymentInfoPropertiesOptionsValue`

GetPropertiesOptions returns the PropertiesOptions field if non-nil, zero value otherwise.

### GetPropertiesOptionsOk

`func (o *DeploymentInfo) GetPropertiesOptionsOk() (*map[string]DeploymentInfoPropertiesOptionsValue, bool)`

GetPropertiesOptionsOk returns a tuple with the PropertiesOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertiesOptions

`func (o *DeploymentInfo) SetPropertiesOptions(v map[string]DeploymentInfoPropertiesOptionsValue)`

SetPropertiesOptions sets PropertiesOptions field to given value.

### HasPropertiesOptions

`func (o *DeploymentInfo) HasPropertiesOptions() bool`

HasPropertiesOptions returns a boolean if a field has been set.

### GetLogLevels

`func (o *DeploymentInfo) GetLogLevels() []DeploymentInfoLogLevelsInner`

GetLogLevels returns the LogLevels field if non-nil, zero value otherwise.

### GetLogLevelsOk

`func (o *DeploymentInfo) GetLogLevelsOk() (*[]DeploymentInfoLogLevelsInner, bool)`

GetLogLevelsOk returns a tuple with the LogLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevels

`func (o *DeploymentInfo) SetLogLevels(v []DeploymentInfoLogLevelsInner)`

SetLogLevels sets LogLevels field to given value.

### HasLogLevels

`func (o *DeploymentInfo) HasLogLevels() bool`

HasLogLevels returns a boolean if a field has been set.

### GetSecureDataGatewayEnabled

`func (o *DeploymentInfo) GetSecureDataGatewayEnabled() bool`

GetSecureDataGatewayEnabled returns the SecureDataGatewayEnabled field if non-nil, zero value otherwise.

### GetSecureDataGatewayEnabledOk

`func (o *DeploymentInfo) GetSecureDataGatewayEnabledOk() (*bool, bool)`

GetSecureDataGatewayEnabledOk returns a tuple with the SecureDataGatewayEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureDataGatewayEnabled

`func (o *DeploymentInfo) SetSecureDataGatewayEnabled(v bool)`

SetSecureDataGatewayEnabled sets SecureDataGatewayEnabled field to given value.

### HasSecureDataGatewayEnabled

`func (o *DeploymentInfo) HasSecureDataGatewayEnabled() bool`

HasSecureDataGatewayEnabled returns a boolean if a field has been set.

### GetStaticIPsEnabled

`func (o *DeploymentInfo) GetStaticIPsEnabled() bool`

GetStaticIPsEnabled returns the StaticIPsEnabled field if non-nil, zero value otherwise.

### GetStaticIPsEnabledOk

`func (o *DeploymentInfo) GetStaticIPsEnabledOk() (*bool, bool)`

GetStaticIPsEnabledOk returns a tuple with the StaticIPsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticIPsEnabled

`func (o *DeploymentInfo) SetStaticIPsEnabled(v bool)`

SetStaticIPsEnabled sets StaticIPsEnabled field to given value.

### HasStaticIPsEnabled

`func (o *DeploymentInfo) HasStaticIPsEnabled() bool`

HasStaticIPsEnabled returns a boolean if a field has been set.

### GetIpAddresses

`func (o *DeploymentInfo) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *DeploymentInfo) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *DeploymentInfo) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *DeploymentInfo) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### GetLoggingCustomLog4JEnabled

`func (o *DeploymentInfo) GetLoggingCustomLog4JEnabled() bool`

GetLoggingCustomLog4JEnabled returns the LoggingCustomLog4JEnabled field if non-nil, zero value otherwise.

### GetLoggingCustomLog4JEnabledOk

`func (o *DeploymentInfo) GetLoggingCustomLog4JEnabledOk() (*bool, bool)`

GetLoggingCustomLog4JEnabledOk returns a tuple with the LoggingCustomLog4JEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingCustomLog4JEnabled

`func (o *DeploymentInfo) SetLoggingCustomLog4JEnabled(v bool)`

SetLoggingCustomLog4JEnabled sets LoggingCustomLog4JEnabled field to given value.

### HasLoggingCustomLog4JEnabled

`func (o *DeploymentInfo) HasLoggingCustomLog4JEnabled() bool`

HasLoggingCustomLog4JEnabled returns a boolean if a field has been set.

### GetLoggingNgEnabled

`func (o *DeploymentInfo) GetLoggingNgEnabled() bool`

GetLoggingNgEnabled returns the LoggingNgEnabled field if non-nil, zero value otherwise.

### GetLoggingNgEnabledOk

`func (o *DeploymentInfo) GetLoggingNgEnabledOk() (*bool, bool)`

GetLoggingNgEnabledOk returns a tuple with the LoggingNgEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingNgEnabled

`func (o *DeploymentInfo) SetLoggingNgEnabled(v bool)`

SetLoggingNgEnabled sets LoggingNgEnabled field to given value.

### HasLoggingNgEnabled

`func (o *DeploymentInfo) HasLoggingNgEnabled() bool`

HasLoggingNgEnabled returns a boolean if a field has been set.

### GetPersistentQueues

`func (o *DeploymentInfo) GetPersistentQueues() bool`

GetPersistentQueues returns the PersistentQueues field if non-nil, zero value otherwise.

### GetPersistentQueuesOk

`func (o *DeploymentInfo) GetPersistentQueuesOk() (*bool, bool)`

GetPersistentQueuesOk returns a tuple with the PersistentQueues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentQueues

`func (o *DeploymentInfo) SetPersistentQueues(v bool)`

SetPersistentQueues sets PersistentQueues field to given value.

### HasPersistentQueues

`func (o *DeploymentInfo) HasPersistentQueues() bool`

HasPersistentQueues returns a boolean if a field has been set.

### GetPersistentQueuesEncrypted

`func (o *DeploymentInfo) GetPersistentQueuesEncrypted() bool`

GetPersistentQueuesEncrypted returns the PersistentQueuesEncrypted field if non-nil, zero value otherwise.

### GetPersistentQueuesEncryptedOk

`func (o *DeploymentInfo) GetPersistentQueuesEncryptedOk() (*bool, bool)`

GetPersistentQueuesEncryptedOk returns a tuple with the PersistentQueuesEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentQueuesEncrypted

`func (o *DeploymentInfo) SetPersistentQueuesEncrypted(v bool)`

SetPersistentQueuesEncrypted sets PersistentQueuesEncrypted field to given value.

### HasPersistentQueuesEncrypted

`func (o *DeploymentInfo) HasPersistentQueuesEncrypted() bool`

HasPersistentQueuesEncrypted returns a boolean if a field has been set.

### GetObjectStoreV1

`func (o *DeploymentInfo) GetObjectStoreV1() bool`

GetObjectStoreV1 returns the ObjectStoreV1 field if non-nil, zero value otherwise.

### GetObjectStoreV1Ok

`func (o *DeploymentInfo) GetObjectStoreV1Ok() (*bool, bool)`

GetObjectStoreV1Ok returns a tuple with the ObjectStoreV1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectStoreV1

`func (o *DeploymentInfo) SetObjectStoreV1(v bool)`

SetObjectStoreV1 sets ObjectStoreV1 field to given value.

### HasObjectStoreV1

`func (o *DeploymentInfo) HasObjectStoreV1() bool`

HasObjectStoreV1 returns a boolean if a field has been set.

### GetRuntime

`func (o *DeploymentInfo) GetRuntime() DeploymentInfoRuntime`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *DeploymentInfo) GetRuntimeOk() (*DeploymentInfoRuntime, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *DeploymentInfo) SetRuntime(v DeploymentInfoRuntime)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *DeploymentInfo) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


