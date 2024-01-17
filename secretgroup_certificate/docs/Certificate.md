# Certificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**ExpirationDate** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Details** | Pointer to [**CertificateDetails**](CertificateDetails.md) |  | [optional] 
**CertificateFileName** | Pointer to **string** |  | [optional] 

## Methods

### NewCertificate

`func NewCertificate() *Certificate`

NewCertificate instantiates a new Certificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateWithDefaults

`func NewCertificateWithDefaults() *Certificate`

NewCertificateWithDefaults instantiates a new Certificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Certificate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Certificate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Certificate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Certificate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExpirationDate

`func (o *Certificate) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *Certificate) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *Certificate) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *Certificate) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetMeta

`func (o *Certificate) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Certificate) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Certificate) SetMeta(v Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Certificate) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetType

`func (o *Certificate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Certificate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Certificate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Certificate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDetails

`func (o *Certificate) GetDetails() CertificateDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Certificate) GetDetailsOk() (*CertificateDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Certificate) SetDetails(v CertificateDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *Certificate) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetCertificateFileName

`func (o *Certificate) GetCertificateFileName() string`

GetCertificateFileName returns the CertificateFileName field if non-nil, zero value otherwise.

### GetCertificateFileNameOk

`func (o *Certificate) GetCertificateFileNameOk() (*string, bool)`

GetCertificateFileNameOk returns a tuple with the CertificateFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateFileName

`func (o *Certificate) SetCertificateFileName(v string)`

SetCertificateFileName sets CertificateFileName field to given value.

### HasCertificateFileName

`func (o *Certificate) HasCertificateFileName() bool`

HasCertificateFileName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


