# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com/apimanager/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAlertForAPIManagerInstance**](DefaultApi.md#CreateAlertForAPIManagerInstance) | **Post** /v1/organizations/{orgId}/environments/{envId}/apis/{apiVersion}/alerts | 
[**DeleteanAlertfromAPImanager**](DefaultApi.md#DeleteanAlertfromAPImanager) | **Delete** /v1/organizations/{orgId}/environments/{envId}/apis/{apiVersion}/alerts/{cloudhubApimAlert} | DeleteanAlertfromAPImanager
[**Getonealert**](DefaultApi.md#Getonealert) | **Get** /v1/organizations/{orgId}/environments/{envId}/apis/{apiVersion}/alerts/{cloudhubApimAlert} | Getonealert
[**ModifyonealertfromAPImanger**](DefaultApi.md#ModifyonealertfromAPImanger) | **Patch** /v1/organizations/{orgId}/environments/{envId}/apis/{apiVersion}/alerts/{cloudhubApimAlert} | ModifyonealertfromAPImanger
[**V1OrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsGet**](DefaultApi.md#V1OrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsGet) | **Get** /v1/organizations/{orgId}/environments/{envId}/apis/{apiVersion}/alerts | GetAlertsfromAPImanager



## CreateAlertForAPIManagerInstance

> AlertWithId CreateAlertForAPIManagerInstance(ctx, orgId, envId, apiVersion).Alert(alert).Execute()





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
    envId := "envId_example" // string | The environment id
    apiVersion := "apiVersion_example" // string | The api version
    alert := *openapiclient.NewAlert("ApiAlertsVersion_example", "Name_example", "Type_example", false, "Severity_example", []openapiclient.Recipient{*openapiclient.NewRecipient("Type_example", "Value_example", "FirstName_example", "LastName_example")}, *openapiclient.NewCondition("ResourceType_example", "Aggregate_example", "Operator_example", int32(123)), *openapiclient.NewPeriod(*openapiclient.NewDuration(int32(123), "Weight_example"), int32(123))) // Alert |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateAlertForAPIManagerInstance(context.Background(), orgId, envId, apiVersion).Alert(alert).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateAlertForAPIManagerInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAlertForAPIManagerInstance`: AlertWithId
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateAlertForAPIManagerInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment id | 
**apiVersion** | **string** | The api version | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAlertForAPIManagerInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **alert** | [**Alert**](Alert.md) |  | 

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


## DeleteanAlertfromAPImanager

> DeleteanAlertfromAPImanager(ctx, orgId, envId, apiVersion, cloudhubApimAlert).Execute()

DeleteanAlertfromAPImanager



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
    envId := "envId_example" // string | the environment id
    apiVersion := "apiVersion_example" // string | the api version
    cloudhubApimAlert := "cloudhubApimAlert_example" // string | the id of the alert

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteanAlertfromAPImanager(context.Background(), orgId, envId, apiVersion, cloudhubApimAlert).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteanAlertfromAPImanager``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | the environment id | 
**apiVersion** | **string** | the api version | 
**cloudhubApimAlert** | **string** | the id of the alert | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteanAlertfromAPImanagerRequest struct via the builder pattern


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


## Getonealert

> AlertWithId Getonealert(ctx, orgId, envId, apiVersion, cloudhubApimAlert).Execute()

Getonealert



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
    envId := "envId_example" // string | the environment id
    apiVersion := "apiVersion_example" // string | the api version
    cloudhubApimAlert := "cloudhubApimAlert_example" // string | the id of the alert

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.Getonealert(context.Background(), orgId, envId, apiVersion, cloudhubApimAlert).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Getonealert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Getonealert`: AlertWithId
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Getonealert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | the environment id | 
**apiVersion** | **string** | the api version | 
**cloudhubApimAlert** | **string** | the id of the alert | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetonealertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**AlertWithId**](AlertWithId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyonealertfromAPImanger

> AlertWithId ModifyonealertfromAPImanger(ctx, orgId, envId, apiVersion, cloudhubApimAlert).AlertWithId(alertWithId).Execute()

ModifyonealertfromAPImanger



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
    envId := "envId_example" // string | the environment id
    apiVersion := "apiVersion_example" // string | the api version
    cloudhubApimAlert := "cloudhubApimAlert_example" // string | the id of the alert
    alertWithId := *openapiclient.NewAlertWithId("ApiAlertsVersion_example", "Name_example", "Type_example", false, "Severity_example", []openapiclient.Recipient{*openapiclient.NewRecipient("Type_example", "Value_example", "FirstName_example", "LastName_example")}, *openapiclient.NewCondition("ResourceType_example", "Aggregate_example", "Operator_example", int32(123)), *openapiclient.NewPeriod(*openapiclient.NewDuration(int32(123), "Weight_example"), int32(123)), "Id_example") // AlertWithId |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ModifyonealertfromAPImanger(context.Background(), orgId, envId, apiVersion, cloudhubApimAlert).AlertWithId(alertWithId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyonealertfromAPImanger``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyonealertfromAPImanger`: AlertWithId
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyonealertfromAPImanger`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | the environment id | 
**apiVersion** | **string** | the api version | 
**cloudhubApimAlert** | **string** | the id of the alert | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyonealertfromAPImangerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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


## V1OrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsGet

> []AlertWithId V1OrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsGet(ctx, orgId, envId, apiVersion).Execute()

GetAlertsfromAPImanager



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
    envId := "envId_example" // string | The environment id
    apiVersion := "apiVersion_example" // string | The api version

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.V1OrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsGet(context.Background(), orgId, envId, apiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.V1OrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1OrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsGet`: []AlertWithId
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.V1OrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment id | 
**apiVersion** | **string** | The api version | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1OrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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

