package config

import (
	"fmt"

	viper "github.com/spf13/viper"
)

// Config of the server
type Config struct {
	// port to listen to. Defaults to 8080
	ServerPort int `mapstructure:"server_port"`
	// Url path prefix to listen to. Defaults to /
	PathPrefix string `mapstructure:"path_prefix"`
	// HTTP read timeout in seconds. Defaults to 60
	HTTPReadTimeout int `mapstructure:"http_read_timeout"`
	// HTTP write timeout in seconds. Defaults to 60
	HTTPWriteTimeout int `mapstructure:"http_write_timeout"`
	// HTTP rate limit per minute. Defaults to 60
	HTTPRateLimit int `mapstructure:"http_rate_limit"`
	// Maximum HTTP burst. Defaults to 30
	HTTPBurst int `mapstructure:"http_burst"`
	// If CorsEnabled is true, CORS becomes active
	CorsEnabled bool `mapstructure:"cors_enabled"`
	// If CorsEnabled is true, CorsAllowedOrigins become whitelisted
	CorsAllowedOrigins []string `mapstructure:"cors_allowed_origins"`
}

func (config Config) validate() error {
	return nil
	/*
		return validation.ValidateStruct(&config,
			validation.Field(&config.SOMETHING, validation.Required),
		)
	*/
}

// LoadConfig returns the config of the server
func LoadConfig() (*Config, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("json")
	v.SetEnvPrefix("PALITRUX")
	v.AutomaticEnv()

	// defaults
	v.SetDefault("server_port", 8080)
	v.SetDefault("path_prefix", "/")
	v.SetDefault("http_read_timeout", 60)
	v.SetDefault("http_write_timeout", 60)
	v.SetDefault("http_rate_limit", 60)
	v.SetDefault("http_burst", 30)
	v.SetDefault("cors_enabled", true)
	v.SetDefault("cors_allowed_origins", []string{"*"})
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("Failed to read the configuration file: %s", err)
	}
	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("config.json is malformed %s", err)
	}

	err := config.validate()
	if err != nil {
		return nil, err
	}

	return &config, nil
}
