# VpnGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The vpn Id | 
**Spec** | Pointer to [**Spec**](Spec.md) |  | [optional] 
**State** | Pointer to [**State**](State.md) |  | [optional] 
**UpdateAvailable** | Pointer to **bool** |  | [optional] [default to false]
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewVpnGet

`func NewVpnGet(id string, ) *VpnGet`

NewVpnGet instantiates a new VpnGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnGetWithDefaults

`func NewVpnGetWithDefaults() *VpnGet`

NewVpnGetWithDefaults instantiates a new VpnGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VpnGet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VpnGet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VpnGet) SetId(v string)`

SetId sets Id field to given value.


### GetSpec

`func (o *VpnGet) GetSpec() Spec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *VpnGet) GetSpecOk() (*Spec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *VpnGet) SetSpec(v Spec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *VpnGet) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetState

`func (o *VpnGet) GetState() State`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VpnGet) GetStateOk() (*State, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VpnGet) SetState(v State)`

SetState sets State field to given value.

### HasState

`func (o *VpnGet) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdateAvailable

`func (o *VpnGet) GetUpdateAvailable() bool`

GetUpdateAvailable returns the UpdateAvailable field if non-nil, zero value otherwise.

### GetUpdateAvailableOk

`func (o *VpnGet) GetUpdateAvailableOk() (*bool, bool)`

GetUpdateAvailableOk returns a tuple with the UpdateAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAvailable

`func (o *VpnGet) SetUpdateAvailable(v bool)`

SetUpdateAvailable sets UpdateAvailable field to given value.

### HasUpdateAvailable

`func (o *VpnGet) HasUpdateAvailable() bool`

HasUpdateAvailable returns a boolean if a field has been set.

### GetName

`func (o *VpnGet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpnGet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpnGet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VpnGet) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


