# ApimPolicyFullData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policies** | Pointer to [**[]ApimPolicyFull**](ApimPolicyFull.md) |  | [optional] 
**Tiers** | Pointer to [**ApimPolicyFullDataTiers**](ApimPolicyFullDataTiers.md) |  | [optional] 

## Methods

### NewApimPolicyFullData

`func NewApimPolicyFullData() *ApimPolicyFullData`

NewApimPolicyFullData instantiates a new ApimPolicyFullData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApimPolicyFullDataWithDefaults

`func NewApimPolicyFullDataWithDefaults() *ApimPolicyFullData`

NewApimPolicyFullDataWithDefaults instantiates a new ApimPolicyFullData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicies

`func (o *ApimPolicyFullData) GetPolicies() []ApimPolicyFull`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *ApimPolicyFullData) GetPoliciesOk() (*[]ApimPolicyFull, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *ApimPolicyFullData) SetPolicies(v []ApimPolicyFull)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *ApimPolicyFullData) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetTiers

`func (o *ApimPolicyFullData) GetTiers() ApimPolicyFullDataTiers`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *ApimPolicyFullData) GetTiersOk() (*ApimPolicyFullDataTiers, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *ApimPolicyFullData) SetTiers(v ApimPolicyFullDataTiers)`

SetTiers sets Tiers field to given value.

### HasTiers

`func (o *ApimPolicyFullData) HasTiers() bool`

HasTiers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


