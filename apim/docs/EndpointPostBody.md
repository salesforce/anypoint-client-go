# EndpointPostBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentType** | Pointer to **string** |  | [optional] 
**MuleVersion4OrAbove** | Pointer to **bool** |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**IsCloudHub** | Pointer to **NullableBool** |  | [optional] 
**ProxyUri** | Pointer to **NullableString** |  | [optional] 
**ProxyRegistrationUri** | Pointer to **NullableString** |  | [optional] 
**ReferencesUserDomain** | Pointer to **NullableString** |  | [optional] 
**ResponseTimeout** | Pointer to **NullableString** |  | [optional] 
**TlsContexts** | Pointer to [**NullableEndpointPostBodyTlsContexts**](EndpointPostBodyTlsContexts.md) |  | [optional] 

## Methods

### NewEndpointPostBody

`func NewEndpointPostBody() *EndpointPostBody`

NewEndpointPostBody instantiates a new EndpointPostBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointPostBodyWithDefaults

`func NewEndpointPostBodyWithDefaults() *EndpointPostBody`

NewEndpointPostBodyWithDefaults instantiates a new EndpointPostBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentType

`func (o *EndpointPostBody) GetDeploymentType() string`

GetDeploymentType returns the DeploymentType field if non-nil, zero value otherwise.

### GetDeploymentTypeOk

`func (o *EndpointPostBody) GetDeploymentTypeOk() (*string, bool)`

GetDeploymentTypeOk returns a tuple with the DeploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentType

`func (o *EndpointPostBody) SetDeploymentType(v string)`

SetDeploymentType sets DeploymentType field to given value.

### HasDeploymentType

`func (o *EndpointPostBody) HasDeploymentType() bool`

HasDeploymentType returns a boolean if a field has been set.

### GetMuleVersion4OrAbove

`func (o *EndpointPostBody) GetMuleVersion4OrAbove() bool`

GetMuleVersion4OrAbove returns the MuleVersion4OrAbove field if non-nil, zero value otherwise.

### GetMuleVersion4OrAboveOk

`func (o *EndpointPostBody) GetMuleVersion4OrAboveOk() (*bool, bool)`

GetMuleVersion4OrAboveOk returns a tuple with the MuleVersion4OrAbove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuleVersion4OrAbove

`func (o *EndpointPostBody) SetMuleVersion4OrAbove(v bool)`

SetMuleVersion4OrAbove sets MuleVersion4OrAbove field to given value.

### HasMuleVersion4OrAbove

`func (o *EndpointPostBody) HasMuleVersion4OrAbove() bool`

HasMuleVersion4OrAbove returns a boolean if a field has been set.

### GetUri

`func (o *EndpointPostBody) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *EndpointPostBody) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *EndpointPostBody) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *EndpointPostBody) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetType

`func (o *EndpointPostBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EndpointPostBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EndpointPostBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EndpointPostBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIsCloudHub

`func (o *EndpointPostBody) GetIsCloudHub() bool`

GetIsCloudHub returns the IsCloudHub field if non-nil, zero value otherwise.

### GetIsCloudHubOk

`func (o *EndpointPostBody) GetIsCloudHubOk() (*bool, bool)`

GetIsCloudHubOk returns a tuple with the IsCloudHub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloudHub

`func (o *EndpointPostBody) SetIsCloudHub(v bool)`

SetIsCloudHub sets IsCloudHub field to given value.

### HasIsCloudHub

`func (o *EndpointPostBody) HasIsCloudHub() bool`

HasIsCloudHub returns a boolean if a field has been set.

### SetIsCloudHubNil

`func (o *EndpointPostBody) SetIsCloudHubNil(b bool)`

 SetIsCloudHubNil sets the value for IsCloudHub to be an explicit nil

### UnsetIsCloudHub
`func (o *EndpointPostBody) UnsetIsCloudHub()`

UnsetIsCloudHub ensures that no value is present for IsCloudHub, not even an explicit nil
### GetProxyUri

`func (o *EndpointPostBody) GetProxyUri() string`

GetProxyUri returns the ProxyUri field if non-nil, zero value otherwise.

### GetProxyUriOk

`func (o *EndpointPostBody) GetProxyUriOk() (*string, bool)`

GetProxyUriOk returns a tuple with the ProxyUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUri

`func (o *EndpointPostBody) SetProxyUri(v string)`

SetProxyUri sets ProxyUri field to given value.

### HasProxyUri

`func (o *EndpointPostBody) HasProxyUri() bool`

HasProxyUri returns a boolean if a field has been set.

### SetProxyUriNil

`func (o *EndpointPostBody) SetProxyUriNil(b bool)`

 SetProxyUriNil sets the value for ProxyUri to be an explicit nil

### UnsetProxyUri
`func (o *EndpointPostBody) UnsetProxyUri()`

UnsetProxyUri ensures that no value is present for ProxyUri, not even an explicit nil
### GetProxyRegistrationUri

`func (o *EndpointPostBody) GetProxyRegistrationUri() string`

GetProxyRegistrationUri returns the ProxyRegistrationUri field if non-nil, zero value otherwise.

### GetProxyRegistrationUriOk

`func (o *EndpointPostBody) GetProxyRegistrationUriOk() (*string, bool)`

GetProxyRegistrationUriOk returns a tuple with the ProxyRegistrationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyRegistrationUri

`func (o *EndpointPostBody) SetProxyRegistrationUri(v string)`

SetProxyRegistrationUri sets ProxyRegistrationUri field to given value.

### HasProxyRegistrationUri

`func (o *EndpointPostBody) HasProxyRegistrationUri() bool`

HasProxyRegistrationUri returns a boolean if a field has been set.

### SetProxyRegistrationUriNil

`func (o *EndpointPostBody) SetProxyRegistrationUriNil(b bool)`

 SetProxyRegistrationUriNil sets the value for ProxyRegistrationUri to be an explicit nil

### UnsetProxyRegistrationUri
`func (o *EndpointPostBody) UnsetProxyRegistrationUri()`

UnsetProxyRegistrationUri ensures that no value is present for ProxyRegistrationUri, not even an explicit nil
### GetReferencesUserDomain

`func (o *EndpointPostBody) GetReferencesUserDomain() string`

GetReferencesUserDomain returns the ReferencesUserDomain field if non-nil, zero value otherwise.

### GetReferencesUserDomainOk

`func (o *EndpointPostBody) GetReferencesUserDomainOk() (*string, bool)`

GetReferencesUserDomainOk returns a tuple with the ReferencesUserDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencesUserDomain

`func (o *EndpointPostBody) SetReferencesUserDomain(v string)`

SetReferencesUserDomain sets ReferencesUserDomain field to given value.

### HasReferencesUserDomain

`func (o *EndpointPostBody) HasReferencesUserDomain() bool`

HasReferencesUserDomain returns a boolean if a field has been set.

### SetReferencesUserDomainNil

`func (o *EndpointPostBody) SetReferencesUserDomainNil(b bool)`

 SetReferencesUserDomainNil sets the value for ReferencesUserDomain to be an explicit nil

### UnsetReferencesUserDomain
`func (o *EndpointPostBody) UnsetReferencesUserDomain()`

UnsetReferencesUserDomain ensures that no value is present for ReferencesUserDomain, not even an explicit nil
### GetResponseTimeout

`func (o *EndpointPostBody) GetResponseTimeout() string`

GetResponseTimeout returns the ResponseTimeout field if non-nil, zero value otherwise.

### GetResponseTimeoutOk

`func (o *EndpointPostBody) GetResponseTimeoutOk() (*string, bool)`

GetResponseTimeoutOk returns a tuple with the ResponseTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimeout

`func (o *EndpointPostBody) SetResponseTimeout(v string)`

SetResponseTimeout sets ResponseTimeout field to given value.

### HasResponseTimeout

`func (o *EndpointPostBody) HasResponseTimeout() bool`

HasResponseTimeout returns a boolean if a field has been set.

### SetResponseTimeoutNil

`func (o *EndpointPostBody) SetResponseTimeoutNil(b bool)`

 SetResponseTimeoutNil sets the value for ResponseTimeout to be an explicit nil

### UnsetResponseTimeout
`func (o *EndpointPostBody) UnsetResponseTimeout()`

UnsetResponseTimeout ensures that no value is present for ResponseTimeout, not even an explicit nil
### GetTlsContexts

`func (o *EndpointPostBody) GetTlsContexts() EndpointPostBodyTlsContexts`

GetTlsContexts returns the TlsContexts field if non-nil, zero value otherwise.

### GetTlsContextsOk

`func (o *EndpointPostBody) GetTlsContextsOk() (*EndpointPostBodyTlsContexts, bool)`

GetTlsContextsOk returns a tuple with the TlsContexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsContexts

`func (o *EndpointPostBody) SetTlsContexts(v EndpointPostBodyTlsContexts)`

SetTlsContexts sets TlsContexts field to given value.

### HasTlsContexts

`func (o *EndpointPostBody) HasTlsContexts() bool`

HasTlsContexts returns a boolean if a field has been set.

### SetTlsContextsNil

`func (o *EndpointPostBody) SetTlsContextsNil(b bool)`

 SetTlsContextsNil sets the value for TlsContexts to be an explicit nil

### UnsetTlsContexts
`func (o *EndpointPostBody) UnsetTlsContexts()`

UnsetTlsContexts ensures that no value is present for TlsContexts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


