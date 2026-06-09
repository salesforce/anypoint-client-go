# DeploymentGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the deployment group. | [optional] 
**Name** | Pointer to **string** | The name of the deployment group. | [optional] 

## Methods

### NewDeploymentGroup

`func NewDeploymentGroup() *DeploymentGroup`

NewDeploymentGroup instantiates a new DeploymentGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentGroupWithDefaults

`func NewDeploymentGroupWithDefaults() *DeploymentGroup`

NewDeploymentGroupWithDefaults instantiates a new DeploymentGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeploymentGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeploymentGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DeploymentGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeploymentGroup) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


