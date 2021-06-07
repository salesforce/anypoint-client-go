# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com/accounts/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsOrgIdRolegroupsGet**](DefaultApi.md#OrganizationsOrgIdRolegroupsGet) | **Get** /organizations/{orgId}/rolegroups | 
[**OrganizationsOrgIdRolegroupsPost**](DefaultApi.md#OrganizationsOrgIdRolegroupsPost) | **Post** /organizations/{orgId}/rolegroups | 
[**OrganizationsOrgIdRolegroupsRolegroupIdDelete**](DefaultApi.md#OrganizationsOrgIdRolegroupsRolegroupIdDelete) | **Delete** /organizations/{orgId}/rolegroups/{rolegroupId} | 
[**OrganizationsOrgIdRolegroupsRolegroupIdGet**](DefaultApi.md#OrganizationsOrgIdRolegroupsRolegroupIdGet) | **Get** /organizations/{orgId}/rolegroups/{rolegroupId} | 
[**OrganizationsOrgIdRolegroupsRolegroupIdPut**](DefaultApi.md#OrganizationsOrgIdRolegroupsRolegroupIdPut) | **Put** /organizations/{orgId}/rolegroups/{rolegroupId} | 



## OrganizationsOrgIdRolegroupsGet

> InlineResponse200 OrganizationsOrgIdRolegroupsGet(ctx, orgId).Execute()





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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdRolegroupsGet(context.Background(), orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdRolegroupsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdRolegroupsGet`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdRolegroupsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdRolegroupsGetRequest struct via the builder pattern


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


## OrganizationsOrgIdRolegroupsPost

> Rolegroup OrganizationsOrgIdRolegroupsPost(ctx, orgId).RolegroupPostBody(rolegroupPostBody).Execute()





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
    rolegroupPostBody := *openapiclient.NewRolegroupPostBody() // RolegroupPostBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdRolegroupsPost(context.Background(), orgId).RolegroupPostBody(rolegroupPostBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdRolegroupsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdRolegroupsPost`: Rolegroup
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdRolegroupsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdRolegroupsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rolegroupPostBody** | [**RolegroupPostBody**](RolegroupPostBody.md) |  | 

### Return type

[**Rolegroup**](Rolegroup.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrgIdRolegroupsRolegroupIdDelete

> []string OrganizationsOrgIdRolegroupsRolegroupIdDelete(ctx, orgId, rolegroupId).Execute()





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
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdRolegroupsRolegroupIdDelete(context.Background(), orgId, rolegroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdRolegroupsRolegroupIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdRolegroupsRolegroupIdDelete`: []string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdRolegroupsRolegroupIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**rolegroupId** | **string** | The id of a rolegroup | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdRolegroupsRolegroupIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**[]string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrgIdRolegroupsRolegroupIdGet

> Rolegroup OrganizationsOrgIdRolegroupsRolegroupIdGet(ctx, orgId, rolegroupId).Execute()





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
    rolegroupId := "rolegroupId_example" // string | The id of rolegroup

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdRolegroupsRolegroupIdGet(context.Background(), orgId, rolegroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdRolegroupsRolegroupIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdRolegroupsRolegroupIdGet`: Rolegroup
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdRolegroupsRolegroupIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**rolegroupId** | **string** | The id of rolegroup | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdRolegroupsRolegroupIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Rolegroup**](Rolegroup.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrgIdRolegroupsRolegroupIdPut

> Rolegroup OrganizationsOrgIdRolegroupsRolegroupIdPut(ctx, orgId, rolegroupId).RolegroupPutBody(rolegroupPutBody).Execute()





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
    rolegroupId := "rolegroupId_example" // string | The id of an rolegroup
    rolegroupPutBody := *openapiclient.NewRolegroupPutBody() // RolegroupPutBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdRolegroupsRolegroupIdPut(context.Background(), orgId, rolegroupId).RolegroupPutBody(rolegroupPutBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdRolegroupsRolegroupIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdRolegroupsRolegroupIdPut`: Rolegroup
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdRolegroupsRolegroupIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**rolegroupId** | **string** | The id of an rolegroup | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdRolegroupsRolegroupIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rolegroupPutBody** | [**RolegroupPutBody**](RolegroupPutBody.md) |  | 

### Return type

[**Rolegroup**](Rolegroup.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

