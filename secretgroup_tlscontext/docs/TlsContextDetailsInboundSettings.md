# TlsContextDetailsInboundSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableClientCertValidation** | Pointer to **bool** | flag that indicates whether the client certificate validation must be enforced. | [optional] 

## Methods

### NewTlsContextDetailsInboundSettings

`func NewTlsContextDetailsInboundSettings() *TlsContextDetailsInboundSettings`

NewTlsContextDetailsInboundSettings instantiates a new TlsContextDetailsInboundSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsContextDetailsInboundSettingsWithDefaults

`func NewTlsContextDetailsInboundSettingsWithDefaults() *TlsContextDetailsInboundSettings`

NewTlsContextDetailsInboundSettingsWithDefaults instantiates a new TlsContextDetailsInboundSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableClientCertValidation

`func (o *TlsContextDetailsInboundSettings) GetEnableClientCertValidation() bool`

GetEnableClientCertValidation returns the EnableClientCertValidation field if non-nil, zero value otherwise.

### GetEnableClientCertValidationOk

`func (o *TlsContextDetailsInboundSettings) GetEnableClientCertValidationOk() (*bool, bool)`

GetEnableClientCertValidationOk returns a tuple with the EnableClientCertValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableClientCertValidation

`func (o *TlsContextDetailsInboundSettings) SetEnableClientCertValidation(v bool)`

SetEnableClientCertValidation sets EnableClientCertValidation field to given value.

### HasEnableClientCertValidation

`func (o *TlsContextDetailsInboundSettings) HasEnableClientCertValidation() bool`

HasEnableClientCertValidation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


