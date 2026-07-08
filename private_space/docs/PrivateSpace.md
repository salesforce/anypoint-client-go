# PrivateSpace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the private space in GUID format. | [optional] 
**Name** | Pointer to **string** | The name of the private space. | [optional] 
**Version** | Pointer to **string** | The version of the private space. | [optional] 
**OrganizationId** | Pointer to **string** | The ID of the organization in GUID format. | [optional] 
**RootOrganizationId** | Pointer to **string** | The root organization ID of the private space in GUID format. | [optional] 
**Status** | Pointer to **string** | The status of the private space. | [optional] 
**StatusMessage** | Pointer to **string** | The last reported infra status message. | [optional] 
**Provisioning** | Pointer to [**PrivateSpaceProvisioning**](PrivateSpaceProvisioning.md) |  | [optional] 
**Region** | Pointer to **string** | The region of the private space. | [optional] 
**Environments** | Pointer to [**PrivateSpaceAssociatedEnvironments**](PrivateSpaceAssociatedEnvironments.md) |  | [optional] 
**Network** | Pointer to [**PrivateSpaceNetwork**](PrivateSpaceNetwork.md) |  | [optional] 
**FirewallRules** | Pointer to [**[]FirewallRule**](FirewallRule.md) | The list of firewall rules for the Private Space network. | [optional] 
**IngressConfiguration** | Pointer to [**IngressConfiguration**](IngressConfiguration.md) |  | [optional] 
**EnableIAMRole** | Pointer to **bool** | If true, application deployed to this space will have the Private Space IAM role attached to the service account. | [optional] 
**EnableEgress** | Pointer to **bool** | If true, egress is enabled for the private space. | [optional] 
**EnableNetworkIsolation** | Pointer to **bool** | If true, network isolation is enabled for the private space. | [optional] 
**GlobalSpaceStatus** | Pointer to [**GlobalSpaceStatus**](GlobalSpaceStatus.md) |  | [optional] 
**MuleAppDeploymentCount** | Pointer to **int32** | The number of MuleSoft applications deployed to the private space. | [optional] 
**DaysLeftForRelaxedQuota** | Pointer to **int32** | The number of days left for the relaxed quota. | [optional] 
**VpcMigrationInProgress** | Pointer to **bool** | If true, a VPC migration is in progress for the private space. | [optional] 

## Methods

### NewPrivateSpace

`func NewPrivateSpace() *PrivateSpace`

NewPrivateSpace instantiates a new PrivateSpace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateSpaceWithDefaults

`func NewPrivateSpaceWithDefaults() *PrivateSpace`

NewPrivateSpaceWithDefaults instantiates a new PrivateSpace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrivateSpace) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrivateSpace) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrivateSpace) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PrivateSpace) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PrivateSpace) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrivateSpace) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrivateSpace) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PrivateSpace) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *PrivateSpace) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PrivateSpace) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PrivateSpace) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PrivateSpace) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetOrganizationId

`func (o *PrivateSpace) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *PrivateSpace) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *PrivateSpace) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *PrivateSpace) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetRootOrganizationId

`func (o *PrivateSpace) GetRootOrganizationId() string`

GetRootOrganizationId returns the RootOrganizationId field if non-nil, zero value otherwise.

### GetRootOrganizationIdOk

`func (o *PrivateSpace) GetRootOrganizationIdOk() (*string, bool)`

GetRootOrganizationIdOk returns a tuple with the RootOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootOrganizationId

`func (o *PrivateSpace) SetRootOrganizationId(v string)`

SetRootOrganizationId sets RootOrganizationId field to given value.

### HasRootOrganizationId

`func (o *PrivateSpace) HasRootOrganizationId() bool`

HasRootOrganizationId returns a boolean if a field has been set.

### GetStatus

`func (o *PrivateSpace) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PrivateSpace) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PrivateSpace) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PrivateSpace) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *PrivateSpace) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *PrivateSpace) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *PrivateSpace) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *PrivateSpace) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetProvisioning

`func (o *PrivateSpace) GetProvisioning() PrivateSpaceProvisioning`

GetProvisioning returns the Provisioning field if non-nil, zero value otherwise.

### GetProvisioningOk

`func (o *PrivateSpace) GetProvisioningOk() (*PrivateSpaceProvisioning, bool)`

GetProvisioningOk returns a tuple with the Provisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioning

`func (o *PrivateSpace) SetProvisioning(v PrivateSpaceProvisioning)`

SetProvisioning sets Provisioning field to given value.

### HasProvisioning

`func (o *PrivateSpace) HasProvisioning() bool`

HasProvisioning returns a boolean if a field has been set.

### GetRegion

`func (o *PrivateSpace) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PrivateSpace) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PrivateSpace) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *PrivateSpace) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetEnvironments

`func (o *PrivateSpace) GetEnvironments() PrivateSpaceAssociatedEnvironments`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *PrivateSpace) GetEnvironmentsOk() (*PrivateSpaceAssociatedEnvironments, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *PrivateSpace) SetEnvironments(v PrivateSpaceAssociatedEnvironments)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *PrivateSpace) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetNetwork

`func (o *PrivateSpace) GetNetwork() PrivateSpaceNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *PrivateSpace) GetNetworkOk() (*PrivateSpaceNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *PrivateSpace) SetNetwork(v PrivateSpaceNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *PrivateSpace) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetFirewallRules

`func (o *PrivateSpace) GetFirewallRules() []FirewallRule`

GetFirewallRules returns the FirewallRules field if non-nil, zero value otherwise.

### GetFirewallRulesOk

`func (o *PrivateSpace) GetFirewallRulesOk() (*[]FirewallRule, bool)`

GetFirewallRulesOk returns a tuple with the FirewallRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallRules

`func (o *PrivateSpace) SetFirewallRules(v []FirewallRule)`

SetFirewallRules sets FirewallRules field to given value.

### HasFirewallRules

`func (o *PrivateSpace) HasFirewallRules() bool`

HasFirewallRules returns a boolean if a field has been set.

### GetIngressConfiguration

`func (o *PrivateSpace) GetIngressConfiguration() IngressConfiguration`

GetIngressConfiguration returns the IngressConfiguration field if non-nil, zero value otherwise.

### GetIngressConfigurationOk

`func (o *PrivateSpace) GetIngressConfigurationOk() (*IngressConfiguration, bool)`

GetIngressConfigurationOk returns a tuple with the IngressConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressConfiguration

`func (o *PrivateSpace) SetIngressConfiguration(v IngressConfiguration)`

SetIngressConfiguration sets IngressConfiguration field to given value.

### HasIngressConfiguration

`func (o *PrivateSpace) HasIngressConfiguration() bool`

HasIngressConfiguration returns a boolean if a field has been set.

### GetEnableIAMRole

`func (o *PrivateSpace) GetEnableIAMRole() bool`

GetEnableIAMRole returns the EnableIAMRole field if non-nil, zero value otherwise.

### GetEnableIAMRoleOk

`func (o *PrivateSpace) GetEnableIAMRoleOk() (*bool, bool)`

GetEnableIAMRoleOk returns a tuple with the EnableIAMRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIAMRole

`func (o *PrivateSpace) SetEnableIAMRole(v bool)`

SetEnableIAMRole sets EnableIAMRole field to given value.

### HasEnableIAMRole

`func (o *PrivateSpace) HasEnableIAMRole() bool`

HasEnableIAMRole returns a boolean if a field has been set.

### GetEnableEgress

`func (o *PrivateSpace) GetEnableEgress() bool`

GetEnableEgress returns the EnableEgress field if non-nil, zero value otherwise.

### GetEnableEgressOk

`func (o *PrivateSpace) GetEnableEgressOk() (*bool, bool)`

GetEnableEgressOk returns a tuple with the EnableEgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEgress

`func (o *PrivateSpace) SetEnableEgress(v bool)`

SetEnableEgress sets EnableEgress field to given value.

### HasEnableEgress

`func (o *PrivateSpace) HasEnableEgress() bool`

HasEnableEgress returns a boolean if a field has been set.

### GetEnableNetworkIsolation

`func (o *PrivateSpace) GetEnableNetworkIsolation() bool`

GetEnableNetworkIsolation returns the EnableNetworkIsolation field if non-nil, zero value otherwise.

### GetEnableNetworkIsolationOk

`func (o *PrivateSpace) GetEnableNetworkIsolationOk() (*bool, bool)`

GetEnableNetworkIsolationOk returns a tuple with the EnableNetworkIsolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNetworkIsolation

`func (o *PrivateSpace) SetEnableNetworkIsolation(v bool)`

SetEnableNetworkIsolation sets EnableNetworkIsolation field to given value.

### HasEnableNetworkIsolation

`func (o *PrivateSpace) HasEnableNetworkIsolation() bool`

HasEnableNetworkIsolation returns a boolean if a field has been set.

### GetGlobalSpaceStatus

`func (o *PrivateSpace) GetGlobalSpaceStatus() GlobalSpaceStatus`

GetGlobalSpaceStatus returns the GlobalSpaceStatus field if non-nil, zero value otherwise.

### GetGlobalSpaceStatusOk

`func (o *PrivateSpace) GetGlobalSpaceStatusOk() (*GlobalSpaceStatus, bool)`

GetGlobalSpaceStatusOk returns a tuple with the GlobalSpaceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalSpaceStatus

`func (o *PrivateSpace) SetGlobalSpaceStatus(v GlobalSpaceStatus)`

SetGlobalSpaceStatus sets GlobalSpaceStatus field to given value.

### HasGlobalSpaceStatus

`func (o *PrivateSpace) HasGlobalSpaceStatus() bool`

HasGlobalSpaceStatus returns a boolean if a field has been set.

### GetMuleAppDeploymentCount

`func (o *PrivateSpace) GetMuleAppDeploymentCount() int32`

GetMuleAppDeploymentCount returns the MuleAppDeploymentCount field if non-nil, zero value otherwise.

### GetMuleAppDeploymentCountOk

`func (o *PrivateSpace) GetMuleAppDeploymentCountOk() (*int32, bool)`

GetMuleAppDeploymentCountOk returns a tuple with the MuleAppDeploymentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuleAppDeploymentCount

`func (o *PrivateSpace) SetMuleAppDeploymentCount(v int32)`

SetMuleAppDeploymentCount sets MuleAppDeploymentCount field to given value.

### HasMuleAppDeploymentCount

`func (o *PrivateSpace) HasMuleAppDeploymentCount() bool`

HasMuleAppDeploymentCount returns a boolean if a field has been set.

### GetDaysLeftForRelaxedQuota

`func (o *PrivateSpace) GetDaysLeftForRelaxedQuota() int32`

GetDaysLeftForRelaxedQuota returns the DaysLeftForRelaxedQuota field if non-nil, zero value otherwise.

### GetDaysLeftForRelaxedQuotaOk

`func (o *PrivateSpace) GetDaysLeftForRelaxedQuotaOk() (*int32, bool)`

GetDaysLeftForRelaxedQuotaOk returns a tuple with the DaysLeftForRelaxedQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysLeftForRelaxedQuota

`func (o *PrivateSpace) SetDaysLeftForRelaxedQuota(v int32)`

SetDaysLeftForRelaxedQuota sets DaysLeftForRelaxedQuota field to given value.

### HasDaysLeftForRelaxedQuota

`func (o *PrivateSpace) HasDaysLeftForRelaxedQuota() bool`

HasDaysLeftForRelaxedQuota returns a boolean if a field has been set.

### GetVpcMigrationInProgress

`func (o *PrivateSpace) GetVpcMigrationInProgress() bool`

GetVpcMigrationInProgress returns the VpcMigrationInProgress field if non-nil, zero value otherwise.

### GetVpcMigrationInProgressOk

`func (o *PrivateSpace) GetVpcMigrationInProgressOk() (*bool, bool)`

GetVpcMigrationInProgressOk returns a tuple with the VpcMigrationInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcMigrationInProgress

`func (o *PrivateSpace) SetVpcMigrationInProgress(v bool)`

SetVpcMigrationInProgress sets VpcMigrationInProgress field to given value.

### HasVpcMigrationInProgress

`func (o *PrivateSpace) HasVpcMigrationInProgress() bool`

HasVpcMigrationInProgress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


