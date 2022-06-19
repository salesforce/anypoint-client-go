# VpnCoreGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | Pointer to [**Spec**](Spec.md) |  | [optional] 
**State** | Pointer to [**State**](State.md) |  | [optional] 
**UpdateAvailable** | Pointer to **bool** |  | [optional] [default to false]
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewVpnCoreGet

`func NewVpnCoreGet() *VpnCoreGet`

NewVpnCoreGet instantiates a new VpnCoreGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnCoreGetWithDefaults

`func NewVpnCoreGetWithDefaults() *VpnCoreGet`

NewVpnCoreGetWithDefaults instantiates a new VpnCoreGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *VpnCoreGet) GetSpec() Spec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *VpnCoreGet) GetSpecOk() (*Spec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *VpnCoreGet) SetSpec(v Spec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *VpnCoreGet) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetState

`func (o *VpnCoreGet) GetState() State`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VpnCoreGet) GetStateOk() (*State, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VpnCoreGet) SetState(v State)`

SetState sets State field to given value.

### HasState

`func (o *VpnCoreGet) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdateAvailable

`func (o *VpnCoreGet) GetUpdateAvailable() bool`

GetUpdateAvailable returns the UpdateAvailable field if non-nil, zero value otherwise.

### GetUpdateAvailableOk

`func (o *VpnCoreGet) GetUpdateAvailableOk() (*bool, bool)`

GetUpdateAvailableOk returns a tuple with the UpdateAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAvailable

`func (o *VpnCoreGet) SetUpdateAvailable(v bool)`

SetUpdateAvailable sets UpdateAvailable field to given value.

### HasUpdateAvailable

`func (o *VpnCoreGet) HasUpdateAvailable() bool`

HasUpdateAvailable returns a boolean if a field has been set.

### GetName

`func (o *VpnCoreGet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpnCoreGet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpnCoreGet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VpnCoreGet) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


