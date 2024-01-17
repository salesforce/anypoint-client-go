# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com/accounts/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsOrgIdTeamsTeamIdRolesDelete**](DefaultApi.md#OrganizationsOrgIdTeamsTeamIdRolesDelete) | **Delete** /organizations/{orgId}/teams/{teamId}/roles | 
[**OrganizationsOrgIdTeamsTeamIdRolesGet**](DefaultApi.md#OrganizationsOrgIdTeamsTeamIdRolesGet) | **Get** /organizations/{orgId}/teams/{teamId}/roles | 
[**OrganizationsOrgIdTeamsTeamIdRolesPost**](DefaultApi.md#OrganizationsOrgIdTeamsTeamIdRolesPost) | **Post** /organizations/{orgId}/teams/{teamId}/roles | 



## OrganizationsOrgIdTeamsTeamIdRolesDelete

> OrganizationsOrgIdTeamsTeamIdRolesDelete(ctx, orgId, teamId).TeamRoleDeleteBody(teamRoleDeleteBody).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/team_roles"
)

func main() {
    orgId := "orgId_example" // string | The ID of the organization in GUID format
    teamId := "teamId_example" // string | The ID of the team in GUID format
    teamRoleDeleteBody := []openapiclient.TeamRoleDeleteBody{*openapiclient.NewTeamRoleDeleteBody()} // []TeamRoleDeleteBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.OrganizationsOrgIdTeamsTeamIdRolesDelete(context.Background(), orgId, teamId).TeamRoleDeleteBody(teamRoleDeleteBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdTeamsTeamIdRolesDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiOrganizationsOrgIdTeamsTeamIdRolesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **teamRoleDeleteBody** | [**[]TeamRoleDeleteBody**](TeamRoleDeleteBody.md) |  | 

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


## OrganizationsOrgIdTeamsTeamIdRolesGet

> TeamRoleCollection OrganizationsOrgIdTeamsTeamIdRolesGet(ctx, orgId, teamId).RoleId(roleId).Search(search).Limit(limit).Offset(offset).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/team_roles"
)

func main() {
    orgId := "orgId_example" // string | The ID of the organization in GUID format
    teamId := "teamId_example" // string | The ID of the team in GUID format
    roleId := "roleId_example" // string | return only role assignments containing one of the supplied role_ids (optional)
    search := "search_example" // string | A search string to use for case-insensitive partial matches on team name (optional)
    limit := int32(56) // int32 | Maximum number of rolegroups to retrieve per request. (optional)
    offset := int32(56) // int32 | The number of records to omit from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OrganizationsOrgIdTeamsTeamIdRolesGet(context.Background(), orgId, teamId).RoleId(roleId).Search(search).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdTeamsTeamIdRolesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdTeamsTeamIdRolesGet`: TeamRoleCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdTeamsTeamIdRolesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**teamId** | **string** | The ID of the team in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdTeamsTeamIdRolesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **roleId** | **string** | return only role assignments containing one of the supplied role_ids | 
 **search** | **string** | A search string to use for case-insensitive partial matches on team name | 
 **limit** | **int32** | Maximum number of rolegroups to retrieve per request. | 
 **offset** | **int32** | The number of records to omit from the response. | 

### Return type

[**TeamRoleCollection**](TeamRoleCollection.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrgIdTeamsTeamIdRolesPost

> OrganizationsOrgIdTeamsTeamIdRolesPost(ctx, orgId, teamId).TeamRolePostBody(teamRolePostBody).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/team_roles"
)

func main() {
    orgId := "orgId_example" // string | The ID of the organization in GUID format
    teamId := "teamId_example" // string | The ID of the team in GUID format
    teamRolePostBody := []openapiclient.TeamRolePostBody{*openapiclient.NewTeamRolePostBody()} // []TeamRolePostBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.OrganizationsOrgIdTeamsTeamIdRolesPost(context.Background(), orgId, teamId).TeamRolePostBody(teamRolePostBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdTeamsTeamIdRolesPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiOrganizationsOrgIdTeamsTeamIdRolesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **teamRolePostBody** | [**[]TeamRolePostBody**](TeamRolePostBody.md) |  | 

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

