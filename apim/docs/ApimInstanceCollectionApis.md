# ApimInstanceCollectionApis

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
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**ProviderId** | Pointer to **string** |  | [optional] 
**Deprecated** | Pointer to **bool** |  | [optional] 
**LastActiveDate** | Pointer to **string** |  | [optional] 
**EndpointUri** | Pointer to **string** |  | [optional] 
**EnvironmentId** | Pointer to **string** |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**Stage** | Pointer to **string** |  | [optional] 
**Technology** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Deployment** | Pointer to [**NullableDeployment**](Deployment.md) |  | [optional] 
**Routing** | Pointer to [**[]Routing**](Routing.md) |  | [optional] 
**Pinned** | Pointer to **bool** |  | [optional] 
**ActiveContractsCount** | Pointer to **int32** |  | [optional] 
**AutodiscoveryInstanceName** | Pointer to **string** |  | [optional] 

## Methods

### NewApimInstanceCollectionApis

`func NewApimInstanceCollectionApis() *ApimInstanceCollectionApis`

NewApimInstanceCollectionApis instantiates a new ApimInstanceCollectionApis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApimInstanceCollectionApisWithDefaults

`func NewApimInstanceCollectionApisWithDefaults() *ApimInstanceCollectionApis`

NewApimInstanceCollectionApisWithDefaults instantiates a new ApimInstanceCollectionApis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudit

`func (o *ApimInstanceCollectionApis) GetAudit() Audit`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *ApimInstanceCollectionApis) GetAuditOk() (*Audit, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *ApimInstanceCollectionApis) SetAudit(v Audit)`

SetAudit sets Audit field to given value.

### HasAudit

`func (o *ApimInstanceCollectionApis) HasAudit() bool`

HasAudit returns a boolean if a field has been set.

### GetMasterOrganizationId

`func (o *ApimInstanceCollectionApis) GetMasterOrganizationId() string`

GetMasterOrganizationId returns the MasterOrganizationId field if non-nil, zero value otherwise.

### GetMasterOrganizationIdOk

`func (o *ApimInstanceCollectionApis) GetMasterOrganizationIdOk() (*string, bool)`

GetMasterOrganizationIdOk returns a tuple with the MasterOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterOrganizationId

`func (o *ApimInstanceCollectionApis) SetMasterOrganizationId(v string)`

SetMasterOrganizationId sets MasterOrganizationId field to given value.

### HasMasterOrganizationId

`func (o *ApimInstanceCollectionApis) HasMasterOrganizationId() bool`

HasMasterOrganizationId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *ApimInstanceCollectionApis) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ApimInstanceCollectionApis) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ApimInstanceCollectionApis) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ApimInstanceCollectionApis) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetId

`func (o *ApimInstanceCollectionApis) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApimInstanceCollectionApis) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApimInstanceCollectionApis) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ApimInstanceCollectionApis) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstanceLabel

`func (o *ApimInstanceCollectionApis) GetInstanceLabel() string`

GetInstanceLabel returns the InstanceLabel field if non-nil, zero value otherwise.

### GetInstanceLabelOk

`func (o *ApimInstanceCollectionApis) GetInstanceLabelOk() (*string, bool)`

GetInstanceLabelOk returns a tuple with the InstanceLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceLabel

`func (o *ApimInstanceCollectionApis) SetInstanceLabel(v string)`

SetInstanceLabel sets InstanceLabel field to given value.

### HasInstanceLabel

`func (o *ApimInstanceCollectionApis) HasInstanceLabel() bool`

HasInstanceLabel returns a boolean if a field has been set.

### GetGroupId

`func (o *ApimInstanceCollectionApis) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ApimInstanceCollectionApis) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ApimInstanceCollectionApis) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ApimInstanceCollectionApis) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetAssetId

`func (o *ApimInstanceCollectionApis) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *ApimInstanceCollectionApis) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *ApimInstanceCollectionApis) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *ApimInstanceCollectionApis) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetAssetVersion

`func (o *ApimInstanceCollectionApis) GetAssetVersion() string`

GetAssetVersion returns the AssetVersion field if non-nil, zero value otherwise.

### GetAssetVersionOk

`func (o *ApimInstanceCollectionApis) GetAssetVersionOk() (*string, bool)`

GetAssetVersionOk returns a tuple with the AssetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetVersion

`func (o *ApimInstanceCollectionApis) SetAssetVersion(v string)`

SetAssetVersion sets AssetVersion field to given value.

### HasAssetVersion

`func (o *ApimInstanceCollectionApis) HasAssetVersion() bool`

HasAssetVersion returns a boolean if a field has been set.

### GetProductVersion

`func (o *ApimInstanceCollectionApis) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *ApimInstanceCollectionApis) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *ApimInstanceCollectionApis) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.

### HasProductVersion

`func (o *ApimInstanceCollectionApis) HasProductVersion() bool`

HasProductVersion returns a boolean if a field has been set.

### GetDescription

`func (o *ApimInstanceCollectionApis) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApimInstanceCollectionApis) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApimInstanceCollectionApis) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApimInstanceCollectionApis) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *ApimInstanceCollectionApis) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ApimInstanceCollectionApis) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ApimInstanceCollectionApis) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ApimInstanceCollectionApis) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOrder

`func (o *ApimInstanceCollectionApis) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ApimInstanceCollectionApis) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ApimInstanceCollectionApis) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ApimInstanceCollectionApis) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetProviderId

`func (o *ApimInstanceCollectionApis) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *ApimInstanceCollectionApis) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *ApimInstanceCollectionApis) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *ApimInstanceCollectionApis) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetDeprecated

`func (o *ApimInstanceCollectionApis) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *ApimInstanceCollectionApis) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *ApimInstanceCollectionApis) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *ApimInstanceCollectionApis) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetLastActiveDate

`func (o *ApimInstanceCollectionApis) GetLastActiveDate() string`

GetLastActiveDate returns the LastActiveDate field if non-nil, zero value otherwise.

### GetLastActiveDateOk

`func (o *ApimInstanceCollectionApis) GetLastActiveDateOk() (*string, bool)`

GetLastActiveDateOk returns a tuple with the LastActiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActiveDate

`func (o *ApimInstanceCollectionApis) SetLastActiveDate(v string)`

SetLastActiveDate sets LastActiveDate field to given value.

### HasLastActiveDate

`func (o *ApimInstanceCollectionApis) HasLastActiveDate() bool`

HasLastActiveDate returns a boolean if a field has been set.

### GetEndpointUri

`func (o *ApimInstanceCollectionApis) GetEndpointUri() string`

GetEndpointUri returns the EndpointUri field if non-nil, zero value otherwise.

### GetEndpointUriOk

`func (o *ApimInstanceCollectionApis) GetEndpointUriOk() (*string, bool)`

GetEndpointUriOk returns a tuple with the EndpointUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointUri

`func (o *ApimInstanceCollectionApis) SetEndpointUri(v string)`

SetEndpointUri sets EndpointUri field to given value.

### HasEndpointUri

`func (o *ApimInstanceCollectionApis) HasEndpointUri() bool`

HasEndpointUri returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *ApimInstanceCollectionApis) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *ApimInstanceCollectionApis) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *ApimInstanceCollectionApis) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *ApimInstanceCollectionApis) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetIsPublic

`func (o *ApimInstanceCollectionApis) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *ApimInstanceCollectionApis) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *ApimInstanceCollectionApis) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *ApimInstanceCollectionApis) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetStage

`func (o *ApimInstanceCollectionApis) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *ApimInstanceCollectionApis) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *ApimInstanceCollectionApis) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *ApimInstanceCollectionApis) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetTechnology

`func (o *ApimInstanceCollectionApis) GetTechnology() string`

GetTechnology returns the Technology field if non-nil, zero value otherwise.

### GetTechnologyOk

`func (o *ApimInstanceCollectionApis) GetTechnologyOk() (*string, bool)`

GetTechnologyOk returns a tuple with the Technology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnology

`func (o *ApimInstanceCollectionApis) SetTechnology(v string)`

SetTechnology sets Technology field to given value.

### HasTechnology

`func (o *ApimInstanceCollectionApis) HasTechnology() bool`

HasTechnology returns a boolean if a field has been set.

### GetStatus

`func (o *ApimInstanceCollectionApis) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApimInstanceCollectionApis) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApimInstanceCollectionApis) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApimInstanceCollectionApis) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDeployment

`func (o *ApimInstanceCollectionApis) GetDeployment() Deployment`

GetDeployment returns the Deployment field if non-nil, zero value otherwise.

### GetDeploymentOk

`func (o *ApimInstanceCollectionApis) GetDeploymentOk() (*Deployment, bool)`

GetDeploymentOk returns a tuple with the Deployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployment

`func (o *ApimInstanceCollectionApis) SetDeployment(v Deployment)`

SetDeployment sets Deployment field to given value.

### HasDeployment

`func (o *ApimInstanceCollectionApis) HasDeployment() bool`

HasDeployment returns a boolean if a field has been set.

### SetDeploymentNil

`func (o *ApimInstanceCollectionApis) SetDeploymentNil(b bool)`

 SetDeploymentNil sets the value for Deployment to be an explicit nil

### UnsetDeployment
`func (o *ApimInstanceCollectionApis) UnsetDeployment()`

UnsetDeployment ensures that no value is present for Deployment, not even an explicit nil
### GetRouting

`func (o *ApimInstanceCollectionApis) GetRouting() []Routing`

GetRouting returns the Routing field if non-nil, zero value otherwise.

### GetRoutingOk

`func (o *ApimInstanceCollectionApis) GetRoutingOk() (*[]Routing, bool)`

GetRoutingOk returns a tuple with the Routing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouting

`func (o *ApimInstanceCollectionApis) SetRouting(v []Routing)`

SetRouting sets Routing field to given value.

### HasRouting

`func (o *ApimInstanceCollectionApis) HasRouting() bool`

HasRouting returns a boolean if a field has been set.

### GetPinned

`func (o *ApimInstanceCollectionApis) GetPinned() bool`

GetPinned returns the Pinned field if non-nil, zero value otherwise.

### GetPinnedOk

`func (o *ApimInstanceCollectionApis) GetPinnedOk() (*bool, bool)`

GetPinnedOk returns a tuple with the Pinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinned

`func (o *ApimInstanceCollectionApis) SetPinned(v bool)`

SetPinned sets Pinned field to given value.

### HasPinned

`func (o *ApimInstanceCollectionApis) HasPinned() bool`

HasPinned returns a boolean if a field has been set.

### GetActiveContractsCount

`func (o *ApimInstanceCollectionApis) GetActiveContractsCount() int32`

GetActiveContractsCount returns the ActiveContractsCount field if non-nil, zero value otherwise.

### GetActiveContractsCountOk

`func (o *ApimInstanceCollectionApis) GetActiveContractsCountOk() (*int32, bool)`

GetActiveContractsCountOk returns a tuple with the ActiveContractsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveContractsCount

`func (o *ApimInstanceCollectionApis) SetActiveContractsCount(v int32)`

SetActiveContractsCount sets ActiveContractsCount field to given value.

### HasActiveContractsCount

`func (o *ApimInstanceCollectionApis) HasActiveContractsCount() bool`

HasActiveContractsCount returns a boolean if a field has been set.

### GetAutodiscoveryInstanceName

`func (o *ApimInstanceCollectionApis) GetAutodiscoveryInstanceName() string`

GetAutodiscoveryInstanceName returns the AutodiscoveryInstanceName field if non-nil, zero value otherwise.

### GetAutodiscoveryInstanceNameOk

`func (o *ApimInstanceCollectionApis) GetAutodiscoveryInstanceNameOk() (*string, bool)`

GetAutodiscoveryInstanceNameOk returns a tuple with the AutodiscoveryInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutodiscoveryInstanceName

`func (o *ApimInstanceCollectionApis) SetAutodiscoveryInstanceName(v string)`

SetAutodiscoveryInstanceName sets AutodiscoveryInstanceName field to given value.

### HasAutodiscoveryInstanceName

`func (o *ApimInstanceCollectionApis) HasAutodiscoveryInstanceName() bool`

HasAutodiscoveryInstanceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


