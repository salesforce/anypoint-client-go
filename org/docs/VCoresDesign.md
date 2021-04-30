# VCoresDesign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assigned** | **float32** | An explanation about the purpose of this instance. | [default to 0]
**Reassigned** | Pointer to **float32** | An explanation about the purpose of this instance. | [optional] [default to 0.0]

## Methods

### NewVCoresDesign

`func NewVCoresDesign(assigned float32, ) *VCoresDesign`

NewVCoresDesign instantiates a new VCoresDesign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVCoresDesignWithDefaults

`func NewVCoresDesignWithDefaults() *VCoresDesign`

NewVCoresDesignWithDefaults instantiates a new VCoresDesign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssigned

`func (o *VCoresDesign) GetAssigned() float32`

GetAssigned returns the Assigned field if non-nil, zero value otherwise.

### GetAssignedOk

`func (o *VCoresDesign) GetAssignedOk() (*float32, bool)`

GetAssignedOk returns a tuple with the Assigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigned

`func (o *VCoresDesign) SetAssigned(v float32)`

SetAssigned sets Assigned field to given value.


### GetReassigned

`func (o *VCoresDesign) GetReassigned() float32`

GetReassigned returns the Reassigned field if non-nil, zero value otherwise.

### GetReassignedOk

`func (o *VCoresDesign) GetReassignedOk() (*float32, bool)`

GetReassignedOk returns a tuple with the Reassigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReassigned

`func (o *VCoresDesign) SetReassigned(v float32)`

SetReassigned sets Reassigned field to given value.

### HasReassigned

`func (o *VCoresDesign) HasReassigned() bool`

HasReassigned returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


