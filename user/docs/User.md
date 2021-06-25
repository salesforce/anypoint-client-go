# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The user Id | 
**Username** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] [default to "Mule"]
**Email** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**IdproviderId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**IsApiConsumer** | Pointer to **bool** |  | [optional] 

## Methods

### NewUser

`func NewUser(id string, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *User) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v string)`

SetId sets Id field to given value.


### GetUsername

`func (o *User) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *User) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *User) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *User) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetFirstName

`func (o *User) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *User) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *User) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *User) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *User) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *User) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *User) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *User) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *User) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEnabled

`func (o *User) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *User) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *User) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *User) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIdproviderId

`func (o *User) GetIdproviderId() string`

GetIdproviderId returns the IdproviderId field if non-nil, zero value otherwise.

### GetIdproviderIdOk

`func (o *User) GetIdproviderIdOk() (*string, bool)`

GetIdproviderIdOk returns a tuple with the IdproviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdproviderId

`func (o *User) SetIdproviderId(v string)`

SetIdproviderId sets IdproviderId field to given value.

### HasIdproviderId

`func (o *User) HasIdproviderId() bool`

HasIdproviderId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *User) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *User) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *User) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *User) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *User) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *User) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *User) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *User) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetIsApiConsumer

`func (o *User) GetIsApiConsumer() bool`

GetIsApiConsumer returns the IsApiConsumer field if non-nil, zero value otherwise.

### GetIsApiConsumerOk

`func (o *User) GetIsApiConsumerOk() (*bool, bool)`

GetIsApiConsumerOk returns a tuple with the IsApiConsumer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApiConsumer

`func (o *User) SetIsApiConsumer(v bool)`

SetIsApiConsumer sets IsApiConsumer field to given value.

### HasIsApiConsumer

`func (o *User) HasIsApiConsumer() bool`

HasIsApiConsumer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


