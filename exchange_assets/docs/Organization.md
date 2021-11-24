# Organization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ParentOrganizationIds** | Pointer to **[]string** |  | [optional] 
**SubOrganizationIds** | Pointer to **[]string** |  | [optional] 
**TenantOrganizationIds** | Pointer to **[]string** |  | [optional] 
**IsMaster** | Pointer to **bool** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**IsMulesoftOrganization** | Pointer to **bool** |  | [optional] 

## Methods

### NewOrganization

`func NewOrganization() *Organization`

NewOrganization instantiates a new Organization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationWithDefaults

`func NewOrganizationWithDefaults() *Organization`

NewOrganizationWithDefaults instantiates a new Organization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Organization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Organization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Organization) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Organization) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Organization) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Organization) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Organization) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Organization) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentOrganizationIds

`func (o *Organization) GetParentOrganizationIds() []string`

GetParentOrganizationIds returns the ParentOrganizationIds field if non-nil, zero value otherwise.

### GetParentOrganizationIdsOk

`func (o *Organization) GetParentOrganizationIdsOk() (*[]string, bool)`

GetParentOrganizationIdsOk returns a tuple with the ParentOrganizationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentOrganizationIds

`func (o *Organization) SetParentOrganizationIds(v []string)`

SetParentOrganizationIds sets ParentOrganizationIds field to given value.

### HasParentOrganizationIds

`func (o *Organization) HasParentOrganizationIds() bool`

HasParentOrganizationIds returns a boolean if a field has been set.

### GetSubOrganizationIds

`func (o *Organization) GetSubOrganizationIds() []string`

GetSubOrganizationIds returns the SubOrganizationIds field if non-nil, zero value otherwise.

### GetSubOrganizationIdsOk

`func (o *Organization) GetSubOrganizationIdsOk() (*[]string, bool)`

GetSubOrganizationIdsOk returns a tuple with the SubOrganizationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubOrganizationIds

`func (o *Organization) SetSubOrganizationIds(v []string)`

SetSubOrganizationIds sets SubOrganizationIds field to given value.

### HasSubOrganizationIds

`func (o *Organization) HasSubOrganizationIds() bool`

HasSubOrganizationIds returns a boolean if a field has been set.

### GetTenantOrganizationIds

`func (o *Organization) GetTenantOrganizationIds() []string`

GetTenantOrganizationIds returns the TenantOrganizationIds field if non-nil, zero value otherwise.

### GetTenantOrganizationIdsOk

`func (o *Organization) GetTenantOrganizationIdsOk() (*[]string, bool)`

GetTenantOrganizationIdsOk returns a tuple with the TenantOrganizationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantOrganizationIds

`func (o *Organization) SetTenantOrganizationIds(v []string)`

SetTenantOrganizationIds sets TenantOrganizationIds field to given value.

### HasTenantOrganizationIds

`func (o *Organization) HasTenantOrganizationIds() bool`

HasTenantOrganizationIds returns a boolean if a field has been set.

### GetIsMaster

`func (o *Organization) GetIsMaster() bool`

GetIsMaster returns the IsMaster field if non-nil, zero value otherwise.

### GetIsMasterOk

`func (o *Organization) GetIsMasterOk() (*bool, bool)`

GetIsMasterOk returns a tuple with the IsMaster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMaster

`func (o *Organization) SetIsMaster(v bool)`

SetIsMaster sets IsMaster field to given value.

### HasIsMaster

`func (o *Organization) HasIsMaster() bool`

HasIsMaster returns a boolean if a field has been set.

### GetDomain

`func (o *Organization) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Organization) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Organization) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *Organization) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetIsMulesoftOrganization

`func (o *Organization) GetIsMulesoftOrganization() bool`

GetIsMulesoftOrganization returns the IsMulesoftOrganization field if non-nil, zero value otherwise.

### GetIsMulesoftOrganizationOk

`func (o *Organization) GetIsMulesoftOrganizationOk() (*bool, bool)`

GetIsMulesoftOrganizationOk returns a tuple with the IsMulesoftOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMulesoftOrganization

`func (o *Organization) SetIsMulesoftOrganization(v bool)`

SetIsMulesoftOrganization sets IsMulesoftOrganization field to given value.

### HasIsMulesoftOrganization

`func (o *Organization) HasIsMulesoftOrganization() bool`

HasIsMulesoftOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


