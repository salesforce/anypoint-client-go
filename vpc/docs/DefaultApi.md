# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com/cloudhub/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsOrgIdVpcsGet**](DefaultApi.md#OrganizationsOrgIdVpcsGet) | **Get** /organizations/{orgId}/vpcs | Returns a list of vpcs.



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

