package main

import (
	"github.com/xUnholy/go-proxy/internal/cmd"
	"github.com/xUnholy/go-proxy/internal/config"
)

func main() {
	config.SetupConfigurationFile()
	cmd.Execute()
}
