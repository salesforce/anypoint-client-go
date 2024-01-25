# Go API client for secretgroup_truststore

Secret Group Truststore API

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
import secretgroup_truststore "github.com/mulesoft-anypoint/anypoint-client-go/secretgroup_truststore"
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
ctx := context.WithValue(context.Background(), secretgroup_truststore.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), secretgroup_truststore.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), secretgroup_truststore.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), secretgroup_truststore.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://anypoint.mulesoft.com/secrets-manager/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**GetSecretGroupTruststoreDetails**](docs/DefaultApi.md#getsecretgrouptruststoredetails) | **Get** /organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId}/truststores/{secretId} | Retrieve truststore details
*DefaultApi* | [**GetSecretGroupTruststores**](docs/DefaultApi.md#getsecretgrouptruststores) | **Get** /organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId}/truststores | Retrieves a secret-groups&#39; collection of truststores.
*DefaultApi* | [**PatchSecretGroupTruststore**](docs/DefaultApi.md#patchsecretgrouptruststore) | **Patch** /organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId}/truststores/{secretId} | Update a given secret-group truststore
*DefaultApi* | [**PostSecretGroupTruststore**](docs/DefaultApi.md#postsecretgrouptruststore) | **Post** /organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId}/truststores | Create a secret-groups&#39; truststore.
*DefaultApi* | [**PutSecretGroupTruststore**](docs/DefaultApi.md#putsecretgrouptruststore) | **Put** /organizations/{orgId}/environments/{envId}/secretGroups/{secretGroupId}/truststores/{secretId} | Update a given secret-group truststore


## Documentation For Models

 - [CertificateDetails](docs/CertificateDetails.md)
 - [CertificateDetailsBasicConstraints](docs/CertificateDetailsBasicConstraints.md)
 - [CertificateEntry](docs/CertificateEntry.md)
 - [CertificateValidity](docs/CertificateValidity.md)
 - [ErrorsResponse](docs/ErrorsResponse.md)
 - [ErrorsResponseErrorsInner](docs/ErrorsResponseErrorsInner.md)
 - [GetSecretGroupTruststoreDetails404Response](docs/GetSecretGroupTruststoreDetails404Response.md)
 - [IssuerSubject](docs/IssuerSubject.md)
 - [Meta](docs/Meta.md)
 - [PostSecretGroupTruststore201Response](docs/PostSecretGroupTruststore201Response.md)
 - [PutSecretGroupTruststore200Response](docs/PutSecretGroupTruststore200Response.md)
 - [Truststore](docs/Truststore.md)
 - [TruststoreDetails](docs/TruststoreDetails.md)
 - [TruststoreSummary](docs/TruststoreSummary.md)


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


