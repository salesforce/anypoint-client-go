# ApiQueryEntitlement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | API Query enabled | [optional] [default to false]
**ProductSKU** | Pointer to **int32** | Product SKU | [optional] [default to 0]
**Sandbox** | Pointer to **int32** | Product SKU | [optional] 
**Production** | Pointer to **int32** | Product SKU | [optional] 

## Methods

### NewApiQueryEntitlement

`func NewApiQueryEntitlement() *ApiQueryEntitlement`

NewApiQueryEntitlement instantiates a new ApiQueryEntitlement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiQueryEntitlementWithDefaults

`func NewApiQueryEntitlementWithDefaults() *ApiQueryEntitlement`

NewApiQueryEntitlementWithDefaults instantiates a new ApiQueryEntitlement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ApiQueryEntitlement) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApiQueryEntitlement) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApiQueryEntitlement) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApiQueryEntitlement) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetProductSKU

`func (o *ApiQueryEntitlement) GetProductSKU() int32`

GetProductSKU returns the ProductSKU field if non-nil, zero value otherwise.

### GetProductSKUOk

`func (o *ApiQueryEntitlement) GetProductSKUOk() (*int32, bool)`

GetProductSKUOk returns a tuple with the ProductSKU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductSKU

`func (o *ApiQueryEntitlement) SetProductSKU(v int32)`

SetProductSKU sets ProductSKU field to given value.

### HasProductSKU

`func (o *ApiQueryEntitlement) HasProductSKU() bool`

HasProductSKU returns a boolean if a field has been set.

### GetSandbox

`func (o *ApiQueryEntitlement) GetSandbox() int32`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *ApiQueryEntitlement) GetSandboxOk() (*int32, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandbox

`func (o *ApiQueryEntitlement) SetSandbox(v int32)`

SetSandbox sets Sandbox field to given value.

### HasSandbox

`func (o *ApiQueryEntitlement) HasSandbox() bool`

HasSandbox returns a boolean if a field has been set.

### GetProduction

`func (o *ApiQueryEntitlement) GetProduction() int32`

GetProduction returns the Production field if non-nil, zero value otherwise.

### GetProductionOk

`func (o *ApiQueryEntitlement) GetProductionOk() (*int32, bool)`

GetProductionOk returns a tuple with the Production field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduction

`func (o *ApiQueryEntitlement) SetProduction(v int32)`

SetProduction sets Production field to given value.

### HasProduction

`func (o *ApiQueryEntitlement) HasProduction() bool`

HasProduction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


