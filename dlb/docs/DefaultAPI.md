# \DefaultAPI

All URIs are relative to *https://anypoint.mulesoft.com/cloudhub/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdDelete**](DefaultAPI.md#OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdDelete) | **Delete** /organizations/{orgId}/vpcs/{vpcId}/loadbalancers/{dlbId} | 
[**OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdGet**](DefaultAPI.md#OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdGet) | **Get** /organizations/{orgId}/vpcs/{vpcId}/loadbalancers/{dlbId} | 
[**OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdPatch**](DefaultAPI.md#OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdPatch) | **Patch** /organizations/{orgId}/vpcs/{vpcId}/loadbalancers/{dlbId} | 
[**OrganizationsOrgIdVpcsVpcIdLoadbalancersGet**](DefaultAPI.md#OrganizationsOrgIdVpcsVpcIdLoadbalancersGet) | **Get** /organizations/{orgId}/vpcs/{vpcId}/loadbalancers | 
[**OrganizationsOrgIdVpcsVpcIdLoadbalancersPost**](DefaultAPI.md#OrganizationsOrgIdVpcsVpcIdLoadbalancersPost) | **Post** /organizations/{orgId}/vpcs/{vpcId}/loadbalancers | 



## OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdDelete

> OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdDelete(ctx, orgId, vpcId, dlbId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/dlb"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	vpcId := "vpcId_example" // string | The ID of the VPC in GUID format
	dlbId := "dlbId_example" // string | The ID of the DLB in GUID format

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdDelete(context.Background(), orgId, vpcId, dlbId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**vpcId** | **string** | The ID of the VPC in GUID format | 
**dlbId** | **string** | The ID of the DLB in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdDeleteRequest struct via the builder pattern


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


## OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdGet

> Dlb OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdGet(ctx, orgId, vpcId, dlbId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/dlb"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	vpcId := "vpcId_example" // string | The ID of the VPC in GUID format
	dlbId := "dlbId_example" // string | The ID of the DLB in GUID format

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdGet(context.Background(), orgId, vpcId, dlbId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdGet`: Dlb
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**vpcId** | **string** | The ID of the VPC in GUID format | 
**dlbId** | **string** | The ID of the DLB in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Dlb**](Dlb.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdPatch

> Dlb OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdPatch(ctx, orgId, vpcId, dlbId).UpdateObject(updateObject).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/dlb"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	vpcId := "vpcId_example" // string | The ID of the VPC in GUID format
	dlbId := "dlbId_example" // string | The ID of the DLB in GUID format
	updateObject := []openapiclient.UpdateObject{*openapiclient.NewUpdateObject()} // []UpdateObject |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdPatch(context.Background(), orgId, vpcId, dlbId).UpdateObject(updateObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdPatch`: Dlb
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**vpcId** | **string** | The ID of the VPC in GUID format | 
**dlbId** | **string** | The ID of the DLB in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateObject** | [**[]UpdateObject**](UpdateObject.md) |  | 

### Return type

[**Dlb**](Dlb.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrgIdVpcsVpcIdLoadbalancersGet

> OrganizationsOrgIdVpcsVpcIdLoadbalancersGet200Response OrganizationsOrgIdVpcsVpcIdLoadbalancersGet(ctx, orgId, vpcId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/dlb"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	vpcId := "vpcId_example" // string | The ID of the VPC in GUID format

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.OrganizationsOrgIdVpcsVpcIdLoadbalancersGet(context.Background(), orgId, vpcId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.OrganizationsOrgIdVpcsVpcIdLoadbalancersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsOrgIdVpcsVpcIdLoadbalancersGet`: OrganizationsOrgIdVpcsVpcIdLoadbalancersGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.OrganizationsOrgIdVpcsVpcIdLoadbalancersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**vpcId** | **string** | The ID of the VPC in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdVpcsVpcIdLoadbalancersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationsOrgIdVpcsVpcIdLoadbalancersGet200Response**](OrganizationsOrgIdVpcsVpcIdLoadbalancersGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrgIdVpcsVpcIdLoadbalancersPost

> Dlb OrganizationsOrgIdVpcsVpcIdLoadbalancersPost(ctx, orgId, vpcId).DlbPostBody(dlbPostBody).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/dlb"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	vpcId := "vpcId_example" // string | The ID of the VPC in GUID format
	dlbPostBody := *openapiclient.NewDlbPostBody() // DlbPostBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.OrganizationsOrgIdVpcsVpcIdLoadbalancersPost(context.Background(), orgId, vpcId).DlbPostBody(dlbPostBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.OrganizationsOrgIdVpcsVpcIdLoadbalancersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsOrgIdVpcsVpcIdLoadbalancersPost`: Dlb
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.OrganizationsOrgIdVpcsVpcIdLoadbalancersPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**vpcId** | **string** | The ID of the VPC in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdVpcsVpcIdLoadbalancersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dlbPostBody** | [**DlbPostBody**](DlbPostBody.md) |  | 

### Return type

[**Dlb**](Dlb.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

