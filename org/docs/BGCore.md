# BGCore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | An explanation about the purpose of this instance. | [default to ""]
**CreatedAt** | **string** | An explanation about the purpose of this instance. | [default to ""]
**Domain** | **string** | An explanation about the purpose of this instance. | [default to ""]
**Entitlements** | [**Entitlements**](Entitlements.md) |  | 
**Environments** | [**[]Environment**](Environment.md) | An explanation about the purpose of this instance. | [default to []]
**Id** | **string** | An explanation about the purpose of this instance. | [default to ""]
**IdproviderId** | **string** | An explanation about the purpose of this instance. | [default to ""]
**IsAutomaticAdminPromotionExempt** | **bool** | An explanation about the purpose of this instance. | [default to false]
**IsFederated** | **bool** | An explanation about the purpose of this instance. | [default to false]
**IsMaster** | **bool** | An explanation about the purpose of this instance. | [default to false]
**MfaRequired** | **string** | An explanation about the purpose of this instance. | [default to ""]
**Name** | **string** | An explanation about the purpose of this instance. | [default to ""]
**OwnerId** | **string** | An explanation about the purpose of this instance. | [default to ""]
**ParentOrganizationIds** | **[]string** | An explanation about the purpose of this instance. | [default to []]
**Properties** | **map[string]interface{}** | An explanation about the purpose of this instance. | [default to {}]
**SubOrganizationIds** | **[]string** | An explanation about the purpose of this instance. | [default to []]
**TenantOrganizationIds** | **[]string** | An explanation about the purpose of this instance. | [default to []]
**UpdatedAt** | **string** | An explanation about the purpose of this instance. | [default to ""]
**Owner** | [**Owner**](Owner.md) |  | 

## Methods

### NewBGCore

`func NewBGCore(clientId string, createdAt string, domain string, entitlements Entitlements, environments []Environment, id string, idproviderId string, isAutomaticAdminPromotionExempt bool, isFederated bool, isMaster bool, mfaRequired string, name string, ownerId string, parentOrganizationIds []string, properties map[string]interface{}, subOrganizationIds []string, tenantOrganizationIds []string, updatedAt string, owner Owner, ) *BGCore`

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


### GetOwner

`func (o *BGCore) GetOwner() Owner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BGCore) GetOwnerOk() (*Owner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BGCore) SetOwner(v Owner)`

SetOwner sets Owner field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


