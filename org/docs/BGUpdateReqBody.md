# BGUpdateReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entitlements** | Pointer to [**EntitlementsCore**](EntitlementsCore.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the organization | [optional] [default to ""]
**OwnerId** | Pointer to **string** | The id of the owner of the organization. | [optional] [default to ""]
**SessionTimeout** | Pointer to **int32** | The session timeout in minutes | [optional] [default to 0]

## Methods

### NewBGUpdateReqBody

`func NewBGUpdateReqBody() *BGUpdateReqBody`

NewBGUpdateReqBody instantiates a new BGUpdateReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBGUpdateReqBodyWithDefaults

`func NewBGUpdateReqBodyWithDefaults() *BGUpdateReqBody`

NewBGUpdateReqBodyWithDefaults instantiates a new BGUpdateReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitlements

`func (o *BGUpdateReqBody) GetEntitlements() EntitlementsCore`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *BGUpdateReqBody) GetEntitlementsOk() (*EntitlementsCore, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *BGUpdateReqBody) SetEntitlements(v EntitlementsCore)`

SetEntitlements sets Entitlements field to given value.

### HasEntitlements

`func (o *BGUpdateReqBody) HasEntitlements() bool`

HasEntitlements returns a boolean if a field has been set.

### GetName

`func (o *BGUpdateReqBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BGUpdateReqBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BGUpdateReqBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BGUpdateReqBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwnerId

`func (o *BGUpdateReqBody) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *BGUpdateReqBody) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *BGUpdateReqBody) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *BGUpdateReqBody) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetSessionTimeout

`func (o *BGUpdateReqBody) GetSessionTimeout() int32`

GetSessionTimeout returns the SessionTimeout field if non-nil, zero value otherwise.

### GetSessionTimeoutOk

`func (o *BGUpdateReqBody) GetSessionTimeoutOk() (*int32, bool)`

GetSessionTimeoutOk returns a tuple with the SessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTimeout

`func (o *BGUpdateReqBody) SetSessionTimeout(v int32)`

SetSessionTimeout sets SessionTimeout field to given value.

### HasSessionTimeout

`func (o *BGUpdateReqBody) HasSessionTimeout() bool`

HasSessionTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


