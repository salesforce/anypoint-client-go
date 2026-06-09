# Audit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to [**AuditDate**](AuditDate.md) |  | [optional] 
**Updated** | Pointer to [**AuditDate**](AuditDate.md) |  | [optional] 

## Methods

### NewAudit

`func NewAudit() *Audit`

NewAudit instantiates a new Audit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditWithDefaults

`func NewAuditWithDefaults() *Audit`

NewAuditWithDefaults instantiates a new Audit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *Audit) GetCreated() AuditDate`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Audit) GetCreatedOk() (*AuditDate, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Audit) SetCreated(v AuditDate)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Audit) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *Audit) GetUpdated() AuditDate`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Audit) GetUpdatedOk() (*AuditDate, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Audit) SetUpdated(v AuditDate)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *Audit) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


