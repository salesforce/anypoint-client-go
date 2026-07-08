# SlaTierPostBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**AutoApprove** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Limits** | Pointer to [**[]SlaLimit**](SlaLimit.md) |  | [optional] 
**ApiVersionId** | Pointer to **string** |  | [optional] 

## Methods

### NewSlaTierPostBody

`func NewSlaTierPostBody() *SlaTierPostBody`

NewSlaTierPostBody instantiates a new SlaTierPostBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlaTierPostBodyWithDefaults

`func NewSlaTierPostBodyWithDefaults() *SlaTierPostBody`

NewSlaTierPostBodyWithDefaults instantiates a new SlaTierPostBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SlaTierPostBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SlaTierPostBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SlaTierPostBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SlaTierPostBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *SlaTierPostBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SlaTierPostBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SlaTierPostBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SlaTierPostBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SlaTierPostBody) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SlaTierPostBody) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAutoApprove

`func (o *SlaTierPostBody) GetAutoApprove() bool`

GetAutoApprove returns the AutoApprove field if non-nil, zero value otherwise.

### GetAutoApproveOk

`func (o *SlaTierPostBody) GetAutoApproveOk() (*bool, bool)`

GetAutoApproveOk returns a tuple with the AutoApprove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApprove

`func (o *SlaTierPostBody) SetAutoApprove(v bool)`

SetAutoApprove sets AutoApprove field to given value.

### HasAutoApprove

`func (o *SlaTierPostBody) HasAutoApprove() bool`

HasAutoApprove returns a boolean if a field has been set.

### GetStatus

`func (o *SlaTierPostBody) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SlaTierPostBody) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SlaTierPostBody) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SlaTierPostBody) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLimits

`func (o *SlaTierPostBody) GetLimits() []SlaLimit`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *SlaTierPostBody) GetLimitsOk() (*[]SlaLimit, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *SlaTierPostBody) SetLimits(v []SlaLimit)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *SlaTierPostBody) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetApiVersionId

`func (o *SlaTierPostBody) GetApiVersionId() string`

GetApiVersionId returns the ApiVersionId field if non-nil, zero value otherwise.

### GetApiVersionIdOk

`func (o *SlaTierPostBody) GetApiVersionIdOk() (*string, bool)`

GetApiVersionIdOk returns a tuple with the ApiVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersionId

`func (o *SlaTierPostBody) SetApiVersionId(v string)`

SetApiVersionId sets ApiVersionId field to given value.

### HasApiVersionId

`func (o *SlaTierPostBody) HasApiVersionId() bool`

HasApiVersionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


