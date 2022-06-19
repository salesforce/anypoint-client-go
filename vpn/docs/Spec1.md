# Spec1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**RemoteAsn** | **int32** |  | 
**RemoteIpAddress** | **string** |  | 
**TunnelConfigs** | [**[]TunnelConfig1**](TunnelConfig1.md) |  | 
**RemoteNetworks** | Pointer to **[]string** |  | [optional] [default to []]

## Methods

### NewSpec1

`func NewSpec1(name string, remoteAsn int32, remoteIpAddress string, tunnelConfigs []TunnelConfig1, ) *Spec1`

NewSpec1 instantiates a new Spec1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpec1WithDefaults

`func NewSpec1WithDefaults() *Spec1`

NewSpec1WithDefaults instantiates a new Spec1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Spec1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Spec1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Spec1) SetName(v string)`

SetName sets Name field to given value.


### GetRemoteAsn

`func (o *Spec1) GetRemoteAsn() int32`

GetRemoteAsn returns the RemoteAsn field if non-nil, zero value otherwise.

### GetRemoteAsnOk

`func (o *Spec1) GetRemoteAsnOk() (*int32, bool)`

GetRemoteAsnOk returns a tuple with the RemoteAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAsn

`func (o *Spec1) SetRemoteAsn(v int32)`

SetRemoteAsn sets RemoteAsn field to given value.


### GetRemoteIpAddress

`func (o *Spec1) GetRemoteIpAddress() string`

GetRemoteIpAddress returns the RemoteIpAddress field if non-nil, zero value otherwise.

### GetRemoteIpAddressOk

`func (o *Spec1) GetRemoteIpAddressOk() (*string, bool)`

GetRemoteIpAddressOk returns a tuple with the RemoteIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIpAddress

`func (o *Spec1) SetRemoteIpAddress(v string)`

SetRemoteIpAddress sets RemoteIpAddress field to given value.


### GetTunnelConfigs

`func (o *Spec1) GetTunnelConfigs() []TunnelConfig1`

GetTunnelConfigs returns the TunnelConfigs field if non-nil, zero value otherwise.

### GetTunnelConfigsOk

`func (o *Spec1) GetTunnelConfigsOk() (*[]TunnelConfig1, bool)`

GetTunnelConfigsOk returns a tuple with the TunnelConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelConfigs

`func (o *Spec1) SetTunnelConfigs(v []TunnelConfig1)`

SetTunnelConfigs sets TunnelConfigs field to given value.


### GetRemoteNetworks

`func (o *Spec1) GetRemoteNetworks() []string`

GetRemoteNetworks returns the RemoteNetworks field if non-nil, zero value otherwise.

### GetRemoteNetworksOk

`func (o *Spec1) GetRemoteNetworksOk() (*[]string, bool)`

GetRemoteNetworksOk returns a tuple with the RemoteNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteNetworks

`func (o *Spec1) SetRemoteNetworks(v []string)`

SetRemoteNetworks sets RemoteNetworks field to given value.

### HasRemoteNetworks

`func (o *Spec1) HasRemoteNetworks() bool`

HasRemoteNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


