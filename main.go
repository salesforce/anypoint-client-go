package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	auth "github.com/mulesoft-consulting/cloudhub-client-go/authenticate"
	vpc "github.com/mulesoft-consulting/cloudhub-client-go/vpc"
)

// Regexp definitions
var keyMatchRegex = regexp.MustCompile(`\"([\w\s]+)\":`)
var wordBarrierRegex = regexp.MustCompile(`([a-z]+)([A-Z])`)

type conventionalMarshaller struct {
	Value interface{}
}

func (c conventionalMarshaller) MarshalJSON() ([]byte, error) {
	marshalled, err := json.Marshal(c.Value)
	converted := keyMatchRegex.ReplaceAllFunc(
		marshalled,
		func(match []byte) []byte {
			noSpaceMatch := []byte(strings.ReplaceAll(string(match), " ", ""))
			return bytes.ToLower(wordBarrierRegex.ReplaceAll(
				noSpaceMatch,
				[]byte(`${1}_${2}`),
			))
		},
	)

	return converted, err
}

func main() {
	orgId := "a9e7fe3f-c09f-4b05-9b2f-c786e009ce94"
	client_id := "7627f157dac94390951b4d804d218289"
	client_secret := "75Db67bE3DFf420A9e038adAff3CEFBd"

	ctx := context.Background()

	cfgauth := auth.NewConfiguration()
	credentials := auth.NewCredentialsWithDefaults()
	credentials.SetClientId(client_id)
	credentials.SetClientSecret(client_secret)
	c := auth.NewAPIClient(cfgauth)

	authres, httpauthr, err := c.DefaultApi.Oauth2TokenPost(ctx).Credentials(*credentials).Execute()
	if err != nil {
		fmt.Printf("ERROR %s", err.Error())
		return
	}
	defer httpauthr.Body.Close()

	fmt.Printf("Auth AccessToken %s\n", authres.GetAccessToken())

	authctx := context.WithValue(ctx, vpc.ContextAccessToken, authres.GetAccessToken())

	cfg := vpc.NewConfiguration()
	//cfg.AddDefaultHeader("Authorization", "Bearer "+authres.GetAccessToken())
	client := vpc.NewAPIClient(cfg)

	vpcObj, httpresponse, err := client.DefaultApi.OrganizationsOrgIdVpcsGet(authctx, orgId).Execute()

	if err != nil {
		b, _ := ioutil.ReadAll(httpresponse.Body)
		fmt.Printf("ERROR \n  statusCode %s  \n Detail %s", err.Error(), b)
		return
	}
	defer httpresponse.Body.Close()

	fmt.Printf("Response Status %d\n", httpresponse.StatusCode)

	encoded, _ := json.Marshal(conventionalMarshaller{vpcObj.GetData()})
	vpcs := make([]map[string]interface{}, 0)
	err = json.Unmarshal(encoded, &vpcs)
	if err != nil {
		fmt.Printf("Error %s\n", err.Error())
		return
	}

	//fmt.Printf("Body: %s \n", httpresponse.Body)
	fmt.Printf("VPC Object\n %s", encoded)

}
