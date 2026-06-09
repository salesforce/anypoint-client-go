# DeploymentDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** | The domain name associated with the deployment. only present for CH1. | [optional] 
**DeploymentUpdateStatus** | Pointer to **string** | The deployment update status, only present for CH1. | [optional] 

## Methods

### NewDeploymentDetails

`func NewDeploymentDetails() *DeploymentDetails`

NewDeploymentDetails instantiates a new DeploymentDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentDetailsWithDefaults

`func NewDeploymentDetailsWithDefaults() *DeploymentDetails`

NewDeploymentDetailsWithDefaults instantiates a new DeploymentDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *DeploymentDetails) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DeploymentDetails) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DeploymentDetails) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *DeploymentDetails) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetDeploymentUpdateStatus

`func (o *DeploymentDetails) GetDeploymentUpdateStatus() string`

GetDeploymentUpdateStatus returns the DeploymentUpdateStatus field if non-nil, zero value otherwise.

### GetDeploymentUpdateStatusOk

`func (o *DeploymentDetails) GetDeploymentUpdateStatusOk() (*string, bool)`

GetDeploymentUpdateStatusOk returns a tuple with the DeploymentUpdateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentUpdateStatus

`func (o *DeploymentDetails) SetDeploymentUpdateStatus(v string)`

SetDeploymentUpdateStatus sets DeploymentUpdateStatus field to given value.

### HasDeploymentUpdateStatus

`func (o *DeploymentDetails) HasDeploymentUpdateStatus() bool`

HasDeploymentUpdateStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


