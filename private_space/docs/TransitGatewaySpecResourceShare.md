# TransitGatewaySpecResourceShare

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | RAM resource share UUID (trailing GUID of the ARN). | [optional] 
**Account** | Pointer to **string** | AWS account id that owns the RAM share. | [optional] 

## Methods

### NewTransitGatewaySpecResourceShare

`func NewTransitGatewaySpecResourceShare() *TransitGatewaySpecResourceShare`

NewTransitGatewaySpecResourceShare instantiates a new TransitGatewaySpecResourceShare object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitGatewaySpecResourceShareWithDefaults

`func NewTransitGatewaySpecResourceShareWithDefaults() *TransitGatewaySpecResourceShare`

NewTransitGatewaySpecResourceShareWithDefaults instantiates a new TransitGatewaySpecResourceShare object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TransitGatewaySpecResourceShare) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransitGatewaySpecResourceShare) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransitGatewaySpecResourceShare) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TransitGatewaySpecResourceShare) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *TransitGatewaySpecResourceShare) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *TransitGatewaySpecResourceShare) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *TransitGatewaySpecResourceShare) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *TransitGatewaySpecResourceShare) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


