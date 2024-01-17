# TlsContextPostBody

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
**CipherSuites** | Pointer to **[]string** | List of acceptable cipher suites for Flex Gateway target if min TLS version is &lt; 1.3. If you are are not using the defaults and select individual ciphers, please select ciphers that match the configured keystore to ensure that TLS can setup a connection. For a keystore with an RSA key (the most common type), select ciphers which contain the string RSA (there are some exceptions). If using ECC ciphers, select ciphers which contain the string \&quot;ECDSA\&quot;. TLS standards and documentation can be consulted for more background information.  | [optional] 
**Insecure** | Pointer to **bool** | Setting this flag to true indicates that certificate validation should not be enforced, i.e. the truststore, even though set, is ignored at runtime. | [optional] 
**MinTlsVersion** | Pointer to **string** | Minimum TLS version supported. | [optional] 
**MaxTlsVersion** | Pointer to **string** | Maximum TLS version supported. | [optional] 
**AlpnProtocols** | Pointer to **[]string** | supported HTTP versions in the most-to-least preferred order. At least one version must be specified. | [optional] 
**InboundSettings** | Pointer to [**TlsContextFlexGatewayBodyInboundSettings**](TlsContextFlexGatewayBodyInboundSettings.md) |  | [optional] 
**OutboundSettings** | Pointer to [**TlsContextFlexGatewayBodyOutboundSettings**](TlsContextFlexGatewayBodyOutboundSettings.md) |  | [optional] 

## Methods

### NewTlsContextPostBody

`func NewTlsContextPostBody() *TlsContextPostBody`

NewTlsContextPostBody instantiates a new TlsContextPostBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsContextPostBodyWithDefaults

`func NewTlsContextPostBodyWithDefaults() *TlsContextPostBody`

NewTlsContextPostBodyWithDefaults instantiates a new TlsContextPostBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpirationDate

`func (o *TlsContextPostBody) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *TlsContextPostBody) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *TlsContextPostBody) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *TlsContextPostBody) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetName

`func (o *TlsContextPostBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TlsContextPostBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TlsContextPostBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TlsContextPostBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAcceptableTlsVersions

`func (o *TlsContextPostBody) GetAcceptableTlsVersions() AcceptableTlsVersions`

GetAcceptableTlsVersions returns the AcceptableTlsVersions field if non-nil, zero value otherwise.

### GetAcceptableTlsVersionsOk

`func (o *TlsContextPostBody) GetAcceptableTlsVersionsOk() (*AcceptableTlsVersions, bool)`

GetAcceptableTlsVersionsOk returns a tuple with the AcceptableTlsVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptableTlsVersions

`func (o *TlsContextPostBody) SetAcceptableTlsVersions(v AcceptableTlsVersions)`

SetAcceptableTlsVersions sets AcceptableTlsVersions field to given value.

### HasAcceptableTlsVersions

`func (o *TlsContextPostBody) HasAcceptableTlsVersions() bool`

HasAcceptableTlsVersions returns a boolean if a field has been set.

### GetAcceptableCipherSuites

`func (o *TlsContextPostBody) GetAcceptableCipherSuites() AcceptableCipherSuites`

GetAcceptableCipherSuites returns the AcceptableCipherSuites field if non-nil, zero value otherwise.

### GetAcceptableCipherSuitesOk

`func (o *TlsContextPostBody) GetAcceptableCipherSuitesOk() (*AcceptableCipherSuites, bool)`

GetAcceptableCipherSuitesOk returns a tuple with the AcceptableCipherSuites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptableCipherSuites

`func (o *TlsContextPostBody) SetAcceptableCipherSuites(v AcceptableCipherSuites)`

SetAcceptableCipherSuites sets AcceptableCipherSuites field to given value.

### HasAcceptableCipherSuites

`func (o *TlsContextPostBody) HasAcceptableCipherSuites() bool`

HasAcceptableCipherSuites returns a boolean if a field has been set.

### GetMutualAuthentication

`func (o *TlsContextPostBody) GetMutualAuthentication() MutualAuthentication`

GetMutualAuthentication returns the MutualAuthentication field if non-nil, zero value otherwise.

### GetMutualAuthenticationOk

`func (o *TlsContextPostBody) GetMutualAuthenticationOk() (*MutualAuthentication, bool)`

GetMutualAuthenticationOk returns a tuple with the MutualAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutualAuthentication

`func (o *TlsContextPostBody) SetMutualAuthentication(v MutualAuthentication)`

SetMutualAuthentication sets MutualAuthentication field to given value.

### HasMutualAuthentication

`func (o *TlsContextPostBody) HasMutualAuthentication() bool`

HasMutualAuthentication returns a boolean if a field has been set.

### GetEnableMutualAuthentication

`func (o *TlsContextPostBody) GetEnableMutualAuthentication() bool`

GetEnableMutualAuthentication returns the EnableMutualAuthentication field if non-nil, zero value otherwise.

### GetEnableMutualAuthenticationOk

`func (o *TlsContextPostBody) GetEnableMutualAuthenticationOk() (*bool, bool)`

GetEnableMutualAuthenticationOk returns a tuple with the EnableMutualAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMutualAuthentication

`func (o *TlsContextPostBody) SetEnableMutualAuthentication(v bool)`

SetEnableMutualAuthentication sets EnableMutualAuthentication field to given value.

### HasEnableMutualAuthentication

`func (o *TlsContextPostBody) HasEnableMutualAuthentication() bool`

HasEnableMutualAuthentication returns a boolean if a field has been set.

### GetTarget

`func (o *TlsContextPostBody) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *TlsContextPostBody) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *TlsContextPostBody) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *TlsContextPostBody) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetKeystore

`func (o *TlsContextPostBody) GetKeystore() SecretPath`

GetKeystore returns the Keystore field if non-nil, zero value otherwise.

### GetKeystoreOk

`func (o *TlsContextPostBody) GetKeystoreOk() (*SecretPath, bool)`

GetKeystoreOk returns a tuple with the Keystore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystore

`func (o *TlsContextPostBody) SetKeystore(v SecretPath)`

SetKeystore sets Keystore field to given value.

### HasKeystore

`func (o *TlsContextPostBody) HasKeystore() bool`

HasKeystore returns a boolean if a field has been set.

### GetTruststore

`func (o *TlsContextPostBody) GetTruststore() SecretPath`

GetTruststore returns the Truststore field if non-nil, zero value otherwise.

### GetTruststoreOk

`func (o *TlsContextPostBody) GetTruststoreOk() (*SecretPath, bool)`

GetTruststoreOk returns a tuple with the Truststore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruststore

`func (o *TlsContextPostBody) SetTruststore(v SecretPath)`

SetTruststore sets Truststore field to given value.

### HasTruststore

`func (o *TlsContextPostBody) HasTruststore() bool`

HasTruststore returns a boolean if a field has been set.

### GetCipherSuites

`func (o *TlsContextPostBody) GetCipherSuites() []string`

GetCipherSuites returns the CipherSuites field if non-nil, zero value otherwise.

### GetCipherSuitesOk

`func (o *TlsContextPostBody) GetCipherSuitesOk() (*[]string, bool)`

GetCipherSuitesOk returns a tuple with the CipherSuites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipherSuites

`func (o *TlsContextPostBody) SetCipherSuites(v []string)`

SetCipherSuites sets CipherSuites field to given value.

### HasCipherSuites

`func (o *TlsContextPostBody) HasCipherSuites() bool`

HasCipherSuites returns a boolean if a field has been set.

### GetInsecure

`func (o *TlsContextPostBody) GetInsecure() bool`

GetInsecure returns the Insecure field if non-nil, zero value otherwise.

### GetInsecureOk

`func (o *TlsContextPostBody) GetInsecureOk() (*bool, bool)`

GetInsecureOk returns a tuple with the Insecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecure

`func (o *TlsContextPostBody) SetInsecure(v bool)`

SetInsecure sets Insecure field to given value.

### HasInsecure

`func (o *TlsContextPostBody) HasInsecure() bool`

HasInsecure returns a boolean if a field has been set.

### GetMinTlsVersion

`func (o *TlsContextPostBody) GetMinTlsVersion() string`

GetMinTlsVersion returns the MinTlsVersion field if non-nil, zero value otherwise.

### GetMinTlsVersionOk

`func (o *TlsContextPostBody) GetMinTlsVersionOk() (*string, bool)`

GetMinTlsVersionOk returns a tuple with the MinTlsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTlsVersion

`func (o *TlsContextPostBody) SetMinTlsVersion(v string)`

SetMinTlsVersion sets MinTlsVersion field to given value.

### HasMinTlsVersion

`func (o *TlsContextPostBody) HasMinTlsVersion() bool`

HasMinTlsVersion returns a boolean if a field has been set.

### GetMaxTlsVersion

`func (o *TlsContextPostBody) GetMaxTlsVersion() string`

GetMaxTlsVersion returns the MaxTlsVersion field if non-nil, zero value otherwise.

### GetMaxTlsVersionOk

`func (o *TlsContextPostBody) GetMaxTlsVersionOk() (*string, bool)`

GetMaxTlsVersionOk returns a tuple with the MaxTlsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTlsVersion

`func (o *TlsContextPostBody) SetMaxTlsVersion(v string)`

SetMaxTlsVersion sets MaxTlsVersion field to given value.

### HasMaxTlsVersion

`func (o *TlsContextPostBody) HasMaxTlsVersion() bool`

HasMaxTlsVersion returns a boolean if a field has been set.

### GetAlpnProtocols

`func (o *TlsContextPostBody) GetAlpnProtocols() []string`

GetAlpnProtocols returns the AlpnProtocols field if non-nil, zero value otherwise.

### GetAlpnProtocolsOk

`func (o *TlsContextPostBody) GetAlpnProtocolsOk() (*[]string, bool)`

GetAlpnProtocolsOk returns a tuple with the AlpnProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlpnProtocols

`func (o *TlsContextPostBody) SetAlpnProtocols(v []string)`

SetAlpnProtocols sets AlpnProtocols field to given value.

### HasAlpnProtocols

`func (o *TlsContextPostBody) HasAlpnProtocols() bool`

HasAlpnProtocols returns a boolean if a field has been set.

### GetInboundSettings

`func (o *TlsContextPostBody) GetInboundSettings() TlsContextFlexGatewayBodyInboundSettings`

GetInboundSettings returns the InboundSettings field if non-nil, zero value otherwise.

### GetInboundSettingsOk

`func (o *TlsContextPostBody) GetInboundSettingsOk() (*TlsContextFlexGatewayBodyInboundSettings, bool)`

GetInboundSettingsOk returns a tuple with the InboundSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundSettings

`func (o *TlsContextPostBody) SetInboundSettings(v TlsContextFlexGatewayBodyInboundSettings)`

SetInboundSettings sets InboundSettings field to given value.

### HasInboundSettings

`func (o *TlsContextPostBody) HasInboundSettings() bool`

HasInboundSettings returns a boolean if a field has been set.

### GetOutboundSettings

`func (o *TlsContextPostBody) GetOutboundSettings() TlsContextFlexGatewayBodyOutboundSettings`

GetOutboundSettings returns the OutboundSettings field if non-nil, zero value otherwise.

### GetOutboundSettingsOk

`func (o *TlsContextPostBody) GetOutboundSettingsOk() (*TlsContextFlexGatewayBodyOutboundSettings, bool)`

GetOutboundSettingsOk returns a tuple with the OutboundSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundSettings

`func (o *TlsContextPostBody) SetOutboundSettings(v TlsContextFlexGatewayBodyOutboundSettings)`

SetOutboundSettings sets OutboundSettings field to given value.

### HasOutboundSettings

`func (o *TlsContextPostBody) HasOutboundSettings() bool`

HasOutboundSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


