# TrackingSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrackingLevel** | Pointer to **string** | The level of tracking applied (e.g., DISABLED, BASIC, FULL). | [optional] 

## Methods

### NewTrackingSettings

`func NewTrackingSettings() *TrackingSettings`

NewTrackingSettings instantiates a new TrackingSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackingSettingsWithDefaults

`func NewTrackingSettingsWithDefaults() *TrackingSettings`

NewTrackingSettingsWithDefaults instantiates a new TrackingSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrackingLevel

`func (o *TrackingSettings) GetTrackingLevel() string`

GetTrackingLevel returns the TrackingLevel field if non-nil, zero value otherwise.

### GetTrackingLevelOk

`func (o *TrackingSettings) GetTrackingLevelOk() (*string, bool)`

GetTrackingLevelOk returns a tuple with the TrackingLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingLevel

`func (o *TrackingSettings) SetTrackingLevel(v string)`

SetTrackingLevel sets TrackingLevel field to given value.

### HasTrackingLevel

`func (o *TrackingSettings) HasTrackingLevel() bool`

HasTrackingLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


