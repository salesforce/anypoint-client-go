# DeploymentMuleVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | The Mule runtime version. | [optional] 
**UpdateId** | Pointer to **string** | The identifier for the runtime update. | [optional] 
**LatestUpdateId** | Pointer to **string** | The identifier for the latest runtime update. | [optional] 
**EndOfSupportDate** | Pointer to **int64** | The timestamp (in milliseconds) representing the end-of-support date. | [optional] 

## Methods

### NewDeploymentMuleVersion

`func NewDeploymentMuleVersion() *DeploymentMuleVersion`

NewDeploymentMuleVersion instantiates a new DeploymentMuleVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentMuleVersionWithDefaults

`func NewDeploymentMuleVersionWithDefaults() *DeploymentMuleVersion`

NewDeploymentMuleVersionWithDefaults instantiates a new DeploymentMuleVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *DeploymentMuleVersion) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DeploymentMuleVersion) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DeploymentMuleVersion) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DeploymentMuleVersion) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetUpdateId

`func (o *DeploymentMuleVersion) GetUpdateId() string`

GetUpdateId returns the UpdateId field if non-nil, zero value otherwise.

### GetUpdateIdOk

`func (o *DeploymentMuleVersion) GetUpdateIdOk() (*string, bool)`

GetUpdateIdOk returns a tuple with the UpdateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateId

`func (o *DeploymentMuleVersion) SetUpdateId(v string)`

SetUpdateId sets UpdateId field to given value.

### HasUpdateId

`func (o *DeploymentMuleVersion) HasUpdateId() bool`

HasUpdateId returns a boolean if a field has been set.

### GetLatestUpdateId

`func (o *DeploymentMuleVersion) GetLatestUpdateId() string`

GetLatestUpdateId returns the LatestUpdateId field if non-nil, zero value otherwise.

### GetLatestUpdateIdOk

`func (o *DeploymentMuleVersion) GetLatestUpdateIdOk() (*string, bool)`

GetLatestUpdateIdOk returns a tuple with the LatestUpdateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestUpdateId

`func (o *DeploymentMuleVersion) SetLatestUpdateId(v string)`

SetLatestUpdateId sets LatestUpdateId field to given value.

### HasLatestUpdateId

`func (o *DeploymentMuleVersion) HasLatestUpdateId() bool`

HasLatestUpdateId returns a boolean if a field has been set.

### GetEndOfSupportDate

`func (o *DeploymentMuleVersion) GetEndOfSupportDate() int64`

GetEndOfSupportDate returns the EndOfSupportDate field if non-nil, zero value otherwise.

### GetEndOfSupportDateOk

`func (o *DeploymentMuleVersion) GetEndOfSupportDateOk() (*int64, bool)`

GetEndOfSupportDateOk returns a tuple with the EndOfSupportDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOfSupportDate

`func (o *DeploymentMuleVersion) SetEndOfSupportDate(v int64)`

SetEndOfSupportDate sets EndOfSupportDate field to given value.

### HasEndOfSupportDate

`func (o *DeploymentMuleVersion) HasEndOfSupportDate() bool`

HasEndOfSupportDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


