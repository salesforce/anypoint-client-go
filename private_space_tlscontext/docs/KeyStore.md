# KeyStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Cn** | Pointer to **string** |  | [optional] 
**San** | Pointer to **[]string** |  | [optional] 
**FileName** | Pointer to **string** |  | [optional] 
**KeyFileName** | Pointer to **string** |  | [optional] 
**CapathFileName** | Pointer to **string** |  | [optional] 
**ExpirationDate** | Pointer to **string** |  | [optional] 

## Methods

### NewKeyStore

`func NewKeyStore() *KeyStore`

NewKeyStore instantiates a new KeyStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyStoreWithDefaults

`func NewKeyStoreWithDefaults() *KeyStore`

NewKeyStoreWithDefaults instantiates a new KeyStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *KeyStore) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KeyStore) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KeyStore) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KeyStore) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCn

`func (o *KeyStore) GetCn() string`

GetCn returns the Cn field if non-nil, zero value otherwise.

### GetCnOk

`func (o *KeyStore) GetCnOk() (*string, bool)`

GetCnOk returns a tuple with the Cn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCn

`func (o *KeyStore) SetCn(v string)`

SetCn sets Cn field to given value.

### HasCn

`func (o *KeyStore) HasCn() bool`

HasCn returns a boolean if a field has been set.

### GetSan

`func (o *KeyStore) GetSan() []string`

GetSan returns the San field if non-nil, zero value otherwise.

### GetSanOk

`func (o *KeyStore) GetSanOk() (*[]string, bool)`

GetSanOk returns a tuple with the San field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSan

`func (o *KeyStore) SetSan(v []string)`

SetSan sets San field to given value.

### HasSan

`func (o *KeyStore) HasSan() bool`

HasSan returns a boolean if a field has been set.

### GetFileName

`func (o *KeyStore) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *KeyStore) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *KeyStore) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *KeyStore) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetKeyFileName

`func (o *KeyStore) GetKeyFileName() string`

GetKeyFileName returns the KeyFileName field if non-nil, zero value otherwise.

### GetKeyFileNameOk

`func (o *KeyStore) GetKeyFileNameOk() (*string, bool)`

GetKeyFileNameOk returns a tuple with the KeyFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyFileName

`func (o *KeyStore) SetKeyFileName(v string)`

SetKeyFileName sets KeyFileName field to given value.

### HasKeyFileName

`func (o *KeyStore) HasKeyFileName() bool`

HasKeyFileName returns a boolean if a field has been set.

### GetCapathFileName

`func (o *KeyStore) GetCapathFileName() string`

GetCapathFileName returns the CapathFileName field if non-nil, zero value otherwise.

### GetCapathFileNameOk

`func (o *KeyStore) GetCapathFileNameOk() (*string, bool)`

GetCapathFileNameOk returns a tuple with the CapathFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapathFileName

`func (o *KeyStore) SetCapathFileName(v string)`

SetCapathFileName sets CapathFileName field to given value.

### HasCapathFileName

`func (o *KeyStore) HasCapathFileName() bool`

HasCapathFileName returns a boolean if a field has been set.

### GetExpirationDate

`func (o *KeyStore) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *KeyStore) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *KeyStore) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *KeyStore) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


