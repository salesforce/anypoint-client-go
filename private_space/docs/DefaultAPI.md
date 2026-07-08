# \DefaultAPI

All URIs are relative to *https://anypoint.mulesoft.com/runtimefabric/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPrivateSpaceVpnConnectionMember**](DefaultAPI.md#AddPrivateSpaceVpnConnectionMember) | **Post** /organizations/{orgId}/privatespaces/{privateSpaceId}/connections/{connectionId}/vpns | 
[**CreatePrivateSpace**](DefaultAPI.md#CreatePrivateSpace) | **Post** /organizations/{orgId}/privatespaces | 
[**CreatePrivateSpaceTransitGateway**](DefaultAPI.md#CreatePrivateSpaceTransitGateway) | **Post** /organizations/{orgId}/privatespaces/{privateSpaceId}/transitgateways | 
[**CreatePrivateSpaceVpnConnection**](DefaultAPI.md#CreatePrivateSpaceVpnConnection) | **Post** /organizations/{orgId}/privatespaces/{privateSpaceId}/connections | 
[**DeletePrivateSpace**](DefaultAPI.md#DeletePrivateSpace) | **Delete** /organizations/{orgId}/privatespaces/{privateSpaceId} | 
[**DeletePrivateSpaceTransitGateway**](DefaultAPI.md#DeletePrivateSpaceTransitGateway) | **Delete** /organizations/{orgId}/privatespaces/{privateSpaceId}/transitgateways/{transitGatewayId} | 
[**DeletePrivateSpaceVpnConnection**](DefaultAPI.md#DeletePrivateSpaceVpnConnection) | **Delete** /organizations/{orgId}/privatespaces/{privateSpaceId}/connections/{connectionId} | 
[**DeletePrivateSpaceVpnConnectionMember**](DefaultAPI.md#DeletePrivateSpaceVpnConnectionMember) | **Delete** /organizations/{orgId}/privatespaces/{privateSpaceId}/connections/{connectionId}/vpns/{vpnId} | 
[**GetOrgTransitGateways**](DefaultAPI.md#GetOrgTransitGateways) | **Get** /organizations/{orgId}/transitgateways | 
[**GetPrivateSpace**](DefaultAPI.md#GetPrivateSpace) | **Get** /organizations/{orgId}/privatespaces/{privateSpaceId} | 
[**GetPrivateSpaceIamRoles**](DefaultAPI.md#GetPrivateSpaceIamRoles) | **Get** /organizations/{orgId}/privatespaces/{privateSpaceId}/iamroles | 
[**GetPrivateSpaceMulesoftAccount**](DefaultAPI.md#GetPrivateSpaceMulesoftAccount) | **Get** /organizations/{orgId}/privatespaces/{privateSpaceId}/accounts | 
[**GetPrivateSpaceRoutes**](DefaultAPI.md#GetPrivateSpaceRoutes) | **Get** /organizations/{orgId}/privatespaces/{privateSpaceId}/routes | 
[**GetPrivateSpaceSupportedVpnConfigs**](DefaultAPI.md#GetPrivateSpaceSupportedVpnConfigs) | **Get** /organizations/{orgId}/privatespaces/supportedVpnConfigs | 
[**GetPrivateSpaceTransitGateways**](DefaultAPI.md#GetPrivateSpaceTransitGateways) | **Get** /organizations/{orgId}/privatespaces/{privateSpaceId}/transitgateways | 
[**GetPrivateSpaceVpnConnection**](DefaultAPI.md#GetPrivateSpaceVpnConnection) | **Get** /organizations/{orgId}/privatespaces/{privateSpaceId}/connections/{connectionId} | 
[**GetPrivateSpaceVpnConnections**](DefaultAPI.md#GetPrivateSpaceVpnConnections) | **Get** /organizations/{orgId}/privatespaces/{privateSpaceId}/connections | 
[**GetPrivateSpaces**](DefaultAPI.md#GetPrivateSpaces) | **Get** /organizations/{orgId}/privatespaces | 
[**UpdateOrgTransitGatewayName**](DefaultAPI.md#UpdateOrgTransitGatewayName) | **Patch** /organizations/{orgId}/transitgateways/{transitGatewayId} | 
[**UpdatePrivateSpace**](DefaultAPI.md#UpdatePrivateSpace) | **Patch** /organizations/{orgId}/privatespaces/{privateSpaceId} | 
[**UpdatePrivateSpaceTransitGatewayRoutes**](DefaultAPI.md#UpdatePrivateSpaceTransitGatewayRoutes) | **Patch** /organizations/{orgId}/privatespaces/{privateSpaceId}/transitgateways/{transitGatewayId} | 
[**UpdatePrivateSpaceVpnConnection**](DefaultAPI.md#UpdatePrivateSpaceVpnConnection) | **Patch** /organizations/{orgId}/privatespaces/{privateSpaceId}/connections/{connectionId} | 
[**UpdatePrivateSpaceVpnConnectionMember**](DefaultAPI.md#UpdatePrivateSpaceVpnConnectionMember) | **Patch** /organizations/{orgId}/privatespaces/{privateSpaceId}/connections/{connectionId}/vpns/{vpnId} | 



## AddPrivateSpaceVpnConnectionMember

> PrivateSpaceVpnConnection AddPrivateSpaceVpnConnectionMember(ctx, orgId, privateSpaceId, connectionId).PrivateSpaceVpnPostBody(privateSpaceVpnPostBody).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/private_space"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	privateSpaceId := "privateSpaceId_example" // string | The ID of the private space in GUID format
	connectionId := "connectionId_example" // string | The ID of the VPN connection in GUID format
	privateSpaceVpnPostBody := *openapiclient.NewPrivateSpaceVpnPostBody("LocalAsn_example", "RemoteIpAddress_example", []openapiclient.PrivateSpaceVpnTunnelPostBody{*openapiclient.NewPrivateSpaceVpnTunnelPostBody("Psk_example", "StartupAction_example")}) // PrivateSpaceVpnPostBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AddPrivateSpaceVpnConnectionMember(context.Background(), orgId, privateSpaceId, connectionId).PrivateSpaceVpnPostBody(privateSpaceVpnPostBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AddPrivateSpaceVpnConnectionMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddPrivateSpaceVpnConnectionMember`: PrivateSpaceVpnConnection
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AddPrivateSpaceVpnConnectionMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**privateSpaceId** | **string** | The ID of the private space in GUID format | 
**connectionId** | **string** | The ID of the VPN connection in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPrivateSpaceVpnConnectionMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **privateSpaceVpnPostBody** | [**PrivateSpaceVpnPostBody**](PrivateSpaceVpnPostBody.md) |  | 

### Return type

[**PrivateSpaceVpnConnection**](PrivateSpaceVpnConnection.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePrivateSpace

> PrivateSpace CreatePrivateSpace(ctx, orgId).PrivateSpacePostBody(privateSpacePostBody).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/private_space"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	privateSpacePostBody := *openapiclient.NewPrivateSpacePostBody() // PrivateSpacePostBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreatePrivateSpace(context.Background(), orgId).PrivateSpacePostBody(privateSpacePostBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreatePrivateSpace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePrivateSpace`: PrivateSpace
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreatePrivateSpace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePrivateSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **privateSpacePostBody** | [**PrivateSpacePostBody**](PrivateSpacePostBody.md) |  | 

### Return type

[**PrivateSpace**](PrivateSpace.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePrivateSpaceTransitGateway

> TransitGateway CreatePrivateSpaceTransitGateway(ctx, orgId, privateSpaceId).TransitGatewayPostBody(transitGatewayPostBody).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/private_space"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	privateSpaceId := "privateSpaceId_example" // string | The ID of the private space in GUID format
	transitGatewayPostBody := *openapiclient.NewTransitGatewayPostBody("Name_example", "ResourceShareId_example", "ResourceShareAccount_example", []string{"Routes_example"}) // TransitGatewayPostBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreatePrivateSpaceTransitGateway(context.Background(), orgId, privateSpaceId).TransitGatewayPostBody(transitGatewayPostBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreatePrivateSpaceTransitGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePrivateSpaceTransitGateway`: TransitGateway
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreatePrivateSpaceTransitGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**privateSpaceId** | **string** | The ID of the private space in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePrivateSpaceTransitGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitGatewayPostBody** | [**TransitGatewayPostBody**](TransitGatewayPostBody.md) |  | 

### Return type

[**TransitGateway**](TransitGateway.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePrivateSpaceVpnConnection

> PrivateSpaceVpnConnection CreatePrivateSpaceVpnConnection(ctx, orgId, privateSpaceId).PrivateSpaceVpnConnectionPostBody(privateSpaceVpnConnectionPostBody).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/private_space"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	privateSpaceId := "privateSpaceId_example" // string | The ID of the private space in GUID format
	privateSpaceVpnConnectionPostBody := *openapiclient.NewPrivateSpaceVpnConnectionPostBody("Name_example", []openapiclient.PrivateSpaceVpnPostBody{*openapiclient.NewPrivateSpaceVpnPostBody("LocalAsn_example", "RemoteIpAddress_example", []openapiclient.PrivateSpaceVpnTunnelPostBody{*openapiclient.NewPrivateSpaceVpnTunnelPostBody("Psk_example", "StartupAction_example")})}) // PrivateSpaceVpnConnectionPostBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreatePrivateSpaceVpnConnection(context.Background(), orgId, privateSpaceId).PrivateSpaceVpnConnectionPostBody(privateSpaceVpnConnectionPostBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreatePrivateSpaceVpnConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePrivateSpaceVpnConnection`: PrivateSpaceVpnConnection
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreatePrivateSpaceVpnConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**privateSpaceId** | **string** | The ID of the private space in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePrivateSpaceVpnConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **privateSpaceVpnConnectionPostBody** | [**PrivateSpaceVpnConnectionPostBody**](PrivateSpaceVpnConnectionPostBody.md) |  | 

### Return type

[**PrivateSpaceVpnConnection**](PrivateSpaceVpnConnection.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePrivateSpace

> DeletePrivateSpace(ctx, orgId, privateSpaceId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/private_space"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	privateSpaceId := "privateSpaceId_example" // string | The ID of the private space in GUID format

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeletePrivateSpace(context.Background(), orgId, privateSpaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeletePrivateSpace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**privateSpaceId** | **string** | The ID of the private space in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePrivateSpaceRequest struct via the builder pattern


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


## DeletePrivateSpaceTransitGateway

> DeletePrivateSpaceTransitGateway(ctx, orgId, privateSpaceId, transitGatewayId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/private_space"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	privateSpaceId := "privateSpaceId_example" // string | The ID of the private space in GUID format
	transitGatewayId := "transitGatewayId_example" // string | The transit gateway id (AWS tgw-...).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeletePrivateSpaceTransitGateway(context.Background(), orgId, privateSpaceId, transitGatewayId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeletePrivateSpaceTransitGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**privateSpaceId** | **string** | The ID of the private space in GUID format | 
**transitGatewayId** | **string** | The transit gateway id (AWS tgw-...). | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePrivateSpaceTransitGatewayRequest struct via the builder pattern


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


## DeletePrivateSpaceVpnConnection

> DeletePrivateSpaceVpnConnection(ctx, orgId, privateSpaceId, connectionId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/private_space"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	privateSpaceId := "privateSpaceId_example" // string | The ID of the private space in GUID format
	connectionId := "connectionId_example" // string | The ID of the VPN connection in GUID format

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeletePrivateSpaceVpnConnection(context.Background(), orgId, privateSpaceId, connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeletePrivateSpaceVpnConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**privateSpaceId** | **string** | The ID of the private space in GUID format | 
**connectionId** | **string** | The ID of the VPN connection in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePrivateSpaceVpnConnectionRequest struct via the builder pattern


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


## DeletePrivateSpaceVpnConnectionMember

> DeletePrivateSpaceVpnConnectionMember(ctx, orgId, privateSpaceId, connectionId, vpnId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/private_space"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	privateSpaceId := "privateSpaceId_example" // string | The ID of the private space in GUID format
	connectionId := "connectionId_example" // string | The ID of the VPN connection in GUID format
	vpnId := "vpnId_example" // string | The ID of the VPN member in GUID format

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeletePrivateSpaceVpnConnectionMember(context.Background(), orgId, privateSpaceId, connectionId, vpnId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeletePrivateSpaceVpnConnectionMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**privateSpaceId** | **string** | The ID of the private space in GUID format | 
**connectionId** | **string** | The ID of the VPN connection in GUID format | 
**vpnId** | **string** | The ID of the VPN member in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePrivateSpaceVpnConnectionMemberRequest struct via the builder pattern


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


## GetOrgTransitGateways

> []TransitGateway GetOrgTransitGateways(ctx, orgId).Region(region).PrivateSpaceId(privateSpaceId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/private_space"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	region := "region_example" // string | AWS region filter. (optional)
	privateSpaceId := "privateSpaceId_example" // string | Private space id filter (GUID). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetOrgTransitGateways(context.Background(), orgId).Region(region).PrivateSpaceId(privateSpaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetOrgTransitGateways``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgTransitGateways`: []TransitGateway
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetOrgTransitGateways`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgTransitGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **region** | **string** | AWS region filter. | 
 **privateSpaceId** | **string** | Private space id filter (GUID). | 

### Return type

[**[]TransitGateway**](TransitGateway.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrivateSpace

> PrivateSpace GetPrivateSpace(ctx, orgId, privateSpaceId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/private_space"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	privateSpaceId := "privateSpaceId_example" // string | The ID of the private space in GUID format

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetPrivateSpace(context.Background(), orgId, privateSpaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetPrivateSpace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrivateSpace`: PrivateSpace
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetPrivateSpace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**privateSpaceId** | **string** | The ID of the private space in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivateSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PrivateSpace**](PrivateSpace.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrivateSpaceIamRoles

> PrivateSpaceIamRoles GetPrivateSpaceIamRoles(ctx, orgId, privateSpaceId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/private_space"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	privateSpaceId := "privateSpaceId_example" // string | The ID of the private space in GUID format

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetPrivateSpaceIamRoles(context.Background(), orgId, privateSpaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetPrivateSpaceIamRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrivateSpaceIamRoles`: PrivateSpaceIamRoles
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetPrivateSpaceIamRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**privateSpaceId** | **string** | The ID of the private space in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivateSpaceIamRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PrivateSpaceIamRoles**](PrivateSpaceIamRoles.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrivateSpaceMulesoftAccount

> string GetPrivateSpaceMulesoftAccount(ctx, orgId, privateSpaceId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/private_space"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	privateSpaceId := "privateSpaceId_example" // string | The ID of the private space in GUID format

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetPrivateSpaceMulesoftAccount(context.Background(), orgId, privateSpaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetPrivateSpaceMulesoftAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrivateSpaceMulesoftAccount`: string
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetPrivateSpaceMulesoftAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**privateSpaceId** | **string** | The ID of the private space in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivateSpaceMulesoftAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrivateSpaceRoutes

> []PrivateSpaceRoute GetPrivateSpaceRoutes(ctx, orgId, privateSpaceId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/private_space"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	privateSpaceId := "privateSpaceId_example" // string | The ID of the private space in GUID format

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetPrivateSpaceRoutes(context.Background(), orgId, privateSpaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetPrivateSpaceRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrivateSpaceRoutes`: []PrivateSpaceRoute
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetPrivateSpaceRoutes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**privateSpaceId** | **string** | The ID of the private space in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivateSpaceRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]PrivateSpaceRoute**](PrivateSpaceRoute.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrivateSpaceSupportedVpnConfigs

> map[string]map[string][]string GetPrivateSpaceSupportedVpnConfigs(ctx, orgId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/private_space"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetPrivateSpaceSupportedVpnConfigs(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetPrivateSpaceSupportedVpnConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrivateSpaceSupportedVpnConfigs`: map[string]map[string][]string
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetPrivateSpaceSupportedVpnConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivateSpaceSupportedVpnConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**map[string]map[string][]string**](map.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrivateSpaceTransitGateways

> []TransitGateway GetPrivateSpaceTransitGateways(ctx, orgId, privateSpaceId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/private_space"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	privateSpaceId := "privateSpaceId_example" // string | The ID of the private space in GUID format

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetPrivateSpaceTransitGateways(context.Background(), orgId, privateSpaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetPrivateSpaceTransitGateways``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrivateSpaceTransitGateways`: []TransitGateway
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetPrivateSpaceTransitGateways`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**privateSpaceId** | **string** | The ID of the private space in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivateSpaceTransitGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]TransitGateway**](TransitGateway.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrivateSpaceVpnConnection

> PrivateSpaceVpnConnection GetPrivateSpaceVpnConnection(ctx, orgId, privateSpaceId, connectionId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/private_space"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	privateSpaceId := "privateSpaceId_example" // string | The ID of the private space in GUID format
	connectionId := "connectionId_example" // string | The ID of the VPN connection in GUID format

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetPrivateSpaceVpnConnection(context.Background(), orgId, privateSpaceId, connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetPrivateSpaceVpnConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrivateSpaceVpnConnection`: PrivateSpaceVpnConnection
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetPrivateSpaceVpnConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**privateSpaceId** | **string** | The ID of the private space in GUID format | 
**connectionId** | **string** | The ID of the VPN connection in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivateSpaceVpnConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PrivateSpaceVpnConnection**](PrivateSpaceVpnConnection.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrivateSpaceVpnConnections

> []PrivateSpaceVpnConnection GetPrivateSpaceVpnConnections(ctx, orgId, privateSpaceId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/private_space"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	privateSpaceId := "privateSpaceId_example" // string | The ID of the private space in GUID format

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetPrivateSpaceVpnConnections(context.Background(), orgId, privateSpaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetPrivateSpaceVpnConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrivateSpaceVpnConnections`: []PrivateSpaceVpnConnection
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetPrivateSpaceVpnConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**privateSpaceId** | **string** | The ID of the private space in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivateSpaceVpnConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]PrivateSpaceVpnConnection**](PrivateSpaceVpnConnection.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrivateSpaces

> PrivateSpaceSummary GetPrivateSpaces(ctx, orgId).Offset(offset).Limit(limit).Sort(sort).Ascending(ascending).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/private_space"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	offset := int32(56) // int32 | The offset specifies the offset of the first row to return (optional)
	limit := int32(56) // int32 | Amount of objects retrieved in the response (optional)
	sort := "sort_example" // string | Property to sort by (optional)
	ascending := "ascending_example" // string | Order for sorting (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetPrivateSpaces(context.Background(), orgId).Offset(offset).Limit(limit).Sort(sort).Ascending(ascending).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetPrivateSpaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrivateSpaces`: PrivateSpaceSummary
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetPrivateSpaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivateSpacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | The offset specifies the offset of the first row to return | 
 **limit** | **int32** | Amount of objects retrieved in the response | 
 **sort** | **string** | Property to sort by | 
 **ascending** | **string** | Order for sorting | 

### Return type

[**PrivateSpaceSummary**](PrivateSpaceSummary.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgTransitGatewayName

> []TransitGateway UpdateOrgTransitGatewayName(ctx, orgId, transitGatewayId).TransitGatewayPatchNameBody(transitGatewayPatchNameBody).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/private_space"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	transitGatewayId := "transitGatewayId_example" // string | The transit gateway id (AWS tgw-...).
	transitGatewayPatchNameBody := *openapiclient.NewTransitGatewayPatchNameBody("Name_example") // TransitGatewayPatchNameBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateOrgTransitGatewayName(context.Background(), orgId, transitGatewayId).TransitGatewayPatchNameBody(transitGatewayPatchNameBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateOrgTransitGatewayName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgTransitGatewayName`: []TransitGateway
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateOrgTransitGatewayName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**transitGatewayId** | **string** | The transit gateway id (AWS tgw-...). | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgTransitGatewayNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitGatewayPatchNameBody** | [**TransitGatewayPatchNameBody**](TransitGatewayPatchNameBody.md) |  | 

### Return type

[**[]TransitGateway**](TransitGateway.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePrivateSpace

> PrivateSpace UpdatePrivateSpace(ctx, orgId, privateSpaceId).PrivateSpacePatchBody(privateSpacePatchBody).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/private_space"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	privateSpaceId := "privateSpaceId_example" // string | The ID of the private space in GUID format
	privateSpacePatchBody := *openapiclient.NewPrivateSpacePatchBody() // PrivateSpacePatchBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdatePrivateSpace(context.Background(), orgId, privateSpaceId).PrivateSpacePatchBody(privateSpacePatchBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdatePrivateSpace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePrivateSpace`: PrivateSpace
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdatePrivateSpace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**privateSpaceId** | **string** | The ID of the private space in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePrivateSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **privateSpacePatchBody** | [**PrivateSpacePatchBody**](PrivateSpacePatchBody.md) |  | 

### Return type

[**PrivateSpace**](PrivateSpace.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePrivateSpaceTransitGatewayRoutes

> []TransitGatewayPatchRoutesResult UpdatePrivateSpaceTransitGatewayRoutes(ctx, orgId, privateSpaceId, transitGatewayId).TransitGatewayPatchRoutesBody(transitGatewayPatchRoutesBody).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/private_space"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	privateSpaceId := "privateSpaceId_example" // string | The ID of the private space in GUID format
	transitGatewayId := "transitGatewayId_example" // string | The transit gateway id (AWS tgw-...).
	transitGatewayPatchRoutesBody := *openapiclient.NewTransitGatewayPatchRoutesBody([]string{"Routes_example"}) // TransitGatewayPatchRoutesBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdatePrivateSpaceTransitGatewayRoutes(context.Background(), orgId, privateSpaceId, transitGatewayId).TransitGatewayPatchRoutesBody(transitGatewayPatchRoutesBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdatePrivateSpaceTransitGatewayRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePrivateSpaceTransitGatewayRoutes`: []TransitGatewayPatchRoutesResult
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdatePrivateSpaceTransitGatewayRoutes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**privateSpaceId** | **string** | The ID of the private space in GUID format | 
**transitGatewayId** | **string** | The transit gateway id (AWS tgw-...). | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePrivateSpaceTransitGatewayRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transitGatewayPatchRoutesBody** | [**TransitGatewayPatchRoutesBody**](TransitGatewayPatchRoutesBody.md) |  | 

### Return type

[**[]TransitGatewayPatchRoutesResult**](TransitGatewayPatchRoutesResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePrivateSpaceVpnConnection

> PrivateSpaceVpnConnection UpdatePrivateSpaceVpnConnection(ctx, orgId, privateSpaceId, connectionId).PrivateSpaceVpnConnectionPatchBody(privateSpaceVpnConnectionPatchBody).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/private_space"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	privateSpaceId := "privateSpaceId_example" // string | The ID of the private space in GUID format
	connectionId := "connectionId_example" // string | The ID of the VPN connection in GUID format
	privateSpaceVpnConnectionPatchBody := *openapiclient.NewPrivateSpaceVpnConnectionPatchBody() // PrivateSpaceVpnConnectionPatchBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdatePrivateSpaceVpnConnection(context.Background(), orgId, privateSpaceId, connectionId).PrivateSpaceVpnConnectionPatchBody(privateSpaceVpnConnectionPatchBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdatePrivateSpaceVpnConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePrivateSpaceVpnConnection`: PrivateSpaceVpnConnection
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdatePrivateSpaceVpnConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**privateSpaceId** | **string** | The ID of the private space in GUID format | 
**connectionId** | **string** | The ID of the VPN connection in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePrivateSpaceVpnConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **privateSpaceVpnConnectionPatchBody** | [**PrivateSpaceVpnConnectionPatchBody**](PrivateSpaceVpnConnectionPatchBody.md) |  | 

### Return type

[**PrivateSpaceVpnConnection**](PrivateSpaceVpnConnection.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePrivateSpaceVpnConnectionMember

> PrivateSpaceVpnConnection UpdatePrivateSpaceVpnConnectionMember(ctx, orgId, privateSpaceId, connectionId, vpnId).PrivateSpaceVpnPatchBody(privateSpaceVpnPatchBody).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/private_space"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	privateSpaceId := "privateSpaceId_example" // string | The ID of the private space in GUID format
	connectionId := "connectionId_example" // string | The ID of the VPN connection in GUID format
	vpnId := "vpnId_example" // string | The ID of the VPN member in GUID format
	privateSpaceVpnPatchBody := *openapiclient.NewPrivateSpaceVpnPatchBody() // PrivateSpaceVpnPatchBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdatePrivateSpaceVpnConnectionMember(context.Background(), orgId, privateSpaceId, connectionId, vpnId).PrivateSpaceVpnPatchBody(privateSpaceVpnPatchBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdatePrivateSpaceVpnConnectionMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePrivateSpaceVpnConnectionMember`: PrivateSpaceVpnConnection
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdatePrivateSpaceVpnConnectionMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**privateSpaceId** | **string** | The ID of the private space in GUID format | 
**connectionId** | **string** | The ID of the VPN connection in GUID format | 
**vpnId** | **string** | The ID of the VPN member in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePrivateSpaceVpnConnectionMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **privateSpaceVpnPatchBody** | [**PrivateSpaceVpnPatchBody**](PrivateSpaceVpnPatchBody.md) |  | 

### Return type

[**PrivateSpaceVpnConnection**](PrivateSpaceVpnConnection.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

