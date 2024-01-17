# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com/accounts/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsOrgIdUsersGet**](DefaultApi.md#OrganizationsOrgIdUsersGet) | **Get** /organizations/{orgId}/users | 
[**OrganizationsOrgIdUsersPost**](DefaultApi.md#OrganizationsOrgIdUsersPost) | **Post** /organizations/{orgId}/users | 
[**OrganizationsOrgIdUsersUserIdDelete**](DefaultApi.md#OrganizationsOrgIdUsersUserIdDelete) | **Delete** /organizations/{orgId}/users/{userId} | 
[**OrganizationsOrgIdUsersUserIdGet**](DefaultApi.md#OrganizationsOrgIdUsersUserIdGet) | **Get** /organizations/{orgId}/users/{userId} | 
[**OrganizationsOrgIdUsersUserIdPut**](DefaultApi.md#OrganizationsOrgIdUsersUserIdPut) | **Put** /organizations/{orgId}/users/{userId} | 



## OrganizationsOrgIdUsersGet

> OrganizationsOrgIdUsersGet200Response OrganizationsOrgIdUsersGet(ctx, orgId).Type_(type_).Limit(limit).Offset(offset).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/user"
)

func main() {
    orgId := "orgId_example" // string | The ID of the organization in GUID format
    type_ := "type__example" // string | A search string to specify the type of user you want to retrieve. (optional)
    limit := int32(56) // int32 | Maximum number of users to retrieve per request. (optional)
    offset := int32(56) // int32 | The number of records to omit from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OrganizationsOrgIdUsersGet(context.Background(), orgId).Type_(type_).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdUsersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdUsersGet`: OrganizationsOrgIdUsersGet200Response
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

[**OrganizationsOrgIdUsersGet200Response**](OrganizationsOrgIdUsersGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrgIdUsersPost

> User OrganizationsOrgIdUsersPost(ctx, orgId).UserPostBody(userPostBody).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/user"
)

func main() {
    orgId := "orgId_example" // string | The ID of the organization in GUID format
    userPostBody := *openapiclient.NewUserPostBody() // UserPostBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OrganizationsOrgIdUsersPost(context.Background(), orgId).UserPostBody(userPostBody).Execute()
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

 **userPostBody** | [**UserPostBody**](UserPostBody.md) |  | 

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


## OrganizationsOrgIdUsersUserIdDelete

> OrganizationsOrgIdUsersUserIdDelete(ctx, orgId, userId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/user"
)

func main() {
    orgId := "orgId_example" // string | The ID of the organization in GUID format
    userId := "userId_example" // string | The ID of the user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.OrganizationsOrgIdUsersUserIdDelete(context.Background(), orgId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdUsersUserIdDelete``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdUsersUserIdDeleteRequest struct via the builder pattern


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


## OrganizationsOrgIdUsersUserIdGet

> User OrganizationsOrgIdUsersUserIdGet(ctx, orgId, userId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/user"
)

func main() {
    orgId := "orgId_example" // string | The ID of the organization in GUID format
    userId := "userId_example" // string | the ID of the user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OrganizationsOrgIdUsersUserIdGet(context.Background(), orgId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdUsersUserIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdUsersUserIdGet`: User
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdUsersUserIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**userId** | **string** | the ID of the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdUsersUserIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**User**](User.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrgIdUsersUserIdPut

> User OrganizationsOrgIdUsersUserIdPut(ctx, orgId, userId).UserPutBody(userPutBody).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/user"
)

func main() {
    orgId := "orgId_example" // string | The ID of the organization in GUID format
    userId := "userId_example" // string | The ID of the user
    userPutBody := *openapiclient.NewUserPutBody() // UserPutBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OrganizationsOrgIdUsersUserIdPut(context.Background(), orgId, userId).UserPutBody(userPutBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdUsersUserIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdUsersUserIdPut`: User
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdUsersUserIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**userId** | **string** | The ID of the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdUsersUserIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userPutBody** | [**UserPutBody**](UserPutBody.md) |  | 

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

