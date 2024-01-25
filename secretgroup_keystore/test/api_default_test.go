/*
Secret Group Keystore API

Testing DefaultApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package secretgroup_keystore

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/secretgroup_keystore"
)

func Test_secretgroup_keystore_DefaultApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DefaultApiService GetSecretGroupKeystoreDetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var envId string
		var secretGroupId string
		var secretId string

		resp, httpRes, err := apiClient.DefaultApi.GetSecretGroupKeystoreDetails(context.Background(), orgId, envId, secretGroupId, secretId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService GetSecretGroupKeystores", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var envId string
		var secretGroupId string

		resp, httpRes, err := apiClient.DefaultApi.GetSecretGroupKeystores(context.Background(), orgId, envId, secretGroupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService PatchSecretGroupKeystore", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var envId string
		var secretGroupId string
		var secretId string

		resp, httpRes, err := apiClient.DefaultApi.PatchSecretGroupKeystore(context.Background(), orgId, envId, secretGroupId, secretId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService PostSecretGroupKeystores", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var envId string
		var secretGroupId string

		resp, httpRes, err := apiClient.DefaultApi.PostSecretGroupKeystores(context.Background(), orgId, envId, secretGroupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService PutSecretGroupKeystore", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var envId string
		var secretGroupId string
		var secretId string

		resp, httpRes, err := apiClient.DefaultApi.PutSecretGroupKeystore(context.Background(), orgId, envId, secretGroupId, secretId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}