# DlbCoreSslEndpoints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateKeyLabel** | Pointer to **string** |  | [optional] 
**PrivateKeyDigest** | Pointer to **string** |  | [optional] 
**PrivateKey** | Pointer to **string** |  | [optional] 
**PublicKeyLabel** | Pointer to **string** |  | [optional] 
**PublicKeyDigest** | Pointer to **string** |  | [optional] 
**PublicKeyCN** | Pointer to **string** |  | [optional] 
**PublicKey** | Pointer to **string** |  | [optional] 
**VerifyClientMode** | Pointer to **string** |  | [optional] [default to "off"]
**Mappings** | Pointer to [**[]DlbCoreMappings**](DlbCoreMappings.md) |  | [optional] 

## Methods

### NewDlbCoreSslEndpoints

`func NewDlbCoreSslEndpoints() *DlbCoreSslEndpoints`

NewDlbCoreSslEndpoints instantiates a new DlbCoreSslEndpoints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDlbCoreSslEndpointsWithDefaults

`func NewDlbCoreSslEndpointsWithDefaults() *DlbCoreSslEndpoints`

NewDlbCoreSslEndpointsWithDefaults instantiates a new DlbCoreSslEndpoints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateKeyLabel

`func (o *DlbCoreSslEndpoints) GetPrivateKeyLabel() string`

GetPrivateKeyLabel returns the PrivateKeyLabel field if non-nil, zero value otherwise.

### GetPrivateKeyLabelOk

`func (o *DlbCoreSslEndpoints) GetPrivateKeyLabelOk() (*string, bool)`

GetPrivateKeyLabelOk returns a tuple with the PrivateKeyLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyLabel

`func (o *DlbCoreSslEndpoints) SetPrivateKeyLabel(v string)`

SetPrivateKeyLabel sets PrivateKeyLabel field to given value.

### HasPrivateKeyLabel

`func (o *DlbCoreSslEndpoints) HasPrivateKeyLabel() bool`

HasPrivateKeyLabel returns a boolean if a field has been set.

### GetPrivateKeyDigest

`func (o *DlbCoreSslEndpoints) GetPrivateKeyDigest() string`

GetPrivateKeyDigest returns the PrivateKeyDigest field if non-nil, zero value otherwise.

### GetPrivateKeyDigestOk

`func (o *DlbCoreSslEndpoints) GetPrivateKeyDigestOk() (*string, bool)`

GetPrivateKeyDigestOk returns a tuple with the PrivateKeyDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyDigest

`func (o *DlbCoreSslEndpoints) SetPrivateKeyDigest(v string)`

SetPrivateKeyDigest sets PrivateKeyDigest field to given value.

### HasPrivateKeyDigest

`func (o *DlbCoreSslEndpoints) HasPrivateKeyDigest() bool`

HasPrivateKeyDigest returns a boolean if a field has been set.

### GetPrivateKey

`func (o *DlbCoreSslEndpoints) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *DlbCoreSslEndpoints) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *DlbCoreSslEndpoints) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *DlbCoreSslEndpoints) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetPublicKeyLabel

`func (o *DlbCoreSslEndpoints) GetPublicKeyLabel() string`

GetPublicKeyLabel returns the PublicKeyLabel field if non-nil, zero value otherwise.

### GetPublicKeyLabelOk

`func (o *DlbCoreSslEndpoints) GetPublicKeyLabelOk() (*string, bool)`

GetPublicKeyLabelOk returns a tuple with the PublicKeyLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyLabel

`func (o *DlbCoreSslEndpoints) SetPublicKeyLabel(v string)`

SetPublicKeyLabel sets PublicKeyLabel field to given value.

### HasPublicKeyLabel

`func (o *DlbCoreSslEndpoints) HasPublicKeyLabel() bool`

HasPublicKeyLabel returns a boolean if a field has been set.

### GetPublicKeyDigest

`func (o *DlbCoreSslEndpoints) GetPublicKeyDigest() string`

GetPublicKeyDigest returns the PublicKeyDigest field if non-nil, zero value otherwise.

### GetPublicKeyDigestOk

`func (o *DlbCoreSslEndpoints) GetPublicKeyDigestOk() (*string, bool)`

GetPublicKeyDigestOk returns a tuple with the PublicKeyDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyDigest

`func (o *DlbCoreSslEndpoints) SetPublicKeyDigest(v string)`

SetPublicKeyDigest sets PublicKeyDigest field to given value.

### HasPublicKeyDigest

`func (o *DlbCoreSslEndpoints) HasPublicKeyDigest() bool`

HasPublicKeyDigest returns a boolean if a field has been set.

### GetPublicKeyCN

`func (o *DlbCoreSslEndpoints) GetPublicKeyCN() string`

GetPublicKeyCN returns the PublicKeyCN field if non-nil, zero value otherwise.

### GetPublicKeyCNOk

`func (o *DlbCoreSslEndpoints) GetPublicKeyCNOk() (*string, bool)`

GetPublicKeyCNOk returns a tuple with the PublicKeyCN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyCN

`func (o *DlbCoreSslEndpoints) SetPublicKeyCN(v string)`

SetPublicKeyCN sets PublicKeyCN field to given value.

### HasPublicKeyCN

`func (o *DlbCoreSslEndpoints) HasPublicKeyCN() bool`

HasPublicKeyCN returns a boolean if a field has been set.

### GetPublicKey

`func (o *DlbCoreSslEndpoints) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *DlbCoreSslEndpoints) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *DlbCoreSslEndpoints) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *DlbCoreSslEndpoints) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetVerifyClientMode

`func (o *DlbCoreSslEndpoints) GetVerifyClientMode() string`

GetVerifyClientMode returns the VerifyClientMode field if non-nil, zero value otherwise.

### GetVerifyClientModeOk

`func (o *DlbCoreSslEndpoints) GetVerifyClientModeOk() (*string, bool)`

GetVerifyClientModeOk returns a tuple with the VerifyClientMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyClientMode

`func (o *DlbCoreSslEndpoints) SetVerifyClientMode(v string)`

SetVerifyClientMode sets VerifyClientMode field to given value.

### HasVerifyClientMode

`func (o *DlbCoreSslEndpoints) HasVerifyClientMode() bool`

HasVerifyClientMode returns a boolean if a field has been set.

### GetMappings

`func (o *DlbCoreSslEndpoints) GetMappings() []DlbCoreMappings`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *DlbCoreSslEndpoints) GetMappingsOk() (*[]DlbCoreMappings, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *DlbCoreSslEndpoints) SetMappings(v []DlbCoreMappings)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *DlbCoreSslEndpoints) HasMappings() bool`

HasMappings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


