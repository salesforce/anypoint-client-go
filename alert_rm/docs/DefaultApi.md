# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com/armui/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Addnewalert**](DefaultApi.md#Addnewalert) | **Post** /v1/alerts/cloudhub | 
[**Deleteonealert**](DefaultApi.md#Deleteonealert) | **Delete** /v1/alerts/cloudhub/{cloudhub-rm-alert} | Deleteonealert
[**ModifyOneAlert**](DefaultApi.md#ModifyOneAlert) | **Put** /v1/alerts/cloudhub/{cloudhub-rm-alert} | 
[**V1AlertsCloudhubCloudhubRmAlertGet**](DefaultApi.md#V1AlertsCloudhubCloudhubRmAlertGet) | **Get** /v1/alerts/cloudhub/{cloudhub-rm-alert} | Get one specific alert
[**V1AlertsGet**](DefaultApi.md#V1AlertsGet) | **Get** /v1/alerts | Get all the alert instances from the run time manager



## Addnewalert

> AlertWithId Addnewalert(ctx).OrgId(orgId).Alert(alert).Authorization(authorization).EnvId(envId).Execute()





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
    orgId := "orgId_example" // string | 
    alert := *openapiclient.NewAlert("Name_example", "Severity_example", []openapiclient.Action{*openapiclient.NewAction("Type_example", "Content_example", "Subject_example", []string{"UserIds_example"})}, *openapiclient.NewCondition("Operator_example", int32(123), int32(123), "ResourceType_example", "Type_example", int32(123), []string{"Resources_example"})) // Alert | Add a new alert
    authorization := "authorization_example" // string |  (optional) (default to "Bearer {token}")
    envId := "envId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.Addnewalert(context.Background()).OrgId(orgId).Alert(alert).Authorization(authorization).EnvId(envId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Addnewalert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Addnewalert`: AlertWithId
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Addnewalert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddnewalertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgId** | **string** |  | 
 **alert** | [**Alert**](Alert.md) | Add a new alert | 
 **authorization** | **string** |  | [default to &quot;Bearer {token}&quot;]
 **envId** | **string** |  | 

### Return type

[**AlertWithId**](AlertWithId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Deleteonealert

> Deleteonealert(ctx, cloudhubRmAlert).OrgId(orgId).EnvId(envId).Authorization(authorization).Execute()

Deleteonealert



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
    orgId := "orgId_example" // string | the org id
    envId := "envId_example" // string | the env id
    cloudhubRmAlert := "cloudhubRmAlert_example" // string | the name of the cloud hub alert
    authorization := "authorization_example" // string |  (optional) (default to "Bearer {token}")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.Deleteonealert(context.Background(), cloudhubRmAlert).OrgId(orgId).EnvId(envId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Deleteonealert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudhubRmAlert** | **string** | the name of the cloud hub alert | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteonealertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgId** | **string** | the org id | 
 **envId** | **string** | the env id | 

 **authorization** | **string** |  | [default to &quot;Bearer {token}&quot;]

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyOneAlert

> AlertWithId ModifyOneAlert(ctx, cloudhubRmAlert).OrgId(orgId).EnvId(envId).Authorization(authorization).AlertWithId(alertWithId).Execute()





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
    orgId := "orgId_example" // string | the org id
    envId := "envId_example" // string | the env id
    cloudhubRmAlert := "cloudhubRmAlert_example" // string | the name of the cloud hub alert
    authorization := "authorization_example" // string | authorization key (optional) (default to "Bearer {token}")
    alertWithId := *openapiclient.NewAlertWithId("Name_example", "Severity_example", []openapiclient.Action{*openapiclient.NewAction("Type_example", "Content_example", "Subject_example", []string{"UserIds_example"})}, *openapiclient.NewCondition("Operator_example", int32(123), int32(123), "ResourceType_example", "Type_example", int32(123), []string{"Resources_example"})) // AlertWithId |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ModifyOneAlert(context.Background(), cloudhubRmAlert).OrgId(orgId).EnvId(envId).Authorization(authorization).AlertWithId(alertWithId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyOneAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyOneAlert`: AlertWithId
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyOneAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudhubRmAlert** | **string** | the name of the cloud hub alert | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyOneAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgId** | **string** | the org id | 
 **envId** | **string** | the env id | 

 **authorization** | **string** | authorization key | [default to &quot;Bearer {token}&quot;]
 **alertWithId** | [**AlertWithId**](AlertWithId.md) |  | 

### Return type

[**AlertWithId**](AlertWithId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AlertsCloudhubCloudhubRmAlertGet

> []AlertWithId V1AlertsCloudhubCloudhubRmAlertGet(ctx, cloudhubRmAlert).OrgId(orgId).EnvId(envId).Authorization(authorization).Execute()

Get one specific alert



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
    orgId := "orgId_example" // string | the org id
    envId := "envId_example" // string | the env id
    cloudhubRmAlert := "cloudhubRmAlert_example" // string | the name of the cloud hub alert
    authorization := "authorization_example" // string |  (optional) (default to "Bearer {token}")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.V1AlertsCloudhubCloudhubRmAlertGet(context.Background(), cloudhubRmAlert).OrgId(orgId).EnvId(envId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.V1AlertsCloudhubCloudhubRmAlertGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AlertsCloudhubCloudhubRmAlertGet`: []AlertWithId
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.V1AlertsCloudhubCloudhubRmAlertGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudhubRmAlert** | **string** | the name of the cloud hub alert | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AlertsCloudhubCloudhubRmAlertGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgId** | **string** | the org id | 
 **envId** | **string** | the env id | 

 **authorization** | **string** |  | [default to &quot;Bearer {token}&quot;]

### Return type

[**[]AlertWithId**](AlertWithId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AlertsGet

> []AlertWithId V1AlertsGet(ctx).OrgId(orgId).EnvId(envId).Execute()

Get all the alert instances from the run time manager



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
    orgId := "orgId_example" // string | The organization ID
    envId := "envId_example" // string | The environment ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.V1AlertsGet(context.Background()).OrgId(orgId).EnvId(envId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.V1AlertsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AlertsGet`: []AlertWithId
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.V1AlertsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1AlertsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgId** | **string** | The organization ID | 
 **envId** | **string** | The environment ID | 

### Return type

[**[]AlertWithId**](AlertWithId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

