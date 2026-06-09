# GetSlaTiersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**Tiers** | Pointer to [**[]SlaTier**](SlaTier.md) |  | [optional] 

## Methods

### NewGetSlaTiersResponse

`func NewGetSlaTiersResponse() *GetSlaTiersResponse`

NewGetSlaTiersResponse instantiates a new GetSlaTiersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSlaTiersResponseWithDefaults

`func NewGetSlaTiersResponseWithDefaults() *GetSlaTiersResponse`

NewGetSlaTiersResponseWithDefaults instantiates a new GetSlaTiersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *GetSlaTiersResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetSlaTiersResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetSlaTiersResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetSlaTiersResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetTiers

`func (o *GetSlaTiersResponse) GetTiers() []SlaTier`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *GetSlaTiersResponse) GetTiersOk() (*[]SlaTier, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *GetSlaTiersResponse) SetTiers(v []SlaTier)`

SetTiers sets Tiers field to given value.

### HasTiers

`func (o *GetSlaTiersResponse) HasTiers() bool`

HasTiers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


