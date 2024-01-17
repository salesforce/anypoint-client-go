# CertificatePinning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PerformCertificatePinning** | Pointer to **bool** |  | [optional] 
**CertificatePinset** | Pointer to [**SecretPath**](SecretPath.md) |  | [optional] 

## Methods

### NewCertificatePinning

`func NewCertificatePinning() *CertificatePinning`

NewCertificatePinning instantiates a new CertificatePinning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificatePinningWithDefaults

`func NewCertificatePinningWithDefaults() *CertificatePinning`

NewCertificatePinningWithDefaults instantiates a new CertificatePinning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPerformCertificatePinning

`func (o *CertificatePinning) GetPerformCertificatePinning() bool`

GetPerformCertificatePinning returns the PerformCertificatePinning field if non-nil, zero value otherwise.

### GetPerformCertificatePinningOk

`func (o *CertificatePinning) GetPerformCertificatePinningOk() (*bool, bool)`

GetPerformCertificatePinningOk returns a tuple with the PerformCertificatePinning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformCertificatePinning

`func (o *CertificatePinning) SetPerformCertificatePinning(v bool)`

SetPerformCertificatePinning sets PerformCertificatePinning field to given value.

### HasPerformCertificatePinning

`func (o *CertificatePinning) HasPerformCertificatePinning() bool`

HasPerformCertificatePinning returns a boolean if a field has been set.

### GetCertificatePinset

`func (o *CertificatePinning) GetCertificatePinset() SecretPath`

GetCertificatePinset returns the CertificatePinset field if non-nil, zero value otherwise.

### GetCertificatePinsetOk

`func (o *CertificatePinning) GetCertificatePinsetOk() (*SecretPath, bool)`

GetCertificatePinsetOk returns a tuple with the CertificatePinset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatePinset

`func (o *CertificatePinning) SetCertificatePinset(v SecretPath)`

SetCertificatePinset sets CertificatePinset field to given value.

### HasCertificatePinset

`func (o *CertificatePinning) HasCertificatePinset() bool`

HasCertificatePinset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


