# Queue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueueId** | Pointer to **string** |  | [optional] 
**ExchangeId** | Pointer to **string** |  | [optional] 
**DefaultTtl** | Pointer to **int32** |  | [optional] 
**DefaultLockTtl** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Encrypted** | Pointer to **bool** |  | [optional] 
**DefaultDeliveryDelay** | Pointer to **int32** |  | [optional] 
**DeadLetterQueueId** | Pointer to **string** |  | [optional] 
**MaxDeliveries** | Pointer to **int32** |  | [optional] 
**Fifo** | Pointer to **bool** |  | [optional] 

## Methods

### NewQueue

`func NewQueue() *Queue`

NewQueue instantiates a new Queue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueWithDefaults

`func NewQueueWithDefaults() *Queue`

NewQueueWithDefaults instantiates a new Queue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueueId

`func (o *Queue) GetQueueId() string`

GetQueueId returns the QueueId field if non-nil, zero value otherwise.

### GetQueueIdOk

`func (o *Queue) GetQueueIdOk() (*string, bool)`

GetQueueIdOk returns a tuple with the QueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueId

`func (o *Queue) SetQueueId(v string)`

SetQueueId sets QueueId field to given value.

### HasQueueId

`func (o *Queue) HasQueueId() bool`

HasQueueId returns a boolean if a field has been set.

### GetExchangeId

`func (o *Queue) GetExchangeId() string`

GetExchangeId returns the ExchangeId field if non-nil, zero value otherwise.

### GetExchangeIdOk

`func (o *Queue) GetExchangeIdOk() (*string, bool)`

GetExchangeIdOk returns a tuple with the ExchangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeId

`func (o *Queue) SetExchangeId(v string)`

SetExchangeId sets ExchangeId field to given value.

### HasExchangeId

`func (o *Queue) HasExchangeId() bool`

HasExchangeId returns a boolean if a field has been set.

### GetDefaultTtl

`func (o *Queue) GetDefaultTtl() int32`

GetDefaultTtl returns the DefaultTtl field if non-nil, zero value otherwise.

### GetDefaultTtlOk

`func (o *Queue) GetDefaultTtlOk() (*int32, bool)`

GetDefaultTtlOk returns a tuple with the DefaultTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTtl

`func (o *Queue) SetDefaultTtl(v int32)`

SetDefaultTtl sets DefaultTtl field to given value.

### HasDefaultTtl

`func (o *Queue) HasDefaultTtl() bool`

HasDefaultTtl returns a boolean if a field has been set.

### GetDefaultLockTtl

`func (o *Queue) GetDefaultLockTtl() int32`

GetDefaultLockTtl returns the DefaultLockTtl field if non-nil, zero value otherwise.

### GetDefaultLockTtlOk

`func (o *Queue) GetDefaultLockTtlOk() (*int32, bool)`

GetDefaultLockTtlOk returns a tuple with the DefaultLockTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLockTtl

`func (o *Queue) SetDefaultLockTtl(v int32)`

SetDefaultLockTtl sets DefaultLockTtl field to given value.

### HasDefaultLockTtl

`func (o *Queue) HasDefaultLockTtl() bool`

HasDefaultLockTtl returns a boolean if a field has been set.

### GetType

`func (o *Queue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Queue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Queue) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Queue) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEncrypted

`func (o *Queue) GetEncrypted() bool`

GetEncrypted returns the Encrypted field if non-nil, zero value otherwise.

### GetEncryptedOk

`func (o *Queue) GetEncryptedOk() (*bool, bool)`

GetEncryptedOk returns a tuple with the Encrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypted

`func (o *Queue) SetEncrypted(v bool)`

SetEncrypted sets Encrypted field to given value.

### HasEncrypted

`func (o *Queue) HasEncrypted() bool`

HasEncrypted returns a boolean if a field has been set.

### GetDefaultDeliveryDelay

`func (o *Queue) GetDefaultDeliveryDelay() int32`

GetDefaultDeliveryDelay returns the DefaultDeliveryDelay field if non-nil, zero value otherwise.

### GetDefaultDeliveryDelayOk

`func (o *Queue) GetDefaultDeliveryDelayOk() (*int32, bool)`

GetDefaultDeliveryDelayOk returns a tuple with the DefaultDeliveryDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDeliveryDelay

`func (o *Queue) SetDefaultDeliveryDelay(v int32)`

SetDefaultDeliveryDelay sets DefaultDeliveryDelay field to given value.

### HasDefaultDeliveryDelay

`func (o *Queue) HasDefaultDeliveryDelay() bool`

HasDefaultDeliveryDelay returns a boolean if a field has been set.

### GetDeadLetterQueueId

`func (o *Queue) GetDeadLetterQueueId() string`

GetDeadLetterQueueId returns the DeadLetterQueueId field if non-nil, zero value otherwise.

### GetDeadLetterQueueIdOk

`func (o *Queue) GetDeadLetterQueueIdOk() (*string, bool)`

GetDeadLetterQueueIdOk returns a tuple with the DeadLetterQueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadLetterQueueId

`func (o *Queue) SetDeadLetterQueueId(v string)`

SetDeadLetterQueueId sets DeadLetterQueueId field to given value.

### HasDeadLetterQueueId

`func (o *Queue) HasDeadLetterQueueId() bool`

HasDeadLetterQueueId returns a boolean if a field has been set.

### GetMaxDeliveries

`func (o *Queue) GetMaxDeliveries() int32`

GetMaxDeliveries returns the MaxDeliveries field if non-nil, zero value otherwise.

### GetMaxDeliveriesOk

`func (o *Queue) GetMaxDeliveriesOk() (*int32, bool)`

GetMaxDeliveriesOk returns a tuple with the MaxDeliveries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDeliveries

`func (o *Queue) SetMaxDeliveries(v int32)`

SetMaxDeliveries sets MaxDeliveries field to given value.

### HasMaxDeliveries

`func (o *Queue) HasMaxDeliveries() bool`

HasMaxDeliveries returns a boolean if a field has been set.

### GetFifo

`func (o *Queue) GetFifo() bool`

GetFifo returns the Fifo field if non-nil, zero value otherwise.

### GetFifoOk

`func (o *Queue) GetFifoOk() (*bool, bool)`

GetFifoOk returns a tuple with the Fifo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFifo

`func (o *Queue) SetFifo(v bool)`

SetFifo sets Fifo field to given value.

### HasFifo

`func (o *Queue) HasFifo() bool`

HasFifo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


