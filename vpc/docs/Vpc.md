# Vpc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The vpc Id | 
**IsDefault** | Pointer to **bool** |  | [optional] [default to false]
**Name** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**SharedWith** | Pointer to **[]string** |  | [optional] [default to []]
**AssociatedEnvironments** | Pointer to **[]string** |  | [optional] [default to []]
**CidrBlock** | Pointer to **string** | An explanation about the purpose of this instance. | [optional] [default to "10.0.0.0/20"]
**FirewallRules** | Pointer to [**[]FirewallRule**](FirewallRule.md) |  | [optional] [default to []]
**InternalDns** | Pointer to [**InternalDns**](InternalDns.md) |  | [optional] 
**VpcRoutes** | Pointer to [**[]VpcRoute**](VpcRoute.md) |  | [optional] [default to []]

## Methods

### NewVpc

`func NewVpc(id string, ) *Vpc`

NewVpc instantiates a new Vpc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpcWithDefaults

`func NewVpcWithDefaults() *Vpc`

NewVpcWithDefaults instantiates a new Vpc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Vpc) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Vpc) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Vpc) SetId(v string)`

SetId sets Id field to given value.


### GetIsDefault

`func (o *Vpc) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *Vpc) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *Vpc) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *Vpc) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetName

`func (o *Vpc) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Vpc) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Vpc) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Vpc) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwnerId

`func (o *Vpc) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *Vpc) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *Vpc) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *Vpc) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetRegion

`func (o *Vpc) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Vpc) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Vpc) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Vpc) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSharedWith

`func (o *Vpc) GetSharedWith() []string`

GetSharedWith returns the SharedWith field if non-nil, zero value otherwise.

### GetSharedWithOk

`func (o *Vpc) GetSharedWithOk() (*[]string, bool)`

GetSharedWithOk returns a tuple with the SharedWith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedWith

`func (o *Vpc) SetSharedWith(v []string)`

SetSharedWith sets SharedWith field to given value.

### HasSharedWith

`func (o *Vpc) HasSharedWith() bool`

HasSharedWith returns a boolean if a field has been set.

### GetAssociatedEnvironments

`func (o *Vpc) GetAssociatedEnvironments() []string`

GetAssociatedEnvironments returns the AssociatedEnvironments field if non-nil, zero value otherwise.

### GetAssociatedEnvironmentsOk

`func (o *Vpc) GetAssociatedEnvironmentsOk() (*[]string, bool)`

GetAssociatedEnvironmentsOk returns a tuple with the AssociatedEnvironments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedEnvironments

`func (o *Vpc) SetAssociatedEnvironments(v []string)`

SetAssociatedEnvironments sets AssociatedEnvironments field to given value.

### HasAssociatedEnvironments

`func (o *Vpc) HasAssociatedEnvironments() bool`

HasAssociatedEnvironments returns a boolean if a field has been set.

### GetCidrBlock

`func (o *Vpc) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *Vpc) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *Vpc) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.

### HasCidrBlock

`func (o *Vpc) HasCidrBlock() bool`

HasCidrBlock returns a boolean if a field has been set.

### GetFirewallRules

`func (o *Vpc) GetFirewallRules() []FirewallRule`

GetFirewallRules returns the FirewallRules field if non-nil, zero value otherwise.

### GetFirewallRulesOk

`func (o *Vpc) GetFirewallRulesOk() (*[]FirewallRule, bool)`

GetFirewallRulesOk returns a tuple with the FirewallRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallRules

`func (o *Vpc) SetFirewallRules(v []FirewallRule)`

SetFirewallRules sets FirewallRules field to given value.

### HasFirewallRules

`func (o *Vpc) HasFirewallRules() bool`

HasFirewallRules returns a boolean if a field has been set.

### GetInternalDns

`func (o *Vpc) GetInternalDns() InternalDns`

GetInternalDns returns the InternalDns field if non-nil, zero value otherwise.

### GetInternalDnsOk

`func (o *Vpc) GetInternalDnsOk() (*InternalDns, bool)`

GetInternalDnsOk returns a tuple with the InternalDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalDns

`func (o *Vpc) SetInternalDns(v InternalDns)`

SetInternalDns sets InternalDns field to given value.

### HasInternalDns

`func (o *Vpc) HasInternalDns() bool`

HasInternalDns returns a boolean if a field has been set.

### GetVpcRoutes

`func (o *Vpc) GetVpcRoutes() []VpcRoute`

GetVpcRoutes returns the VpcRoutes field if non-nil, zero value otherwise.

### GetVpcRoutesOk

`func (o *Vpc) GetVpcRoutesOk() (*[]VpcRoute, bool)`

GetVpcRoutesOk returns a tuple with the VpcRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcRoutes

`func (o *Vpc) SetVpcRoutes(v []VpcRoute)`

SetVpcRoutes sets VpcRoutes field to given value.

### HasVpcRoutes

`func (o *Vpc) HasVpcRoutes() bool`

HasVpcRoutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


