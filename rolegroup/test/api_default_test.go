/*
RoleGroup API

Testing DefaultApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package rolegroup

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/rolegroup"
)

func Test_rolegroup_DefaultApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DefaultApiService OrganizationsOrgIdRolegroupsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.DefaultApi.OrganizationsOrgIdRolegroupsGet(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService OrganizationsOrgIdRolegroupsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.DefaultApi.OrganizationsOrgIdRolegroupsPost(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService OrganizationsOrgIdRolegroupsRolegroupIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var rolegroupId string

		resp, httpRes, err := apiClient.DefaultApi.OrganizationsOrgIdRolegroupsRolegroupIdDelete(context.Background(), orgId, rolegroupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService OrganizationsOrgIdRolegroupsRolegroupIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var rolegroupId string

		resp, httpRes, err := apiClient.DefaultApi.OrganizationsOrgIdRolegroupsRolegroupIdGet(context.Background(), orgId, rolegroupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService OrganizationsOrgIdRolegroupsRolegroupIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var rolegroupId string

		resp, httpRes, err := apiClient.DefaultApi.OrganizationsOrgIdRolegroupsRolegroupIdPut(context.Background(), orgId, rolegroupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
