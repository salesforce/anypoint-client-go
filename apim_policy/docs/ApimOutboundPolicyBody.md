# ApimOutboundPolicyBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationData** | **map[string]interface{}** |  | 
**ApiVersionId** | **int32** | The numeric API instance id (same as &#x60;apiId&#x60; in &#x60;ApimPolicy&#x60;). | 
**GroupId** | **string** |  | 
**AssetId** | **string** |  | 
**AssetVersion** | **string** |  | 
**Label** | Pointer to **NullableString** |  | [optional] 
**UpstreamIds** | **[]string** | Upstream identifiers (UUID strings) the policy must bind to. The API creates one policy record per id and returns them all in the response array.  | 

## Methods

### NewApimOutboundPolicyBody

`func NewApimOutboundPolicyBody(configurationData map[string]interface{}, apiVersionId int32, groupId string, assetId string, assetVersion string, upstreamIds []string, ) *ApimOutboundPolicyBody`

NewApimOutboundPolicyBody instantiates a new ApimOutboundPolicyBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApimOutboundPolicyBodyWithDefaults

`func NewApimOutboundPolicyBodyWithDefaults() *ApimOutboundPolicyBody`

NewApimOutboundPolicyBodyWithDefaults instantiates a new ApimOutboundPolicyBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurationData

`func (o *ApimOutboundPolicyBody) GetConfigurationData() map[string]interface{}`

GetConfigurationData returns the ConfigurationData field if non-nil, zero value otherwise.

### GetConfigurationDataOk

`func (o *ApimOutboundPolicyBody) GetConfigurationDataOk() (*map[string]interface{}, bool)`

GetConfigurationDataOk returns a tuple with the ConfigurationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationData

`func (o *ApimOutboundPolicyBody) SetConfigurationData(v map[string]interface{})`

SetConfigurationData sets ConfigurationData field to given value.


### GetApiVersionId

`func (o *ApimOutboundPolicyBody) GetApiVersionId() int32`

GetApiVersionId returns the ApiVersionId field if non-nil, zero value otherwise.

### GetApiVersionIdOk

`func (o *ApimOutboundPolicyBody) GetApiVersionIdOk() (*int32, bool)`

GetApiVersionIdOk returns a tuple with the ApiVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersionId

`func (o *ApimOutboundPolicyBody) SetApiVersionId(v int32)`

SetApiVersionId sets ApiVersionId field to given value.


### GetGroupId

`func (o *ApimOutboundPolicyBody) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ApimOutboundPolicyBody) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ApimOutboundPolicyBody) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetAssetId

`func (o *ApimOutboundPolicyBody) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *ApimOutboundPolicyBody) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *ApimOutboundPolicyBody) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetAssetVersion

`func (o *ApimOutboundPolicyBody) GetAssetVersion() string`

GetAssetVersion returns the AssetVersion field if non-nil, zero value otherwise.

### GetAssetVersionOk

`func (o *ApimOutboundPolicyBody) GetAssetVersionOk() (*string, bool)`

GetAssetVersionOk returns a tuple with the AssetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetVersion

`func (o *ApimOutboundPolicyBody) SetAssetVersion(v string)`

SetAssetVersion sets AssetVersion field to given value.


### GetLabel

`func (o *ApimOutboundPolicyBody) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ApimOutboundPolicyBody) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ApimOutboundPolicyBody) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ApimOutboundPolicyBody) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *ApimOutboundPolicyBody) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *ApimOutboundPolicyBody) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetUpstreamIds

`func (o *ApimOutboundPolicyBody) GetUpstreamIds() []string`

GetUpstreamIds returns the UpstreamIds field if non-nil, zero value otherwise.

### GetUpstreamIdsOk

`func (o *ApimOutboundPolicyBody) GetUpstreamIdsOk() (*[]string, bool)`

GetUpstreamIdsOk returns a tuple with the UpstreamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamIds

`func (o *ApimOutboundPolicyBody) SetUpstreamIds(v []string)`

SetUpstreamIds sets UpstreamIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


