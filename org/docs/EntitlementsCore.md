# EntitlementsCore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GlobalDeployment** | **bool** | An explanation about the purpose of this instance. | [default to false]
**CreateEnvironments** | **bool** | An explanation about the purpose of this instance. | [default to false]
**CreateSubOrgs** | **bool** | An explanation about the purpose of this instance. | [default to false]
**LoadBalancer** | [**LoadBalancer**](LoadBalancer.md) |  | [default to {}]
**StaticIps** | [**StaticIps**](StaticIps.md) |  | [default to {}]
**VCoresDesign** | [**VCoresDesign**](VCoresDesign.md) |  | [default to {}]
**VCoresProduction** | [**VCoresProduction**](VCoresProduction.md) |  | [default to {}]
**VCoresSandbox** | [**VCoresSandbox**](VCoresSandbox.md) |  | [default to {}]
**Vpcs** | [**Vpcs**](Vpcs.md) |  | [default to {}]
**Vpns** | [**Vpns**](Vpns.md) |  | [default to {}]

## Methods

### NewEntitlementsCore

`func NewEntitlementsCore(globalDeployment bool, createEnvironments bool, createSubOrgs bool, loadBalancer LoadBalancer, staticIps StaticIps, vCoresDesign VCoresDesign, vCoresProduction VCoresProduction, vCoresSandbox VCoresSandbox, vpcs Vpcs, vpns Vpns, ) *EntitlementsCore`

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


### GetLoadBalancer

`func (o *EntitlementsCore) GetLoadBalancer() LoadBalancer`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *EntitlementsCore) GetLoadBalancerOk() (*LoadBalancer, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *EntitlementsCore) SetLoadBalancer(v LoadBalancer)`

SetLoadBalancer sets LoadBalancer field to given value.


### GetStaticIps

`func (o *EntitlementsCore) GetStaticIps() StaticIps`

GetStaticIps returns the StaticIps field if non-nil, zero value otherwise.

### GetStaticIpsOk

`func (o *EntitlementsCore) GetStaticIpsOk() (*StaticIps, bool)`

GetStaticIpsOk returns a tuple with the StaticIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticIps

`func (o *EntitlementsCore) SetStaticIps(v StaticIps)`

SetStaticIps sets StaticIps field to given value.


### GetVCoresDesign

`func (o *EntitlementsCore) GetVCoresDesign() VCoresDesign`

GetVCoresDesign returns the VCoresDesign field if non-nil, zero value otherwise.

### GetVCoresDesignOk

`func (o *EntitlementsCore) GetVCoresDesignOk() (*VCoresDesign, bool)`

GetVCoresDesignOk returns a tuple with the VCoresDesign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCoresDesign

`func (o *EntitlementsCore) SetVCoresDesign(v VCoresDesign)`

SetVCoresDesign sets VCoresDesign field to given value.


### GetVCoresProduction

`func (o *EntitlementsCore) GetVCoresProduction() VCoresProduction`

GetVCoresProduction returns the VCoresProduction field if non-nil, zero value otherwise.

### GetVCoresProductionOk

`func (o *EntitlementsCore) GetVCoresProductionOk() (*VCoresProduction, bool)`

GetVCoresProductionOk returns a tuple with the VCoresProduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCoresProduction

`func (o *EntitlementsCore) SetVCoresProduction(v VCoresProduction)`

SetVCoresProduction sets VCoresProduction field to given value.


### GetVCoresSandbox

`func (o *EntitlementsCore) GetVCoresSandbox() VCoresSandbox`

GetVCoresSandbox returns the VCoresSandbox field if non-nil, zero value otherwise.

### GetVCoresSandboxOk

`func (o *EntitlementsCore) GetVCoresSandboxOk() (*VCoresSandbox, bool)`

GetVCoresSandboxOk returns a tuple with the VCoresSandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCoresSandbox

`func (o *EntitlementsCore) SetVCoresSandbox(v VCoresSandbox)`

SetVCoresSandbox sets VCoresSandbox field to given value.


### GetVpcs

`func (o *EntitlementsCore) GetVpcs() Vpcs`

GetVpcs returns the Vpcs field if non-nil, zero value otherwise.

### GetVpcsOk

`func (o *EntitlementsCore) GetVpcsOk() (*Vpcs, bool)`

GetVpcsOk returns a tuple with the Vpcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcs

`func (o *EntitlementsCore) SetVpcs(v Vpcs)`

SetVpcs sets Vpcs field to given value.


### GetVpns

`func (o *EntitlementsCore) GetVpns() Vpns`

GetVpns returns the Vpns field if non-nil, zero value otherwise.

### GetVpnsOk

`func (o *EntitlementsCore) GetVpnsOk() (*Vpns, bool)`

GetVpnsOk returns a tuple with the Vpns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpns

`func (o *EntitlementsCore) SetVpns(v Vpns)`

SetVpns sets Vpns field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


