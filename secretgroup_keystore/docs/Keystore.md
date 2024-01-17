# Keystore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpirationDate** | Pointer to **string** |  | [optional] 
**KeystoreFileName** | Pointer to **string** | File name of the keystore that is stored in this secret | [optional] 
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Algorithm** | Pointer to **string** | Algorithm used to create the keystore manager factory which will make use of this keystore | [optional] 
**Details** | Pointer to [**KeystoreDetails**](KeystoreDetails.md) |  | [optional] 
**Type** | Pointer to **string** | Type of keystore supported | [optional] 
**Alias** | Pointer to **string** | The alias name of the entry that contains the certificate | [optional] 
**CertificateFileName** | Pointer to **string** | The file name of the certificate file that is stored in this keystore | [optional] 
**KeyFileName** | Pointer to **string** | The file name of the encrypted private key that is stored in this keystore | [optional] 
**CapathFileName** | Pointer to **string** | The file name of the CA file that is stored in this keystore | [optional] 

## Methods

### NewKeystore

`func NewKeystore() *Keystore`

NewKeystore instantiates a new Keystore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeystoreWithDefaults

`func NewKeystoreWithDefaults() *Keystore`

NewKeystoreWithDefaults instantiates a new Keystore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpirationDate

`func (o *Keystore) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *Keystore) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *Keystore) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *Keystore) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetKeystoreFileName

`func (o *Keystore) GetKeystoreFileName() string`

GetKeystoreFileName returns the KeystoreFileName field if non-nil, zero value otherwise.

### GetKeystoreFileNameOk

`func (o *Keystore) GetKeystoreFileNameOk() (*string, bool)`

GetKeystoreFileNameOk returns a tuple with the KeystoreFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystoreFileName

`func (o *Keystore) SetKeystoreFileName(v string)`

SetKeystoreFileName sets KeystoreFileName field to given value.

### HasKeystoreFileName

`func (o *Keystore) HasKeystoreFileName() bool`

HasKeystoreFileName returns a boolean if a field has been set.

### GetMeta

`func (o *Keystore) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Keystore) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Keystore) SetMeta(v Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Keystore) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetName

`func (o *Keystore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Keystore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Keystore) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Keystore) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAlgorithm

`func (o *Keystore) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *Keystore) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *Keystore) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *Keystore) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetDetails

`func (o *Keystore) GetDetails() KeystoreDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Keystore) GetDetailsOk() (*KeystoreDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Keystore) SetDetails(v KeystoreDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *Keystore) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetType

`func (o *Keystore) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Keystore) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Keystore) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Keystore) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAlias

`func (o *Keystore) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *Keystore) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *Keystore) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *Keystore) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetCertificateFileName

`func (o *Keystore) GetCertificateFileName() string`

GetCertificateFileName returns the CertificateFileName field if non-nil, zero value otherwise.

### GetCertificateFileNameOk

`func (o *Keystore) GetCertificateFileNameOk() (*string, bool)`

GetCertificateFileNameOk returns a tuple with the CertificateFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateFileName

`func (o *Keystore) SetCertificateFileName(v string)`

SetCertificateFileName sets CertificateFileName field to given value.

### HasCertificateFileName

`func (o *Keystore) HasCertificateFileName() bool`

HasCertificateFileName returns a boolean if a field has been set.

### GetKeyFileName

`func (o *Keystore) GetKeyFileName() string`

GetKeyFileName returns the KeyFileName field if non-nil, zero value otherwise.

### GetKeyFileNameOk

`func (o *Keystore) GetKeyFileNameOk() (*string, bool)`

GetKeyFileNameOk returns a tuple with the KeyFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyFileName

`func (o *Keystore) SetKeyFileName(v string)`

SetKeyFileName sets KeyFileName field to given value.

### HasKeyFileName

`func (o *Keystore) HasKeyFileName() bool`

HasKeyFileName returns a boolean if a field has been set.

### GetCapathFileName

`func (o *Keystore) GetCapathFileName() string`

GetCapathFileName returns the CapathFileName field if non-nil, zero value otherwise.

### GetCapathFileNameOk

`func (o *Keystore) GetCapathFileNameOk() (*string, bool)`

GetCapathFileNameOk returns a tuple with the CapathFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapathFileName

`func (o *Keystore) SetCapathFileName(v string)`

SetCapathFileName sets CapathFileName field to given value.

### HasCapathFileName

`func (o *Keystore) HasCapathFileName() bool`

HasCapathFileName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


