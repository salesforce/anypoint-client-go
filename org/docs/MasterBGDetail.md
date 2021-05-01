# MasterBGDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | An explanation about the purpose of this instance. | [optional] [default to ""]
**CreatedAt** | Pointer to **string** | An explanation about the purpose of this instance. | [optional] [default to ""]
**Domain** | Pointer to **string** | An explanation about the purpose of this instance. | [optional] [default to ""]
**Entitlements** | Pointer to [**Entitlements**](Entitlements.md) |  | [optional] [default to {}]
**Environments** | Pointer to [**[]Environment**](Environment.md) | An explanation about the purpose of this instance. | [optional] [default to []]
**Id** | Pointer to **string** | An explanation about the purpose of this instance. | [optional] [default to ""]
**IdproviderId** | Pointer to **string** | An explanation about the purpose of this instance. | [optional] [default to ""]
**IsAutomaticAdminPromotionExempt** | Pointer to **bool** | An explanation about the purpose of this instance. | [optional] [default to false]
**IsFederated** | Pointer to **bool** | An explanation about the purpose of this instance. | [optional] [default to false]
**IsMaster** | Pointer to **bool** | An explanation about the purpose of this instance. | [optional] [default to false]
**MfaRequired** | Pointer to **string** | An explanation about the purpose of this instance. | [optional] [default to ""]
**Name** | Pointer to **string** | An explanation about the purpose of this instance. | [optional] [default to ""]
**OwnerId** | Pointer to **string** | An explanation about the purpose of this instance. | [optional] [default to ""]
**ParentOrganizationIds** | Pointer to **[]string** | An explanation about the purpose of this instance. | [optional] [default to []]
**Properties** | Pointer to **map[string]interface{}** | An explanation about the purpose of this instance. | [optional] [default to {}]
**SubOrganizationIds** | Pointer to **[]string** | An explanation about the purpose of this instance. | [optional] [default to []]
**TenantOrganizationIds** | Pointer to **[]string** | An explanation about the purpose of this instance. | [optional] [default to []]
**UpdatedAt** | Pointer to **string** | An explanation about the purpose of this instance. | [optional] [default to ""]
**Owner** | Pointer to [**Owner**](Owner.md) |  | [optional] 
**SessionTimeout** | Pointer to **int32** | An explanation about the purpose of this instance. | [optional] [default to 0]
**Subscription** | Pointer to [**Subscription**](Subscription.md) |  | [optional] 

## Methods

### NewMasterBGDetail

`func NewMasterBGDetail() *MasterBGDetail`

NewMasterBGDetail instantiates a new MasterBGDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMasterBGDetailWithDefaults

`func NewMasterBGDetailWithDefaults() *MasterBGDetail`

NewMasterBGDetailWithDefaults instantiates a new MasterBGDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *MasterBGDetail) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *MasterBGDetail) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *MasterBGDetail) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *MasterBGDetail) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MasterBGDetail) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MasterBGDetail) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MasterBGDetail) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MasterBGDetail) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDomain

`func (o *MasterBGDetail) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *MasterBGDetail) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *MasterBGDetail) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *MasterBGDetail) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetEntitlements

`func (o *MasterBGDetail) GetEntitlements() Entitlements`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *MasterBGDetail) GetEntitlementsOk() (*Entitlements, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *MasterBGDetail) SetEntitlements(v Entitlements)`

SetEntitlements sets Entitlements field to given value.

### HasEntitlements

`func (o *MasterBGDetail) HasEntitlements() bool`

HasEntitlements returns a boolean if a field has been set.

### GetEnvironments

`func (o *MasterBGDetail) GetEnvironments() []Environment`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *MasterBGDetail) GetEnvironmentsOk() (*[]Environment, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *MasterBGDetail) SetEnvironments(v []Environment)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *MasterBGDetail) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetId

`func (o *MasterBGDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MasterBGDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MasterBGDetail) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MasterBGDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdproviderId

`func (o *MasterBGDetail) GetIdproviderId() string`

GetIdproviderId returns the IdproviderId field if non-nil, zero value otherwise.

### GetIdproviderIdOk

`func (o *MasterBGDetail) GetIdproviderIdOk() (*string, bool)`

GetIdproviderIdOk returns a tuple with the IdproviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdproviderId

`func (o *MasterBGDetail) SetIdproviderId(v string)`

SetIdproviderId sets IdproviderId field to given value.

### HasIdproviderId

`func (o *MasterBGDetail) HasIdproviderId() bool`

HasIdproviderId returns a boolean if a field has been set.

### GetIsAutomaticAdminPromotionExempt

`func (o *MasterBGDetail) GetIsAutomaticAdminPromotionExempt() bool`

GetIsAutomaticAdminPromotionExempt returns the IsAutomaticAdminPromotionExempt field if non-nil, zero value otherwise.

### GetIsAutomaticAdminPromotionExemptOk

`func (o *MasterBGDetail) GetIsAutomaticAdminPromotionExemptOk() (*bool, bool)`

GetIsAutomaticAdminPromotionExemptOk returns a tuple with the IsAutomaticAdminPromotionExempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomaticAdminPromotionExempt

`func (o *MasterBGDetail) SetIsAutomaticAdminPromotionExempt(v bool)`

SetIsAutomaticAdminPromotionExempt sets IsAutomaticAdminPromotionExempt field to given value.

### HasIsAutomaticAdminPromotionExempt

`func (o *MasterBGDetail) HasIsAutomaticAdminPromotionExempt() bool`

HasIsAutomaticAdminPromotionExempt returns a boolean if a field has been set.

### GetIsFederated

`func (o *MasterBGDetail) GetIsFederated() bool`

GetIsFederated returns the IsFederated field if non-nil, zero value otherwise.

### GetIsFederatedOk

`func (o *MasterBGDetail) GetIsFederatedOk() (*bool, bool)`

GetIsFederatedOk returns a tuple with the IsFederated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFederated

`func (o *MasterBGDetail) SetIsFederated(v bool)`

SetIsFederated sets IsFederated field to given value.

### HasIsFederated

`func (o *MasterBGDetail) HasIsFederated() bool`

HasIsFederated returns a boolean if a field has been set.

### GetIsMaster

`func (o *MasterBGDetail) GetIsMaster() bool`

GetIsMaster returns the IsMaster field if non-nil, zero value otherwise.

### GetIsMasterOk

`func (o *MasterBGDetail) GetIsMasterOk() (*bool, bool)`

GetIsMasterOk returns a tuple with the IsMaster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMaster

`func (o *MasterBGDetail) SetIsMaster(v bool)`

SetIsMaster sets IsMaster field to given value.

### HasIsMaster

`func (o *MasterBGDetail) HasIsMaster() bool`

HasIsMaster returns a boolean if a field has been set.

### GetMfaRequired

`func (o *MasterBGDetail) GetMfaRequired() string`

GetMfaRequired returns the MfaRequired field if non-nil, zero value otherwise.

### GetMfaRequiredOk

`func (o *MasterBGDetail) GetMfaRequiredOk() (*string, bool)`

GetMfaRequiredOk returns a tuple with the MfaRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaRequired

`func (o *MasterBGDetail) SetMfaRequired(v string)`

SetMfaRequired sets MfaRequired field to given value.

### HasMfaRequired

`func (o *MasterBGDetail) HasMfaRequired() bool`

HasMfaRequired returns a boolean if a field has been set.

### GetName

`func (o *MasterBGDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MasterBGDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MasterBGDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MasterBGDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwnerId

`func (o *MasterBGDetail) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *MasterBGDetail) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *MasterBGDetail) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *MasterBGDetail) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetParentOrganizationIds

`func (o *MasterBGDetail) GetParentOrganizationIds() []string`

GetParentOrganizationIds returns the ParentOrganizationIds field if non-nil, zero value otherwise.

### GetParentOrganizationIdsOk

`func (o *MasterBGDetail) GetParentOrganizationIdsOk() (*[]string, bool)`

GetParentOrganizationIdsOk returns a tuple with the ParentOrganizationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentOrganizationIds

`func (o *MasterBGDetail) SetParentOrganizationIds(v []string)`

SetParentOrganizationIds sets ParentOrganizationIds field to given value.

### HasParentOrganizationIds

`func (o *MasterBGDetail) HasParentOrganizationIds() bool`

HasParentOrganizationIds returns a boolean if a field has been set.

### GetProperties

`func (o *MasterBGDetail) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *MasterBGDetail) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *MasterBGDetail) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *MasterBGDetail) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetSubOrganizationIds

`func (o *MasterBGDetail) GetSubOrganizationIds() []string`

GetSubOrganizationIds returns the SubOrganizationIds field if non-nil, zero value otherwise.

### GetSubOrganizationIdsOk

`func (o *MasterBGDetail) GetSubOrganizationIdsOk() (*[]string, bool)`

GetSubOrganizationIdsOk returns a tuple with the SubOrganizationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubOrganizationIds

`func (o *MasterBGDetail) SetSubOrganizationIds(v []string)`

SetSubOrganizationIds sets SubOrganizationIds field to given value.

### HasSubOrganizationIds

`func (o *MasterBGDetail) HasSubOrganizationIds() bool`

HasSubOrganizationIds returns a boolean if a field has been set.

### GetTenantOrganizationIds

`func (o *MasterBGDetail) GetTenantOrganizationIds() []string`

GetTenantOrganizationIds returns the TenantOrganizationIds field if non-nil, zero value otherwise.

### GetTenantOrganizationIdsOk

`func (o *MasterBGDetail) GetTenantOrganizationIdsOk() (*[]string, bool)`

GetTenantOrganizationIdsOk returns a tuple with the TenantOrganizationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantOrganizationIds

`func (o *MasterBGDetail) SetTenantOrganizationIds(v []string)`

SetTenantOrganizationIds sets TenantOrganizationIds field to given value.

### HasTenantOrganizationIds

`func (o *MasterBGDetail) HasTenantOrganizationIds() bool`

HasTenantOrganizationIds returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MasterBGDetail) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MasterBGDetail) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MasterBGDetail) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MasterBGDetail) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetOwner

`func (o *MasterBGDetail) GetOwner() Owner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *MasterBGDetail) GetOwnerOk() (*Owner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *MasterBGDetail) SetOwner(v Owner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *MasterBGDetail) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetSessionTimeout

`func (o *MasterBGDetail) GetSessionTimeout() int32`

GetSessionTimeout returns the SessionTimeout field if non-nil, zero value otherwise.

### GetSessionTimeoutOk

`func (o *MasterBGDetail) GetSessionTimeoutOk() (*int32, bool)`

GetSessionTimeoutOk returns a tuple with the SessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTimeout

`func (o *MasterBGDetail) SetSessionTimeout(v int32)`

SetSessionTimeout sets SessionTimeout field to given value.

### HasSessionTimeout

`func (o *MasterBGDetail) HasSessionTimeout() bool`

HasSessionTimeout returns a boolean if a field has been set.

### GetSubscription

`func (o *MasterBGDetail) GetSubscription() Subscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *MasterBGDetail) GetSubscriptionOk() (*Subscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *MasterBGDetail) SetSubscription(v Subscription)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *MasterBGDetail) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


