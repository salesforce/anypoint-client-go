# FlexGatewayTargetSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Replicas** | Pointer to [**FlexGatewayTargetSummaryReplicas**](FlexGatewayTargetSummaryReplicas.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**LastUpdate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewFlexGatewayTargetSummary

`func NewFlexGatewayTargetSummary() *FlexGatewayTargetSummary`

NewFlexGatewayTargetSummary instantiates a new FlexGatewayTargetSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlexGatewayTargetSummaryWithDefaults

`func NewFlexGatewayTargetSummaryWithDefaults() *FlexGatewayTargetSummary`

NewFlexGatewayTargetSummaryWithDefaults instantiates a new FlexGatewayTargetSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *FlexGatewayTargetSummary) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *FlexGatewayTargetSummary) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *FlexGatewayTargetSummary) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *FlexGatewayTargetSummary) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetId

`func (o *FlexGatewayTargetSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlexGatewayTargetSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlexGatewayTargetSummary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FlexGatewayTargetSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FlexGatewayTargetSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlexGatewayTargetSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlexGatewayTargetSummary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FlexGatewayTargetSummary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *FlexGatewayTargetSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FlexGatewayTargetSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FlexGatewayTargetSummary) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FlexGatewayTargetSummary) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReplicas

`func (o *FlexGatewayTargetSummary) GetReplicas() FlexGatewayTargetSummaryReplicas`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *FlexGatewayTargetSummary) GetReplicasOk() (*FlexGatewayTargetSummaryReplicas, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *FlexGatewayTargetSummary) SetReplicas(v FlexGatewayTargetSummaryReplicas)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *FlexGatewayTargetSummary) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetTags

`func (o *FlexGatewayTargetSummary) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FlexGatewayTargetSummary) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FlexGatewayTargetSummary) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FlexGatewayTargetSummary) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetLastUpdate

`func (o *FlexGatewayTargetSummary) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *FlexGatewayTargetSummary) GetLastUpdateOk() (*time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *FlexGatewayTargetSummary) SetLastUpdate(v time.Time)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *FlexGatewayTargetSummary) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


