# TlsContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**TrustStore** | Pointer to [**TrustStore**](TrustStore.md) |  | [optional] 
**KeyStore** | Pointer to [**KeyStore**](KeyStore.md) |  | [optional] 
**Ciphers** | Pointer to [**Ciphers**](Ciphers.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewTlsContext

`func NewTlsContext() *TlsContext`

NewTlsContext instantiates a new TlsContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsContextWithDefaults

`func NewTlsContextWithDefaults() *TlsContext`

NewTlsContextWithDefaults instantiates a new TlsContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TlsContext) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TlsContext) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TlsContext) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TlsContext) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TlsContext) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TlsContext) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TlsContext) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TlsContext) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTrustStore

`func (o *TlsContext) GetTrustStore() TrustStore`

GetTrustStore returns the TrustStore field if non-nil, zero value otherwise.

### GetTrustStoreOk

`func (o *TlsContext) GetTrustStoreOk() (*TrustStore, bool)`

GetTrustStoreOk returns a tuple with the TrustStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStore

`func (o *TlsContext) SetTrustStore(v TrustStore)`

SetTrustStore sets TrustStore field to given value.

### HasTrustStore

`func (o *TlsContext) HasTrustStore() bool`

HasTrustStore returns a boolean if a field has been set.

### GetKeyStore

`func (o *TlsContext) GetKeyStore() KeyStore`

GetKeyStore returns the KeyStore field if non-nil, zero value otherwise.

### GetKeyStoreOk

`func (o *TlsContext) GetKeyStoreOk() (*KeyStore, bool)`

GetKeyStoreOk returns a tuple with the KeyStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyStore

`func (o *TlsContext) SetKeyStore(v KeyStore)`

SetKeyStore sets KeyStore field to given value.

### HasKeyStore

`func (o *TlsContext) HasKeyStore() bool`

HasKeyStore returns a boolean if a field has been set.

### GetCiphers

`func (o *TlsContext) GetCiphers() Ciphers`

GetCiphers returns the Ciphers field if non-nil, zero value otherwise.

### GetCiphersOk

`func (o *TlsContext) GetCiphersOk() (*Ciphers, bool)`

GetCiphersOk returns a tuple with the Ciphers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiphers

`func (o *TlsContext) SetCiphers(v Ciphers)`

SetCiphers sets Ciphers field to given value.

### HasCiphers

`func (o *TlsContext) HasCiphers() bool`

HasCiphers returns a boolean if a field has been set.

### GetType

`func (o *TlsContext) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TlsContext) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TlsContext) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TlsContext) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


