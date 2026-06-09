# PrivateSpaceSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to [**[]PrivateSpaceSummaryContentItem**](PrivateSpaceSummaryContentItem.md) |  | [optional] 
**Pageable** | Pointer to [**PrivateSpaceSummaryPageable**](PrivateSpaceSummaryPageable.md) |  | [optional] 
**Last** | Pointer to **bool** |  | [optional] 
**TotalElements** | Pointer to **int32** |  | [optional] 
**TotalPages** | Pointer to **int32** |  | [optional] 
**First** | Pointer to **bool** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Number** | Pointer to **int32** |  | [optional] 
**Sort** | Pointer to [**PrivateSpaceSummaryPageableSort**](PrivateSpaceSummaryPageableSort.md) |  | [optional] 
**NumberOfElements** | Pointer to **int32** |  | [optional] 
**Empty** | Pointer to **bool** |  | [optional] 

## Methods

### NewPrivateSpaceSummary

`func NewPrivateSpaceSummary() *PrivateSpaceSummary`

NewPrivateSpaceSummary instantiates a new PrivateSpaceSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateSpaceSummaryWithDefaults

`func NewPrivateSpaceSummaryWithDefaults() *PrivateSpaceSummary`

NewPrivateSpaceSummaryWithDefaults instantiates a new PrivateSpaceSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *PrivateSpaceSummary) GetContent() []PrivateSpaceSummaryContentItem`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PrivateSpaceSummary) GetContentOk() (*[]PrivateSpaceSummaryContentItem, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PrivateSpaceSummary) SetContent(v []PrivateSpaceSummaryContentItem)`

SetContent sets Content field to given value.

### HasContent

`func (o *PrivateSpaceSummary) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetPageable

`func (o *PrivateSpaceSummary) GetPageable() PrivateSpaceSummaryPageable`

GetPageable returns the Pageable field if non-nil, zero value otherwise.

### GetPageableOk

`func (o *PrivateSpaceSummary) GetPageableOk() (*PrivateSpaceSummaryPageable, bool)`

GetPageableOk returns a tuple with the Pageable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageable

`func (o *PrivateSpaceSummary) SetPageable(v PrivateSpaceSummaryPageable)`

SetPageable sets Pageable field to given value.

### HasPageable

`func (o *PrivateSpaceSummary) HasPageable() bool`

HasPageable returns a boolean if a field has been set.

### GetLast

`func (o *PrivateSpaceSummary) GetLast() bool`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *PrivateSpaceSummary) GetLastOk() (*bool, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *PrivateSpaceSummary) SetLast(v bool)`

SetLast sets Last field to given value.

### HasLast

`func (o *PrivateSpaceSummary) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetTotalElements

`func (o *PrivateSpaceSummary) GetTotalElements() int32`

GetTotalElements returns the TotalElements field if non-nil, zero value otherwise.

### GetTotalElementsOk

`func (o *PrivateSpaceSummary) GetTotalElementsOk() (*int32, bool)`

GetTotalElementsOk returns a tuple with the TotalElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalElements

`func (o *PrivateSpaceSummary) SetTotalElements(v int32)`

SetTotalElements sets TotalElements field to given value.

### HasTotalElements

`func (o *PrivateSpaceSummary) HasTotalElements() bool`

HasTotalElements returns a boolean if a field has been set.

### GetTotalPages

`func (o *PrivateSpaceSummary) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *PrivateSpaceSummary) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *PrivateSpaceSummary) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *PrivateSpaceSummary) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetFirst

`func (o *PrivateSpaceSummary) GetFirst() bool`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *PrivateSpaceSummary) GetFirstOk() (*bool, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *PrivateSpaceSummary) SetFirst(v bool)`

SetFirst sets First field to given value.

### HasFirst

`func (o *PrivateSpaceSummary) HasFirst() bool`

HasFirst returns a boolean if a field has been set.

### GetSize

`func (o *PrivateSpaceSummary) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PrivateSpaceSummary) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PrivateSpaceSummary) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *PrivateSpaceSummary) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetNumber

`func (o *PrivateSpaceSummary) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *PrivateSpaceSummary) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *PrivateSpaceSummary) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *PrivateSpaceSummary) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetSort

`func (o *PrivateSpaceSummary) GetSort() PrivateSpaceSummaryPageableSort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *PrivateSpaceSummary) GetSortOk() (*PrivateSpaceSummaryPageableSort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *PrivateSpaceSummary) SetSort(v PrivateSpaceSummaryPageableSort)`

SetSort sets Sort field to given value.

### HasSort

`func (o *PrivateSpaceSummary) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetNumberOfElements

`func (o *PrivateSpaceSummary) GetNumberOfElements() int32`

GetNumberOfElements returns the NumberOfElements field if non-nil, zero value otherwise.

### GetNumberOfElementsOk

`func (o *PrivateSpaceSummary) GetNumberOfElementsOk() (*int32, bool)`

GetNumberOfElementsOk returns a tuple with the NumberOfElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfElements

`func (o *PrivateSpaceSummary) SetNumberOfElements(v int32)`

SetNumberOfElements sets NumberOfElements field to given value.

### HasNumberOfElements

`func (o *PrivateSpaceSummary) HasNumberOfElements() bool`

HasNumberOfElements returns a boolean if a field has been set.

### GetEmpty

`func (o *PrivateSpaceSummary) GetEmpty() bool`

GetEmpty returns the Empty field if non-nil, zero value otherwise.

### GetEmptyOk

`func (o *PrivateSpaceSummary) GetEmptyOk() (*bool, bool)`

GetEmptyOk returns a tuple with the Empty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmpty

`func (o *PrivateSpaceSummary) SetEmpty(v bool)`

SetEmpty sets Empty field to given value.

### HasEmpty

`func (o *PrivateSpaceSummary) HasEmpty() bool`

HasEmpty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


