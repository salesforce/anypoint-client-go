# TierContractResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audit** | Pointer to [**Audit**](Audit.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Limits** | Pointer to [**[]TierContractResponseLimitsInner**](TierContractResponseLimitsInner.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**AutoApprove** | Pointer to **bool** |  | [optional] 

## Methods

### NewTierContractResponse

`func NewTierContractResponse() *TierContractResponse`

NewTierContractResponse instantiates a new TierContractResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTierContractResponseWithDefaults

`func NewTierContractResponseWithDefaults() *TierContractResponse`

NewTierContractResponseWithDefaults instantiates a new TierContractResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudit

`func (o *TierContractResponse) GetAudit() Audit`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *TierContractResponse) GetAuditOk() (*Audit, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *TierContractResponse) SetAudit(v Audit)`

SetAudit sets Audit field to given value.

### HasAudit

`func (o *TierContractResponse) HasAudit() bool`

HasAudit returns a boolean if a field has been set.

### GetId

`func (o *TierContractResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TierContractResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TierContractResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TierContractResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TierContractResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TierContractResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TierContractResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TierContractResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *TierContractResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TierContractResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TierContractResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TierContractResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TierContractResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TierContractResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLimits

`func (o *TierContractResponse) GetLimits() []TierContractResponseLimitsInner`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *TierContractResponse) GetLimitsOk() (*[]TierContractResponseLimitsInner, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *TierContractResponse) SetLimits(v []TierContractResponseLimitsInner)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *TierContractResponse) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetStatus

`func (o *TierContractResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TierContractResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TierContractResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TierContractResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAutoApprove

`func (o *TierContractResponse) GetAutoApprove() bool`

GetAutoApprove returns the AutoApprove field if non-nil, zero value otherwise.

### GetAutoApproveOk

`func (o *TierContractResponse) GetAutoApproveOk() (*bool, bool)`

GetAutoApproveOk returns a tuple with the AutoApprove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApprove

`func (o *TierContractResponse) SetAutoApprove(v bool)`

SetAutoApprove sets AutoApprove field to given value.

### HasAutoApprove

`func (o *TierContractResponse) HasAutoApprove() bool`

HasAutoApprove returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


