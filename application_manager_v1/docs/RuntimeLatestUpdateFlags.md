# RuntimeLatestUpdateFlags

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PerformLog4jReplacement** | Pointer to **bool** |  | [optional] 
**SkipMMCPairing** | Pointer to **bool** |  | [optional] 
**Log4j1Used** | Pointer to **bool** |  | [optional] 
**UsingMuleAgent** | Pointer to **bool** |  | [optional] 
**VpnSupported** | Pointer to **bool** |  | [optional] 
**MonitoringSupported** | Pointer to **bool** |  | [optional] 
**ObjectStoreV1** | Pointer to **bool** |  | [optional] 
**LoggingNgSupported** | Pointer to **bool** |  | [optional] 
**DiagnosticsSupported** | Pointer to **bool** |  | [optional] 
**PersistentQueuesSupported** | Pointer to **bool** |  | [optional] 
**AsyncScheduleSupported** | Pointer to **bool** |  | [optional] 
**MuleLivenessSupported** | Pointer to **bool** |  | [optional] 
**MuleReadinessCheckSupported** | Pointer to **bool** |  | [optional] 
**OcsServiceSupported** | Pointer to **bool** |  | [optional] 
**InjectLog4j2Supported** | Pointer to **bool** |  | [optional] 
**InjectMuleLogsWhenFail** | Pointer to **bool** |  | [optional] 
**AmiVersionData** | Pointer to [**RuntimeLatestUpdateFlagsAmiVersionData**](RuntimeLatestUpdateFlagsAmiVersionData.md) |  | [optional] 

## Methods

### NewRuntimeLatestUpdateFlags

`func NewRuntimeLatestUpdateFlags() *RuntimeLatestUpdateFlags`

NewRuntimeLatestUpdateFlags instantiates a new RuntimeLatestUpdateFlags object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuntimeLatestUpdateFlagsWithDefaults

`func NewRuntimeLatestUpdateFlagsWithDefaults() *RuntimeLatestUpdateFlags`

NewRuntimeLatestUpdateFlagsWithDefaults instantiates a new RuntimeLatestUpdateFlags object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPerformLog4jReplacement

`func (o *RuntimeLatestUpdateFlags) GetPerformLog4jReplacement() bool`

GetPerformLog4jReplacement returns the PerformLog4jReplacement field if non-nil, zero value otherwise.

### GetPerformLog4jReplacementOk

`func (o *RuntimeLatestUpdateFlags) GetPerformLog4jReplacementOk() (*bool, bool)`

GetPerformLog4jReplacementOk returns a tuple with the PerformLog4jReplacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformLog4jReplacement

`func (o *RuntimeLatestUpdateFlags) SetPerformLog4jReplacement(v bool)`

SetPerformLog4jReplacement sets PerformLog4jReplacement field to given value.

### HasPerformLog4jReplacement

`func (o *RuntimeLatestUpdateFlags) HasPerformLog4jReplacement() bool`

HasPerformLog4jReplacement returns a boolean if a field has been set.

### GetSkipMMCPairing

`func (o *RuntimeLatestUpdateFlags) GetSkipMMCPairing() bool`

GetSkipMMCPairing returns the SkipMMCPairing field if non-nil, zero value otherwise.

### GetSkipMMCPairingOk

`func (o *RuntimeLatestUpdateFlags) GetSkipMMCPairingOk() (*bool, bool)`

GetSkipMMCPairingOk returns a tuple with the SkipMMCPairing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipMMCPairing

`func (o *RuntimeLatestUpdateFlags) SetSkipMMCPairing(v bool)`

SetSkipMMCPairing sets SkipMMCPairing field to given value.

### HasSkipMMCPairing

`func (o *RuntimeLatestUpdateFlags) HasSkipMMCPairing() bool`

HasSkipMMCPairing returns a boolean if a field has been set.

### GetLog4j1Used

`func (o *RuntimeLatestUpdateFlags) GetLog4j1Used() bool`

GetLog4j1Used returns the Log4j1Used field if non-nil, zero value otherwise.

### GetLog4j1UsedOk

`func (o *RuntimeLatestUpdateFlags) GetLog4j1UsedOk() (*bool, bool)`

GetLog4j1UsedOk returns a tuple with the Log4j1Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog4j1Used

`func (o *RuntimeLatestUpdateFlags) SetLog4j1Used(v bool)`

SetLog4j1Used sets Log4j1Used field to given value.

### HasLog4j1Used

`func (o *RuntimeLatestUpdateFlags) HasLog4j1Used() bool`

HasLog4j1Used returns a boolean if a field has been set.

### GetUsingMuleAgent

`func (o *RuntimeLatestUpdateFlags) GetUsingMuleAgent() bool`

GetUsingMuleAgent returns the UsingMuleAgent field if non-nil, zero value otherwise.

### GetUsingMuleAgentOk

`func (o *RuntimeLatestUpdateFlags) GetUsingMuleAgentOk() (*bool, bool)`

GetUsingMuleAgentOk returns a tuple with the UsingMuleAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsingMuleAgent

`func (o *RuntimeLatestUpdateFlags) SetUsingMuleAgent(v bool)`

SetUsingMuleAgent sets UsingMuleAgent field to given value.

### HasUsingMuleAgent

`func (o *RuntimeLatestUpdateFlags) HasUsingMuleAgent() bool`

HasUsingMuleAgent returns a boolean if a field has been set.

### GetVpnSupported

`func (o *RuntimeLatestUpdateFlags) GetVpnSupported() bool`

GetVpnSupported returns the VpnSupported field if non-nil, zero value otherwise.

### GetVpnSupportedOk

`func (o *RuntimeLatestUpdateFlags) GetVpnSupportedOk() (*bool, bool)`

GetVpnSupportedOk returns a tuple with the VpnSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnSupported

`func (o *RuntimeLatestUpdateFlags) SetVpnSupported(v bool)`

SetVpnSupported sets VpnSupported field to given value.

### HasVpnSupported

`func (o *RuntimeLatestUpdateFlags) HasVpnSupported() bool`

HasVpnSupported returns a boolean if a field has been set.

### GetMonitoringSupported

`func (o *RuntimeLatestUpdateFlags) GetMonitoringSupported() bool`

GetMonitoringSupported returns the MonitoringSupported field if non-nil, zero value otherwise.

### GetMonitoringSupportedOk

`func (o *RuntimeLatestUpdateFlags) GetMonitoringSupportedOk() (*bool, bool)`

GetMonitoringSupportedOk returns a tuple with the MonitoringSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringSupported

`func (o *RuntimeLatestUpdateFlags) SetMonitoringSupported(v bool)`

SetMonitoringSupported sets MonitoringSupported field to given value.

### HasMonitoringSupported

`func (o *RuntimeLatestUpdateFlags) HasMonitoringSupported() bool`

HasMonitoringSupported returns a boolean if a field has been set.

### GetObjectStoreV1

`func (o *RuntimeLatestUpdateFlags) GetObjectStoreV1() bool`

GetObjectStoreV1 returns the ObjectStoreV1 field if non-nil, zero value otherwise.

### GetObjectStoreV1Ok

`func (o *RuntimeLatestUpdateFlags) GetObjectStoreV1Ok() (*bool, bool)`

GetObjectStoreV1Ok returns a tuple with the ObjectStoreV1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectStoreV1

`func (o *RuntimeLatestUpdateFlags) SetObjectStoreV1(v bool)`

SetObjectStoreV1 sets ObjectStoreV1 field to given value.

### HasObjectStoreV1

`func (o *RuntimeLatestUpdateFlags) HasObjectStoreV1() bool`

HasObjectStoreV1 returns a boolean if a field has been set.

### GetLoggingNgSupported

`func (o *RuntimeLatestUpdateFlags) GetLoggingNgSupported() bool`

GetLoggingNgSupported returns the LoggingNgSupported field if non-nil, zero value otherwise.

### GetLoggingNgSupportedOk

`func (o *RuntimeLatestUpdateFlags) GetLoggingNgSupportedOk() (*bool, bool)`

GetLoggingNgSupportedOk returns a tuple with the LoggingNgSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingNgSupported

`func (o *RuntimeLatestUpdateFlags) SetLoggingNgSupported(v bool)`

SetLoggingNgSupported sets LoggingNgSupported field to given value.

### HasLoggingNgSupported

`func (o *RuntimeLatestUpdateFlags) HasLoggingNgSupported() bool`

HasLoggingNgSupported returns a boolean if a field has been set.

### GetDiagnosticsSupported

`func (o *RuntimeLatestUpdateFlags) GetDiagnosticsSupported() bool`

GetDiagnosticsSupported returns the DiagnosticsSupported field if non-nil, zero value otherwise.

### GetDiagnosticsSupportedOk

`func (o *RuntimeLatestUpdateFlags) GetDiagnosticsSupportedOk() (*bool, bool)`

GetDiagnosticsSupportedOk returns a tuple with the DiagnosticsSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagnosticsSupported

`func (o *RuntimeLatestUpdateFlags) SetDiagnosticsSupported(v bool)`

SetDiagnosticsSupported sets DiagnosticsSupported field to given value.

### HasDiagnosticsSupported

`func (o *RuntimeLatestUpdateFlags) HasDiagnosticsSupported() bool`

HasDiagnosticsSupported returns a boolean if a field has been set.

### GetPersistentQueuesSupported

`func (o *RuntimeLatestUpdateFlags) GetPersistentQueuesSupported() bool`

GetPersistentQueuesSupported returns the PersistentQueuesSupported field if non-nil, zero value otherwise.

### GetPersistentQueuesSupportedOk

`func (o *RuntimeLatestUpdateFlags) GetPersistentQueuesSupportedOk() (*bool, bool)`

GetPersistentQueuesSupportedOk returns a tuple with the PersistentQueuesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentQueuesSupported

`func (o *RuntimeLatestUpdateFlags) SetPersistentQueuesSupported(v bool)`

SetPersistentQueuesSupported sets PersistentQueuesSupported field to given value.

### HasPersistentQueuesSupported

`func (o *RuntimeLatestUpdateFlags) HasPersistentQueuesSupported() bool`

HasPersistentQueuesSupported returns a boolean if a field has been set.

### GetAsyncScheduleSupported

`func (o *RuntimeLatestUpdateFlags) GetAsyncScheduleSupported() bool`

GetAsyncScheduleSupported returns the AsyncScheduleSupported field if non-nil, zero value otherwise.

### GetAsyncScheduleSupportedOk

`func (o *RuntimeLatestUpdateFlags) GetAsyncScheduleSupportedOk() (*bool, bool)`

GetAsyncScheduleSupportedOk returns a tuple with the AsyncScheduleSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsyncScheduleSupported

`func (o *RuntimeLatestUpdateFlags) SetAsyncScheduleSupported(v bool)`

SetAsyncScheduleSupported sets AsyncScheduleSupported field to given value.

### HasAsyncScheduleSupported

`func (o *RuntimeLatestUpdateFlags) HasAsyncScheduleSupported() bool`

HasAsyncScheduleSupported returns a boolean if a field has been set.

### GetMuleLivenessSupported

`func (o *RuntimeLatestUpdateFlags) GetMuleLivenessSupported() bool`

GetMuleLivenessSupported returns the MuleLivenessSupported field if non-nil, zero value otherwise.

### GetMuleLivenessSupportedOk

`func (o *RuntimeLatestUpdateFlags) GetMuleLivenessSupportedOk() (*bool, bool)`

GetMuleLivenessSupportedOk returns a tuple with the MuleLivenessSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuleLivenessSupported

`func (o *RuntimeLatestUpdateFlags) SetMuleLivenessSupported(v bool)`

SetMuleLivenessSupported sets MuleLivenessSupported field to given value.

### HasMuleLivenessSupported

`func (o *RuntimeLatestUpdateFlags) HasMuleLivenessSupported() bool`

HasMuleLivenessSupported returns a boolean if a field has been set.

### GetMuleReadinessCheckSupported

`func (o *RuntimeLatestUpdateFlags) GetMuleReadinessCheckSupported() bool`

GetMuleReadinessCheckSupported returns the MuleReadinessCheckSupported field if non-nil, zero value otherwise.

### GetMuleReadinessCheckSupportedOk

`func (o *RuntimeLatestUpdateFlags) GetMuleReadinessCheckSupportedOk() (*bool, bool)`

GetMuleReadinessCheckSupportedOk returns a tuple with the MuleReadinessCheckSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuleReadinessCheckSupported

`func (o *RuntimeLatestUpdateFlags) SetMuleReadinessCheckSupported(v bool)`

SetMuleReadinessCheckSupported sets MuleReadinessCheckSupported field to given value.

### HasMuleReadinessCheckSupported

`func (o *RuntimeLatestUpdateFlags) HasMuleReadinessCheckSupported() bool`

HasMuleReadinessCheckSupported returns a boolean if a field has been set.

### GetOcsServiceSupported

`func (o *RuntimeLatestUpdateFlags) GetOcsServiceSupported() bool`

GetOcsServiceSupported returns the OcsServiceSupported field if non-nil, zero value otherwise.

### GetOcsServiceSupportedOk

`func (o *RuntimeLatestUpdateFlags) GetOcsServiceSupportedOk() (*bool, bool)`

GetOcsServiceSupportedOk returns a tuple with the OcsServiceSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcsServiceSupported

`func (o *RuntimeLatestUpdateFlags) SetOcsServiceSupported(v bool)`

SetOcsServiceSupported sets OcsServiceSupported field to given value.

### HasOcsServiceSupported

`func (o *RuntimeLatestUpdateFlags) HasOcsServiceSupported() bool`

HasOcsServiceSupported returns a boolean if a field has been set.

### GetInjectLog4j2Supported

`func (o *RuntimeLatestUpdateFlags) GetInjectLog4j2Supported() bool`

GetInjectLog4j2Supported returns the InjectLog4j2Supported field if non-nil, zero value otherwise.

### GetInjectLog4j2SupportedOk

`func (o *RuntimeLatestUpdateFlags) GetInjectLog4j2SupportedOk() (*bool, bool)`

GetInjectLog4j2SupportedOk returns a tuple with the InjectLog4j2Supported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInjectLog4j2Supported

`func (o *RuntimeLatestUpdateFlags) SetInjectLog4j2Supported(v bool)`

SetInjectLog4j2Supported sets InjectLog4j2Supported field to given value.

### HasInjectLog4j2Supported

`func (o *RuntimeLatestUpdateFlags) HasInjectLog4j2Supported() bool`

HasInjectLog4j2Supported returns a boolean if a field has been set.

### GetInjectMuleLogsWhenFail

`func (o *RuntimeLatestUpdateFlags) GetInjectMuleLogsWhenFail() bool`

GetInjectMuleLogsWhenFail returns the InjectMuleLogsWhenFail field if non-nil, zero value otherwise.

### GetInjectMuleLogsWhenFailOk

`func (o *RuntimeLatestUpdateFlags) GetInjectMuleLogsWhenFailOk() (*bool, bool)`

GetInjectMuleLogsWhenFailOk returns a tuple with the InjectMuleLogsWhenFail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInjectMuleLogsWhenFail

`func (o *RuntimeLatestUpdateFlags) SetInjectMuleLogsWhenFail(v bool)`

SetInjectMuleLogsWhenFail sets InjectMuleLogsWhenFail field to given value.

### HasInjectMuleLogsWhenFail

`func (o *RuntimeLatestUpdateFlags) HasInjectMuleLogsWhenFail() bool`

HasInjectMuleLogsWhenFail returns a boolean if a field has been set.

### GetAmiVersionData

`func (o *RuntimeLatestUpdateFlags) GetAmiVersionData() RuntimeLatestUpdateFlagsAmiVersionData`

GetAmiVersionData returns the AmiVersionData field if non-nil, zero value otherwise.

### GetAmiVersionDataOk

`func (o *RuntimeLatestUpdateFlags) GetAmiVersionDataOk() (*RuntimeLatestUpdateFlagsAmiVersionData, bool)`

GetAmiVersionDataOk returns a tuple with the AmiVersionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmiVersionData

`func (o *RuntimeLatestUpdateFlags) SetAmiVersionData(v RuntimeLatestUpdateFlagsAmiVersionData)`

SetAmiVersionData sets AmiVersionData field to given value.

### HasAmiVersionData

`func (o *RuntimeLatestUpdateFlags) HasAmiVersionData() bool`

HasAmiVersionData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


