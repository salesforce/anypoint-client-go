# Condition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to **string** |  | [optional] 
**PeriodCount** | Pointer to **int32** |  | [optional] 
**Value** | Pointer to **int32** |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**PeriodMins** | Pointer to **int32** |  | [optional] 
**Resources** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCondition

`func NewCondition() *Condition`

NewCondition instantiates a new Condition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionWithDefaults

`func NewConditionWithDefaults() *Condition`

NewConditionWithDefaults instantiates a new Condition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *Condition) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *Condition) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *Condition) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *Condition) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetPeriodCount

`func (o *Condition) GetPeriodCount() int32`

GetPeriodCount returns the PeriodCount field if non-nil, zero value otherwise.

### GetPeriodCountOk

`func (o *Condition) GetPeriodCountOk() (*int32, bool)`

GetPeriodCountOk returns a tuple with the PeriodCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodCount

`func (o *Condition) SetPeriodCount(v int32)`

SetPeriodCount sets PeriodCount field to given value.

### HasPeriodCount

`func (o *Condition) HasPeriodCount() bool`

HasPeriodCount returns a boolean if a field has been set.

### GetValue

`func (o *Condition) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Condition) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Condition) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *Condition) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetResourceType

`func (o *Condition) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *Condition) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *Condition) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *Condition) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetType

`func (o *Condition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Condition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Condition) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Condition) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPeriodMins

`func (o *Condition) GetPeriodMins() int32`

GetPeriodMins returns the PeriodMins field if non-nil, zero value otherwise.

### GetPeriodMinsOk

`func (o *Condition) GetPeriodMinsOk() (*int32, bool)`

GetPeriodMinsOk returns a tuple with the PeriodMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodMins

`func (o *Condition) SetPeriodMins(v int32)`

SetPeriodMins sets PeriodMins field to given value.

### HasPeriodMins

`func (o *Condition) HasPeriodMins() bool`

HasPeriodMins returns a boolean if a field has been set.

### GetResources

`func (o *Condition) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *Condition) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *Condition) SetResources(v []string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *Condition) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


