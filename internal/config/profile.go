package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func SetDefaults() {
	viper.SetDefault("Proxy.Address", "localhost")
	viper.SetDefault("Proxy.Port", 3128)
}

func LoadConfiguration(proxyProfilePath string) (Configuration, error) {
	viper.SetConfigName(proxyProfilePath)
	viper.AddConfigPath(os.Getenv("HOME"))
	var configuration Configuration
	if err := viper.ReadInConfig(); err != nil {
		return configuration, fmt.Errorf("Error reading config file, %q", err)
	}
	if err := viper.Unmarshal(&configuration); err != nil {
		return configuration, fmt.Errorf("unable to decode into struct, %q", err)
	}
	return configuration, nil
}

func SaveConfiguration(proxyProfilePath string) error {
	// TODO: Can write to TOML / YAML / JSON
	SetDefaults()
	if err := viper.WriteConfigAs(proxyProfilePath); err != nil {
		return fmt.Errorf("unable to write config, %q", err)
	}
	return nil
}

func PrintConfiguration(proxyProfilePath string) error {
	return nil
}

func ValidateRequiredFields(c *Configuration) error {
	// TODO: Add required fields
	if c.ProxyConfig.Credentials.Username == "" {
		return fmt.Errorf("username is required")
	}
	return nil
}
