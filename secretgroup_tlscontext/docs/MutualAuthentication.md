# MutualAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificatePolicies** | Pointer to **[]string** | List of Object identifier (OID). OIDs are intended to be globally unique. They are formed by taking a unique numeric string (e.g. 1.3.5.7.9.24.68) and adding additional digits in a unique fashion (e.g. 1.3.5.7.9.24.68.1, 1.3.5.7.9.24.68.2, 1.3.5.7.9.24.68.1.1, etc.) An institution will acquire an arc (eg 1.3.5.7.9.24.68) and then extend the arc (called subarcs) as indicated above to create additional OID’s and arcs. There is no limit to the length of an OID, and virtually no computational burden to having a long OID.  | [optional] 
**CertCheckingStrength** | Pointer to **string** | allows application to control if strict or lax certificate checking will be performed during chain-of-trust processing | [optional] 
**VerificationDepth** | Pointer to **int32** | maximum allowed chain length for the certificates | [optional] 
**PerformDomainChecking** | Pointer to **bool** | Whether or not to perform domain checking | [optional] [default to false]
**CertificatePolicyChecking** | Pointer to **bool** | Controls certificate policy processing as defined in RFC 3280, 5280. A certificate can contain zero or more policies. A policy is represented as an object identifier (OID). In an end entity certificate, this policy information indicate the policy under which the certificate has been issued and the purposes for which the certificate may be used. In a CA certificate, this policy information limits the set of policies for certification paths that include this certificate. Applications with specific policy requirements are expected to have a list of those policies that they will accept and to compare the policy OIDs in the certificate to that list. If this extension is critical, the path validation software MUST be able to interpret this extension (including the optional qualifier), or MUST reject the certificate  | [optional] [default to false]
**RequireInitialExplicitPolicy** | Pointer to **bool** | Indicates if the path must be valid for at least one of the certificate policies in the user-initial-policy-set. | [optional] [default to false]
**RevocationChecking** | Pointer to **bool** | Indicates if certificate revocation checking should be enabled or not | [optional] [default to false]
**RevocationCheckingMethod** | Pointer to **string** | Protocol used for certificate revocation checking. Must be set if revocationChecking is set to &#39;true&#39;. | [optional] 
**CrlDistributorConfig** | Pointer to [**SecretPath**](SecretPath.md) |  | [optional] 
**RequireCrlForAllCa** | Pointer to **bool** | Indicates if a valid CRL file must be in effect for every immediate and root Certificate Authority (CA) in the chain-of-trust | [optional] [default to false]
**SendTruststore** | Pointer to **bool** | Should the truststore i.e. trusted certificate authorities be sent to far-end during mutual authentication | [optional] [default to false]
**CertificatePinning** | Pointer to [**CertificatePinning**](CertificatePinning.md) |  | [optional] 
**AuthenticationOverrides** | Pointer to [**AuthenticationOverrides**](AuthenticationOverrides.md) |  | [optional] 

## Methods

### NewMutualAuthentication

`func NewMutualAuthentication() *MutualAuthentication`

NewMutualAuthentication instantiates a new MutualAuthentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMutualAuthenticationWithDefaults

`func NewMutualAuthenticationWithDefaults() *MutualAuthentication`

NewMutualAuthenticationWithDefaults instantiates a new MutualAuthentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificatePolicies

`func (o *MutualAuthentication) GetCertificatePolicies() []string`

GetCertificatePolicies returns the CertificatePolicies field if non-nil, zero value otherwise.

### GetCertificatePoliciesOk

`func (o *MutualAuthentication) GetCertificatePoliciesOk() (*[]string, bool)`

GetCertificatePoliciesOk returns a tuple with the CertificatePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatePolicies

`func (o *MutualAuthentication) SetCertificatePolicies(v []string)`

SetCertificatePolicies sets CertificatePolicies field to given value.

### HasCertificatePolicies

`func (o *MutualAuthentication) HasCertificatePolicies() bool`

HasCertificatePolicies returns a boolean if a field has been set.

### GetCertCheckingStrength

`func (o *MutualAuthentication) GetCertCheckingStrength() string`

GetCertCheckingStrength returns the CertCheckingStrength field if non-nil, zero value otherwise.

### GetCertCheckingStrengthOk

`func (o *MutualAuthentication) GetCertCheckingStrengthOk() (*string, bool)`

GetCertCheckingStrengthOk returns a tuple with the CertCheckingStrength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertCheckingStrength

`func (o *MutualAuthentication) SetCertCheckingStrength(v string)`

SetCertCheckingStrength sets CertCheckingStrength field to given value.

### HasCertCheckingStrength

`func (o *MutualAuthentication) HasCertCheckingStrength() bool`

HasCertCheckingStrength returns a boolean if a field has been set.

### GetVerificationDepth

`func (o *MutualAuthentication) GetVerificationDepth() int32`

GetVerificationDepth returns the VerificationDepth field if non-nil, zero value otherwise.

### GetVerificationDepthOk

`func (o *MutualAuthentication) GetVerificationDepthOk() (*int32, bool)`

GetVerificationDepthOk returns a tuple with the VerificationDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationDepth

`func (o *MutualAuthentication) SetVerificationDepth(v int32)`

SetVerificationDepth sets VerificationDepth field to given value.

### HasVerificationDepth

`func (o *MutualAuthentication) HasVerificationDepth() bool`

HasVerificationDepth returns a boolean if a field has been set.

### GetPerformDomainChecking

`func (o *MutualAuthentication) GetPerformDomainChecking() bool`

GetPerformDomainChecking returns the PerformDomainChecking field if non-nil, zero value otherwise.

### GetPerformDomainCheckingOk

`func (o *MutualAuthentication) GetPerformDomainCheckingOk() (*bool, bool)`

GetPerformDomainCheckingOk returns a tuple with the PerformDomainChecking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformDomainChecking

`func (o *MutualAuthentication) SetPerformDomainChecking(v bool)`

SetPerformDomainChecking sets PerformDomainChecking field to given value.

### HasPerformDomainChecking

`func (o *MutualAuthentication) HasPerformDomainChecking() bool`

HasPerformDomainChecking returns a boolean if a field has been set.

### GetCertificatePolicyChecking

`func (o *MutualAuthentication) GetCertificatePolicyChecking() bool`

GetCertificatePolicyChecking returns the CertificatePolicyChecking field if non-nil, zero value otherwise.

### GetCertificatePolicyCheckingOk

`func (o *MutualAuthentication) GetCertificatePolicyCheckingOk() (*bool, bool)`

GetCertificatePolicyCheckingOk returns a tuple with the CertificatePolicyChecking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatePolicyChecking

`func (o *MutualAuthentication) SetCertificatePolicyChecking(v bool)`

SetCertificatePolicyChecking sets CertificatePolicyChecking field to given value.

### HasCertificatePolicyChecking

`func (o *MutualAuthentication) HasCertificatePolicyChecking() bool`

HasCertificatePolicyChecking returns a boolean if a field has been set.

### GetRequireInitialExplicitPolicy

`func (o *MutualAuthentication) GetRequireInitialExplicitPolicy() bool`

GetRequireInitialExplicitPolicy returns the RequireInitialExplicitPolicy field if non-nil, zero value otherwise.

### GetRequireInitialExplicitPolicyOk

`func (o *MutualAuthentication) GetRequireInitialExplicitPolicyOk() (*bool, bool)`

GetRequireInitialExplicitPolicyOk returns a tuple with the RequireInitialExplicitPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireInitialExplicitPolicy

`func (o *MutualAuthentication) SetRequireInitialExplicitPolicy(v bool)`

SetRequireInitialExplicitPolicy sets RequireInitialExplicitPolicy field to given value.

### HasRequireInitialExplicitPolicy

`func (o *MutualAuthentication) HasRequireInitialExplicitPolicy() bool`

HasRequireInitialExplicitPolicy returns a boolean if a field has been set.

### GetRevocationChecking

`func (o *MutualAuthentication) GetRevocationChecking() bool`

GetRevocationChecking returns the RevocationChecking field if non-nil, zero value otherwise.

### GetRevocationCheckingOk

`func (o *MutualAuthentication) GetRevocationCheckingOk() (*bool, bool)`

GetRevocationCheckingOk returns a tuple with the RevocationChecking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationChecking

`func (o *MutualAuthentication) SetRevocationChecking(v bool)`

SetRevocationChecking sets RevocationChecking field to given value.

### HasRevocationChecking

`func (o *MutualAuthentication) HasRevocationChecking() bool`

HasRevocationChecking returns a boolean if a field has been set.

### GetRevocationCheckingMethod

`func (o *MutualAuthentication) GetRevocationCheckingMethod() string`

GetRevocationCheckingMethod returns the RevocationCheckingMethod field if non-nil, zero value otherwise.

### GetRevocationCheckingMethodOk

`func (o *MutualAuthentication) GetRevocationCheckingMethodOk() (*string, bool)`

GetRevocationCheckingMethodOk returns a tuple with the RevocationCheckingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationCheckingMethod

`func (o *MutualAuthentication) SetRevocationCheckingMethod(v string)`

SetRevocationCheckingMethod sets RevocationCheckingMethod field to given value.

### HasRevocationCheckingMethod

`func (o *MutualAuthentication) HasRevocationCheckingMethod() bool`

HasRevocationCheckingMethod returns a boolean if a field has been set.

### GetCrlDistributorConfig

`func (o *MutualAuthentication) GetCrlDistributorConfig() SecretPath`

GetCrlDistributorConfig returns the CrlDistributorConfig field if non-nil, zero value otherwise.

### GetCrlDistributorConfigOk

`func (o *MutualAuthentication) GetCrlDistributorConfigOk() (*SecretPath, bool)`

GetCrlDistributorConfigOk returns a tuple with the CrlDistributorConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlDistributorConfig

`func (o *MutualAuthentication) SetCrlDistributorConfig(v SecretPath)`

SetCrlDistributorConfig sets CrlDistributorConfig field to given value.

### HasCrlDistributorConfig

`func (o *MutualAuthentication) HasCrlDistributorConfig() bool`

HasCrlDistributorConfig returns a boolean if a field has been set.

### GetRequireCrlForAllCa

`func (o *MutualAuthentication) GetRequireCrlForAllCa() bool`

GetRequireCrlForAllCa returns the RequireCrlForAllCa field if non-nil, zero value otherwise.

### GetRequireCrlForAllCaOk

`func (o *MutualAuthentication) GetRequireCrlForAllCaOk() (*bool, bool)`

GetRequireCrlForAllCaOk returns a tuple with the RequireCrlForAllCa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireCrlForAllCa

`func (o *MutualAuthentication) SetRequireCrlForAllCa(v bool)`

SetRequireCrlForAllCa sets RequireCrlForAllCa field to given value.

### HasRequireCrlForAllCa

`func (o *MutualAuthentication) HasRequireCrlForAllCa() bool`

HasRequireCrlForAllCa returns a boolean if a field has been set.

### GetSendTruststore

`func (o *MutualAuthentication) GetSendTruststore() bool`

GetSendTruststore returns the SendTruststore field if non-nil, zero value otherwise.

### GetSendTruststoreOk

`func (o *MutualAuthentication) GetSendTruststoreOk() (*bool, bool)`

GetSendTruststoreOk returns a tuple with the SendTruststore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendTruststore

`func (o *MutualAuthentication) SetSendTruststore(v bool)`

SetSendTruststore sets SendTruststore field to given value.

### HasSendTruststore

`func (o *MutualAuthentication) HasSendTruststore() bool`

HasSendTruststore returns a boolean if a field has been set.

### GetCertificatePinning

`func (o *MutualAuthentication) GetCertificatePinning() CertificatePinning`

GetCertificatePinning returns the CertificatePinning field if non-nil, zero value otherwise.

### GetCertificatePinningOk

`func (o *MutualAuthentication) GetCertificatePinningOk() (*CertificatePinning, bool)`

GetCertificatePinningOk returns a tuple with the CertificatePinning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatePinning

`func (o *MutualAuthentication) SetCertificatePinning(v CertificatePinning)`

SetCertificatePinning sets CertificatePinning field to given value.

### HasCertificatePinning

`func (o *MutualAuthentication) HasCertificatePinning() bool`

HasCertificatePinning returns a boolean if a field has been set.

### GetAuthenticationOverrides

`func (o *MutualAuthentication) GetAuthenticationOverrides() AuthenticationOverrides`

GetAuthenticationOverrides returns the AuthenticationOverrides field if non-nil, zero value otherwise.

### GetAuthenticationOverridesOk

`func (o *MutualAuthentication) GetAuthenticationOverridesOk() (*AuthenticationOverrides, bool)`

GetAuthenticationOverridesOk returns a tuple with the AuthenticationOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationOverrides

`func (o *MutualAuthentication) SetAuthenticationOverrides(v AuthenticationOverrides)`

SetAuthenticationOverrides sets AuthenticationOverrides field to given value.

### HasAuthenticationOverrides

`func (o *MutualAuthentication) HasAuthenticationOverrides() bool`

HasAuthenticationOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


