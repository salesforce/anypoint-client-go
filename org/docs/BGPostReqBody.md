# BGPostReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entitlements** | [**EntitlementsCore**](EntitlementsCore.md) |  | 
**Name** | **string** | The name of the organization | [default to ""]
**OwnerId** | **string** | The id of the owner of the organization. | [default to ""]
**ParentOrganizationId** | **string** | a reference to the parent organization | [default to ""]

## Methods

### NewBGPostReqBody

`func NewBGPostReqBody(entitlements EntitlementsCore, name string, ownerId string, parentOrganizationId string, ) *BGPostReqBody`

NewBGPostReqBody instantiates a new BGPostReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBGPostReqBodyWithDefaults

`func NewBGPostReqBodyWithDefaults() *BGPostReqBody`

NewBGPostReqBodyWithDefaults instantiates a new BGPostReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitlements

`func (o *BGPostReqBody) GetEntitlements() EntitlementsCore`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *BGPostReqBody) GetEntitlementsOk() (*EntitlementsCore, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *BGPostReqBody) SetEntitlements(v EntitlementsCore)`

SetEntitlements sets Entitlements field to given value.


### GetName

`func (o *BGPostReqBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BGPostReqBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BGPostReqBody) SetName(v string)`

SetName sets Name field to given value.


### GetOwnerId

`func (o *BGPostReqBody) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *BGPostReqBody) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *BGPostReqBody) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetParentOrganizationId

`func (o *BGPostReqBody) GetParentOrganizationId() string`

GetParentOrganizationId returns the ParentOrganizationId field if non-nil, zero value otherwise.

### GetParentOrganizationIdOk

`func (o *BGPostReqBody) GetParentOrganizationIdOk() (*string, bool)`

GetParentOrganizationIdOk returns a tuple with the ParentOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentOrganizationId

`func (o *BGPostReqBody) SetParentOrganizationId(v string)`

SetParentOrganizationId sets ParentOrganizationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


