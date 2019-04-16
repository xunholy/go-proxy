package config

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/viper"
)

func SetDefaults() {
	viper.SetDefault("Proxy.Address", "localhost")
	viper.SetDefault("Proxy.Port", 3128)
	viper.SetDefault("Proxy.Running", false)
}

func LoadConfiguration(proxyProfilePath string) (Configuration, error) {
	ext := filepath.Ext(proxyProfilePath)
	if ext != "" {
		proxyProfilePath = proxyProfilePath[0 : len(proxyProfilePath)-len(ext)]
	}
	viper.SetConfigName(proxyProfilePath)
	viper.AddConfigPath(os.Getenv("HOME"))
	var configuration Configuration
	if err := viper.ReadInConfig(); err != nil {
		return configuration, fmt.Errorf("failed reading config file, %s", err)
	}
	if err := viper.Unmarshal(&configuration); err != nil {
		return configuration, fmt.Errorf("unable to decode into struct, %s", err)
	}
	return configuration, nil
}

func SaveConfiguration(proxyProfilePath string) error {
	SetDefaults()
	ext := filepath.Ext(proxyProfilePath)
	if ext == "" {
		proxyProfilePath += ".yaml"
	}
	if err := viper.WriteConfigAs(path.Join(os.Getenv("HOME"), proxyProfilePath)); err != nil {
		return fmt.Errorf("unable to write config, %s", err)
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
