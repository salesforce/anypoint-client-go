# AuthenticationOverrides

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateBadFormat** | Pointer to **bool** | Allow processing of certificates with bad format | [optional] [default to false]
**CertificateBadSignature** | Pointer to **bool** | Allow processing of certificates with bad signature | [optional] [default to false]
**CertificateNotYetValid** | Pointer to **bool** | Allow processing of certificates that are not yet valid | [optional] [default to false]
**CertificateHasExpired** | Pointer to **bool** | Allow processing of certificates that are expired | [optional] [default to false]
**AllowSelfSigned** | Pointer to **bool** | Allow self signed certificates | [optional] [default to false]
**CertificateUnresolved** | Pointer to **bool** | Allow unresolved certificates | [optional] [default to false]
**CertificateUntrusted** | Pointer to **bool** | Allow untrusted certificates | [optional] [default to false]
**InvalidCa** | Pointer to **bool** | Allow invalid certificate authority certificates | [optional] [default to false]
**InvalidPurpose** | Pointer to **bool** | Allow certificates with invalid purpose | [optional] [default to false]
**Other** | Pointer to **bool** | Override any miscellaneous error condition encountered | [optional] [default to false]

## Methods

### NewAuthenticationOverrides

`func NewAuthenticationOverrides() *AuthenticationOverrides`

NewAuthenticationOverrides instantiates a new AuthenticationOverrides object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationOverridesWithDefaults

`func NewAuthenticationOverridesWithDefaults() *AuthenticationOverrides`

NewAuthenticationOverridesWithDefaults instantiates a new AuthenticationOverrides object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateBadFormat

`func (o *AuthenticationOverrides) GetCertificateBadFormat() bool`

GetCertificateBadFormat returns the CertificateBadFormat field if non-nil, zero value otherwise.

### GetCertificateBadFormatOk

`func (o *AuthenticationOverrides) GetCertificateBadFormatOk() (*bool, bool)`

GetCertificateBadFormatOk returns a tuple with the CertificateBadFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateBadFormat

`func (o *AuthenticationOverrides) SetCertificateBadFormat(v bool)`

SetCertificateBadFormat sets CertificateBadFormat field to given value.

### HasCertificateBadFormat

`func (o *AuthenticationOverrides) HasCertificateBadFormat() bool`

HasCertificateBadFormat returns a boolean if a field has been set.

### GetCertificateBadSignature

`func (o *AuthenticationOverrides) GetCertificateBadSignature() bool`

GetCertificateBadSignature returns the CertificateBadSignature field if non-nil, zero value otherwise.

### GetCertificateBadSignatureOk

`func (o *AuthenticationOverrides) GetCertificateBadSignatureOk() (*bool, bool)`

GetCertificateBadSignatureOk returns a tuple with the CertificateBadSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateBadSignature

`func (o *AuthenticationOverrides) SetCertificateBadSignature(v bool)`

SetCertificateBadSignature sets CertificateBadSignature field to given value.

### HasCertificateBadSignature

`func (o *AuthenticationOverrides) HasCertificateBadSignature() bool`

HasCertificateBadSignature returns a boolean if a field has been set.

### GetCertificateNotYetValid

`func (o *AuthenticationOverrides) GetCertificateNotYetValid() bool`

GetCertificateNotYetValid returns the CertificateNotYetValid field if non-nil, zero value otherwise.

### GetCertificateNotYetValidOk

`func (o *AuthenticationOverrides) GetCertificateNotYetValidOk() (*bool, bool)`

GetCertificateNotYetValidOk returns a tuple with the CertificateNotYetValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateNotYetValid

`func (o *AuthenticationOverrides) SetCertificateNotYetValid(v bool)`

SetCertificateNotYetValid sets CertificateNotYetValid field to given value.

### HasCertificateNotYetValid

`func (o *AuthenticationOverrides) HasCertificateNotYetValid() bool`

HasCertificateNotYetValid returns a boolean if a field has been set.

### GetCertificateHasExpired

`func (o *AuthenticationOverrides) GetCertificateHasExpired() bool`

GetCertificateHasExpired returns the CertificateHasExpired field if non-nil, zero value otherwise.

### GetCertificateHasExpiredOk

`func (o *AuthenticationOverrides) GetCertificateHasExpiredOk() (*bool, bool)`

GetCertificateHasExpiredOk returns a tuple with the CertificateHasExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateHasExpired

`func (o *AuthenticationOverrides) SetCertificateHasExpired(v bool)`

SetCertificateHasExpired sets CertificateHasExpired field to given value.

### HasCertificateHasExpired

`func (o *AuthenticationOverrides) HasCertificateHasExpired() bool`

HasCertificateHasExpired returns a boolean if a field has been set.

### GetAllowSelfSigned

`func (o *AuthenticationOverrides) GetAllowSelfSigned() bool`

GetAllowSelfSigned returns the AllowSelfSigned field if non-nil, zero value otherwise.

### GetAllowSelfSignedOk

`func (o *AuthenticationOverrides) GetAllowSelfSignedOk() (*bool, bool)`

GetAllowSelfSignedOk returns a tuple with the AllowSelfSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSelfSigned

`func (o *AuthenticationOverrides) SetAllowSelfSigned(v bool)`

SetAllowSelfSigned sets AllowSelfSigned field to given value.

### HasAllowSelfSigned

`func (o *AuthenticationOverrides) HasAllowSelfSigned() bool`

HasAllowSelfSigned returns a boolean if a field has been set.

### GetCertificateUnresolved

`func (o *AuthenticationOverrides) GetCertificateUnresolved() bool`

GetCertificateUnresolved returns the CertificateUnresolved field if non-nil, zero value otherwise.

### GetCertificateUnresolvedOk

`func (o *AuthenticationOverrides) GetCertificateUnresolvedOk() (*bool, bool)`

GetCertificateUnresolvedOk returns a tuple with the CertificateUnresolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateUnresolved

`func (o *AuthenticationOverrides) SetCertificateUnresolved(v bool)`

SetCertificateUnresolved sets CertificateUnresolved field to given value.

### HasCertificateUnresolved

`func (o *AuthenticationOverrides) HasCertificateUnresolved() bool`

HasCertificateUnresolved returns a boolean if a field has been set.

### GetCertificateUntrusted

`func (o *AuthenticationOverrides) GetCertificateUntrusted() bool`

GetCertificateUntrusted returns the CertificateUntrusted field if non-nil, zero value otherwise.

### GetCertificateUntrustedOk

`func (o *AuthenticationOverrides) GetCertificateUntrustedOk() (*bool, bool)`

GetCertificateUntrustedOk returns a tuple with the CertificateUntrusted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateUntrusted

`func (o *AuthenticationOverrides) SetCertificateUntrusted(v bool)`

SetCertificateUntrusted sets CertificateUntrusted field to given value.

### HasCertificateUntrusted

`func (o *AuthenticationOverrides) HasCertificateUntrusted() bool`

HasCertificateUntrusted returns a boolean if a field has been set.

### GetInvalidCa

`func (o *AuthenticationOverrides) GetInvalidCa() bool`

GetInvalidCa returns the InvalidCa field if non-nil, zero value otherwise.

### GetInvalidCaOk

`func (o *AuthenticationOverrides) GetInvalidCaOk() (*bool, bool)`

GetInvalidCaOk returns a tuple with the InvalidCa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidCa

`func (o *AuthenticationOverrides) SetInvalidCa(v bool)`

SetInvalidCa sets InvalidCa field to given value.

### HasInvalidCa

`func (o *AuthenticationOverrides) HasInvalidCa() bool`

HasInvalidCa returns a boolean if a field has been set.

### GetInvalidPurpose

`func (o *AuthenticationOverrides) GetInvalidPurpose() bool`

GetInvalidPurpose returns the InvalidPurpose field if non-nil, zero value otherwise.

### GetInvalidPurposeOk

`func (o *AuthenticationOverrides) GetInvalidPurposeOk() (*bool, bool)`

GetInvalidPurposeOk returns a tuple with the InvalidPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidPurpose

`func (o *AuthenticationOverrides) SetInvalidPurpose(v bool)`

SetInvalidPurpose sets InvalidPurpose field to given value.

### HasInvalidPurpose

`func (o *AuthenticationOverrides) HasInvalidPurpose() bool`

HasInvalidPurpose returns a boolean if a field has been set.

### GetOther

`func (o *AuthenticationOverrides) GetOther() bool`

GetOther returns the Other field if non-nil, zero value otherwise.

### GetOtherOk

`func (o *AuthenticationOverrides) GetOtherOk() (*bool, bool)`

GetOtherOk returns a tuple with the Other field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOther

`func (o *AuthenticationOverrides) SetOther(v bool)`

SetOther sets Other field to given value.

### HasOther

`func (o *AuthenticationOverrides) HasOther() bool`

HasOther returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


