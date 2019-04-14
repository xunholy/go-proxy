package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/xUnholy/go-proxy/internal/config"
)

var (
	proxyProfile = "profile"
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
		Short: "This command will save proxy configuration to a profile",
		Run:   configSaveCmd,
	}
	saveConfigCmd.Flags().StringVarP(&proxyProfile, "file", "f", proxyProfile, "save custom proxy profile")
	var printConfigCmd = &cobra.Command{
		Use:   "print",
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
	if err := config.SaveConfiguration(proxyProfile); err != nil {
		log.Fatalln(err)
	}
}

func configPrintCmd(cmd *cobra.Command, args []string) {
	if err := config.PrintConfiguration(proxyProfile); err != nil {
		log.Fatalln(err)
	}
}
