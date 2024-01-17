# Routing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** |  | [optional] 
**Rules** | Pointer to [**RoutingRules**](RoutingRules.md) |  | [optional] 
**Upstreams** | Pointer to [**[]RoutingUpstreamsInner**](RoutingUpstreamsInner.md) |  | [optional] 

## Methods

### NewRouting

`func NewRouting() *Routing`

NewRouting instantiates a new Routing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingWithDefaults

`func NewRoutingWithDefaults() *Routing`

NewRoutingWithDefaults instantiates a new Routing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *Routing) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Routing) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Routing) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Routing) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetRules

`func (o *Routing) GetRules() RoutingRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *Routing) GetRulesOk() (*RoutingRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *Routing) SetRules(v RoutingRules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *Routing) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetUpstreams

`func (o *Routing) GetUpstreams() []RoutingUpstreamsInner`

GetUpstreams returns the Upstreams field if non-nil, zero value otherwise.

### GetUpstreamsOk

`func (o *Routing) GetUpstreamsOk() (*[]RoutingUpstreamsInner, bool)`

GetUpstreamsOk returns a tuple with the Upstreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreams

`func (o *Routing) SetUpstreams(v []RoutingUpstreamsInner)`

SetUpstreams sets Upstreams field to given value.

### HasUpstreams

`func (o *Routing) HasUpstreams() bool`

HasUpstreams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


