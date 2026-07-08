# ExchangeBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueueId** | Pointer to **string** |  | [optional] 
**ExchangeId** | Pointer to **string** |  | [optional] 
**Fifo** | Pointer to **bool** |  | [optional] 

## Methods

### NewExchangeBinding

`func NewExchangeBinding() *ExchangeBinding`

NewExchangeBinding instantiates a new ExchangeBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeBindingWithDefaults

`func NewExchangeBindingWithDefaults() *ExchangeBinding`

NewExchangeBindingWithDefaults instantiates a new ExchangeBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueueId

`func (o *ExchangeBinding) GetQueueId() string`

GetQueueId returns the QueueId field if non-nil, zero value otherwise.

### GetQueueIdOk

`func (o *ExchangeBinding) GetQueueIdOk() (*string, bool)`

GetQueueIdOk returns a tuple with the QueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueId

`func (o *ExchangeBinding) SetQueueId(v string)`

SetQueueId sets QueueId field to given value.

### HasQueueId

`func (o *ExchangeBinding) HasQueueId() bool`

HasQueueId returns a boolean if a field has been set.

### GetExchangeId

`func (o *ExchangeBinding) GetExchangeId() string`

GetExchangeId returns the ExchangeId field if non-nil, zero value otherwise.

### GetExchangeIdOk

`func (o *ExchangeBinding) GetExchangeIdOk() (*string, bool)`

GetExchangeIdOk returns a tuple with the ExchangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeId

`func (o *ExchangeBinding) SetExchangeId(v string)`

SetExchangeId sets ExchangeId field to given value.

### HasExchangeId

`func (o *ExchangeBinding) HasExchangeId() bool`

HasExchangeId returns a boolean if a field has been set.

### GetFifo

`func (o *ExchangeBinding) GetFifo() bool`

GetFifo returns the Fifo field if non-nil, zero value otherwise.

### GetFifoOk

`func (o *ExchangeBinding) GetFifoOk() (*bool, bool)`

GetFifoOk returns a tuple with the Fifo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFifo

`func (o *ExchangeBinding) SetFifo(v bool)`

SetFifo sets Fifo field to given value.

### HasFifo

`func (o *ExchangeBinding) HasFifo() bool`

HasFifo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


