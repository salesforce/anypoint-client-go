# PrivateSpaceRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | Pointer to **string** | Destination CIDR. | [optional] 
**Status** | Pointer to **string** | Route status. | [optional] 
**Target** | Pointer to **string** | Route target (transit gateway id, VPN id, etc.). | [optional] 

## Methods

### NewPrivateSpaceRoute

`func NewPrivateSpaceRoute() *PrivateSpaceRoute`

NewPrivateSpaceRoute instantiates a new PrivateSpaceRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateSpaceRouteWithDefaults

`func NewPrivateSpaceRouteWithDefaults() *PrivateSpaceRoute`

NewPrivateSpaceRouteWithDefaults instantiates a new PrivateSpaceRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *PrivateSpaceRoute) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *PrivateSpaceRoute) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *PrivateSpaceRoute) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *PrivateSpaceRoute) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetStatus

`func (o *PrivateSpaceRoute) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PrivateSpaceRoute) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PrivateSpaceRoute) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PrivateSpaceRoute) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTarget

`func (o *PrivateSpaceRoute) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *PrivateSpaceRoute) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *PrivateSpaceRoute) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *PrivateSpaceRoute) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


