# FirewallRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CidrBlock** | Pointer to **string** | The CIDR block of the firewall rule. | [optional] 
**Protocol** | Pointer to **string** | The protocol of the firewall rule. | [optional] 
**FromPort** | Pointer to **int32** | The from port of the firewall rule. | [optional] 
**ToPort** | Pointer to **int32** | The to port of the firewall rule. | [optional] 
**Type** | Pointer to **string** | Type of the firewall rule. Allowed values are [inbound, outbound]. | [optional] 

## Methods

### NewFirewallRule

`func NewFirewallRule() *FirewallRule`

NewFirewallRule instantiates a new FirewallRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallRuleWithDefaults

`func NewFirewallRuleWithDefaults() *FirewallRule`

NewFirewallRuleWithDefaults instantiates a new FirewallRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidrBlock

`func (o *FirewallRule) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *FirewallRule) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *FirewallRule) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.

### HasCidrBlock

`func (o *FirewallRule) HasCidrBlock() bool`

HasCidrBlock returns a boolean if a field has been set.

### GetProtocol

`func (o *FirewallRule) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *FirewallRule) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *FirewallRule) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *FirewallRule) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetFromPort

`func (o *FirewallRule) GetFromPort() int32`

GetFromPort returns the FromPort field if non-nil, zero value otherwise.

### GetFromPortOk

`func (o *FirewallRule) GetFromPortOk() (*int32, bool)`

GetFromPortOk returns a tuple with the FromPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromPort

`func (o *FirewallRule) SetFromPort(v int32)`

SetFromPort sets FromPort field to given value.

### HasFromPort

`func (o *FirewallRule) HasFromPort() bool`

HasFromPort returns a boolean if a field has been set.

### GetToPort

`func (o *FirewallRule) GetToPort() int32`

GetToPort returns the ToPort field if non-nil, zero value otherwise.

### GetToPortOk

`func (o *FirewallRule) GetToPortOk() (*int32, bool)`

GetToPortOk returns a tuple with the ToPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToPort

`func (o *FirewallRule) SetToPort(v int32)`

SetToPort sets ToPort field to given value.

### HasToPort

`func (o *FirewallRule) HasToPort() bool`

HasToPort returns a boolean if a field has been set.

### GetType

`func (o *FirewallRule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FirewallRule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FirewallRule) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FirewallRule) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


