# AssetsPost400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** |  | [optional] [default to 400]
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewAssetsPost400Response

`func NewAssetsPost400Response() *AssetsPost400Response`

NewAssetsPost400Response instantiates a new AssetsPost400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetsPost400ResponseWithDefaults

`func NewAssetsPost400ResponseWithDefaults() *AssetsPost400Response`

NewAssetsPost400ResponseWithDefaults instantiates a new AssetsPost400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AssetsPost400Response) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AssetsPost400Response) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AssetsPost400Response) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AssetsPost400Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *AssetsPost400Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AssetsPost400Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AssetsPost400Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AssetsPost400Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


