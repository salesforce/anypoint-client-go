# UpstreamCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**Upstreams** | Pointer to [**[]UpstreamDetails**](UpstreamDetails.md) |  | [optional] 

## Methods

### NewUpstreamCollection

`func NewUpstreamCollection() *UpstreamCollection`

NewUpstreamCollection instantiates a new UpstreamCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpstreamCollectionWithDefaults

`func NewUpstreamCollectionWithDefaults() *UpstreamCollection`

NewUpstreamCollectionWithDefaults instantiates a new UpstreamCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *UpstreamCollection) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *UpstreamCollection) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *UpstreamCollection) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *UpstreamCollection) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUpstreams

`func (o *UpstreamCollection) GetUpstreams() []UpstreamDetails`

GetUpstreams returns the Upstreams field if non-nil, zero value otherwise.

### GetUpstreamsOk

`func (o *UpstreamCollection) GetUpstreamsOk() (*[]UpstreamDetails, bool)`

GetUpstreamsOk returns a tuple with the Upstreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreams

`func (o *UpstreamCollection) SetUpstreams(v []UpstreamDetails)`

SetUpstreams sets Upstreams field to given value.

### HasUpstreams

`func (o *UpstreamCollection) HasUpstreams() bool`

HasUpstreams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


