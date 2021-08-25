# ContextParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Org** | Pointer to **string** |  | [optional] 
**EnvId** | Pointer to **string** |  | [optional] 

## Methods

### NewContextParams

`func NewContextParams() *ContextParams`

NewContextParams instantiates a new ContextParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextParamsWithDefaults

`func NewContextParamsWithDefaults() *ContextParams`

NewContextParamsWithDefaults instantiates a new ContextParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrg

`func (o *ContextParams) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *ContextParams) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *ContextParams) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *ContextParams) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetEnvId

`func (o *ContextParams) GetEnvId() string`

GetEnvId returns the EnvId field if non-nil, zero value otherwise.

### GetEnvIdOk

`func (o *ContextParams) GetEnvIdOk() (*string, bool)`

GetEnvIdOk returns a tuple with the EnvId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvId

`func (o *ContextParams) SetEnvId(v string)`

SetEnvId sets EnvId field to given value.

### HasEnvId

`func (o *ContextParams) HasEnvId() bool`

HasEnvId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


