# TlsContextPostBodyTlsConfigKeyStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** |  | [default to "JKS"]
**Certificate** | **string** |  | 
**Key** | **string** |  | 
**KeyPassphrase** | Pointer to **string** |  | [optional] 
**KeyFileName** | **string** |  | 
**CertificateFileName** | **string** |  | 
**Capath** | Pointer to **string** |  | [optional] 
**CapathFileName** | Pointer to **string** |  | [optional] 
**KeystoreBase64** | **string** |  | 
**StorePassphrase** | **string** |  | 
**Alias** | Pointer to **string** |  | [optional] 
**KeystoreFileName** | Pointer to **string** |  | [optional] 

## Methods

### NewTlsContextPostBodyTlsConfigKeyStore

`func NewTlsContextPostBodyTlsConfigKeyStore(source string, certificate string, key string, keyFileName string, certificateFileName string, keystoreBase64 string, storePassphrase string, ) *TlsContextPostBodyTlsConfigKeyStore`

NewTlsContextPostBodyTlsConfigKeyStore instantiates a new TlsContextPostBodyTlsConfigKeyStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsContextPostBodyTlsConfigKeyStoreWithDefaults

`func NewTlsContextPostBodyTlsConfigKeyStoreWithDefaults() *TlsContextPostBodyTlsConfigKeyStore`

NewTlsContextPostBodyTlsConfigKeyStoreWithDefaults instantiates a new TlsContextPostBodyTlsConfigKeyStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *TlsContextPostBodyTlsConfigKeyStore) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TlsContextPostBodyTlsConfigKeyStore) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TlsContextPostBodyTlsConfigKeyStore) SetSource(v string)`

SetSource sets Source field to given value.


### GetCertificate

`func (o *TlsContextPostBodyTlsConfigKeyStore) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *TlsContextPostBodyTlsConfigKeyStore) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *TlsContextPostBodyTlsConfigKeyStore) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### GetKey

`func (o *TlsContextPostBodyTlsConfigKeyStore) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TlsContextPostBodyTlsConfigKeyStore) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TlsContextPostBodyTlsConfigKeyStore) SetKey(v string)`

SetKey sets Key field to given value.


### GetKeyPassphrase

`func (o *TlsContextPostBodyTlsConfigKeyStore) GetKeyPassphrase() string`

GetKeyPassphrase returns the KeyPassphrase field if non-nil, zero value otherwise.

### GetKeyPassphraseOk

`func (o *TlsContextPostBodyTlsConfigKeyStore) GetKeyPassphraseOk() (*string, bool)`

GetKeyPassphraseOk returns a tuple with the KeyPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPassphrase

`func (o *TlsContextPostBodyTlsConfigKeyStore) SetKeyPassphrase(v string)`

SetKeyPassphrase sets KeyPassphrase field to given value.

### HasKeyPassphrase

`func (o *TlsContextPostBodyTlsConfigKeyStore) HasKeyPassphrase() bool`

HasKeyPassphrase returns a boolean if a field has been set.

### GetKeyFileName

`func (o *TlsContextPostBodyTlsConfigKeyStore) GetKeyFileName() string`

GetKeyFileName returns the KeyFileName field if non-nil, zero value otherwise.

### GetKeyFileNameOk

`func (o *TlsContextPostBodyTlsConfigKeyStore) GetKeyFileNameOk() (*string, bool)`

GetKeyFileNameOk returns a tuple with the KeyFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyFileName

`func (o *TlsContextPostBodyTlsConfigKeyStore) SetKeyFileName(v string)`

SetKeyFileName sets KeyFileName field to given value.


### GetCertificateFileName

`func (o *TlsContextPostBodyTlsConfigKeyStore) GetCertificateFileName() string`

GetCertificateFileName returns the CertificateFileName field if non-nil, zero value otherwise.

### GetCertificateFileNameOk

`func (o *TlsContextPostBodyTlsConfigKeyStore) GetCertificateFileNameOk() (*string, bool)`

GetCertificateFileNameOk returns a tuple with the CertificateFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateFileName

`func (o *TlsContextPostBodyTlsConfigKeyStore) SetCertificateFileName(v string)`

SetCertificateFileName sets CertificateFileName field to given value.


### GetCapath

`func (o *TlsContextPostBodyTlsConfigKeyStore) GetCapath() string`

GetCapath returns the Capath field if non-nil, zero value otherwise.

### GetCapathOk

`func (o *TlsContextPostBodyTlsConfigKeyStore) GetCapathOk() (*string, bool)`

GetCapathOk returns a tuple with the Capath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapath

`func (o *TlsContextPostBodyTlsConfigKeyStore) SetCapath(v string)`

SetCapath sets Capath field to given value.

### HasCapath

`func (o *TlsContextPostBodyTlsConfigKeyStore) HasCapath() bool`

HasCapath returns a boolean if a field has been set.

### GetCapathFileName

`func (o *TlsContextPostBodyTlsConfigKeyStore) GetCapathFileName() string`

GetCapathFileName returns the CapathFileName field if non-nil, zero value otherwise.

### GetCapathFileNameOk

`func (o *TlsContextPostBodyTlsConfigKeyStore) GetCapathFileNameOk() (*string, bool)`

GetCapathFileNameOk returns a tuple with the CapathFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapathFileName

`func (o *TlsContextPostBodyTlsConfigKeyStore) SetCapathFileName(v string)`

SetCapathFileName sets CapathFileName field to given value.

### HasCapathFileName

`func (o *TlsContextPostBodyTlsConfigKeyStore) HasCapathFileName() bool`

HasCapathFileName returns a boolean if a field has been set.

### GetKeystoreBase64

`func (o *TlsContextPostBodyTlsConfigKeyStore) GetKeystoreBase64() string`

GetKeystoreBase64 returns the KeystoreBase64 field if non-nil, zero value otherwise.

### GetKeystoreBase64Ok

`func (o *TlsContextPostBodyTlsConfigKeyStore) GetKeystoreBase64Ok() (*string, bool)`

GetKeystoreBase64Ok returns a tuple with the KeystoreBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystoreBase64

`func (o *TlsContextPostBodyTlsConfigKeyStore) SetKeystoreBase64(v string)`

SetKeystoreBase64 sets KeystoreBase64 field to given value.


### GetStorePassphrase

`func (o *TlsContextPostBodyTlsConfigKeyStore) GetStorePassphrase() string`

GetStorePassphrase returns the StorePassphrase field if non-nil, zero value otherwise.

### GetStorePassphraseOk

`func (o *TlsContextPostBodyTlsConfigKeyStore) GetStorePassphraseOk() (*string, bool)`

GetStorePassphraseOk returns a tuple with the StorePassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorePassphrase

`func (o *TlsContextPostBodyTlsConfigKeyStore) SetStorePassphrase(v string)`

SetStorePassphrase sets StorePassphrase field to given value.


### GetAlias

`func (o *TlsContextPostBodyTlsConfigKeyStore) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *TlsContextPostBodyTlsConfigKeyStore) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *TlsContextPostBodyTlsConfigKeyStore) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *TlsContextPostBodyTlsConfigKeyStore) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetKeystoreFileName

`func (o *TlsContextPostBodyTlsConfigKeyStore) GetKeystoreFileName() string`

GetKeystoreFileName returns the KeystoreFileName field if non-nil, zero value otherwise.

### GetKeystoreFileNameOk

`func (o *TlsContextPostBodyTlsConfigKeyStore) GetKeystoreFileNameOk() (*string, bool)`

GetKeystoreFileNameOk returns a tuple with the KeystoreFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystoreFileName

`func (o *TlsContextPostBodyTlsConfigKeyStore) SetKeystoreFileName(v string)`

SetKeystoreFileName sets KeystoreFileName field to given value.

### HasKeystoreFileName

`func (o *TlsContextPostBodyTlsConfigKeyStore) HasKeystoreFileName() bool`

HasKeystoreFileName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


