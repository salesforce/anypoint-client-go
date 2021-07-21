# Condition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | **string** |  | 
**PeriodCount** | **int32** |  | 
**Value** | **int32** |  | 
**ResourceType** | **string** |  | 
**Type** | **string** |  | 
**PeriodMins** | **int32** |  | 
**Resources** | **[]string** |  | 

## Methods

### NewCondition

`func NewCondition(operator string, periodCount int32, value int32, resourceType string, type_ string, periodMins int32, resources []string, ) *Condition`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


