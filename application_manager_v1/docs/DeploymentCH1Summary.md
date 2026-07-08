# DeploymentCH1Summary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | Pointer to **string** |  | [optional] 
**VersionId** | Pointer to **string** |  | [optional] 
**LoggingNgEnabled** | Pointer to **bool** |  | [optional] 
**LoggingCustomLog4JEnabled** | Pointer to **bool** |  | [optional] 
**UpdateRuntimeConfig** | Pointer to **bool** |  | [optional] 
**DeploymentErrorMessage** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**FullDomain** | Pointer to **string** |  | [optional] 
**Properties** | Pointer to **map[string]string** |  | [optional] 
**PropertiesOptions** | Pointer to [**map[string]NewDeploymentStructPropertiesOptionsValue**](NewDeploymentStructPropertiesOptionsValue.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Workers** | Pointer to [**Workers**](Workers.md) |  | [optional] 
**LastUpdateTime** | Pointer to **int64** |  | [optional] 
**FileName** | Pointer to **string** |  | [optional] 
**PersistentQueues** | Pointer to **bool** |  | [optional] 
**PersistentQueuesEncryptionEnabled** | Pointer to **bool** |  | [optional] 
**PersistentQueuesEncrypted** | Pointer to **bool** |  | [optional] 
**MonitoringEnabled** | Pointer to **bool** |  | [optional] 
**MonitoringAutoRestart** | Pointer to **bool** |  | [optional] 
**StaticIPsEnabled** | Pointer to **bool** |  | [optional] 
**HasFile** | Pointer to **bool** |  | [optional] 
**SecureDataGatewayEnabled** | Pointer to **bool** |  | [optional] 
**DeploymentGroup** | Pointer to [**DeploymentGroup**](DeploymentGroup.md) |  | [optional] 
**TrackingSettings** | Pointer to [**TrackingSettings**](TrackingSettings.md) |  | [optional] 
**LogLevels** | Pointer to [**[]LogLevel**](LogLevel.md) |  | [optional] 
**IpAddresses** | Pointer to [**[]IpAddress**](IpAddress.md) |  | [optional] 
**MuleVersion** | Pointer to [**MuleVersion**](MuleVersion.md) |  | [optional] 
**PreviousMuleVersion** | Pointer to [**MuleVersion**](MuleVersion.md) |  | [optional] 

## Methods

### NewDeploymentCH1Summary

`func NewDeploymentCH1Summary() *DeploymentCH1Summary`

NewDeploymentCH1Summary instantiates a new DeploymentCH1Summary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentCH1SummaryWithDefaults

`func NewDeploymentCH1SummaryWithDefaults() *DeploymentCH1Summary`

NewDeploymentCH1SummaryWithDefaults instantiates a new DeploymentCH1Summary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *DeploymentCH1Summary) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *DeploymentCH1Summary) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *DeploymentCH1Summary) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *DeploymentCH1Summary) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetVersionId

`func (o *DeploymentCH1Summary) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *DeploymentCH1Summary) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *DeploymentCH1Summary) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *DeploymentCH1Summary) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetLoggingNgEnabled

`func (o *DeploymentCH1Summary) GetLoggingNgEnabled() bool`

GetLoggingNgEnabled returns the LoggingNgEnabled field if non-nil, zero value otherwise.

### GetLoggingNgEnabledOk

`func (o *DeploymentCH1Summary) GetLoggingNgEnabledOk() (*bool, bool)`

GetLoggingNgEnabledOk returns a tuple with the LoggingNgEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingNgEnabled

`func (o *DeploymentCH1Summary) SetLoggingNgEnabled(v bool)`

SetLoggingNgEnabled sets LoggingNgEnabled field to given value.

### HasLoggingNgEnabled

`func (o *DeploymentCH1Summary) HasLoggingNgEnabled() bool`

HasLoggingNgEnabled returns a boolean if a field has been set.

### GetLoggingCustomLog4JEnabled

`func (o *DeploymentCH1Summary) GetLoggingCustomLog4JEnabled() bool`

GetLoggingCustomLog4JEnabled returns the LoggingCustomLog4JEnabled field if non-nil, zero value otherwise.

### GetLoggingCustomLog4JEnabledOk

`func (o *DeploymentCH1Summary) GetLoggingCustomLog4JEnabledOk() (*bool, bool)`

GetLoggingCustomLog4JEnabledOk returns a tuple with the LoggingCustomLog4JEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingCustomLog4JEnabled

`func (o *DeploymentCH1Summary) SetLoggingCustomLog4JEnabled(v bool)`

SetLoggingCustomLog4JEnabled sets LoggingCustomLog4JEnabled field to given value.

### HasLoggingCustomLog4JEnabled

`func (o *DeploymentCH1Summary) HasLoggingCustomLog4JEnabled() bool`

HasLoggingCustomLog4JEnabled returns a boolean if a field has been set.

### GetUpdateRuntimeConfig

`func (o *DeploymentCH1Summary) GetUpdateRuntimeConfig() bool`

GetUpdateRuntimeConfig returns the UpdateRuntimeConfig field if non-nil, zero value otherwise.

### GetUpdateRuntimeConfigOk

`func (o *DeploymentCH1Summary) GetUpdateRuntimeConfigOk() (*bool, bool)`

GetUpdateRuntimeConfigOk returns a tuple with the UpdateRuntimeConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateRuntimeConfig

`func (o *DeploymentCH1Summary) SetUpdateRuntimeConfig(v bool)`

SetUpdateRuntimeConfig sets UpdateRuntimeConfig field to given value.

### HasUpdateRuntimeConfig

`func (o *DeploymentCH1Summary) HasUpdateRuntimeConfig() bool`

HasUpdateRuntimeConfig returns a boolean if a field has been set.

### GetDeploymentErrorMessage

`func (o *DeploymentCH1Summary) GetDeploymentErrorMessage() string`

GetDeploymentErrorMessage returns the DeploymentErrorMessage field if non-nil, zero value otherwise.

### GetDeploymentErrorMessageOk

`func (o *DeploymentCH1Summary) GetDeploymentErrorMessageOk() (*string, bool)`

GetDeploymentErrorMessageOk returns a tuple with the DeploymentErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentErrorMessage

`func (o *DeploymentCH1Summary) SetDeploymentErrorMessage(v string)`

SetDeploymentErrorMessage sets DeploymentErrorMessage field to given value.

### HasDeploymentErrorMessage

`func (o *DeploymentCH1Summary) HasDeploymentErrorMessage() bool`

HasDeploymentErrorMessage returns a boolean if a field has been set.

### GetDomain

`func (o *DeploymentCH1Summary) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DeploymentCH1Summary) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DeploymentCH1Summary) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *DeploymentCH1Summary) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetFullDomain

`func (o *DeploymentCH1Summary) GetFullDomain() string`

GetFullDomain returns the FullDomain field if non-nil, zero value otherwise.

### GetFullDomainOk

`func (o *DeploymentCH1Summary) GetFullDomainOk() (*string, bool)`

GetFullDomainOk returns a tuple with the FullDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullDomain

`func (o *DeploymentCH1Summary) SetFullDomain(v string)`

SetFullDomain sets FullDomain field to given value.

### HasFullDomain

`func (o *DeploymentCH1Summary) HasFullDomain() bool`

HasFullDomain returns a boolean if a field has been set.

### GetProperties

`func (o *DeploymentCH1Summary) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *DeploymentCH1Summary) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *DeploymentCH1Summary) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *DeploymentCH1Summary) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetPropertiesOptions

`func (o *DeploymentCH1Summary) GetPropertiesOptions() map[string]NewDeploymentStructPropertiesOptionsValue`

GetPropertiesOptions returns the PropertiesOptions field if non-nil, zero value otherwise.

### GetPropertiesOptionsOk

`func (o *DeploymentCH1Summary) GetPropertiesOptionsOk() (*map[string]NewDeploymentStructPropertiesOptionsValue, bool)`

GetPropertiesOptionsOk returns a tuple with the PropertiesOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertiesOptions

`func (o *DeploymentCH1Summary) SetPropertiesOptions(v map[string]NewDeploymentStructPropertiesOptionsValue)`

SetPropertiesOptions sets PropertiesOptions field to given value.

### HasPropertiesOptions

`func (o *DeploymentCH1Summary) HasPropertiesOptions() bool`

HasPropertiesOptions returns a boolean if a field has been set.

### GetStatus

`func (o *DeploymentCH1Summary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeploymentCH1Summary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeploymentCH1Summary) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeploymentCH1Summary) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWorkers

`func (o *DeploymentCH1Summary) GetWorkers() Workers`

GetWorkers returns the Workers field if non-nil, zero value otherwise.

### GetWorkersOk

`func (o *DeploymentCH1Summary) GetWorkersOk() (*Workers, bool)`

GetWorkersOk returns a tuple with the Workers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkers

`func (o *DeploymentCH1Summary) SetWorkers(v Workers)`

SetWorkers sets Workers field to given value.

### HasWorkers

`func (o *DeploymentCH1Summary) HasWorkers() bool`

HasWorkers returns a boolean if a field has been set.

### GetLastUpdateTime

`func (o *DeploymentCH1Summary) GetLastUpdateTime() int64`

GetLastUpdateTime returns the LastUpdateTime field if non-nil, zero value otherwise.

### GetLastUpdateTimeOk

`func (o *DeploymentCH1Summary) GetLastUpdateTimeOk() (*int64, bool)`

GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTime

`func (o *DeploymentCH1Summary) SetLastUpdateTime(v int64)`

SetLastUpdateTime sets LastUpdateTime field to given value.

### HasLastUpdateTime

`func (o *DeploymentCH1Summary) HasLastUpdateTime() bool`

HasLastUpdateTime returns a boolean if a field has been set.

### GetFileName

`func (o *DeploymentCH1Summary) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *DeploymentCH1Summary) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *DeploymentCH1Summary) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *DeploymentCH1Summary) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetPersistentQueues

`func (o *DeploymentCH1Summary) GetPersistentQueues() bool`

GetPersistentQueues returns the PersistentQueues field if non-nil, zero value otherwise.

### GetPersistentQueuesOk

`func (o *DeploymentCH1Summary) GetPersistentQueuesOk() (*bool, bool)`

GetPersistentQueuesOk returns a tuple with the PersistentQueues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentQueues

`func (o *DeploymentCH1Summary) SetPersistentQueues(v bool)`

SetPersistentQueues sets PersistentQueues field to given value.

### HasPersistentQueues

`func (o *DeploymentCH1Summary) HasPersistentQueues() bool`

HasPersistentQueues returns a boolean if a field has been set.

### GetPersistentQueuesEncryptionEnabled

`func (o *DeploymentCH1Summary) GetPersistentQueuesEncryptionEnabled() bool`

GetPersistentQueuesEncryptionEnabled returns the PersistentQueuesEncryptionEnabled field if non-nil, zero value otherwise.

### GetPersistentQueuesEncryptionEnabledOk

`func (o *DeploymentCH1Summary) GetPersistentQueuesEncryptionEnabledOk() (*bool, bool)`

GetPersistentQueuesEncryptionEnabledOk returns a tuple with the PersistentQueuesEncryptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentQueuesEncryptionEnabled

`func (o *DeploymentCH1Summary) SetPersistentQueuesEncryptionEnabled(v bool)`

SetPersistentQueuesEncryptionEnabled sets PersistentQueuesEncryptionEnabled field to given value.

### HasPersistentQueuesEncryptionEnabled

`func (o *DeploymentCH1Summary) HasPersistentQueuesEncryptionEnabled() bool`

HasPersistentQueuesEncryptionEnabled returns a boolean if a field has been set.

### GetPersistentQueuesEncrypted

`func (o *DeploymentCH1Summary) GetPersistentQueuesEncrypted() bool`

GetPersistentQueuesEncrypted returns the PersistentQueuesEncrypted field if non-nil, zero value otherwise.

### GetPersistentQueuesEncryptedOk

`func (o *DeploymentCH1Summary) GetPersistentQueuesEncryptedOk() (*bool, bool)`

GetPersistentQueuesEncryptedOk returns a tuple with the PersistentQueuesEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentQueuesEncrypted

`func (o *DeploymentCH1Summary) SetPersistentQueuesEncrypted(v bool)`

SetPersistentQueuesEncrypted sets PersistentQueuesEncrypted field to given value.

### HasPersistentQueuesEncrypted

`func (o *DeploymentCH1Summary) HasPersistentQueuesEncrypted() bool`

HasPersistentQueuesEncrypted returns a boolean if a field has been set.

### GetMonitoringEnabled

`func (o *DeploymentCH1Summary) GetMonitoringEnabled() bool`

GetMonitoringEnabled returns the MonitoringEnabled field if non-nil, zero value otherwise.

### GetMonitoringEnabledOk

`func (o *DeploymentCH1Summary) GetMonitoringEnabledOk() (*bool, bool)`

GetMonitoringEnabledOk returns a tuple with the MonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringEnabled

`func (o *DeploymentCH1Summary) SetMonitoringEnabled(v bool)`

SetMonitoringEnabled sets MonitoringEnabled field to given value.

### HasMonitoringEnabled

`func (o *DeploymentCH1Summary) HasMonitoringEnabled() bool`

HasMonitoringEnabled returns a boolean if a field has been set.

### GetMonitoringAutoRestart

`func (o *DeploymentCH1Summary) GetMonitoringAutoRestart() bool`

GetMonitoringAutoRestart returns the MonitoringAutoRestart field if non-nil, zero value otherwise.

### GetMonitoringAutoRestartOk

`func (o *DeploymentCH1Summary) GetMonitoringAutoRestartOk() (*bool, bool)`

GetMonitoringAutoRestartOk returns a tuple with the MonitoringAutoRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringAutoRestart

`func (o *DeploymentCH1Summary) SetMonitoringAutoRestart(v bool)`

SetMonitoringAutoRestart sets MonitoringAutoRestart field to given value.

### HasMonitoringAutoRestart

`func (o *DeploymentCH1Summary) HasMonitoringAutoRestart() bool`

HasMonitoringAutoRestart returns a boolean if a field has been set.

### GetStaticIPsEnabled

`func (o *DeploymentCH1Summary) GetStaticIPsEnabled() bool`

GetStaticIPsEnabled returns the StaticIPsEnabled field if non-nil, zero value otherwise.

### GetStaticIPsEnabledOk

`func (o *DeploymentCH1Summary) GetStaticIPsEnabledOk() (*bool, bool)`

GetStaticIPsEnabledOk returns a tuple with the StaticIPsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticIPsEnabled

`func (o *DeploymentCH1Summary) SetStaticIPsEnabled(v bool)`

SetStaticIPsEnabled sets StaticIPsEnabled field to given value.

### HasStaticIPsEnabled

`func (o *DeploymentCH1Summary) HasStaticIPsEnabled() bool`

HasStaticIPsEnabled returns a boolean if a field has been set.

### GetHasFile

`func (o *DeploymentCH1Summary) GetHasFile() bool`

GetHasFile returns the HasFile field if non-nil, zero value otherwise.

### GetHasFileOk

`func (o *DeploymentCH1Summary) GetHasFileOk() (*bool, bool)`

GetHasFileOk returns a tuple with the HasFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFile

`func (o *DeploymentCH1Summary) SetHasFile(v bool)`

SetHasFile sets HasFile field to given value.

### HasHasFile

`func (o *DeploymentCH1Summary) HasHasFile() bool`

HasHasFile returns a boolean if a field has been set.

### GetSecureDataGatewayEnabled

`func (o *DeploymentCH1Summary) GetSecureDataGatewayEnabled() bool`

GetSecureDataGatewayEnabled returns the SecureDataGatewayEnabled field if non-nil, zero value otherwise.

### GetSecureDataGatewayEnabledOk

`func (o *DeploymentCH1Summary) GetSecureDataGatewayEnabledOk() (*bool, bool)`

GetSecureDataGatewayEnabledOk returns a tuple with the SecureDataGatewayEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureDataGatewayEnabled

`func (o *DeploymentCH1Summary) SetSecureDataGatewayEnabled(v bool)`

SetSecureDataGatewayEnabled sets SecureDataGatewayEnabled field to given value.

### HasSecureDataGatewayEnabled

`func (o *DeploymentCH1Summary) HasSecureDataGatewayEnabled() bool`

HasSecureDataGatewayEnabled returns a boolean if a field has been set.

### GetDeploymentGroup

`func (o *DeploymentCH1Summary) GetDeploymentGroup() DeploymentGroup`

GetDeploymentGroup returns the DeploymentGroup field if non-nil, zero value otherwise.

### GetDeploymentGroupOk

`func (o *DeploymentCH1Summary) GetDeploymentGroupOk() (*DeploymentGroup, bool)`

GetDeploymentGroupOk returns a tuple with the DeploymentGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentGroup

`func (o *DeploymentCH1Summary) SetDeploymentGroup(v DeploymentGroup)`

SetDeploymentGroup sets DeploymentGroup field to given value.

### HasDeploymentGroup

`func (o *DeploymentCH1Summary) HasDeploymentGroup() bool`

HasDeploymentGroup returns a boolean if a field has been set.

### GetTrackingSettings

`func (o *DeploymentCH1Summary) GetTrackingSettings() TrackingSettings`

GetTrackingSettings returns the TrackingSettings field if non-nil, zero value otherwise.

### GetTrackingSettingsOk

`func (o *DeploymentCH1Summary) GetTrackingSettingsOk() (*TrackingSettings, bool)`

GetTrackingSettingsOk returns a tuple with the TrackingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingSettings

`func (o *DeploymentCH1Summary) SetTrackingSettings(v TrackingSettings)`

SetTrackingSettings sets TrackingSettings field to given value.

### HasTrackingSettings

`func (o *DeploymentCH1Summary) HasTrackingSettings() bool`

HasTrackingSettings returns a boolean if a field has been set.

### GetLogLevels

`func (o *DeploymentCH1Summary) GetLogLevels() []LogLevel`

GetLogLevels returns the LogLevels field if non-nil, zero value otherwise.

### GetLogLevelsOk

`func (o *DeploymentCH1Summary) GetLogLevelsOk() (*[]LogLevel, bool)`

GetLogLevelsOk returns a tuple with the LogLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevels

`func (o *DeploymentCH1Summary) SetLogLevels(v []LogLevel)`

SetLogLevels sets LogLevels field to given value.

### HasLogLevels

`func (o *DeploymentCH1Summary) HasLogLevels() bool`

HasLogLevels returns a boolean if a field has been set.

### GetIpAddresses

`func (o *DeploymentCH1Summary) GetIpAddresses() []IpAddress`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *DeploymentCH1Summary) GetIpAddressesOk() (*[]IpAddress, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *DeploymentCH1Summary) SetIpAddresses(v []IpAddress)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *DeploymentCH1Summary) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### GetMuleVersion

`func (o *DeploymentCH1Summary) GetMuleVersion() MuleVersion`

GetMuleVersion returns the MuleVersion field if non-nil, zero value otherwise.

### GetMuleVersionOk

`func (o *DeploymentCH1Summary) GetMuleVersionOk() (*MuleVersion, bool)`

GetMuleVersionOk returns a tuple with the MuleVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuleVersion

`func (o *DeploymentCH1Summary) SetMuleVersion(v MuleVersion)`

SetMuleVersion sets MuleVersion field to given value.

### HasMuleVersion

`func (o *DeploymentCH1Summary) HasMuleVersion() bool`

HasMuleVersion returns a boolean if a field has been set.

### GetPreviousMuleVersion

`func (o *DeploymentCH1Summary) GetPreviousMuleVersion() MuleVersion`

GetPreviousMuleVersion returns the PreviousMuleVersion field if non-nil, zero value otherwise.

### GetPreviousMuleVersionOk

`func (o *DeploymentCH1Summary) GetPreviousMuleVersionOk() (*MuleVersion, bool)`

GetPreviousMuleVersionOk returns a tuple with the PreviousMuleVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousMuleVersion

`func (o *DeploymentCH1Summary) SetPreviousMuleVersion(v MuleVersion)`

SetPreviousMuleVersion sets PreviousMuleVersion field to given value.

### HasPreviousMuleVersion

`func (o *DeploymentCH1Summary) HasPreviousMuleVersion() bool`

HasPreviousMuleVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


