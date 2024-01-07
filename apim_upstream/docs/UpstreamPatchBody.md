# UpstreamPatchBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 
**TlsContext** | Pointer to [**NullableUpstreamPostBodyTlsContext**](UpstreamPostBodyTlsContext.md) |  | [optional] 

## Methods

### NewUpstreamPatchBody

`func NewUpstreamPatchBody() *UpstreamPatchBody`

NewUpstreamPatchBody instantiates a new UpstreamPatchBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpstreamPatchBodyWithDefaults

`func NewUpstreamPatchBodyWithDefaults() *UpstreamPatchBody`

NewUpstreamPatchBodyWithDefaults instantiates a new UpstreamPatchBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *UpstreamPatchBody) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UpstreamPatchBody) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UpstreamPatchBody) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *UpstreamPatchBody) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetUri

`func (o *UpstreamPatchBody) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *UpstreamPatchBody) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *UpstreamPatchBody) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *UpstreamPatchBody) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetTlsContext

`func (o *UpstreamPatchBody) GetTlsContext() UpstreamPostBodyTlsContext`

GetTlsContext returns the TlsContext field if non-nil, zero value otherwise.

### GetTlsContextOk

`func (o *UpstreamPatchBody) GetTlsContextOk() (*UpstreamPostBodyTlsContext, bool)`

GetTlsContextOk returns a tuple with the TlsContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsContext

`func (o *UpstreamPatchBody) SetTlsContext(v UpstreamPostBodyTlsContext)`

SetTlsContext sets TlsContext field to given value.

### HasTlsContext

`func (o *UpstreamPatchBody) HasTlsContext() bool`

HasTlsContext returns a boolean if a field has been set.

### SetTlsContextNil

`func (o *UpstreamPatchBody) SetTlsContextNil(b bool)`

 SetTlsContextNil sets the value for TlsContext to be an explicit nil

### UnsetTlsContext
`func (o *UpstreamPatchBody) UnsetTlsContext()`

UnsetTlsContext ensures that no value is present for TlsContext, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


