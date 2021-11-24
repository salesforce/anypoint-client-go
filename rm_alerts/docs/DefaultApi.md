# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com/cloudhub/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlertsAlertIdDelete**](DefaultApi.md#AlertsAlertIdDelete) | **Delete** /alerts/{alertId} | Delete a specific alert
[**AlertsAlertIdGet**](DefaultApi.md#AlertsAlertIdGet) | **Get** /alerts/{alertId} | Get one specific alert
[**AlertsGet**](DefaultApi.md#AlertsGet) | **Get** /alerts | Get all alets for a given environment.
[**AlertsPost**](DefaultApi.md#AlertsPost) | **Post** /alerts | 
[**ModifyOneAlert**](DefaultApi.md#ModifyOneAlert) | **Put** /alerts/{alertId} | 



## AlertsAlertIdDelete

> AlertsAlertIdDelete(ctx, alertId).Execute()

Delete a specific alert



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
    alertId := "alertId_example" // string | the alert Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AlertsAlertIdDelete(context.Background(), alertId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AlertsAlertIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **string** | the alert Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertsAlertIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## AlertsAlertIdGet

> []Alert AlertsAlertIdGet(ctx, alertId).Execute()

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
    alertId := "alertId_example" // string | the alert Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AlertsAlertIdGet(context.Background(), alertId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AlertsAlertIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsAlertIdGet`: []Alert
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AlertsAlertIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **string** | the alert Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertsAlertIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Alert**](Alert.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertsGet

> []Alert AlertsGet(ctx).XANYPNTENVID(xANYPNTENVID).Offset(offset).Limit(limit).Resource(resource).Execute()

Get all alets for a given environment.



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
    xANYPNTENVID := "xANYPNTENVID_example" // string | The ID of your current environment
    offset := int32(56) // int32 | Offset to return alerts from (optional)
    limit := int32(56) // int32 | Maximum number of alerts to return (optional)
    resource := "resource_example" // string | Only return alerts which are connected to this resource (application name) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AlertsGet(context.Background()).XANYPNTENVID(xANYPNTENVID).Offset(offset).Limit(limit).Resource(resource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AlertsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsGet`: []Alert
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AlertsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xANYPNTENVID** | **string** | The ID of your current environment | 
 **offset** | **int32** | Offset to return alerts from | 
 **limit** | **int32** | Maximum number of alerts to return | 
 **resource** | **string** | Only return alerts which are connected to this resource (application name) | 

### Return type

[**[]Alert**](Alert.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertsPost

> Alert AlertsPost(ctx).AlertBody(alertBody).XANYPNTENVID(xANYPNTENVID).Execute()





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
    alertBody := *openapiclient.NewAlertBody() // AlertBody | Add a new alert
    xANYPNTENVID := "xANYPNTENVID_example" // string | The ID of your current environment (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AlertsPost(context.Background()).AlertBody(alertBody).XANYPNTENVID(xANYPNTENVID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AlertsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsPost`: Alert
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AlertsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alertBody** | [**AlertBody**](AlertBody.md) | Add a new alert | 
 **xANYPNTENVID** | **string** | The ID of your current environment | 

### Return type

[**Alert**](Alert.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyOneAlert

> Alert ModifyOneAlert(ctx, alertId).AlertBody(alertBody).Execute()





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
    alertId := "alertId_example" // string | the alert Id
    alertBody := *openapiclient.NewAlertBody() // AlertBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ModifyOneAlert(context.Background(), alertId).AlertBody(alertBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyOneAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyOneAlert`: Alert
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyOneAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **string** | the alert Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyOneAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alertBody** | [**AlertBody**](AlertBody.md) |  | 

### Return type

[**Alert**](Alert.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

