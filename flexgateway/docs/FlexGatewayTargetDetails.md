# FlexGatewayTargetDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Replicas** | Pointer to [**[]FlexGatewayTargetDetailsReplicas**](FlexGatewayTargetDetailsReplicas.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**LastUpdate** | Pointer to **time.Time** |  | [optional] 
**Versions** | Pointer to **[]string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewFlexGatewayTargetDetails

`func NewFlexGatewayTargetDetails() *FlexGatewayTargetDetails`

NewFlexGatewayTargetDetails instantiates a new FlexGatewayTargetDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlexGatewayTargetDetailsWithDefaults

`func NewFlexGatewayTargetDetailsWithDefaults() *FlexGatewayTargetDetails`

NewFlexGatewayTargetDetailsWithDefaults instantiates a new FlexGatewayTargetDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *FlexGatewayTargetDetails) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *FlexGatewayTargetDetails) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *FlexGatewayTargetDetails) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *FlexGatewayTargetDetails) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetId

`func (o *FlexGatewayTargetDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlexGatewayTargetDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlexGatewayTargetDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FlexGatewayTargetDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FlexGatewayTargetDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlexGatewayTargetDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlexGatewayTargetDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FlexGatewayTargetDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *FlexGatewayTargetDetails) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FlexGatewayTargetDetails) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FlexGatewayTargetDetails) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FlexGatewayTargetDetails) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReplicas

`func (o *FlexGatewayTargetDetails) GetReplicas() []FlexGatewayTargetDetailsReplicas`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *FlexGatewayTargetDetails) GetReplicasOk() (*[]FlexGatewayTargetDetailsReplicas, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *FlexGatewayTargetDetails) SetReplicas(v []FlexGatewayTargetDetailsReplicas)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *FlexGatewayTargetDetails) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetTags

`func (o *FlexGatewayTargetDetails) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FlexGatewayTargetDetails) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FlexGatewayTargetDetails) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FlexGatewayTargetDetails) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetLastUpdate

`func (o *FlexGatewayTargetDetails) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *FlexGatewayTargetDetails) GetLastUpdateOk() (*time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *FlexGatewayTargetDetails) SetLastUpdate(v time.Time)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *FlexGatewayTargetDetails) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.

### GetVersions

`func (o *FlexGatewayTargetDetails) GetVersions() []string`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *FlexGatewayTargetDetails) GetVersionsOk() (*[]string, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *FlexGatewayTargetDetails) SetVersions(v []string)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *FlexGatewayTargetDetails) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetVersion

`func (o *FlexGatewayTargetDetails) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FlexGatewayTargetDetails) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FlexGatewayTargetDetails) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *FlexGatewayTargetDetails) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


