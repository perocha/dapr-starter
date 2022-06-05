package config

import (
	"time"

	"github.com/spf13/viper"
)

// Provider is the interface for config providers
type Provider interface {
	ConfigFileUsed() string
	Get(key string) interface{}
	GetBool(key string) bool
	GetDuration(key string) time.Duration
	GetFloat64(key string) float64
	GetInt(key string) int
	GetInt64(key string) int64
	GetSizeInBytes(key string) uint
	GetString(key string) string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringMapStringSlice(key string) map[string][]string
	GetStringSlice(key string) []string
	GetTime(key string) time.Time
	InConfig(key string) bool
	IsSet(key string) bool
}

// defaultConfig is the default config provider
var defaultConfig *viper.Viper

// Config returns a default config providers
func Config() Provider {
	return defaultConfig
}

// NewViperProvider returns a new viper config provider
func LoadConfigProvider(appName string) (Provider, error) {
	prov, err := readViperConfig(appName)
	if err != nil {
		return nil, err
	}
	return prov, nil
}

// readViperConfig reads the config file and returns a viper config provider
func readViperConfig(appName string) (*viper.Viper, error) {
	v := viper.New()
	v.SetEnvPrefix(appName)
	v.AutomaticEnv()

	v.SetConfigType("yaml")
	v.SetConfigName("config")
	v.AddConfigPath(".")

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return v, nil
}
