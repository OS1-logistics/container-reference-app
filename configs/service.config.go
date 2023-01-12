package config

import (
	"fmt"
	"os"
)

var (
	ServiceConf *ServiceConfig //exported Configuration that all packages can use
)

func LoadConfig() (*ServiceConfig, error) {

	const (
		Env             = "ENV"
		StackBaseUrl    = "STACK_BASE_URL"
		AppClientId     = "APP_CLIENT_ID"
		AppClientSecret = "APP_CLIENT_SECRET"
		AppApiKey       = "APP_API_KEY"
	)

	ServiceConf = &ServiceConfig{}
	if _, t := os.LookupEnv(Env); t {
		ServiceConf.Env = os.Getenv(Env)
	} else {
		panic(fmt.Sprintf("Environment variable %s not set", Env))
	}

	if _, t := os.LookupEnv(StackBaseUrl); t {
		ServiceConf.StackBaseUrl = os.Getenv(StackBaseUrl)
	} else {
		panic(fmt.Sprintf("Environment variable %s not set", AppApiKey))
	}

	if _, t := os.LookupEnv(AppClientId); t {
		ServiceConf.APP.ClientId = os.Getenv(AppClientId)
	} else {
		panic(fmt.Sprintf("Environment variable %s not set", AppClientId))
	}

	if _, t := os.LookupEnv(AppClientSecret); t {
		ServiceConf.APP.ClientSecret = os.Getenv(AppClientSecret)
	} else {
		panic(fmt.Sprintf("Environment variable %s not set", AppClientSecret))
	}

	if _, t := os.LookupEnv(AppApiKey); t {
		ServiceConf.APP.ApiKey = os.Getenv(AppApiKey)
	} else {
		panic(fmt.Sprintf("Environment variable %s not set", AppApiKey))
	}

	// TODO: use viper to mape enviornment variables/ config yaml to ServiceConf
	return ServiceConf, nil
}

type ServiceConfig struct {
	APP          app
	Env          string
	StackBaseUrl string
}

type app struct {
	ClientId     string
	ClientSecret string
	ApiKey       string
}
