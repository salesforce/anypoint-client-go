/*
Secret Group Truststore API

Secret Group Truststore API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package secretgroup_truststore

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"os"
)


// DefaultApiService DefaultApi service
type DefaultApiService service

type DefaultApiGetSecretGroupTruststoreDetailsRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	orgId string
	envId string
	secretGroupId string
	secretId string
}

func (r DefaultApiGetSecretGroupTruststoreDetailsRequest) Execute() (*Truststore, *http.Response, error) {
	return r.ApiService.GetSecretGroupTruststoreDetailsExecute(r)
}

/*
GetSecretGroupTruststoreDetails Retrieve truststore details

Retrieves truststore details by id for a given secret group in a given organization and environment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId The organization Id
 @param envId The environment id
 @param secretGroupId The secret group id
 @param secretId The truststore id
 @return DefaultApiGetSecretGroupTruststoreDetailsRequest
*/
func (a *DefaultApiService) GetSecretGroupTruststoreDetails(ctx context.Context, orgId string, envId string, secretGroupId string, secretId string) DefaultApiGetSecretGroupTruststoreDetailsRequest {
	return DefaultApiGetSecretGroupTruststoreDetailsRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		envId: envId,
		secretGroupId: secretGroupId,
		secretId: secretId,
	}
}

// Execute executes the request
//  @return Truststore
func (a *DefaultApiService) GetSecretGroupTruststoreDetailsExecute(r DefaultApiGetSecretGroupTruststoreDetailsRequest) (*Truststore, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Truststore
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.GetSecretGroupTruststoreDetails")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId}/truststores/{secretId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"envId"+"}", url.PathEscape(parameterValueToString(r.envId, "envId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"secretGroupId"+"}", url.PathEscape(parameterValueToString(r.secretGroupId, "secretGroupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"secretId"+"}", url.PathEscape(parameterValueToString(r.secretId, "secretId")), -1)

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
		if localVarHTTPResponse.StatusCode == 404 {
			var v GetSecretGroupTruststoreDetails404Response
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

type DefaultApiGetSecretGroupTruststoresRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	orgId string
	envId string
	secretGroupId string
	type_ *string
}

// Filter the elements on the response to be of a specific type from {PEM, JKS, JCEKS, PKCS12}
func (r DefaultApiGetSecretGroupTruststoresRequest) Type_(type_ string) DefaultApiGetSecretGroupTruststoresRequest {
	r.type_ = &type_
	return r
}

func (r DefaultApiGetSecretGroupTruststoresRequest) Execute() ([]TruststoreSummary, *http.Response, error) {
	return r.ApiService.GetSecretGroupTruststoresExecute(r)
}

/*
GetSecretGroupTruststores Retrieves a secret-groups' collection of truststores.

Retrieves a secret-groups' collection of truststores.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId The organization Id
 @param envId The environment id
 @param secretGroupId The secret group id
 @return DefaultApiGetSecretGroupTruststoresRequest
*/
func (a *DefaultApiService) GetSecretGroupTruststores(ctx context.Context, orgId string, envId string, secretGroupId string) DefaultApiGetSecretGroupTruststoresRequest {
	return DefaultApiGetSecretGroupTruststoresRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		envId: envId,
		secretGroupId: secretGroupId,
	}
}

// Execute executes the request
//  @return []TruststoreSummary
func (a *DefaultApiService) GetSecretGroupTruststoresExecute(r DefaultApiGetSecretGroupTruststoresRequest) ([]TruststoreSummary, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []TruststoreSummary
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.GetSecretGroupTruststores")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId}/truststores"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"envId"+"}", url.PathEscape(parameterValueToString(r.envId, "envId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"secretGroupId"+"}", url.PathEscape(parameterValueToString(r.secretGroupId, "secretGroupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "")
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorsResponse
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

type DefaultApiPatchSecretGroupTruststoreRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	orgId string
	envId string
	secretGroupId string
	secretId string
	body *map[string]interface{}
}

func (r DefaultApiPatchSecretGroupTruststoreRequest) Body(body map[string]interface{}) DefaultApiPatchSecretGroupTruststoreRequest {
	r.body = &body
	return r
}

func (r DefaultApiPatchSecretGroupTruststoreRequest) Execute() (*PutSecretGroupTruststore200Response, *http.Response, error) {
	return r.ApiService.PatchSecretGroupTruststoreExecute(r)
}

/*
PatchSecretGroupTruststore Update a given secret-group truststore

Update truststore details for a given secret-group in a given organization and environment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId The organization Id
 @param envId The environment id
 @param secretGroupId The secret group id
 @param secretId The truststore id
 @return DefaultApiPatchSecretGroupTruststoreRequest
*/
func (a *DefaultApiService) PatchSecretGroupTruststore(ctx context.Context, orgId string, envId string, secretGroupId string, secretId string) DefaultApiPatchSecretGroupTruststoreRequest {
	return DefaultApiPatchSecretGroupTruststoreRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		envId: envId,
		secretGroupId: secretGroupId,
		secretId: secretId,
	}
}

// Execute executes the request
//  @return PutSecretGroupTruststore200Response
func (a *DefaultApiService) PatchSecretGroupTruststoreExecute(r DefaultApiPatchSecretGroupTruststoreRequest) (*PutSecretGroupTruststore200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PutSecretGroupTruststore200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.PatchSecretGroupTruststore")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId}/truststores/{secretId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"envId"+"}", url.PathEscape(parameterValueToString(r.envId, "envId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"secretGroupId"+"}", url.PathEscape(parameterValueToString(r.secretGroupId, "secretGroupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"secretId"+"}", url.PathEscape(parameterValueToString(r.secretId, "secretId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json:"}

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
	localVarPostBody = r.body
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
			var v GetSecretGroupTruststoreDetails404Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorsResponse
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

type DefaultApiPostSecretGroupTruststoreRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	orgId string
	envId string
	secretGroupId string
	allowExpiredCert *bool
	expirationDate *string
	name *string
	type_ *string
	trustStore *os.File
	algorithm *string
	storePassphrase *string
}

// With &#39;true&#39; to allow uploading expired certificates
func (r DefaultApiPostSecretGroupTruststoreRequest) AllowExpiredCert(allowExpiredCert bool) DefaultApiPostSecretGroupTruststoreRequest {
	r.allowExpiredCert = &allowExpiredCert
	return r
}

// Date on which this secret should expire. If not set, by default, it will be set to notAfter date of the public certificate from this keystore. Once the secret expires, a grant can not be requested for it. 
func (r DefaultApiPostSecretGroupTruststoreRequest) ExpirationDate(expirationDate string) DefaultApiPostSecretGroupTruststoreRequest {
	r.expirationDate = &expirationDate
	return r
}

// The name of the truststore instance
func (r DefaultApiPostSecretGroupTruststoreRequest) Name(name string) DefaultApiPostSecretGroupTruststoreRequest {
	r.name = &name
	return r
}

func (r DefaultApiPostSecretGroupTruststoreRequest) Type_(type_ string) DefaultApiPostSecretGroupTruststoreRequest {
	r.type_ = &type_
	return r
}

// File containing one or more trusted certificate entries
func (r DefaultApiPostSecretGroupTruststoreRequest) TrustStore(trustStore *os.File) DefaultApiPostSecretGroupTruststoreRequest {
	r.trustStore = trustStore
	return r
}

// The algorithm used to initialize TrustManagerFactory
func (r DefaultApiPostSecretGroupTruststoreRequest) Algorithm(algorithm string) DefaultApiPostSecretGroupTruststoreRequest {
	r.algorithm = &algorithm
	return r
}

// The passphrase with which the trustStore file is protected
func (r DefaultApiPostSecretGroupTruststoreRequest) StorePassphrase(storePassphrase string) DefaultApiPostSecretGroupTruststoreRequest {
	r.storePassphrase = &storePassphrase
	return r
}

func (r DefaultApiPostSecretGroupTruststoreRequest) Execute() (*PostSecretGroupTruststore201Response, *http.Response, error) {
	return r.ApiService.PostSecretGroupTruststoreExecute(r)
}

/*
PostSecretGroupTruststore Create a secret-groups' truststore.

Create a secret-groups' truststore.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId The organization Id
 @param envId The environment id
 @param secretGroupId The secret group id
 @return DefaultApiPostSecretGroupTruststoreRequest
*/
func (a *DefaultApiService) PostSecretGroupTruststore(ctx context.Context, orgId string, envId string, secretGroupId string) DefaultApiPostSecretGroupTruststoreRequest {
	return DefaultApiPostSecretGroupTruststoreRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		envId: envId,
		secretGroupId: secretGroupId,
	}
}

// Execute executes the request
//  @return PostSecretGroupTruststore201Response
func (a *DefaultApiService) PostSecretGroupTruststoreExecute(r DefaultApiPostSecretGroupTruststoreRequest) (*PostSecretGroupTruststore201Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostSecretGroupTruststore201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.PostSecretGroupTruststore")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId}/truststores"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"envId"+"}", url.PathEscape(parameterValueToString(r.envId, "envId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"secretGroupId"+"}", url.PathEscape(parameterValueToString(r.secretGroupId, "secretGroupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.allowExpiredCert == nil {
		return localVarReturnValue, nil, reportError("allowExpiredCert is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "allowExpiredCert", r.allowExpiredCert, "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	if r.expirationDate != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "expirationDate", r.expirationDate, "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "name", r.name, "")
	}
	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "type", r.type_, "")
	}
	var trustStoreLocalVarFormFileName string
	var trustStoreLocalVarFileName     string
	var trustStoreLocalVarFileBytes    []byte

	trustStoreLocalVarFormFileName = "trustStore"


	trustStoreLocalVarFile := r.trustStore

	if trustStoreLocalVarFile != nil {
		fbs, _ := io.ReadAll(trustStoreLocalVarFile)

		trustStoreLocalVarFileBytes = fbs
		trustStoreLocalVarFileName = trustStoreLocalVarFile.Name()
		trustStoreLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: trustStoreLocalVarFileBytes, fileName: trustStoreLocalVarFileName, formFileName: trustStoreLocalVarFormFileName})
	}
	if r.algorithm != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "algorithm", r.algorithm, "")
	}
	if r.storePassphrase != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "storePassphrase", r.storePassphrase, "")
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorsResponse
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

type DefaultApiPutSecretGroupTruststoreRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	orgId string
	envId string
	secretGroupId string
	secretId string
	allowExpiredCert *bool
	expirationDate *string
	name *string
	type_ *string
	trustStore *os.File
	algorithm *string
	storePassphrase *string
}

// With &#39;true&#39; to allow uploading expired certificates
func (r DefaultApiPutSecretGroupTruststoreRequest) AllowExpiredCert(allowExpiredCert bool) DefaultApiPutSecretGroupTruststoreRequest {
	r.allowExpiredCert = &allowExpiredCert
	return r
}

// Date on which this secret should expire. If not set, by default, it will be set to notAfter date of the public certificate from this keystore. Once the secret expires, a grant can not be requested for it. 
func (r DefaultApiPutSecretGroupTruststoreRequest) ExpirationDate(expirationDate string) DefaultApiPutSecretGroupTruststoreRequest {
	r.expirationDate = &expirationDate
	return r
}

// The name of the truststore instance
func (r DefaultApiPutSecretGroupTruststoreRequest) Name(name string) DefaultApiPutSecretGroupTruststoreRequest {
	r.name = &name
	return r
}

func (r DefaultApiPutSecretGroupTruststoreRequest) Type_(type_ string) DefaultApiPutSecretGroupTruststoreRequest {
	r.type_ = &type_
	return r
}

// File containing one or more trusted certificate entries
func (r DefaultApiPutSecretGroupTruststoreRequest) TrustStore(trustStore *os.File) DefaultApiPutSecretGroupTruststoreRequest {
	r.trustStore = trustStore
	return r
}

// The algorithm used to initialize TrustManagerFactory
func (r DefaultApiPutSecretGroupTruststoreRequest) Algorithm(algorithm string) DefaultApiPutSecretGroupTruststoreRequest {
	r.algorithm = &algorithm
	return r
}

// The passphrase with which the trustStore file is protected
func (r DefaultApiPutSecretGroupTruststoreRequest) StorePassphrase(storePassphrase string) DefaultApiPutSecretGroupTruststoreRequest {
	r.storePassphrase = &storePassphrase
	return r
}

func (r DefaultApiPutSecretGroupTruststoreRequest) Execute() (*PutSecretGroupTruststore200Response, *http.Response, error) {
	return r.ApiService.PutSecretGroupTruststoreExecute(r)
}

/*
PutSecretGroupTruststore Update a given secret-group truststore

Update truststore details for a given secret-group in a given organization and environment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId The organization Id
 @param envId The environment id
 @param secretGroupId The secret group id
 @param secretId The truststore id
 @return DefaultApiPutSecretGroupTruststoreRequest
*/
func (a *DefaultApiService) PutSecretGroupTruststore(ctx context.Context, orgId string, envId string, secretGroupId string, secretId string) DefaultApiPutSecretGroupTruststoreRequest {
	return DefaultApiPutSecretGroupTruststoreRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		envId: envId,
		secretGroupId: secretGroupId,
		secretId: secretId,
	}
}

// Execute executes the request
//  @return PutSecretGroupTruststore200Response
func (a *DefaultApiService) PutSecretGroupTruststoreExecute(r DefaultApiPutSecretGroupTruststoreRequest) (*PutSecretGroupTruststore200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PutSecretGroupTruststore200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.PutSecretGroupTruststore")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId}/truststores/{secretId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"envId"+"}", url.PathEscape(parameterValueToString(r.envId, "envId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"secretGroupId"+"}", url.PathEscape(parameterValueToString(r.secretGroupId, "secretGroupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"secretId"+"}", url.PathEscape(parameterValueToString(r.secretId, "secretId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.allowExpiredCert == nil {
		return localVarReturnValue, nil, reportError("allowExpiredCert is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "allowExpiredCert", r.allowExpiredCert, "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	if r.expirationDate != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "expirationDate", r.expirationDate, "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "name", r.name, "")
	}
	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "type", r.type_, "")
	}
	var trustStoreLocalVarFormFileName string
	var trustStoreLocalVarFileName     string
	var trustStoreLocalVarFileBytes    []byte

	trustStoreLocalVarFormFileName = "trustStore"


	trustStoreLocalVarFile := r.trustStore

	if trustStoreLocalVarFile != nil {
		fbs, _ := io.ReadAll(trustStoreLocalVarFile)

		trustStoreLocalVarFileBytes = fbs
		trustStoreLocalVarFileName = trustStoreLocalVarFile.Name()
		trustStoreLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: trustStoreLocalVarFileBytes, fileName: trustStoreLocalVarFileName, formFileName: trustStoreLocalVarFormFileName})
	}
	if r.algorithm != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "algorithm", r.algorithm, "")
	}
	if r.storePassphrase != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "storePassphrase", r.storePassphrase, "")
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v GetSecretGroupTruststoreDetails404Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorsResponse
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
