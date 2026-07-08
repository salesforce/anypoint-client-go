# ContractDetailsApi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audit** | Pointer to [**Audit**](Audit.md) |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 

## Methods

### NewContractDetailsApi

`func NewContractDetailsApi() *ContractDetailsApi`

NewContractDetailsApi instantiates a new ContractDetailsApi object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractDetailsApiWithDefaults

`func NewContractDetailsApiWithDefaults() *ContractDetailsApi`

NewContractDetailsApiWithDefaults instantiates a new ContractDetailsApi object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudit

`func (o *ContractDetailsApi) GetAudit() Audit`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *ContractDetailsApi) GetAuditOk() (*Audit, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *ContractDetailsApi) SetAudit(v Audit)`

SetAudit sets Audit field to given value.

### HasAudit

`func (o *ContractDetailsApi) HasAudit() bool`

HasAudit returns a boolean if a field has been set.

### GetOrganizationId

`func (o *ContractDetailsApi) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ContractDetailsApi) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ContractDetailsApi) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ContractDetailsApi) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetId

`func (o *ContractDetailsApi) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContractDetailsApi) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContractDetailsApi) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ContractDetailsApi) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


