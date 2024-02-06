# VpcCore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDefault** | Pointer to **bool** |  | [optional] [default to false]
**Name** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**SharedWith** | Pointer to **[]string** |  | [optional] [default to []]
**AssociatedEnvironments** | Pointer to **[]string** |  | [optional] [default to []]
**CidrBlock** | Pointer to **string** | An explanation about the purpose of this instance. | [optional] [default to "10.0.0.0/20"]
**FirewallRules** | Pointer to [**[]FirewallRule**](FirewallRule.md) |  | [optional] [default to []]
**InternalDns** | Pointer to [**InternalDns**](InternalDns.md) |  | [optional] [default to {"dnsServers":[],"specialDomains":[]}]
**VpcRoutes** | Pointer to [**[]VpcRoute**](VpcRoute.md) |  | [optional] [default to []]

## Methods

### NewVpcCore

`func NewVpcCore() *VpcCore`

NewVpcCore instantiates a new VpcCore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpcCoreWithDefaults

`func NewVpcCoreWithDefaults() *VpcCore`

NewVpcCoreWithDefaults instantiates a new VpcCore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDefault

`func (o *VpcCore) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *VpcCore) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *VpcCore) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *VpcCore) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetName

`func (o *VpcCore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpcCore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpcCore) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VpcCore) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwnerId

`func (o *VpcCore) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *VpcCore) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *VpcCore) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *VpcCore) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetRegion

`func (o *VpcCore) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *VpcCore) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *VpcCore) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *VpcCore) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSharedWith

`func (o *VpcCore) GetSharedWith() []string`

GetSharedWith returns the SharedWith field if non-nil, zero value otherwise.

### GetSharedWithOk

`func (o *VpcCore) GetSharedWithOk() (*[]string, bool)`

GetSharedWithOk returns a tuple with the SharedWith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedWith

`func (o *VpcCore) SetSharedWith(v []string)`

SetSharedWith sets SharedWith field to given value.

### HasSharedWith

`func (o *VpcCore) HasSharedWith() bool`

HasSharedWith returns a boolean if a field has been set.

### GetAssociatedEnvironments

`func (o *VpcCore) GetAssociatedEnvironments() []string`

GetAssociatedEnvironments returns the AssociatedEnvironments field if non-nil, zero value otherwise.

### GetAssociatedEnvironmentsOk

`func (o *VpcCore) GetAssociatedEnvironmentsOk() (*[]string, bool)`

GetAssociatedEnvironmentsOk returns a tuple with the AssociatedEnvironments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedEnvironments

`func (o *VpcCore) SetAssociatedEnvironments(v []string)`

SetAssociatedEnvironments sets AssociatedEnvironments field to given value.

### HasAssociatedEnvironments

`func (o *VpcCore) HasAssociatedEnvironments() bool`

HasAssociatedEnvironments returns a boolean if a field has been set.

### GetCidrBlock

`func (o *VpcCore) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *VpcCore) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *VpcCore) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.

### HasCidrBlock

`func (o *VpcCore) HasCidrBlock() bool`

HasCidrBlock returns a boolean if a field has been set.

### GetFirewallRules

`func (o *VpcCore) GetFirewallRules() []FirewallRule`

GetFirewallRules returns the FirewallRules field if non-nil, zero value otherwise.

### GetFirewallRulesOk

`func (o *VpcCore) GetFirewallRulesOk() (*[]FirewallRule, bool)`

GetFirewallRulesOk returns a tuple with the FirewallRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallRules

`func (o *VpcCore) SetFirewallRules(v []FirewallRule)`

SetFirewallRules sets FirewallRules field to given value.

### HasFirewallRules

`func (o *VpcCore) HasFirewallRules() bool`

HasFirewallRules returns a boolean if a field has been set.

### GetInternalDns

`func (o *VpcCore) GetInternalDns() InternalDns`

GetInternalDns returns the InternalDns field if non-nil, zero value otherwise.

### GetInternalDnsOk

`func (o *VpcCore) GetInternalDnsOk() (*InternalDns, bool)`

GetInternalDnsOk returns a tuple with the InternalDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalDns

`func (o *VpcCore) SetInternalDns(v InternalDns)`

SetInternalDns sets InternalDns field to given value.

### HasInternalDns

`func (o *VpcCore) HasInternalDns() bool`

HasInternalDns returns a boolean if a field has been set.

### GetVpcRoutes

`func (o *VpcCore) GetVpcRoutes() []VpcRoute`

GetVpcRoutes returns the VpcRoutes field if non-nil, zero value otherwise.

### GetVpcRoutesOk

`func (o *VpcCore) GetVpcRoutesOk() (*[]VpcRoute, bool)`

GetVpcRoutesOk returns a tuple with the VpcRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcRoutes

`func (o *VpcCore) SetVpcRoutes(v []VpcRoute)`

SetVpcRoutes sets VpcRoutes field to given value.

### HasVpcRoutes

`func (o *VpcCore) HasVpcRoutes() bool`

HasVpcRoutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


