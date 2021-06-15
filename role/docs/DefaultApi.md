# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com/accounts/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsOrgIdRolegroupsRolegroupIdRolesDelete**](DefaultApi.md#OrganizationsOrgIdRolegroupsRolegroupIdRolesDelete) | **Delete** /organizations/{orgId}/rolegroups/{rolegroupId}/roles | 
[**OrganizationsOrgIdRolegroupsRolegroupIdRolesGet**](DefaultApi.md#OrganizationsOrgIdRolegroupsRolegroupIdRolesGet) | **Get** /organizations/{orgId}/rolegroups/{rolegroupId}/roles | 
[**OrganizationsOrgIdRolegroupsRolegroupIdRolesPost**](DefaultApi.md#OrganizationsOrgIdRolegroupsRolegroupIdRolesPost) | **Post** /organizations/{orgId}/rolegroups/{rolegroupId}/roles | 
[**RolesGet**](DefaultApi.md#RolesGet) | **Get** /roles | 



## OrganizationsOrgIdRolegroupsRolegroupIdRolesDelete

> []int32 OrganizationsOrgIdRolegroupsRolegroupIdRolesDelete(ctx, orgId, rolegroupId).RequestBody(requestBody).Execute()





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
    rolegroupId := "rolegroupId_example" // string | The id of a rolegroup
    requestBody := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdRolegroupsRolegroupIdRolesDelete(context.Background(), orgId, rolegroupId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdRolegroupsRolegroupIdRolesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdRolegroupsRolegroupIdRolesDelete`: []int32
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdRolegroupsRolegroupIdRolesDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**rolegroupId** | **string** | The id of a rolegroup | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdRolegroupsRolegroupIdRolesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **[]map[string]interface{}** |  | 

### Return type

**[]int32**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrgIdRolegroupsRolegroupIdRolesGet

> InlineResponse2001 OrganizationsOrgIdRolegroupsRolegroupIdRolesGet(ctx, orgId, rolegroupId).Execute()





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
    rolegroupId := "rolegroupId_example" // string | The id of a rolegroup

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdRolegroupsRolegroupIdRolesGet(context.Background(), orgId, rolegroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdRolegroupsRolegroupIdRolesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdRolegroupsRolegroupIdRolesGet`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdRolegroupsRolegroupIdRolesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**rolegroupId** | **string** | The id of a rolegroup | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdRolegroupsRolegroupIdRolesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrgIdRolegroupsRolegroupIdRolesPost

> []map[string]interface{} OrganizationsOrgIdRolegroupsRolegroupIdRolesPost(ctx, orgId, rolegroupId).RequestBody(requestBody).Execute()





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
    rolegroupId := "rolegroupId_example" // string | The id of a rolegroup
    requestBody := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdRolegroupsRolegroupIdRolesPost(context.Background(), orgId, rolegroupId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdRolegroupsRolegroupIdRolesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdRolegroupsRolegroupIdRolesPost`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdRolegroupsRolegroupIdRolesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**rolegroupId** | **string** | The id of a rolegroup | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdRolegroupsRolegroupIdRolesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **[]map[string]interface{}** |  | 

### Return type

**[]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RolesGet

> InlineResponse200 RolesGet(ctx).Name(name).Description(description).IncludeInternal(includeInternal).Search(search).Offset(offset).Limit(limit).Ascending(ascending).Execute()





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
    name := "name_example" // string | search by role name (optional)
    description := "description_example" // string | search by role description (optional)
    includeInternal := true // bool | include internal roles in search (optional)
    search := "search_example" // string | A search string to use for partial matches of role names (optional)
    offset := int32(56) // int32 | Pagination parameter to start returning results from this position of matches (optional)
    limit := int32(56) // int32 | Pagination parameter for how many results to return (optional)
    ascending := true // bool | Sort order for filtering (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.RolesGet(context.Background()).Name(name).Description(description).IncludeInternal(includeInternal).Search(search).Offset(offset).Limit(limit).Ascending(ascending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RolesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RolesGet`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RolesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRolesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | search by role name | 
 **description** | **string** | search by role description | 
 **includeInternal** | **bool** | include internal roles in search | 
 **search** | **string** | A search string to use for partial matches of role names | 
 **offset** | **int32** | Pagination parameter to start returning results from this position of matches | 
 **limit** | **int32** | Pagination parameter for how many results to return | 
 **ascending** | **bool** | Sort order for filtering | 

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

