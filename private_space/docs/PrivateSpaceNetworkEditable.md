# PrivateSpaceNetworkEditable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InternalDns** | Pointer to [**PrivateSpaceNetworkCreateInternalDns**](PrivateSpaceNetworkCreateInternalDns.md) |  | [optional] 
**ReservedCidrs** | Pointer to **[]string** | The list of reserved CIDR blocks for your private space to prevent IP address overlap. Required when you want to connect your private space to your corporate network (either on-premises or in the cloud). Use CIDR notation and commas.  | [optional] 

## Methods

### NewPrivateSpaceNetworkEditable

`func NewPrivateSpaceNetworkEditable() *PrivateSpaceNetworkEditable`

NewPrivateSpaceNetworkEditable instantiates a new PrivateSpaceNetworkEditable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateSpaceNetworkEditableWithDefaults

`func NewPrivateSpaceNetworkEditableWithDefaults() *PrivateSpaceNetworkEditable`

NewPrivateSpaceNetworkEditableWithDefaults instantiates a new PrivateSpaceNetworkEditable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInternalDns

`func (o *PrivateSpaceNetworkEditable) GetInternalDns() PrivateSpaceNetworkCreateInternalDns`

GetInternalDns returns the InternalDns field if non-nil, zero value otherwise.

### GetInternalDnsOk

`func (o *PrivateSpaceNetworkEditable) GetInternalDnsOk() (*PrivateSpaceNetworkCreateInternalDns, bool)`

GetInternalDnsOk returns a tuple with the InternalDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalDns

`func (o *PrivateSpaceNetworkEditable) SetInternalDns(v PrivateSpaceNetworkCreateInternalDns)`

SetInternalDns sets InternalDns field to given value.

### HasInternalDns

`func (o *PrivateSpaceNetworkEditable) HasInternalDns() bool`

HasInternalDns returns a boolean if a field has been set.

### GetReservedCidrs

`func (o *PrivateSpaceNetworkEditable) GetReservedCidrs() []string`

GetReservedCidrs returns the ReservedCidrs field if non-nil, zero value otherwise.

### GetReservedCidrsOk

`func (o *PrivateSpaceNetworkEditable) GetReservedCidrsOk() (*[]string, bool)`

GetReservedCidrsOk returns a tuple with the ReservedCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedCidrs

`func (o *PrivateSpaceNetworkEditable) SetReservedCidrs(v []string)`

SetReservedCidrs sets ReservedCidrs field to given value.

### HasReservedCidrs

`func (o *PrivateSpaceNetworkEditable) HasReservedCidrs() bool`

HasReservedCidrs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


