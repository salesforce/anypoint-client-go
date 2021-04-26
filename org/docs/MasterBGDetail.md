# MasterBGDetail

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
**SessionTimeout** | **int32** | An explanation about the purpose of this instance. | [default to 0]
**Subscription** | **map[string]map[string]interface{}** | An explanation about the purpose of this instance. | [default to {}]

## Methods

### NewMasterBGDetail

`func NewMasterBGDetail(clientId string, createdAt string, domain string, entitlements map[string]map[string]interface{}, environments []AnyOfmap, id string, idproviderId string, isAutomaticAdminPromotionExempt bool, isFederated bool, isMaster bool, mfaRequired string, name string, ownerId string, parentOrganizationIds []string, properties map[string]map[string]interface{}, subOrganizationIds []AnyOfstring, tenantOrganizationIds []string, updatedAt string, sessionTimeout int32, subscription map[string]map[string]interface{}, ) *MasterBGDetail`

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


### GetEntitlements

`func (o *MasterBGDetail) GetEntitlements() map[string]map[string]interface{}`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *MasterBGDetail) GetEntitlementsOk() (*map[string]map[string]interface{}, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *MasterBGDetail) SetEntitlements(v map[string]map[string]interface{})`

SetEntitlements sets Entitlements field to given value.


### GetEnvironments

`func (o *MasterBGDetail) GetEnvironments() []AnyOfmap`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *MasterBGDetail) GetEnvironmentsOk() (*[]AnyOfmap, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *MasterBGDetail) SetEnvironments(v []AnyOfmap)`

SetEnvironments sets Environments field to given value.


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


### GetProperties

`func (o *MasterBGDetail) GetProperties() map[string]map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *MasterBGDetail) GetPropertiesOk() (*map[string]map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *MasterBGDetail) SetProperties(v map[string]map[string]interface{})`

SetProperties sets Properties field to given value.


### GetSubOrganizationIds

`func (o *MasterBGDetail) GetSubOrganizationIds() []AnyOfstring`

GetSubOrganizationIds returns the SubOrganizationIds field if non-nil, zero value otherwise.

### GetSubOrganizationIdsOk

`func (o *MasterBGDetail) GetSubOrganizationIdsOk() (*[]AnyOfstring, bool)`

GetSubOrganizationIdsOk returns a tuple with the SubOrganizationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubOrganizationIds

`func (o *MasterBGDetail) SetSubOrganizationIds(v []AnyOfstring)`

SetSubOrganizationIds sets SubOrganizationIds field to given value.


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


### GetSubscription

`func (o *MasterBGDetail) GetSubscription() map[string]map[string]interface{}`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *MasterBGDetail) GetSubscriptionOk() (*map[string]map[string]interface{}, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *MasterBGDetail) SetSubscription(v map[string]map[string]interface{})`

SetSubscription sets Subscription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


