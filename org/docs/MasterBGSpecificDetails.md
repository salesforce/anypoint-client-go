# MasterBGSpecificDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionTimeout** | Pointer to **int32** | An explanation about the purpose of this instance. | [optional] [default to 0]
**Subscription** | Pointer to [**Subscription**](Subscription.md) |  | [optional] 

## Methods

### NewMasterBGSpecificDetails

`func NewMasterBGSpecificDetails() *MasterBGSpecificDetails`

NewMasterBGSpecificDetails instantiates a new MasterBGSpecificDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMasterBGSpecificDetailsWithDefaults

`func NewMasterBGSpecificDetailsWithDefaults() *MasterBGSpecificDetails`

NewMasterBGSpecificDetailsWithDefaults instantiates a new MasterBGSpecificDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionTimeout

`func (o *MasterBGSpecificDetails) GetSessionTimeout() int32`

GetSessionTimeout returns the SessionTimeout field if non-nil, zero value otherwise.

### GetSessionTimeoutOk

`func (o *MasterBGSpecificDetails) GetSessionTimeoutOk() (*int32, bool)`

GetSessionTimeoutOk returns a tuple with the SessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTimeout

`func (o *MasterBGSpecificDetails) SetSessionTimeout(v int32)`

SetSessionTimeout sets SessionTimeout field to given value.

### HasSessionTimeout

`func (o *MasterBGSpecificDetails) HasSessionTimeout() bool`

HasSessionTimeout returns a boolean if a field has been set.

### GetSubscription

`func (o *MasterBGSpecificDetails) GetSubscription() Subscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *MasterBGSpecificDetails) GetSubscriptionOk() (*Subscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *MasterBGSpecificDetails) SetSubscription(v Subscription)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *MasterBGSpecificDetails) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


