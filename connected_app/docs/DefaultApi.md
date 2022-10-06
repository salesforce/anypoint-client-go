# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com/cloudhub/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountsApiConnectedApplicationsConnAppIdDelete**](DefaultApi.md#AccountsApiConnectedApplicationsConnAppIdDelete) | **Delete** /accounts/api/connectedApplications/{connAppId} | 
[**AccountsApiConnectedApplicationsConnAppIdGet**](DefaultApi.md#AccountsApiConnectedApplicationsConnAppIdGet) | **Get** /accounts/api/connectedApplications/{connAppId} | 
[**AccountsApiConnectedApplicationsConnAppIdPatch**](DefaultApi.md#AccountsApiConnectedApplicationsConnAppIdPatch) | **Patch** /accounts/api/connectedApplications/{connAppId} | 
[**AccountsApiConnectedApplicationsConnAppIdScopesGet**](DefaultApi.md#AccountsApiConnectedApplicationsConnAppIdScopesGet) | **Get** /accounts/api/connectedApplications/{connAppId}/scopes | 
[**AccountsApiConnectedApplicationsConnAppIdScopesPut**](DefaultApi.md#AccountsApiConnectedApplicationsConnAppIdScopesPut) | **Put** /accounts/api/connectedApplications/{connAppId}/scopes | 
[**AccountsApiConnectedApplicationsGet**](DefaultApi.md#AccountsApiConnectedApplicationsGet) | **Get** /accounts/api/connectedApplications | 
[**AccountsApiConnectedApplicationsPost**](DefaultApi.md#AccountsApiConnectedApplicationsPost) | **Post** /accounts/api/connectedApplications | 



## AccountsApiConnectedApplicationsConnAppIdDelete

> AccountsApiConnectedApplicationsConnAppIdDelete(ctx, connAppId).Execute()





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
    connAppId := "connAppId_example" // string | The ID of the connected app

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AccountsApiConnectedApplicationsConnAppIdDelete(context.Background(), connAppId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AccountsApiConnectedApplicationsConnAppIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAccountsApiConnectedApplicationsConnAppIdDeleteRequest struct via the builder pattern


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


## AccountsApiConnectedApplicationsConnAppIdGet

> ConnectedAppRespExt AccountsApiConnectedApplicationsConnAppIdGet(ctx, connAppId).Execute()





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
    connAppId := "connAppId_example" // string | The ID of the connected app

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AccountsApiConnectedApplicationsConnAppIdGet(context.Background(), connAppId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AccountsApiConnectedApplicationsConnAppIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsApiConnectedApplicationsConnAppIdGet`: ConnectedAppRespExt
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AccountsApiConnectedApplicationsConnAppIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connAppId** | **string** | The ID of the connected app | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsApiConnectedApplicationsConnAppIdGetRequest struct via the builder pattern


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


## AccountsApiConnectedApplicationsConnAppIdPatch

> ConnectedAppRespExt AccountsApiConnectedApplicationsConnAppIdPatch(ctx, connAppId).ConnectedAppPatchExt(connectedAppPatchExt).Execute()





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
    connAppId := "connAppId_example" // string | The ID of the connected app
    connectedAppPatchExt := *openapiclient.NewConnectedAppPatchExt() // ConnectedAppPatchExt |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AccountsApiConnectedApplicationsConnAppIdPatch(context.Background(), connAppId).ConnectedAppPatchExt(connectedAppPatchExt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AccountsApiConnectedApplicationsConnAppIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsApiConnectedApplicationsConnAppIdPatch`: ConnectedAppRespExt
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AccountsApiConnectedApplicationsConnAppIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connAppId** | **string** | The ID of the connected app | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsApiConnectedApplicationsConnAppIdPatchRequest struct via the builder pattern


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


## AccountsApiConnectedApplicationsConnAppIdScopesGet

> InlineResponse2001 AccountsApiConnectedApplicationsConnAppIdScopesGet(ctx, connAppId).Execute()





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
    connAppId := "connAppId_example" // string | The ID of the connected app

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AccountsApiConnectedApplicationsConnAppIdScopesGet(context.Background(), connAppId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AccountsApiConnectedApplicationsConnAppIdScopesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsApiConnectedApplicationsConnAppIdScopesGet`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AccountsApiConnectedApplicationsConnAppIdScopesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connAppId** | **string** | The ID of the connected app | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsApiConnectedApplicationsConnAppIdScopesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountsApiConnectedApplicationsConnAppIdScopesPut

> AccountsApiConnectedApplicationsConnAppIdScopesPut(ctx, connAppId).ConnectedAppScopesPutBody(connectedAppScopesPutBody).Execute()





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
    connAppId := "connAppId_example" // string | The ID of the connected app
    connectedAppScopesPutBody := *openapiclient.NewConnectedAppScopesPutBody() // ConnectedAppScopesPutBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AccountsApiConnectedApplicationsConnAppIdScopesPut(context.Background(), connAppId).ConnectedAppScopesPutBody(connectedAppScopesPutBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AccountsApiConnectedApplicationsConnAppIdScopesPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAccountsApiConnectedApplicationsConnAppIdScopesPutRequest struct via the builder pattern


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


## AccountsApiConnectedApplicationsGet

> InlineResponse200 AccountsApiConnectedApplicationsGet(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AccountsApiConnectedApplicationsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AccountsApiConnectedApplicationsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsApiConnectedApplicationsGet`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AccountsApiConnectedApplicationsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsApiConnectedApplicationsGetRequest struct via the builder pattern


### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountsApiConnectedApplicationsPost

> ConnectedAppRespExt AccountsApiConnectedApplicationsPost(ctx).ConnectedAppCore(connectedAppCore).Execute()





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
    connectedAppCore := *openapiclient.NewConnectedAppCore() // ConnectedAppCore |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AccountsApiConnectedApplicationsPost(context.Background()).ConnectedAppCore(connectedAppCore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AccountsApiConnectedApplicationsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsApiConnectedApplicationsPost`: ConnectedAppRespExt
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AccountsApiConnectedApplicationsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountsApiConnectedApplicationsPostRequest struct via the builder pattern


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

