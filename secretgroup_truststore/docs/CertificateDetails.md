# CertificateDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Issuer** | Pointer to [**IssuerSubject**](IssuerSubject.md) |  | [optional] 
**Subject** | Pointer to [**IssuerSubject**](IssuerSubject.md) |  | [optional] 
**SubjectAlternativeName** | Pointer to **[]string** | Collection of subject alternative names from the SubjectAltName x509 extension | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** | Serial number assigned by the CA to this certificate, in hex format | [optional] 
**SignatureAlgorithm** | Pointer to **string** | Name of the signature algorithm | [optional] 
**PublicKeyAlgorithm** | Pointer to **string** | The standard algorithm name for the public key of this certificate | [optional] 
**BasicConstraints** | Pointer to [**CertificateDetailsBasicConstraints**](CertificateDetailsBasicConstraints.md) |  | [optional] 
**Validity** | Pointer to [**CertificateValidity**](CertificateValidity.md) |  | [optional] 
**KeyUsage** | Pointer to **[]string** | A list of values defining the purpose of the public key i.e. the key usage extensions from this certificate | [optional] 
**ExtendedKeyUsage** | Pointer to **[]string** | A list of values providing details about the extended key usage extensions from this certificate. | [optional] 
**CertificateType** | Pointer to **string** | The type of this certificate | [optional] 

## Methods

### NewCertificateDetails

`func NewCertificateDetails() *CertificateDetails`

NewCertificateDetails instantiates a new CertificateDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateDetailsWithDefaults

`func NewCertificateDetailsWithDefaults() *CertificateDetails`

NewCertificateDetailsWithDefaults instantiates a new CertificateDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuer

`func (o *CertificateDetails) GetIssuer() IssuerSubject`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *CertificateDetails) GetIssuerOk() (*IssuerSubject, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *CertificateDetails) SetIssuer(v IssuerSubject)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *CertificateDetails) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetSubject

`func (o *CertificateDetails) GetSubject() IssuerSubject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CertificateDetails) GetSubjectOk() (*IssuerSubject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CertificateDetails) SetSubject(v IssuerSubject)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CertificateDetails) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetSubjectAlternativeName

`func (o *CertificateDetails) GetSubjectAlternativeName() []string`

GetSubjectAlternativeName returns the SubjectAlternativeName field if non-nil, zero value otherwise.

### GetSubjectAlternativeNameOk

`func (o *CertificateDetails) GetSubjectAlternativeNameOk() (*[]string, bool)`

GetSubjectAlternativeNameOk returns a tuple with the SubjectAlternativeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAlternativeName

`func (o *CertificateDetails) SetSubjectAlternativeName(v []string)`

SetSubjectAlternativeName sets SubjectAlternativeName field to given value.

### HasSubjectAlternativeName

`func (o *CertificateDetails) HasSubjectAlternativeName() bool`

HasSubjectAlternativeName returns a boolean if a field has been set.

### GetVersion

`func (o *CertificateDetails) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CertificateDetails) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CertificateDetails) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CertificateDetails) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSerialNumber

`func (o *CertificateDetails) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *CertificateDetails) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *CertificateDetails) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *CertificateDetails) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSignatureAlgorithm

`func (o *CertificateDetails) GetSignatureAlgorithm() string`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *CertificateDetails) GetSignatureAlgorithmOk() (*string, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *CertificateDetails) SetSignatureAlgorithm(v string)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.

### HasSignatureAlgorithm

`func (o *CertificateDetails) HasSignatureAlgorithm() bool`

HasSignatureAlgorithm returns a boolean if a field has been set.

### GetPublicKeyAlgorithm

`func (o *CertificateDetails) GetPublicKeyAlgorithm() string`

GetPublicKeyAlgorithm returns the PublicKeyAlgorithm field if non-nil, zero value otherwise.

### GetPublicKeyAlgorithmOk

`func (o *CertificateDetails) GetPublicKeyAlgorithmOk() (*string, bool)`

GetPublicKeyAlgorithmOk returns a tuple with the PublicKeyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyAlgorithm

`func (o *CertificateDetails) SetPublicKeyAlgorithm(v string)`

SetPublicKeyAlgorithm sets PublicKeyAlgorithm field to given value.

### HasPublicKeyAlgorithm

`func (o *CertificateDetails) HasPublicKeyAlgorithm() bool`

HasPublicKeyAlgorithm returns a boolean if a field has been set.

### GetBasicConstraints

`func (o *CertificateDetails) GetBasicConstraints() CertificateDetailsBasicConstraints`

GetBasicConstraints returns the BasicConstraints field if non-nil, zero value otherwise.

### GetBasicConstraintsOk

`func (o *CertificateDetails) GetBasicConstraintsOk() (*CertificateDetailsBasicConstraints, bool)`

GetBasicConstraintsOk returns a tuple with the BasicConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicConstraints

`func (o *CertificateDetails) SetBasicConstraints(v CertificateDetailsBasicConstraints)`

SetBasicConstraints sets BasicConstraints field to given value.

### HasBasicConstraints

`func (o *CertificateDetails) HasBasicConstraints() bool`

HasBasicConstraints returns a boolean if a field has been set.

### GetValidity

`func (o *CertificateDetails) GetValidity() CertificateValidity`

GetValidity returns the Validity field if non-nil, zero value otherwise.

### GetValidityOk

`func (o *CertificateDetails) GetValidityOk() (*CertificateValidity, bool)`

GetValidityOk returns a tuple with the Validity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidity

`func (o *CertificateDetails) SetValidity(v CertificateValidity)`

SetValidity sets Validity field to given value.

### HasValidity

`func (o *CertificateDetails) HasValidity() bool`

HasValidity returns a boolean if a field has been set.

### GetKeyUsage

`func (o *CertificateDetails) GetKeyUsage() []string`

GetKeyUsage returns the KeyUsage field if non-nil, zero value otherwise.

### GetKeyUsageOk

`func (o *CertificateDetails) GetKeyUsageOk() (*[]string, bool)`

GetKeyUsageOk returns a tuple with the KeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUsage

`func (o *CertificateDetails) SetKeyUsage(v []string)`

SetKeyUsage sets KeyUsage field to given value.

### HasKeyUsage

`func (o *CertificateDetails) HasKeyUsage() bool`

HasKeyUsage returns a boolean if a field has been set.

### GetExtendedKeyUsage

`func (o *CertificateDetails) GetExtendedKeyUsage() []string`

GetExtendedKeyUsage returns the ExtendedKeyUsage field if non-nil, zero value otherwise.

### GetExtendedKeyUsageOk

`func (o *CertificateDetails) GetExtendedKeyUsageOk() (*[]string, bool)`

GetExtendedKeyUsageOk returns a tuple with the ExtendedKeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedKeyUsage

`func (o *CertificateDetails) SetExtendedKeyUsage(v []string)`

SetExtendedKeyUsage sets ExtendedKeyUsage field to given value.

### HasExtendedKeyUsage

`func (o *CertificateDetails) HasExtendedKeyUsage() bool`

HasExtendedKeyUsage returns a boolean if a field has been set.

### GetCertificateType

`func (o *CertificateDetails) GetCertificateType() string`

GetCertificateType returns the CertificateType field if non-nil, zero value otherwise.

### GetCertificateTypeOk

`func (o *CertificateDetails) GetCertificateTypeOk() (*string, bool)`

GetCertificateTypeOk returns a tuple with the CertificateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateType

`func (o *CertificateDetails) SetCertificateType(v string)`

SetCertificateType sets CertificateType field to given value.

### HasCertificateType

`func (o *CertificateDetails) HasCertificateType() bool`

HasCertificateType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


