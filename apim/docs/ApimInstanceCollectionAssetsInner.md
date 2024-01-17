# ApimInstanceCollectionAssetsInner

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
**Apis** | Pointer to [**[]ApimInstanceCollectionAssetsInnerApisInner**](ApimInstanceCollectionAssetsInnerApisInner.md) |  | [optional] 
**TotalApis** | Pointer to **int32** |  | [optional] 
**AutodiscoveryApiName** | Pointer to **string** |  | [optional] 

## Methods

### NewApimInstanceCollectionAssetsInner

`func NewApimInstanceCollectionAssetsInner() *ApimInstanceCollectionAssetsInner`

NewApimInstanceCollectionAssetsInner instantiates a new ApimInstanceCollectionAssetsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApimInstanceCollectionAssetsInnerWithDefaults

`func NewApimInstanceCollectionAssetsInnerWithDefaults() *ApimInstanceCollectionAssetsInner`

NewApimInstanceCollectionAssetsInnerWithDefaults instantiates a new ApimInstanceCollectionAssetsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudit

`func (o *ApimInstanceCollectionAssetsInner) GetAudit() Audit`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *ApimInstanceCollectionAssetsInner) GetAuditOk() (*Audit, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *ApimInstanceCollectionAssetsInner) SetAudit(v Audit)`

SetAudit sets Audit field to given value.

### HasAudit

`func (o *ApimInstanceCollectionAssetsInner) HasAudit() bool`

HasAudit returns a boolean if a field has been set.

### GetMasterOrganizationId

`func (o *ApimInstanceCollectionAssetsInner) GetMasterOrganizationId() string`

GetMasterOrganizationId returns the MasterOrganizationId field if non-nil, zero value otherwise.

### GetMasterOrganizationIdOk

`func (o *ApimInstanceCollectionAssetsInner) GetMasterOrganizationIdOk() (*string, bool)`

GetMasterOrganizationIdOk returns a tuple with the MasterOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterOrganizationId

`func (o *ApimInstanceCollectionAssetsInner) SetMasterOrganizationId(v string)`

SetMasterOrganizationId sets MasterOrganizationId field to given value.

### HasMasterOrganizationId

`func (o *ApimInstanceCollectionAssetsInner) HasMasterOrganizationId() bool`

HasMasterOrganizationId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *ApimInstanceCollectionAssetsInner) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ApimInstanceCollectionAssetsInner) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ApimInstanceCollectionAssetsInner) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ApimInstanceCollectionAssetsInner) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetId

`func (o *ApimInstanceCollectionAssetsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApimInstanceCollectionAssetsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApimInstanceCollectionAssetsInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ApimInstanceCollectionAssetsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ApimInstanceCollectionAssetsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApimInstanceCollectionAssetsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApimInstanceCollectionAssetsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApimInstanceCollectionAssetsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExchangeAssetName

`func (o *ApimInstanceCollectionAssetsInner) GetExchangeAssetName() string`

GetExchangeAssetName returns the ExchangeAssetName field if non-nil, zero value otherwise.

### GetExchangeAssetNameOk

`func (o *ApimInstanceCollectionAssetsInner) GetExchangeAssetNameOk() (*string, bool)`

GetExchangeAssetNameOk returns a tuple with the ExchangeAssetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeAssetName

`func (o *ApimInstanceCollectionAssetsInner) SetExchangeAssetName(v string)`

SetExchangeAssetName sets ExchangeAssetName field to given value.

### HasExchangeAssetName

`func (o *ApimInstanceCollectionAssetsInner) HasExchangeAssetName() bool`

HasExchangeAssetName returns a boolean if a field has been set.

### GetGroupId

`func (o *ApimInstanceCollectionAssetsInner) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ApimInstanceCollectionAssetsInner) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ApimInstanceCollectionAssetsInner) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ApimInstanceCollectionAssetsInner) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetAssetId

`func (o *ApimInstanceCollectionAssetsInner) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *ApimInstanceCollectionAssetsInner) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *ApimInstanceCollectionAssetsInner) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *ApimInstanceCollectionAssetsInner) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetApis

`func (o *ApimInstanceCollectionAssetsInner) GetApis() []ApimInstanceCollectionAssetsInnerApisInner`

GetApis returns the Apis field if non-nil, zero value otherwise.

### GetApisOk

`func (o *ApimInstanceCollectionAssetsInner) GetApisOk() (*[]ApimInstanceCollectionAssetsInnerApisInner, bool)`

GetApisOk returns a tuple with the Apis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApis

`func (o *ApimInstanceCollectionAssetsInner) SetApis(v []ApimInstanceCollectionAssetsInnerApisInner)`

SetApis sets Apis field to given value.

### HasApis

`func (o *ApimInstanceCollectionAssetsInner) HasApis() bool`

HasApis returns a boolean if a field has been set.

### GetTotalApis

`func (o *ApimInstanceCollectionAssetsInner) GetTotalApis() int32`

GetTotalApis returns the TotalApis field if non-nil, zero value otherwise.

### GetTotalApisOk

`func (o *ApimInstanceCollectionAssetsInner) GetTotalApisOk() (*int32, bool)`

GetTotalApisOk returns a tuple with the TotalApis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalApis

`func (o *ApimInstanceCollectionAssetsInner) SetTotalApis(v int32)`

SetTotalApis sets TotalApis field to given value.

### HasTotalApis

`func (o *ApimInstanceCollectionAssetsInner) HasTotalApis() bool`

HasTotalApis returns a boolean if a field has been set.

### GetAutodiscoveryApiName

`func (o *ApimInstanceCollectionAssetsInner) GetAutodiscoveryApiName() string`

GetAutodiscoveryApiName returns the AutodiscoveryApiName field if non-nil, zero value otherwise.

### GetAutodiscoveryApiNameOk

`func (o *ApimInstanceCollectionAssetsInner) GetAutodiscoveryApiNameOk() (*string, bool)`

GetAutodiscoveryApiNameOk returns a tuple with the AutodiscoveryApiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutodiscoveryApiName

`func (o *ApimInstanceCollectionAssetsInner) SetAutodiscoveryApiName(v string)`

SetAutodiscoveryApiName sets AutodiscoveryApiName field to given value.

### HasAutodiscoveryApiName

`func (o *ApimInstanceCollectionAssetsInner) HasAutodiscoveryApiName() bool`

HasAutodiscoveryApiName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


