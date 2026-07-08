# LogLevel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageName** | Pointer to **string** |  | [optional] 
**Level** | Pointer to **string** |  | [optional] 

## Methods

### NewLogLevel

`func NewLogLevel() *LogLevel`

NewLogLevel instantiates a new LogLevel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogLevelWithDefaults

`func NewLogLevelWithDefaults() *LogLevel`

NewLogLevelWithDefaults instantiates a new LogLevel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageName

`func (o *LogLevel) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *LogLevel) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *LogLevel) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.

### HasPackageName

`func (o *LogLevel) HasPackageName() bool`

HasPackageName returns a boolean if a field has been set.

### GetLevel

`func (o *LogLevel) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *LogLevel) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *LogLevel) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *LogLevel) HasLevel() bool`

HasLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


