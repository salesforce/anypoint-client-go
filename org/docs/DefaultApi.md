# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com/accounts/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsOrgIdDelete**](DefaultApi.md#OrganizationsOrgIdDelete) | **Delete** /organizations/{orgId} | Delete a Business Group by its id.
[**OrganizationsOrgIdGet**](DefaultApi.md#OrganizationsOrgIdGet) | **Get** /organizations/{orgId} | Returns the business Group instance with the given id.
[**OrganizationsOrgIdPut**](DefaultApi.md#OrganizationsOrgIdPut) | **Put** /organizations/{orgId} | Put a Business Group by its id.
[**OrganizationsPost**](DefaultApi.md#OrganizationsPost) | **Post** /organizations | Creates a new Business Group within an organization.



## OrganizationsOrgIdDelete

> map[string]interface{} OrganizationsOrgIdDelete(ctx, orgId).Execute()

Delete a Business Group by its id.



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
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdDelete(context.Background(), orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdDelete`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrgIdGet

> MasterBGDetail OrganizationsOrgIdGet(ctx, orgId).Execute()

Returns the business Group instance with the given id.



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
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdGet(context.Background(), orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdGet`: MasterBGDetail
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MasterBGDetail**](MasterBGDetail.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrgIdPut

> BGCore OrganizationsOrgIdPut(ctx, orgId).BGPutReqBody(bGPutReqBody).Execute()

Put a Business Group by its id.



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
    bGPutReqBody := *openapiclient.NewBGPutReqBody("TODO", "TODO", "TODO", "ParentOrganizationId_example", int32(123)) // BGPutReqBody | Business Group Request Body Object

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdPut(context.Background(), orgId).BGPutReqBody(bGPutReqBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdPut`: BGCore
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bGPutReqBody** | [**BGPutReqBody**](BGPutReqBody.md) | Business Group Request Body Object | 

### Return type

[**BGCore**](BGCore.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsPost

> BGCore OrganizationsPost(ctx).BGPostReqBody(bGPostReqBody).Execute()

Creates a new Business Group within an organization.



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
    bGPostReqBody := *openapiclient.NewBGPostReqBody("TODO", "TODO", "TODO", "ParentOrganizationId_example") // BGPostReqBody | Business Group Request Body Object

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsPost(context.Background()).BGPostReqBody(bGPostReqBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsPost`: BGCore
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bGPostReqBody** | [**BGPostReqBody**](BGPostReqBody.md) | Business Group Request Body Object | 

### Return type

[**BGCore**](BGCore.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/html, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

