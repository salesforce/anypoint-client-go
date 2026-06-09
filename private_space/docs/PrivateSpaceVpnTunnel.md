# PrivateSpaceVpnTunnel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Psk** | Pointer to **string** | Pre-shared key used to authenticate the IPsec tunnel. | [optional] 
**PtpCidr** | Pointer to **string** | Point-to-point /30 CIDR used for the inside-tunnel addresses. | [optional] 
**StartupAction** | Pointer to **string** | Tunnel startup action. Allowed values are &#39;start&#39; or &#39;add&#39;. | [optional] 
**IsLogsEnabled** | Pointer to **bool** | True when tunnel logs are enabled. | [optional] 

## Methods

### NewPrivateSpaceVpnTunnel

`func NewPrivateSpaceVpnTunnel() *PrivateSpaceVpnTunnel`

NewPrivateSpaceVpnTunnel instantiates a new PrivateSpaceVpnTunnel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateSpaceVpnTunnelWithDefaults

`func NewPrivateSpaceVpnTunnelWithDefaults() *PrivateSpaceVpnTunnel`

NewPrivateSpaceVpnTunnelWithDefaults instantiates a new PrivateSpaceVpnTunnel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPsk

`func (o *PrivateSpaceVpnTunnel) GetPsk() string`

GetPsk returns the Psk field if non-nil, zero value otherwise.

### GetPskOk

`func (o *PrivateSpaceVpnTunnel) GetPskOk() (*string, bool)`

GetPskOk returns a tuple with the Psk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsk

`func (o *PrivateSpaceVpnTunnel) SetPsk(v string)`

SetPsk sets Psk field to given value.

### HasPsk

`func (o *PrivateSpaceVpnTunnel) HasPsk() bool`

HasPsk returns a boolean if a field has been set.

### GetPtpCidr

`func (o *PrivateSpaceVpnTunnel) GetPtpCidr() string`

GetPtpCidr returns the PtpCidr field if non-nil, zero value otherwise.

### GetPtpCidrOk

`func (o *PrivateSpaceVpnTunnel) GetPtpCidrOk() (*string, bool)`

GetPtpCidrOk returns a tuple with the PtpCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtpCidr

`func (o *PrivateSpaceVpnTunnel) SetPtpCidr(v string)`

SetPtpCidr sets PtpCidr field to given value.

### HasPtpCidr

`func (o *PrivateSpaceVpnTunnel) HasPtpCidr() bool`

HasPtpCidr returns a boolean if a field has been set.

### GetStartupAction

`func (o *PrivateSpaceVpnTunnel) GetStartupAction() string`

GetStartupAction returns the StartupAction field if non-nil, zero value otherwise.

### GetStartupActionOk

`func (o *PrivateSpaceVpnTunnel) GetStartupActionOk() (*string, bool)`

GetStartupActionOk returns a tuple with the StartupAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupAction

`func (o *PrivateSpaceVpnTunnel) SetStartupAction(v string)`

SetStartupAction sets StartupAction field to given value.

### HasStartupAction

`func (o *PrivateSpaceVpnTunnel) HasStartupAction() bool`

HasStartupAction returns a boolean if a field has been set.

### GetIsLogsEnabled

`func (o *PrivateSpaceVpnTunnel) GetIsLogsEnabled() bool`

GetIsLogsEnabled returns the IsLogsEnabled field if non-nil, zero value otherwise.

### GetIsLogsEnabledOk

`func (o *PrivateSpaceVpnTunnel) GetIsLogsEnabledOk() (*bool, bool)`

GetIsLogsEnabledOk returns a tuple with the IsLogsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLogsEnabled

`func (o *PrivateSpaceVpnTunnel) SetIsLogsEnabled(v bool)`

SetIsLogsEnabled sets IsLogsEnabled field to given value.

### HasIsLogsEnabled

`func (o *PrivateSpaceVpnTunnel) HasIsLogsEnabled() bool`

HasIsLogsEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


