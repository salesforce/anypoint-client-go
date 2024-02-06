# UpstreamTlsContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecretGroupId** | Pointer to **string** |  | [optional] 
**TlsContextId** | Pointer to **string** |  | [optional] 
**Audit** | Pointer to [**Audit**](Audit.md) |  | [optional] 

## Methods

### NewUpstreamTlsContext

`func NewUpstreamTlsContext() *UpstreamTlsContext`

NewUpstreamTlsContext instantiates a new UpstreamTlsContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpstreamTlsContextWithDefaults

`func NewUpstreamTlsContextWithDefaults() *UpstreamTlsContext`

NewUpstreamTlsContextWithDefaults instantiates a new UpstreamTlsContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecretGroupId

`func (o *UpstreamTlsContext) GetSecretGroupId() string`

GetSecretGroupId returns the SecretGroupId field if non-nil, zero value otherwise.

### GetSecretGroupIdOk

`func (o *UpstreamTlsContext) GetSecretGroupIdOk() (*string, bool)`

GetSecretGroupIdOk returns a tuple with the SecretGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretGroupId

`func (o *UpstreamTlsContext) SetSecretGroupId(v string)`

SetSecretGroupId sets SecretGroupId field to given value.

### HasSecretGroupId

`func (o *UpstreamTlsContext) HasSecretGroupId() bool`

HasSecretGroupId returns a boolean if a field has been set.

### GetTlsContextId

`func (o *UpstreamTlsContext) GetTlsContextId() string`

GetTlsContextId returns the TlsContextId field if non-nil, zero value otherwise.

### GetTlsContextIdOk

`func (o *UpstreamTlsContext) GetTlsContextIdOk() (*string, bool)`

GetTlsContextIdOk returns a tuple with the TlsContextId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsContextId

`func (o *UpstreamTlsContext) SetTlsContextId(v string)`

SetTlsContextId sets TlsContextId field to given value.

### HasTlsContextId

`func (o *UpstreamTlsContext) HasTlsContextId() bool`

HasTlsContextId returns a boolean if a field has been set.

### GetAudit

`func (o *UpstreamTlsContext) GetAudit() Audit`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *UpstreamTlsContext) GetAuditOk() (*Audit, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *UpstreamTlsContext) SetAudit(v Audit)`

SetAudit sets Audit field to given value.

### HasAudit

`func (o *UpstreamTlsContext) HasAudit() bool`

HasAudit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


