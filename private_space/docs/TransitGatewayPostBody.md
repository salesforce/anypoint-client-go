# TransitGatewayPostBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name for the transit gateway attachment. | 
**ResourceShareId** | **string** | RAM resource share UUID (trailing GUID of the ARN, NOT the full ARN). | 
**ResourceShareAccount** | **string** | AWS account id that owns the RAM share (your AWS account, NOT MuleSoft&#39;s account). | 
**Routes** | **[]string** | Static routes (CIDR strings) to push into the private space route table. Must not equal or be more specific than the private space CIDR. | 

## Methods

### NewTransitGatewayPostBody

`func NewTransitGatewayPostBody(name string, resourceShareId string, resourceShareAccount string, routes []string, ) *TransitGatewayPostBody`

NewTransitGatewayPostBody instantiates a new TransitGatewayPostBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitGatewayPostBodyWithDefaults

`func NewTransitGatewayPostBodyWithDefaults() *TransitGatewayPostBody`

NewTransitGatewayPostBodyWithDefaults instantiates a new TransitGatewayPostBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TransitGatewayPostBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TransitGatewayPostBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TransitGatewayPostBody) SetName(v string)`

SetName sets Name field to given value.


### GetResourceShareId

`func (o *TransitGatewayPostBody) GetResourceShareId() string`

GetResourceShareId returns the ResourceShareId field if non-nil, zero value otherwise.

### GetResourceShareIdOk

`func (o *TransitGatewayPostBody) GetResourceShareIdOk() (*string, bool)`

GetResourceShareIdOk returns a tuple with the ResourceShareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceShareId

`func (o *TransitGatewayPostBody) SetResourceShareId(v string)`

SetResourceShareId sets ResourceShareId field to given value.


### GetResourceShareAccount

`func (o *TransitGatewayPostBody) GetResourceShareAccount() string`

GetResourceShareAccount returns the ResourceShareAccount field if non-nil, zero value otherwise.

### GetResourceShareAccountOk

`func (o *TransitGatewayPostBody) GetResourceShareAccountOk() (*string, bool)`

GetResourceShareAccountOk returns a tuple with the ResourceShareAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceShareAccount

`func (o *TransitGatewayPostBody) SetResourceShareAccount(v string)`

SetResourceShareAccount sets ResourceShareAccount field to given value.


### GetRoutes

`func (o *TransitGatewayPostBody) GetRoutes() []string`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *TransitGatewayPostBody) GetRoutesOk() (*[]string, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *TransitGatewayPostBody) SetRoutes(v []string)`

SetRoutes sets Routes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


