# DlbPostBodySslEndpointsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateKeyLabel** | Pointer to **string** |  | [optional] 
**PrivateKey** | Pointer to **string** |  | [optional] 
**PublicKeyLabel** | Pointer to **string** |  | [optional] 
**PublicKey** | Pointer to **string** |  | [optional] 
**VerifyClientMode** | Pointer to **string** |  | [optional] [default to "off"]
**ClientCertLabel** | Pointer to **string** |  | [optional] 
**ClientCert** | Pointer to **string** |  | [optional] 
**RevocationListLabel** | Pointer to **string** |  | [optional] 
**RevocationList** | Pointer to **string** |  | [optional] 
**Mappings** | Pointer to [**[]DlbPostBodySslEndpointsInnerMappingsInner**](DlbPostBodySslEndpointsInnerMappingsInner.md) |  | [optional] 

## Methods

### NewDlbPostBodySslEndpointsInner

`func NewDlbPostBodySslEndpointsInner() *DlbPostBodySslEndpointsInner`

NewDlbPostBodySslEndpointsInner instantiates a new DlbPostBodySslEndpointsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDlbPostBodySslEndpointsInnerWithDefaults

`func NewDlbPostBodySslEndpointsInnerWithDefaults() *DlbPostBodySslEndpointsInner`

NewDlbPostBodySslEndpointsInnerWithDefaults instantiates a new DlbPostBodySslEndpointsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateKeyLabel

`func (o *DlbPostBodySslEndpointsInner) GetPrivateKeyLabel() string`

GetPrivateKeyLabel returns the PrivateKeyLabel field if non-nil, zero value otherwise.

### GetPrivateKeyLabelOk

`func (o *DlbPostBodySslEndpointsInner) GetPrivateKeyLabelOk() (*string, bool)`

GetPrivateKeyLabelOk returns a tuple with the PrivateKeyLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyLabel

`func (o *DlbPostBodySslEndpointsInner) SetPrivateKeyLabel(v string)`

SetPrivateKeyLabel sets PrivateKeyLabel field to given value.

### HasPrivateKeyLabel

`func (o *DlbPostBodySslEndpointsInner) HasPrivateKeyLabel() bool`

HasPrivateKeyLabel returns a boolean if a field has been set.

### GetPrivateKey

`func (o *DlbPostBodySslEndpointsInner) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *DlbPostBodySslEndpointsInner) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *DlbPostBodySslEndpointsInner) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *DlbPostBodySslEndpointsInner) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetPublicKeyLabel

`func (o *DlbPostBodySslEndpointsInner) GetPublicKeyLabel() string`

GetPublicKeyLabel returns the PublicKeyLabel field if non-nil, zero value otherwise.

### GetPublicKeyLabelOk

`func (o *DlbPostBodySslEndpointsInner) GetPublicKeyLabelOk() (*string, bool)`

GetPublicKeyLabelOk returns a tuple with the PublicKeyLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyLabel

`func (o *DlbPostBodySslEndpointsInner) SetPublicKeyLabel(v string)`

SetPublicKeyLabel sets PublicKeyLabel field to given value.

### HasPublicKeyLabel

`func (o *DlbPostBodySslEndpointsInner) HasPublicKeyLabel() bool`

HasPublicKeyLabel returns a boolean if a field has been set.

### GetPublicKey

`func (o *DlbPostBodySslEndpointsInner) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *DlbPostBodySslEndpointsInner) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *DlbPostBodySslEndpointsInner) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *DlbPostBodySslEndpointsInner) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetVerifyClientMode

`func (o *DlbPostBodySslEndpointsInner) GetVerifyClientMode() string`

GetVerifyClientMode returns the VerifyClientMode field if non-nil, zero value otherwise.

### GetVerifyClientModeOk

`func (o *DlbPostBodySslEndpointsInner) GetVerifyClientModeOk() (*string, bool)`

GetVerifyClientModeOk returns a tuple with the VerifyClientMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyClientMode

`func (o *DlbPostBodySslEndpointsInner) SetVerifyClientMode(v string)`

SetVerifyClientMode sets VerifyClientMode field to given value.

### HasVerifyClientMode

`func (o *DlbPostBodySslEndpointsInner) HasVerifyClientMode() bool`

HasVerifyClientMode returns a boolean if a field has been set.

### GetClientCertLabel

`func (o *DlbPostBodySslEndpointsInner) GetClientCertLabel() string`

GetClientCertLabel returns the ClientCertLabel field if non-nil, zero value otherwise.

### GetClientCertLabelOk

`func (o *DlbPostBodySslEndpointsInner) GetClientCertLabelOk() (*string, bool)`

GetClientCertLabelOk returns a tuple with the ClientCertLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertLabel

`func (o *DlbPostBodySslEndpointsInner) SetClientCertLabel(v string)`

SetClientCertLabel sets ClientCertLabel field to given value.

### HasClientCertLabel

`func (o *DlbPostBodySslEndpointsInner) HasClientCertLabel() bool`

HasClientCertLabel returns a boolean if a field has been set.

### GetClientCert

`func (o *DlbPostBodySslEndpointsInner) GetClientCert() string`

GetClientCert returns the ClientCert field if non-nil, zero value otherwise.

### GetClientCertOk

`func (o *DlbPostBodySslEndpointsInner) GetClientCertOk() (*string, bool)`

GetClientCertOk returns a tuple with the ClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCert

`func (o *DlbPostBodySslEndpointsInner) SetClientCert(v string)`

SetClientCert sets ClientCert field to given value.

### HasClientCert

`func (o *DlbPostBodySslEndpointsInner) HasClientCert() bool`

HasClientCert returns a boolean if a field has been set.

### GetRevocationListLabel

`func (o *DlbPostBodySslEndpointsInner) GetRevocationListLabel() string`

GetRevocationListLabel returns the RevocationListLabel field if non-nil, zero value otherwise.

### GetRevocationListLabelOk

`func (o *DlbPostBodySslEndpointsInner) GetRevocationListLabelOk() (*string, bool)`

GetRevocationListLabelOk returns a tuple with the RevocationListLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationListLabel

`func (o *DlbPostBodySslEndpointsInner) SetRevocationListLabel(v string)`

SetRevocationListLabel sets RevocationListLabel field to given value.

### HasRevocationListLabel

`func (o *DlbPostBodySslEndpointsInner) HasRevocationListLabel() bool`

HasRevocationListLabel returns a boolean if a field has been set.

### GetRevocationList

`func (o *DlbPostBodySslEndpointsInner) GetRevocationList() string`

GetRevocationList returns the RevocationList field if non-nil, zero value otherwise.

### GetRevocationListOk

`func (o *DlbPostBodySslEndpointsInner) GetRevocationListOk() (*string, bool)`

GetRevocationListOk returns a tuple with the RevocationList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationList

`func (o *DlbPostBodySslEndpointsInner) SetRevocationList(v string)`

SetRevocationList sets RevocationList field to given value.

### HasRevocationList

`func (o *DlbPostBodySslEndpointsInner) HasRevocationList() bool`

HasRevocationList returns a boolean if a field has been set.

### GetMappings

`func (o *DlbPostBodySslEndpointsInner) GetMappings() []DlbPostBodySslEndpointsInnerMappingsInner`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *DlbPostBodySslEndpointsInner) GetMappingsOk() (*[]DlbPostBodySslEndpointsInnerMappingsInner, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *DlbPostBodySslEndpointsInner) SetMappings(v []DlbPostBodySslEndpointsInnerMappingsInner)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *DlbPostBodySslEndpointsInner) HasMappings() bool`

HasMappings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


