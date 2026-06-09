# PrivateSpaceVpnTunnelPostBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Psk** | **string** | Pre-shared key used to authenticate the IPsec tunnel. | 
**PtpCidr** | Pointer to **string** | Point-to-point /30 CIDR used for the inside-tunnel addresses. | [optional] 
**StartupAction** | **string** | Tunnel startup action. Allowed values are &#39;start&#39; or &#39;add&#39;. | 

## Methods

### NewPrivateSpaceVpnTunnelPostBody

`func NewPrivateSpaceVpnTunnelPostBody(psk string, startupAction string, ) *PrivateSpaceVpnTunnelPostBody`

NewPrivateSpaceVpnTunnelPostBody instantiates a new PrivateSpaceVpnTunnelPostBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateSpaceVpnTunnelPostBodyWithDefaults

`func NewPrivateSpaceVpnTunnelPostBodyWithDefaults() *PrivateSpaceVpnTunnelPostBody`

NewPrivateSpaceVpnTunnelPostBodyWithDefaults instantiates a new PrivateSpaceVpnTunnelPostBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPsk

`func (o *PrivateSpaceVpnTunnelPostBody) GetPsk() string`

GetPsk returns the Psk field if non-nil, zero value otherwise.

### GetPskOk

`func (o *PrivateSpaceVpnTunnelPostBody) GetPskOk() (*string, bool)`

GetPskOk returns a tuple with the Psk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsk

`func (o *PrivateSpaceVpnTunnelPostBody) SetPsk(v string)`

SetPsk sets Psk field to given value.


### GetPtpCidr

`func (o *PrivateSpaceVpnTunnelPostBody) GetPtpCidr() string`

GetPtpCidr returns the PtpCidr field if non-nil, zero value otherwise.

### GetPtpCidrOk

`func (o *PrivateSpaceVpnTunnelPostBody) GetPtpCidrOk() (*string, bool)`

GetPtpCidrOk returns a tuple with the PtpCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtpCidr

`func (o *PrivateSpaceVpnTunnelPostBody) SetPtpCidr(v string)`

SetPtpCidr sets PtpCidr field to given value.

### HasPtpCidr

`func (o *PrivateSpaceVpnTunnelPostBody) HasPtpCidr() bool`

HasPtpCidr returns a boolean if a field has been set.

### GetStartupAction

`func (o *PrivateSpaceVpnTunnelPostBody) GetStartupAction() string`

GetStartupAction returns the StartupAction field if non-nil, zero value otherwise.

### GetStartupActionOk

`func (o *PrivateSpaceVpnTunnelPostBody) GetStartupActionOk() (*string, bool)`

GetStartupActionOk returns a tuple with the StartupAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupAction

`func (o *PrivateSpaceVpnTunnelPostBody) SetStartupAction(v string)`

SetStartupAction sets StartupAction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


