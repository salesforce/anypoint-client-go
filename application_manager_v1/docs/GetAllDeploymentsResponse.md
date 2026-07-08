# GetAllDeploymentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Deployment**](Deployment.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetAllDeploymentsResponse

`func NewGetAllDeploymentsResponse() *GetAllDeploymentsResponse`

NewGetAllDeploymentsResponse instantiates a new GetAllDeploymentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllDeploymentsResponseWithDefaults

`func NewGetAllDeploymentsResponseWithDefaults() *GetAllDeploymentsResponse`

NewGetAllDeploymentsResponseWithDefaults instantiates a new GetAllDeploymentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetAllDeploymentsResponse) GetData() []Deployment`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAllDeploymentsResponse) GetDataOk() (*[]Deployment, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAllDeploymentsResponse) SetData(v []Deployment)`

SetData sets Data field to given value.

### HasData

`func (o *GetAllDeploymentsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotal

`func (o *GetAllDeploymentsResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetAllDeploymentsResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetAllDeploymentsResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetAllDeploymentsResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


