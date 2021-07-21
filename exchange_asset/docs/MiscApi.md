# \MiscApi

All URIs are relative to *https://anypoint.mulesoft.com/exchange/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Createanewasset**](MiscApi.md#Createanewasset) | **Post** /v1/assets | Createanewasset
[**Deleteanasset**](MiscApi.md#Deleteanasset) | **Delete** /v2/assets/{cloudhub-org}/{cloudhub-ex-asset}/{cloudhub-ex-version} | Deleteanasset
[**GetAsset**](MiscApi.md#GetAsset) | **Get** /v2/assets | GetAsset
[**GetonespecificAsset**](MiscApi.md#GetonespecificAsset) | **Get** /v2/assets/{cloudhub-org}/{cloudhub-ex-asset}/asset | GetonespecificAsset
[**Modifyanasset**](MiscApi.md#Modifyanasset) | **Patch** /v2/assets/{cloudhub-org}/{cloudhub-ex-asset} | Modifyanasset



## Createanewasset

> Createanewasset(ctx).OrganizationId(organizationId).AssetId(assetId).Version(version).Name(name).Classifier(classifier).ApiVersion(apiVersion).Main(main).GroupId(groupId).Asset(asset).Authorization(authorization).Execute()

Createanewasset



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
    organizationId := "organizationId_example" // string | 
    assetId := "assetId_example" // string | 
    version := "version_example" // string | 
    name := "name_example" // string | 
    classifier := "classifier_example" // string | 
    apiVersion := "apiVersion_example" // string | 
    main := "main_example" // string | 
    groupId := "groupId_example" // string | 
    asset := "asset_example" // string | 
    authorization := "authorization_example" // string |  (optional) (default to "Bearer {token}")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MiscApi.Createanewasset(context.Background()).OrganizationId(organizationId).AssetId(assetId).Version(version).Name(name).Classifier(classifier).ApiVersion(apiVersion).Main(main).GroupId(groupId).Asset(asset).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiscApi.Createanewasset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateanewassetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **string** |  | 
 **assetId** | **string** |  | 
 **version** | **string** |  | 
 **name** | **string** |  | 
 **classifier** | **string** |  | 
 **apiVersion** | **string** |  | 
 **main** | **string** |  | 
 **groupId** | **string** |  | 
 **asset** | **string** |  | 
 **authorization** | **string** |  | [default to &quot;Bearer {token}&quot;]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Deleteanasset

> Deleteanasset(ctx, cloudhubOrg, cloudhubExAsset, cloudhubExVersion).XDeleteType(xDeleteType).Authorization(authorization).Execute()

Deleteanasset



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
    xDeleteType := "xDeleteType_example" // string | 
    cloudhubOrg := "cloudhubOrg_example" // string | 
    cloudhubExAsset := "cloudhubExAsset_example" // string | 
    cloudhubExVersion := "cloudhubExVersion_example" // string | 
    authorization := "authorization_example" // string |  (optional) (default to "Bearer {token}")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MiscApi.Deleteanasset(context.Background(), cloudhubOrg, cloudhubExAsset, cloudhubExVersion).XDeleteType(xDeleteType).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiscApi.Deleteanasset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudhubOrg** | **string** |  | 
**cloudhubExAsset** | **string** |  | 
**cloudhubExVersion** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteanassetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xDeleteType** | **string** |  | 



 **authorization** | **string** |  | [default to &quot;Bearer {token}&quot;]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAsset

> GetAsset(ctx).Search(search).Types(types).Domain(domain).MasterOrganizationId(masterOrganizationId).Offset(offset).Limit(limit).SharedWithMe(sharedWithMe).IncludeSnapshots(includeSnapshots).Authorization(authorization).Execute()

GetAsset



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
    search := "search_example" // string | 
    types := "types_example" // string | 
    domain := "domain_example" // string | 
    masterOrganizationId := "masterOrganizationId_example" // string | 
    offset := int32(56) // int32 | 
    limit := int32(56) // int32 | 
    sharedWithMe := "sharedWithMe_example" // string | 
    includeSnapshots := true // bool | 
    authorization := "authorization_example" // string |  (optional) (default to "Bearer {token}")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MiscApi.GetAsset(context.Background()).Search(search).Types(types).Domain(domain).MasterOrganizationId(masterOrganizationId).Offset(offset).Limit(limit).SharedWithMe(sharedWithMe).IncludeSnapshots(includeSnapshots).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiscApi.GetAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** |  | 
 **types** | **string** |  | 
 **domain** | **string** |  | 
 **masterOrganizationId** | **string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **sharedWithMe** | **string** |  | 
 **includeSnapshots** | **bool** |  | 
 **authorization** | **string** |  | [default to &quot;Bearer {token}&quot;]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetonespecificAsset

> GetonespecificAsset(ctx, cloudhubOrg, cloudhubExAsset).Authorization(authorization).Execute()

GetonespecificAsset



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
    cloudhubOrg := "cloudhubOrg_example" // string | 
    cloudhubExAsset := "cloudhubExAsset_example" // string | 
    authorization := "authorization_example" // string |  (optional) (default to "Bearer {token}")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MiscApi.GetonespecificAsset(context.Background(), cloudhubOrg, cloudhubExAsset).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiscApi.GetonespecificAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudhubOrg** | **string** |  | 
**cloudhubExAsset** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetonespecificAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** |  | [default to &quot;Bearer {token}&quot;]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Modifyanasset

> Modifyanasset(ctx, cloudhubOrg, cloudhubExAsset).ModifyanassetRequest(modifyanassetRequest).Authorization(authorization).Execute()

Modifyanasset



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
    cloudhubOrg := "cloudhubOrg_example" // string | 
    cloudhubExAsset := "cloudhubExAsset_example" // string | 
    modifyanassetRequest := *openapiclient.NewModifyanassetRequest("Description_example") // ModifyanassetRequest | 
    authorization := "authorization_example" // string |  (optional) (default to "Bearer {token}")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MiscApi.Modifyanasset(context.Background(), cloudhubOrg, cloudhubExAsset).ModifyanassetRequest(modifyanassetRequest).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiscApi.Modifyanasset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudhubOrg** | **string** |  | 
**cloudhubExAsset** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyanassetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **modifyanassetRequest** | [**ModifyanassetRequest**](ModifyanassetRequest.md) |  | 
 **authorization** | **string** |  | [default to &quot;Bearer {token}&quot;]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

