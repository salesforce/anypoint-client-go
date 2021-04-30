# Vpcs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assigned** | **int32** | An explanation about the purpose of this instance. | [default to 0]
**Reassigned** | Pointer to **int32** | An explanation about the purpose of this instance. | [optional] [default to 0]

## Methods

### NewVpcs

`func NewVpcs(assigned int32, ) *Vpcs`

NewVpcs instantiates a new Vpcs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpcsWithDefaults

`func NewVpcsWithDefaults() *Vpcs`

NewVpcsWithDefaults instantiates a new Vpcs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssigned

`func (o *Vpcs) GetAssigned() int32`

GetAssigned returns the Assigned field if non-nil, zero value otherwise.

### GetAssignedOk

`func (o *Vpcs) GetAssignedOk() (*int32, bool)`

GetAssignedOk returns a tuple with the Assigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigned

`func (o *Vpcs) SetAssigned(v int32)`

SetAssigned sets Assigned field to given value.


### GetReassigned

`func (o *Vpcs) GetReassigned() int32`

GetReassigned returns the Reassigned field if non-nil, zero value otherwise.

### GetReassignedOk

`func (o *Vpcs) GetReassignedOk() (*int32, bool)`

GetReassignedOk returns a tuple with the Reassigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReassigned

`func (o *Vpcs) SetReassigned(v int32)`

SetReassigned sets Reassigned field to given value.

### HasReassigned

`func (o *Vpcs) HasReassigned() bool`

HasReassigned returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


