# BGCore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | The client id of the organization | [optional] [default to ""]
**CreatedAt** | Pointer to **string** | The creation date of the organization | [optional] [default to ""]
**Domain** | Pointer to **string** | The domain of the organization | [optional] [default to ""]
**Entitlements** | Pointer to [**Entitlements**](Entitlements.md) |  | [optional] 
**Environments** | Pointer to [**[]Environment**](Environment.md) | List of environments. | [optional] 
**Id** | Pointer to **string** | The id of the organization | [optional] [default to ""]
**IdproviderId** | Pointer to **string** | The id of the idprovider of the organization. | [optional] [default to ""]
**IsAutomaticAdminPromotionExempt** | Pointer to **bool** | Whether the organization is exempt from automatic admin promotion. | [optional] [default to false]
**IsFederated** | Pointer to **bool** | Whether the organization is federated. | [optional] [default to false]
**IsRoot** | Pointer to **bool** | Whether the organization is a root organization. | [optional] [default to false]
**IsMaster** | Pointer to **bool** | Whether the organization is a master organization. | [optional] [default to false]
**MfaRequired** | Pointer to **string** | Whether multi-factor authentication is required for the organization. | [optional] [default to ""]
**Name** | Pointer to **string** | The name of the organization. | [optional] [default to ""]
**OwnerId** | Pointer to **string** | The id of the owner of the organization. | [optional] [default to ""]
**ParentOrganizationIds** | Pointer to **[]string** | The ids of the parent organizations. | [optional] 
**Properties** | Pointer to **map[string]interface{}** | The properties of the organization. | [optional] 
**SubOrganizationIds** | Pointer to **[]string** | The ids of the sub organizations. | [optional] 
**TenantOrganizationIds** | Pointer to **[]string** | The ids of the tenant organizations. | [optional] 
**UpdatedAt** | Pointer to **string** | The last update date of the organization. | [optional] [default to ""]
**Owner** | Pointer to [**User**](User.md) |  | [optional] 

## Methods

### NewBGCore

`func NewBGCore() *BGCore`

NewBGCore instantiates a new BGCore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBGCoreWithDefaults

`func NewBGCoreWithDefaults() *BGCore`

NewBGCoreWithDefaults instantiates a new BGCore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *BGCore) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *BGCore) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *BGCore) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *BGCore) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BGCore) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BGCore) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BGCore) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BGCore) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDomain

`func (o *BGCore) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *BGCore) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *BGCore) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *BGCore) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetEntitlements

`func (o *BGCore) GetEntitlements() Entitlements`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *BGCore) GetEntitlementsOk() (*Entitlements, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *BGCore) SetEntitlements(v Entitlements)`

SetEntitlements sets Entitlements field to given value.

### HasEntitlements

`func (o *BGCore) HasEntitlements() bool`

HasEntitlements returns a boolean if a field has been set.

### GetEnvironments

`func (o *BGCore) GetEnvironments() []Environment`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *BGCore) GetEnvironmentsOk() (*[]Environment, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *BGCore) SetEnvironments(v []Environment)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *BGCore) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetId

`func (o *BGCore) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BGCore) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BGCore) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BGCore) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdproviderId

`func (o *BGCore) GetIdproviderId() string`

GetIdproviderId returns the IdproviderId field if non-nil, zero value otherwise.

### GetIdproviderIdOk

`func (o *BGCore) GetIdproviderIdOk() (*string, bool)`

GetIdproviderIdOk returns a tuple with the IdproviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdproviderId

`func (o *BGCore) SetIdproviderId(v string)`

SetIdproviderId sets IdproviderId field to given value.

### HasIdproviderId

`func (o *BGCore) HasIdproviderId() bool`

HasIdproviderId returns a boolean if a field has been set.

### GetIsAutomaticAdminPromotionExempt

`func (o *BGCore) GetIsAutomaticAdminPromotionExempt() bool`

GetIsAutomaticAdminPromotionExempt returns the IsAutomaticAdminPromotionExempt field if non-nil, zero value otherwise.

### GetIsAutomaticAdminPromotionExemptOk

`func (o *BGCore) GetIsAutomaticAdminPromotionExemptOk() (*bool, bool)`

GetIsAutomaticAdminPromotionExemptOk returns a tuple with the IsAutomaticAdminPromotionExempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomaticAdminPromotionExempt

`func (o *BGCore) SetIsAutomaticAdminPromotionExempt(v bool)`

SetIsAutomaticAdminPromotionExempt sets IsAutomaticAdminPromotionExempt field to given value.

### HasIsAutomaticAdminPromotionExempt

`func (o *BGCore) HasIsAutomaticAdminPromotionExempt() bool`

HasIsAutomaticAdminPromotionExempt returns a boolean if a field has been set.

### GetIsFederated

`func (o *BGCore) GetIsFederated() bool`

GetIsFederated returns the IsFederated field if non-nil, zero value otherwise.

### GetIsFederatedOk

`func (o *BGCore) GetIsFederatedOk() (*bool, bool)`

GetIsFederatedOk returns a tuple with the IsFederated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFederated

`func (o *BGCore) SetIsFederated(v bool)`

SetIsFederated sets IsFederated field to given value.

### HasIsFederated

`func (o *BGCore) HasIsFederated() bool`

HasIsFederated returns a boolean if a field has been set.

### GetIsRoot

`func (o *BGCore) GetIsRoot() bool`

GetIsRoot returns the IsRoot field if non-nil, zero value otherwise.

### GetIsRootOk

`func (o *BGCore) GetIsRootOk() (*bool, bool)`

GetIsRootOk returns a tuple with the IsRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRoot

`func (o *BGCore) SetIsRoot(v bool)`

SetIsRoot sets IsRoot field to given value.

### HasIsRoot

`func (o *BGCore) HasIsRoot() bool`

HasIsRoot returns a boolean if a field has been set.

### GetIsMaster

`func (o *BGCore) GetIsMaster() bool`

GetIsMaster returns the IsMaster field if non-nil, zero value otherwise.

### GetIsMasterOk

`func (o *BGCore) GetIsMasterOk() (*bool, bool)`

GetIsMasterOk returns a tuple with the IsMaster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMaster

`func (o *BGCore) SetIsMaster(v bool)`

SetIsMaster sets IsMaster field to given value.

### HasIsMaster

`func (o *BGCore) HasIsMaster() bool`

HasIsMaster returns a boolean if a field has been set.

### GetMfaRequired

`func (o *BGCore) GetMfaRequired() string`

GetMfaRequired returns the MfaRequired field if non-nil, zero value otherwise.

### GetMfaRequiredOk

`func (o *BGCore) GetMfaRequiredOk() (*string, bool)`

GetMfaRequiredOk returns a tuple with the MfaRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaRequired

`func (o *BGCore) SetMfaRequired(v string)`

SetMfaRequired sets MfaRequired field to given value.

### HasMfaRequired

`func (o *BGCore) HasMfaRequired() bool`

HasMfaRequired returns a boolean if a field has been set.

### GetName

`func (o *BGCore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BGCore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BGCore) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BGCore) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwnerId

`func (o *BGCore) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *BGCore) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *BGCore) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *BGCore) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetParentOrganizationIds

`func (o *BGCore) GetParentOrganizationIds() []string`

GetParentOrganizationIds returns the ParentOrganizationIds field if non-nil, zero value otherwise.

### GetParentOrganizationIdsOk

`func (o *BGCore) GetParentOrganizationIdsOk() (*[]string, bool)`

GetParentOrganizationIdsOk returns a tuple with the ParentOrganizationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentOrganizationIds

`func (o *BGCore) SetParentOrganizationIds(v []string)`

SetParentOrganizationIds sets ParentOrganizationIds field to given value.

### HasParentOrganizationIds

`func (o *BGCore) HasParentOrganizationIds() bool`

HasParentOrganizationIds returns a boolean if a field has been set.

### GetProperties

`func (o *BGCore) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *BGCore) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *BGCore) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *BGCore) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetSubOrganizationIds

`func (o *BGCore) GetSubOrganizationIds() []string`

GetSubOrganizationIds returns the SubOrganizationIds field if non-nil, zero value otherwise.

### GetSubOrganizationIdsOk

`func (o *BGCore) GetSubOrganizationIdsOk() (*[]string, bool)`

GetSubOrganizationIdsOk returns a tuple with the SubOrganizationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubOrganizationIds

`func (o *BGCore) SetSubOrganizationIds(v []string)`

SetSubOrganizationIds sets SubOrganizationIds field to given value.

### HasSubOrganizationIds

`func (o *BGCore) HasSubOrganizationIds() bool`

HasSubOrganizationIds returns a boolean if a field has been set.

### GetTenantOrganizationIds

`func (o *BGCore) GetTenantOrganizationIds() []string`

GetTenantOrganizationIds returns the TenantOrganizationIds field if non-nil, zero value otherwise.

### GetTenantOrganizationIdsOk

`func (o *BGCore) GetTenantOrganizationIdsOk() (*[]string, bool)`

GetTenantOrganizationIdsOk returns a tuple with the TenantOrganizationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantOrganizationIds

`func (o *BGCore) SetTenantOrganizationIds(v []string)`

SetTenantOrganizationIds sets TenantOrganizationIds field to given value.

### HasTenantOrganizationIds

`func (o *BGCore) HasTenantOrganizationIds() bool`

HasTenantOrganizationIds returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *BGCore) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BGCore) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BGCore) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BGCore) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetOwner

`func (o *BGCore) GetOwner() User`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BGCore) GetOwnerOk() (*User, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BGCore) SetOwner(v User)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BGCore) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


