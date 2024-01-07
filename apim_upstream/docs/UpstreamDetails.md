# UpstreamDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audit** | Pointer to [**Audit**](Audit.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 
**TlsContext** | Pointer to [**NullableUpstreamDetailsTlsContext**](UpstreamDetailsTlsContext.md) |  | [optional] 

## Methods

### NewUpstreamDetails

`func NewUpstreamDetails() *UpstreamDetails`

NewUpstreamDetails instantiates a new UpstreamDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpstreamDetailsWithDefaults

`func NewUpstreamDetailsWithDefaults() *UpstreamDetails`

NewUpstreamDetailsWithDefaults instantiates a new UpstreamDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudit

`func (o *UpstreamDetails) GetAudit() Audit`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *UpstreamDetails) GetAuditOk() (*Audit, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *UpstreamDetails) SetAudit(v Audit)`

SetAudit sets Audit field to given value.

### HasAudit

`func (o *UpstreamDetails) HasAudit() bool`

HasAudit returns a boolean if a field has been set.

### GetId

`func (o *UpstreamDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpstreamDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpstreamDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpstreamDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *UpstreamDetails) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UpstreamDetails) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UpstreamDetails) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *UpstreamDetails) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetUri

`func (o *UpstreamDetails) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *UpstreamDetails) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *UpstreamDetails) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *UpstreamDetails) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetTlsContext

`func (o *UpstreamDetails) GetTlsContext() UpstreamDetailsTlsContext`

GetTlsContext returns the TlsContext field if non-nil, zero value otherwise.

### GetTlsContextOk

`func (o *UpstreamDetails) GetTlsContextOk() (*UpstreamDetailsTlsContext, bool)`

GetTlsContextOk returns a tuple with the TlsContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsContext

`func (o *UpstreamDetails) SetTlsContext(v UpstreamDetailsTlsContext)`

SetTlsContext sets TlsContext field to given value.

### HasTlsContext

`func (o *UpstreamDetails) HasTlsContext() bool`

HasTlsContext returns a boolean if a field has been set.

### SetTlsContextNil

`func (o *UpstreamDetails) SetTlsContextNil(b bool)`

 SetTlsContextNil sets the value for TlsContext to be an explicit nil

### UnsetTlsContext
`func (o *UpstreamDetails) UnsetTlsContext()`

UnsetTlsContext ensures that no value is present for TlsContext, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


