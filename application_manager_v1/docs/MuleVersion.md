# MuleVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | The Mule runtime version. | [optional] 
**UpdateId** | Pointer to **string** | Identifier for the current update. | [optional] 
**LatestUpdateId** | Pointer to **string** | Identifier for the latest available update. | [optional] 
**EndOfSupportDate** | Pointer to **int64** | Timestamp (in milliseconds) representing the end-of-support date. | [optional] 
**EndOfLifeDate** | Pointer to **int64** | Timestamp (in milliseconds) representing the end-of-life date. | [optional] 
**UpdateVersion** | Pointer to **string** | The update version string. | [optional] 
**ReleaseChannel** | Pointer to **string** | The release channel for the runtime. | [optional] 
**JavaVersion** | Pointer to **string** | The version of Java used by the runtime. | [optional] 

## Methods

### NewMuleVersion

`func NewMuleVersion() *MuleVersion`

NewMuleVersion instantiates a new MuleVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMuleVersionWithDefaults

`func NewMuleVersionWithDefaults() *MuleVersion`

NewMuleVersionWithDefaults instantiates a new MuleVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *MuleVersion) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MuleVersion) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MuleVersion) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MuleVersion) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetUpdateId

`func (o *MuleVersion) GetUpdateId() string`

GetUpdateId returns the UpdateId field if non-nil, zero value otherwise.

### GetUpdateIdOk

`func (o *MuleVersion) GetUpdateIdOk() (*string, bool)`

GetUpdateIdOk returns a tuple with the UpdateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateId

`func (o *MuleVersion) SetUpdateId(v string)`

SetUpdateId sets UpdateId field to given value.

### HasUpdateId

`func (o *MuleVersion) HasUpdateId() bool`

HasUpdateId returns a boolean if a field has been set.

### GetLatestUpdateId

`func (o *MuleVersion) GetLatestUpdateId() string`

GetLatestUpdateId returns the LatestUpdateId field if non-nil, zero value otherwise.

### GetLatestUpdateIdOk

`func (o *MuleVersion) GetLatestUpdateIdOk() (*string, bool)`

GetLatestUpdateIdOk returns a tuple with the LatestUpdateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestUpdateId

`func (o *MuleVersion) SetLatestUpdateId(v string)`

SetLatestUpdateId sets LatestUpdateId field to given value.

### HasLatestUpdateId

`func (o *MuleVersion) HasLatestUpdateId() bool`

HasLatestUpdateId returns a boolean if a field has been set.

### GetEndOfSupportDate

`func (o *MuleVersion) GetEndOfSupportDate() int64`

GetEndOfSupportDate returns the EndOfSupportDate field if non-nil, zero value otherwise.

### GetEndOfSupportDateOk

`func (o *MuleVersion) GetEndOfSupportDateOk() (*int64, bool)`

GetEndOfSupportDateOk returns a tuple with the EndOfSupportDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOfSupportDate

`func (o *MuleVersion) SetEndOfSupportDate(v int64)`

SetEndOfSupportDate sets EndOfSupportDate field to given value.

### HasEndOfSupportDate

`func (o *MuleVersion) HasEndOfSupportDate() bool`

HasEndOfSupportDate returns a boolean if a field has been set.

### GetEndOfLifeDate

`func (o *MuleVersion) GetEndOfLifeDate() int64`

GetEndOfLifeDate returns the EndOfLifeDate field if non-nil, zero value otherwise.

### GetEndOfLifeDateOk

`func (o *MuleVersion) GetEndOfLifeDateOk() (*int64, bool)`

GetEndOfLifeDateOk returns a tuple with the EndOfLifeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOfLifeDate

`func (o *MuleVersion) SetEndOfLifeDate(v int64)`

SetEndOfLifeDate sets EndOfLifeDate field to given value.

### HasEndOfLifeDate

`func (o *MuleVersion) HasEndOfLifeDate() bool`

HasEndOfLifeDate returns a boolean if a field has been set.

### GetUpdateVersion

`func (o *MuleVersion) GetUpdateVersion() string`

GetUpdateVersion returns the UpdateVersion field if non-nil, zero value otherwise.

### GetUpdateVersionOk

`func (o *MuleVersion) GetUpdateVersionOk() (*string, bool)`

GetUpdateVersionOk returns a tuple with the UpdateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateVersion

`func (o *MuleVersion) SetUpdateVersion(v string)`

SetUpdateVersion sets UpdateVersion field to given value.

### HasUpdateVersion

`func (o *MuleVersion) HasUpdateVersion() bool`

HasUpdateVersion returns a boolean if a field has been set.

### GetReleaseChannel

`func (o *MuleVersion) GetReleaseChannel() string`

GetReleaseChannel returns the ReleaseChannel field if non-nil, zero value otherwise.

### GetReleaseChannelOk

`func (o *MuleVersion) GetReleaseChannelOk() (*string, bool)`

GetReleaseChannelOk returns a tuple with the ReleaseChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseChannel

`func (o *MuleVersion) SetReleaseChannel(v string)`

SetReleaseChannel sets ReleaseChannel field to given value.

### HasReleaseChannel

`func (o *MuleVersion) HasReleaseChannel() bool`

HasReleaseChannel returns a boolean if a field has been set.

### GetJavaVersion

`func (o *MuleVersion) GetJavaVersion() string`

GetJavaVersion returns the JavaVersion field if non-nil, zero value otherwise.

### GetJavaVersionOk

`func (o *MuleVersion) GetJavaVersionOk() (*string, bool)`

GetJavaVersionOk returns a tuple with the JavaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavaVersion

`func (o *MuleVersion) SetJavaVersion(v string)`

SetJavaVersion sets JavaVersion field to given value.

### HasJavaVersion

`func (o *MuleVersion) HasJavaVersion() bool`

HasJavaVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


