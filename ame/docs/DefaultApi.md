# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com/mq/admin/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAME**](DefaultApi.md#CreateAME) | **Put** /organizations/{orgId}/environments/{envId}/regions/{regionId}/destinations/exchanges/{exchangeId} | 
[**DeleteAME**](DefaultApi.md#DeleteAME) | **Delete** /organizations/{orgId}/environments/{envId}/regions/{regionId}/destinations/exchanges/{exchangeId} | 
[**GetAME**](DefaultApi.md#GetAME) | **Get** /organizations/{orgId}/environments/{envId}/regions/{regionId}/destinations/exchanges/{exchangeId} | 
[**UpdateAME**](DefaultApi.md#UpdateAME) | **Patch** /organizations/{orgId}/environments/{envId}/regions/{regionId}/destinations/exchanges/{exchangeId} | 



## CreateAME

> Exchange CreateAME(ctx, orgId, envId, regionId, exchangeId).ExchangeBody(exchangeBody).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/ame"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    regionId := "regionId_example" // string | The region id
    exchangeId := "exchangeId_example" // string | The id of a specific exchange
    exchangeBody := *openapiclient.NewExchangeBody() // ExchangeBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateAME(context.Background(), orgId, envId, regionId, exchangeId).ExchangeBody(exchangeBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateAME``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAME`: Exchange
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateAME`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment id | 
**regionId** | **string** | The region id | 
**exchangeId** | **string** | The id of a specific exchange | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAMERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **exchangeBody** | [**ExchangeBody**](ExchangeBody.md) |  | 

### Return type

[**Exchange**](Exchange.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAME

> DeleteAME(ctx, orgId, envId, regionId, exchangeId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/ame"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    regionId := "regionId_example" // string | The region id
    exchangeId := "exchangeId_example" // string | The id of a specific exchange

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeleteAME(context.Background(), orgId, envId, regionId, exchangeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteAME``: %v\n", err)
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
**regionId** | **string** | The region id | 
**exchangeId** | **string** | The id of a specific exchange | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAMERequest struct via the builder pattern


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


## GetAME

> Exchange GetAME(ctx, orgId, envId, regionId, exchangeId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/ame"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    regionId := "regionId_example" // string | The region id
    exchangeId := "exchangeId_example" // string | The id of a specific exchange

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetAME(context.Background(), orgId, envId, regionId, exchangeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAME``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAME`: Exchange
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAME`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment id | 
**regionId** | **string** | The region id | 
**exchangeId** | **string** | The id of a specific exchange | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAMERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**Exchange**](Exchange.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAME

> Exchange UpdateAME(ctx, orgId, envId, regionId, exchangeId).ExchangeBody(exchangeBody).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/ame"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    regionId := "regionId_example" // string | The region id
    exchangeId := "exchangeId_example" // string | The id of a specific exchange
    exchangeBody := *openapiclient.NewExchangeBody() // ExchangeBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateAME(context.Background(), orgId, envId, regionId, exchangeId).ExchangeBody(exchangeBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateAME``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAME`: Exchange
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateAME`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment id | 
**regionId** | **string** | The region id | 
**exchangeId** | **string** | The id of a specific exchange | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAMERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **exchangeBody** | [**ExchangeBody**](ExchangeBody.md) |  | 

### Return type

[**Exchange**](Exchange.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

