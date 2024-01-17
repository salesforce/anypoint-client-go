# CrlDistribCfgsDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompleteCrlIssuerUrl** | Pointer to **string** | URL from where complete CRL file is retrieved | [optional] 
**Frequency** | Pointer to **int32** | How frequently should the distributor site be checked for new crl files(in minutes) | [optional] 
**ExpirationDate** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DistributorCertificate** | Pointer to [**SecretPath**](SecretPath.md) |  | [optional] 
**DeltaCrlIssuerUrl** | Pointer to **string** | URL from where the changes in CRL file can be retrieved | [optional] 
**CaCertificate** | Pointer to [**SecretPath**](SecretPath.md) |  | [optional] 

## Methods

### NewCrlDistribCfgsDetails

`func NewCrlDistribCfgsDetails() *CrlDistribCfgsDetails`

NewCrlDistribCfgsDetails instantiates a new CrlDistribCfgsDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrlDistribCfgsDetailsWithDefaults

`func NewCrlDistribCfgsDetailsWithDefaults() *CrlDistribCfgsDetails`

NewCrlDistribCfgsDetailsWithDefaults instantiates a new CrlDistribCfgsDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompleteCrlIssuerUrl

`func (o *CrlDistribCfgsDetails) GetCompleteCrlIssuerUrl() string`

GetCompleteCrlIssuerUrl returns the CompleteCrlIssuerUrl field if non-nil, zero value otherwise.

### GetCompleteCrlIssuerUrlOk

`func (o *CrlDistribCfgsDetails) GetCompleteCrlIssuerUrlOk() (*string, bool)`

GetCompleteCrlIssuerUrlOk returns a tuple with the CompleteCrlIssuerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteCrlIssuerUrl

`func (o *CrlDistribCfgsDetails) SetCompleteCrlIssuerUrl(v string)`

SetCompleteCrlIssuerUrl sets CompleteCrlIssuerUrl field to given value.

### HasCompleteCrlIssuerUrl

`func (o *CrlDistribCfgsDetails) HasCompleteCrlIssuerUrl() bool`

HasCompleteCrlIssuerUrl returns a boolean if a field has been set.

### GetFrequency

`func (o *CrlDistribCfgsDetails) GetFrequency() int32`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *CrlDistribCfgsDetails) GetFrequencyOk() (*int32, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *CrlDistribCfgsDetails) SetFrequency(v int32)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *CrlDistribCfgsDetails) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetExpirationDate

`func (o *CrlDistribCfgsDetails) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *CrlDistribCfgsDetails) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *CrlDistribCfgsDetails) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *CrlDistribCfgsDetails) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetMeta

`func (o *CrlDistribCfgsDetails) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CrlDistribCfgsDetails) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CrlDistribCfgsDetails) SetMeta(v Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CrlDistribCfgsDetails) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetName

`func (o *CrlDistribCfgsDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CrlDistribCfgsDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CrlDistribCfgsDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CrlDistribCfgsDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDistributorCertificate

`func (o *CrlDistribCfgsDetails) GetDistributorCertificate() SecretPath`

GetDistributorCertificate returns the DistributorCertificate field if non-nil, zero value otherwise.

### GetDistributorCertificateOk

`func (o *CrlDistribCfgsDetails) GetDistributorCertificateOk() (*SecretPath, bool)`

GetDistributorCertificateOk returns a tuple with the DistributorCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributorCertificate

`func (o *CrlDistribCfgsDetails) SetDistributorCertificate(v SecretPath)`

SetDistributorCertificate sets DistributorCertificate field to given value.

### HasDistributorCertificate

`func (o *CrlDistribCfgsDetails) HasDistributorCertificate() bool`

HasDistributorCertificate returns a boolean if a field has been set.

### GetDeltaCrlIssuerUrl

`func (o *CrlDistribCfgsDetails) GetDeltaCrlIssuerUrl() string`

GetDeltaCrlIssuerUrl returns the DeltaCrlIssuerUrl field if non-nil, zero value otherwise.

### GetDeltaCrlIssuerUrlOk

`func (o *CrlDistribCfgsDetails) GetDeltaCrlIssuerUrlOk() (*string, bool)`

GetDeltaCrlIssuerUrlOk returns a tuple with the DeltaCrlIssuerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeltaCrlIssuerUrl

`func (o *CrlDistribCfgsDetails) SetDeltaCrlIssuerUrl(v string)`

SetDeltaCrlIssuerUrl sets DeltaCrlIssuerUrl field to given value.

### HasDeltaCrlIssuerUrl

`func (o *CrlDistribCfgsDetails) HasDeltaCrlIssuerUrl() bool`

HasDeltaCrlIssuerUrl returns a boolean if a field has been set.

### GetCaCertificate

`func (o *CrlDistribCfgsDetails) GetCaCertificate() SecretPath`

GetCaCertificate returns the CaCertificate field if non-nil, zero value otherwise.

### GetCaCertificateOk

`func (o *CrlDistribCfgsDetails) GetCaCertificateOk() (*SecretPath, bool)`

GetCaCertificateOk returns a tuple with the CaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificate

`func (o *CrlDistribCfgsDetails) SetCaCertificate(v SecretPath)`

SetCaCertificate sets CaCertificate field to given value.

### HasCaCertificate

`func (o *CrlDistribCfgsDetails) HasCaCertificate() bool`

HasCaCertificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


