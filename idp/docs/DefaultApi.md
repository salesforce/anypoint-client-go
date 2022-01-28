# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com/accounts/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsOrgIdIdentityProvidersGet**](DefaultApi.md#OrganizationsOrgIdIdentityProvidersGet) | **Get** /organizations/{orgId}/identityProviders | 
[**OrganizationsOrgIdIdentityProvidersIdpIdDelete**](DefaultApi.md#OrganizationsOrgIdIdentityProvidersIdpIdDelete) | **Delete** /organizations/{orgId}/identityProviders/{idpId} | 
[**OrganizationsOrgIdIdentityProvidersIdpIdGet**](DefaultApi.md#OrganizationsOrgIdIdentityProvidersIdpIdGet) | **Get** /organizations/{orgId}/identityProviders/{idpId} | 
[**OrganizationsOrgIdIdentityProvidersIdpIdPatch**](DefaultApi.md#OrganizationsOrgIdIdentityProvidersIdpIdPatch) | **Patch** /organizations/{orgId}/identityProviders/{idpId} | 
[**OrganizationsOrgIdIdentityProvidersPost**](DefaultApi.md#OrganizationsOrgIdIdentityProvidersPost) | **Post** /organizations/{orgId}/identityProviders | 



## OrganizationsOrgIdIdentityProvidersGet

> InlineResponse200 OrganizationsOrgIdIdentityProvidersGet(ctx, orgId).Execute()





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
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdIdentityProvidersGet(context.Background(), orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdIdentityProvidersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdIdentityProvidersGet`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdIdentityProvidersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdIdentityProvidersGetRequest struct via the builder pattern


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


## OrganizationsOrgIdIdentityProvidersIdpIdDelete

> OrganizationsOrgIdIdentityProvidersIdpIdDelete(ctx, orgId, idpId).Execute()





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
    idpId := "idpId_example" // string | The ID of the Identity Provider in GUID format

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdIdentityProvidersIdpIdDelete(context.Background(), orgId, idpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdIdentityProvidersIdpIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**idpId** | **string** | The ID of the Identity Provider in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdIdentityProvidersIdpIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: plain/text

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrgIdIdentityProvidersIdpIdGet

> Idp OrganizationsOrgIdIdentityProvidersIdpIdGet(ctx, orgId, idpId).Execute()





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
    idpId := "idpId_example" // string | The ID of the Identity Provider in GUID format

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdIdentityProvidersIdpIdGet(context.Background(), orgId, idpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdIdentityProvidersIdpIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdIdentityProvidersIdpIdGet`: Idp
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdIdentityProvidersIdpIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**idpId** | **string** | The ID of the Identity Provider in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdIdentityProvidersIdpIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Idp**](Idp.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: plain/text, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrgIdIdentityProvidersIdpIdPatch

> Idp OrganizationsOrgIdIdentityProvidersIdpIdPatch(ctx, orgId, idpId).IdpPatchBody(idpPatchBody).Execute()





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
    idpId := "idpId_example" // string | The ID of the Identity Provider in GUID format
    idpPatchBody := *openapiclient.NewIdpPatchBody() // IdpPatchBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdIdentityProvidersIdpIdPatch(context.Background(), orgId, idpId).IdpPatchBody(idpPatchBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdIdentityProvidersIdpIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdIdentityProvidersIdpIdPatch`: Idp
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdIdentityProvidersIdpIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**idpId** | **string** | The ID of the Identity Provider in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdIdentityProvidersIdpIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **idpPatchBody** | [**IdpPatchBody**](IdpPatchBody.md) |  | 

### Return type

[**Idp**](Idp.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, plain/text

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrgIdIdentityProvidersPost

> Idp OrganizationsOrgIdIdentityProvidersPost(ctx, orgId).IdpPostBody(idpPostBody).Execute()





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
    idpPostBody := *openapiclient.NewIdpPostBody() // IdpPostBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdIdentityProvidersPost(context.Background(), orgId).IdpPostBody(idpPostBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdIdentityProvidersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdIdentityProvidersPost`: Idp
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdIdentityProvidersPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdIdentityProvidersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idpPostBody** | [**IdpPostBody**](IdpPostBody.md) |  | 

### Return type

[**Idp**](Idp.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

