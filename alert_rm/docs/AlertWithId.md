# AlertWithId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Severity** | **string** |  | 
**Actions** | [**[]Action**](Action.md) |  | 
**Condition** | [**Condition**](Condition.md) |  | 

## Methods

### NewAlertWithId

`func NewAlertWithId(name string, severity string, actions []Action, condition Condition, ) *AlertWithId`

NewAlertWithId instantiates a new AlertWithId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertWithIdWithDefaults

`func NewAlertWithIdWithDefaults() *AlertWithId`

NewAlertWithIdWithDefaults instantiates a new AlertWithId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AlertWithId) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertWithId) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertWithId) SetName(v string)`

SetName sets Name field to given value.


### GetSeverity

`func (o *AlertWithId) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AlertWithId) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AlertWithId) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetActions

`func (o *AlertWithId) GetActions() []Action`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *AlertWithId) GetActionsOk() (*[]Action, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *AlertWithId) SetActions(v []Action)`

SetActions sets Actions field to given value.


### GetCondition

`func (o *AlertWithId) GetCondition() Condition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *AlertWithId) GetConditionOk() (*Condition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *AlertWithId) SetCondition(v Condition)`

SetCondition sets Condition field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


