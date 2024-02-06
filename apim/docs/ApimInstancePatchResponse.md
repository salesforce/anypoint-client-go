# ApimInstancePatchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audit** | Pointer to [**Audit**](Audit.md) |  | [optional] 
**MasterOrganizationId** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**InstanceLabel** | Pointer to **string** |  | [optional] 
**ProviderId** | Pointer to **NullableString** |  | [optional] 
**EndpointUri** | Pointer to **string** |  | [optional] 
**EnvironmentId** | Pointer to **string** |  | [optional] 
**Technology** | Pointer to **string** |  | [optional] 
**Routing** | Pointer to [**[]Routing**](Routing.md) |  | [optional] 
**Endpoint** | Pointer to [**Endpoint**](Endpoint.md) |  | [optional] 
**Deployment** | Pointer to [**NullableDeployment**](Deployment.md) |  | [optional] 

## Methods

### NewApimInstancePatchResponse

`func NewApimInstancePatchResponse() *ApimInstancePatchResponse`

NewApimInstancePatchResponse instantiates a new ApimInstancePatchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApimInstancePatchResponseWithDefaults

`func NewApimInstancePatchResponseWithDefaults() *ApimInstancePatchResponse`

NewApimInstancePatchResponseWithDefaults instantiates a new ApimInstancePatchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudit

`func (o *ApimInstancePatchResponse) GetAudit() Audit`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *ApimInstancePatchResponse) GetAuditOk() (*Audit, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *ApimInstancePatchResponse) SetAudit(v Audit)`

SetAudit sets Audit field to given value.

### HasAudit

`func (o *ApimInstancePatchResponse) HasAudit() bool`

HasAudit returns a boolean if a field has been set.

### GetMasterOrganizationId

`func (o *ApimInstancePatchResponse) GetMasterOrganizationId() string`

GetMasterOrganizationId returns the MasterOrganizationId field if non-nil, zero value otherwise.

### GetMasterOrganizationIdOk

`func (o *ApimInstancePatchResponse) GetMasterOrganizationIdOk() (*string, bool)`

GetMasterOrganizationIdOk returns a tuple with the MasterOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterOrganizationId

`func (o *ApimInstancePatchResponse) SetMasterOrganizationId(v string)`

SetMasterOrganizationId sets MasterOrganizationId field to given value.

### HasMasterOrganizationId

`func (o *ApimInstancePatchResponse) HasMasterOrganizationId() bool`

HasMasterOrganizationId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *ApimInstancePatchResponse) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ApimInstancePatchResponse) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ApimInstancePatchResponse) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ApimInstancePatchResponse) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetId

`func (o *ApimInstancePatchResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApimInstancePatchResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApimInstancePatchResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ApimInstancePatchResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstanceLabel

`func (o *ApimInstancePatchResponse) GetInstanceLabel() string`

GetInstanceLabel returns the InstanceLabel field if non-nil, zero value otherwise.

### GetInstanceLabelOk

`func (o *ApimInstancePatchResponse) GetInstanceLabelOk() (*string, bool)`

GetInstanceLabelOk returns a tuple with the InstanceLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceLabel

`func (o *ApimInstancePatchResponse) SetInstanceLabel(v string)`

SetInstanceLabel sets InstanceLabel field to given value.

### HasInstanceLabel

`func (o *ApimInstancePatchResponse) HasInstanceLabel() bool`

HasInstanceLabel returns a boolean if a field has been set.

### GetProviderId

`func (o *ApimInstancePatchResponse) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *ApimInstancePatchResponse) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *ApimInstancePatchResponse) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *ApimInstancePatchResponse) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### SetProviderIdNil

`func (o *ApimInstancePatchResponse) SetProviderIdNil(b bool)`

 SetProviderIdNil sets the value for ProviderId to be an explicit nil

### UnsetProviderId
`func (o *ApimInstancePatchResponse) UnsetProviderId()`

UnsetProviderId ensures that no value is present for ProviderId, not even an explicit nil
### GetEndpointUri

`func (o *ApimInstancePatchResponse) GetEndpointUri() string`

GetEndpointUri returns the EndpointUri field if non-nil, zero value otherwise.

### GetEndpointUriOk

`func (o *ApimInstancePatchResponse) GetEndpointUriOk() (*string, bool)`

GetEndpointUriOk returns a tuple with the EndpointUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointUri

`func (o *ApimInstancePatchResponse) SetEndpointUri(v string)`

SetEndpointUri sets EndpointUri field to given value.

### HasEndpointUri

`func (o *ApimInstancePatchResponse) HasEndpointUri() bool`

HasEndpointUri returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *ApimInstancePatchResponse) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *ApimInstancePatchResponse) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *ApimInstancePatchResponse) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *ApimInstancePatchResponse) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetTechnology

`func (o *ApimInstancePatchResponse) GetTechnology() string`

GetTechnology returns the Technology field if non-nil, zero value otherwise.

### GetTechnologyOk

`func (o *ApimInstancePatchResponse) GetTechnologyOk() (*string, bool)`

GetTechnologyOk returns a tuple with the Technology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnology

`func (o *ApimInstancePatchResponse) SetTechnology(v string)`

SetTechnology sets Technology field to given value.

### HasTechnology

`func (o *ApimInstancePatchResponse) HasTechnology() bool`

HasTechnology returns a boolean if a field has been set.

### GetRouting

`func (o *ApimInstancePatchResponse) GetRouting() []Routing`

GetRouting returns the Routing field if non-nil, zero value otherwise.

### GetRoutingOk

`func (o *ApimInstancePatchResponse) GetRoutingOk() (*[]Routing, bool)`

GetRoutingOk returns a tuple with the Routing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouting

`func (o *ApimInstancePatchResponse) SetRouting(v []Routing)`

SetRouting sets Routing field to given value.

### HasRouting

`func (o *ApimInstancePatchResponse) HasRouting() bool`

HasRouting returns a boolean if a field has been set.

### GetEndpoint

`func (o *ApimInstancePatchResponse) GetEndpoint() Endpoint`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ApimInstancePatchResponse) GetEndpointOk() (*Endpoint, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ApimInstancePatchResponse) SetEndpoint(v Endpoint)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *ApimInstancePatchResponse) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetDeployment

`func (o *ApimInstancePatchResponse) GetDeployment() Deployment`

GetDeployment returns the Deployment field if non-nil, zero value otherwise.

### GetDeploymentOk

`func (o *ApimInstancePatchResponse) GetDeploymentOk() (*Deployment, bool)`

GetDeploymentOk returns a tuple with the Deployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployment

`func (o *ApimInstancePatchResponse) SetDeployment(v Deployment)`

SetDeployment sets Deployment field to given value.

### HasDeployment

`func (o *ApimInstancePatchResponse) HasDeployment() bool`

HasDeployment returns a boolean if a field has been set.

### SetDeploymentNil

`func (o *ApimInstancePatchResponse) SetDeploymentNil(b bool)`

 SetDeploymentNil sets the value for Deployment to be an explicit nil

### UnsetDeployment
`func (o *ApimInstancePatchResponse) UnsetDeployment()`

UnsetDeployment ensures that no value is present for Deployment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


