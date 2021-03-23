# InternalDns

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsServers** | **[]string** |  | [default to []]
**SpecialDomains** | **[]string** |  | [default to []]

## Methods

### NewInternalDns

`func NewInternalDns(dnsServers []string, specialDomains []string, ) *InternalDns`

NewInternalDns instantiates a new InternalDns object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalDnsWithDefaults

`func NewInternalDnsWithDefaults() *InternalDns`

NewInternalDnsWithDefaults instantiates a new InternalDns object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsServers

`func (o *InternalDns) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *InternalDns) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *InternalDns) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.


### GetSpecialDomains

`func (o *InternalDns) GetSpecialDomains() []string`

GetSpecialDomains returns the SpecialDomains field if non-nil, zero value otherwise.

### GetSpecialDomainsOk

`func (o *InternalDns) GetSpecialDomainsOk() (*[]string, bool)`

GetSpecialDomainsOk returns a tuple with the SpecialDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialDomains

`func (o *InternalDns) SetSpecialDomains(v []string)`

SetSpecialDomains sets SpecialDomains field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


