# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com/cloudhub/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsOrgIdVpcsGet**](DefaultApi.md#OrganizationsOrgIdVpcsGet) | **Get** /organizations/{orgId}/vpcs | Returns a list of vpcs.
[**OrganizationsOrgIdVpcsPost**](DefaultApi.md#OrganizationsOrgIdVpcsPost) | **Post** /organizations/{orgId}/vpcs | Creates a new VPC in the provided organization.
[**OrganizationsOrgIdVpcsVpcIdDelete**](DefaultApi.md#OrganizationsOrgIdVpcsVpcIdDelete) | **Delete** /organizations/{orgId}/vpcs/{vpcId} | Delete a VPC by its id.
[**OrganizationsOrgIdVpcsVpcIdGet**](DefaultApi.md#OrganizationsOrgIdVpcsVpcIdGet) | **Get** /organizations/{orgId}/vpcs/{vpcId} | Returns the vpc instance with the given id.
[**OrganizationsOrgIdVpcsVpcIdPut**](DefaultApi.md#OrganizationsOrgIdVpcsVpcIdPut) | **Put** /organizations/{orgId}/vpcs/{vpcId} | Update the VPC configuration



## OrganizationsOrgIdVpcsGet

> InlineResponse200 OrganizationsOrgIdVpcsGet(ctx, orgId).Execute()

Returns a list of vpcs.



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdVpcsGet(context.Background(), orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdVpcsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdVpcsGet`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdVpcsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdVpcsGetRequest struct via the builder pattern


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


## OrganizationsOrgIdVpcsPost

> Vpc OrganizationsOrgIdVpcsPost(ctx, orgId).VpcCore(vpcCore).Execute()

Creates a new VPC in the provided organization.



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
    vpcCore := *openapiclient.NewVpcCore() // VpcCore | VPC Object don't

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdVpcsPost(context.Background(), orgId).VpcCore(vpcCore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdVpcsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdVpcsPost`: Vpc
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdVpcsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdVpcsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vpcCore** | [**VpcCore**](VpcCore.md) | VPC Object don&#39;t | 

### Return type

[**Vpc**](Vpc.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrgIdVpcsVpcIdDelete

> OrganizationsOrgIdVpcsVpcIdDelete(ctx, orgId, vpcId).Execute()

Delete a VPC by its id.



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
    vpcId := "vpcId_example" // string | The VPC Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdVpcsVpcIdDelete(context.Background(), orgId, vpcId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdVpcsVpcIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**vpcId** | **string** | The VPC Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdVpcsVpcIdDeleteRequest struct via the builder pattern


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


## OrganizationsOrgIdVpcsVpcIdGet

> Vpc OrganizationsOrgIdVpcsVpcIdGet(ctx, orgId, vpcId).Execute()

Returns the vpc instance with the given id.



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
    vpcId := "vpcId_example" // string | The VPC Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdVpcsVpcIdGet(context.Background(), orgId, vpcId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdVpcsVpcIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdVpcsVpcIdGet`: Vpc
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdVpcsVpcIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**vpcId** | **string** | The VPC Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdVpcsVpcIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Vpc**](Vpc.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrgIdVpcsVpcIdPut

> OrganizationsOrgIdVpcsVpcIdPut(ctx, orgId, vpcId).VpcCore(vpcCore).Execute()

Update the VPC configuration



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
    vpcId := "vpcId_example" // string | The VPC Id
    vpcCore := *openapiclient.NewVpcCore() // VpcCore | VPC Object don't

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdVpcsVpcIdPut(context.Background(), orgId, vpcId).VpcCore(vpcCore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdVpcsVpcIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**vpcId** | **string** | The VPC Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdVpcsVpcIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **vpcCore** | [**VpcCore**](VpcCore.md) | VPC Object don&#39;t | 

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

