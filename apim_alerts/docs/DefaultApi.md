# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com/apimanager/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAlertForAPIManagerInstance**](DefaultApi.md#CreateAlertForAPIManagerInstance) | **Post** /organizations/{orgId}/environments/{envId}/apis/{apiVersion}/alerts | 
[**DeleteanAlertfromAPImanager**](DefaultApi.md#DeleteanAlertfromAPImanager) | **Delete** /organizations/{orgId}/environments/{envId}/apis/{apiVersion}/alerts/{alertId} | DeleteanAlertfromAPImanager
[**ModifyonealertfromAPImanger**](DefaultApi.md#ModifyonealertfromAPImanger) | **Patch** /organizations/{orgId}/environments/{envId}/apis/{apiVersion}/alerts/{alertId} | ModifyonealertfromAPImanger
[**OrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsAlertIdGet**](DefaultApi.md#OrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsAlertIdGet) | **Get** /organizations/{orgId}/environments/{envId}/apis/{apiVersion}/alerts/{alertId} | Getonealert
[**OrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsGet**](DefaultApi.md#OrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsGet) | **Get** /organizations/{orgId}/environments/{envId}/apis/{apiVersion}/alerts | GetAlertsfromAPImanager



## CreateAlertForAPIManagerInstance

> Alert CreateAlertForAPIManagerInstance(ctx, orgId, envId, apiVersion).AlertCore(alertCore).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/apim_alerts"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    apiVersion := "apiVersion_example" // string | The api version
    alertCore := *openapiclient.NewAlertCore("ApiAlertsVersion_example", "Name_example", "Type_example", false, "Severity_example", []openapiclient.Recipient{*openapiclient.NewRecipient("Type_example", "Value_example", "FirstName_example", "LastName_example")}, *openapiclient.NewCondition("ResourceType_example", "Aggregate_example", "Operator_example", int32(123)), *openapiclient.NewPeriod(*openapiclient.NewDuration(int32(123), "Weight_example"), int32(123))) // AlertCore |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateAlertForAPIManagerInstance(context.Background(), orgId, envId, apiVersion).AlertCore(alertCore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateAlertForAPIManagerInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAlertForAPIManagerInstance`: Alert
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



 **alertCore** | [**AlertCore**](AlertCore.md) |  | 

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


## DeleteanAlertfromAPImanager

> DeleteanAlertfromAPImanager(ctx, orgId, envId, apiVersion, alertId).Execute()

DeleteanAlertfromAPImanager



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/apim_alerts"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | the environment id
    apiVersion := "apiVersion_example" // string | the api version
    alertId := "alertId_example" // string | the id of the alert

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeleteanAlertfromAPImanager(context.Background(), orgId, envId, apiVersion, alertId).Execute()
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
**alertId** | **string** | the id of the alert | 

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


## ModifyonealertfromAPImanger

> Alert ModifyonealertfromAPImanger(ctx, orgId, envId, apiVersion, alertId).AlertCore(alertCore).Execute()

ModifyonealertfromAPImanger



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/apim_alerts"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | the environment id
    apiVersion := "apiVersion_example" // string | the api version
    alertId := "alertId_example" // string | the id of the alert
    alertCore := *openapiclient.NewAlertCore("ApiAlertsVersion_example", "Name_example", "Type_example", false, "Severity_example", []openapiclient.Recipient{*openapiclient.NewRecipient("Type_example", "Value_example", "FirstName_example", "LastName_example")}, *openapiclient.NewCondition("ResourceType_example", "Aggregate_example", "Operator_example", int32(123)), *openapiclient.NewPeriod(*openapiclient.NewDuration(int32(123), "Weight_example"), int32(123))) // AlertCore |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ModifyonealertfromAPImanger(context.Background(), orgId, envId, apiVersion, alertId).AlertCore(alertCore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyonealertfromAPImanger``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyonealertfromAPImanger`: Alert
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
**alertId** | **string** | the id of the alert | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyonealertfromAPImangerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **alertCore** | [**AlertCore**](AlertCore.md) |  | 

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


## OrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsAlertIdGet

> Alert OrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsAlertIdGet(ctx, orgId, envId, apiVersion, alertId).Execute()

Getonealert



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/apim_alerts"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | the environment id
    apiVersion := "apiVersion_example" // string | the api version
    alertId := "alertId_example" // string | the id of the alert

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsAlertIdGet(context.Background(), orgId, envId, apiVersion, alertId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsAlertIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsAlertIdGet`: Alert
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsAlertIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | the environment id | 
**apiVersion** | **string** | the api version | 
**alertId** | **string** | the id of the alert | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsAlertIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**Alert**](Alert.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsGet

> []Alert OrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsGet(ctx, orgId, envId, apiVersion).Execute()

GetAlertsfromAPImanager



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/apim_alerts"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    apiVersion := "apiVersion_example" // string | The api version

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsGet(context.Background(), orgId, envId, apiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsGet`: []Alert
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsGet`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiOrganizationsOrgIdEnvironmentsEnvIdApisApiVersionAlertsGetRequest struct via the builder pattern


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

