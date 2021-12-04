# AlertBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Condition** | Pointer to [**Condition**](Condition.md) |  | [optional] 
**Actions** | Pointer to [**[]Action**](Action.md) |  | [optional] 

## Methods

### NewAlertBody

`func NewAlertBody() *AlertBody`

NewAlertBody instantiates a new AlertBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertBodyWithDefaults

`func NewAlertBodyWithDefaults() *AlertBody`

NewAlertBodyWithDefaults instantiates a new AlertBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AlertBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AlertBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSeverity

`func (o *AlertBody) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AlertBody) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AlertBody) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AlertBody) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetCondition

`func (o *AlertBody) GetCondition() Condition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *AlertBody) GetConditionOk() (*Condition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *AlertBody) SetCondition(v Condition)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *AlertBody) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetActions

`func (o *AlertBody) GetActions() []Action`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *AlertBody) GetActionsOk() (*[]Action, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *AlertBody) SetActions(v []Action)`

SetActions sets Actions field to given value.

### HasActions

`func (o *AlertBody) HasActions() bool`

HasActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


