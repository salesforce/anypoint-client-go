# ApimPolicyBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationData** | Pointer to **map[string]interface{}** |  | [optional] 
**GroupId** | Pointer to **string** |  | [optional] 
**AssetId** | Pointer to **string** |  | [optional] 
**AssetVersion** | Pointer to **string** |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**PointcutData** | Pointer to [**[]PointcutDataItem**](PointcutDataItem.md) |  | [optional] 

## Methods

### NewApimPolicyBody

`func NewApimPolicyBody() *ApimPolicyBody`

NewApimPolicyBody instantiates a new ApimPolicyBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApimPolicyBodyWithDefaults

`func NewApimPolicyBodyWithDefaults() *ApimPolicyBody`

NewApimPolicyBodyWithDefaults instantiates a new ApimPolicyBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurationData

`func (o *ApimPolicyBody) GetConfigurationData() map[string]interface{}`

GetConfigurationData returns the ConfigurationData field if non-nil, zero value otherwise.

### GetConfigurationDataOk

`func (o *ApimPolicyBody) GetConfigurationDataOk() (*map[string]interface{}, bool)`

GetConfigurationDataOk returns a tuple with the ConfigurationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationData

`func (o *ApimPolicyBody) SetConfigurationData(v map[string]interface{})`

SetConfigurationData sets ConfigurationData field to given value.

### HasConfigurationData

`func (o *ApimPolicyBody) HasConfigurationData() bool`

HasConfigurationData returns a boolean if a field has been set.

### GetGroupId

`func (o *ApimPolicyBody) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ApimPolicyBody) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ApimPolicyBody) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ApimPolicyBody) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetAssetId

`func (o *ApimPolicyBody) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *ApimPolicyBody) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *ApimPolicyBody) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *ApimPolicyBody) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetAssetVersion

`func (o *ApimPolicyBody) GetAssetVersion() string`

GetAssetVersion returns the AssetVersion field if non-nil, zero value otherwise.

### GetAssetVersionOk

`func (o *ApimPolicyBody) GetAssetVersionOk() (*string, bool)`

GetAssetVersionOk returns a tuple with the AssetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetVersion

`func (o *ApimPolicyBody) SetAssetVersion(v string)`

SetAssetVersion sets AssetVersion field to given value.

### HasAssetVersion

`func (o *ApimPolicyBody) HasAssetVersion() bool`

HasAssetVersion returns a boolean if a field has been set.

### GetOrder

`func (o *ApimPolicyBody) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ApimPolicyBody) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ApimPolicyBody) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ApimPolicyBody) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPointcutData

`func (o *ApimPolicyBody) GetPointcutData() []PointcutDataItem`

GetPointcutData returns the PointcutData field if non-nil, zero value otherwise.

### GetPointcutDataOk

`func (o *ApimPolicyBody) GetPointcutDataOk() (*[]PointcutDataItem, bool)`

GetPointcutDataOk returns a tuple with the PointcutData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointcutData

`func (o *ApimPolicyBody) SetPointcutData(v []PointcutDataItem)`

SetPointcutData sets PointcutData field to given value.

### HasPointcutData

`func (o *ApimPolicyBody) HasPointcutData() bool`

HasPointcutData returns a boolean if a field has been set.

### SetPointcutDataNil

`func (o *ApimPolicyBody) SetPointcutDataNil(b bool)`

 SetPointcutDataNil sets the value for PointcutData to be an explicit nil

### UnsetPointcutData
`func (o *ApimPolicyBody) UnsetPointcutData()`

UnsetPointcutData ensures that no value is present for PointcutData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


