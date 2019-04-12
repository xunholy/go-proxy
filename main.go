package main

import (
	"github.com/xUnholy/go-proxy/internal/cmd"
	"github.com/xUnholy/go-proxy/internal/profile"
)

func main() {
	profile.SetupConfigurationFile()
	cmd.Execute()
}
