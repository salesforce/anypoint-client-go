# ApiQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | API Query enabled | [optional] [default to false]
**ProductSKU** | Pointer to **int32** | Product SKU | [optional] [default to 0]

## Methods

### NewApiQuery

`func NewApiQuery() *ApiQuery`

NewApiQuery instantiates a new ApiQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiQueryWithDefaults

`func NewApiQueryWithDefaults() *ApiQuery`

NewApiQueryWithDefaults instantiates a new ApiQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ApiQuery) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApiQuery) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApiQuery) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApiQuery) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetProductSKU

`func (o *ApiQuery) GetProductSKU() int32`

GetProductSKU returns the ProductSKU field if non-nil, zero value otherwise.

### GetProductSKUOk

`func (o *ApiQuery) GetProductSKUOk() (*int32, bool)`

GetProductSKUOk returns a tuple with the ProductSKU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductSKU

`func (o *ApiQuery) SetProductSKU(v int32)`

SetProductSKU sets ProductSKU field to given value.

### HasProductSKU

`func (o *ApiQuery) HasProductSKU() bool`

HasProductSKU returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


