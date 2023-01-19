package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var (
	ServiceConf    *ServiceConfig //exported Configuration that all packages can use
	ErrEnvVarEmpty = errors.New("getenv: environment variable empty")
)

func LoadConfig() (*ServiceConfig, error) {

	const (
		Env             = "ENV"
		StackBaseUrl    = "STACK_BASE_URL"
		AppName         = "APP_NAME"
		AppId           = "APP_ID"
		AppClientId     = "APP_CLIENT_ID"
		AppClientSecret = "APP_CLIENT_SECRET"
		AppApiKey       = "APP_API_KEY"
		AuthEnabled     = "AUTH_ENABLED"
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

	if _, t := os.LookupEnv(AppName); t {
		ServiceConf.APP.AppName = os.Getenv(AppName)
	} else {
		panic(fmt.Sprintf("Environment variable %s not set", AppName))
	}

	if _, t := os.LookupEnv(AppId); t {
		ServiceConf.APP.AppId = os.Getenv(AppId)
	} else {
		panic(fmt.Sprintf("Environment variable %s not set", AppId))
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

	if _, t := os.LookupEnv(AuthEnabled); t {
		ServiceConf.APP.AuthEnabled, _ = strconv.ParseBool(os.Getenv(AuthEnabled))
	} else {
		panic(fmt.Sprintf("Environment variable %s not set", AuthEnabled))
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
	AppName      string
	AppId        string
	ClientId     string
	ClientSecret string
	ApiKey       string
	AuthEnabled  bool
}
