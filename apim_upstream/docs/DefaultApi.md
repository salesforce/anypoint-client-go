# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com/apimanager/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApimInstanceUpstream**](DefaultApi.md#DeleteApimInstanceUpstream) | **Delete** /organizations/{orgId}/environments/{envId}/apis/{envApiId}/upstreams/{upstreamId} | Delete a specific Upstream of the given API Manager Instance
[**GetApimInstanceUpstream**](DefaultApi.md#GetApimInstanceUpstream) | **Get** /organizations/{orgId}/environments/{envId}/apis/{envApiId}/upstreams/{upstreamId} | Retrieve a specific upstream for a given API Manager instance
[**GetApimInstanceUpstreams**](DefaultApi.md#GetApimInstanceUpstreams) | **Get** /organizations/{orgId}/environments/{envId}/apis/{envApiId}/upstreams | Retrieve all upstreams of a given API Manager instance
[**PatchApimInstanceUpstream**](DefaultApi.md#PatchApimInstanceUpstream) | **Patch** /organizations/{orgId}/environments/{envId}/apis/{envApiId}/upstreams/{upstreamId} | Update a specific upstream in a given API Manager instance
[**PostApimInstanceUpstream**](DefaultApi.md#PostApimInstanceUpstream) | **Post** /organizations/{orgId}/environments/{envId}/apis/{envApiId}/upstreams | Creates an upstream for a given API Manager instance



## DeleteApimInstanceUpstream

> DeleteApimInstanceUpstream(ctx, orgId, envId, envApiId, upstreamId).Execute()

Delete a specific Upstream of the given API Manager Instance



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
    envApiId := "envApiId_example" // string | The api id specific to a given environment
    upstreamId := "upstreamId_example" // string | The upstream id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteApimInstanceUpstream(context.Background(), orgId, envId, envApiId, upstreamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteApimInstanceUpstream``: %v\n", err)
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
**envApiId** | **string** | The api id specific to a given environment | 
**upstreamId** | **string** | The upstream id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApimInstanceUpstreamRequest struct via the builder pattern


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


## GetApimInstanceUpstream

> UpstreamDetails GetApimInstanceUpstream(ctx, orgId, envId, envApiId, upstreamId).Execute()

Retrieve a specific upstream for a given API Manager instance



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
    envApiId := "envApiId_example" // string | The api id specific to a given environment
    upstreamId := "upstreamId_example" // string | The upstream id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetApimInstanceUpstream(context.Background(), orgId, envId, envApiId, upstreamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetApimInstanceUpstream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApimInstanceUpstream`: UpstreamDetails
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetApimInstanceUpstream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment id | 
**envApiId** | **string** | The api id specific to a given environment | 
**upstreamId** | **string** | The upstream id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApimInstanceUpstreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**UpstreamDetails**](UpstreamDetails.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApimInstanceUpstreams

> UpstreamCollection GetApimInstanceUpstreams(ctx, orgId, envId, envApiId).Execute()

Retrieve all upstreams of a given API Manager instance



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
    envApiId := "envApiId_example" // string | The api id specific to a given environment

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetApimInstanceUpstreams(context.Background(), orgId, envId, envApiId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetApimInstanceUpstreams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApimInstanceUpstreams`: UpstreamCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetApimInstanceUpstreams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment id | 
**envApiId** | **string** | The api id specific to a given environment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApimInstanceUpstreamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**UpstreamCollection**](UpstreamCollection.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchApimInstanceUpstream

> Upstream PatchApimInstanceUpstream(ctx, orgId, envId, envApiId, upstreamId).UpstreamPatchBody(upstreamPatchBody).Execute()

Update a specific upstream in a given API Manager instance



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
    envApiId := "envApiId_example" // string | The api id specific to a given environment
    upstreamId := "upstreamId_example" // string | The upstream id
    upstreamPatchBody := *openapiclient.NewUpstreamPatchBody() // UpstreamPatchBody | Patch API Manager Instance Upstream Body (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.PatchApimInstanceUpstream(context.Background(), orgId, envId, envApiId, upstreamId).UpstreamPatchBody(upstreamPatchBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PatchApimInstanceUpstream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchApimInstanceUpstream`: Upstream
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PatchApimInstanceUpstream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment id | 
**envApiId** | **string** | The api id specific to a given environment | 
**upstreamId** | **string** | The upstream id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchApimInstanceUpstreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **upstreamPatchBody** | [**UpstreamPatchBody**](UpstreamPatchBody.md) | Patch API Manager Instance Upstream Body | 

### Return type

[**Upstream**](Upstream.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApimInstanceUpstream

> UpstreamDetails PostApimInstanceUpstream(ctx, orgId, envId, envApiId).UpstreamPostBody(upstreamPostBody).Execute()

Creates an upstream for a given API Manager instance



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
    envApiId := "envApiId_example" // string | The api id specific to a given environment
    upstreamPostBody := *openapiclient.NewUpstreamPostBody() // UpstreamPostBody | Post API Manager Instance Upstream Body (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.PostApimInstanceUpstream(context.Background(), orgId, envId, envApiId).UpstreamPostBody(upstreamPostBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PostApimInstanceUpstream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostApimInstanceUpstream`: UpstreamDetails
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PostApimInstanceUpstream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment id | 
**envApiId** | **string** | The api id specific to a given environment | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApimInstanceUpstreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **upstreamPostBody** | [**UpstreamPostBody**](UpstreamPostBody.md) | Post API Manager Instance Upstream Body | 

### Return type

[**UpstreamDetails**](UpstreamDetails.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

