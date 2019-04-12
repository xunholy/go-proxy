package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func SetupConfigurationFile() {
	viper.SetConfigName("example_profile")
	viper.AddConfigPath("./example")
	viper.AddConfigPath(os.Getenv("HOME"))
}

func SetDefaults() {
	viper.SetDefault("Proxy.Address", "localhost")
	viper.SetDefault("Proxy.Port", 3128)
}

func LoadConfigurtation() {
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
	if err := viper.WriteConfig(); err != nil {
		log.Fatalf("unable to write config, %q", err)
	}
}
