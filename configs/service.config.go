package config

import (
	"os"
)

var (
	ServiceConf *ServiceConfig //exported Configuration that all packages can use
)

func LoadConfig() (*ServiceConfig, error) {

	ServiceConf = &ServiceConfig{}
	if _, t := os.LookupEnv("ENV"); t {
		ServiceConf.Env = os.Getenv("ENV")
	} else {
		panic("ENV not set")
	}

	if _, t := os.LookupEnv("STACK_BASE_URL"); t {
		ServiceConf.StackBaseUrl = os.Getenv("STACK_BASE_URL")
	} else {
		panic("STACK_BASE_URL not set")
	}

	if _, t := os.LookupEnv("APP_CLIENT_ID"); t {
		ServiceConf.APP.ClientId = os.Getenv("APP_CLIENT_ID")
	} else {
		panic("APP_CLIENT_ID not set")
	}

	if _, t := os.LookupEnv("APP_CLIENT_SECRET"); t {
		ServiceConf.APP.ClientSecret = os.Getenv("APP_CLIENT_SECRET")
	} else {
		panic("APP_CLIENT_SECRET not set")
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
}
