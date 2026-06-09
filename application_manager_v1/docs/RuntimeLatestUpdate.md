# RuntimeLatestUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier of the update. | [optional] 
**Name** | Pointer to **string** | The name of the update. | [optional] 
**ReleaseDate** | Pointer to **float32** | The release date as a timestamp (in milliseconds). | [optional] 
**ReleaseNotes** | Pointer to **string** | URL or text of the release notes. | [optional] 
**Flags** | Pointer to [**RuntimeLatestUpdateFlags**](RuntimeLatestUpdateFlags.md) |  | [optional] 

## Methods

### NewRuntimeLatestUpdate

`func NewRuntimeLatestUpdate() *RuntimeLatestUpdate`

NewRuntimeLatestUpdate instantiates a new RuntimeLatestUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuntimeLatestUpdateWithDefaults

`func NewRuntimeLatestUpdateWithDefaults() *RuntimeLatestUpdate`

NewRuntimeLatestUpdateWithDefaults instantiates a new RuntimeLatestUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RuntimeLatestUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RuntimeLatestUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RuntimeLatestUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RuntimeLatestUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RuntimeLatestUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RuntimeLatestUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RuntimeLatestUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RuntimeLatestUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReleaseDate

`func (o *RuntimeLatestUpdate) GetReleaseDate() float32`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *RuntimeLatestUpdate) GetReleaseDateOk() (*float32, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *RuntimeLatestUpdate) SetReleaseDate(v float32)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *RuntimeLatestUpdate) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetReleaseNotes

`func (o *RuntimeLatestUpdate) GetReleaseNotes() string`

GetReleaseNotes returns the ReleaseNotes field if non-nil, zero value otherwise.

### GetReleaseNotesOk

`func (o *RuntimeLatestUpdate) GetReleaseNotesOk() (*string, bool)`

GetReleaseNotesOk returns a tuple with the ReleaseNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseNotes

`func (o *RuntimeLatestUpdate) SetReleaseNotes(v string)`

SetReleaseNotes sets ReleaseNotes field to given value.

### HasReleaseNotes

`func (o *RuntimeLatestUpdate) HasReleaseNotes() bool`

HasReleaseNotes returns a boolean if a field has been set.

### GetFlags

`func (o *RuntimeLatestUpdate) GetFlags() RuntimeLatestUpdateFlags`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *RuntimeLatestUpdate) GetFlagsOk() (*RuntimeLatestUpdateFlags, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *RuntimeLatestUpdate) SetFlags(v RuntimeLatestUpdateFlags)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *RuntimeLatestUpdate) HasFlags() bool`

HasFlags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


