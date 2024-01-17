# DlbExtras

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
**IpAddressesInfo** | Pointer to [**[]DlbExtrasIpAddressesInfoInner**](DlbExtrasIpAddressesInfoInner.md) |  | [optional] 
**DoubleStaticIps** | Pointer to **bool** |  | [optional] 
**EnableStreaming** | Pointer to **bool** |  | [optional] 
**ForwardClientCertificate** | Pointer to **bool** |  | [optional] 

## Methods

### NewDlbExtras

`func NewDlbExtras() *DlbExtras`

NewDlbExtras instantiates a new DlbExtras object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDlbExtrasWithDefaults

`func NewDlbExtrasWithDefaults() *DlbExtras`

NewDlbExtrasWithDefaults instantiates a new DlbExtras object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DlbExtras) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DlbExtras) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DlbExtras) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DlbExtras) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDomain

`func (o *DlbExtras) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DlbExtras) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DlbExtras) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *DlbExtras) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetDeploymentId

`func (o *DlbExtras) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *DlbExtras) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *DlbExtras) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *DlbExtras) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### GetInstanceConfig

`func (o *DlbExtras) GetInstanceConfig() InstanceConfig`

GetInstanceConfig returns the InstanceConfig field if non-nil, zero value otherwise.

### GetInstanceConfigOk

`func (o *DlbExtras) GetInstanceConfigOk() (*InstanceConfig, bool)`

GetInstanceConfigOk returns a tuple with the InstanceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceConfig

`func (o *DlbExtras) SetInstanceConfig(v InstanceConfig)`

SetInstanceConfig sets InstanceConfig field to given value.

### HasInstanceConfig

`func (o *DlbExtras) HasInstanceConfig() bool`

HasInstanceConfig returns a boolean if a field has been set.

### GetIpAddresses

`func (o *DlbExtras) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *DlbExtras) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *DlbExtras) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *DlbExtras) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### GetStaticIPsDisabled

`func (o *DlbExtras) GetStaticIPsDisabled() bool`

GetStaticIPsDisabled returns the StaticIPsDisabled field if non-nil, zero value otherwise.

### GetStaticIPsDisabledOk

`func (o *DlbExtras) GetStaticIPsDisabledOk() (*bool, bool)`

GetStaticIPsDisabledOk returns a tuple with the StaticIPsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticIPsDisabled

`func (o *DlbExtras) SetStaticIPsDisabled(v bool)`

SetStaticIPsDisabled sets StaticIPsDisabled field to given value.

### HasStaticIPsDisabled

`func (o *DlbExtras) HasStaticIPsDisabled() bool`

HasStaticIPsDisabled returns a boolean if a field has been set.

### GetVpcId

`func (o *DlbExtras) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *DlbExtras) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *DlbExtras) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *DlbExtras) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### GetWorkers

`func (o *DlbExtras) GetWorkers() int32`

GetWorkers returns the Workers field if non-nil, zero value otherwise.

### GetWorkersOk

`func (o *DlbExtras) GetWorkersOk() (*int32, bool)`

GetWorkersOk returns a tuple with the Workers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkers

`func (o *DlbExtras) SetWorkers(v int32)`

SetWorkers sets Workers field to given value.

### HasWorkers

`func (o *DlbExtras) HasWorkers() bool`

HasWorkers returns a boolean if a field has been set.

### GetDefaultCipherSuite

`func (o *DlbExtras) GetDefaultCipherSuite() string`

GetDefaultCipherSuite returns the DefaultCipherSuite field if non-nil, zero value otherwise.

### GetDefaultCipherSuiteOk

`func (o *DlbExtras) GetDefaultCipherSuiteOk() (*string, bool)`

GetDefaultCipherSuiteOk returns a tuple with the DefaultCipherSuite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCipherSuite

`func (o *DlbExtras) SetDefaultCipherSuite(v string)`

SetDefaultCipherSuite sets DefaultCipherSuite field to given value.

### HasDefaultCipherSuite

`func (o *DlbExtras) HasDefaultCipherSuite() bool`

HasDefaultCipherSuite returns a boolean if a field has been set.

### GetKeepUrlEncoding

`func (o *DlbExtras) GetKeepUrlEncoding() bool`

GetKeepUrlEncoding returns the KeepUrlEncoding field if non-nil, zero value otherwise.

### GetKeepUrlEncodingOk

`func (o *DlbExtras) GetKeepUrlEncodingOk() (*bool, bool)`

GetKeepUrlEncodingOk returns a tuple with the KeepUrlEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepUrlEncoding

`func (o *DlbExtras) SetKeepUrlEncoding(v bool)`

SetKeepUrlEncoding sets KeepUrlEncoding field to given value.

### HasKeepUrlEncoding

`func (o *DlbExtras) HasKeepUrlEncoding() bool`

HasKeepUrlEncoding returns a boolean if a field has been set.

### GetUpstreamTlsv12

`func (o *DlbExtras) GetUpstreamTlsv12() bool`

GetUpstreamTlsv12 returns the UpstreamTlsv12 field if non-nil, zero value otherwise.

### GetUpstreamTlsv12Ok

`func (o *DlbExtras) GetUpstreamTlsv12Ok() (*bool, bool)`

GetUpstreamTlsv12Ok returns a tuple with the UpstreamTlsv12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamTlsv12

`func (o *DlbExtras) SetUpstreamTlsv12(v bool)`

SetUpstreamTlsv12 sets UpstreamTlsv12 field to given value.

### HasUpstreamTlsv12

`func (o *DlbExtras) HasUpstreamTlsv12() bool`

HasUpstreamTlsv12 returns a boolean if a field has been set.

### GetProxyReadTimeout

`func (o *DlbExtras) GetProxyReadTimeout() int32`

GetProxyReadTimeout returns the ProxyReadTimeout field if non-nil, zero value otherwise.

### GetProxyReadTimeoutOk

`func (o *DlbExtras) GetProxyReadTimeoutOk() (*int32, bool)`

GetProxyReadTimeoutOk returns a tuple with the ProxyReadTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyReadTimeout

`func (o *DlbExtras) SetProxyReadTimeout(v int32)`

SetProxyReadTimeout sets ProxyReadTimeout field to given value.

### HasProxyReadTimeout

`func (o *DlbExtras) HasProxyReadTimeout() bool`

HasProxyReadTimeout returns a boolean if a field has been set.

### GetIpAddressesInfo

`func (o *DlbExtras) GetIpAddressesInfo() []DlbExtrasIpAddressesInfoInner`

GetIpAddressesInfo returns the IpAddressesInfo field if non-nil, zero value otherwise.

### GetIpAddressesInfoOk

`func (o *DlbExtras) GetIpAddressesInfoOk() (*[]DlbExtrasIpAddressesInfoInner, bool)`

GetIpAddressesInfoOk returns a tuple with the IpAddressesInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddressesInfo

`func (o *DlbExtras) SetIpAddressesInfo(v []DlbExtrasIpAddressesInfoInner)`

SetIpAddressesInfo sets IpAddressesInfo field to given value.

### HasIpAddressesInfo

`func (o *DlbExtras) HasIpAddressesInfo() bool`

HasIpAddressesInfo returns a boolean if a field has been set.

### GetDoubleStaticIps

`func (o *DlbExtras) GetDoubleStaticIps() bool`

GetDoubleStaticIps returns the DoubleStaticIps field if non-nil, zero value otherwise.

### GetDoubleStaticIpsOk

`func (o *DlbExtras) GetDoubleStaticIpsOk() (*bool, bool)`

GetDoubleStaticIpsOk returns a tuple with the DoubleStaticIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoubleStaticIps

`func (o *DlbExtras) SetDoubleStaticIps(v bool)`

SetDoubleStaticIps sets DoubleStaticIps field to given value.

### HasDoubleStaticIps

`func (o *DlbExtras) HasDoubleStaticIps() bool`

HasDoubleStaticIps returns a boolean if a field has been set.

### GetEnableStreaming

`func (o *DlbExtras) GetEnableStreaming() bool`

GetEnableStreaming returns the EnableStreaming field if non-nil, zero value otherwise.

### GetEnableStreamingOk

`func (o *DlbExtras) GetEnableStreamingOk() (*bool, bool)`

GetEnableStreamingOk returns a tuple with the EnableStreaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableStreaming

`func (o *DlbExtras) SetEnableStreaming(v bool)`

SetEnableStreaming sets EnableStreaming field to given value.

### HasEnableStreaming

`func (o *DlbExtras) HasEnableStreaming() bool`

HasEnableStreaming returns a boolean if a field has been set.

### GetForwardClientCertificate

`func (o *DlbExtras) GetForwardClientCertificate() bool`

GetForwardClientCertificate returns the ForwardClientCertificate field if non-nil, zero value otherwise.

### GetForwardClientCertificateOk

`func (o *DlbExtras) GetForwardClientCertificateOk() (*bool, bool)`

GetForwardClientCertificateOk returns a tuple with the ForwardClientCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardClientCertificate

`func (o *DlbExtras) SetForwardClientCertificate(v bool)`

SetForwardClientCertificate sets ForwardClientCertificate field to given value.

### HasForwardClientCertificate

`func (o *DlbExtras) HasForwardClientCertificate() bool`

HasForwardClientCertificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


