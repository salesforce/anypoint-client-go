# RoutingPostBodyInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** |  | [optional] 
**Rules** | Pointer to [**RoutingRules**](RoutingRules.md) |  | [optional] 
**Upstreams** | Pointer to [**RoutingPostBodyInnerUpstreams**](RoutingPostBodyInnerUpstreams.md) |  | [optional] 

## Methods

### NewRoutingPostBodyInner

`func NewRoutingPostBodyInner() *RoutingPostBodyInner`

NewRoutingPostBodyInner instantiates a new RoutingPostBodyInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingPostBodyInnerWithDefaults

`func NewRoutingPostBodyInnerWithDefaults() *RoutingPostBodyInner`

NewRoutingPostBodyInnerWithDefaults instantiates a new RoutingPostBodyInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *RoutingPostBodyInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *RoutingPostBodyInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *RoutingPostBodyInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *RoutingPostBodyInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetRules

`func (o *RoutingPostBodyInner) GetRules() RoutingRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *RoutingPostBodyInner) GetRulesOk() (*RoutingRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *RoutingPostBodyInner) SetRules(v RoutingRules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *RoutingPostBodyInner) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetUpstreams

`func (o *RoutingPostBodyInner) GetUpstreams() RoutingPostBodyInnerUpstreams`

GetUpstreams returns the Upstreams field if non-nil, zero value otherwise.

### GetUpstreamsOk

`func (o *RoutingPostBodyInner) GetUpstreamsOk() (*RoutingPostBodyInnerUpstreams, bool)`

GetUpstreamsOk returns a tuple with the Upstreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreams

`func (o *RoutingPostBodyInner) SetUpstreams(v RoutingPostBodyInnerUpstreams)`

SetUpstreams sets Upstreams field to given value.

### HasUpstreams

`func (o *RoutingPostBodyInner) HasUpstreams() bool`

HasUpstreams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


