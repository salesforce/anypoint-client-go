# PrivateSpaceVpn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | User-chosen display name for the VPN member. | [optional] 
**VpnId** | Pointer to **string** | VPN member id (GUID). | [optional] 
**ConnectionId** | Pointer to **string** | Parent connection id (GUID). | [optional] 
**ConnectionName** | Pointer to **string** | Parent connection display name. | [optional] 
**VpnConnectionStatus** | Pointer to **string** | Current status of this VPN member. | [optional] 
**LocalAsn** | Pointer to **int32** | Local BGP ASN announced by Anypoint. | [optional] 
**RemoteAsn** | Pointer to **int32** | Remote BGP ASN announced by the customer endpoint. | [optional] 
**RemoteIpAddress** | Pointer to **string** | Public IP address of the customer VPN endpoint. | [optional] 
**StaticRoutes** | Pointer to **[]string** | Static routes carried over the tunnel (CIDR strings). Empty when BGP is used. | [optional] 
**VpnTunnels** | Pointer to [**[]PrivateSpaceVpnTunnel**](PrivateSpaceVpnTunnel.md) | IPsec tunnels composing the VPN member (typically two for high availability). | [optional] 

## Methods

### NewPrivateSpaceVpn

`func NewPrivateSpaceVpn() *PrivateSpaceVpn`

NewPrivateSpaceVpn instantiates a new PrivateSpaceVpn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateSpaceVpnWithDefaults

`func NewPrivateSpaceVpnWithDefaults() *PrivateSpaceVpn`

NewPrivateSpaceVpnWithDefaults instantiates a new PrivateSpaceVpn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PrivateSpaceVpn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrivateSpaceVpn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrivateSpaceVpn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PrivateSpaceVpn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVpnId

`func (o *PrivateSpaceVpn) GetVpnId() string`

GetVpnId returns the VpnId field if non-nil, zero value otherwise.

### GetVpnIdOk

`func (o *PrivateSpaceVpn) GetVpnIdOk() (*string, bool)`

GetVpnIdOk returns a tuple with the VpnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnId

`func (o *PrivateSpaceVpn) SetVpnId(v string)`

SetVpnId sets VpnId field to given value.

### HasVpnId

`func (o *PrivateSpaceVpn) HasVpnId() bool`

HasVpnId returns a boolean if a field has been set.

### GetConnectionId

`func (o *PrivateSpaceVpn) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *PrivateSpaceVpn) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *PrivateSpaceVpn) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *PrivateSpaceVpn) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetConnectionName

`func (o *PrivateSpaceVpn) GetConnectionName() string`

GetConnectionName returns the ConnectionName field if non-nil, zero value otherwise.

### GetConnectionNameOk

`func (o *PrivateSpaceVpn) GetConnectionNameOk() (*string, bool)`

GetConnectionNameOk returns a tuple with the ConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionName

`func (o *PrivateSpaceVpn) SetConnectionName(v string)`

SetConnectionName sets ConnectionName field to given value.

### HasConnectionName

`func (o *PrivateSpaceVpn) HasConnectionName() bool`

HasConnectionName returns a boolean if a field has been set.

### GetVpnConnectionStatus

`func (o *PrivateSpaceVpn) GetVpnConnectionStatus() string`

GetVpnConnectionStatus returns the VpnConnectionStatus field if non-nil, zero value otherwise.

### GetVpnConnectionStatusOk

`func (o *PrivateSpaceVpn) GetVpnConnectionStatusOk() (*string, bool)`

GetVpnConnectionStatusOk returns a tuple with the VpnConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnConnectionStatus

`func (o *PrivateSpaceVpn) SetVpnConnectionStatus(v string)`

SetVpnConnectionStatus sets VpnConnectionStatus field to given value.

### HasVpnConnectionStatus

`func (o *PrivateSpaceVpn) HasVpnConnectionStatus() bool`

HasVpnConnectionStatus returns a boolean if a field has been set.

### GetLocalAsn

`func (o *PrivateSpaceVpn) GetLocalAsn() int32`

GetLocalAsn returns the LocalAsn field if non-nil, zero value otherwise.

### GetLocalAsnOk

`func (o *PrivateSpaceVpn) GetLocalAsnOk() (*int32, bool)`

GetLocalAsnOk returns a tuple with the LocalAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAsn

`func (o *PrivateSpaceVpn) SetLocalAsn(v int32)`

SetLocalAsn sets LocalAsn field to given value.

### HasLocalAsn

`func (o *PrivateSpaceVpn) HasLocalAsn() bool`

HasLocalAsn returns a boolean if a field has been set.

### GetRemoteAsn

`func (o *PrivateSpaceVpn) GetRemoteAsn() int32`

GetRemoteAsn returns the RemoteAsn field if non-nil, zero value otherwise.

### GetRemoteAsnOk

`func (o *PrivateSpaceVpn) GetRemoteAsnOk() (*int32, bool)`

GetRemoteAsnOk returns a tuple with the RemoteAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAsn

`func (o *PrivateSpaceVpn) SetRemoteAsn(v int32)`

SetRemoteAsn sets RemoteAsn field to given value.

### HasRemoteAsn

`func (o *PrivateSpaceVpn) HasRemoteAsn() bool`

HasRemoteAsn returns a boolean if a field has been set.

### GetRemoteIpAddress

`func (o *PrivateSpaceVpn) GetRemoteIpAddress() string`

GetRemoteIpAddress returns the RemoteIpAddress field if non-nil, zero value otherwise.

### GetRemoteIpAddressOk

`func (o *PrivateSpaceVpn) GetRemoteIpAddressOk() (*string, bool)`

GetRemoteIpAddressOk returns a tuple with the RemoteIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIpAddress

`func (o *PrivateSpaceVpn) SetRemoteIpAddress(v string)`

SetRemoteIpAddress sets RemoteIpAddress field to given value.

### HasRemoteIpAddress

`func (o *PrivateSpaceVpn) HasRemoteIpAddress() bool`

HasRemoteIpAddress returns a boolean if a field has been set.

### GetStaticRoutes

`func (o *PrivateSpaceVpn) GetStaticRoutes() []string`

GetStaticRoutes returns the StaticRoutes field if non-nil, zero value otherwise.

### GetStaticRoutesOk

`func (o *PrivateSpaceVpn) GetStaticRoutesOk() (*[]string, bool)`

GetStaticRoutesOk returns a tuple with the StaticRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticRoutes

`func (o *PrivateSpaceVpn) SetStaticRoutes(v []string)`

SetStaticRoutes sets StaticRoutes field to given value.

### HasStaticRoutes

`func (o *PrivateSpaceVpn) HasStaticRoutes() bool`

HasStaticRoutes returns a boolean if a field has been set.

### GetVpnTunnels

`func (o *PrivateSpaceVpn) GetVpnTunnels() []PrivateSpaceVpnTunnel`

GetVpnTunnels returns the VpnTunnels field if non-nil, zero value otherwise.

### GetVpnTunnelsOk

`func (o *PrivateSpaceVpn) GetVpnTunnelsOk() (*[]PrivateSpaceVpnTunnel, bool)`

GetVpnTunnelsOk returns a tuple with the VpnTunnels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnTunnels

`func (o *PrivateSpaceVpn) SetVpnTunnels(v []PrivateSpaceVpnTunnel)`

SetVpnTunnels sets VpnTunnels field to given value.

### HasVpnTunnels

`func (o *PrivateSpaceVpn) HasVpnTunnels() bool`

HasVpnTunnels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


