# DlbCoreSslEndpointsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateKeyLabel** | Pointer to **string** |  | [optional] 
**PrivateKeyDigest** | Pointer to **string** |  | [optional] 
**PublicKeyLabel** | Pointer to **string** |  | [optional] 
**PublicKeyDigest** | Pointer to **string** |  | [optional] 
**PublicKeyCN** | Pointer to **string** |  | [optional] 
**ClientCertLabel** | Pointer to **string** |  | [optional] 
**ClientCertDigest** | Pointer to **string** |  | [optional] 
**ClientCertCN** | Pointer to **string** |  | [optional] 
**RevocationListLabel** | Pointer to **string** |  | [optional] 
**RevocationListDigest** | Pointer to **string** |  | [optional] 
**VerifyClientMode** | Pointer to **string** |  | [optional] [default to "off"]
**Mappings** | Pointer to [**[]DlbCoreSslEndpointsInnerMappingsInner**](DlbCoreSslEndpointsInnerMappingsInner.md) |  | [optional] 

## Methods

### NewDlbCoreSslEndpointsInner

`func NewDlbCoreSslEndpointsInner() *DlbCoreSslEndpointsInner`

NewDlbCoreSslEndpointsInner instantiates a new DlbCoreSslEndpointsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDlbCoreSslEndpointsInnerWithDefaults

`func NewDlbCoreSslEndpointsInnerWithDefaults() *DlbCoreSslEndpointsInner`

NewDlbCoreSslEndpointsInnerWithDefaults instantiates a new DlbCoreSslEndpointsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateKeyLabel

`func (o *DlbCoreSslEndpointsInner) GetPrivateKeyLabel() string`

GetPrivateKeyLabel returns the PrivateKeyLabel field if non-nil, zero value otherwise.

### GetPrivateKeyLabelOk

`func (o *DlbCoreSslEndpointsInner) GetPrivateKeyLabelOk() (*string, bool)`

GetPrivateKeyLabelOk returns a tuple with the PrivateKeyLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyLabel

`func (o *DlbCoreSslEndpointsInner) SetPrivateKeyLabel(v string)`

SetPrivateKeyLabel sets PrivateKeyLabel field to given value.

### HasPrivateKeyLabel

`func (o *DlbCoreSslEndpointsInner) HasPrivateKeyLabel() bool`

HasPrivateKeyLabel returns a boolean if a field has been set.

### GetPrivateKeyDigest

`func (o *DlbCoreSslEndpointsInner) GetPrivateKeyDigest() string`

GetPrivateKeyDigest returns the PrivateKeyDigest field if non-nil, zero value otherwise.

### GetPrivateKeyDigestOk

`func (o *DlbCoreSslEndpointsInner) GetPrivateKeyDigestOk() (*string, bool)`

GetPrivateKeyDigestOk returns a tuple with the PrivateKeyDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyDigest

`func (o *DlbCoreSslEndpointsInner) SetPrivateKeyDigest(v string)`

SetPrivateKeyDigest sets PrivateKeyDigest field to given value.

### HasPrivateKeyDigest

`func (o *DlbCoreSslEndpointsInner) HasPrivateKeyDigest() bool`

HasPrivateKeyDigest returns a boolean if a field has been set.

### GetPublicKeyLabel

`func (o *DlbCoreSslEndpointsInner) GetPublicKeyLabel() string`

GetPublicKeyLabel returns the PublicKeyLabel field if non-nil, zero value otherwise.

### GetPublicKeyLabelOk

`func (o *DlbCoreSslEndpointsInner) GetPublicKeyLabelOk() (*string, bool)`

GetPublicKeyLabelOk returns a tuple with the PublicKeyLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyLabel

`func (o *DlbCoreSslEndpointsInner) SetPublicKeyLabel(v string)`

SetPublicKeyLabel sets PublicKeyLabel field to given value.

### HasPublicKeyLabel

`func (o *DlbCoreSslEndpointsInner) HasPublicKeyLabel() bool`

HasPublicKeyLabel returns a boolean if a field has been set.

### GetPublicKeyDigest

`func (o *DlbCoreSslEndpointsInner) GetPublicKeyDigest() string`

GetPublicKeyDigest returns the PublicKeyDigest field if non-nil, zero value otherwise.

### GetPublicKeyDigestOk

`func (o *DlbCoreSslEndpointsInner) GetPublicKeyDigestOk() (*string, bool)`

GetPublicKeyDigestOk returns a tuple with the PublicKeyDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyDigest

`func (o *DlbCoreSslEndpointsInner) SetPublicKeyDigest(v string)`

SetPublicKeyDigest sets PublicKeyDigest field to given value.

### HasPublicKeyDigest

`func (o *DlbCoreSslEndpointsInner) HasPublicKeyDigest() bool`

HasPublicKeyDigest returns a boolean if a field has been set.

### GetPublicKeyCN

`func (o *DlbCoreSslEndpointsInner) GetPublicKeyCN() string`

GetPublicKeyCN returns the PublicKeyCN field if non-nil, zero value otherwise.

### GetPublicKeyCNOk

`func (o *DlbCoreSslEndpointsInner) GetPublicKeyCNOk() (*string, bool)`

GetPublicKeyCNOk returns a tuple with the PublicKeyCN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyCN

`func (o *DlbCoreSslEndpointsInner) SetPublicKeyCN(v string)`

SetPublicKeyCN sets PublicKeyCN field to given value.

### HasPublicKeyCN

`func (o *DlbCoreSslEndpointsInner) HasPublicKeyCN() bool`

HasPublicKeyCN returns a boolean if a field has been set.

### GetClientCertLabel

`func (o *DlbCoreSslEndpointsInner) GetClientCertLabel() string`

GetClientCertLabel returns the ClientCertLabel field if non-nil, zero value otherwise.

### GetClientCertLabelOk

`func (o *DlbCoreSslEndpointsInner) GetClientCertLabelOk() (*string, bool)`

GetClientCertLabelOk returns a tuple with the ClientCertLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertLabel

`func (o *DlbCoreSslEndpointsInner) SetClientCertLabel(v string)`

SetClientCertLabel sets ClientCertLabel field to given value.

### HasClientCertLabel

`func (o *DlbCoreSslEndpointsInner) HasClientCertLabel() bool`

HasClientCertLabel returns a boolean if a field has been set.

### GetClientCertDigest

`func (o *DlbCoreSslEndpointsInner) GetClientCertDigest() string`

GetClientCertDigest returns the ClientCertDigest field if non-nil, zero value otherwise.

### GetClientCertDigestOk

`func (o *DlbCoreSslEndpointsInner) GetClientCertDigestOk() (*string, bool)`

GetClientCertDigestOk returns a tuple with the ClientCertDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertDigest

`func (o *DlbCoreSslEndpointsInner) SetClientCertDigest(v string)`

SetClientCertDigest sets ClientCertDigest field to given value.

### HasClientCertDigest

`func (o *DlbCoreSslEndpointsInner) HasClientCertDigest() bool`

HasClientCertDigest returns a boolean if a field has been set.

### GetClientCertCN

`func (o *DlbCoreSslEndpointsInner) GetClientCertCN() string`

GetClientCertCN returns the ClientCertCN field if non-nil, zero value otherwise.

### GetClientCertCNOk

`func (o *DlbCoreSslEndpointsInner) GetClientCertCNOk() (*string, bool)`

GetClientCertCNOk returns a tuple with the ClientCertCN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertCN

`func (o *DlbCoreSslEndpointsInner) SetClientCertCN(v string)`

SetClientCertCN sets ClientCertCN field to given value.

### HasClientCertCN

`func (o *DlbCoreSslEndpointsInner) HasClientCertCN() bool`

HasClientCertCN returns a boolean if a field has been set.

### GetRevocationListLabel

`func (o *DlbCoreSslEndpointsInner) GetRevocationListLabel() string`

GetRevocationListLabel returns the RevocationListLabel field if non-nil, zero value otherwise.

### GetRevocationListLabelOk

`func (o *DlbCoreSslEndpointsInner) GetRevocationListLabelOk() (*string, bool)`

GetRevocationListLabelOk returns a tuple with the RevocationListLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationListLabel

`func (o *DlbCoreSslEndpointsInner) SetRevocationListLabel(v string)`

SetRevocationListLabel sets RevocationListLabel field to given value.

### HasRevocationListLabel

`func (o *DlbCoreSslEndpointsInner) HasRevocationListLabel() bool`

HasRevocationListLabel returns a boolean if a field has been set.

### GetRevocationListDigest

`func (o *DlbCoreSslEndpointsInner) GetRevocationListDigest() string`

GetRevocationListDigest returns the RevocationListDigest field if non-nil, zero value otherwise.

### GetRevocationListDigestOk

`func (o *DlbCoreSslEndpointsInner) GetRevocationListDigestOk() (*string, bool)`

GetRevocationListDigestOk returns a tuple with the RevocationListDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationListDigest

`func (o *DlbCoreSslEndpointsInner) SetRevocationListDigest(v string)`

SetRevocationListDigest sets RevocationListDigest field to given value.

### HasRevocationListDigest

`func (o *DlbCoreSslEndpointsInner) HasRevocationListDigest() bool`

HasRevocationListDigest returns a boolean if a field has been set.

### GetVerifyClientMode

`func (o *DlbCoreSslEndpointsInner) GetVerifyClientMode() string`

GetVerifyClientMode returns the VerifyClientMode field if non-nil, zero value otherwise.

### GetVerifyClientModeOk

`func (o *DlbCoreSslEndpointsInner) GetVerifyClientModeOk() (*string, bool)`

GetVerifyClientModeOk returns a tuple with the VerifyClientMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyClientMode

`func (o *DlbCoreSslEndpointsInner) SetVerifyClientMode(v string)`

SetVerifyClientMode sets VerifyClientMode field to given value.

### HasVerifyClientMode

`func (o *DlbCoreSslEndpointsInner) HasVerifyClientMode() bool`

HasVerifyClientMode returns a boolean if a field has been set.

### GetMappings

`func (o *DlbCoreSslEndpointsInner) GetMappings() []DlbCoreSslEndpointsInnerMappingsInner`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *DlbCoreSslEndpointsInner) GetMappingsOk() (*[]DlbCoreSslEndpointsInnerMappingsInner, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *DlbCoreSslEndpointsInner) SetMappings(v []DlbCoreSslEndpointsInnerMappingsInner)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *DlbCoreSslEndpointsInner) HasMappings() bool`

HasMappings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


