# TunnelConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Psk** | **string** | Pre-shared secret key, this is used for authentication the VPN-tunnel. | [default to ""]
**PtpCidr** | **string** | Point-to-point Classless Inter-Domain Routing (CIDR) | [default to ""]

## Methods

### NewTunnelConfig

`func NewTunnelConfig(psk string, ptpCidr string, ) *TunnelConfig`

NewTunnelConfig instantiates a new TunnelConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTunnelConfigWithDefaults

`func NewTunnelConfigWithDefaults() *TunnelConfig`

NewTunnelConfigWithDefaults instantiates a new TunnelConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPsk

`func (o *TunnelConfig) GetPsk() string`

GetPsk returns the Psk field if non-nil, zero value otherwise.

### GetPskOk

`func (o *TunnelConfig) GetPskOk() (*string, bool)`

GetPskOk returns a tuple with the Psk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsk

`func (o *TunnelConfig) SetPsk(v string)`

SetPsk sets Psk field to given value.


### GetPtpCidr

`func (o *TunnelConfig) GetPtpCidr() string`

GetPtpCidr returns the PtpCidr field if non-nil, zero value otherwise.

### GetPtpCidrOk

`func (o *TunnelConfig) GetPtpCidrOk() (*string, bool)`

GetPtpCidrOk returns a tuple with the PtpCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtpCidr

`func (o *TunnelConfig) SetPtpCidr(v string)`

SetPtpCidr sets PtpCidr field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


