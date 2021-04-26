# BGDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | An explanation about the purpose of this instance. | [default to ""]
**CreatedAt** | **string** | An explanation about the purpose of this instance. | [default to ""]
**Domain** | **string** | An explanation about the purpose of this instance. | [default to ""]
**Entitlements** | **map[string]map[string]interface{}** | An explanation about the purpose of this instance. | [default to {}]
**Environments** | [**[]AnyOfmap**](AnyOfmap.md) | An explanation about the purpose of this instance. | [default to []]
**Id** | **string** | An explanation about the purpose of this instance. | [default to ""]
**IdproviderId** | **string** | An explanation about the purpose of this instance. | [default to ""]
**IsAutomaticAdminPromotionExempt** | **bool** | An explanation about the purpose of this instance. | [default to false]
**IsFederated** | **bool** | An explanation about the purpose of this instance. | [default to false]
**IsMaster** | **bool** | An explanation about the purpose of this instance. | [default to false]
**MfaRequired** | **string** | An explanation about the purpose of this instance. | [default to ""]
**Name** | **string** | An explanation about the purpose of this instance. | [default to ""]
**OwnerId** | **string** | An explanation about the purpose of this instance. | [default to ""]
**ParentOrganizationIds** | **[]string** | An explanation about the purpose of this instance. | [default to []]
**Properties** | **map[string]map[string]interface{}** | An explanation about the purpose of this instance. | [default to {}]
**SubOrganizationIds** | [**[]AnyOfstring**](AnyOfstring.md) | An explanation about the purpose of this instance. | [default to []]
**TenantOrganizationIds** | **[]string** | An explanation about the purpose of this instance. | [default to []]
**UpdatedAt** | **string** | An explanation about the purpose of this instance. | [default to ""]
**Deleted** | **bool** | An explanation about the purpose of this instance. | [default to false]
**Email** | **string** | An explanation about the purpose of this instance. | [default to ""]
**Enabled** | **bool** | An explanation about the purpose of this instance. | [default to false]
**FirstName** | **string** | An explanation about the purpose of this instance. | [default to ""]
**LastLogin** | **string** | An explanation about the purpose of this instance. | [default to ""]
**LastName** | **string** | An explanation about the purpose of this instance. | [default to ""]
**MfaVerificationExcluded** | **bool** | An explanation about the purpose of this instance. | [default to false]
**MfaVerifiersConfigured** | **string** | An explanation about the purpose of this instance. | [default to ""]
**OrganizationId** | **string** | An explanation about the purpose of this instance. | [default to ""]
**PhoneNumber** | **string** | An explanation about the purpose of this instance. | [default to ""]
**Type** | **string** | An explanation about the purpose of this instance. | [default to ""]
**Username** | **string** | An explanation about the purpose of this instance. | [default to ""]

## Methods

### NewBGDetail

`func NewBGDetail(clientId string, createdAt string, domain string, entitlements map[string]map[string]interface{}, environments []AnyOfmap, id string, idproviderId string, isAutomaticAdminPromotionExempt bool, isFederated bool, isMaster bool, mfaRequired string, name string, ownerId string, parentOrganizationIds []string, properties map[string]map[string]interface{}, subOrganizationIds []AnyOfstring, tenantOrganizationIds []string, updatedAt string, deleted bool, email string, enabled bool, firstName string, lastLogin string, lastName string, mfaVerificationExcluded bool, mfaVerifiersConfigured string, organizationId string, phoneNumber string, type_ string, username string, ) *BGDetail`

NewBGDetail instantiates a new BGDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBGDetailWithDefaults

`func NewBGDetailWithDefaults() *BGDetail`

NewBGDetailWithDefaults instantiates a new BGDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *BGDetail) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *BGDetail) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *BGDetail) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetCreatedAt

`func (o *BGDetail) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BGDetail) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BGDetail) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDomain

`func (o *BGDetail) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *BGDetail) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *BGDetail) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetEntitlements

`func (o *BGDetail) GetEntitlements() map[string]map[string]interface{}`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *BGDetail) GetEntitlementsOk() (*map[string]map[string]interface{}, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *BGDetail) SetEntitlements(v map[string]map[string]interface{})`

SetEntitlements sets Entitlements field to given value.


### GetEnvironments

`func (o *BGDetail) GetEnvironments() []AnyOfmap`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *BGDetail) GetEnvironmentsOk() (*[]AnyOfmap, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *BGDetail) SetEnvironments(v []AnyOfmap)`

SetEnvironments sets Environments field to given value.


### GetId

`func (o *BGDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BGDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BGDetail) SetId(v string)`

SetId sets Id field to given value.


### GetIdproviderId

`func (o *BGDetail) GetIdproviderId() string`

GetIdproviderId returns the IdproviderId field if non-nil, zero value otherwise.

### GetIdproviderIdOk

`func (o *BGDetail) GetIdproviderIdOk() (*string, bool)`

GetIdproviderIdOk returns a tuple with the IdproviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdproviderId

`func (o *BGDetail) SetIdproviderId(v string)`

SetIdproviderId sets IdproviderId field to given value.


### GetIsAutomaticAdminPromotionExempt

`func (o *BGDetail) GetIsAutomaticAdminPromotionExempt() bool`

GetIsAutomaticAdminPromotionExempt returns the IsAutomaticAdminPromotionExempt field if non-nil, zero value otherwise.

### GetIsAutomaticAdminPromotionExemptOk

`func (o *BGDetail) GetIsAutomaticAdminPromotionExemptOk() (*bool, bool)`

GetIsAutomaticAdminPromotionExemptOk returns a tuple with the IsAutomaticAdminPromotionExempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomaticAdminPromotionExempt

`func (o *BGDetail) SetIsAutomaticAdminPromotionExempt(v bool)`

SetIsAutomaticAdminPromotionExempt sets IsAutomaticAdminPromotionExempt field to given value.


### GetIsFederated

`func (o *BGDetail) GetIsFederated() bool`

GetIsFederated returns the IsFederated field if non-nil, zero value otherwise.

### GetIsFederatedOk

`func (o *BGDetail) GetIsFederatedOk() (*bool, bool)`

GetIsFederatedOk returns a tuple with the IsFederated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFederated

`func (o *BGDetail) SetIsFederated(v bool)`

SetIsFederated sets IsFederated field to given value.


### GetIsMaster

`func (o *BGDetail) GetIsMaster() bool`

GetIsMaster returns the IsMaster field if non-nil, zero value otherwise.

### GetIsMasterOk

`func (o *BGDetail) GetIsMasterOk() (*bool, bool)`

GetIsMasterOk returns a tuple with the IsMaster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMaster

`func (o *BGDetail) SetIsMaster(v bool)`

SetIsMaster sets IsMaster field to given value.


### GetMfaRequired

`func (o *BGDetail) GetMfaRequired() string`

GetMfaRequired returns the MfaRequired field if non-nil, zero value otherwise.

### GetMfaRequiredOk

`func (o *BGDetail) GetMfaRequiredOk() (*string, bool)`

GetMfaRequiredOk returns a tuple with the MfaRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaRequired

`func (o *BGDetail) SetMfaRequired(v string)`

SetMfaRequired sets MfaRequired field to given value.


### GetName

`func (o *BGDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BGDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BGDetail) SetName(v string)`

SetName sets Name field to given value.


### GetOwnerId

`func (o *BGDetail) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *BGDetail) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *BGDetail) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetParentOrganizationIds

`func (o *BGDetail) GetParentOrganizationIds() []string`

GetParentOrganizationIds returns the ParentOrganizationIds field if non-nil, zero value otherwise.

### GetParentOrganizationIdsOk

`func (o *BGDetail) GetParentOrganizationIdsOk() (*[]string, bool)`

GetParentOrganizationIdsOk returns a tuple with the ParentOrganizationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentOrganizationIds

`func (o *BGDetail) SetParentOrganizationIds(v []string)`

SetParentOrganizationIds sets ParentOrganizationIds field to given value.


### GetProperties

`func (o *BGDetail) GetProperties() map[string]map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *BGDetail) GetPropertiesOk() (*map[string]map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *BGDetail) SetProperties(v map[string]map[string]interface{})`

SetProperties sets Properties field to given value.


### GetSubOrganizationIds

`func (o *BGDetail) GetSubOrganizationIds() []AnyOfstring`

GetSubOrganizationIds returns the SubOrganizationIds field if non-nil, zero value otherwise.

### GetSubOrganizationIdsOk

`func (o *BGDetail) GetSubOrganizationIdsOk() (*[]AnyOfstring, bool)`

GetSubOrganizationIdsOk returns a tuple with the SubOrganizationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubOrganizationIds

`func (o *BGDetail) SetSubOrganizationIds(v []AnyOfstring)`

SetSubOrganizationIds sets SubOrganizationIds field to given value.


### GetTenantOrganizationIds

`func (o *BGDetail) GetTenantOrganizationIds() []string`

GetTenantOrganizationIds returns the TenantOrganizationIds field if non-nil, zero value otherwise.

### GetTenantOrganizationIdsOk

`func (o *BGDetail) GetTenantOrganizationIdsOk() (*[]string, bool)`

GetTenantOrganizationIdsOk returns a tuple with the TenantOrganizationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantOrganizationIds

`func (o *BGDetail) SetTenantOrganizationIds(v []string)`

SetTenantOrganizationIds sets TenantOrganizationIds field to given value.


### GetUpdatedAt

`func (o *BGDetail) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BGDetail) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BGDetail) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeleted

`func (o *BGDetail) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *BGDetail) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *BGDetail) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.


### GetEmail

`func (o *BGDetail) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BGDetail) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BGDetail) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetEnabled

`func (o *BGDetail) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BGDetail) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BGDetail) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetFirstName

`func (o *BGDetail) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *BGDetail) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *BGDetail) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastLogin

`func (o *BGDetail) GetLastLogin() string`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *BGDetail) GetLastLoginOk() (*string, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *BGDetail) SetLastLogin(v string)`

SetLastLogin sets LastLogin field to given value.


### GetLastName

`func (o *BGDetail) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *BGDetail) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *BGDetail) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetMfaVerificationExcluded

`func (o *BGDetail) GetMfaVerificationExcluded() bool`

GetMfaVerificationExcluded returns the MfaVerificationExcluded field if non-nil, zero value otherwise.

### GetMfaVerificationExcludedOk

`func (o *BGDetail) GetMfaVerificationExcludedOk() (*bool, bool)`

GetMfaVerificationExcludedOk returns a tuple with the MfaVerificationExcluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaVerificationExcluded

`func (o *BGDetail) SetMfaVerificationExcluded(v bool)`

SetMfaVerificationExcluded sets MfaVerificationExcluded field to given value.


### GetMfaVerifiersConfigured

`func (o *BGDetail) GetMfaVerifiersConfigured() string`

GetMfaVerifiersConfigured returns the MfaVerifiersConfigured field if non-nil, zero value otherwise.

### GetMfaVerifiersConfiguredOk

`func (o *BGDetail) GetMfaVerifiersConfiguredOk() (*string, bool)`

GetMfaVerifiersConfiguredOk returns a tuple with the MfaVerifiersConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaVerifiersConfigured

`func (o *BGDetail) SetMfaVerifiersConfigured(v string)`

SetMfaVerifiersConfigured sets MfaVerifiersConfigured field to given value.


### GetOrganizationId

`func (o *BGDetail) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *BGDetail) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *BGDetail) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetPhoneNumber

`func (o *BGDetail) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *BGDetail) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *BGDetail) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.


### GetType

`func (o *BGDetail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BGDetail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BGDetail) SetType(v string)`

SetType sets Type field to given value.


### GetUsername

`func (o *BGDetail) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *BGDetail) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *BGDetail) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


