# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com/accounts/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsOrgIdTeamsTeamIdMembersGet**](DefaultApi.md#OrganizationsOrgIdTeamsTeamIdMembersGet) | **Get** /organizations/{orgId}/teams/{teamId}/members | 
[**OrganizationsOrgIdTeamsTeamIdMembersUserIdDelete**](DefaultApi.md#OrganizationsOrgIdTeamsTeamIdMembersUserIdDelete) | **Delete** /organizations/{orgId}/teams/{teamId}/members/{userId} | 
[**OrganizationsOrgIdTeamsTeamIdMembersUserIdPut**](DefaultApi.md#OrganizationsOrgIdTeamsTeamIdMembersUserIdPut) | **Put** /organizations/{orgId}/teams/{teamId}/members/{userId} | 



## OrganizationsOrgIdTeamsTeamIdMembersGet

> TeamMemberCollection OrganizationsOrgIdTeamsTeamIdMembersGet(ctx, orgId, teamId).MembershipType(membershipType).IdentityType(identityType).MemberIds(memberIds).Search(search).Limit(limit).Offset(offset).Sort(sort).Ascending(ascending).Execute()





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
    membershipType := "membershipType_example" // string | Include the group access mappings that grant the provided membership type By default, all group access mappings are returned (optional)
    identityType := "identityType_example" // string | A search string to use for case-insensitive partial matches on external group name (optional)
    memberIds := "memberIds_example" // string | Include the members of the team that have ids in this list (optional)
    search := "search_example" // string | A search string to use for case-insensitive partial matches on team name (optional)
    limit := int32(56) // int32 | Maximum number of rolegroups to retrieve per request. (optional)
    offset := int32(56) // int32 | The number of records to omit from the response. (optional)
    sort := "sort_example" // string | The field to sort on. default identity_type (optional)
    ascending := true // bool | Whether to sort ascending or descending. Default true (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdTeamsTeamIdMembersGet(context.Background(), orgId, teamId).MembershipType(membershipType).IdentityType(identityType).MemberIds(memberIds).Search(search).Limit(limit).Offset(offset).Sort(sort).Ascending(ascending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdTeamsTeamIdMembersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrgIdTeamsTeamIdMembersGet`: TeamMemberCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrganizationsOrgIdTeamsTeamIdMembersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**teamId** | **string** | The ID of the team in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdTeamsTeamIdMembersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **membershipType** | **string** | Include the group access mappings that grant the provided membership type By default, all group access mappings are returned | 
 **identityType** | **string** | A search string to use for case-insensitive partial matches on external group name | 
 **memberIds** | **string** | Include the members of the team that have ids in this list | 
 **search** | **string** | A search string to use for case-insensitive partial matches on team name | 
 **limit** | **int32** | Maximum number of rolegroups to retrieve per request. | 
 **offset** | **int32** | The number of records to omit from the response. | 
 **sort** | **string** | The field to sort on. default identity_type | 
 **ascending** | **bool** | Whether to sort ascending or descending. Default true | 

### Return type

[**TeamMemberCollection**](TeamMemberCollection.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrgIdTeamsTeamIdMembersUserIdDelete

> OrganizationsOrgIdTeamsTeamIdMembersUserIdDelete(ctx, orgId, teamId, userId).Execute()





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
    userId := "userId_example" // string | The ID of the user in GUID format

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdTeamsTeamIdMembersUserIdDelete(context.Background(), orgId, teamId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdTeamsTeamIdMembersUserIdDelete``: %v\n", err)
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
**userId** | **string** | The ID of the user in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdTeamsTeamIdMembersUserIdDeleteRequest struct via the builder pattern


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


## OrganizationsOrgIdTeamsTeamIdMembersUserIdPut

> OrganizationsOrgIdTeamsTeamIdMembersUserIdPut(ctx, orgId, teamId, userId).TeamMemberPutBody(teamMemberPutBody).Execute()





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
    userId := "userId_example" // string | The ID of the user in GUID format
    teamMemberPutBody := *openapiclient.NewTeamMemberPutBody() // TeamMemberPutBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OrganizationsOrgIdTeamsTeamIdMembersUserIdPut(context.Background(), orgId, teamId, userId).TeamMemberPutBody(teamMemberPutBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrganizationsOrgIdTeamsTeamIdMembersUserIdPut``: %v\n", err)
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
**userId** | **string** | The ID of the user in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrgIdTeamsTeamIdMembersUserIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **teamMemberPutBody** | [**TeamMemberPutBody**](TeamMemberPutBody.md) |  | 

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

