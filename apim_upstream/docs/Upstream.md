# Upstream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 
**TlsContext** | Pointer to [**UpstreamTlsContext**](UpstreamTlsContext.md) |  | [optional] 

## Methods

### NewUpstream

`func NewUpstream() *Upstream`

NewUpstream instantiates a new Upstream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpstreamWithDefaults

`func NewUpstreamWithDefaults() *Upstream`

NewUpstreamWithDefaults instantiates a new Upstream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Upstream) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Upstream) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Upstream) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Upstream) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *Upstream) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Upstream) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Upstream) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Upstream) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetUri

`func (o *Upstream) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *Upstream) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *Upstream) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *Upstream) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetTlsContext

`func (o *Upstream) GetTlsContext() UpstreamTlsContext`

GetTlsContext returns the TlsContext field if non-nil, zero value otherwise.

### GetTlsContextOk

`func (o *Upstream) GetTlsContextOk() (*UpstreamTlsContext, bool)`

GetTlsContextOk returns a tuple with the TlsContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsContext

`func (o *Upstream) SetTlsContext(v UpstreamTlsContext)`

SetTlsContext sets TlsContext field to given value.

### HasTlsContext

`func (o *Upstream) HasTlsContext() bool`

HasTlsContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


