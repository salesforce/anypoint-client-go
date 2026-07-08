# TransitGatewayStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gateway** | Pointer to **string** | Gateway status (e.g. refreshing, available). | [optional] 
**Attachment** | Pointer to **string** | Attachment status (e.g. refreshing, available). | [optional] 
**TgwResource** | Pointer to **string** | AWS console URL pointing at the transit gateway resource. | [optional] 
**Routes** | Pointer to **[]string** | Static routes pushed into the private space route table. | [optional] 

## Methods

### NewTransitGatewayStatus

`func NewTransitGatewayStatus() *TransitGatewayStatus`

NewTransitGatewayStatus instantiates a new TransitGatewayStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitGatewayStatusWithDefaults

`func NewTransitGatewayStatusWithDefaults() *TransitGatewayStatus`

NewTransitGatewayStatusWithDefaults instantiates a new TransitGatewayStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGateway

`func (o *TransitGatewayStatus) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *TransitGatewayStatus) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *TransitGatewayStatus) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *TransitGatewayStatus) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetAttachment

`func (o *TransitGatewayStatus) GetAttachment() string`

GetAttachment returns the Attachment field if non-nil, zero value otherwise.

### GetAttachmentOk

`func (o *TransitGatewayStatus) GetAttachmentOk() (*string, bool)`

GetAttachmentOk returns a tuple with the Attachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachment

`func (o *TransitGatewayStatus) SetAttachment(v string)`

SetAttachment sets Attachment field to given value.

### HasAttachment

`func (o *TransitGatewayStatus) HasAttachment() bool`

HasAttachment returns a boolean if a field has been set.

### GetTgwResource

`func (o *TransitGatewayStatus) GetTgwResource() string`

GetTgwResource returns the TgwResource field if non-nil, zero value otherwise.

### GetTgwResourceOk

`func (o *TransitGatewayStatus) GetTgwResourceOk() (*string, bool)`

GetTgwResourceOk returns a tuple with the TgwResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgwResource

`func (o *TransitGatewayStatus) SetTgwResource(v string)`

SetTgwResource sets TgwResource field to given value.

### HasTgwResource

`func (o *TransitGatewayStatus) HasTgwResource() bool`

HasTgwResource returns a boolean if a field has been set.

### GetRoutes

`func (o *TransitGatewayStatus) GetRoutes() []string`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *TransitGatewayStatus) GetRoutesOk() (*[]string, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *TransitGatewayStatus) SetRoutes(v []string)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *TransitGatewayStatus) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


