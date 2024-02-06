# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com/secrets-manager/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSecretGroup**](DefaultApi.md#DeleteSecretGroup) | **Delete** /organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId} | Delete a secret group
[**GetEnvSecretGroups**](DefaultApi.md#GetEnvSecretGroups) | **Get** /organizations/{orgId}/environments/{envId}/secretGroups | Retrieves a collection of secret groups
[**GetSecretGroup**](DefaultApi.md#GetSecretGroup) | **Get** /organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId} | Retrieve a secret group
[**PatchSecretGroup**](DefaultApi.md#PatchSecretGroup) | **Patch** /organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId} | Update a secret group
[**PostSecretGroup**](DefaultApi.md#PostSecretGroup) | **Post** /organizations/{orgId}/environments/{envId}/secretGroups | Creates a secret group



## DeleteSecretGroup

> DeleteSecretGroup(ctx, orgId, envId, secretGroupId).Execute()

Delete a secret group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/secretgroup"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    secretGroupId := "secretGroupId_example" // string | The secret group id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeleteSecretGroup(context.Background(), orgId, envId, secretGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSecretGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment id | 
**secretGroupId** | **string** | The secret group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecretGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvSecretGroups

> []SecretGroup GetEnvSecretGroups(ctx, orgId, envId).Downloadable(downloadable).Execute()

Retrieves a collection of secret groups



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/secretgroup"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    downloadable := true // bool | Filter and fetch list of secret groups based on value of 'downloadable' flag. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetEnvSecretGroups(context.Background(), orgId, envId).Downloadable(downloadable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetEnvSecretGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvSecretGroups`: []SecretGroup
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetEnvSecretGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvSecretGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **downloadable** | **bool** | Filter and fetch list of secret groups based on value of &#39;downloadable&#39; flag. | 

### Return type

[**[]SecretGroup**](SecretGroup.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecretGroup

> SecretGroup GetSecretGroup(ctx, orgId, envId, secretGroupId).Execute()

Retrieve a secret group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/secretgroup"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    secretGroupId := "secretGroupId_example" // string | The secret group id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetSecretGroup(context.Background(), orgId, envId, secretGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSecretGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecretGroup`: SecretGroup
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSecretGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment id | 
**secretGroupId** | **string** | The secret group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**SecretGroup**](SecretGroup.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSecretGroup

> SecretGroupPatchResponse PatchSecretGroup(ctx, orgId, envId, secretGroupId).SecretGroupPatchBody(secretGroupPatchBody).Execute()

Update a secret group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/secretgroup"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    secretGroupId := "secretGroupId_example" // string | The secret group id
    secretGroupPatchBody := *openapiclient.NewSecretGroupPatchBody() // SecretGroupPatchBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.PatchSecretGroup(context.Background(), orgId, envId, secretGroupId).SecretGroupPatchBody(secretGroupPatchBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PatchSecretGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchSecretGroup`: SecretGroupPatchResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PatchSecretGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment id | 
**secretGroupId** | **string** | The secret group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSecretGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **secretGroupPatchBody** | [**SecretGroupPatchBody**](SecretGroupPatchBody.md) |  | 

### Return type

[**SecretGroupPatchResponse**](SecretGroupPatchResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSecretGroup

> SecretGroupPostResponse PostSecretGroup(ctx, orgId, envId).SecretGroupPostBody(secretGroupPostBody).Execute()

Creates a secret group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/secretgroup"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    secretGroupPostBody := *openapiclient.NewSecretGroupPostBody() // SecretGroupPostBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.PostSecretGroup(context.Background(), orgId, envId).SecretGroupPostBody(secretGroupPostBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PostSecretGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostSecretGroup`: SecretGroupPostResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PostSecretGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSecretGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **secretGroupPostBody** | [**SecretGroupPostBody**](SecretGroupPostBody.md) |  | 

### Return type

[**SecretGroupPostResponse**](SecretGroupPostResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

