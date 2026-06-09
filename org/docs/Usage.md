# Usage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductionVCoresConsumed** | Pointer to **float32** | Production VCores consumed | [optional] [default to 0]
**SandboxVCoresConsumed** | Pointer to **float32** | Sandbox VCores consumed | [optional] [default to 0]
**DesignVCoresConsumed** | Pointer to **float32** | Design VCores consumed | [optional] [default to 0]
**StaticIpsConsumed** | Pointer to **float32** | Static IPs consumed | [optional] [default to 0]
**VpcsConsumed** | Pointer to **int32** | VPCs consumed | [optional] [default to 0]
**VpnsConsumed** | Pointer to **int32** | VPNs consumed | [optional] [default to 0]
**NetworkConnectionsConsumed** | Pointer to **int32** | Network connections consumed | [optional] [default to 0]
**LoadBalancersConsumed** | Pointer to **int32** | Load balancers consumed | [optional] [default to 0]
**LoadbalancerWorkersConsumed** | Pointer to **int32** | Load balancer workers consumed | [optional] [default to 0]

## Methods

### NewUsage

`func NewUsage() *Usage`

NewUsage instantiates a new Usage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageWithDefaults

`func NewUsageWithDefaults() *Usage`

NewUsageWithDefaults instantiates a new Usage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductionVCoresConsumed

`func (o *Usage) GetProductionVCoresConsumed() float32`

GetProductionVCoresConsumed returns the ProductionVCoresConsumed field if non-nil, zero value otherwise.

### GetProductionVCoresConsumedOk

`func (o *Usage) GetProductionVCoresConsumedOk() (*float32, bool)`

GetProductionVCoresConsumedOk returns a tuple with the ProductionVCoresConsumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductionVCoresConsumed

`func (o *Usage) SetProductionVCoresConsumed(v float32)`

SetProductionVCoresConsumed sets ProductionVCoresConsumed field to given value.

### HasProductionVCoresConsumed

`func (o *Usage) HasProductionVCoresConsumed() bool`

HasProductionVCoresConsumed returns a boolean if a field has been set.

### GetSandboxVCoresConsumed

`func (o *Usage) GetSandboxVCoresConsumed() float32`

GetSandboxVCoresConsumed returns the SandboxVCoresConsumed field if non-nil, zero value otherwise.

### GetSandboxVCoresConsumedOk

`func (o *Usage) GetSandboxVCoresConsumedOk() (*float32, bool)`

GetSandboxVCoresConsumedOk returns a tuple with the SandboxVCoresConsumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandboxVCoresConsumed

`func (o *Usage) SetSandboxVCoresConsumed(v float32)`

SetSandboxVCoresConsumed sets SandboxVCoresConsumed field to given value.

### HasSandboxVCoresConsumed

`func (o *Usage) HasSandboxVCoresConsumed() bool`

HasSandboxVCoresConsumed returns a boolean if a field has been set.

### GetDesignVCoresConsumed

`func (o *Usage) GetDesignVCoresConsumed() float32`

GetDesignVCoresConsumed returns the DesignVCoresConsumed field if non-nil, zero value otherwise.

### GetDesignVCoresConsumedOk

`func (o *Usage) GetDesignVCoresConsumedOk() (*float32, bool)`

GetDesignVCoresConsumedOk returns a tuple with the DesignVCoresConsumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesignVCoresConsumed

`func (o *Usage) SetDesignVCoresConsumed(v float32)`

SetDesignVCoresConsumed sets DesignVCoresConsumed field to given value.

### HasDesignVCoresConsumed

`func (o *Usage) HasDesignVCoresConsumed() bool`

HasDesignVCoresConsumed returns a boolean if a field has been set.

### GetStaticIpsConsumed

`func (o *Usage) GetStaticIpsConsumed() float32`

GetStaticIpsConsumed returns the StaticIpsConsumed field if non-nil, zero value otherwise.

### GetStaticIpsConsumedOk

`func (o *Usage) GetStaticIpsConsumedOk() (*float32, bool)`

GetStaticIpsConsumedOk returns a tuple with the StaticIpsConsumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticIpsConsumed

`func (o *Usage) SetStaticIpsConsumed(v float32)`

SetStaticIpsConsumed sets StaticIpsConsumed field to given value.

### HasStaticIpsConsumed

`func (o *Usage) HasStaticIpsConsumed() bool`

HasStaticIpsConsumed returns a boolean if a field has been set.

### GetVpcsConsumed

`func (o *Usage) GetVpcsConsumed() int32`

GetVpcsConsumed returns the VpcsConsumed field if non-nil, zero value otherwise.

### GetVpcsConsumedOk

`func (o *Usage) GetVpcsConsumedOk() (*int32, bool)`

GetVpcsConsumedOk returns a tuple with the VpcsConsumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcsConsumed

`func (o *Usage) SetVpcsConsumed(v int32)`

SetVpcsConsumed sets VpcsConsumed field to given value.

### HasVpcsConsumed

`func (o *Usage) HasVpcsConsumed() bool`

HasVpcsConsumed returns a boolean if a field has been set.

### GetVpnsConsumed

`func (o *Usage) GetVpnsConsumed() int32`

GetVpnsConsumed returns the VpnsConsumed field if non-nil, zero value otherwise.

### GetVpnsConsumedOk

`func (o *Usage) GetVpnsConsumedOk() (*int32, bool)`

GetVpnsConsumedOk returns a tuple with the VpnsConsumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnsConsumed

`func (o *Usage) SetVpnsConsumed(v int32)`

SetVpnsConsumed sets VpnsConsumed field to given value.

### HasVpnsConsumed

`func (o *Usage) HasVpnsConsumed() bool`

HasVpnsConsumed returns a boolean if a field has been set.

### GetNetworkConnectionsConsumed

`func (o *Usage) GetNetworkConnectionsConsumed() int32`

GetNetworkConnectionsConsumed returns the NetworkConnectionsConsumed field if non-nil, zero value otherwise.

### GetNetworkConnectionsConsumedOk

`func (o *Usage) GetNetworkConnectionsConsumedOk() (*int32, bool)`

GetNetworkConnectionsConsumedOk returns a tuple with the NetworkConnectionsConsumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConnectionsConsumed

`func (o *Usage) SetNetworkConnectionsConsumed(v int32)`

SetNetworkConnectionsConsumed sets NetworkConnectionsConsumed field to given value.

### HasNetworkConnectionsConsumed

`func (o *Usage) HasNetworkConnectionsConsumed() bool`

HasNetworkConnectionsConsumed returns a boolean if a field has been set.

### GetLoadBalancersConsumed

`func (o *Usage) GetLoadBalancersConsumed() int32`

GetLoadBalancersConsumed returns the LoadBalancersConsumed field if non-nil, zero value otherwise.

### GetLoadBalancersConsumedOk

`func (o *Usage) GetLoadBalancersConsumedOk() (*int32, bool)`

GetLoadBalancersConsumedOk returns a tuple with the LoadBalancersConsumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancersConsumed

`func (o *Usage) SetLoadBalancersConsumed(v int32)`

SetLoadBalancersConsumed sets LoadBalancersConsumed field to given value.

### HasLoadBalancersConsumed

`func (o *Usage) HasLoadBalancersConsumed() bool`

HasLoadBalancersConsumed returns a boolean if a field has been set.

### GetLoadbalancerWorkersConsumed

`func (o *Usage) GetLoadbalancerWorkersConsumed() int32`

GetLoadbalancerWorkersConsumed returns the LoadbalancerWorkersConsumed field if non-nil, zero value otherwise.

### GetLoadbalancerWorkersConsumedOk

`func (o *Usage) GetLoadbalancerWorkersConsumedOk() (*int32, bool)`

GetLoadbalancerWorkersConsumedOk returns a tuple with the LoadbalancerWorkersConsumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadbalancerWorkersConsumed

`func (o *Usage) SetLoadbalancerWorkersConsumed(v int32)`

SetLoadbalancerWorkersConsumed sets LoadbalancerWorkersConsumed field to given value.

### HasLoadbalancerWorkersConsumed

`func (o *Usage) HasLoadbalancerWorkersConsumed() bool`

HasLoadbalancerWorkersConsumed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


