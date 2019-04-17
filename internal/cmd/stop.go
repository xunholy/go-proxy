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
		if err := config.UnsetAllConfiguration(&cfg); err != nil {
			log.Fatal(err)
		}
	}
	emptyURL, err := url.Parse("")
	if err != nil {
		log.Fatal(err)
	}
	env.UpdateGlobalEnvironmentVariables(emptyURL)
	cmds := execute.Command{Cmd: "pkill", Args: []string{"cntlm"}}
	_, err = execute.RunCommand(cmds)
	if err != nil {
		log.Fatalf("Couldn't kill CNTLM process %q", err)
	}
	viper.Set("Proxy.Running", false)
	config.SaveConfiguration(proxyProfile)
	fmt.Println("CNTLM proxy stopped")
}
