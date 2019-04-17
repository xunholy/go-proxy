package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/viper"
)

func SetDefaults() {
	viper.SetDefault("Proxy.Address", "localhost")
	viper.SetDefault("Proxy.Running", false)
	viper.SetDefault("Proxy.Domain", "")
	viper.SetDefault("Proxy.ProxyAddress", "")
	viper.SetDefault("Proxy.NoProxy", "")
	viper.SetDefault("Proxy.Port", 3128)
	viper.SetDefault("Proxy.Credentials.Username", "")
	viper.SetDefault("Proxy.Credentials.Password", "")
	viper.SetDefault("Proxy.Credentials.PassLM", "")
	viper.SetDefault("Proxy.Credentials.PassNT", "")
	viper.SetDefault("Proxy.Credentials.PassNTLMv2", "")
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
	if err := ValidateRequiredFields(&configuration); err != nil {
		return configuration, err
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

func PrintConfiguration() error {
	cfg, err := json.MarshalIndent(viper.AllSettings(), "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(cfg))
	return nil
}

func ValidateRequiredFields(c *Configuration) error {
	if c.Proxy.Credentials.Username == "" {
		return fmt.Errorf("username is a required field")
	}
	if len(c.Proxy.ProxyAddress) == 0 {
		return fmt.Errorf("proxy address is a required field")
	}
	if c.Proxy.Domain == "" {
		return fmt.Errorf("domain is a required field")
	}
	if c.Proxy.Credentials.Password == "" {
		if c.Proxy.Credentials.PassLM == "" && c.Proxy.Credentials.PassNT == "" && c.Proxy.Credentials.PassNTLMv2 == "" {
			return fmt.Errorf("credentials is a required field")
		}
	}
	return nil
}
