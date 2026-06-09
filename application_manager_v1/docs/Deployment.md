# Deployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A unique identifier for the deployment. | [optional] 
**Artifact** | Pointer to [**DeploymentArtifact**](DeploymentArtifact.md) |  | [optional] 
**LastReportedStatus** | Pointer to **string** | The last reported status of the deployment (e.g., STARTED, UNDEPLOYED). | [optional] 
**Details** | Pointer to [**DeploymentDetails**](DeploymentDetails.md) |  | [optional] 
**Target** | Pointer to [**DeploymentTarget**](DeploymentTarget.md) |  | [optional] 
**Application** | Pointer to [**DeploymentApplication**](DeploymentApplication.md) |  | [optional] 
**MuleVersion** | Pointer to [**DeploymentMuleVersion**](DeploymentMuleVersion.md) |  | [optional] 

## Methods

### NewDeployment

`func NewDeployment() *Deployment`

NewDeployment instantiates a new Deployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentWithDefaults

`func NewDeploymentWithDefaults() *Deployment`

NewDeploymentWithDefaults instantiates a new Deployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Deployment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Deployment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Deployment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Deployment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetArtifact

`func (o *Deployment) GetArtifact() DeploymentArtifact`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *Deployment) GetArtifactOk() (*DeploymentArtifact, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *Deployment) SetArtifact(v DeploymentArtifact)`

SetArtifact sets Artifact field to given value.

### HasArtifact

`func (o *Deployment) HasArtifact() bool`

HasArtifact returns a boolean if a field has been set.

### GetLastReportedStatus

`func (o *Deployment) GetLastReportedStatus() string`

GetLastReportedStatus returns the LastReportedStatus field if non-nil, zero value otherwise.

### GetLastReportedStatusOk

`func (o *Deployment) GetLastReportedStatusOk() (*string, bool)`

GetLastReportedStatusOk returns a tuple with the LastReportedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReportedStatus

`func (o *Deployment) SetLastReportedStatus(v string)`

SetLastReportedStatus sets LastReportedStatus field to given value.

### HasLastReportedStatus

`func (o *Deployment) HasLastReportedStatus() bool`

HasLastReportedStatus returns a boolean if a field has been set.

### GetDetails

`func (o *Deployment) GetDetails() DeploymentDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Deployment) GetDetailsOk() (*DeploymentDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Deployment) SetDetails(v DeploymentDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *Deployment) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetTarget

`func (o *Deployment) GetTarget() DeploymentTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *Deployment) GetTargetOk() (*DeploymentTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *Deployment) SetTarget(v DeploymentTarget)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *Deployment) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetApplication

`func (o *Deployment) GetApplication() DeploymentApplication`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *Deployment) GetApplicationOk() (*DeploymentApplication, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *Deployment) SetApplication(v DeploymentApplication)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *Deployment) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetMuleVersion

`func (o *Deployment) GetMuleVersion() DeploymentMuleVersion`

GetMuleVersion returns the MuleVersion field if non-nil, zero value otherwise.

### GetMuleVersionOk

`func (o *Deployment) GetMuleVersionOk() (*DeploymentMuleVersion, bool)`

GetMuleVersionOk returns a tuple with the MuleVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuleVersion

`func (o *Deployment) SetMuleVersion(v DeploymentMuleVersion)`

SetMuleVersion sets MuleVersion field to given value.

### HasMuleVersion

`func (o *Deployment) HasMuleVersion() bool`

HasMuleVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


