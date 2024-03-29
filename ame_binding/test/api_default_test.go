/*
Anypoint MQ Exchange Binding specfication

Testing DefaultApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package ame_binding

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/ame_binding"
)

func Test_ame_binding_DefaultApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DefaultApiService CreateAMEBinding", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var envId string
		var regionId string
		var exchangeId string
		var queueId string

		resp, httpRes, err := apiClient.DefaultApi.CreateAMEBinding(context.Background(), orgId, envId, regionId, exchangeId, queueId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService CreateAMEBindingRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var envId string
		var regionId string
		var exchangeId string
		var queueId string

		resp, httpRes, err := apiClient.DefaultApi.CreateAMEBindingRule(context.Background(), orgId, envId, regionId, exchangeId, queueId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService DeleteAMEBinding", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var envId string
		var regionId string
		var exchangeId string
		var queueId string

		httpRes, err := apiClient.DefaultApi.DeleteAMEBinding(context.Background(), orgId, envId, regionId, exchangeId, queueId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService DeleteAMEBindingRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var envId string
		var regionId string
		var exchangeId string
		var queueId string

		httpRes, err := apiClient.DefaultApi.DeleteAMEBindingRule(context.Background(), orgId, envId, regionId, exchangeId, queueId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService GetAMEBinding", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var envId string
		var regionId string
		var exchangeId string
		var queueId string

		resp, httpRes, err := apiClient.DefaultApi.GetAMEBinding(context.Background(), orgId, envId, regionId, exchangeId, queueId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
