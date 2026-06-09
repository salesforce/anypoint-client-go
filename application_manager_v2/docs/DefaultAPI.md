# \DefaultAPI

All URIs are relative to *https://anypoint.mulesoft.com/amc/application-manager/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDeployment**](DefaultAPI.md#DeleteDeployment) | **Delete** /organizations/{orgId}/environments/{envId}/deployments/{deploymentId} | Delete deployment
[**GetAllDeployments**](DefaultAPI.md#GetAllDeployments) | **Get** /organizations/{orgId}/environments/{envId}/deployments | List deployments
[**GetDeploymentById**](DefaultAPI.md#GetDeploymentById) | **Get** /organizations/{orgId}/environments/{envId}/deployments/{deploymentId} | Get a deploymnt
[**PatchDeployment**](DefaultAPI.md#PatchDeployment) | **Patch** /organizations/{orgId}/environments/{envId}/deployments/{deploymentId} | Patch a deployment
[**PostDeployment**](DefaultAPI.md#PostDeployment) | **Post** /organizations/{orgId}/environments/{envId}/deployments | Create deployment



## DeleteDeployment

> DeleteDeployment(ctx, orgId, envId, deploymentId).Execute()

Delete deployment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/application_manager_v2"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	envId := "envId_example" // string | The ID of the environment in GUID format
	deploymentId := "deploymentId_example" // string | The ID of the deployment in GUID format.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteDeployment(context.Background(), orgId, envId, deploymentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteDeployment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**envId** | **string** | The ID of the environment in GUID format | 
**deploymentId** | **string** | The ID of the deployment in GUID format. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeploymentRequest struct via the builder pattern


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


## GetAllDeployments

> DeploymentResponsePaged GetAllDeployments(ctx, orgId, envId).Provider(provider).TargetId(targetId).Offset(offset).Limit(limit).Execute()

List deployments



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/application_manager_v2"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	envId := "envId_example" // string | The ID of the environment in GUID format
	provider := "provider_example" // string | The provider of the target the deployments are deployed to. (optional)
	targetId := "targetId_example" // string | The id of the target the deployments are deployed to. (optional)
	offset := int32(56) // int32 | Offset used to retrieve the results. (optional)
	limit := int32(56) // int32 | Maximum records to retrieve. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetAllDeployments(context.Background(), orgId, envId).Provider(provider).TargetId(targetId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetAllDeployments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllDeployments`: DeploymentResponsePaged
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetAllDeployments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**envId** | **string** | The ID of the environment in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllDeploymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **provider** | **string** | The provider of the target the deployments are deployed to. | 
 **targetId** | **string** | The id of the target the deployments are deployed to. | 
 **offset** | **int32** | Offset used to retrieve the results. | 
 **limit** | **int32** | Maximum records to retrieve. | 

### Return type

[**DeploymentResponsePaged**](DeploymentResponsePaged.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploymentById

> Deployment GetDeploymentById(ctx, orgId, envId, deploymentId).Execute()

Get a deploymnt



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/application_manager_v2"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	envId := "envId_example" // string | The ID of the environment in GUID format
	deploymentId := "deploymentId_example" // string | The ID of the deployment in GUID format.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetDeploymentById(context.Background(), orgId, envId, deploymentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetDeploymentById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentById`: Deployment
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetDeploymentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**envId** | **string** | The ID of the environment in GUID format | 
**deploymentId** | **string** | The ID of the deployment in GUID format. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Deployment**](Deployment.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchDeployment

> Deployment PatchDeployment(ctx, orgId, envId, deploymentId).DeploymentRequestBody(deploymentRequestBody).Execute()

Patch a deployment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/application_manager_v2"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	envId := "envId_example" // string | The ID of the environment in GUID format
	deploymentId := "deploymentId_example" // string | The ID of the deployment in GUID format.
	deploymentRequestBody := *openapiclient.NewDeploymentRequestBody() // DeploymentRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PatchDeployment(context.Background(), orgId, envId, deploymentId).DeploymentRequestBody(deploymentRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PatchDeployment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchDeployment`: Deployment
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PatchDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**envId** | **string** | The ID of the environment in GUID format | 
**deploymentId** | **string** | The ID of the deployment in GUID format. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **deploymentRequestBody** | [**DeploymentRequestBody**](DeploymentRequestBody.md) |  | 

### Return type

[**Deployment**](Deployment.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDeployment

> Deployment PostDeployment(ctx, orgId, envId).DeploymentRequestBody(deploymentRequestBody).Execute()

Create deployment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/application_manager_v2"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	envId := "envId_example" // string | The ID of the environment in GUID format
	deploymentRequestBody := *openapiclient.NewDeploymentRequestBody() // DeploymentRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PostDeployment(context.Background(), orgId, envId).DeploymentRequestBody(deploymentRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PostDeployment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDeployment`: Deployment
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PostDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**envId** | **string** | The ID of the environment in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deploymentRequestBody** | [**DeploymentRequestBody**](DeploymentRequestBody.md) |  | 

### Return type

[**Deployment**](Deployment.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

