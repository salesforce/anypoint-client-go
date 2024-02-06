# EndpointTlsContexts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audit** | Pointer to [**Audit**](Audit.md) |  | [optional] 
**Inbound** | Pointer to [**NullableEndpointTlsContextsInbound**](EndpointTlsContextsInbound.md) |  | [optional] 
**Outbound** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewEndpointTlsContexts

`func NewEndpointTlsContexts() *EndpointTlsContexts`

NewEndpointTlsContexts instantiates a new EndpointTlsContexts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointTlsContextsWithDefaults

`func NewEndpointTlsContextsWithDefaults() *EndpointTlsContexts`

NewEndpointTlsContextsWithDefaults instantiates a new EndpointTlsContexts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudit

`func (o *EndpointTlsContexts) GetAudit() Audit`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *EndpointTlsContexts) GetAuditOk() (*Audit, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *EndpointTlsContexts) SetAudit(v Audit)`

SetAudit sets Audit field to given value.

### HasAudit

`func (o *EndpointTlsContexts) HasAudit() bool`

HasAudit returns a boolean if a field has been set.

### GetInbound

`func (o *EndpointTlsContexts) GetInbound() EndpointTlsContextsInbound`

GetInbound returns the Inbound field if non-nil, zero value otherwise.

### GetInboundOk

`func (o *EndpointTlsContexts) GetInboundOk() (*EndpointTlsContextsInbound, bool)`

GetInboundOk returns a tuple with the Inbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbound

`func (o *EndpointTlsContexts) SetInbound(v EndpointTlsContextsInbound)`

SetInbound sets Inbound field to given value.

### HasInbound

`func (o *EndpointTlsContexts) HasInbound() bool`

HasInbound returns a boolean if a field has been set.

### SetInboundNil

`func (o *EndpointTlsContexts) SetInboundNil(b bool)`

 SetInboundNil sets the value for Inbound to be an explicit nil

### UnsetInbound
`func (o *EndpointTlsContexts) UnsetInbound()`

UnsetInbound ensures that no value is present for Inbound, not even an explicit nil
### GetOutbound

`func (o *EndpointTlsContexts) GetOutbound() map[string]interface{}`

GetOutbound returns the Outbound field if non-nil, zero value otherwise.

### GetOutboundOk

`func (o *EndpointTlsContexts) GetOutboundOk() (*map[string]interface{}, bool)`

GetOutboundOk returns a tuple with the Outbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutbound

`func (o *EndpointTlsContexts) SetOutbound(v map[string]interface{})`

SetOutbound sets Outbound field to given value.

### HasOutbound

`func (o *EndpointTlsContexts) HasOutbound() bool`

HasOutbound returns a boolean if a field has been set.

### SetOutboundNil

`func (o *EndpointTlsContexts) SetOutboundNil(b bool)`

 SetOutboundNil sets the value for Outbound to be an explicit nil

### UnsetOutbound
`func (o *EndpointTlsContexts) UnsetOutbound()`

UnsetOutbound ensures that no value is present for Outbound, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


