# \DefaultAPI

All URIs are relative to *https://anypoint.mulesoft.com/accounts*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2Oauth2TokenPost**](DefaultAPI.md#ApiV2Oauth2TokenPost) | **Post** /api/v2/oauth2/token | Returns access token
[**LoginPost**](DefaultAPI.md#LoginPost) | **Post** /login | Returns access token



## ApiV2Oauth2TokenPost

> ApiV2Oauth2TokenPost200Response ApiV2Oauth2TokenPost(ctx).Credentials(credentials).Execute()

Returns access token



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/authorization"
)

func main() {
	credentials := *openapiclient.NewCredentials() // Credentials | Request body containing credentials

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiV2Oauth2TokenPost(context.Background()).Credentials(credentials).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV2Oauth2TokenPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2Oauth2TokenPost`: ApiV2Oauth2TokenPost200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiV2Oauth2TokenPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2Oauth2TokenPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **credentials** | [**Credentials**](Credentials.md) | Request body containing credentials | 

### Return type

[**ApiV2Oauth2TokenPost200Response**](ApiV2Oauth2TokenPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoginPost

> LoginPost200Response LoginPost(ctx).UserPwdCredentials(userPwdCredentials).Execute()

Returns access token



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/authorization"
)

func main() {
	userPwdCredentials := *openapiclient.NewUserPwdCredentials() // UserPwdCredentials | Request body containing credentials

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.LoginPost(context.Background()).UserPwdCredentials(userPwdCredentials).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.LoginPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LoginPost`: LoginPost200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.LoginPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoginPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userPwdCredentials** | [**UserPwdCredentials**](UserPwdCredentials.md) | Request body containing credentials | 

### Return type

[**LoginPost200Response**](LoginPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

