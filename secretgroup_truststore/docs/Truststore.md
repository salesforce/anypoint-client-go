# Truststore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**ExpirationDate** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 
**Type** | Pointer to **string** | Type of truststore supported | [optional] 
**Details** | Pointer to [**TruststoreDetails**](TruststoreDetails.md) |  | [optional] 
**TruststoreFileName** | Pointer to **string** | File name of the truststore that is stored in this secret | [optional] 
**Algorithm** | Pointer to **string** | Algorithm used to create the truststore manager factory which will make use of this truststore | [optional] 

## Methods

### NewTruststore

`func NewTruststore() *Truststore`

NewTruststore instantiates a new Truststore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTruststoreWithDefaults

`func NewTruststoreWithDefaults() *Truststore`

NewTruststoreWithDefaults instantiates a new Truststore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Truststore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Truststore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Truststore) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Truststore) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExpirationDate

`func (o *Truststore) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *Truststore) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *Truststore) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *Truststore) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetMeta

`func (o *Truststore) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Truststore) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Truststore) SetMeta(v Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Truststore) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetType

`func (o *Truststore) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Truststore) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Truststore) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Truststore) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDetails

`func (o *Truststore) GetDetails() TruststoreDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Truststore) GetDetailsOk() (*TruststoreDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Truststore) SetDetails(v TruststoreDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *Truststore) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetTruststoreFileName

`func (o *Truststore) GetTruststoreFileName() string`

GetTruststoreFileName returns the TruststoreFileName field if non-nil, zero value otherwise.

### GetTruststoreFileNameOk

`func (o *Truststore) GetTruststoreFileNameOk() (*string, bool)`

GetTruststoreFileNameOk returns a tuple with the TruststoreFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruststoreFileName

`func (o *Truststore) SetTruststoreFileName(v string)`

SetTruststoreFileName sets TruststoreFileName field to given value.

### HasTruststoreFileName

`func (o *Truststore) HasTruststoreFileName() bool`

HasTruststoreFileName returns a boolean if a field has been set.

### GetAlgorithm

`func (o *Truststore) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *Truststore) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *Truststore) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *Truststore) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


