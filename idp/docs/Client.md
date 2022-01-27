# Client

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Urls** | Pointer to [**Urls2**](Urls2.md) |  | [optional] 
**Credentials** | Pointer to [**Credentials**](Credentials.md) |  | [optional] 

## Methods

### NewClient

`func NewClient() *Client`

NewClient instantiates a new Client object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientWithDefaults

`func NewClientWithDefaults() *Client`

NewClientWithDefaults instantiates a new Client object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrls

`func (o *Client) GetUrls() Urls2`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *Client) GetUrlsOk() (*Urls2, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *Client) SetUrls(v Urls2)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *Client) HasUrls() bool`

HasUrls returns a boolean if a field has been set.

### GetCredentials

`func (o *Client) GetCredentials() Credentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *Client) GetCredentialsOk() (*Credentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *Client) SetCredentials(v Credentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *Client) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


