# CertificateEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** | Alias associated with the certificate entry | [optional] 
**Certificate** | Pointer to [**CertificateDetails**](CertificateDetails.md) |  | [optional] 

## Methods

### NewCertificateEntry

`func NewCertificateEntry() *CertificateEntry`

NewCertificateEntry instantiates a new CertificateEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateEntryWithDefaults

`func NewCertificateEntryWithDefaults() *CertificateEntry`

NewCertificateEntryWithDefaults instantiates a new CertificateEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *CertificateEntry) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *CertificateEntry) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *CertificateEntry) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *CertificateEntry) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetCertificate

`func (o *CertificateEntry) GetCertificate() CertificateDetails`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CertificateEntry) GetCertificateOk() (*CertificateDetails, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CertificateEntry) SetCertificate(v CertificateDetails)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *CertificateEntry) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


