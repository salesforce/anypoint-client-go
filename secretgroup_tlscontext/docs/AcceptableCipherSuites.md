# AcceptableCipherSuites

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aes128GcmSha256** | Pointer to **bool** | Allowed to be enabled only if tlsV1Dot2 is enabled. | [optional] [default to false]
**Aes128Sha256** | Pointer to **bool** | Allowed to be enabled only if tlsV1Dot2 is enabled. | [optional] [default to false]
**Aes256GcmSha384** | Pointer to **bool** | Allowed to be enabled only if tlsV1Dot2 is enabled. | [optional] [default to false]
**Aes256Sha256** | Pointer to **bool** | Allowed to be enabled only if tlsV1Dot2 is enabled. | [optional] [default to false]
**DheRsaAes128GcmSha256** | Pointer to **bool** | Allowed to be enabled only if tlsV1Dot2 is enabled. | [optional] [default to false]
**DheRsaAes128Sha256** | Pointer to **bool** | Allowed to be enabled only if tlsV1Dot2 is enabled. | [optional] [default to false]
**DheRsaAes256GcmSha384** | Pointer to **bool** | Allowed to be enabled only if tlsV1Dot2 is enabled. | [optional] [default to false]
**DheRsaAes256Sha256** | Pointer to **bool** | Allowed to be enabled only if tlsV1Dot2 is enabled. | [optional] [default to false]
**EcdheEcdsaAes128GcmSha256** | Pointer to **bool** | Allowed to be enabled only if tlsV1Dot2 is enabled. | [optional] [default to false]
**EcdheEcdsaAes128Sha1** | Pointer to **bool** | Allowed to be enabled only if tlsV1Dot2 is enabled. | [optional] [default to false]
**EcdheEcdsaAes256GcmSha384** | Pointer to **bool** | Allowed to be enabled only if tlsV1Dot2 is enabled. | [optional] [default to false]
**EcdheEcdsaAes256Sha1** | Pointer to **bool** | Allowed to be enabled only if tlsV1Dot2 is enabled. | [optional] [default to false]
**EcdheRsaAes128GcmSha256** | Pointer to **bool** | Allowed to be enabled only if tlsV1Dot2 is enabled. | [optional] [default to false]
**EcdheRsaAes128Sha1** | Pointer to **bool** | Allowed to be enabled only if tlsV1Dot2 is enabled. | [optional] [default to false]
**EcdheRsaAes256GcmSha384** | Pointer to **bool** | Allowed to be enabled only if tlsV1Dot2 is enabled. | [optional] [default to false]
**EcdheRsaAes256Sha1** | Pointer to **bool** | Allowed to be enabled only if tlsV1Dot2 is enabled. | [optional] [default to false]
**EcdheEcdsaChacha20Poly1305** | Pointer to **bool** | Allowed to be enabled only if tlsV1Dot2 is enabled. | [optional] [default to false]
**EcdheRsaChacha20Poly1305** | Pointer to **bool** | Allowed to be enabled only if tlsV1Dot2 is enabled. | [optional] [default to false]
**DheRsaChacha20Poly1305** | Pointer to **bool** | Allowed to be enabled only if tlsV1Dot2 is enabled. | [optional] [default to false]
**TlsAes256GcmSha384** | Pointer to **bool** | Allowed to be enabled only if tlsV1Dot2 is enabled. | [optional] [default to false]
**TlsChacha20Poly1305Sha256** | Pointer to **bool** | Allowed to be enabled only if tlsV1Dot2 is enabled. | [optional] [default to false]
**TlsAes128GcmSha256** | Pointer to **bool** | Allowed to be enabled only if tlsV1Dot2 is enabled. | [optional] [default to false]

## Methods

### NewAcceptableCipherSuites

`func NewAcceptableCipherSuites() *AcceptableCipherSuites`

NewAcceptableCipherSuites instantiates a new AcceptableCipherSuites object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcceptableCipherSuitesWithDefaults

`func NewAcceptableCipherSuitesWithDefaults() *AcceptableCipherSuites`

NewAcceptableCipherSuitesWithDefaults instantiates a new AcceptableCipherSuites object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAes128GcmSha256

`func (o *AcceptableCipherSuites) GetAes128GcmSha256() bool`

GetAes128GcmSha256 returns the Aes128GcmSha256 field if non-nil, zero value otherwise.

### GetAes128GcmSha256Ok

`func (o *AcceptableCipherSuites) GetAes128GcmSha256Ok() (*bool, bool)`

GetAes128GcmSha256Ok returns a tuple with the Aes128GcmSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAes128GcmSha256

`func (o *AcceptableCipherSuites) SetAes128GcmSha256(v bool)`

SetAes128GcmSha256 sets Aes128GcmSha256 field to given value.

### HasAes128GcmSha256

`func (o *AcceptableCipherSuites) HasAes128GcmSha256() bool`

HasAes128GcmSha256 returns a boolean if a field has been set.

### GetAes128Sha256

`func (o *AcceptableCipherSuites) GetAes128Sha256() bool`

GetAes128Sha256 returns the Aes128Sha256 field if non-nil, zero value otherwise.

### GetAes128Sha256Ok

`func (o *AcceptableCipherSuites) GetAes128Sha256Ok() (*bool, bool)`

GetAes128Sha256Ok returns a tuple with the Aes128Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAes128Sha256

`func (o *AcceptableCipherSuites) SetAes128Sha256(v bool)`

SetAes128Sha256 sets Aes128Sha256 field to given value.

### HasAes128Sha256

`func (o *AcceptableCipherSuites) HasAes128Sha256() bool`

HasAes128Sha256 returns a boolean if a field has been set.

### GetAes256GcmSha384

`func (o *AcceptableCipherSuites) GetAes256GcmSha384() bool`

GetAes256GcmSha384 returns the Aes256GcmSha384 field if non-nil, zero value otherwise.

### GetAes256GcmSha384Ok

`func (o *AcceptableCipherSuites) GetAes256GcmSha384Ok() (*bool, bool)`

GetAes256GcmSha384Ok returns a tuple with the Aes256GcmSha384 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAes256GcmSha384

`func (o *AcceptableCipherSuites) SetAes256GcmSha384(v bool)`

SetAes256GcmSha384 sets Aes256GcmSha384 field to given value.

### HasAes256GcmSha384

`func (o *AcceptableCipherSuites) HasAes256GcmSha384() bool`

HasAes256GcmSha384 returns a boolean if a field has been set.

### GetAes256Sha256

`func (o *AcceptableCipherSuites) GetAes256Sha256() bool`

GetAes256Sha256 returns the Aes256Sha256 field if non-nil, zero value otherwise.

### GetAes256Sha256Ok

`func (o *AcceptableCipherSuites) GetAes256Sha256Ok() (*bool, bool)`

GetAes256Sha256Ok returns a tuple with the Aes256Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAes256Sha256

`func (o *AcceptableCipherSuites) SetAes256Sha256(v bool)`

SetAes256Sha256 sets Aes256Sha256 field to given value.

### HasAes256Sha256

`func (o *AcceptableCipherSuites) HasAes256Sha256() bool`

HasAes256Sha256 returns a boolean if a field has been set.

### GetDheRsaAes128GcmSha256

`func (o *AcceptableCipherSuites) GetDheRsaAes128GcmSha256() bool`

GetDheRsaAes128GcmSha256 returns the DheRsaAes128GcmSha256 field if non-nil, zero value otherwise.

### GetDheRsaAes128GcmSha256Ok

`func (o *AcceptableCipherSuites) GetDheRsaAes128GcmSha256Ok() (*bool, bool)`

GetDheRsaAes128GcmSha256Ok returns a tuple with the DheRsaAes128GcmSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDheRsaAes128GcmSha256

`func (o *AcceptableCipherSuites) SetDheRsaAes128GcmSha256(v bool)`

SetDheRsaAes128GcmSha256 sets DheRsaAes128GcmSha256 field to given value.

### HasDheRsaAes128GcmSha256

`func (o *AcceptableCipherSuites) HasDheRsaAes128GcmSha256() bool`

HasDheRsaAes128GcmSha256 returns a boolean if a field has been set.

### GetDheRsaAes128Sha256

`func (o *AcceptableCipherSuites) GetDheRsaAes128Sha256() bool`

GetDheRsaAes128Sha256 returns the DheRsaAes128Sha256 field if non-nil, zero value otherwise.

### GetDheRsaAes128Sha256Ok

`func (o *AcceptableCipherSuites) GetDheRsaAes128Sha256Ok() (*bool, bool)`

GetDheRsaAes128Sha256Ok returns a tuple with the DheRsaAes128Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDheRsaAes128Sha256

`func (o *AcceptableCipherSuites) SetDheRsaAes128Sha256(v bool)`

SetDheRsaAes128Sha256 sets DheRsaAes128Sha256 field to given value.

### HasDheRsaAes128Sha256

`func (o *AcceptableCipherSuites) HasDheRsaAes128Sha256() bool`

HasDheRsaAes128Sha256 returns a boolean if a field has been set.

### GetDheRsaAes256GcmSha384

`func (o *AcceptableCipherSuites) GetDheRsaAes256GcmSha384() bool`

GetDheRsaAes256GcmSha384 returns the DheRsaAes256GcmSha384 field if non-nil, zero value otherwise.

### GetDheRsaAes256GcmSha384Ok

`func (o *AcceptableCipherSuites) GetDheRsaAes256GcmSha384Ok() (*bool, bool)`

GetDheRsaAes256GcmSha384Ok returns a tuple with the DheRsaAes256GcmSha384 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDheRsaAes256GcmSha384

`func (o *AcceptableCipherSuites) SetDheRsaAes256GcmSha384(v bool)`

SetDheRsaAes256GcmSha384 sets DheRsaAes256GcmSha384 field to given value.

### HasDheRsaAes256GcmSha384

`func (o *AcceptableCipherSuites) HasDheRsaAes256GcmSha384() bool`

HasDheRsaAes256GcmSha384 returns a boolean if a field has been set.

### GetDheRsaAes256Sha256

`func (o *AcceptableCipherSuites) GetDheRsaAes256Sha256() bool`

GetDheRsaAes256Sha256 returns the DheRsaAes256Sha256 field if non-nil, zero value otherwise.

### GetDheRsaAes256Sha256Ok

`func (o *AcceptableCipherSuites) GetDheRsaAes256Sha256Ok() (*bool, bool)`

GetDheRsaAes256Sha256Ok returns a tuple with the DheRsaAes256Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDheRsaAes256Sha256

`func (o *AcceptableCipherSuites) SetDheRsaAes256Sha256(v bool)`

SetDheRsaAes256Sha256 sets DheRsaAes256Sha256 field to given value.

### HasDheRsaAes256Sha256

`func (o *AcceptableCipherSuites) HasDheRsaAes256Sha256() bool`

HasDheRsaAes256Sha256 returns a boolean if a field has been set.

### GetEcdheEcdsaAes128GcmSha256

`func (o *AcceptableCipherSuites) GetEcdheEcdsaAes128GcmSha256() bool`

GetEcdheEcdsaAes128GcmSha256 returns the EcdheEcdsaAes128GcmSha256 field if non-nil, zero value otherwise.

### GetEcdheEcdsaAes128GcmSha256Ok

`func (o *AcceptableCipherSuites) GetEcdheEcdsaAes128GcmSha256Ok() (*bool, bool)`

GetEcdheEcdsaAes128GcmSha256Ok returns a tuple with the EcdheEcdsaAes128GcmSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcdheEcdsaAes128GcmSha256

`func (o *AcceptableCipherSuites) SetEcdheEcdsaAes128GcmSha256(v bool)`

SetEcdheEcdsaAes128GcmSha256 sets EcdheEcdsaAes128GcmSha256 field to given value.

### HasEcdheEcdsaAes128GcmSha256

`func (o *AcceptableCipherSuites) HasEcdheEcdsaAes128GcmSha256() bool`

HasEcdheEcdsaAes128GcmSha256 returns a boolean if a field has been set.

### GetEcdheEcdsaAes128Sha1

`func (o *AcceptableCipherSuites) GetEcdheEcdsaAes128Sha1() bool`

GetEcdheEcdsaAes128Sha1 returns the EcdheEcdsaAes128Sha1 field if non-nil, zero value otherwise.

### GetEcdheEcdsaAes128Sha1Ok

`func (o *AcceptableCipherSuites) GetEcdheEcdsaAes128Sha1Ok() (*bool, bool)`

GetEcdheEcdsaAes128Sha1Ok returns a tuple with the EcdheEcdsaAes128Sha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcdheEcdsaAes128Sha1

`func (o *AcceptableCipherSuites) SetEcdheEcdsaAes128Sha1(v bool)`

SetEcdheEcdsaAes128Sha1 sets EcdheEcdsaAes128Sha1 field to given value.

### HasEcdheEcdsaAes128Sha1

`func (o *AcceptableCipherSuites) HasEcdheEcdsaAes128Sha1() bool`

HasEcdheEcdsaAes128Sha1 returns a boolean if a field has been set.

### GetEcdheEcdsaAes256GcmSha384

`func (o *AcceptableCipherSuites) GetEcdheEcdsaAes256GcmSha384() bool`

GetEcdheEcdsaAes256GcmSha384 returns the EcdheEcdsaAes256GcmSha384 field if non-nil, zero value otherwise.

### GetEcdheEcdsaAes256GcmSha384Ok

`func (o *AcceptableCipherSuites) GetEcdheEcdsaAes256GcmSha384Ok() (*bool, bool)`

GetEcdheEcdsaAes256GcmSha384Ok returns a tuple with the EcdheEcdsaAes256GcmSha384 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcdheEcdsaAes256GcmSha384

`func (o *AcceptableCipherSuites) SetEcdheEcdsaAes256GcmSha384(v bool)`

SetEcdheEcdsaAes256GcmSha384 sets EcdheEcdsaAes256GcmSha384 field to given value.

### HasEcdheEcdsaAes256GcmSha384

`func (o *AcceptableCipherSuites) HasEcdheEcdsaAes256GcmSha384() bool`

HasEcdheEcdsaAes256GcmSha384 returns a boolean if a field has been set.

### GetEcdheEcdsaAes256Sha1

`func (o *AcceptableCipherSuites) GetEcdheEcdsaAes256Sha1() bool`

GetEcdheEcdsaAes256Sha1 returns the EcdheEcdsaAes256Sha1 field if non-nil, zero value otherwise.

### GetEcdheEcdsaAes256Sha1Ok

`func (o *AcceptableCipherSuites) GetEcdheEcdsaAes256Sha1Ok() (*bool, bool)`

GetEcdheEcdsaAes256Sha1Ok returns a tuple with the EcdheEcdsaAes256Sha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcdheEcdsaAes256Sha1

`func (o *AcceptableCipherSuites) SetEcdheEcdsaAes256Sha1(v bool)`

SetEcdheEcdsaAes256Sha1 sets EcdheEcdsaAes256Sha1 field to given value.

### HasEcdheEcdsaAes256Sha1

`func (o *AcceptableCipherSuites) HasEcdheEcdsaAes256Sha1() bool`

HasEcdheEcdsaAes256Sha1 returns a boolean if a field has been set.

### GetEcdheRsaAes128GcmSha256

`func (o *AcceptableCipherSuites) GetEcdheRsaAes128GcmSha256() bool`

GetEcdheRsaAes128GcmSha256 returns the EcdheRsaAes128GcmSha256 field if non-nil, zero value otherwise.

### GetEcdheRsaAes128GcmSha256Ok

`func (o *AcceptableCipherSuites) GetEcdheRsaAes128GcmSha256Ok() (*bool, bool)`

GetEcdheRsaAes128GcmSha256Ok returns a tuple with the EcdheRsaAes128GcmSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcdheRsaAes128GcmSha256

`func (o *AcceptableCipherSuites) SetEcdheRsaAes128GcmSha256(v bool)`

SetEcdheRsaAes128GcmSha256 sets EcdheRsaAes128GcmSha256 field to given value.

### HasEcdheRsaAes128GcmSha256

`func (o *AcceptableCipherSuites) HasEcdheRsaAes128GcmSha256() bool`

HasEcdheRsaAes128GcmSha256 returns a boolean if a field has been set.

### GetEcdheRsaAes128Sha1

`func (o *AcceptableCipherSuites) GetEcdheRsaAes128Sha1() bool`

GetEcdheRsaAes128Sha1 returns the EcdheRsaAes128Sha1 field if non-nil, zero value otherwise.

### GetEcdheRsaAes128Sha1Ok

`func (o *AcceptableCipherSuites) GetEcdheRsaAes128Sha1Ok() (*bool, bool)`

GetEcdheRsaAes128Sha1Ok returns a tuple with the EcdheRsaAes128Sha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcdheRsaAes128Sha1

`func (o *AcceptableCipherSuites) SetEcdheRsaAes128Sha1(v bool)`

SetEcdheRsaAes128Sha1 sets EcdheRsaAes128Sha1 field to given value.

### HasEcdheRsaAes128Sha1

`func (o *AcceptableCipherSuites) HasEcdheRsaAes128Sha1() bool`

HasEcdheRsaAes128Sha1 returns a boolean if a field has been set.

### GetEcdheRsaAes256GcmSha384

`func (o *AcceptableCipherSuites) GetEcdheRsaAes256GcmSha384() bool`

GetEcdheRsaAes256GcmSha384 returns the EcdheRsaAes256GcmSha384 field if non-nil, zero value otherwise.

### GetEcdheRsaAes256GcmSha384Ok

`func (o *AcceptableCipherSuites) GetEcdheRsaAes256GcmSha384Ok() (*bool, bool)`

GetEcdheRsaAes256GcmSha384Ok returns a tuple with the EcdheRsaAes256GcmSha384 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcdheRsaAes256GcmSha384

`func (o *AcceptableCipherSuites) SetEcdheRsaAes256GcmSha384(v bool)`

SetEcdheRsaAes256GcmSha384 sets EcdheRsaAes256GcmSha384 field to given value.

### HasEcdheRsaAes256GcmSha384

`func (o *AcceptableCipherSuites) HasEcdheRsaAes256GcmSha384() bool`

HasEcdheRsaAes256GcmSha384 returns a boolean if a field has been set.

### GetEcdheRsaAes256Sha1

`func (o *AcceptableCipherSuites) GetEcdheRsaAes256Sha1() bool`

GetEcdheRsaAes256Sha1 returns the EcdheRsaAes256Sha1 field if non-nil, zero value otherwise.

### GetEcdheRsaAes256Sha1Ok

`func (o *AcceptableCipherSuites) GetEcdheRsaAes256Sha1Ok() (*bool, bool)`

GetEcdheRsaAes256Sha1Ok returns a tuple with the EcdheRsaAes256Sha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcdheRsaAes256Sha1

`func (o *AcceptableCipherSuites) SetEcdheRsaAes256Sha1(v bool)`

SetEcdheRsaAes256Sha1 sets EcdheRsaAes256Sha1 field to given value.

### HasEcdheRsaAes256Sha1

`func (o *AcceptableCipherSuites) HasEcdheRsaAes256Sha1() bool`

HasEcdheRsaAes256Sha1 returns a boolean if a field has been set.

### GetEcdheEcdsaChacha20Poly1305

`func (o *AcceptableCipherSuites) GetEcdheEcdsaChacha20Poly1305() bool`

GetEcdheEcdsaChacha20Poly1305 returns the EcdheEcdsaChacha20Poly1305 field if non-nil, zero value otherwise.

### GetEcdheEcdsaChacha20Poly1305Ok

`func (o *AcceptableCipherSuites) GetEcdheEcdsaChacha20Poly1305Ok() (*bool, bool)`

GetEcdheEcdsaChacha20Poly1305Ok returns a tuple with the EcdheEcdsaChacha20Poly1305 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcdheEcdsaChacha20Poly1305

`func (o *AcceptableCipherSuites) SetEcdheEcdsaChacha20Poly1305(v bool)`

SetEcdheEcdsaChacha20Poly1305 sets EcdheEcdsaChacha20Poly1305 field to given value.

### HasEcdheEcdsaChacha20Poly1305

`func (o *AcceptableCipherSuites) HasEcdheEcdsaChacha20Poly1305() bool`

HasEcdheEcdsaChacha20Poly1305 returns a boolean if a field has been set.

### GetEcdheRsaChacha20Poly1305

`func (o *AcceptableCipherSuites) GetEcdheRsaChacha20Poly1305() bool`

GetEcdheRsaChacha20Poly1305 returns the EcdheRsaChacha20Poly1305 field if non-nil, zero value otherwise.

### GetEcdheRsaChacha20Poly1305Ok

`func (o *AcceptableCipherSuites) GetEcdheRsaChacha20Poly1305Ok() (*bool, bool)`

GetEcdheRsaChacha20Poly1305Ok returns a tuple with the EcdheRsaChacha20Poly1305 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcdheRsaChacha20Poly1305

`func (o *AcceptableCipherSuites) SetEcdheRsaChacha20Poly1305(v bool)`

SetEcdheRsaChacha20Poly1305 sets EcdheRsaChacha20Poly1305 field to given value.

### HasEcdheRsaChacha20Poly1305

`func (o *AcceptableCipherSuites) HasEcdheRsaChacha20Poly1305() bool`

HasEcdheRsaChacha20Poly1305 returns a boolean if a field has been set.

### GetDheRsaChacha20Poly1305

`func (o *AcceptableCipherSuites) GetDheRsaChacha20Poly1305() bool`

GetDheRsaChacha20Poly1305 returns the DheRsaChacha20Poly1305 field if non-nil, zero value otherwise.

### GetDheRsaChacha20Poly1305Ok

`func (o *AcceptableCipherSuites) GetDheRsaChacha20Poly1305Ok() (*bool, bool)`

GetDheRsaChacha20Poly1305Ok returns a tuple with the DheRsaChacha20Poly1305 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDheRsaChacha20Poly1305

`func (o *AcceptableCipherSuites) SetDheRsaChacha20Poly1305(v bool)`

SetDheRsaChacha20Poly1305 sets DheRsaChacha20Poly1305 field to given value.

### HasDheRsaChacha20Poly1305

`func (o *AcceptableCipherSuites) HasDheRsaChacha20Poly1305() bool`

HasDheRsaChacha20Poly1305 returns a boolean if a field has been set.

### GetTlsAes256GcmSha384

`func (o *AcceptableCipherSuites) GetTlsAes256GcmSha384() bool`

GetTlsAes256GcmSha384 returns the TlsAes256GcmSha384 field if non-nil, zero value otherwise.

### GetTlsAes256GcmSha384Ok

`func (o *AcceptableCipherSuites) GetTlsAes256GcmSha384Ok() (*bool, bool)`

GetTlsAes256GcmSha384Ok returns a tuple with the TlsAes256GcmSha384 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsAes256GcmSha384

`func (o *AcceptableCipherSuites) SetTlsAes256GcmSha384(v bool)`

SetTlsAes256GcmSha384 sets TlsAes256GcmSha384 field to given value.

### HasTlsAes256GcmSha384

`func (o *AcceptableCipherSuites) HasTlsAes256GcmSha384() bool`

HasTlsAes256GcmSha384 returns a boolean if a field has been set.

### GetTlsChacha20Poly1305Sha256

`func (o *AcceptableCipherSuites) GetTlsChacha20Poly1305Sha256() bool`

GetTlsChacha20Poly1305Sha256 returns the TlsChacha20Poly1305Sha256 field if non-nil, zero value otherwise.

### GetTlsChacha20Poly1305Sha256Ok

`func (o *AcceptableCipherSuites) GetTlsChacha20Poly1305Sha256Ok() (*bool, bool)`

GetTlsChacha20Poly1305Sha256Ok returns a tuple with the TlsChacha20Poly1305Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsChacha20Poly1305Sha256

`func (o *AcceptableCipherSuites) SetTlsChacha20Poly1305Sha256(v bool)`

SetTlsChacha20Poly1305Sha256 sets TlsChacha20Poly1305Sha256 field to given value.

### HasTlsChacha20Poly1305Sha256

`func (o *AcceptableCipherSuites) HasTlsChacha20Poly1305Sha256() bool`

HasTlsChacha20Poly1305Sha256 returns a boolean if a field has been set.

### GetTlsAes128GcmSha256

`func (o *AcceptableCipherSuites) GetTlsAes128GcmSha256() bool`

GetTlsAes128GcmSha256 returns the TlsAes128GcmSha256 field if non-nil, zero value otherwise.

### GetTlsAes128GcmSha256Ok

`func (o *AcceptableCipherSuites) GetTlsAes128GcmSha256Ok() (*bool, bool)`

GetTlsAes128GcmSha256Ok returns a tuple with the TlsAes128GcmSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsAes128GcmSha256

`func (o *AcceptableCipherSuites) SetTlsAes128GcmSha256(v bool)`

SetTlsAes128GcmSha256 sets TlsAes128GcmSha256 field to given value.

### HasTlsAes128GcmSha256

`func (o *AcceptableCipherSuites) HasTlsAes128GcmSha256() bool`

HasTlsAes128GcmSha256 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


