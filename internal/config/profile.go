package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

func SetupConfigurationFile() {
	var configuration Configuration
	LoadConfigurtation()
	ValidateRequiredFields(&configuration)
}

func SetDefaults() {
	viper.SetDefault("Proxy.Address", "localhost")
	viper.SetDefault("Proxy.Port", 3128)
}

func LoadConfigurtation() {
	// TODO: File name & path should be parameterized
	viper.SetConfigName("example_profile")
	viper.AddConfigPath("./example") // Remove line, only for testing purposes
	viper.AddConfigPath(os.Getenv("HOME"))
	var configuration Configuration
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %q", err)
	}
	if err := viper.Unmarshal(&configuration); err != nil {
		log.Fatalf("unable to decode into struct, %q", err)
	}
	log.Printf("All configuration %v", configuration)
}

func SaveConfiguration() {
	// TODO: Can write to TOML / YAML / JSON
	if err := viper.WriteConfig(); err != nil {
		log.Fatalf("unable to write config, %q", err)
	}
}

func ValidateRequiredFields(c *Configuration) error {
	// TODO: Add required fields
	if c.ProxyConfig.Credentials.Username == "" {
		return fmt.Errorf("username is required")
	}
	return nil
}
