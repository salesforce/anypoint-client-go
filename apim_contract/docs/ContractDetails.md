# ContractDetails

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
**Application** | Pointer to [**ApplicationContractResponse**](ApplicationContractResponse.md) |  | [optional] 
**TierId** | Pointer to **int32** |  | [optional] 
**Tier** | Pointer to [**TierContractResponse**](TierContractResponse.md) |  | [optional] 
**RequestedTierId** | Pointer to **NullableInt32** |  | [optional] 
**RequestedTier** | Pointer to [**TierContractResponse**](TierContractResponse.md) |  | [optional] 
**Condition** | Pointer to **string** |  | [optional] [default to "APPLIED"]
**ApiId** | Pointer to **int32** |  | [optional] 
**Api** | Pointer to [**ContractDetailsApi**](ContractDetailsApi.md) |  | [optional] 

## Methods

### NewContractDetails

`func NewContractDetails() *ContractDetails`

NewContractDetails instantiates a new ContractDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractDetailsWithDefaults

`func NewContractDetailsWithDefaults() *ContractDetails`

NewContractDetailsWithDefaults instantiates a new ContractDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudit

`func (o *ContractDetails) GetAudit() Audit`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *ContractDetails) GetAuditOk() (*Audit, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *ContractDetails) SetAudit(v Audit)`

SetAudit sets Audit field to given value.

### HasAudit

`func (o *ContractDetails) HasAudit() bool`

HasAudit returns a boolean if a field has been set.

### GetMasterOrganizationId

`func (o *ContractDetails) GetMasterOrganizationId() string`

GetMasterOrganizationId returns the MasterOrganizationId field if non-nil, zero value otherwise.

### GetMasterOrganizationIdOk

`func (o *ContractDetails) GetMasterOrganizationIdOk() (*string, bool)`

GetMasterOrganizationIdOk returns a tuple with the MasterOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterOrganizationId

`func (o *ContractDetails) SetMasterOrganizationId(v string)`

SetMasterOrganizationId sets MasterOrganizationId field to given value.

### HasMasterOrganizationId

`func (o *ContractDetails) HasMasterOrganizationId() bool`

HasMasterOrganizationId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *ContractDetails) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ContractDetails) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ContractDetails) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ContractDetails) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetId

`func (o *ContractDetails) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContractDetails) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContractDetails) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ContractDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *ContractDetails) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ContractDetails) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ContractDetails) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ContractDetails) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetApprovedDate

`func (o *ContractDetails) GetApprovedDate() string`

GetApprovedDate returns the ApprovedDate field if non-nil, zero value otherwise.

### GetApprovedDateOk

`func (o *ContractDetails) GetApprovedDateOk() (*string, bool)`

GetApprovedDateOk returns a tuple with the ApprovedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedDate

`func (o *ContractDetails) SetApprovedDate(v string)`

SetApprovedDate sets ApprovedDate field to given value.

### HasApprovedDate

`func (o *ContractDetails) HasApprovedDate() bool`

HasApprovedDate returns a boolean if a field has been set.

### SetApprovedDateNil

`func (o *ContractDetails) SetApprovedDateNil(b bool)`

 SetApprovedDateNil sets the value for ApprovedDate to be an explicit nil

### UnsetApprovedDate
`func (o *ContractDetails) UnsetApprovedDate()`

UnsetApprovedDate ensures that no value is present for ApprovedDate, not even an explicit nil
### GetRejectedDate

`func (o *ContractDetails) GetRejectedDate() string`

GetRejectedDate returns the RejectedDate field if non-nil, zero value otherwise.

### GetRejectedDateOk

`func (o *ContractDetails) GetRejectedDateOk() (*string, bool)`

GetRejectedDateOk returns a tuple with the RejectedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedDate

`func (o *ContractDetails) SetRejectedDate(v string)`

SetRejectedDate sets RejectedDate field to given value.

### HasRejectedDate

`func (o *ContractDetails) HasRejectedDate() bool`

HasRejectedDate returns a boolean if a field has been set.

### SetRejectedDateNil

`func (o *ContractDetails) SetRejectedDateNil(b bool)`

 SetRejectedDateNil sets the value for RejectedDate to be an explicit nil

### UnsetRejectedDate
`func (o *ContractDetails) UnsetRejectedDate()`

UnsetRejectedDate ensures that no value is present for RejectedDate, not even an explicit nil
### GetRevokedDate

`func (o *ContractDetails) GetRevokedDate() string`

GetRevokedDate returns the RevokedDate field if non-nil, zero value otherwise.

### GetRevokedDateOk

`func (o *ContractDetails) GetRevokedDateOk() (*string, bool)`

GetRevokedDateOk returns a tuple with the RevokedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedDate

`func (o *ContractDetails) SetRevokedDate(v string)`

SetRevokedDate sets RevokedDate field to given value.

### HasRevokedDate

`func (o *ContractDetails) HasRevokedDate() bool`

HasRevokedDate returns a boolean if a field has been set.

### SetRevokedDateNil

`func (o *ContractDetails) SetRevokedDateNil(b bool)`

 SetRevokedDateNil sets the value for RevokedDate to be an explicit nil

### UnsetRevokedDate
`func (o *ContractDetails) UnsetRevokedDate()`

UnsetRevokedDate ensures that no value is present for RevokedDate, not even an explicit nil
### GetApplicationId

`func (o *ContractDetails) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ContractDetails) GetApplicationIdOk() (*int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ContractDetails) SetApplicationId(v int32)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *ContractDetails) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetApplication

`func (o *ContractDetails) GetApplication() ApplicationContractResponse`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *ContractDetails) GetApplicationOk() (*ApplicationContractResponse, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *ContractDetails) SetApplication(v ApplicationContractResponse)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *ContractDetails) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetTierId

`func (o *ContractDetails) GetTierId() int32`

GetTierId returns the TierId field if non-nil, zero value otherwise.

### GetTierIdOk

`func (o *ContractDetails) GetTierIdOk() (*int32, bool)`

GetTierIdOk returns a tuple with the TierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierId

`func (o *ContractDetails) SetTierId(v int32)`

SetTierId sets TierId field to given value.

### HasTierId

`func (o *ContractDetails) HasTierId() bool`

HasTierId returns a boolean if a field has been set.

### GetTier

`func (o *ContractDetails) GetTier() TierContractResponse`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *ContractDetails) GetTierOk() (*TierContractResponse, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *ContractDetails) SetTier(v TierContractResponse)`

SetTier sets Tier field to given value.

### HasTier

`func (o *ContractDetails) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetRequestedTierId

`func (o *ContractDetails) GetRequestedTierId() int32`

GetRequestedTierId returns the RequestedTierId field if non-nil, zero value otherwise.

### GetRequestedTierIdOk

`func (o *ContractDetails) GetRequestedTierIdOk() (*int32, bool)`

GetRequestedTierIdOk returns a tuple with the RequestedTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedTierId

`func (o *ContractDetails) SetRequestedTierId(v int32)`

SetRequestedTierId sets RequestedTierId field to given value.

### HasRequestedTierId

`func (o *ContractDetails) HasRequestedTierId() bool`

HasRequestedTierId returns a boolean if a field has been set.

### SetRequestedTierIdNil

`func (o *ContractDetails) SetRequestedTierIdNil(b bool)`

 SetRequestedTierIdNil sets the value for RequestedTierId to be an explicit nil

### UnsetRequestedTierId
`func (o *ContractDetails) UnsetRequestedTierId()`

UnsetRequestedTierId ensures that no value is present for RequestedTierId, not even an explicit nil
### GetRequestedTier

`func (o *ContractDetails) GetRequestedTier() TierContractResponse`

GetRequestedTier returns the RequestedTier field if non-nil, zero value otherwise.

### GetRequestedTierOk

`func (o *ContractDetails) GetRequestedTierOk() (*TierContractResponse, bool)`

GetRequestedTierOk returns a tuple with the RequestedTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedTier

`func (o *ContractDetails) SetRequestedTier(v TierContractResponse)`

SetRequestedTier sets RequestedTier field to given value.

### HasRequestedTier

`func (o *ContractDetails) HasRequestedTier() bool`

HasRequestedTier returns a boolean if a field has been set.

### GetCondition

`func (o *ContractDetails) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ContractDetails) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ContractDetails) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *ContractDetails) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetApiId

`func (o *ContractDetails) GetApiId() int32`

GetApiId returns the ApiId field if non-nil, zero value otherwise.

### GetApiIdOk

`func (o *ContractDetails) GetApiIdOk() (*int32, bool)`

GetApiIdOk returns a tuple with the ApiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiId

`func (o *ContractDetails) SetApiId(v int32)`

SetApiId sets ApiId field to given value.

### HasApiId

`func (o *ContractDetails) HasApiId() bool`

HasApiId returns a boolean if a field has been set.

### GetApi

`func (o *ContractDetails) GetApi() ContractDetailsApi`

GetApi returns the Api field if non-nil, zero value otherwise.

### GetApiOk

`func (o *ContractDetails) GetApiOk() (*ContractDetailsApi, bool)`

GetApiOk returns a tuple with the Api field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApi

`func (o *ContractDetails) SetApi(v ContractDetailsApi)`

SetApi sets Api field to given value.

### HasApi

`func (o *ContractDetails) HasApi() bool`

HasApi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


