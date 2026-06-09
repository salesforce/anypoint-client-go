# SlaTierPutBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**AutoApprove** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Limits** | Pointer to [**[]SlaLimit**](SlaLimit.md) |  | [optional] 
**ApiVersionId** | Pointer to **string** |  | [optional] 
**ApiId** | Pointer to **int32** |  | [optional] 
**ApplicationCount** | Pointer to **int32** |  | [optional] 
**MasterOrganizationId** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**Audit** | Pointer to [**SlaTierAudit**](SlaTierAudit.md) |  | [optional] 

## Methods

### NewSlaTierPutBody

`func NewSlaTierPutBody() *SlaTierPutBody`

NewSlaTierPutBody instantiates a new SlaTierPutBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlaTierPutBodyWithDefaults

`func NewSlaTierPutBodyWithDefaults() *SlaTierPutBody`

NewSlaTierPutBodyWithDefaults instantiates a new SlaTierPutBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SlaTierPutBody) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SlaTierPutBody) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SlaTierPutBody) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SlaTierPutBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SlaTierPutBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SlaTierPutBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SlaTierPutBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SlaTierPutBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *SlaTierPutBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SlaTierPutBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SlaTierPutBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SlaTierPutBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SlaTierPutBody) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SlaTierPutBody) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAutoApprove

`func (o *SlaTierPutBody) GetAutoApprove() bool`

GetAutoApprove returns the AutoApprove field if non-nil, zero value otherwise.

### GetAutoApproveOk

`func (o *SlaTierPutBody) GetAutoApproveOk() (*bool, bool)`

GetAutoApproveOk returns a tuple with the AutoApprove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApprove

`func (o *SlaTierPutBody) SetAutoApprove(v bool)`

SetAutoApprove sets AutoApprove field to given value.

### HasAutoApprove

`func (o *SlaTierPutBody) HasAutoApprove() bool`

HasAutoApprove returns a boolean if a field has been set.

### GetStatus

`func (o *SlaTierPutBody) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SlaTierPutBody) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SlaTierPutBody) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SlaTierPutBody) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLimits

`func (o *SlaTierPutBody) GetLimits() []SlaLimit`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *SlaTierPutBody) GetLimitsOk() (*[]SlaLimit, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *SlaTierPutBody) SetLimits(v []SlaLimit)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *SlaTierPutBody) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetApiVersionId

`func (o *SlaTierPutBody) GetApiVersionId() string`

GetApiVersionId returns the ApiVersionId field if non-nil, zero value otherwise.

### GetApiVersionIdOk

`func (o *SlaTierPutBody) GetApiVersionIdOk() (*string, bool)`

GetApiVersionIdOk returns a tuple with the ApiVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersionId

`func (o *SlaTierPutBody) SetApiVersionId(v string)`

SetApiVersionId sets ApiVersionId field to given value.

### HasApiVersionId

`func (o *SlaTierPutBody) HasApiVersionId() bool`

HasApiVersionId returns a boolean if a field has been set.

### GetApiId

`func (o *SlaTierPutBody) GetApiId() int32`

GetApiId returns the ApiId field if non-nil, zero value otherwise.

### GetApiIdOk

`func (o *SlaTierPutBody) GetApiIdOk() (*int32, bool)`

GetApiIdOk returns a tuple with the ApiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiId

`func (o *SlaTierPutBody) SetApiId(v int32)`

SetApiId sets ApiId field to given value.

### HasApiId

`func (o *SlaTierPutBody) HasApiId() bool`

HasApiId returns a boolean if a field has been set.

### GetApplicationCount

`func (o *SlaTierPutBody) GetApplicationCount() int32`

GetApplicationCount returns the ApplicationCount field if non-nil, zero value otherwise.

### GetApplicationCountOk

`func (o *SlaTierPutBody) GetApplicationCountOk() (*int32, bool)`

GetApplicationCountOk returns a tuple with the ApplicationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCount

`func (o *SlaTierPutBody) SetApplicationCount(v int32)`

SetApplicationCount sets ApplicationCount field to given value.

### HasApplicationCount

`func (o *SlaTierPutBody) HasApplicationCount() bool`

HasApplicationCount returns a boolean if a field has been set.

### GetMasterOrganizationId

`func (o *SlaTierPutBody) GetMasterOrganizationId() string`

GetMasterOrganizationId returns the MasterOrganizationId field if non-nil, zero value otherwise.

### GetMasterOrganizationIdOk

`func (o *SlaTierPutBody) GetMasterOrganizationIdOk() (*string, bool)`

GetMasterOrganizationIdOk returns a tuple with the MasterOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterOrganizationId

`func (o *SlaTierPutBody) SetMasterOrganizationId(v string)`

SetMasterOrganizationId sets MasterOrganizationId field to given value.

### HasMasterOrganizationId

`func (o *SlaTierPutBody) HasMasterOrganizationId() bool`

HasMasterOrganizationId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *SlaTierPutBody) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *SlaTierPutBody) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *SlaTierPutBody) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *SlaTierPutBody) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetAudit

`func (o *SlaTierPutBody) GetAudit() SlaTierAudit`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *SlaTierPutBody) GetAuditOk() (*SlaTierAudit, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *SlaTierPutBody) SetAudit(v SlaTierAudit)`

SetAudit sets Audit field to given value.

### HasAudit

`func (o *SlaTierPutBody) HasAudit() bool`

HasAudit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


