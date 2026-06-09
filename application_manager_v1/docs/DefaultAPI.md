# \DefaultAPI

All URIs are relative to *https://anypoint.mulesoft.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDeployment**](DefaultAPI.md#DeleteDeployment) | **Delete** /cloudhub/api/v2/applications/{domain} | Delete a single application
[**GetAllDeployments**](DefaultAPI.md#GetAllDeployments) | **Get** /armui/api/v1/applications | List deployments
[**GetCH1Deployments**](DefaultAPI.md#GetCH1Deployments) | **Get** /cloudhub/api/v2/applications | List Cloudhub1 deployments
[**GetDeploymentCH1Details**](DefaultAPI.md#GetDeploymentCH1Details) | **Get** /cloudhub/api/v2/applications/{domain} | Read deployment details
[**GetMuleVersions**](DefaultAPI.md#GetMuleVersions) | **Get** /cloudhub/api/mule-versions | Read available mule versions
[**PostDeployment**](DefaultAPI.md#PostDeployment) | **Post** /cloudhub/api/v2/applications | Create a new deployment
[**PutDeployment**](DefaultAPI.md#PutDeployment) | **Put** /cloudhub/api/v2/applications/{domain} | Update a single application.
[**PutDeployments**](DefaultAPI.md#PutDeployments) | **Put** /cloudhub/api/v2/applications | Updates deployment states in bulk



## DeleteDeployment

> DeleteDeployment(ctx, domain).XAnypntEnvId(xAnypntEnvId).XAnypntOrgId(xAnypntOrgId).Execute()

Delete a single application



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/application_manager_v1"
)

func main() {
	xAnypntEnvId := "xAnypntEnvId_example" // string | The environment id where the application is deployed
	xAnypntOrgId := "xAnypntOrgId_example" // string | The org id where the application is deployed
	domain := "domain_example" // string | The application domain in the path

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteDeployment(context.Background(), domain).XAnypntEnvId(xAnypntEnvId).XAnypntOrgId(xAnypntOrgId).Execute()
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
**domain** | **string** | The application domain in the path | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAnypntEnvId** | **string** | The environment id where the application is deployed | 
 **xAnypntOrgId** | **string** | The org id where the application is deployed | 


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

> GetAllDeploymentsResponse GetAllDeployments(ctx).XAnypntEnvId(xAnypntEnvId).XAnypntOrgId(xAnypntOrgId).Execute()

List deployments



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/application_manager_v1"
)

func main() {
	xAnypntEnvId := "xAnypntEnvId_example" // string | The environment id where the applications are deployed
	xAnypntOrgId := "xAnypntOrgId_example" // string | The org id where the applications are deployed

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetAllDeployments(context.Background()).XAnypntEnvId(xAnypntEnvId).XAnypntOrgId(xAnypntOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetAllDeployments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllDeployments`: GetAllDeploymentsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetAllDeployments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllDeploymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAnypntEnvId** | **string** | The environment id where the applications are deployed | 
 **xAnypntOrgId** | **string** | The org id where the applications are deployed | 

### Return type

[**GetAllDeploymentsResponse**](GetAllDeploymentsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCH1Deployments

> []DeploymentCH1Summary GetCH1Deployments(ctx).XAnypntEnvId(xAnypntEnvId).XAnypntOrgId(xAnypntOrgId).RetrieveStatistics(retrieveStatistics).Period(period).RetrieveLogLevels(retrieveLogLevels).RetrieveTrackingSettings(retrieveTrackingSettings).RetrieveIpAddresses(retrieveIpAddresses).Execute()

List Cloudhub1 deployments



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/application_manager_v1"
)

func main() {
	xAnypntEnvId := "xAnypntEnvId_example" // string | The environment id where the application is deployed
	xAnypntOrgId := "xAnypntOrgId_example" // string | The org id where the application is deployed
	retrieveStatistics := true // bool | Include statistics for each application in the response. (optional)
	period := int64(789) // int64 | Number of milliseconds of statistics to gather. (optional)
	retrieveLogLevels := true // bool | Include log levels for each application in the response. (optional)
	retrieveTrackingSettings := true // bool | Include tracking settings for each application in the response. (optional)
	retrieveIpAddresses := true // bool | Include ip address allocations for each application in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetCH1Deployments(context.Background()).XAnypntEnvId(xAnypntEnvId).XAnypntOrgId(xAnypntOrgId).RetrieveStatistics(retrieveStatistics).Period(period).RetrieveLogLevels(retrieveLogLevels).RetrieveTrackingSettings(retrieveTrackingSettings).RetrieveIpAddresses(retrieveIpAddresses).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetCH1Deployments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCH1Deployments`: []DeploymentCH1Summary
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetCH1Deployments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCH1DeploymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAnypntEnvId** | **string** | The environment id where the application is deployed | 
 **xAnypntOrgId** | **string** | The org id where the application is deployed | 
 **retrieveStatistics** | **bool** | Include statistics for each application in the response. | 
 **period** | **int64** | Number of milliseconds of statistics to gather. | 
 **retrieveLogLevels** | **bool** | Include log levels for each application in the response. | 
 **retrieveTrackingSettings** | **bool** | Include tracking settings for each application in the response. | 
 **retrieveIpAddresses** | **bool** | Include ip address allocations for each application in the response. | 

### Return type

[**[]DeploymentCH1Summary**](DeploymentCH1Summary.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploymentCH1Details

> DeploymentCH1Details GetDeploymentCH1Details(ctx, domain).XAnypntEnvId(xAnypntEnvId).XAnypntOrgId(xAnypntOrgId).Execute()

Read deployment details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/application_manager_v1"
)

func main() {
	xAnypntEnvId := "xAnypntEnvId_example" // string | The environment id where the application is deployed
	xAnypntOrgId := "xAnypntOrgId_example" // string | The org id where the application is deployed
	domain := "domain_example" // string | The application domain in the path

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetDeploymentCH1Details(context.Background(), domain).XAnypntEnvId(xAnypntEnvId).XAnypntOrgId(xAnypntOrgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetDeploymentCH1Details``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentCH1Details`: DeploymentCH1Details
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetDeploymentCH1Details`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | The application domain in the path | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentCH1DetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAnypntEnvId** | **string** | The environment id where the application is deployed | 
 **xAnypntOrgId** | **string** | The org id where the application is deployed | 


### Return type

[**DeploymentCH1Details**](DeploymentCH1Details.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMuleVersions

> RuntimeCollection GetMuleVersions(ctx).Execute()

Read available mule versions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/application_manager_v1"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetMuleVersions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetMuleVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMuleVersions`: RuntimeCollection
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetMuleVersions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMuleVersionsRequest struct via the builder pattern


### Return type

[**RuntimeCollection**](RuntimeCollection.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDeployment

> DeploymentCH1Details PostDeployment(ctx).XAnypntEnvId(xAnypntEnvId).XAnypntOrgId(xAnypntOrgId).AppInfoJson(appInfoJson).File(file).AutoStart(autoStart).Execute()

Create a new deployment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/application_manager_v1"
)

func main() {
	xAnypntEnvId := "xAnypntEnvId_example" // string | The environment id where the application is deployed
	xAnypntOrgId := "xAnypntOrgId_example" // string | The org id where the application is deployed
	appInfoJson := "appInfoJson_example" // string | A JSON string representing the deployment configuration. Look at \\\"NewDeploymentStruct\\\" for structure.  (optional)
	file := os.NewFile(1234, "some_file") // *os.File | The Mule application artifact to be deployed. (optional)
	autoStart := true // bool | Indicates whether the application should be automatically started after deployment. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PostDeployment(context.Background()).XAnypntEnvId(xAnypntEnvId).XAnypntOrgId(xAnypntOrgId).AppInfoJson(appInfoJson).File(file).AutoStart(autoStart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PostDeployment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDeployment`: DeploymentCH1Details
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PostDeployment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAnypntEnvId** | **string** | The environment id where the application is deployed | 
 **xAnypntOrgId** | **string** | The org id where the application is deployed | 
 **appInfoJson** | **string** | A JSON string representing the deployment configuration. Look at \\\&quot;NewDeploymentStruct\\\&quot; for structure.  | 
 **file** | ***os.File** | The Mule application artifact to be deployed. | 
 **autoStart** | **bool** | Indicates whether the application should be automatically started after deployment. | 

### Return type

[**DeploymentCH1Details**](DeploymentCH1Details.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data, application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutDeployment

> DeploymentCH1Details PutDeployment(ctx, domain).XAnypntEnvId(xAnypntEnvId).XAnypntOrgId(xAnypntOrgId).AppInfoJson(appInfoJson).File(file).Execute()

Update a single application.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/application_manager_v1"
)

func main() {
	xAnypntEnvId := "xAnypntEnvId_example" // string | The environment id where the application is deployed
	xAnypntOrgId := "xAnypntOrgId_example" // string | The org id where the application is deployed
	domain := "domain_example" // string | The application domain in the path
	appInfoJson := "appInfoJson_example" // string | A JSON string representing the deployment configuration. Look at \\\"NewDeploymentStruct\\\" for structure.  (optional)
	file := os.NewFile(1234, "some_file") // *os.File | The Mule application artifact to be deployed. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PutDeployment(context.Background(), domain).XAnypntEnvId(xAnypntEnvId).XAnypntOrgId(xAnypntOrgId).AppInfoJson(appInfoJson).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PutDeployment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutDeployment`: DeploymentCH1Details
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PutDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | The application domain in the path | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAnypntEnvId** | **string** | The environment id where the application is deployed | 
 **xAnypntOrgId** | **string** | The org id where the application is deployed | 

 **appInfoJson** | **string** | A JSON string representing the deployment configuration. Look at \\\&quot;NewDeploymentStruct\\\&quot; for structure.  | 
 **file** | ***os.File** | The Mule application artifact to be deployed. | 

### Return type

[**DeploymentCH1Details**](DeploymentCH1Details.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data, application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutDeployments

> PutDeployments(ctx).XAnypntEnvId(xAnypntEnvId).XAnypntOrgId(xAnypntOrgId).UpdateDeploymentBulkBody(updateDeploymentBulkBody).Execute()

Updates deployment states in bulk



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/application_manager_v1"
)

func main() {
	xAnypntEnvId := "xAnypntEnvId_example" // string | The environment id where the application is deployed
	xAnypntOrgId := "xAnypntOrgId_example" // string | The org id where the application is deployed
	updateDeploymentBulkBody := *openapiclient.NewUpdateDeploymentBulkBody() // UpdateDeploymentBulkBody | Application deployment body containing configuration and artifact file.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.PutDeployments(context.Background()).XAnypntEnvId(xAnypntEnvId).XAnypntOrgId(xAnypntOrgId).UpdateDeploymentBulkBody(updateDeploymentBulkBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PutDeployments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutDeploymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAnypntEnvId** | **string** | The environment id where the application is deployed | 
 **xAnypntOrgId** | **string** | The org id where the application is deployed | 
 **updateDeploymentBulkBody** | [**UpdateDeploymentBulkBody**](UpdateDeploymentBulkBody.md) | Application deployment body containing configuration and artifact file. | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

