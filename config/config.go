package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// Load the config from environment or config file
func Load() {
	viper.AutomaticEnv()
	viper.SetConfigName("application")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")
	err := viper.ReadInConfig()
	if err != nil {
		return
	}

	ValidateAppConfig()
}

type config struct {
	Host                                   string
	Port                                   int
	Database                               DatabaseConfig
	PostgresQueryTimeoutInMillisecond      int
	PostgresSavedQueryTimeoutInMillisecond int
}

var appConfig config

func ValidateAppConfig() {
	appConfig = config{
		Host:                                   viperString("app_host", "localhost"),
		Port:                                   viperInt("app_port", 8000),
		Database:                               getDatabaseConfig(),
		PostgresQueryTimeoutInMillisecond:      viperInt("postgres_query_timeout_in_millisecond", 500),
		PostgresSavedQueryTimeoutInMillisecond: viperInt("postgres_saved_query_timeout_in_millisecond", 500),
	}
}

func Host() string {
	return appConfig.Host
}

func Database() DatabaseConfig {
	return appConfig.Database
}

// Port returns the application port
func Port() int {
	return appConfig.Port
}

// Addr returns the interface address in host:port format
func Addr() string {
	return fmt.Sprintf("127.0.0.1:8000")
	//return fmt.Sprintf("%s:%d", Host(), Port())
}

func viperString(config string, defaultVal ...string) string {
	if len(defaultVal) > 0 {
		viper.SetDefault(config, defaultVal[0])
	}

	return viper.GetString(config)
}

func viperInt(config string, defaultVal ...int) int {
	if len(defaultVal) > 0 {
		viper.SetDefault(config, defaultVal[0])
	}

	return viper.GetInt(config)
}
