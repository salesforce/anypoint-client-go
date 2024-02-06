# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com/secrets-manager/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSecretGroupKeystoreDetails**](DefaultApi.md#GetSecretGroupKeystoreDetails) | **Get** /organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId}/keystores/{secretId} | Retrieve keystore details
[**GetSecretGroupKeystores**](DefaultApi.md#GetSecretGroupKeystores) | **Get** /organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId}/keystores | Retrieves a secret-groups&#39; collection of keystores.
[**PatchSecretGroupKeystore**](DefaultApi.md#PatchSecretGroupKeystore) | **Patch** /organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId}/keystores/{secretId} | Update a given secret-group keystore
[**PostSecretGroupKeystores**](DefaultApi.md#PostSecretGroupKeystores) | **Post** /organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId}/keystores | Create a secret-groups&#39; keystore.
[**PutSecretGroupKeystore**](DefaultApi.md#PutSecretGroupKeystore) | **Put** /organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId}/keystores/{secretId} | Update a given secret-group keystore



## GetSecretGroupKeystoreDetails

> Keystore GetSecretGroupKeystoreDetails(ctx, orgId, envId, secretGroupId, secretId).Execute()

Retrieve keystore details



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/secretgroup_keystore"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    secretGroupId := "secretGroupId_example" // string | The secret group id
    secretId := "secretId_example" // string | The keystore id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetSecretGroupKeystoreDetails(context.Background(), orgId, envId, secretGroupId, secretId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSecretGroupKeystoreDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecretGroupKeystoreDetails`: Keystore
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSecretGroupKeystoreDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment id | 
**secretGroupId** | **string** | The secret group id | 
**secretId** | **string** | The keystore id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretGroupKeystoreDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**Keystore**](Keystore.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecretGroupKeystores

> []KeystoreSummary GetSecretGroupKeystores(ctx, orgId, envId, secretGroupId).Type_(type_).Execute()

Retrieves a secret-groups' collection of keystores.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/secretgroup_keystore"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    secretGroupId := "secretGroupId_example" // string | The secret group id
    type_ := "type__example" // string | Filter the elements on the response to be of a specific type from {PEM, JKS, JCEKS, PKCS12} (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetSecretGroupKeystores(context.Background(), orgId, envId, secretGroupId).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSecretGroupKeystores``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecretGroupKeystores`: []KeystoreSummary
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSecretGroupKeystores`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetSecretGroupKeystoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **type_** | **string** | Filter the elements on the response to be of a specific type from {PEM, JKS, JCEKS, PKCS12} | 

### Return type

[**[]KeystoreSummary**](KeystoreSummary.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSecretGroupKeystore

> PutSecretGroupKeystore200Response PatchSecretGroupKeystore(ctx, orgId, envId, secretGroupId, secretId).Body(body).Execute()

Update a given secret-group keystore



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/secretgroup_keystore"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    secretGroupId := "secretGroupId_example" // string | The secret group id
    secretId := "secretId_example" // string | The keystore id
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.PatchSecretGroupKeystore(context.Background(), orgId, envId, secretGroupId, secretId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PatchSecretGroupKeystore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchSecretGroupKeystore`: PutSecretGroupKeystore200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PatchSecretGroupKeystore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment id | 
**secretGroupId** | **string** | The secret group id | 
**secretId** | **string** | The keystore id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSecretGroupKeystoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | **map[string]interface{}** |  | 

### Return type

[**PutSecretGroupKeystore200Response**](PutSecretGroupKeystore200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json:
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSecretGroupKeystores

> PostSecretGroupKeystores201Response PostSecretGroupKeystores(ctx, orgId, envId, secretGroupId).AllowExpiredCert(allowExpiredCert).ExpirationDate(expirationDate).Key(key).Name(name).KeyPassphrase(keyPassphrase).Certificate(certificate).Type_(type_).Capath(capath).KeyStore(keyStore).Algorithm(algorithm).StorePassphrase(storePassphrase).Alias(alias).Execute()

Create a secret-groups' keystore.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/secretgroup_keystore"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    secretGroupId := "secretGroupId_example" // string | The secret group id
    allowExpiredCert := true // bool | With 'true' to allow uploading expired certificates
    expirationDate := time.Now() // string | Date on which this secret should expire. If not set, by default, it will be set to notAfter date of the public certificate from this keystore. Once the secret expires, a grant can not be requested for it.  (optional)
    key := os.NewFile(1234, "some_file") // *os.File | The encrypted private key. Required in case of PEM type.  (optional)
    name := "name_example" // string | The name of this keystore instance.  (optional)
    keyPassphrase := "keyPassphrase_example" // string | Passphrase with which private key for a particular alias is protected  (optional)
    certificate := os.NewFile(1234, "some_file") // *os.File | The public certificate. Required in the case of PEM type.  (optional)
    type_ := "type__example" // string |  (optional)
    capath := os.NewFile(1234, "some_file") // *os.File | The concatenated chain of CA certificates, except the leaf, leading up to the root CA. Can only be set in case of PEM type.  (optional)
    keyStore := os.NewFile(1234, "some_file") // *os.File | File containing one or more certificate entries Required in case of JKS, JCEKS and PKCS12 types  (optional)
    algorithm := "algorithm_example" // string | The algorithm used to initialize KeyManagerFactory Required in case of JKS, JCEKS and PKCS12 types  (optional)
    storePassphrase := "storePassphrase_example" // string | Passphrase with which keystore is protected Required in case of JKS, JCEKS and PKCS12 types  (optional)
    alias := "alias_example" // string | The alias name of the entry that contains the certificate. Required in case of JKS, JCEKS and PKCS12 types  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.PostSecretGroupKeystores(context.Background(), orgId, envId, secretGroupId).AllowExpiredCert(allowExpiredCert).ExpirationDate(expirationDate).Key(key).Name(name).KeyPassphrase(keyPassphrase).Certificate(certificate).Type_(type_).Capath(capath).KeyStore(keyStore).Algorithm(algorithm).StorePassphrase(storePassphrase).Alias(alias).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PostSecretGroupKeystores``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostSecretGroupKeystores`: PostSecretGroupKeystores201Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PostSecretGroupKeystores`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiPostSecretGroupKeystoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **allowExpiredCert** | **bool** | With &#39;true&#39; to allow uploading expired certificates | 
 **expirationDate** | **string** | Date on which this secret should expire. If not set, by default, it will be set to notAfter date of the public certificate from this keystore. Once the secret expires, a grant can not be requested for it.  | 
 **key** | ***os.File** | The encrypted private key. Required in case of PEM type.  | 
 **name** | **string** | The name of this keystore instance.  | 
 **keyPassphrase** | **string** | Passphrase with which private key for a particular alias is protected  | 
 **certificate** | ***os.File** | The public certificate. Required in the case of PEM type.  | 
 **type_** | **string** |  | 
 **capath** | ***os.File** | The concatenated chain of CA certificates, except the leaf, leading up to the root CA. Can only be set in case of PEM type.  | 
 **keyStore** | ***os.File** | File containing one or more certificate entries Required in case of JKS, JCEKS and PKCS12 types  | 
 **algorithm** | **string** | The algorithm used to initialize KeyManagerFactory Required in case of JKS, JCEKS and PKCS12 types  | 
 **storePassphrase** | **string** | Passphrase with which keystore is protected Required in case of JKS, JCEKS and PKCS12 types  | 
 **alias** | **string** | The alias name of the entry that contains the certificate. Required in case of JKS, JCEKS and PKCS12 types  | 

### Return type

[**PostSecretGroupKeystores201Response**](PostSecretGroupKeystores201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSecretGroupKeystore

> PutSecretGroupKeystore200Response PutSecretGroupKeystore(ctx, orgId, envId, secretGroupId, secretId).AllowExpiredCert(allowExpiredCert).ExpirationDate(expirationDate).Key(key).Name(name).KeyPassphrase(keyPassphrase).Certificate(certificate).Type_(type_).Capath(capath).KeyStore(keyStore).Algorithm(algorithm).StorePassphrase(storePassphrase).Alias(alias).Execute()

Update a given secret-group keystore



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/secretgroup_keystore"
)

func main() {
    orgId := "orgId_example" // string | The organization Id
    envId := "envId_example" // string | The environment id
    secretGroupId := "secretGroupId_example" // string | The secret group id
    secretId := "secretId_example" // string | The keystore id
    allowExpiredCert := true // bool | With 'true' to allow uploading expired certificates
    expirationDate := time.Now() // string | Date on which this secret should expire. If not set, by default, it will be set to notAfter date of the public certificate from this keystore. Once the secret expires, a grant can not be requested for it.  (optional)
    key := os.NewFile(1234, "some_file") // *os.File | The encrypted private key. Required in case of PEM type.  (optional)
    name := "name_example" // string | The name of this keystore instance.  (optional)
    keyPassphrase := "keyPassphrase_example" // string | Passphrase with which private key for a particular alias is protected  (optional)
    certificate := os.NewFile(1234, "some_file") // *os.File | The public certificate. Required in the case of PEM type.  (optional)
    type_ := "type__example" // string |  (optional)
    capath := os.NewFile(1234, "some_file") // *os.File | The concatenated chain of CA certificates, except the leaf, leading up to the root CA. Can only be set in case of PEM type.  (optional)
    keyStore := os.NewFile(1234, "some_file") // *os.File | File containing one or more certificate entries Required in case of JKS, JCEKS and PKCS12 types  (optional)
    algorithm := "algorithm_example" // string | The algorithm used to initialize KeyManagerFactory Required in case of JKS, JCEKS and PKCS12 types  (optional)
    storePassphrase := "storePassphrase_example" // string | Passphrase with which keystore is protected Required in case of JKS, JCEKS and PKCS12 types  (optional)
    alias := "alias_example" // string | The alias name of the entry that contains the certificate. Required in case of JKS, JCEKS and PKCS12 types  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.PutSecretGroupKeystore(context.Background(), orgId, envId, secretGroupId, secretId).AllowExpiredCert(allowExpiredCert).ExpirationDate(expirationDate).Key(key).Name(name).KeyPassphrase(keyPassphrase).Certificate(certificate).Type_(type_).Capath(capath).KeyStore(keyStore).Algorithm(algorithm).StorePassphrase(storePassphrase).Alias(alias).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PutSecretGroupKeystore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutSecretGroupKeystore`: PutSecretGroupKeystore200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PutSecretGroupKeystore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**envId** | **string** | The environment id | 
**secretGroupId** | **string** | The secret group id | 
**secretId** | **string** | The keystore id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSecretGroupKeystoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **allowExpiredCert** | **bool** | With &#39;true&#39; to allow uploading expired certificates | 
 **expirationDate** | **string** | Date on which this secret should expire. If not set, by default, it will be set to notAfter date of the public certificate from this keystore. Once the secret expires, a grant can not be requested for it.  | 
 **key** | ***os.File** | The encrypted private key. Required in case of PEM type.  | 
 **name** | **string** | The name of this keystore instance.  | 
 **keyPassphrase** | **string** | Passphrase with which private key for a particular alias is protected  | 
 **certificate** | ***os.File** | The public certificate. Required in the case of PEM type.  | 
 **type_** | **string** |  | 
 **capath** | ***os.File** | The concatenated chain of CA certificates, except the leaf, leading up to the root CA. Can only be set in case of PEM type.  | 
 **keyStore** | ***os.File** | File containing one or more certificate entries Required in case of JKS, JCEKS and PKCS12 types  | 
 **algorithm** | **string** | The algorithm used to initialize KeyManagerFactory Required in case of JKS, JCEKS and PKCS12 types  | 
 **storePassphrase** | **string** | Passphrase with which keystore is protected Required in case of JKS, JCEKS and PKCS12 types  | 
 **alias** | **string** | The alias name of the entry that contains the certificate. Required in case of JKS, JCEKS and PKCS12 types  | 

### Return type

[**PutSecretGroupKeystore200Response**](PutSecretGroupKeystore200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

