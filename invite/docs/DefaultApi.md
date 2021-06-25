# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com/accounts/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsOrgIdInvitesDelete**](DefaultApi.md#OrganizationsOrgIdInvitesDelete) | **Delete** /organizations/{orgId}/invites | 
[**OrganizationsOrgIdInvitesGet**](DefaultApi.md#OrganizationsOrgIdInvitesGet) | **Get** /organizations/{orgId}/invites | 
[**OrganizationsOrgIdInvitesPost**](DefaultApi.md#OrganizationsOrgIdInvitesPost) | **Post** /organizations/{orgId}/invites | 



## OrganizationsOrgIdInvitesDelete

> OrganizationsOrgIdInvitesDelete(ctx, orgId).RequestBody(requestBody).Execute()





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
    requestBody := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdInvitesDelete(context.Background(), orgId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdInvitesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdInvitesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrgIdInvitesGet

> InlineResponse200 OrganizationsOrgIdInvitesGet(ctx, orgId).Search(search).Limit(limit).Offset(offset).Execute()





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
    search := "search_example" // string | A search string to use for partial matches of invited emails (optional)
    limit := int32(56) // int32 | Pagination parameter for choosing how many results include in the response (optional)
    offset := int32(56) // int32 | Pagination parameter to start returning results from the specified position of matches (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdInvitesGet(context.Background(), orgId).Search(search).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdInvitesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdInvitesGet`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdInvitesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdInvitesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** | A search string to use for partial matches of invited emails | 
 **limit** | **int32** | Pagination parameter for choosing how many results include in the response | 
 **offset** | **int32** | Pagination parameter to start returning results from the specified position of matches | 

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


## OrganizationsOrgIdInvitesPost

> Invite OrganizationsOrgIdInvitesPost(ctx, orgId).InvitePostBody(invitePostBody).Execute()





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
    invitePostBody := *openapiclient.NewInvitePostBody() // InvitePostBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdInvitesPost(context.Background(), orgId).InvitePostBody(invitePostBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdInvitesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdInvitesPost`: Invite
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdInvitesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdInvitesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invitePostBody** | [**InvitePostBody**](InvitePostBody.md) |  | 

### Return type

[**Invite**](Invite.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

