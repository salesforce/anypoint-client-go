# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com/accounts/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsOrgIdEnvironmentsEnvironmentIdDelete**](DefaultApi.md#OrganizationsOrgIdEnvironmentsEnvironmentIdDelete) | **Delete** /organizations/{orgId}/environments/{environmentId} | 
[**OrganizationsOrgIdEnvironmentsEnvironmentIdGet**](DefaultApi.md#OrganizationsOrgIdEnvironmentsEnvironmentIdGet) | **Get** /organizations/{orgId}/environments/{environmentId} | 
[**OrganizationsOrgIdEnvironmentsEnvironmentIdPut**](DefaultApi.md#OrganizationsOrgIdEnvironmentsEnvironmentIdPut) | **Put** /organizations/{orgId}/environments/{environmentId} | 
[**OrganizationsOrgIdEnvironmentsGet**](DefaultApi.md#OrganizationsOrgIdEnvironmentsGet) | **Get** /organizations/{orgId}/environments | 
[**OrganizationsOrgIdEnvironmentsPost**](DefaultApi.md#OrganizationsOrgIdEnvironmentsPost) | **Post** /organizations/{orgId}/environments | 



## OrganizationsOrgIdEnvironmentsEnvironmentIdDelete

> OrganizationsOrgIdEnvironmentsEnvironmentIdDelete(ctx, orgId, environmentId).Execute()





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
    orgId := "orgId_example" // string | The ID of the organization in GUID format
    environmentId := "environmentId_example" // string | The id of an environment

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdEnvironmentsEnvironmentIdDelete(context.Background(), orgId, environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdEnvironmentsEnvironmentIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**environmentId** | **string** | The id of an environment | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdEnvironmentsEnvironmentIdDeleteRequest struct via the builder pattern


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


## OrganizationsOrgIdEnvironmentsEnvironmentIdGet

> Env OrganizationsOrgIdEnvironmentsEnvironmentIdGet(ctx, orgId, environmentId).Execute()





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
    orgId := "orgId_example" // string | The ID of the organization in GUID format
    environmentId := "environmentId_example" // string | The id of an environment

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdEnvironmentsEnvironmentIdGet(context.Background(), orgId, environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdEnvironmentsEnvironmentIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdEnvironmentsEnvironmentIdGet`: Env
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdEnvironmentsEnvironmentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**environmentId** | **string** | The id of an environment | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdEnvironmentsEnvironmentIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Env**](Env.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrgIdEnvironmentsEnvironmentIdPut

> Env OrganizationsOrgIdEnvironmentsEnvironmentIdPut(ctx, orgId, environmentId).EnvCore(envCore).Execute()





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
    orgId := "orgId_example" // string | The ID of the organization in GUID format
    environmentId := "environmentId_example" // string | The id of an environment
    envCore := *openapiclient.NewEnvCore() // EnvCore |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdEnvironmentsEnvironmentIdPut(context.Background(), orgId, environmentId).EnvCore(envCore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdEnvironmentsEnvironmentIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdEnvironmentsEnvironmentIdPut`: Env
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdEnvironmentsEnvironmentIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**environmentId** | **string** | The id of an environment | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdEnvironmentsEnvironmentIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envCore** | [**EnvCore**](EnvCore.md) |  | 

### Return type

[**Env**](Env.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrgIdEnvironmentsGet

> InlineResponse200 OrganizationsOrgIdEnvironmentsGet(ctx, orgId).Execute()





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
    orgId := "orgId_example" // string | The ID of the organization in GUID format

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdEnvironmentsGet(context.Background(), orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdEnvironmentsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdEnvironmentsGet`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdEnvironmentsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdEnvironmentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## OrganizationsOrgIdEnvironmentsPost

> Env OrganizationsOrgIdEnvironmentsPost(ctx, orgId).InlineObject(inlineObject).Execute()





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
    orgId := "orgId_example" // string | The ID of the organization in GUID format
    inlineObject := *openapiclient.NewInlineObject() // InlineObject |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdEnvironmentsPost(context.Background(), orgId).InlineObject(inlineObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdEnvironmentsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdEnvironmentsPost`: Env
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdEnvironmentsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdEnvironmentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject** | [**InlineObject**](InlineObject.md) |  | 

### Return type

[**Env**](Env.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

