# DlbPostBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** | dlb state | [optional] [default to "STOPPED"]
**IpWhitelist** | Pointer to **[]string** |  | [optional] 
**HttpMode** | Pointer to **string** |  | [optional] [default to "redirect"]
**Tlsv1** | Pointer to **bool** |  | [optional] 
**SslEndpoints** | Pointer to [**[]DlbPostBodySslEndpoints**](DlbPostBodySslEndpoints.md) |  | [optional] 

## Methods

### NewDlbPostBody

`func NewDlbPostBody() *DlbPostBody`

NewDlbPostBody instantiates a new DlbPostBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDlbPostBodyWithDefaults

`func NewDlbPostBodyWithDefaults() *DlbPostBody`

NewDlbPostBodyWithDefaults instantiates a new DlbPostBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DlbPostBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DlbPostBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DlbPostBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DlbPostBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *DlbPostBody) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DlbPostBody) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DlbPostBody) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DlbPostBody) HasState() bool`

HasState returns a boolean if a field has been set.

### GetIpWhitelist

`func (o *DlbPostBody) GetIpWhitelist() []string`

GetIpWhitelist returns the IpWhitelist field if non-nil, zero value otherwise.

### GetIpWhitelistOk

`func (o *DlbPostBody) GetIpWhitelistOk() (*[]string, bool)`

GetIpWhitelistOk returns a tuple with the IpWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpWhitelist

`func (o *DlbPostBody) SetIpWhitelist(v []string)`

SetIpWhitelist sets IpWhitelist field to given value.

### HasIpWhitelist

`func (o *DlbPostBody) HasIpWhitelist() bool`

HasIpWhitelist returns a boolean if a field has been set.

### GetHttpMode

`func (o *DlbPostBody) GetHttpMode() string`

GetHttpMode returns the HttpMode field if non-nil, zero value otherwise.

### GetHttpModeOk

`func (o *DlbPostBody) GetHttpModeOk() (*string, bool)`

GetHttpModeOk returns a tuple with the HttpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMode

`func (o *DlbPostBody) SetHttpMode(v string)`

SetHttpMode sets HttpMode field to given value.

### HasHttpMode

`func (o *DlbPostBody) HasHttpMode() bool`

HasHttpMode returns a boolean if a field has been set.

### GetTlsv1

`func (o *DlbPostBody) GetTlsv1() bool`

GetTlsv1 returns the Tlsv1 field if non-nil, zero value otherwise.

### GetTlsv1Ok

`func (o *DlbPostBody) GetTlsv1Ok() (*bool, bool)`

GetTlsv1Ok returns a tuple with the Tlsv1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsv1

`func (o *DlbPostBody) SetTlsv1(v bool)`

SetTlsv1 sets Tlsv1 field to given value.

### HasTlsv1

`func (o *DlbPostBody) HasTlsv1() bool`

HasTlsv1 returns a boolean if a field has been set.

### GetSslEndpoints

`func (o *DlbPostBody) GetSslEndpoints() []DlbPostBodySslEndpoints`

GetSslEndpoints returns the SslEndpoints field if non-nil, zero value otherwise.

### GetSslEndpointsOk

`func (o *DlbPostBody) GetSslEndpointsOk() (*[]DlbPostBodySslEndpoints, bool)`

GetSslEndpointsOk returns a tuple with the SslEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslEndpoints

`func (o *DlbPostBody) SetSslEndpoints(v []DlbPostBodySslEndpoints)`

SetSslEndpoints sets SslEndpoints field to given value.

### HasSslEndpoints

`func (o *DlbPostBody) HasSslEndpoints() bool`

HasSslEndpoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


