# IngressConfigurationLogs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to [**[]Filter**](Filter.md) | The filters of the logs configuration. | [optional] 
**PortLogLevel** | Pointer to **string** | The port log level of the ingress configuration. | [optional] 

## Methods

### NewIngressConfigurationLogs

`func NewIngressConfigurationLogs() *IngressConfigurationLogs`

NewIngressConfigurationLogs instantiates a new IngressConfigurationLogs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngressConfigurationLogsWithDefaults

`func NewIngressConfigurationLogsWithDefaults() *IngressConfigurationLogs`

NewIngressConfigurationLogsWithDefaults instantiates a new IngressConfigurationLogs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *IngressConfigurationLogs) GetFilters() []Filter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *IngressConfigurationLogs) GetFiltersOk() (*[]Filter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *IngressConfigurationLogs) SetFilters(v []Filter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *IngressConfigurationLogs) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetPortLogLevel

`func (o *IngressConfigurationLogs) GetPortLogLevel() string`

GetPortLogLevel returns the PortLogLevel field if non-nil, zero value otherwise.

### GetPortLogLevelOk

`func (o *IngressConfigurationLogs) GetPortLogLevelOk() (*string, bool)`

GetPortLogLevelOk returns a tuple with the PortLogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortLogLevel

`func (o *IngressConfigurationLogs) SetPortLogLevel(v string)`

SetPortLogLevel sets PortLogLevel field to given value.

### HasPortLogLevel

`func (o *IngressConfigurationLogs) HasPortLogLevel() bool`

HasPortLogLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


