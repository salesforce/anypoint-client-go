# GetContractsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**Contracts** | Pointer to [**[]Contract**](Contract.md) |  | [optional] 

## Methods

### NewGetContractsResponse

`func NewGetContractsResponse() *GetContractsResponse`

NewGetContractsResponse instantiates a new GetContractsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetContractsResponseWithDefaults

`func NewGetContractsResponseWithDefaults() *GetContractsResponse`

NewGetContractsResponseWithDefaults instantiates a new GetContractsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *GetContractsResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetContractsResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetContractsResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetContractsResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetContracts

`func (o *GetContractsResponse) GetContracts() []Contract`

GetContracts returns the Contracts field if non-nil, zero value otherwise.

### GetContractsOk

`func (o *GetContractsResponse) GetContractsOk() (*[]Contract, bool)`

GetContractsOk returns a tuple with the Contracts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContracts

`func (o *GetContractsResponse) SetContracts(v []Contract)`

SetContracts sets Contracts field to given value.

### HasContracts

`func (o *GetContractsResponse) HasContracts() bool`

HasContracts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


