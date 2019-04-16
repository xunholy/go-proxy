package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
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
	if cmd.Flags().Changed("all") {
		if err := config.UnsetAllConfiguration(proxyProfile); err != nil {
			log.Fatal(err)
		}
	}
	env.UpdateGlobalEnvironmentVariables("")
	cmds := execute.Command{Cmd: "pkill", Args: []string{"cntlm"}}
	_, err := execute.RunCommand(cmds)
	if err != nil {
		log.Fatalf("Couldn't kill CNTLM process %q", err)
	}
	fmt.Println("CNTLM proxy stopped")
}
