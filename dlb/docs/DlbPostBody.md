# DlbPostBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** | dlb state | [optional] [default to "STOPPED"]
**Domain** | Pointer to **string** | The DNS domain for the Load Balancer | [optional] 
**IpWhitelist** | Pointer to **[]string** |  | [optional] 
**IpAllowlist** | Pointer to **[]string** |  | [optional] 
**HttpMode** | Pointer to **string** |  | [optional] [default to "redirect"]
**DefaultSslEndpoint** | Pointer to **int32** |  | [optional] [default to 0]
**Tlsv1** | Pointer to **bool** |  | [optional] 
**KeepUrlEncoding** | Pointer to **bool** |  | [optional] 
**UpstreamTlsv12** | Pointer to **bool** |  | [optional] 
**StaticIPsDisabled** | Pointer to **bool** |  | [optional] 
**DoubleStaticIps** | Pointer to **bool** |  | [optional] 
**EnableStreaming** | Pointer to **bool** |  | [optional] 
**ForwardClientCertificate** | Pointer to **bool** |  | [optional] 
**Workers** | Pointer to **int32** | the dedicated load balancer&#39;s number of workers. | [optional] [default to 2]
**ProxyReadTimeout** | Pointer to **string** | Sets the Mule application response timeout value. | [optional] [default to "300"]
**SslEndpoints** | Pointer to [**[]DlbPostBodySslEndpoints**](DlbPostBodySslEndpoints.md) |  | [optional] 

## Methods

### NewDlbPostBody

`func NewDlbPostBody() *DlbPostBody`

NewDlbPostBody instantiates a new DlbPostBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDlbPostBodyWithDefaults

`func NewDlbPostBodyWithDefaults() *DlbPostBody`

NewDlbPostBodyWithDefaults instantiates a new DlbPostBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DlbPostBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DlbPostBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DlbPostBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DlbPostBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *DlbPostBody) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DlbPostBody) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DlbPostBody) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DlbPostBody) HasState() bool`

HasState returns a boolean if a field has been set.

### GetDomain

`func (o *DlbPostBody) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DlbPostBody) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DlbPostBody) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *DlbPostBody) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetIpWhitelist

`func (o *DlbPostBody) GetIpWhitelist() []string`

GetIpWhitelist returns the IpWhitelist field if non-nil, zero value otherwise.

### GetIpWhitelistOk

`func (o *DlbPostBody) GetIpWhitelistOk() (*[]string, bool)`

GetIpWhitelistOk returns a tuple with the IpWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpWhitelist

`func (o *DlbPostBody) SetIpWhitelist(v []string)`

SetIpWhitelist sets IpWhitelist field to given value.

### HasIpWhitelist

`func (o *DlbPostBody) HasIpWhitelist() bool`

HasIpWhitelist returns a boolean if a field has been set.

### GetIpAllowlist

`func (o *DlbPostBody) GetIpAllowlist() []string`

GetIpAllowlist returns the IpAllowlist field if non-nil, zero value otherwise.

### GetIpAllowlistOk

`func (o *DlbPostBody) GetIpAllowlistOk() (*[]string, bool)`

GetIpAllowlistOk returns a tuple with the IpAllowlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAllowlist

`func (o *DlbPostBody) SetIpAllowlist(v []string)`

SetIpAllowlist sets IpAllowlist field to given value.

### HasIpAllowlist

`func (o *DlbPostBody) HasIpAllowlist() bool`

HasIpAllowlist returns a boolean if a field has been set.

### GetHttpMode

`func (o *DlbPostBody) GetHttpMode() string`

GetHttpMode returns the HttpMode field if non-nil, zero value otherwise.

### GetHttpModeOk

`func (o *DlbPostBody) GetHttpModeOk() (*string, bool)`

GetHttpModeOk returns a tuple with the HttpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMode

`func (o *DlbPostBody) SetHttpMode(v string)`

SetHttpMode sets HttpMode field to given value.

### HasHttpMode

`func (o *DlbPostBody) HasHttpMode() bool`

HasHttpMode returns a boolean if a field has been set.

### GetDefaultSslEndpoint

`func (o *DlbPostBody) GetDefaultSslEndpoint() int32`

GetDefaultSslEndpoint returns the DefaultSslEndpoint field if non-nil, zero value otherwise.

### GetDefaultSslEndpointOk

`func (o *DlbPostBody) GetDefaultSslEndpointOk() (*int32, bool)`

GetDefaultSslEndpointOk returns a tuple with the DefaultSslEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSslEndpoint

`func (o *DlbPostBody) SetDefaultSslEndpoint(v int32)`

SetDefaultSslEndpoint sets DefaultSslEndpoint field to given value.

### HasDefaultSslEndpoint

`func (o *DlbPostBody) HasDefaultSslEndpoint() bool`

HasDefaultSslEndpoint returns a boolean if a field has been set.

### GetTlsv1

`func (o *DlbPostBody) GetTlsv1() bool`

GetTlsv1 returns the Tlsv1 field if non-nil, zero value otherwise.

### GetTlsv1Ok

`func (o *DlbPostBody) GetTlsv1Ok() (*bool, bool)`

GetTlsv1Ok returns a tuple with the Tlsv1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsv1

`func (o *DlbPostBody) SetTlsv1(v bool)`

SetTlsv1 sets Tlsv1 field to given value.

### HasTlsv1

`func (o *DlbPostBody) HasTlsv1() bool`

HasTlsv1 returns a boolean if a field has been set.

### GetKeepUrlEncoding

`func (o *DlbPostBody) GetKeepUrlEncoding() bool`

GetKeepUrlEncoding returns the KeepUrlEncoding field if non-nil, zero value otherwise.

### GetKeepUrlEncodingOk

`func (o *DlbPostBody) GetKeepUrlEncodingOk() (*bool, bool)`

GetKeepUrlEncodingOk returns a tuple with the KeepUrlEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepUrlEncoding

`func (o *DlbPostBody) SetKeepUrlEncoding(v bool)`

SetKeepUrlEncoding sets KeepUrlEncoding field to given value.

### HasKeepUrlEncoding

`func (o *DlbPostBody) HasKeepUrlEncoding() bool`

HasKeepUrlEncoding returns a boolean if a field has been set.

### GetUpstreamTlsv12

`func (o *DlbPostBody) GetUpstreamTlsv12() bool`

GetUpstreamTlsv12 returns the UpstreamTlsv12 field if non-nil, zero value otherwise.

### GetUpstreamTlsv12Ok

`func (o *DlbPostBody) GetUpstreamTlsv12Ok() (*bool, bool)`

GetUpstreamTlsv12Ok returns a tuple with the UpstreamTlsv12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamTlsv12

`func (o *DlbPostBody) SetUpstreamTlsv12(v bool)`

SetUpstreamTlsv12 sets UpstreamTlsv12 field to given value.

### HasUpstreamTlsv12

`func (o *DlbPostBody) HasUpstreamTlsv12() bool`

HasUpstreamTlsv12 returns a boolean if a field has been set.

### GetStaticIPsDisabled

`func (o *DlbPostBody) GetStaticIPsDisabled() bool`

GetStaticIPsDisabled returns the StaticIPsDisabled field if non-nil, zero value otherwise.

### GetStaticIPsDisabledOk

`func (o *DlbPostBody) GetStaticIPsDisabledOk() (*bool, bool)`

GetStaticIPsDisabledOk returns a tuple with the StaticIPsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticIPsDisabled

`func (o *DlbPostBody) SetStaticIPsDisabled(v bool)`

SetStaticIPsDisabled sets StaticIPsDisabled field to given value.

### HasStaticIPsDisabled

`func (o *DlbPostBody) HasStaticIPsDisabled() bool`

HasStaticIPsDisabled returns a boolean if a field has been set.

### GetDoubleStaticIps

`func (o *DlbPostBody) GetDoubleStaticIps() bool`

GetDoubleStaticIps returns the DoubleStaticIps field if non-nil, zero value otherwise.

### GetDoubleStaticIpsOk

`func (o *DlbPostBody) GetDoubleStaticIpsOk() (*bool, bool)`

GetDoubleStaticIpsOk returns a tuple with the DoubleStaticIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoubleStaticIps

`func (o *DlbPostBody) SetDoubleStaticIps(v bool)`

SetDoubleStaticIps sets DoubleStaticIps field to given value.

### HasDoubleStaticIps

`func (o *DlbPostBody) HasDoubleStaticIps() bool`

HasDoubleStaticIps returns a boolean if a field has been set.

### GetEnableStreaming

`func (o *DlbPostBody) GetEnableStreaming() bool`

GetEnableStreaming returns the EnableStreaming field if non-nil, zero value otherwise.

### GetEnableStreamingOk

`func (o *DlbPostBody) GetEnableStreamingOk() (*bool, bool)`

GetEnableStreamingOk returns a tuple with the EnableStreaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableStreaming

`func (o *DlbPostBody) SetEnableStreaming(v bool)`

SetEnableStreaming sets EnableStreaming field to given value.

### HasEnableStreaming

`func (o *DlbPostBody) HasEnableStreaming() bool`

HasEnableStreaming returns a boolean if a field has been set.

### GetForwardClientCertificate

`func (o *DlbPostBody) GetForwardClientCertificate() bool`

GetForwardClientCertificate returns the ForwardClientCertificate field if non-nil, zero value otherwise.

### GetForwardClientCertificateOk

`func (o *DlbPostBody) GetForwardClientCertificateOk() (*bool, bool)`

GetForwardClientCertificateOk returns a tuple with the ForwardClientCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardClientCertificate

`func (o *DlbPostBody) SetForwardClientCertificate(v bool)`

SetForwardClientCertificate sets ForwardClientCertificate field to given value.

### HasForwardClientCertificate

`func (o *DlbPostBody) HasForwardClientCertificate() bool`

HasForwardClientCertificate returns a boolean if a field has been set.

### GetWorkers

`func (o *DlbPostBody) GetWorkers() int32`

GetWorkers returns the Workers field if non-nil, zero value otherwise.

### GetWorkersOk

`func (o *DlbPostBody) GetWorkersOk() (*int32, bool)`

GetWorkersOk returns a tuple with the Workers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkers

`func (o *DlbPostBody) SetWorkers(v int32)`

SetWorkers sets Workers field to given value.

### HasWorkers

`func (o *DlbPostBody) HasWorkers() bool`

HasWorkers returns a boolean if a field has been set.

### GetProxyReadTimeout

`func (o *DlbPostBody) GetProxyReadTimeout() string`

GetProxyReadTimeout returns the ProxyReadTimeout field if non-nil, zero value otherwise.

### GetProxyReadTimeoutOk

`func (o *DlbPostBody) GetProxyReadTimeoutOk() (*string, bool)`

GetProxyReadTimeoutOk returns a tuple with the ProxyReadTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyReadTimeout

`func (o *DlbPostBody) SetProxyReadTimeout(v string)`

SetProxyReadTimeout sets ProxyReadTimeout field to given value.

### HasProxyReadTimeout

`func (o *DlbPostBody) HasProxyReadTimeout() bool`

HasProxyReadTimeout returns a boolean if a field has been set.

### GetSslEndpoints

`func (o *DlbPostBody) GetSslEndpoints() []DlbPostBodySslEndpoints`

GetSslEndpoints returns the SslEndpoints field if non-nil, zero value otherwise.

### GetSslEndpointsOk

`func (o *DlbPostBody) GetSslEndpointsOk() (*[]DlbPostBodySslEndpoints, bool)`

GetSslEndpointsOk returns a tuple with the SslEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslEndpoints

`func (o *DlbPostBody) SetSslEndpoints(v []DlbPostBodySslEndpoints)`

SetSslEndpoints sets SslEndpoints field to given value.

### HasSslEndpoints

`func (o *DlbPostBody) HasSslEndpoints() bool`

HasSslEndpoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


