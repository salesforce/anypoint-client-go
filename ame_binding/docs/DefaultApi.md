# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com/mq/admin/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAMEBinding**](DefaultApi.md#CreateAMEBinding) | **Put** /organizations/{orgId}/environments/{envId}/regions/{regionId}/bindings/exchanges/{exchangeId}/queues/{queueId} | 
[**CreateAMEBindingRule**](DefaultApi.md#CreateAMEBindingRule) | **Put** /organizations/{orgId}/environments/{envId}/regions/{regionId}/bindings/exchanges/{exchangeId}/queues/{queueId}/rules/routing | 
[**DeleteAMEBinding**](DefaultApi.md#DeleteAMEBinding) | **Delete** /organizations/{orgId}/environments/{envId}/regions/{regionId}/bindings/exchanges/{exchangeId}/queues/{queueId} | 
[**DeleteAMEBindingRule**](DefaultApi.md#DeleteAMEBindingRule) | **Delete** /organizations/{orgId}/environments/{envId}/regions/{regionId}/bindings/exchanges/{exchangeId}/queues/{queueId}/rules/routing | 
[**GetAMEBinding**](DefaultApi.md#GetAMEBinding) | **Get** /organizations/{orgId}/environments/{envId}/regions/{regionId}/bindings/exchanges/{exchangeId}/queues/{queueId} | 



## CreateAMEBinding

> ExchangeBinding CreateAMEBinding(ctx, orgId, envId, regionId, exchangeId, queueId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/ame_binding"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    regionId := "regionId_example" // string | The region id
    exchangeId := "exchangeId_example" // string | The id of a specific exchange
    queueId := "queueId_example" // string | The id of a specific exchange queue binding

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateAMEBinding(context.Background(), orgId, envId, regionId, exchangeId, queueId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateAMEBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAMEBinding`: ExchangeBinding
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateAMEBinding`: %v\n", resp)
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
**queueId** | **string** | The id of a specific exchange queue binding | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAMEBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

[**ExchangeBinding**](ExchangeBinding.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAMEBindingRule

> ExchangeBindingRules CreateAMEBindingRule(ctx, orgId, envId, regionId, exchangeId, queueId).AMEBindingRuleBody(aMEBindingRuleBody).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/ame_binding"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    regionId := "regionId_example" // string | The region id
    exchangeId := "exchangeId_example" // string | The id of a specific exchange
    queueId := "queueId_example" // string | The id of a specific exchange queue binding
    aMEBindingRuleBody := *openapiclient.NewAMEBindingRuleBody() // AMEBindingRuleBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateAMEBindingRule(context.Background(), orgId, envId, regionId, exchangeId, queueId).AMEBindingRuleBody(aMEBindingRuleBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateAMEBindingRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAMEBindingRule`: ExchangeBindingRules
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateAMEBindingRule`: %v\n", resp)
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
**queueId** | **string** | The id of a specific exchange queue binding | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAMEBindingRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **aMEBindingRuleBody** | [**AMEBindingRuleBody**](AMEBindingRuleBody.md) |  | 

### Return type

[**ExchangeBindingRules**](ExchangeBindingRules.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAMEBinding

> DeleteAMEBinding(ctx, orgId, envId, regionId, exchangeId, queueId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/ame_binding"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    regionId := "regionId_example" // string | The region id
    exchangeId := "exchangeId_example" // string | The id of a specific exchange
    queueId := "queueId_example" // string | The id of a specific exchange queue binding

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeleteAMEBinding(context.Background(), orgId, envId, regionId, exchangeId, queueId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteAMEBinding``: %v\n", err)
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
**queueId** | **string** | The id of a specific exchange queue binding | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAMEBindingRequest struct via the builder pattern


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


## DeleteAMEBindingRule

> DeleteAMEBindingRule(ctx, orgId, envId, regionId, exchangeId, queueId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/ame_binding"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    regionId := "regionId_example" // string | The region id
    exchangeId := "exchangeId_example" // string | The id of a specific exchange
    queueId := "queueId_example" // string | The id of a specific exchange queue binding

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeleteAMEBindingRule(context.Background(), orgId, envId, regionId, exchangeId, queueId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteAMEBindingRule``: %v\n", err)
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
**queueId** | **string** | The id of a specific exchange queue binding | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAMEBindingRuleRequest struct via the builder pattern


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


## GetAMEBinding

> ExchangeBindingWithRules GetAMEBinding(ctx, orgId, envId, regionId, exchangeId, queueId).Inclusion(inclusion).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/ame_binding"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    regionId := "regionId_example" // string | The region id
    exchangeId := "exchangeId_example" // string | The id of a specific exchange
    queueId := "queueId_example" // string | The id of a specific exchange queue binding
    inclusion := "inclusion_example" // string | Defines what to fetch (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetAMEBinding(context.Background(), orgId, envId, regionId, exchangeId, queueId).Inclusion(inclusion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAMEBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAMEBinding`: ExchangeBindingWithRules
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAMEBinding`: %v\n", resp)
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
**queueId** | **string** | The id of a specific exchange queue binding | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAMEBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **inclusion** | **string** | Defines what to fetch | 

### Return type

[**ExchangeBindingWithRules**](ExchangeBindingWithRules.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

