# TunnelConfig1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Psk** | **string** | Pre-shared secret key, this is used for authentication the VPN-tunnel. | [default to ""]
**PtpCidr** | **string** | Point-to-point Classless Inter-Domain Routing (CIDR) | [default to ""]
**RekeyMarginInSeconds** | **int32** | The margin time in seconds before connection expiration or keying-channel expiration, during which the local side of the VPN connection attempts to negotiate a replacement. The exact time of the rekey is randomly selected based on the value of Rekey fuzz. Relevant only locally, the remote side does not need to agree on it. | [default to 0]
**RekeyFuzz** | **int32** | The maximum percentage by which marginbytes, marginpackets and margintime are randomly increased to randomize rekeying intervals (important for hosts with many connections). | [default to 0]

## Methods

### NewTunnelConfig1

`func NewTunnelConfig1(psk string, ptpCidr string, rekeyMarginInSeconds int32, rekeyFuzz int32, ) *TunnelConfig1`

NewTunnelConfig1 instantiates a new TunnelConfig1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTunnelConfig1WithDefaults

`func NewTunnelConfig1WithDefaults() *TunnelConfig1`

NewTunnelConfig1WithDefaults instantiates a new TunnelConfig1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPsk

`func (o *TunnelConfig1) GetPsk() string`

GetPsk returns the Psk field if non-nil, zero value otherwise.

### GetPskOk

`func (o *TunnelConfig1) GetPskOk() (*string, bool)`

GetPskOk returns a tuple with the Psk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsk

`func (o *TunnelConfig1) SetPsk(v string)`

SetPsk sets Psk field to given value.


### GetPtpCidr

`func (o *TunnelConfig1) GetPtpCidr() string`

GetPtpCidr returns the PtpCidr field if non-nil, zero value otherwise.

### GetPtpCidrOk

`func (o *TunnelConfig1) GetPtpCidrOk() (*string, bool)`

GetPtpCidrOk returns a tuple with the PtpCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtpCidr

`func (o *TunnelConfig1) SetPtpCidr(v string)`

SetPtpCidr sets PtpCidr field to given value.


### GetRekeyMarginInSeconds

`func (o *TunnelConfig1) GetRekeyMarginInSeconds() int32`

GetRekeyMarginInSeconds returns the RekeyMarginInSeconds field if non-nil, zero value otherwise.

### GetRekeyMarginInSecondsOk

`func (o *TunnelConfig1) GetRekeyMarginInSecondsOk() (*int32, bool)`

GetRekeyMarginInSecondsOk returns a tuple with the RekeyMarginInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRekeyMarginInSeconds

`func (o *TunnelConfig1) SetRekeyMarginInSeconds(v int32)`

SetRekeyMarginInSeconds sets RekeyMarginInSeconds field to given value.


### GetRekeyFuzz

`func (o *TunnelConfig1) GetRekeyFuzz() int32`

GetRekeyFuzz returns the RekeyFuzz field if non-nil, zero value otherwise.

### GetRekeyFuzzOk

`func (o *TunnelConfig1) GetRekeyFuzzOk() (*int32, bool)`

GetRekeyFuzzOk returns a tuple with the RekeyFuzz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRekeyFuzz

`func (o *TunnelConfig1) SetRekeyFuzz(v int32)`

SetRekeyFuzz sets RekeyFuzz field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


