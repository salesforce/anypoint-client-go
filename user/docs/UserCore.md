# UserCore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**MemberOfOrganizations** | Pointer to [**[]Organization**](Organization.md) |  | [optional] 
**ContributorOfOrganizations** | Pointer to [**[]Organization**](Organization.md) |  | [optional] 
**Properties** | Pointer to [**Properties**](Properties.md) |  | [optional] 
**OrganizationPreferences** | Pointer to **map[string]interface{}** |  | [optional] 
**PrimaryOrganization** | Pointer to [**PrimaryOrganization**](PrimaryOrganization.md) |  | [optional] 

## Methods

### NewUserCore

`func NewUserCore() *UserCore`

NewUserCore instantiates a new UserCore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCoreWithDefaults

`func NewUserCoreWithDefaults() *UserCore`

NewUserCoreWithDefaults instantiates a new UserCore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *UserCore) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserCore) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserCore) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserCore) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetFirstName

`func (o *UserCore) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserCore) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserCore) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UserCore) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *UserCore) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserCore) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserCore) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UserCore) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *UserCore) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *UserCore) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *UserCore) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *UserCore) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetEmail

`func (o *UserCore) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserCore) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserCore) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserCore) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetOrganizationId

`func (o *UserCore) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *UserCore) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *UserCore) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *UserCore) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetEnabled

`func (o *UserCore) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UserCore) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UserCore) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UserCore) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDeleted

`func (o *UserCore) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *UserCore) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *UserCore) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *UserCore) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetLastLogin

`func (o *UserCore) GetLastLogin() string`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *UserCore) GetLastLoginOk() (*string, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *UserCore) SetLastLogin(v string)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *UserCore) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### GetMfaVerificationExcluded

`func (o *UserCore) GetMfaVerificationExcluded() bool`

GetMfaVerificationExcluded returns the MfaVerificationExcluded field if non-nil, zero value otherwise.

### GetMfaVerificationExcludedOk

`func (o *UserCore) GetMfaVerificationExcludedOk() (*bool, bool)`

GetMfaVerificationExcludedOk returns a tuple with the MfaVerificationExcluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaVerificationExcluded

`func (o *UserCore) SetMfaVerificationExcluded(v bool)`

SetMfaVerificationExcluded sets MfaVerificationExcluded field to given value.

### HasMfaVerificationExcluded

`func (o *UserCore) HasMfaVerificationExcluded() bool`

HasMfaVerificationExcluded returns a boolean if a field has been set.

### GetMfaVerifiersConfigured

`func (o *UserCore) GetMfaVerifiersConfigured() string`

GetMfaVerifiersConfigured returns the MfaVerifiersConfigured field if non-nil, zero value otherwise.

### GetMfaVerifiersConfiguredOk

`func (o *UserCore) GetMfaVerifiersConfiguredOk() (*string, bool)`

GetMfaVerifiersConfiguredOk returns a tuple with the MfaVerifiersConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaVerifiersConfigured

`func (o *UserCore) SetMfaVerifiersConfigured(v string)`

SetMfaVerifiersConfigured sets MfaVerifiersConfigured field to given value.

### HasMfaVerifiersConfigured

`func (o *UserCore) HasMfaVerifiersConfigured() bool`

HasMfaVerifiersConfigured returns a boolean if a field has been set.

### GetIdproviderId

`func (o *UserCore) GetIdproviderId() string`

GetIdproviderId returns the IdproviderId field if non-nil, zero value otherwise.

### GetIdproviderIdOk

`func (o *UserCore) GetIdproviderIdOk() (*string, bool)`

GetIdproviderIdOk returns a tuple with the IdproviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdproviderId

`func (o *UserCore) SetIdproviderId(v string)`

SetIdproviderId sets IdproviderId field to given value.

### HasIdproviderId

`func (o *UserCore) HasIdproviderId() bool`

HasIdproviderId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *UserCore) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserCore) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserCore) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UserCore) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *UserCore) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserCore) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserCore) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UserCore) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetIsFederated

`func (o *UserCore) GetIsFederated() bool`

GetIsFederated returns the IsFederated field if non-nil, zero value otherwise.

### GetIsFederatedOk

`func (o *UserCore) GetIsFederatedOk() (*bool, bool)`

GetIsFederatedOk returns a tuple with the IsFederated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFederated

`func (o *UserCore) SetIsFederated(v bool)`

SetIsFederated sets IsFederated field to given value.

### HasIsFederated

`func (o *UserCore) HasIsFederated() bool`

HasIsFederated returns a boolean if a field has been set.

### GetType

`func (o *UserCore) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserCore) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserCore) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserCore) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOrganization

`func (o *UserCore) GetOrganization() Organization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *UserCore) GetOrganizationOk() (*Organization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *UserCore) SetOrganization(v Organization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *UserCore) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetMemberOfOrganizations

`func (o *UserCore) GetMemberOfOrganizations() []Organization`

GetMemberOfOrganizations returns the MemberOfOrganizations field if non-nil, zero value otherwise.

### GetMemberOfOrganizationsOk

`func (o *UserCore) GetMemberOfOrganizationsOk() (*[]Organization, bool)`

GetMemberOfOrganizationsOk returns a tuple with the MemberOfOrganizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberOfOrganizations

`func (o *UserCore) SetMemberOfOrganizations(v []Organization)`

SetMemberOfOrganizations sets MemberOfOrganizations field to given value.

### HasMemberOfOrganizations

`func (o *UserCore) HasMemberOfOrganizations() bool`

HasMemberOfOrganizations returns a boolean if a field has been set.

### GetContributorOfOrganizations

`func (o *UserCore) GetContributorOfOrganizations() []Organization`

GetContributorOfOrganizations returns the ContributorOfOrganizations field if non-nil, zero value otherwise.

### GetContributorOfOrganizationsOk

`func (o *UserCore) GetContributorOfOrganizationsOk() (*[]Organization, bool)`

GetContributorOfOrganizationsOk returns a tuple with the ContributorOfOrganizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributorOfOrganizations

`func (o *UserCore) SetContributorOfOrganizations(v []Organization)`

SetContributorOfOrganizations sets ContributorOfOrganizations field to given value.

### HasContributorOfOrganizations

`func (o *UserCore) HasContributorOfOrganizations() bool`

HasContributorOfOrganizations returns a boolean if a field has been set.

### GetProperties

`func (o *UserCore) GetProperties() Properties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *UserCore) GetPropertiesOk() (*Properties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *UserCore) SetProperties(v Properties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *UserCore) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetOrganizationPreferences

`func (o *UserCore) GetOrganizationPreferences() map[string]interface{}`

GetOrganizationPreferences returns the OrganizationPreferences field if non-nil, zero value otherwise.

### GetOrganizationPreferencesOk

`func (o *UserCore) GetOrganizationPreferencesOk() (*map[string]interface{}, bool)`

GetOrganizationPreferencesOk returns a tuple with the OrganizationPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationPreferences

`func (o *UserCore) SetOrganizationPreferences(v map[string]interface{})`

SetOrganizationPreferences sets OrganizationPreferences field to given value.

### HasOrganizationPreferences

`func (o *UserCore) HasOrganizationPreferences() bool`

HasOrganizationPreferences returns a boolean if a field has been set.

### GetPrimaryOrganization

`func (o *UserCore) GetPrimaryOrganization() PrimaryOrganization`

GetPrimaryOrganization returns the PrimaryOrganization field if non-nil, zero value otherwise.

### GetPrimaryOrganizationOk

`func (o *UserCore) GetPrimaryOrganizationOk() (*PrimaryOrganization, bool)`

GetPrimaryOrganizationOk returns a tuple with the PrimaryOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryOrganization

`func (o *UserCore) SetPrimaryOrganization(v PrimaryOrganization)`

SetPrimaryOrganization sets PrimaryOrganization field to given value.

### HasPrimaryOrganization

`func (o *UserCore) HasPrimaryOrganization() bool`

HasPrimaryOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


