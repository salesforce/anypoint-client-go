# \DefaultAPI

All URIs are relative to *https://anypoint.mulesoft.com/apimanager*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiInstanceTier**](DefaultAPI.md#CreateApiInstanceTier) | **Post** /api/v1/organizations/{orgId}/environments/{envId}/apis/{apiId}/tiers | Creates a new SLA tier for a given API instance
[**DeleteApiInstanceTier**](DefaultAPI.md#DeleteApiInstanceTier) | **Delete** /api/v1/organizations/{orgId}/environments/{envId}/apis/{apiId}/tiers/{tierId} | Deletes an SLA tier for a given API instance
[**GetApiInstanceTiers**](DefaultAPI.md#GetApiInstanceTiers) | **Get** /api/v1/organizations/{orgId}/environments/{envId}/apis/{apiId}/tiers | Retrieves all SLA tiers for a given API instance
[**UpdateApiInstanceTier**](DefaultAPI.md#UpdateApiInstanceTier) | **Put** /api/v1/organizations/{orgId}/environments/{envId}/apis/{apiId}/tiers/{tierId} | Updates an SLA tier for a given API instance



## CreateApiInstanceTier

> SlaTier CreateApiInstanceTier(ctx, orgId, envId, apiId).SlaTierPostBody(slaTierPostBody).Execute()

Creates a new SLA tier for a given API instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/apim_tier"
)

func main() {
	orgId := "orgId_example" // string | The organization Id
	envId := "envId_example" // string | The environment id
	apiId := "apiId_example" // string | The API Manager instance id
	slaTierPostBody := *openapiclient.NewSlaTierPostBody() // SlaTierPostBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateApiInstanceTier(context.Background(), orgId, envId, apiId).SlaTierPostBody(slaTierPostBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateApiInstanceTier``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApiInstanceTier`: SlaTier
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateApiInstanceTier`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment id | 
**apiId** | **string** | The API Manager instance id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiInstanceTierRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **slaTierPostBody** | [**SlaTierPostBody**](SlaTierPostBody.md) |  | 

### Return type

[**SlaTier**](SlaTier.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApiInstanceTier

> DeleteApiInstanceTier(ctx, orgId, envId, apiId, tierId).Execute()

Deletes an SLA tier for a given API instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/apim_tier"
)

func main() {
	orgId := "orgId_example" // string | The organization Id
	envId := "envId_example" // string | The environment id
	apiId := "apiId_example" // string | The API Manager instance id
	tierId := int32(56) // int32 | The SLA tier id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteApiInstanceTier(context.Background(), orgId, envId, apiId, tierId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteApiInstanceTier``: %v\n", err)
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
**apiId** | **string** | The API Manager instance id | 
**tierId** | **int32** | The SLA tier id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiInstanceTierRequest struct via the builder pattern


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


## GetApiInstanceTiers

> GetSlaTiersResponse GetApiInstanceTiers(ctx, orgId, envId, apiId).Limit(limit).Offset(offset).Execute()

Retrieves all SLA tiers for a given API instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/apim_tier"
)

func main() {
	orgId := "orgId_example" // string | The organization Id
	envId := "envId_example" // string | The environment id
	apiId := "apiId_example" // string | The API Manager instance id
	limit := int32(56) // int32 | The maximum number of tiers to return (optional)
	offset := int32(56) // int32 | The offset of the first tier to return (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetApiInstanceTiers(context.Background(), orgId, envId, apiId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetApiInstanceTiers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiInstanceTiers`: GetSlaTiersResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetApiInstanceTiers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment id | 
**apiId** | **string** | The API Manager instance id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiInstanceTiersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **int32** | The maximum number of tiers to return | 
 **offset** | **int32** | The offset of the first tier to return | 

### Return type

[**GetSlaTiersResponse**](GetSlaTiersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApiInstanceTier

> SlaTier UpdateApiInstanceTier(ctx, orgId, envId, apiId, tierId).SlaTierPutBody(slaTierPutBody).Execute()

Updates an SLA tier for a given API instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/apim_tier"
)

func main() {
	orgId := "orgId_example" // string | The organization Id
	envId := "envId_example" // string | The environment id
	apiId := "apiId_example" // string | The API Manager instance id
	tierId := int32(56) // int32 | The SLA tier id
	slaTierPutBody := *openapiclient.NewSlaTierPutBody() // SlaTierPutBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateApiInstanceTier(context.Background(), orgId, envId, apiId, tierId).SlaTierPutBody(slaTierPutBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateApiInstanceTier``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApiInstanceTier`: SlaTier
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateApiInstanceTier`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment id | 
**apiId** | **string** | The API Manager instance id | 
**tierId** | **int32** | The SLA tier id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiInstanceTierRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **slaTierPutBody** | [**SlaTierPutBody**](SlaTierPutBody.md) |  | 

### Return type

[**SlaTier**](SlaTier.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

