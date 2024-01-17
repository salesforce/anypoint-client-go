# TlsContextDetailsOutboundSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkipServerCertValidation** | Pointer to **bool** | flag that indicates whether the server certificate validation must be skipped. | [optional] 

## Methods

### NewTlsContextDetailsOutboundSettings

`func NewTlsContextDetailsOutboundSettings() *TlsContextDetailsOutboundSettings`

NewTlsContextDetailsOutboundSettings instantiates a new TlsContextDetailsOutboundSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsContextDetailsOutboundSettingsWithDefaults

`func NewTlsContextDetailsOutboundSettingsWithDefaults() *TlsContextDetailsOutboundSettings`

NewTlsContextDetailsOutboundSettingsWithDefaults instantiates a new TlsContextDetailsOutboundSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkipServerCertValidation

`func (o *TlsContextDetailsOutboundSettings) GetSkipServerCertValidation() bool`

GetSkipServerCertValidation returns the SkipServerCertValidation field if non-nil, zero value otherwise.

### GetSkipServerCertValidationOk

`func (o *TlsContextDetailsOutboundSettings) GetSkipServerCertValidationOk() (*bool, bool)`

GetSkipServerCertValidationOk returns a tuple with the SkipServerCertValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipServerCertValidation

`func (o *TlsContextDetailsOutboundSettings) SetSkipServerCertValidation(v bool)`

SetSkipServerCertValidation sets SkipServerCertValidation field to given value.

### HasSkipServerCertValidation

`func (o *TlsContextDetailsOutboundSettings) HasSkipServerCertValidation() bool`

HasSkipServerCertValidation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


