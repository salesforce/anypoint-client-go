# VpnPostReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**RemoteAsn** | Pointer to **int32** |  | [optional] 
**RemoteIpAddress** | Pointer to **string** |  | [optional] 
**RemoteNetworks** | Pointer to **[]string** |  | [optional] [default to []]
**TunnelConfigs** | Pointer to [**[]TunnelConfig**](TunnelConfig.md) |  | [optional] 

## Methods

### NewVpnPostReqBody

`func NewVpnPostReqBody() *VpnPostReqBody`

NewVpnPostReqBody instantiates a new VpnPostReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnPostReqBodyWithDefaults

`func NewVpnPostReqBodyWithDefaults() *VpnPostReqBody`

NewVpnPostReqBodyWithDefaults instantiates a new VpnPostReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VpnPostReqBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpnPostReqBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpnPostReqBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VpnPostReqBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRemoteAsn

`func (o *VpnPostReqBody) GetRemoteAsn() int32`

GetRemoteAsn returns the RemoteAsn field if non-nil, zero value otherwise.

### GetRemoteAsnOk

`func (o *VpnPostReqBody) GetRemoteAsnOk() (*int32, bool)`

GetRemoteAsnOk returns a tuple with the RemoteAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAsn

`func (o *VpnPostReqBody) SetRemoteAsn(v int32)`

SetRemoteAsn sets RemoteAsn field to given value.

### HasRemoteAsn

`func (o *VpnPostReqBody) HasRemoteAsn() bool`

HasRemoteAsn returns a boolean if a field has been set.

### GetRemoteIpAddress

`func (o *VpnPostReqBody) GetRemoteIpAddress() string`

GetRemoteIpAddress returns the RemoteIpAddress field if non-nil, zero value otherwise.

### GetRemoteIpAddressOk

`func (o *VpnPostReqBody) GetRemoteIpAddressOk() (*string, bool)`

GetRemoteIpAddressOk returns a tuple with the RemoteIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIpAddress

`func (o *VpnPostReqBody) SetRemoteIpAddress(v string)`

SetRemoteIpAddress sets RemoteIpAddress field to given value.

### HasRemoteIpAddress

`func (o *VpnPostReqBody) HasRemoteIpAddress() bool`

HasRemoteIpAddress returns a boolean if a field has been set.

### GetRemoteNetworks

`func (o *VpnPostReqBody) GetRemoteNetworks() []string`

GetRemoteNetworks returns the RemoteNetworks field if non-nil, zero value otherwise.

### GetRemoteNetworksOk

`func (o *VpnPostReqBody) GetRemoteNetworksOk() (*[]string, bool)`

GetRemoteNetworksOk returns a tuple with the RemoteNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteNetworks

`func (o *VpnPostReqBody) SetRemoteNetworks(v []string)`

SetRemoteNetworks sets RemoteNetworks field to given value.

### HasRemoteNetworks

`func (o *VpnPostReqBody) HasRemoteNetworks() bool`

HasRemoteNetworks returns a boolean if a field has been set.

### GetTunnelConfigs

`func (o *VpnPostReqBody) GetTunnelConfigs() []TunnelConfig`

GetTunnelConfigs returns the TunnelConfigs field if non-nil, zero value otherwise.

### GetTunnelConfigsOk

`func (o *VpnPostReqBody) GetTunnelConfigsOk() (*[]TunnelConfig, bool)`

GetTunnelConfigsOk returns a tuple with the TunnelConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelConfigs

`func (o *VpnPostReqBody) SetTunnelConfigs(v []TunnelConfig)`

SetTunnelConfigs sets TunnelConfigs field to given value.

### HasTunnelConfigs

`func (o *VpnPostReqBody) HasTunnelConfigs() bool`

HasTunnelConfigs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


