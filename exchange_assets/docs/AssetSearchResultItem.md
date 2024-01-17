# AssetSearchResultItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** |  | [optional] 
**AssetId** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**MinorVersion** | Pointer to **string** |  | [optional] 
**VersionGroup** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ContactName** | Pointer to **NullableString** |  | [optional] 
**ContactEmail** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**IsSnapshot** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ExternalFile** | Pointer to [**ExternalFile**](ExternalFile.md) |  | [optional] 
**CreatedDate** | Pointer to **string** |  | [optional] 
**UpdatedDate** | Pointer to **string** |  | [optional] 
**MinMuleVersion** | Pointer to **NullableString** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Categories** | Pointer to [**[]Category**](Category.md) |  | [optional] 
**Files** | Pointer to [**[]*os.File**](*os.File.md) |  | [optional] 
**CustomFields** | Pointer to [**[]CustomField**](CustomField.md) |  | [optional] 
**Rating** | Pointer to **int32** |  | [optional] 
**NumberOfRates** | Pointer to **int32** |  | [optional] 
**CreatedBy** | Pointer to [**CreatedBy**](CreatedBy.md) |  | [optional] 
**Organization** | Pointer to [**Organization**](Organization.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Icon** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**ModifiedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewAssetSearchResultItem

`func NewAssetSearchResultItem() *AssetSearchResultItem`

NewAssetSearchResultItem instantiates a new AssetSearchResultItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetSearchResultItemWithDefaults

`func NewAssetSearchResultItemWithDefaults() *AssetSearchResultItem`

NewAssetSearchResultItemWithDefaults instantiates a new AssetSearchResultItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *AssetSearchResultItem) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *AssetSearchResultItem) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *AssetSearchResultItem) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *AssetSearchResultItem) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetAssetId

`func (o *AssetSearchResultItem) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *AssetSearchResultItem) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *AssetSearchResultItem) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *AssetSearchResultItem) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetVersion

`func (o *AssetSearchResultItem) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AssetSearchResultItem) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AssetSearchResultItem) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AssetSearchResultItem) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetMinorVersion

`func (o *AssetSearchResultItem) GetMinorVersion() string`

GetMinorVersion returns the MinorVersion field if non-nil, zero value otherwise.

### GetMinorVersionOk

`func (o *AssetSearchResultItem) GetMinorVersionOk() (*string, bool)`

GetMinorVersionOk returns a tuple with the MinorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinorVersion

`func (o *AssetSearchResultItem) SetMinorVersion(v string)`

SetMinorVersion sets MinorVersion field to given value.

### HasMinorVersion

`func (o *AssetSearchResultItem) HasMinorVersion() bool`

HasMinorVersion returns a boolean if a field has been set.

### GetVersionGroup

`func (o *AssetSearchResultItem) GetVersionGroup() string`

GetVersionGroup returns the VersionGroup field if non-nil, zero value otherwise.

### GetVersionGroupOk

`func (o *AssetSearchResultItem) GetVersionGroupOk() (*string, bool)`

GetVersionGroupOk returns a tuple with the VersionGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionGroup

`func (o *AssetSearchResultItem) SetVersionGroup(v string)`

SetVersionGroup sets VersionGroup field to given value.

### HasVersionGroup

`func (o *AssetSearchResultItem) HasVersionGroup() bool`

HasVersionGroup returns a boolean if a field has been set.

### GetDescription

`func (o *AssetSearchResultItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AssetSearchResultItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AssetSearchResultItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AssetSearchResultItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsPublic

`func (o *AssetSearchResultItem) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *AssetSearchResultItem) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *AssetSearchResultItem) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *AssetSearchResultItem) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetName

`func (o *AssetSearchResultItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssetSearchResultItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssetSearchResultItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AssetSearchResultItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetContactName

`func (o *AssetSearchResultItem) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *AssetSearchResultItem) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *AssetSearchResultItem) SetContactName(v string)`

SetContactName sets ContactName field to given value.

### HasContactName

`func (o *AssetSearchResultItem) HasContactName() bool`

HasContactName returns a boolean if a field has been set.

### SetContactNameNil

`func (o *AssetSearchResultItem) SetContactNameNil(b bool)`

 SetContactNameNil sets the value for ContactName to be an explicit nil

### UnsetContactName
`func (o *AssetSearchResultItem) UnsetContactName()`

UnsetContactName ensures that no value is present for ContactName, not even an explicit nil
### GetContactEmail

`func (o *AssetSearchResultItem) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *AssetSearchResultItem) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *AssetSearchResultItem) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *AssetSearchResultItem) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### SetContactEmailNil

`func (o *AssetSearchResultItem) SetContactEmailNil(b bool)`

 SetContactEmailNil sets the value for ContactEmail to be an explicit nil

### UnsetContactEmail
`func (o *AssetSearchResultItem) UnsetContactEmail()`

UnsetContactEmail ensures that no value is present for ContactEmail, not even an explicit nil
### GetType

`func (o *AssetSearchResultItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AssetSearchResultItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AssetSearchResultItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AssetSearchResultItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIsSnapshot

`func (o *AssetSearchResultItem) GetIsSnapshot() bool`

GetIsSnapshot returns the IsSnapshot field if non-nil, zero value otherwise.

### GetIsSnapshotOk

`func (o *AssetSearchResultItem) GetIsSnapshotOk() (*bool, bool)`

GetIsSnapshotOk returns a tuple with the IsSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSnapshot

`func (o *AssetSearchResultItem) SetIsSnapshot(v bool)`

SetIsSnapshot sets IsSnapshot field to given value.

### HasIsSnapshot

`func (o *AssetSearchResultItem) HasIsSnapshot() bool`

HasIsSnapshot returns a boolean if a field has been set.

### GetStatus

`func (o *AssetSearchResultItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AssetSearchResultItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AssetSearchResultItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AssetSearchResultItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetExternalFile

`func (o *AssetSearchResultItem) GetExternalFile() ExternalFile`

GetExternalFile returns the ExternalFile field if non-nil, zero value otherwise.

### GetExternalFileOk

`func (o *AssetSearchResultItem) GetExternalFileOk() (*ExternalFile, bool)`

GetExternalFileOk returns a tuple with the ExternalFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalFile

`func (o *AssetSearchResultItem) SetExternalFile(v ExternalFile)`

SetExternalFile sets ExternalFile field to given value.

### HasExternalFile

`func (o *AssetSearchResultItem) HasExternalFile() bool`

HasExternalFile returns a boolean if a field has been set.

### GetCreatedDate

`func (o *AssetSearchResultItem) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *AssetSearchResultItem) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *AssetSearchResultItem) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *AssetSearchResultItem) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetUpdatedDate

`func (o *AssetSearchResultItem) GetUpdatedDate() string`

GetUpdatedDate returns the UpdatedDate field if non-nil, zero value otherwise.

### GetUpdatedDateOk

`func (o *AssetSearchResultItem) GetUpdatedDateOk() (*string, bool)`

GetUpdatedDateOk returns a tuple with the UpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDate

`func (o *AssetSearchResultItem) SetUpdatedDate(v string)`

SetUpdatedDate sets UpdatedDate field to given value.

### HasUpdatedDate

`func (o *AssetSearchResultItem) HasUpdatedDate() bool`

HasUpdatedDate returns a boolean if a field has been set.

### GetMinMuleVersion

`func (o *AssetSearchResultItem) GetMinMuleVersion() string`

GetMinMuleVersion returns the MinMuleVersion field if non-nil, zero value otherwise.

### GetMinMuleVersionOk

`func (o *AssetSearchResultItem) GetMinMuleVersionOk() (*string, bool)`

GetMinMuleVersionOk returns a tuple with the MinMuleVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinMuleVersion

`func (o *AssetSearchResultItem) SetMinMuleVersion(v string)`

SetMinMuleVersion sets MinMuleVersion field to given value.

### HasMinMuleVersion

`func (o *AssetSearchResultItem) HasMinMuleVersion() bool`

HasMinMuleVersion returns a boolean if a field has been set.

### SetMinMuleVersionNil

`func (o *AssetSearchResultItem) SetMinMuleVersionNil(b bool)`

 SetMinMuleVersionNil sets the value for MinMuleVersion to be an explicit nil

### UnsetMinMuleVersion
`func (o *AssetSearchResultItem) UnsetMinMuleVersion()`

UnsetMinMuleVersion ensures that no value is present for MinMuleVersion, not even an explicit nil
### GetLabels

`func (o *AssetSearchResultItem) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AssetSearchResultItem) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AssetSearchResultItem) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AssetSearchResultItem) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetCategories

`func (o *AssetSearchResultItem) GetCategories() []Category`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *AssetSearchResultItem) GetCategoriesOk() (*[]Category, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *AssetSearchResultItem) SetCategories(v []Category)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *AssetSearchResultItem) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetFiles

`func (o *AssetSearchResultItem) GetFiles() []*os.File`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *AssetSearchResultItem) GetFilesOk() (*[]*os.File, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *AssetSearchResultItem) SetFiles(v []*os.File)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *AssetSearchResultItem) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetCustomFields

`func (o *AssetSearchResultItem) GetCustomFields() []CustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *AssetSearchResultItem) GetCustomFieldsOk() (*[]CustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *AssetSearchResultItem) SetCustomFields(v []CustomField)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *AssetSearchResultItem) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRating

`func (o *AssetSearchResultItem) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *AssetSearchResultItem) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *AssetSearchResultItem) SetRating(v int32)`

SetRating sets Rating field to given value.

### HasRating

`func (o *AssetSearchResultItem) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetNumberOfRates

`func (o *AssetSearchResultItem) GetNumberOfRates() int32`

GetNumberOfRates returns the NumberOfRates field if non-nil, zero value otherwise.

### GetNumberOfRatesOk

`func (o *AssetSearchResultItem) GetNumberOfRatesOk() (*int32, bool)`

GetNumberOfRatesOk returns a tuple with the NumberOfRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfRates

`func (o *AssetSearchResultItem) SetNumberOfRates(v int32)`

SetNumberOfRates sets NumberOfRates field to given value.

### HasNumberOfRates

`func (o *AssetSearchResultItem) HasNumberOfRates() bool`

HasNumberOfRates returns a boolean if a field has been set.

### GetCreatedBy

`func (o *AssetSearchResultItem) GetCreatedBy() CreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AssetSearchResultItem) GetCreatedByOk() (*CreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AssetSearchResultItem) SetCreatedBy(v CreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AssetSearchResultItem) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetOrganization

`func (o *AssetSearchResultItem) GetOrganization() Organization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *AssetSearchResultItem) GetOrganizationOk() (*Organization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *AssetSearchResultItem) SetOrganization(v Organization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *AssetSearchResultItem) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetId

`func (o *AssetSearchResultItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssetSearchResultItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssetSearchResultItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AssetSearchResultItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIcon

`func (o *AssetSearchResultItem) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *AssetSearchResultItem) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *AssetSearchResultItem) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *AssetSearchResultItem) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *AssetSearchResultItem) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *AssetSearchResultItem) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetCreatedAt

`func (o *AssetSearchResultItem) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AssetSearchResultItem) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AssetSearchResultItem) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AssetSearchResultItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *AssetSearchResultItem) GetModifiedAt() string`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *AssetSearchResultItem) GetModifiedAtOk() (*string, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *AssetSearchResultItem) SetModifiedAt(v string)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *AssetSearchResultItem) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


