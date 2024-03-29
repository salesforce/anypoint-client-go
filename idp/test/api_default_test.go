/*
Identity Provider Management API

Testing DefaultApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package idp

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/idp"
)

func Test_idp_DefaultApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DefaultApiService OrganizationsOrgIdIdentityProvidersGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.DefaultApi.OrganizationsOrgIdIdentityProvidersGet(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService OrganizationsOrgIdIdentityProvidersIdpIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var idpId string

		httpRes, err := apiClient.DefaultApi.OrganizationsOrgIdIdentityProvidersIdpIdDelete(context.Background(), orgId, idpId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService OrganizationsOrgIdIdentityProvidersIdpIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var idpId string

		resp, httpRes, err := apiClient.DefaultApi.OrganizationsOrgIdIdentityProvidersIdpIdGet(context.Background(), orgId, idpId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService OrganizationsOrgIdIdentityProvidersIdpIdPatch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var idpId string

		resp, httpRes, err := apiClient.DefaultApi.OrganizationsOrgIdIdentityProvidersIdpIdPatch(context.Background(), orgId, idpId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService OrganizationsOrgIdIdentityProvidersPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.DefaultApi.OrganizationsOrgIdIdentityProvidersPost(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
