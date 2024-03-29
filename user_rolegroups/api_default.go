/*
User RoleGroups API

Description of the User Group API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package user_rolegroups

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// DefaultApiService DefaultApi service
type DefaultApiService service

type DefaultApiOrganizationsOrgIdUsersUserIdRolegroupsGetRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	orgId string
	userId string
	limit *int32
	offset *int32
}

// Maximum number of rolegroups to retrieve per request.
func (r DefaultApiOrganizationsOrgIdUsersUserIdRolegroupsGetRequest) Limit(limit int32) DefaultApiOrganizationsOrgIdUsersUserIdRolegroupsGetRequest {
	r.limit = &limit
	return r
}

// The number of records to omit from the response.
func (r DefaultApiOrganizationsOrgIdUsersUserIdRolegroupsGetRequest) Offset(offset int32) DefaultApiOrganizationsOrgIdUsersUserIdRolegroupsGetRequest {
	r.offset = &offset
	return r
}

func (r DefaultApiOrganizationsOrgIdUsersUserIdRolegroupsGetRequest) Execute() (*OrganizationsOrgIdUsersUserIdRolegroupsGet200Response, *http.Response, error) {
	return r.ApiService.OrganizationsOrgIdUsersUserIdRolegroupsGetExecute(r)
}

/*
OrganizationsOrgIdUsersUserIdRolegroupsGet Method for OrganizationsOrgIdUsersUserIdRolegroupsGet

Returns the users' rolegroups.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId The ID of the organization in GUID format
 @param userId The ID of the user in GUID format
 @return DefaultApiOrganizationsOrgIdUsersUserIdRolegroupsGetRequest
*/
func (a *DefaultApiService) OrganizationsOrgIdUsersUserIdRolegroupsGet(ctx context.Context, orgId string, userId string) DefaultApiOrganizationsOrgIdUsersUserIdRolegroupsGetRequest {
	return DefaultApiOrganizationsOrgIdUsersUserIdRolegroupsGetRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		userId: userId,
	}
}

// Execute executes the request
//  @return OrganizationsOrgIdUsersUserIdRolegroupsGet200Response
func (a *DefaultApiService) OrganizationsOrgIdUsersUserIdRolegroupsGetExecute(r DefaultApiOrganizationsOrgIdUsersUserIdRolegroupsGetRequest) (*OrganizationsOrgIdUsersUserIdRolegroupsGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OrganizationsOrgIdUsersUserIdRolegroupsGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.OrganizationsOrgIdUsersUserIdRolegroupsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/users/{userId}/rolegroups"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type DefaultApiOrganizationsOrgIdUsersUserIdRolegroupsRolegroupIdDeleteRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	orgId string
	userId string
	rolegroupId string
}

func (r DefaultApiOrganizationsOrgIdUsersUserIdRolegroupsRolegroupIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.OrganizationsOrgIdUsersUserIdRolegroupsRolegroupIdDeleteExecute(r)
}

/*
OrganizationsOrgIdUsersUserIdRolegroupsRolegroupIdDelete Method for OrganizationsOrgIdUsersUserIdRolegroupsRolegroupIdDelete

Delete user rolegroup assignment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId The ID of the organization in GUID format
 @param userId The ID of the user
 @param rolegroupId The ID of the user's rolegroup
 @return DefaultApiOrganizationsOrgIdUsersUserIdRolegroupsRolegroupIdDeleteRequest
*/
func (a *DefaultApiService) OrganizationsOrgIdUsersUserIdRolegroupsRolegroupIdDelete(ctx context.Context, orgId string, userId string, rolegroupId string) DefaultApiOrganizationsOrgIdUsersUserIdRolegroupsRolegroupIdDeleteRequest {
	return DefaultApiOrganizationsOrgIdUsersUserIdRolegroupsRolegroupIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		userId: userId,
		rolegroupId: rolegroupId,
	}
}

// Execute executes the request
func (a *DefaultApiService) OrganizationsOrgIdUsersUserIdRolegroupsRolegroupIdDeleteExecute(r DefaultApiOrganizationsOrgIdUsersUserIdRolegroupsRolegroupIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.OrganizationsOrgIdUsersUserIdRolegroupsRolegroupIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/users/{userId}/rolegroups/{rolegroupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rolegroupId"+"}", url.PathEscape(parameterValueToString(r.rolegroupId, "rolegroupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type DefaultApiOrganizationsOrgIdUsersUserIdRolegroupsRolegroupIdPostRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	orgId string
	userId string
	rolegroupId string
}

func (r DefaultApiOrganizationsOrgIdUsersUserIdRolegroupsRolegroupIdPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.OrganizationsOrgIdUsersUserIdRolegroupsRolegroupIdPostExecute(r)
}

/*
OrganizationsOrgIdUsersUserIdRolegroupsRolegroupIdPost Method for OrganizationsOrgIdUsersUserIdRolegroupsRolegroupIdPost

Assign rolegroup to user

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId The ID of the organization in GUID format
 @param userId The ID of the user
 @param rolegroupId The ID of the user's rolegroup
 @return DefaultApiOrganizationsOrgIdUsersUserIdRolegroupsRolegroupIdPostRequest
*/
func (a *DefaultApiService) OrganizationsOrgIdUsersUserIdRolegroupsRolegroupIdPost(ctx context.Context, orgId string, userId string, rolegroupId string) DefaultApiOrganizationsOrgIdUsersUserIdRolegroupsRolegroupIdPostRequest {
	return DefaultApiOrganizationsOrgIdUsersUserIdRolegroupsRolegroupIdPostRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		userId: userId,
		rolegroupId: rolegroupId,
	}
}

// Execute executes the request
func (a *DefaultApiService) OrganizationsOrgIdUsersUserIdRolegroupsRolegroupIdPostExecute(r DefaultApiOrganizationsOrgIdUsersUserIdRolegroupsRolegroupIdPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.OrganizationsOrgIdUsersUserIdRolegroupsRolegroupIdPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/users/{userId}/rolegroups/{rolegroupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rolegroupId"+"}", url.PathEscape(parameterValueToString(r.rolegroupId, "rolegroupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
