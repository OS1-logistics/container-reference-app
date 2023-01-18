package domain

import (
	"fmt"

	"github.com/golang/glog"
	config "github.com/os1-logistics/container-reference-app/configs"
	aaa "github.com/os1-logistics/container-reference-app/internal/pkg/domain/aaa"
	authorizer "github.com/os1-logistics/container-reference-app/internal/pkg/domain/authorizer"
	container "github.com/os1-logistics/container-reference-app/internal/pkg/domain/container"
)

func NewAAAClient(tenantId string) *aaa.APIClient {
	glog.Info("NewAAAClient: config.ServiceConf.StackBaseUrl ", config.ServiceConf.StackBaseUrl)
	cfg := aaa.NewConfiguration()
	cfg.Host = fmt.Sprintf("%s.%s", tenantId, config.ServiceConf.StackBaseUrl)
	cfg.Scheme = "https"
	return aaa.NewAPIClient(cfg)
}

func NewAuthorizerClient(tenantId string) *authorizer.APIClient {
	glog.Info("NewAuthorizerClient: config.ServiceConf.StackBaseUrl ", config.ServiceConf.StackBaseUrl)
	cfg := authorizer.NewConfiguration()
	cfg.Host = fmt.Sprintf("%s.%s", tenantId, config.ServiceConf.StackBaseUrl)
	cfg.Scheme = "https"
	return authorizer.NewAPIClient(cfg)
}

func NewContainerClient(tenantId string) *container.APIClient {
	glog.Info("NewContainerClient: config.ServiceConf.StackBaseUrl: ", config.ServiceConf.StackBaseUrl)
	cfg := container.NewConfiguration()
	cfg.Host = fmt.Sprintf("%s.%s", tenantId, config.ServiceConf.StackBaseUrl)
	cfg.Scheme = "https"
	return container.NewAPIClient(cfg)
}
