# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com/secrets-manager/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSecretGroupCrlDistribCfgsDetails**](DefaultApi.md#GetSecretGroupCrlDistribCfgsDetails) | **Get** /organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId}/crlDistributorConfigs/{secretId} | Retrieve crl-distributor-configs details
[**GetSecretGroupCrlDistribCfgsList**](DefaultApi.md#GetSecretGroupCrlDistribCfgsList) | **Get** /organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId}/crlDistributorConfigs | Retrieves a secret-groups&#39; collection of crl-distributor-configs.
[**PostSecretGroupCrlDistribCfgs**](DefaultApi.md#PostSecretGroupCrlDistribCfgs) | **Post** /organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId}/crlDistributorConfigs | Create a secret-groups&#39; crl-distributor-configs.
[**PutSecretGroupTlsContext**](DefaultApi.md#PutSecretGroupTlsContext) | **Put** /organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId}/crlDistributorConfigs/{secretId} | Update a given secret-group tls-context



## GetSecretGroupCrlDistribCfgsDetails

> CrlDistribCfgsDetails GetSecretGroupCrlDistribCfgsDetails(ctx, orgId, envId, secretGroupId, secretId).Execute()

Retrieve crl-distributor-configs details



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/secretgroup_crl_distributor_configs"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    secretGroupId := "secretGroupId_example" // string | The secret group id
    secretId := "secretId_example" // string | The crl distributor configurations id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetSecretGroupCrlDistribCfgsDetails(context.Background(), orgId, envId, secretGroupId, secretId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSecretGroupCrlDistribCfgsDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecretGroupCrlDistribCfgsDetails`: CrlDistribCfgsDetails
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSecretGroupCrlDistribCfgsDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment id | 
**secretGroupId** | **string** | The secret group id | 
**secretId** | **string** | The crl distributor configurations id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretGroupCrlDistribCfgsDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**CrlDistribCfgsDetails**](CrlDistribCfgsDetails.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecretGroupCrlDistribCfgsList

> []CrlDistribCfgSummary GetSecretGroupCrlDistribCfgsList(ctx, orgId, envId, secretGroupId).Execute()

Retrieves a secret-groups' collection of crl-distributor-configs.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/secretgroup_crl_distributor_configs"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    secretGroupId := "secretGroupId_example" // string | The secret group id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetSecretGroupCrlDistribCfgsList(context.Background(), orgId, envId, secretGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSecretGroupCrlDistribCfgsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecretGroupCrlDistribCfgsList`: []CrlDistribCfgSummary
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSecretGroupCrlDistribCfgsList`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetSecretGroupCrlDistribCfgsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]CrlDistribCfgSummary**](CrlDistribCfgSummary.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSecretGroupCrlDistribCfgs

> PostSecretGroupCrlDistribCfgs201Response PostSecretGroupCrlDistribCfgs(ctx, orgId, envId, secretGroupId).CrlDistribCfgsReqBody(crlDistribCfgsReqBody).Execute()

Create a secret-groups' crl-distributor-configs.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/secretgroup_crl_distributor_configs"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    secretGroupId := "secretGroupId_example" // string | The secret group id
    crlDistribCfgsReqBody := *openapiclient.NewCrlDistribCfgsReqBody() // CrlDistribCfgsReqBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.PostSecretGroupCrlDistribCfgs(context.Background(), orgId, envId, secretGroupId).CrlDistribCfgsReqBody(crlDistribCfgsReqBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PostSecretGroupCrlDistribCfgs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostSecretGroupCrlDistribCfgs`: PostSecretGroupCrlDistribCfgs201Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PostSecretGroupCrlDistribCfgs`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiPostSecretGroupCrlDistribCfgsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **crlDistribCfgsReqBody** | [**CrlDistribCfgsReqBody**](CrlDistribCfgsReqBody.md) |  | 

### Return type

[**PostSecretGroupCrlDistribCfgs201Response**](PostSecretGroupCrlDistribCfgs201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSecretGroupTlsContext

> PutSecretGroupTlsContext200Response PutSecretGroupTlsContext(ctx, orgId, envId, secretGroupId, secretId).CrlDistribCfgsReqBody(crlDistribCfgsReqBody).Execute()

Update a given secret-group tls-context



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/secretgroup_crl_distributor_configs"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    secretGroupId := "secretGroupId_example" // string | The secret group id
    secretId := "secretId_example" // string | The crl distributor configurations id
    crlDistribCfgsReqBody := *openapiclient.NewCrlDistribCfgsReqBody() // CrlDistribCfgsReqBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.PutSecretGroupTlsContext(context.Background(), orgId, envId, secretGroupId, secretId).CrlDistribCfgsReqBody(crlDistribCfgsReqBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PutSecretGroupTlsContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutSecretGroupTlsContext`: PutSecretGroupTlsContext200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PutSecretGroupTlsContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment id | 
**secretGroupId** | **string** | The secret group id | 
**secretId** | **string** | The crl distributor configurations id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSecretGroupTlsContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **crlDistribCfgsReqBody** | [**CrlDistribCfgsReqBody**](CrlDistribCfgsReqBody.md) |  | 

### Return type

[**PutSecretGroupTlsContext200Response**](PutSecretGroupTlsContext200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

