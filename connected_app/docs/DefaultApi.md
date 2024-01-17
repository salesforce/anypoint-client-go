# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com/accounts/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConnectedApplicationsConnAppIdDelete**](DefaultApi.md#ConnectedApplicationsConnAppIdDelete) | **Delete** /connectedApplications/{connAppId} | 
[**ConnectedApplicationsConnAppIdGet**](DefaultApi.md#ConnectedApplicationsConnAppIdGet) | **Get** /connectedApplications/{connAppId} | 
[**ConnectedApplicationsConnAppIdPatch**](DefaultApi.md#ConnectedApplicationsConnAppIdPatch) | **Patch** /connectedApplications/{connAppId} | 
[**ConnectedApplicationsConnAppIdScopesGet**](DefaultApi.md#ConnectedApplicationsConnAppIdScopesGet) | **Get** /connectedApplications/{connAppId}/scopes | 
[**ConnectedApplicationsConnAppIdScopesPut**](DefaultApi.md#ConnectedApplicationsConnAppIdScopesPut) | **Put** /connectedApplications/{connAppId}/scopes | 
[**ConnectedApplicationsGet**](DefaultApi.md#ConnectedApplicationsGet) | **Get** /connectedApplications | 
[**ConnectedApplicationsPost**](DefaultApi.md#ConnectedApplicationsPost) | **Post** /connectedApplications | 



## ConnectedApplicationsConnAppIdDelete

> ConnectedApplicationsConnAppIdDelete(ctx, connAppId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/connected_app"
)

func main() {
    connAppId := "connAppId_example" // string | The ID of the connected app

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.ConnectedApplicationsConnAppIdDelete(context.Background(), connAppId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ConnectedApplicationsConnAppIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connAppId** | **string** | The ID of the connected app | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectedApplicationsConnAppIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectedApplicationsConnAppIdGet

> ConnectedAppRespExt ConnectedApplicationsConnAppIdGet(ctx, connAppId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/connected_app"
)

func main() {
    connAppId := "connAppId_example" // string | The ID of the connected app

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ConnectedApplicationsConnAppIdGet(context.Background(), connAppId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ConnectedApplicationsConnAppIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectedApplicationsConnAppIdGet`: ConnectedAppRespExt
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ConnectedApplicationsConnAppIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connAppId** | **string** | The ID of the connected app | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectedApplicationsConnAppIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConnectedAppRespExt**](ConnectedAppRespExt.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectedApplicationsConnAppIdPatch

> ConnectedAppRespExt ConnectedApplicationsConnAppIdPatch(ctx, connAppId).ConnectedAppPatchExt(connectedAppPatchExt).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/connected_app"
)

func main() {
    connAppId := "connAppId_example" // string | The ID of the connected app
    connectedAppPatchExt := *openapiclient.NewConnectedAppPatchExt() // ConnectedAppPatchExt |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ConnectedApplicationsConnAppIdPatch(context.Background(), connAppId).ConnectedAppPatchExt(connectedAppPatchExt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ConnectedApplicationsConnAppIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectedApplicationsConnAppIdPatch`: ConnectedAppRespExt
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ConnectedApplicationsConnAppIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connAppId** | **string** | The ID of the connected app | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectedApplicationsConnAppIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectedAppPatchExt** | [**ConnectedAppPatchExt**](ConnectedAppPatchExt.md) |  | 

### Return type

[**ConnectedAppRespExt**](ConnectedAppRespExt.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectedApplicationsConnAppIdScopesGet

> ConnectedApplicationsConnAppIdScopesGet200Response ConnectedApplicationsConnAppIdScopesGet(ctx, connAppId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/connected_app"
)

func main() {
    connAppId := "connAppId_example" // string | The ID of the connected app

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ConnectedApplicationsConnAppIdScopesGet(context.Background(), connAppId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ConnectedApplicationsConnAppIdScopesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectedApplicationsConnAppIdScopesGet`: ConnectedApplicationsConnAppIdScopesGet200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ConnectedApplicationsConnAppIdScopesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connAppId** | **string** | The ID of the connected app | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectedApplicationsConnAppIdScopesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConnectedApplicationsConnAppIdScopesGet200Response**](ConnectedApplicationsConnAppIdScopesGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectedApplicationsConnAppIdScopesPut

> ConnectedApplicationsConnAppIdScopesPut(ctx, connAppId).ConnectedAppScopesPutBody(connectedAppScopesPutBody).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/connected_app"
)

func main() {
    connAppId := "connAppId_example" // string | The ID of the connected app
    connectedAppScopesPutBody := *openapiclient.NewConnectedAppScopesPutBody() // ConnectedAppScopesPutBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.ConnectedApplicationsConnAppIdScopesPut(context.Background(), connAppId).ConnectedAppScopesPutBody(connectedAppScopesPutBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ConnectedApplicationsConnAppIdScopesPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connAppId** | **string** | The ID of the connected app | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectedApplicationsConnAppIdScopesPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectedAppScopesPutBody** | [**ConnectedAppScopesPutBody**](ConnectedAppScopesPutBody.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectedApplicationsGet

> ConnectedApplicationsGet200Response ConnectedApplicationsGet(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/connected_app"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ConnectedApplicationsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ConnectedApplicationsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectedApplicationsGet`: ConnectedApplicationsGet200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ConnectedApplicationsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiConnectedApplicationsGetRequest struct via the builder pattern


### Return type

[**ConnectedApplicationsGet200Response**](ConnectedApplicationsGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectedApplicationsPost

> ConnectedAppRespExt ConnectedApplicationsPost(ctx).ConnectedAppCore(connectedAppCore).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/connected_app"
)

func main() {
    connectedAppCore := *openapiclient.NewConnectedAppCore() // ConnectedAppCore |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ConnectedApplicationsPost(context.Background()).ConnectedAppCore(connectedAppCore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ConnectedApplicationsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectedApplicationsPost`: ConnectedAppRespExt
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ConnectedApplicationsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConnectedApplicationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectedAppCore** | [**ConnectedAppCore**](ConnectedAppCore.md) |  | 

### Return type

[**ConnectedAppRespExt**](ConnectedAppRespExt.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

