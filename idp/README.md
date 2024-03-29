# Go API client for idp

Description of Identity Provider API

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import idp "github.com/mulesoft-anypoint/anypoint-client-go/idp"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), idp.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), idp.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), idp.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), idp.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://anypoint.mulesoft.com/accounts/api*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**OrganizationsOrgIdIdentityProvidersGet**](docs/DefaultApi.md#organizationsorgididentityprovidersget) | **Get** /organizations/{orgId}/identityProviders | 
*DefaultApi* | [**OrganizationsOrgIdIdentityProvidersIdpIdDelete**](docs/DefaultApi.md#organizationsorgididentityprovidersidpiddelete) | **Delete** /organizations/{orgId}/identityProviders/{idpId} | 
*DefaultApi* | [**OrganizationsOrgIdIdentityProvidersIdpIdGet**](docs/DefaultApi.md#organizationsorgididentityprovidersidpidget) | **Get** /organizations/{orgId}/identityProviders/{idpId} | 
*DefaultApi* | [**OrganizationsOrgIdIdentityProvidersIdpIdPatch**](docs/DefaultApi.md#organizationsorgididentityprovidersidpidpatch) | **Patch** /organizations/{orgId}/identityProviders/{idpId} | 
*DefaultApi* | [**OrganizationsOrgIdIdentityProvidersPost**](docs/DefaultApi.md#organizationsorgididentityproviderspost) | **Post** /organizations/{orgId}/identityProviders | 


## Documentation For Models

 - [ClaimsMapping](docs/ClaimsMapping.md)
 - [ClaimsMapping1](docs/ClaimsMapping1.md)
 - [ClaimsMapping2](docs/ClaimsMapping2.md)
 - [Client](docs/Client.md)
 - [Client1](docs/Client1.md)
 - [Credentials](docs/Credentials.md)
 - [Credentials1](docs/Credentials1.md)
 - [Idp](docs/Idp.md)
 - [IdpCommonProps](docs/IdpCommonProps.md)
 - [IdpOIDC](docs/IdpOIDC.md)
 - [IdpPatchBody](docs/IdpPatchBody.md)
 - [IdpPatchBodyType](docs/IdpPatchBodyType.md)
 - [IdpPostBody](docs/IdpPostBody.md)
 - [IdpPostBodyType](docs/IdpPostBodyType.md)
 - [IdpSAML](docs/IdpSAML.md)
 - [IdpSummary](docs/IdpSummary.md)
 - [ModelType](docs/ModelType.md)
 - [OidcProvider](docs/OidcProvider.md)
 - [OidcProvider1](docs/OidcProvider1.md)
 - [OrganizationsOrgIdIdentityProvidersGet200Response](docs/OrganizationsOrgIdIdentityProvidersGet200Response.md)
 - [OrganizationsOrgIdIdentityProvidersIdpIdPatch400Response](docs/OrganizationsOrgIdIdentityProvidersIdpIdPatch400Response.md)
 - [OrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner](docs/OrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner.md)
 - [Saml](docs/Saml.md)
 - [Saml1](docs/Saml1.md)
 - [ServiceProvider](docs/ServiceProvider.md)
 - [ServiceProvider1](docs/ServiceProvider1.md)
 - [Urls](docs/Urls.md)
 - [Urls1](docs/Urls1.md)
 - [Urls2](docs/Urls2.md)
 - [Urls3](docs/Urls3.md)
 - [Urls4](docs/Urls4.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### bearerAuth

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



