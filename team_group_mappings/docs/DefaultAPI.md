# \DefaultAPI

All URIs are relative to *https://anypoint.mulesoft.com/accounts/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTeamGroupMappings**](DefaultAPI.md#GetTeamGroupMappings) | **Get** /organizations/{orgId}/teams/{teamId}/groupmappings | 
[**PutTeamGroupMappings**](DefaultAPI.md#PutTeamGroupMappings) | **Put** /organizations/{orgId}/teams/{teamId}/groupmappings | 



## GetTeamGroupMappings

> TeamGroupMappingsCollection GetTeamGroupMappings(ctx, orgId, teamId).Limit(limit).Offset(offset).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/team_group_mappings"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	teamId := "teamId_example" // string | The ID of the team in GUID format
	limit := int32(56) // int32 | Maximum number of rolegroups to retrieve per request. (optional)
	offset := int32(56) // int32 | The number of records to omit from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetTeamGroupMappings(context.Background(), orgId, teamId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetTeamGroupMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamGroupMappings`: TeamGroupMappingsCollection
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetTeamGroupMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**teamId** | **string** | The ID of the team in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamGroupMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Maximum number of rolegroups to retrieve per request. | 
 **offset** | **int32** | The number of records to omit from the response. | 

### Return type

[**TeamGroupMappingsCollection**](TeamGroupMappingsCollection.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutTeamGroupMappings

> PutTeamGroupMappings(ctx, orgId, teamId).TeamGroupMappingPutBody(teamGroupMappingPutBody).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/team_group_mappings"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	teamId := "teamId_example" // string | The ID of the team in GUID format
	teamGroupMappingPutBody := []openapiclient.TeamGroupMappingPutBody{*openapiclient.NewTeamGroupMappingPutBody()} // []TeamGroupMappingPutBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.PutTeamGroupMappings(context.Background(), orgId, teamId).TeamGroupMappingPutBody(teamGroupMappingPutBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PutTeamGroupMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**teamId** | **string** | The ID of the team in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutTeamGroupMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **teamGroupMappingPutBody** | [**[]TeamGroupMappingPutBody**](TeamGroupMappingPutBody.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

