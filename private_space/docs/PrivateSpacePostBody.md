# PrivateSpacePostBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the private space. | [optional] 
**Environments** | Pointer to [**PrivateSpaceAssociatedEnvironments**](PrivateSpaceAssociatedEnvironments.md) |  | [optional] 
**Network** | Pointer to [**PrivateSpaceNetworkCreate**](PrivateSpaceNetworkCreate.md) |  | [optional] 
**FirewallRules** | Pointer to [**[]FirewallRule**](FirewallRule.md) | The list of firewall rules for the Private Space network. | [optional] 

## Methods

### NewPrivateSpacePostBody

`func NewPrivateSpacePostBody() *PrivateSpacePostBody`

NewPrivateSpacePostBody instantiates a new PrivateSpacePostBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateSpacePostBodyWithDefaults

`func NewPrivateSpacePostBodyWithDefaults() *PrivateSpacePostBody`

NewPrivateSpacePostBodyWithDefaults instantiates a new PrivateSpacePostBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PrivateSpacePostBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrivateSpacePostBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrivateSpacePostBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PrivateSpacePostBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnvironments

`func (o *PrivateSpacePostBody) GetEnvironments() PrivateSpaceAssociatedEnvironments`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *PrivateSpacePostBody) GetEnvironmentsOk() (*PrivateSpaceAssociatedEnvironments, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *PrivateSpacePostBody) SetEnvironments(v PrivateSpaceAssociatedEnvironments)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *PrivateSpacePostBody) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetNetwork

`func (o *PrivateSpacePostBody) GetNetwork() PrivateSpaceNetworkCreate`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *PrivateSpacePostBody) GetNetworkOk() (*PrivateSpaceNetworkCreate, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *PrivateSpacePostBody) SetNetwork(v PrivateSpaceNetworkCreate)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *PrivateSpacePostBody) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetFirewallRules

`func (o *PrivateSpacePostBody) GetFirewallRules() []FirewallRule`

GetFirewallRules returns the FirewallRules field if non-nil, zero value otherwise.

### GetFirewallRulesOk

`func (o *PrivateSpacePostBody) GetFirewallRulesOk() (*[]FirewallRule, bool)`

GetFirewallRulesOk returns a tuple with the FirewallRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallRules

`func (o *PrivateSpacePostBody) SetFirewallRules(v []FirewallRule)`

SetFirewallRules sets FirewallRules field to given value.

### HasFirewallRules

`func (o *PrivateSpacePostBody) HasFirewallRules() bool`

HasFirewallRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


