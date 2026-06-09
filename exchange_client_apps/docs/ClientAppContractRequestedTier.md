# ClientAppContractRequestedTier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Limits** | Pointer to [**[]ClientAppContractRequestedTierLimitsInner**](ClientAppContractRequestedTierLimitsInner.md) |  | [optional] 

## Methods

### NewClientAppContractRequestedTier

`func NewClientAppContractRequestedTier() *ClientAppContractRequestedTier`

NewClientAppContractRequestedTier instantiates a new ClientAppContractRequestedTier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientAppContractRequestedTierWithDefaults

`func NewClientAppContractRequestedTierWithDefaults() *ClientAppContractRequestedTier`

NewClientAppContractRequestedTierWithDefaults instantiates a new ClientAppContractRequestedTier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClientAppContractRequestedTier) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClientAppContractRequestedTier) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClientAppContractRequestedTier) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ClientAppContractRequestedTier) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ClientAppContractRequestedTier) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClientAppContractRequestedTier) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClientAppContractRequestedTier) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClientAppContractRequestedTier) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ClientAppContractRequestedTier) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClientAppContractRequestedTier) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClientAppContractRequestedTier) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClientAppContractRequestedTier) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLimits

`func (o *ClientAppContractRequestedTier) GetLimits() []ClientAppContractRequestedTierLimitsInner`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *ClientAppContractRequestedTier) GetLimitsOk() (*[]ClientAppContractRequestedTierLimitsInner, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *ClientAppContractRequestedTier) SetLimits(v []ClientAppContractRequestedTierLimitsInner)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *ClientAppContractRequestedTier) HasLimits() bool`

HasLimits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


