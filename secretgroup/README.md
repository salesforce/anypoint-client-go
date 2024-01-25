# Go API client for secretgroup

Secret Group API

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
import secretgroup "github.com/mulesoft-anypoint/anypoint-client-go/secretgroup"
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
ctx := context.WithValue(context.Background(), secretgroup.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), secretgroup.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), secretgroup.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), secretgroup.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://anypoint.mulesoft.com/secrets-manager/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**DeleteSecretGroup**](docs/DefaultApi.md#deletesecretgroup) | **Delete** /organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId} | Delete a secret group
*DefaultApi* | [**GetEnvSecretGroups**](docs/DefaultApi.md#getenvsecretgroups) | **Get** /organizations/{orgId}/environments/{envId}/secretGroups | Retrieves a collection of secret groups
*DefaultApi* | [**GetSecretGroup**](docs/DefaultApi.md#getsecretgroup) | **Get** /organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId} | Retrieve a secret group
*DefaultApi* | [**PatchSecretGroup**](docs/DefaultApi.md#patchsecretgroup) | **Patch** /organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId} | Update a secret group
*DefaultApi* | [**PostSecretGroup**](docs/DefaultApi.md#postsecretgroup) | **Post** /organizations/{orgId}/environments/{envId}/secretGroups | Creates a secret group


## Documentation For Models

 - [ErrorsResponse](docs/ErrorsResponse.md)
 - [ErrorsResponseErrorsInner](docs/ErrorsResponseErrorsInner.md)
 - [GetEnvSecretGroups404Response](docs/GetEnvSecretGroups404Response.md)
 - [Meta](docs/Meta.md)
 - [SecretGroup](docs/SecretGroup.md)
 - [SecretGroupPatchBody](docs/SecretGroupPatchBody.md)
 - [SecretGroupPatchResponse](docs/SecretGroupPatchResponse.md)
 - [SecretGroupPostBody](docs/SecretGroupPostBody.md)
 - [SecretGroupPostResponse](docs/SecretGroupPostResponse.md)


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


