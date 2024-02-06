# AutomatedPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audit** | Pointer to [**Audit**](Audit.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**RuleOfApplication** | Pointer to [**AutomatedPolicyRuleOfApplication**](AutomatedPolicyRuleOfApplication.md) |  | [optional] 
**GroupId** | Pointer to **string** |  | [optional] 
**AssetId** | Pointer to **string** |  | [optional] 
**AssetVersion** | Pointer to **string** |  | [optional] 
**ConfigurationData** | Pointer to **map[string]interface{}** |  | [optional] 
**PointcutData** | Pointer to [**[]PointcutDataItem**](PointcutDataItem.md) |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**ImplementationAssets** | Pointer to [**[]ImplementationAsset**](ImplementationAsset.md) |  | [optional] 

## Methods

### NewAutomatedPolicy

`func NewAutomatedPolicy() *AutomatedPolicy`

NewAutomatedPolicy instantiates a new AutomatedPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomatedPolicyWithDefaults

`func NewAutomatedPolicyWithDefaults() *AutomatedPolicy`

NewAutomatedPolicyWithDefaults instantiates a new AutomatedPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudit

`func (o *AutomatedPolicy) GetAudit() Audit`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *AutomatedPolicy) GetAuditOk() (*Audit, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *AutomatedPolicy) SetAudit(v Audit)`

SetAudit sets Audit field to given value.

### HasAudit

`func (o *AutomatedPolicy) HasAudit() bool`

HasAudit returns a boolean if a field has been set.

### GetId

`func (o *AutomatedPolicy) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutomatedPolicy) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutomatedPolicy) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AutomatedPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRuleOfApplication

`func (o *AutomatedPolicy) GetRuleOfApplication() AutomatedPolicyRuleOfApplication`

GetRuleOfApplication returns the RuleOfApplication field if non-nil, zero value otherwise.

### GetRuleOfApplicationOk

`func (o *AutomatedPolicy) GetRuleOfApplicationOk() (*AutomatedPolicyRuleOfApplication, bool)`

GetRuleOfApplicationOk returns a tuple with the RuleOfApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleOfApplication

`func (o *AutomatedPolicy) SetRuleOfApplication(v AutomatedPolicyRuleOfApplication)`

SetRuleOfApplication sets RuleOfApplication field to given value.

### HasRuleOfApplication

`func (o *AutomatedPolicy) HasRuleOfApplication() bool`

HasRuleOfApplication returns a boolean if a field has been set.

### GetGroupId

`func (o *AutomatedPolicy) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *AutomatedPolicy) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *AutomatedPolicy) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *AutomatedPolicy) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetAssetId

`func (o *AutomatedPolicy) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *AutomatedPolicy) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *AutomatedPolicy) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *AutomatedPolicy) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetAssetVersion

`func (o *AutomatedPolicy) GetAssetVersion() string`

GetAssetVersion returns the AssetVersion field if non-nil, zero value otherwise.

### GetAssetVersionOk

`func (o *AutomatedPolicy) GetAssetVersionOk() (*string, bool)`

GetAssetVersionOk returns a tuple with the AssetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetVersion

`func (o *AutomatedPolicy) SetAssetVersion(v string)`

SetAssetVersion sets AssetVersion field to given value.

### HasAssetVersion

`func (o *AutomatedPolicy) HasAssetVersion() bool`

HasAssetVersion returns a boolean if a field has been set.

### GetConfigurationData

`func (o *AutomatedPolicy) GetConfigurationData() map[string]interface{}`

GetConfigurationData returns the ConfigurationData field if non-nil, zero value otherwise.

### GetConfigurationDataOk

`func (o *AutomatedPolicy) GetConfigurationDataOk() (*map[string]interface{}, bool)`

GetConfigurationDataOk returns a tuple with the ConfigurationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationData

`func (o *AutomatedPolicy) SetConfigurationData(v map[string]interface{})`

SetConfigurationData sets ConfigurationData field to given value.

### HasConfigurationData

`func (o *AutomatedPolicy) HasConfigurationData() bool`

HasConfigurationData returns a boolean if a field has been set.

### GetPointcutData

`func (o *AutomatedPolicy) GetPointcutData() []PointcutDataItem`

GetPointcutData returns the PointcutData field if non-nil, zero value otherwise.

### GetPointcutDataOk

`func (o *AutomatedPolicy) GetPointcutDataOk() (*[]PointcutDataItem, bool)`

GetPointcutDataOk returns a tuple with the PointcutData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointcutData

`func (o *AutomatedPolicy) SetPointcutData(v []PointcutDataItem)`

SetPointcutData sets PointcutData field to given value.

### HasPointcutData

`func (o *AutomatedPolicy) HasPointcutData() bool`

HasPointcutData returns a boolean if a field has been set.

### SetPointcutDataNil

`func (o *AutomatedPolicy) SetPointcutDataNil(b bool)`

 SetPointcutDataNil sets the value for PointcutData to be an explicit nil

### UnsetPointcutData
`func (o *AutomatedPolicy) UnsetPointcutData()`

UnsetPointcutData ensures that no value is present for PointcutData, not even an explicit nil
### GetOrder

`func (o *AutomatedPolicy) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *AutomatedPolicy) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *AutomatedPolicy) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *AutomatedPolicy) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetDisabled

`func (o *AutomatedPolicy) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *AutomatedPolicy) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *AutomatedPolicy) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *AutomatedPolicy) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetImplementationAssets

`func (o *AutomatedPolicy) GetImplementationAssets() []ImplementationAsset`

GetImplementationAssets returns the ImplementationAssets field if non-nil, zero value otherwise.

### GetImplementationAssetsOk

`func (o *AutomatedPolicy) GetImplementationAssetsOk() (*[]ImplementationAsset, bool)`

GetImplementationAssetsOk returns a tuple with the ImplementationAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementationAssets

`func (o *AutomatedPolicy) SetImplementationAssets(v []ImplementationAsset)`

SetImplementationAssets sets ImplementationAssets field to given value.

### HasImplementationAssets

`func (o *AutomatedPolicy) HasImplementationAssets() bool`

HasImplementationAssets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


