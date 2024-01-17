# ConnectedApplicationsGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ConnectedAppRespExt**](ConnectedAppRespExt.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewConnectedApplicationsGet200Response

`func NewConnectedApplicationsGet200Response() *ConnectedApplicationsGet200Response`

NewConnectedApplicationsGet200Response instantiates a new ConnectedApplicationsGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectedApplicationsGet200ResponseWithDefaults

`func NewConnectedApplicationsGet200ResponseWithDefaults() *ConnectedApplicationsGet200Response`

NewConnectedApplicationsGet200ResponseWithDefaults instantiates a new ConnectedApplicationsGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ConnectedApplicationsGet200Response) GetData() []ConnectedAppRespExt`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ConnectedApplicationsGet200Response) GetDataOk() (*[]ConnectedAppRespExt, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ConnectedApplicationsGet200Response) SetData(v []ConnectedAppRespExt)`

SetData sets Data field to given value.

### HasData

`func (o *ConnectedApplicationsGet200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotal

`func (o *ConnectedApplicationsGet200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ConnectedApplicationsGet200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ConnectedApplicationsGet200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ConnectedApplicationsGet200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


