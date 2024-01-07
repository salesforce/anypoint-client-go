# UpstreamPostBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 
**TlsContext** | Pointer to [**NullableUpstreamPostBodyTlsContext**](UpstreamPostBodyTlsContext.md) |  | [optional] 

## Methods

### NewUpstreamPostBody

`func NewUpstreamPostBody() *UpstreamPostBody`

NewUpstreamPostBody instantiates a new UpstreamPostBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpstreamPostBodyWithDefaults

`func NewUpstreamPostBodyWithDefaults() *UpstreamPostBody`

NewUpstreamPostBodyWithDefaults instantiates a new UpstreamPostBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *UpstreamPostBody) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UpstreamPostBody) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UpstreamPostBody) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *UpstreamPostBody) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetUri

`func (o *UpstreamPostBody) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *UpstreamPostBody) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *UpstreamPostBody) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *UpstreamPostBody) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetTlsContext

`func (o *UpstreamPostBody) GetTlsContext() UpstreamPostBodyTlsContext`

GetTlsContext returns the TlsContext field if non-nil, zero value otherwise.

### GetTlsContextOk

`func (o *UpstreamPostBody) GetTlsContextOk() (*UpstreamPostBodyTlsContext, bool)`

GetTlsContextOk returns a tuple with the TlsContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsContext

`func (o *UpstreamPostBody) SetTlsContext(v UpstreamPostBodyTlsContext)`

SetTlsContext sets TlsContext field to given value.

### HasTlsContext

`func (o *UpstreamPostBody) HasTlsContext() bool`

HasTlsContext returns a boolean if a field has been set.

### SetTlsContextNil

`func (o *UpstreamPostBody) SetTlsContextNil(b bool)`

 SetTlsContextNil sets the value for TlsContext to be an explicit nil

### UnsetTlsContext
`func (o *UpstreamPostBody) UnsetTlsContext()`

UnsetTlsContext ensures that no value is present for TlsContext, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


