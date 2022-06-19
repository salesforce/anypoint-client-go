# VpnPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The vpn Id | 
**Spec** | Pointer to [**Spec1**](Spec1.md) |  | [optional] 
**State** | Pointer to [**State1**](State1.md) |  | [optional] 
**UpdateAvailable** | Pointer to **bool** |  | [optional] [default to false]
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewVpnPost

`func NewVpnPost(id string, ) *VpnPost`

NewVpnPost instantiates a new VpnPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnPostWithDefaults

`func NewVpnPostWithDefaults() *VpnPost`

NewVpnPostWithDefaults instantiates a new VpnPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VpnPost) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VpnPost) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VpnPost) SetId(v string)`

SetId sets Id field to given value.


### GetSpec

`func (o *VpnPost) GetSpec() Spec1`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *VpnPost) GetSpecOk() (*Spec1, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *VpnPost) SetSpec(v Spec1)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *VpnPost) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetState

`func (o *VpnPost) GetState() State1`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VpnPost) GetStateOk() (*State1, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VpnPost) SetState(v State1)`

SetState sets State field to given value.

### HasState

`func (o *VpnPost) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdateAvailable

`func (o *VpnPost) GetUpdateAvailable() bool`

GetUpdateAvailable returns the UpdateAvailable field if non-nil, zero value otherwise.

### GetUpdateAvailableOk

`func (o *VpnPost) GetUpdateAvailableOk() (*bool, bool)`

GetUpdateAvailableOk returns a tuple with the UpdateAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAvailable

`func (o *VpnPost) SetUpdateAvailable(v bool)`

SetUpdateAvailable sets UpdateAvailable field to given value.

### HasUpdateAvailable

`func (o *VpnPost) HasUpdateAvailable() bool`

HasUpdateAvailable returns a boolean if a field has been set.

### GetName

`func (o *VpnPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpnPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpnPost) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VpnPost) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


