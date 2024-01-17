# CrlDistribCfgsReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpirationDate** | Pointer to **string** | Date on which this secret should expire. If not set, by default, it will be set to one year from the date on which this secret is created/updated. Once the secret expires, a grant can not be requested for it.  | [optional] 
**CompleteCrlIssuerUrl** | Pointer to **string** | URL where complete CRL file should be retrieved | [optional] 
**Frequency** | Pointer to **int32** | How frequently should the distributor site be checked for new crl files(in minutes) | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DistributorCertificate** | Pointer to [**SecretPath**](SecretPath.md) |  | [optional] 
**DeltaCrlIssuerUrl** | Pointer to **string** |  | [optional] 
**CaCertificate** | Pointer to [**SecretPath**](SecretPath.md) |  | [optional] 

## Methods

### NewCrlDistribCfgsReqBody

`func NewCrlDistribCfgsReqBody() *CrlDistribCfgsReqBody`

NewCrlDistribCfgsReqBody instantiates a new CrlDistribCfgsReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrlDistribCfgsReqBodyWithDefaults

`func NewCrlDistribCfgsReqBodyWithDefaults() *CrlDistribCfgsReqBody`

NewCrlDistribCfgsReqBodyWithDefaults instantiates a new CrlDistribCfgsReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpirationDate

`func (o *CrlDistribCfgsReqBody) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *CrlDistribCfgsReqBody) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *CrlDistribCfgsReqBody) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *CrlDistribCfgsReqBody) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetCompleteCrlIssuerUrl

`func (o *CrlDistribCfgsReqBody) GetCompleteCrlIssuerUrl() string`

GetCompleteCrlIssuerUrl returns the CompleteCrlIssuerUrl field if non-nil, zero value otherwise.

### GetCompleteCrlIssuerUrlOk

`func (o *CrlDistribCfgsReqBody) GetCompleteCrlIssuerUrlOk() (*string, bool)`

GetCompleteCrlIssuerUrlOk returns a tuple with the CompleteCrlIssuerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteCrlIssuerUrl

`func (o *CrlDistribCfgsReqBody) SetCompleteCrlIssuerUrl(v string)`

SetCompleteCrlIssuerUrl sets CompleteCrlIssuerUrl field to given value.

### HasCompleteCrlIssuerUrl

`func (o *CrlDistribCfgsReqBody) HasCompleteCrlIssuerUrl() bool`

HasCompleteCrlIssuerUrl returns a boolean if a field has been set.

### GetFrequency

`func (o *CrlDistribCfgsReqBody) GetFrequency() int32`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *CrlDistribCfgsReqBody) GetFrequencyOk() (*int32, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *CrlDistribCfgsReqBody) SetFrequency(v int32)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *CrlDistribCfgsReqBody) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetName

`func (o *CrlDistribCfgsReqBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CrlDistribCfgsReqBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CrlDistribCfgsReqBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CrlDistribCfgsReqBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDistributorCertificate

`func (o *CrlDistribCfgsReqBody) GetDistributorCertificate() SecretPath`

GetDistributorCertificate returns the DistributorCertificate field if non-nil, zero value otherwise.

### GetDistributorCertificateOk

`func (o *CrlDistribCfgsReqBody) GetDistributorCertificateOk() (*SecretPath, bool)`

GetDistributorCertificateOk returns a tuple with the DistributorCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributorCertificate

`func (o *CrlDistribCfgsReqBody) SetDistributorCertificate(v SecretPath)`

SetDistributorCertificate sets DistributorCertificate field to given value.

### HasDistributorCertificate

`func (o *CrlDistribCfgsReqBody) HasDistributorCertificate() bool`

HasDistributorCertificate returns a boolean if a field has been set.

### GetDeltaCrlIssuerUrl

`func (o *CrlDistribCfgsReqBody) GetDeltaCrlIssuerUrl() string`

GetDeltaCrlIssuerUrl returns the DeltaCrlIssuerUrl field if non-nil, zero value otherwise.

### GetDeltaCrlIssuerUrlOk

`func (o *CrlDistribCfgsReqBody) GetDeltaCrlIssuerUrlOk() (*string, bool)`

GetDeltaCrlIssuerUrlOk returns a tuple with the DeltaCrlIssuerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeltaCrlIssuerUrl

`func (o *CrlDistribCfgsReqBody) SetDeltaCrlIssuerUrl(v string)`

SetDeltaCrlIssuerUrl sets DeltaCrlIssuerUrl field to given value.

### HasDeltaCrlIssuerUrl

`func (o *CrlDistribCfgsReqBody) HasDeltaCrlIssuerUrl() bool`

HasDeltaCrlIssuerUrl returns a boolean if a field has been set.

### GetCaCertificate

`func (o *CrlDistribCfgsReqBody) GetCaCertificate() SecretPath`

GetCaCertificate returns the CaCertificate field if non-nil, zero value otherwise.

### GetCaCertificateOk

`func (o *CrlDistribCfgsReqBody) GetCaCertificateOk() (*SecretPath, bool)`

GetCaCertificateOk returns a tuple with the CaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificate

`func (o *CrlDistribCfgsReqBody) SetCaCertificate(v SecretPath)`

SetCaCertificate sets CaCertificate field to given value.

### HasCaCertificate

`func (o *CrlDistribCfgsReqBody) HasCaCertificate() bool`

HasCaCertificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


