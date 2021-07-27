# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com/accounts/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsOrgIdUsersUserIdRolegroupsGet**](DefaultApi.md#OrganizationsOrgIdUsersUserIdRolegroupsGet) | **Get** /organizations/{orgId}/users/{userId}/rolegroups | 
[**OrganizationsOrgIdUsersUserIdRolegroupsRolegroupIdDelete**](DefaultApi.md#OrganizationsOrgIdUsersUserIdRolegroupsRolegroupIdDelete) | **Delete** /organizations/{orgId}/users/{userId}/rolegroups/{rolegroupId} | 
[**OrganizationsOrgIdUsersUserIdRolegroupsRolegroupIdPost**](DefaultApi.md#OrganizationsOrgIdUsersUserIdRolegroupsRolegroupIdPost) | **Post** /organizations/{orgId}/users/{userId}/rolegroups/{rolegroupId} | 



## OrganizationsOrgIdUsersUserIdRolegroupsGet

> InlineResponse200 OrganizationsOrgIdUsersUserIdRolegroupsGet(ctx, orgId, userId).Limit(limit).Offset(offset).Execute()





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
    userId := "userId_example" // string | The ID of the user in GUID format
    limit := int32(56) // int32 | Maximum number of rolegroups to retrieve per request. (optional)
    offset := int32(56) // int32 | The number of records to omit from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdUsersUserIdRolegroupsGet(context.Background(), orgId, userId).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdUsersUserIdRolegroupsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdUsersUserIdRolegroupsGet`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdUsersUserIdRolegroupsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**userId** | **string** | The ID of the user in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdUsersUserIdRolegroupsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Maximum number of rolegroups to retrieve per request. | 
 **offset** | **int32** | The number of records to omit from the response. | 

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


## OrganizationsOrgIdUsersUserIdRolegroupsRolegroupIdDelete

> OrganizationsOrgIdUsersUserIdRolegroupsRolegroupIdDelete(ctx, orgId, userId, rolegroupId).Execute()





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
    userId := "userId_example" // string | The ID of the user
    rolegroupId := "rolegroupId_example" // string | The ID of the user's rolegroup

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdUsersUserIdRolegroupsRolegroupIdDelete(context.Background(), orgId, userId, rolegroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdUsersUserIdRolegroupsRolegroupIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**userId** | **string** | The ID of the user | 
**rolegroupId** | **string** | The ID of the user&#39;s rolegroup | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdUsersUserIdRolegroupsRolegroupIdDeleteRequest struct via the builder pattern


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


## OrganizationsOrgIdUsersUserIdRolegroupsRolegroupIdPost

> OrganizationsOrgIdUsersUserIdRolegroupsRolegroupIdPost(ctx, orgId, userId, rolegroupId).Execute()





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
    userId := "userId_example" // string | The ID of the user
    rolegroupId := "rolegroupId_example" // string | The ID of the user's rolegroup

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdUsersUserIdRolegroupsRolegroupIdPost(context.Background(), orgId, userId, rolegroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdUsersUserIdRolegroupsRolegroupIdPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**userId** | **string** | The ID of the user | 
**rolegroupId** | **string** | The ID of the user&#39;s rolegroup | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdUsersUserIdRolegroupsRolegroupIdPostRequest struct via the builder pattern


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

