# PrivateSpaceVpnPostBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Optional display name for the VPN member. | [optional] 
**LocalAsn** | **string** | Local BGP ASN announced by Anypoint (string-encoded integer). | 
**RemoteAsn** | Pointer to **string** | Remote BGP ASN announced by the customer endpoint (string-encoded integer). Omit when using static routes. | [optional] 
**RemoteIpAddress** | **string** | Public IP address of the customer VPN endpoint. | 
**StaticRoutes** | Pointer to **[]string** | Static routes carried over the tunnel (CIDR strings). Mutually exclusive with BGP. | [optional] 
**VpnTunnels** | [**[]PrivateSpaceVpnTunnelPostBody**](PrivateSpaceVpnTunnelPostBody.md) | IPsec tunnels composing the VPN member. | 

## Methods

### NewPrivateSpaceVpnPostBody

`func NewPrivateSpaceVpnPostBody(localAsn string, remoteIpAddress string, vpnTunnels []PrivateSpaceVpnTunnelPostBody, ) *PrivateSpaceVpnPostBody`

NewPrivateSpaceVpnPostBody instantiates a new PrivateSpaceVpnPostBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateSpaceVpnPostBodyWithDefaults

`func NewPrivateSpaceVpnPostBodyWithDefaults() *PrivateSpaceVpnPostBody`

NewPrivateSpaceVpnPostBodyWithDefaults instantiates a new PrivateSpaceVpnPostBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PrivateSpaceVpnPostBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrivateSpaceVpnPostBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrivateSpaceVpnPostBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PrivateSpaceVpnPostBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLocalAsn

`func (o *PrivateSpaceVpnPostBody) GetLocalAsn() string`

GetLocalAsn returns the LocalAsn field if non-nil, zero value otherwise.

### GetLocalAsnOk

`func (o *PrivateSpaceVpnPostBody) GetLocalAsnOk() (*string, bool)`

GetLocalAsnOk returns a tuple with the LocalAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAsn

`func (o *PrivateSpaceVpnPostBody) SetLocalAsn(v string)`

SetLocalAsn sets LocalAsn field to given value.


### GetRemoteAsn

`func (o *PrivateSpaceVpnPostBody) GetRemoteAsn() string`

GetRemoteAsn returns the RemoteAsn field if non-nil, zero value otherwise.

### GetRemoteAsnOk

`func (o *PrivateSpaceVpnPostBody) GetRemoteAsnOk() (*string, bool)`

GetRemoteAsnOk returns a tuple with the RemoteAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAsn

`func (o *PrivateSpaceVpnPostBody) SetRemoteAsn(v string)`

SetRemoteAsn sets RemoteAsn field to given value.

### HasRemoteAsn

`func (o *PrivateSpaceVpnPostBody) HasRemoteAsn() bool`

HasRemoteAsn returns a boolean if a field has been set.

### GetRemoteIpAddress

`func (o *PrivateSpaceVpnPostBody) GetRemoteIpAddress() string`

GetRemoteIpAddress returns the RemoteIpAddress field if non-nil, zero value otherwise.

### GetRemoteIpAddressOk

`func (o *PrivateSpaceVpnPostBody) GetRemoteIpAddressOk() (*string, bool)`

GetRemoteIpAddressOk returns a tuple with the RemoteIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIpAddress

`func (o *PrivateSpaceVpnPostBody) SetRemoteIpAddress(v string)`

SetRemoteIpAddress sets RemoteIpAddress field to given value.


### GetStaticRoutes

`func (o *PrivateSpaceVpnPostBody) GetStaticRoutes() []string`

GetStaticRoutes returns the StaticRoutes field if non-nil, zero value otherwise.

### GetStaticRoutesOk

`func (o *PrivateSpaceVpnPostBody) GetStaticRoutesOk() (*[]string, bool)`

GetStaticRoutesOk returns a tuple with the StaticRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticRoutes

`func (o *PrivateSpaceVpnPostBody) SetStaticRoutes(v []string)`

SetStaticRoutes sets StaticRoutes field to given value.

### HasStaticRoutes

`func (o *PrivateSpaceVpnPostBody) HasStaticRoutes() bool`

HasStaticRoutes returns a boolean if a field has been set.

### GetVpnTunnels

`func (o *PrivateSpaceVpnPostBody) GetVpnTunnels() []PrivateSpaceVpnTunnelPostBody`

GetVpnTunnels returns the VpnTunnels field if non-nil, zero value otherwise.

### GetVpnTunnelsOk

`func (o *PrivateSpaceVpnPostBody) GetVpnTunnelsOk() (*[]PrivateSpaceVpnTunnelPostBody, bool)`

GetVpnTunnelsOk returns a tuple with the VpnTunnels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnTunnels

`func (o *PrivateSpaceVpnPostBody) SetVpnTunnels(v []PrivateSpaceVpnTunnelPostBody)`

SetVpnTunnels sets VpnTunnels field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


