# IngressConfigurationDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | The status of the deployment configuration. | [optional] 
**LastSeenTimestamp** | Pointer to **int64** | The last seen timestamp of the deployment configuration. | [optional] 

## Methods

### NewIngressConfigurationDeployment

`func NewIngressConfigurationDeployment() *IngressConfigurationDeployment`

NewIngressConfigurationDeployment instantiates a new IngressConfigurationDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngressConfigurationDeploymentWithDefaults

`func NewIngressConfigurationDeploymentWithDefaults() *IngressConfigurationDeployment`

NewIngressConfigurationDeploymentWithDefaults instantiates a new IngressConfigurationDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *IngressConfigurationDeployment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IngressConfigurationDeployment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IngressConfigurationDeployment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IngressConfigurationDeployment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLastSeenTimestamp

`func (o *IngressConfigurationDeployment) GetLastSeenTimestamp() int64`

GetLastSeenTimestamp returns the LastSeenTimestamp field if non-nil, zero value otherwise.

### GetLastSeenTimestampOk

`func (o *IngressConfigurationDeployment) GetLastSeenTimestampOk() (*int64, bool)`

GetLastSeenTimestampOk returns a tuple with the LastSeenTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenTimestamp

`func (o *IngressConfigurationDeployment) SetLastSeenTimestamp(v int64)`

SetLastSeenTimestamp sets LastSeenTimestamp field to given value.

### HasLastSeenTimestamp

`func (o *IngressConfigurationDeployment) HasLastSeenTimestamp() bool`

HasLastSeenTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


