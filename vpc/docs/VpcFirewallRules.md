# VpcFirewallRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CidrBlock** | **string** | An explanation about the purpose of this instance. | [default to ""]
**FromPort** | **int32** | An explanation about the purpose of this instance. | [default to 0]
**Protocol** | **string** | An explanation about the purpose of this instance. | [default to ""]
**ToPort** | **int32** | An explanation about the purpose of this instance. | [default to 0]

## Methods

### NewVpcFirewallRules

`func NewVpcFirewallRules(cidrBlock string, fromPort int32, protocol string, toPort int32, ) *VpcFirewallRules`

NewVpcFirewallRules instantiates a new VpcFirewallRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpcFirewallRulesWithDefaults

`func NewVpcFirewallRulesWithDefaults() *VpcFirewallRules`

NewVpcFirewallRulesWithDefaults instantiates a new VpcFirewallRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidrBlock

`func (o *VpcFirewallRules) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *VpcFirewallRules) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *VpcFirewallRules) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.


### GetFromPort

`func (o *VpcFirewallRules) GetFromPort() int32`

GetFromPort returns the FromPort field if non-nil, zero value otherwise.

### GetFromPortOk

`func (o *VpcFirewallRules) GetFromPortOk() (*int32, bool)`

GetFromPortOk returns a tuple with the FromPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromPort

`func (o *VpcFirewallRules) SetFromPort(v int32)`

SetFromPort sets FromPort field to given value.


### GetProtocol

`func (o *VpcFirewallRules) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *VpcFirewallRules) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *VpcFirewallRules) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetToPort

`func (o *VpcFirewallRules) GetToPort() int32`

GetToPort returns the ToPort field if non-nil, zero value otherwise.

### GetToPortOk

`func (o *VpcFirewallRules) GetToPortOk() (*int32, bool)`

GetToPortOk returns a tuple with the ToPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToPort

`func (o *VpcFirewallRules) SetToPort(v int32)`

SetToPort sets ToPort field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


