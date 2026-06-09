# DeploymentDetailsVpnConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | Pointer to **string** | The VPN network address. | [optional] 
**Mask** | Pointer to **string** | The subnet mask for the VPN configuration. | [optional] 

## Methods

### NewDeploymentDetailsVpnConfig

`func NewDeploymentDetailsVpnConfig() *DeploymentDetailsVpnConfig`

NewDeploymentDetailsVpnConfig instantiates a new DeploymentDetailsVpnConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentDetailsVpnConfigWithDefaults

`func NewDeploymentDetailsVpnConfigWithDefaults() *DeploymentDetailsVpnConfig`

NewDeploymentDetailsVpnConfigWithDefaults instantiates a new DeploymentDetailsVpnConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *DeploymentDetailsVpnConfig) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *DeploymentDetailsVpnConfig) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *DeploymentDetailsVpnConfig) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *DeploymentDetailsVpnConfig) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetMask

`func (o *DeploymentDetailsVpnConfig) GetMask() string`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *DeploymentDetailsVpnConfig) GetMaskOk() (*string, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *DeploymentDetailsVpnConfig) SetMask(v string)`

SetMask sets Mask field to given value.

### HasMask

`func (o *DeploymentDetailsVpnConfig) HasMask() bool`

HasMask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


