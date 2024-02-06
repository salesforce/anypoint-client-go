# TlsContextFlexGatewayBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**ExpirationDate** | Pointer to **string** | Date on which this secret should expire. If not set, by default, it will be set to one year from the date on which this secret is created/updated. Once the secret expires, a grant can not be requested for it.  | [optional] 
**MinTlsVersion** | Pointer to **string** | Minimum TLS version supported. | [optional] 
**MaxTlsVersion** | Pointer to **string** | Maximum TLS version supported. | [optional] 
**Target** | Pointer to **string** | The target engine | [optional] 
**AlpnProtocols** | Pointer to **[]string** | supported HTTP versions in the most-to-least preferred order. At least one version must be specified. | [optional] 
**InboundSettings** | Pointer to [**TlsContextFlexGatewayBodyInboundSettings**](TlsContextFlexGatewayBodyInboundSettings.md) |  | [optional] 
**OutboundSettings** | Pointer to [**TlsContextFlexGatewayBodyOutboundSettings**](TlsContextFlexGatewayBodyOutboundSettings.md) |  | [optional] 
**CipherSuites** | Pointer to **[]string** | List of acceptable cipher suites for Flex Gateway target if min TLS version is &lt; 1.3. If you are are not using the defaults and select individual ciphers, please select ciphers that match the configured keystore to ensure that TLS can setup a connection. For a keystore with an RSA key (the most common type), select ciphers which contain the string RSA (there are some exceptions). If using ECC ciphers, select ciphers which contain the string \&quot;ECDSA\&quot;. TLS standards and documentation can be consulted for more background information.  | [optional] 
**Keystore** | Pointer to [**SecretPath**](SecretPath.md) |  | [optional] 
**Truststore** | Pointer to [**SecretPath**](SecretPath.md) |  | [optional] 

## Methods

### NewTlsContextFlexGatewayBody

`func NewTlsContextFlexGatewayBody() *TlsContextFlexGatewayBody`

NewTlsContextFlexGatewayBody instantiates a new TlsContextFlexGatewayBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsContextFlexGatewayBodyWithDefaults

`func NewTlsContextFlexGatewayBodyWithDefaults() *TlsContextFlexGatewayBody`

NewTlsContextFlexGatewayBodyWithDefaults instantiates a new TlsContextFlexGatewayBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TlsContextFlexGatewayBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TlsContextFlexGatewayBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TlsContextFlexGatewayBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TlsContextFlexGatewayBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExpirationDate

`func (o *TlsContextFlexGatewayBody) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *TlsContextFlexGatewayBody) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *TlsContextFlexGatewayBody) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *TlsContextFlexGatewayBody) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetMinTlsVersion

`func (o *TlsContextFlexGatewayBody) GetMinTlsVersion() string`

GetMinTlsVersion returns the MinTlsVersion field if non-nil, zero value otherwise.

### GetMinTlsVersionOk

`func (o *TlsContextFlexGatewayBody) GetMinTlsVersionOk() (*string, bool)`

GetMinTlsVersionOk returns a tuple with the MinTlsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTlsVersion

`func (o *TlsContextFlexGatewayBody) SetMinTlsVersion(v string)`

SetMinTlsVersion sets MinTlsVersion field to given value.

### HasMinTlsVersion

`func (o *TlsContextFlexGatewayBody) HasMinTlsVersion() bool`

HasMinTlsVersion returns a boolean if a field has been set.

### GetMaxTlsVersion

`func (o *TlsContextFlexGatewayBody) GetMaxTlsVersion() string`

GetMaxTlsVersion returns the MaxTlsVersion field if non-nil, zero value otherwise.

### GetMaxTlsVersionOk

`func (o *TlsContextFlexGatewayBody) GetMaxTlsVersionOk() (*string, bool)`

GetMaxTlsVersionOk returns a tuple with the MaxTlsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTlsVersion

`func (o *TlsContextFlexGatewayBody) SetMaxTlsVersion(v string)`

SetMaxTlsVersion sets MaxTlsVersion field to given value.

### HasMaxTlsVersion

`func (o *TlsContextFlexGatewayBody) HasMaxTlsVersion() bool`

HasMaxTlsVersion returns a boolean if a field has been set.

### GetTarget

`func (o *TlsContextFlexGatewayBody) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *TlsContextFlexGatewayBody) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *TlsContextFlexGatewayBody) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *TlsContextFlexGatewayBody) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetAlpnProtocols

`func (o *TlsContextFlexGatewayBody) GetAlpnProtocols() []string`

GetAlpnProtocols returns the AlpnProtocols field if non-nil, zero value otherwise.

### GetAlpnProtocolsOk

`func (o *TlsContextFlexGatewayBody) GetAlpnProtocolsOk() (*[]string, bool)`

GetAlpnProtocolsOk returns a tuple with the AlpnProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlpnProtocols

`func (o *TlsContextFlexGatewayBody) SetAlpnProtocols(v []string)`

SetAlpnProtocols sets AlpnProtocols field to given value.

### HasAlpnProtocols

`func (o *TlsContextFlexGatewayBody) HasAlpnProtocols() bool`

HasAlpnProtocols returns a boolean if a field has been set.

### GetInboundSettings

`func (o *TlsContextFlexGatewayBody) GetInboundSettings() TlsContextFlexGatewayBodyInboundSettings`

GetInboundSettings returns the InboundSettings field if non-nil, zero value otherwise.

### GetInboundSettingsOk

`func (o *TlsContextFlexGatewayBody) GetInboundSettingsOk() (*TlsContextFlexGatewayBodyInboundSettings, bool)`

GetInboundSettingsOk returns a tuple with the InboundSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundSettings

`func (o *TlsContextFlexGatewayBody) SetInboundSettings(v TlsContextFlexGatewayBodyInboundSettings)`

SetInboundSettings sets InboundSettings field to given value.

### HasInboundSettings

`func (o *TlsContextFlexGatewayBody) HasInboundSettings() bool`

HasInboundSettings returns a boolean if a field has been set.

### GetOutboundSettings

`func (o *TlsContextFlexGatewayBody) GetOutboundSettings() TlsContextFlexGatewayBodyOutboundSettings`

GetOutboundSettings returns the OutboundSettings field if non-nil, zero value otherwise.

### GetOutboundSettingsOk

`func (o *TlsContextFlexGatewayBody) GetOutboundSettingsOk() (*TlsContextFlexGatewayBodyOutboundSettings, bool)`

GetOutboundSettingsOk returns a tuple with the OutboundSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundSettings

`func (o *TlsContextFlexGatewayBody) SetOutboundSettings(v TlsContextFlexGatewayBodyOutboundSettings)`

SetOutboundSettings sets OutboundSettings field to given value.

### HasOutboundSettings

`func (o *TlsContextFlexGatewayBody) HasOutboundSettings() bool`

HasOutboundSettings returns a boolean if a field has been set.

### GetCipherSuites

`func (o *TlsContextFlexGatewayBody) GetCipherSuites() []string`

GetCipherSuites returns the CipherSuites field if non-nil, zero value otherwise.

### GetCipherSuitesOk

`func (o *TlsContextFlexGatewayBody) GetCipherSuitesOk() (*[]string, bool)`

GetCipherSuitesOk returns a tuple with the CipherSuites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipherSuites

`func (o *TlsContextFlexGatewayBody) SetCipherSuites(v []string)`

SetCipherSuites sets CipherSuites field to given value.

### HasCipherSuites

`func (o *TlsContextFlexGatewayBody) HasCipherSuites() bool`

HasCipherSuites returns a boolean if a field has been set.

### GetKeystore

`func (o *TlsContextFlexGatewayBody) GetKeystore() SecretPath`

GetKeystore returns the Keystore field if non-nil, zero value otherwise.

### GetKeystoreOk

`func (o *TlsContextFlexGatewayBody) GetKeystoreOk() (*SecretPath, bool)`

GetKeystoreOk returns a tuple with the Keystore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystore

`func (o *TlsContextFlexGatewayBody) SetKeystore(v SecretPath)`

SetKeystore sets Keystore field to given value.

### HasKeystore

`func (o *TlsContextFlexGatewayBody) HasKeystore() bool`

HasKeystore returns a boolean if a field has been set.

### GetTruststore

`func (o *TlsContextFlexGatewayBody) GetTruststore() SecretPath`

GetTruststore returns the Truststore field if non-nil, zero value otherwise.

### GetTruststoreOk

`func (o *TlsContextFlexGatewayBody) GetTruststoreOk() (*SecretPath, bool)`

GetTruststoreOk returns a tuple with the Truststore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruststore

`func (o *TlsContextFlexGatewayBody) SetTruststore(v SecretPath)`

SetTruststore sets Truststore field to given value.

### HasTruststore

`func (o *TlsContextFlexGatewayBody) HasTruststore() bool`

HasTruststore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


