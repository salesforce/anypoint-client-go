/*
Exchange Assets

Description of the Exchange Assets API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package exchange_assets

import (
	"encoding/json"
	"os"
)

// checks if the Asset type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Asset{}

// Asset struct for Asset
type Asset struct {
	GroupId *string `json:"groupId,omitempty"`
	AssetId *string `json:"assetId,omitempty"`
	Version *string `json:"version,omitempty"`
	MinorVersion *string `json:"minorVersion,omitempty"`
	VersionGroup *string `json:"versionGroup,omitempty"`
	Description *string `json:"description,omitempty"`
	OrganizationId *string `json:"organizationId,omitempty"`
	IsPublic *bool `json:"isPublic,omitempty"`
	Name *string `json:"name,omitempty"`
	ContactName NullableString `json:"contactName,omitempty"`
	ContactEmail NullableString `json:"contactEmail,omitempty"`
	Type *string `json:"type,omitempty"`
	IsSnapshot *bool `json:"isSnapshot,omitempty"`
	Status *string `json:"status,omitempty"`
	Labels []string `json:"labels,omitempty"`
	Attributes []Attribute `json:"attributes,omitempty"`
	Categories []Category `json:"categories,omitempty"`
	Files []*os.File `json:"files,omitempty"`
	CustomFields []CustomField `json:"customFields,omitempty"`
	CreatedById *string `json:"createdById,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
	Versions []string `json:"versions,omitempty"`
	Dependencies []Dependency `json:"dependencies,omitempty"`
}

// NewAsset instantiates a new Asset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAsset() *Asset {
	this := Asset{}
	return &this
}

// NewAssetWithDefaults instantiates a new Asset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetWithDefaults() *Asset {
	this := Asset{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *Asset) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *Asset) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *Asset) SetGroupId(v string) {
	o.GroupId = &v
}

// GetAssetId returns the AssetId field value if set, zero value otherwise.
func (o *Asset) GetAssetId() string {
	if o == nil || IsNil(o.AssetId) {
		var ret string
		return ret
	}
	return *o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetAssetIdOk() (*string, bool) {
	if o == nil || IsNil(o.AssetId) {
		return nil, false
	}
	return o.AssetId, true
}

// HasAssetId returns a boolean if a field has been set.
func (o *Asset) HasAssetId() bool {
	if o != nil && !IsNil(o.AssetId) {
		return true
	}

	return false
}

// SetAssetId gets a reference to the given string and assigns it to the AssetId field.
func (o *Asset) SetAssetId(v string) {
	o.AssetId = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Asset) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Asset) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Asset) SetVersion(v string) {
	o.Version = &v
}

// GetMinorVersion returns the MinorVersion field value if set, zero value otherwise.
func (o *Asset) GetMinorVersion() string {
	if o == nil || IsNil(o.MinorVersion) {
		var ret string
		return ret
	}
	return *o.MinorVersion
}

// GetMinorVersionOk returns a tuple with the MinorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetMinorVersionOk() (*string, bool) {
	if o == nil || IsNil(o.MinorVersion) {
		return nil, false
	}
	return o.MinorVersion, true
}

// HasMinorVersion returns a boolean if a field has been set.
func (o *Asset) HasMinorVersion() bool {
	if o != nil && !IsNil(o.MinorVersion) {
		return true
	}

	return false
}

// SetMinorVersion gets a reference to the given string and assigns it to the MinorVersion field.
func (o *Asset) SetMinorVersion(v string) {
	o.MinorVersion = &v
}

// GetVersionGroup returns the VersionGroup field value if set, zero value otherwise.
func (o *Asset) GetVersionGroup() string {
	if o == nil || IsNil(o.VersionGroup) {
		var ret string
		return ret
	}
	return *o.VersionGroup
}

// GetVersionGroupOk returns a tuple with the VersionGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetVersionGroupOk() (*string, bool) {
	if o == nil || IsNil(o.VersionGroup) {
		return nil, false
	}
	return o.VersionGroup, true
}

// HasVersionGroup returns a boolean if a field has been set.
func (o *Asset) HasVersionGroup() bool {
	if o != nil && !IsNil(o.VersionGroup) {
		return true
	}

	return false
}

// SetVersionGroup gets a reference to the given string and assigns it to the VersionGroup field.
func (o *Asset) SetVersionGroup(v string) {
	o.VersionGroup = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Asset) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Asset) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Asset) SetDescription(v string) {
	o.Description = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *Asset) GetOrganizationId() string {
	if o == nil || IsNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *Asset) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *Asset) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *Asset) GetIsPublic() bool {
	if o == nil || IsNil(o.IsPublic) {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetIsPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPublic) {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *Asset) HasIsPublic() bool {
	if o != nil && !IsNil(o.IsPublic) {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *Asset) SetIsPublic(v bool) {
	o.IsPublic = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Asset) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Asset) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Asset) SetName(v string) {
	o.Name = &v
}

// GetContactName returns the ContactName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Asset) GetContactName() string {
	if o == nil || IsNil(o.ContactName.Get()) {
		var ret string
		return ret
	}
	return *o.ContactName.Get()
}

// GetContactNameOk returns a tuple with the ContactName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Asset) GetContactNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContactName.Get(), o.ContactName.IsSet()
}

// HasContactName returns a boolean if a field has been set.
func (o *Asset) HasContactName() bool {
	if o != nil && o.ContactName.IsSet() {
		return true
	}

	return false
}

// SetContactName gets a reference to the given NullableString and assigns it to the ContactName field.
func (o *Asset) SetContactName(v string) {
	o.ContactName.Set(&v)
}
// SetContactNameNil sets the value for ContactName to be an explicit nil
func (o *Asset) SetContactNameNil() {
	o.ContactName.Set(nil)
}

// UnsetContactName ensures that no value is present for ContactName, not even an explicit nil
func (o *Asset) UnsetContactName() {
	o.ContactName.Unset()
}

// GetContactEmail returns the ContactEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Asset) GetContactEmail() string {
	if o == nil || IsNil(o.ContactEmail.Get()) {
		var ret string
		return ret
	}
	return *o.ContactEmail.Get()
}

// GetContactEmailOk returns a tuple with the ContactEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Asset) GetContactEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContactEmail.Get(), o.ContactEmail.IsSet()
}

// HasContactEmail returns a boolean if a field has been set.
func (o *Asset) HasContactEmail() bool {
	if o != nil && o.ContactEmail.IsSet() {
		return true
	}

	return false
}

// SetContactEmail gets a reference to the given NullableString and assigns it to the ContactEmail field.
func (o *Asset) SetContactEmail(v string) {
	o.ContactEmail.Set(&v)
}
// SetContactEmailNil sets the value for ContactEmail to be an explicit nil
func (o *Asset) SetContactEmailNil() {
	o.ContactEmail.Set(nil)
}

// UnsetContactEmail ensures that no value is present for ContactEmail, not even an explicit nil
func (o *Asset) UnsetContactEmail() {
	o.ContactEmail.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Asset) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Asset) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Asset) SetType(v string) {
	o.Type = &v
}

// GetIsSnapshot returns the IsSnapshot field value if set, zero value otherwise.
func (o *Asset) GetIsSnapshot() bool {
	if o == nil || IsNil(o.IsSnapshot) {
		var ret bool
		return ret
	}
	return *o.IsSnapshot
}

// GetIsSnapshotOk returns a tuple with the IsSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetIsSnapshotOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSnapshot) {
		return nil, false
	}
	return o.IsSnapshot, true
}

// HasIsSnapshot returns a boolean if a field has been set.
func (o *Asset) HasIsSnapshot() bool {
	if o != nil && !IsNil(o.IsSnapshot) {
		return true
	}

	return false
}

// SetIsSnapshot gets a reference to the given bool and assigns it to the IsSnapshot field.
func (o *Asset) SetIsSnapshot(v bool) {
	o.IsSnapshot = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Asset) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Asset) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Asset) SetStatus(v string) {
	o.Status = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *Asset) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *Asset) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *Asset) SetLabels(v []string) {
	o.Labels = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *Asset) GetAttributes() []Attribute {
	if o == nil || IsNil(o.Attributes) {
		var ret []Attribute
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetAttributesOk() ([]Attribute, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Asset) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []Attribute and assigns it to the Attributes field.
func (o *Asset) SetAttributes(v []Attribute) {
	o.Attributes = v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *Asset) GetCategories() []Category {
	if o == nil || IsNil(o.Categories) {
		var ret []Category
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetCategoriesOk() ([]Category, bool) {
	if o == nil || IsNil(o.Categories) {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *Asset) HasCategories() bool {
	if o != nil && !IsNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []Category and assigns it to the Categories field.
func (o *Asset) SetCategories(v []Category) {
	o.Categories = v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *Asset) GetFiles() []*os.File {
	if o == nil || IsNil(o.Files) {
		var ret []*os.File
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetFilesOk() ([]*os.File, bool) {
	if o == nil || IsNil(o.Files) {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *Asset) HasFiles() bool {
	if o != nil && !IsNil(o.Files) {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []*os.File and assigns it to the Files field.
func (o *Asset) SetFiles(v []*os.File) {
	o.Files = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *Asset) GetCustomFields() []CustomField {
	if o == nil || IsNil(o.CustomFields) {
		var ret []CustomField
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetCustomFieldsOk() ([]CustomField, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *Asset) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given []CustomField and assigns it to the CustomFields field.
func (o *Asset) SetCustomFields(v []CustomField) {
	o.CustomFields = v
}

// GetCreatedById returns the CreatedById field value if set, zero value otherwise.
func (o *Asset) GetCreatedById() string {
	if o == nil || IsNil(o.CreatedById) {
		var ret string
		return ret
	}
	return *o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetCreatedByIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedById) {
		return nil, false
	}
	return o.CreatedById, true
}

// HasCreatedById returns a boolean if a field has been set.
func (o *Asset) HasCreatedById() bool {
	if o != nil && !IsNil(o.CreatedById) {
		return true
	}

	return false
}

// SetCreatedById gets a reference to the given string and assigns it to the CreatedById field.
func (o *Asset) SetCreatedById(v string) {
	o.CreatedById = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Asset) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Asset) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *Asset) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *Asset) GetVersions() []string {
	if o == nil || IsNil(o.Versions) {
		var ret []string
		return ret
	}
	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetVersionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *Asset) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given []string and assigns it to the Versions field.
func (o *Asset) SetVersions(v []string) {
	o.Versions = v
}

// GetDependencies returns the Dependencies field value if set, zero value otherwise.
func (o *Asset) GetDependencies() []Dependency {
	if o == nil || IsNil(o.Dependencies) {
		var ret []Dependency
		return ret
	}
	return o.Dependencies
}

// GetDependenciesOk returns a tuple with the Dependencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetDependenciesOk() ([]Dependency, bool) {
	if o == nil || IsNil(o.Dependencies) {
		return nil, false
	}
	return o.Dependencies, true
}

// HasDependencies returns a boolean if a field has been set.
func (o *Asset) HasDependencies() bool {
	if o != nil && !IsNil(o.Dependencies) {
		return true
	}

	return false
}

// SetDependencies gets a reference to the given []Dependency and assigns it to the Dependencies field.
func (o *Asset) SetDependencies(v []Dependency) {
	o.Dependencies = v
}

func (o Asset) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Asset) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !IsNil(o.AssetId) {
		toSerialize["assetId"] = o.AssetId
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.MinorVersion) {
		toSerialize["minorVersion"] = o.MinorVersion
	}
	if !IsNil(o.VersionGroup) {
		toSerialize["versionGroup"] = o.VersionGroup
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if !IsNil(o.IsPublic) {
		toSerialize["isPublic"] = o.IsPublic
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.ContactName.IsSet() {
		toSerialize["contactName"] = o.ContactName.Get()
	}
	if o.ContactEmail.IsSet() {
		toSerialize["contactEmail"] = o.ContactEmail.Get()
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.IsSnapshot) {
		toSerialize["isSnapshot"] = o.IsSnapshot
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Categories) {
		toSerialize["categories"] = o.Categories
	}
	if !IsNil(o.Files) {
		toSerialize["files"] = o.Files
	}
	if !IsNil(o.CustomFields) {
		toSerialize["customFields"] = o.CustomFields
	}
	if !IsNil(o.CreatedById) {
		toSerialize["createdById"] = o.CreatedById
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	if !IsNil(o.Dependencies) {
		toSerialize["dependencies"] = o.Dependencies
	}
	return toSerialize, nil
}

type NullableAsset struct {
	value *Asset
	isSet bool
}

func (v NullableAsset) Get() *Asset {
	return v.value
}

func (v *NullableAsset) Set(val *Asset) {
	v.value = val
	v.isSet = true
}

func (v NullableAsset) IsSet() bool {
	return v.isSet
}

func (v *NullableAsset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsset(val *Asset) *NullableAsset {
	return &NullableAsset{value: val, isSet: true}
}

func (v NullableAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAsset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


