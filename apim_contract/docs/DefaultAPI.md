# \DefaultAPI

All URIs are relative to *https://anypoint.mulesoft.com/apimanager*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiContract**](DefaultAPI.md#CreateApiContract) | **Post** /api/v1/organizations/{orgId}/environments/{envId}/apis/{apiId}/contracts | Creates a new contract for a given API in a given organization and environment
[**DeleteApiContract**](DefaultAPI.md#DeleteApiContract) | **Delete** /api/v1/organizations/{orgId}/environments/{envId}/apis/{apiId}/contracts/{contractId} | Deletes a contract for a given API in a given organization and environment
[**GetApiContract**](DefaultAPI.md#GetApiContract) | **Get** /api/v1/organizations/{orgId}/environments/{envId}/apis/{apiId}/contracts/{contractId} | Retrieves a contract for a given API in a given organization and environment
[**GetApiContracts**](DefaultAPI.md#GetApiContracts) | **Get** /api/v1/organizations/{orgId}/environments/{envId}/apis/{apiId}/contracts | Retrieves all contracts for a given API in a given organization and environment
[**UpdateApiContract**](DefaultAPI.md#UpdateApiContract) | **Patch** /api/v1/organizations/{orgId}/environments/{envId}/apis/{apiId}/contracts/{contractId} | Updates a contract for a given API in a given organization and environment
[**UpdateApiContractStatus**](DefaultAPI.md#UpdateApiContractStatus) | **Post** /xapi/v1/organizations/{orgId}/environments/{envId}/apis/{apiId}/contracts/{contractId}/{action} | Performs an action on a contract for a given API in a given organization and environment



## CreateApiContract

> ContractDetails CreateApiContract(ctx, orgId, envId, apiId).PostApiContract(postApiContract).Execute()

Creates a new contract for a given API in a given organization and environment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/apim_contract"
)

func main() {
	orgId := "orgId_example" // string | The organization Id
	envId := "envId_example" // string | The environment id
	apiId := "apiId_example" // string | The api manager instance id for a given environment
	postApiContract := *openapiclient.NewPostApiContract() // PostApiContract | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateApiContract(context.Background(), orgId, envId, apiId).PostApiContract(postApiContract).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateApiContract``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApiContract`: ContractDetails
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateApiContract`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment id | 
**apiId** | **string** | The api manager instance id for a given environment | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **postApiContract** | [**PostApiContract**](PostApiContract.md) |  | 

### Return type

[**ContractDetails**](ContractDetails.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApiContract

> DeleteApiContract(ctx, orgId, envId, apiId, contractId).Execute()

Deletes a contract for a given API in a given organization and environment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/apim_contract"
)

func main() {
	orgId := "orgId_example" // string | The organization Id
	envId := "envId_example" // string | The environment id
	apiId := "apiId_example" // string | The api manager instance id for a given environment
	contractId := "contractId_example" // string | The contract id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteApiContract(context.Background(), orgId, envId, apiId, contractId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteApiContract``: %v\n", err)
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
**apiId** | **string** | The api manager instance id for a given environment | 
**contractId** | **string** | The contract id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiContractRequest struct via the builder pattern


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


## GetApiContract

> ContractDetails GetApiContract(ctx, orgId, envId, apiId, contractId).Execute()

Retrieves a contract for a given API in a given organization and environment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/apim_contract"
)

func main() {
	orgId := "orgId_example" // string | The organization Id
	envId := "envId_example" // string | The environment id
	apiId := "apiId_example" // string | The api manager instance id for a given environment
	contractId := "contractId_example" // string | The contract id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetApiContract(context.Background(), orgId, envId, apiId, contractId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetApiContract``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiContract`: ContractDetails
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetApiContract`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment id | 
**apiId** | **string** | The api manager instance id for a given environment | 
**contractId** | **string** | The contract id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**ContractDetails**](ContractDetails.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiContracts

> GetContractsResponse GetApiContracts(ctx, orgId, envId, apiId).Limit(limit).Offset(offset).Sort(sort).Order(order).Ascending(ascending).Status(status).Query(query).IncludeExtraApplicationData(includeExtraApplicationData).Execute()

Retrieves all contracts for a given API in a given organization and environment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/apim_contract"
)

func main() {
	orgId := "orgId_example" // string | The organization Id
	envId := "envId_example" // string | The environment id
	apiId := "apiId_example" // string | The api manager instance id for a given environment
	limit := int32(56) // int32 | The maximum number of contracts to return (optional)
	offset := int32(56) // int32 | The offset of the first contract to return (optional)
	sort := "sort_example" // string | The field to sort by (optional)
	order := "order_example" // string | The order to sort by (optional)
	ascending := true // bool | The order to sort by (optional)
	status := "status_example" // string | The status of the contract (optional)
	query := "query_example" // string | A string that will be checked for a partial or similar matches of the name, description, label and tags (optional)
	includeExtraApplicationData := true // bool | Whether to include extra application data (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetApiContracts(context.Background(), orgId, envId, apiId).Limit(limit).Offset(offset).Sort(sort).Order(order).Ascending(ascending).Status(status).Query(query).IncludeExtraApplicationData(includeExtraApplicationData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetApiContracts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiContracts`: GetContractsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetApiContracts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment id | 
**apiId** | **string** | The api manager instance id for a given environment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiContractsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **int32** | The maximum number of contracts to return | 
 **offset** | **int32** | The offset of the first contract to return | 
 **sort** | **string** | The field to sort by | 
 **order** | **string** | The order to sort by | 
 **ascending** | **bool** | The order to sort by | 
 **status** | **string** | The status of the contract | 
 **query** | **string** | A string that will be checked for a partial or similar matches of the name, description, label and tags | 
 **includeExtraApplicationData** | **bool** | Whether to include extra application data | 

### Return type

[**GetContractsResponse**](GetContractsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApiContract

> ContractDetails UpdateApiContract(ctx, orgId, envId, apiId, contractId).PatchApiContract(patchApiContract).Execute()

Updates a contract for a given API in a given organization and environment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/apim_contract"
)

func main() {
	orgId := "orgId_example" // string | The organization Id
	envId := "envId_example" // string | The environment id
	apiId := "apiId_example" // string | The api manager instance id for a given environment
	contractId := "contractId_example" // string | The contract id
	patchApiContract := *openapiclient.NewPatchApiContract() // PatchApiContract | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateApiContract(context.Background(), orgId, envId, apiId, contractId).PatchApiContract(patchApiContract).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateApiContract``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApiContract`: ContractDetails
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateApiContract`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment id | 
**apiId** | **string** | The api manager instance id for a given environment | 
**contractId** | **string** | The contract id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **patchApiContract** | [**PatchApiContract**](PatchApiContract.md) |  | 

### Return type

[**ContractDetails**](ContractDetails.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApiContractStatus

> ContractDetails UpdateApiContractStatus(ctx, orgId, envId, apiId, contractId, action).Execute()

Performs an action on a contract for a given API in a given organization and environment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/apim_contract"
)

func main() {
	orgId := "orgId_example" // string | The organization Id
	envId := "envId_example" // string | The environment id
	apiId := "apiId_example" // string | The api manager instance id for a given environment
	contractId := "contractId_example" // string | The contract id
	action := "action_example" // string | The action to be performed on the contract

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateApiContractStatus(context.Background(), orgId, envId, apiId, contractId, action).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateApiContractStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApiContractStatus`: ContractDetails
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateApiContractStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment id | 
**apiId** | **string** | The api manager instance id for a given environment | 
**contractId** | **string** | The contract id | 
**action** | **string** | The action to be performed on the contract | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiContractStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

[**ContractDetails**](ContractDetails.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

