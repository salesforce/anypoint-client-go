# RolesGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Role**](Role.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewRolesGet200Response

`func NewRolesGet200Response() *RolesGet200Response`

NewRolesGet200Response instantiates a new RolesGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolesGet200ResponseWithDefaults

`func NewRolesGet200ResponseWithDefaults() *RolesGet200Response`

NewRolesGet200ResponseWithDefaults instantiates a new RolesGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RolesGet200Response) GetData() []Role`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RolesGet200Response) GetDataOk() (*[]Role, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RolesGet200Response) SetData(v []Role)`

SetData sets Data field to given value.

### HasData

`func (o *RolesGet200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotal

`func (o *RolesGet200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *RolesGet200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *RolesGet200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *RolesGet200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


