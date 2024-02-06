# ApimInstanceCollectionAssets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audit** | Pointer to [**Audit**](Audit.md) |  | [optional] 
**MasterOrganizationId** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ExchangeAssetName** | Pointer to **string** |  | [optional] 
**GroupId** | Pointer to **string** |  | [optional] 
**AssetId** | Pointer to **string** |  | [optional] 
**Apis** | Pointer to [**[]ApimInstanceCollectionApis**](ApimInstanceCollectionApis.md) |  | [optional] 
**TotalApis** | Pointer to **int32** |  | [optional] 
**AutodiscoveryApiName** | Pointer to **string** |  | [optional] 

## Methods

### NewApimInstanceCollectionAssets

`func NewApimInstanceCollectionAssets() *ApimInstanceCollectionAssets`

NewApimInstanceCollectionAssets instantiates a new ApimInstanceCollectionAssets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApimInstanceCollectionAssetsWithDefaults

`func NewApimInstanceCollectionAssetsWithDefaults() *ApimInstanceCollectionAssets`

NewApimInstanceCollectionAssetsWithDefaults instantiates a new ApimInstanceCollectionAssets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudit

`func (o *ApimInstanceCollectionAssets) GetAudit() Audit`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *ApimInstanceCollectionAssets) GetAuditOk() (*Audit, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *ApimInstanceCollectionAssets) SetAudit(v Audit)`

SetAudit sets Audit field to given value.

### HasAudit

`func (o *ApimInstanceCollectionAssets) HasAudit() bool`

HasAudit returns a boolean if a field has been set.

### GetMasterOrganizationId

`func (o *ApimInstanceCollectionAssets) GetMasterOrganizationId() string`

GetMasterOrganizationId returns the MasterOrganizationId field if non-nil, zero value otherwise.

### GetMasterOrganizationIdOk

`func (o *ApimInstanceCollectionAssets) GetMasterOrganizationIdOk() (*string, bool)`

GetMasterOrganizationIdOk returns a tuple with the MasterOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterOrganizationId

`func (o *ApimInstanceCollectionAssets) SetMasterOrganizationId(v string)`

SetMasterOrganizationId sets MasterOrganizationId field to given value.

### HasMasterOrganizationId

`func (o *ApimInstanceCollectionAssets) HasMasterOrganizationId() bool`

HasMasterOrganizationId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *ApimInstanceCollectionAssets) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ApimInstanceCollectionAssets) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ApimInstanceCollectionAssets) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ApimInstanceCollectionAssets) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetId

`func (o *ApimInstanceCollectionAssets) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApimInstanceCollectionAssets) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApimInstanceCollectionAssets) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ApimInstanceCollectionAssets) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ApimInstanceCollectionAssets) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApimInstanceCollectionAssets) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApimInstanceCollectionAssets) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApimInstanceCollectionAssets) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExchangeAssetName

`func (o *ApimInstanceCollectionAssets) GetExchangeAssetName() string`

GetExchangeAssetName returns the ExchangeAssetName field if non-nil, zero value otherwise.

### GetExchangeAssetNameOk

`func (o *ApimInstanceCollectionAssets) GetExchangeAssetNameOk() (*string, bool)`

GetExchangeAssetNameOk returns a tuple with the ExchangeAssetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeAssetName

`func (o *ApimInstanceCollectionAssets) SetExchangeAssetName(v string)`

SetExchangeAssetName sets ExchangeAssetName field to given value.

### HasExchangeAssetName

`func (o *ApimInstanceCollectionAssets) HasExchangeAssetName() bool`

HasExchangeAssetName returns a boolean if a field has been set.

### GetGroupId

`func (o *ApimInstanceCollectionAssets) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ApimInstanceCollectionAssets) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ApimInstanceCollectionAssets) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ApimInstanceCollectionAssets) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetAssetId

`func (o *ApimInstanceCollectionAssets) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *ApimInstanceCollectionAssets) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *ApimInstanceCollectionAssets) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *ApimInstanceCollectionAssets) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetApis

`func (o *ApimInstanceCollectionAssets) GetApis() []ApimInstanceCollectionApis`

GetApis returns the Apis field if non-nil, zero value otherwise.

### GetApisOk

`func (o *ApimInstanceCollectionAssets) GetApisOk() (*[]ApimInstanceCollectionApis, bool)`

GetApisOk returns a tuple with the Apis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApis

`func (o *ApimInstanceCollectionAssets) SetApis(v []ApimInstanceCollectionApis)`

SetApis sets Apis field to given value.

### HasApis

`func (o *ApimInstanceCollectionAssets) HasApis() bool`

HasApis returns a boolean if a field has been set.

### GetTotalApis

`func (o *ApimInstanceCollectionAssets) GetTotalApis() int32`

GetTotalApis returns the TotalApis field if non-nil, zero value otherwise.

### GetTotalApisOk

`func (o *ApimInstanceCollectionAssets) GetTotalApisOk() (*int32, bool)`

GetTotalApisOk returns a tuple with the TotalApis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalApis

`func (o *ApimInstanceCollectionAssets) SetTotalApis(v int32)`

SetTotalApis sets TotalApis field to given value.

### HasTotalApis

`func (o *ApimInstanceCollectionAssets) HasTotalApis() bool`

HasTotalApis returns a boolean if a field has been set.

### GetAutodiscoveryApiName

`func (o *ApimInstanceCollectionAssets) GetAutodiscoveryApiName() string`

GetAutodiscoveryApiName returns the AutodiscoveryApiName field if non-nil, zero value otherwise.

### GetAutodiscoveryApiNameOk

`func (o *ApimInstanceCollectionAssets) GetAutodiscoveryApiNameOk() (*string, bool)`

GetAutodiscoveryApiNameOk returns a tuple with the AutodiscoveryApiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutodiscoveryApiName

`func (o *ApimInstanceCollectionAssets) SetAutodiscoveryApiName(v string)`

SetAutodiscoveryApiName sets AutodiscoveryApiName field to given value.

### HasAutodiscoveryApiName

`func (o *ApimInstanceCollectionAssets) HasAutodiscoveryApiName() bool`

HasAutodiscoveryApiName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


