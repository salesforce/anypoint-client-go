# InternalDnsSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsServers** | **[]string** |  | [default to []]
**SpecialDomains** | **[]string** |  | [default to []]

## Methods

### NewInternalDnsSchema

`func NewInternalDnsSchema(dnsServers []string, specialDomains []string, ) *InternalDnsSchema`

NewInternalDnsSchema instantiates a new InternalDnsSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalDnsSchemaWithDefaults

`func NewInternalDnsSchemaWithDefaults() *InternalDnsSchema`

NewInternalDnsSchemaWithDefaults instantiates a new InternalDnsSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsServers

`func (o *InternalDnsSchema) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *InternalDnsSchema) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *InternalDnsSchema) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.


### GetSpecialDomains

`func (o *InternalDnsSchema) GetSpecialDomains() []string`

GetSpecialDomains returns the SpecialDomains field if non-nil, zero value otherwise.

### GetSpecialDomainsOk

`func (o *InternalDnsSchema) GetSpecialDomainsOk() (*[]string, bool)`

GetSpecialDomainsOk returns a tuple with the SpecialDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialDomains

`func (o *InternalDnsSchema) SetSpecialDomains(v []string)`

SetSpecialDomains sets SpecialDomains field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


