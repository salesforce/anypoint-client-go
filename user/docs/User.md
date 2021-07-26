# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The user Id | 
**Username** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] [default to "Mule"]
**PhoneNumber** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**LastLogin** | Pointer to **string** |  | [optional] 
**MfaVerificationExcluded** | Pointer to **bool** |  | [optional] 
**MfaVerifiersConfigured** | Pointer to **string** |  | [optional] 
**IdproviderId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**IsFederated** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to [**Organization**](Organization.md) |  | [optional] 
**MemberOfOrganizations** | Pointer to [**[]Org**](Org.md) |  | [optional] 
**ContributorOfOrganizations** | Pointer to [**[]Org**](Org.md) |  | [optional] 
**Properties** | Pointer to [**Properties**](Properties.md) |  | [optional] 
**OrganizationPreferences** | Pointer to **map[string]interface{}** |  | [optional] 
**PrimaryOrganization** | Pointer to [**PrimaryOrganization**](PrimaryOrganization.md) |  | [optional] 

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

### GetPhoneNumber

`func (o *User) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *User) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *User) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *User) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

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

### GetOrganizationId

`func (o *User) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *User) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *User) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *User) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

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

### GetDeleted

`func (o *User) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *User) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *User) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *User) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetLastLogin

`func (o *User) GetLastLogin() string`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *User) GetLastLoginOk() (*string, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *User) SetLastLogin(v string)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *User) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### GetMfaVerificationExcluded

`func (o *User) GetMfaVerificationExcluded() bool`

GetMfaVerificationExcluded returns the MfaVerificationExcluded field if non-nil, zero value otherwise.

### GetMfaVerificationExcludedOk

`func (o *User) GetMfaVerificationExcludedOk() (*bool, bool)`

GetMfaVerificationExcludedOk returns a tuple with the MfaVerificationExcluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaVerificationExcluded

`func (o *User) SetMfaVerificationExcluded(v bool)`

SetMfaVerificationExcluded sets MfaVerificationExcluded field to given value.

### HasMfaVerificationExcluded

`func (o *User) HasMfaVerificationExcluded() bool`

HasMfaVerificationExcluded returns a boolean if a field has been set.

### GetMfaVerifiersConfigured

`func (o *User) GetMfaVerifiersConfigured() string`

GetMfaVerifiersConfigured returns the MfaVerifiersConfigured field if non-nil, zero value otherwise.

### GetMfaVerifiersConfiguredOk

`func (o *User) GetMfaVerifiersConfiguredOk() (*string, bool)`

GetMfaVerifiersConfiguredOk returns a tuple with the MfaVerifiersConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaVerifiersConfigured

`func (o *User) SetMfaVerifiersConfigured(v string)`

SetMfaVerifiersConfigured sets MfaVerifiersConfigured field to given value.

### HasMfaVerifiersConfigured

`func (o *User) HasMfaVerifiersConfigured() bool`

HasMfaVerifiersConfigured returns a boolean if a field has been set.

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

### GetIsFederated

`func (o *User) GetIsFederated() bool`

GetIsFederated returns the IsFederated field if non-nil, zero value otherwise.

### GetIsFederatedOk

`func (o *User) GetIsFederatedOk() (*bool, bool)`

GetIsFederatedOk returns a tuple with the IsFederated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFederated

`func (o *User) SetIsFederated(v bool)`

SetIsFederated sets IsFederated field to given value.

### HasIsFederated

`func (o *User) HasIsFederated() bool`

HasIsFederated returns a boolean if a field has been set.

### GetType

`func (o *User) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *User) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *User) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *User) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOrganization

`func (o *User) GetOrganization() Organization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *User) GetOrganizationOk() (*Organization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *User) SetOrganization(v Organization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *User) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetMemberOfOrganizations

`func (o *User) GetMemberOfOrganizations() []Org`

GetMemberOfOrganizations returns the MemberOfOrganizations field if non-nil, zero value otherwise.

### GetMemberOfOrganizationsOk

`func (o *User) GetMemberOfOrganizationsOk() (*[]Org, bool)`

GetMemberOfOrganizationsOk returns a tuple with the MemberOfOrganizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberOfOrganizations

`func (o *User) SetMemberOfOrganizations(v []Org)`

SetMemberOfOrganizations sets MemberOfOrganizations field to given value.

### HasMemberOfOrganizations

`func (o *User) HasMemberOfOrganizations() bool`

HasMemberOfOrganizations returns a boolean if a field has been set.

### GetContributorOfOrganizations

`func (o *User) GetContributorOfOrganizations() []Org`

GetContributorOfOrganizations returns the ContributorOfOrganizations field if non-nil, zero value otherwise.

### GetContributorOfOrganizationsOk

`func (o *User) GetContributorOfOrganizationsOk() (*[]Org, bool)`

GetContributorOfOrganizationsOk returns a tuple with the ContributorOfOrganizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributorOfOrganizations

`func (o *User) SetContributorOfOrganizations(v []Org)`

SetContributorOfOrganizations sets ContributorOfOrganizations field to given value.

### HasContributorOfOrganizations

`func (o *User) HasContributorOfOrganizations() bool`

HasContributorOfOrganizations returns a boolean if a field has been set.

### GetProperties

`func (o *User) GetProperties() Properties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *User) GetPropertiesOk() (*Properties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *User) SetProperties(v Properties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *User) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetOrganizationPreferences

`func (o *User) GetOrganizationPreferences() map[string]interface{}`

GetOrganizationPreferences returns the OrganizationPreferences field if non-nil, zero value otherwise.

### GetOrganizationPreferencesOk

`func (o *User) GetOrganizationPreferencesOk() (*map[string]interface{}, bool)`

GetOrganizationPreferencesOk returns a tuple with the OrganizationPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationPreferences

`func (o *User) SetOrganizationPreferences(v map[string]interface{})`

SetOrganizationPreferences sets OrganizationPreferences field to given value.

### HasOrganizationPreferences

`func (o *User) HasOrganizationPreferences() bool`

HasOrganizationPreferences returns a boolean if a field has been set.

### GetPrimaryOrganization

`func (o *User) GetPrimaryOrganization() PrimaryOrganization`

GetPrimaryOrganization returns the PrimaryOrganization field if non-nil, zero value otherwise.

### GetPrimaryOrganizationOk

`func (o *User) GetPrimaryOrganizationOk() (*PrimaryOrganization, bool)`

GetPrimaryOrganizationOk returns a tuple with the PrimaryOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryOrganization

`func (o *User) SetPrimaryOrganization(v PrimaryOrganization)`

SetPrimaryOrganization sets PrimaryOrganization field to given value.

### HasPrimaryOrganization

`func (o *User) HasPrimaryOrganization() bool`

HasPrimaryOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


