# Owner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | An explanation about the purpose of this instance. | [optional] [default to ""]
**Deleted** | Pointer to **bool** | An explanation about the purpose of this instance. | [optional] [default to false]
**Email** | Pointer to **string** | An explanation about the purpose of this instance. | [optional] [default to ""]
**Enabled** | Pointer to **bool** | An explanation about the purpose of this instance. | [optional] [default to false]
**FirstName** | Pointer to **string** | An explanation about the purpose of this instance. | [optional] [default to ""]
**Id** | Pointer to **string** | An explanation about the purpose of this instance. | [optional] [default to ""]
**IdproviderId** | Pointer to **string** | An explanation about the purpose of this instance. | [optional] [default to ""]
**LastLogin** | Pointer to **string** | An explanation about the purpose of this instance. | [optional] [default to ""]
**LastName** | Pointer to **string** | An explanation about the purpose of this instance. | [optional] [default to ""]
**MfaVerificationExcluded** | Pointer to **bool** | An explanation about the purpose of this instance. | [optional] [default to false]
**MfaVerifiersConfigured** | Pointer to **string** | An explanation about the purpose of this instance. | [optional] [default to ""]
**OrganizationId** | Pointer to **string** | An explanation about the purpose of this instance. | [optional] [default to ""]
**PhoneNumber** | Pointer to **string** | An explanation about the purpose of this instance. | [optional] [default to ""]
**Type** | Pointer to **string** | An explanation about the purpose of this instance. | [optional] [default to ""]
**UpdatedAt** | Pointer to **string** | An explanation about the purpose of this instance. | [optional] [default to ""]
**Username** | Pointer to **string** | An explanation about the purpose of this instance. | [optional] [default to ""]

## Methods

### NewOwner

`func NewOwner() *Owner`

NewOwner instantiates a new Owner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOwnerWithDefaults

`func NewOwnerWithDefaults() *Owner`

NewOwnerWithDefaults instantiates a new Owner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Owner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Owner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Owner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Owner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeleted

`func (o *Owner) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Owner) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Owner) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Owner) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetEmail

`func (o *Owner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Owner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Owner) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Owner) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEnabled

`func (o *Owner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Owner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Owner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Owner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFirstName

`func (o *Owner) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Owner) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Owner) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *Owner) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetId

`func (o *Owner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Owner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Owner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Owner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdproviderId

`func (o *Owner) GetIdproviderId() string`

GetIdproviderId returns the IdproviderId field if non-nil, zero value otherwise.

### GetIdproviderIdOk

`func (o *Owner) GetIdproviderIdOk() (*string, bool)`

GetIdproviderIdOk returns a tuple with the IdproviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdproviderId

`func (o *Owner) SetIdproviderId(v string)`

SetIdproviderId sets IdproviderId field to given value.

### HasIdproviderId

`func (o *Owner) HasIdproviderId() bool`

HasIdproviderId returns a boolean if a field has been set.

### GetLastLogin

`func (o *Owner) GetLastLogin() string`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *Owner) GetLastLoginOk() (*string, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *Owner) SetLastLogin(v string)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *Owner) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### GetLastName

`func (o *Owner) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Owner) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Owner) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *Owner) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetMfaVerificationExcluded

`func (o *Owner) GetMfaVerificationExcluded() bool`

GetMfaVerificationExcluded returns the MfaVerificationExcluded field if non-nil, zero value otherwise.

### GetMfaVerificationExcludedOk

`func (o *Owner) GetMfaVerificationExcludedOk() (*bool, bool)`

GetMfaVerificationExcludedOk returns a tuple with the MfaVerificationExcluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaVerificationExcluded

`func (o *Owner) SetMfaVerificationExcluded(v bool)`

SetMfaVerificationExcluded sets MfaVerificationExcluded field to given value.

### HasMfaVerificationExcluded

`func (o *Owner) HasMfaVerificationExcluded() bool`

HasMfaVerificationExcluded returns a boolean if a field has been set.

### GetMfaVerifiersConfigured

`func (o *Owner) GetMfaVerifiersConfigured() string`

GetMfaVerifiersConfigured returns the MfaVerifiersConfigured field if non-nil, zero value otherwise.

### GetMfaVerifiersConfiguredOk

`func (o *Owner) GetMfaVerifiersConfiguredOk() (*string, bool)`

GetMfaVerifiersConfiguredOk returns a tuple with the MfaVerifiersConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaVerifiersConfigured

`func (o *Owner) SetMfaVerifiersConfigured(v string)`

SetMfaVerifiersConfigured sets MfaVerifiersConfigured field to given value.

### HasMfaVerifiersConfigured

`func (o *Owner) HasMfaVerifiersConfigured() bool`

HasMfaVerifiersConfigured returns a boolean if a field has been set.

### GetOrganizationId

`func (o *Owner) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Owner) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *Owner) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *Owner) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *Owner) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *Owner) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *Owner) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *Owner) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetType

`func (o *Owner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Owner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Owner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Owner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Owner) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Owner) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Owner) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Owner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUsername

`func (o *Owner) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Owner) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Owner) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Owner) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


