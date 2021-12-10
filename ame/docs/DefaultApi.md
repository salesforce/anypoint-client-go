# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com/mq/admin/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsExchangesExchangeIdDelete**](DefaultApi.md#OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsExchangesExchangeIdDelete) | **Delete** /organizations/{orgId}/environments/{envId}/regions/{regionId}/destinations/exchanges/{exchangeId} | 
[**OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsExchangesExchangeIdGet**](DefaultApi.md#OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsExchangesExchangeIdGet) | **Get** /organizations/{orgId}/environments/{envId}/regions/{regionId}/destinations/exchanges/{exchangeId} | 
[**OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsExchangesExchangeIdPatch**](DefaultApi.md#OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsExchangesExchangeIdPatch) | **Patch** /organizations/{orgId}/environments/{envId}/regions/{regionId}/destinations/exchanges/{exchangeId} | 
[**OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsExchangesExchangeIdPut**](DefaultApi.md#OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsExchangesExchangeIdPut) | **Put** /organizations/{orgId}/environments/{envId}/regions/{regionId}/destinations/exchanges/{exchangeId} | 



## OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsExchangesExchangeIdDelete

> OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsExchangesExchangeIdDelete(ctx, orgId, envId, regionId, exchangeId).Execute()





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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsExchangesExchangeIdDelete(context.Background(), orgId, envId, regionId, exchangeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsExchangesExchangeIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiOrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsExchangesExchangeIdDeleteRequest struct via the builder pattern


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


## OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsExchangesExchangeIdGet

> Exchange OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsExchangesExchangeIdGet(ctx, orgId, envId, regionId, exchangeId).Execute()





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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsExchangesExchangeIdGet(context.Background(), orgId, envId, regionId, exchangeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsExchangesExchangeIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsExchangesExchangeIdGet`: Exchange
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsExchangesExchangeIdGet`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiOrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsExchangesExchangeIdGetRequest struct via the builder pattern


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


## OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsExchangesExchangeIdPatch

> Exchange OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsExchangesExchangeIdPatch(ctx, orgId, envId, regionId, exchangeId).ExchangeBody(exchangeBody).Execute()





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
    exchangeBody := *openapiclient.NewExchangeBody() // ExchangeBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsExchangesExchangeIdPatch(context.Background(), orgId, envId, regionId, exchangeId).ExchangeBody(exchangeBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsExchangesExchangeIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsExchangesExchangeIdPatch`: Exchange
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsExchangesExchangeIdPatch`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiOrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsExchangesExchangeIdPatchRequest struct via the builder pattern


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


## OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsExchangesExchangeIdPut

> Exchange OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsExchangesExchangeIdPut(ctx, orgId, envId, regionId, exchangeId).ExchangeBody(exchangeBody).Execute()





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
    exchangeBody := *openapiclient.NewExchangeBody() // ExchangeBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsExchangesExchangeIdPut(context.Background(), orgId, envId, regionId, exchangeId).ExchangeBody(exchangeBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsExchangesExchangeIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsExchangesExchangeIdPut`: Exchange
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsExchangesExchangeIdPut`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiOrganizationsOrgIdEnvironmentsEnvIdRegionsRegionIdDestinationsExchangesExchangeIdPutRequest struct via the builder pattern


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

