# \DefaultAPI

All URIs are relative to *https://anypoint.mulesoft.com/exchange/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteExchangeClientApp**](DefaultAPI.md#DeleteExchangeClientApp) | **Delete** /organizations/{orgId}/applications/{appId} | Delete Exchange Client App by ID
[**GetExchangeClientApp**](DefaultAPI.md#GetExchangeClientApp) | **Get** /organizations/{orgId}/applications/{appId} | Get Exchange Client App by ID
[**GetExchangeClientAppContracts**](DefaultAPI.md#GetExchangeClientAppContracts) | **Get** /organizations/{orgId}/applications/{appId}/contracts | Get Exchange Client Application Contracts
[**GetExchangeClientApps**](DefaultAPI.md#GetExchangeClientApps) | **Get** /organizations/{orgId}/applications | Read Exchange Client Apps
[**PostExchangeClientApp**](DefaultAPI.md#PostExchangeClientApp) | **Post** /organizations/{orgId}/applications | Create a new Exchange Client App



## DeleteExchangeClientApp

> DeleteExchangeClientApp(ctx, orgId, appId).Execute()

Delete Exchange Client App by ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/exchange_client_apps"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	appId := int32(56) // int32 | The ID of the Exchange App

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteExchangeClientApp(context.Background(), orgId, appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteExchangeClientApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**appId** | **int32** | The ID of the Exchange App | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteExchangeClientAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExchangeClientApp

> ClientApp GetExchangeClientApp(ctx, orgId, appId).Execute()

Get Exchange Client App by ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/exchange_client_apps"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	appId := int32(56) // int32 | The ID of the Exchange App

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetExchangeClientApp(context.Background(), orgId, appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetExchangeClientApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExchangeClientApp`: ClientApp
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetExchangeClientApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**appId** | **int32** | The ID of the Exchange App | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExchangeClientAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClientApp**](ClientApp.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExchangeClientAppContracts

> []ClientAppContract GetExchangeClientAppContracts(ctx, orgId, appId).IncludeContractsForApiVersion(includeContractsForApiVersion).Execute()

Get Exchange Client Application Contracts



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/exchange_client_apps"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	appId := int32(56) // int32 | The ID of the Exchange App
	includeContractsForApiVersion := int32(56) // int32 | This field is used to filter by API version (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetExchangeClientAppContracts(context.Background(), orgId, appId).IncludeContractsForApiVersion(includeContractsForApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetExchangeClientAppContracts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExchangeClientAppContracts`: []ClientAppContract
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetExchangeClientAppContracts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**appId** | **int32** | The ID of the Exchange App | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExchangeClientAppContractsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeContractsForApiVersion** | **int32** | This field is used to filter by API version | 

### Return type

[**[]ClientAppContract**](ClientAppContract.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExchangeClientApps

> []ClientApp GetExchangeClientApps(ctx, orgId).TargetAdminSite(targetAdminSite).Query(query).Offset(offset).Limit(limit).Execute()

Read Exchange Client Apps



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/exchange_client_apps"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	targetAdminSite := true // bool | MUST be set to true in order to get the Client ID and Secret. Without this query parameter, the same call will return all the Client Applications but not the Client ID and Secret (and other info). (optional)
	query := "query_example" // string | Filter results that matches the input (optional)
	offset := int32(56) // int32 | The offset specifies the offset of the first row to return (optional)
	limit := int32(56) // int32 | Amount of objects retrieved in the response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetExchangeClientApps(context.Background(), orgId).TargetAdminSite(targetAdminSite).Query(query).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetExchangeClientApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExchangeClientApps`: []ClientApp
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetExchangeClientApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExchangeClientAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **targetAdminSite** | **bool** | MUST be set to true in order to get the Client ID and Secret. Without this query parameter, the same call will return all the Client Applications but not the Client ID and Secret (and other info). | 
 **query** | **string** | Filter results that matches the input | 
 **offset** | **int32** | The offset specifies the offset of the first row to return | 
 **limit** | **int32** | Amount of objects retrieved in the response | 

### Return type

[**[]ClientApp**](ClientApp.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostExchangeClientApp

> ClientApp PostExchangeClientApp(ctx, orgId).ApiInstanceId(apiInstanceId).PostExchangeAppsBody(postExchangeAppsBody).Execute()

Create a new Exchange Client App



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/exchange_client_apps"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	apiInstanceId := "apiInstanceId_example" // string | The API Manager Instance Id (optional)
	postExchangeAppsBody := *openapiclient.NewPostExchangeAppsBody() // PostExchangeAppsBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PostExchangeClientApp(context.Background(), orgId).ApiInstanceId(apiInstanceId).PostExchangeAppsBody(postExchangeAppsBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PostExchangeClientApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostExchangeClientApp`: ClientApp
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PostExchangeClientApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostExchangeClientAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiInstanceId** | **string** | The API Manager Instance Id | 
 **postExchangeAppsBody** | [**PostExchangeAppsBody**](PostExchangeAppsBody.md) |  | 

### Return type

[**ClientApp**](ClientApp.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

