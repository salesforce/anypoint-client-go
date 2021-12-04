# Period

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | [**Duration**](Duration.md) |  | 
**Repeat** | **int32** |  | 

## Methods

### NewPeriod

`func NewPeriod(duration Duration, repeat int32, ) *Period`

NewPeriod instantiates a new Period object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeriodWithDefaults

`func NewPeriodWithDefaults() *Period`

NewPeriodWithDefaults instantiates a new Period object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *Period) GetDuration() Duration`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *Period) GetDurationOk() (*Duration, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *Period) SetDuration(v Duration)`

SetDuration sets Duration field to given value.


### GetRepeat

`func (o *Period) GetRepeat() int32`

GetRepeat returns the Repeat field if non-nil, zero value otherwise.

### GetRepeatOk

`func (o *Period) GetRepeatOk() (*int32, bool)`

GetRepeatOk returns a tuple with the Repeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeat

`func (o *Period) SetRepeat(v int32)`

SetRepeat sets Repeat field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


