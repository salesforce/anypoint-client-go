# \DefaultAPI

All URIs are relative to *https://anypoint.mulesoft.com/exchange/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssetsOrgIdAssetIdAssetGet**](DefaultAPI.md#AssetsOrgIdAssetIdAssetGet) | **Get** /assets/{orgId}/{assetId}/asset | Get one specific Asset
[**AssetsOrgIdAssetIdPatch**](DefaultAPI.md#AssetsOrgIdAssetIdPatch) | **Patch** /assets/{orgId}/{assetId} | update Asset name and description
[**AssetsOrgIdAssetIdVersionAssetGet**](DefaultAPI.md#AssetsOrgIdAssetIdVersionAssetGet) | **Get** /assets/{orgId}/{assetId}/{version}/asset | Get Asset by Version
[**AssetsOrgIdAssetIdVersionDelete**](DefaultAPI.md#AssetsOrgIdAssetIdVersionDelete) | **Delete** /assets/{orgId}/{assetId}/{version} | Delete an asset
[**AssetsPost**](DefaultAPI.md#AssetsPost) | **Post** /assets | Create a new asset
[**AssetsSearchGet**](DefaultAPI.md#AssetsSearchGet) | **Get** /assets/search | Search for assets
[**PostLLMAsset**](DefaultAPI.md#PostLLMAsset) | **Post** /organizations/{orgId}/assets/{groupId}/{assetId}/{version} | Publish an LLM-typed Exchange asset



## AssetsOrgIdAssetIdAssetGet

> Asset AssetsOrgIdAssetIdAssetGet(ctx, orgId, assetId).Execute()

Get one specific Asset



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/exchange_assets"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	assetId := "assetId_example" // string | The ID of the asset

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AssetsOrgIdAssetIdAssetGet(context.Background(), orgId, assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AssetsOrgIdAssetIdAssetGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsOrgIdAssetIdAssetGet`: Asset
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AssetsOrgIdAssetIdAssetGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**assetId** | **string** | The ID of the asset | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsOrgIdAssetIdAssetGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Asset**](Asset.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsOrgIdAssetIdPatch

> AssetsOrgIdAssetIdPatch(ctx, orgId, assetId).PatchAssetNameAndDescr(patchAssetNameAndDescr).Execute()

update Asset name and description



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/exchange_assets"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	assetId := "assetId_example" // string | The ID of the asset
	patchAssetNameAndDescr := *openapiclient.NewPatchAssetNameAndDescr() // PatchAssetNameAndDescr |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.AssetsOrgIdAssetIdPatch(context.Background(), orgId, assetId).PatchAssetNameAndDescr(patchAssetNameAndDescr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AssetsOrgIdAssetIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**assetId** | **string** | The ID of the asset | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsOrgIdAssetIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchAssetNameAndDescr** | [**PatchAssetNameAndDescr**](PatchAssetNameAndDescr.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsOrgIdAssetIdVersionAssetGet

> Asset AssetsOrgIdAssetIdVersionAssetGet(ctx, orgId, assetId, version).Execute()

Get Asset by Version



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/exchange_assets"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	assetId := "assetId_example" // string | The ID of the asset
	version := "version_example" // string | The version of the asset

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AssetsOrgIdAssetIdVersionAssetGet(context.Background(), orgId, assetId, version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AssetsOrgIdAssetIdVersionAssetGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsOrgIdAssetIdVersionAssetGet`: Asset
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AssetsOrgIdAssetIdVersionAssetGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**assetId** | **string** | The ID of the asset | 
**version** | **string** | The version of the asset | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsOrgIdAssetIdVersionAssetGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Asset**](Asset.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsOrgIdAssetIdVersionDelete

> AssetsOrgIdAssetIdVersionDelete(ctx, orgId, assetId, version).XDeleteType(xDeleteType).Execute()

Delete an asset



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/exchange_assets"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	assetId := "assetId_example" // string | The ID of the asset
	version := "version_example" // string | The version of the asset
	xDeleteType := "xDeleteType_example" // string | It could be 'soft-delete' or 'hard-delete', that mean the asset will be logical deleted or physical deleted It's supposed to if it is not specified, the type will be 'soft-delete' (optional) (default to "soft-delete")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.AssetsOrgIdAssetIdVersionDelete(context.Background(), orgId, assetId, version).XDeleteType(xDeleteType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AssetsOrgIdAssetIdVersionDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**assetId** | **string** | The ID of the asset | 
**version** | **string** | The version of the asset | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssetsOrgIdAssetIdVersionDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xDeleteType** | **string** | It could be &#39;soft-delete&#39; or &#39;hard-delete&#39;, that mean the asset will be logical deleted or physical deleted It&#39;s supposed to if it is not specified, the type will be &#39;soft-delete&#39; | [default to &quot;soft-delete&quot;]

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsPost

> PostAssetResponse AssetsPost(ctx).XStrictPackage(xStrictPackage).OrganizationId(organizationId).AssetId(assetId).Version(version).Name(name).Classifier(classifier).GroupId(groupId).XAllowedApiSpecFormats(xAllowedApiSpecFormats).ApiVersion(apiVersion).Main(main).Asset(asset).Dependencies(dependencies).OriginalFormatVersion(originalFormatVersion).Metadata(metadata).Tags(tags).AssetLink(assetLink).Execute()

Create a new asset



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/exchange_assets"
)

func main() {
	xStrictPackage := true // bool | Indicates if file is immutable. (default to false)
	organizationId := "organizationId_example" // string | The id of the organization the asset will belong to
	assetId := "assetId_example" // string | The id of the asset
	version := "version_example" // string | The version of the asset being created (must follow Semver syntax)
	name := "name_example" // string | The visible name of the asset
	classifier := "classifier_example" // string | The type of the asset to be created
	groupId := "groupId_example" // string | The id of the business group the asset will belong to
	xAllowedApiSpecFormats := "xAllowedApiSpecFormats_example" // string | Specify API Spec formats that assets are allowed to use (optional)
	apiVersion := "apiVersion_example" // string | The product version of API assets. Required for \\\"raml\\\", \\\"oas\\\", \\\"wsdl\\\" and \\\"http\\\" assets (optional)
	main := "main_example" // string | The main file of the asset. Required for \\\"raml\\\", \\\"raml-fragment\\\", \\\"oas\\\" and \\\"wsdl\\\". (optional)
	asset := os.NewFile(1234, "some_file") // *os.File | The asset file. Required for \\\"raml\\\", \\\"raml-fragment\\\", \\\"oas\\\" and \\\"wsdl\\\". Maximum size of 5 MB. This field must be the last field of the multipart. (optional)
	dependencies := "dependencies_example" // string | Required for \\\"api-group\\\" classifier only, They are APIs included in it, as a JSON array of objects. Because the field must be of String type, the stringified value of the JSON array must be passed as parameter. (optional)
	originalFormatVersion := "originalFormatVersion_example" // string | The version of the format of the api specification. ie ‘2.0’ for OAS 2.0 (optional)
	metadata := "metadata_example" // string | A design center object describing asset projectId, branchId and commitId. Because the field must be of String type, the stringified value of the JSON object must be passed as parameter. (optional)
	tags := "tags_example" // string | An array of strings to be saved as asset's tags. Because the field must be of String type, the stringified value of the JSON array must be passed as parameter. (optional)
	assetLink := "assetLink_example" // string | The link of the asset. Required for \\\"wsdl\\\" or \\\"http\\\" assets (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AssetsPost(context.Background()).XStrictPackage(xStrictPackage).OrganizationId(organizationId).AssetId(assetId).Version(version).Name(name).Classifier(classifier).GroupId(groupId).XAllowedApiSpecFormats(xAllowedApiSpecFormats).ApiVersion(apiVersion).Main(main).Asset(asset).Dependencies(dependencies).OriginalFormatVersion(originalFormatVersion).Metadata(metadata).Tags(tags).AssetLink(assetLink).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AssetsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsPost`: PostAssetResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AssetsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssetsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xStrictPackage** | **bool** | Indicates if file is immutable. | [default to false]
 **organizationId** | **string** | The id of the organization the asset will belong to | 
 **assetId** | **string** | The id of the asset | 
 **version** | **string** | The version of the asset being created (must follow Semver syntax) | 
 **name** | **string** | The visible name of the asset | 
 **classifier** | **string** | The type of the asset to be created | 
 **groupId** | **string** | The id of the business group the asset will belong to | 
 **xAllowedApiSpecFormats** | **string** | Specify API Spec formats that assets are allowed to use | 
 **apiVersion** | **string** | The product version of API assets. Required for \\\&quot;raml\\\&quot;, \\\&quot;oas\\\&quot;, \\\&quot;wsdl\\\&quot; and \\\&quot;http\\\&quot; assets | 
 **main** | **string** | The main file of the asset. Required for \\\&quot;raml\\\&quot;, \\\&quot;raml-fragment\\\&quot;, \\\&quot;oas\\\&quot; and \\\&quot;wsdl\\\&quot;. | 
 **asset** | ***os.File** | The asset file. Required for \\\&quot;raml\\\&quot;, \\\&quot;raml-fragment\\\&quot;, \\\&quot;oas\\\&quot; and \\\&quot;wsdl\\\&quot;. Maximum size of 5 MB. This field must be the last field of the multipart. | 
 **dependencies** | **string** | Required for \\\&quot;api-group\\\&quot; classifier only, They are APIs included in it, as a JSON array of objects. Because the field must be of String type, the stringified value of the JSON array must be passed as parameter. | 
 **originalFormatVersion** | **string** | The version of the format of the api specification. ie ‘2.0’ for OAS 2.0 | 
 **metadata** | **string** | A design center object describing asset projectId, branchId and commitId. Because the field must be of String type, the stringified value of the JSON object must be passed as parameter. | 
 **tags** | **string** | An array of strings to be saved as asset&#39;s tags. Because the field must be of String type, the stringified value of the JSON array must be passed as parameter. | 
 **assetLink** | **string** | The link of the asset. Required for \\\&quot;wsdl\\\&quot; or \\\&quot;http\\\&quot; assets | 

### Return type

[**PostAssetResponse**](PostAssetResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetsSearchGet

> []AssetSearchResultItem AssetsSearchGet(ctx).Types(types).Search(search).Domain(domain).MasterOrganizationId(masterOrganizationId).OrganizationId(organizationId).Offset(offset).Limit(limit).Sort(sort).Ascending(ascending).SharedWithMe(sharedWithMe).IncludeSnapshots(includeSnapshots).Execute()

Search for assets



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/exchange_assets"
)

func main() {
	types := "types_example" // string | Filter results that matches the input with the asset type
	search := "search_example" // string | Filter results that partially match the input with the asset name (optional)
	domain := "domain_example" // string | Filter results by organization using its domain (optional)
	masterOrganizationId := "masterOrganizationId_example" // string | Filter results by master organization id. (optional)
	organizationId := "organizationId_example" // string | Filter results by organizations. For more than one organization, & organizationId=1& organizationId=2, etc... (optional)
	offset := int32(56) // int32 | The offset specifies the offset of the first row to return (optional)
	limit := int32(56) // int32 | Amount of objects retrieved in the response (optional)
	sort := "sort_example" // string | Property to sort by (optional)
	ascending := "ascending_example" // string | Order for sorting (optional)
	sharedWithMe := true // bool | Retrieve only the assets that has been shared with you (optional)
	includeSnapshots := true // bool | Include snapshots in the results (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AssetsSearchGet(context.Background()).Types(types).Search(search).Domain(domain).MasterOrganizationId(masterOrganizationId).OrganizationId(organizationId).Offset(offset).Limit(limit).Sort(sort).Ascending(ascending).SharedWithMe(sharedWithMe).IncludeSnapshots(includeSnapshots).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AssetsSearchGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssetsSearchGet`: []AssetSearchResultItem
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AssetsSearchGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssetsSearchGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **types** | **string** | Filter results that matches the input with the asset type | 
 **search** | **string** | Filter results that partially match the input with the asset name | 
 **domain** | **string** | Filter results by organization using its domain | 
 **masterOrganizationId** | **string** | Filter results by master organization id. | 
 **organizationId** | **string** | Filter results by organizations. For more than one organization, &amp; organizationId&#x3D;1&amp; organizationId&#x3D;2, etc... | 
 **offset** | **int32** | The offset specifies the offset of the first row to return | 
 **limit** | **int32** | Amount of objects retrieved in the response | 
 **sort** | **string** | Property to sort by | 
 **ascending** | **string** | Order for sorting | 
 **sharedWithMe** | **bool** | Retrieve only the assets that has been shared with you | 
 **includeSnapshots** | **bool** | Include snapshots in the results | 

### Return type

[**[]AssetSearchResultItem**](AssetSearchResultItem.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLLMAsset

> PostLLMAssetResponse PostLLMAsset(ctx, orgId, groupId, assetId, version).Name(name).Type_(type_).Status(status).XStrictPackage(xStrictPackage).PropertiesPlatform(propertiesPlatform).Tags(tags).Execute()

Publish an LLM-typed Exchange asset



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/exchange_assets"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	groupId := "groupId_example" // string | The business group id (often equals orgId for root-level assets)
	assetId := "assetId_example" // string | The asset slug
	version := "version_example" // string | The asset version (must follow Semver, e.g. 1.0.0)
	name := "name_example" // string | The visible name of the asset.
	type_ := "type__example" // string | Always \\\"llm\\\" for this resource.
	status := "status_example" // string | Publication status. Defaults to \\\"published\\\". (default to "published")
	xStrictPackage := true // bool | Indicates if the asset package is immutable. (optional) (default to false)
	propertiesPlatform := "propertiesPlatform_example" // string | The LLM vendor identifier. Free string — Anypoint does not enforce an enum. Known values include \\\"openai\\\", \\\"bedrock\\\", \\\"anthropic\\\", \\\"azure-openai\\\", \\\"gemini\\\", \\\"other\\\". Defaults server-side to \\\"other\\\" if omitted. (optional)
	tags := "tags_example" // string | Stringified JSON array of asset tags. Optional. The field must be sent as a String; pass the stringified value. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PostLLMAsset(context.Background(), orgId, groupId, assetId, version).Name(name).Type_(type_).Status(status).XStrictPackage(xStrictPackage).PropertiesPlatform(propertiesPlatform).Tags(tags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PostLLMAsset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLLMAsset`: PostLLMAssetResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PostLLMAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**groupId** | **string** | The business group id (often equals orgId for root-level assets) | 
**assetId** | **string** | The asset slug | 
**version** | **string** | The asset version (must follow Semver, e.g. 1.0.0) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostLLMAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **name** | **string** | The visible name of the asset. | 
 **type_** | **string** | Always \\\&quot;llm\\\&quot; for this resource. | 
 **status** | **string** | Publication status. Defaults to \\\&quot;published\\\&quot;. | [default to &quot;published&quot;]
 **xStrictPackage** | **bool** | Indicates if the asset package is immutable. | [default to false]
 **propertiesPlatform** | **string** | The LLM vendor identifier. Free string — Anypoint does not enforce an enum. Known values include \\\&quot;openai\\\&quot;, \\\&quot;bedrock\\\&quot;, \\\&quot;anthropic\\\&quot;, \\\&quot;azure-openai\\\&quot;, \\\&quot;gemini\\\&quot;, \\\&quot;other\\\&quot;. Defaults server-side to \\\&quot;other\\\&quot; if omitted. | 
 **tags** | **string** | Stringified JSON array of asset tags. Optional. The field must be sent as a String; pass the stringified value. | 

### Return type

[**PostLLMAssetResponse**](PostLLMAssetResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

