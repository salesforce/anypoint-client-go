# DesignCenter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Api** | Pointer to **bool** | API entitlements | [optional] [default to false]
**Mozart** | Pointer to **bool** | Mozart entitlements | [optional] [default to false]

## Methods

### NewDesignCenter

`func NewDesignCenter() *DesignCenter`

NewDesignCenter instantiates a new DesignCenter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesignCenterWithDefaults

`func NewDesignCenterWithDefaults() *DesignCenter`

NewDesignCenterWithDefaults instantiates a new DesignCenter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApi

`func (o *DesignCenter) GetApi() bool`

GetApi returns the Api field if non-nil, zero value otherwise.

### GetApiOk

`func (o *DesignCenter) GetApiOk() (*bool, bool)`

GetApiOk returns a tuple with the Api field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApi

`func (o *DesignCenter) SetApi(v bool)`

SetApi sets Api field to given value.

### HasApi

`func (o *DesignCenter) HasApi() bool`

HasApi returns a boolean if a field has been set.

### GetMozart

`func (o *DesignCenter) GetMozart() bool`

GetMozart returns the Mozart field if non-nil, zero value otherwise.

### GetMozartOk

`func (o *DesignCenter) GetMozartOk() (*bool, bool)`

GetMozartOk returns a tuple with the Mozart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMozart

`func (o *DesignCenter) SetMozart(v bool)`

SetMozart sets Mozart field to given value.

### HasMozart

`func (o *DesignCenter) HasMozart() bool`

HasMozart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


