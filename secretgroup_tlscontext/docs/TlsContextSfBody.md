# TlsContextSfBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpirationDate** | Pointer to **string** | Date on which this secret should expire. If not set, by default, it will be set to one year from the date on which this secret is created/updated. Once the secret expires, a grant can not be requested for it.  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**AcceptableTlsVersions** | Pointer to [**AcceptableTlsVersions**](AcceptableTlsVersions.md) |  | [optional] 
**AcceptableCipherSuites** | Pointer to [**AcceptableCipherSuites**](AcceptableCipherSuites.md) |  | [optional] 
**MutualAuthentication** | Pointer to [**MutualAuthentication**](MutualAuthentication.md) |  | [optional] 
**EnableMutualAuthentication** | Pointer to **bool** | This flag is to enable client authentication. To set this flag to true, both keystore and truststore must be set. | [optional] 
**Target** | Pointer to **string** | The target engine | [optional] 
**Keystore** | Pointer to [**SecretPath**](SecretPath.md) |  | [optional] 
**Truststore** | Pointer to [**SecretPath**](SecretPath.md) |  | [optional] 

## Methods

### NewTlsContextSfBody

`func NewTlsContextSfBody() *TlsContextSfBody`

NewTlsContextSfBody instantiates a new TlsContextSfBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsContextSfBodyWithDefaults

`func NewTlsContextSfBodyWithDefaults() *TlsContextSfBody`

NewTlsContextSfBodyWithDefaults instantiates a new TlsContextSfBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpirationDate

`func (o *TlsContextSfBody) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *TlsContextSfBody) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *TlsContextSfBody) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *TlsContextSfBody) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetName

`func (o *TlsContextSfBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TlsContextSfBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TlsContextSfBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TlsContextSfBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAcceptableTlsVersions

`func (o *TlsContextSfBody) GetAcceptableTlsVersions() AcceptableTlsVersions`

GetAcceptableTlsVersions returns the AcceptableTlsVersions field if non-nil, zero value otherwise.

### GetAcceptableTlsVersionsOk

`func (o *TlsContextSfBody) GetAcceptableTlsVersionsOk() (*AcceptableTlsVersions, bool)`

GetAcceptableTlsVersionsOk returns a tuple with the AcceptableTlsVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptableTlsVersions

`func (o *TlsContextSfBody) SetAcceptableTlsVersions(v AcceptableTlsVersions)`

SetAcceptableTlsVersions sets AcceptableTlsVersions field to given value.

### HasAcceptableTlsVersions

`func (o *TlsContextSfBody) HasAcceptableTlsVersions() bool`

HasAcceptableTlsVersions returns a boolean if a field has been set.

### GetAcceptableCipherSuites

`func (o *TlsContextSfBody) GetAcceptableCipherSuites() AcceptableCipherSuites`

GetAcceptableCipherSuites returns the AcceptableCipherSuites field if non-nil, zero value otherwise.

### GetAcceptableCipherSuitesOk

`func (o *TlsContextSfBody) GetAcceptableCipherSuitesOk() (*AcceptableCipherSuites, bool)`

GetAcceptableCipherSuitesOk returns a tuple with the AcceptableCipherSuites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptableCipherSuites

`func (o *TlsContextSfBody) SetAcceptableCipherSuites(v AcceptableCipherSuites)`

SetAcceptableCipherSuites sets AcceptableCipherSuites field to given value.

### HasAcceptableCipherSuites

`func (o *TlsContextSfBody) HasAcceptableCipherSuites() bool`

HasAcceptableCipherSuites returns a boolean if a field has been set.

### GetMutualAuthentication

`func (o *TlsContextSfBody) GetMutualAuthentication() MutualAuthentication`

GetMutualAuthentication returns the MutualAuthentication field if non-nil, zero value otherwise.

### GetMutualAuthenticationOk

`func (o *TlsContextSfBody) GetMutualAuthenticationOk() (*MutualAuthentication, bool)`

GetMutualAuthenticationOk returns a tuple with the MutualAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutualAuthentication

`func (o *TlsContextSfBody) SetMutualAuthentication(v MutualAuthentication)`

SetMutualAuthentication sets MutualAuthentication field to given value.

### HasMutualAuthentication

`func (o *TlsContextSfBody) HasMutualAuthentication() bool`

HasMutualAuthentication returns a boolean if a field has been set.

### GetEnableMutualAuthentication

`func (o *TlsContextSfBody) GetEnableMutualAuthentication() bool`

GetEnableMutualAuthentication returns the EnableMutualAuthentication field if non-nil, zero value otherwise.

### GetEnableMutualAuthenticationOk

`func (o *TlsContextSfBody) GetEnableMutualAuthenticationOk() (*bool, bool)`

GetEnableMutualAuthenticationOk returns a tuple with the EnableMutualAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMutualAuthentication

`func (o *TlsContextSfBody) SetEnableMutualAuthentication(v bool)`

SetEnableMutualAuthentication sets EnableMutualAuthentication field to given value.

### HasEnableMutualAuthentication

`func (o *TlsContextSfBody) HasEnableMutualAuthentication() bool`

HasEnableMutualAuthentication returns a boolean if a field has been set.

### GetTarget

`func (o *TlsContextSfBody) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *TlsContextSfBody) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *TlsContextSfBody) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *TlsContextSfBody) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetKeystore

`func (o *TlsContextSfBody) GetKeystore() SecretPath`

GetKeystore returns the Keystore field if non-nil, zero value otherwise.

### GetKeystoreOk

`func (o *TlsContextSfBody) GetKeystoreOk() (*SecretPath, bool)`

GetKeystoreOk returns a tuple with the Keystore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystore

`func (o *TlsContextSfBody) SetKeystore(v SecretPath)`

SetKeystore sets Keystore field to given value.

### HasKeystore

`func (o *TlsContextSfBody) HasKeystore() bool`

HasKeystore returns a boolean if a field has been set.

### GetTruststore

`func (o *TlsContextSfBody) GetTruststore() SecretPath`

GetTruststore returns the Truststore field if non-nil, zero value otherwise.

### GetTruststoreOk

`func (o *TlsContextSfBody) GetTruststoreOk() (*SecretPath, bool)`

GetTruststoreOk returns a tuple with the Truststore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruststore

`func (o *TlsContextSfBody) SetTruststore(v SecretPath)`

SetTruststore sets Truststore field to given value.

### HasTruststore

`func (o *TlsContextSfBody) HasTruststore() bool`

HasTruststore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


