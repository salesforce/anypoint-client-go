# VpnTunnel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedRouteCount** | **int32** | Amount of routes that have been accepted by the VPN. | [default to 0]
**LastStatusChange** | **string** | The last time the status of the VPN changed. | [default to ""]
**LocalExternalIpAddress** | **string** | The local external ip address. | [default to ""]
**LocalPtpIpAddress** | **string** | Point-to-point Ip address of the VPN that connects with the VPC. This address can be used by internal VM&#39;s/apps to connect with another server/app. | [default to ""]
**RemotePtpIpAddress** | **string** | The remote point-to-point ip address which a connecting tool like Cisco Anyconnect can connect with to use the VPC. | [default to ""]
**Psk** | **string** | Pre-shared secret key, this is used for authentication the VPN-tunnel. | [default to ""]
**Status** | **string** | The current status of the VPN tunnel | [default to ""]
**StatusMessage** | **string** | Message of the status | [default to ""]

## Methods

### NewVpnTunnel

`func NewVpnTunnel(acceptedRouteCount int32, lastStatusChange string, localExternalIpAddress string, localPtpIpAddress string, remotePtpIpAddress string, psk string, status string, statusMessage string, ) *VpnTunnel`

NewVpnTunnel instantiates a new VpnTunnel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnTunnelWithDefaults

`func NewVpnTunnelWithDefaults() *VpnTunnel`

NewVpnTunnelWithDefaults instantiates a new VpnTunnel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptedRouteCount

`func (o *VpnTunnel) GetAcceptedRouteCount() int32`

GetAcceptedRouteCount returns the AcceptedRouteCount field if non-nil, zero value otherwise.

### GetAcceptedRouteCountOk

`func (o *VpnTunnel) GetAcceptedRouteCountOk() (*int32, bool)`

GetAcceptedRouteCountOk returns a tuple with the AcceptedRouteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedRouteCount

`func (o *VpnTunnel) SetAcceptedRouteCount(v int32)`

SetAcceptedRouteCount sets AcceptedRouteCount field to given value.


### GetLastStatusChange

`func (o *VpnTunnel) GetLastStatusChange() string`

GetLastStatusChange returns the LastStatusChange field if non-nil, zero value otherwise.

### GetLastStatusChangeOk

`func (o *VpnTunnel) GetLastStatusChangeOk() (*string, bool)`

GetLastStatusChangeOk returns a tuple with the LastStatusChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatusChange

`func (o *VpnTunnel) SetLastStatusChange(v string)`

SetLastStatusChange sets LastStatusChange field to given value.


### GetLocalExternalIpAddress

`func (o *VpnTunnel) GetLocalExternalIpAddress() string`

GetLocalExternalIpAddress returns the LocalExternalIpAddress field if non-nil, zero value otherwise.

### GetLocalExternalIpAddressOk

`func (o *VpnTunnel) GetLocalExternalIpAddressOk() (*string, bool)`

GetLocalExternalIpAddressOk returns a tuple with the LocalExternalIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalExternalIpAddress

`func (o *VpnTunnel) SetLocalExternalIpAddress(v string)`

SetLocalExternalIpAddress sets LocalExternalIpAddress field to given value.


### GetLocalPtpIpAddress

`func (o *VpnTunnel) GetLocalPtpIpAddress() string`

GetLocalPtpIpAddress returns the LocalPtpIpAddress field if non-nil, zero value otherwise.

### GetLocalPtpIpAddressOk

`func (o *VpnTunnel) GetLocalPtpIpAddressOk() (*string, bool)`

GetLocalPtpIpAddressOk returns a tuple with the LocalPtpIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPtpIpAddress

`func (o *VpnTunnel) SetLocalPtpIpAddress(v string)`

SetLocalPtpIpAddress sets LocalPtpIpAddress field to given value.


### GetRemotePtpIpAddress

`func (o *VpnTunnel) GetRemotePtpIpAddress() string`

GetRemotePtpIpAddress returns the RemotePtpIpAddress field if non-nil, zero value otherwise.

### GetRemotePtpIpAddressOk

`func (o *VpnTunnel) GetRemotePtpIpAddressOk() (*string, bool)`

GetRemotePtpIpAddressOk returns a tuple with the RemotePtpIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePtpIpAddress

`func (o *VpnTunnel) SetRemotePtpIpAddress(v string)`

SetRemotePtpIpAddress sets RemotePtpIpAddress field to given value.


### GetPsk

`func (o *VpnTunnel) GetPsk() string`

GetPsk returns the Psk field if non-nil, zero value otherwise.

### GetPskOk

`func (o *VpnTunnel) GetPskOk() (*string, bool)`

GetPskOk returns a tuple with the Psk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsk

`func (o *VpnTunnel) SetPsk(v string)`

SetPsk sets Psk field to given value.


### GetStatus

`func (o *VpnTunnel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VpnTunnel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VpnTunnel) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusMessage

`func (o *VpnTunnel) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *VpnTunnel) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *VpnTunnel) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


