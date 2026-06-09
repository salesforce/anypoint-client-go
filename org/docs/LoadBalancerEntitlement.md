# LoadBalancerEntitlement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assigned** | Pointer to **int32** | Load Balancers assigned | [optional] [default to 0]
**Reassigned** | Pointer to **int32** | Load Balancers reassigned | [optional] [default to 0]

## Methods

### NewLoadBalancerEntitlement

`func NewLoadBalancerEntitlement() *LoadBalancerEntitlement`

NewLoadBalancerEntitlement instantiates a new LoadBalancerEntitlement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerEntitlementWithDefaults

`func NewLoadBalancerEntitlementWithDefaults() *LoadBalancerEntitlement`

NewLoadBalancerEntitlementWithDefaults instantiates a new LoadBalancerEntitlement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssigned

`func (o *LoadBalancerEntitlement) GetAssigned() int32`

GetAssigned returns the Assigned field if non-nil, zero value otherwise.

### GetAssignedOk

`func (o *LoadBalancerEntitlement) GetAssignedOk() (*int32, bool)`

GetAssignedOk returns a tuple with the Assigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigned

`func (o *LoadBalancerEntitlement) SetAssigned(v int32)`

SetAssigned sets Assigned field to given value.

### HasAssigned

`func (o *LoadBalancerEntitlement) HasAssigned() bool`

HasAssigned returns a boolean if a field has been set.

### GetReassigned

`func (o *LoadBalancerEntitlement) GetReassigned() int32`

GetReassigned returns the Reassigned field if non-nil, zero value otherwise.

### GetReassignedOk

`func (o *LoadBalancerEntitlement) GetReassignedOk() (*int32, bool)`

GetReassignedOk returns a tuple with the Reassigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReassigned

`func (o *LoadBalancerEntitlement) SetReassigned(v int32)`

SetReassigned sets Reassigned field to given value.

### HasReassigned

`func (o *LoadBalancerEntitlement) HasReassigned() bool`

HasReassigned returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


