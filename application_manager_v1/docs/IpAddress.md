# IpAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegionId** | Pointer to **string** | The region identifier for the IP address. | [optional] 
**RegionName** | Pointer to **string** | The region name for the IP address. | [optional] 
**Type** | Pointer to **string** | The type of IP address (ex. vpc) | [optional] 
**Address** | Pointer to **string** | The IP address. | [optional] 
**State** | Pointer to **string** | The state of the IP address (ex. available, in use). | [optional] 

## Methods

### NewIpAddress

`func NewIpAddress() *IpAddress`

NewIpAddress instantiates a new IpAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpAddressWithDefaults

`func NewIpAddressWithDefaults() *IpAddress`

NewIpAddressWithDefaults instantiates a new IpAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegionId

`func (o *IpAddress) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *IpAddress) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *IpAddress) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *IpAddress) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetRegionName

`func (o *IpAddress) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *IpAddress) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *IpAddress) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *IpAddress) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### GetType

`func (o *IpAddress) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IpAddress) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IpAddress) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IpAddress) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAddress

`func (o *IpAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *IpAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *IpAddress) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *IpAddress) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetState

`func (o *IpAddress) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IpAddress) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IpAddress) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *IpAddress) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


