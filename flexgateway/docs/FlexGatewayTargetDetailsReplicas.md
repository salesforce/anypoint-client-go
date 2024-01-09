# FlexGatewayTargetDetailsReplicas

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**CertificateExpirationDates** | Pointer to [**[]time.Time**](time.Time.md) |  | [optional] 

## Methods

### NewFlexGatewayTargetDetailsReplicas

`func NewFlexGatewayTargetDetailsReplicas() *FlexGatewayTargetDetailsReplicas`

NewFlexGatewayTargetDetailsReplicas instantiates a new FlexGatewayTargetDetailsReplicas object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlexGatewayTargetDetailsReplicasWithDefaults

`func NewFlexGatewayTargetDetailsReplicasWithDefaults() *FlexGatewayTargetDetailsReplicas`

NewFlexGatewayTargetDetailsReplicasWithDefaults instantiates a new FlexGatewayTargetDetailsReplicas object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *FlexGatewayTargetDetailsReplicas) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FlexGatewayTargetDetailsReplicas) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FlexGatewayTargetDetailsReplicas) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FlexGatewayTargetDetailsReplicas) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCount

`func (o *FlexGatewayTargetDetailsReplicas) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *FlexGatewayTargetDetailsReplicas) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *FlexGatewayTargetDetailsReplicas) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *FlexGatewayTargetDetailsReplicas) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCertificateExpirationDates

`func (o *FlexGatewayTargetDetailsReplicas) GetCertificateExpirationDates() []time.Time`

GetCertificateExpirationDates returns the CertificateExpirationDates field if non-nil, zero value otherwise.

### GetCertificateExpirationDatesOk

`func (o *FlexGatewayTargetDetailsReplicas) GetCertificateExpirationDatesOk() (*[]time.Time, bool)`

GetCertificateExpirationDatesOk returns a tuple with the CertificateExpirationDates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateExpirationDates

`func (o *FlexGatewayTargetDetailsReplicas) SetCertificateExpirationDates(v []time.Time)`

SetCertificateExpirationDates sets CertificateExpirationDates field to given value.

### HasCertificateExpirationDates

`func (o *FlexGatewayTargetDetailsReplicas) HasCertificateExpirationDates() bool`

HasCertificateExpirationDates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


