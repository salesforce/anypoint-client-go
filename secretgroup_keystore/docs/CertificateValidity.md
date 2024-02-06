# CertificateValidity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotBefore** | Pointer to **string** |  | [optional] 
**NotAfter** | Pointer to **string** |  | [optional] 

## Methods

### NewCertificateValidity

`func NewCertificateValidity() *CertificateValidity`

NewCertificateValidity instantiates a new CertificateValidity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateValidityWithDefaults

`func NewCertificateValidityWithDefaults() *CertificateValidity`

NewCertificateValidityWithDefaults instantiates a new CertificateValidity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotBefore

`func (o *CertificateValidity) GetNotBefore() string`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *CertificateValidity) GetNotBeforeOk() (*string, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *CertificateValidity) SetNotBefore(v string)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *CertificateValidity) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

### GetNotAfter

`func (o *CertificateValidity) GetNotAfter() string`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *CertificateValidity) GetNotAfterOk() (*string, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *CertificateValidity) SetNotAfter(v string)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *CertificateValidity) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


