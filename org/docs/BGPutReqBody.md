# BGPutReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entitlements** | [**Entitlements**](Entitlements.md) |  | 
**Name** | [**Name**](Name.md) |  | 
**OwnerId** | [**OwnerId**](OwnerId.md) |  | 
**ParentOrganizationId** | **string** | An explanation about the purpose of this instance. | [default to ""]
**SessionTimeout** | **int32** | An explanation about the purpose of this instance. | [default to 0]

## Methods

### NewBGPutReqBody

`func NewBGPutReqBody(entitlements Entitlements, name Name, ownerId OwnerId, parentOrganizationId string, sessionTimeout int32, ) *BGPutReqBody`

NewBGPutReqBody instantiates a new BGPutReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBGPutReqBodyWithDefaults

`func NewBGPutReqBodyWithDefaults() *BGPutReqBody`

NewBGPutReqBodyWithDefaults instantiates a new BGPutReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitlements

`func (o *BGPutReqBody) GetEntitlements() Entitlements`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *BGPutReqBody) GetEntitlementsOk() (*Entitlements, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *BGPutReqBody) SetEntitlements(v Entitlements)`

SetEntitlements sets Entitlements field to given value.


### GetName

`func (o *BGPutReqBody) GetName() Name`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BGPutReqBody) GetNameOk() (*Name, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BGPutReqBody) SetName(v Name)`

SetName sets Name field to given value.


### GetOwnerId

`func (o *BGPutReqBody) GetOwnerId() OwnerId`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *BGPutReqBody) GetOwnerIdOk() (*OwnerId, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *BGPutReqBody) SetOwnerId(v OwnerId)`

SetOwnerId sets OwnerId field to given value.


### GetParentOrganizationId

`func (o *BGPutReqBody) GetParentOrganizationId() string`

GetParentOrganizationId returns the ParentOrganizationId field if non-nil, zero value otherwise.

### GetParentOrganizationIdOk

`func (o *BGPutReqBody) GetParentOrganizationIdOk() (*string, bool)`

GetParentOrganizationIdOk returns a tuple with the ParentOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentOrganizationId

`func (o *BGPutReqBody) SetParentOrganizationId(v string)`

SetParentOrganizationId sets ParentOrganizationId field to given value.


### GetSessionTimeout

`func (o *BGPutReqBody) GetSessionTimeout() int32`

GetSessionTimeout returns the SessionTimeout field if non-nil, zero value otherwise.

### GetSessionTimeoutOk

`func (o *BGPutReqBody) GetSessionTimeoutOk() (*int32, bool)`

GetSessionTimeoutOk returns a tuple with the SessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTimeout

`func (o *BGPutReqBody) SetSessionTimeout(v int32)`

SetSessionTimeout sets SessionTimeout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


