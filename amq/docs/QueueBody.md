# QueueBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultTtl** | Pointer to **NullableInt32** |  | [optional] 
**DefaultLockTtl** | Pointer to **NullableInt32** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Encrypted** | Pointer to **NullableBool** |  | [optional] 
**DefaultDeliveryDelay** | Pointer to **NullableInt32** |  | [optional] 
**DeadLetterQueueId** | Pointer to **NullableString** |  | [optional] 
**MaxDeliveries** | Pointer to **NullableInt32** |  | [optional] 
**Fifo** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewQueueBody

`func NewQueueBody() *QueueBody`

NewQueueBody instantiates a new QueueBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueBodyWithDefaults

`func NewQueueBodyWithDefaults() *QueueBody`

NewQueueBodyWithDefaults instantiates a new QueueBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultTtl

`func (o *QueueBody) GetDefaultTtl() int32`

GetDefaultTtl returns the DefaultTtl field if non-nil, zero value otherwise.

### GetDefaultTtlOk

`func (o *QueueBody) GetDefaultTtlOk() (*int32, bool)`

GetDefaultTtlOk returns a tuple with the DefaultTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTtl

`func (o *QueueBody) SetDefaultTtl(v int32)`

SetDefaultTtl sets DefaultTtl field to given value.

### HasDefaultTtl

`func (o *QueueBody) HasDefaultTtl() bool`

HasDefaultTtl returns a boolean if a field has been set.

### SetDefaultTtlNil

`func (o *QueueBody) SetDefaultTtlNil(b bool)`

 SetDefaultTtlNil sets the value for DefaultTtl to be an explicit nil

### UnsetDefaultTtl
`func (o *QueueBody) UnsetDefaultTtl()`

UnsetDefaultTtl ensures that no value is present for DefaultTtl, not even an explicit nil
### GetDefaultLockTtl

`func (o *QueueBody) GetDefaultLockTtl() int32`

GetDefaultLockTtl returns the DefaultLockTtl field if non-nil, zero value otherwise.

### GetDefaultLockTtlOk

`func (o *QueueBody) GetDefaultLockTtlOk() (*int32, bool)`

GetDefaultLockTtlOk returns a tuple with the DefaultLockTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLockTtl

`func (o *QueueBody) SetDefaultLockTtl(v int32)`

SetDefaultLockTtl sets DefaultLockTtl field to given value.

### HasDefaultLockTtl

`func (o *QueueBody) HasDefaultLockTtl() bool`

HasDefaultLockTtl returns a boolean if a field has been set.

### SetDefaultLockTtlNil

`func (o *QueueBody) SetDefaultLockTtlNil(b bool)`

 SetDefaultLockTtlNil sets the value for DefaultLockTtl to be an explicit nil

### UnsetDefaultLockTtl
`func (o *QueueBody) UnsetDefaultLockTtl()`

UnsetDefaultLockTtl ensures that no value is present for DefaultLockTtl, not even an explicit nil
### GetType

`func (o *QueueBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QueueBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QueueBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *QueueBody) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *QueueBody) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *QueueBody) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetEncrypted

`func (o *QueueBody) GetEncrypted() bool`

GetEncrypted returns the Encrypted field if non-nil, zero value otherwise.

### GetEncryptedOk

`func (o *QueueBody) GetEncryptedOk() (*bool, bool)`

GetEncryptedOk returns a tuple with the Encrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypted

`func (o *QueueBody) SetEncrypted(v bool)`

SetEncrypted sets Encrypted field to given value.

### HasEncrypted

`func (o *QueueBody) HasEncrypted() bool`

HasEncrypted returns a boolean if a field has been set.

### SetEncryptedNil

`func (o *QueueBody) SetEncryptedNil(b bool)`

 SetEncryptedNil sets the value for Encrypted to be an explicit nil

### UnsetEncrypted
`func (o *QueueBody) UnsetEncrypted()`

UnsetEncrypted ensures that no value is present for Encrypted, not even an explicit nil
### GetDefaultDeliveryDelay

`func (o *QueueBody) GetDefaultDeliveryDelay() int32`

GetDefaultDeliveryDelay returns the DefaultDeliveryDelay field if non-nil, zero value otherwise.

### GetDefaultDeliveryDelayOk

`func (o *QueueBody) GetDefaultDeliveryDelayOk() (*int32, bool)`

GetDefaultDeliveryDelayOk returns a tuple with the DefaultDeliveryDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDeliveryDelay

`func (o *QueueBody) SetDefaultDeliveryDelay(v int32)`

SetDefaultDeliveryDelay sets DefaultDeliveryDelay field to given value.

### HasDefaultDeliveryDelay

`func (o *QueueBody) HasDefaultDeliveryDelay() bool`

HasDefaultDeliveryDelay returns a boolean if a field has been set.

### SetDefaultDeliveryDelayNil

`func (o *QueueBody) SetDefaultDeliveryDelayNil(b bool)`

 SetDefaultDeliveryDelayNil sets the value for DefaultDeliveryDelay to be an explicit nil

### UnsetDefaultDeliveryDelay
`func (o *QueueBody) UnsetDefaultDeliveryDelay()`

UnsetDefaultDeliveryDelay ensures that no value is present for DefaultDeliveryDelay, not even an explicit nil
### GetDeadLetterQueueId

`func (o *QueueBody) GetDeadLetterQueueId() string`

GetDeadLetterQueueId returns the DeadLetterQueueId field if non-nil, zero value otherwise.

### GetDeadLetterQueueIdOk

`func (o *QueueBody) GetDeadLetterQueueIdOk() (*string, bool)`

GetDeadLetterQueueIdOk returns a tuple with the DeadLetterQueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadLetterQueueId

`func (o *QueueBody) SetDeadLetterQueueId(v string)`

SetDeadLetterQueueId sets DeadLetterQueueId field to given value.

### HasDeadLetterQueueId

`func (o *QueueBody) HasDeadLetterQueueId() bool`

HasDeadLetterQueueId returns a boolean if a field has been set.

### SetDeadLetterQueueIdNil

`func (o *QueueBody) SetDeadLetterQueueIdNil(b bool)`

 SetDeadLetterQueueIdNil sets the value for DeadLetterQueueId to be an explicit nil

### UnsetDeadLetterQueueId
`func (o *QueueBody) UnsetDeadLetterQueueId()`

UnsetDeadLetterQueueId ensures that no value is present for DeadLetterQueueId, not even an explicit nil
### GetMaxDeliveries

`func (o *QueueBody) GetMaxDeliveries() int32`

GetMaxDeliveries returns the MaxDeliveries field if non-nil, zero value otherwise.

### GetMaxDeliveriesOk

`func (o *QueueBody) GetMaxDeliveriesOk() (*int32, bool)`

GetMaxDeliveriesOk returns a tuple with the MaxDeliveries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDeliveries

`func (o *QueueBody) SetMaxDeliveries(v int32)`

SetMaxDeliveries sets MaxDeliveries field to given value.

### HasMaxDeliveries

`func (o *QueueBody) HasMaxDeliveries() bool`

HasMaxDeliveries returns a boolean if a field has been set.

### SetMaxDeliveriesNil

`func (o *QueueBody) SetMaxDeliveriesNil(b bool)`

 SetMaxDeliveriesNil sets the value for MaxDeliveries to be an explicit nil

### UnsetMaxDeliveries
`func (o *QueueBody) UnsetMaxDeliveries()`

UnsetMaxDeliveries ensures that no value is present for MaxDeliveries, not even an explicit nil
### GetFifo

`func (o *QueueBody) GetFifo() bool`

GetFifo returns the Fifo field if non-nil, zero value otherwise.

### GetFifoOk

`func (o *QueueBody) GetFifoOk() (*bool, bool)`

GetFifoOk returns a tuple with the Fifo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFifo

`func (o *QueueBody) SetFifo(v bool)`

SetFifo sets Fifo field to given value.

### HasFifo

`func (o *QueueBody) HasFifo() bool`

HasFifo returns a boolean if a field has been set.

### SetFifoNil

`func (o *QueueBody) SetFifoNil(b bool)`

 SetFifoNil sets the value for Fifo to be an explicit nil

### UnsetFifo
`func (o *QueueBody) UnsetFifo()`

UnsetFifo ensures that no value is present for Fifo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


