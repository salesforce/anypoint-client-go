/*
Team Roles API

Testing DefaultApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package team_roles

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/team_roles"
)

func Test_team_roles_DefaultApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DefaultApiService OrganizationsOrgIdTeamsTeamIdRolesDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var teamId string

		httpRes, err := apiClient.DefaultApi.OrganizationsOrgIdTeamsTeamIdRolesDelete(context.Background(), orgId, teamId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService OrganizationsOrgIdTeamsTeamIdRolesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var teamId string

		resp, httpRes, err := apiClient.DefaultApi.OrganizationsOrgIdTeamsTeamIdRolesGet(context.Background(), orgId, teamId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService OrganizationsOrgIdTeamsTeamIdRolesPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var teamId string

		httpRes, err := apiClient.DefaultApi.OrganizationsOrgIdTeamsTeamIdRolesPost(context.Background(), orgId, teamId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
