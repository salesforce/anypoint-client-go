# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com/secrets-manager/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSecretGroupTruststoreDetails**](DefaultApi.md#GetSecretGroupTruststoreDetails) | **Get** /organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId}/truststores/{secretId} | Retrieve truststore details
[**GetSecretGroupTruststores**](DefaultApi.md#GetSecretGroupTruststores) | **Get** /organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId}/truststores | Retrieves a secret-groups&#39; collection of truststores.
[**PatchSecretGroupTruststore**](DefaultApi.md#PatchSecretGroupTruststore) | **Patch** /organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId}/truststores/{secretId} | Update a given secret-group truststore
[**PostSecretGroupTruststore**](DefaultApi.md#PostSecretGroupTruststore) | **Post** /organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId}/truststores | Create a secret-groups&#39; truststore.
[**PutSecretGroupTruststore**](DefaultApi.md#PutSecretGroupTruststore) | **Put** /organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId}/truststores/{secretId} | Update a given secret-group truststore



## GetSecretGroupTruststoreDetails

> Truststore GetSecretGroupTruststoreDetails(ctx, orgId, envId, secretGroupId, secretId).Execute()

Retrieve truststore details



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/secretgroup_truststore"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    secretGroupId := "secretGroupId_example" // string | The secret group id
    secretId := "secretId_example" // string | The truststore id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetSecretGroupTruststoreDetails(context.Background(), orgId, envId, secretGroupId, secretId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSecretGroupTruststoreDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecretGroupTruststoreDetails`: Truststore
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSecretGroupTruststoreDetails`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetSecretGroupTruststoreDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**Truststore**](Truststore.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecretGroupTruststores

> []TruststoreSummary GetSecretGroupTruststores(ctx, orgId, envId, secretGroupId).Type_(type_).Execute()

Retrieves a secret-groups' collection of truststores.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/secretgroup_truststore"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    secretGroupId := "secretGroupId_example" // string | The secret group id
    type_ := "type__example" // string | Filter the elements on the response to be of a specific type from {PEM, JKS, JCEKS, PKCS12} (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetSecretGroupTruststores(context.Background(), orgId, envId, secretGroupId).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSecretGroupTruststores``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecretGroupTruststores`: []TruststoreSummary
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSecretGroupTruststores`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetSecretGroupTruststoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **type_** | **string** | Filter the elements on the response to be of a specific type from {PEM, JKS, JCEKS, PKCS12} | 

### Return type

[**[]TruststoreSummary**](TruststoreSummary.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSecretGroupTruststore

> PutSecretGroupTruststore200Response PatchSecretGroupTruststore(ctx, orgId, envId, secretGroupId, secretId).Body(body).Execute()

Update a given secret-group truststore



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/secretgroup_truststore"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    secretGroupId := "secretGroupId_example" // string | The secret group id
    secretId := "secretId_example" // string | The truststore id
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.PatchSecretGroupTruststore(context.Background(), orgId, envId, secretGroupId, secretId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PatchSecretGroupTruststore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchSecretGroupTruststore`: PutSecretGroupTruststore200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PatchSecretGroupTruststore`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiPatchSecretGroupTruststoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | **map[string]interface{}** |  | 

### Return type

[**PutSecretGroupTruststore200Response**](PutSecretGroupTruststore200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json:
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSecretGroupTruststore

> PostSecretGroupTruststore201Response PostSecretGroupTruststore(ctx, orgId, envId, secretGroupId).AllowExpiredCert(allowExpiredCert).ExpirationDate(expirationDate).Name(name).Type_(type_).TrustStore(trustStore).Algorithm(algorithm).StorePassphrase(storePassphrase).Execute()

Create a secret-groups' truststore.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/secretgroup_truststore"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    secretGroupId := "secretGroupId_example" // string | The secret group id
    allowExpiredCert := true // bool | With 'true' to allow uploading expired certificates
    expirationDate := "expirationDate_example" // string | Date on which this secret should expire. If not set, by default, it will be set to notAfter date of the public certificate from this keystore. Once the secret expires, a grant can not be requested for it.  (optional)
    name := "name_example" // string | The name of the truststore instance (optional)
    type_ := "type__example" // string |  (optional)
    trustStore := os.NewFile(1234, "some_file") // *os.File | File containing one or more trusted certificate entries (optional)
    algorithm := "algorithm_example" // string | The algorithm used to initialize TrustManagerFactory (optional)
    storePassphrase := "storePassphrase_example" // string | The passphrase with which the trustStore file is protected (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.PostSecretGroupTruststore(context.Background(), orgId, envId, secretGroupId).AllowExpiredCert(allowExpiredCert).ExpirationDate(expirationDate).Name(name).Type_(type_).TrustStore(trustStore).Algorithm(algorithm).StorePassphrase(storePassphrase).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PostSecretGroupTruststore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostSecretGroupTruststore`: PostSecretGroupTruststore201Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PostSecretGroupTruststore`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiPostSecretGroupTruststoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **allowExpiredCert** | **bool** | With &#39;true&#39; to allow uploading expired certificates | 
 **expirationDate** | **string** | Date on which this secret should expire. If not set, by default, it will be set to notAfter date of the public certificate from this keystore. Once the secret expires, a grant can not be requested for it.  | 
 **name** | **string** | The name of the truststore instance | 
 **type_** | **string** |  | 
 **trustStore** | ***os.File** | File containing one or more trusted certificate entries | 
 **algorithm** | **string** | The algorithm used to initialize TrustManagerFactory | 
 **storePassphrase** | **string** | The passphrase with which the trustStore file is protected | 

### Return type

[**PostSecretGroupTruststore201Response**](PostSecretGroupTruststore201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSecretGroupTruststore

> PutSecretGroupTruststore200Response PutSecretGroupTruststore(ctx, orgId, envId, secretGroupId, secretId).AllowExpiredCert(allowExpiredCert).ExpirationDate(expirationDate).Name(name).Type_(type_).TrustStore(trustStore).Algorithm(algorithm).StorePassphrase(storePassphrase).Execute()

Update a given secret-group truststore



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/secretgroup_truststore"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    secretGroupId := "secretGroupId_example" // string | The secret group id
    secretId := "secretId_example" // string | The truststore id
    allowExpiredCert := true // bool | With 'true' to allow uploading expired certificates
    expirationDate := "expirationDate_example" // string | Date on which this secret should expire. If not set, by default, it will be set to notAfter date of the public certificate from this keystore. Once the secret expires, a grant can not be requested for it.  (optional)
    name := "name_example" // string | The name of the truststore instance (optional)
    type_ := "type__example" // string |  (optional)
    trustStore := os.NewFile(1234, "some_file") // *os.File | File containing one or more trusted certificate entries (optional)
    algorithm := "algorithm_example" // string | The algorithm used to initialize TrustManagerFactory (optional)
    storePassphrase := "storePassphrase_example" // string | The passphrase with which the trustStore file is protected (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.PutSecretGroupTruststore(context.Background(), orgId, envId, secretGroupId, secretId).AllowExpiredCert(allowExpiredCert).ExpirationDate(expirationDate).Name(name).Type_(type_).TrustStore(trustStore).Algorithm(algorithm).StorePassphrase(storePassphrase).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PutSecretGroupTruststore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutSecretGroupTruststore`: PutSecretGroupTruststore200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PutSecretGroupTruststore`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiPutSecretGroupTruststoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **allowExpiredCert** | **bool** | With &#39;true&#39; to allow uploading expired certificates | 
 **expirationDate** | **string** | Date on which this secret should expire. If not set, by default, it will be set to notAfter date of the public certificate from this keystore. Once the secret expires, a grant can not be requested for it.  | 
 **name** | **string** | The name of the truststore instance | 
 **type_** | **string** |  | 
 **trustStore** | ***os.File** | File containing one or more trusted certificate entries | 
 **algorithm** | **string** | The algorithm used to initialize TrustManagerFactory | 
 **storePassphrase** | **string** | The passphrase with which the trustStore file is protected | 

### Return type

[**PutSecretGroupTruststore200Response**](PutSecretGroupTruststore200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

