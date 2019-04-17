package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xUnholy/go-proxy/internal/config"
	tools "github.com/xUnholy/go-proxy/internal/tools"
)

func SetupUnsetCli() *cobra.Command {
	// Top-level command
	var unsetCmd = &cobra.Command{
		Use:   "unset",
		Short: "Unset CNTLM Proxy Config",
	}

	var unsetNpmCmd = &cobra.Command{
		Use:   "npm",
		Short: "This command will unset the NPM proxy values. Both https-proxy and proxy will be unset",
		Run:   unsetNpmCmd,
	}

	var unsetGitCmd = &cobra.Command{
		Use:   "git",
		Short: "This command will unset the GIT global proxy values. Both http.proxy and https.proxy will be unset",
		Run:   unsetGitCmd,
	}

	// add subcommands to set
	unsetCmd.AddCommand(unsetNpmCmd, unsetGitCmd)
	return unsetCmd
}

func unsetNpmCmd(cmd *cobra.Command, args []string) {
	_, err := config.LoadConfiguration(proxyProfile)
	if err != nil {
		log.Fatal(err)
	}
	_ = tools.DisableNPMProxyConfiguration()
	viper.Set("Proxy.Tools.Npm", false)
	config.SaveConfiguration(proxyProfile)
	fmt.Println("Unset npm config successfully")
}

func unsetGitCmd(cmd *cobra.Command, args []string) {
	_, err := config.LoadConfiguration(proxyProfile)
	if err != nil {
		log.Fatal(err)
	}
	_ = tools.DisableGITProxyConfiguration()
	viper.Set("Proxy.Tools.Git", false)
	config.SaveConfiguration(proxyProfile)
	fmt.Println("Unset git config successfully")
}
