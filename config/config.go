package config

import "github.com/spf13/viper"

type Config struct {
	AppPort      string `mapstructure:"APP_PORT"`
	DaprPort     string `mapstructure:"DAPR_HTTP_PORT"`
	DaprGrpcPort string `mapstructure:"DAPR_GRPC_PORT"`
	Namespace    string `mapstructure:"NAMESPACE"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
