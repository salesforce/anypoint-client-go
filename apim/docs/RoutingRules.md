# RoutingRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Methods** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Headers** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewRoutingRules

`func NewRoutingRules() *RoutingRules`

NewRoutingRules instantiates a new RoutingRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingRulesWithDefaults

`func NewRoutingRulesWithDefaults() *RoutingRules`

NewRoutingRulesWithDefaults instantiates a new RoutingRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethods

`func (o *RoutingRules) GetMethods() string`

GetMethods returns the Methods field if non-nil, zero value otherwise.

### GetMethodsOk

`func (o *RoutingRules) GetMethodsOk() (*string, bool)`

GetMethodsOk returns a tuple with the Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethods

`func (o *RoutingRules) SetMethods(v string)`

SetMethods sets Methods field to given value.

### HasMethods

`func (o *RoutingRules) HasMethods() bool`

HasMethods returns a boolean if a field has been set.

### GetHost

`func (o *RoutingRules) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *RoutingRules) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *RoutingRules) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *RoutingRules) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPath

`func (o *RoutingRules) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RoutingRules) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RoutingRules) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *RoutingRules) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetHeaders

`func (o *RoutingRules) GetHeaders() map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *RoutingRules) GetHeadersOk() (*map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *RoutingRules) SetHeaders(v map[string]interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *RoutingRules) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


