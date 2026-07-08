# \DefaultAPI

All URIs are relative to *https://anypoint.mulesoft.com/accounts/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBG**](DefaultAPI.md#CreateBG) | **Post** /organizations | Creates a new Business Group.
[**DeleteBG**](DefaultAPI.md#DeleteBG) | **Delete** /organizations/{orgId} | Delete a Business Group by its id.
[**GetBG**](DefaultAPI.md#GetBG) | **Get** /organizations/{orgId} | Returns the Business Group instance with the given id.
[**GetBGUsage**](DefaultAPI.md#GetBGUsage) | **Get** /cs/organizations/{orgId}/usage | Returns the usage for a given Business Group.
[**GetClientCredentials**](DefaultAPI.md#GetClientCredentials) | **Get** /organizations/{orgId}/clients/{clientId} | 
[**UpdateBG**](DefaultAPI.md#UpdateBG) | **Put** /organizations/{orgId} | Update a Business Group by its id.



## CreateBG

> BGCore CreateBG(ctx).BGPostReqBody(bGPostReqBody).Execute()

Creates a new Business Group.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/org"
)

func main() {
	bGPostReqBody := *openapiclient.NewBGPostReqBody(*openapiclient.NewEntitlementsCore(), "Name_example", "OwnerId_example", "ParentOrganizationId_example") // BGPostReqBody | Business Group Request Body Object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateBG(context.Background()).BGPostReqBody(bGPostReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateBG``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBG`: BGCore
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateBG`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBGRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bGPostReqBody** | [**BGPostReqBody**](BGPostReqBody.md) | Business Group Request Body Object | 

### Return type

[**BGCore**](BGCore.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/html, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBG

> map[string]interface{} DeleteBG(ctx, orgId).Execute()

Delete a Business Group by its id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/org"
)

func main() {
	orgId := "orgId_example" // string | The organization Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DeleteBG(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteBG``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBG`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteBG`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBGRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBG

> MasterBGDetail GetBG(ctx, orgId).Execute()

Returns the Business Group instance with the given id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/org"
)

func main() {
	orgId := "orgId_example" // string | The organization Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetBG(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetBG``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBG`: MasterBGDetail
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetBG`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBGRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MasterBGDetail**](MasterBGDetail.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBGUsage

> Usage GetBGUsage(ctx, orgId).Execute()

Returns the usage for a given Business Group.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/org"
)

func main() {
	orgId := "orgId_example" // string | The organization Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetBGUsage(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetBGUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBGUsage`: Usage
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetBGUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBGUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Usage**](Usage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientCredentials

> Client GetClientCredentials(ctx, orgId, clientId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/org"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	clientId := "clientId_example" // string | The client ID of the client

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetClientCredentials(context.Background(), orgId, clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetClientCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientCredentials`: Client
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetClientCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**clientId** | **string** | The client ID of the client | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Client**](Client.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBG

> BGCore UpdateBG(ctx, orgId).BGUpdateReqBody(bGUpdateReqBody).Execute()

Update a Business Group by its id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/org"
)

func main() {
	orgId := "orgId_example" // string | The organization Id
	bGUpdateReqBody := *openapiclient.NewBGUpdateReqBody() // BGUpdateReqBody | Business Group Request Body Object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateBG(context.Background(), orgId).BGUpdateReqBody(bGUpdateReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateBG``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBG`: BGCore
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateBG`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBGRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bGUpdateReqBody** | [**BGUpdateReqBody**](BGUpdateReqBody.md) | Business Group Request Body Object | 

### Return type

[**BGCore**](BGCore.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

