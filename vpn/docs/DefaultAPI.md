# \DefaultAPI

All URIs are relative to *https://anypoint.mulesoft.com/cloudhub/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsOrgIdVpcsVpcIdIpsecGet**](DefaultAPI.md#OrganizationsOrgIdVpcsVpcIdIpsecGet) | **Get** /organizations/{orgId}/vpcs/{vpcId}/ipsec | Returns a list of vpns.
[**OrganizationsOrgIdVpcsVpcIdIpsecPost**](DefaultAPI.md#OrganizationsOrgIdVpcsVpcIdIpsecPost) | **Post** /organizations/{orgId}/vpcs/{vpcId}/ipsec | Creates a VPN.
[**OrganizationsOrgIdVpcsVpcIdIpsecVpnIdDelete**](DefaultAPI.md#OrganizationsOrgIdVpcsVpcIdIpsecVpnIdDelete) | **Delete** /organizations/{orgId}/vpcs/{vpcId}/ipsec/{vpnId} | Delete a VPN connection
[**OrganizationsOrgIdVpcsVpcIdIpsecVpnIdGet**](DefaultAPI.md#OrganizationsOrgIdVpcsVpcIdIpsecVpnIdGet) | **Get** /organizations/{orgId}/vpcs/{vpcId}/ipsec/{vpnId} | Returns a a specific vpn



## OrganizationsOrgIdVpcsVpcIdIpsecGet

> OrganizationsOrgIdVpcsVpcIdIpsecGet200Response OrganizationsOrgIdVpcsVpcIdIpsecGet(ctx, orgId, vpcId).Execute()

Returns a list of vpns.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/vpn"
)

func main() {
	orgId := "orgId_example" // string | The organization Id
	vpcId := "vpcId_example" // string | The vpc Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.OrganizationsOrgIdVpcsVpcIdIpsecGet(context.Background(), orgId, vpcId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.OrganizationsOrgIdVpcsVpcIdIpsecGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsOrgIdVpcsVpcIdIpsecGet`: OrganizationsOrgIdVpcsVpcIdIpsecGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.OrganizationsOrgIdVpcsVpcIdIpsecGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**vpcId** | **string** | The vpc Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdVpcsVpcIdIpsecGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationsOrgIdVpcsVpcIdIpsecGet200Response**](OrganizationsOrgIdVpcsVpcIdIpsecGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrgIdVpcsVpcIdIpsecPost

> VpnPost OrganizationsOrgIdVpcsVpcIdIpsecPost(ctx, orgId, vpcId).VpnPostReqBody(vpnPostReqBody).Execute()

Creates a VPN.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/vpn"
)

func main() {
	orgId := "orgId_example" // string | The organization Id
	vpcId := "vpcId_example" // string | The vpc Id
	vpnPostReqBody := *openapiclient.NewVpnPostReqBody() // VpnPostReqBody | VPN request body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.OrganizationsOrgIdVpcsVpcIdIpsecPost(context.Background(), orgId, vpcId).VpnPostReqBody(vpnPostReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.OrganizationsOrgIdVpcsVpcIdIpsecPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsOrgIdVpcsVpcIdIpsecPost`: VpnPost
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.OrganizationsOrgIdVpcsVpcIdIpsecPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**vpcId** | **string** | The vpc Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdVpcsVpcIdIpsecPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **vpnPostReqBody** | [**VpnPostReqBody**](VpnPostReqBody.md) | VPN request body | 

### Return type

[**VpnPost**](VpnPost.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrgIdVpcsVpcIdIpsecVpnIdDelete

> OrganizationsOrgIdVpcsVpcIdIpsecVpnIdDelete(ctx, orgId, vpcId, vpnId).Execute()

Delete a VPN connection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/vpn"
)

func main() {
	orgId := "orgId_example" // string | The organization Id
	vpcId := "vpcId_example" // string | The vpc Id
	vpnId := "vpnId_example" // string | The vpn Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.OrganizationsOrgIdVpcsVpcIdIpsecVpnIdDelete(context.Background(), orgId, vpcId, vpnId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.OrganizationsOrgIdVpcsVpcIdIpsecVpnIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**vpcId** | **string** | The vpc Id | 
**vpnId** | **string** | The vpn Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdVpcsVpcIdIpsecVpnIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrgIdVpcsVpcIdIpsecVpnIdGet

> VpnGet OrganizationsOrgIdVpcsVpcIdIpsecVpnIdGet(ctx, orgId, vpcId, vpnId).Execute()

Returns a a specific vpn



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/vpn"
)

func main() {
	orgId := "orgId_example" // string | The organization Id
	vpcId := "vpcId_example" // string | The vpc Id
	vpnId := "vpnId_example" // string | The vpn Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.OrganizationsOrgIdVpcsVpcIdIpsecVpnIdGet(context.Background(), orgId, vpcId, vpnId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.OrganizationsOrgIdVpcsVpcIdIpsecVpnIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsOrgIdVpcsVpcIdIpsecVpnIdGet`: VpnGet
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.OrganizationsOrgIdVpcsVpcIdIpsecVpnIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization Id | 
**vpcId** | **string** | The vpc Id | 
**vpnId** | **string** | The vpn Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdVpcsVpcIdIpsecVpnIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**VpnGet**](VpnGet.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

