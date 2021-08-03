# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com/accounts/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsOrgIdTeamsGet**](DefaultApi.md#OrganizationsOrgIdTeamsGet) | **Get** /organizations/{orgId}/teams | 
[**OrganizationsOrgIdTeamsPost**](DefaultApi.md#OrganizationsOrgIdTeamsPost) | **Post** /organizations/{orgId}/teams | 
[**OrganizationsOrgIdTeamsTeamIdDelete**](DefaultApi.md#OrganizationsOrgIdTeamsTeamIdDelete) | **Delete** /organizations/{orgId}/teams/{teamId} | 
[**OrganizationsOrgIdTeamsTeamIdGet**](DefaultApi.md#OrganizationsOrgIdTeamsTeamIdGet) | **Get** /organizations/{orgId}/teams/{teamId} | 
[**OrganizationsOrgIdTeamsTeamIdPatch**](DefaultApi.md#OrganizationsOrgIdTeamsTeamIdPatch) | **Patch** /organizations/{orgId}/teams/{teamId} | 
[**OrganizationsOrgIdTeamsTeamIdPut**](DefaultApi.md#OrganizationsOrgIdTeamsTeamIdPut) | **Put** /organizations/{orgId}/teams/{teamId} | 



## OrganizationsOrgIdTeamsGet

> TeamCollection OrganizationsOrgIdTeamsGet(ctx, orgId).AncestorTeamId(ancestorTeamId).ParentTeamId(parentTeamId).TeamId(teamId).TeamType(teamType).Search(search).Limit(limit).Offset(offset).Sort(sort).Ascending(ascending).Execute()





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
    orgId := "orgId_example" // string | The ID of the organization in GUID format
    ancestorTeamId := "ancestorTeamId_example" // string | team_id that must appear in the team's ancestor_team_ids. (optional)
    parentTeamId := "parentTeamId_example" // string | team_id of the immediate parent of the team to return. (optional)
    teamId := "teamId_example" // string | id of the team to return. (optional)
    teamType := "teamType_example" // string | return only teams that are of this type (optional)
    search := "search_example" // string | A search string to use for case-insensitive partial matches on team name (optional)
    limit := int32(56) // int32 | Maximum number of rolegroups to retrieve per request. (optional)
    offset := int32(56) // int32 | The number of records to omit from the response. (optional)
    sort := "sort_example" // string | The field to sort on. default team_name (optional)
    ascending := true // bool | Whether to sort ascending or descending. Default true (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdTeamsGet(context.Background(), orgId).AncestorTeamId(ancestorTeamId).ParentTeamId(parentTeamId).TeamId(teamId).TeamType(teamType).Search(search).Limit(limit).Offset(offset).Sort(sort).Ascending(ascending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdTeamsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdTeamsGet`: TeamCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdTeamsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdTeamsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ancestorTeamId** | **string** | team_id that must appear in the team&#39;s ancestor_team_ids. | 
 **parentTeamId** | **string** | team_id of the immediate parent of the team to return. | 
 **teamId** | **string** | id of the team to return. | 
 **teamType** | **string** | return only teams that are of this type | 
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


## OrganizationsOrgIdTeamsPost

> Team OrganizationsOrgIdTeamsPost(ctx, orgId).TeamPostBody(teamPostBody).Execute()





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
    orgId := "orgId_example" // string | The ID of the organization in GUID format
    teamPostBody := *openapiclient.NewTeamPostBody() // TeamPostBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdTeamsPost(context.Background(), orgId).TeamPostBody(teamPostBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdTeamsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdTeamsPost`: Team
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdTeamsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdTeamsPostRequest struct via the builder pattern


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


## OrganizationsOrgIdTeamsTeamIdDelete

> OrganizationsOrgIdTeamsTeamIdDelete(ctx, orgId, teamId).Execute()





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
    orgId := "orgId_example" // string | The ID of the organization in GUID format
    teamId := "teamId_example" // string | The ID of the team in GUID format

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdTeamsTeamIdDelete(context.Background(), orgId, teamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdTeamsTeamIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiOrganizationsOrgIdTeamsTeamIdDeleteRequest struct via the builder pattern


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


## OrganizationsOrgIdTeamsTeamIdGet

> Team OrganizationsOrgIdTeamsTeamIdGet(ctx, orgId, teamId).Execute()





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
    orgId := "orgId_example" // string | The ID of the organization in GUID format
    teamId := "teamId_example" // string | The ID of the team in GUID format

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdTeamsTeamIdGet(context.Background(), orgId, teamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdTeamsTeamIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdTeamsTeamIdGet`: Team
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdTeamsTeamIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**teamId** | **string** | The ID of the team in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdTeamsTeamIdGetRequest struct via the builder pattern


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


## OrganizationsOrgIdTeamsTeamIdPatch

> Team OrganizationsOrgIdTeamsTeamIdPatch(ctx, orgId, teamId).TeamPatchBody(teamPatchBody).Execute()





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
    orgId := "orgId_example" // string | The ID of the organization in GUID format
    teamId := "teamId_example" // string | The ID of the team in GUID format
    teamPatchBody := *openapiclient.NewTeamPatchBody() // TeamPatchBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdTeamsTeamIdPatch(context.Background(), orgId, teamId).TeamPatchBody(teamPatchBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdTeamsTeamIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdTeamsTeamIdPatch`: Team
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdTeamsTeamIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**teamId** | **string** | The ID of the team in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdTeamsTeamIdPatchRequest struct via the builder pattern


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


## OrganizationsOrgIdTeamsTeamIdPut

> Team OrganizationsOrgIdTeamsTeamIdPut(ctx, orgId, teamId).TeamPutBody(teamPutBody).Execute()





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
    orgId := "orgId_example" // string | The ID of the organization in GUID format
    teamId := "teamId_example" // string | The ID of the team in GUID format
    teamPutBody := *openapiclient.NewTeamPutBody() // TeamPutBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdTeamsTeamIdPut(context.Background(), orgId, teamId).TeamPutBody(teamPutBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdTeamsTeamIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdTeamsTeamIdPut`: Team
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdTeamsTeamIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**teamId** | **string** | The ID of the team in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdTeamsTeamIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **teamPutBody** | [**TeamPutBody**](TeamPutBody.md) |  | 

### Return type

[**Team**](team.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

