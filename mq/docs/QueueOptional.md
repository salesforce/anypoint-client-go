# QueueOptional

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

### NewQueueOptional

`func NewQueueOptional() *QueueOptional`

NewQueueOptional instantiates a new QueueOptional object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueOptionalWithDefaults

`func NewQueueOptionalWithDefaults() *QueueOptional`

NewQueueOptionalWithDefaults instantiates a new QueueOptional object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultTtl

`func (o *QueueOptional) GetDefaultTtl() int32`

GetDefaultTtl returns the DefaultTtl field if non-nil, zero value otherwise.

### GetDefaultTtlOk

`func (o *QueueOptional) GetDefaultTtlOk() (*int32, bool)`

GetDefaultTtlOk returns a tuple with the DefaultTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTtl

`func (o *QueueOptional) SetDefaultTtl(v int32)`

SetDefaultTtl sets DefaultTtl field to given value.

### HasDefaultTtl

`func (o *QueueOptional) HasDefaultTtl() bool`

HasDefaultTtl returns a boolean if a field has been set.

### GetDefaultLockTtl

`func (o *QueueOptional) GetDefaultLockTtl() int32`

GetDefaultLockTtl returns the DefaultLockTtl field if non-nil, zero value otherwise.

### GetDefaultLockTtlOk

`func (o *QueueOptional) GetDefaultLockTtlOk() (*int32, bool)`

GetDefaultLockTtlOk returns a tuple with the DefaultLockTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLockTtl

`func (o *QueueOptional) SetDefaultLockTtl(v int32)`

SetDefaultLockTtl sets DefaultLockTtl field to given value.

### HasDefaultLockTtl

`func (o *QueueOptional) HasDefaultLockTtl() bool`

HasDefaultLockTtl returns a boolean if a field has been set.

### GetType

`func (o *QueueOptional) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QueueOptional) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QueueOptional) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *QueueOptional) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEncrypted

`func (o *QueueOptional) GetEncrypted() bool`

GetEncrypted returns the Encrypted field if non-nil, zero value otherwise.

### GetEncryptedOk

`func (o *QueueOptional) GetEncryptedOk() (*bool, bool)`

GetEncryptedOk returns a tuple with the Encrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypted

`func (o *QueueOptional) SetEncrypted(v bool)`

SetEncrypted sets Encrypted field to given value.

### HasEncrypted

`func (o *QueueOptional) HasEncrypted() bool`

HasEncrypted returns a boolean if a field has been set.

### GetDefaultDeliveryDelay

`func (o *QueueOptional) GetDefaultDeliveryDelay() int32`

GetDefaultDeliveryDelay returns the DefaultDeliveryDelay field if non-nil, zero value otherwise.

### GetDefaultDeliveryDelayOk

`func (o *QueueOptional) GetDefaultDeliveryDelayOk() (*int32, bool)`

GetDefaultDeliveryDelayOk returns a tuple with the DefaultDeliveryDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDeliveryDelay

`func (o *QueueOptional) SetDefaultDeliveryDelay(v int32)`

SetDefaultDeliveryDelay sets DefaultDeliveryDelay field to given value.

### HasDefaultDeliveryDelay

`func (o *QueueOptional) HasDefaultDeliveryDelay() bool`

HasDefaultDeliveryDelay returns a boolean if a field has been set.

### GetFifo

`func (o *QueueOptional) GetFifo() bool`

GetFifo returns the Fifo field if non-nil, zero value otherwise.

### GetFifoOk

`func (o *QueueOptional) GetFifoOk() (*bool, bool)`

GetFifoOk returns a tuple with the Fifo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFifo

`func (o *QueueOptional) SetFifo(v bool)`

SetFifo sets Fifo field to given value.

### HasFifo

`func (o *QueueOptional) HasFifo() bool`

HasFifo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


