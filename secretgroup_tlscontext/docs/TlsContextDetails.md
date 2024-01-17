# TlsContextDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Target** | Pointer to **string** | The target application | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ExpirationDate** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 
**Keystore** | Pointer to [**SecretPath**](SecretPath.md) |  | [optional] 
**Truststore** | Pointer to [**SecretPath**](SecretPath.md) |  | [optional] 
**EnableMutualAuthentication** | Pointer to **bool** | This flag is to enable client authentication | [optional] 
**MutualAuthentication** | Pointer to [**MutualAuthentication**](MutualAuthentication.md) |  | [optional] 
**AcceptableCipherSuites** | Pointer to [**AcceptableCipherSuites**](AcceptableCipherSuites.md) |  | [optional] 
**AcceptableTlsVersions** | Pointer to [**AcceptableTlsVersions**](AcceptableTlsVersions.md) |  | [optional] 
**CipherSuites** | Pointer to **[]string** | List of enabled cipher suites for Mule target | [optional] 
**Insecure** | Pointer to **bool** | Setting this flag to true indicates that certificate validation should not be enforced, i.e. the truststore, even though set, is ignored at runtime. | [optional] [default to false]
**MinTlsVersion** | Pointer to **string** | Minimum TLS version supported. | [optional] 
**MaxTlsVersion** | Pointer to **string** | Maximum TLS version supported. | [optional] 
**AlpnProtocols** | Pointer to **[]string** | supported HTTP versions in the most-to-least preferred order. At least one version must be specified. | [optional] 
**InboundSettings** | Pointer to [**TlsContextDetailsInboundSettings**](TlsContextDetailsInboundSettings.md) |  | [optional] 
**OutboundSettings** | Pointer to [**TlsContextDetailsOutboundSettings**](TlsContextDetailsOutboundSettings.md) |  | [optional] 

## Methods

### NewTlsContextDetails

`func NewTlsContextDetails() *TlsContextDetails`

NewTlsContextDetails instantiates a new TlsContextDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsContextDetailsWithDefaults

`func NewTlsContextDetailsWithDefaults() *TlsContextDetails`

NewTlsContextDetailsWithDefaults instantiates a new TlsContextDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTarget

`func (o *TlsContextDetails) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *TlsContextDetails) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *TlsContextDetails) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *TlsContextDetails) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetName

`func (o *TlsContextDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TlsContextDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TlsContextDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TlsContextDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExpirationDate

`func (o *TlsContextDetails) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *TlsContextDetails) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *TlsContextDetails) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *TlsContextDetails) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetMeta

`func (o *TlsContextDetails) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *TlsContextDetails) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *TlsContextDetails) SetMeta(v Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *TlsContextDetails) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetKeystore

`func (o *TlsContextDetails) GetKeystore() SecretPath`

GetKeystore returns the Keystore field if non-nil, zero value otherwise.

### GetKeystoreOk

`func (o *TlsContextDetails) GetKeystoreOk() (*SecretPath, bool)`

GetKeystoreOk returns a tuple with the Keystore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystore

`func (o *TlsContextDetails) SetKeystore(v SecretPath)`

SetKeystore sets Keystore field to given value.

### HasKeystore

`func (o *TlsContextDetails) HasKeystore() bool`

HasKeystore returns a boolean if a field has been set.

### GetTruststore

`func (o *TlsContextDetails) GetTruststore() SecretPath`

GetTruststore returns the Truststore field if non-nil, zero value otherwise.

### GetTruststoreOk

`func (o *TlsContextDetails) GetTruststoreOk() (*SecretPath, bool)`

GetTruststoreOk returns a tuple with the Truststore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruststore

`func (o *TlsContextDetails) SetTruststore(v SecretPath)`

SetTruststore sets Truststore field to given value.

### HasTruststore

`func (o *TlsContextDetails) HasTruststore() bool`

HasTruststore returns a boolean if a field has been set.

### GetEnableMutualAuthentication

`func (o *TlsContextDetails) GetEnableMutualAuthentication() bool`

GetEnableMutualAuthentication returns the EnableMutualAuthentication field if non-nil, zero value otherwise.

### GetEnableMutualAuthenticationOk

`func (o *TlsContextDetails) GetEnableMutualAuthenticationOk() (*bool, bool)`

GetEnableMutualAuthenticationOk returns a tuple with the EnableMutualAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMutualAuthentication

`func (o *TlsContextDetails) SetEnableMutualAuthentication(v bool)`

SetEnableMutualAuthentication sets EnableMutualAuthentication field to given value.

### HasEnableMutualAuthentication

`func (o *TlsContextDetails) HasEnableMutualAuthentication() bool`

HasEnableMutualAuthentication returns a boolean if a field has been set.

### GetMutualAuthentication

`func (o *TlsContextDetails) GetMutualAuthentication() MutualAuthentication`

GetMutualAuthentication returns the MutualAuthentication field if non-nil, zero value otherwise.

### GetMutualAuthenticationOk

`func (o *TlsContextDetails) GetMutualAuthenticationOk() (*MutualAuthentication, bool)`

GetMutualAuthenticationOk returns a tuple with the MutualAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutualAuthentication

`func (o *TlsContextDetails) SetMutualAuthentication(v MutualAuthentication)`

SetMutualAuthentication sets MutualAuthentication field to given value.

### HasMutualAuthentication

`func (o *TlsContextDetails) HasMutualAuthentication() bool`

HasMutualAuthentication returns a boolean if a field has been set.

### GetAcceptableCipherSuites

`func (o *TlsContextDetails) GetAcceptableCipherSuites() AcceptableCipherSuites`

GetAcceptableCipherSuites returns the AcceptableCipherSuites field if non-nil, zero value otherwise.

### GetAcceptableCipherSuitesOk

`func (o *TlsContextDetails) GetAcceptableCipherSuitesOk() (*AcceptableCipherSuites, bool)`

GetAcceptableCipherSuitesOk returns a tuple with the AcceptableCipherSuites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptableCipherSuites

`func (o *TlsContextDetails) SetAcceptableCipherSuites(v AcceptableCipherSuites)`

SetAcceptableCipherSuites sets AcceptableCipherSuites field to given value.

### HasAcceptableCipherSuites

`func (o *TlsContextDetails) HasAcceptableCipherSuites() bool`

HasAcceptableCipherSuites returns a boolean if a field has been set.

### GetAcceptableTlsVersions

`func (o *TlsContextDetails) GetAcceptableTlsVersions() AcceptableTlsVersions`

GetAcceptableTlsVersions returns the AcceptableTlsVersions field if non-nil, zero value otherwise.

### GetAcceptableTlsVersionsOk

`func (o *TlsContextDetails) GetAcceptableTlsVersionsOk() (*AcceptableTlsVersions, bool)`

GetAcceptableTlsVersionsOk returns a tuple with the AcceptableTlsVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptableTlsVersions

`func (o *TlsContextDetails) SetAcceptableTlsVersions(v AcceptableTlsVersions)`

SetAcceptableTlsVersions sets AcceptableTlsVersions field to given value.

### HasAcceptableTlsVersions

`func (o *TlsContextDetails) HasAcceptableTlsVersions() bool`

HasAcceptableTlsVersions returns a boolean if a field has been set.

### GetCipherSuites

`func (o *TlsContextDetails) GetCipherSuites() []string`

GetCipherSuites returns the CipherSuites field if non-nil, zero value otherwise.

### GetCipherSuitesOk

`func (o *TlsContextDetails) GetCipherSuitesOk() (*[]string, bool)`

GetCipherSuitesOk returns a tuple with the CipherSuites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipherSuites

`func (o *TlsContextDetails) SetCipherSuites(v []string)`

SetCipherSuites sets CipherSuites field to given value.

### HasCipherSuites

`func (o *TlsContextDetails) HasCipherSuites() bool`

HasCipherSuites returns a boolean if a field has been set.

### GetInsecure

`func (o *TlsContextDetails) GetInsecure() bool`

GetInsecure returns the Insecure field if non-nil, zero value otherwise.

### GetInsecureOk

`func (o *TlsContextDetails) GetInsecureOk() (*bool, bool)`

GetInsecureOk returns a tuple with the Insecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecure

`func (o *TlsContextDetails) SetInsecure(v bool)`

SetInsecure sets Insecure field to given value.

### HasInsecure

`func (o *TlsContextDetails) HasInsecure() bool`

HasInsecure returns a boolean if a field has been set.

### GetMinTlsVersion

`func (o *TlsContextDetails) GetMinTlsVersion() string`

GetMinTlsVersion returns the MinTlsVersion field if non-nil, zero value otherwise.

### GetMinTlsVersionOk

`func (o *TlsContextDetails) GetMinTlsVersionOk() (*string, bool)`

GetMinTlsVersionOk returns a tuple with the MinTlsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTlsVersion

`func (o *TlsContextDetails) SetMinTlsVersion(v string)`

SetMinTlsVersion sets MinTlsVersion field to given value.

### HasMinTlsVersion

`func (o *TlsContextDetails) HasMinTlsVersion() bool`

HasMinTlsVersion returns a boolean if a field has been set.

### GetMaxTlsVersion

`func (o *TlsContextDetails) GetMaxTlsVersion() string`

GetMaxTlsVersion returns the MaxTlsVersion field if non-nil, zero value otherwise.

### GetMaxTlsVersionOk

`func (o *TlsContextDetails) GetMaxTlsVersionOk() (*string, bool)`

GetMaxTlsVersionOk returns a tuple with the MaxTlsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTlsVersion

`func (o *TlsContextDetails) SetMaxTlsVersion(v string)`

SetMaxTlsVersion sets MaxTlsVersion field to given value.

### HasMaxTlsVersion

`func (o *TlsContextDetails) HasMaxTlsVersion() bool`

HasMaxTlsVersion returns a boolean if a field has been set.

### GetAlpnProtocols

`func (o *TlsContextDetails) GetAlpnProtocols() []string`

GetAlpnProtocols returns the AlpnProtocols field if non-nil, zero value otherwise.

### GetAlpnProtocolsOk

`func (o *TlsContextDetails) GetAlpnProtocolsOk() (*[]string, bool)`

GetAlpnProtocolsOk returns a tuple with the AlpnProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlpnProtocols

`func (o *TlsContextDetails) SetAlpnProtocols(v []string)`

SetAlpnProtocols sets AlpnProtocols field to given value.

### HasAlpnProtocols

`func (o *TlsContextDetails) HasAlpnProtocols() bool`

HasAlpnProtocols returns a boolean if a field has been set.

### GetInboundSettings

`func (o *TlsContextDetails) GetInboundSettings() TlsContextDetailsInboundSettings`

GetInboundSettings returns the InboundSettings field if non-nil, zero value otherwise.

### GetInboundSettingsOk

`func (o *TlsContextDetails) GetInboundSettingsOk() (*TlsContextDetailsInboundSettings, bool)`

GetInboundSettingsOk returns a tuple with the InboundSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundSettings

`func (o *TlsContextDetails) SetInboundSettings(v TlsContextDetailsInboundSettings)`

SetInboundSettings sets InboundSettings field to given value.

### HasInboundSettings

`func (o *TlsContextDetails) HasInboundSettings() bool`

HasInboundSettings returns a boolean if a field has been set.

### GetOutboundSettings

`func (o *TlsContextDetails) GetOutboundSettings() TlsContextDetailsOutboundSettings`

GetOutboundSettings returns the OutboundSettings field if non-nil, zero value otherwise.

### GetOutboundSettingsOk

`func (o *TlsContextDetails) GetOutboundSettingsOk() (*TlsContextDetailsOutboundSettings, bool)`

GetOutboundSettingsOk returns a tuple with the OutboundSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundSettings

`func (o *TlsContextDetails) SetOutboundSettings(v TlsContextDetailsOutboundSettings)`

SetOutboundSettings sets OutboundSettings field to given value.

### HasOutboundSettings

`func (o *TlsContextDetails) HasOutboundSettings() bool`

HasOutboundSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


