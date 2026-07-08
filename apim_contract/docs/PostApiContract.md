# PostApiContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to **int32** |  | [optional] 
**RequestedTierId** | Pointer to **int32** |  | [optional] 
**RequestAccessInfo** | Pointer to [**PostApiContractRequestAccessInfo**](PostApiContractRequestAccessInfo.md) |  | [optional] 

## Methods

### NewPostApiContract

`func NewPostApiContract() *PostApiContract`

NewPostApiContract instantiates a new PostApiContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiContractWithDefaults

`func NewPostApiContractWithDefaults() *PostApiContract`

NewPostApiContractWithDefaults instantiates a new PostApiContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *PostApiContract) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *PostApiContract) GetApplicationIdOk() (*int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *PostApiContract) SetApplicationId(v int32)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *PostApiContract) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetRequestedTierId

`func (o *PostApiContract) GetRequestedTierId() int32`

GetRequestedTierId returns the RequestedTierId field if non-nil, zero value otherwise.

### GetRequestedTierIdOk

`func (o *PostApiContract) GetRequestedTierIdOk() (*int32, bool)`

GetRequestedTierIdOk returns a tuple with the RequestedTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedTierId

`func (o *PostApiContract) SetRequestedTierId(v int32)`

SetRequestedTierId sets RequestedTierId field to given value.

### HasRequestedTierId

`func (o *PostApiContract) HasRequestedTierId() bool`

HasRequestedTierId returns a boolean if a field has been set.

### GetRequestAccessInfo

`func (o *PostApiContract) GetRequestAccessInfo() PostApiContractRequestAccessInfo`

GetRequestAccessInfo returns the RequestAccessInfo field if non-nil, zero value otherwise.

### GetRequestAccessInfoOk

`func (o *PostApiContract) GetRequestAccessInfoOk() (*PostApiContractRequestAccessInfo, bool)`

GetRequestAccessInfoOk returns a tuple with the RequestAccessInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestAccessInfo

`func (o *PostApiContract) SetRequestAccessInfo(v PostApiContractRequestAccessInfo)`

SetRequestAccessInfo sets RequestAccessInfo field to given value.

### HasRequestAccessInfo

`func (o *PostApiContract) HasRequestAccessInfo() bool`

HasRequestAccessInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


