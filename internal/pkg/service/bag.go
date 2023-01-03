package service

import (
	"context"
	"fmt"

	aaa_client "github.com/os1-logistics/container-reference-app/internal/pkg/clients/aaa"
	container_client "github.com/os1-logistics/container-reference-app/internal/pkg/clients/container"
)

type BagServiceInterface interface {
	GetBags()
}

type BagService struct {
	containerApiClient container_client.APIClient
	aaaApiClient       aaa_client.APIClient
}

func NewBagService() BagService {

	aaaCfg := aaa_client.NewConfiguration()
	aaaCfg.Host = "alpha.preprod.fxtrt.io"
	aaaCfg.Scheme = "https"

	aaaClient := aaa_client.NewAPIClient(aaaCfg)

	containerCfg := container_client.NewConfiguration()
	containerCfg.Host = "alpha.preprod.fxtrt.io"
	containerCfg.Scheme = "https"

	containerClient := container_client.NewAPIClient(containerCfg)

	ctx := context.Background()
	ApiAuthClientCredentialsRequest := aaaClient.AuthenticationApi.AuthClientCredentials(ctx)

	ClientCredentialsRequest := aaa_client.ClientCredentialsRequest{
		ClientId:     "platform:app:ContainerReferenceApp-backend",
		ClientSecret: "****",
		Audience:     *aaa_client.NewNullableString(aaa_client.PtrString("container:reference:app")),
	}

	request := ApiAuthClientCredentialsRequest.ClientCredentialsRequest(ClientCredentialsRequest)
	request = request.XCOREOSREQUESTID("1234")
	request = request.XCOREOSTID("alpha")

	success, response, _ := aaaClient.AuthenticationApi.AuthClientCredentialsExecute(request)
	fmt.Println("=====================================")
	if response.StatusCode == 200 && success != nil {
		fmt.Println(success.Data.GetAccessToken())
	} else {
		fmt.Println("Unable to get token")
	}

	return BagService{
		aaaApiClient:       *aaaClient,
		containerApiClient: *containerClient,
	}
}

func (s BagService) GetBags() {

	fmt.Println("invoked GetBags")
}
