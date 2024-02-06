# PolicyConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PropertyName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Options** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Optional** | Pointer to **bool** |  | [optional] 
**Sensitive** | Pointer to **bool** |  | [optional] 
**AllowMultiple** | Pointer to **bool** |  | [optional] 
**Configuration** | Pointer to [**[]PolicyConfigurationConfigurationInner**](PolicyConfigurationConfigurationInner.md) |  | [optional] 

## Methods

### NewPolicyConfiguration

`func NewPolicyConfiguration() *PolicyConfiguration`

NewPolicyConfiguration instantiates a new PolicyConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyConfigurationWithDefaults

`func NewPolicyConfigurationWithDefaults() *PolicyConfiguration`

NewPolicyConfigurationWithDefaults instantiates a new PolicyConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPropertyName

`func (o *PolicyConfiguration) GetPropertyName() string`

GetPropertyName returns the PropertyName field if non-nil, zero value otherwise.

### GetPropertyNameOk

`func (o *PolicyConfiguration) GetPropertyNameOk() (*string, bool)`

GetPropertyNameOk returns a tuple with the PropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyName

`func (o *PolicyConfiguration) SetPropertyName(v string)`

SetPropertyName sets PropertyName field to given value.

### HasPropertyName

`func (o *PolicyConfiguration) HasPropertyName() bool`

HasPropertyName returns a boolean if a field has been set.

### GetName

`func (o *PolicyConfiguration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyConfiguration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyConfiguration) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PolicyConfiguration) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PolicyConfiguration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyConfiguration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyConfiguration) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyConfiguration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *PolicyConfiguration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PolicyConfiguration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PolicyConfiguration) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PolicyConfiguration) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOptions

`func (o *PolicyConfiguration) GetOptions() []map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *PolicyConfiguration) GetOptionsOk() (*[]map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *PolicyConfiguration) SetOptions(v []map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *PolicyConfiguration) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetOptional

`func (o *PolicyConfiguration) GetOptional() bool`

GetOptional returns the Optional field if non-nil, zero value otherwise.

### GetOptionalOk

`func (o *PolicyConfiguration) GetOptionalOk() (*bool, bool)`

GetOptionalOk returns a tuple with the Optional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptional

`func (o *PolicyConfiguration) SetOptional(v bool)`

SetOptional sets Optional field to given value.

### HasOptional

`func (o *PolicyConfiguration) HasOptional() bool`

HasOptional returns a boolean if a field has been set.

### GetSensitive

`func (o *PolicyConfiguration) GetSensitive() bool`

GetSensitive returns the Sensitive field if non-nil, zero value otherwise.

### GetSensitiveOk

`func (o *PolicyConfiguration) GetSensitiveOk() (*bool, bool)`

GetSensitiveOk returns a tuple with the Sensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitive

`func (o *PolicyConfiguration) SetSensitive(v bool)`

SetSensitive sets Sensitive field to given value.

### HasSensitive

`func (o *PolicyConfiguration) HasSensitive() bool`

HasSensitive returns a boolean if a field has been set.

### GetAllowMultiple

`func (o *PolicyConfiguration) GetAllowMultiple() bool`

GetAllowMultiple returns the AllowMultiple field if non-nil, zero value otherwise.

### GetAllowMultipleOk

`func (o *PolicyConfiguration) GetAllowMultipleOk() (*bool, bool)`

GetAllowMultipleOk returns a tuple with the AllowMultiple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultiple

`func (o *PolicyConfiguration) SetAllowMultiple(v bool)`

SetAllowMultiple sets AllowMultiple field to given value.

### HasAllowMultiple

`func (o *PolicyConfiguration) HasAllowMultiple() bool`

HasAllowMultiple returns a boolean if a field has been set.

### GetConfiguration

`func (o *PolicyConfiguration) GetConfiguration() []PolicyConfigurationConfigurationInner`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *PolicyConfiguration) GetConfigurationOk() (*[]PolicyConfigurationConfigurationInner, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *PolicyConfiguration) SetConfiguration(v []PolicyConfigurationConfigurationInner)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *PolicyConfiguration) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


