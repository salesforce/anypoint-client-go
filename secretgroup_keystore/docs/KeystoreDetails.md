# KeystoreDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to [**CertificateDetails**](CertificateDetails.md) |  | [optional] 
**Capath** | Pointer to [**[]CertificateDetails**](CertificateDetails.md) |  | [optional] 

## Methods

### NewKeystoreDetails

`func NewKeystoreDetails() *KeystoreDetails`

NewKeystoreDetails instantiates a new KeystoreDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeystoreDetailsWithDefaults

`func NewKeystoreDetailsWithDefaults() *KeystoreDetails`

NewKeystoreDetailsWithDefaults instantiates a new KeystoreDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *KeystoreDetails) GetCertificate() CertificateDetails`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *KeystoreDetails) GetCertificateOk() (*CertificateDetails, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *KeystoreDetails) SetCertificate(v CertificateDetails)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *KeystoreDetails) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetCapath

`func (o *KeystoreDetails) GetCapath() []CertificateDetails`

GetCapath returns the Capath field if non-nil, zero value otherwise.

### GetCapathOk

`func (o *KeystoreDetails) GetCapathOk() (*[]CertificateDetails, bool)`

GetCapathOk returns a tuple with the Capath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapath

`func (o *KeystoreDetails) SetCapath(v []CertificateDetails)`

SetCapath sets Capath field to given value.

### HasCapath

`func (o *KeystoreDetails) HasCapath() bool`

HasCapath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


