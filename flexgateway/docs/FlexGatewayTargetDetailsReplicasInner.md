# FlexGatewayTargetDetailsReplicasInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**CertificateExpirationDates** | Pointer to [**[]time.Time**](time.Time.md) |  | [optional] 

## Methods

### NewFlexGatewayTargetDetailsReplicasInner

`func NewFlexGatewayTargetDetailsReplicasInner() *FlexGatewayTargetDetailsReplicasInner`

NewFlexGatewayTargetDetailsReplicasInner instantiates a new FlexGatewayTargetDetailsReplicasInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlexGatewayTargetDetailsReplicasInnerWithDefaults

`func NewFlexGatewayTargetDetailsReplicasInnerWithDefaults() *FlexGatewayTargetDetailsReplicasInner`

NewFlexGatewayTargetDetailsReplicasInnerWithDefaults instantiates a new FlexGatewayTargetDetailsReplicasInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *FlexGatewayTargetDetailsReplicasInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FlexGatewayTargetDetailsReplicasInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FlexGatewayTargetDetailsReplicasInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FlexGatewayTargetDetailsReplicasInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCount

`func (o *FlexGatewayTargetDetailsReplicasInner) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *FlexGatewayTargetDetailsReplicasInner) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *FlexGatewayTargetDetailsReplicasInner) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *FlexGatewayTargetDetailsReplicasInner) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCertificateExpirationDates

`func (o *FlexGatewayTargetDetailsReplicasInner) GetCertificateExpirationDates() []time.Time`

GetCertificateExpirationDates returns the CertificateExpirationDates field if non-nil, zero value otherwise.

### GetCertificateExpirationDatesOk

`func (o *FlexGatewayTargetDetailsReplicasInner) GetCertificateExpirationDatesOk() (*[]time.Time, bool)`

GetCertificateExpirationDatesOk returns a tuple with the CertificateExpirationDates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateExpirationDates

`func (o *FlexGatewayTargetDetailsReplicasInner) SetCertificateExpirationDates(v []time.Time)`

SetCertificateExpirationDates sets CertificateExpirationDates field to given value.

### HasCertificateExpirationDates

`func (o *FlexGatewayTargetDetailsReplicasInner) HasCertificateExpirationDates() bool`

HasCertificateExpirationDates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


