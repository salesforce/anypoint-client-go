# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com/apimanager/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApimInstance**](DefaultApi.md#DeleteApimInstance) | **Delete** /organizations/{orgId}/environments/{envId}/apis/{envApiId} | Delete a specific API Manager Instance
[**GetApimInstanceDetails**](DefaultApi.md#GetApimInstanceDetails) | **Get** /organizations/{orgId}/environments/{envId}/apis/{envApiId} | Retrieves a specific API Manager Instance
[**GetEnvApimInstances**](DefaultApi.md#GetEnvApimInstances) | **Get** /organizations/{orgId}/environments/{envId}/apis | Retrieves a collection of API Manager Instances
[**PatchApimInstance**](DefaultApi.md#PatchApimInstance) | **Patch** /organizations/{orgId}/environments/{envId}/apis/{envApiId} | Patches a specific API Manager Instance
[**PostApimInstance**](DefaultApi.md#PostApimInstance) | **Post** /organizations/{orgId}/environments/{envId}/apis | Creates an API Manager Instance



## DeleteApimInstance

> DeleteApimInstance(ctx, orgId, envId, envApiId).Execute()

Delete a specific API Manager Instance



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    envApiId := "envApiId_example" // string | The api manager instance id for a given environment

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteApimInstance(context.Background(), orgId, envId, envApiId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteApimInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment id | 
**envApiId** | **string** | The api manager instance id for a given environment | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApimInstanceRequest struct via the builder pattern


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


## GetApimInstanceDetails

> ApimInstanceDetails GetApimInstanceDetails(ctx, orgId, envId, envApiId).IncludeProxyTemplate(includeProxyTemplate).IncludeValidation(includeValidation).IncludeTlsContexts(includeTlsContexts).Execute()

Retrieves a specific API Manager Instance



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    envApiId := "envApiId_example" // string | The api manager instance id for a given environment
    includeProxyTemplate := true // bool | Include the configured proxyTemplate to its associated endpoint (optional)
    includeValidation := true // bool | Include the configured validation to its associated endpoint (optional)
    includeTlsContexts := true // bool | Include the configured TLS contexts (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetApimInstanceDetails(context.Background(), orgId, envId, envApiId).IncludeProxyTemplate(includeProxyTemplate).IncludeValidation(includeValidation).IncludeTlsContexts(includeTlsContexts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetApimInstanceDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApimInstanceDetails`: ApimInstanceDetails
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetApimInstanceDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment id | 
**envApiId** | **string** | The api manager instance id for a given environment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApimInstanceDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **includeProxyTemplate** | **bool** | Include the configured proxyTemplate to its associated endpoint | 
 **includeValidation** | **bool** | Include the configured validation to its associated endpoint | 
 **includeTlsContexts** | **bool** | Include the configured TLS contexts | 

### Return type

[**ApimInstanceDetails**](ApimInstanceDetails.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvApimInstances

> ApimInstanceCollection GetEnvApimInstances(ctx, orgId, envId).Query(query).GroupId(groupId).AssetId(assetId).AssetVersion(assetVersion).InstanceLabel(instanceLabel).ProductVersion(productVersion).AutodiscoveryInstanceName(autodiscoveryInstanceName).Filters(filters).Limit(limit).Offset(offset).Sort(sort).Ascending(ascending).Execute()

Retrieves a collection of API Manager Instances



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    query := "query_example" // string | A string that will be checked for a partial or similar matches of the name, description, label and tags (optional)
    groupId := "groupId_example" // string | A string that will be checked for an exact match of the groupId (optional)
    assetId := "assetId_example" // string | A string that will be checked for an exact match of the assetId (optional)
    assetVersion := "assetVersion_example" // string | A string that will be checked for an exact match of the assetVersion (optional)
    instanceLabel := "instanceLabel_example" // string | A string that will be checked for an exact match of the instanceLabel (optional)
    productVersion := "productVersion_example" // string | A string that will be checked for an exact match of the productVersion (optional)
    autodiscoveryInstanceName := "autodiscoveryInstanceName_example" // string | A string that will be checked for an exact match of the autodiscoveryInstanceName (optional)
    filters := []string{"Inner_example"} // []string | Comma-separated list of filters, which can be \"active\" and/or \"pinned\" (optional)
    limit := int32(56) // int32 | Maximum number of rolegroups to retrieve per request. (optional)
    offset := int32(56) // int32 | The number of records to omit from the response. (optional)
    sort := "sort_example" // string | Default value is name (optional)
    ascending := true // bool | To activate ascending sorting (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetEnvApimInstances(context.Background(), orgId, envId).Query(query).GroupId(groupId).AssetId(assetId).AssetVersion(assetVersion).InstanceLabel(instanceLabel).ProductVersion(productVersion).AutodiscoveryInstanceName(autodiscoveryInstanceName).Filters(filters).Limit(limit).Offset(offset).Sort(sort).Ascending(ascending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetEnvApimInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvApimInstances`: ApimInstanceCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetEnvApimInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvApimInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **query** | **string** | A string that will be checked for a partial or similar matches of the name, description, label and tags | 
 **groupId** | **string** | A string that will be checked for an exact match of the groupId | 
 **assetId** | **string** | A string that will be checked for an exact match of the assetId | 
 **assetVersion** | **string** | A string that will be checked for an exact match of the assetVersion | 
 **instanceLabel** | **string** | A string that will be checked for an exact match of the instanceLabel | 
 **productVersion** | **string** | A string that will be checked for an exact match of the productVersion | 
 **autodiscoveryInstanceName** | **string** | A string that will be checked for an exact match of the autodiscoveryInstanceName | 
 **filters** | **[]string** | Comma-separated list of filters, which can be \&quot;active\&quot; and/or \&quot;pinned\&quot; | 
 **limit** | **int32** | Maximum number of rolegroups to retrieve per request. | 
 **offset** | **int32** | The number of records to omit from the response. | 
 **sort** | **string** | Default value is name | 
 **ascending** | **bool** | To activate ascending sorting | 

### Return type

[**ApimInstanceCollection**](ApimInstanceCollection.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchApimInstance

> ApimInstancePatchResponse PatchApimInstance(ctx, orgId, envId, envApiId).Force(force).UpdateApisInSamePort(updateApisInSamePort).Body(body).Execute()

Patches a specific API Manager Instance



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    envApiId := "envApiId_example" // string | The api manager instance id for a given environment
    force := true // bool | Allows patching the API autodiscoveryInstanceName. You may want to change the 'api.version' configuration on all Mule 2 & Mule 3 applications tracking this API. (optional)
    updateApisInSamePort := true // bool | For APIs deployed to Flex, if endpoint proxyUri or inbound TLS Contexts are being updated, apis in the same port are also updated (optional)
    body := map[string]interface{}(Object) // map[string]interface{} | Patch API Manager Instance Body (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.PatchApimInstance(context.Background(), orgId, envId, envApiId).Force(force).UpdateApisInSamePort(updateApisInSamePort).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PatchApimInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchApimInstance`: ApimInstancePatchResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PatchApimInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment id | 
**envApiId** | **string** | The api manager instance id for a given environment | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchApimInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **force** | **bool** | Allows patching the API autodiscoveryInstanceName. You may want to change the &#39;api.version&#39; configuration on all Mule 2 &amp; Mule 3 applications tracking this API. | 
 **updateApisInSamePort** | **bool** | For APIs deployed to Flex, if endpoint proxyUri or inbound TLS Contexts are being updated, apis in the same port are also updated | 
 **body** | **map[string]interface{}** | Patch API Manager Instance Body | 

### Return type

[**ApimInstancePatchResponse**](ApimInstancePatchResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApimInstance

> ApimInstancePostResponse PostApimInstance(ctx, orgId, envId).ApimInstancePostBody(apimInstancePostBody).Execute()

Creates an API Manager Instance



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    apimInstancePostBody := *openapiclient.NewApimInstancePostBody() // ApimInstancePostBody | Post API Manager Instance Body (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.PostApimInstance(context.Background(), orgId, envId).ApimInstancePostBody(apimInstancePostBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PostApimInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostApimInstance`: ApimInstancePostResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PostApimInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApimInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apimInstancePostBody** | [**ApimInstancePostBody**](ApimInstancePostBody.md) | Post API Manager Instance Body | 

### Return type

[**ApimInstancePostResponse**](ApimInstancePostResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

