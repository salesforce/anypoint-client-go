# Alert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ApiAlertsVersion** | **string** |  | 
**Name** | **string** |  | 
**Type** | **string** |  | 
**Enabled** | **bool** |  | 
**Severity** | **string** |  | 
**Recipients** | [**[]Recipient**](Recipient.md) |  | 
**Condition** | [**Condition**](Condition.md) |  | 
**Period** | [**Period**](Period.md) |  | 

## Methods

### NewAlert

`func NewAlert(id string, apiAlertsVersion string, name string, type_ string, enabled bool, severity string, recipients []Recipient, condition Condition, period Period, ) *Alert`

NewAlert instantiates a new Alert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertWithDefaults

`func NewAlertWithDefaults() *Alert`

NewAlertWithDefaults instantiates a new Alert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Alert) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Alert) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Alert) SetId(v string)`

SetId sets Id field to given value.


### GetApiAlertsVersion

`func (o *Alert) GetApiAlertsVersion() string`

GetApiAlertsVersion returns the ApiAlertsVersion field if non-nil, zero value otherwise.

### GetApiAlertsVersionOk

`func (o *Alert) GetApiAlertsVersionOk() (*string, bool)`

GetApiAlertsVersionOk returns a tuple with the ApiAlertsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiAlertsVersion

`func (o *Alert) SetApiAlertsVersion(v string)`

SetApiAlertsVersion sets ApiAlertsVersion field to given value.


### GetName

`func (o *Alert) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Alert) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Alert) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *Alert) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Alert) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Alert) SetType(v string)`

SetType sets Type field to given value.


### GetEnabled

`func (o *Alert) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Alert) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Alert) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSeverity

`func (o *Alert) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *Alert) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *Alert) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetRecipients

`func (o *Alert) GetRecipients() []Recipient`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *Alert) GetRecipientsOk() (*[]Recipient, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *Alert) SetRecipients(v []Recipient)`

SetRecipients sets Recipients field to given value.


### GetCondition

`func (o *Alert) GetCondition() Condition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *Alert) GetConditionOk() (*Condition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *Alert) SetCondition(v Condition)`

SetCondition sets Condition field to given value.


### GetPeriod

`func (o *Alert) GetPeriod() Period`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *Alert) GetPeriodOk() (*Period, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *Alert) SetPeriod(v Period)`

SetPeriod sets Period field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


