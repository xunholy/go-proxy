package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/xUnholy/go-proxy/internal/config"
)

var (
	defaultProfile = "profile"
	proxyProfile   = "profile"
)

func SetupConfigCli() *cobra.Command {
	var configCmd = &cobra.Command{
		Use:   "config",
		Short: "Configure custom proxy profile",
	}

	var loadConfigCmd = &cobra.Command{
		Use:   "load",
		Short: "This command will load a custom proxy profile",
		Run:   configLoadCmd,
	}
	loadConfigCmd.Flags().StringVarP(&proxyProfile, "file", "f", proxyProfile, "load custom proxy profile")
	var saveConfigCmd = &cobra.Command{
		Use:   "save",
		Short: "This command will save the proxy configuration to a profile",
		Long: "The save command allows the user to save the current state of the proxy and its configuration, the state will be saved " +
			"by default in the profile.yaml file. Saving to a custom file can be done using the --file flag and providing a custom file name. " +
			"Supported file types are [ YAML | TOML | JSON ], if none is provided it will default to YAML.",
		Run: configSaveCmd,
	}
	saveConfigCmd.Flags().StringVarP(&proxyProfile, "file", "f", proxyProfile, "save custom proxy profile (Default Values)")
	var printConfigCmd = &cobra.Command{
		Use:   "list",
		Short: "This command will print the proxy configuration profile",
		Run:   configPrintCmd,
	}
	configCmd.AddCommand(loadConfigCmd, saveConfigCmd, printConfigCmd)
	return configCmd
}

func configLoadCmd(cmd *cobra.Command, args []string) {
	cfg, err := config.LoadConfiguration(proxyProfile)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(cfg)
}

func configSaveCmd(cmd *cobra.Command, args []string) {
	_, err := config.LoadConfiguration(defaultProfile)
	if err != nil {
		log.Fatalln(err)
	}
	if err := config.SaveConfiguration(proxyProfile); err != nil {
		log.Fatalln(err)
	}
}

func configPrintCmd(cmd *cobra.Command, args []string) {
	_, err := config.LoadConfiguration(defaultProfile)
	if err != nil {
		log.Fatalln(err)
	}
	if err := config.PrintConfiguration(); err != nil {
		log.Fatalln(err)
	}
}
