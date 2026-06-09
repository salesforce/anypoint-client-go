# \DefaultAPI

All URIs are relative to *https://anypoint.mulesoft.com/accounts/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTeam**](DefaultAPI.md#CreateTeam) | **Post** /organizations/{orgId}/teams | Create a new team
[**DeleteTeam**](DefaultAPI.md#DeleteTeam) | **Delete** /organizations/{orgId}/teams/{teamId} | Delete a team by id
[**GetTeam**](DefaultAPI.md#GetTeam) | **Get** /organizations/{orgId}/teams/{teamId} | Read a team by id
[**GetTeams**](DefaultAPI.md#GetTeams) | **Get** /organizations/{orgId}/teams | Read all Teams for a given org
[**MoveTeam**](DefaultAPI.md#MoveTeam) | **Put** /organizations/{orgId}/teams/{teamId}/parent | Move a team
[**UpdateTeam**](DefaultAPI.md#UpdateTeam) | **Patch** /organizations/{orgId}/teams/{teamId} | Update a team by id



## CreateTeam

> Team CreateTeam(ctx, orgId).TeamPostBody(teamPostBody).Execute()

Create a new team



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/team"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	teamPostBody := *openapiclient.NewTeamPostBody() // TeamPostBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateTeam(context.Background(), orgId).TeamPostBody(teamPostBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTeam`: Team
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **teamPostBody** | [**TeamPostBody**](TeamPostBody.md) |  | 

### Return type

[**Team**](Team.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTeam

> DeleteTeam(ctx, orgId, teamId).Execute()

Delete a team by id



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/team"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	teamId := "teamId_example" // string | The ID of the team in GUID format

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteTeam(context.Background(), orgId, teamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteTeam``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteTeamRequest struct via the builder pattern


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


## GetTeam

> Team GetTeam(ctx, orgId, teamId).Execute()

Read a team by id



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/team"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	teamId := "teamId_example" // string | The ID of the team in GUID format

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetTeam(context.Background(), orgId, teamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeam`: Team
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**teamId** | **string** | The ID of the team in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Team**](Team.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeams

> TeamCollection GetTeams(ctx, orgId).AncestorTeamId(ancestorTeamId).ParentTeamId(parentTeamId).TeamId(teamId).TeamType(teamType).Search(search).Limit(limit).Offset(offset).Sort(sort).Ascending(ascending).Execute()

Read all Teams for a given org



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/team"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	ancestorTeamId := []string{"Inner_example"} // []string | team_id that must appear in the team's ancestor_team_ids. (optional)
	parentTeamId := []string{"Inner_example"} // []string | team_id of the immediate parent of the team to return. (optional)
	teamId := "teamId_example" // string | id of the team to return. (optional)
	teamType := "teamType_example" // string | return only teams that are of the given type. Enum are internal, shared, private, external (optional)
	search := "search_example" // string | A search string to use for case-insensitive partial matches on team name (optional)
	limit := int32(56) // int32 | Maximum number of rolegroups to retrieve per request. (optional)
	offset := int32(56) // int32 | The number of records to omit from the response. (optional)
	sort := "sort_example" // string | The field to sort on. default team_name (optional)
	ascending := true // bool | Whether to sort ascending or descending. Default true (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetTeams(context.Background(), orgId).AncestorTeamId(ancestorTeamId).ParentTeamId(parentTeamId).TeamId(teamId).TeamType(teamType).Search(search).Limit(limit).Offset(offset).Sort(sort).Ascending(ascending).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetTeams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeams`: TeamCollection
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetTeams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ancestorTeamId** | **[]string** | team_id that must appear in the team&#39;s ancestor_team_ids. | 
 **parentTeamId** | **[]string** | team_id of the immediate parent of the team to return. | 
 **teamId** | **string** | id of the team to return. | 
 **teamType** | **string** | return only teams that are of the given type. Enum are internal, shared, private, external | 
 **search** | **string** | A search string to use for case-insensitive partial matches on team name | 
 **limit** | **int32** | Maximum number of rolegroups to retrieve per request. | 
 **offset** | **int32** | The number of records to omit from the response. | 
 **sort** | **string** | The field to sort on. default team_name | 
 **ascending** | **bool** | Whether to sort ascending or descending. Default true | 

### Return type

[**TeamCollection**](TeamCollection.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveTeam

> PutTeamResponse MoveTeam(ctx, orgId, teamId).TeamPutBody(teamPutBody).Execute()

Move a team



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/team"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	teamId := "teamId_example" // string | The ID of the team in GUID format
	teamPutBody := *openapiclient.NewTeamPutBody() // TeamPutBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.MoveTeam(context.Background(), orgId, teamId).TeamPutBody(teamPutBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.MoveTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MoveTeam`: PutTeamResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.MoveTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**teamId** | **string** | The ID of the team in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **teamPutBody** | [**TeamPutBody**](TeamPutBody.md) |  | 

### Return type

[**PutTeamResponse**](PutTeamResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTeam

> Team UpdateTeam(ctx, orgId, teamId).TeamPatchBody(teamPatchBody).Execute()

Update a team by id



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/team"
)

func main() {
	orgId := "orgId_example" // string | The ID of the organization in GUID format
	teamId := "teamId_example" // string | The ID of the team in GUID format
	teamPatchBody := *openapiclient.NewTeamPatchBody() // TeamPatchBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateTeam(context.Background(), orgId, teamId).TeamPatchBody(teamPatchBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTeam`: Team
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**teamId** | **string** | The ID of the team in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **teamPatchBody** | [**TeamPatchBody**](TeamPatchBody.md) |  | 

### Return type

[**Team**](Team.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

