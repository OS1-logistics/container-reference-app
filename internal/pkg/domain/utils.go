package domain

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang/glog"
	config "github.com/os1-logistics/container-reference-app/configs"
	cache "github.com/os1-logistics/container-reference-app/internal/pkg/cache"
	"github.com/os1-logistics/container-reference-app/internal/pkg/common"
	aaa "github.com/os1-logistics/container-reference-app/internal/pkg/domain/aaa"
)

func GetToken(tenantId string) (string, error) {

	if tenantId == "" {
		tenantId = "alpha"
	}

	tenantTokenKey := fmt.Sprintf("%s:%s:os1:token", common.AppName, tenantId)
	token, isPresent := cache.ServiceCache.Get(tenantTokenKey)
	if isPresent && token != nil {
		glog.Info("Token found in cache")
		fmt.Println(token)
		return token.(string), nil
	}

	aaaClient := NewAAAClient(tenantId)

	ctx := context.Background()
	ApiAuthClientCredentialsRequest := aaaClient.AuthenticationApi.AuthClientCredentials(ctx)

	ClientCredentialsRequest := aaa.ClientCredentialsRequest{
		ClientId:     config.ServiceConf.APP.ClientId,
		ClientSecret: config.ServiceConf.APP.ClientSecret,
		Audience:     *aaa.NewNullableString(aaa.PtrString("container:reference:app")),
	}

	request := ApiAuthClientCredentialsRequest.ClientCredentialsRequest(ClientCredentialsRequest)
	request = request.XCOREOSREQUESTID("1234")
	request = request.XCOREOSTID("alpha")

	success, response, _ := aaaClient.AuthenticationApi.AuthClientCredentialsExecute(request)
	var accessToken string
	glog.Info("Generating token for tenant: ", tenantId)
	if response.StatusCode == 200 && success != nil {
		accessToken = success.Data.GetAccessToken()
		// cache duration is 10 minutes less than the actual token expiry
		duration := time.Duration(success.Data.GetExpiresIn()-600) * time.Second
		cache.ServiceCache.SetWithExpiry(tenantTokenKey, accessToken, duration)
		return accessToken, nil
	}

	glog.Info("Unable to get token")

	return "", errors.New("Unable to get token")
}
