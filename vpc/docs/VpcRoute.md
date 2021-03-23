# VpcRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CIDR** | **string** |  | [default to ""]
**NextHop** | **string** |  | [default to ""]

## Methods

### NewVpcRoute

`func NewVpcRoute(cIDR string, nextHop string, ) *VpcRoute`

NewVpcRoute instantiates a new VpcRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpcRouteWithDefaults

`func NewVpcRouteWithDefaults() *VpcRoute`

NewVpcRouteWithDefaults instantiates a new VpcRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCIDR

`func (o *VpcRoute) GetCIDR() string`

GetCIDR returns the CIDR field if non-nil, zero value otherwise.

### GetCIDROk

`func (o *VpcRoute) GetCIDROk() (*string, bool)`

GetCIDROk returns a tuple with the CIDR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCIDR

`func (o *VpcRoute) SetCIDR(v string)`

SetCIDR sets CIDR field to given value.


### GetNextHop

`func (o *VpcRoute) GetNextHop() string`

GetNextHop returns the NextHop field if non-nil, zero value otherwise.

### GetNextHopOk

`func (o *VpcRoute) GetNextHopOk() (*string, bool)`

GetNextHopOk returns a tuple with the NextHop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHop

`func (o *VpcRoute) SetNextHop(v string)`

SetNextHop sets NextHop field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


