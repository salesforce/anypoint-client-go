# ApiGovernanceEntitlement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | API Governance enabled | [optional] 
**ApisPerMonth** | Pointer to **int32** | APIs per month | [optional] 

## Methods

### NewApiGovernanceEntitlement

`func NewApiGovernanceEntitlement() *ApiGovernanceEntitlement`

NewApiGovernanceEntitlement instantiates a new ApiGovernanceEntitlement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiGovernanceEntitlementWithDefaults

`func NewApiGovernanceEntitlementWithDefaults() *ApiGovernanceEntitlement`

NewApiGovernanceEntitlementWithDefaults instantiates a new ApiGovernanceEntitlement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ApiGovernanceEntitlement) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApiGovernanceEntitlement) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApiGovernanceEntitlement) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApiGovernanceEntitlement) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetApisPerMonth

`func (o *ApiGovernanceEntitlement) GetApisPerMonth() int32`

GetApisPerMonth returns the ApisPerMonth field if non-nil, zero value otherwise.

### GetApisPerMonthOk

`func (o *ApiGovernanceEntitlement) GetApisPerMonthOk() (*int32, bool)`

GetApisPerMonthOk returns a tuple with the ApisPerMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApisPerMonth

`func (o *ApiGovernanceEntitlement) SetApisPerMonth(v int32)`

SetApisPerMonth sets ApisPerMonth field to given value.

### HasApisPerMonth

`func (o *ApiGovernanceEntitlement) HasApisPerMonth() bool`

HasApisPerMonth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


