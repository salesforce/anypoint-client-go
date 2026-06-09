# GlobalSpaceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Space** | Pointer to [**PrivateSpaceComponentStatus**](PrivateSpaceComponentStatus.md) |  | [optional] 
**Cluster** | Pointer to [**[]GlobalSpaceStatusClusterInner**](GlobalSpaceStatusClusterInner.md) | The status of the space cluster component. | [optional] 
**Network** | Pointer to [**PrivateSpaceComponentStatus**](PrivateSpaceComponentStatus.md) |  | [optional] 

## Methods

### NewGlobalSpaceStatus

`func NewGlobalSpaceStatus() *GlobalSpaceStatus`

NewGlobalSpaceStatus instantiates a new GlobalSpaceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalSpaceStatusWithDefaults

`func NewGlobalSpaceStatusWithDefaults() *GlobalSpaceStatus`

NewGlobalSpaceStatusWithDefaults instantiates a new GlobalSpaceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpace

`func (o *GlobalSpaceStatus) GetSpace() PrivateSpaceComponentStatus`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *GlobalSpaceStatus) GetSpaceOk() (*PrivateSpaceComponentStatus, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *GlobalSpaceStatus) SetSpace(v PrivateSpaceComponentStatus)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *GlobalSpaceStatus) HasSpace() bool`

HasSpace returns a boolean if a field has been set.

### GetCluster

`func (o *GlobalSpaceStatus) GetCluster() []GlobalSpaceStatusClusterInner`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *GlobalSpaceStatus) GetClusterOk() (*[]GlobalSpaceStatusClusterInner, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *GlobalSpaceStatus) SetCluster(v []GlobalSpaceStatusClusterInner)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *GlobalSpaceStatus) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetNetwork

`func (o *GlobalSpaceStatus) GetNetwork() PrivateSpaceComponentStatus`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GlobalSpaceStatus) GetNetworkOk() (*PrivateSpaceComponentStatus, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GlobalSpaceStatus) SetNetwork(v PrivateSpaceComponentStatus)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GlobalSpaceStatus) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


