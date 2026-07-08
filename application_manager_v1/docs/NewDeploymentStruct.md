# NewDeploymentStruct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** |  | [optional] 
**MuleVersion** | Pointer to [**NewDeploymentStructMuleVersion**](NewDeploymentStructMuleVersion.md) |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**MonitoringEnabled** | Pointer to **bool** |  | [optional] 
**MonitoringAutoRestart** | Pointer to **bool** |  | [optional] 
**Workers** | Pointer to [**NewDeploymentStructWorkers**](NewDeploymentStructWorkers.md) |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** | Application properties | [optional] 
**PropertiesOptions** | Pointer to [**map[string]NewDeploymentStructPropertiesOptionsValue**](NewDeploymentStructPropertiesOptionsValue.md) | Additional option on the properties, like security configuratoin for each prop | [optional] 
**LogLevels** | Pointer to [**[]NewDeploymentStructLogLevelsInner**](NewDeploymentStructLogLevelsInner.md) |  | [optional] 
**SecureDataGatewayEnabled** | Pointer to **bool** |  | [optional] 
**StaticIPsEnabled** | Pointer to **bool** |  | [optional] 
**IpAddresses** | Pointer to **[]string** |  | [optional] 
**LoggingCustomLog4JEnabled** | Pointer to **bool** |  | [optional] 
**LoggingNgEnabled** | Pointer to **bool** |  | [optional] 
**PersistentQueues** | Pointer to **bool** |  | [optional] 
**PersistentQueuesEncrypted** | Pointer to **bool** |  | [optional] 
**ObjectStoreV1** | Pointer to **bool** | if set to false, Object Store v2 will be used. | [optional] 
**Runtime** | Pointer to [**NewDeploymentStructRuntime**](NewDeploymentStructRuntime.md) |  | [optional] 

## Methods

### NewNewDeploymentStruct

`func NewNewDeploymentStruct() *NewDeploymentStruct`

NewNewDeploymentStruct instantiates a new NewDeploymentStruct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewDeploymentStructWithDefaults

`func NewNewDeploymentStructWithDefaults() *NewDeploymentStruct`

NewNewDeploymentStructWithDefaults instantiates a new NewDeploymentStruct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *NewDeploymentStruct) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *NewDeploymentStruct) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *NewDeploymentStruct) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *NewDeploymentStruct) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetMuleVersion

`func (o *NewDeploymentStruct) GetMuleVersion() NewDeploymentStructMuleVersion`

GetMuleVersion returns the MuleVersion field if non-nil, zero value otherwise.

### GetMuleVersionOk

`func (o *NewDeploymentStruct) GetMuleVersionOk() (*NewDeploymentStructMuleVersion, bool)`

GetMuleVersionOk returns a tuple with the MuleVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuleVersion

`func (o *NewDeploymentStruct) SetMuleVersion(v NewDeploymentStructMuleVersion)`

SetMuleVersion sets MuleVersion field to given value.

### HasMuleVersion

`func (o *NewDeploymentStruct) HasMuleVersion() bool`

HasMuleVersion returns a boolean if a field has been set.

### GetRegion

`func (o *NewDeploymentStruct) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *NewDeploymentStruct) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *NewDeploymentStruct) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *NewDeploymentStruct) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetMonitoringEnabled

`func (o *NewDeploymentStruct) GetMonitoringEnabled() bool`

GetMonitoringEnabled returns the MonitoringEnabled field if non-nil, zero value otherwise.

### GetMonitoringEnabledOk

`func (o *NewDeploymentStruct) GetMonitoringEnabledOk() (*bool, bool)`

GetMonitoringEnabledOk returns a tuple with the MonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringEnabled

`func (o *NewDeploymentStruct) SetMonitoringEnabled(v bool)`

SetMonitoringEnabled sets MonitoringEnabled field to given value.

### HasMonitoringEnabled

`func (o *NewDeploymentStruct) HasMonitoringEnabled() bool`

HasMonitoringEnabled returns a boolean if a field has been set.

### GetMonitoringAutoRestart

`func (o *NewDeploymentStruct) GetMonitoringAutoRestart() bool`

GetMonitoringAutoRestart returns the MonitoringAutoRestart field if non-nil, zero value otherwise.

### GetMonitoringAutoRestartOk

`func (o *NewDeploymentStruct) GetMonitoringAutoRestartOk() (*bool, bool)`

GetMonitoringAutoRestartOk returns a tuple with the MonitoringAutoRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringAutoRestart

`func (o *NewDeploymentStruct) SetMonitoringAutoRestart(v bool)`

SetMonitoringAutoRestart sets MonitoringAutoRestart field to given value.

### HasMonitoringAutoRestart

`func (o *NewDeploymentStruct) HasMonitoringAutoRestart() bool`

HasMonitoringAutoRestart returns a boolean if a field has been set.

### GetWorkers

`func (o *NewDeploymentStruct) GetWorkers() NewDeploymentStructWorkers`

GetWorkers returns the Workers field if non-nil, zero value otherwise.

### GetWorkersOk

`func (o *NewDeploymentStruct) GetWorkersOk() (*NewDeploymentStructWorkers, bool)`

GetWorkersOk returns a tuple with the Workers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkers

`func (o *NewDeploymentStruct) SetWorkers(v NewDeploymentStructWorkers)`

SetWorkers sets Workers field to given value.

### HasWorkers

`func (o *NewDeploymentStruct) HasWorkers() bool`

HasWorkers returns a boolean if a field has been set.

### GetProperties

`func (o *NewDeploymentStruct) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *NewDeploymentStruct) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *NewDeploymentStruct) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *NewDeploymentStruct) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetPropertiesOptions

`func (o *NewDeploymentStruct) GetPropertiesOptions() map[string]NewDeploymentStructPropertiesOptionsValue`

GetPropertiesOptions returns the PropertiesOptions field if non-nil, zero value otherwise.

### GetPropertiesOptionsOk

`func (o *NewDeploymentStruct) GetPropertiesOptionsOk() (*map[string]NewDeploymentStructPropertiesOptionsValue, bool)`

GetPropertiesOptionsOk returns a tuple with the PropertiesOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertiesOptions

`func (o *NewDeploymentStruct) SetPropertiesOptions(v map[string]NewDeploymentStructPropertiesOptionsValue)`

SetPropertiesOptions sets PropertiesOptions field to given value.

### HasPropertiesOptions

`func (o *NewDeploymentStruct) HasPropertiesOptions() bool`

HasPropertiesOptions returns a boolean if a field has been set.

### GetLogLevels

`func (o *NewDeploymentStruct) GetLogLevels() []NewDeploymentStructLogLevelsInner`

GetLogLevels returns the LogLevels field if non-nil, zero value otherwise.

### GetLogLevelsOk

`func (o *NewDeploymentStruct) GetLogLevelsOk() (*[]NewDeploymentStructLogLevelsInner, bool)`

GetLogLevelsOk returns a tuple with the LogLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevels

`func (o *NewDeploymentStruct) SetLogLevels(v []NewDeploymentStructLogLevelsInner)`

SetLogLevels sets LogLevels field to given value.

### HasLogLevels

`func (o *NewDeploymentStruct) HasLogLevels() bool`

HasLogLevels returns a boolean if a field has been set.

### GetSecureDataGatewayEnabled

`func (o *NewDeploymentStruct) GetSecureDataGatewayEnabled() bool`

GetSecureDataGatewayEnabled returns the SecureDataGatewayEnabled field if non-nil, zero value otherwise.

### GetSecureDataGatewayEnabledOk

`func (o *NewDeploymentStruct) GetSecureDataGatewayEnabledOk() (*bool, bool)`

GetSecureDataGatewayEnabledOk returns a tuple with the SecureDataGatewayEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureDataGatewayEnabled

`func (o *NewDeploymentStruct) SetSecureDataGatewayEnabled(v bool)`

SetSecureDataGatewayEnabled sets SecureDataGatewayEnabled field to given value.

### HasSecureDataGatewayEnabled

`func (o *NewDeploymentStruct) HasSecureDataGatewayEnabled() bool`

HasSecureDataGatewayEnabled returns a boolean if a field has been set.

### GetStaticIPsEnabled

`func (o *NewDeploymentStruct) GetStaticIPsEnabled() bool`

GetStaticIPsEnabled returns the StaticIPsEnabled field if non-nil, zero value otherwise.

### GetStaticIPsEnabledOk

`func (o *NewDeploymentStruct) GetStaticIPsEnabledOk() (*bool, bool)`

GetStaticIPsEnabledOk returns a tuple with the StaticIPsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticIPsEnabled

`func (o *NewDeploymentStruct) SetStaticIPsEnabled(v bool)`

SetStaticIPsEnabled sets StaticIPsEnabled field to given value.

### HasStaticIPsEnabled

`func (o *NewDeploymentStruct) HasStaticIPsEnabled() bool`

HasStaticIPsEnabled returns a boolean if a field has been set.

### GetIpAddresses

`func (o *NewDeploymentStruct) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *NewDeploymentStruct) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *NewDeploymentStruct) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *NewDeploymentStruct) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### GetLoggingCustomLog4JEnabled

`func (o *NewDeploymentStruct) GetLoggingCustomLog4JEnabled() bool`

GetLoggingCustomLog4JEnabled returns the LoggingCustomLog4JEnabled field if non-nil, zero value otherwise.

### GetLoggingCustomLog4JEnabledOk

`func (o *NewDeploymentStruct) GetLoggingCustomLog4JEnabledOk() (*bool, bool)`

GetLoggingCustomLog4JEnabledOk returns a tuple with the LoggingCustomLog4JEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingCustomLog4JEnabled

`func (o *NewDeploymentStruct) SetLoggingCustomLog4JEnabled(v bool)`

SetLoggingCustomLog4JEnabled sets LoggingCustomLog4JEnabled field to given value.

### HasLoggingCustomLog4JEnabled

`func (o *NewDeploymentStruct) HasLoggingCustomLog4JEnabled() bool`

HasLoggingCustomLog4JEnabled returns a boolean if a field has been set.

### GetLoggingNgEnabled

`func (o *NewDeploymentStruct) GetLoggingNgEnabled() bool`

GetLoggingNgEnabled returns the LoggingNgEnabled field if non-nil, zero value otherwise.

### GetLoggingNgEnabledOk

`func (o *NewDeploymentStruct) GetLoggingNgEnabledOk() (*bool, bool)`

GetLoggingNgEnabledOk returns a tuple with the LoggingNgEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingNgEnabled

`func (o *NewDeploymentStruct) SetLoggingNgEnabled(v bool)`

SetLoggingNgEnabled sets LoggingNgEnabled field to given value.

### HasLoggingNgEnabled

`func (o *NewDeploymentStruct) HasLoggingNgEnabled() bool`

HasLoggingNgEnabled returns a boolean if a field has been set.

### GetPersistentQueues

`func (o *NewDeploymentStruct) GetPersistentQueues() bool`

GetPersistentQueues returns the PersistentQueues field if non-nil, zero value otherwise.

### GetPersistentQueuesOk

`func (o *NewDeploymentStruct) GetPersistentQueuesOk() (*bool, bool)`

GetPersistentQueuesOk returns a tuple with the PersistentQueues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentQueues

`func (o *NewDeploymentStruct) SetPersistentQueues(v bool)`

SetPersistentQueues sets PersistentQueues field to given value.

### HasPersistentQueues

`func (o *NewDeploymentStruct) HasPersistentQueues() bool`

HasPersistentQueues returns a boolean if a field has been set.

### GetPersistentQueuesEncrypted

`func (o *NewDeploymentStruct) GetPersistentQueuesEncrypted() bool`

GetPersistentQueuesEncrypted returns the PersistentQueuesEncrypted field if non-nil, zero value otherwise.

### GetPersistentQueuesEncryptedOk

`func (o *NewDeploymentStruct) GetPersistentQueuesEncryptedOk() (*bool, bool)`

GetPersistentQueuesEncryptedOk returns a tuple with the PersistentQueuesEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentQueuesEncrypted

`func (o *NewDeploymentStruct) SetPersistentQueuesEncrypted(v bool)`

SetPersistentQueuesEncrypted sets PersistentQueuesEncrypted field to given value.

### HasPersistentQueuesEncrypted

`func (o *NewDeploymentStruct) HasPersistentQueuesEncrypted() bool`

HasPersistentQueuesEncrypted returns a boolean if a field has been set.

### GetObjectStoreV1

`func (o *NewDeploymentStruct) GetObjectStoreV1() bool`

GetObjectStoreV1 returns the ObjectStoreV1 field if non-nil, zero value otherwise.

### GetObjectStoreV1Ok

`func (o *NewDeploymentStruct) GetObjectStoreV1Ok() (*bool, bool)`

GetObjectStoreV1Ok returns a tuple with the ObjectStoreV1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectStoreV1

`func (o *NewDeploymentStruct) SetObjectStoreV1(v bool)`

SetObjectStoreV1 sets ObjectStoreV1 field to given value.

### HasObjectStoreV1

`func (o *NewDeploymentStruct) HasObjectStoreV1() bool`

HasObjectStoreV1 returns a boolean if a field has been set.

### GetRuntime

`func (o *NewDeploymentStruct) GetRuntime() NewDeploymentStructRuntime`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *NewDeploymentStruct) GetRuntimeOk() (*NewDeploymentStructRuntime, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *NewDeploymentStruct) SetRuntime(v NewDeploymentStructRuntime)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *NewDeploymentStruct) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


