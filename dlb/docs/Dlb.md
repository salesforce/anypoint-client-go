# Dlb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The dlb Id | [optional] 
**Domain** | Pointer to **string** | dlb domain | [optional] 
**DeploymentId** | Pointer to **string** | the dlb deployment id | [optional] 
**InstanceConfig** | Pointer to [**InstanceConfig**](InstanceConfig.md) |  | [optional] 
**IpAddresses** | Pointer to **[]string** |  | [optional] 
**StaticIPsDisabled** | Pointer to **bool** |  | [optional] 
**VpcId** | Pointer to **string** |  | [optional] 
**Workers** | Pointer to **int32** |  | [optional] 
**DefaultCipherSuite** | Pointer to **string** |  | [optional] 
**KeepUrlEncoding** | Pointer to **bool** |  | [optional] 
**UpstreamTlsv12** | Pointer to **bool** |  | [optional] 
**ProxyReadTimeout** | Pointer to **int32** |  | [optional] 
**IpAddressesInfo** | Pointer to [**[]DlbExtrasIpAddressesInfo**](DlbExtrasIpAddressesInfo.md) |  | [optional] 
**DoubleStaticIps** | Pointer to **bool** |  | [optional] 
**EnableStreaming** | Pointer to **bool** |  | [optional] 
**ForwardClientCertificate** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** | dlb state | [optional] [default to "STOPPED"]
**IpWhitelist** | Pointer to **[]string** |  | [optional] 
**IpAllowlist** | Pointer to **[]string** |  | [optional] 
**HttpMode** | Pointer to **string** |  | [optional] [default to "redirect"]
**DefaultSslEndpoint** | Pointer to **int32** |  | [optional] [default to 0]
**Tlsv1** | Pointer to **bool** |  | [optional] 
**SslEndpoints** | Pointer to [**[]DlbCoreSslEndpoints**](DlbCoreSslEndpoints.md) |  | [optional] 

## Methods

### NewDlb

`func NewDlb() *Dlb`

NewDlb instantiates a new Dlb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDlbWithDefaults

`func NewDlbWithDefaults() *Dlb`

NewDlbWithDefaults instantiates a new Dlb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Dlb) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Dlb) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Dlb) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Dlb) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDomain

`func (o *Dlb) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Dlb) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Dlb) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *Dlb) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetDeploymentId

`func (o *Dlb) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *Dlb) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *Dlb) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *Dlb) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### GetInstanceConfig

`func (o *Dlb) GetInstanceConfig() InstanceConfig`

GetInstanceConfig returns the InstanceConfig field if non-nil, zero value otherwise.

### GetInstanceConfigOk

`func (o *Dlb) GetInstanceConfigOk() (*InstanceConfig, bool)`

GetInstanceConfigOk returns a tuple with the InstanceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceConfig

`func (o *Dlb) SetInstanceConfig(v InstanceConfig)`

SetInstanceConfig sets InstanceConfig field to given value.

### HasInstanceConfig

`func (o *Dlb) HasInstanceConfig() bool`

HasInstanceConfig returns a boolean if a field has been set.

### GetIpAddresses

`func (o *Dlb) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *Dlb) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *Dlb) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *Dlb) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### GetStaticIPsDisabled

`func (o *Dlb) GetStaticIPsDisabled() bool`

GetStaticIPsDisabled returns the StaticIPsDisabled field if non-nil, zero value otherwise.

### GetStaticIPsDisabledOk

`func (o *Dlb) GetStaticIPsDisabledOk() (*bool, bool)`

GetStaticIPsDisabledOk returns a tuple with the StaticIPsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticIPsDisabled

`func (o *Dlb) SetStaticIPsDisabled(v bool)`

SetStaticIPsDisabled sets StaticIPsDisabled field to given value.

### HasStaticIPsDisabled

`func (o *Dlb) HasStaticIPsDisabled() bool`

HasStaticIPsDisabled returns a boolean if a field has been set.

### GetVpcId

`func (o *Dlb) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *Dlb) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *Dlb) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *Dlb) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### GetWorkers

`func (o *Dlb) GetWorkers() int32`

GetWorkers returns the Workers field if non-nil, zero value otherwise.

### GetWorkersOk

`func (o *Dlb) GetWorkersOk() (*int32, bool)`

GetWorkersOk returns a tuple with the Workers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkers

`func (o *Dlb) SetWorkers(v int32)`

SetWorkers sets Workers field to given value.

### HasWorkers

`func (o *Dlb) HasWorkers() bool`

HasWorkers returns a boolean if a field has been set.

### GetDefaultCipherSuite

`func (o *Dlb) GetDefaultCipherSuite() string`

GetDefaultCipherSuite returns the DefaultCipherSuite field if non-nil, zero value otherwise.

### GetDefaultCipherSuiteOk

`func (o *Dlb) GetDefaultCipherSuiteOk() (*string, bool)`

GetDefaultCipherSuiteOk returns a tuple with the DefaultCipherSuite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCipherSuite

`func (o *Dlb) SetDefaultCipherSuite(v string)`

SetDefaultCipherSuite sets DefaultCipherSuite field to given value.

### HasDefaultCipherSuite

`func (o *Dlb) HasDefaultCipherSuite() bool`

HasDefaultCipherSuite returns a boolean if a field has been set.

### GetKeepUrlEncoding

`func (o *Dlb) GetKeepUrlEncoding() bool`

GetKeepUrlEncoding returns the KeepUrlEncoding field if non-nil, zero value otherwise.

### GetKeepUrlEncodingOk

`func (o *Dlb) GetKeepUrlEncodingOk() (*bool, bool)`

GetKeepUrlEncodingOk returns a tuple with the KeepUrlEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepUrlEncoding

`func (o *Dlb) SetKeepUrlEncoding(v bool)`

SetKeepUrlEncoding sets KeepUrlEncoding field to given value.

### HasKeepUrlEncoding

`func (o *Dlb) HasKeepUrlEncoding() bool`

HasKeepUrlEncoding returns a boolean if a field has been set.

### GetUpstreamTlsv12

`func (o *Dlb) GetUpstreamTlsv12() bool`

GetUpstreamTlsv12 returns the UpstreamTlsv12 field if non-nil, zero value otherwise.

### GetUpstreamTlsv12Ok

`func (o *Dlb) GetUpstreamTlsv12Ok() (*bool, bool)`

GetUpstreamTlsv12Ok returns a tuple with the UpstreamTlsv12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamTlsv12

`func (o *Dlb) SetUpstreamTlsv12(v bool)`

SetUpstreamTlsv12 sets UpstreamTlsv12 field to given value.

### HasUpstreamTlsv12

`func (o *Dlb) HasUpstreamTlsv12() bool`

HasUpstreamTlsv12 returns a boolean if a field has been set.

### GetProxyReadTimeout

`func (o *Dlb) GetProxyReadTimeout() int32`

GetProxyReadTimeout returns the ProxyReadTimeout field if non-nil, zero value otherwise.

### GetProxyReadTimeoutOk

`func (o *Dlb) GetProxyReadTimeoutOk() (*int32, bool)`

GetProxyReadTimeoutOk returns a tuple with the ProxyReadTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyReadTimeout

`func (o *Dlb) SetProxyReadTimeout(v int32)`

SetProxyReadTimeout sets ProxyReadTimeout field to given value.

### HasProxyReadTimeout

`func (o *Dlb) HasProxyReadTimeout() bool`

HasProxyReadTimeout returns a boolean if a field has been set.

### GetIpAddressesInfo

`func (o *Dlb) GetIpAddressesInfo() []DlbExtrasIpAddressesInfo`

GetIpAddressesInfo returns the IpAddressesInfo field if non-nil, zero value otherwise.

### GetIpAddressesInfoOk

`func (o *Dlb) GetIpAddressesInfoOk() (*[]DlbExtrasIpAddressesInfo, bool)`

GetIpAddressesInfoOk returns a tuple with the IpAddressesInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddressesInfo

`func (o *Dlb) SetIpAddressesInfo(v []DlbExtrasIpAddressesInfo)`

SetIpAddressesInfo sets IpAddressesInfo field to given value.

### HasIpAddressesInfo

`func (o *Dlb) HasIpAddressesInfo() bool`

HasIpAddressesInfo returns a boolean if a field has been set.

### GetDoubleStaticIps

`func (o *Dlb) GetDoubleStaticIps() bool`

GetDoubleStaticIps returns the DoubleStaticIps field if non-nil, zero value otherwise.

### GetDoubleStaticIpsOk

`func (o *Dlb) GetDoubleStaticIpsOk() (*bool, bool)`

GetDoubleStaticIpsOk returns a tuple with the DoubleStaticIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoubleStaticIps

`func (o *Dlb) SetDoubleStaticIps(v bool)`

SetDoubleStaticIps sets DoubleStaticIps field to given value.

### HasDoubleStaticIps

`func (o *Dlb) HasDoubleStaticIps() bool`

HasDoubleStaticIps returns a boolean if a field has been set.

### GetEnableStreaming

`func (o *Dlb) GetEnableStreaming() bool`

GetEnableStreaming returns the EnableStreaming field if non-nil, zero value otherwise.

### GetEnableStreamingOk

`func (o *Dlb) GetEnableStreamingOk() (*bool, bool)`

GetEnableStreamingOk returns a tuple with the EnableStreaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableStreaming

`func (o *Dlb) SetEnableStreaming(v bool)`

SetEnableStreaming sets EnableStreaming field to given value.

### HasEnableStreaming

`func (o *Dlb) HasEnableStreaming() bool`

HasEnableStreaming returns a boolean if a field has been set.

### GetForwardClientCertificate

`func (o *Dlb) GetForwardClientCertificate() bool`

GetForwardClientCertificate returns the ForwardClientCertificate field if non-nil, zero value otherwise.

### GetForwardClientCertificateOk

`func (o *Dlb) GetForwardClientCertificateOk() (*bool, bool)`

GetForwardClientCertificateOk returns a tuple with the ForwardClientCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardClientCertificate

`func (o *Dlb) SetForwardClientCertificate(v bool)`

SetForwardClientCertificate sets ForwardClientCertificate field to given value.

### HasForwardClientCertificate

`func (o *Dlb) HasForwardClientCertificate() bool`

HasForwardClientCertificate returns a boolean if a field has been set.

### GetName

`func (o *Dlb) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Dlb) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Dlb) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Dlb) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *Dlb) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Dlb) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Dlb) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Dlb) HasState() bool`

HasState returns a boolean if a field has been set.

### GetIpWhitelist

`func (o *Dlb) GetIpWhitelist() []string`

GetIpWhitelist returns the IpWhitelist field if non-nil, zero value otherwise.

### GetIpWhitelistOk

`func (o *Dlb) GetIpWhitelistOk() (*[]string, bool)`

GetIpWhitelistOk returns a tuple with the IpWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpWhitelist

`func (o *Dlb) SetIpWhitelist(v []string)`

SetIpWhitelist sets IpWhitelist field to given value.

### HasIpWhitelist

`func (o *Dlb) HasIpWhitelist() bool`

HasIpWhitelist returns a boolean if a field has been set.

### GetIpAllowlist

`func (o *Dlb) GetIpAllowlist() []string`

GetIpAllowlist returns the IpAllowlist field if non-nil, zero value otherwise.

### GetIpAllowlistOk

`func (o *Dlb) GetIpAllowlistOk() (*[]string, bool)`

GetIpAllowlistOk returns a tuple with the IpAllowlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAllowlist

`func (o *Dlb) SetIpAllowlist(v []string)`

SetIpAllowlist sets IpAllowlist field to given value.

### HasIpAllowlist

`func (o *Dlb) HasIpAllowlist() bool`

HasIpAllowlist returns a boolean if a field has been set.

### GetHttpMode

`func (o *Dlb) GetHttpMode() string`

GetHttpMode returns the HttpMode field if non-nil, zero value otherwise.

### GetHttpModeOk

`func (o *Dlb) GetHttpModeOk() (*string, bool)`

GetHttpModeOk returns a tuple with the HttpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMode

`func (o *Dlb) SetHttpMode(v string)`

SetHttpMode sets HttpMode field to given value.

### HasHttpMode

`func (o *Dlb) HasHttpMode() bool`

HasHttpMode returns a boolean if a field has been set.

### GetDefaultSslEndpoint

`func (o *Dlb) GetDefaultSslEndpoint() int32`

GetDefaultSslEndpoint returns the DefaultSslEndpoint field if non-nil, zero value otherwise.

### GetDefaultSslEndpointOk

`func (o *Dlb) GetDefaultSslEndpointOk() (*int32, bool)`

GetDefaultSslEndpointOk returns a tuple with the DefaultSslEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSslEndpoint

`func (o *Dlb) SetDefaultSslEndpoint(v int32)`

SetDefaultSslEndpoint sets DefaultSslEndpoint field to given value.

### HasDefaultSslEndpoint

`func (o *Dlb) HasDefaultSslEndpoint() bool`

HasDefaultSslEndpoint returns a boolean if a field has been set.

### GetTlsv1

`func (o *Dlb) GetTlsv1() bool`

GetTlsv1 returns the Tlsv1 field if non-nil, zero value otherwise.

### GetTlsv1Ok

`func (o *Dlb) GetTlsv1Ok() (*bool, bool)`

GetTlsv1Ok returns a tuple with the Tlsv1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsv1

`func (o *Dlb) SetTlsv1(v bool)`

SetTlsv1 sets Tlsv1 field to given value.

### HasTlsv1

`func (o *Dlb) HasTlsv1() bool`

HasTlsv1 returns a boolean if a field has been set.

### GetSslEndpoints

`func (o *Dlb) GetSslEndpoints() []DlbCoreSslEndpoints`

GetSslEndpoints returns the SslEndpoints field if non-nil, zero value otherwise.

### GetSslEndpointsOk

`func (o *Dlb) GetSslEndpointsOk() (*[]DlbCoreSslEndpoints, bool)`

GetSslEndpointsOk returns a tuple with the SslEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslEndpoints

`func (o *Dlb) SetSslEndpoints(v []DlbCoreSslEndpoints)`

SetSslEndpoints sets SslEndpoints field to given value.

### HasSslEndpoints

`func (o *Dlb) HasSslEndpoints() bool`

HasSslEndpoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


