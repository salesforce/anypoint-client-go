# ApimPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audit** | Pointer to [**Audit**](Audit.md) |  | [optional] 
**MasterOrganizationId** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**PolicyTemplateId** | Pointer to **string** |  | [optional] 
**ConfigurationData** | Pointer to **map[string]interface{}** |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**PointcutData** | Pointer to [**[]PointcutDataItem**](PointcutDataItem.md) |  | [optional] 
**GroupId** | Pointer to **string** |  | [optional] 
**AssetId** | Pointer to **string** |  | [optional] 
**AssetVersion** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**ApiId** | Pointer to **int32** |  | [optional] 

## Methods

### NewApimPolicy

`func NewApimPolicy() *ApimPolicy`

NewApimPolicy instantiates a new ApimPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApimPolicyWithDefaults

`func NewApimPolicyWithDefaults() *ApimPolicy`

NewApimPolicyWithDefaults instantiates a new ApimPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudit

`func (o *ApimPolicy) GetAudit() Audit`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *ApimPolicy) GetAuditOk() (*Audit, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *ApimPolicy) SetAudit(v Audit)`

SetAudit sets Audit field to given value.

### HasAudit

`func (o *ApimPolicy) HasAudit() bool`

HasAudit returns a boolean if a field has been set.

### GetMasterOrganizationId

`func (o *ApimPolicy) GetMasterOrganizationId() string`

GetMasterOrganizationId returns the MasterOrganizationId field if non-nil, zero value otherwise.

### GetMasterOrganizationIdOk

`func (o *ApimPolicy) GetMasterOrganizationIdOk() (*string, bool)`

GetMasterOrganizationIdOk returns a tuple with the MasterOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterOrganizationId

`func (o *ApimPolicy) SetMasterOrganizationId(v string)`

SetMasterOrganizationId sets MasterOrganizationId field to given value.

### HasMasterOrganizationId

`func (o *ApimPolicy) HasMasterOrganizationId() bool`

HasMasterOrganizationId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *ApimPolicy) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ApimPolicy) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ApimPolicy) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ApimPolicy) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetId

`func (o *ApimPolicy) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApimPolicy) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApimPolicy) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ApimPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPolicyTemplateId

`func (o *ApimPolicy) GetPolicyTemplateId() string`

GetPolicyTemplateId returns the PolicyTemplateId field if non-nil, zero value otherwise.

### GetPolicyTemplateIdOk

`func (o *ApimPolicy) GetPolicyTemplateIdOk() (*string, bool)`

GetPolicyTemplateIdOk returns a tuple with the PolicyTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyTemplateId

`func (o *ApimPolicy) SetPolicyTemplateId(v string)`

SetPolicyTemplateId sets PolicyTemplateId field to given value.

### HasPolicyTemplateId

`func (o *ApimPolicy) HasPolicyTemplateId() bool`

HasPolicyTemplateId returns a boolean if a field has been set.

### GetConfigurationData

`func (o *ApimPolicy) GetConfigurationData() map[string]interface{}`

GetConfigurationData returns the ConfigurationData field if non-nil, zero value otherwise.

### GetConfigurationDataOk

`func (o *ApimPolicy) GetConfigurationDataOk() (*map[string]interface{}, bool)`

GetConfigurationDataOk returns a tuple with the ConfigurationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationData

`func (o *ApimPolicy) SetConfigurationData(v map[string]interface{})`

SetConfigurationData sets ConfigurationData field to given value.

### HasConfigurationData

`func (o *ApimPolicy) HasConfigurationData() bool`

HasConfigurationData returns a boolean if a field has been set.

### GetOrder

`func (o *ApimPolicy) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ApimPolicy) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ApimPolicy) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ApimPolicy) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetDisabled

`func (o *ApimPolicy) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ApimPolicy) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ApimPolicy) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ApimPolicy) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetPointcutData

`func (o *ApimPolicy) GetPointcutData() []PointcutDataItem`

GetPointcutData returns the PointcutData field if non-nil, zero value otherwise.

### GetPointcutDataOk

`func (o *ApimPolicy) GetPointcutDataOk() (*[]PointcutDataItem, bool)`

GetPointcutDataOk returns a tuple with the PointcutData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointcutData

`func (o *ApimPolicy) SetPointcutData(v []PointcutDataItem)`

SetPointcutData sets PointcutData field to given value.

### HasPointcutData

`func (o *ApimPolicy) HasPointcutData() bool`

HasPointcutData returns a boolean if a field has been set.

### SetPointcutDataNil

`func (o *ApimPolicy) SetPointcutDataNil(b bool)`

 SetPointcutDataNil sets the value for PointcutData to be an explicit nil

### UnsetPointcutData
`func (o *ApimPolicy) UnsetPointcutData()`

UnsetPointcutData ensures that no value is present for PointcutData, not even an explicit nil
### GetGroupId

`func (o *ApimPolicy) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ApimPolicy) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ApimPolicy) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ApimPolicy) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetAssetId

`func (o *ApimPolicy) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *ApimPolicy) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *ApimPolicy) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *ApimPolicy) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetAssetVersion

`func (o *ApimPolicy) GetAssetVersion() string`

GetAssetVersion returns the AssetVersion field if non-nil, zero value otherwise.

### GetAssetVersionOk

`func (o *ApimPolicy) GetAssetVersionOk() (*string, bool)`

GetAssetVersionOk returns a tuple with the AssetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetVersion

`func (o *ApimPolicy) SetAssetVersion(v string)`

SetAssetVersion sets AssetVersion field to given value.

### HasAssetVersion

`func (o *ApimPolicy) HasAssetVersion() bool`

HasAssetVersion returns a boolean if a field has been set.

### GetType

`func (o *ApimPolicy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApimPolicy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApimPolicy) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApimPolicy) HasType() bool`

HasType returns a boolean if a field has been set.

### GetApiId

`func (o *ApimPolicy) GetApiId() int32`

GetApiId returns the ApiId field if non-nil, zero value otherwise.

### GetApiIdOk

`func (o *ApimPolicy) GetApiIdOk() (*int32, bool)`

GetApiIdOk returns a tuple with the ApiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiId

`func (o *ApimPolicy) SetApiId(v int32)`

SetApiId sets ApiId field to given value.

### HasApiId

`func (o *ApimPolicy) HasApiId() bool`

HasApiId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


