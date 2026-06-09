# PrivateSpaceIamRoles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Roles** | Pointer to **[]string** |  | [optional] 
**OrganizationId** | Pointer to **string** | The ID of the organization in GUID format. | [optional] 
**SpaceId** | Pointer to **string** | The ID of the private space in GUID format. | [optional] 

## Methods

### NewPrivateSpaceIamRoles

`func NewPrivateSpaceIamRoles() *PrivateSpaceIamRoles`

NewPrivateSpaceIamRoles instantiates a new PrivateSpaceIamRoles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateSpaceIamRolesWithDefaults

`func NewPrivateSpaceIamRolesWithDefaults() *PrivateSpaceIamRoles`

NewPrivateSpaceIamRolesWithDefaults instantiates a new PrivateSpaceIamRoles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoles

`func (o *PrivateSpaceIamRoles) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *PrivateSpaceIamRoles) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *PrivateSpaceIamRoles) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *PrivateSpaceIamRoles) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetOrganizationId

`func (o *PrivateSpaceIamRoles) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *PrivateSpaceIamRoles) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *PrivateSpaceIamRoles) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *PrivateSpaceIamRoles) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetSpaceId

`func (o *PrivateSpaceIamRoles) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *PrivateSpaceIamRoles) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *PrivateSpaceIamRoles) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *PrivateSpaceIamRoles) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


