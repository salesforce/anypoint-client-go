# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com/mq/admin/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdBindingsExchangesExchangeIdQueuesQueueIdDelete**](DefaultApi.md#OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdBindingsExchangesExchangeIdQueuesQueueIdDelete) | **Delete** /organizations/{orgId}/environments/{envId}/regions/{regionId}/bindings/exchanges/{exchangeId}/queues/{queueId} | 
[**OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdBindingsExchangesExchangeIdQueuesQueueIdGet**](DefaultApi.md#OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdBindingsExchangesExchangeIdQueuesQueueIdGet) | **Get** /organizations/{orgId}/environments/{envId}/regions/{regionId}/bindings/exchanges/{exchangeId}/queues/{queueId} | 
[**OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdBindingsExchangesExchangeIdQueuesQueueIdPut**](DefaultApi.md#OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdBindingsExchangesExchangeIdQueuesQueueIdPut) | **Put** /organizations/{orgId}/environments/{envId}/regions/{regionId}/bindings/exchanges/{exchangeId}/queues/{queueId} | 



## OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdBindingsExchangesExchangeIdQueuesQueueIdDelete

> OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdBindingsExchangesExchangeIdQueuesQueueIdDelete(ctx, orgId, envId, regionId, exchangeId, queueId).Execute()





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
    exchangeId := "exchangeId_example" // string | The id of a specific exchange
    queueId := "queueId_example" // string | The id of a specific exchange queue binding

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdBindingsExchangesExchangeIdQueuesQueueIdDelete(context.Background(), orgId, envId, regionId, exchangeId, queueId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdBindingsExchangesExchangeIdQueuesQueueIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiOrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdBindingsExchangesExchangeIdQueuesQueueIdDeleteRequest struct via the builder pattern


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


## OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdBindingsExchangesExchangeIdQueuesQueueIdGet

> ExchangeBinding OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdBindingsExchangesExchangeIdQueuesQueueIdGet(ctx, orgId, envId, regionId, exchangeId, queueId).Execute()





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
    exchangeId := "exchangeId_example" // string | The id of a specific exchange
    queueId := "queueId_example" // string | The id of a specific exchange queue binding

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdBindingsExchangesExchangeIdQueuesQueueIdGet(context.Background(), orgId, envId, regionId, exchangeId, queueId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdBindingsExchangesExchangeIdQueuesQueueIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdBindingsExchangesExchangeIdQueuesQueueIdGet`: ExchangeBinding
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdBindingsExchangesExchangeIdQueuesQueueIdGet`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiOrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdBindingsExchangesExchangeIdQueuesQueueIdGetRequest struct via the builder pattern


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


## OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdBindingsExchangesExchangeIdQueuesQueueIdPut

> ExchangeBinding OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdBindingsExchangesExchangeIdQueuesQueueIdPut(ctx, orgId, envId, regionId, exchangeId, queueId).Execute()





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
    exchangeId := "exchangeId_example" // string | The id of a specific exchange
    queueId := "queueId_example" // string | The id of a specific exchange queue binding

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdBindingsExchangesExchangeIdQueuesQueueIdPut(context.Background(), orgId, envId, regionId, exchangeId, queueId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdBindingsExchangesExchangeIdQueuesQueueIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdBindingsExchangesExchangeIdQueuesQueueIdPut`: ExchangeBinding
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdBindingsExchangesExchangeIdQueuesQueueIdPut`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiOrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdBindingsExchangesExchangeIdQueuesQueueIdPutRequest struct via the builder pattern


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

