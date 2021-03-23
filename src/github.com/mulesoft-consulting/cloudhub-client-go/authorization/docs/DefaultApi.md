# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com/accounts/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Oauth2TokenPost**](DefaultApi.md#Oauth2TokenPost) | **Post** /oauth2/token | Returns access token



## Oauth2TokenPost

> InlineResponse200 Oauth2TokenPost(ctx).Credentials(credentials).Execute()

Returns access token



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
    credentials := *openapiclient.NewCredentials() // Credentials | Request body containing credentials

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.Oauth2TokenPost(context.Background()).Credentials(credentials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Oauth2TokenPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Oauth2TokenPost`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Oauth2TokenPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOauth2TokenPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **credentials** | [**Credentials**](Credentials.md) | Request body containing credentials | 

### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

