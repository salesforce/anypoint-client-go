# ApimInstancePostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | Pointer to **string** |  | [optional] 
**InstanceLabel** | Pointer to **string** |  | [optional] 
**ProviderId** | Pointer to **NullableString** |  | [optional] 
**Technology** | Pointer to **string** |  | [optional] 
**AssetVersion** | Pointer to **string** |  | [optional] 
**ProductVersion** | Pointer to **string** |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**Stage** | Pointer to **string** |  | [optional] 
**Audit** | Pointer to [**Audit**](Audit.md) |  | [optional] 
**MasterOrganizationId** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**GroupId** | Pointer to **string** |  | [optional] 
**AssetId** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Endpoint** | Pointer to [**Endpoint**](Endpoint.md) |  | [optional] 
**AutodiscoveryInstanceName** | Pointer to **string** |  | [optional] 

## Methods

### NewApimInstancePostResponse

`func NewApimInstancePostResponse() *ApimInstancePostResponse`

NewApimInstancePostResponse instantiates a new ApimInstancePostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApimInstancePostResponseWithDefaults

`func NewApimInstancePostResponseWithDefaults() *ApimInstancePostResponse`

NewApimInstancePostResponseWithDefaults instantiates a new ApimInstancePostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *ApimInstancePostResponse) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *ApimInstancePostResponse) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *ApimInstancePostResponse) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *ApimInstancePostResponse) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetInstanceLabel

`func (o *ApimInstancePostResponse) GetInstanceLabel() string`

GetInstanceLabel returns the InstanceLabel field if non-nil, zero value otherwise.

### GetInstanceLabelOk

`func (o *ApimInstancePostResponse) GetInstanceLabelOk() (*string, bool)`

GetInstanceLabelOk returns a tuple with the InstanceLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceLabel

`func (o *ApimInstancePostResponse) SetInstanceLabel(v string)`

SetInstanceLabel sets InstanceLabel field to given value.

### HasInstanceLabel

`func (o *ApimInstancePostResponse) HasInstanceLabel() bool`

HasInstanceLabel returns a boolean if a field has been set.

### GetProviderId

`func (o *ApimInstancePostResponse) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *ApimInstancePostResponse) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *ApimInstancePostResponse) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *ApimInstancePostResponse) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### SetProviderIdNil

`func (o *ApimInstancePostResponse) SetProviderIdNil(b bool)`

 SetProviderIdNil sets the value for ProviderId to be an explicit nil

### UnsetProviderId
`func (o *ApimInstancePostResponse) UnsetProviderId()`

UnsetProviderId ensures that no value is present for ProviderId, not even an explicit nil
### GetTechnology

`func (o *ApimInstancePostResponse) GetTechnology() string`

GetTechnology returns the Technology field if non-nil, zero value otherwise.

### GetTechnologyOk

`func (o *ApimInstancePostResponse) GetTechnologyOk() (*string, bool)`

GetTechnologyOk returns a tuple with the Technology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnology

`func (o *ApimInstancePostResponse) SetTechnology(v string)`

SetTechnology sets Technology field to given value.

### HasTechnology

`func (o *ApimInstancePostResponse) HasTechnology() bool`

HasTechnology returns a boolean if a field has been set.

### GetAssetVersion

`func (o *ApimInstancePostResponse) GetAssetVersion() string`

GetAssetVersion returns the AssetVersion field if non-nil, zero value otherwise.

### GetAssetVersionOk

`func (o *ApimInstancePostResponse) GetAssetVersionOk() (*string, bool)`

GetAssetVersionOk returns a tuple with the AssetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetVersion

`func (o *ApimInstancePostResponse) SetAssetVersion(v string)`

SetAssetVersion sets AssetVersion field to given value.

### HasAssetVersion

`func (o *ApimInstancePostResponse) HasAssetVersion() bool`

HasAssetVersion returns a boolean if a field has been set.

### GetProductVersion

`func (o *ApimInstancePostResponse) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *ApimInstancePostResponse) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *ApimInstancePostResponse) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.

### HasProductVersion

`func (o *ApimInstancePostResponse) HasProductVersion() bool`

HasProductVersion returns a boolean if a field has been set.

### GetOrder

`func (o *ApimInstancePostResponse) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ApimInstancePostResponse) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ApimInstancePostResponse) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ApimInstancePostResponse) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetStage

`func (o *ApimInstancePostResponse) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *ApimInstancePostResponse) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *ApimInstancePostResponse) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *ApimInstancePostResponse) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetAudit

`func (o *ApimInstancePostResponse) GetAudit() Audit`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *ApimInstancePostResponse) GetAuditOk() (*Audit, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *ApimInstancePostResponse) SetAudit(v Audit)`

SetAudit sets Audit field to given value.

### HasAudit

`func (o *ApimInstancePostResponse) HasAudit() bool`

HasAudit returns a boolean if a field has been set.

### GetMasterOrganizationId

`func (o *ApimInstancePostResponse) GetMasterOrganizationId() string`

GetMasterOrganizationId returns the MasterOrganizationId field if non-nil, zero value otherwise.

### GetMasterOrganizationIdOk

`func (o *ApimInstancePostResponse) GetMasterOrganizationIdOk() (*string, bool)`

GetMasterOrganizationIdOk returns a tuple with the MasterOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterOrganizationId

`func (o *ApimInstancePostResponse) SetMasterOrganizationId(v string)`

SetMasterOrganizationId sets MasterOrganizationId field to given value.

### HasMasterOrganizationId

`func (o *ApimInstancePostResponse) HasMasterOrganizationId() bool`

HasMasterOrganizationId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *ApimInstancePostResponse) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ApimInstancePostResponse) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ApimInstancePostResponse) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ApimInstancePostResponse) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetId

`func (o *ApimInstancePostResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApimInstancePostResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApimInstancePostResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ApimInstancePostResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGroupId

`func (o *ApimInstancePostResponse) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ApimInstancePostResponse) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ApimInstancePostResponse) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ApimInstancePostResponse) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetAssetId

`func (o *ApimInstancePostResponse) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *ApimInstancePostResponse) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *ApimInstancePostResponse) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *ApimInstancePostResponse) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetTags

`func (o *ApimInstancePostResponse) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ApimInstancePostResponse) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ApimInstancePostResponse) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ApimInstancePostResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEndpoint

`func (o *ApimInstancePostResponse) GetEndpoint() Endpoint`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ApimInstancePostResponse) GetEndpointOk() (*Endpoint, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ApimInstancePostResponse) SetEndpoint(v Endpoint)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *ApimInstancePostResponse) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetAutodiscoveryInstanceName

`func (o *ApimInstancePostResponse) GetAutodiscoveryInstanceName() string`

GetAutodiscoveryInstanceName returns the AutodiscoveryInstanceName field if non-nil, zero value otherwise.

### GetAutodiscoveryInstanceNameOk

`func (o *ApimInstancePostResponse) GetAutodiscoveryInstanceNameOk() (*string, bool)`

GetAutodiscoveryInstanceNameOk returns a tuple with the AutodiscoveryInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutodiscoveryInstanceName

`func (o *ApimInstancePostResponse) SetAutodiscoveryInstanceName(v string)`

SetAutodiscoveryInstanceName sets AutodiscoveryInstanceName field to given value.

### HasAutodiscoveryInstanceName

`func (o *ApimInstancePostResponse) HasAutodiscoveryInstanceName() bool`

HasAutodiscoveryInstanceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


