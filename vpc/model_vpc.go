/*
VPC API

Description of the VPC API

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vpc

import (
	"encoding/json"
)

// checks if the Vpc type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Vpc{}

// Vpc struct for Vpc
type Vpc struct {
	// The vpc Id
	Id string `json:"id"`
	IsDefault *bool `json:"isDefault,omitempty"`
	Name *string `json:"name,omitempty"`
	OwnerId *string `json:"ownerId,omitempty"`
	Region *string `json:"region,omitempty"`
	SharedWith []string `json:"sharedWith,omitempty"`
	AssociatedEnvironments []string `json:"associatedEnvironments,omitempty"`
	// An explanation about the purpose of this instance.
	CidrBlock *string `json:"cidrBlock,omitempty"`
	FirewallRules []FirewallRule `json:"firewallRules,omitempty"`
	InternalDns *InternalDns `json:"internalDns,omitempty"`
	VpcRoutes []VpcRoute `json:"vpcRoutes,omitempty"`
}

// NewVpc instantiates a new Vpc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVpc(id string) *Vpc {
	this := Vpc{}
	this.Id = id
	var isDefault bool = false
	this.IsDefault = &isDefault
	var cidrBlock string = "10.0.0.0/20"
	this.CidrBlock = &cidrBlock
	var internalDns InternalDns = {"dnsServers":[],"specialDomains":[]}
	this.InternalDns = &internalDns
	return &this
}

// NewVpcWithDefaults instantiates a new Vpc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVpcWithDefaults() *Vpc {
	this := Vpc{}
	var isDefault bool = false
	this.IsDefault = &isDefault
	var cidrBlock string = "10.0.0.0/20"
	this.CidrBlock = &cidrBlock
	var internalDns InternalDns = {"dnsServers":[],"specialDomains":[]}
	this.InternalDns = &internalDns
	return &this
}

// GetId returns the Id field value
func (o *Vpc) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Vpc) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Vpc) SetId(v string) {
	o.Id = v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *Vpc) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vpc) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *Vpc) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *Vpc) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Vpc) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vpc) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Vpc) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Vpc) SetName(v string) {
	o.Name = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *Vpc) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vpc) GetOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *Vpc) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *Vpc) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *Vpc) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vpc) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *Vpc) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *Vpc) SetRegion(v string) {
	o.Region = &v
}

// GetSharedWith returns the SharedWith field value if set, zero value otherwise.
func (o *Vpc) GetSharedWith() []string {
	if o == nil || IsNil(o.SharedWith) {
		var ret []string
		return ret
	}
	return o.SharedWith
}

// GetSharedWithOk returns a tuple with the SharedWith field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vpc) GetSharedWithOk() ([]string, bool) {
	if o == nil || IsNil(o.SharedWith) {
		return nil, false
	}
	return o.SharedWith, true
}

// HasSharedWith returns a boolean if a field has been set.
func (o *Vpc) HasSharedWith() bool {
	if o != nil && !IsNil(o.SharedWith) {
		return true
	}

	return false
}

// SetSharedWith gets a reference to the given []string and assigns it to the SharedWith field.
func (o *Vpc) SetSharedWith(v []string) {
	o.SharedWith = v
}

// GetAssociatedEnvironments returns the AssociatedEnvironments field value if set, zero value otherwise.
func (o *Vpc) GetAssociatedEnvironments() []string {
	if o == nil || IsNil(o.AssociatedEnvironments) {
		var ret []string
		return ret
	}
	return o.AssociatedEnvironments
}

// GetAssociatedEnvironmentsOk returns a tuple with the AssociatedEnvironments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vpc) GetAssociatedEnvironmentsOk() ([]string, bool) {
	if o == nil || IsNil(o.AssociatedEnvironments) {
		return nil, false
	}
	return o.AssociatedEnvironments, true
}

// HasAssociatedEnvironments returns a boolean if a field has been set.
func (o *Vpc) HasAssociatedEnvironments() bool {
	if o != nil && !IsNil(o.AssociatedEnvironments) {
		return true
	}

	return false
}

// SetAssociatedEnvironments gets a reference to the given []string and assigns it to the AssociatedEnvironments field.
func (o *Vpc) SetAssociatedEnvironments(v []string) {
	o.AssociatedEnvironments = v
}

// GetCidrBlock returns the CidrBlock field value if set, zero value otherwise.
func (o *Vpc) GetCidrBlock() string {
	if o == nil || IsNil(o.CidrBlock) {
		var ret string
		return ret
	}
	return *o.CidrBlock
}

// GetCidrBlockOk returns a tuple with the CidrBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vpc) GetCidrBlockOk() (*string, bool) {
	if o == nil || IsNil(o.CidrBlock) {
		return nil, false
	}
	return o.CidrBlock, true
}

// HasCidrBlock returns a boolean if a field has been set.
func (o *Vpc) HasCidrBlock() bool {
	if o != nil && !IsNil(o.CidrBlock) {
		return true
	}

	return false
}

// SetCidrBlock gets a reference to the given string and assigns it to the CidrBlock field.
func (o *Vpc) SetCidrBlock(v string) {
	o.CidrBlock = &v
}

// GetFirewallRules returns the FirewallRules field value if set, zero value otherwise.
func (o *Vpc) GetFirewallRules() []FirewallRule {
	if o == nil || IsNil(o.FirewallRules) {
		var ret []FirewallRule
		return ret
	}
	return o.FirewallRules
}

// GetFirewallRulesOk returns a tuple with the FirewallRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vpc) GetFirewallRulesOk() ([]FirewallRule, bool) {
	if o == nil || IsNil(o.FirewallRules) {
		return nil, false
	}
	return o.FirewallRules, true
}

// HasFirewallRules returns a boolean if a field has been set.
func (o *Vpc) HasFirewallRules() bool {
	if o != nil && !IsNil(o.FirewallRules) {
		return true
	}

	return false
}

// SetFirewallRules gets a reference to the given []FirewallRule and assigns it to the FirewallRules field.
func (o *Vpc) SetFirewallRules(v []FirewallRule) {
	o.FirewallRules = v
}

// GetInternalDns returns the InternalDns field value if set, zero value otherwise.
func (o *Vpc) GetInternalDns() InternalDns {
	if o == nil || IsNil(o.InternalDns) {
		var ret InternalDns
		return ret
	}
	return *o.InternalDns
}

// GetInternalDnsOk returns a tuple with the InternalDns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vpc) GetInternalDnsOk() (*InternalDns, bool) {
	if o == nil || IsNil(o.InternalDns) {
		return nil, false
	}
	return o.InternalDns, true
}

// HasInternalDns returns a boolean if a field has been set.
func (o *Vpc) HasInternalDns() bool {
	if o != nil && !IsNil(o.InternalDns) {
		return true
	}

	return false
}

// SetInternalDns gets a reference to the given InternalDns and assigns it to the InternalDns field.
func (o *Vpc) SetInternalDns(v InternalDns) {
	o.InternalDns = &v
}

// GetVpcRoutes returns the VpcRoutes field value if set, zero value otherwise.
func (o *Vpc) GetVpcRoutes() []VpcRoute {
	if o == nil || IsNil(o.VpcRoutes) {
		var ret []VpcRoute
		return ret
	}
	return o.VpcRoutes
}

// GetVpcRoutesOk returns a tuple with the VpcRoutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vpc) GetVpcRoutesOk() ([]VpcRoute, bool) {
	if o == nil || IsNil(o.VpcRoutes) {
		return nil, false
	}
	return o.VpcRoutes, true
}

// HasVpcRoutes returns a boolean if a field has been set.
func (o *Vpc) HasVpcRoutes() bool {
	if o != nil && !IsNil(o.VpcRoutes) {
		return true
	}

	return false
}

// SetVpcRoutes gets a reference to the given []VpcRoute and assigns it to the VpcRoutes field.
func (o *Vpc) SetVpcRoutes(v []VpcRoute) {
	o.VpcRoutes = v
}

func (o Vpc) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Vpc) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.IsDefault) {
		toSerialize["isDefault"] = o.IsDefault
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OwnerId) {
		toSerialize["ownerId"] = o.OwnerId
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.SharedWith) {
		toSerialize["sharedWith"] = o.SharedWith
	}
	if !IsNil(o.AssociatedEnvironments) {
		toSerialize["associatedEnvironments"] = o.AssociatedEnvironments
	}
	if !IsNil(o.CidrBlock) {
		toSerialize["cidrBlock"] = o.CidrBlock
	}
	if !IsNil(o.FirewallRules) {
		toSerialize["firewallRules"] = o.FirewallRules
	}
	if !IsNil(o.InternalDns) {
		toSerialize["internalDns"] = o.InternalDns
	}
	if !IsNil(o.VpcRoutes) {
		toSerialize["vpcRoutes"] = o.VpcRoutes
	}
	return toSerialize, nil
}

type NullableVpc struct {
	value *Vpc
	isSet bool
}

func (v NullableVpc) Get() *Vpc {
	return v.value
}

func (v *NullableVpc) Set(val *Vpc) {
	v.value = val
	v.isSet = true
}

func (v NullableVpc) IsSet() bool {
	return v.isSet
}

func (v *NullableVpc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVpc(val *Vpc) *NullableVpc {
	return &NullableVpc{value: val, isSet: true}
}

func (v NullableVpc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVpc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


