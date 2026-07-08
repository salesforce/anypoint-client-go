# PrivateSpacePatchBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environments** | Pointer to [**PrivateSpaceAssociatedEnvironments**](PrivateSpaceAssociatedEnvironments.md) |  | [optional] 
**Network** | Pointer to [**PrivateSpaceNetworkEditable**](PrivateSpaceNetworkEditable.md) |  | [optional] 
**FirewallRules** | Pointer to [**[]FirewallRule**](FirewallRule.md) | The list of firewall rules for the Private Space network. | [optional] 
**LogForwarding** | Pointer to [**PrivateSpacePatchBodyLogForwarding**](PrivateSpacePatchBodyLogForwarding.md) |  | [optional] 
**IngressConfiguration** | Pointer to [**IngressConfiguration**](IngressConfiguration.md) |  | [optional] 
**EnableIAMRole** | Pointer to **bool** | If true, application deployed to this space will have the Private Space IAM role attached to the service account. | [optional] 
**EnableEgress** | Pointer to **bool** | If true, egress is enabled for the private space. | [optional] 
**EnableNetworkIsolation** | Pointer to **bool** | If true, network isolation is enabled for the private space. | [optional] 

## Methods

### NewPrivateSpacePatchBody

`func NewPrivateSpacePatchBody() *PrivateSpacePatchBody`

NewPrivateSpacePatchBody instantiates a new PrivateSpacePatchBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateSpacePatchBodyWithDefaults

`func NewPrivateSpacePatchBodyWithDefaults() *PrivateSpacePatchBody`

NewPrivateSpacePatchBodyWithDefaults instantiates a new PrivateSpacePatchBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironments

`func (o *PrivateSpacePatchBody) GetEnvironments() PrivateSpaceAssociatedEnvironments`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *PrivateSpacePatchBody) GetEnvironmentsOk() (*PrivateSpaceAssociatedEnvironments, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *PrivateSpacePatchBody) SetEnvironments(v PrivateSpaceAssociatedEnvironments)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *PrivateSpacePatchBody) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetNetwork

`func (o *PrivateSpacePatchBody) GetNetwork() PrivateSpaceNetworkEditable`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *PrivateSpacePatchBody) GetNetworkOk() (*PrivateSpaceNetworkEditable, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *PrivateSpacePatchBody) SetNetwork(v PrivateSpaceNetworkEditable)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *PrivateSpacePatchBody) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetFirewallRules

`func (o *PrivateSpacePatchBody) GetFirewallRules() []FirewallRule`

GetFirewallRules returns the FirewallRules field if non-nil, zero value otherwise.

### GetFirewallRulesOk

`func (o *PrivateSpacePatchBody) GetFirewallRulesOk() (*[]FirewallRule, bool)`

GetFirewallRulesOk returns a tuple with the FirewallRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallRules

`func (o *PrivateSpacePatchBody) SetFirewallRules(v []FirewallRule)`

SetFirewallRules sets FirewallRules field to given value.

### HasFirewallRules

`func (o *PrivateSpacePatchBody) HasFirewallRules() bool`

HasFirewallRules returns a boolean if a field has been set.

### GetLogForwarding

`func (o *PrivateSpacePatchBody) GetLogForwarding() PrivateSpacePatchBodyLogForwarding`

GetLogForwarding returns the LogForwarding field if non-nil, zero value otherwise.

### GetLogForwardingOk

`func (o *PrivateSpacePatchBody) GetLogForwardingOk() (*PrivateSpacePatchBodyLogForwarding, bool)`

GetLogForwardingOk returns a tuple with the LogForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogForwarding

`func (o *PrivateSpacePatchBody) SetLogForwarding(v PrivateSpacePatchBodyLogForwarding)`

SetLogForwarding sets LogForwarding field to given value.

### HasLogForwarding

`func (o *PrivateSpacePatchBody) HasLogForwarding() bool`

HasLogForwarding returns a boolean if a field has been set.

### GetIngressConfiguration

`func (o *PrivateSpacePatchBody) GetIngressConfiguration() IngressConfiguration`

GetIngressConfiguration returns the IngressConfiguration field if non-nil, zero value otherwise.

### GetIngressConfigurationOk

`func (o *PrivateSpacePatchBody) GetIngressConfigurationOk() (*IngressConfiguration, bool)`

GetIngressConfigurationOk returns a tuple with the IngressConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressConfiguration

`func (o *PrivateSpacePatchBody) SetIngressConfiguration(v IngressConfiguration)`

SetIngressConfiguration sets IngressConfiguration field to given value.

### HasIngressConfiguration

`func (o *PrivateSpacePatchBody) HasIngressConfiguration() bool`

HasIngressConfiguration returns a boolean if a field has been set.

### GetEnableIAMRole

`func (o *PrivateSpacePatchBody) GetEnableIAMRole() bool`

GetEnableIAMRole returns the EnableIAMRole field if non-nil, zero value otherwise.

### GetEnableIAMRoleOk

`func (o *PrivateSpacePatchBody) GetEnableIAMRoleOk() (*bool, bool)`

GetEnableIAMRoleOk returns a tuple with the EnableIAMRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIAMRole

`func (o *PrivateSpacePatchBody) SetEnableIAMRole(v bool)`

SetEnableIAMRole sets EnableIAMRole field to given value.

### HasEnableIAMRole

`func (o *PrivateSpacePatchBody) HasEnableIAMRole() bool`

HasEnableIAMRole returns a boolean if a field has been set.

### GetEnableEgress

`func (o *PrivateSpacePatchBody) GetEnableEgress() bool`

GetEnableEgress returns the EnableEgress field if non-nil, zero value otherwise.

### GetEnableEgressOk

`func (o *PrivateSpacePatchBody) GetEnableEgressOk() (*bool, bool)`

GetEnableEgressOk returns a tuple with the EnableEgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEgress

`func (o *PrivateSpacePatchBody) SetEnableEgress(v bool)`

SetEnableEgress sets EnableEgress field to given value.

### HasEnableEgress

`func (o *PrivateSpacePatchBody) HasEnableEgress() bool`

HasEnableEgress returns a boolean if a field has been set.

### GetEnableNetworkIsolation

`func (o *PrivateSpacePatchBody) GetEnableNetworkIsolation() bool`

GetEnableNetworkIsolation returns the EnableNetworkIsolation field if non-nil, zero value otherwise.

### GetEnableNetworkIsolationOk

`func (o *PrivateSpacePatchBody) GetEnableNetworkIsolationOk() (*bool, bool)`

GetEnableNetworkIsolationOk returns a tuple with the EnableNetworkIsolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNetworkIsolation

`func (o *PrivateSpacePatchBody) SetEnableNetworkIsolation(v bool)`

SetEnableNetworkIsolation sets EnableNetworkIsolation field to given value.

### HasEnableNetworkIsolation

`func (o *PrivateSpacePatchBody) HasEnableNetworkIsolation() bool`

HasEnableNetworkIsolation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


