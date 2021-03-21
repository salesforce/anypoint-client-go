package main

import (
	"context"
	"fmt"

	vpc "github.com/mulesoft-consulting/cloudhub-client-go/vpc"
)

func main() {
	orgId := "0b94e66b-0a10-4445-bdc0-9e892bd1c609"

	ctx := context.Background()

	cfg := vpc.NewConfiguration()
	cfg.AddDefaultHeader("Authorization", "Bearer 5729f864-7061-4e57-9270-5f1ab7e75f3d")
	client := vpc.NewAPIClient(cfg)

	_, httpresponse, err := client.DefaultApi.OrganizationsOrgIdVpcsGet(ctx, orgId).Execute()

	if err != nil {
		fmt.Errorf(err.Error())
	}

	fmt.Printf("Response Status %d", httpresponse.StatusCode)
	fmt.Printf("Body: %s", httpresponse.Body)

}
