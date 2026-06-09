# Tier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audit** | Pointer to [**Audit**](Audit.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Limits** | Pointer to [**[]Limit**](Limit.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**AutoApprove** | Pointer to **bool** |  | [optional] 

## Methods

### NewTier

`func NewTier() *Tier`

NewTier instantiates a new Tier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTierWithDefaults

`func NewTierWithDefaults() *Tier`

NewTierWithDefaults instantiates a new Tier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudit

`func (o *Tier) GetAudit() Audit`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *Tier) GetAuditOk() (*Audit, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *Tier) SetAudit(v Audit)`

SetAudit sets Audit field to given value.

### HasAudit

`func (o *Tier) HasAudit() bool`

HasAudit returns a boolean if a field has been set.

### GetId

`func (o *Tier) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Tier) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Tier) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Tier) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Tier) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Tier) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Tier) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Tier) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Tier) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Tier) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Tier) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Tier) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Tier) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Tier) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLimits

`func (o *Tier) GetLimits() []Limit`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *Tier) GetLimitsOk() (*[]Limit, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *Tier) SetLimits(v []Limit)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *Tier) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetStatus

`func (o *Tier) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Tier) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Tier) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Tier) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAutoApprove

`func (o *Tier) GetAutoApprove() bool`

GetAutoApprove returns the AutoApprove field if non-nil, zero value otherwise.

### GetAutoApproveOk

`func (o *Tier) GetAutoApproveOk() (*bool, bool)`

GetAutoApproveOk returns a tuple with the AutoApprove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApprove

`func (o *Tier) SetAutoApprove(v bool)`

SetAutoApprove sets AutoApprove field to given value.

### HasAutoApprove

`func (o *Tier) HasAutoApprove() bool`

HasAutoApprove returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


