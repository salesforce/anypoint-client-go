/*
 * Flex Gateway API
 *
 * Description of the Flex Gateway API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package flexgateway

import (
	"encoding/json"
	"time"
)

// FlexGatewayTargetDetails struct for FlexGatewayTargetDetails
type FlexGatewayTargetDetails struct {
	OrganizationId *string `json:"organizationId,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Status *string `json:"status,omitempty"`
	Replicas *[]FlexGatewayTargetDetailsReplicas `json:"replicas,omitempty"`
	Tags *[]string `json:"tags,omitempty"`
	LastUpdate *time.Time `json:"lastUpdate,omitempty"`
	Versions *[]string `json:"versions,omitempty"`
	Version *string `json:"version,omitempty"`
}

// NewFlexGatewayTargetDetails instantiates a new FlexGatewayTargetDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlexGatewayTargetDetails() *FlexGatewayTargetDetails {
	this := FlexGatewayTargetDetails{}
	return &this
}

// NewFlexGatewayTargetDetailsWithDefaults instantiates a new FlexGatewayTargetDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlexGatewayTargetDetailsWithDefaults() *FlexGatewayTargetDetails {
	this := FlexGatewayTargetDetails{}
	return &this
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *FlexGatewayTargetDetails) GetOrganizationId() string {
	if o == nil || o.OrganizationId == nil {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexGatewayTargetDetails) GetOrganizationIdOk() (*string, bool) {
	if o == nil || o.OrganizationId == nil {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *FlexGatewayTargetDetails) HasOrganizationId() bool {
	if o != nil && o.OrganizationId != nil {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *FlexGatewayTargetDetails) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FlexGatewayTargetDetails) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexGatewayTargetDetails) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FlexGatewayTargetDetails) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FlexGatewayTargetDetails) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FlexGatewayTargetDetails) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexGatewayTargetDetails) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FlexGatewayTargetDetails) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FlexGatewayTargetDetails) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FlexGatewayTargetDetails) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexGatewayTargetDetails) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FlexGatewayTargetDetails) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *FlexGatewayTargetDetails) SetStatus(v string) {
	o.Status = &v
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *FlexGatewayTargetDetails) GetReplicas() []FlexGatewayTargetDetailsReplicas {
	if o == nil || o.Replicas == nil {
		var ret []FlexGatewayTargetDetailsReplicas
		return ret
	}
	return *o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexGatewayTargetDetails) GetReplicasOk() (*[]FlexGatewayTargetDetailsReplicas, bool) {
	if o == nil || o.Replicas == nil {
		return nil, false
	}
	return o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *FlexGatewayTargetDetails) HasReplicas() bool {
	if o != nil && o.Replicas != nil {
		return true
	}

	return false
}

// SetReplicas gets a reference to the given []FlexGatewayTargetDetailsReplicas and assigns it to the Replicas field.
func (o *FlexGatewayTargetDetails) SetReplicas(v []FlexGatewayTargetDetailsReplicas) {
	o.Replicas = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FlexGatewayTargetDetails) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexGatewayTargetDetails) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FlexGatewayTargetDetails) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *FlexGatewayTargetDetails) SetTags(v []string) {
	o.Tags = &v
}

// GetLastUpdate returns the LastUpdate field value if set, zero value otherwise.
func (o *FlexGatewayTargetDetails) GetLastUpdate() time.Time {
	if o == nil || o.LastUpdate == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdate
}

// GetLastUpdateOk returns a tuple with the LastUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexGatewayTargetDetails) GetLastUpdateOk() (*time.Time, bool) {
	if o == nil || o.LastUpdate == nil {
		return nil, false
	}
	return o.LastUpdate, true
}

// HasLastUpdate returns a boolean if a field has been set.
func (o *FlexGatewayTargetDetails) HasLastUpdate() bool {
	if o != nil && o.LastUpdate != nil {
		return true
	}

	return false
}

// SetLastUpdate gets a reference to the given time.Time and assigns it to the LastUpdate field.
func (o *FlexGatewayTargetDetails) SetLastUpdate(v time.Time) {
	o.LastUpdate = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *FlexGatewayTargetDetails) GetVersions() []string {
	if o == nil || o.Versions == nil {
		var ret []string
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexGatewayTargetDetails) GetVersionsOk() (*[]string, bool) {
	if o == nil || o.Versions == nil {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *FlexGatewayTargetDetails) HasVersions() bool {
	if o != nil && o.Versions != nil {
		return true
	}

	return false
}

// SetVersions gets a reference to the given []string and assigns it to the Versions field.
func (o *FlexGatewayTargetDetails) SetVersions(v []string) {
	o.Versions = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *FlexGatewayTargetDetails) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexGatewayTargetDetails) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *FlexGatewayTargetDetails) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *FlexGatewayTargetDetails) SetVersion(v string) {
	o.Version = &v
}

func (o FlexGatewayTargetDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OrganizationId != nil {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Replicas != nil {
		toSerialize["replicas"] = o.Replicas
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.LastUpdate != nil {
		toSerialize["lastUpdate"] = o.LastUpdate
	}
	if o.Versions != nil {
		toSerialize["versions"] = o.Versions
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableFlexGatewayTargetDetails struct {
	value *FlexGatewayTargetDetails
	isSet bool
}

func (v NullableFlexGatewayTargetDetails) Get() *FlexGatewayTargetDetails {
	return v.value
}

func (v *NullableFlexGatewayTargetDetails) Set(val *FlexGatewayTargetDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableFlexGatewayTargetDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableFlexGatewayTargetDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlexGatewayTargetDetails(val *FlexGatewayTargetDetails) *NullableFlexGatewayTargetDetails {
	return &NullableFlexGatewayTargetDetails{value: val, isSet: true}
}

func (v NullableFlexGatewayTargetDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlexGatewayTargetDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


