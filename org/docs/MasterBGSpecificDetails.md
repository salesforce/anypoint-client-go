# MasterBGSpecificDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionTimeout** | Pointer to **int32** | The session timeout in minutes. | [optional] [default to 0]
**Subscription** | Pointer to [**Subscription**](Subscription.md) |  | [optional] 
**GdotId** | Pointer to **string** | The gdot id of the organization | [optional] [default to ""]
**DeletedAt** | Pointer to **NullableString** | The deleted date of the organization | [optional] 
**OrgType** | Pointer to **string** | The type of the organization | [optional] [default to ""]

## Methods

### NewMasterBGSpecificDetails

`func NewMasterBGSpecificDetails() *MasterBGSpecificDetails`

NewMasterBGSpecificDetails instantiates a new MasterBGSpecificDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMasterBGSpecificDetailsWithDefaults

`func NewMasterBGSpecificDetailsWithDefaults() *MasterBGSpecificDetails`

NewMasterBGSpecificDetailsWithDefaults instantiates a new MasterBGSpecificDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionTimeout

`func (o *MasterBGSpecificDetails) GetSessionTimeout() int32`

GetSessionTimeout returns the SessionTimeout field if non-nil, zero value otherwise.

### GetSessionTimeoutOk

`func (o *MasterBGSpecificDetails) GetSessionTimeoutOk() (*int32, bool)`

GetSessionTimeoutOk returns a tuple with the SessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTimeout

`func (o *MasterBGSpecificDetails) SetSessionTimeout(v int32)`

SetSessionTimeout sets SessionTimeout field to given value.

### HasSessionTimeout

`func (o *MasterBGSpecificDetails) HasSessionTimeout() bool`

HasSessionTimeout returns a boolean if a field has been set.

### GetSubscription

`func (o *MasterBGSpecificDetails) GetSubscription() Subscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *MasterBGSpecificDetails) GetSubscriptionOk() (*Subscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *MasterBGSpecificDetails) SetSubscription(v Subscription)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *MasterBGSpecificDetails) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetGdotId

`func (o *MasterBGSpecificDetails) GetGdotId() string`

GetGdotId returns the GdotId field if non-nil, zero value otherwise.

### GetGdotIdOk

`func (o *MasterBGSpecificDetails) GetGdotIdOk() (*string, bool)`

GetGdotIdOk returns a tuple with the GdotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGdotId

`func (o *MasterBGSpecificDetails) SetGdotId(v string)`

SetGdotId sets GdotId field to given value.

### HasGdotId

`func (o *MasterBGSpecificDetails) HasGdotId() bool`

HasGdotId returns a boolean if a field has been set.

### GetDeletedAt

`func (o *MasterBGSpecificDetails) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *MasterBGSpecificDetails) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *MasterBGSpecificDetails) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *MasterBGSpecificDetails) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *MasterBGSpecificDetails) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *MasterBGSpecificDetails) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetOrgType

`func (o *MasterBGSpecificDetails) GetOrgType() string`

GetOrgType returns the OrgType field if non-nil, zero value otherwise.

### GetOrgTypeOk

`func (o *MasterBGSpecificDetails) GetOrgTypeOk() (*string, bool)`

GetOrgTypeOk returns a tuple with the OrgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgType

`func (o *MasterBGSpecificDetails) SetOrgType(v string)`

SetOrgType sets OrgType field to given value.

### HasOrgType

`func (o *MasterBGSpecificDetails) HasOrgType() bool`

HasOrgType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


