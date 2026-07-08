# PostLLMAssetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicationStatusLink** | Pointer to **string** | URL pointing to the publication status of the newly created asset. Exchange publish is synchronous from the caller&#39;s perspective; the asset is queryable immediately. This link is the resource location, not an async completion handle. | [optional] 

## Methods

### NewPostLLMAssetResponse

`func NewPostLLMAssetResponse() *PostLLMAssetResponse`

NewPostLLMAssetResponse instantiates a new PostLLMAssetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostLLMAssetResponseWithDefaults

`func NewPostLLMAssetResponseWithDefaults() *PostLLMAssetResponse`

NewPostLLMAssetResponseWithDefaults instantiates a new PostLLMAssetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicationStatusLink

`func (o *PostLLMAssetResponse) GetPublicationStatusLink() string`

GetPublicationStatusLink returns the PublicationStatusLink field if non-nil, zero value otherwise.

### GetPublicationStatusLinkOk

`func (o *PostLLMAssetResponse) GetPublicationStatusLinkOk() (*string, bool)`

GetPublicationStatusLinkOk returns a tuple with the PublicationStatusLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationStatusLink

`func (o *PostLLMAssetResponse) SetPublicationStatusLink(v string)`

SetPublicationStatusLink sets PublicationStatusLink field to given value.

### HasPublicationStatusLink

`func (o *PostLLMAssetResponse) HasPublicationStatusLink() bool`

HasPublicationStatusLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


