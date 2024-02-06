# UpdateObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewUpdateObject

`func NewUpdateObject() *UpdateObject`

NewUpdateObject instantiates a new UpdateObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateObjectWithDefaults

`func NewUpdateObjectWithDefaults() *UpdateObject`

NewUpdateObjectWithDefaults instantiates a new UpdateObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *UpdateObject) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *UpdateObject) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *UpdateObject) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *UpdateObject) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetPath

`func (o *UpdateObject) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *UpdateObject) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *UpdateObject) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *UpdateObject) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetValue

`func (o *UpdateObject) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateObject) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateObject) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *UpdateObject) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *UpdateObject) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *UpdateObject) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


