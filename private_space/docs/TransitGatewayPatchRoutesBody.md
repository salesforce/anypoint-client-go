# TransitGatewayPatchRoutesBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Routes** | **[]string** | Replacement set of static routes (CIDR strings). | 

## Methods

### NewTransitGatewayPatchRoutesBody

`func NewTransitGatewayPatchRoutesBody(routes []string, ) *TransitGatewayPatchRoutesBody`

NewTransitGatewayPatchRoutesBody instantiates a new TransitGatewayPatchRoutesBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitGatewayPatchRoutesBodyWithDefaults

`func NewTransitGatewayPatchRoutesBodyWithDefaults() *TransitGatewayPatchRoutesBody`

NewTransitGatewayPatchRoutesBodyWithDefaults instantiates a new TransitGatewayPatchRoutesBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoutes

`func (o *TransitGatewayPatchRoutesBody) GetRoutes() []string`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *TransitGatewayPatchRoutesBody) GetRoutesOk() (*[]string, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *TransitGatewayPatchRoutesBody) SetRoutes(v []string)`

SetRoutes sets Routes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


