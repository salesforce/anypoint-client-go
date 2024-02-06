# ApimPolicyFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyTemplateId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**PolicyId** | Pointer to **int32** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**PointcutData** | Pointer to [**[]PointcutDataItem**](PointcutDataItem.md) |  | [optional] 
**Configuration** | Pointer to **map[string]interface{}** |  | [optional] 
**Template** | Pointer to [**ApimPolicyFullTemplate**](ApimPolicyFullTemplate.md) |  | [optional] 

## Methods

### NewApimPolicyFull

`func NewApimPolicyFull() *ApimPolicyFull`

NewApimPolicyFull instantiates a new ApimPolicyFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApimPolicyFullWithDefaults

`func NewApimPolicyFullWithDefaults() *ApimPolicyFull`

NewApimPolicyFullWithDefaults instantiates a new ApimPolicyFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyTemplateId

`func (o *ApimPolicyFull) GetPolicyTemplateId() string`

GetPolicyTemplateId returns the PolicyTemplateId field if non-nil, zero value otherwise.

### GetPolicyTemplateIdOk

`func (o *ApimPolicyFull) GetPolicyTemplateIdOk() (*string, bool)`

GetPolicyTemplateIdOk returns a tuple with the PolicyTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyTemplateId

`func (o *ApimPolicyFull) SetPolicyTemplateId(v string)`

SetPolicyTemplateId sets PolicyTemplateId field to given value.

### HasPolicyTemplateId

`func (o *ApimPolicyFull) HasPolicyTemplateId() bool`

HasPolicyTemplateId returns a boolean if a field has been set.

### GetType

`func (o *ApimPolicyFull) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApimPolicyFull) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApimPolicyFull) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApimPolicyFull) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPolicyId

`func (o *ApimPolicyFull) GetPolicyId() int32`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *ApimPolicyFull) GetPolicyIdOk() (*int32, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *ApimPolicyFull) SetPolicyId(v int32)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *ApimPolicyFull) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetVersion

`func (o *ApimPolicyFull) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApimPolicyFull) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApimPolicyFull) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ApimPolicyFull) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetOrder

`func (o *ApimPolicyFull) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ApimPolicyFull) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ApimPolicyFull) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ApimPolicyFull) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPointcutData

`func (o *ApimPolicyFull) GetPointcutData() []PointcutDataItem`

GetPointcutData returns the PointcutData field if non-nil, zero value otherwise.

### GetPointcutDataOk

`func (o *ApimPolicyFull) GetPointcutDataOk() (*[]PointcutDataItem, bool)`

GetPointcutDataOk returns a tuple with the PointcutData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointcutData

`func (o *ApimPolicyFull) SetPointcutData(v []PointcutDataItem)`

SetPointcutData sets PointcutData field to given value.

### HasPointcutData

`func (o *ApimPolicyFull) HasPointcutData() bool`

HasPointcutData returns a boolean if a field has been set.

### SetPointcutDataNil

`func (o *ApimPolicyFull) SetPointcutDataNil(b bool)`

 SetPointcutDataNil sets the value for PointcutData to be an explicit nil

### UnsetPointcutData
`func (o *ApimPolicyFull) UnsetPointcutData()`

UnsetPointcutData ensures that no value is present for PointcutData, not even an explicit nil
### GetConfiguration

`func (o *ApimPolicyFull) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ApimPolicyFull) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ApimPolicyFull) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *ApimPolicyFull) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetTemplate

`func (o *ApimPolicyFull) GetTemplate() ApimPolicyFullTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *ApimPolicyFull) GetTemplateOk() (*ApimPolicyFullTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *ApimPolicyFull) SetTemplate(v ApimPolicyFullTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *ApimPolicyFull) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


