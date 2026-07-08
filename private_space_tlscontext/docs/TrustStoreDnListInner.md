# TrustStoreDnListInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Issuer** | Pointer to [**TrustStoreDnListInnerIssuer**](TrustStoreDnListInnerIssuer.md) |  | [optional] 
**Subject** | Pointer to [**TrustStoreDnListInnerIssuer**](TrustStoreDnListInnerIssuer.md) |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**SignatureAlgorithm** | Pointer to **string** |  | [optional] 
**PublicKeyAlgorithm** | Pointer to **string** |  | [optional] 
**BasicConstraints** | Pointer to [**TrustStoreDnListInnerBasicConstraints**](TrustStoreDnListInnerBasicConstraints.md) |  | [optional] 
**Validity** | Pointer to [**TrustStoreDnListInnerValidity**](TrustStoreDnListInnerValidity.md) |  | [optional] 
**KeyUsage** | Pointer to **[]string** |  | [optional] 
**CertificateType** | Pointer to **string** |  | [optional] 

## Methods

### NewTrustStoreDnListInner

`func NewTrustStoreDnListInner() *TrustStoreDnListInner`

NewTrustStoreDnListInner instantiates a new TrustStoreDnListInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustStoreDnListInnerWithDefaults

`func NewTrustStoreDnListInnerWithDefaults() *TrustStoreDnListInner`

NewTrustStoreDnListInnerWithDefaults instantiates a new TrustStoreDnListInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuer

`func (o *TrustStoreDnListInner) GetIssuer() TrustStoreDnListInnerIssuer`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *TrustStoreDnListInner) GetIssuerOk() (*TrustStoreDnListInnerIssuer, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *TrustStoreDnListInner) SetIssuer(v TrustStoreDnListInnerIssuer)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *TrustStoreDnListInner) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetSubject

`func (o *TrustStoreDnListInner) GetSubject() TrustStoreDnListInnerIssuer`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *TrustStoreDnListInner) GetSubjectOk() (*TrustStoreDnListInnerIssuer, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *TrustStoreDnListInner) SetSubject(v TrustStoreDnListInnerIssuer)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *TrustStoreDnListInner) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetVersion

`func (o *TrustStoreDnListInner) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TrustStoreDnListInner) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TrustStoreDnListInner) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *TrustStoreDnListInner) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSerialNumber

`func (o *TrustStoreDnListInner) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *TrustStoreDnListInner) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *TrustStoreDnListInner) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *TrustStoreDnListInner) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSignatureAlgorithm

`func (o *TrustStoreDnListInner) GetSignatureAlgorithm() string`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *TrustStoreDnListInner) GetSignatureAlgorithmOk() (*string, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *TrustStoreDnListInner) SetSignatureAlgorithm(v string)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.

### HasSignatureAlgorithm

`func (o *TrustStoreDnListInner) HasSignatureAlgorithm() bool`

HasSignatureAlgorithm returns a boolean if a field has been set.

### GetPublicKeyAlgorithm

`func (o *TrustStoreDnListInner) GetPublicKeyAlgorithm() string`

GetPublicKeyAlgorithm returns the PublicKeyAlgorithm field if non-nil, zero value otherwise.

### GetPublicKeyAlgorithmOk

`func (o *TrustStoreDnListInner) GetPublicKeyAlgorithmOk() (*string, bool)`

GetPublicKeyAlgorithmOk returns a tuple with the PublicKeyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyAlgorithm

`func (o *TrustStoreDnListInner) SetPublicKeyAlgorithm(v string)`

SetPublicKeyAlgorithm sets PublicKeyAlgorithm field to given value.

### HasPublicKeyAlgorithm

`func (o *TrustStoreDnListInner) HasPublicKeyAlgorithm() bool`

HasPublicKeyAlgorithm returns a boolean if a field has been set.

### GetBasicConstraints

`func (o *TrustStoreDnListInner) GetBasicConstraints() TrustStoreDnListInnerBasicConstraints`

GetBasicConstraints returns the BasicConstraints field if non-nil, zero value otherwise.

### GetBasicConstraintsOk

`func (o *TrustStoreDnListInner) GetBasicConstraintsOk() (*TrustStoreDnListInnerBasicConstraints, bool)`

GetBasicConstraintsOk returns a tuple with the BasicConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicConstraints

`func (o *TrustStoreDnListInner) SetBasicConstraints(v TrustStoreDnListInnerBasicConstraints)`

SetBasicConstraints sets BasicConstraints field to given value.

### HasBasicConstraints

`func (o *TrustStoreDnListInner) HasBasicConstraints() bool`

HasBasicConstraints returns a boolean if a field has been set.

### GetValidity

`func (o *TrustStoreDnListInner) GetValidity() TrustStoreDnListInnerValidity`

GetValidity returns the Validity field if non-nil, zero value otherwise.

### GetValidityOk

`func (o *TrustStoreDnListInner) GetValidityOk() (*TrustStoreDnListInnerValidity, bool)`

GetValidityOk returns a tuple with the Validity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidity

`func (o *TrustStoreDnListInner) SetValidity(v TrustStoreDnListInnerValidity)`

SetValidity sets Validity field to given value.

### HasValidity

`func (o *TrustStoreDnListInner) HasValidity() bool`

HasValidity returns a boolean if a field has been set.

### GetKeyUsage

`func (o *TrustStoreDnListInner) GetKeyUsage() []string`

GetKeyUsage returns the KeyUsage field if non-nil, zero value otherwise.

### GetKeyUsageOk

`func (o *TrustStoreDnListInner) GetKeyUsageOk() (*[]string, bool)`

GetKeyUsageOk returns a tuple with the KeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUsage

`func (o *TrustStoreDnListInner) SetKeyUsage(v []string)`

SetKeyUsage sets KeyUsage field to given value.

### HasKeyUsage

`func (o *TrustStoreDnListInner) HasKeyUsage() bool`

HasKeyUsage returns a boolean if a field has been set.

### GetCertificateType

`func (o *TrustStoreDnListInner) GetCertificateType() string`

GetCertificateType returns the CertificateType field if non-nil, zero value otherwise.

### GetCertificateTypeOk

`func (o *TrustStoreDnListInner) GetCertificateTypeOk() (*string, bool)`

GetCertificateTypeOk returns a tuple with the CertificateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateType

`func (o *TrustStoreDnListInner) SetCertificateType(v string)`

SetCertificateType sets CertificateType field to given value.

### HasCertificateType

`func (o *TrustStoreDnListInner) HasCertificateType() bool`

HasCertificateType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


