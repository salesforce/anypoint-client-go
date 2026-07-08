# UpdateDeploymentBulkBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | * UPDATE: Update all applications in the provided list to the latest update of the Mule Version that the application is currently using * START: Start all applications in the provided list if they are stopped * STOP: Stop all applications in the provided list if they are started * RESTART: Restart all applications in the provided list with zero downtime * DELETE: Delete all applications in the provided list  | [optional] 
**Domains** | Pointer to **[]string** | The list of application domains | [optional] 

## Methods

### NewUpdateDeploymentBulkBody

`func NewUpdateDeploymentBulkBody() *UpdateDeploymentBulkBody`

NewUpdateDeploymentBulkBody instantiates a new UpdateDeploymentBulkBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeploymentBulkBodyWithDefaults

`func NewUpdateDeploymentBulkBodyWithDefaults() *UpdateDeploymentBulkBody`

NewUpdateDeploymentBulkBodyWithDefaults instantiates a new UpdateDeploymentBulkBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *UpdateDeploymentBulkBody) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UpdateDeploymentBulkBody) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UpdateDeploymentBulkBody) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *UpdateDeploymentBulkBody) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetDomains

`func (o *UpdateDeploymentBulkBody) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *UpdateDeploymentBulkBody) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *UpdateDeploymentBulkBody) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *UpdateDeploymentBulkBody) HasDomains() bool`

HasDomains returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


