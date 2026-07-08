# TlsContextPostBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**TlsConfig** | Pointer to [**TlsContextPostBodyTlsConfig**](TlsContextPostBodyTlsConfig.md) |  | [optional] 
**Ciphers** | Pointer to [**Ciphers**](Ciphers.md) |  | [optional] 

## Methods

### NewTlsContextPostBody

`func NewTlsContextPostBody() *TlsContextPostBody`

NewTlsContextPostBody instantiates a new TlsContextPostBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsContextPostBodyWithDefaults

`func NewTlsContextPostBodyWithDefaults() *TlsContextPostBody`

NewTlsContextPostBodyWithDefaults instantiates a new TlsContextPostBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TlsContextPostBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TlsContextPostBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TlsContextPostBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TlsContextPostBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTlsConfig

`func (o *TlsContextPostBody) GetTlsConfig() TlsContextPostBodyTlsConfig`

GetTlsConfig returns the TlsConfig field if non-nil, zero value otherwise.

### GetTlsConfigOk

`func (o *TlsContextPostBody) GetTlsConfigOk() (*TlsContextPostBodyTlsConfig, bool)`

GetTlsConfigOk returns a tuple with the TlsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsConfig

`func (o *TlsContextPostBody) SetTlsConfig(v TlsContextPostBodyTlsConfig)`

SetTlsConfig sets TlsConfig field to given value.

### HasTlsConfig

`func (o *TlsContextPostBody) HasTlsConfig() bool`

HasTlsConfig returns a boolean if a field has been set.

### GetCiphers

`func (o *TlsContextPostBody) GetCiphers() Ciphers`

GetCiphers returns the Ciphers field if non-nil, zero value otherwise.

### GetCiphersOk

`func (o *TlsContextPostBody) GetCiphersOk() (*Ciphers, bool)`

GetCiphersOk returns a tuple with the Ciphers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiphers

`func (o *TlsContextPostBody) SetCiphers(v Ciphers)`

SetCiphers sets Ciphers field to given value.

### HasCiphers

`func (o *TlsContextPostBody) HasCiphers() bool`

HasCiphers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


