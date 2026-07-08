# PrivateSpaceVpnConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Connection id (GUID). | [optional] 
**Name** | Pointer to **string** | User-chosen display name for the VPN connection. | [optional] 
**Vpns** | Pointer to [**[]PrivateSpaceVpn**](PrivateSpaceVpn.md) | VPN members composing the connection. | [optional] 

## Methods

### NewPrivateSpaceVpnConnection

`func NewPrivateSpaceVpnConnection() *PrivateSpaceVpnConnection`

NewPrivateSpaceVpnConnection instantiates a new PrivateSpaceVpnConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateSpaceVpnConnectionWithDefaults

`func NewPrivateSpaceVpnConnectionWithDefaults() *PrivateSpaceVpnConnection`

NewPrivateSpaceVpnConnectionWithDefaults instantiates a new PrivateSpaceVpnConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrivateSpaceVpnConnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrivateSpaceVpnConnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrivateSpaceVpnConnection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PrivateSpaceVpnConnection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PrivateSpaceVpnConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrivateSpaceVpnConnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrivateSpaceVpnConnection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PrivateSpaceVpnConnection) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVpns

`func (o *PrivateSpaceVpnConnection) GetVpns() []PrivateSpaceVpn`

GetVpns returns the Vpns field if non-nil, zero value otherwise.

### GetVpnsOk

`func (o *PrivateSpaceVpnConnection) GetVpnsOk() (*[]PrivateSpaceVpn, bool)`

GetVpnsOk returns a tuple with the Vpns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpns

`func (o *PrivateSpaceVpnConnection) SetVpns(v []PrivateSpaceVpn)`

SetVpns sets Vpns field to given value.

### HasVpns

`func (o *PrivateSpaceVpnConnection) HasVpns() bool`

HasVpns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


