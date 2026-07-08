# ClientAppContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**RequestedTier** | Pointer to [**ClientAppContractRequestedTier**](ClientAppContractRequestedTier.md) |  | [optional] 

## Methods

### NewClientAppContract

`func NewClientAppContract() *ClientAppContract`

NewClientAppContract instantiates a new ClientAppContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientAppContractWithDefaults

`func NewClientAppContractWithDefaults() *ClientAppContract`

NewClientAppContractWithDefaults instantiates a new ClientAppContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClientAppContract) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClientAppContract) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClientAppContract) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ClientAppContract) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *ClientAppContract) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClientAppContract) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClientAppContract) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClientAppContract) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRequestedTier

`func (o *ClientAppContract) GetRequestedTier() ClientAppContractRequestedTier`

GetRequestedTier returns the RequestedTier field if non-nil, zero value otherwise.

### GetRequestedTierOk

`func (o *ClientAppContract) GetRequestedTierOk() (*ClientAppContractRequestedTier, bool)`

GetRequestedTierOk returns a tuple with the RequestedTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedTier

`func (o *ClientAppContract) SetRequestedTier(v ClientAppContractRequestedTier)`

SetRequestedTier sets RequestedTier field to given value.

### HasRequestedTier

`func (o *ClientAppContract) HasRequestedTier() bool`

HasRequestedTier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


