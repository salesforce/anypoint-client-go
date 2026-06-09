# TrustStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileName** | Pointer to **string** |  | [optional] 
**ExpirationDate** | Pointer to **string** |  | [optional] 
**DnList** | Pointer to [**[]TrustStoreDnListInner**](TrustStoreDnListInner.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewTrustStore

`func NewTrustStore() *TrustStore`

NewTrustStore instantiates a new TrustStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustStoreWithDefaults

`func NewTrustStoreWithDefaults() *TrustStore`

NewTrustStoreWithDefaults instantiates a new TrustStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileName

`func (o *TrustStore) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *TrustStore) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *TrustStore) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *TrustStore) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetExpirationDate

`func (o *TrustStore) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *TrustStore) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *TrustStore) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *TrustStore) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetDnList

`func (o *TrustStore) GetDnList() []TrustStoreDnListInner`

GetDnList returns the DnList field if non-nil, zero value otherwise.

### GetDnListOk

`func (o *TrustStore) GetDnListOk() (*[]TrustStoreDnListInner, bool)`

GetDnListOk returns a tuple with the DnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnList

`func (o *TrustStore) SetDnList(v []TrustStoreDnListInner)`

SetDnList sets DnList field to given value.

### HasDnList

`func (o *TrustStore) HasDnList() bool`

HasDnList returns a boolean if a field has been set.

### GetType

`func (o *TrustStore) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TrustStore) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TrustStore) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TrustStore) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


