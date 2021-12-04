# File

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Classifier** | Pointer to **string** |  | [optional] 
**Packaging** | Pointer to **string** |  | [optional] 
**ExternalLink** | Pointer to **string** |  | [optional] 
**DownloadURL** | Pointer to **string** |  | [optional] 
**Md5** | Pointer to **string** |  | [optional] 
**Sha1** | Pointer to **string** |  | [optional] 
**CreatedDate** | Pointer to **string** |  | [optional] 
**MainFile** | Pointer to **NullableString** |  | [optional] 
**IsGenerated** | Pointer to **bool** |  | [optional] 

## Methods

### NewFile

`func NewFile() *File`

NewFile instantiates a new File object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileWithDefaults

`func NewFileWithDefaults() *File`

NewFileWithDefaults instantiates a new File object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassifier

`func (o *File) GetClassifier() string`

GetClassifier returns the Classifier field if non-nil, zero value otherwise.

### GetClassifierOk

`func (o *File) GetClassifierOk() (*string, bool)`

GetClassifierOk returns a tuple with the Classifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifier

`func (o *File) SetClassifier(v string)`

SetClassifier sets Classifier field to given value.

### HasClassifier

`func (o *File) HasClassifier() bool`

HasClassifier returns a boolean if a field has been set.

### GetPackaging

`func (o *File) GetPackaging() string`

GetPackaging returns the Packaging field if non-nil, zero value otherwise.

### GetPackagingOk

`func (o *File) GetPackagingOk() (*string, bool)`

GetPackagingOk returns a tuple with the Packaging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackaging

`func (o *File) SetPackaging(v string)`

SetPackaging sets Packaging field to given value.

### HasPackaging

`func (o *File) HasPackaging() bool`

HasPackaging returns a boolean if a field has been set.

### GetExternalLink

`func (o *File) GetExternalLink() string`

GetExternalLink returns the ExternalLink field if non-nil, zero value otherwise.

### GetExternalLinkOk

`func (o *File) GetExternalLinkOk() (*string, bool)`

GetExternalLinkOk returns a tuple with the ExternalLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLink

`func (o *File) SetExternalLink(v string)`

SetExternalLink sets ExternalLink field to given value.

### HasExternalLink

`func (o *File) HasExternalLink() bool`

HasExternalLink returns a boolean if a field has been set.

### GetDownloadURL

`func (o *File) GetDownloadURL() string`

GetDownloadURL returns the DownloadURL field if non-nil, zero value otherwise.

### GetDownloadURLOk

`func (o *File) GetDownloadURLOk() (*string, bool)`

GetDownloadURLOk returns a tuple with the DownloadURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadURL

`func (o *File) SetDownloadURL(v string)`

SetDownloadURL sets DownloadURL field to given value.

### HasDownloadURL

`func (o *File) HasDownloadURL() bool`

HasDownloadURL returns a boolean if a field has been set.

### GetMd5

`func (o *File) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *File) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *File) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *File) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### GetSha1

`func (o *File) GetSha1() string`

GetSha1 returns the Sha1 field if non-nil, zero value otherwise.

### GetSha1Ok

`func (o *File) GetSha1Ok() (*string, bool)`

GetSha1Ok returns a tuple with the Sha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha1

`func (o *File) SetSha1(v string)`

SetSha1 sets Sha1 field to given value.

### HasSha1

`func (o *File) HasSha1() bool`

HasSha1 returns a boolean if a field has been set.

### GetCreatedDate

`func (o *File) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *File) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *File) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *File) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetMainFile

`func (o *File) GetMainFile() string`

GetMainFile returns the MainFile field if non-nil, zero value otherwise.

### GetMainFileOk

`func (o *File) GetMainFileOk() (*string, bool)`

GetMainFileOk returns a tuple with the MainFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainFile

`func (o *File) SetMainFile(v string)`

SetMainFile sets MainFile field to given value.

### HasMainFile

`func (o *File) HasMainFile() bool`

HasMainFile returns a boolean if a field has been set.

### SetMainFileNil

`func (o *File) SetMainFileNil(b bool)`

 SetMainFileNil sets the value for MainFile to be an explicit nil

### UnsetMainFile
`func (o *File) UnsetMainFile()`

UnsetMainFile ensures that no value is present for MainFile, not even an explicit nil
### GetIsGenerated

`func (o *File) GetIsGenerated() bool`

GetIsGenerated returns the IsGenerated field if non-nil, zero value otherwise.

### GetIsGeneratedOk

`func (o *File) GetIsGeneratedOk() (*bool, bool)`

GetIsGeneratedOk returns a tuple with the IsGenerated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGenerated

`func (o *File) SetIsGenerated(v bool)`

SetIsGenerated sets IsGenerated field to given value.

### HasIsGenerated

`func (o *File) HasIsGenerated() bool`

HasIsGenerated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


