# State

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VpnConnectionStatus** | **string** |  | [default to ""]
**CreatedAt** | Pointer to **string** |  | [optional] [default to ""]
**RemoteAsn** | Pointer to **int32** |  | [optional] [default to 0]
**LocalAsn** | Pointer to **int32** |  | [optional] [default to 0]
**VpnTunnels** | Pointer to [**[]VpnTunnel**](VpnTunnel.md) |  | [optional] 
**FailedReason** | Pointer to **string** |  | [optional] [default to ""]

## Methods

### NewState

`func NewState(vpnConnectionStatus string, ) *State`

NewState instantiates a new State object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStateWithDefaults

`func NewStateWithDefaults() *State`

NewStateWithDefaults instantiates a new State object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpnConnectionStatus

`func (o *State) GetVpnConnectionStatus() string`

GetVpnConnectionStatus returns the VpnConnectionStatus field if non-nil, zero value otherwise.

### GetVpnConnectionStatusOk

`func (o *State) GetVpnConnectionStatusOk() (*string, bool)`

GetVpnConnectionStatusOk returns a tuple with the VpnConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnConnectionStatus

`func (o *State) SetVpnConnectionStatus(v string)`

SetVpnConnectionStatus sets VpnConnectionStatus field to given value.


### GetCreatedAt

`func (o *State) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *State) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *State) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *State) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetRemoteAsn

`func (o *State) GetRemoteAsn() int32`

GetRemoteAsn returns the RemoteAsn field if non-nil, zero value otherwise.

### GetRemoteAsnOk

`func (o *State) GetRemoteAsnOk() (*int32, bool)`

GetRemoteAsnOk returns a tuple with the RemoteAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAsn

`func (o *State) SetRemoteAsn(v int32)`

SetRemoteAsn sets RemoteAsn field to given value.

### HasRemoteAsn

`func (o *State) HasRemoteAsn() bool`

HasRemoteAsn returns a boolean if a field has been set.

### GetLocalAsn

`func (o *State) GetLocalAsn() int32`

GetLocalAsn returns the LocalAsn field if non-nil, zero value otherwise.

### GetLocalAsnOk

`func (o *State) GetLocalAsnOk() (*int32, bool)`

GetLocalAsnOk returns a tuple with the LocalAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAsn

`func (o *State) SetLocalAsn(v int32)`

SetLocalAsn sets LocalAsn field to given value.

### HasLocalAsn

`func (o *State) HasLocalAsn() bool`

HasLocalAsn returns a boolean if a field has been set.

### GetVpnTunnels

`func (o *State) GetVpnTunnels() []VpnTunnel`

GetVpnTunnels returns the VpnTunnels field if non-nil, zero value otherwise.

### GetVpnTunnelsOk

`func (o *State) GetVpnTunnelsOk() (*[]VpnTunnel, bool)`

GetVpnTunnelsOk returns a tuple with the VpnTunnels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnTunnels

`func (o *State) SetVpnTunnels(v []VpnTunnel)`

SetVpnTunnels sets VpnTunnels field to given value.

### HasVpnTunnels

`func (o *State) HasVpnTunnels() bool`

HasVpnTunnels returns a boolean if a field has been set.

### GetFailedReason

`func (o *State) GetFailedReason() string`

GetFailedReason returns the FailedReason field if non-nil, zero value otherwise.

### GetFailedReasonOk

`func (o *State) GetFailedReasonOk() (*string, bool)`

GetFailedReasonOk returns a tuple with the FailedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedReason

`func (o *State) SetFailedReason(v string)`

SetFailedReason sets FailedReason field to given value.

### HasFailedReason

`func (o *State) HasFailedReason() bool`

HasFailedReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


