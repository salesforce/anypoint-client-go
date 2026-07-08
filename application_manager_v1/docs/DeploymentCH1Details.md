# DeploymentCH1Details

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** | The short domain name of the application. | [optional] 
**FullDomain** | Pointer to **string** | The full domain name including the CH1 suffix. | [optional] 
**Description** | Pointer to **string** | A description of the application. | [optional] 
**Properties** | Pointer to **map[string]interface{}** | Deployment properties specific to the app configuration. | [optional] 
**PropertiesOptions** | Pointer to **map[string]interface{}** | Options for the deployment properties, such as security flags. | [optional] 
**Status** | Pointer to **string** | Current status of the deployed application (e.g. STARTED). | [optional] 
**Workers** | Pointer to [**Workers**](Workers.md) |  | [optional] 
**LastUpdateTime** | Pointer to **int64** | Timestamp (in milliseconds) when the application was last updated. | [optional] 
**FileName** | Pointer to **string** | The name of the deployed application file. | [optional] 
**MuleVersion** | Pointer to [**MuleVersion**](MuleVersion.md) |  | [optional] 
**PreviousMuleVersion** | Pointer to [**MuleVersion**](MuleVersion.md) |  | [optional] 
**Tenants** | Pointer to **int32** | The number of tenants associated with the deployment. | [optional] 
**UserId** | Pointer to **string** | Identifier of the user who deployed the application. | [optional] 
**UserName** | Pointer to **string** | Username of the deploying user. | [optional] 
**VpnConfig** | Pointer to [**DeploymentCH1DetailsVpnConfig**](DeploymentCH1DetailsVpnConfig.md) |  | [optional] 
**DeploymentGroup** | Pointer to [**DeploymentGroup**](DeploymentGroup.md) |  | [optional] 
**PersistentQueues** | Pointer to **bool** | Indicates whether persistent queues are enabled. | [optional] 
**PersistentQueuesEncryptionEnabled** | Pointer to **bool** | Indicates whether encryption for persistent queues is enabled. | [optional] 
**PersistentQueuesEncrypted** | Pointer to **bool** | Indicates whether the persistent queues are encrypted. | [optional] 
**MonitoringEnabled** | Pointer to **bool** | Specifies if monitoring is enabled for the application. | [optional] 
**MonitoringAutoRestart** | Pointer to **bool** | Specifies if the application is set to auto-restart when monitoring triggers. | [optional] 
**StaticIPsEnabled** | Pointer to **bool** | Indicates whether static IPs are enabled for the deployment. | [optional] 
**Multitenanted** | Pointer to **bool** | Indicates if the deployment supports multiple tenants. | [optional] 
**HasFile** | Pointer to **bool** | Indicates whether a deployment file exists for the application. | [optional] 
**SecureDataGatewayEnabled** | Pointer to **bool** | Specifies if a secure data gateway is enabled for the deployment. | [optional] 
**VpnEnabled** | Pointer to **bool** | Indicates if VPN is enabled for the deployment. | [optional] 
**TrackingSettings** | Pointer to [**TrackingSettings**](TrackingSettings.md) |  | [optional] 
**LogLevels** | Pointer to [**[]LogLevel**](LogLevel.md) | List of log level configurations for the application. | [optional] 
**IpAddresses** | Pointer to [**[]IpAddress**](IpAddress.md) | List of IP addresses associated with the deployment. | [optional] 

## Methods

### NewDeploymentCH1Details

`func NewDeploymentCH1Details() *DeploymentCH1Details`

NewDeploymentCH1Details instantiates a new DeploymentCH1Details object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentCH1DetailsWithDefaults

`func NewDeploymentCH1DetailsWithDefaults() *DeploymentCH1Details`

NewDeploymentCH1DetailsWithDefaults instantiates a new DeploymentCH1Details object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *DeploymentCH1Details) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DeploymentCH1Details) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DeploymentCH1Details) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *DeploymentCH1Details) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetFullDomain

`func (o *DeploymentCH1Details) GetFullDomain() string`

GetFullDomain returns the FullDomain field if non-nil, zero value otherwise.

### GetFullDomainOk

`func (o *DeploymentCH1Details) GetFullDomainOk() (*string, bool)`

GetFullDomainOk returns a tuple with the FullDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullDomain

`func (o *DeploymentCH1Details) SetFullDomain(v string)`

SetFullDomain sets FullDomain field to given value.

### HasFullDomain

`func (o *DeploymentCH1Details) HasFullDomain() bool`

HasFullDomain returns a boolean if a field has been set.

### GetDescription

`func (o *DeploymentCH1Details) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeploymentCH1Details) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeploymentCH1Details) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeploymentCH1Details) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProperties

`func (o *DeploymentCH1Details) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *DeploymentCH1Details) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *DeploymentCH1Details) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *DeploymentCH1Details) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetPropertiesOptions

`func (o *DeploymentCH1Details) GetPropertiesOptions() map[string]interface{}`

GetPropertiesOptions returns the PropertiesOptions field if non-nil, zero value otherwise.

### GetPropertiesOptionsOk

`func (o *DeploymentCH1Details) GetPropertiesOptionsOk() (*map[string]interface{}, bool)`

GetPropertiesOptionsOk returns a tuple with the PropertiesOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertiesOptions

`func (o *DeploymentCH1Details) SetPropertiesOptions(v map[string]interface{})`

SetPropertiesOptions sets PropertiesOptions field to given value.

### HasPropertiesOptions

`func (o *DeploymentCH1Details) HasPropertiesOptions() bool`

HasPropertiesOptions returns a boolean if a field has been set.

### GetStatus

`func (o *DeploymentCH1Details) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeploymentCH1Details) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeploymentCH1Details) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeploymentCH1Details) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWorkers

`func (o *DeploymentCH1Details) GetWorkers() Workers`

GetWorkers returns the Workers field if non-nil, zero value otherwise.

### GetWorkersOk

`func (o *DeploymentCH1Details) GetWorkersOk() (*Workers, bool)`

GetWorkersOk returns a tuple with the Workers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkers

`func (o *DeploymentCH1Details) SetWorkers(v Workers)`

SetWorkers sets Workers field to given value.

### HasWorkers

`func (o *DeploymentCH1Details) HasWorkers() bool`

HasWorkers returns a boolean if a field has been set.

### GetLastUpdateTime

`func (o *DeploymentCH1Details) GetLastUpdateTime() int64`

GetLastUpdateTime returns the LastUpdateTime field if non-nil, zero value otherwise.

### GetLastUpdateTimeOk

`func (o *DeploymentCH1Details) GetLastUpdateTimeOk() (*int64, bool)`

GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTime

`func (o *DeploymentCH1Details) SetLastUpdateTime(v int64)`

SetLastUpdateTime sets LastUpdateTime field to given value.

### HasLastUpdateTime

`func (o *DeploymentCH1Details) HasLastUpdateTime() bool`

HasLastUpdateTime returns a boolean if a field has been set.

### GetFileName

`func (o *DeploymentCH1Details) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *DeploymentCH1Details) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *DeploymentCH1Details) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *DeploymentCH1Details) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetMuleVersion

`func (o *DeploymentCH1Details) GetMuleVersion() MuleVersion`

GetMuleVersion returns the MuleVersion field if non-nil, zero value otherwise.

### GetMuleVersionOk

`func (o *DeploymentCH1Details) GetMuleVersionOk() (*MuleVersion, bool)`

GetMuleVersionOk returns a tuple with the MuleVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuleVersion

`func (o *DeploymentCH1Details) SetMuleVersion(v MuleVersion)`

SetMuleVersion sets MuleVersion field to given value.

### HasMuleVersion

`func (o *DeploymentCH1Details) HasMuleVersion() bool`

HasMuleVersion returns a boolean if a field has been set.

### GetPreviousMuleVersion

`func (o *DeploymentCH1Details) GetPreviousMuleVersion() MuleVersion`

GetPreviousMuleVersion returns the PreviousMuleVersion field if non-nil, zero value otherwise.

### GetPreviousMuleVersionOk

`func (o *DeploymentCH1Details) GetPreviousMuleVersionOk() (*MuleVersion, bool)`

GetPreviousMuleVersionOk returns a tuple with the PreviousMuleVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousMuleVersion

`func (o *DeploymentCH1Details) SetPreviousMuleVersion(v MuleVersion)`

SetPreviousMuleVersion sets PreviousMuleVersion field to given value.

### HasPreviousMuleVersion

`func (o *DeploymentCH1Details) HasPreviousMuleVersion() bool`

HasPreviousMuleVersion returns a boolean if a field has been set.

### GetTenants

`func (o *DeploymentCH1Details) GetTenants() int32`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *DeploymentCH1Details) GetTenantsOk() (*int32, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *DeploymentCH1Details) SetTenants(v int32)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *DeploymentCH1Details) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetUserId

`func (o *DeploymentCH1Details) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DeploymentCH1Details) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DeploymentCH1Details) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *DeploymentCH1Details) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserName

`func (o *DeploymentCH1Details) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *DeploymentCH1Details) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *DeploymentCH1Details) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *DeploymentCH1Details) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetVpnConfig

`func (o *DeploymentCH1Details) GetVpnConfig() DeploymentCH1DetailsVpnConfig`

GetVpnConfig returns the VpnConfig field if non-nil, zero value otherwise.

### GetVpnConfigOk

`func (o *DeploymentCH1Details) GetVpnConfigOk() (*DeploymentCH1DetailsVpnConfig, bool)`

GetVpnConfigOk returns a tuple with the VpnConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnConfig

`func (o *DeploymentCH1Details) SetVpnConfig(v DeploymentCH1DetailsVpnConfig)`

SetVpnConfig sets VpnConfig field to given value.

### HasVpnConfig

`func (o *DeploymentCH1Details) HasVpnConfig() bool`

HasVpnConfig returns a boolean if a field has been set.

### GetDeploymentGroup

`func (o *DeploymentCH1Details) GetDeploymentGroup() DeploymentGroup`

GetDeploymentGroup returns the DeploymentGroup field if non-nil, zero value otherwise.

### GetDeploymentGroupOk

`func (o *DeploymentCH1Details) GetDeploymentGroupOk() (*DeploymentGroup, bool)`

GetDeploymentGroupOk returns a tuple with the DeploymentGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentGroup

`func (o *DeploymentCH1Details) SetDeploymentGroup(v DeploymentGroup)`

SetDeploymentGroup sets DeploymentGroup field to given value.

### HasDeploymentGroup

`func (o *DeploymentCH1Details) HasDeploymentGroup() bool`

HasDeploymentGroup returns a boolean if a field has been set.

### GetPersistentQueues

`func (o *DeploymentCH1Details) GetPersistentQueues() bool`

GetPersistentQueues returns the PersistentQueues field if non-nil, zero value otherwise.

### GetPersistentQueuesOk

`func (o *DeploymentCH1Details) GetPersistentQueuesOk() (*bool, bool)`

GetPersistentQueuesOk returns a tuple with the PersistentQueues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentQueues

`func (o *DeploymentCH1Details) SetPersistentQueues(v bool)`

SetPersistentQueues sets PersistentQueues field to given value.

### HasPersistentQueues

`func (o *DeploymentCH1Details) HasPersistentQueues() bool`

HasPersistentQueues returns a boolean if a field has been set.

### GetPersistentQueuesEncryptionEnabled

`func (o *DeploymentCH1Details) GetPersistentQueuesEncryptionEnabled() bool`

GetPersistentQueuesEncryptionEnabled returns the PersistentQueuesEncryptionEnabled field if non-nil, zero value otherwise.

### GetPersistentQueuesEncryptionEnabledOk

`func (o *DeploymentCH1Details) GetPersistentQueuesEncryptionEnabledOk() (*bool, bool)`

GetPersistentQueuesEncryptionEnabledOk returns a tuple with the PersistentQueuesEncryptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentQueuesEncryptionEnabled

`func (o *DeploymentCH1Details) SetPersistentQueuesEncryptionEnabled(v bool)`

SetPersistentQueuesEncryptionEnabled sets PersistentQueuesEncryptionEnabled field to given value.

### HasPersistentQueuesEncryptionEnabled

`func (o *DeploymentCH1Details) HasPersistentQueuesEncryptionEnabled() bool`

HasPersistentQueuesEncryptionEnabled returns a boolean if a field has been set.

### GetPersistentQueuesEncrypted

`func (o *DeploymentCH1Details) GetPersistentQueuesEncrypted() bool`

GetPersistentQueuesEncrypted returns the PersistentQueuesEncrypted field if non-nil, zero value otherwise.

### GetPersistentQueuesEncryptedOk

`func (o *DeploymentCH1Details) GetPersistentQueuesEncryptedOk() (*bool, bool)`

GetPersistentQueuesEncryptedOk returns a tuple with the PersistentQueuesEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentQueuesEncrypted

`func (o *DeploymentCH1Details) SetPersistentQueuesEncrypted(v bool)`

SetPersistentQueuesEncrypted sets PersistentQueuesEncrypted field to given value.

### HasPersistentQueuesEncrypted

`func (o *DeploymentCH1Details) HasPersistentQueuesEncrypted() bool`

HasPersistentQueuesEncrypted returns a boolean if a field has been set.

### GetMonitoringEnabled

`func (o *DeploymentCH1Details) GetMonitoringEnabled() bool`

GetMonitoringEnabled returns the MonitoringEnabled field if non-nil, zero value otherwise.

### GetMonitoringEnabledOk

`func (o *DeploymentCH1Details) GetMonitoringEnabledOk() (*bool, bool)`

GetMonitoringEnabledOk returns a tuple with the MonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringEnabled

`func (o *DeploymentCH1Details) SetMonitoringEnabled(v bool)`

SetMonitoringEnabled sets MonitoringEnabled field to given value.

### HasMonitoringEnabled

`func (o *DeploymentCH1Details) HasMonitoringEnabled() bool`

HasMonitoringEnabled returns a boolean if a field has been set.

### GetMonitoringAutoRestart

`func (o *DeploymentCH1Details) GetMonitoringAutoRestart() bool`

GetMonitoringAutoRestart returns the MonitoringAutoRestart field if non-nil, zero value otherwise.

### GetMonitoringAutoRestartOk

`func (o *DeploymentCH1Details) GetMonitoringAutoRestartOk() (*bool, bool)`

GetMonitoringAutoRestartOk returns a tuple with the MonitoringAutoRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringAutoRestart

`func (o *DeploymentCH1Details) SetMonitoringAutoRestart(v bool)`

SetMonitoringAutoRestart sets MonitoringAutoRestart field to given value.

### HasMonitoringAutoRestart

`func (o *DeploymentCH1Details) HasMonitoringAutoRestart() bool`

HasMonitoringAutoRestart returns a boolean if a field has been set.

### GetStaticIPsEnabled

`func (o *DeploymentCH1Details) GetStaticIPsEnabled() bool`

GetStaticIPsEnabled returns the StaticIPsEnabled field if non-nil, zero value otherwise.

### GetStaticIPsEnabledOk

`func (o *DeploymentCH1Details) GetStaticIPsEnabledOk() (*bool, bool)`

GetStaticIPsEnabledOk returns a tuple with the StaticIPsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticIPsEnabled

`func (o *DeploymentCH1Details) SetStaticIPsEnabled(v bool)`

SetStaticIPsEnabled sets StaticIPsEnabled field to given value.

### HasStaticIPsEnabled

`func (o *DeploymentCH1Details) HasStaticIPsEnabled() bool`

HasStaticIPsEnabled returns a boolean if a field has been set.

### GetMultitenanted

`func (o *DeploymentCH1Details) GetMultitenanted() bool`

GetMultitenanted returns the Multitenanted field if non-nil, zero value otherwise.

### GetMultitenantedOk

`func (o *DeploymentCH1Details) GetMultitenantedOk() (*bool, bool)`

GetMultitenantedOk returns a tuple with the Multitenanted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultitenanted

`func (o *DeploymentCH1Details) SetMultitenanted(v bool)`

SetMultitenanted sets Multitenanted field to given value.

### HasMultitenanted

`func (o *DeploymentCH1Details) HasMultitenanted() bool`

HasMultitenanted returns a boolean if a field has been set.

### GetHasFile

`func (o *DeploymentCH1Details) GetHasFile() bool`

GetHasFile returns the HasFile field if non-nil, zero value otherwise.

### GetHasFileOk

`func (o *DeploymentCH1Details) GetHasFileOk() (*bool, bool)`

GetHasFileOk returns a tuple with the HasFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFile

`func (o *DeploymentCH1Details) SetHasFile(v bool)`

SetHasFile sets HasFile field to given value.

### HasHasFile

`func (o *DeploymentCH1Details) HasHasFile() bool`

HasHasFile returns a boolean if a field has been set.

### GetSecureDataGatewayEnabled

`func (o *DeploymentCH1Details) GetSecureDataGatewayEnabled() bool`

GetSecureDataGatewayEnabled returns the SecureDataGatewayEnabled field if non-nil, zero value otherwise.

### GetSecureDataGatewayEnabledOk

`func (o *DeploymentCH1Details) GetSecureDataGatewayEnabledOk() (*bool, bool)`

GetSecureDataGatewayEnabledOk returns a tuple with the SecureDataGatewayEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureDataGatewayEnabled

`func (o *DeploymentCH1Details) SetSecureDataGatewayEnabled(v bool)`

SetSecureDataGatewayEnabled sets SecureDataGatewayEnabled field to given value.

### HasSecureDataGatewayEnabled

`func (o *DeploymentCH1Details) HasSecureDataGatewayEnabled() bool`

HasSecureDataGatewayEnabled returns a boolean if a field has been set.

### GetVpnEnabled

`func (o *DeploymentCH1Details) GetVpnEnabled() bool`

GetVpnEnabled returns the VpnEnabled field if non-nil, zero value otherwise.

### GetVpnEnabledOk

`func (o *DeploymentCH1Details) GetVpnEnabledOk() (*bool, bool)`

GetVpnEnabledOk returns a tuple with the VpnEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnEnabled

`func (o *DeploymentCH1Details) SetVpnEnabled(v bool)`

SetVpnEnabled sets VpnEnabled field to given value.

### HasVpnEnabled

`func (o *DeploymentCH1Details) HasVpnEnabled() bool`

HasVpnEnabled returns a boolean if a field has been set.

### GetTrackingSettings

`func (o *DeploymentCH1Details) GetTrackingSettings() TrackingSettings`

GetTrackingSettings returns the TrackingSettings field if non-nil, zero value otherwise.

### GetTrackingSettingsOk

`func (o *DeploymentCH1Details) GetTrackingSettingsOk() (*TrackingSettings, bool)`

GetTrackingSettingsOk returns a tuple with the TrackingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingSettings

`func (o *DeploymentCH1Details) SetTrackingSettings(v TrackingSettings)`

SetTrackingSettings sets TrackingSettings field to given value.

### HasTrackingSettings

`func (o *DeploymentCH1Details) HasTrackingSettings() bool`

HasTrackingSettings returns a boolean if a field has been set.

### GetLogLevels

`func (o *DeploymentCH1Details) GetLogLevels() []LogLevel`

GetLogLevels returns the LogLevels field if non-nil, zero value otherwise.

### GetLogLevelsOk

`func (o *DeploymentCH1Details) GetLogLevelsOk() (*[]LogLevel, bool)`

GetLogLevelsOk returns a tuple with the LogLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevels

`func (o *DeploymentCH1Details) SetLogLevels(v []LogLevel)`

SetLogLevels sets LogLevels field to given value.

### HasLogLevels

`func (o *DeploymentCH1Details) HasLogLevels() bool`

HasLogLevels returns a boolean if a field has been set.

### GetIpAddresses

`func (o *DeploymentCH1Details) GetIpAddresses() []IpAddress`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *DeploymentCH1Details) GetIpAddressesOk() (*[]IpAddress, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *DeploymentCH1Details) SetIpAddresses(v []IpAddress)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *DeploymentCH1Details) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


