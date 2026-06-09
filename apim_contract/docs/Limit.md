# Limit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaximumRequests** | Pointer to **int64** |  | [optional] 
**TimePeriodInMilliseconds** | Pointer to **int64** |  | [optional] 
**Visible** | Pointer to **bool** |  | [optional] 

## Methods

### NewLimit

`func NewLimit() *Limit`

NewLimit instantiates a new Limit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLimitWithDefaults

`func NewLimitWithDefaults() *Limit`

NewLimitWithDefaults instantiates a new Limit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaximumRequests

`func (o *Limit) GetMaximumRequests() int64`

GetMaximumRequests returns the MaximumRequests field if non-nil, zero value otherwise.

### GetMaximumRequestsOk

`func (o *Limit) GetMaximumRequestsOk() (*int64, bool)`

GetMaximumRequestsOk returns a tuple with the MaximumRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumRequests

`func (o *Limit) SetMaximumRequests(v int64)`

SetMaximumRequests sets MaximumRequests field to given value.

### HasMaximumRequests

`func (o *Limit) HasMaximumRequests() bool`

HasMaximumRequests returns a boolean if a field has been set.

### GetTimePeriodInMilliseconds

`func (o *Limit) GetTimePeriodInMilliseconds() int64`

GetTimePeriodInMilliseconds returns the TimePeriodInMilliseconds field if non-nil, zero value otherwise.

### GetTimePeriodInMillisecondsOk

`func (o *Limit) GetTimePeriodInMillisecondsOk() (*int64, bool)`

GetTimePeriodInMillisecondsOk returns a tuple with the TimePeriodInMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePeriodInMilliseconds

`func (o *Limit) SetTimePeriodInMilliseconds(v int64)`

SetTimePeriodInMilliseconds sets TimePeriodInMilliseconds field to given value.

### HasTimePeriodInMilliseconds

`func (o *Limit) HasTimePeriodInMilliseconds() bool`

HasTimePeriodInMilliseconds returns a boolean if a field has been set.

### GetVisible

`func (o *Limit) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *Limit) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *Limit) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *Limit) HasVisible() bool`

HasVisible returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


