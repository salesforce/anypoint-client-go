# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFlexGatewayRegistrationToken**](DefaultApi.md#GetFlexGatewayRegistrationToken) | **Post** /standalone/api/v1/organizations/{orgId}/environments/{envId}/gatewaytokens | Retrieves the flex gateway registration token.
[**GetFlexGatewayTargetApis**](DefaultApi.md#GetFlexGatewayTargetApis) | **Get** /apimanager/xapi/v1/organizations/{orgId}/environments/{envId}/flex-gateway-targets/{flexGatewayTargetId}/apis | Retrieves all APIs within a particular flex gateway target
[**GetFlexGatewayTargetById**](DefaultApi.md#GetFlexGatewayTargetById) | **Get** /apimanager/xapi/v1/organizations/{orgId}/environments/{envId}/flex-gateway-targets/{flexGatewayTargetId} | Retrieves a particular flex gateway by Id
[**GetFlexGatewayTargets**](DefaultApi.md#GetFlexGatewayTargets) | **Get** /apimanager/xapi/v1/organizations/{orgId}/environments/{envId}/flex-gateway-targets | Retrieves all flex gateways



## GetFlexGatewayRegistrationToken

> FlexGatewayRegistrationToken GetFlexGatewayRegistrationToken(ctx, orgId, envId).Execute()

Retrieves the flex gateway registration token.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/flexgateway"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetFlexGatewayRegistrationToken(context.Background(), orgId, envId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetFlexGatewayRegistrationToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFlexGatewayRegistrationToken`: FlexGatewayRegistrationToken
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetFlexGatewayRegistrationToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlexGatewayRegistrationTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FlexGatewayRegistrationToken**](FlexGatewayRegistrationToken.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlexGatewayTargetApis

> FlexGatewayTargetApis GetFlexGatewayTargetApis(ctx, orgId, envId, flexGatewayTargetId).Execute()

Retrieves all APIs within a particular flex gateway target



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/flexgateway"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment Id
    flexGatewayTargetId := "flexGatewayTargetId_example" // string | The flex gateway target Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetFlexGatewayTargetApis(context.Background(), orgId, envId, flexGatewayTargetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetFlexGatewayTargetApis``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFlexGatewayTargetApis`: FlexGatewayTargetApis
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetFlexGatewayTargetApis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment Id | 
**flexGatewayTargetId** | **string** | The flex gateway target Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlexGatewayTargetApisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**FlexGatewayTargetApis**](FlexGatewayTargetApis.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlexGatewayTargetById

> FlexGatewayTargetDetails GetFlexGatewayTargetById(ctx, orgId, envId, flexGatewayTargetId).Execute()

Retrieves a particular flex gateway by Id



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/flexgateway"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment Id
    flexGatewayTargetId := "flexGatewayTargetId_example" // string | The flex gateway target Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetFlexGatewayTargetById(context.Background(), orgId, envId, flexGatewayTargetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetFlexGatewayTargetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFlexGatewayTargetById`: FlexGatewayTargetDetails
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetFlexGatewayTargetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment Id | 
**flexGatewayTargetId** | **string** | The flex gateway target Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlexGatewayTargetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**FlexGatewayTargetDetails**](FlexGatewayTargetDetails.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlexGatewayTargets

> []FlexGatewayTargetSummary GetFlexGatewayTargets(ctx, orgId, envId).Execute()

Retrieves all flex gateways



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/flexgateway"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetFlexGatewayTargets(context.Background(), orgId, envId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetFlexGatewayTargets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFlexGatewayTargets`: []FlexGatewayTargetSummary
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetFlexGatewayTargets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlexGatewayTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]FlexGatewayTargetSummary**](FlexGatewayTargetSummary.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

