/*
 * Dedicated Load Balancer API
 *
 * Description of the DLB API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dlb

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// DefaultApiService DefaultApi service
type DefaultApiService service

type DefaultApiApiOrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdDeleteRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	orgId string
	vpcId string
	dlbId string
}


func (r DefaultApiApiOrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdDeleteRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdDeleteExecute(r)
}

/*
 * OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdDelete Method for OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdDelete
 * Deletes a DLB
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgId The ID of the organization in GUID format
 * @param vpcId The ID of the VPC in GUID format
 * @param dlbId The ID of the DLB in GUID format
 * @return DefaultApiApiOrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdDeleteRequest
 */
func (a *DefaultApiService) OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdDelete(ctx _context.Context, orgId string, vpcId string, dlbId string) DefaultApiApiOrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdDeleteRequest {
	return DefaultApiApiOrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		vpcId: vpcId,
		dlbId: dlbId,
	}
}

/*
 * Execute executes the request
 */
func (a *DefaultApiService) OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdDeleteExecute(r DefaultApiApiOrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdDeleteRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdDelete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/vpcs/{vpcId}/loadbalancers/{dlbId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", _neturl.PathEscape(parameterToString(r.orgId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"vpcId"+"}", _neturl.PathEscape(parameterToString(r.vpcId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"dlbId"+"}", _neturl.PathEscape(parameterToString(r.dlbId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type DefaultApiApiOrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdGetRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	orgId string
	vpcId string
	dlbId string
}


func (r DefaultApiApiOrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdGetRequest) Execute() (Dlb, *_nethttp.Response, error) {
	return r.ApiService.OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdGetExecute(r)
}

/*
 * OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdGet Method for OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdGet
 * Retrieves a DLB by id
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgId The ID of the organization in GUID format
 * @param vpcId The ID of the VPC in GUID format
 * @param dlbId The ID of the DLB in GUID format
 * @return DefaultApiApiOrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdGetRequest
 */
func (a *DefaultApiService) OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdGet(ctx _context.Context, orgId string, vpcId string, dlbId string) DefaultApiApiOrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdGetRequest {
	return DefaultApiApiOrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdGetRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		vpcId: vpcId,
		dlbId: dlbId,
	}
}

/*
 * Execute executes the request
 * @return Dlb
 */
func (a *DefaultApiService) OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdGetExecute(r DefaultApiApiOrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdGetRequest) (Dlb, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Dlb
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/vpcs/{vpcId}/loadbalancers/{dlbId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", _neturl.PathEscape(parameterToString(r.orgId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"vpcId"+"}", _neturl.PathEscape(parameterToString(r.vpcId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"dlbId"+"}", _neturl.PathEscape(parameterToString(r.dlbId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type DefaultApiApiOrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdPatchRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	orgId string
	vpcId string
	dlbId string
	requestBody *[]map[string]interface{}
}

func (r DefaultApiApiOrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdPatchRequest) RequestBody(requestBody []map[string]interface{}) DefaultApiApiOrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdPatchRequest {
	r.requestBody = &requestBody
	return r
}

func (r DefaultApiApiOrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdPatchRequest) Execute() (Dlb, *_nethttp.Response, error) {
	return r.ApiService.OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdPatchExecute(r)
}

/*
 * OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdPatch Method for OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdPatch
 * Updates a dlb. uses JSON Patch body object
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgId The ID of the organization in GUID format
 * @param vpcId The ID of the VPC in GUID format
 * @param dlbId The ID of the DLB in GUID format
 * @return DefaultApiApiOrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdPatchRequest
 */
func (a *DefaultApiService) OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdPatch(ctx _context.Context, orgId string, vpcId string, dlbId string) DefaultApiApiOrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdPatchRequest {
	return DefaultApiApiOrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdPatchRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		vpcId: vpcId,
		dlbId: dlbId,
	}
}

/*
 * Execute executes the request
 * @return Dlb
 */
func (a *DefaultApiService) OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdPatchExecute(r DefaultApiApiOrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdPatchRequest) (Dlb, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Dlb
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.OrganizationsOrgIdVpcsVpcIdLoadbalancersDlbIdPatch")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/vpcs/{vpcId}/loadbalancers/{dlbId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", _neturl.PathEscape(parameterToString(r.orgId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"vpcId"+"}", _neturl.PathEscape(parameterToString(r.vpcId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"dlbId"+"}", _neturl.PathEscape(parameterToString(r.dlbId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.requestBody
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v InlineResponse400
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type DefaultApiApiOrganizationsOrgIdVpcsVpcIdLoadbalancersGetRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	orgId string
	vpcId string
}


func (r DefaultApiApiOrganizationsOrgIdVpcsVpcIdLoadbalancersGetRequest) Execute() (InlineResponse200, *_nethttp.Response, error) {
	return r.ApiService.OrganizationsOrgIdVpcsVpcIdLoadbalancersGetExecute(r)
}

/*
 * OrganizationsOrgIdVpcsVpcIdLoadbalancersGet Method for OrganizationsOrgIdVpcsVpcIdLoadbalancersGet
 * Returns all loadbalancers in the given vpc
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgId The ID of the organization in GUID format
 * @param vpcId The ID of the VPC in GUID format
 * @return DefaultApiApiOrganizationsOrgIdVpcsVpcIdLoadbalancersGetRequest
 */
func (a *DefaultApiService) OrganizationsOrgIdVpcsVpcIdLoadbalancersGet(ctx _context.Context, orgId string, vpcId string) DefaultApiApiOrganizationsOrgIdVpcsVpcIdLoadbalancersGetRequest {
	return DefaultApiApiOrganizationsOrgIdVpcsVpcIdLoadbalancersGetRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		vpcId: vpcId,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse200
 */
func (a *DefaultApiService) OrganizationsOrgIdVpcsVpcIdLoadbalancersGetExecute(r DefaultApiApiOrganizationsOrgIdVpcsVpcIdLoadbalancersGetRequest) (InlineResponse200, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.OrganizationsOrgIdVpcsVpcIdLoadbalancersGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/vpcs/{vpcId}/loadbalancers"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", _neturl.PathEscape(parameterToString(r.orgId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"vpcId"+"}", _neturl.PathEscape(parameterToString(r.vpcId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type DefaultApiApiOrganizationsOrgIdVpcsVpcIdLoadbalancersPostRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	orgId string
	vpcId string
	dlbCore *DlbCore
}

func (r DefaultApiApiOrganizationsOrgIdVpcsVpcIdLoadbalancersPostRequest) DlbCore(dlbCore DlbCore) DefaultApiApiOrganizationsOrgIdVpcsVpcIdLoadbalancersPostRequest {
	r.dlbCore = &dlbCore
	return r
}

func (r DefaultApiApiOrganizationsOrgIdVpcsVpcIdLoadbalancersPostRequest) Execute() (InlineResponse200, *_nethttp.Response, error) {
	return r.ApiService.OrganizationsOrgIdVpcsVpcIdLoadbalancersPostExecute(r)
}

/*
 * OrganizationsOrgIdVpcsVpcIdLoadbalancersPost Method for OrganizationsOrgIdVpcsVpcIdLoadbalancersPost
 * create a DLB
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgId The ID of the organization in GUID format
 * @param vpcId The ID of the VPC in GUID format
 * @return DefaultApiApiOrganizationsOrgIdVpcsVpcIdLoadbalancersPostRequest
 */
func (a *DefaultApiService) OrganizationsOrgIdVpcsVpcIdLoadbalancersPost(ctx _context.Context, orgId string, vpcId string) DefaultApiApiOrganizationsOrgIdVpcsVpcIdLoadbalancersPostRequest {
	return DefaultApiApiOrganizationsOrgIdVpcsVpcIdLoadbalancersPostRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		vpcId: vpcId,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse200
 */
func (a *DefaultApiService) OrganizationsOrgIdVpcsVpcIdLoadbalancersPostExecute(r DefaultApiApiOrganizationsOrgIdVpcsVpcIdLoadbalancersPostRequest) (InlineResponse200, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.OrganizationsOrgIdVpcsVpcIdLoadbalancersPost")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/vpcs/{vpcId}/loadbalancers"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", _neturl.PathEscape(parameterToString(r.orgId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"vpcId"+"}", _neturl.PathEscape(parameterToString(r.vpcId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.dlbCore
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v InlineResponse400
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
