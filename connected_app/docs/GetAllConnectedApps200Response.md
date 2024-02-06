# GetAllConnectedApps200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ConnectedAppRespExt**](ConnectedAppRespExt.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetAllConnectedApps200Response

`func NewGetAllConnectedApps200Response() *GetAllConnectedApps200Response`

NewGetAllConnectedApps200Response instantiates a new GetAllConnectedApps200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllConnectedApps200ResponseWithDefaults

`func NewGetAllConnectedApps200ResponseWithDefaults() *GetAllConnectedApps200Response`

NewGetAllConnectedApps200ResponseWithDefaults instantiates a new GetAllConnectedApps200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetAllConnectedApps200Response) GetData() []ConnectedAppRespExt`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAllConnectedApps200Response) GetDataOk() (*[]ConnectedAppRespExt, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAllConnectedApps200Response) SetData(v []ConnectedAppRespExt)`

SetData sets Data field to given value.

### HasData

`func (o *GetAllConnectedApps200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotal

`func (o *GetAllConnectedApps200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetAllConnectedApps200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetAllConnectedApps200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetAllConnectedApps200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


