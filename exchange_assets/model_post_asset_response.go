/*
Exchange Assets

Description of the Exchange Assets API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package exchange_assets

import (
	"encoding/json"
)

// checks if the PostAssetResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostAssetResponse{}

// PostAssetResponse struct for PostAssetResponse
type PostAssetResponse struct {
	OrganizationId *string `json:"organizationId,omitempty"`
	GroupId *string `json:"groupId,omitempty"`
	AssetId *string `json:"assetId,omitempty"`
	Version *string `json:"version,omitempty"`
	VersionGroup *string `json:"versionGroup,omitempty"`
	Classifier *string `json:"classifier,omitempty"`
	Metadata *Metadata `json:"metadata,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewPostAssetResponse instantiates a new PostAssetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostAssetResponse() *PostAssetResponse {
	this := PostAssetResponse{}
	return &this
}

// NewPostAssetResponseWithDefaults instantiates a new PostAssetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostAssetResponseWithDefaults() *PostAssetResponse {
	this := PostAssetResponse{}
	return &this
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *PostAssetResponse) GetOrganizationId() string {
	if o == nil || IsNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAssetResponse) GetOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *PostAssetResponse) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *PostAssetResponse) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *PostAssetResponse) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAssetResponse) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *PostAssetResponse) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *PostAssetResponse) SetGroupId(v string) {
	o.GroupId = &v
}

// GetAssetId returns the AssetId field value if set, zero value otherwise.
func (o *PostAssetResponse) GetAssetId() string {
	if o == nil || IsNil(o.AssetId) {
		var ret string
		return ret
	}
	return *o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAssetResponse) GetAssetIdOk() (*string, bool) {
	if o == nil || IsNil(o.AssetId) {
		return nil, false
	}
	return o.AssetId, true
}

// HasAssetId returns a boolean if a field has been set.
func (o *PostAssetResponse) HasAssetId() bool {
	if o != nil && !IsNil(o.AssetId) {
		return true
	}

	return false
}

// SetAssetId gets a reference to the given string and assigns it to the AssetId field.
func (o *PostAssetResponse) SetAssetId(v string) {
	o.AssetId = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *PostAssetResponse) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAssetResponse) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *PostAssetResponse) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *PostAssetResponse) SetVersion(v string) {
	o.Version = &v
}

// GetVersionGroup returns the VersionGroup field value if set, zero value otherwise.
func (o *PostAssetResponse) GetVersionGroup() string {
	if o == nil || IsNil(o.VersionGroup) {
		var ret string
		return ret
	}
	return *o.VersionGroup
}

// GetVersionGroupOk returns a tuple with the VersionGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAssetResponse) GetVersionGroupOk() (*string, bool) {
	if o == nil || IsNil(o.VersionGroup) {
		return nil, false
	}
	return o.VersionGroup, true
}

// HasVersionGroup returns a boolean if a field has been set.
func (o *PostAssetResponse) HasVersionGroup() bool {
	if o != nil && !IsNil(o.VersionGroup) {
		return true
	}

	return false
}

// SetVersionGroup gets a reference to the given string and assigns it to the VersionGroup field.
func (o *PostAssetResponse) SetVersionGroup(v string) {
	o.VersionGroup = &v
}

// GetClassifier returns the Classifier field value if set, zero value otherwise.
func (o *PostAssetResponse) GetClassifier() string {
	if o == nil || IsNil(o.Classifier) {
		var ret string
		return ret
	}
	return *o.Classifier
}

// GetClassifierOk returns a tuple with the Classifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAssetResponse) GetClassifierOk() (*string, bool) {
	if o == nil || IsNil(o.Classifier) {
		return nil, false
	}
	return o.Classifier, true
}

// HasClassifier returns a boolean if a field has been set.
func (o *PostAssetResponse) HasClassifier() bool {
	if o != nil && !IsNil(o.Classifier) {
		return true
	}

	return false
}

// SetClassifier gets a reference to the given string and assigns it to the Classifier field.
func (o *PostAssetResponse) SetClassifier(v string) {
	o.Classifier = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PostAssetResponse) GetMetadata() Metadata {
	if o == nil || IsNil(o.Metadata) {
		var ret Metadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAssetResponse) GetMetadataOk() (*Metadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PostAssetResponse) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given Metadata and assigns it to the Metadata field.
func (o *PostAssetResponse) SetMetadata(v Metadata) {
	o.Metadata = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PostAssetResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAssetResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PostAssetResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PostAssetResponse) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PostAssetResponse) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAssetResponse) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PostAssetResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PostAssetResponse) SetType(v string) {
	o.Type = &v
}

func (o PostAssetResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostAssetResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if !IsNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !IsNil(o.AssetId) {
		toSerialize["assetId"] = o.AssetId
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.VersionGroup) {
		toSerialize["versionGroup"] = o.VersionGroup
	}
	if !IsNil(o.Classifier) {
		toSerialize["classifier"] = o.Classifier
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullablePostAssetResponse struct {
	value *PostAssetResponse
	isSet bool
}

func (v NullablePostAssetResponse) Get() *PostAssetResponse {
	return v.value
}

func (v *NullablePostAssetResponse) Set(val *PostAssetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePostAssetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePostAssetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostAssetResponse(val *PostAssetResponse) *NullablePostAssetResponse {
	return &NullablePostAssetResponse{value: val, isSet: true}
}

func (v NullablePostAssetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostAssetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


