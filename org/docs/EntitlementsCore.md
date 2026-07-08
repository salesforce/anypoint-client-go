# EntitlementsCore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GlobalDeployment** | Pointer to **bool** | Global Deployment | [optional] [default to false]
**CreateEnvironments** | Pointer to **bool** | Create Environments | [optional] [default to false]
**CreateSubOrgs** | Pointer to **bool** | Create Sub Orgs | [optional] [default to false]
**LoadBalancer** | Pointer to [**LoadBalancerEntitlement**](LoadBalancerEntitlement.md) |  | [optional] 
**StaticIps** | Pointer to [**StaticIpsEntitlement**](StaticIpsEntitlement.md) |  | [optional] 
**VCoresDesign** | Pointer to [**VCoresDesignEntitlement**](VCoresDesignEntitlement.md) |  | [optional] 
**VCoresProduction** | Pointer to [**VCoresProductionEntitlement**](VCoresProductionEntitlement.md) |  | [optional] 
**VCoresSandbox** | Pointer to [**VCoresSandboxEntitlement**](VCoresSandboxEntitlement.md) |  | [optional] 
**Vpcs** | Pointer to [**VpcsEntitlement**](VpcsEntitlement.md) |  | [optional] 
**Vpns** | Pointer to [**VpnsEntitlement**](VpnsEntitlement.md) |  | [optional] 

## Methods

### NewEntitlementsCore

`func NewEntitlementsCore() *EntitlementsCore`

NewEntitlementsCore instantiates a new EntitlementsCore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementsCoreWithDefaults

`func NewEntitlementsCoreWithDefaults() *EntitlementsCore`

NewEntitlementsCoreWithDefaults instantiates a new EntitlementsCore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlobalDeployment

`func (o *EntitlementsCore) GetGlobalDeployment() bool`

GetGlobalDeployment returns the GlobalDeployment field if non-nil, zero value otherwise.

### GetGlobalDeploymentOk

`func (o *EntitlementsCore) GetGlobalDeploymentOk() (*bool, bool)`

GetGlobalDeploymentOk returns a tuple with the GlobalDeployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalDeployment

`func (o *EntitlementsCore) SetGlobalDeployment(v bool)`

SetGlobalDeployment sets GlobalDeployment field to given value.

### HasGlobalDeployment

`func (o *EntitlementsCore) HasGlobalDeployment() bool`

HasGlobalDeployment returns a boolean if a field has been set.

### GetCreateEnvironments

`func (o *EntitlementsCore) GetCreateEnvironments() bool`

GetCreateEnvironments returns the CreateEnvironments field if non-nil, zero value otherwise.

### GetCreateEnvironmentsOk

`func (o *EntitlementsCore) GetCreateEnvironmentsOk() (*bool, bool)`

GetCreateEnvironmentsOk returns a tuple with the CreateEnvironments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateEnvironments

`func (o *EntitlementsCore) SetCreateEnvironments(v bool)`

SetCreateEnvironments sets CreateEnvironments field to given value.

### HasCreateEnvironments

`func (o *EntitlementsCore) HasCreateEnvironments() bool`

HasCreateEnvironments returns a boolean if a field has been set.

### GetCreateSubOrgs

`func (o *EntitlementsCore) GetCreateSubOrgs() bool`

GetCreateSubOrgs returns the CreateSubOrgs field if non-nil, zero value otherwise.

### GetCreateSubOrgsOk

`func (o *EntitlementsCore) GetCreateSubOrgsOk() (*bool, bool)`

GetCreateSubOrgsOk returns a tuple with the CreateSubOrgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateSubOrgs

`func (o *EntitlementsCore) SetCreateSubOrgs(v bool)`

SetCreateSubOrgs sets CreateSubOrgs field to given value.

### HasCreateSubOrgs

`func (o *EntitlementsCore) HasCreateSubOrgs() bool`

HasCreateSubOrgs returns a boolean if a field has been set.

### GetLoadBalancer

`func (o *EntitlementsCore) GetLoadBalancer() LoadBalancerEntitlement`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *EntitlementsCore) GetLoadBalancerOk() (*LoadBalancerEntitlement, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *EntitlementsCore) SetLoadBalancer(v LoadBalancerEntitlement)`

SetLoadBalancer sets LoadBalancer field to given value.

### HasLoadBalancer

`func (o *EntitlementsCore) HasLoadBalancer() bool`

HasLoadBalancer returns a boolean if a field has been set.

### GetStaticIps

`func (o *EntitlementsCore) GetStaticIps() StaticIpsEntitlement`

GetStaticIps returns the StaticIps field if non-nil, zero value otherwise.

### GetStaticIpsOk

`func (o *EntitlementsCore) GetStaticIpsOk() (*StaticIpsEntitlement, bool)`

GetStaticIpsOk returns a tuple with the StaticIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticIps

`func (o *EntitlementsCore) SetStaticIps(v StaticIpsEntitlement)`

SetStaticIps sets StaticIps field to given value.

### HasStaticIps

`func (o *EntitlementsCore) HasStaticIps() bool`

HasStaticIps returns a boolean if a field has been set.

### GetVCoresDesign

`func (o *EntitlementsCore) GetVCoresDesign() VCoresDesignEntitlement`

GetVCoresDesign returns the VCoresDesign field if non-nil, zero value otherwise.

### GetVCoresDesignOk

`func (o *EntitlementsCore) GetVCoresDesignOk() (*VCoresDesignEntitlement, bool)`

GetVCoresDesignOk returns a tuple with the VCoresDesign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCoresDesign

`func (o *EntitlementsCore) SetVCoresDesign(v VCoresDesignEntitlement)`

SetVCoresDesign sets VCoresDesign field to given value.

### HasVCoresDesign

`func (o *EntitlementsCore) HasVCoresDesign() bool`

HasVCoresDesign returns a boolean if a field has been set.

### GetVCoresProduction

`func (o *EntitlementsCore) GetVCoresProduction() VCoresProductionEntitlement`

GetVCoresProduction returns the VCoresProduction field if non-nil, zero value otherwise.

### GetVCoresProductionOk

`func (o *EntitlementsCore) GetVCoresProductionOk() (*VCoresProductionEntitlement, bool)`

GetVCoresProductionOk returns a tuple with the VCoresProduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCoresProduction

`func (o *EntitlementsCore) SetVCoresProduction(v VCoresProductionEntitlement)`

SetVCoresProduction sets VCoresProduction field to given value.

### HasVCoresProduction

`func (o *EntitlementsCore) HasVCoresProduction() bool`

HasVCoresProduction returns a boolean if a field has been set.

### GetVCoresSandbox

`func (o *EntitlementsCore) GetVCoresSandbox() VCoresSandboxEntitlement`

GetVCoresSandbox returns the VCoresSandbox field if non-nil, zero value otherwise.

### GetVCoresSandboxOk

`func (o *EntitlementsCore) GetVCoresSandboxOk() (*VCoresSandboxEntitlement, bool)`

GetVCoresSandboxOk returns a tuple with the VCoresSandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCoresSandbox

`func (o *EntitlementsCore) SetVCoresSandbox(v VCoresSandboxEntitlement)`

SetVCoresSandbox sets VCoresSandbox field to given value.

### HasVCoresSandbox

`func (o *EntitlementsCore) HasVCoresSandbox() bool`

HasVCoresSandbox returns a boolean if a field has been set.

### GetVpcs

`func (o *EntitlementsCore) GetVpcs() VpcsEntitlement`

GetVpcs returns the Vpcs field if non-nil, zero value otherwise.

### GetVpcsOk

`func (o *EntitlementsCore) GetVpcsOk() (*VpcsEntitlement, bool)`

GetVpcsOk returns a tuple with the Vpcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcs

`func (o *EntitlementsCore) SetVpcs(v VpcsEntitlement)`

SetVpcs sets Vpcs field to given value.

### HasVpcs

`func (o *EntitlementsCore) HasVpcs() bool`

HasVpcs returns a boolean if a field has been set.

### GetVpns

`func (o *EntitlementsCore) GetVpns() VpnsEntitlement`

GetVpns returns the Vpns field if non-nil, zero value otherwise.

### GetVpnsOk

`func (o *EntitlementsCore) GetVpnsOk() (*VpnsEntitlement, bool)`

GetVpnsOk returns a tuple with the Vpns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpns

`func (o *EntitlementsCore) SetVpns(v VpnsEntitlement)`

SetVpns sets Vpns field to given value.

### HasVpns

`func (o *EntitlementsCore) HasVpns() bool`

HasVpns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


