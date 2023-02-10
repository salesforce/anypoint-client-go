# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com/mq/admin/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAMQ**](DefaultApi.md#CreateAMQ) | **Put** /organizations/{orgId}/environments/{envId}/regions/{regionId}/destinations/queues/{queueId} | 
[**DeleteAMQ**](DefaultApi.md#DeleteAMQ) | **Delete** /organizations/{orgId}/environments/{envId}/regions/{regionId}/destinations/queues/{queueId} | 
[**GetAMQ**](DefaultApi.md#GetAMQ) | **Get** /organizations/{orgId}/environments/{envId}/regions/{regionId}/destinations/queues/{queueId} | 
[**GetAMQList**](DefaultApi.md#GetAMQList) | **Get** /organizations/{orgId}/environments/{envId}/regions/{regionId}/destinations | 
[**UpdateAMQ**](DefaultApi.md#UpdateAMQ) | **Patch** /organizations/{orgId}/environments/{envId}/regions/{regionId}/destinations/queues/{queueId} | 



## CreateAMQ

> Queue CreateAMQ(ctx, orgId, envId, regionId, queueId).QueueBody(queueBody).Execute()





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
    regionId := "regionId_example" // string | The region id
    queueId := "queueId_example" // string | The id of a specific queue
    queueBody := *openapiclient.NewQueueBody() // QueueBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateAMQ(context.Background(), orgId, envId, regionId, queueId).QueueBody(queueBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateAMQ``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAMQ`: Queue
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateAMQ`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment id | 
**regionId** | **string** | The region id | 
**queueId** | **string** | The id of a specific queue | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAMQRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **queueBody** | [**QueueBody**](QueueBody.md) |  | 

### Return type

[**Queue**](Queue.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAMQ

> DeleteAMQ(ctx, orgId, envId, regionId, queueId).Execute()





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
    regionId := "regionId_example" // string | The region id
    queueId := "queueId_example" // string | The id of a specific queue

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteAMQ(context.Background(), orgId, envId, regionId, queueId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteAMQ``: %v\n", err)
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
**queueId** | **string** | The id of a specific queue | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAMQRequest struct via the builder pattern


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


## GetAMQ

> Queue GetAMQ(ctx, orgId, envId, regionId, queueId).Execute()





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
    regionId := "regionId_example" // string | The region id
    queueId := "queueId_example" // string | The id of a specific queue

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetAMQ(context.Background(), orgId, envId, regionId, queueId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAMQ``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAMQ`: Queue
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAMQ`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment id | 
**regionId** | **string** | The region id | 
**queueId** | **string** | The id of a specific queue | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAMQRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**Queue**](Queue.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAMQList

> []Queue GetAMQList(ctx, orgId, envId, regionId).Inclusion(inclusion).DestinationType(destinationType).Offset(offset).Limit(limit).StartsWith(startsWith).DestinationIds(destinationIds).Execute()





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
    regionId := "regionId_example" // string | The region id for MQ
    inclusion := "inclusion_example" // string | Defines what to fetch (optional)
    destinationType := "destinationType_example" // string | Defines what to fetch (optional)
    offset := int32(56) // int32 | Skip over a number of elements by specifying an offset value for the query. (optional) (default to 0)
    limit := int32(56) // int32 | Limit the number of elements in the response. (optional) (default to 20)
    startsWith := "startsWith_example" // string | Searchs the field from the left using the passed string. (optional)
    destinationIds := []string{"Inner_example"} // []string | Includes only results with the given Ids. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetAMQList(context.Background(), orgId, envId, regionId).Inclusion(inclusion).DestinationType(destinationType).Offset(offset).Limit(limit).StartsWith(startsWith).DestinationIds(destinationIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAMQList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAMQList`: []Queue
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAMQList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment id | 
**regionId** | **string** | The region id for MQ | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAMQListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **inclusion** | **string** | Defines what to fetch | 
 **destinationType** | **string** | Defines what to fetch | 
 **offset** | **int32** | Skip over a number of elements by specifying an offset value for the query. | [default to 0]
 **limit** | **int32** | Limit the number of elements in the response. | [default to 20]
 **startsWith** | **string** | Searchs the field from the left using the passed string. | 
 **destinationIds** | **[]string** | Includes only results with the given Ids. | 

### Return type

[**[]Queue**](Queue.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAMQ

> Queue UpdateAMQ(ctx, orgId, envId, regionId, queueId).QueueBody(queueBody).Execute()





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
    regionId := "regionId_example" // string | The region id
    queueId := "queueId_example" // string | The id of a specific queue
    queueBody := *openapiclient.NewQueueBody() // QueueBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateAMQ(context.Background(), orgId, envId, regionId, queueId).QueueBody(queueBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateAMQ``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAMQ`: Queue
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateAMQ`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment id | 
**regionId** | **string** | The region id | 
**queueId** | **string** | The id of a specific queue | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAMQRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **queueBody** | [**QueueBody**](QueueBody.md) |  | 

### Return type

[**Queue**](Queue.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

