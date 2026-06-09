# PrivateSpaceNetworkCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | Pointer to **string** | The region of the Private Space. Required when create a Private Space network. | [optional] 
**CidrBlock** | Pointer to **string** | The CIDR block of the VPC where the Private Space will be created. Required when create a Private Space in a new VPC. Default is 10.0.0.0/16. | [optional] 
**InternalDns** | Pointer to [**PrivateSpaceNetworkCreateInternalDns**](PrivateSpaceNetworkCreateInternalDns.md) |  | [optional] 
**ReservedCidrs** | Pointer to **[]string** | The list of reserved CIDR blocks for your private space to prevent IP address overlap. Required when you want to connect your private space to your corporate network (either on-premises or in the cloud). Use CIDR notation and commas.  | [optional] 

## Methods

### NewPrivateSpaceNetworkCreate

`func NewPrivateSpaceNetworkCreate() *PrivateSpaceNetworkCreate`

NewPrivateSpaceNetworkCreate instantiates a new PrivateSpaceNetworkCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateSpaceNetworkCreateWithDefaults

`func NewPrivateSpaceNetworkCreateWithDefaults() *PrivateSpaceNetworkCreate`

NewPrivateSpaceNetworkCreateWithDefaults instantiates a new PrivateSpaceNetworkCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *PrivateSpaceNetworkCreate) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PrivateSpaceNetworkCreate) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PrivateSpaceNetworkCreate) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *PrivateSpaceNetworkCreate) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetCidrBlock

`func (o *PrivateSpaceNetworkCreate) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *PrivateSpaceNetworkCreate) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *PrivateSpaceNetworkCreate) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.

### HasCidrBlock

`func (o *PrivateSpaceNetworkCreate) HasCidrBlock() bool`

HasCidrBlock returns a boolean if a field has been set.

### GetInternalDns

`func (o *PrivateSpaceNetworkCreate) GetInternalDns() PrivateSpaceNetworkCreateInternalDns`

GetInternalDns returns the InternalDns field if non-nil, zero value otherwise.

### GetInternalDnsOk

`func (o *PrivateSpaceNetworkCreate) GetInternalDnsOk() (*PrivateSpaceNetworkCreateInternalDns, bool)`

GetInternalDnsOk returns a tuple with the InternalDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalDns

`func (o *PrivateSpaceNetworkCreate) SetInternalDns(v PrivateSpaceNetworkCreateInternalDns)`

SetInternalDns sets InternalDns field to given value.

### HasInternalDns

`func (o *PrivateSpaceNetworkCreate) HasInternalDns() bool`

HasInternalDns returns a boolean if a field has been set.

### GetReservedCidrs

`func (o *PrivateSpaceNetworkCreate) GetReservedCidrs() []string`

GetReservedCidrs returns the ReservedCidrs field if non-nil, zero value otherwise.

### GetReservedCidrsOk

`func (o *PrivateSpaceNetworkCreate) GetReservedCidrsOk() (*[]string, bool)`

GetReservedCidrsOk returns a tuple with the ReservedCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedCidrs

`func (o *PrivateSpaceNetworkCreate) SetReservedCidrs(v []string)`

SetReservedCidrs sets ReservedCidrs field to given value.

### HasReservedCidrs

`func (o *PrivateSpaceNetworkCreate) HasReservedCidrs() bool`

HasReservedCidrs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


