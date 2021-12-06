# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com/cloudhub/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdDelete**](DefaultApi.md#OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdDelete) | **Delete** /organizations/{orgId}/vpcs/{vpcId}/loadbalancers/{dlbId} | 
[**OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdGet**](DefaultApi.md#OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdGet) | **Get** /organizations/{orgId}/vpcs/{vpcId}/loadbalancers/{dlbId} | 
[**OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdPatch**](DefaultApi.md#OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdPatch) | **Patch** /organizations/{orgId}/vpcs/{vpcId}/loadbalancers/{dlbId} | 
[**OrganizationsOrgIdVpcsVpcIdLoadbalancersGet**](DefaultApi.md#OrganizationsOrgIdVpcsVpcIdLoadbalancersGet) | **Get** /organizations/{orgId}/vpcs/{vpcId}/loadbalancers | 
[**OrganizationsOrgIdVpcsVpcIdLoadbalancersPost**](DefaultApi.md#OrganizationsOrgIdVpcsVpcIdLoadbalancersPost) | **Post** /organizations/{orgId}/vpcs/{vpcId}/loadbalancers | 



## OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdDelete

> OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdDelete(ctx, orgId, vpcId, dlbId).Execute()





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
    orgId := "orgId_example" // string | The ID of the organization in GUID format
    vpcId := "vpcId_example" // string | The ID of the VPC in GUID format
    dlbId := "dlbId_example" // string | The ID of the DLB in GUID format

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdDelete(context.Background(), orgId, vpcId, dlbId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdDelete``: %v\n", err)
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
    openapiclient "./openapi"
)

func main() {
    orgId := "orgId_example" // string | The ID of the organization in GUID format
    vpcId := "vpcId_example" // string | The ID of the VPC in GUID format
    dlbId := "dlbId_example" // string | The ID of the DLB in GUID format

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdGet(context.Background(), orgId, vpcId, dlbId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdGet`: Dlb
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdGet`: %v\n", resp)
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

> Dlb OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdPatch(ctx, orgId, vpcId, dlbId).RequestBody(requestBody).Execute()





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
    orgId := "orgId_example" // string | The ID of the organization in GUID format
    vpcId := "vpcId_example" // string | The ID of the VPC in GUID format
    dlbId := "dlbId_example" // string | The ID of the DLB in GUID format
    requestBody := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdPatch(context.Background(), orgId, vpcId, dlbId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdPatch`: Dlb
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdPatch`: %v\n", resp)
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



 **requestBody** | **[]map[string]interface{}** |  | 

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

> InlineResponse200 OrganizationsOrgIdVpcsVpcIdLoadbalancersGet(ctx, orgId, vpcId).Execute()





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
    orgId := "orgId_example" // string | The ID of the organization in GUID format
    vpcId := "vpcId_example" // string | The ID of the VPC in GUID format

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdVpcsVpcIdLoadbalancersGet(context.Background(), orgId, vpcId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdVpcsVpcIdLoadbalancersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdVpcsVpcIdLoadbalancersGet`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdVpcsVpcIdLoadbalancersGet`: %v\n", resp)
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

[**InlineResponse200**](InlineResponse200.md)

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
    openapiclient "./openapi"
)

func main() {
    orgId := "orgId_example" // string | The ID of the organization in GUID format
    vpcId := "vpcId_example" // string | The ID of the VPC in GUID format
    dlbPostBody := *openapiclient.NewDlbPostBody() // DlbPostBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdVpcsVpcIdLoadbalancersPost(context.Background(), orgId, vpcId).DlbPostBody(dlbPostBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdVpcsVpcIdLoadbalancersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdVpcsVpcIdLoadbalancersPost`: Dlb
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdVpcsVpcIdLoadbalancersPost`: %v\n", resp)
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

