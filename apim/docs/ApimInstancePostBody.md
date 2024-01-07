# ApimInstancePostBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Technology** | Pointer to **string** |  | [optional] 
**Endpoint** | Pointer to [**NullableEndpointPostBody**](EndpointPostBody.md) |  | [optional] 
**Spec** | Pointer to [**Spec**](Spec.md) |  | [optional] 
**Routing** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Deployment** | Pointer to [**NullableDeploymentPostBody**](DeploymentPostBody.md) |  | [optional] 
**InstanceLabel** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewApimInstancePostBody

`func NewApimInstancePostBody() *ApimInstancePostBody`

NewApimInstancePostBody instantiates a new ApimInstancePostBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApimInstancePostBodyWithDefaults

`func NewApimInstancePostBodyWithDefaults() *ApimInstancePostBody`

NewApimInstancePostBodyWithDefaults instantiates a new ApimInstancePostBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTechnology

`func (o *ApimInstancePostBody) GetTechnology() string`

GetTechnology returns the Technology field if non-nil, zero value otherwise.

### GetTechnologyOk

`func (o *ApimInstancePostBody) GetTechnologyOk() (*string, bool)`

GetTechnologyOk returns a tuple with the Technology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnology

`func (o *ApimInstancePostBody) SetTechnology(v string)`

SetTechnology sets Technology field to given value.

### HasTechnology

`func (o *ApimInstancePostBody) HasTechnology() bool`

HasTechnology returns a boolean if a field has been set.

### GetEndpoint

`func (o *ApimInstancePostBody) GetEndpoint() EndpointPostBody`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ApimInstancePostBody) GetEndpointOk() (*EndpointPostBody, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ApimInstancePostBody) SetEndpoint(v EndpointPostBody)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *ApimInstancePostBody) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### SetEndpointNil

`func (o *ApimInstancePostBody) SetEndpointNil(b bool)`

 SetEndpointNil sets the value for Endpoint to be an explicit nil

### UnsetEndpoint
`func (o *ApimInstancePostBody) UnsetEndpoint()`

UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
### GetSpec

`func (o *ApimInstancePostBody) GetSpec() Spec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ApimInstancePostBody) GetSpecOk() (*Spec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ApimInstancePostBody) SetSpec(v Spec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ApimInstancePostBody) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetRouting

`func (o *ApimInstancePostBody) GetRouting() []map[string]interface{}`

GetRouting returns the Routing field if non-nil, zero value otherwise.

### GetRoutingOk

`func (o *ApimInstancePostBody) GetRoutingOk() (*[]map[string]interface{}, bool)`

GetRoutingOk returns a tuple with the Routing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouting

`func (o *ApimInstancePostBody) SetRouting(v []map[string]interface{})`

SetRouting sets Routing field to given value.

### HasRouting

`func (o *ApimInstancePostBody) HasRouting() bool`

HasRouting returns a boolean if a field has been set.

### SetRoutingNil

`func (o *ApimInstancePostBody) SetRoutingNil(b bool)`

 SetRoutingNil sets the value for Routing to be an explicit nil

### UnsetRouting
`func (o *ApimInstancePostBody) UnsetRouting()`

UnsetRouting ensures that no value is present for Routing, not even an explicit nil
### GetDeployment

`func (o *ApimInstancePostBody) GetDeployment() DeploymentPostBody`

GetDeployment returns the Deployment field if non-nil, zero value otherwise.

### GetDeploymentOk

`func (o *ApimInstancePostBody) GetDeploymentOk() (*DeploymentPostBody, bool)`

GetDeploymentOk returns a tuple with the Deployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployment

`func (o *ApimInstancePostBody) SetDeployment(v DeploymentPostBody)`

SetDeployment sets Deployment field to given value.

### HasDeployment

`func (o *ApimInstancePostBody) HasDeployment() bool`

HasDeployment returns a boolean if a field has been set.

### SetDeploymentNil

`func (o *ApimInstancePostBody) SetDeploymentNil(b bool)`

 SetDeploymentNil sets the value for Deployment to be an explicit nil

### UnsetDeployment
`func (o *ApimInstancePostBody) UnsetDeployment()`

UnsetDeployment ensures that no value is present for Deployment, not even an explicit nil
### GetInstanceLabel

`func (o *ApimInstancePostBody) GetInstanceLabel() string`

GetInstanceLabel returns the InstanceLabel field if non-nil, zero value otherwise.

### GetInstanceLabelOk

`func (o *ApimInstancePostBody) GetInstanceLabelOk() (*string, bool)`

GetInstanceLabelOk returns a tuple with the InstanceLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceLabel

`func (o *ApimInstancePostBody) SetInstanceLabel(v string)`

SetInstanceLabel sets InstanceLabel field to given value.

### HasInstanceLabel

`func (o *ApimInstancePostBody) HasInstanceLabel() bool`

HasInstanceLabel returns a boolean if a field has been set.

### SetInstanceLabelNil

`func (o *ApimInstancePostBody) SetInstanceLabelNil(b bool)`

 SetInstanceLabelNil sets the value for InstanceLabel to be an explicit nil

### UnsetInstanceLabel
`func (o *ApimInstancePostBody) UnsetInstanceLabel()`

UnsetInstanceLabel ensures that no value is present for InstanceLabel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


