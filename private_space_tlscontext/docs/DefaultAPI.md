# \DefaultAPI

All URIs are relative to *https://anypoint.mulesoft.com/runtimefabric/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTlsContext**](DefaultAPI.md#CreateTlsContext) | **Post** /organizations/{orgId}/privatespaces/{privateSpaceId}/tlsContexts | Create a TLS Context
[**DeleteTlsContext**](DefaultAPI.md#DeleteTlsContext) | **Delete** /organizations/{orgId}/privatespaces/{privateSpaceId}/tlsContexts/{tlsContextId} | Delete TLS Context
[**GetTlsContext**](DefaultAPI.md#GetTlsContext) | **Get** /organizations/{orgId}/privatespaces/{privateSpaceId}/tlsContexts/{tlsContextId} | Get TLS Context
[**GetTlsContexts**](DefaultAPI.md#GetTlsContexts) | **Get** /organizations/{orgId}/privatespaces/{privateSpaceId}/tlsContexts | Get TLS Contexts
[**UpdateTlsContext**](DefaultAPI.md#UpdateTlsContext) | **Patch** /organizations/{orgId}/privatespaces/{privateSpaceId}/tlsContexts/{tlsContextId} | Update TLS Context



## CreateTlsContext

> TlsContext CreateTlsContext(ctx, orgId, privateSpaceId).TlsContextPostBody(tlsContextPostBody).Execute()

Create a TLS Context



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/private_space_tlscontext"
)

func main() {
	orgId := "orgId_example" // string | Organization ID
	privateSpaceId := "privateSpaceId_example" // string | Private Space ID
	tlsContextPostBody := *openapiclient.NewTlsContextPostBody() // TlsContextPostBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateTlsContext(context.Background(), orgId, privateSpaceId).TlsContextPostBody(tlsContextPostBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateTlsContext``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTlsContext`: TlsContext
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateTlsContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**privateSpaceId** | **string** | Private Space ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTlsContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tlsContextPostBody** | [**TlsContextPostBody**](TlsContextPostBody.md) |  | 

### Return type

[**TlsContext**](TlsContext.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTlsContext

> DeleteTlsContext(ctx, orgId, privateSpaceId, tlsContextId).Execute()

Delete TLS Context



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/private_space_tlscontext"
)

func main() {
	orgId := "orgId_example" // string | Organization ID
	privateSpaceId := "privateSpaceId_example" // string | Private Space ID
	tlsContextId := "tlsContextId_example" // string | TLS Context ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteTlsContext(context.Background(), orgId, privateSpaceId, tlsContextId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteTlsContext``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**privateSpaceId** | **string** | Private Space ID | 
**tlsContextId** | **string** | TLS Context ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTlsContextRequest struct via the builder pattern


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


## GetTlsContext

> TlsContext GetTlsContext(ctx, orgId, privateSpaceId, tlsContextId).Execute()

Get TLS Context



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/private_space_tlscontext"
)

func main() {
	orgId := "orgId_example" // string | Organization ID
	privateSpaceId := "privateSpaceId_example" // string | Private Space ID
	tlsContextId := "tlsContextId_example" // string | TLS Context ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetTlsContext(context.Background(), orgId, privateSpaceId, tlsContextId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetTlsContext``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTlsContext`: TlsContext
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetTlsContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**privateSpaceId** | **string** | Private Space ID | 
**tlsContextId** | **string** | TLS Context ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTlsContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**TlsContext**](TlsContext.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTlsContexts

> []TlsContext GetTlsContexts(ctx, orgId, privateSpaceId).Execute()

Get TLS Contexts



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/private_space_tlscontext"
)

func main() {
	orgId := "orgId_example" // string | Organization ID
	privateSpaceId := "privateSpaceId_example" // string | Private Space ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetTlsContexts(context.Background(), orgId, privateSpaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetTlsContexts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTlsContexts`: []TlsContext
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetTlsContexts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**privateSpaceId** | **string** | Private Space ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTlsContextsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]TlsContext**](TlsContext.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTlsContext

> TlsContext UpdateTlsContext(ctx, orgId, privateSpaceId, tlsContextId).Body(body).Execute()

Update TLS Context



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/private_space_tlscontext"
)

func main() {
	orgId := "orgId_example" // string | Organization ID
	privateSpaceId := "privateSpaceId_example" // string | Private Space ID
	tlsContextId := "tlsContextId_example" // string | TLS Context ID
	body := TlsContextPostBody(987) // TlsContextPostBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateTlsContext(context.Background(), orgId, privateSpaceId, tlsContextId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateTlsContext``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTlsContext`: TlsContext
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateTlsContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**privateSpaceId** | **string** | Private Space ID | 
**tlsContextId** | **string** | TLS Context ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTlsContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | **TlsContextPostBody** |  | 

### Return type

[**TlsContext**](TlsContext.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

