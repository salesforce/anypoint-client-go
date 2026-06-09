# Contract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audit** | Pointer to [**Audit**](Audit.md) |  | [optional] 
**MasterOrganizationId** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ApprovedDate** | Pointer to **NullableString** |  | [optional] 
**RejectedDate** | Pointer to **NullableString** |  | [optional] 
**RevokedDate** | Pointer to **NullableString** |  | [optional] 
**ApplicationId** | Pointer to **int32** |  | [optional] 
**Application** | Pointer to [**Application**](Application.md) |  | [optional] 
**TierId** | Pointer to **NullableInt32** |  | [optional] 
**Tier** | Pointer to [**Tier**](Tier.md) |  | [optional] 
**RequestedTierId** | Pointer to **NullableInt32** |  | [optional] 
**RequestedTier** | Pointer to **map[string]interface{}** |  | [optional] 
**Terms** | Pointer to **NullableString** |  | [optional] 
**GroupInstanceId** | Pointer to **NullableString** |  | [optional] 
**GroupInstance** | Pointer to **map[string]interface{}** |  | [optional] 
**PartyId** | Pointer to **NullableString** |  | [optional] 
**PartyName** | Pointer to **NullableString** |  | [optional] 
**ApiId** | Pointer to **int32** |  | [optional] 
**Api** | Pointer to [**ApiReference**](ApiReference.md) |  | [optional] 

## Methods

### NewContract

`func NewContract() *Contract`

NewContract instantiates a new Contract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractWithDefaults

`func NewContractWithDefaults() *Contract`

NewContractWithDefaults instantiates a new Contract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudit

`func (o *Contract) GetAudit() Audit`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *Contract) GetAuditOk() (*Audit, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *Contract) SetAudit(v Audit)`

SetAudit sets Audit field to given value.

### HasAudit

`func (o *Contract) HasAudit() bool`

HasAudit returns a boolean if a field has been set.

### GetMasterOrganizationId

`func (o *Contract) GetMasterOrganizationId() string`

GetMasterOrganizationId returns the MasterOrganizationId field if non-nil, zero value otherwise.

### GetMasterOrganizationIdOk

`func (o *Contract) GetMasterOrganizationIdOk() (*string, bool)`

GetMasterOrganizationIdOk returns a tuple with the MasterOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterOrganizationId

`func (o *Contract) SetMasterOrganizationId(v string)`

SetMasterOrganizationId sets MasterOrganizationId field to given value.

### HasMasterOrganizationId

`func (o *Contract) HasMasterOrganizationId() bool`

HasMasterOrganizationId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *Contract) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Contract) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *Contract) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *Contract) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetId

`func (o *Contract) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Contract) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Contract) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Contract) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *Contract) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Contract) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Contract) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Contract) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetApprovedDate

`func (o *Contract) GetApprovedDate() string`

GetApprovedDate returns the ApprovedDate field if non-nil, zero value otherwise.

### GetApprovedDateOk

`func (o *Contract) GetApprovedDateOk() (*string, bool)`

GetApprovedDateOk returns a tuple with the ApprovedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedDate

`func (o *Contract) SetApprovedDate(v string)`

SetApprovedDate sets ApprovedDate field to given value.

### HasApprovedDate

`func (o *Contract) HasApprovedDate() bool`

HasApprovedDate returns a boolean if a field has been set.

### SetApprovedDateNil

`func (o *Contract) SetApprovedDateNil(b bool)`

 SetApprovedDateNil sets the value for ApprovedDate to be an explicit nil

### UnsetApprovedDate
`func (o *Contract) UnsetApprovedDate()`

UnsetApprovedDate ensures that no value is present for ApprovedDate, not even an explicit nil
### GetRejectedDate

`func (o *Contract) GetRejectedDate() string`

GetRejectedDate returns the RejectedDate field if non-nil, zero value otherwise.

### GetRejectedDateOk

`func (o *Contract) GetRejectedDateOk() (*string, bool)`

GetRejectedDateOk returns a tuple with the RejectedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedDate

`func (o *Contract) SetRejectedDate(v string)`

SetRejectedDate sets RejectedDate field to given value.

### HasRejectedDate

`func (o *Contract) HasRejectedDate() bool`

HasRejectedDate returns a boolean if a field has been set.

### SetRejectedDateNil

`func (o *Contract) SetRejectedDateNil(b bool)`

 SetRejectedDateNil sets the value for RejectedDate to be an explicit nil

### UnsetRejectedDate
`func (o *Contract) UnsetRejectedDate()`

UnsetRejectedDate ensures that no value is present for RejectedDate, not even an explicit nil
### GetRevokedDate

`func (o *Contract) GetRevokedDate() string`

GetRevokedDate returns the RevokedDate field if non-nil, zero value otherwise.

### GetRevokedDateOk

`func (o *Contract) GetRevokedDateOk() (*string, bool)`

GetRevokedDateOk returns a tuple with the RevokedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedDate

`func (o *Contract) SetRevokedDate(v string)`

SetRevokedDate sets RevokedDate field to given value.

### HasRevokedDate

`func (o *Contract) HasRevokedDate() bool`

HasRevokedDate returns a boolean if a field has been set.

### SetRevokedDateNil

`func (o *Contract) SetRevokedDateNil(b bool)`

 SetRevokedDateNil sets the value for RevokedDate to be an explicit nil

### UnsetRevokedDate
`func (o *Contract) UnsetRevokedDate()`

UnsetRevokedDate ensures that no value is present for RevokedDate, not even an explicit nil
### GetApplicationId

`func (o *Contract) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *Contract) GetApplicationIdOk() (*int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *Contract) SetApplicationId(v int32)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *Contract) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetApplication

`func (o *Contract) GetApplication() Application`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *Contract) GetApplicationOk() (*Application, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *Contract) SetApplication(v Application)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *Contract) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetTierId

`func (o *Contract) GetTierId() int32`

GetTierId returns the TierId field if non-nil, zero value otherwise.

### GetTierIdOk

`func (o *Contract) GetTierIdOk() (*int32, bool)`

GetTierIdOk returns a tuple with the TierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierId

`func (o *Contract) SetTierId(v int32)`

SetTierId sets TierId field to given value.

### HasTierId

`func (o *Contract) HasTierId() bool`

HasTierId returns a boolean if a field has been set.

### SetTierIdNil

`func (o *Contract) SetTierIdNil(b bool)`

 SetTierIdNil sets the value for TierId to be an explicit nil

### UnsetTierId
`func (o *Contract) UnsetTierId()`

UnsetTierId ensures that no value is present for TierId, not even an explicit nil
### GetTier

`func (o *Contract) GetTier() Tier`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *Contract) GetTierOk() (*Tier, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *Contract) SetTier(v Tier)`

SetTier sets Tier field to given value.

### HasTier

`func (o *Contract) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetRequestedTierId

`func (o *Contract) GetRequestedTierId() int32`

GetRequestedTierId returns the RequestedTierId field if non-nil, zero value otherwise.

### GetRequestedTierIdOk

`func (o *Contract) GetRequestedTierIdOk() (*int32, bool)`

GetRequestedTierIdOk returns a tuple with the RequestedTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedTierId

`func (o *Contract) SetRequestedTierId(v int32)`

SetRequestedTierId sets RequestedTierId field to given value.

### HasRequestedTierId

`func (o *Contract) HasRequestedTierId() bool`

HasRequestedTierId returns a boolean if a field has been set.

### SetRequestedTierIdNil

`func (o *Contract) SetRequestedTierIdNil(b bool)`

 SetRequestedTierIdNil sets the value for RequestedTierId to be an explicit nil

### UnsetRequestedTierId
`func (o *Contract) UnsetRequestedTierId()`

UnsetRequestedTierId ensures that no value is present for RequestedTierId, not even an explicit nil
### GetRequestedTier

`func (o *Contract) GetRequestedTier() map[string]interface{}`

GetRequestedTier returns the RequestedTier field if non-nil, zero value otherwise.

### GetRequestedTierOk

`func (o *Contract) GetRequestedTierOk() (*map[string]interface{}, bool)`

GetRequestedTierOk returns a tuple with the RequestedTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedTier

`func (o *Contract) SetRequestedTier(v map[string]interface{})`

SetRequestedTier sets RequestedTier field to given value.

### HasRequestedTier

`func (o *Contract) HasRequestedTier() bool`

HasRequestedTier returns a boolean if a field has been set.

### SetRequestedTierNil

`func (o *Contract) SetRequestedTierNil(b bool)`

 SetRequestedTierNil sets the value for RequestedTier to be an explicit nil

### UnsetRequestedTier
`func (o *Contract) UnsetRequestedTier()`

UnsetRequestedTier ensures that no value is present for RequestedTier, not even an explicit nil
### GetTerms

`func (o *Contract) GetTerms() string`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *Contract) GetTermsOk() (*string, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *Contract) SetTerms(v string)`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *Contract) HasTerms() bool`

HasTerms returns a boolean if a field has been set.

### SetTermsNil

`func (o *Contract) SetTermsNil(b bool)`

 SetTermsNil sets the value for Terms to be an explicit nil

### UnsetTerms
`func (o *Contract) UnsetTerms()`

UnsetTerms ensures that no value is present for Terms, not even an explicit nil
### GetGroupInstanceId

`func (o *Contract) GetGroupInstanceId() string`

GetGroupInstanceId returns the GroupInstanceId field if non-nil, zero value otherwise.

### GetGroupInstanceIdOk

`func (o *Contract) GetGroupInstanceIdOk() (*string, bool)`

GetGroupInstanceIdOk returns a tuple with the GroupInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupInstanceId

`func (o *Contract) SetGroupInstanceId(v string)`

SetGroupInstanceId sets GroupInstanceId field to given value.

### HasGroupInstanceId

`func (o *Contract) HasGroupInstanceId() bool`

HasGroupInstanceId returns a boolean if a field has been set.

### SetGroupInstanceIdNil

`func (o *Contract) SetGroupInstanceIdNil(b bool)`

 SetGroupInstanceIdNil sets the value for GroupInstanceId to be an explicit nil

### UnsetGroupInstanceId
`func (o *Contract) UnsetGroupInstanceId()`

UnsetGroupInstanceId ensures that no value is present for GroupInstanceId, not even an explicit nil
### GetGroupInstance

`func (o *Contract) GetGroupInstance() map[string]interface{}`

GetGroupInstance returns the GroupInstance field if non-nil, zero value otherwise.

### GetGroupInstanceOk

`func (o *Contract) GetGroupInstanceOk() (*map[string]interface{}, bool)`

GetGroupInstanceOk returns a tuple with the GroupInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupInstance

`func (o *Contract) SetGroupInstance(v map[string]interface{})`

SetGroupInstance sets GroupInstance field to given value.

### HasGroupInstance

`func (o *Contract) HasGroupInstance() bool`

HasGroupInstance returns a boolean if a field has been set.

### SetGroupInstanceNil

`func (o *Contract) SetGroupInstanceNil(b bool)`

 SetGroupInstanceNil sets the value for GroupInstance to be an explicit nil

### UnsetGroupInstance
`func (o *Contract) UnsetGroupInstance()`

UnsetGroupInstance ensures that no value is present for GroupInstance, not even an explicit nil
### GetPartyId

`func (o *Contract) GetPartyId() string`

GetPartyId returns the PartyId field if non-nil, zero value otherwise.

### GetPartyIdOk

`func (o *Contract) GetPartyIdOk() (*string, bool)`

GetPartyIdOk returns a tuple with the PartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyId

`func (o *Contract) SetPartyId(v string)`

SetPartyId sets PartyId field to given value.

### HasPartyId

`func (o *Contract) HasPartyId() bool`

HasPartyId returns a boolean if a field has been set.

### SetPartyIdNil

`func (o *Contract) SetPartyIdNil(b bool)`

 SetPartyIdNil sets the value for PartyId to be an explicit nil

### UnsetPartyId
`func (o *Contract) UnsetPartyId()`

UnsetPartyId ensures that no value is present for PartyId, not even an explicit nil
### GetPartyName

`func (o *Contract) GetPartyName() string`

GetPartyName returns the PartyName field if non-nil, zero value otherwise.

### GetPartyNameOk

`func (o *Contract) GetPartyNameOk() (*string, bool)`

GetPartyNameOk returns a tuple with the PartyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyName

`func (o *Contract) SetPartyName(v string)`

SetPartyName sets PartyName field to given value.

### HasPartyName

`func (o *Contract) HasPartyName() bool`

HasPartyName returns a boolean if a field has been set.

### SetPartyNameNil

`func (o *Contract) SetPartyNameNil(b bool)`

 SetPartyNameNil sets the value for PartyName to be an explicit nil

### UnsetPartyName
`func (o *Contract) UnsetPartyName()`

UnsetPartyName ensures that no value is present for PartyName, not even an explicit nil
### GetApiId

`func (o *Contract) GetApiId() int32`

GetApiId returns the ApiId field if non-nil, zero value otherwise.

### GetApiIdOk

`func (o *Contract) GetApiIdOk() (*int32, bool)`

GetApiIdOk returns a tuple with the ApiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiId

`func (o *Contract) SetApiId(v int32)`

SetApiId sets ApiId field to given value.

### HasApiId

`func (o *Contract) HasApiId() bool`

HasApiId returns a boolean if a field has been set.

### GetApi

`func (o *Contract) GetApi() ApiReference`

GetApi returns the Api field if non-nil, zero value otherwise.

### GetApiOk

`func (o *Contract) GetApiOk() (*ApiReference, bool)`

GetApiOk returns a tuple with the Api field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApi

`func (o *Contract) SetApi(v ApiReference)`

SetApi sets Api field to given value.

### HasApi

`func (o *Contract) HasApi() bool`

HasApi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


