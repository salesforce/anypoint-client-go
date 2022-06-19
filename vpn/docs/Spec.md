# Spec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**RemoteAsn** | Pointer to **float32** |  | [optional] 
**RemoteIpAddress** | **string** |  | 
**TunnelConfigs** | Pointer to [**[]TunnelConfig1**](TunnelConfig1.md) |  | [optional] 
**RemoteNetworks** | **[]string** |  | [default to []]

## Methods

### NewSpec

`func NewSpec(name string, remoteIpAddress string, remoteNetworks []string, ) *Spec`

NewSpec instantiates a new Spec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpecWithDefaults

`func NewSpecWithDefaults() *Spec`

NewSpecWithDefaults instantiates a new Spec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Spec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Spec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Spec) SetName(v string)`

SetName sets Name field to given value.


### GetRemoteAsn

`func (o *Spec) GetRemoteAsn() float32`

GetRemoteAsn returns the RemoteAsn field if non-nil, zero value otherwise.

### GetRemoteAsnOk

`func (o *Spec) GetRemoteAsnOk() (*float32, bool)`

GetRemoteAsnOk returns a tuple with the RemoteAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAsn

`func (o *Spec) SetRemoteAsn(v float32)`

SetRemoteAsn sets RemoteAsn field to given value.

### HasRemoteAsn

`func (o *Spec) HasRemoteAsn() bool`

HasRemoteAsn returns a boolean if a field has been set.

### GetRemoteIpAddress

`func (o *Spec) GetRemoteIpAddress() string`

GetRemoteIpAddress returns the RemoteIpAddress field if non-nil, zero value otherwise.

### GetRemoteIpAddressOk

`func (o *Spec) GetRemoteIpAddressOk() (*string, bool)`

GetRemoteIpAddressOk returns a tuple with the RemoteIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIpAddress

`func (o *Spec) SetRemoteIpAddress(v string)`

SetRemoteIpAddress sets RemoteIpAddress field to given value.


### GetTunnelConfigs

`func (o *Spec) GetTunnelConfigs() []TunnelConfig1`

GetTunnelConfigs returns the TunnelConfigs field if non-nil, zero value otherwise.

### GetTunnelConfigsOk

`func (o *Spec) GetTunnelConfigsOk() (*[]TunnelConfig1, bool)`

GetTunnelConfigsOk returns a tuple with the TunnelConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelConfigs

`func (o *Spec) SetTunnelConfigs(v []TunnelConfig1)`

SetTunnelConfigs sets TunnelConfigs field to given value.

### HasTunnelConfigs

`func (o *Spec) HasTunnelConfigs() bool`

HasTunnelConfigs returns a boolean if a field has been set.

### GetRemoteNetworks

`func (o *Spec) GetRemoteNetworks() []string`

GetRemoteNetworks returns the RemoteNetworks field if non-nil, zero value otherwise.

### GetRemoteNetworksOk

`func (o *Spec) GetRemoteNetworksOk() (*[]string, bool)`

GetRemoteNetworksOk returns a tuple with the RemoteNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteNetworks

`func (o *Spec) SetRemoteNetworks(v []string)`

SetRemoteNetworks sets RemoteNetworks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


