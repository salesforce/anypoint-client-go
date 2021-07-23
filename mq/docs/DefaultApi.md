# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com/mq/admin/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsGet**](DefaultApi.md#V1OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsGet) | **Get** /v1/organizations/{orgId}/environments/{envId}/regions/{regionId}/destinations | 
[**V1OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdDelete**](DefaultApi.md#V1OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdDelete) | **Delete** /v1/organizations/{orgId}/environments/{envId}/regions/{regionId}/destinations/queues/{queueId} | 
[**V1OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdGet**](DefaultApi.md#V1OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdGet) | **Get** /v1/organizations/{orgId}/environments/{envId}/regions/{regionId}/destinations/queues/{queueId} | 
[**V1OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdPatch**](DefaultApi.md#V1OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdPatch) | **Patch** /v1/organizations/{orgId}/environments/{envId}/regions/{regionId}/destinations/queues/{queueId} | 
[**V1OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdPut**](DefaultApi.md#V1OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdPut) | **Put** /v1/organizations/{orgId}/environments/{envId}/regions/{regionId}/destinations/queues/{queueId} | 



## V1OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsGet

> []Queue V1OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsGet(ctx, orgId, envId, regionId).Execute()





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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.V1OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsGet(context.Background(), orgId, envId, regionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.V1OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsGet`: []Queue
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.V1OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsGet`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiV1OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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


## V1OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdDelete

> V1OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdDelete(ctx, orgId, envId, regionId, queueId).Execute()





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
    resp, r, err := api_client.DefaultApi.V1OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdDelete(context.Background(), orgId, envId, regionId, queueId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.V1OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdDeleteRequest struct via the builder pattern


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


## V1OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdGet

> Queue V1OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdGet(ctx, orgId, envId, regionId, queueId).Execute()





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
    resp, r, err := api_client.DefaultApi.V1OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdGet(context.Background(), orgId, envId, regionId, queueId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.V1OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdGet`: Queue
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.V1OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdGet`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiV1OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdGetRequest struct via the builder pattern


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


## V1OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdPatch

> Queue V1OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdPatch(ctx, orgId, envId, regionId, queueId).QueueOptional(queueOptional).Execute()





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
    queueOptional := *openapiclient.NewQueueOptional() // QueueOptional |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.V1OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdPatch(context.Background(), orgId, envId, regionId, queueId).QueueOptional(queueOptional).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.V1OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdPatch`: Queue
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.V1OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdPatch`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiV1OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **queueOptional** | [**QueueOptional**](QueueOptional.md) |  | 

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


## V1OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdPut

> Queue V1OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdPut(ctx, orgId, envId, regionId, queueId).QueueOptional(queueOptional).Execute()





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
    queueOptional := *openapiclient.NewQueueOptional() // QueueOptional |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.V1OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdPut(context.Background(), orgId, envId, regionId, queueId).QueueOptional(queueOptional).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.V1OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdPut`: Queue
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.V1OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdPut`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiV1OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **queueOptional** | [**QueueOptional**](QueueOptional.md) |  | 

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

