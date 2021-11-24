# QueueBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultTtl** | Pointer to **int32** |  | [optional] 
**DefaultLockTtl** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Encrypted** | Pointer to **bool** |  | [optional] 
**DefaultDeliveryDelay** | Pointer to **int32** |  | [optional] 
**Fifo** | Pointer to **bool** |  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


