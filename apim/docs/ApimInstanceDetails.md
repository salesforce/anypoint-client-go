# ApimInstanceDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audit** | Pointer to [**Audit**](Audit.md) |  | [optional] 
**MasterOrganizationId** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**InstanceLabel** | Pointer to **string** |  | [optional] 
**GroupId** | Pointer to **string** |  | [optional] 
**AssetId** | Pointer to **string** |  | [optional] 
**AssetVersion** | Pointer to **string** |  | [optional] 
**ProductVersion** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**ProviderId** | Pointer to **NullableString** |  | [optional] 
**Deprecated** | Pointer to **bool** |  | [optional] 
**LastActiveDate** | Pointer to **NullableTime** |  | [optional] 
**EndpointUri** | Pointer to **string** |  | [optional] 
**EnvironmentId** | Pointer to **string** |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**Stage** | Pointer to **string** |  | [optional] 
**Technology** | Pointer to **string** |  | [optional] 
**Endpoint** | Pointer to [**Endpoint**](Endpoint.md) |  | [optional] 
**Deployment** | Pointer to [**NullableDeployment**](Deployment.md) |  | [optional] 
**Routing** | Pointer to [**[]Routing**](Routing.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**AutodiscoveryInstanceName** | Pointer to **string** |  | [optional] 

## Methods

### NewApimInstanceDetails

`func NewApimInstanceDetails() *ApimInstanceDetails`

NewApimInstanceDetails instantiates a new ApimInstanceDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApimInstanceDetailsWithDefaults

`func NewApimInstanceDetailsWithDefaults() *ApimInstanceDetails`

NewApimInstanceDetailsWithDefaults instantiates a new ApimInstanceDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudit

`func (o *ApimInstanceDetails) GetAudit() Audit`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *ApimInstanceDetails) GetAuditOk() (*Audit, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *ApimInstanceDetails) SetAudit(v Audit)`

SetAudit sets Audit field to given value.

### HasAudit

`func (o *ApimInstanceDetails) HasAudit() bool`

HasAudit returns a boolean if a field has been set.

### GetMasterOrganizationId

`func (o *ApimInstanceDetails) GetMasterOrganizationId() string`

GetMasterOrganizationId returns the MasterOrganizationId field if non-nil, zero value otherwise.

### GetMasterOrganizationIdOk

`func (o *ApimInstanceDetails) GetMasterOrganizationIdOk() (*string, bool)`

GetMasterOrganizationIdOk returns a tuple with the MasterOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterOrganizationId

`func (o *ApimInstanceDetails) SetMasterOrganizationId(v string)`

SetMasterOrganizationId sets MasterOrganizationId field to given value.

### HasMasterOrganizationId

`func (o *ApimInstanceDetails) HasMasterOrganizationId() bool`

HasMasterOrganizationId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *ApimInstanceDetails) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ApimInstanceDetails) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ApimInstanceDetails) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ApimInstanceDetails) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetId

`func (o *ApimInstanceDetails) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApimInstanceDetails) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApimInstanceDetails) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ApimInstanceDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstanceLabel

`func (o *ApimInstanceDetails) GetInstanceLabel() string`

GetInstanceLabel returns the InstanceLabel field if non-nil, zero value otherwise.

### GetInstanceLabelOk

`func (o *ApimInstanceDetails) GetInstanceLabelOk() (*string, bool)`

GetInstanceLabelOk returns a tuple with the InstanceLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceLabel

`func (o *ApimInstanceDetails) SetInstanceLabel(v string)`

SetInstanceLabel sets InstanceLabel field to given value.

### HasInstanceLabel

`func (o *ApimInstanceDetails) HasInstanceLabel() bool`

HasInstanceLabel returns a boolean if a field has been set.

### GetGroupId

`func (o *ApimInstanceDetails) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ApimInstanceDetails) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ApimInstanceDetails) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ApimInstanceDetails) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetAssetId

`func (o *ApimInstanceDetails) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *ApimInstanceDetails) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *ApimInstanceDetails) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *ApimInstanceDetails) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetAssetVersion

`func (o *ApimInstanceDetails) GetAssetVersion() string`

GetAssetVersion returns the AssetVersion field if non-nil, zero value otherwise.

### GetAssetVersionOk

`func (o *ApimInstanceDetails) GetAssetVersionOk() (*string, bool)`

GetAssetVersionOk returns a tuple with the AssetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetVersion

`func (o *ApimInstanceDetails) SetAssetVersion(v string)`

SetAssetVersion sets AssetVersion field to given value.

### HasAssetVersion

`func (o *ApimInstanceDetails) HasAssetVersion() bool`

HasAssetVersion returns a boolean if a field has been set.

### GetProductVersion

`func (o *ApimInstanceDetails) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *ApimInstanceDetails) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *ApimInstanceDetails) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.

### HasProductVersion

`func (o *ApimInstanceDetails) HasProductVersion() bool`

HasProductVersion returns a boolean if a field has been set.

### GetDescription

`func (o *ApimInstanceDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApimInstanceDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApimInstanceDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApimInstanceDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ApimInstanceDetails) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ApimInstanceDetails) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTags

`func (o *ApimInstanceDetails) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ApimInstanceDetails) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ApimInstanceDetails) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ApimInstanceDetails) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOrder

`func (o *ApimInstanceDetails) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ApimInstanceDetails) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ApimInstanceDetails) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ApimInstanceDetails) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetProviderId

`func (o *ApimInstanceDetails) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *ApimInstanceDetails) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *ApimInstanceDetails) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *ApimInstanceDetails) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### SetProviderIdNil

`func (o *ApimInstanceDetails) SetProviderIdNil(b bool)`

 SetProviderIdNil sets the value for ProviderId to be an explicit nil

### UnsetProviderId
`func (o *ApimInstanceDetails) UnsetProviderId()`

UnsetProviderId ensures that no value is present for ProviderId, not even an explicit nil
### GetDeprecated

`func (o *ApimInstanceDetails) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *ApimInstanceDetails) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *ApimInstanceDetails) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *ApimInstanceDetails) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetLastActiveDate

`func (o *ApimInstanceDetails) GetLastActiveDate() time.Time`

GetLastActiveDate returns the LastActiveDate field if non-nil, zero value otherwise.

### GetLastActiveDateOk

`func (o *ApimInstanceDetails) GetLastActiveDateOk() (*time.Time, bool)`

GetLastActiveDateOk returns a tuple with the LastActiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActiveDate

`func (o *ApimInstanceDetails) SetLastActiveDate(v time.Time)`

SetLastActiveDate sets LastActiveDate field to given value.

### HasLastActiveDate

`func (o *ApimInstanceDetails) HasLastActiveDate() bool`

HasLastActiveDate returns a boolean if a field has been set.

### SetLastActiveDateNil

`func (o *ApimInstanceDetails) SetLastActiveDateNil(b bool)`

 SetLastActiveDateNil sets the value for LastActiveDate to be an explicit nil

### UnsetLastActiveDate
`func (o *ApimInstanceDetails) UnsetLastActiveDate()`

UnsetLastActiveDate ensures that no value is present for LastActiveDate, not even an explicit nil
### GetEndpointUri

`func (o *ApimInstanceDetails) GetEndpointUri() string`

GetEndpointUri returns the EndpointUri field if non-nil, zero value otherwise.

### GetEndpointUriOk

`func (o *ApimInstanceDetails) GetEndpointUriOk() (*string, bool)`

GetEndpointUriOk returns a tuple with the EndpointUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointUri

`func (o *ApimInstanceDetails) SetEndpointUri(v string)`

SetEndpointUri sets EndpointUri field to given value.

### HasEndpointUri

`func (o *ApimInstanceDetails) HasEndpointUri() bool`

HasEndpointUri returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *ApimInstanceDetails) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *ApimInstanceDetails) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *ApimInstanceDetails) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *ApimInstanceDetails) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetIsPublic

`func (o *ApimInstanceDetails) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *ApimInstanceDetails) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *ApimInstanceDetails) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *ApimInstanceDetails) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetStage

`func (o *ApimInstanceDetails) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *ApimInstanceDetails) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *ApimInstanceDetails) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *ApimInstanceDetails) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetTechnology

`func (o *ApimInstanceDetails) GetTechnology() string`

GetTechnology returns the Technology field if non-nil, zero value otherwise.

### GetTechnologyOk

`func (o *ApimInstanceDetails) GetTechnologyOk() (*string, bool)`

GetTechnologyOk returns a tuple with the Technology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnology

`func (o *ApimInstanceDetails) SetTechnology(v string)`

SetTechnology sets Technology field to given value.

### HasTechnology

`func (o *ApimInstanceDetails) HasTechnology() bool`

HasTechnology returns a boolean if a field has been set.

### GetEndpoint

`func (o *ApimInstanceDetails) GetEndpoint() Endpoint`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ApimInstanceDetails) GetEndpointOk() (*Endpoint, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ApimInstanceDetails) SetEndpoint(v Endpoint)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *ApimInstanceDetails) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetDeployment

`func (o *ApimInstanceDetails) GetDeployment() Deployment`

GetDeployment returns the Deployment field if non-nil, zero value otherwise.

### GetDeploymentOk

`func (o *ApimInstanceDetails) GetDeploymentOk() (*Deployment, bool)`

GetDeploymentOk returns a tuple with the Deployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployment

`func (o *ApimInstanceDetails) SetDeployment(v Deployment)`

SetDeployment sets Deployment field to given value.

### HasDeployment

`func (o *ApimInstanceDetails) HasDeployment() bool`

HasDeployment returns a boolean if a field has been set.

### SetDeploymentNil

`func (o *ApimInstanceDetails) SetDeploymentNil(b bool)`

 SetDeploymentNil sets the value for Deployment to be an explicit nil

### UnsetDeployment
`func (o *ApimInstanceDetails) UnsetDeployment()`

UnsetDeployment ensures that no value is present for Deployment, not even an explicit nil
### GetRouting

`func (o *ApimInstanceDetails) GetRouting() []Routing`

GetRouting returns the Routing field if non-nil, zero value otherwise.

### GetRoutingOk

`func (o *ApimInstanceDetails) GetRoutingOk() (*[]Routing, bool)`

GetRoutingOk returns a tuple with the Routing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouting

`func (o *ApimInstanceDetails) SetRouting(v []Routing)`

SetRouting sets Routing field to given value.

### HasRouting

`func (o *ApimInstanceDetails) HasRouting() bool`

HasRouting returns a boolean if a field has been set.

### GetStatus

`func (o *ApimInstanceDetails) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApimInstanceDetails) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApimInstanceDetails) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApimInstanceDetails) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAutodiscoveryInstanceName

`func (o *ApimInstanceDetails) GetAutodiscoveryInstanceName() string`

GetAutodiscoveryInstanceName returns the AutodiscoveryInstanceName field if non-nil, zero value otherwise.

### GetAutodiscoveryInstanceNameOk

`func (o *ApimInstanceDetails) GetAutodiscoveryInstanceNameOk() (*string, bool)`

GetAutodiscoveryInstanceNameOk returns a tuple with the AutodiscoveryInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutodiscoveryInstanceName

`func (o *ApimInstanceDetails) SetAutodiscoveryInstanceName(v string)`

SetAutodiscoveryInstanceName sets AutodiscoveryInstanceName field to given value.

### HasAutodiscoveryInstanceName

`func (o *ApimInstanceDetails) HasAutodiscoveryInstanceName() bool`

HasAutodiscoveryInstanceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


