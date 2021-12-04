# Asset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** |  | [optional] 
**AssetId** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**MinorVersion** | Pointer to **string** |  | [optional] 
**VersionGroup** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ContactName** | Pointer to **NullableString** |  | [optional] 
**ContactEmail** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**IsSnapshot** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Attributes** | Pointer to [**[]Attribute**](Attribute.md) |  | [optional] 
**Categories** | Pointer to [**[]Category**](Category.md) |  | [optional] 
**Files** | Pointer to [**[]*os.File**](*os.File.md) |  | [optional] 
**CustomFields** | Pointer to [**[]CustomField**](CustomField.md) |  | [optional] 
**CreatedById** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Versions** | Pointer to **[]string** |  | [optional] 
**Dependencies** | Pointer to [**[]Dependency**](Dependency.md) |  | [optional] 

## Methods

### NewAsset

`func NewAsset() *Asset`

NewAsset instantiates a new Asset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetWithDefaults

`func NewAssetWithDefaults() *Asset`

NewAssetWithDefaults instantiates a new Asset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *Asset) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *Asset) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *Asset) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *Asset) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetAssetId

`func (o *Asset) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *Asset) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *Asset) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *Asset) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetVersion

`func (o *Asset) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Asset) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Asset) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Asset) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetMinorVersion

`func (o *Asset) GetMinorVersion() string`

GetMinorVersion returns the MinorVersion field if non-nil, zero value otherwise.

### GetMinorVersionOk

`func (o *Asset) GetMinorVersionOk() (*string, bool)`

GetMinorVersionOk returns a tuple with the MinorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinorVersion

`func (o *Asset) SetMinorVersion(v string)`

SetMinorVersion sets MinorVersion field to given value.

### HasMinorVersion

`func (o *Asset) HasMinorVersion() bool`

HasMinorVersion returns a boolean if a field has been set.

### GetVersionGroup

`func (o *Asset) GetVersionGroup() string`

GetVersionGroup returns the VersionGroup field if non-nil, zero value otherwise.

### GetVersionGroupOk

`func (o *Asset) GetVersionGroupOk() (*string, bool)`

GetVersionGroupOk returns a tuple with the VersionGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionGroup

`func (o *Asset) SetVersionGroup(v string)`

SetVersionGroup sets VersionGroup field to given value.

### HasVersionGroup

`func (o *Asset) HasVersionGroup() bool`

HasVersionGroup returns a boolean if a field has been set.

### GetDescription

`func (o *Asset) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Asset) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Asset) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Asset) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOrganizationId

`func (o *Asset) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Asset) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *Asset) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *Asset) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetIsPublic

`func (o *Asset) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *Asset) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *Asset) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *Asset) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetName

`func (o *Asset) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Asset) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Asset) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Asset) HasName() bool`

HasName returns a boolean if a field has been set.

### GetContactName

`func (o *Asset) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *Asset) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *Asset) SetContactName(v string)`

SetContactName sets ContactName field to given value.

### HasContactName

`func (o *Asset) HasContactName() bool`

HasContactName returns a boolean if a field has been set.

### SetContactNameNil

`func (o *Asset) SetContactNameNil(b bool)`

 SetContactNameNil sets the value for ContactName to be an explicit nil

### UnsetContactName
`func (o *Asset) UnsetContactName()`

UnsetContactName ensures that no value is present for ContactName, not even an explicit nil
### GetContactEmail

`func (o *Asset) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *Asset) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *Asset) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *Asset) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### SetContactEmailNil

`func (o *Asset) SetContactEmailNil(b bool)`

 SetContactEmailNil sets the value for ContactEmail to be an explicit nil

### UnsetContactEmail
`func (o *Asset) UnsetContactEmail()`

UnsetContactEmail ensures that no value is present for ContactEmail, not even an explicit nil
### GetType

`func (o *Asset) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Asset) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Asset) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Asset) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIsSnapshot

`func (o *Asset) GetIsSnapshot() bool`

GetIsSnapshot returns the IsSnapshot field if non-nil, zero value otherwise.

### GetIsSnapshotOk

`func (o *Asset) GetIsSnapshotOk() (*bool, bool)`

GetIsSnapshotOk returns a tuple with the IsSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSnapshot

`func (o *Asset) SetIsSnapshot(v bool)`

SetIsSnapshot sets IsSnapshot field to given value.

### HasIsSnapshot

`func (o *Asset) HasIsSnapshot() bool`

HasIsSnapshot returns a boolean if a field has been set.

### GetStatus

`func (o *Asset) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Asset) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Asset) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Asset) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLabels

`func (o *Asset) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Asset) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Asset) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *Asset) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetAttributes

`func (o *Asset) GetAttributes() []Attribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Asset) GetAttributesOk() (*[]Attribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Asset) SetAttributes(v []Attribute)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Asset) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetCategories

`func (o *Asset) GetCategories() []Category`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *Asset) GetCategoriesOk() (*[]Category, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *Asset) SetCategories(v []Category)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *Asset) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetFiles

`func (o *Asset) GetFiles() []*os.File`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *Asset) GetFilesOk() (*[]*os.File, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *Asset) SetFiles(v []*os.File)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *Asset) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetCustomFields

`func (o *Asset) GetCustomFields() []CustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Asset) GetCustomFieldsOk() (*[]CustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Asset) SetCustomFields(v []CustomField)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Asset) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreatedById

`func (o *Asset) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *Asset) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *Asset) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *Asset) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Asset) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Asset) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Asset) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Asset) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetVersions

`func (o *Asset) GetVersions() []string`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *Asset) GetVersionsOk() (*[]string, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *Asset) SetVersions(v []string)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *Asset) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetDependencies

`func (o *Asset) GetDependencies() []Dependency`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *Asset) GetDependenciesOk() (*[]Dependency, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *Asset) SetDependencies(v []Dependency)`

SetDependencies sets Dependencies field to given value.

### HasDependencies

`func (o *Asset) HasDependencies() bool`

HasDependencies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


