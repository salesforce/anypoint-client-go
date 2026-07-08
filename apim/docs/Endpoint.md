# Endpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audit** | Pointer to [**Audit**](Audit.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Uri** | Pointer to **NullableString** |  | [optional] 
**ApiGatewayVersion** | Pointer to **NullableString** |  | [optional] 
**ProxyUri** | Pointer to **NullableString** |  | [optional] 
**ProxyRegistrationUri** | Pointer to **NullableString** |  | [optional] 
**LastActiveDate** | Pointer to **NullableString** |  | [optional] 
**IsCloudHub** | Pointer to **NullableBool** |  | [optional] 
**DeploymentType** | Pointer to **string** |  | [optional] 
**PoliciesVersion** | Pointer to **NullableString** |  | [optional] 
**ReferencesUserDomain** | Pointer to **NullableString** |  | [optional] 
**ResponseTimeout** | Pointer to **NullableString** |  | [optional] 
**WsdlConfig** | Pointer to **NullableString** |  | [optional] 
**MuleVersion4OrAbove** | Pointer to **NullableBool** |  | [optional] 
**ApiVersionId** | Pointer to **int32** |  | [optional] 
**Validation** | Pointer to **NullableString** |  | [optional] 
**Console** | Pointer to [**NullableConsole**](Console.md) |  | [optional] 
**TlsContexts** | Pointer to [**EndpointTlsContexts**](EndpointTlsContexts.md) |  | [optional] 

## Methods

### NewEndpoint

`func NewEndpoint() *Endpoint`

NewEndpoint instantiates a new Endpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointWithDefaults

`func NewEndpointWithDefaults() *Endpoint`

NewEndpointWithDefaults instantiates a new Endpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudit

`func (o *Endpoint) GetAudit() Audit`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *Endpoint) GetAuditOk() (*Audit, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *Endpoint) SetAudit(v Audit)`

SetAudit sets Audit field to given value.

### HasAudit

`func (o *Endpoint) HasAudit() bool`

HasAudit returns a boolean if a field has been set.

### GetId

`func (o *Endpoint) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Endpoint) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Endpoint) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Endpoint) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Endpoint) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Endpoint) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Endpoint) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Endpoint) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *Endpoint) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *Endpoint) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUri

`func (o *Endpoint) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *Endpoint) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *Endpoint) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *Endpoint) HasUri() bool`

HasUri returns a boolean if a field has been set.

### SetUriNil

`func (o *Endpoint) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *Endpoint) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil
### GetApiGatewayVersion

`func (o *Endpoint) GetApiGatewayVersion() string`

GetApiGatewayVersion returns the ApiGatewayVersion field if non-nil, zero value otherwise.

### GetApiGatewayVersionOk

`func (o *Endpoint) GetApiGatewayVersionOk() (*string, bool)`

GetApiGatewayVersionOk returns a tuple with the ApiGatewayVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiGatewayVersion

`func (o *Endpoint) SetApiGatewayVersion(v string)`

SetApiGatewayVersion sets ApiGatewayVersion field to given value.

### HasApiGatewayVersion

`func (o *Endpoint) HasApiGatewayVersion() bool`

HasApiGatewayVersion returns a boolean if a field has been set.

### SetApiGatewayVersionNil

`func (o *Endpoint) SetApiGatewayVersionNil(b bool)`

 SetApiGatewayVersionNil sets the value for ApiGatewayVersion to be an explicit nil

### UnsetApiGatewayVersion
`func (o *Endpoint) UnsetApiGatewayVersion()`

UnsetApiGatewayVersion ensures that no value is present for ApiGatewayVersion, not even an explicit nil
### GetProxyUri

`func (o *Endpoint) GetProxyUri() string`

GetProxyUri returns the ProxyUri field if non-nil, zero value otherwise.

### GetProxyUriOk

`func (o *Endpoint) GetProxyUriOk() (*string, bool)`

GetProxyUriOk returns a tuple with the ProxyUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUri

`func (o *Endpoint) SetProxyUri(v string)`

SetProxyUri sets ProxyUri field to given value.

### HasProxyUri

`func (o *Endpoint) HasProxyUri() bool`

HasProxyUri returns a boolean if a field has been set.

### SetProxyUriNil

`func (o *Endpoint) SetProxyUriNil(b bool)`

 SetProxyUriNil sets the value for ProxyUri to be an explicit nil

### UnsetProxyUri
`func (o *Endpoint) UnsetProxyUri()`

UnsetProxyUri ensures that no value is present for ProxyUri, not even an explicit nil
### GetProxyRegistrationUri

`func (o *Endpoint) GetProxyRegistrationUri() string`

GetProxyRegistrationUri returns the ProxyRegistrationUri field if non-nil, zero value otherwise.

### GetProxyRegistrationUriOk

`func (o *Endpoint) GetProxyRegistrationUriOk() (*string, bool)`

GetProxyRegistrationUriOk returns a tuple with the ProxyRegistrationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyRegistrationUri

`func (o *Endpoint) SetProxyRegistrationUri(v string)`

SetProxyRegistrationUri sets ProxyRegistrationUri field to given value.

### HasProxyRegistrationUri

`func (o *Endpoint) HasProxyRegistrationUri() bool`

HasProxyRegistrationUri returns a boolean if a field has been set.

### SetProxyRegistrationUriNil

`func (o *Endpoint) SetProxyRegistrationUriNil(b bool)`

 SetProxyRegistrationUriNil sets the value for ProxyRegistrationUri to be an explicit nil

### UnsetProxyRegistrationUri
`func (o *Endpoint) UnsetProxyRegistrationUri()`

UnsetProxyRegistrationUri ensures that no value is present for ProxyRegistrationUri, not even an explicit nil
### GetLastActiveDate

`func (o *Endpoint) GetLastActiveDate() string`

GetLastActiveDate returns the LastActiveDate field if non-nil, zero value otherwise.

### GetLastActiveDateOk

`func (o *Endpoint) GetLastActiveDateOk() (*string, bool)`

GetLastActiveDateOk returns a tuple with the LastActiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActiveDate

`func (o *Endpoint) SetLastActiveDate(v string)`

SetLastActiveDate sets LastActiveDate field to given value.

### HasLastActiveDate

`func (o *Endpoint) HasLastActiveDate() bool`

HasLastActiveDate returns a boolean if a field has been set.

### SetLastActiveDateNil

`func (o *Endpoint) SetLastActiveDateNil(b bool)`

 SetLastActiveDateNil sets the value for LastActiveDate to be an explicit nil

### UnsetLastActiveDate
`func (o *Endpoint) UnsetLastActiveDate()`

UnsetLastActiveDate ensures that no value is present for LastActiveDate, not even an explicit nil
### GetIsCloudHub

`func (o *Endpoint) GetIsCloudHub() bool`

GetIsCloudHub returns the IsCloudHub field if non-nil, zero value otherwise.

### GetIsCloudHubOk

`func (o *Endpoint) GetIsCloudHubOk() (*bool, bool)`

GetIsCloudHubOk returns a tuple with the IsCloudHub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloudHub

`func (o *Endpoint) SetIsCloudHub(v bool)`

SetIsCloudHub sets IsCloudHub field to given value.

### HasIsCloudHub

`func (o *Endpoint) HasIsCloudHub() bool`

HasIsCloudHub returns a boolean if a field has been set.

### SetIsCloudHubNil

`func (o *Endpoint) SetIsCloudHubNil(b bool)`

 SetIsCloudHubNil sets the value for IsCloudHub to be an explicit nil

### UnsetIsCloudHub
`func (o *Endpoint) UnsetIsCloudHub()`

UnsetIsCloudHub ensures that no value is present for IsCloudHub, not even an explicit nil
### GetDeploymentType

`func (o *Endpoint) GetDeploymentType() string`

GetDeploymentType returns the DeploymentType field if non-nil, zero value otherwise.

### GetDeploymentTypeOk

`func (o *Endpoint) GetDeploymentTypeOk() (*string, bool)`

GetDeploymentTypeOk returns a tuple with the DeploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentType

`func (o *Endpoint) SetDeploymentType(v string)`

SetDeploymentType sets DeploymentType field to given value.

### HasDeploymentType

`func (o *Endpoint) HasDeploymentType() bool`

HasDeploymentType returns a boolean if a field has been set.

### GetPoliciesVersion

`func (o *Endpoint) GetPoliciesVersion() string`

GetPoliciesVersion returns the PoliciesVersion field if non-nil, zero value otherwise.

### GetPoliciesVersionOk

`func (o *Endpoint) GetPoliciesVersionOk() (*string, bool)`

GetPoliciesVersionOk returns a tuple with the PoliciesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoliciesVersion

`func (o *Endpoint) SetPoliciesVersion(v string)`

SetPoliciesVersion sets PoliciesVersion field to given value.

### HasPoliciesVersion

`func (o *Endpoint) HasPoliciesVersion() bool`

HasPoliciesVersion returns a boolean if a field has been set.

### SetPoliciesVersionNil

`func (o *Endpoint) SetPoliciesVersionNil(b bool)`

 SetPoliciesVersionNil sets the value for PoliciesVersion to be an explicit nil

### UnsetPoliciesVersion
`func (o *Endpoint) UnsetPoliciesVersion()`

UnsetPoliciesVersion ensures that no value is present for PoliciesVersion, not even an explicit nil
### GetReferencesUserDomain

`func (o *Endpoint) GetReferencesUserDomain() string`

GetReferencesUserDomain returns the ReferencesUserDomain field if non-nil, zero value otherwise.

### GetReferencesUserDomainOk

`func (o *Endpoint) GetReferencesUserDomainOk() (*string, bool)`

GetReferencesUserDomainOk returns a tuple with the ReferencesUserDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencesUserDomain

`func (o *Endpoint) SetReferencesUserDomain(v string)`

SetReferencesUserDomain sets ReferencesUserDomain field to given value.

### HasReferencesUserDomain

`func (o *Endpoint) HasReferencesUserDomain() bool`

HasReferencesUserDomain returns a boolean if a field has been set.

### SetReferencesUserDomainNil

`func (o *Endpoint) SetReferencesUserDomainNil(b bool)`

 SetReferencesUserDomainNil sets the value for ReferencesUserDomain to be an explicit nil

### UnsetReferencesUserDomain
`func (o *Endpoint) UnsetReferencesUserDomain()`

UnsetReferencesUserDomain ensures that no value is present for ReferencesUserDomain, not even an explicit nil
### GetResponseTimeout

`func (o *Endpoint) GetResponseTimeout() string`

GetResponseTimeout returns the ResponseTimeout field if non-nil, zero value otherwise.

### GetResponseTimeoutOk

`func (o *Endpoint) GetResponseTimeoutOk() (*string, bool)`

GetResponseTimeoutOk returns a tuple with the ResponseTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimeout

`func (o *Endpoint) SetResponseTimeout(v string)`

SetResponseTimeout sets ResponseTimeout field to given value.

### HasResponseTimeout

`func (o *Endpoint) HasResponseTimeout() bool`

HasResponseTimeout returns a boolean if a field has been set.

### SetResponseTimeoutNil

`func (o *Endpoint) SetResponseTimeoutNil(b bool)`

 SetResponseTimeoutNil sets the value for ResponseTimeout to be an explicit nil

### UnsetResponseTimeout
`func (o *Endpoint) UnsetResponseTimeout()`

UnsetResponseTimeout ensures that no value is present for ResponseTimeout, not even an explicit nil
### GetWsdlConfig

`func (o *Endpoint) GetWsdlConfig() string`

GetWsdlConfig returns the WsdlConfig field if non-nil, zero value otherwise.

### GetWsdlConfigOk

`func (o *Endpoint) GetWsdlConfigOk() (*string, bool)`

GetWsdlConfigOk returns a tuple with the WsdlConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWsdlConfig

`func (o *Endpoint) SetWsdlConfig(v string)`

SetWsdlConfig sets WsdlConfig field to given value.

### HasWsdlConfig

`func (o *Endpoint) HasWsdlConfig() bool`

HasWsdlConfig returns a boolean if a field has been set.

### SetWsdlConfigNil

`func (o *Endpoint) SetWsdlConfigNil(b bool)`

 SetWsdlConfigNil sets the value for WsdlConfig to be an explicit nil

### UnsetWsdlConfig
`func (o *Endpoint) UnsetWsdlConfig()`

UnsetWsdlConfig ensures that no value is present for WsdlConfig, not even an explicit nil
### GetMuleVersion4OrAbove

`func (o *Endpoint) GetMuleVersion4OrAbove() bool`

GetMuleVersion4OrAbove returns the MuleVersion4OrAbove field if non-nil, zero value otherwise.

### GetMuleVersion4OrAboveOk

`func (o *Endpoint) GetMuleVersion4OrAboveOk() (*bool, bool)`

GetMuleVersion4OrAboveOk returns a tuple with the MuleVersion4OrAbove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuleVersion4OrAbove

`func (o *Endpoint) SetMuleVersion4OrAbove(v bool)`

SetMuleVersion4OrAbove sets MuleVersion4OrAbove field to given value.

### HasMuleVersion4OrAbove

`func (o *Endpoint) HasMuleVersion4OrAbove() bool`

HasMuleVersion4OrAbove returns a boolean if a field has been set.

### SetMuleVersion4OrAboveNil

`func (o *Endpoint) SetMuleVersion4OrAboveNil(b bool)`

 SetMuleVersion4OrAboveNil sets the value for MuleVersion4OrAbove to be an explicit nil

### UnsetMuleVersion4OrAbove
`func (o *Endpoint) UnsetMuleVersion4OrAbove()`

UnsetMuleVersion4OrAbove ensures that no value is present for MuleVersion4OrAbove, not even an explicit nil
### GetApiVersionId

`func (o *Endpoint) GetApiVersionId() int32`

GetApiVersionId returns the ApiVersionId field if non-nil, zero value otherwise.

### GetApiVersionIdOk

`func (o *Endpoint) GetApiVersionIdOk() (*int32, bool)`

GetApiVersionIdOk returns a tuple with the ApiVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersionId

`func (o *Endpoint) SetApiVersionId(v int32)`

SetApiVersionId sets ApiVersionId field to given value.

### HasApiVersionId

`func (o *Endpoint) HasApiVersionId() bool`

HasApiVersionId returns a boolean if a field has been set.

### GetValidation

`func (o *Endpoint) GetValidation() string`

GetValidation returns the Validation field if non-nil, zero value otherwise.

### GetValidationOk

`func (o *Endpoint) GetValidationOk() (*string, bool)`

GetValidationOk returns a tuple with the Validation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidation

`func (o *Endpoint) SetValidation(v string)`

SetValidation sets Validation field to given value.

### HasValidation

`func (o *Endpoint) HasValidation() bool`

HasValidation returns a boolean if a field has been set.

### SetValidationNil

`func (o *Endpoint) SetValidationNil(b bool)`

 SetValidationNil sets the value for Validation to be an explicit nil

### UnsetValidation
`func (o *Endpoint) UnsetValidation()`

UnsetValidation ensures that no value is present for Validation, not even an explicit nil
### GetConsole

`func (o *Endpoint) GetConsole() Console`

GetConsole returns the Console field if non-nil, zero value otherwise.

### GetConsoleOk

`func (o *Endpoint) GetConsoleOk() (*Console, bool)`

GetConsoleOk returns a tuple with the Console field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsole

`func (o *Endpoint) SetConsole(v Console)`

SetConsole sets Console field to given value.

### HasConsole

`func (o *Endpoint) HasConsole() bool`

HasConsole returns a boolean if a field has been set.

### SetConsoleNil

`func (o *Endpoint) SetConsoleNil(b bool)`

 SetConsoleNil sets the value for Console to be an explicit nil

### UnsetConsole
`func (o *Endpoint) UnsetConsole()`

UnsetConsole ensures that no value is present for Console, not even an explicit nil
### GetTlsContexts

`func (o *Endpoint) GetTlsContexts() EndpointTlsContexts`

GetTlsContexts returns the TlsContexts field if non-nil, zero value otherwise.

### GetTlsContextsOk

`func (o *Endpoint) GetTlsContextsOk() (*EndpointTlsContexts, bool)`

GetTlsContextsOk returns a tuple with the TlsContexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsContexts

`func (o *Endpoint) SetTlsContexts(v EndpointTlsContexts)`

SetTlsContexts sets TlsContexts field to given value.

### HasTlsContexts

`func (o *Endpoint) HasTlsContexts() bool`

HasTlsContexts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


