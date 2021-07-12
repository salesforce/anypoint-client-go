# Organization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**IdproviderId** | Pointer to **string** |  | [optional] 
**IsFederated** | Pointer to **bool** |  | [optional] 
**ParentOrganizationIds** | Pointer to **[]string** |  | [optional] 
**SubOrganizationIds** | Pointer to **[]string** |  | [optional] 
**TenantOrganizationIds** | Pointer to **[]string** |  | [optional] 
**MfaRequired** | Pointer to **string** |  | [optional] 
**IsAutomaticAdminPromotionExempt** | Pointer to **bool** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**IsMaster** | Pointer to **bool** |  | [optional] 
**Subscription** | Pointer to [**Subscription**](Subscription.md) |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** |  | [optional] 
**Entitlements** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewOrganization

`func NewOrganization() *Organization`

NewOrganization instantiates a new Organization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationWithDefaults

`func NewOrganizationWithDefaults() *Organization`

NewOrganizationWithDefaults instantiates a new Organization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Organization) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Organization) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Organization) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Organization) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *Organization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Organization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Organization) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Organization) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Organization) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Organization) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Organization) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Organization) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Organization) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Organization) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Organization) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Organization) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetOwnerId

`func (o *Organization) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *Organization) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *Organization) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *Organization) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetClientId

`func (o *Organization) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Organization) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Organization) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *Organization) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetIdproviderId

`func (o *Organization) GetIdproviderId() string`

GetIdproviderId returns the IdproviderId field if non-nil, zero value otherwise.

### GetIdproviderIdOk

`func (o *Organization) GetIdproviderIdOk() (*string, bool)`

GetIdproviderIdOk returns a tuple with the IdproviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdproviderId

`func (o *Organization) SetIdproviderId(v string)`

SetIdproviderId sets IdproviderId field to given value.

### HasIdproviderId

`func (o *Organization) HasIdproviderId() bool`

HasIdproviderId returns a boolean if a field has been set.

### GetIsFederated

`func (o *Organization) GetIsFederated() bool`

GetIsFederated returns the IsFederated field if non-nil, zero value otherwise.

### GetIsFederatedOk

`func (o *Organization) GetIsFederatedOk() (*bool, bool)`

GetIsFederatedOk returns a tuple with the IsFederated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFederated

`func (o *Organization) SetIsFederated(v bool)`

SetIsFederated sets IsFederated field to given value.

### HasIsFederated

`func (o *Organization) HasIsFederated() bool`

HasIsFederated returns a boolean if a field has been set.

### GetParentOrganizationIds

`func (o *Organization) GetParentOrganizationIds() []string`

GetParentOrganizationIds returns the ParentOrganizationIds field if non-nil, zero value otherwise.

### GetParentOrganizationIdsOk

`func (o *Organization) GetParentOrganizationIdsOk() (*[]string, bool)`

GetParentOrganizationIdsOk returns a tuple with the ParentOrganizationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentOrganizationIds

`func (o *Organization) SetParentOrganizationIds(v []string)`

SetParentOrganizationIds sets ParentOrganizationIds field to given value.

### HasParentOrganizationIds

`func (o *Organization) HasParentOrganizationIds() bool`

HasParentOrganizationIds returns a boolean if a field has been set.

### GetSubOrganizationIds

`func (o *Organization) GetSubOrganizationIds() []string`

GetSubOrganizationIds returns the SubOrganizationIds field if non-nil, zero value otherwise.

### GetSubOrganizationIdsOk

`func (o *Organization) GetSubOrganizationIdsOk() (*[]string, bool)`

GetSubOrganizationIdsOk returns a tuple with the SubOrganizationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubOrganizationIds

`func (o *Organization) SetSubOrganizationIds(v []string)`

SetSubOrganizationIds sets SubOrganizationIds field to given value.

### HasSubOrganizationIds

`func (o *Organization) HasSubOrganizationIds() bool`

HasSubOrganizationIds returns a boolean if a field has been set.

### GetTenantOrganizationIds

`func (o *Organization) GetTenantOrganizationIds() []string`

GetTenantOrganizationIds returns the TenantOrganizationIds field if non-nil, zero value otherwise.

### GetTenantOrganizationIdsOk

`func (o *Organization) GetTenantOrganizationIdsOk() (*[]string, bool)`

GetTenantOrganizationIdsOk returns a tuple with the TenantOrganizationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantOrganizationIds

`func (o *Organization) SetTenantOrganizationIds(v []string)`

SetTenantOrganizationIds sets TenantOrganizationIds field to given value.

### HasTenantOrganizationIds

`func (o *Organization) HasTenantOrganizationIds() bool`

HasTenantOrganizationIds returns a boolean if a field has been set.

### GetMfaRequired

`func (o *Organization) GetMfaRequired() string`

GetMfaRequired returns the MfaRequired field if non-nil, zero value otherwise.

### GetMfaRequiredOk

`func (o *Organization) GetMfaRequiredOk() (*string, bool)`

GetMfaRequiredOk returns a tuple with the MfaRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaRequired

`func (o *Organization) SetMfaRequired(v string)`

SetMfaRequired sets MfaRequired field to given value.

### HasMfaRequired

`func (o *Organization) HasMfaRequired() bool`

HasMfaRequired returns a boolean if a field has been set.

### GetIsAutomaticAdminPromotionExempt

`func (o *Organization) GetIsAutomaticAdminPromotionExempt() bool`

GetIsAutomaticAdminPromotionExempt returns the IsAutomaticAdminPromotionExempt field if non-nil, zero value otherwise.

### GetIsAutomaticAdminPromotionExemptOk

`func (o *Organization) GetIsAutomaticAdminPromotionExemptOk() (*bool, bool)`

GetIsAutomaticAdminPromotionExemptOk returns a tuple with the IsAutomaticAdminPromotionExempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomaticAdminPromotionExempt

`func (o *Organization) SetIsAutomaticAdminPromotionExempt(v bool)`

SetIsAutomaticAdminPromotionExempt sets IsAutomaticAdminPromotionExempt field to given value.

### HasIsAutomaticAdminPromotionExempt

`func (o *Organization) HasIsAutomaticAdminPromotionExempt() bool`

HasIsAutomaticAdminPromotionExempt returns a boolean if a field has been set.

### GetDomain

`func (o *Organization) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Organization) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Organization) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *Organization) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetIsMaster

`func (o *Organization) GetIsMaster() bool`

GetIsMaster returns the IsMaster field if non-nil, zero value otherwise.

### GetIsMasterOk

`func (o *Organization) GetIsMasterOk() (*bool, bool)`

GetIsMasterOk returns a tuple with the IsMaster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMaster

`func (o *Organization) SetIsMaster(v bool)`

SetIsMaster sets IsMaster field to given value.

### HasIsMaster

`func (o *Organization) HasIsMaster() bool`

HasIsMaster returns a boolean if a field has been set.

### GetSubscription

`func (o *Organization) GetSubscription() Subscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *Organization) GetSubscriptionOk() (*Subscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *Organization) SetSubscription(v Subscription)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *Organization) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetProperties

`func (o *Organization) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Organization) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Organization) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Organization) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetEntitlements

`func (o *Organization) GetEntitlements() map[string]interface{}`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *Organization) GetEntitlementsOk() (*map[string]interface{}, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *Organization) SetEntitlements(v map[string]interface{})`

SetEntitlements sets Entitlements field to given value.

### HasEntitlements

`func (o *Organization) HasEntitlements() bool`

HasEntitlements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


