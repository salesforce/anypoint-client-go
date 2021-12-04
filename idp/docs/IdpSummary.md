# IdpSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderId** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ModelType**](ModelType.md) |  | [optional] 

## Methods

### NewIdpSummary

`func NewIdpSummary() *IdpSummary`

NewIdpSummary instantiates a new IdpSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpSummaryWithDefaults

`func NewIdpSummaryWithDefaults() *IdpSummary`

NewIdpSummaryWithDefaults instantiates a new IdpSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderId

`func (o *IdpSummary) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *IdpSummary) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *IdpSummary) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *IdpSummary) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetOrgId

`func (o *IdpSummary) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *IdpSummary) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *IdpSummary) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *IdpSummary) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetName

`func (o *IdpSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdpSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdpSummary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdpSummary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *IdpSummary) GetType() ModelType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IdpSummary) GetTypeOk() (*ModelType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IdpSummary) SetType(v ModelType)`

SetType sets Type field to given value.

### HasType

`func (o *IdpSummary) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


