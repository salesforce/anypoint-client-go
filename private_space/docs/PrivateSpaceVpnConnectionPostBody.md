# PrivateSpaceVpnConnectionPostBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name for the VPN connection. | 
**Vpns** | [**[]PrivateSpaceVpnPostBody**](PrivateSpaceVpnPostBody.md) | VPN members to create under this connection. | 

## Methods

### NewPrivateSpaceVpnConnectionPostBody

`func NewPrivateSpaceVpnConnectionPostBody(name string, vpns []PrivateSpaceVpnPostBody, ) *PrivateSpaceVpnConnectionPostBody`

NewPrivateSpaceVpnConnectionPostBody instantiates a new PrivateSpaceVpnConnectionPostBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateSpaceVpnConnectionPostBodyWithDefaults

`func NewPrivateSpaceVpnConnectionPostBodyWithDefaults() *PrivateSpaceVpnConnectionPostBody`

NewPrivateSpaceVpnConnectionPostBodyWithDefaults instantiates a new PrivateSpaceVpnConnectionPostBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PrivateSpaceVpnConnectionPostBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrivateSpaceVpnConnectionPostBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrivateSpaceVpnConnectionPostBody) SetName(v string)`

SetName sets Name field to given value.


### GetVpns

`func (o *PrivateSpaceVpnConnectionPostBody) GetVpns() []PrivateSpaceVpnPostBody`

GetVpns returns the Vpns field if non-nil, zero value otherwise.

### GetVpnsOk

`func (o *PrivateSpaceVpnConnectionPostBody) GetVpnsOk() (*[]PrivateSpaceVpnPostBody, bool)`

GetVpnsOk returns a tuple with the Vpns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpns

`func (o *PrivateSpaceVpnConnectionPostBody) SetVpns(v []PrivateSpaceVpnPostBody)`

SetVpns sets Vpns field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


