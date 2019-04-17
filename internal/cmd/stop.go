package cmd

import (
	"fmt"
	"log"
	"net/url"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xUnholy/go-proxy/pkg/execute"

	"github.com/xUnholy/go-proxy/internal/config"
	env "github.com/xUnholy/go-proxy/internal/environment"
	tools "github.com/xUnholy/go-proxy/internal/tools"
)

func SetupStopCli() *cobra.Command {
	// Top-level command
	var stopCmd = &cobra.Command{
		Use:   "stop",
		Short: "Stop CNTLM Proxy",
		Run:   stopCmd,
	}
	stopCmd.Flags().Bool("all", true, "unset all proxy configuration")
	return stopCmd
}

func stopCmd(cmd *cobra.Command, args []string) {
	cfg, err := config.LoadConfiguration(proxyProfile)
	if err != nil {
		log.Fatal(err)
	}
	if cmd.Flags().Changed("all") {
		if err = unsetAllConfiguration(&cfg); err != nil {
			log.Fatal(err)
		}
	}
	emptyURL, err := url.Parse("")
	if err != nil {
		log.Fatal(err)
	}
	err = env.UpdateGlobalEnvironmentVariables(emptyURL)
	if err != nil {
		log.Fatal(err)
	}
	cmds := execute.Command{Cmd: "pkill", Args: []string{"cntlm"}}
	_, err = execute.RunCommand(cmds)
	if err != nil {
		log.Fatalf("Couldn't kill CNTLM process %q", err)
	}
	viper.Set("Proxy.Running", false)
	err = config.SaveConfiguration(proxyProfile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("CNTLM proxy stopped")
}

func unsetAllConfiguration(cfg *config.Configuration) error {
	// TODO: Consider a better way to loop through Tools as the list of supported CLI's may grow
	if !cfg.Proxy.Tools.Git {

		if err := tools.DisableGITProxyConfiguration(); err != nil {
			return err
		}
	}
	if !cfg.Proxy.Tools.Npm {
		if err := tools.DisableNPMProxyConfiguration(); err != nil {
			return err
		}
	}
	return nil
}
