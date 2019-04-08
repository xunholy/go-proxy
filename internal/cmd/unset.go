package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	git "github.com/xUnholy/go-proxy/internal/git"
	npm "github.com/xUnholy/go-proxy/internal/npm"
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
	err := npm.DisableProxyConfiguration()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Unset npm config successfully")
}

func unsetGitCmd(cmd *cobra.Command, args []string) {
	err := git.DisableProxyConfiguration()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Unset git config successfully")
}
