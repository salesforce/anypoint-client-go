# ExchangeBindingWithRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueueId** | Pointer to **string** |  | [optional] 
**ExchangeId** | Pointer to **string** |  | [optional] 
**Fifo** | Pointer to **bool** |  | [optional] 
**RoutingRules** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewExchangeBindingWithRules

`func NewExchangeBindingWithRules() *ExchangeBindingWithRules`

NewExchangeBindingWithRules instantiates a new ExchangeBindingWithRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeBindingWithRulesWithDefaults

`func NewExchangeBindingWithRulesWithDefaults() *ExchangeBindingWithRules`

NewExchangeBindingWithRulesWithDefaults instantiates a new ExchangeBindingWithRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueueId

`func (o *ExchangeBindingWithRules) GetQueueId() string`

GetQueueId returns the QueueId field if non-nil, zero value otherwise.

### GetQueueIdOk

`func (o *ExchangeBindingWithRules) GetQueueIdOk() (*string, bool)`

GetQueueIdOk returns a tuple with the QueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueId

`func (o *ExchangeBindingWithRules) SetQueueId(v string)`

SetQueueId sets QueueId field to given value.

### HasQueueId

`func (o *ExchangeBindingWithRules) HasQueueId() bool`

HasQueueId returns a boolean if a field has been set.

### GetExchangeId

`func (o *ExchangeBindingWithRules) GetExchangeId() string`

GetExchangeId returns the ExchangeId field if non-nil, zero value otherwise.

### GetExchangeIdOk

`func (o *ExchangeBindingWithRules) GetExchangeIdOk() (*string, bool)`

GetExchangeIdOk returns a tuple with the ExchangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeId

`func (o *ExchangeBindingWithRules) SetExchangeId(v string)`

SetExchangeId sets ExchangeId field to given value.

### HasExchangeId

`func (o *ExchangeBindingWithRules) HasExchangeId() bool`

HasExchangeId returns a boolean if a field has been set.

### GetFifo

`func (o *ExchangeBindingWithRules) GetFifo() bool`

GetFifo returns the Fifo field if non-nil, zero value otherwise.

### GetFifoOk

`func (o *ExchangeBindingWithRules) GetFifoOk() (*bool, bool)`

GetFifoOk returns a tuple with the Fifo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFifo

`func (o *ExchangeBindingWithRules) SetFifo(v bool)`

SetFifo sets Fifo field to given value.

### HasFifo

`func (o *ExchangeBindingWithRules) HasFifo() bool`

HasFifo returns a boolean if a field has been set.

### GetRoutingRules

`func (o *ExchangeBindingWithRules) GetRoutingRules() []map[string]interface{}`

GetRoutingRules returns the RoutingRules field if non-nil, zero value otherwise.

### GetRoutingRulesOk

`func (o *ExchangeBindingWithRules) GetRoutingRulesOk() (*[]map[string]interface{}, bool)`

GetRoutingRulesOk returns a tuple with the RoutingRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingRules

`func (o *ExchangeBindingWithRules) SetRoutingRules(v []map[string]interface{})`

SetRoutingRules sets RoutingRules field to given value.

### HasRoutingRules

`func (o *ExchangeBindingWithRules) HasRoutingRules() bool`

HasRoutingRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


