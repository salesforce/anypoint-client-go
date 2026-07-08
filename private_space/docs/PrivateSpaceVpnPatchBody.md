# PrivateSpaceVpnPatchBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | New display name for the VPN member. | [optional] 
**StaticRoutes** | Pointer to **[]string** | Replacement set of static routes (CIDR strings). | [optional] 
**VpnTunnels** | Pointer to [**[]PrivateSpaceVpnTunnelPatchBody**](PrivateSpaceVpnTunnelPatchBody.md) | Replacement startupAction for each tunnel position. Order must match the existing tunnel order. | [optional] 

## Methods

### NewPrivateSpaceVpnPatchBody

`func NewPrivateSpaceVpnPatchBody() *PrivateSpaceVpnPatchBody`

NewPrivateSpaceVpnPatchBody instantiates a new PrivateSpaceVpnPatchBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateSpaceVpnPatchBodyWithDefaults

`func NewPrivateSpaceVpnPatchBodyWithDefaults() *PrivateSpaceVpnPatchBody`

NewPrivateSpaceVpnPatchBodyWithDefaults instantiates a new PrivateSpaceVpnPatchBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PrivateSpaceVpnPatchBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrivateSpaceVpnPatchBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrivateSpaceVpnPatchBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PrivateSpaceVpnPatchBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStaticRoutes

`func (o *PrivateSpaceVpnPatchBody) GetStaticRoutes() []string`

GetStaticRoutes returns the StaticRoutes field if non-nil, zero value otherwise.

### GetStaticRoutesOk

`func (o *PrivateSpaceVpnPatchBody) GetStaticRoutesOk() (*[]string, bool)`

GetStaticRoutesOk returns a tuple with the StaticRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticRoutes

`func (o *PrivateSpaceVpnPatchBody) SetStaticRoutes(v []string)`

SetStaticRoutes sets StaticRoutes field to given value.

### HasStaticRoutes

`func (o *PrivateSpaceVpnPatchBody) HasStaticRoutes() bool`

HasStaticRoutes returns a boolean if a field has been set.

### GetVpnTunnels

`func (o *PrivateSpaceVpnPatchBody) GetVpnTunnels() []PrivateSpaceVpnTunnelPatchBody`

GetVpnTunnels returns the VpnTunnels field if non-nil, zero value otherwise.

### GetVpnTunnelsOk

`func (o *PrivateSpaceVpnPatchBody) GetVpnTunnelsOk() (*[]PrivateSpaceVpnTunnelPatchBody, bool)`

GetVpnTunnelsOk returns a tuple with the VpnTunnels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnTunnels

`func (o *PrivateSpaceVpnPatchBody) SetVpnTunnels(v []PrivateSpaceVpnTunnelPatchBody)`

SetVpnTunnels sets VpnTunnels field to given value.

### HasVpnTunnels

`func (o *PrivateSpaceVpnPatchBody) HasVpnTunnels() bool`

HasVpnTunnels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


