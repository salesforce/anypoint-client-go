# ClaimsMapping2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsernameAttribute** | Pointer to **string** | Field name in the SAML AttributeStatements that maps to username. By default, the NameID attribute in the SAML assertion is used. | [optional] 
**FirstnameAttribute** | Pointer to **string** | Field name in the SAML AttributeStatements that maps to First Name. By default, the firstname attribute in the SAML assertion is used. | [optional] 
**LastnameAttribute** | Pointer to **string** | Field name in the SAML AttributeStatements that maps to Last Name. By default, the lastname attribute in the SAML assertion is used. | [optional] 
**EmailAttribute** | Pointer to **string** | Field name in the SAML AttributeStatements that maps to Email. By default, the email attribute in the SAML assertion is used. | [optional] 
**GroupAttribute** | Pointer to **string** | Field name in the SAML AttributeStatements that maps to Group. | [optional] 

## Methods

### NewClaimsMapping2

`func NewClaimsMapping2() *ClaimsMapping2`

NewClaimsMapping2 instantiates a new ClaimsMapping2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClaimsMapping2WithDefaults

`func NewClaimsMapping2WithDefaults() *ClaimsMapping2`

NewClaimsMapping2WithDefaults instantiates a new ClaimsMapping2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsernameAttribute

`func (o *ClaimsMapping2) GetUsernameAttribute() string`

GetUsernameAttribute returns the UsernameAttribute field if non-nil, zero value otherwise.

### GetUsernameAttributeOk

`func (o *ClaimsMapping2) GetUsernameAttributeOk() (*string, bool)`

GetUsernameAttributeOk returns a tuple with the UsernameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameAttribute

`func (o *ClaimsMapping2) SetUsernameAttribute(v string)`

SetUsernameAttribute sets UsernameAttribute field to given value.

### HasUsernameAttribute

`func (o *ClaimsMapping2) HasUsernameAttribute() bool`

HasUsernameAttribute returns a boolean if a field has been set.

### GetFirstnameAttribute

`func (o *ClaimsMapping2) GetFirstnameAttribute() string`

GetFirstnameAttribute returns the FirstnameAttribute field if non-nil, zero value otherwise.

### GetFirstnameAttributeOk

`func (o *ClaimsMapping2) GetFirstnameAttributeOk() (*string, bool)`

GetFirstnameAttributeOk returns a tuple with the FirstnameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstnameAttribute

`func (o *ClaimsMapping2) SetFirstnameAttribute(v string)`

SetFirstnameAttribute sets FirstnameAttribute field to given value.

### HasFirstnameAttribute

`func (o *ClaimsMapping2) HasFirstnameAttribute() bool`

HasFirstnameAttribute returns a boolean if a field has been set.

### GetLastnameAttribute

`func (o *ClaimsMapping2) GetLastnameAttribute() string`

GetLastnameAttribute returns the LastnameAttribute field if non-nil, zero value otherwise.

### GetLastnameAttributeOk

`func (o *ClaimsMapping2) GetLastnameAttributeOk() (*string, bool)`

GetLastnameAttributeOk returns a tuple with the LastnameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastnameAttribute

`func (o *ClaimsMapping2) SetLastnameAttribute(v string)`

SetLastnameAttribute sets LastnameAttribute field to given value.

### HasLastnameAttribute

`func (o *ClaimsMapping2) HasLastnameAttribute() bool`

HasLastnameAttribute returns a boolean if a field has been set.

### GetEmailAttribute

`func (o *ClaimsMapping2) GetEmailAttribute() string`

GetEmailAttribute returns the EmailAttribute field if non-nil, zero value otherwise.

### GetEmailAttributeOk

`func (o *ClaimsMapping2) GetEmailAttributeOk() (*string, bool)`

GetEmailAttributeOk returns a tuple with the EmailAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAttribute

`func (o *ClaimsMapping2) SetEmailAttribute(v string)`

SetEmailAttribute sets EmailAttribute field to given value.

### HasEmailAttribute

`func (o *ClaimsMapping2) HasEmailAttribute() bool`

HasEmailAttribute returns a boolean if a field has been set.

### GetGroupAttribute

`func (o *ClaimsMapping2) GetGroupAttribute() string`

GetGroupAttribute returns the GroupAttribute field if non-nil, zero value otherwise.

### GetGroupAttributeOk

`func (o *ClaimsMapping2) GetGroupAttributeOk() (*string, bool)`

GetGroupAttributeOk returns a tuple with the GroupAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupAttribute

`func (o *ClaimsMapping2) SetGroupAttribute(v string)`

SetGroupAttribute sets GroupAttribute field to given value.

### HasGroupAttribute

`func (o *ClaimsMapping2) HasGroupAttribute() bool`

HasGroupAttribute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


