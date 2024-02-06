# SecretGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Downloadable** | Pointer to **bool** |  | [optional] 
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 

## Methods

### NewSecretGroup

`func NewSecretGroup() *SecretGroup`

NewSecretGroup instantiates a new SecretGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretGroupWithDefaults

`func NewSecretGroupWithDefaults() *SecretGroup`

NewSecretGroupWithDefaults instantiates a new SecretGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SecretGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecretGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecretGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecretGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDownloadable

`func (o *SecretGroup) GetDownloadable() bool`

GetDownloadable returns the Downloadable field if non-nil, zero value otherwise.

### GetDownloadableOk

`func (o *SecretGroup) GetDownloadableOk() (*bool, bool)`

GetDownloadableOk returns a tuple with the Downloadable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadable

`func (o *SecretGroup) SetDownloadable(v bool)`

SetDownloadable sets Downloadable field to given value.

### HasDownloadable

`func (o *SecretGroup) HasDownloadable() bool`

HasDownloadable returns a boolean if a field has been set.

### GetMeta

`func (o *SecretGroup) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SecretGroup) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SecretGroup) SetMeta(v Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SecretGroup) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


