# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com/apimanager*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApimPolicy**](DefaultApi.md#DeleteApimPolicy) | **Delete** /api/v1/organizations/{orgId}/environments/{envId}/apis/{apiId}/policies/{apiPolicyId} | Delete a specific api manager instance policy.
[**DisableApimPolicy**](DefaultApi.md#DisableApimPolicy) | **Post** /xapi/v1/organizations/{orgId}/environments/{envId}/apis/{apiId}/policies/{apiPolicyId}/disable | Disable a specific api manager instance policy.
[**EnableApimPolicy**](DefaultApi.md#EnableApimPolicy) | **Post** /xapi/v1/organizations/{orgId}/environments/{envId}/apis/{apiId}/policies/{apiPolicyId}/enable | Enable a specific api manager instance policy.
[**GetApimPolicies**](DefaultApi.md#GetApimPolicies) | **Get** /api/v1/organizations/{orgId}/environments/{envId}/apis/{apiId}/policies | Retrieve all of api manager instance policies.
[**GetApimPolicy**](DefaultApi.md#GetApimPolicy) | **Get** /api/v1/organizations/{orgId}/environments/{envId}/apis/{apiId}/policies/{apiPolicyId} | Retrieve a specific api manager instance policy.
[**GetOrgAutomatedPolicies**](DefaultApi.md#GetOrgAutomatedPolicies) | **Get** /api/v1/organizations/{orgId}/automated-policies | Retrieve all automated policies of a given organization
[**GetOrgCustomPolicyTemplates**](DefaultApi.md#GetOrgCustomPolicyTemplates) | **Get** /api/v1/organizations/{orgId}/custom-policy-templates | Retrieve all or part of custom policy templates of a given organization
[**GetOrgExchangePolicyTemplateDetails**](DefaultApi.md#GetOrgExchangePolicyTemplateDetails) | **Get** /xapi/v1/organizations/{orgId}/exchange-policy-templates/{groupId}/{assetId}/{assetVersion} | Retrieve details of exchange policy template of a given organization
[**GetOrgExchangePolicyTemplates**](DefaultApi.md#GetOrgExchangePolicyTemplates) | **Get** /xapi/v1/organizations/{orgId}/exchange-policy-templates | Retrieve all or part of exchange policy templates of a given organization
[**PatchApimPolicy**](DefaultApi.md#PatchApimPolicy) | **Patch** /api/v1/organizations/{orgId}/environments/{envId}/apis/{apiId}/policies/{apiPolicyId} | Update a specific api manager instance policy.
[**PostApimPolicy**](DefaultApi.md#PostApimPolicy) | **Post** /api/v1/organizations/{orgId}/environments/{envId}/apis/{apiId}/policies | Create an api manager instance policy.



## DeleteApimPolicy

> DeleteApimPolicy(ctx, orgId, envId, apiId, apiPolicyId).Execute()

Delete a specific api manager instance policy.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/apim_policy"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment Id
    apiId := "apiId_example" // string | The api manager instance Id
    apiPolicyId := "apiPolicyId_example" // string | The api manager instance policy Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeleteApimPolicy(context.Background(), orgId, envId, apiId, apiPolicyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteApimPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment Id | 
**apiId** | **string** | The api manager instance Id | 
**apiPolicyId** | **string** | The api manager instance policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApimPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableApimPolicy

> ApimPolicy DisableApimPolicy(ctx, orgId, envId, apiId, apiPolicyId).Execute()

Disable a specific api manager instance policy.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/apim_policy"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment Id
    apiId := "apiId_example" // string | The api manager instance Id
    apiPolicyId := "apiPolicyId_example" // string | The api manager instance policy Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DisableApimPolicy(context.Background(), orgId, envId, apiId, apiPolicyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DisableApimPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisableApimPolicy`: ApimPolicy
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DisableApimPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment Id | 
**apiId** | **string** | The api manager instance Id | 
**apiPolicyId** | **string** | The api manager instance policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableApimPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**ApimPolicy**](ApimPolicy.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableApimPolicy

> ApimPolicy EnableApimPolicy(ctx, orgId, envId, apiId, apiPolicyId).Execute()

Enable a specific api manager instance policy.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/apim_policy"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment Id
    apiId := "apiId_example" // string | The api manager instance Id
    apiPolicyId := "apiPolicyId_example" // string | The api manager instance policy Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.EnableApimPolicy(context.Background(), orgId, envId, apiId, apiPolicyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.EnableApimPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableApimPolicy`: ApimPolicy
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.EnableApimPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment Id | 
**apiId** | **string** | The api manager instance Id | 
**apiPolicyId** | **string** | The api manager instance policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableApimPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**ApimPolicy**](ApimPolicy.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApimPolicies

> ApimPolicyCollection GetApimPolicies(ctx, orgId, envId, apiId).FullInfo(fullInfo).Execute()

Retrieve all of api manager instance policies.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/apim_policy"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment Id
    apiId := "apiId_example" // string | The api manager instance Id
    fullInfo := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetApimPolicies(context.Background(), orgId, envId, apiId).FullInfo(fullInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetApimPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApimPolicies`: ApimPolicyCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetApimPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment Id | 
**apiId** | **string** | The api manager instance Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApimPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fullInfo** | **bool** |  | [default to false]

### Return type

[**ApimPolicyCollection**](ApimPolicyCollection.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApimPolicy

> ApimPolicy GetApimPolicy(ctx, orgId, envId, apiId, apiPolicyId).Execute()

Retrieve a specific api manager instance policy.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/apim_policy"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment Id
    apiId := "apiId_example" // string | The api manager instance Id
    apiPolicyId := "apiPolicyId_example" // string | The api manager instance policy Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetApimPolicy(context.Background(), orgId, envId, apiId, apiPolicyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetApimPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApimPolicy`: ApimPolicy
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetApimPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment Id | 
**apiId** | **string** | The api manager instance Id | 
**apiPolicyId** | **string** | The api manager instance policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApimPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**ApimPolicy**](ApimPolicy.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgAutomatedPolicies

> AutomatedPolicyCollection GetOrgAutomatedPolicies(ctx, orgId).EnvironmentId(environmentId).Execute()

Retrieve all automated policies of a given organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/apim_policy"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    environmentId := "environmentId_example" // string | A environment id (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetOrgAutomatedPolicies(context.Background(), orgId).EnvironmentId(environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetOrgAutomatedPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrgAutomatedPolicies`: AutomatedPolicyCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetOrgAutomatedPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgAutomatedPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentId** | **string** | A environment id | 

### Return type

[**AutomatedPolicyCollection**](AutomatedPolicyCollection.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgCustomPolicyTemplates

> GetOrgCustomPolicyTemplates(ctx, orgId).Query(query).Offset(offset).Ascending(ascending).Sort(sort).Limit(limit).Execute()

Retrieve all or part of custom policy templates of a given organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/apim_policy"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    query := "query_example" // string | Search criteria. (optional)
    offset := int32(56) // int32 | Skip over a number of elements by specifying an offset value for the query. (optional)
    ascending := true // bool | Order for sorting. (optional)
    sort := "sort_example" // string | Property to sort by. (optional)
    limit := int32(56) // int32 | Limit the number of elements on the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.GetOrgCustomPolicyTemplates(context.Background(), orgId).Query(query).Offset(offset).Ascending(ascending).Sort(sort).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetOrgCustomPolicyTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgCustomPolicyTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **string** | Search criteria. | 
 **offset** | **int32** | Skip over a number of elements by specifying an offset value for the query. | 
 **ascending** | **bool** | Order for sorting. | 
 **sort** | **string** | Property to sort by. | 
 **limit** | **int32** | Limit the number of elements on the response. | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgExchangePolicyTemplateDetails

> ExchangePolicyTemplate GetOrgExchangePolicyTemplateDetails(ctx, orgId, groupId, assetId, assetVersion).IncludeAllVersions(includeAllVersions).SplitModel(splitModel).Execute()

Retrieve details of exchange policy template of a given organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/apim_policy"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    groupId := "groupId_example" // string | The group Id
    assetId := "assetId_example" // string | The asset Id
    assetVersion := "assetVersion_example" // string | The asset version
    includeAllVersions := true // bool | Whether to include all versions of the asset. (optional)
    splitModel := true // bool | Whether to include asset split model. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetOrgExchangePolicyTemplateDetails(context.Background(), orgId, groupId, assetId, assetVersion).IncludeAllVersions(includeAllVersions).SplitModel(splitModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetOrgExchangePolicyTemplateDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrgExchangePolicyTemplateDetails`: ExchangePolicyTemplate
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetOrgExchangePolicyTemplateDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**groupId** | **string** | The group Id | 
**assetId** | **string** | The asset Id | 
**assetVersion** | **string** | The asset version | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgExchangePolicyTemplateDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **includeAllVersions** | **bool** | Whether to include all versions of the asset. | 
 **splitModel** | **bool** | Whether to include asset split model. | 

### Return type

[**ExchangePolicyTemplate**](ExchangePolicyTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgExchangePolicyTemplates

> []ExchangePolicyTemplate GetOrgExchangePolicyTemplates(ctx, orgId).EnvironmentId(environmentId).SplitModel(splitModel).Latest(latest).ApiInstanceId(apiInstanceId).IncludeConfiguration(includeConfiguration).AutomatedOnly(automatedOnly).Execute()

Retrieve all or part of exchange policy templates of a given organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/apim_policy"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    environmentId := "environmentId_example" // string | The environment id. (optional)
    splitModel := true // bool | Whether to include asset split model. (optional)
    latest := true // bool | include only latest versions. (optional)
    apiInstanceId := "apiInstanceId_example" // string | include only templates used for api instance id. (optional)
    includeConfiguration := true // bool | whether to include configuration. (optional)
    automatedOnly := true // bool | whether to include automated policies only. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetOrgExchangePolicyTemplates(context.Background(), orgId).EnvironmentId(environmentId).SplitModel(splitModel).Latest(latest).ApiInstanceId(apiInstanceId).IncludeConfiguration(includeConfiguration).AutomatedOnly(automatedOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetOrgExchangePolicyTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrgExchangePolicyTemplates`: []ExchangePolicyTemplate
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetOrgExchangePolicyTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgExchangePolicyTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentId** | **string** | The environment id. | 
 **splitModel** | **bool** | Whether to include asset split model. | 
 **latest** | **bool** | include only latest versions. | 
 **apiInstanceId** | **string** | include only templates used for api instance id. | 
 **includeConfiguration** | **bool** | whether to include configuration. | 
 **automatedOnly** | **bool** | whether to include automated policies only. | 

### Return type

[**[]ExchangePolicyTemplate**](ExchangePolicyTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchApimPolicy

> ApimPolicy PatchApimPolicy(ctx, orgId, envId, apiId, apiPolicyId).Body(body).Execute()

Update a specific api manager instance policy.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/apim_policy"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment Id
    apiId := "apiId_example" // string | The api manager instance Id
    apiPolicyId := "apiPolicyId_example" // string | The api manager instance policy Id
    body := map[string]interface{}{ ... } // map[string]interface{} | policy content (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.PatchApimPolicy(context.Background(), orgId, envId, apiId, apiPolicyId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PatchApimPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchApimPolicy`: ApimPolicy
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PatchApimPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment Id | 
**apiId** | **string** | The api manager instance Id | 
**apiPolicyId** | **string** | The api manager instance policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchApimPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | **map[string]interface{}** | policy content | 

### Return type

[**ApimPolicy**](ApimPolicy.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApimPolicy

> ApimPolicy PostApimPolicy(ctx, orgId, envId, apiId).ApimPolicyBody(apimPolicyBody).Execute()

Create an api manager instance policy.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/apim_policy"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment Id
    apiId := "apiId_example" // string | The api manager instance Id
    apimPolicyBody := *openapiclient.NewApimPolicyBody() // ApimPolicyBody | policy content (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.PostApimPolicy(context.Background(), orgId, envId, apiId).ApimPolicyBody(apimPolicyBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PostApimPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostApimPolicy`: ApimPolicy
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PostApimPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment Id | 
**apiId** | **string** | The api manager instance Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApimPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **apimPolicyBody** | [**ApimPolicyBody**](ApimPolicyBody.md) | policy content | 

### Return type

[**ApimPolicy**](ApimPolicy.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

