# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com/accounts/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsOrgIdRolegroupsRolegroupIdRolesDelete**](DefaultApi.md#OrganizationsOrgIdRolegroupsRolegroupIdRolesDelete) | **Delete** /organizations/{orgId}/rolegroups/{rolegroupId}/roles | 
[**OrganizationsOrgIdRolegroupsRolegroupIdRolesGet**](DefaultApi.md#OrganizationsOrgIdRolegroupsRolegroupIdRolesGet) | **Get** /organizations/{orgId}/rolegroups/{rolegroupId}/roles | 
[**OrganizationsOrgIdRolegroupsRolegroupIdRolesPost**](DefaultApi.md#OrganizationsOrgIdRolegroupsRolegroupIdRolesPost) | **Post** /organizations/{orgId}/rolegroups/{rolegroupId}/roles | 
[**RolesGet**](DefaultApi.md#RolesGet) | **Get** /roles | 



## OrganizationsOrgIdRolegroupsRolegroupIdRolesDelete

> []int32 OrganizationsOrgIdRolegroupsRolegroupIdRolesDelete(ctx, orgId, rolegroupId).RoleToDelete(roleToDelete).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/role"
)

func main() {
    orgId := "orgId_example" // string | The ID of the organization in GUID format
    rolegroupId := "rolegroupId_example" // string | The id of a rolegroup
    roleToDelete := []openapiclient.RoleToDelete{*openapiclient.NewRoleToDelete()} // []RoleToDelete |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OrganizationsOrgIdRolegroupsRolegroupIdRolesDelete(context.Background(), orgId, rolegroupId).RoleToDelete(roleToDelete).Execute()
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


 **roleToDelete** | [**[]RoleToDelete**](RoleToDelete.md) |  | 

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

> OrganizationsOrgIdRolegroupsRolegroupIdRolesGet200Response OrganizationsOrgIdRolegroupsRolegroupIdRolesGet(ctx, orgId, rolegroupId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/role"
)

func main() {
    orgId := "orgId_example" // string | The ID of the organization in GUID format
    rolegroupId := "rolegroupId_example" // string | The id of a rolegroup

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OrganizationsOrgIdRolegroupsRolegroupIdRolesGet(context.Background(), orgId, rolegroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdRolegroupsRolegroupIdRolesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdRolegroupsRolegroupIdRolesGet`: OrganizationsOrgIdRolegroupsRolegroupIdRolesGet200Response
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

[**OrganizationsOrgIdRolegroupsRolegroupIdRolesGet200Response**](OrganizationsOrgIdRolegroupsRolegroupIdRolesGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrgIdRolegroupsRolegroupIdRolesPost

> []RolePostResponseItem OrganizationsOrgIdRolegroupsRolegroupIdRolesPost(ctx, orgId, rolegroupId).RoleToAssign(roleToAssign).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/role"
)

func main() {
    orgId := "orgId_example" // string | The ID of the organization in GUID format
    rolegroupId := "rolegroupId_example" // string | The id of a rolegroup
    roleToAssign := []openapiclient.RoleToAssign{*openapiclient.NewRoleToAssign()} // []RoleToAssign |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OrganizationsOrgIdRolegroupsRolegroupIdRolesPost(context.Background(), orgId, rolegroupId).RoleToAssign(roleToAssign).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdRolegroupsRolegroupIdRolesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdRolegroupsRolegroupIdRolesPost`: []RolePostResponseItem
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


 **roleToAssign** | [**[]RoleToAssign**](RoleToAssign.md) |  | 

### Return type

[**[]RolePostResponseItem**](RolePostResponseItem.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RolesGet

> RolesGet200Response RolesGet(ctx).Name(name).Description(description).IncludeInternal(includeInternal).Search(search).Offset(offset).Limit(limit).Ascending(ascending).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/role"
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RolesGet(context.Background()).Name(name).Description(description).IncludeInternal(includeInternal).Search(search).Offset(offset).Limit(limit).Ascending(ascending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RolesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RolesGet`: RolesGet200Response
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

[**RolesGet200Response**](RolesGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

