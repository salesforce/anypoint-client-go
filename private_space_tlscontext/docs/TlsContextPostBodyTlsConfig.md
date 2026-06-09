# TlsContextPostBodyTlsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyStore** | Pointer to [**TlsContextPostBodyTlsConfigKeyStore**](TlsContextPostBodyTlsConfigKeyStore.md) |  | [optional] 
**TrustStore** | Pointer to [**[]TlsContextPostBodyTrustStorePEM**](TlsContextPostBodyTrustStorePEM.md) |  | [optional] 

## Methods

### NewTlsContextPostBodyTlsConfig

`func NewTlsContextPostBodyTlsConfig() *TlsContextPostBodyTlsConfig`

NewTlsContextPostBodyTlsConfig instantiates a new TlsContextPostBodyTlsConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsContextPostBodyTlsConfigWithDefaults

`func NewTlsContextPostBodyTlsConfigWithDefaults() *TlsContextPostBodyTlsConfig`

NewTlsContextPostBodyTlsConfigWithDefaults instantiates a new TlsContextPostBodyTlsConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyStore

`func (o *TlsContextPostBodyTlsConfig) GetKeyStore() TlsContextPostBodyTlsConfigKeyStore`

GetKeyStore returns the KeyStore field if non-nil, zero value otherwise.

### GetKeyStoreOk

`func (o *TlsContextPostBodyTlsConfig) GetKeyStoreOk() (*TlsContextPostBodyTlsConfigKeyStore, bool)`

GetKeyStoreOk returns a tuple with the KeyStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyStore

`func (o *TlsContextPostBodyTlsConfig) SetKeyStore(v TlsContextPostBodyTlsConfigKeyStore)`

SetKeyStore sets KeyStore field to given value.

### HasKeyStore

`func (o *TlsContextPostBodyTlsConfig) HasKeyStore() bool`

HasKeyStore returns a boolean if a field has been set.

### GetTrustStore

`func (o *TlsContextPostBodyTlsConfig) GetTrustStore() []TlsContextPostBodyTrustStorePEM`

GetTrustStore returns the TrustStore field if non-nil, zero value otherwise.

### GetTrustStoreOk

`func (o *TlsContextPostBodyTlsConfig) GetTrustStoreOk() (*[]TlsContextPostBodyTrustStorePEM, bool)`

GetTrustStoreOk returns a tuple with the TrustStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustStore

`func (o *TlsContextPostBodyTlsConfig) SetTrustStore(v []TlsContextPostBodyTrustStorePEM)`

SetTrustStore sets TrustStore field to given value.

### HasTrustStore

`func (o *TlsContextPostBodyTlsConfig) HasTrustStore() bool`

HasTrustStore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


