package profile

import (
	"log"
	"os"

	"github.com/spf13/viper"
	"github.com/xUnholy/go-proxy/internal/config"
)

func SetupConfigurationFile() {
	viper.SetConfigName(".proxy_profile")
	viper.AddConfigPath(os.Getenv("HOME"))

	var configuration config.Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&configuration); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	log.Printf("proxy port is %v", configuration.Proxy.Port)
	log.Printf("All configuration %v", configuration)
}
