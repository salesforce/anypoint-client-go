# DlbCore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** | dlb state | [optional] [default to "STOPPED"]
**IpWhitelist** | Pointer to **[]string** |  | [optional] 
**IpAllowlist** | Pointer to **[]string** |  | [optional] 
**HttpMode** | Pointer to **string** |  | [optional] [default to "redirect"]
**DefaultSslEndpoint** | Pointer to **int32** |  | [optional] [default to 0]
**Tlsv1** | Pointer to **bool** |  | [optional] 
**SslEndpoints** | Pointer to [**[]DlbCoreSslEndpointsInner**](DlbCoreSslEndpointsInner.md) |  | [optional] 

## Methods

### NewDlbCore

`func NewDlbCore() *DlbCore`

NewDlbCore instantiates a new DlbCore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDlbCoreWithDefaults

`func NewDlbCoreWithDefaults() *DlbCore`

NewDlbCoreWithDefaults instantiates a new DlbCore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DlbCore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DlbCore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DlbCore) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DlbCore) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *DlbCore) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DlbCore) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DlbCore) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DlbCore) HasState() bool`

HasState returns a boolean if a field has been set.

### GetIpWhitelist

`func (o *DlbCore) GetIpWhitelist() []string`

GetIpWhitelist returns the IpWhitelist field if non-nil, zero value otherwise.

### GetIpWhitelistOk

`func (o *DlbCore) GetIpWhitelistOk() (*[]string, bool)`

GetIpWhitelistOk returns a tuple with the IpWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpWhitelist

`func (o *DlbCore) SetIpWhitelist(v []string)`

SetIpWhitelist sets IpWhitelist field to given value.

### HasIpWhitelist

`func (o *DlbCore) HasIpWhitelist() bool`

HasIpWhitelist returns a boolean if a field has been set.

### GetIpAllowlist

`func (o *DlbCore) GetIpAllowlist() []string`

GetIpAllowlist returns the IpAllowlist field if non-nil, zero value otherwise.

### GetIpAllowlistOk

`func (o *DlbCore) GetIpAllowlistOk() (*[]string, bool)`

GetIpAllowlistOk returns a tuple with the IpAllowlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAllowlist

`func (o *DlbCore) SetIpAllowlist(v []string)`

SetIpAllowlist sets IpAllowlist field to given value.

### HasIpAllowlist

`func (o *DlbCore) HasIpAllowlist() bool`

HasIpAllowlist returns a boolean if a field has been set.

### GetHttpMode

`func (o *DlbCore) GetHttpMode() string`

GetHttpMode returns the HttpMode field if non-nil, zero value otherwise.

### GetHttpModeOk

`func (o *DlbCore) GetHttpModeOk() (*string, bool)`

GetHttpModeOk returns a tuple with the HttpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMode

`func (o *DlbCore) SetHttpMode(v string)`

SetHttpMode sets HttpMode field to given value.

### HasHttpMode

`func (o *DlbCore) HasHttpMode() bool`

HasHttpMode returns a boolean if a field has been set.

### GetDefaultSslEndpoint

`func (o *DlbCore) GetDefaultSslEndpoint() int32`

GetDefaultSslEndpoint returns the DefaultSslEndpoint field if non-nil, zero value otherwise.

### GetDefaultSslEndpointOk

`func (o *DlbCore) GetDefaultSslEndpointOk() (*int32, bool)`

GetDefaultSslEndpointOk returns a tuple with the DefaultSslEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSslEndpoint

`func (o *DlbCore) SetDefaultSslEndpoint(v int32)`

SetDefaultSslEndpoint sets DefaultSslEndpoint field to given value.

### HasDefaultSslEndpoint

`func (o *DlbCore) HasDefaultSslEndpoint() bool`

HasDefaultSslEndpoint returns a boolean if a field has been set.

### GetTlsv1

`func (o *DlbCore) GetTlsv1() bool`

GetTlsv1 returns the Tlsv1 field if non-nil, zero value otherwise.

### GetTlsv1Ok

`func (o *DlbCore) GetTlsv1Ok() (*bool, bool)`

GetTlsv1Ok returns a tuple with the Tlsv1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsv1

`func (o *DlbCore) SetTlsv1(v bool)`

SetTlsv1 sets Tlsv1 field to given value.

### HasTlsv1

`func (o *DlbCore) HasTlsv1() bool`

HasTlsv1 returns a boolean if a field has been set.

### GetSslEndpoints

`func (o *DlbCore) GetSslEndpoints() []DlbCoreSslEndpointsInner`

GetSslEndpoints returns the SslEndpoints field if non-nil, zero value otherwise.

### GetSslEndpointsOk

`func (o *DlbCore) GetSslEndpointsOk() (*[]DlbCoreSslEndpointsInner, bool)`

GetSslEndpointsOk returns a tuple with the SslEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslEndpoints

`func (o *DlbCore) SetSslEndpoints(v []DlbCoreSslEndpointsInner)`

SetSslEndpoints sets SslEndpoints field to given value.

### HasSslEndpoints

`func (o *DlbCore) HasSslEndpoints() bool`

HasSslEndpoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


