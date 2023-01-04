package service

import (
	"fmt"

	domain "github.com/os1-logistics/container-reference-app/internal/pkg/domain"
	aaa "github.com/os1-logistics/container-reference-app/internal/pkg/domain/aaa"
	container "github.com/os1-logistics/container-reference-app/internal/pkg/domain/container"
)

type BagServiceInterface interface {
	GetBags()
}

type BagService struct {
	containerApiClient container.APIClient
	aaaApiClient       aaa.APIClient
}

func NewBagService() BagService {

	containerClient := domain.NewContainerClient("alpha")
	aaaClient := domain.NewAAAClient("alpha")

	domain.GetToken("alpha")
	// ctx := context.Background()
	// ApiAuthClientCredentialsRequest := aaaClient.AuthenticationApi.AuthClientCredentials(ctx)

	// ClientCredentialsRequest := aaa.ClientCredentialsRequest{
	// 	ClientId:     "platform:app:ContainerReferenceApp-backend",
	// 	ClientSecret: "****",
	// 	Audience:     *aaa.NewNullableString(aaa.PtrString("container:reference:app")),
	// }

	// request := ApiAuthClientCredentialsRequest.ClientCredentialsRequest(ClientCredentialsRequest)
	// request = request.XCOREOSREQUESTID("1234")
	// request = request.XCOREOSTID("alpha")

	// success, response, _ := aaaClient.AuthenticationApi.AuthClientCredentialsExecute(request)
	// fmt.Println("=====================================")
	// if response.StatusCode == 200 && success != nil {
	// 	fmt.Println(success.Data.GetAccessToken())
	// } else {
	// 	fmt.Println("Unable to get token")
	// }

	return BagService{
		aaaApiClient:       *aaaClient,
		containerApiClient: *containerClient,
	}
}

func (s BagService) GetBags() {
	domain.GetToken("alpha")
	fmt.Println("invoked GetBags")
}
