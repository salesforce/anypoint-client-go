# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com/secrets-manager/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSecretGroupCertificateDetails**](DefaultApi.md#GetSecretGroupCertificateDetails) | **Get** /organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId}/certificates/{secretId} | Retrieve certificate details
[**GetSecretGroupCertificates**](DefaultApi.md#GetSecretGroupCertificates) | **Get** /organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId}/certificates | Retrieves a secret-groups&#39; collection of certificates.
[**PatchSecretGroupCertificate**](DefaultApi.md#PatchSecretGroupCertificate) | **Patch** /organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId}/certificates/{secretId} | Update a given secret-group certificate
[**PostSecretGroupCertificate**](DefaultApi.md#PostSecretGroupCertificate) | **Post** /organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId}/certificates | Create a secret-groups&#39; certificate.
[**PutSecretGroupCertificate**](DefaultApi.md#PutSecretGroupCertificate) | **Put** /organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId}/certificates/{secretId} | Update a given secret-group certificate



## GetSecretGroupCertificateDetails

> Certificate GetSecretGroupCertificateDetails(ctx, orgId, envId, secretGroupId, secretId).Execute()

Retrieve certificate details



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/secretgroup_certificate"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    secretGroupId := "secretGroupId_example" // string | The secret group id
    secretId := "secretId_example" // string | The truststore id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetSecretGroupCertificateDetails(context.Background(), orgId, envId, secretGroupId, secretId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSecretGroupCertificateDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecretGroupCertificateDetails`: Certificate
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSecretGroupCertificateDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment id | 
**secretGroupId** | **string** | The secret group id | 
**secretId** | **string** | The truststore id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretGroupCertificateDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**Certificate**](Certificate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecretGroupCertificates

> []CertificateSummary GetSecretGroupCertificates(ctx, orgId, envId, secretGroupId).Execute()

Retrieves a secret-groups' collection of certificates.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/secretgroup_certificate"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    secretGroupId := "secretGroupId_example" // string | The secret group id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetSecretGroupCertificates(context.Background(), orgId, envId, secretGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSecretGroupCertificates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecretGroupCertificates`: []CertificateSummary
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSecretGroupCertificates`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetSecretGroupCertificatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]CertificateSummary**](CertificateSummary.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSecretGroupCertificate

> PutSecretGroupCertificate200Response PatchSecretGroupCertificate(ctx, orgId, envId, secretGroupId, secretId).Body(body).Execute()

Update a given secret-group certificate



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/secretgroup_certificate"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    secretGroupId := "secretGroupId_example" // string | The secret group id
    secretId := "secretId_example" // string | The truststore id
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.PatchSecretGroupCertificate(context.Background(), orgId, envId, secretGroupId, secretId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PatchSecretGroupCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchSecretGroupCertificate`: PutSecretGroupCertificate200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PatchSecretGroupCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment id | 
**secretGroupId** | **string** | The secret group id | 
**secretId** | **string** | The truststore id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSecretGroupCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | **map[string]interface{}** |  | 

### Return type

[**PutSecretGroupCertificate200Response**](PutSecretGroupCertificate200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json:
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSecretGroupCertificate

> PostSecretGroupCertificate201Response PostSecretGroupCertificate(ctx, orgId, envId, secretGroupId).AllowExpiredCert(allowExpiredCert).ExpirationDate(expirationDate).CertStore(certStore).Type_(type_).Name(name).Execute()

Create a secret-groups' certificate.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/secretgroup_certificate"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    secretGroupId := "secretGroupId_example" // string | The secret group id
    allowExpiredCert := true // bool | With 'true' to allow uploading expired certificates
    expirationDate := "expirationDate_example" // string | Date on which this secret should expire. If not set, by default, it will be set to notAfter date of this certificate.  Once the secret expires, a grant can not be requested for it.  (optional)
    certStore := os.NewFile(1234, "some_file") // *os.File | The file containing the certificate in PEM format (optional)
    type_ := "type__example" // string | Type of certificate supported (optional)
    name := "name_example" // string | The name of the certificate (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.PostSecretGroupCertificate(context.Background(), orgId, envId, secretGroupId).AllowExpiredCert(allowExpiredCert).ExpirationDate(expirationDate).CertStore(certStore).Type_(type_).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PostSecretGroupCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostSecretGroupCertificate`: PostSecretGroupCertificate201Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PostSecretGroupCertificate`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiPostSecretGroupCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **allowExpiredCert** | **bool** | With &#39;true&#39; to allow uploading expired certificates | 
 **expirationDate** | **string** | Date on which this secret should expire. If not set, by default, it will be set to notAfter date of this certificate.  Once the secret expires, a grant can not be requested for it.  | 
 **certStore** | ***os.File** | The file containing the certificate in PEM format | 
 **type_** | **string** | Type of certificate supported | 
 **name** | **string** | The name of the certificate | 

### Return type

[**PostSecretGroupCertificate201Response**](PostSecretGroupCertificate201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSecretGroupCertificate

> PutSecretGroupCertificate200Response PutSecretGroupCertificate(ctx, orgId, envId, secretGroupId, secretId).AllowExpiredCert(allowExpiredCert).ExpirationDate(expirationDate).CertStore(certStore).Type_(type_).Name(name).Execute()

Update a given secret-group certificate



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/secretgroup_certificate"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    secretGroupId := "secretGroupId_example" // string | The secret group id
    secretId := "secretId_example" // string | The truststore id
    allowExpiredCert := true // bool | With 'true' to allow uploading expired certificates
    expirationDate := "expirationDate_example" // string | Date on which this secret should expire. If not set, by default, it will be set to notAfter date of this certificate.  Once the secret expires, a grant can not be requested for it.  (optional)
    certStore := os.NewFile(1234, "some_file") // *os.File | The file containing the certificate in PEM format (optional)
    type_ := "type__example" // string | Type of certificate supported (optional)
    name := "name_example" // string | The name of the certificate (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.PutSecretGroupCertificate(context.Background(), orgId, envId, secretGroupId, secretId).AllowExpiredCert(allowExpiredCert).ExpirationDate(expirationDate).CertStore(certStore).Type_(type_).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PutSecretGroupCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutSecretGroupCertificate`: PutSecretGroupCertificate200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PutSecretGroupCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment id | 
**secretGroupId** | **string** | The secret group id | 
**secretId** | **string** | The truststore id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSecretGroupCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **allowExpiredCert** | **bool** | With &#39;true&#39; to allow uploading expired certificates | 
 **expirationDate** | **string** | Date on which this secret should expire. If not set, by default, it will be set to notAfter date of this certificate.  Once the secret expires, a grant can not be requested for it.  | 
 **certStore** | ***os.File** | The file containing the certificate in PEM format | 
 **type_** | **string** | Type of certificate supported | 
 **name** | **string** | The name of the certificate | 

### Return type

[**PutSecretGroupCertificate200Response**](PutSecretGroupCertificate200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

