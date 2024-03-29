/*
Anypoint MQ Exchange specfication

Anypoint MQ Exchange API specification

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ame

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

type DefaultApiCreateAMERequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	orgId string
	envId string
	regionId string
	exchangeId string
	exchangeBody *ExchangeBody
}

func (r DefaultApiCreateAMERequest) ExchangeBody(exchangeBody ExchangeBody) DefaultApiCreateAMERequest {
	r.exchangeBody = &exchangeBody
	return r
}

func (r DefaultApiCreateAMERequest) Execute() (*Exchange, *http.Response, error) {
	return r.ApiService.CreateAMEExecute(r)
}

/*
CreateAME Method for CreateAME

Create exchange

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId The organization Id
 @param envId The environment id
 @param regionId The region id
 @param exchangeId The id of a specific exchange
 @return DefaultApiCreateAMERequest
*/
func (a *DefaultApiService) CreateAME(ctx context.Context, orgId string, envId string, regionId string, exchangeId string) DefaultApiCreateAMERequest {
	return DefaultApiCreateAMERequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		envId: envId,
		regionId: regionId,
		exchangeId: exchangeId,
	}
}

// Execute executes the request
//  @return Exchange
func (a *DefaultApiService) CreateAMEExecute(r DefaultApiCreateAMERequest) (*Exchange, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Exchange
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.CreateAME")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/environments/{envId}/regions/{regionId}/destinations/exchanges/{exchangeId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"envId"+"}", url.PathEscape(parameterValueToString(r.envId, "envId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"regionId"+"}", url.PathEscape(parameterValueToString(r.regionId, "regionId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"exchangeId"+"}", url.PathEscape(parameterValueToString(r.exchangeId, "exchangeId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.exchangeBody
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

type DefaultApiDeleteAMERequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	orgId string
	envId string
	regionId string
	exchangeId string
}

func (r DefaultApiDeleteAMERequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteAMEExecute(r)
}

/*
DeleteAME Method for DeleteAME

Delete an exchange

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId The organization Id
 @param envId The environment id
 @param regionId The region id
 @param exchangeId The id of a specific exchange
 @return DefaultApiDeleteAMERequest
*/
func (a *DefaultApiService) DeleteAME(ctx context.Context, orgId string, envId string, regionId string, exchangeId string) DefaultApiDeleteAMERequest {
	return DefaultApiDeleteAMERequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		envId: envId,
		regionId: regionId,
		exchangeId: exchangeId,
	}
}

// Execute executes the request
func (a *DefaultApiService) DeleteAMEExecute(r DefaultApiDeleteAMERequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.DeleteAME")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/environments/{envId}/regions/{regionId}/destinations/exchanges/{exchangeId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"envId"+"}", url.PathEscape(parameterValueToString(r.envId, "envId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"regionId"+"}", url.PathEscape(parameterValueToString(r.regionId, "regionId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"exchangeId"+"}", url.PathEscape(parameterValueToString(r.exchangeId, "exchangeId")), -1)

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

type DefaultApiGetAMERequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	orgId string
	envId string
	regionId string
	exchangeId string
}

func (r DefaultApiGetAMERequest) Execute() (*Exchange, *http.Response, error) {
	return r.ApiService.GetAMEExecute(r)
}

/*
GetAME Method for GetAME

Get details about an exchange

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId The organization Id
 @param envId The environment id
 @param regionId The region id
 @param exchangeId The id of a specific exchange
 @return DefaultApiGetAMERequest
*/
func (a *DefaultApiService) GetAME(ctx context.Context, orgId string, envId string, regionId string, exchangeId string) DefaultApiGetAMERequest {
	return DefaultApiGetAMERequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		envId: envId,
		regionId: regionId,
		exchangeId: exchangeId,
	}
}

// Execute executes the request
//  @return Exchange
func (a *DefaultApiService) GetAMEExecute(r DefaultApiGetAMERequest) (*Exchange, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Exchange
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.GetAME")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/environments/{envId}/regions/{regionId}/destinations/exchanges/{exchangeId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"envId"+"}", url.PathEscape(parameterValueToString(r.envId, "envId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"regionId"+"}", url.PathEscape(parameterValueToString(r.regionId, "regionId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"exchangeId"+"}", url.PathEscape(parameterValueToString(r.exchangeId, "exchangeId")), -1)

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

type DefaultApiUpdateAMERequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	orgId string
	envId string
	regionId string
	exchangeId string
	exchangeBody *ExchangeBody
}

func (r DefaultApiUpdateAMERequest) ExchangeBody(exchangeBody ExchangeBody) DefaultApiUpdateAMERequest {
	r.exchangeBody = &exchangeBody
	return r
}

func (r DefaultApiUpdateAMERequest) Execute() (*Exchange, *http.Response, error) {
	return r.ApiService.UpdateAMEExecute(r)
}

/*
UpdateAME Method for UpdateAME

Modify an exchange's properties

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId The organization Id
 @param envId The environment id
 @param regionId The region id
 @param exchangeId The id of a specific exchange
 @return DefaultApiUpdateAMERequest
*/
func (a *DefaultApiService) UpdateAME(ctx context.Context, orgId string, envId string, regionId string, exchangeId string) DefaultApiUpdateAMERequest {
	return DefaultApiUpdateAMERequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		envId: envId,
		regionId: regionId,
		exchangeId: exchangeId,
	}
}

// Execute executes the request
//  @return Exchange
func (a *DefaultApiService) UpdateAMEExecute(r DefaultApiUpdateAMERequest) (*Exchange, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Exchange
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.UpdateAME")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/environments/{envId}/regions/{regionId}/destinations/exchanges/{exchangeId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"envId"+"}", url.PathEscape(parameterValueToString(r.envId, "envId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"regionId"+"}", url.PathEscape(parameterValueToString(r.regionId, "regionId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"exchangeId"+"}", url.PathEscape(parameterValueToString(r.exchangeId, "exchangeId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.exchangeBody
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v UpdateAME404Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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
