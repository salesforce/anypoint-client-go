# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com/accounts/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsOrgIdUsersDelete**](DefaultApi.md#OrganizationsOrgIdUsersDelete) | **Delete** /organizations/{orgId}/users | 
[**OrganizationsOrgIdUsersGet**](DefaultApi.md#OrganizationsOrgIdUsersGet) | **Get** /organizations/{orgId}/users | 
[**OrganizationsOrgIdUsersPost**](DefaultApi.md#OrganizationsOrgIdUsersPost) | **Post** /organizations/{orgId}/users | 
[**OrganizationsOrgIdUsersPut**](DefaultApi.md#OrganizationsOrgIdUsersPut) | **Put** /organizations/{orgId}/users | 



## OrganizationsOrgIdUsersDelete

> OrganizationsOrgIdUsersDelete(ctx, orgId).RequestBody(requestBody).Execute()





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
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdUsersDelete(context.Background(), orgId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdUsersDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiOrganizationsOrgIdUsersDeleteRequest struct via the builder pattern


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


## OrganizationsOrgIdUsersGet

> InlineResponse200 OrganizationsOrgIdUsersGet(ctx, orgId).Type_(type_).Limit(limit).Offset(offset).Execute()





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
    type_ := "type__example" // string | A search string to specify the type of user you want to retrieve. (optional)
    limit := int32(56) // int32 | Maximum number of users to retrieve per request. (optional)
    offset := int32(56) // int32 | The number of records to omit from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdUsersGet(context.Background(), orgId).Type_(type_).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdUsersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdUsersGet`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdUsersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdUsersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | A search string to specify the type of user you want to retrieve. | 
 **limit** | **int32** | Maximum number of users to retrieve per request. | 
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


## OrganizationsOrgIdUsersPost

> User OrganizationsOrgIdUsersPost(ctx, orgId).UsersPostBody(usersPostBody).Execute()





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
    usersPostBody := *openapiclient.NewUsersPostBody() // UsersPostBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdUsersPost(context.Background(), orgId).UsersPostBody(usersPostBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdUsersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdUsersPost`: User
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdUsersPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdUsersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **usersPostBody** | [**UsersPostBody**](UsersPostBody.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrgIdUsersPut

> OrganizationsOrgIdUsersPut(ctx, orgId).RequestBody(requestBody).Execute()





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
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdUsersPut(context.Background(), orgId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdUsersPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiOrganizationsOrgIdUsersPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]map[string]interface{}** |  | 

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

