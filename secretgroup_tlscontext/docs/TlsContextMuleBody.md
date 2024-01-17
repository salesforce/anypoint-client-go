# TlsContextMuleBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**ExpirationDate** | Pointer to **string** | Date on which this secret should expire. If not set, by default, it will be set to one year from the date on which this secret is created/updated. Once the secret expires, a grant can not be requested for it.  | [optional] 
**Target** | Pointer to **string** | The target engine. The enum type SecurityFabric is used for the Anypoint Security products. | [optional] 
**Keystore** | Pointer to [**SecretPath**](SecretPath.md) |  | [optional] 
**Truststore** | Pointer to [**SecretPath**](SecretPath.md) |  | [optional] 
**CipherSuites** | Pointer to **[]string** | List of enabled cipher suites for Mule target | [optional] 
**AcceptableTlsVersions** | Pointer to [**AcceptableTlsVersions**](AcceptableTlsVersions.md) |  | [optional] 
**Insecure** | Pointer to **bool** | Setting this flag to true indicates that certificate validation should not be enforced, i.e. the truststore, even though set, is ignored at runtime. | [optional] 

## Methods

### NewTlsContextMuleBody

`func NewTlsContextMuleBody() *TlsContextMuleBody`

NewTlsContextMuleBody instantiates a new TlsContextMuleBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsContextMuleBodyWithDefaults

`func NewTlsContextMuleBodyWithDefaults() *TlsContextMuleBody`

NewTlsContextMuleBodyWithDefaults instantiates a new TlsContextMuleBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TlsContextMuleBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TlsContextMuleBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TlsContextMuleBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TlsContextMuleBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExpirationDate

`func (o *TlsContextMuleBody) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *TlsContextMuleBody) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *TlsContextMuleBody) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *TlsContextMuleBody) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetTarget

`func (o *TlsContextMuleBody) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *TlsContextMuleBody) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *TlsContextMuleBody) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *TlsContextMuleBody) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetKeystore

`func (o *TlsContextMuleBody) GetKeystore() SecretPath`

GetKeystore returns the Keystore field if non-nil, zero value otherwise.

### GetKeystoreOk

`func (o *TlsContextMuleBody) GetKeystoreOk() (*SecretPath, bool)`

GetKeystoreOk returns a tuple with the Keystore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystore

`func (o *TlsContextMuleBody) SetKeystore(v SecretPath)`

SetKeystore sets Keystore field to given value.

### HasKeystore

`func (o *TlsContextMuleBody) HasKeystore() bool`

HasKeystore returns a boolean if a field has been set.

### GetTruststore

`func (o *TlsContextMuleBody) GetTruststore() SecretPath`

GetTruststore returns the Truststore field if non-nil, zero value otherwise.

### GetTruststoreOk

`func (o *TlsContextMuleBody) GetTruststoreOk() (*SecretPath, bool)`

GetTruststoreOk returns a tuple with the Truststore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruststore

`func (o *TlsContextMuleBody) SetTruststore(v SecretPath)`

SetTruststore sets Truststore field to given value.

### HasTruststore

`func (o *TlsContextMuleBody) HasTruststore() bool`

HasTruststore returns a boolean if a field has been set.

### GetCipherSuites

`func (o *TlsContextMuleBody) GetCipherSuites() []string`

GetCipherSuites returns the CipherSuites field if non-nil, zero value otherwise.

### GetCipherSuitesOk

`func (o *TlsContextMuleBody) GetCipherSuitesOk() (*[]string, bool)`

GetCipherSuitesOk returns a tuple with the CipherSuites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipherSuites

`func (o *TlsContextMuleBody) SetCipherSuites(v []string)`

SetCipherSuites sets CipherSuites field to given value.

### HasCipherSuites

`func (o *TlsContextMuleBody) HasCipherSuites() bool`

HasCipherSuites returns a boolean if a field has been set.

### GetAcceptableTlsVersions

`func (o *TlsContextMuleBody) GetAcceptableTlsVersions() AcceptableTlsVersions`

GetAcceptableTlsVersions returns the AcceptableTlsVersions field if non-nil, zero value otherwise.

### GetAcceptableTlsVersionsOk

`func (o *TlsContextMuleBody) GetAcceptableTlsVersionsOk() (*AcceptableTlsVersions, bool)`

GetAcceptableTlsVersionsOk returns a tuple with the AcceptableTlsVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptableTlsVersions

`func (o *TlsContextMuleBody) SetAcceptableTlsVersions(v AcceptableTlsVersions)`

SetAcceptableTlsVersions sets AcceptableTlsVersions field to given value.

### HasAcceptableTlsVersions

`func (o *TlsContextMuleBody) HasAcceptableTlsVersions() bool`

HasAcceptableTlsVersions returns a boolean if a field has been set.

### GetInsecure

`func (o *TlsContextMuleBody) GetInsecure() bool`

GetInsecure returns the Insecure field if non-nil, zero value otherwise.

### GetInsecureOk

`func (o *TlsContextMuleBody) GetInsecureOk() (*bool, bool)`

GetInsecureOk returns a tuple with the Insecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecure

`func (o *TlsContextMuleBody) SetInsecure(v bool)`

SetInsecure sets Insecure field to given value.

### HasInsecure

`func (o *TlsContextMuleBody) HasInsecure() bool`

HasInsecure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


