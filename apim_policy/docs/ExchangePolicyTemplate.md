# ExchangePolicyTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audit** | Pointer to [**Audit**](Audit.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**GroupId** | Pointer to **string** |  | [optional] 
**AssetId** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**IsOOTB** | Pointer to **bool** |  | [optional] 
**Stage** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**YamlMd5** | Pointer to **string** |  | [optional] 
**JarMd5** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**MinMuleVersion** | Pointer to **string** |  | [optional] 
**SupportedPoliciesVersions** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**ViolationCategory** | Pointer to **string** |  | [optional] 
**ResourceLevelSupported** | Pointer to **bool** |  | [optional] 
**EncryptionSupported** | Pointer to **bool** |  | [optional] 
**Standalone** | Pointer to **bool** |  | [optional] 
**RequiredCharacteristics** | Pointer to **[]string** |  | [optional] 
**IdentityManagement** | Pointer to [**ExchangePolicyTemplateIdentityManagement**](ExchangePolicyTemplateIdentityManagement.md) |  | [optional] 
**ProvidedCharacteristics** | Pointer to **[]string** |  | [optional] 
**RamlSnippet** | Pointer to **string** |  | [optional] 
**RamlV1Snippet** | Pointer to **string** |  | [optional] 
**OasV2Snippet** | Pointer to **string** |  | [optional] 
**OasV3Snippet** | Pointer to **string** |  | [optional] 
**Applicable** | Pointer to **bool** |  | [optional] 
**Configuration** | Pointer to [**[]PolicyConfiguration**](PolicyConfiguration.md) |  | [optional] 
**AllVersions** | Pointer to [**[]ExchangePolicyTemplateAllVersionsInner**](ExchangePolicyTemplateAllVersionsInner.md) |  | [optional] 

## Methods

### NewExchangePolicyTemplate

`func NewExchangePolicyTemplate() *ExchangePolicyTemplate`

NewExchangePolicyTemplate instantiates a new ExchangePolicyTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangePolicyTemplateWithDefaults

`func NewExchangePolicyTemplateWithDefaults() *ExchangePolicyTemplate`

NewExchangePolicyTemplateWithDefaults instantiates a new ExchangePolicyTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudit

`func (o *ExchangePolicyTemplate) GetAudit() Audit`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *ExchangePolicyTemplate) GetAuditOk() (*Audit, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *ExchangePolicyTemplate) SetAudit(v Audit)`

SetAudit sets Audit field to given value.

### HasAudit

`func (o *ExchangePolicyTemplate) HasAudit() bool`

HasAudit returns a boolean if a field has been set.

### GetId

`func (o *ExchangePolicyTemplate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExchangePolicyTemplate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExchangePolicyTemplate) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ExchangePolicyTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGroupId

`func (o *ExchangePolicyTemplate) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ExchangePolicyTemplate) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ExchangePolicyTemplate) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ExchangePolicyTemplate) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetAssetId

`func (o *ExchangePolicyTemplate) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *ExchangePolicyTemplate) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *ExchangePolicyTemplate) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *ExchangePolicyTemplate) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetVersion

`func (o *ExchangePolicyTemplate) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ExchangePolicyTemplate) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ExchangePolicyTemplate) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ExchangePolicyTemplate) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetName

`func (o *ExchangePolicyTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExchangePolicyTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExchangePolicyTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExchangePolicyTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ExchangePolicyTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExchangePolicyTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExchangePolicyTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExchangePolicyTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *ExchangePolicyTemplate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExchangePolicyTemplate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExchangePolicyTemplate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ExchangePolicyTemplate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIsOOTB

`func (o *ExchangePolicyTemplate) GetIsOOTB() bool`

GetIsOOTB returns the IsOOTB field if non-nil, zero value otherwise.

### GetIsOOTBOk

`func (o *ExchangePolicyTemplate) GetIsOOTBOk() (*bool, bool)`

GetIsOOTBOk returns a tuple with the IsOOTB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOOTB

`func (o *ExchangePolicyTemplate) SetIsOOTB(v bool)`

SetIsOOTB sets IsOOTB field to given value.

### HasIsOOTB

`func (o *ExchangePolicyTemplate) HasIsOOTB() bool`

HasIsOOTB returns a boolean if a field has been set.

### GetStage

`func (o *ExchangePolicyTemplate) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *ExchangePolicyTemplate) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *ExchangePolicyTemplate) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *ExchangePolicyTemplate) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetStatus

`func (o *ExchangePolicyTemplate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExchangePolicyTemplate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExchangePolicyTemplate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExchangePolicyTemplate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetYamlMd5

`func (o *ExchangePolicyTemplate) GetYamlMd5() string`

GetYamlMd5 returns the YamlMd5 field if non-nil, zero value otherwise.

### GetYamlMd5Ok

`func (o *ExchangePolicyTemplate) GetYamlMd5Ok() (*string, bool)`

GetYamlMd5Ok returns a tuple with the YamlMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYamlMd5

`func (o *ExchangePolicyTemplate) SetYamlMd5(v string)`

SetYamlMd5 sets YamlMd5 field to given value.

### HasYamlMd5

`func (o *ExchangePolicyTemplate) HasYamlMd5() bool`

HasYamlMd5 returns a boolean if a field has been set.

### GetJarMd5

`func (o *ExchangePolicyTemplate) GetJarMd5() string`

GetJarMd5 returns the JarMd5 field if non-nil, zero value otherwise.

### GetJarMd5Ok

`func (o *ExchangePolicyTemplate) GetJarMd5Ok() (*string, bool)`

GetJarMd5Ok returns a tuple with the JarMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJarMd5

`func (o *ExchangePolicyTemplate) SetJarMd5(v string)`

SetJarMd5 sets JarMd5 field to given value.

### HasJarMd5

`func (o *ExchangePolicyTemplate) HasJarMd5() bool`

HasJarMd5 returns a boolean if a field has been set.

### GetOrgId

`func (o *ExchangePolicyTemplate) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ExchangePolicyTemplate) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ExchangePolicyTemplate) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ExchangePolicyTemplate) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetMinMuleVersion

`func (o *ExchangePolicyTemplate) GetMinMuleVersion() string`

GetMinMuleVersion returns the MinMuleVersion field if non-nil, zero value otherwise.

### GetMinMuleVersionOk

`func (o *ExchangePolicyTemplate) GetMinMuleVersionOk() (*string, bool)`

GetMinMuleVersionOk returns a tuple with the MinMuleVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinMuleVersion

`func (o *ExchangePolicyTemplate) SetMinMuleVersion(v string)`

SetMinMuleVersion sets MinMuleVersion field to given value.

### HasMinMuleVersion

`func (o *ExchangePolicyTemplate) HasMinMuleVersion() bool`

HasMinMuleVersion returns a boolean if a field has been set.

### GetSupportedPoliciesVersions

`func (o *ExchangePolicyTemplate) GetSupportedPoliciesVersions() string`

GetSupportedPoliciesVersions returns the SupportedPoliciesVersions field if non-nil, zero value otherwise.

### GetSupportedPoliciesVersionsOk

`func (o *ExchangePolicyTemplate) GetSupportedPoliciesVersionsOk() (*string, bool)`

GetSupportedPoliciesVersionsOk returns a tuple with the SupportedPoliciesVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedPoliciesVersions

`func (o *ExchangePolicyTemplate) SetSupportedPoliciesVersions(v string)`

SetSupportedPoliciesVersions sets SupportedPoliciesVersions field to given value.

### HasSupportedPoliciesVersions

`func (o *ExchangePolicyTemplate) HasSupportedPoliciesVersions() bool`

HasSupportedPoliciesVersions returns a boolean if a field has been set.

### GetCategory

`func (o *ExchangePolicyTemplate) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ExchangePolicyTemplate) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ExchangePolicyTemplate) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ExchangePolicyTemplate) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetViolationCategory

`func (o *ExchangePolicyTemplate) GetViolationCategory() string`

GetViolationCategory returns the ViolationCategory field if non-nil, zero value otherwise.

### GetViolationCategoryOk

`func (o *ExchangePolicyTemplate) GetViolationCategoryOk() (*string, bool)`

GetViolationCategoryOk returns a tuple with the ViolationCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolationCategory

`func (o *ExchangePolicyTemplate) SetViolationCategory(v string)`

SetViolationCategory sets ViolationCategory field to given value.

### HasViolationCategory

`func (o *ExchangePolicyTemplate) HasViolationCategory() bool`

HasViolationCategory returns a boolean if a field has been set.

### GetResourceLevelSupported

`func (o *ExchangePolicyTemplate) GetResourceLevelSupported() bool`

GetResourceLevelSupported returns the ResourceLevelSupported field if non-nil, zero value otherwise.

### GetResourceLevelSupportedOk

`func (o *ExchangePolicyTemplate) GetResourceLevelSupportedOk() (*bool, bool)`

GetResourceLevelSupportedOk returns a tuple with the ResourceLevelSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLevelSupported

`func (o *ExchangePolicyTemplate) SetResourceLevelSupported(v bool)`

SetResourceLevelSupported sets ResourceLevelSupported field to given value.

### HasResourceLevelSupported

`func (o *ExchangePolicyTemplate) HasResourceLevelSupported() bool`

HasResourceLevelSupported returns a boolean if a field has been set.

### GetEncryptionSupported

`func (o *ExchangePolicyTemplate) GetEncryptionSupported() bool`

GetEncryptionSupported returns the EncryptionSupported field if non-nil, zero value otherwise.

### GetEncryptionSupportedOk

`func (o *ExchangePolicyTemplate) GetEncryptionSupportedOk() (*bool, bool)`

GetEncryptionSupportedOk returns a tuple with the EncryptionSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionSupported

`func (o *ExchangePolicyTemplate) SetEncryptionSupported(v bool)`

SetEncryptionSupported sets EncryptionSupported field to given value.

### HasEncryptionSupported

`func (o *ExchangePolicyTemplate) HasEncryptionSupported() bool`

HasEncryptionSupported returns a boolean if a field has been set.

### GetStandalone

`func (o *ExchangePolicyTemplate) GetStandalone() bool`

GetStandalone returns the Standalone field if non-nil, zero value otherwise.

### GetStandaloneOk

`func (o *ExchangePolicyTemplate) GetStandaloneOk() (*bool, bool)`

GetStandaloneOk returns a tuple with the Standalone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandalone

`func (o *ExchangePolicyTemplate) SetStandalone(v bool)`

SetStandalone sets Standalone field to given value.

### HasStandalone

`func (o *ExchangePolicyTemplate) HasStandalone() bool`

HasStandalone returns a boolean if a field has been set.

### GetRequiredCharacteristics

`func (o *ExchangePolicyTemplate) GetRequiredCharacteristics() []string`

GetRequiredCharacteristics returns the RequiredCharacteristics field if non-nil, zero value otherwise.

### GetRequiredCharacteristicsOk

`func (o *ExchangePolicyTemplate) GetRequiredCharacteristicsOk() (*[]string, bool)`

GetRequiredCharacteristicsOk returns a tuple with the RequiredCharacteristics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredCharacteristics

`func (o *ExchangePolicyTemplate) SetRequiredCharacteristics(v []string)`

SetRequiredCharacteristics sets RequiredCharacteristics field to given value.

### HasRequiredCharacteristics

`func (o *ExchangePolicyTemplate) HasRequiredCharacteristics() bool`

HasRequiredCharacteristics returns a boolean if a field has been set.

### GetIdentityManagement

`func (o *ExchangePolicyTemplate) GetIdentityManagement() ExchangePolicyTemplateIdentityManagement`

GetIdentityManagement returns the IdentityManagement field if non-nil, zero value otherwise.

### GetIdentityManagementOk

`func (o *ExchangePolicyTemplate) GetIdentityManagementOk() (*ExchangePolicyTemplateIdentityManagement, bool)`

GetIdentityManagementOk returns a tuple with the IdentityManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityManagement

`func (o *ExchangePolicyTemplate) SetIdentityManagement(v ExchangePolicyTemplateIdentityManagement)`

SetIdentityManagement sets IdentityManagement field to given value.

### HasIdentityManagement

`func (o *ExchangePolicyTemplate) HasIdentityManagement() bool`

HasIdentityManagement returns a boolean if a field has been set.

### GetProvidedCharacteristics

`func (o *ExchangePolicyTemplate) GetProvidedCharacteristics() []string`

GetProvidedCharacteristics returns the ProvidedCharacteristics field if non-nil, zero value otherwise.

### GetProvidedCharacteristicsOk

`func (o *ExchangePolicyTemplate) GetProvidedCharacteristicsOk() (*[]string, bool)`

GetProvidedCharacteristicsOk returns a tuple with the ProvidedCharacteristics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvidedCharacteristics

`func (o *ExchangePolicyTemplate) SetProvidedCharacteristics(v []string)`

SetProvidedCharacteristics sets ProvidedCharacteristics field to given value.

### HasProvidedCharacteristics

`func (o *ExchangePolicyTemplate) HasProvidedCharacteristics() bool`

HasProvidedCharacteristics returns a boolean if a field has been set.

### GetRamlSnippet

`func (o *ExchangePolicyTemplate) GetRamlSnippet() string`

GetRamlSnippet returns the RamlSnippet field if non-nil, zero value otherwise.

### GetRamlSnippetOk

`func (o *ExchangePolicyTemplate) GetRamlSnippetOk() (*string, bool)`

GetRamlSnippetOk returns a tuple with the RamlSnippet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamlSnippet

`func (o *ExchangePolicyTemplate) SetRamlSnippet(v string)`

SetRamlSnippet sets RamlSnippet field to given value.

### HasRamlSnippet

`func (o *ExchangePolicyTemplate) HasRamlSnippet() bool`

HasRamlSnippet returns a boolean if a field has been set.

### GetRamlV1Snippet

`func (o *ExchangePolicyTemplate) GetRamlV1Snippet() string`

GetRamlV1Snippet returns the RamlV1Snippet field if non-nil, zero value otherwise.

### GetRamlV1SnippetOk

`func (o *ExchangePolicyTemplate) GetRamlV1SnippetOk() (*string, bool)`

GetRamlV1SnippetOk returns a tuple with the RamlV1Snippet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamlV1Snippet

`func (o *ExchangePolicyTemplate) SetRamlV1Snippet(v string)`

SetRamlV1Snippet sets RamlV1Snippet field to given value.

### HasRamlV1Snippet

`func (o *ExchangePolicyTemplate) HasRamlV1Snippet() bool`

HasRamlV1Snippet returns a boolean if a field has been set.

### GetOasV2Snippet

`func (o *ExchangePolicyTemplate) GetOasV2Snippet() string`

GetOasV2Snippet returns the OasV2Snippet field if non-nil, zero value otherwise.

### GetOasV2SnippetOk

`func (o *ExchangePolicyTemplate) GetOasV2SnippetOk() (*string, bool)`

GetOasV2SnippetOk returns a tuple with the OasV2Snippet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOasV2Snippet

`func (o *ExchangePolicyTemplate) SetOasV2Snippet(v string)`

SetOasV2Snippet sets OasV2Snippet field to given value.

### HasOasV2Snippet

`func (o *ExchangePolicyTemplate) HasOasV2Snippet() bool`

HasOasV2Snippet returns a boolean if a field has been set.

### GetOasV3Snippet

`func (o *ExchangePolicyTemplate) GetOasV3Snippet() string`

GetOasV3Snippet returns the OasV3Snippet field if non-nil, zero value otherwise.

### GetOasV3SnippetOk

`func (o *ExchangePolicyTemplate) GetOasV3SnippetOk() (*string, bool)`

GetOasV3SnippetOk returns a tuple with the OasV3Snippet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOasV3Snippet

`func (o *ExchangePolicyTemplate) SetOasV3Snippet(v string)`

SetOasV3Snippet sets OasV3Snippet field to given value.

### HasOasV3Snippet

`func (o *ExchangePolicyTemplate) HasOasV3Snippet() bool`

HasOasV3Snippet returns a boolean if a field has been set.

### GetApplicable

`func (o *ExchangePolicyTemplate) GetApplicable() bool`

GetApplicable returns the Applicable field if non-nil, zero value otherwise.

### GetApplicableOk

`func (o *ExchangePolicyTemplate) GetApplicableOk() (*bool, bool)`

GetApplicableOk returns a tuple with the Applicable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicable

`func (o *ExchangePolicyTemplate) SetApplicable(v bool)`

SetApplicable sets Applicable field to given value.

### HasApplicable

`func (o *ExchangePolicyTemplate) HasApplicable() bool`

HasApplicable returns a boolean if a field has been set.

### GetConfiguration

`func (o *ExchangePolicyTemplate) GetConfiguration() []PolicyConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ExchangePolicyTemplate) GetConfigurationOk() (*[]PolicyConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ExchangePolicyTemplate) SetConfiguration(v []PolicyConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *ExchangePolicyTemplate) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetAllVersions

`func (o *ExchangePolicyTemplate) GetAllVersions() []ExchangePolicyTemplateAllVersionsInner`

GetAllVersions returns the AllVersions field if non-nil, zero value otherwise.

### GetAllVersionsOk

`func (o *ExchangePolicyTemplate) GetAllVersionsOk() (*[]ExchangePolicyTemplateAllVersionsInner, bool)`

GetAllVersionsOk returns a tuple with the AllVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllVersions

`func (o *ExchangePolicyTemplate) SetAllVersions(v []ExchangePolicyTemplateAllVersionsInner)`

SetAllVersions sets AllVersions field to given value.

### HasAllVersions

`func (o *ExchangePolicyTemplate) HasAllVersions() bool`

HasAllVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


