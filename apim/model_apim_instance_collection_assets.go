/*
 * API Manager API
 *
 * API Manager API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apim

import (
	"encoding/json"
)

// ApimInstanceCollectionAssets struct for ApimInstanceCollectionAssets
type ApimInstanceCollectionAssets struct {
	Audit *Audit `json:"audit,omitempty"`
	MasterOrganizationId *string `json:"masterOrganizationId,omitempty"`
	OrganizationId *string `json:"organizationId,omitempty"`
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	ExchangeAssetName *string `json:"exchangeAssetName,omitempty"`
	GroupId *string `json:"groupId,omitempty"`
	AssetId *string `json:"assetId,omitempty"`
	Apis *[]ApimInstanceCollectionApis `json:"apis,omitempty"`
	TotalApis *int32 `json:"totalApis,omitempty"`
	AutodiscoveryApiName *string `json:"autodiscoveryApiName,omitempty"`
}

// NewApimInstanceCollectionAssets instantiates a new ApimInstanceCollectionAssets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApimInstanceCollectionAssets() *ApimInstanceCollectionAssets {
	this := ApimInstanceCollectionAssets{}
	return &this
}

// NewApimInstanceCollectionAssetsWithDefaults instantiates a new ApimInstanceCollectionAssets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApimInstanceCollectionAssetsWithDefaults() *ApimInstanceCollectionAssets {
	this := ApimInstanceCollectionAssets{}
	return &this
}

// GetAudit returns the Audit field value if set, zero value otherwise.
func (o *ApimInstanceCollectionAssets) GetAudit() Audit {
	if o == nil || o.Audit == nil {
		var ret Audit
		return ret
	}
	return *o.Audit
}

// GetAuditOk returns a tuple with the Audit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstanceCollectionAssets) GetAuditOk() (*Audit, bool) {
	if o == nil || o.Audit == nil {
		return nil, false
	}
	return o.Audit, true
}

// HasAudit returns a boolean if a field has been set.
func (o *ApimInstanceCollectionAssets) HasAudit() bool {
	if o != nil && o.Audit != nil {
		return true
	}

	return false
}

// SetAudit gets a reference to the given Audit and assigns it to the Audit field.
func (o *ApimInstanceCollectionAssets) SetAudit(v Audit) {
	o.Audit = &v
}

// GetMasterOrganizationId returns the MasterOrganizationId field value if set, zero value otherwise.
func (o *ApimInstanceCollectionAssets) GetMasterOrganizationId() string {
	if o == nil || o.MasterOrganizationId == nil {
		var ret string
		return ret
	}
	return *o.MasterOrganizationId
}

// GetMasterOrganizationIdOk returns a tuple with the MasterOrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstanceCollectionAssets) GetMasterOrganizationIdOk() (*string, bool) {
	if o == nil || o.MasterOrganizationId == nil {
		return nil, false
	}
	return o.MasterOrganizationId, true
}

// HasMasterOrganizationId returns a boolean if a field has been set.
func (o *ApimInstanceCollectionAssets) HasMasterOrganizationId() bool {
	if o != nil && o.MasterOrganizationId != nil {
		return true
	}

	return false
}

// SetMasterOrganizationId gets a reference to the given string and assigns it to the MasterOrganizationId field.
func (o *ApimInstanceCollectionAssets) SetMasterOrganizationId(v string) {
	o.MasterOrganizationId = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *ApimInstanceCollectionAssets) GetOrganizationId() string {
	if o == nil || o.OrganizationId == nil {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstanceCollectionAssets) GetOrganizationIdOk() (*string, bool) {
	if o == nil || o.OrganizationId == nil {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *ApimInstanceCollectionAssets) HasOrganizationId() bool {
	if o != nil && o.OrganizationId != nil {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *ApimInstanceCollectionAssets) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApimInstanceCollectionAssets) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstanceCollectionAssets) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApimInstanceCollectionAssets) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ApimInstanceCollectionAssets) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApimInstanceCollectionAssets) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstanceCollectionAssets) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApimInstanceCollectionAssets) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApimInstanceCollectionAssets) SetName(v string) {
	o.Name = &v
}

// GetExchangeAssetName returns the ExchangeAssetName field value if set, zero value otherwise.
func (o *ApimInstanceCollectionAssets) GetExchangeAssetName() string {
	if o == nil || o.ExchangeAssetName == nil {
		var ret string
		return ret
	}
	return *o.ExchangeAssetName
}

// GetExchangeAssetNameOk returns a tuple with the ExchangeAssetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstanceCollectionAssets) GetExchangeAssetNameOk() (*string, bool) {
	if o == nil || o.ExchangeAssetName == nil {
		return nil, false
	}
	return o.ExchangeAssetName, true
}

// HasExchangeAssetName returns a boolean if a field has been set.
func (o *ApimInstanceCollectionAssets) HasExchangeAssetName() bool {
	if o != nil && o.ExchangeAssetName != nil {
		return true
	}

	return false
}

// SetExchangeAssetName gets a reference to the given string and assigns it to the ExchangeAssetName field.
func (o *ApimInstanceCollectionAssets) SetExchangeAssetName(v string) {
	o.ExchangeAssetName = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *ApimInstanceCollectionAssets) GetGroupId() string {
	if o == nil || o.GroupId == nil {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstanceCollectionAssets) GetGroupIdOk() (*string, bool) {
	if o == nil || o.GroupId == nil {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *ApimInstanceCollectionAssets) HasGroupId() bool {
	if o != nil && o.GroupId != nil {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *ApimInstanceCollectionAssets) SetGroupId(v string) {
	o.GroupId = &v
}

// GetAssetId returns the AssetId field value if set, zero value otherwise.
func (o *ApimInstanceCollectionAssets) GetAssetId() string {
	if o == nil || o.AssetId == nil {
		var ret string
		return ret
	}
	return *o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstanceCollectionAssets) GetAssetIdOk() (*string, bool) {
	if o == nil || o.AssetId == nil {
		return nil, false
	}
	return o.AssetId, true
}

// HasAssetId returns a boolean if a field has been set.
func (o *ApimInstanceCollectionAssets) HasAssetId() bool {
	if o != nil && o.AssetId != nil {
		return true
	}

	return false
}

// SetAssetId gets a reference to the given string and assigns it to the AssetId field.
func (o *ApimInstanceCollectionAssets) SetAssetId(v string) {
	o.AssetId = &v
}

// GetApis returns the Apis field value if set, zero value otherwise.
func (o *ApimInstanceCollectionAssets) GetApis() []ApimInstanceCollectionApis {
	if o == nil || o.Apis == nil {
		var ret []ApimInstanceCollectionApis
		return ret
	}
	return *o.Apis
}

// GetApisOk returns a tuple with the Apis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstanceCollectionAssets) GetApisOk() (*[]ApimInstanceCollectionApis, bool) {
	if o == nil || o.Apis == nil {
		return nil, false
	}
	return o.Apis, true
}

// HasApis returns a boolean if a field has been set.
func (o *ApimInstanceCollectionAssets) HasApis() bool {
	if o != nil && o.Apis != nil {
		return true
	}

	return false
}

// SetApis gets a reference to the given []ApimInstanceCollectionApis and assigns it to the Apis field.
func (o *ApimInstanceCollectionAssets) SetApis(v []ApimInstanceCollectionApis) {
	o.Apis = &v
}

// GetTotalApis returns the TotalApis field value if set, zero value otherwise.
func (o *ApimInstanceCollectionAssets) GetTotalApis() int32 {
	if o == nil || o.TotalApis == nil {
		var ret int32
		return ret
	}
	return *o.TotalApis
}

// GetTotalApisOk returns a tuple with the TotalApis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstanceCollectionAssets) GetTotalApisOk() (*int32, bool) {
	if o == nil || o.TotalApis == nil {
		return nil, false
	}
	return o.TotalApis, true
}

// HasTotalApis returns a boolean if a field has been set.
func (o *ApimInstanceCollectionAssets) HasTotalApis() bool {
	if o != nil && o.TotalApis != nil {
		return true
	}

	return false
}

// SetTotalApis gets a reference to the given int32 and assigns it to the TotalApis field.
func (o *ApimInstanceCollectionAssets) SetTotalApis(v int32) {
	o.TotalApis = &v
}

// GetAutodiscoveryApiName returns the AutodiscoveryApiName field value if set, zero value otherwise.
func (o *ApimInstanceCollectionAssets) GetAutodiscoveryApiName() string {
	if o == nil || o.AutodiscoveryApiName == nil {
		var ret string
		return ret
	}
	return *o.AutodiscoveryApiName
}

// GetAutodiscoveryApiNameOk returns a tuple with the AutodiscoveryApiName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApimInstanceCollectionAssets) GetAutodiscoveryApiNameOk() (*string, bool) {
	if o == nil || o.AutodiscoveryApiName == nil {
		return nil, false
	}
	return o.AutodiscoveryApiName, true
}

// HasAutodiscoveryApiName returns a boolean if a field has been set.
func (o *ApimInstanceCollectionAssets) HasAutodiscoveryApiName() bool {
	if o != nil && o.AutodiscoveryApiName != nil {
		return true
	}

	return false
}

// SetAutodiscoveryApiName gets a reference to the given string and assigns it to the AutodiscoveryApiName field.
func (o *ApimInstanceCollectionAssets) SetAutodiscoveryApiName(v string) {
	o.AutodiscoveryApiName = &v
}

func (o ApimInstanceCollectionAssets) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Audit != nil {
		toSerialize["audit"] = o.Audit
	}
	if o.MasterOrganizationId != nil {
		toSerialize["masterOrganizationId"] = o.MasterOrganizationId
	}
	if o.OrganizationId != nil {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ExchangeAssetName != nil {
		toSerialize["exchangeAssetName"] = o.ExchangeAssetName
	}
	if o.GroupId != nil {
		toSerialize["groupId"] = o.GroupId
	}
	if o.AssetId != nil {
		toSerialize["assetId"] = o.AssetId
	}
	if o.Apis != nil {
		toSerialize["apis"] = o.Apis
	}
	if o.TotalApis != nil {
		toSerialize["totalApis"] = o.TotalApis
	}
	if o.AutodiscoveryApiName != nil {
		toSerialize["autodiscoveryApiName"] = o.AutodiscoveryApiName
	}
	return json.Marshal(toSerialize)
}

type NullableApimInstanceCollectionAssets struct {
	value *ApimInstanceCollectionAssets
	isSet bool
}

func (v NullableApimInstanceCollectionAssets) Get() *ApimInstanceCollectionAssets {
	return v.value
}

func (v *NullableApimInstanceCollectionAssets) Set(val *ApimInstanceCollectionAssets) {
	v.value = val
	v.isSet = true
}

func (v NullableApimInstanceCollectionAssets) IsSet() bool {
	return v.isSet
}

func (v *NullableApimInstanceCollectionAssets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApimInstanceCollectionAssets(val *ApimInstanceCollectionAssets) *NullableApimInstanceCollectionAssets {
	return &NullableApimInstanceCollectionAssets{value: val, isSet: true}
}

func (v NullableApimInstanceCollectionAssets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApimInstanceCollectionAssets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


