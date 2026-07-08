# DeploymentTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The target type (e.g., CLOUDHUB, MC). | [optional] 
**Subtype** | Pointer to **string** | The target subtype (if applicable). | [optional] 
**Provider** | Pointer to **string** | The provider of the target platform. | [optional] 
**Id** | Pointer to **string** | The identifier of the target. | [optional] 
**Name** | Pointer to **string** | The name of the target. | [optional] 
**Status** | Pointer to **string** | The current status of the target. | [optional] 

## Methods

### NewDeploymentTarget

`func NewDeploymentTarget() *DeploymentTarget`

NewDeploymentTarget instantiates a new DeploymentTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentTargetWithDefaults

`func NewDeploymentTargetWithDefaults() *DeploymentTarget`

NewDeploymentTargetWithDefaults instantiates a new DeploymentTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DeploymentTarget) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeploymentTarget) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeploymentTarget) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DeploymentTarget) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSubtype

`func (o *DeploymentTarget) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *DeploymentTarget) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *DeploymentTarget) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *DeploymentTarget) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetProvider

`func (o *DeploymentTarget) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *DeploymentTarget) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *DeploymentTarget) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *DeploymentTarget) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetId

`func (o *DeploymentTarget) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentTarget) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentTarget) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeploymentTarget) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DeploymentTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentTarget) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeploymentTarget) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *DeploymentTarget) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeploymentTarget) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeploymentTarget) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeploymentTarget) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


