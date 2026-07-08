# PrivateSpaceNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | Pointer to **string** | The region of the Private Space. Required when create a Private Space network. | [optional] 
**CidrBlock** | Pointer to **string** | The CIDR block of the VPC where the Private Space will be created. Required when create a Private Space in a new VPC. Default is 10.0.0.0/16. | [optional] 
**InternalDns** | Pointer to [**PrivateSpaceNetworkCreateInternalDns**](PrivateSpaceNetworkCreateInternalDns.md) |  | [optional] 
**InboundStaticIps** | Pointer to **[]string** | The inbound static IPs of the Private Space network. | [optional] 
**OutboundStaticIps** | Pointer to **[]string** | The outbound static IPs of the Private Space network. | [optional] 
**DnsTarget** | Pointer to **string** | The DNS target of the Private Space network. | [optional] 
**InternalDnsTarget** | Pointer to **string** | The internal DNS target of the Private Space network. | [optional] 
**ReservedCidrs** | Pointer to **[]string** | The list of reserved CIDR blocks for your private space to prevent IP address overlap when you connect your private space to your corporate network. | [optional] 

## Methods

### NewPrivateSpaceNetwork

`func NewPrivateSpaceNetwork() *PrivateSpaceNetwork`

NewPrivateSpaceNetwork instantiates a new PrivateSpaceNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateSpaceNetworkWithDefaults

`func NewPrivateSpaceNetworkWithDefaults() *PrivateSpaceNetwork`

NewPrivateSpaceNetworkWithDefaults instantiates a new PrivateSpaceNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *PrivateSpaceNetwork) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PrivateSpaceNetwork) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PrivateSpaceNetwork) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *PrivateSpaceNetwork) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetCidrBlock

`func (o *PrivateSpaceNetwork) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *PrivateSpaceNetwork) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *PrivateSpaceNetwork) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.

### HasCidrBlock

`func (o *PrivateSpaceNetwork) HasCidrBlock() bool`

HasCidrBlock returns a boolean if a field has been set.

### GetInternalDns

`func (o *PrivateSpaceNetwork) GetInternalDns() PrivateSpaceNetworkCreateInternalDns`

GetInternalDns returns the InternalDns field if non-nil, zero value otherwise.

### GetInternalDnsOk

`func (o *PrivateSpaceNetwork) GetInternalDnsOk() (*PrivateSpaceNetworkCreateInternalDns, bool)`

GetInternalDnsOk returns a tuple with the InternalDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalDns

`func (o *PrivateSpaceNetwork) SetInternalDns(v PrivateSpaceNetworkCreateInternalDns)`

SetInternalDns sets InternalDns field to given value.

### HasInternalDns

`func (o *PrivateSpaceNetwork) HasInternalDns() bool`

HasInternalDns returns a boolean if a field has been set.

### GetInboundStaticIps

`func (o *PrivateSpaceNetwork) GetInboundStaticIps() []string`

GetInboundStaticIps returns the InboundStaticIps field if non-nil, zero value otherwise.

### GetInboundStaticIpsOk

`func (o *PrivateSpaceNetwork) GetInboundStaticIpsOk() (*[]string, bool)`

GetInboundStaticIpsOk returns a tuple with the InboundStaticIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundStaticIps

`func (o *PrivateSpaceNetwork) SetInboundStaticIps(v []string)`

SetInboundStaticIps sets InboundStaticIps field to given value.

### HasInboundStaticIps

`func (o *PrivateSpaceNetwork) HasInboundStaticIps() bool`

HasInboundStaticIps returns a boolean if a field has been set.

### GetOutboundStaticIps

`func (o *PrivateSpaceNetwork) GetOutboundStaticIps() []string`

GetOutboundStaticIps returns the OutboundStaticIps field if non-nil, zero value otherwise.

### GetOutboundStaticIpsOk

`func (o *PrivateSpaceNetwork) GetOutboundStaticIpsOk() (*[]string, bool)`

GetOutboundStaticIpsOk returns a tuple with the OutboundStaticIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundStaticIps

`func (o *PrivateSpaceNetwork) SetOutboundStaticIps(v []string)`

SetOutboundStaticIps sets OutboundStaticIps field to given value.

### HasOutboundStaticIps

`func (o *PrivateSpaceNetwork) HasOutboundStaticIps() bool`

HasOutboundStaticIps returns a boolean if a field has been set.

### GetDnsTarget

`func (o *PrivateSpaceNetwork) GetDnsTarget() string`

GetDnsTarget returns the DnsTarget field if non-nil, zero value otherwise.

### GetDnsTargetOk

`func (o *PrivateSpaceNetwork) GetDnsTargetOk() (*string, bool)`

GetDnsTargetOk returns a tuple with the DnsTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsTarget

`func (o *PrivateSpaceNetwork) SetDnsTarget(v string)`

SetDnsTarget sets DnsTarget field to given value.

### HasDnsTarget

`func (o *PrivateSpaceNetwork) HasDnsTarget() bool`

HasDnsTarget returns a boolean if a field has been set.

### GetInternalDnsTarget

`func (o *PrivateSpaceNetwork) GetInternalDnsTarget() string`

GetInternalDnsTarget returns the InternalDnsTarget field if non-nil, zero value otherwise.

### GetInternalDnsTargetOk

`func (o *PrivateSpaceNetwork) GetInternalDnsTargetOk() (*string, bool)`

GetInternalDnsTargetOk returns a tuple with the InternalDnsTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalDnsTarget

`func (o *PrivateSpaceNetwork) SetInternalDnsTarget(v string)`

SetInternalDnsTarget sets InternalDnsTarget field to given value.

### HasInternalDnsTarget

`func (o *PrivateSpaceNetwork) HasInternalDnsTarget() bool`

HasInternalDnsTarget returns a boolean if a field has been set.

### GetReservedCidrs

`func (o *PrivateSpaceNetwork) GetReservedCidrs() []string`

GetReservedCidrs returns the ReservedCidrs field if non-nil, zero value otherwise.

### GetReservedCidrsOk

`func (o *PrivateSpaceNetwork) GetReservedCidrsOk() (*[]string, bool)`

GetReservedCidrsOk returns a tuple with the ReservedCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedCidrs

`func (o *PrivateSpaceNetwork) SetReservedCidrs(v []string)`

SetReservedCidrs sets ReservedCidrs field to given value.

### HasReservedCidrs

`func (o *PrivateSpaceNetwork) HasReservedCidrs() bool`

HasReservedCidrs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


