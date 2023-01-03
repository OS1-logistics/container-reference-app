package service

import (
	"context"
	"fmt"

	aaa_client "github.com/os1-logistics/container-reference-app/internal/pkg/clients/aaa"
)

type BagServiceInterface interface {
	GetBags()
}

type BagService struct {
}

func NewBagService() BagService {

	cfg := &aaa_client.Configuration{}
	apiClient := aaa_client.NewAPIClient(cfg)

	ctx := context.Background()
	ApiAuthClientCredentialsRequest := apiClient.AuthenticationApi.AuthClientCredentials(ctx)

	ClientCredentialsRequest := aaa_client.ClientCredentialsRequest{
		ClientId:     "client_id",
		ClientSecret: "client_secret",
		Audience:     *aaa_client.NewNullableString(aaa_client.PtrString("audience")),
	}

	ApiAuthClientCredentialsRequest.ClientCredentialsRequest(ClientCredentialsRequest)
	apiClient.AuthenticationApi.AuthClientCredentialsExecute(ApiAuthClientCredentialsRequest)

	return BagService{}
}

func (s BagService) GetBags() {

	fmt.Println("invoked GetBags")
}
