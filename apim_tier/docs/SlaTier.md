# SlaTier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audit** | Pointer to [**SlaTierAudit**](SlaTierAudit.md) |  | [optional] 
**MasterOrganizationId** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**AutoApprove** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Limits** | Pointer to [**[]SlaLimit**](SlaLimit.md) |  | [optional] 
**ApplicationCount** | Pointer to **int32** |  | [optional] 
**ApiId** | Pointer to **int32** |  | [optional] 
**ApiVersionId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSlaTier

`func NewSlaTier() *SlaTier`

NewSlaTier instantiates a new SlaTier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlaTierWithDefaults

`func NewSlaTierWithDefaults() *SlaTier`

NewSlaTierWithDefaults instantiates a new SlaTier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudit

`func (o *SlaTier) GetAudit() SlaTierAudit`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *SlaTier) GetAuditOk() (*SlaTierAudit, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *SlaTier) SetAudit(v SlaTierAudit)`

SetAudit sets Audit field to given value.

### HasAudit

`func (o *SlaTier) HasAudit() bool`

HasAudit returns a boolean if a field has been set.

### GetMasterOrganizationId

`func (o *SlaTier) GetMasterOrganizationId() string`

GetMasterOrganizationId returns the MasterOrganizationId field if non-nil, zero value otherwise.

### GetMasterOrganizationIdOk

`func (o *SlaTier) GetMasterOrganizationIdOk() (*string, bool)`

GetMasterOrganizationIdOk returns a tuple with the MasterOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterOrganizationId

`func (o *SlaTier) SetMasterOrganizationId(v string)`

SetMasterOrganizationId sets MasterOrganizationId field to given value.

### HasMasterOrganizationId

`func (o *SlaTier) HasMasterOrganizationId() bool`

HasMasterOrganizationId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *SlaTier) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *SlaTier) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *SlaTier) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *SlaTier) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetId

`func (o *SlaTier) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SlaTier) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SlaTier) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SlaTier) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SlaTier) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SlaTier) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SlaTier) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SlaTier) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *SlaTier) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SlaTier) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SlaTier) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SlaTier) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SlaTier) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SlaTier) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAutoApprove

`func (o *SlaTier) GetAutoApprove() bool`

GetAutoApprove returns the AutoApprove field if non-nil, zero value otherwise.

### GetAutoApproveOk

`func (o *SlaTier) GetAutoApproveOk() (*bool, bool)`

GetAutoApproveOk returns a tuple with the AutoApprove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApprove

`func (o *SlaTier) SetAutoApprove(v bool)`

SetAutoApprove sets AutoApprove field to given value.

### HasAutoApprove

`func (o *SlaTier) HasAutoApprove() bool`

HasAutoApprove returns a boolean if a field has been set.

### GetStatus

`func (o *SlaTier) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SlaTier) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SlaTier) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SlaTier) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLimits

`func (o *SlaTier) GetLimits() []SlaLimit`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *SlaTier) GetLimitsOk() (*[]SlaLimit, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *SlaTier) SetLimits(v []SlaLimit)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *SlaTier) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetApplicationCount

`func (o *SlaTier) GetApplicationCount() int32`

GetApplicationCount returns the ApplicationCount field if non-nil, zero value otherwise.

### GetApplicationCountOk

`func (o *SlaTier) GetApplicationCountOk() (*int32, bool)`

GetApplicationCountOk returns a tuple with the ApplicationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCount

`func (o *SlaTier) SetApplicationCount(v int32)`

SetApplicationCount sets ApplicationCount field to given value.

### HasApplicationCount

`func (o *SlaTier) HasApplicationCount() bool`

HasApplicationCount returns a boolean if a field has been set.

### GetApiId

`func (o *SlaTier) GetApiId() int32`

GetApiId returns the ApiId field if non-nil, zero value otherwise.

### GetApiIdOk

`func (o *SlaTier) GetApiIdOk() (*int32, bool)`

GetApiIdOk returns a tuple with the ApiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiId

`func (o *SlaTier) SetApiId(v int32)`

SetApiId sets ApiId field to given value.

### HasApiId

`func (o *SlaTier) HasApiId() bool`

HasApiId returns a boolean if a field has been set.

### GetApiVersionId

`func (o *SlaTier) GetApiVersionId() string`

GetApiVersionId returns the ApiVersionId field if non-nil, zero value otherwise.

### GetApiVersionIdOk

`func (o *SlaTier) GetApiVersionIdOk() (*string, bool)`

GetApiVersionIdOk returns a tuple with the ApiVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersionId

`func (o *SlaTier) SetApiVersionId(v string)`

SetApiVersionId sets ApiVersionId field to given value.

### HasApiVersionId

`func (o *SlaTier) HasApiVersionId() bool`

HasApiVersionId returns a boolean if a field has been set.

### SetApiVersionIdNil

`func (o *SlaTier) SetApiVersionIdNil(b bool)`

 SetApiVersionIdNil sets the value for ApiVersionId to be an explicit nil

### UnsetApiVersionId
`func (o *SlaTier) UnsetApiVersionId()`

UnsetApiVersionId ensures that no value is present for ApiVersionId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


