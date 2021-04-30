# Vpns

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assigned** | **int32** | An explanation about the purpose of this instance. | [default to 0]
**Reassigned** | Pointer to **int32** | An explanation about the purpose of this instance. | [optional] [default to 0]

## Methods

### NewVpns

`func NewVpns(assigned int32, ) *Vpns`

NewVpns instantiates a new Vpns object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnsWithDefaults

`func NewVpnsWithDefaults() *Vpns`

NewVpnsWithDefaults instantiates a new Vpns object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssigned

`func (o *Vpns) GetAssigned() int32`

GetAssigned returns the Assigned field if non-nil, zero value otherwise.

### GetAssignedOk

`func (o *Vpns) GetAssignedOk() (*int32, bool)`

GetAssignedOk returns a tuple with the Assigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigned

`func (o *Vpns) SetAssigned(v int32)`

SetAssigned sets Assigned field to given value.


### GetReassigned

`func (o *Vpns) GetReassigned() int32`

GetReassigned returns the Reassigned field if non-nil, zero value otherwise.

### GetReassignedOk

`func (o *Vpns) GetReassignedOk() (*int32, bool)`

GetReassignedOk returns a tuple with the Reassigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReassigned

`func (o *Vpns) SetReassigned(v int32)`

SetReassigned sets Reassigned field to given value.

### HasReassigned

`func (o *Vpns) HasReassigned() bool`

HasReassigned returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


