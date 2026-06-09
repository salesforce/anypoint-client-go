# DeploymentArtifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastUpdateTime** | Pointer to **NullableInt64** | The timestamp (in milliseconds) when the artifact was last updated. | [optional] 
**CreateTime** | Pointer to **NullableInt64** | The creation timestamp (in milliseconds); may be null. | [optional] 
**Name** | Pointer to **string** | The name of the artifact. | [optional] 
**FileName** | Pointer to **string** | The file name of the artifact. | [optional] 

## Methods

### NewDeploymentArtifact

`func NewDeploymentArtifact() *DeploymentArtifact`

NewDeploymentArtifact instantiates a new DeploymentArtifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentArtifactWithDefaults

`func NewDeploymentArtifactWithDefaults() *DeploymentArtifact`

NewDeploymentArtifactWithDefaults instantiates a new DeploymentArtifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastUpdateTime

`func (o *DeploymentArtifact) GetLastUpdateTime() int64`

GetLastUpdateTime returns the LastUpdateTime field if non-nil, zero value otherwise.

### GetLastUpdateTimeOk

`func (o *DeploymentArtifact) GetLastUpdateTimeOk() (*int64, bool)`

GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTime

`func (o *DeploymentArtifact) SetLastUpdateTime(v int64)`

SetLastUpdateTime sets LastUpdateTime field to given value.

### HasLastUpdateTime

`func (o *DeploymentArtifact) HasLastUpdateTime() bool`

HasLastUpdateTime returns a boolean if a field has been set.

### SetLastUpdateTimeNil

`func (o *DeploymentArtifact) SetLastUpdateTimeNil(b bool)`

 SetLastUpdateTimeNil sets the value for LastUpdateTime to be an explicit nil

### UnsetLastUpdateTime
`func (o *DeploymentArtifact) UnsetLastUpdateTime()`

UnsetLastUpdateTime ensures that no value is present for LastUpdateTime, not even an explicit nil
### GetCreateTime

`func (o *DeploymentArtifact) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *DeploymentArtifact) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *DeploymentArtifact) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *DeploymentArtifact) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### SetCreateTimeNil

`func (o *DeploymentArtifact) SetCreateTimeNil(b bool)`

 SetCreateTimeNil sets the value for CreateTime to be an explicit nil

### UnsetCreateTime
`func (o *DeploymentArtifact) UnsetCreateTime()`

UnsetCreateTime ensures that no value is present for CreateTime, not even an explicit nil
### GetName

`func (o *DeploymentArtifact) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentArtifact) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentArtifact) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeploymentArtifact) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFileName

`func (o *DeploymentArtifact) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *DeploymentArtifact) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *DeploymentArtifact) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *DeploymentArtifact) HasFileName() bool`

HasFileName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


