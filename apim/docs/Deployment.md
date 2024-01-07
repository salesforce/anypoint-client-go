# Deployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audit** | Pointer to [**Audit**](Audit.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**ApplicationId** | Pointer to **string** |  | [optional] 
**ApplicationName** | Pointer to **NullableString** |  | [optional] 
**GatewayVersion** | Pointer to **NullableString** |  | [optional] 
**EnvironmentName** | Pointer to **NullableString** |  | [optional] 
**EnvironmentId** | Pointer to **NullableString** |  | [optional] 
**TargetId** | Pointer to **string** |  | [optional] 
**TargetName** | Pointer to **string** |  | [optional] 
**DeploymentId** | Pointer to **NullableString** |  | [optional] 
**UpdatedDate** | Pointer to **time.Time** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**ExpectedStatus** | Pointer to **string** |  | [optional] 
**ApiId** | Pointer to **int32** |  | [optional] 

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

### GetAudit

`func (o *Deployment) GetAudit() Audit`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *Deployment) GetAuditOk() (*Audit, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *Deployment) SetAudit(v Audit)`

SetAudit sets Audit field to given value.

### HasAudit

`func (o *Deployment) HasAudit() bool`

HasAudit returns a boolean if a field has been set.

### GetId

`func (o *Deployment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Deployment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Deployment) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Deployment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetApplicationId

`func (o *Deployment) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *Deployment) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *Deployment) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *Deployment) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetApplicationName

`func (o *Deployment) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *Deployment) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *Deployment) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *Deployment) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### SetApplicationNameNil

`func (o *Deployment) SetApplicationNameNil(b bool)`

 SetApplicationNameNil sets the value for ApplicationName to be an explicit nil

### UnsetApplicationName
`func (o *Deployment) UnsetApplicationName()`

UnsetApplicationName ensures that no value is present for ApplicationName, not even an explicit nil
### GetGatewayVersion

`func (o *Deployment) GetGatewayVersion() string`

GetGatewayVersion returns the GatewayVersion field if non-nil, zero value otherwise.

### GetGatewayVersionOk

`func (o *Deployment) GetGatewayVersionOk() (*string, bool)`

GetGatewayVersionOk returns a tuple with the GatewayVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayVersion

`func (o *Deployment) SetGatewayVersion(v string)`

SetGatewayVersion sets GatewayVersion field to given value.

### HasGatewayVersion

`func (o *Deployment) HasGatewayVersion() bool`

HasGatewayVersion returns a boolean if a field has been set.

### SetGatewayVersionNil

`func (o *Deployment) SetGatewayVersionNil(b bool)`

 SetGatewayVersionNil sets the value for GatewayVersion to be an explicit nil

### UnsetGatewayVersion
`func (o *Deployment) UnsetGatewayVersion()`

UnsetGatewayVersion ensures that no value is present for GatewayVersion, not even an explicit nil
### GetEnvironmentName

`func (o *Deployment) GetEnvironmentName() string`

GetEnvironmentName returns the EnvironmentName field if non-nil, zero value otherwise.

### GetEnvironmentNameOk

`func (o *Deployment) GetEnvironmentNameOk() (*string, bool)`

GetEnvironmentNameOk returns a tuple with the EnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentName

`func (o *Deployment) SetEnvironmentName(v string)`

SetEnvironmentName sets EnvironmentName field to given value.

### HasEnvironmentName

`func (o *Deployment) HasEnvironmentName() bool`

HasEnvironmentName returns a boolean if a field has been set.

### SetEnvironmentNameNil

`func (o *Deployment) SetEnvironmentNameNil(b bool)`

 SetEnvironmentNameNil sets the value for EnvironmentName to be an explicit nil

### UnsetEnvironmentName
`func (o *Deployment) UnsetEnvironmentName()`

UnsetEnvironmentName ensures that no value is present for EnvironmentName, not even an explicit nil
### GetEnvironmentId

`func (o *Deployment) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *Deployment) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *Deployment) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *Deployment) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### SetEnvironmentIdNil

`func (o *Deployment) SetEnvironmentIdNil(b bool)`

 SetEnvironmentIdNil sets the value for EnvironmentId to be an explicit nil

### UnsetEnvironmentId
`func (o *Deployment) UnsetEnvironmentId()`

UnsetEnvironmentId ensures that no value is present for EnvironmentId, not even an explicit nil
### GetTargetId

`func (o *Deployment) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *Deployment) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *Deployment) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *Deployment) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetTargetName

`func (o *Deployment) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *Deployment) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *Deployment) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *Deployment) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetDeploymentId

`func (o *Deployment) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *Deployment) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *Deployment) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *Deployment) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### SetDeploymentIdNil

`func (o *Deployment) SetDeploymentIdNil(b bool)`

 SetDeploymentIdNil sets the value for DeploymentId to be an explicit nil

### UnsetDeploymentId
`func (o *Deployment) UnsetDeploymentId()`

UnsetDeploymentId ensures that no value is present for DeploymentId, not even an explicit nil
### GetUpdatedDate

`func (o *Deployment) GetUpdatedDate() time.Time`

GetUpdatedDate returns the UpdatedDate field if non-nil, zero value otherwise.

### GetUpdatedDateOk

`func (o *Deployment) GetUpdatedDateOk() (*time.Time, bool)`

GetUpdatedDateOk returns a tuple with the UpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDate

`func (o *Deployment) SetUpdatedDate(v time.Time)`

SetUpdatedDate sets UpdatedDate field to given value.

### HasUpdatedDate

`func (o *Deployment) HasUpdatedDate() bool`

HasUpdatedDate returns a boolean if a field has been set.

### GetType

`func (o *Deployment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Deployment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Deployment) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Deployment) HasType() bool`

HasType returns a boolean if a field has been set.

### GetExpectedStatus

`func (o *Deployment) GetExpectedStatus() string`

GetExpectedStatus returns the ExpectedStatus field if non-nil, zero value otherwise.

### GetExpectedStatusOk

`func (o *Deployment) GetExpectedStatusOk() (*string, bool)`

GetExpectedStatusOk returns a tuple with the ExpectedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedStatus

`func (o *Deployment) SetExpectedStatus(v string)`

SetExpectedStatus sets ExpectedStatus field to given value.

### HasExpectedStatus

`func (o *Deployment) HasExpectedStatus() bool`

HasExpectedStatus returns a boolean if a field has been set.

### GetApiId

`func (o *Deployment) GetApiId() int32`

GetApiId returns the ApiId field if non-nil, zero value otherwise.

### GetApiIdOk

`func (o *Deployment) GetApiIdOk() (*int32, bool)`

GetApiIdOk returns a tuple with the ApiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiId

`func (o *Deployment) SetApiId(v int32)`

SetApiId sets ApiId field to given value.

### HasApiId

`func (o *Deployment) HasApiId() bool`

HasApiId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


