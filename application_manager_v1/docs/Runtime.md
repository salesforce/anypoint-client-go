# Runtime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | The runtime version identifier. | [optional] 
**DisplayName** | Pointer to **string** | The display name for the runtime. | [optional] 
**Recommended** | Pointer to **bool** | Indicates whether this runtime is recommended. | [optional] 
**EndOfSupportDate** | Pointer to **float32** | The timestamp (in milliseconds) when support ends for this runtime. | [optional] 
**EndOfLifeDate** | Pointer to **float32** | The timestamp (in milliseconds) when this runtime reaches end of life. | [optional] 
**LatestUpdate** | Pointer to [**RuntimeLatestUpdate**](RuntimeLatestUpdate.md) |  | [optional] 
**State** | Pointer to **string** | The state of the runtime (for example, ACTIVE). | [optional] 
**ReleaseChannel** | Pointer to **string** | The release channel of the runtime (for example, EDGE, LTS). | [optional] 
**JavaVersion** | Pointer to **string** | The Java version used by the runtime. | [optional] 
**Default** | Pointer to **bool** | Indicates whether this runtime is the default. | [optional] 

## Methods

### NewRuntime

`func NewRuntime() *Runtime`

NewRuntime instantiates a new Runtime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuntimeWithDefaults

`func NewRuntimeWithDefaults() *Runtime`

NewRuntimeWithDefaults instantiates a new Runtime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *Runtime) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Runtime) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Runtime) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Runtime) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDisplayName

`func (o *Runtime) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Runtime) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Runtime) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Runtime) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetRecommended

`func (o *Runtime) GetRecommended() bool`

GetRecommended returns the Recommended field if non-nil, zero value otherwise.

### GetRecommendedOk

`func (o *Runtime) GetRecommendedOk() (*bool, bool)`

GetRecommendedOk returns a tuple with the Recommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommended

`func (o *Runtime) SetRecommended(v bool)`

SetRecommended sets Recommended field to given value.

### HasRecommended

`func (o *Runtime) HasRecommended() bool`

HasRecommended returns a boolean if a field has been set.

### GetEndOfSupportDate

`func (o *Runtime) GetEndOfSupportDate() float32`

GetEndOfSupportDate returns the EndOfSupportDate field if non-nil, zero value otherwise.

### GetEndOfSupportDateOk

`func (o *Runtime) GetEndOfSupportDateOk() (*float32, bool)`

GetEndOfSupportDateOk returns a tuple with the EndOfSupportDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOfSupportDate

`func (o *Runtime) SetEndOfSupportDate(v float32)`

SetEndOfSupportDate sets EndOfSupportDate field to given value.

### HasEndOfSupportDate

`func (o *Runtime) HasEndOfSupportDate() bool`

HasEndOfSupportDate returns a boolean if a field has been set.

### GetEndOfLifeDate

`func (o *Runtime) GetEndOfLifeDate() float32`

GetEndOfLifeDate returns the EndOfLifeDate field if non-nil, zero value otherwise.

### GetEndOfLifeDateOk

`func (o *Runtime) GetEndOfLifeDateOk() (*float32, bool)`

GetEndOfLifeDateOk returns a tuple with the EndOfLifeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOfLifeDate

`func (o *Runtime) SetEndOfLifeDate(v float32)`

SetEndOfLifeDate sets EndOfLifeDate field to given value.

### HasEndOfLifeDate

`func (o *Runtime) HasEndOfLifeDate() bool`

HasEndOfLifeDate returns a boolean if a field has been set.

### GetLatestUpdate

`func (o *Runtime) GetLatestUpdate() RuntimeLatestUpdate`

GetLatestUpdate returns the LatestUpdate field if non-nil, zero value otherwise.

### GetLatestUpdateOk

`func (o *Runtime) GetLatestUpdateOk() (*RuntimeLatestUpdate, bool)`

GetLatestUpdateOk returns a tuple with the LatestUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestUpdate

`func (o *Runtime) SetLatestUpdate(v RuntimeLatestUpdate)`

SetLatestUpdate sets LatestUpdate field to given value.

### HasLatestUpdate

`func (o *Runtime) HasLatestUpdate() bool`

HasLatestUpdate returns a boolean if a field has been set.

### GetState

`func (o *Runtime) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Runtime) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Runtime) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Runtime) HasState() bool`

HasState returns a boolean if a field has been set.

### GetReleaseChannel

`func (o *Runtime) GetReleaseChannel() string`

GetReleaseChannel returns the ReleaseChannel field if non-nil, zero value otherwise.

### GetReleaseChannelOk

`func (o *Runtime) GetReleaseChannelOk() (*string, bool)`

GetReleaseChannelOk returns a tuple with the ReleaseChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseChannel

`func (o *Runtime) SetReleaseChannel(v string)`

SetReleaseChannel sets ReleaseChannel field to given value.

### HasReleaseChannel

`func (o *Runtime) HasReleaseChannel() bool`

HasReleaseChannel returns a boolean if a field has been set.

### GetJavaVersion

`func (o *Runtime) GetJavaVersion() string`

GetJavaVersion returns the JavaVersion field if non-nil, zero value otherwise.

### GetJavaVersionOk

`func (o *Runtime) GetJavaVersionOk() (*string, bool)`

GetJavaVersionOk returns a tuple with the JavaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavaVersion

`func (o *Runtime) SetJavaVersion(v string)`

SetJavaVersion sets JavaVersion field to given value.

### HasJavaVersion

`func (o *Runtime) HasJavaVersion() bool`

HasJavaVersion returns a boolean if a field has been set.

### GetDefault

`func (o *Runtime) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *Runtime) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *Runtime) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *Runtime) HasDefault() bool`

HasDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


