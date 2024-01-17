# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com/secrets-manager/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSecretGroupTlsContextDetails**](DefaultApi.md#GetSecretGroupTlsContextDetails) | **Get** /organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId}/tlsContexts/{secretId} | Retrieve tls-context details
[**GetSecretGroupTlsContexts**](DefaultApi.md#GetSecretGroupTlsContexts) | **Get** /organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId}/tlsContexts | Retrieves a secret-groups&#39; collection of tls-contexts.
[**PatchSecretGroupTlsContext**](DefaultApi.md#PatchSecretGroupTlsContext) | **Patch** /organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId}/tlsContexts/{secretId} | Update a given secret-group tls-context
[**PostSecretGroupTlsContext**](DefaultApi.md#PostSecretGroupTlsContext) | **Post** /organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId}/tlsContexts | Create a secret-groups&#39; tls-context.
[**PutSecretGroupTlsContext**](DefaultApi.md#PutSecretGroupTlsContext) | **Put** /organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId}/tlsContexts/{secretId} | Update a given secret-group tls-context



## GetSecretGroupTlsContextDetails

> TlsContextDetails GetSecretGroupTlsContextDetails(ctx, orgId, envId, secretGroupId, secretId).Execute()

Retrieve tls-context details



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/secretgroup_tlscontext"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    secretGroupId := "secretGroupId_example" // string | The secret group id
    secretId := "secretId_example" // string | The keystore id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetSecretGroupTlsContextDetails(context.Background(), orgId, envId, secretGroupId, secretId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSecretGroupTlsContextDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecretGroupTlsContextDetails`: TlsContextDetails
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSecretGroupTlsContextDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment id | 
**secretGroupId** | **string** | The secret group id | 
**secretId** | **string** | The keystore id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretGroupTlsContextDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**TlsContextDetails**](TlsContextDetails.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecretGroupTlsContexts

> []TlsContextSummary GetSecretGroupTlsContexts(ctx, orgId, envId, secretGroupId).Execute()

Retrieves a secret-groups' collection of tls-contexts.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/secretgroup_tlscontext"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    secretGroupId := "secretGroupId_example" // string | The secret group id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetSecretGroupTlsContexts(context.Background(), orgId, envId, secretGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSecretGroupTlsContexts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecretGroupTlsContexts`: []TlsContextSummary
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSecretGroupTlsContexts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment id | 
**secretGroupId** | **string** | The secret group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretGroupTlsContextsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]TlsContextSummary**](TlsContextSummary.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSecretGroupTlsContext

> PutSecretGroupTlsContext200Response PatchSecretGroupTlsContext(ctx, orgId, envId, secretGroupId, secretId).Body(body).Execute()

Update a given secret-group tls-context



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/secretgroup_tlscontext"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    secretGroupId := "secretGroupId_example" // string | The secret group id
    secretId := "secretId_example" // string | The keystore id
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.PatchSecretGroupTlsContext(context.Background(), orgId, envId, secretGroupId, secretId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PatchSecretGroupTlsContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchSecretGroupTlsContext`: PutSecretGroupTlsContext200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PatchSecretGroupTlsContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment id | 
**secretGroupId** | **string** | The secret group id | 
**secretId** | **string** | The keystore id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSecretGroupTlsContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | **map[string]interface{}** |  | 

### Return type

[**PutSecretGroupTlsContext200Response**](PutSecretGroupTlsContext200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSecretGroupTlsContext

> PostSecretGroupTlsContext201Response PostSecretGroupTlsContext(ctx, orgId, envId, secretGroupId).TlsContextPostBody(tlsContextPostBody).Execute()

Create a secret-groups' tls-context.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/secretgroup_tlscontext"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    secretGroupId := "secretGroupId_example" // string | The secret group id
    tlsContextPostBody := openapiclient.TlsContextPostBody{TlsContextFlexGatewayBody: openapiclient.NewTlsContextFlexGatewayBody()} // TlsContextPostBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.PostSecretGroupTlsContext(context.Background(), orgId, envId, secretGroupId).TlsContextPostBody(tlsContextPostBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PostSecretGroupTlsContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostSecretGroupTlsContext`: PostSecretGroupTlsContext201Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PostSecretGroupTlsContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment id | 
**secretGroupId** | **string** | The secret group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSecretGroupTlsContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **tlsContextPostBody** | [**TlsContextPostBody**](TlsContextPostBody.md) |  | 

### Return type

[**PostSecretGroupTlsContext201Response**](PostSecretGroupTlsContext201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSecretGroupTlsContext

> PutSecretGroupTlsContext200Response PutSecretGroupTlsContext(ctx, orgId, envId, secretGroupId, secretId).TlsContextPutBody(tlsContextPutBody).Execute()

Update a given secret-group tls-context



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/secretgroup_tlscontext"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    secretGroupId := "secretGroupId_example" // string | The secret group id
    secretId := "secretId_example" // string | The keystore id
    tlsContextPutBody := openapiclient.TlsContextPutBody{TlsContextFlexGatewayBody: openapiclient.NewTlsContextFlexGatewayBody()} // TlsContextPutBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.PutSecretGroupTlsContext(context.Background(), orgId, envId, secretGroupId, secretId).TlsContextPutBody(tlsContextPutBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PutSecretGroupTlsContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutSecretGroupTlsContext`: PutSecretGroupTlsContext200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PutSecretGroupTlsContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment id | 
**secretGroupId** | **string** | The secret group id | 
**secretId** | **string** | The keystore id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSecretGroupTlsContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **tlsContextPutBody** | [**TlsContextPutBody**](TlsContextPutBody.md) |  | 

### Return type

[**PutSecretGroupTlsContext200Response**](PutSecretGroupTlsContext200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

