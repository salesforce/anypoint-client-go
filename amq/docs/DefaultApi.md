# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com/mq/admin/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsGet**](DefaultApi.md#OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsGet) | **Get** /organizations/{orgId}/environments/{envId}/regions/{regionId}/destinations | 
[**OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdDelete**](DefaultApi.md#OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdDelete) | **Delete** /organizations/{orgId}/environments/{envId}/regions/{regionId}/destinations/queues/{queueId} | 
[**OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdGet**](DefaultApi.md#OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdGet) | **Get** /organizations/{orgId}/environments/{envId}/regions/{regionId}/destinations/queues/{queueId} | 
[**OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdPatch**](DefaultApi.md#OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdPatch) | **Patch** /organizations/{orgId}/environments/{envId}/regions/{regionId}/destinations/queues/{queueId} | 
[**OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdPut**](DefaultApi.md#OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdPut) | **Put** /organizations/{orgId}/environments/{envId}/regions/{regionId}/destinations/queues/{queueId} | 



## OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsGet

> []Queue OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsGet(ctx, orgId, envId, regionId).Execute()





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
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsGet(context.Background(), orgId, envId, regionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsGet`: []Queue
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsGet`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiOrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsGetRequest struct via the builder pattern


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


## OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdDelete

> OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdDelete(ctx, orgId, envId, regionId, queueId).Execute()





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
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdDelete(context.Background(), orgId, envId, regionId, queueId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiOrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdDeleteRequest struct via the builder pattern


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


## OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdGet

> Queue OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdGet(ctx, orgId, envId, regionId, queueId).Execute()





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
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdGet(context.Background(), orgId, envId, regionId, queueId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdGet`: Queue
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdGet`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiOrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdGetRequest struct via the builder pattern


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


## OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdPatch

> Queue OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdPatch(ctx, orgId, envId, regionId, queueId).QueueBody(queueBody).Execute()





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
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdPatch(context.Background(), orgId, envId, regionId, queueId).QueueBody(queueBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdPatch`: Queue
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdPatch`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiOrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdPatchRequest struct via the builder pattern


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


## OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdPut

> Queue OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdPut(ctx, orgId, envId, regionId, queueId).QueueBody(queueBody).Execute()





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
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdPut(context.Background(), orgId, envId, regionId, queueId).QueueBody(queueBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdPut`: Queue
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdPut`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiOrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsQueuesQueueIdPutRequest struct via the builder pattern


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

