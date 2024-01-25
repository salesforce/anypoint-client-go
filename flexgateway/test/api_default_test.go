/*
Flex Gateway API

Testing DefaultApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package flexgateway

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/flexgateway"
)

func Test_flexgateway_DefaultApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DefaultApiService GetFlexGatewayRegistrationToken", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var envId string

		resp, httpRes, err := apiClient.DefaultApi.GetFlexGatewayRegistrationToken(context.Background(), orgId, envId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService GetFlexGatewayTargetApis", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var envId string
		var flexGatewayTargetId string

		resp, httpRes, err := apiClient.DefaultApi.GetFlexGatewayTargetApis(context.Background(), orgId, envId, flexGatewayTargetId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService GetFlexGatewayTargetById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var envId string
		var flexGatewayTargetId string

		resp, httpRes, err := apiClient.DefaultApi.GetFlexGatewayTargetById(context.Background(), orgId, envId, flexGatewayTargetId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService GetFlexGatewayTargets", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var envId string

		resp, httpRes, err := apiClient.DefaultApi.GetFlexGatewayTargets(context.Background(), orgId, envId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}